package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type FenxiaoProductsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewFenxiaoProductsGetRequest() (req *FenxiaoProductsGetRequest) {
	request := topsdk.Request{MethodName: "taobao.fenxiao.products.get", Params: make(map[string]interface{}, 13)}
	req = &FenxiaoProductsGetRequest{
		Request: &request,
	}
	return
}

// 结束修改时间
func (req *FenxiaoProductsGetRequest) SetEndModified(endModified string) {
	req.Request.Params["end_modified"] = endModified
}

func (req *FenxiaoProductsGetRequest) GetEndModified() string {
	endModified, found := req.Request.Params["end_modified"]
	if found {
		return endModified.(string)
	}
	return ""
}

// 开始修改时间
func (req *FenxiaoProductsGetRequest) SetStartModified(startModified string) {
	req.Request.Params["start_modified"] = startModified
}

func (req *FenxiaoProductsGetRequest) GetStartModified() string {
	startModified, found := req.Request.Params["start_modified"]
	if found {
		return startModified.(string)
	}
	return ""
}

// 指定查询额外的信息，可选值：skus（sku数据）、images（多图），多个可选值用逗号分割。
func (req *FenxiaoProductsGetRequest) SetFields(fields string) {
	req.Request.Params["fields"] = fields
}

func (req *FenxiaoProductsGetRequest) GetFields() string {
	fields, found := req.Request.Params["fields"]
	if found {
		return fields.(string)
	}
	return ""
}

// 查询产品列表时，查询入参“是否需要授权”
func (req *FenxiaoProductsGetRequest) SetIsAuthZ(isAuthz string) {
	req.Request.Params["is_authz"] = isAuthz
}

func (req *FenxiaoProductsGetRequest) GetIsAuthZ() string {
	isAuthz, found := req.Request.Params["is_authz"]
	if found {
		return isAuthz.(string)
	}
	return ""
}

// 可根据导入的商品ID列表查询，优先级次于产品ID、sku_numbers，高于其他分页查询条件。最大限制20，用逗号分割，例如：“1001,1002,1003,1004,1005”
func (req *FenxiaoProductsGetRequest) SetItemIds(itemIds string) {
	req.Request.Params["item_ids"] = itemIds
}

func (req *FenxiaoProductsGetRequest) GetItemIds() string {
	itemIds, found := req.Request.Params["item_ids"]
	if found {
		return itemIds.(string)
	}
	return ""
}

// 产品ID列表（最大限制30），用逗号分割，例如：“1001,1002,1003,1004,1005”
func (req *FenxiaoProductsGetRequest) SetPids(pids string) {
	req.Request.Params["pids"] = pids
}

func (req *FenxiaoProductsGetRequest) GetPids() string {
	pids, found := req.Request.Params["pids"]
	if found {
		return pids.(string)
	}
	return ""
}

//商家编码
func (req *FenxiaoProductsGetRequest) SetOuterId(outerId string) {
	req.Request.Params["outer_id"] = outerId
}

func (req *FenxiaoProductsGetRequest) GetOuterId() string {
	outerId, found := req.Request.Params["outer_id"]
	if found {
		return outerId.(string)
	}
	return ""
}

// 产品线ID
func (req *FenxiaoProductsGetRequest) SetProductcatId(productcatId uint64) {
	req.Request.Params["productcat_id"] = productcatId
}

func (req *FenxiaoProductsGetRequest) GetProductcatId() uint64 {
	productcatId, found := req.Request.Params["productcat_id"]
	if found {
		return productcatId.(uint64)
	}
	return 0
}

//sku商家编码
func (req *FenxiaoProductsGetRequest) SetSkuNumber(skuNumber string) {
	req.Request.Params["sku_number"] = skuNumber
}

func (req *FenxiaoProductsGetRequest) GetSkuNumber() string {
	skuNumber, found := req.Request.Params["sku_number"]
	if found {
		return skuNumber.(string)
	}
	return ""
}

//产品状态，可选值：up（上架）、down（下架），不传默认查询所有
func (req *FenxiaoProductsGetRequest) SetStatus(status string) {
	req.Request.Params["status"] = status
}

func (req *FenxiaoProductsGetRequest) GetStatus() string {
	status, found := req.Request.Params["status"]
	if found {
		return status.(string)
	}
	return ""
}

//页码（大于0的整数，默认1）
func (req *FenxiaoProductsGetRequest) SetPageNo(page_no uint16) {
	req.Request.Params["page_no"] = page_no
}

func (req *FenxiaoProductsGetRequest) GetPageNo() uint16 {
	page_no, found := req.Request.Params["page_no"]
	if found {
		return page_no.(uint16)
	}
	return 0
}

// 每页记录数（默认20，最大50）
func (req *FenxiaoProductsGetRequest) SetPageSize(page_size uint16) {
	req.Request.Params["page_size"] = page_size
}

func (req *FenxiaoProductsGetRequest) GetPageSize() uint16 {
	page_size, found := req.Request.Params["page_size"]
	if found {
		return page_size.(uint16)
	}
	return 0
}
