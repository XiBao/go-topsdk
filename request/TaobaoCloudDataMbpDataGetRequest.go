package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type TaobaoCloudDataMbpDataGetRequest struct {
    Request *topsdk.Request
}

func NewTaobaoCloudDataMbpDataGetRequest() (req *TaobaoCloudDataMbpDataGetRequest) {
    request := topsdk.Request{MethodName:"taobao.clouddata.mbp.data.get", Params:make(map[string]interface{}, 2)}
    req = &TaobaoCloudDataMbpDataGetRequest {
        Request: &request,
    }
    return
}

func (req *TaobaoCloudDataMbpDataGetRequest) SetSqlId(sqlId uint64) {
    req.Request.Params["sql_id"] = sqlId
}

func (req *TaobaoCloudDataMbpDataGetRequest) GetSqlId() uint64 {
    sqlId, found := req.Request.Params["sql_id"]
    if found {
        return sqlId.(uint64)
    }
    return 0
}

func (req *TaobaoCloudDataMbpDataGetRequest) SetParameters(parameters string) {
    req.Request.Params["parameter"] = parameters
}

func (req *TaobaoCloudDataMbpDataGetRequest) GetParameters() string {
    parameters, found := req.Request.Params["parameter"]
    if found {
        return parameters.(string)
    }
    return ""
}
