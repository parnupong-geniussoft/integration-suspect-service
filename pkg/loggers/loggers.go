package loggers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type Logger interface {
	SystemLogger(ctx *fiber.Ctx, start time.Time, err error)
	DbLogger(ctx *fiber.Ctx) error
	OnBeforeRequest(client *resty.Client, request *resty.Request) error
	OnAfterResponse(client *resty.Client, resp *resty.Response) error
	OnErrorRetryHook(resp *resty.Response, errors error)
}

type logger struct {
	Db *sqlx.DB
}

func NewLogger(db *sqlx.DB) Logger {
	return &logger{
		Db: db,
	}
}

func (l *logger) SystemLogger(ctx *fiber.Ctx, start time.Time, err error) {
	stop := time.Now()
	latency := stop.Sub(start)
	status := ctx.Response().StatusCode()
	method := ctx.Method()
	ip := ctx.IP()
	path := ctx.Path()

	errMessage := "-"
	if err != nil {
		errMessage = err.Error()
	}

	log.Printf("%s | %3d | %10s | %-15s | %-5s | %-60s | %s",
		stop.Format("15:04:05"),
		status,
		latency,
		ip,
		method,
		path,
		errMessage,
	)
}

func (l *logger) DbLogger(ctx *fiber.Ctx) error {
	timeNow := time.Now()

	xcc := fmt.Sprintf("%d", timeNow.UnixNano())
	ctx.Request().Header.Add("x_correlation_id", xcc)

	data := LoggerStruct{
		CreatedAt:      timeNow,
		Level:          "info",
		Type:           "request",
		Method:         ctx.Method(),
		Path:           ctx.Path(),
		Ip:             ctx.IP(),
		DurationMs:     0,
		RequestDate:    timeNow,
		XCorrelationId: xcc,
	}

	data.GetReferenceId(ctx)
	ctx.Locals("referenceId", data.ReferenceId)

	data.MaskBodyRequest(ctx)
	data.HeaderConvert(ctx)
	loggerDbErr := SaveLoggerDb(data, l.Db)
	if loggerDbErr != nil {
		return errors.New("Can't save logger to db: " + loggerDbErr.Error())
	}

	if err := ctx.Next(); err != nil {
		data.HandleError(ctx, err)
		data.HeaderConvertResponse(ctx)
		loggerDbErr := SaveLoggerDb(data, l.Db)
		if loggerDbErr != nil {
			return errors.New("Can't save logger to db: " + loggerDbErr.Error())
		}
		return err
	}

	data.HeaderConvertResponse(ctx)
	data.HandleResponse(ctx)
	loggerDbErr = SaveLoggerDb(data, l.Db)
	if loggerDbErr != nil {
		return errors.New("Can't save logger to db: " + loggerDbErr.Error())
	}

	return nil
}

// p = path , maker, body
func HandlerBodyMask(p string, mk []MaskData, body []byte) []byte {
	maskData := FindMasker(p, mk)
	maskedBody, err := MaskBody(body, maskData.Key)
	if err != nil {
		return body
	} else {
		return maskedBody
	}

}

func MaskBody(x []byte, keys []string) ([]byte, error) {
	var bodyMap map[string]interface{}

	if err := json.Unmarshal(x, &bodyMap); err != nil {
		return nil, fmt.Errorf("maskLoop: %w", err)
	}

	for _, key := range keys {
		if secret, exists := bodyMap[key]; exists && secret != "" {
			bodyMap[key] = "*****"
		}
	}

	mBody, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, fmt.Errorf("maskLoop: %w", err)
	}

	return mBody, nil
}

func FindMasker(path string, masker []MaskData) MaskData {
	for _, v := range masker {
		if v.Path == path {
			return v
		}
	}
	return MaskData{}
}

func SaveLoggerDb(data LoggerStruct, Db *sqlx.DB) error {
	_, err := Db.NamedExec(`INSERT INTO logger (
		created_at, level, type, method, path, ip, message, duration_ms, header,
		request_date, request, x_correlation_id, reference_id
	) VALUES (
		:created_at, :level, :type, :method, :path, :ip, :message, :duration_ms, :header,
		:request_date, :request, :x_correlation_id, :reference_id
	)`, data)
	if err != nil {
		fmt.Println("logger save", err)
		return err
	}
	return nil
}

func SaveLoggerDbAsync(data LoggerStruct, Db *sqlx.DB) {
	go func(data LoggerStruct) {
		_, err := Db.NamedExec(`INSERT INTO logger (
		created_at, level, type, method, path, ip, message, duration_ms, header,
		request_date, request, x_correlation_id, reference_id
	) VALUES (
		:created_at, :level, :type, :method, :path, :ip, :message, :duration_ms, :header,
		:request_date, :request, :x_correlation_id, :reference_id
	)`, data)
		if err != nil {
			fmt.Println("logger save", err)
		}
	}(data)
}
