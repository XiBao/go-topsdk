package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type FenxiaoDistributorProductsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewFenxiaoDistributorProductsGetRequest() (req *FenxiaoDistributorProductsGetRequest) {
	request := topsdk.Request{MethodName: "taobao.fenxiao.distributor.products.get", Params: make(map[string]interface{}, 13)}
	req = &FenxiaoDistributorProductsGetRequest{
		Request: &request,
	}
	return
}

// 查询时间类型，默认更新时间。MODIFIED:更新时间，CREATE:创建时间
func (req *FenxiaoDistributorProductsGetRequest) SetSupplierNick(supplierNick string) {
	req.Request.Params["supplier_nick"] = supplierNick
}

func (req *FenxiaoDistributorProductsGetRequest) GetSupplierNick() string {
	supplierNick, found := req.Request.Params["supplier_nick"]
	if found {
		return supplierNick.(string)
	}
	return ""
}

// 结束时间
func (req *FenxiaoDistributorProductsGetRequest) SetEndTime(endTime string) {
	req.Request.Params["end_time"] = endTime
}

func (req *FenxiaoDistributorProductsGetRequest) GetEndTime() string {
	endTime, found := req.Request.Params["end_time"]
	if found {
		return endTime.(string)
	}
	return ""
}

// 开始时间
func (req *FenxiaoDistributorProductsGetRequest) SetStartTime(startTime string) {
	req.Request.Params["start_time"] = startTime
}

func (req *FenxiaoDistributorProductsGetRequest) GetStartTime() string {
	startTime, found := req.Request.Params["start_time"]
	if found {
		return startTime.(string)
	}
	return ""
}

// 指定查询额外的信息，可选值：skus（sku数据）、images（多图），多个可选值用逗号分割。
func (req *FenxiaoDistributorProductsGetRequest) SetFields(fields string) {
	req.Request.Params["fields"] = fields
}

func (req *FenxiaoDistributorProductsGetRequest) GetFields() string {
	fields, found := req.Request.Params["fields"]
	if found {
		return fields.(string)
	}
	return ""
}

// 分销方式，分页查询时，必传。AGENT：代销，DEALER：经销
func (req *FenxiaoDistributorProductsGetRequest) SetTradeType(tradeType string) {
	req.Request.Params["trade_type"] = tradeType
}

func (req *FenxiaoDistributorProductsGetRequest) GetTradeType() string {
	tradeType, found := req.Request.Params["trade_type"]
	if found {
		return tradeType.(string)
	}
	return ""
}

// 下载状态，默认未下载。UNDOWNLOAD：未下载，DOWNLOADED：已下载。
func (req *FenxiaoDistributorProductsGetRequest) SetDownloadStatus(downloadStatus string) {
	req.Request.Params["download_status"] = downloadStatus
}

func (req *FenxiaoDistributorProductsGetRequest) GetDownloadStatus() string {
	downloadStatus, found := req.Request.Params["download_status"]
	if found {
		return downloadStatus.(string)
	}
	return ""
}

// 根据商品ID列表查询，优先级次于产品ID列表，高于其他分页查询条件。如果商品不是分销商品，自动过滤。最大限制20，用逗号分割，例如：“1001,1002,1003,1004,1005”
func (req *FenxiaoDistributorProductsGetRequest) SetItemIds(itemIds string) {
	req.Request.Params["item_ids"] = itemIds
}

func (req *FenxiaoDistributorProductsGetRequest) GetItemIds() string {
	itemIds, found := req.Request.Params["item_ids"]
	if found {
		return itemIds.(string)
	}
	return ""
}

// 产品ID列表，优先级最高，传了忽略其他查询条件。用逗号分割，例如：“1001,1002,1003,1004,1005”
func (req *FenxiaoDistributorProductsGetRequest) SetPids(pids string) {
	req.Request.Params["pids"] = pids
}

func (req *FenxiaoDistributorProductsGetRequest) GetPids() string {
	pids, found := req.Request.Params["pids"]
	if found {
		return pids.(string)
	}
	return ""
}

// 查询时间类型，默认更新时间。MODIFIED:更新时间，CREATE:创建时间
func (req *FenxiaoDistributorProductsGetRequest) SetTimeType(timeType string) {
	req.Request.Params["time_type"] = timeType
}

func (req *FenxiaoDistributorProductsGetRequest) GetTimeType() string {
	timeType, found := req.Request.Params["time_type"]
	if found {
		return timeType.(string)
	}
	return ""
}

// 产品线ID
func (req *FenxiaoDistributorProductsGetRequest) SetProductcatId(productcatId uint64) {
	req.Request.Params["productcat_id"] = productcatId
}

func (req *FenxiaoDistributorProductsGetRequest) GetProductcatId() uint64 {
	productcatId, found := req.Request.Params["productcat_id"]
	if found {
		return productcatId.(uint64)
	}
	return 0
}

// 排序。QUANTITY_DESC：库存降序，CREATE_TIME_DESC，创建时间降序。
func (req *FenxiaoDistributorProductsGetRequest) SetOrderBy(orderBy string) {
	req.Request.Params["order_by"] = orderBy
}

func (req *FenxiaoDistributorProductsGetRequest) GetOrderBy() string {
	orderBy, found := req.Request.Params["order_by"]
	if found {
		return orderBy.(string)
	}
	return ""
}

//页码（大于0的整数，默认1）
func (req *FenxiaoDistributorProductsGetRequest) SetPageNo(page_no uint16) {
	req.Request.Params["page_no"] = page_no
}

func (req *FenxiaoDistributorProductsGetRequest) GetPageNo() uint16 {
	page_no, found := req.Request.Params["page_no"]
	if found {
		return page_no.(uint16)
	}
	return 0
}

// 每页记录数（默认20，最大50）
func (req *FenxiaoDistributorProductsGetRequest) SetPageSize(page_size uint16) {
	req.Request.Params["page_size"] = page_size
}

func (req *FenxiaoDistributorProductsGetRequest) GetPageSize() uint16 {
	page_size, found := req.Request.Params["page_size"]
	if found {
		return page_size.(uint16)
	}
	return 0
}
