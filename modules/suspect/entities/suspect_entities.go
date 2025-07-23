package entities

type ListMaster struct {
	EntityID       string `json:"entityID"`
	SourceSystem   string `json:"sourceSystem"`
	EntityTP       string `json:"entityTP"`
	ListSupTPCD    string `json:"listSupTPCD"`
	CollateralFlag string `json:"collateralFlag"`
	StartDate      string `json:"startDate"`
	EndDate        string `json:"endDate"`
	VerifiedDate   string `json:"verifiedDate"`
	UpdateBy       string `json:"updateBy"`
	UpdateDate     string `json:"updateDate"`
	ActiveFlag     string `json:"activeFlag"`
	Name           string `json:"name"`
}

type AddressInfo struct {
	EntityID     string `json:"entityID"`
	SysID        string `json:"sysID"`
	SourceSystem string `json:"sourceSystem" validate:"required"`
	Type         string `json:"type" validate:"required"`
	Typedesc     string `json:"typedesc"`
	Fulladdress  string `json:"fulladdress"`
	BuildingName string `json:"buildingName"`
	HomeNo       string `json:"homeNo"`
	Moo          string `json:"moo"`
	Soi          string `json:"soi"`
	Road         string `json:"road"`
	SubDistrict  string `json:"subDistrict"`
	District     string `json:"district"`
	Province     string `json:"province"`
	PostCode     string `json:"postCode"`
	StartDate    string `json:"startDate"`
	EndDate      string `json:"endDate"`
	VerifiedDate string `json:"verifiedDate"`
	UpdateBy     string `json:"updateBy" validate:"required"`
	UpdateDate   string `json:"updateDate"`
	ActiveFlag   string `json:"activeFlag"`
}

type BirthInfo struct {
	EntityID     string `json:"entityID"`
	SysID        string `json:"sysID"`
	SourceSystem string `json:"sourceSystem" validate:"required"`
	Type         string `json:"type" validate:"required"`
	BirthDate    string `json:"birthDate"`    // required_if
	BirthCountry string `json:"birthCountry"` // required_if
	StartDate    string `json:"startDate"`
	EndDate      string `json:"endDate"`
	VerifiedDate string `json:"verifiedDate"`
	UpdateBy     string `json:"updateBy" validate:"required"`
	UpdateDate   string `json:"updateDate"`
	ActiveFlag   string `json:"activeFlag"`
}

type IdentificationInfo struct {
	EntityID     string `json:"entityID"`
	SysId        string `json:"sysId"`
	SourceSystem string `json:"sourceSystem"`
	Type         string `json:"type"`
	Number       string `json:"number"`
	IdCountry    string `json:"IdCountry"`
	IdTypeDesc   string `json:"IdTypeDesc"`
	StartDate    string `json:"startDate"`
	EndDate      string `json:"endDate"`
	VerifiedDate string `json:"verifiedDate"`
	UpdateBy     string `json:"updateBy"`
	UpdateDate   string `json:"updateDate"`
	ActiveFlag   string `json:"activeFlag"`
}

type NameInfo struct {
	EntityID     string `json:"entityID"`
	SysId        string `json:"sysId"`
	SourceSystem string `json:"sourceSystem"`
	Type         string `json:"type"`
	Title        string `json:"title"`
	TitleDesc    string `json:"titleDesc"`
	FirstName    string `json:"firstname"`
	MiddleName   string `json:"middlename"`
	LastName     string `json:"lastname"`
	FullName     string `json:"fullname"`
	OrgName      string `json:"orgname"`
	StartDate    string `json:"startDate"`
	EndDate      string `json:"endDate"`
	VerifiedDate string `json:"verifiedDate"`
	UpdateBy     string `json:"updateBy"`
	UpdateDate   string `json:"updateDate"`
	ActiveFlag   string `json:"activeFlag"`
}

type ResidentInfo struct{}

type FraudInfo struct {
	EntityID        string `json:"entityID"`
	SysId           string `json:"sysId"`
	SourceSystem    string `json:"sourceSystem"`
	Source          string `json:"source"`
	SourceDesc      string `json:"sourceDesc"`
	SourceKtbDept   string `json:"sourceKtbDept"`
	FraudAreaCode   string `json:"fraudAreaCode"`
	FraudAreaDesc   string `json:"fraudAreaDesc"`
	FraudTypeCode   string `json:"fraudTypeCode"`
	FraudTypeDesc   string `json:"fraudTypeDesc"`
	FraudDegree     string `json:"fraudDegree"`
	FraudCategory   string `json:"fraudCategory"`
	ResultCheck     string `json:"resultCheck"`
	ResultCheckDesc string `json:"resultCheckDesc"`
	FromBcCbc       string `json:"fromBC_CBC"`
	Remark          string `json:"remark"`
	DataDate        string `json:"dataDate"`
	RecordDate      string `json:"recordDate"`
	StartDate       string `json:"startDate"`
	EndDate         string `json:"endDate"`
	VerifiedDate    string `json:"verifiedDate"`
	UpdateBy        string `json:"updateBy"`
	UpdateDate      string `json:"updateDate"`
	ActiveFlag      string `json:"activeFlag"`
}

