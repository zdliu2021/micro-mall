package request

type Page struct {
	PageNum  int    `json:"page"`                                                       // 当前页码
	PageSize int    `json:"limit"`                                                      //每页记录数
	Keyword  string `json:"key"`                                                        //检索关键字
	Order    string `json:"micro-mall-micro-mall-order-proto" binding:"oneof=desc asc"` //排序方式
	Sidx     string `json:"sidx"`                                                       // 排序字段
}
