//   File Name:  response.go
//    Description:
//    Author:      Chenghu
//    Date:       2024/6/12 15:38
//    Change Activity:

package view

type HttpResponse struct {
	Data      any    `json:"data"`
	ErrorCode int    `json:"errorCode"`
	Msg       string `json:"msg"`
}

//func (h HttpResponse) String() string {
//	res, _ := json.Marshal(h)
//	return string(res)
//}
