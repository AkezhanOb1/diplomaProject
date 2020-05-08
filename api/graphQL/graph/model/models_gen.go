// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type BusinessCategory struct {
	BusinessCategoryID   int64  `json:"businessCategoryID"`
	BusinessCategoryName string `json:"businessCategoryName"`
}

type BusinessCategoryRequest struct {
	BusinessCategoryID int64 `json:"businessCategoryID"`
}

type BusinessCompanies struct {
	BusinessCompanies []BusinessCompany `json:"businessCompanies"`
}

type BusinessCompany struct {
	BusinessCompanyID         int64  `json:"businessCompanyID"`
	BusinessCompanyName       string `json:"businessCompanyName"`
	BusinessCompanyCategoryID int64  `json:"businessCompanyCategoryID"`
}

type BusinessCompanyOperationHour struct {
	CompanyOperationHourID int64  `json:"companyOperationHourID"`
	BusinessCompanyID      int64  `json:"businessCompanyID"`
	DayOfWeek              int64  `json:"dayOfWeek"`
	OpenTime               string `json:"openTime"`
	CloseTime              string `json:"closeTime"`
}

type BusinessCompanyOperationHours struct {
	BusinessCompanyOperationHour []BusinessCompanyOperationHour `json:"businessCompanyOperationHour"`
}

type BusinessCompanyServiceOperationHour struct {
	ServiceOperationHourID int64  `json:"serviceOperationHourID"`
	BusinessCompanyID      int64  `json:"businessCompanyID"`
	BusinessServiceID      int64  `json:"businessServiceID"`
	DayOfWeek              int64  `json:"dayOfWeek"`
	OpenTime               string `json:"openTime"`
	CloseTime              string `json:"closeTime"`
}

type BusinessCompanyServiceOperationHours struct {
	BusinessCompanyServiceOperationHour []BusinessCompanyServiceOperationHour `json:"businessCompanyServiceOperationHour"`
}

type BusinessOwner struct {
	BusinessOwnerID                int64  `json:"businessOwnerID"`
	BusinessOwnerName              string `json:"businessOwnerName"`
	BusinessOwnerEmail             string `json:"businessOwnerEmail"`
	BusinessOwnerPhoneNumberPrefix string `json:"businessOwnerPhoneNumberPrefix"`
	BusinessOwnerPhoneNumber       string `json:"businessOwnerPhoneNumber"`
}

type BusinessService struct {
	BusinessServiceID   int64   `json:"businessServiceID"`
	BusinessServiceName string  `json:"businessServiceName"`
	SubCategories       []int64 `json:"subCategories"`
}

type BusinessServiceOrder struct {
	BusinessServiceOrderID  int64  `json:"businessServiceOrderID"`
	ClientID                int64  `json:"clientID"`
	BusinessServiceID       int64  `json:"businessServiceID"`
	OrderDate               string `json:"orderDate"`
	CreatedAt               string `json:"createdAt"`
	PrePaid                 bool   `json:"prePaid"`
	ClientFirstName         string `json:"clientFirstName"`
	ClientPhoneNumber       string `json:"clientPhoneNumber"`
	ClientPhoneNumberPrefix string `json:"clientPhoneNumberPrefix"`
	ClientCommentary        string `json:"clientCommentary"`
}

type BusinessServices struct {
	BusinessServices []BusinessService `json:"businessServices"`
}

type BusinessSubCategories struct {
	BusinessSubCategories []BusinessSubCategory `json:"businessSubCategories"`
}

type BusinessSubCategoriesUnderCategoryRequest struct {
	BusinessCategoryID int64 `json:"businessCategoryID"`
}

type BusinessSubCategory struct {
	BusinessSubCategoryID   int64  `json:"businessSubCategoryID"`
	BusinessSubCategoryName string `json:"businessSubCategoryName"`
	BusinessCategoryID      int64  `json:"businessCategoryID"`
}

type BusinessSubCategoryRequest struct {
	BusinessSubCategoryID int64 `json:"businessSubCategoryID"`
}

type CompanyService struct {
	CompanyServiceID       int64   `json:"companyServiceID"`
	CompanyServiceName     string  `json:"companyServiceName"`
	CompanyServiceDuration int64   `json:"companyServiceDuration"`
	CompanyServicePrice    float64 `json:"companyServicePrice"`
	BusinessServiceID      *int64  `json:"businessServiceID"`
	BusinessServiceName    *string `json:"businessServiceName"`
	BusinessCompanyID      *int64  `json:"businessCompanyID"`
	BusinessCompanyName    *string `json:"businessCompanyName"`
}

type CompanyServices struct {
	CompanyServices []CompanyService `json:"companyServices"`
}

