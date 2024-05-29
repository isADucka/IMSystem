/*
 * @Author: cyy
 * @Description: --
 */
package domain

type Endport struct {
	Ip          string       `json:"ip"`
	Port        string       `json:"port"`
	ActiveSorce float64      `json:"-"`
	StaticSorce float64      `json:"-"`
	Stat        *Stat        `json:"-"`
	window      *stateWindow `json:"-"`
}


//新建一个endPort节点
func NewEndport(ip, port string) *Endport {
	ed := &Endport{
		Ip:   ip,
		Port: port,
	}
	ed.window = newStateWindow()
	ed.Stat = ed.window.getStat()
	go func() {
		for stat := range ed.window.statChan {
			ed.window.appendStat(stat)

		}
	}()
	return ed
}

func (ed *Endport) UpdateStat(s *Stat) {
	ed.window.statChan <- s
}

//计算endPoint的分数
func(ed *Endport) CalculateScore(){
	if(ed.Stat!=nil){
		ed.ActiveSorce=ed.Stat.CalculateACtiveSorce()
		ed.StaticSorce=ed.Stat.ConnectNum
	}
}