package ktb

import (
	"context"
	"errors"
	"fmt"
	"integration-suspect-service/modules/suspect/entities"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

type KtbClient interface {
	SubmitKtbSuspect(ctx *fiber.Ctx, ktbSubmitBody *entities.KtbIndividualSubmitSuspectRequest, xCorrelationID string, referenceId string) (*entities.SuspectResponse, error)
}

type ktbClient struct {
	Client *resty.Client
}

func NewKtbClient(client *resty.Client) KtbClient {
	return &ktbClient{
		Client: client,
	}
}

func (c *ktbClient) SubmitKtbSuspect(ctx *fiber.Ctx, ktbSubmitBody *entities.KtbIndividualSubmitSuspectRequest, xCorrelationID string, referenceId string) (*entities.SuspectResponse, error) {
	resp, err := c.Client.R().
		SetContext(context.WithValue(ctx.Context(), "referenceId", referenceId)).
		EnableTrace().
		SetBody(ktbSubmitBody).
		SetHeader("Content-Type", "application/json").
		SetHeader("x_correlation_id", xCorrelationID).
		SetResult(&entities.SuspectResponse{}).
		Post("/susp-addListDetail")

	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}

	result := resp.Result().(*entities.SuspectResponse)
	if resp.IsError() {
		fmt.Println("suspect response : ", resp.RawResponse)
		return nil, errors.New(resp.Status())
	}

	return result, nil
}
