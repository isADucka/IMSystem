/*
 * @Author: cyy
 * @Description: 工具类,对返回值的封装
 */
package ipconf

import "github.com/awai/IMSystem/ipconf/domain"




// 封装返回值
func Result(ed []*domain.Endport) Response{
	return Response{
		Message: "ok",
		Code: 0,
		Data: ed,
	}
}

// 返回endpoints节点活跃度前5的
func Top5EndPoints(endpoints []*domain.Endport) []*domain.Endport{
	if len(endpoints)<5 {
		return endpoints
	}
	return endpoints[:5]
}


