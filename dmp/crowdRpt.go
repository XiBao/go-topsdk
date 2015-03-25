package dmp

type CrowdRpt struct {
	AppSystem       string `json:"app_system,omitempty" codec:",omitempty"`        // 应用系统
	Finprice        uint   `json:"finprice,omitempty" codec:",omitempty"`          // 消耗(单位:分)
	Click           uint   `json:"click,omitempty" codec:",omitempty"`             // click
	AdzoneId        string `json:"adzone_id,omitempty" codec:",omitempty"`         // 广告位ID
	ECollectItemcnt string `json:"e_collect_itemcnt,omitempty" codec:",omitempty"` // 广告位名称
	EAlipayAmt      uint   `json:"e_alipay_amt,omitempty" codec:",omitempty"`      // 三天直接+间接 alipay成交金额(单位:分)
	AdzoneName      string `json:"adzone_name,omitempty" codec:",omitempty"`       // 广告位名称
	RecordOn        string `json:"the_date,omitempty" codec:",omitempty"`          // 日期
	AppSystemName   string `json:"app_system_name,omitempty" codec:",omitempty"`   // 应用系统名称
	CrowdId         uint64 `json:"crowd_id,omitempty" codec:",omitempty"`          // 人群id
	MemberId        uint64 `json:"member_id,omitempty" codec:",omitempty"`         // 用户id
	Pv              uint64 `json:"pv,omitempty" codec:",omitempty"`                // pv
	CrowdName       string `json:"crowd_name,omitempty" codec:",omitempty"`        // 人群名称
}