type CreateBusinessCompanyRequest struct {
	BusinessCompanyName       string `json:"businessCompanyName"`
	BusinessCompanyCategoryID int64  `json:"businessCompanyCategoryID"`
}

type CreateBusinessOwnerRequest struct {
	BusinessCompanyID              int64  `json:"businessCompanyID"`
	BusinessOwnerName              string `json:"businessOwnerName"`
	BusinessOwnerEmail             string `json:"businessOwnerEmail"`
	BusinessOwnerPassword          string `json:"businessOwnerPassword"`
	BusinessOwnerPhoneNumberPrefix string `json:"businessOwnerPhoneNumberPrefix"`
	BusinessOwnerPhoneNumber       string `json:"businessOwnerPhoneNumber"`
}

type CreateBusinessOwnerResponse struct {
	BusinessOwner *BusinessOwner `json:"businessOwner"`
	Token         *Token         `json:"token"`
}

type CreateBusinessServiceOrderRequest struct {
	ClientID                int64  `json:"clientID"`
	BusinessServiceID       int64  `json:"businessServiceID"`
	OrderDate               string `json:"orderDate"`
	PrePaid                 bool   `json:"prePaid"`
	ClientFirstName         string `json:"clientFirstName"`
	ClientPhoneNumber       string `json:"clientPhoneNumber"`
	ClientPhoneNumberPrefix string `json:"clientPhoneNumberPrefix"`
	ClientCommentary        string `json:"clientCommentary"`
}

type CreateBusinessServiceOrderResponse struct {
	BusinessServiceOrder *BusinessServiceOrder `json:"businessServiceOrder"`
}

type GetBusinessServiceOrderRequest struct {
	BusinessServiceOrderID int64 `json:"businessServiceOrderID"`
}

type GetBusinessServiceOrderResponse struct {
	BusinessServiceOrder *BusinessServiceOrder `json:"businessServiceOrder"`
}

type GetBusinessServiceOrdersResponse struct {
	BusinessServicesOrders []BusinessServiceOrder `json:"businessServicesOrders"`
}

type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int64  `json:"expiresIn"`
	TokenType    string `json:"tokenType"`
}

type BusinessCompanyOperationHourResponse struct {
	BusinessCompanyOperationHour *BusinessCompanyOperationHour `json:"businessCompanyOperationHour"`
}

type BusinessCompanyService struct {
	CompanyServiceID       int64   `json:"companyServiceID"`
	CompanyServiceName     string  `json:"companyServiceName"`
	CompanyServiceDuration int64   `json:"companyServiceDuration"`
	CompanyServicePrice    float64 `json:"companyServicePrice"`
}

type BusinessCompanyServiceOperationHourResponse struct {
	BusinessCompanyServiceOperationHour *BusinessCompanyServiceOperationHour `json:"businessCompanyServiceOperationHour"`
}

type CreateBusinessCompanyOperationHoursRequest struct {
	BusinessCompanyID int64  `json:"businessCompanyID"`
	DayOfWeek         int64  `json:"dayOfWeek"`
	OpenTime          string `json:"openTime"`
	CloseTime         string `json:"closeTime"`
}

type CreateBusinessCompanyOperationHoursResponse struct {
	BusinessCompanyOperationHour *BusinessCompanyOperationHour `json:"businessCompanyOperationHour"`
}

type CreateBusinessCompanyServiceOperationHoursRequest struct {
	BusinessCompanyID int64  `json:"businessCompanyID"`
	BusinessServiceID int64  `json:"businessServiceID"`
	DayOfWeek         int64  `json:"dayOfWeek"`
	OpenTime          string `json:"openTime"`
	CloseTime         string `json:"closeTime"`
}

type CreateBusinessCompanyServiceOperationHoursResponse struct {
	BusinessCompanyServiceOperationHour *BusinessCompanyServiceOperationHour `json:"businessCompanyServiceOperationHour"`
}

type CreateBusinessServiceRequest struct {
	BusinessServiceName          string  `json:"businessServiceName"`
	BusinessServiceSubCategories []int64 `json:"businessServiceSubCategories"`
}

type CreateBusinessServiceResponse struct {
	BusinessService *BusinessService `json:"businessService"`
}

type CreateCompanyServiceRequest struct {
	CompanyServiceName     string  `json:"companyServiceName"`
	CompanyServiceDuration int64   `json:"companyServiceDuration"`
	CompanyServicePrice    float64 `json:"companyServicePrice"`
	BusinessServiceID      int64   `json:"businessServiceID"`
	BusinessCompanyID      int64   `json:"businessCompanyID"`
}

type CreateCompanyServiceResponse struct {
	CompanyService *CompanyService `json:"companyService"`
}

type DeleteBusinessCompanyOperationHoursRequest struct {
	CompanyOperationHourID int64 `json:"companyOperationHourID"`
}

