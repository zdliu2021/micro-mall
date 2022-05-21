package request

type MergePurchaseRequest struct {
	PurchaseId int64   `json:"purchaseId"`
	Items      []int64 `json:"items"`
}

type ReceivePurchaseRequest struct {
	Ids []int64 `json:"ids"`
}

type SaveWareRequest struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Areacode string `json:"areacode"`
}

type SavePurchaseDetailRequest struct {
	PurchaseId int32   `json:"purchaseId"`
	SkuId      int32   `json:"skuId"`
	SkuNum     int32   `json:"skuNum"`
	SkuPrice   float64 `json:"skuPrice"`
	WareId     int     `json:"wareId"`
	Status     int     `json:"status"`
}

type SavePurchaseRequest struct {
	AssigneeId   int32   `json:"assigneeId"`
	AssigneeName string  `json:"assigneeName"`
	Phone        string  `json:"phone"`
	Priority     int32   `json:"priority"`
	Status       int     `json:"status"`
	WareId       int32   `json:"wareId"`
	Amount       float64 `json:"amount"`
	CreateTime   string  `json:"createTime"`
	UpdateTime   string  `json:"updateTime"`
}
type UpdatePurchaseRequest struct {
	Id           int    `json:"id"`
	AssigneeId   int    `json:"assigneeId"`
	AssigneeName string `json:"assigneeName"`
	Phone        string `json:"phone"`
	Status       int    `json:"status"`
}
