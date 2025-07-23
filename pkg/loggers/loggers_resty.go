package loggers

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

func (l *logger) OnBeforeRequest(client *resty.Client, request *resty.Request) error {
	bodyBytes, _ := json.Marshal(request.Body)
	timeNow := time.Now()
	headers := request.Header
	headersJSON, err := json.Marshal(headers)
	if err != nil {
		return err
	}
	headersBytes := []byte(headersJSON)

	xCorrelationIdHeader := headers.Get("x_correlation_id")

	refId := ""
	if ctxRefId := request.Context().Value("referenceId"); ctxRefId != nil {
		refId = fmt.Sprintf("%v", ctxRefId)
	}

	data := LoggerStruct{
		CreatedAt:      timeNow,
		Level:          "info",
		Type:           "outgoing-request",
		Method:         request.Method,
		Path:           request.URL,
		Ip:             "0.0.0.0",
		DurationMs:     0,
		RequestDate:    timeNow,
		Header:         headersBytes,
		Request:        string(bodyBytes),
		XCorrelationId: xCorrelationIdHeader,
		ReferenceId:    refId,
	}

	loggerDbErr := SaveLoggerDb(data, l.Db)
	if loggerDbErr != nil {
		return errors.New("Can't save logger to db: " + loggerDbErr.Error())
	}
	request.SetHeader("x_correlation_id", xCorrelationIdHeader)
	return nil
}

func (l *logger) OnAfterResponse(client *resty.Client, resp *resty.Response) error {
	timeNow := time.Now()
	headers := resp.Header()
	ResponseCode := fmt.Sprintf("%d", resp.StatusCode())
	headers.Add("response_code", ResponseCode)
	headerReq := resp.Request.Header

	appNumHeader := headerReq.Get("application_number")
	if appNumHeader != "" {
		headers.Add("application_number", appNumHeader)
	}

	appNumOriginHeader := headerReq.Get("application_original_number")
	if appNumHeader != "" {
		headers.Add("application_original_number", appNumOriginHeader)
	}

	headersJSON, err := json.Marshal(headers)
	if err != nil {
		return err
	}

	headersBytes := []byte(headersJSON)
	ip := resp.Request.RawRequest.RemoteAddr

	xCorrelationIdHeader := headerReq.Get("x_correlation_id")
	correlation := fmt.Sprintf("%d", timeNow.UnixNano())
	if xCorrelationIdHeader != "" {
		correlation = xCorrelationIdHeader
	}

	refId := ""
	if ctxRefId := resp.Request.Context().Value("referenceId"); ctxRefId != nil {
		refId = fmt.Sprintf("%v", ctxRefId)
	}

	data := LoggerStruct{
		CreatedAt:      timeNow,
		Level:          "info",
		Type:           "outgoing-response",
		Method:         resp.Request.Method,
		Path:           resp.Request.URL,
		Ip:             ip,
		Header:         headersBytes,
		DurationMs:     resp.Time().Milliseconds(),
		RequestDate:    timeNow,
		Request:        string(resp.Body()),
		XCorrelationId: correlation,
		ReferenceId:    refId,
	}
	loggerDbErr := SaveLoggerDb(data, l.Db)
	if loggerDbErr != nil {
		return errors.New("Can't save logger to db: " + loggerDbErr.Error())
	}

	return nil
}

func (l *logger) OnErrorRetryHook(resp *resty.Response, errors error) {
	timeNow := time.Now()
	headers := resp.Header()
	headersJSON, err := json.Marshal(headers)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	headersBytes := []byte(headersJSON)
	ip := resp.Request.RawRequest.RemoteAddr

	refId := ""
	if ctxRefId := resp.Request.Context().Value("referenceId"); ctxRefId != nil {
		refId = fmt.Sprintf("%v", ctxRefId)
	}

	data := LoggerStruct{
		CreatedAt:      timeNow,
		Level:          "info",
		Type:           "outgoing-response-error",
		Method:         resp.Request.Method,
		Path:           resp.Request.URL,
		Ip:             ip,
		Header:         headersBytes,
		DurationMs:     resp.Time().Milliseconds(),
		RequestDate:    timeNow,
		Request:        string(resp.Body()),
		XCorrelationId: fmt.Sprintf("%d", timeNow.UnixNano()),
		ReferenceId:    refId,
	}
	SaveLoggerDbAsync(data, l.Db)
}