type DeleteBusinessCompanyOperationHoursResponse struct {
	BusinessCompanyOperationHour *BusinessCompanyOperationHour `json:"businessCompanyOperationHour"`
}

type DeleteBusinessCompanyServiceOperationHoursRequest struct {
	OperationHourID int64 `json:"operationHourID"`
}

type DeleteBusinessCompanyServiceOperationHoursResponse struct {
	BusinessCompanyServiceOperationHour *BusinessCompanyServiceOperationHour `json:"businessCompanyServiceOperationHour"`
}

type DeleteCompanyServiceRequest struct {
	CompanyServiceID int64 `json:"companyServiceID"`
}

type DeleteCompanyServiceResponse struct {
	CompanyService *CompanyService `json:"companyService"`
}

type GenerateTokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GenerateTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int64  `json:"expiresIn"`
	TokenType    string `json:"tokenType"`
}

type GetBusinessCompanyOperationHoursRequest struct {
	BusinessCompanyID int64 `json:"businessCompanyID"`
}

type GetBusinessCompanyRequest struct {
	BusinessCompanyID int64 `json:"businessCompanyID"`
}

type GetBusinessCompanyServiceOperationHoursRequest struct {
	ServiceID int64 `json:"serviceID"`
}

type GetBusinessCompanyServicesRequest struct {
	BusinessCompanyID int64 `json:"businessCompanyID"`
}

type GetBusinessCompanyServicesResponse struct {
	BusinessCompanyService []BusinessCompanyService `json:"businessCompanyService"`
}

type GetBusinessOwnerCompaniesRequest struct {
	Email string `json:"email"`
}

type GetBusinessOwnerCompaniesResponse struct {
	Companies []BusinessCompany `json:"companies"`
}

type GetBusinessServiceRequest struct {
	BusinessServiceID int64 `json:"businessServiceID"`
}

type GetBusinessServicesUnderSubCategoryRequest struct {
	SubCategoryID int64 `json:"subCategoryID"`
}

type GetCompanyServiceRequest struct {
	CompanyServiceID int64 `json:"companyServiceID"`
}

type GetCompanyServicesUnderSubCategoryRequest struct {
	SubCategoryID int64 `json:"subCategoryID"`
}

type GetGetBusinessCompanyOperationHourByDayRequest struct {
	BusinessCompanyID int64 `json:"businessCompanyID"`
	DayOfWeek         int64 `json:"dayOfWeek"`
}

type GetGetBusinessCompanyServiceOperationHourByDayRequest struct {
	ServiceID int64 `json:"serviceID"`
	DayOfWeek int64 `json:"dayOfWeek"`
}

type RetrieveTokenInfoRequst struct {
	AccessToken string `json:"accessToken"`
}

type RetrieveTokenInfoResponse struct {
	Email     string `json:"email"`
	ExpiresAt int64  `json:"expiresAt"`
	IssuedAt  int64  `json:"issuedAt"`
}

type UpdateBusinessCompanyOperationHoursRequest struct {
	CompanyOperationHourID int64  `json:"companyOperationHourID"`
	BusinessCompanyID      int64  `json:"businessCompanyID"`
	DayOfWeek              int64  `json:"dayOfWeek"`
	OpenTime               string `json:"openTime"`
	CloseTime              string `json:"closeTime"`
}

type UpdateBusinessCompanyOperationHoursResponse struct {
	BusinessCompanyOperationHour *BusinessCompanyOperationHour `json:"businessCompanyOperationHour"`
}

type UpdateBusinessCompanyServiceOperationHoursRequest struct {
	OperationHourID   int64  `json:"operationHourID"`
	BusinessCompanyID int64  `json:"businessCompanyID"`
	BusinessServiceID int64  `json:"businessServiceID"`
	DayOfWeek         int64  `json:"dayOfWeek"`
	OpenTime          string `json:"openTime"`
	CloseTime         string `json:"closeTime"`
}

type UpdateBusinessCompanyServiceOperationHoursResponse struct {
	BusinessCompanyServiceOperationHour *BusinessCompanyServiceOperationHour `json:"businessCompanyServiceOperationHour"`
}

type UpdateCompanyServiceRequest struct {
	CompanyServiceID       int64   `json:"companyServiceID"`
	CompanyServiceName     string  `json:"companyServiceName"`
	CompanyServiceDuration int64   `json:"companyServiceDuration"`
	CompanyServicePrice    float64 `json:"companyServicePrice"`
	BusinessServiceID      int64   `json:"businessServiceID"`
	BusinessCompanyID      int64   `json:"businessCompanyID"`
}

type UpdateCompanyServiceResponse struct {
	CompanyService *CompanyService `json:"companyService"`
}
