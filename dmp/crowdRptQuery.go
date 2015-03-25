package dmp

type CrowdRptQuery struct {
	StartDate string `json:"start_date"` // 报表的查询截至时间
	EndDate   string `json:"end_date"`   // 报表的查询起始时间
}
