package resty_factory

import (
	"crypto/tls"
	"integration-suspect-service/configs"
	"integration-suspect-service/pkg/loggers"
	"net/http"

	"github.com/go-resty/resty/v2"
)

func NewRestyClient(cfg *configs.Configs, logger loggers.Logger, baseUrl string) *resty.Client {
	client := resty.New().
		EnableTrace().
		SetBaseURL(baseUrl).
		SetTLSClientConfig(&tls.Config{
			InsecureSkipVerify: true,
		}).
		SetRetryCount(cfg.Retry.RetryCount).
		SetRetryWaitTime(cfg.Retry.RetryMinWaitTimeSecond).
		SetRetryMaxWaitTime(cfg.Retry.RetryMaxWaitTimeSecond).
		AddRetryCondition(func(r *resty.Response, err error) bool {
			return r.StatusCode() != http.StatusOK && r.StatusCode() > 0
		}).
		AddRetryHook(func(r *resty.Response, err error) {
			logger.OnErrorRetryHook(r, err)
		})

	client.OnBeforeRequest(logger.OnBeforeRequest)
	client.OnAfterResponse(logger.OnAfterResponse)

	return client
}
