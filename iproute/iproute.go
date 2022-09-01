package iproute

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"strconv"
	"sync"

	"github.com/gotti/meshover/status"
	"go.uber.org/zap"
)

//tableNumStart sets initial usable table number for rt_tables, and table number is incremented by each host
const tableNumStart = 532
const tableNumEnd = 599

//Instance contains ideal status of routing table
type Instance struct{
	currenttableNum int
	mtx sync.Mutex
	sbrs []*SBRStatus
	log *zap.Logger
}

//SBRDiff contains sufficient info about SBR
type SBRDiff struct {
	HostName string
	Diff status.DiffType
	TunName string
	Gathering []net.IPNet
}

//NewInstance creates a Instance
func NewInstance(log *zap.Logger) Instance{
	return Instance{currenttableNum: tableNumStart, mtx: sync.Mutex{}, sbrs: []*SBRStatus{}, log: log}
}

//UpdateSBRPeer updates SBR peers
func (i *Instance)UpdateSBRPeer(diff []SBRDiff) error{
	for _, d := range diff{
		switch d.Diff{
		case status.DiffTypeAdd:{
			for _, e := range d.Gathering{
				i.addSBRPeer(d.TunName, e)
			}
		}
		case status.DiffTypeDelete:{
			i.delSBRPeer(d.TunName)
		}
		case status.DiffTypeChange:{
			i.delSBRPeer(d.TunName)
			for _, e := range d.Gathering{
				i.addSBRPeer(d.TunName, e)
			}
		}
		}
	}
	return nil
}

//addSBRPeer adds Source Based Routing Peer
func (i *Instance)addSBRPeer(tunName string, gathering net.IPNet){
	i.mtx.Lock()
	defer i.mtx.Unlock()
	s := &SBRStatus{tableNum: i.currenttableNum, oppositeTunName: tunName, gatheringRange: gathering}
	i.currenttableNum = i.currenttableNum + 1
	i.sbrs = append(i.sbrs, s)
	if err := addRouteAndTable(s); err != nil {
		i.log.Error("failed to add route and table", zap.Error(err))
	}
}

//updateSBRPeer adds Source Based Routing Peer
func (i *Instance)updateSBRPeer(tunName string, gathering net.IPNet){
	i.mtx.Lock()
	defer i.mtx.Unlock()
	s := &SBRStatus{tableNum: i.currenttableNum, oppositeTunName: tunName, gatheringRange: gathering}
	i.currenttableNum = i.currenttableNum + 1
	i.sbrs = append(i.sbrs, s)
	if err := addRouteAndTable(s); err != nil {
		i.log.Error("failed to add route and table", zap.Error(err))
	}
}

//delSBRPeer deletes all Source Based Routing Peer by tunName
func (i *Instance)delSBRPeer(tunName string){
	i.mtx.Lock()
	defer i.mtx.Unlock()
	for _, d := range i.sbrs{
		if d.oppositeTunName != tunName{
			continue
		}
		if err := delRuleIfExist(d.tableNum); err != nil {
			i.log.Error("failed to delete rule", zap.Error(err))
		}
		if err := delRouteIfExist(d.tableNum); err != nil {
			i.log.Error("failed to delete route", zap.Error(err))
		}
	}
}

//Clean deletes all created rule and routes
func (i *Instance)Clean() error{
	for _, s := range i.sbrs{
		if err := delRuleIfExist(s.tableNum); err != nil {
			return fmt.Errorf("failed to delete rule of tablenum=%d, err=%w", s.tableNum, err)
		}
		if err := delRouteIfExist(s.tableNum); err != nil {
			return fmt.Errorf("failed to delete route of tablenum=%d, err=%w", s.tableNum, err)
		}
	}
	return nil
}

//ForceClean deletes all potentially exist rule and routes
func ForceClean() error{
	for i := tableNumStart; i < tableNumEnd; i++{
		if err := delRuleIfExist(i); err != nil {
			log.Printf("failed to delete rule of tablenum=%d, err=%s\n", i, err)
		}
		if err := delRouteIfExist(i); err != nil {
			log.Printf("failed to delete route of tablenum=%d, err=%s\n", i, err)
		}
	}
	return nil
}

//SBRStatus contains single SBR status
type SBRStatus struct{
	tableNum int
	oppositeTunName string
	gatheringRange net.IPNet
}

func addRouteAndTable(s *SBRStatus) error{
	strtable := strconv.Itoa(s.tableNum)
	if err := delRuleIfExist(s.tableNum); err != nil {
		return fmt.Errorf("failed to delete rule of tablenum=%s, err=%w", strtable, err)
	}
	if err := delRouteIfExist(s.tableNum); err != nil {
		return fmt.Errorf("failed to delete route of tablenum=%s, err=%w", strtable, err)
	}
	if err := exec.Command("bash", "-c", "ip rule add from " + s.gatheringRange.String() + " table " + strtable).Run(); err != nil {
		return fmt.Errorf("failed to add rule, err=%w", err)
	}
	if err := exec.Command("bash", "-c", "ip route add default dev " + s.oppositeTunName + " table " + strtable).Run(); err != nil {
		return fmt.Errorf("failed to add default route, err=%w", err)
	}
	return nil
}

func delRuleIfExist(tableNumInt int) error{
	tableNum := strconv.Itoa(tableNumInt)
	ruo, err := exec.Command("bash", "-c", "ip rule ls table " + tableNum).CombinedOutput()
	if err != nil || len(ruo) != 0{
		o, err := exec.Command("bash", "-c", "ip rule del table " + tableNum).CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed to del rules of specified table number, output=%s, err=%w", string(o), err)
		}
	}
	return nil
}

func delRouteIfExist(tableNumInt int) error{
	tableNum := strconv.Itoa(tableNumInt)
	roo, err := exec.Command("bash",  "-c", "ip route ls table " + tableNum).Output()
	if err == nil && len(roo) != 0 {
		o, err := exec.Command("bash", "-c", "ip route del default table " + tableNum).CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed to del routes of specified table number, output=%s, err=%w", string(o), err)
		}
	}
	return nil
}
