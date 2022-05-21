package response

type GetWareInfoListResponse struct {
	TotalCount int `json:"totalCount"`
	PageSize   int `json:"pageSize"`
	TotalPage  int `json:"totalPage"`
	CurrPage   int `json:"currPage"`
	List       []struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Address  string `json:"address"`
		Areacode string `json:"areacode"`
	} `json:"list"`
}

type GetSkuWareInfoResponse struct {
	TotalCount int `json:"totalCount"`
	PageSize   int `json:"pageSize"`
	TotalPage  int `json:"totalPage"`
	CurrPage   int `json:"currPage"`
	List       []struct {
		Id          int    `json:"id"`
		SkuId       int    `json:"skuId"`
		WareId      int    `json:"wareId"`
		Stock       int    `json:"stock"`
		SkuName     string `json:"skuName"`
		StockLocked int    `json:"stockLocked"`
	} `json:"list"`
}

type GetPurchaseDetailedInfo struct {
	TotalCount int `json:"totalCount"`
	PageSize   int `json:"pageSize"`
	TotalPage  int `json:"totalPage"`
	CurrPage   int `json:"currPage"`
	List       []struct {
		Id         int     `json:"id"`
		PurchaseId int     `json:"purchaseId"`
		SkuId      int     `json:"skuId"`
		SkuNum     int     `json:"skuNum"`
		SkuPrice   float64 `json:"skuPrice"`
		WareId     int     `json:"wareId"`
		Status     int     `json:"status"`
	} `json:"list"`
}

type GetUnReceivedPurchaseInfo struct {
	TotalCount int `json:"totalCount"`
	PageSize   int `json:"pageSize"`
	TotalPage  int `json:"totalPage"`
	CurrPage   int `json:"currPage"`
	List       []struct {
		Id           int     `json:"id"`
		AssigneeId   int     `json:"assigneeId"`
		AssigneeName string  `json:"assigneeName"`
		Phone        string  `json:"phone"`
		Priority     int     `json:"priority"`
		Status       int     `json:"status"`
		WareId       int     `json:"wareId"`
		Amount       float64 `json:"amount"`
		CreateTime   string  `json:"createTime"`
		UpdateTime   string  `json:"updateTime"`
	} `json:"list"`
}

type GetPurchaseListResponse struct {
	List []struct {
		Id           int64   `json:"id"`
		AssigneeId   int64   `json:"assigneeId"`
		AssigneeName string  `json:"assigneeName"`
		Phone        string  `json:"phone"`
		Priority     int     `json:"priority"`
		Status       int     `json:"status"`
		WareId       int64   `json:"wareId"`
		Amount       float64 `json:"amount"`
		CreateTime   string  `json:"createTime"`
		UpdateTime   string  `json:"updateTime"`
	} `json:"list"`

	TotalCount int32 `json:"totalCount"`
}
