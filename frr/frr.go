package frr

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/gotti/meshover/spec"
	"github.com/gotti/meshover/status"
)

const frrImage = "frrouting/frr:v8.3.1"

// Backend defines backend of frr, such as docker, nerdctl...
type Backend interface {
	//Run creates frr instance
	Run(ctx context.Context) error
	//Kill kills frr instance
	Kill() error
	//execCommand executes a command inside frr
	execCommand(ctx context.Context, command []string) ([]byte, error)
	//UpdatePeers updates peer information
	UpdatePeers(ctx context.Context, peer []status.FrrPeerDiffrence) error
}

// BackendType defines type of backend, such as docker, containerd
type BackendType string

var (
	//BackendNone do nothing
	BackendNone = BackendType("none")
	//BackendDockerSDK run frr container with docker sdk
	BackendDockerSDK = BackendType("dockersdk")
	//BackendNerdCtl run frr container with containerd(nerdctl)
	BackendNerdCtl = BackendType("nerdctl")
)

// InstanceConfig is frr instance configuration
type InstanceConfig struct {
	hostname     string
	overlayIP    string
	bgpdTemplate string
	daemonConfig string
	vtyshConfig  string
}

// NewFrrConfig creates InstanceConfig
func NewFrrConfig(hostName string, overlayIP string, bgpdTemplate string, daemonConfig string, vtyshConfig string) InstanceConfig {
	return InstanceConfig{hostname: hostName, overlayIP: overlayIP, bgpdTemplate: bgpdTemplate, daemonConfig: daemonConfig, vtyshConfig: vtyshConfig}
}

// DockerInstance implements Backend interface
type DockerInstance struct {
	ctx         context.Context
	frrConfig   InstanceConfig
	asn         *spec.ASN
	configDir   string
	containerID string
	client      *client.Client
}

type frrPeerConfig struct {
	HostName string
	IPAddr   string
	ASN      string
	Peers    []frrConfigPeer
}

type frrConfigPeer struct {
	Add    bool
	IPAddr string
	IFName string
	ASN    string
}

// NewDockerInstance creates docker instance
func NewDockerInstance(ctx context.Context, asn *spec.ASN, config InstanceConfig) (*DockerInstance, error) {
	d, err := ioutil.TempDir(os.TempDir(), "meshover-frr")
	if err != nil {
		return nil, fmt.Errorf("failed to create tempdir, err=%w", err)
	}

	return &DockerInstance{ctx: ctx, asn: asn, frrConfig: config, configDir: d}, nil
}

// Kill kill container
func (f *DockerInstance) Kill() error {
	err := f.client.ContainerKill(f.ctx, f.containerID, "SIGKILL")
	if err != nil {
		return fmt.Errorf("failed to clean frr container, err=%w", err)
	}
	return nil
}

func (f *DockerInstance) newFrrPeerConfig(p []status.FrrPeerDiffrence) *frrPeerConfig {
	fc := new(frrPeerConfig)
	fc.ASN = f.asn.Format()
	fc.HostName = f.frrConfig.hostname
	fc.IPAddr = f.frrConfig.overlayIP
	//TODO CRUDがクッソ雑
	for _, pp := range p {
		fc.Peers = append(fc.Peers, frrConfigPeer{Add: true, IPAddr: pp.NewPeer.GetAddress()[0].ToNetIPNet().IP.String(), ASN: pp.NewPeer.GetAsnumber().Format(), IFName: pp.TunName})
	}
	return fc
}

