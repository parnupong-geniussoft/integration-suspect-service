package repositories

import (
	"integration-suspect-service/modules/suspect/entities"
	"integration-suspect-service/pkg/clients/ktb"

	"github.com/gofiber/fiber/v2"
)

type SuspectRepository interface {
	SubmitKtbSuspect(ctx *fiber.Ctx, ktbSubmitBody *entities.KtbIndividualSubmitSuspectRequest, xCorrelationID string, referenceId string) (*entities.SuspectResponse, error)
}

type suspectRepo struct {
	KtbClient ktb.KtbClient
}

func NewSuspectRepository(ktbClient ktb.KtbClient) SuspectRepository {
	return &suspectRepo{
		KtbClient: ktbClient,
	}
}

func (r *suspectRepo) SubmitKtbSuspect(ctx *fiber.Ctx, ktbSubmitBody *entities.KtbIndividualSubmitSuspectRequest, xCorrelationID string, referenceId string) (*entities.SuspectResponse, error) {
	return r.KtbClient.SubmitKtbSuspect(ctx, ktbSubmitBody, xCorrelationID, referenceId)
}
