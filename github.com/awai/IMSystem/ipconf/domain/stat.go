/*
 * @Author: cyy
 * @Description: --
 */
package domain

import "math"

type Stat struct {
	ConnectNum   float64  //这个代表的是网关gateway总体持有的长连接数量的剩余值
	MessageBytes float64  //网关gateway每秒收发消息的总字节数的剩余值
}


func(s *Stat) CalculateACtiveSorce() float64{
	return getGB(s.MessageBytes)
}

/**
* 转换为GB 并且四舍五入报留两个小数
*/
func getGB(m float64) float64{
	value:=(m/(1<<30))
	return math.Trunc(value*1e2+0.5)*1e-2
}

/**
* Clone了一个Stat
*/
func (s *Stat) Clone() *Stat{
	return &Stat{
		ConnectNum: s.ConnectNum,
		MessageBytes: s.MessageBytes,
	}
}

func(s *Stat) Avg(num float64){
	s.ConnectNum/=num
	s.MessageBytes/=num
}


/**
* 加上新的连接
*/
func(s *Stat) Add(st *Stat){
	if st!=nil{
		s.ConnectNum+=st.ConnectNum
		s.MessageBytes+=st.MessageBytes
	}
}

/*
* 减去即将端开的连接
*/
func(s *Stat) Sub(st *Stat){
	if st!=nil{
		s.ConnectNum-=st.ConnectNum
		s.MessageBytes-=st.MessageBytes
	}
}