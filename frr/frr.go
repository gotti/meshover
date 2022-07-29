package frr

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/gotti/meshover/spec"
)

type FrrConfig struct {
	hostname     string
	overlayIP    string
	bgpdTemplate string
	daemonConfig string
	vtyshConfig  string
}

type FrrInstance struct {
	frrConfig   FrrConfig
	asn         *spec.ASN
	configDir   string
	containerID string
	client      *client.Client
}

type frrPeerConfig struct {
	HostName string
	IPAddr   string
	ASN      string
	Peers    []FrrConfigPeer
}

type FrrConfigPeer struct {
	Add    bool
	IPAddr string
	ASN    string
}

func NewInstance(asn *spec.ASN, hostName string, overlayIP string, bgpdTemplate string, daemons string, vtysh string) (*FrrInstance, error) {
	d, err := ioutil.TempDir(os.TempDir(), "meshover-frr")
	if err != nil {
		return nil, fmt.Errorf("failed to create tempdir, err=%w", err)
	}

	return &FrrInstance{asn: asn, frrConfig: FrrConfig{hostname: hostName, overlayIP: overlayIP, bgpdTemplate: bgpdTemplate, daemonConfig: daemons, vtyshConfig: vtysh}, configDir: d}, nil
}

func (f *FrrInstance) newFrrPeerConfig(p *spec.Peers) *frrPeerConfig {
	fc := new(frrPeerConfig)
	fc.ASN = fmt.Sprintf("%d", f.asn.GetNumber())
	fc.HostName = f.frrConfig.hostname
	fc.IPAddr = f.frrConfig.overlayIP
	for _, pp := range p.GetPeers() {
		fc.Peers = append(fc.Peers, FrrConfigPeer{Add: true, IPAddr: pp.GetAddress().GetIpaddress(), ASN: fmt.Sprintf("%d", pp.GetAsnumber().GetNumber())})
	}
	return fc
}

func (f *FrrInstance) execCommand(ctx context.Context, command []string) error {
	res, err := f.client.ContainerExecCreate(ctx, f.containerID, types.ExecConfig{
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          command,
	})
	if err != nil {
		return fmt.Errorf("failed to run ContainerExecCreate, err=%w", err)
	}
	respo, err := f.client.ContainerExecAttach(ctx, res.ID, types.ExecStartCheck{})
	if err != nil {
		return fmt.Errorf("failed to run ContainerExecAttach, err=%w", err)
	}
	respo.Close()
	return nil
}

func (f *FrrInstance) UpdatePeers(ctx context.Context, peer *spec.Peers) error {
	tmpl, err := template.New("frr").Parse(f.frrConfig.bgpdTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template, err=%w", err)
	}
	fb, err := os.Create(filepath.Join(f.configDir, "frr.conf"))
	if err != nil {
		return fmt.Errorf("failed to create bgpd.conf, err=%w", err)
	}
	cfg := f.newFrrPeerConfig(peer)
	if err := tmpl.Execute(fb, cfg); err != nil {
		return fmt.Errorf("failed to execute template, err=%w", err)
	}
	if err := f.execCommand(ctx, []string{"vtysh", "-b"}); err != nil {
		return fmt.Errorf("failed to execute command inside the docker container, err=%w", err)
	}
	vv, err := os.ReadFile(fb.Name())
	if err != nil {
		return err
	}
	log.Println(string(vv))
	return nil
}

func (f *FrrInstance) Up(ctx context.Context) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("failed to create docker client, err=%w", err)
	}
	f.client = cli

	reader, err := f.client.ImagePull(ctx, "frrouting/frr:v8.3.0", types.ImagePullOptions{})
	if err != nil {
		return fmt.Errorf("failed to image pull, err=%w", err)
	}
	io.Copy(os.Stdout, reader)

	if err := f.client.ContainerKill(ctx, "meshover0-frr", "SIGTERM"); err != nil {
		log.Println("failed to kill container", err)
	}

	time.Sleep(5 * time.Second)

	fd, err := os.Create(filepath.Join(f.configDir, "daemons"))
	if err != nil {
		return fmt.Errorf("failed to create daemons config file")
	}
	fd.Write([]byte(f.frrConfig.daemonConfig))

	fv, err := os.Create(filepath.Join(f.configDir, "vtysh.conf"))
	if err != nil {
		return fmt.Errorf("failed to create daemons config file")
	}
	fv.Write([]byte(f.frrConfig.vtyshConfig))

	resp, err := f.client.ContainerCreate(ctx, &container.Config{
		Image: "frrouting/frr:v8.3.0",
	}, &container.HostConfig{
		CapAdd:      []string{"CAP_NET_ADMIN", "CAP_NET_RAW", "CAP_SYS_ADMIN"},
		NetworkMode: "host",
		Privileged:  true,
		AutoRemove:  true,
		Binds:       []string{fmt.Sprintf("%s:/etc/frr", f.configDir)},
	}, nil, nil, "meshover0-frr")
	if err != nil {
		return fmt.Errorf("failed to create container, err=%w", err)
	}
	f.containerID = resp.ID

	if err := f.client.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return fmt.Errorf("failed to start container, err=%w", err)
	}
	defer func() {
		err := f.client.ContainerKill(ctx, resp.ID, "SIGTERM")
		if err != nil {
			fmt.Println("failed to kill container", err)
		}
	}()

	out, err := f.client.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		return fmt.Errorf("failed to get container logs, err=%w", err)
	}
	defer out.Close()
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	io.Copy(os.Stdout, out)
	return nil
}
