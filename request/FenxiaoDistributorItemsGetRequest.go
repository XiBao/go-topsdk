package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type FenxiaoDistributorItemsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewFenxiaoDistributorItemsGetRequest() (req *FenxiaoDistributorItemsGetRequest) {
	request := topsdk.Request{MethodName: "taobao.fenxiao.distributor.items.get", Params: make(map[string]interface{}, 6)}
	req = &FenxiaoDistributorItemsGetRequest{
		Request: &request,
	}
	return
}

// 设置结束时间,空为不设置。
func (req *FenxiaoDistributorItemsGetRequest) SetEndModified(endModified string) {
	req.Request.Params["end_modified"] = endModified
}

func (req *FenxiaoDistributorItemsGetRequest) GetEndModified() string {
	endModified, found := req.Request.Params["end_modified"]
	if found {
		return endModified.(string)
	}
	return ""
}

// 设置开始时间。空为不设置。
func (req *FenxiaoDistributorItemsGetRequest) SetStartModified(startModified string) {
	req.Request.Params["start_modified"] = startModified
}

func (req *FenxiaoDistributorItemsGetRequest) GetStartModified() string {
	startModified, found := req.Request.Params["start_modified"]
	if found {
		return startModified.(string)
	}
	return ""
}

// 分销商ID
func (req *FenxiaoDistributorItemsGetRequest) SetDistributorId(distributorId uint64) {
	req.Request.Params["distributor_id"] = distributorId
}

func (req *FenxiaoDistributorItemsGetRequest) GetDistributorId() uint64 {
	distributorId, found := req.Request.Params["distributor_id"]
	if found {
		return distributorId.(uint64)
	}
	return 0
}

// 产品ID
func (req *FenxiaoDistributorItemsGetRequest) SetProductId(productId uint64) {
	req.Request.Params["product_id"] = productId
}

func (req *FenxiaoDistributorItemsGetRequest) GetProductId() uint64 {
	productId, found := req.Request.Params["product_id"]
	if found {
		return productId.(uint64)
	}
	return 0
}

//页码（大于0的整数，默认1）
func (req *FenxiaoDistributorItemsGetRequest) SetPageNo(page_no uint16) {
	req.Request.Params["page_no"] = page_no
}

func (req *FenxiaoDistributorItemsGetRequest) GetPageNo() uint16 {
	page_no, found := req.Request.Params["page_no"]
	if found {
		return page_no.(uint16)
	}
	return 0
}

// 每页记录数（默认20，最大50）
func (req *FenxiaoDistributorItemsGetRequest) SetPageSize(page_size uint16) {
	req.Request.Params["page_size"] = page_size
}

func (req *FenxiaoDistributorItemsGetRequest) GetPageSize() uint16 {
	page_size, found := req.Request.Params["page_size"]
	if found {
		return page_size.(uint16)
	}
	return 0
}
