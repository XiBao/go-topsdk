package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type DmpDatauploadInsertrecordsRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpDatauploadInsertrecordsRequest() (req *DmpDatauploadInsertrecordsRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.dataupload.insertrecords", Params: make(map[string]interface{}, 5)}
	req = &DmpDatauploadInsertrecordsRequest{
		Request: &request,
	}
	return
}

// 数据块Id
func (req *DmpDatauploadInsertrecordsRequest) SetBlockid(blockid uint) {
	req.Request.Params["blockid"] = blockid
}

func (req *DmpDatauploadInsertrecordsRequest) GetBlockid() uint {
	blockid, found := req.Request.Params["blockid"]
	if found {
		return blockid.(uint)
	}
	return 0
}

// 表名
func (req *DmpDatauploadInsertrecordsRequest) SetTablecode(tablecode string) {
	req.Request.Params["tablecode"] = tablecode
}

func (req *DmpDatauploadInsertrecordsRequest) GetTablecode() string {
	tablecode, found := req.Request.Params["tablecode"]
	if found {
		return tablecode.(string)
	}
	return ""
}

// 列分隔符
func (req *DmpDatauploadInsertrecordsRequest) SetColumnsplit(columnsplit string) {
	req.Request.Params["columnsplit"] = columnsplit
}

func (req *DmpDatauploadInsertrecordsRequest) GetColumnsplit() string {
	columnsplit, found := req.Request.Params["columnsplit"]
	if found {
		return columnsplit.(string)
	}
	return ""
}

// 行分隔符
func (req *DmpDatauploadInsertrecordsRequest) SetRowsplit(rowsplit string) {
	req.Request.Params["rowsplit"] = rowsplit
}

func (req *DmpDatauploadInsertrecordsRequest) GetRowsplit() string {
	rowsplit, found := req.Request.Params["rowsplit"]
	if found {
		return rowsplit.(string)
	}
	return ""
}

// 上传具体数据值
func (req *DmpDatauploadInsertrecordsRequest) SetRecords(records string) {
	req.Request.Params["records"] = records
}

func (req *DmpDatauploadInsertrecordsRequest) GetRecords() string {
	records, found := req.Request.Params["records"]
	if found {
		return records.(string)
	}
	return ""
}
