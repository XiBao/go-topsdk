package request

import (
	"encoding/json"
	"fmt"
	"github.com/XiBao/topsdk"
)

type TopSecret struct {
	Interval    int
	Secret      string
	Version     int
	MaxInterval int
	AppConfig   map[string]string
}

func TopSecretGet(req *TopSecretGetRequest) (topSecret *TopSecret, err error) {
	c := topsdk.NewClient(req.ApiKey.Key, req.ApiKey.Secret)
	result, err := c.Execute(req.Request, req.Session)
	if err != nil {
		return nil, err
	}
	v, ok := result.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid result: %v", result)
	}
	var appConfig map[string]string
	err = json.Unmarshal([]byte(StringAssert(v["app_config"])), &appConfig)
	if err != nil {
		return nil, err
	}

	topSecret = &TopSecret{
		Interval:    IntAssert(v["interval"]),
		Secret:      StringAssert(v["secret"]),
		Version:     IntAssert(v["secret_version"]),
		MaxInterval: IntAssert(v["max_interval"]),
		AppConfig:   appConfig,
	}
	return topSecret, nil
}
