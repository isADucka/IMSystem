/*
 * @Author: cyy
 * @Description: --
 */
package domain

import (
	"sort"
	"sync"

	"github.com/awai/IMSystem/ipconf/source"
)

var dp *Dispatcher

type Dispatcher struct {
	candidateTable map[string]*Endport
	sync.RWMutex
}

func Init() {
	dp := &Dispatcher{}
	dp.candidateTable = make(map[string]*Endport)
	go func() {
		for event := range source.EventChan() {
			switch event.Type {
			case source.ADD_NODE_ENVENT:
				dp.addNode(event)
			case source.DEL_NODE_ENVENT:
				dp.delNode(event)
			}
		}
	}()
}

// 添加候选节点
func (dp *Dispatcher) addNode(event *source.Event) {
	dp.Lock()
	defer dp.Unlock()
	ed := NewEndport(event.IP, event.Port)
	ed.UpdateStat(&Stat{
		ConnectNum:   event.ConnectNum,
		MessageBytes: event.MessageBytes,
	})
	dp.candidateTable[event.Key()] = ed

}

// 删除候选节点
func (dp *Dispatcher) delNode(event *source.Event) {
	dp.Lock()
	defer dp.Unlock()
	delete(dp.candidateTable, event.Key())
}

// 进行ip调度
func Dispatch(icf *IpConfContext) []*Endport {
	edpoints := dp.getCandidateEndport()
	//进行计算得分
	for _, ed := range edpoints {
		ed.CalculateScore()
	}

	sort.Slice(edpoints, func(i, j int) bool {
		//优先activeSorce
		if edpoints[i].ActiveSorce > edpoints[j].ActiveSorce {
			return true
		}
		//activeSorce一样的情况下，用staticSorce排序
		if (edpoints[i].ActiveSorce == edpoints[j].ActiveSorce) && (edpoints[i].StaticSorce > edpoints[j].StaticSorce) {
			return true
		}
		return false
	})
	return edpoints

}

// 获取所有的候选节点
func (dp *Dispatcher) getCandidateEndport() []*Endport {
	dp.RLock()
	defer dp.RUnlock()
	candidateList := make([]*Endport, 0, len(dp.candidateTable))
	for _, ed := range dp.candidateTable {
		candidateList = append(candidateList, ed)
	}
	return candidateList
}