func (f *DockerInstance) execCommand(ctx context.Context, command []string) ([]byte, error) {
	res, err := f.client.ContainerExecCreate(ctx, f.containerID, types.ExecConfig{
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          command,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to run ContainerExecCreate, err=%w", err)
	}
	respo, err := f.client.ContainerExecAttach(ctx, res.ID, types.ExecStartCheck{
		Tty:    true,
		Detach: false})
	if err != nil {
		return nil, fmt.Errorf("failed to run ContainerExecAttach, err=%w", err)
	}
	defer respo.Close()
	d, err := ioutil.ReadAll(respo.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read all, err=%w", err)
	}
	return d, nil
}

// UpdatePeers updates frr bgp peer
func (f *DockerInstance) UpdatePeers(ctx context.Context, peer []status.FrrPeerDiffrence) error {
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
	if d, err := f.execCommand(ctx, []string{"vtysh", "-b"}); err != nil {
		return fmt.Errorf("failed to execute command inside the docker container, result=%s, err=%w", d, err)
	}
	vv, err := os.ReadFile(fb.Name())
	if err != nil {
		return err
	}
	log.Println(string(vv))
	return nil
}

// Run executes docker instance
func (f *DockerInstance) Run(ctx context.Context) error {
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
	defer reader.Close()

	if err := f.client.ContainerKill(ctx, "meshover0-frr", "SIGTERM"); err != nil {
		log.Println("failed to kill container", err)
	}

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
		Image: frrImage,
	}, &container.HostConfig{
		CapAdd:      []string{"CAP_NET_ADMIN", "CAP_NET_RAW", "CAP_SYS_ADMIN"},
		NetworkMode: "host",
		Privileged:  true,
		AutoRemove:  true,
		Binds:       []string{fmt.Sprintf("%s:/etc/frr", f.configDir)},
	}, nil, nil, "meshover0-frr")
	if err != nil {
		log.Println("failed to create container", err)
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

// NerdCtlInstance implements Backend interface
type NerdCtlInstance struct {
	ctx         context.Context
	frrConfig   InstanceConfig
	asn         *spec.ASN
	configDir   string
	containerID string
}

// NewNerdCtlInstance creates NerdCtlInstance
func NewNerdCtlInstance(ctx context.Context, asn *spec.ASN, config InstanceConfig) (*NerdCtlInstance, error) {
	d, err := ioutil.TempDir(os.TempDir(), "meshover-frr")
	if err != nil {
		return nil, fmt.Errorf("failed to create tempdir, err=%w", err)
	}

	return &NerdCtlInstance{ctx: ctx, asn: asn, frrConfig: config, configDir: d}, nil
}

// Run starts with nerdctl
func (f *NerdCtlInstance) Run(ctx context.Context) error {
	if err := f.Kill(); err != nil {
		log.Println("failed to kill before running frr instance, err", err)
	}
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
	out, err := exec.Command("bash", "-c", fmt.Sprintf("nerdctl -n meshover run -d --net=host --cap-add=CAP_NET_ADMIN --cap-add=CAP_NET_RAW --cap-add=CAP_SYS_ADMIN -v %s:/etc/frr --privileged --name meshover0-frr frrouting/frr:v8.3.1", f.configDir)).CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run frr container, output=%s err=%w", out, err)
	}
	f.containerID = string(out)
	return nil
}

// Kill kills by nerdctl
func (f *NerdCtlInstance) Kill() error {
	out, err := exec.Command("bash", "-c", "nerdctl -n meshover kill meshover0-frr").CombinedOutput()
	if err != nil {
		log.Printf("failed to kill container, err=%e\n", err)
	}
	if string(out) != f.containerID {
		log.Printf("may failed to kill container, expected=%s, got=%s", f.containerID, string(out))
	}
	out2, err := exec.Command("bash", "-c", "nerdctl -n meshover rm meshover0-frr").CombinedOutput()
	if err != nil {
		log.Printf("may failed to remove container, output=%s, err=%e", out2, err)
	}
	if string(out2) != f.containerID {
		log.Printf("may failed to remove container, expected=%s, got=%s", f.containerID, string(out))
	}
	return nil
}

func (f *NerdCtlInstance) execCommand(ctx context.Context, command []string) ([]byte, error) {
	for i, c := range command {
		command[i] = "\"" + c + "\""
	}
	flatten := ""
	for _, c := range command {
		flatten += " " + c
	}
	command = append([]string{"-c", "nerdctl -n meshover exec meshover0-frr" + flatten}, command...)
	out, err := exec.Command("bash", command...).CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to execute command, output=%s, err=%w", string(out), err)
	}
	return out, nil
}

// UpdatePeers updates peers
func (f *NerdCtlInstance) UpdatePeers(ctx context.Context, peer []status.FrrPeerDiffrence) error {
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
	if d, err := f.execCommand(ctx, []string{"vtysh", "-b"}); err != nil {
		return fmt.Errorf("failed to execute command inside the containerd, result=%s, err=%w", d, err)
	}
	vv, err := os.ReadFile(fb.Name())
	if err != nil {
		return err
	}
	log.Println(string(vv))
	return nil
}

func (f *NerdCtlInstance) newFrrPeerConfig(p []status.FrrPeerDiffrence) *frrPeerConfig {
	fc := new(frrPeerConfig)
	fc.ASN = f.asn.Format()
	fc.HostName = f.frrConfig.hostname
	fc.IPAddr = f.frrConfig.overlayIP
	//TODO CRUDがクッソ雑
	for _, pp := range p {
		fc.Peers = append(fc.Peers, frrConfigPeer{Add: true, IPAddr: pp.NewPeer.GetAddress()[0].ToNetIPNet().IP.String(), ASN: pp.NewPeer.GetAsnumber().Format(), IFName: pp.TunName})
	}
	return fc
}

// DummyInstance implements Backend interface and do nothing
type DummyInstance struct {
}

// NewDummyInstance creates DummyInstance
func NewDummyInstance() *DummyInstance {
	return &DummyInstance{}
}

// Run do nothing and always return nil as error
func (f *DummyInstance) Run(ctx context.Context) error { return nil }

// Kill do nothing and always return nil as error
func (f *DummyInstance) Kill() error { return nil }

// execCommand do nothing and always return nil as bytes and error
func (f *DummyInstance) execCommand(ctx context.Context, command []string) ([]byte, error) {
	return []byte("{}"), nil
}

// UpdatePeers updates peer information and always return nil as error
func (f *DummyInstance) UpdatePeers(ctx context.Context, peer []status.FrrPeerDiffrence) error {
	return nil
}
