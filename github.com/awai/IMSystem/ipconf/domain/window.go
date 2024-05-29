/*
 * @Author: cyy
 * @Description: --
 */
package domain

type stateWindow struct {
	stateQueue []*Stat
	statChan   chan *Stat
	sumStat    *Stat
	idx        int64
}

const(
	WINDOW_SIZE=5
)

func newStateWindow() *stateWindow{
	return &stateWindow{
		stateQueue: make([]*Stat, WINDOW_SIZE),
		statChan: make(chan *Stat),
		sumStat: &Stat{},
	}
}

func(sw *stateWindow) getStat() *Stat{
	res:=sw.sumStat.Clone()
	res.Avg(WINDOW_SIZE) 
	return res
}

/**
* 添加新的Stat时，处理影响的数据
*/
func(sw *stateWindow) appendStat(s *Stat){
	//减去即将被删除的stat
	sw.sumStat.Sub(sw.stateQueue[sw.idx%WINDOW_SIZE])
	//更新最新的stat
	sw.stateQueue[sw.idx%WINDOW_SIZE]=s
	//重新计算窗口和
	sw.sumStat.Add(s)
	sw.idx++
}

