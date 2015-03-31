package tanx

type Qualification struct {
	ElementId   uint64 `json:"element_id,omitempty" codec:",omitempty"`   // 上传的客户资质元素id
	UserId      uint64 `json:"user_id,omitempty" codec:",omitempty"`      // 和资质平台交互的userId
	EndTime     string `json:"end_time,omitempty" codec:",omitempty"`     // 资质失效时间
	StartTime   string `json:"start_time,omitempty" codec:",omitempty"`   // 资质生效时间
	Name        string `json:"name,omitempty" codec:",omitempty"`         // 为本次上传客户资质的操作取一个名称,如果为空则系统自动生成一个(选填)
	Supplement  string `json:"supplement,omitempty" codec:",omitempty"`   // 用户附加内容(明星资质的备注)
	UrlContents string `json:"url_contents,omitempty" codec:",omitempty"` // 资质内容,如果是图片请先调用uploadQualificationPic转换成url上传，在填入返回url
}
