package usecases

import (
	"fmt"
	"integration-suspect-service/modules/suspect/entities"
	"integration-suspect-service/modules/suspect/repositories"
	"time"

	"github.com/gofiber/fiber/v2"
)

type SuspectUsecase interface {
	SubmitKtbSuspect(spr *entities.IngSuspectRequest, xCorrelationID string, ctx *fiber.Ctx) (*entities.SuspectResponse, error)
}

type suspectUsecase struct {
	SuspectRepo repositories.SuspectRepository
}

func NewSuspectUsecase(suspectRepo repositories.SuspectRepository) SuspectUsecase {
	return &suspectUsecase{
		SuspectRepo: suspectRepo,
	}
}

func (u *suspectUsecase) SubmitKtbSuspect(spr *entities.IngSuspectRequest, xCorrelationID string, ctx *fiber.Ctx) (*entities.SuspectResponse, error) {
	sourceDesc := entities.SOURCE_DESC
	if spr.SourceDesc != "" {
		sourceDesc = spr.SourceDesc
	}

	fmt.Println("sourceDesc", sourceDesc)

	ktbSubmitBody := &entities.KtbIndividualSubmitSuspectRequest{
		Control: &entities.SuspectControl{
			RequestID:             spr.UpdateBy + time.Now().Format("20060102150405"),
			RequesterName:         "SUSP",
			RequesterLanguage:     "100",
			RequesterLocale:       "en",
			PageStartIndex:        0,
			PageEndIndex:          0,
			AvailableResultsCount: 0,
		},
		ListDetail: &entities.ListDetail{
			EntityID: "",
			ListMaster: &entities.ListMaster{
				EntityID:     "",
				SourceSystem: "SUSP",
				EntityTP:     "PERSON",
				ListSupTPCD:  "1000028",
				UpdateBy:     spr.UpdateBy,
				Name:         "",
			},
			ListIdentificationInfo: &[]entities.IdentificationInfo{},
			ListNameInfo: &[]entities.NameInfo{{
				SourceSystem: "SUSP",
				FirstName:    spr.FirstName,
				LastName:     spr.LastName,
				OrgName:      spr.CompanyName,
				UpdateBy:     spr.UpdateBy,
				Type:         "1",
			}},
			ListFraudInfo: &[]entities.FraudInfo{{
				SourceSystem:  "SUSP",
				Source:        spr.Source,
				FraudAreaCode: entities.FRAUD_AREA_CODE,
				FraudTypeCode: spr.FraudTypeCode,
				FraudDegree:   spr.FraudDegree,
				DataDate:      spr.DataDate,
				RecordDate:    spr.RecordDate,
				UpdateBy:      spr.UpdateBy,
				UpdateDate:    spr.UpdateDate,
				SourceDesc:    sourceDesc,
			}},
			ListResidentInfo:     &[]entities.ResidentInfo{},
			ListAddressInfo:      &[]entities.AddressInfo{},
			ListContactInfo:      &[]entities.ContactInfo{},
			ListCollateralInfo:   &[]entities.CollateralInfo{},
			ListRelationShipInfo: &[]entities.RelationShipInfo{},
			ListBirthInfo:        &[]entities.BirthInfo{},
		},
	}

	if spr.CitizenID != "" {
		ktbSubmitBody.AppendNumber(spr, spr.CitizenID, entities.CODE_CITIZEN_ID)
	}

	if spr.PassportID != "" {
		ktbSubmitBody.AppendNumber(spr, spr.PassportID, entities.CODE_PASSPORT_ID)
	}

	if spr.JuristicID != "" {
		ktbSubmitBody.AppendNumber(spr, spr.JuristicID, entities.CODE_JURISTIC_ID) // TODO
	}

	if spr.BirthDate != "" {
		ktbSubmitBody.ListDetail.ListBirthInfo = &[]entities.BirthInfo{{
			SourceSystem: "SUSP",
			BirthDate:    spr.BirthDate,
			UpdateBy:     spr.UpdateBy,
			Type:         "1",
			EntityID:     "",
			SysID:        "",
		}}
	}

	referenceId := ctx.Locals("referenceId").(string)

	result, err := u.SuspectRepo.SubmitKtbSuspect(ctx, ktbSubmitBody, xCorrelationID, referenceId)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}

	return result, nil
}
