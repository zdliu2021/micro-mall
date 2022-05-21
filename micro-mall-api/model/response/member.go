package response

type MemberEntity struct {
	Id          int64  `json:"id"`          // id
	LevelId     int64  `json:"levelId"`     // 会员等级ID
	UserName    string `json:"userName"`    // 用户名
	Password    string `json:"password"`    // 密码
	Nickname    string `json:"nickname"`    // 昵称
	Mobile      string `json:"mobile"`      // 手机号码
	Email       string `json:"email"`       // 邮箱
	Header      string `json:"header"`      // 头像
	Gender      int32  `json:"gender"`      // 性别
	Birth       string `json:"birth"`       // 生日
	City        string `json:"city"`        // 所在城市
	Job         string `json:"job"`         // 职业
	Sign        string `json:"sign"`        // 个性签名
	SourceType  int32  `json:"sourceType"`  // 用户来源
	Integration int32  `json:"integration"` // 积分
	Growth      int32  `json:"growth"`      // 成长值
	Status      int32  `json:"status"`      // 启用状态
}

type MemberLevelEntity struct {
	Id                    int64   `json:"id"`
	Name                  string  ` json:"name"`
	GrowthPoint           int32   ` json:"growthPoint"`
	DefaultStatus         int32   `json:"defaultStatus"`
	FreeFreightPoint      float32 `json:"freeFreightPoint"`
	CommentGrowthPoint    int32   ` json:"commentGrowthPoint"`
	PriviledgeFreeFreight int32   ` json:"priviledgeFreeFreight"`
	PriviledgeMemberPrice int32   `json:"priviledgeMemberPrice"`
	PriviledgeBirthday    int32   ` json:"priviledgeBirthday"`
	Note                  string  `json:"note"`
}
