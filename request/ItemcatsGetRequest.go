package TopRequest

import (
	"github.com/XiBao/topsdk"
	"strconv"
	"strings"
)

type ItemcatsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewItemcatsGetRequest() (req *ItemcatsGetRequest) {
	request := topsdk.Request{MethodName: "taobao.itemcats.get", Params: make(map[string]interface{}, 3)}
	req = &ItemcatsGetRequest{
		Request: &request,
	}
	return
}

// 需要返回的字段列表，见ItemCat，默认返回：cid,parent_cid,name,is_parent；增量类目信息,根据fields传入的参数返回相应的结果。 features字段： 1、如果存在attr_key=freeze表示该类目被冻结了，attr_value=0,5，value可能存在2个值（也可能只有1个），用逗号分割，0表示禁编辑，5表示禁止发布
func (req *ItemcatsGetRequest) SetFields(fields string) {
	req.Request.Params["fields"] = fields
}

func (req *ItemcatsGetRequest) GetFields() string {
	fields, found := req.Request.Params["fields"]
	if found {
		return fields.(string)
	}
	return ""
}

// 父商品类目 id，0表示根节点, 传输该参数返回所有子类目。 (cids、parent_cid至少传一个)
func (req *ItemcatsGetRequest) SetParentCid(parentCid uint64) {
	req.Request.Params["parent_cid"] = parentCid
}

func (req *ItemcatsGetRequest) GetParentCid() uint64 {
	parentCid, found := req.Request.Params["parent_cid"]
	if found {
		return parentCid.(uint64)
	}
	return 0
}

// 商品所属类目ID列表，用半角逗号(,)分隔 例如:(18957,19562,) (cids、parent_cid至少传一个)
func (req *ItemcatsGetRequest) SetCids(cids []uint64) {
	var cidStr []string
	for _, cid := range cids {
		cidStr = append(cidStr, strconv.FormatUint(cid, 10))
	}
	req.Request.Params["cids"] = strings.Join(cidStr, ",")
}

func (req *ItemcatsGetRequest) GetCids() []uint64 {
	cidsStr, found := req.Request.Params["cids"]
	if found {
		var cids []uint64
		cidArr := strings.Split(cidsStr.(string), ",")
		for _, cidStr := range cidArr {
			cid, err := strconv.ParseUint(cidStr, 10, 64)
			if err != nil {
				continue
			}
			cids = append(cids, cid)
		}
		return cids
	}
	return nil
}
