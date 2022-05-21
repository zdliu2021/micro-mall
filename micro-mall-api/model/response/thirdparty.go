package response

type GetOSSTokenResponse struct {
	AccessId  string `json:"accessid,omitempty"`
	Host      string `json:"host,omitempty"`
	Expire    int64  `json:"expire,omitempty"`
	Signature string `json:"signature,omitempty"`
	Policy    string `json:"policy,omitempty"`
	Dir       string `json:"dir,omitempty"`
	Callback  string `json:"callback,omitempty"`
}
