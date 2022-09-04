package frr

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/gotti/meshover/spec"
)

type bgpEntry struct {
	BgpNeighborAddr                                string `json:"bgpNeighborAddr"`
	RemoteAs                                       int64  `json:"remoteAs"`
	LocalAs                                        int64  `json:"localAs"`
	NbrExternalLink                                bool   `json:"nbrExternalLink"`
	Hostname                                       string `json:"hostname"`
	PeerGroup                                      string `json:"peerGroup"`
	BgpVersion                                     int    `json:"bgpVersion"`
	RemoteRouterID                                 string `json:"remoteRouterId"`
	LocalRouterID                                  string `json:"localRouterId"`
	BgpState                                       string `json:"bgpState"`
	BgpTimerUpMsec                                 int    `json:"bgpTimerUpMsec"`
	BgpTimerUpString                               string `json:"bgpTimerUpString"`
	BgpTimerUpEstablishedEpoch                     int    `json:"bgpTimerUpEstablishedEpoch"`
	BgpTimerLastRead                               int    `json:"bgpTimerLastRead"`
	BgpTimerLastWrite                              int    `json:"bgpTimerLastWrite"`
	BgpInUpdateElapsedTimeMsecs                    int    `json:"bgpInUpdateElapsedTimeMsecs"`
	BgpTimerHoldTimeMsecs                          int    `json:"bgpTimerHoldTimeMsecs"`
	BgpTimerKeepAliveIntervalMsecs                 int    `json:"bgpTimerKeepAliveIntervalMsecs"`
	ExtendedOptionalParametersLength               bool   `json:"extendedOptionalParametersLength"`
	BgpTimerConfiguredConditionalAdvertisementsSec int    `json:"bgpTimerConfiguredConditionalAdvertisementsSec"`
}

func deleteInvalid(original []byte, invalidCodes []byte) []byte {
	tmp := make([]byte, len(original))
	copy(tmp, original)
	for _, d := range invalidCodes {
		for i := 0; i < len(tmp); i++ {
			if tmp[i] == d {
				tmp = append(tmp[:i], tmp[i+1:]...)
				i--
			}
		}
	}
	return tmp
}

// GetBGPStatus executes bgp neighbor command inside frr container and returns []*spec.PeerBGPStatus
func (f *FrrInstance) GetBGPStatus(ctx context.Context) ([]*spec.PeerBGPStatus, error) {
	j, err := f.execCommand(ctx, []string{"vtysh", "-c", "show bgp neighbor json"})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch bgp neighbor json on frr, err=%w", err)
	}
	var bgpStat map[string]bgpEntry
	if err := json.Unmarshal(j, &bgpStat); err != nil {
		return nil, fmt.Errorf("failed to unmarshal, output=%s, err=%w", string(j), err)
	}

	var ret []*spec.PeerBGPStatus
	for _, v := range bgpStat {
		var d *spec.PeerBGPStatus
		switch v.BgpState {
		case "Established":
			d = &spec.PeerBGPStatus{
				RemoteHostname:  v.Hostname,
				LocalAS:         &spec.ASN{Number: uint32(v.LocalAs)},
				RemoteAS:        &spec.ASN{Number: uint32(v.RemoteAs)},
				BgpNeighborAddr: spec.NewAddress(net.ParseIP(v.BgpNeighborAddr)),
				BGPState:        spec.BGPStates_BGPStateEstablished,
			}
		case "Active":
			d = &spec.PeerBGPStatus{
				BGPState: spec.BGPStates_BGPStateActive,
				LocalAS:  &spec.ASN{Number: uint32(v.LocalAs)},
			}
		case "Connect":
			d = &spec.PeerBGPStatus{
				BGPState: spec.BGPStates_BGPStateConnect,
				LocalAS:  &spec.ASN{Number: uint32(v.LocalAs)},
			}
		case "Idle":
			d = &spec.PeerBGPStatus{
				BGPState: spec.BGPStates_BGPStateActive,
				LocalAS:  &spec.ASN{Number: uint32(v.LocalAs)},
			}
		case "Unknown":
		default:
			log.Println("unknown bgp state", v)
		}
		ret = append(ret, d)
	}

	return ret, nil
}