type ContactInfo struct{}

type CollateralInfo struct{}

type RelationShipInfo struct{}

type ListDetail struct {
	EntityID               string                `json:"entityID"`
	ListMaster             *ListMaster           `json:"listMaster" validate:"required"`
	ListAddressInfo        *[]AddressInfo        `json:"listAddressInfo"`
	ListBirthInfo          *[]BirthInfo          `json:"listBirthInfo"`
	ListIdentificationInfo *[]IdentificationInfo `json:"listIdentificationInfo"`
	ListNameInfo           *[]NameInfo           `json:"listNameInfo"`
	ListResidentInfo       *[]ResidentInfo       `json:"listResidentInfo"`
	ListFraudInfo          *[]FraudInfo          `json:"listFraudInfo"`
	ListContactInfo        *[]ContactInfo        `json:"listContactInfo"`
	ListCollateralInfo     *[]CollateralInfo     `json:"listCollateralInfo"`
	ListRelationShipInfo   *[]RelationShipInfo   `json:"listRelationShipInfo"`
}

type SuspectControl struct {
	RequestID             string `json:"requestId"`
	RequesterName         string `json:"requesterName"`
	RequesterLanguage     string `json:"requesterLanguage"`
	RequesterLocale       string `json:"requesterLocale"`
	PageStartIndex        uint16 `json:"pageStartIndex"`
	PageEndIndex          uint16 `json:"pageEndIndex"`
	AvailableResultsCount uint16 `json:"availableResultsCount"`
}

type KtbIndividualSubmitSuspectRequest struct {
	Control    *SuspectControl `json:"control,omitempty"`
	ListDetail *ListDetail     `json:"listDetail"`
}

type SuspectResponse struct {
	Control    *SuspectControl `json:"control,omitempty"`
	Status     string          `json:"status" example:"SUCCESS"`
	StatusDesc string          `json:"statusDesc" example:"Successful"`
	Errors     string          `json:"errors,omitempty" example:""`
}

type IngSuspectRequest struct {
	EntityTP      string `json:"entityTP" validate:"required,oneof=PERSON ENTITY"`
	FirstName     string `json:"first_name,omitempty" validate:"omitempty,isName=EntityTP"`
	LastName      string `json:"last_name,omitempty" validate:"omitempty,isName=EntityTP"`
	CompanyName   string `json:"company_name,omitempty" validate:"isCompanyName=EntityTP"`
	BirthDate     string `json:"birth_date,omitempty" validate:"omitempty,dateformat"`
	CitizenID     string `json:"citizen_id,omitempty" validate:"omitempty,isCitizenID=EntityTP"`
	JuristicID    string `json:"juristic_id,omitempty" validate:"omitempty,isJuristicID=EntityTP"`
	PassportID    string `json:"passport_id,omitempty" validate:"omitempty,noSpecialChar"`
	FraudTypeCode string `json:"fraud_type_code" validate:"required,isNegative"`
	FraudDegree   string `json:"fraud_degree,omitempty" validate:"omitempty,oneof=B G"`
	DataDate      string `json:"data_date,omitempty" validate:"omitempty,dateformat"`
	RecordDate    string `json:"record_date,omitempty" validate:"omitempty,dateformat"`
	UpdateBy      string `json:"update_by,omitempty" validate:"omitempty,noSpecialChar"`
	UpdateDate    string `json:"update_date" validate:"required,dateformat"`
	Source        string `json:"source" validate:"required,isNegative"`
	SourceDesc    string `json:"source_des,omitempty"`
	Remark        string `json:"remark,omitempty"`
	ActiveFlag    string `json:"active_flag,omitempty"`
}

func (ktb *KtbIndividualSubmitSuspectRequest) AppendNumber(spr *IngSuspectRequest, number string, codeType string) {
	*ktb.ListDetail.ListIdentificationInfo = append(*ktb.ListDetail.ListIdentificationInfo, IdentificationInfo{
		SourceSystem: "SUSP",
		Type:         codeType,
		UpdateBy:     spr.UpdateBy,
		Number:       number,
		UpdateDate:   spr.UpdateDate, // TODO
		ActiveFlag:   spr.ActiveFlag, // TODO
	})
}
