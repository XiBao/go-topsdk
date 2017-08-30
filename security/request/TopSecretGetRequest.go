package request

import (
	"github.com/XiBao/taobao"
	"github.com/XiBao/topsdk"
)

type TopSecretGetRequest struct {
	Request *topsdk.Request
	ApiKey  *taobao.ApiKey
	Session string
}

// create new request
func NewTopSecretGetRequest(apiKey *taobao.ApiKey, session string) (req *TopSecretGetRequest) {
	request := topsdk.Request{MethodName: "taobao.top.secret.get", Params: make(map[string]interface{})}
	req = &TopSecretGetRequest{
		Request: &request,
		ApiKey:  apiKey,
		Session: session,
	}
	return
}

func (req *TopSecretGetRequest) SetSecretVersion(secretVersion int) {
	req.Request.Params["secret_version"] = secretVersion
}

func (req *TopSecretGetRequest) GetSecretVersion() int {
	secretVersion, found := req.Request.Params["secret_version"]
	if found {
		return secretVersion.(int)
	}
	return 0
}

func (req *TopSecretGetRequest) SetRandomNum(randomNum string) {
	req.Request.Params["random_num"] = randomNum
}

func (req *TopSecretGetRequest) GetRandomNum() string {
	startDate, found := req.Request.Params["random_num"]
	if found {
		return startDate.(string)
	}
	return ""
}

func (req *TopSecretGetRequest) SetCustomerUserId(customerUserId uint64) {
	req.Request.Params["customer_user_id"] = customerUserId
}

func (req *TopSecretGetRequest) GetCustomerUserId() uint64 {
	customerUserId, found := req.Request.Params["customer_user_id"]
	if found {
		return customerUserId.(uint64)
	}
	return 0
}
