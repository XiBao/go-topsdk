package tanx

type Advertiser struct {
	Uid                  uint64 `json:"-" codec:",omitempty"`
	EnglishName          string `json:"english_name,omitempty" codec:",omitempty"`           // 英文名称
	NickName             string `json:"nick_name,omitempty" codec:",omitempty"`              // 昵称
	AdvertiserType       uint8  `json:"advertiser_type,omitempty" codec:",omitempty"`        //用二进制存储广告主属性1.品牌广告主 2. VIP 4. 世界500强
	UserType             uint8  `json:"user_type" codec:",omitempty"`                        // 广告主类别（0-淘宝，1-天猫，2-dsp公司，3-dsp个人）
	AdvertiserName       string `json:"advertiser_name,omitempty" codec:",omitempty"`        // 广告主名称
	AdvertiserId         uint64 `json:"advertiser_id,omitempty" codec:",omitempty"`          // 广告主id
	Status               uint8  `json:"status,omitempty" codec:",omitempty"`                 // 审核状态 0待审核，-1拒绝，1审核通过
	Domains              string `json:"domains,omitempty" codec:",omitempty"`                // 主域名
	ExceptAdvertiserType uint8  `json:"except_advertiser_type,omitempty" codec:",omitempty"` // 查询时排除的用户特性
	MemberId             uint64 `json:"member_id,omitempty" codec:",omitempty"`              // dspid，这里无意义
	Comments             string `json:"comments,omitempty" codec:",omitempty"`               // 备注
}
