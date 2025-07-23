package loggers

import (
	"encoding/json"
	"fmt"
	"integration-suspect-service/pkg/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

type MaskData struct {
	Path string   `json:"path"`
	Key  []string `json:"key"`
}

type ReferenceData struct {
	Path string   `json:"path"`
	Key  []string `json:"key"`
}

var MaskersRequest = []MaskData{{Path: "/v1/integration-api/request_token", Key: []string{"client_secret"}}}

var MaskersResponse = []MaskData{{Path: "/v1/integration-api/request_token", Key: []string{"access_token"}}}

var referenceKey = []ReferenceData{
	{
		Path: "/v1/integration-api/fraud_check",
		Key: []string{
			"Application_Number",
		},
	},
	{
		Path: "/v1/integration-api/fraud_staff_action",
		Key: []string{
			"Application_Number",
		},
	},
	{
		Path: "/dgl",
		Key: []string{
			"Application_Number",
		},
	},
}

type LoggerStruct struct {
	CreatedAt      time.Time `db:"created_at"`
	Level          string    `db:"level"`
	Type           string    `db:"type"`
	Method         string    `db:"method"`
	Path           string    `db:"path"`
	Ip             string    `db:"ip"`
	Message        string    `db:"message"`
	DurationMs     int64     `db:"duration_ms"`
	Header         []byte    `db:"header"`
	Request        string    `db:"request"`
	RequestDate    time.Time `db:"request_date"`
	XCorrelationId string    `db:"x_correlation_id"`
	ReferenceId    string    `db:"reference_id"`
}

func (data *LoggerStruct) HandleResponse(ctx *fiber.Ctx) {
	mBody := HandlerBodyMask(ctx.Path(), MaskersResponse, ctx.Response().Body())
	data.Request = string(mBody)
	data.CreatedAt = time.Now()
	durationMs := utils.DurationMS(data.RequestDate)
	data.Type = "response"
	data.DurationMs = durationMs
}

func (data *LoggerStruct) MaskBodyRequest(ctx *fiber.Ctx) {
	mBody := HandlerBodyMask(ctx.Path(), MaskersRequest, ctx.Body())
	data.Request = string(mBody)
}

func (data *LoggerStruct) HandleError(ctx *fiber.Ctx, err error) {
	data.CreatedAt = time.Now()
	durationMs := utils.DurationMS(data.RequestDate)
	data.Type = "error"
	data.DurationMs = durationMs
	data.Message = err.Error()
}

func (data *LoggerStruct) HeaderConvert(ctx *fiber.Ctx) {
	headers := ctx.Request().Header
	headerMap := make(map[string]string)
	headers.VisitAll(func(key, value []byte) {
		if string(key) == "Authorization" {
			headerMap[string(key)] = "*****"
			return
		}
		headerMap[string(key)] = string(value)
	})
	headerConverted, _ := json.Marshal(headerMap)

	data.Header = headerConverted
}

func (data *LoggerStruct) HeaderConvertResponse(ctx *fiber.Ctx) {
	headers := ctx.Response().Header
	headerMap := make(map[string]string)
	headers.VisitAll(func(key, value []byte) {
		if string(key) == "Authorization" {
			headerMap[string(key)] = "*****"
			return
		}
		headerMap[string(key)] = string(value)
	})

	headerConverted, _ := json.Marshal(headerMap)

	data.Header = headerConverted
}

func (data *LoggerStruct) GetReferenceId(ctx *fiber.Ctx) {
	var referenceId *string
	var err error

	referenceKey := findReferenceKey(ctx.Path(), referenceKey)
	if ctx.Path() == "/v1/integration-api/suspect/add_suspect" {
		referenceId, err = getReferenceIdFromBodyForAddSuspect(ctx.Body())
		if err != nil {
			return
		}
	} else {
		referenceId, err = getReferenceIdFromBody(ctx.Body(), referenceKey.Key)
		if err != nil {
			return
		}
	}

	data.ReferenceId = *referenceId
}

func findReferenceKey(path string, referenceData []ReferenceData) ReferenceData {
	for _, v := range referenceData {
		if v.Path == path {
			return v
		}
	}
	return ReferenceData{}
}

func getReferenceIdFromBody(x []byte, keys []string) (*string, error) {
	var bodyMap map[string]interface{}
	var referenceId string

	if err := json.Unmarshal(x, &bodyMap); err != nil {
		return nil, fmt.Errorf("referenceIdFromBodyLoop: %w", err)
	}

	for _, key := range keys {
		value, exists := bodyMap[key]
		if exists && value != "" {
			strValue, ok := value.(string)
			if ok {
				referenceId = strValue
			}
		}
	}

	return &referenceId, nil
}

func getReferenceIdFromBodyForAddSuspect(x []byte) (*string, error) {
	var bodyMap map[string]interface{}
	var referenceId string

	if err := json.Unmarshal(x, &bodyMap); err != nil {
		return nil, fmt.Errorf("referenceIdFromBodyLoop: %w", err)
	}

	switch bodyMap["entityTP"] {
	case "PERSON":
		citizenId, exists := bodyMap["citizen_id"]
		if exists && citizenId != "" {
			strValue, ok := citizenId.(string)
			if ok {
				referenceId = strValue
			}
		} else {
			passportId, exists := bodyMap["passport_id"]
			if exists && passportId != "" {
				strValue, ok := passportId.(string)
				if ok {
					referenceId = strValue
				}
			}
		}
	case "ENTITY":
		juristicId, exists := bodyMap["juristic_id"]
		if exists && juristicId != "" {
			strValue, ok := juristicId.(string)
			if ok {
				referenceId = strValue
			}
		}
	}

	return &referenceId, nil
}
