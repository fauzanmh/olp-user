package auth

import (
	"context"
	"fmt"
	"net/http"

	entity "github.com/fauzanmh/olp-user/entity/microservice"
	appInit "github.com/fauzanmh/olp-user/init"
	"github.com/fauzanmh/olp-user/pkg/helper"
	"go.uber.org/zap"
)

type Auth struct {
	config *appInit.Config
}

// NewProviderAuth :nodoc:
func NewProviderAuth(config *appInit.Config) AuthAdapter {
	return &Auth{
		config: config,
	}
}

// --- CreateUser --- //
func (auth *Auth) CreateUser(ctx context.Context, req *entity.CreateUserRequest) (err error) {
	url := fmt.Sprintf("%s%s", auth.config.Microservice.Auth.BaseURL, auth.config.Microservice.Auth.CreateUser)
	httpMethod := http.MethodPost
	headers := map[string]string{}
	payload, err := helper.ParseRequestBody(req)
	if err != nil {
		return
	}
	client := helper.APICall{
		URL:       url,
		Method:    httpMethod,
		FormParam: payload,
		Header:    headers,
	}

	// log request
	zap.S().Named("auth.create-user.request").Info(helper.ConstructRequestLog(url, httpMethod, headers, req))

	response, err := client.CallWithJson(ctx)
	if err != nil {
		return
	}

	// log response
	zap.S().Named("auth.create-user.response").Info(response.Body)

	if response.StatusCode != 200 {
		err = fmt.Errorf("got unsucceful response from auth api")
		return
	}

	return
}

// --- DeleteUser --- //
func (auth *Auth) DeleteUser(ctx context.Context, id int64) (err error) {
	url := fmt.Sprintf("%s%s/%d", auth.config.Microservice.Auth.BaseURL, auth.config.Microservice.Auth.DeleteUser, id)
	httpMethod := http.MethodDelete
	headers := map[string]string{}
	client := helper.APICall{
		URL:    url,
		Method: httpMethod,
		Header: headers,
	}

	// log request
	zap.S().Named("auth.delete-user.request").Info(helper.ConstructRequestLog(url, httpMethod, headers, nil))

	response, err := client.CallWithJson(ctx)
	if err != nil {
		return
	}

	// log response
	zap.S().Named("auth.delete-user.response").Info(response.Body)

	if response.StatusCode != 200 {
		err = fmt.Errorf("got unsucceful response from auth api")
		return
	}

	return
}
