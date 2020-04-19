package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"log"

	"github.com/AkezhanOb1/diplomaProject/api/graphQL/graph/generated"
	"github.com/AkezhanOb1/diplomaProject/api/graphQL/graph/model"
	"github.com/AkezhanOb1/diplomaProject/pkg"
	c "github.com/AkezhanOb1/diplomaProject/services/client/business/category"
	bc "github.com/AkezhanOb1/diplomaProject/services/client/business/company"
	cs "github.com/AkezhanOb1/diplomaProject/services/client/business/companyService"
	bo "github.com/AkezhanOb1/diplomaProject/services/client/business/owner"
	bs "github.com/AkezhanOb1/diplomaProject/services/client/business/service"
	sc "github.com/AkezhanOb1/diplomaProject/services/client/business/subCategories"
)

func (r *mutationResolver) CreateBusinessCompany(ctx context.Context, input model.CreateBusinessCompanyRequest) (*model.BusinessCompany, error) {
	newBusinessCompany, err := bc.CreateBusinessCompany(ctx, input)
	if err != nil {
		return nil, err
	}

	var resp = &model.BusinessCompany{
		BusinessCompanyID:         newBusinessCompany.BusinessCompany.BusinessCompanyID,
		BusinessCompanyName:       newBusinessCompany.BusinessCompany.BusinessCompanyName,
		BusinessCompanyCategoryID: newBusinessCompany.BusinessCompany.BusinessCompanyCategoryID,
	}

	return resp, nil
}

func (r *mutationResolver) CreateBusinessOwner(ctx context.Context, input model.CreateBusinessOwnerRequest) (*model.BusinessOwner, error) {
	newBusinessOwner, err := bo.CreateBusinessOwner(ctx, input)
	if err != nil {
		return nil, err
	}

	var resp = &model.BusinessOwner{
		BusinessOwnerID:                newBusinessOwner.BusinessOwner.BusinessOwnerID,
		BusinessOwnerName:              input.BusinessOwnerName,
		BusinessOwnerEmail:             input.BusinessOwnerEmail,
		BusinessOwnerPhoneNumberPrefix: input.BusinessOwnerPhoneNumberPrefix,
		BusinessOwnerPhoneNumber:       input.BusinessOwnerPhoneNumber,
	}

	return resp, nil
}

func (r *mutationResolver) CreateBusinessService(ctx context.Context, input model.CreateBusinessServiceRequest) (*model.CreateBusinessServiceResponse, error) {
	newBusinessService, err := bs.CreateBusinessService(ctx, input.BusinessServiceName, input.BusinessServiceSubCategories)
	if err != nil {
		return nil, err
	}

	var resp = &model.CreateBusinessServiceResponse{
		BusinessService: &model.BusinessService{
			BusinessServiceID:   newBusinessService.BusinessService.GetBusinessServiceID(),
			BusinessServiceName: newBusinessService.BusinessService.GetBusinessServiceName(),
			SubCategories:       newBusinessService.BusinessService.GetSubCategories(),
		},
	}

	return resp, nil
}

func (r *mutationResolver) CreateCompanyService(ctx context.Context, input model.CreateCompanyServiceRequest) (*model.CreateCompanyServiceResponse, error) {
	newCompanyService, err := cs.CreateBusinessService(ctx, input)
	if err != nil {
		return nil, err
	}

	var service = newCompanyService.GetCompanyService()
	var resp = &model.CreateCompanyServiceResponse{
		CompanyService: &model.CompanyService{
			CompanyServiceID:       service.GetCompanyServiceID(),
			CompanyServiceName:     service.GetCompanyServiceName(),
			CompanyServiceDuration: service.GetCompanyServiceDuration(),
			CompanyServicePrice:    service.GetCompanyServicePrice(),
			BusinessServiceID:      &service.BusinessServiceID,
			BusinessCompanyID:      &service.BusinessCompanyID,
		},
	}

	return resp, nil
}

func (r *mutationResolver) UpdateCompanyService(ctx context.Context, input model.UpdateCompanyServiceRequest) (*model.UpdateCompanyServiceResponse, error) {
	updatedService, err := cs.UpdateCompanyService(ctx, input)
	if err != nil {
		return nil, err
	}

	var service = updatedService.GetCompanyService()
	var resp = &model.UpdateCompanyServiceResponse{
		CompanyService: &model.CompanyService{
			CompanyServiceID:       service.GetCompanyServiceID(),
			CompanyServiceName:     service.GetCompanyServiceName(),
			CompanyServiceDuration: service.GetCompanyServiceDuration(),
			CompanyServicePrice:    service.GetCompanyServicePrice(),
			BusinessServiceID:      &service.BusinessServiceID,
			BusinessCompanyID:      &service.BusinessCompanyID,
		},
	}

	return resp, nil
}

func (r *mutationResolver) DeleteCompanyService(ctx context.Context, input model.DeleteCompanyServiceRequest) (*model.DeleteCompanyServiceResponse, error) {
	deletedService, err := cs.DeleteCompanyService(ctx, input)
	if err != nil {
		return nil, err
	}

	var service = deletedService.GetCompanyService()
	var resp = &model.DeleteCompanyServiceResponse{
		CompanyService: &model.CompanyService{
			CompanyServiceID:       service.GetCompanyServiceID(),
			CompanyServiceName:     service.GetCompanyServiceName(),
			CompanyServiceDuration: service.GetCompanyServiceDuration(),
			CompanyServicePrice:    service.GetCompanyServicePrice(),
			BusinessServiceID:      &service.BusinessServiceID,
			BusinessCompanyID:      &service.BusinessCompanyID,
		},
	}

	return resp, nil
}

func (r *mutationResolver) CreateBusinessCompanyOperationHours(ctx context.Context, input model.CreateBusinessCompanyOperationHoursRequest) (*model.CreateBusinessCompanyOperationHoursResponse, error) {
	newOperationHours, err := bc.CreateBusinessCompanyOperationHour(ctx, input)
	if err != nil {
		return nil, err
	}

	var operationHours = newOperationHours.GetBusinessCompanyOperationHour()
	var resp = &model.CreateBusinessCompanyOperationHoursResponse{
		BusinessCompanyOperationHour: &model.BusinessCompanyOperationHour{
			CompanyOperationHourID: operationHours.GetCompanyOperationHourID(),
			BusinessCompanyID:      operationHours.GetBusinessCompanyID(),
			DayOfWeek:              operationHours.GetDayOfWeek(),
			OpenTime:               operationHours.GetOpenTime(),
			CloseTime:              operationHours.GetCloseTime(),
		},
	}

	return resp, nil
}

func (r *mutationResolver) UpdateBusinessCompanyOperationHours(ctx context.Context, input model.UpdateBusinessCompanyOperationHoursRequest) (*model.UpdateBusinessCompanyOperationHoursResponse, error) {
	updatesOperationHour, err := bc.UpdateBusinessCompanyOperationHour(ctx, input)
	if err != nil {
		return nil, err
	}

	var operationHours = updatesOperationHour.GetBusinessCompanyOperationHour()
	var resp = &model.UpdateBusinessCompanyOperationHoursResponse{
		BusinessCompanyOperationHour: &model.BusinessCompanyOperationHour{
			CompanyOperationHourID: operationHours.GetCompanyOperationHourID(),
			BusinessCompanyID:      operationHours.GetBusinessCompanyID(),
			DayOfWeek:              operationHours.GetDayOfWeek(),
			OpenTime:               operationHours.GetOpenTime(),
			CloseTime:              operationHours.GetCloseTime(),
		},
	}

	return resp, nil
}

func (r *mutationResolver) DeleteBusinessCompanyOperationHours(ctx context.Context, input model.DeleteBusinessCompanyOperationHoursRequest) (*model.DeleteBusinessCompanyOperationHoursResponse, error) {
	deletedOperationHour, err := bc.DeleteBusinessCompanyOperationHour(ctx, input)
	if err != nil {
		return nil, err
	}

	var operationHours = deletedOperationHour.GetBusinessCompanyOperationHour()
	var resp = &model.DeleteBusinessCompanyOperationHoursResponse{
		BusinessCompanyOperationHour: &model.BusinessCompanyOperationHour{
			CompanyOperationHourID: operationHours.GetCompanyOperationHourID(),
			BusinessCompanyID:      operationHours.GetBusinessCompanyID(),
			DayOfWeek:              operationHours.GetDayOfWeek(),
			OpenTime:               operationHours.GetOpenTime(),
			CloseTime:              operationHours.GetCloseTime(),
		},
	}

	return resp, nil
}

func (r *queryResolver) GetBusinessCompany(ctx context.Context, input model.GetBusinessCompanyRequest) (*model.BusinessCompany, error) {
	company, err := bc.GetBusinessCompany(ctx, input.BusinessCompanyID)
	if err != nil {
		return nil, err
	}
	var resp = &model.BusinessCompany{
		BusinessCompanyID:         company.BusinessCompany.BusinessCompanyID,
		BusinessCompanyName:       company.BusinessCompany.BusinessCompanyName,
		BusinessCompanyCategoryID: company.BusinessCompany.BusinessCompanyCategoryID,
	}

	return resp, nil
}

func (r *queryResolver) GetBusinessCompanies(ctx context.Context) (*model.BusinessCompanies, error) {
	companies, err := bc.GetBusinessCompanies(ctx)
	if err != nil {
		return nil, err
	}

	b, err := pkg.Serializer(companies)
	if err != nil {
		return nil, err
	}

	var resp model.BusinessCompanies
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *queryResolver) GetBusinessCompanyServices(ctx context.Context, input *model.GetBusinessCompanyServicesRequest) (*model.GetBusinessCompanyServicesResponse, error) {
	services, err := bc.GetBusinessCompanyServices(ctx, input.BusinessCompanyID)
	if err != nil {
		return nil, err
	}

	b, err := pkg.Serializer(services)
	if err != nil {
		return nil, err
	}

	var resp model.GetBusinessCompanyServicesResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *queryResolver) GetBusinessCompanyOperationHourByDay(ctx context.Context, input *model.GetGetBusinessCompanyOperationHourByDayRequest) (*model.BusinessCompanyOperationHourResponse, error) {
	operationHours, err := bc.GetBusinessCompanyOperationHourByDay(ctx, input.BusinessCompanyID, input.DayOfWeek)
	if err != nil {
		return nil, err
	}

	b, err := pkg.Serializer(operationHours)
	if err != nil {
		return nil, err
	}

	var resp model.BusinessCompanyOperationHourResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *queryResolver) GetBusinessCompanyOperationHours(ctx context.Context, input *model.GetBusinessCompanyOperationHoursRequest) (*model.BusinessCompanyOperationHours, error) {
	operationHours, err := bc.GetBusinessCompanyOperationHours(ctx, input.BusinessCompanyID)
	if err != nil {
		return nil, err
	}

	b, err := pkg.Serializer(operationHours)
	if err != nil {
		return nil, err
	}

	var resp model.BusinessCompanyOperationHours
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *queryResolver) GetBusinessCategory(ctx context.Context, input model.BusinessCategoryRequest) (*model.BusinessCategory, error) {
	category, err := c.GetBusinessCategory(ctx, input.BusinessCategoryID)
	if err != nil {
		return nil, err
	}
	var resp = &model.BusinessCategory{
		BusinessCategoryID:   category.BusinessCategory.BusinessCategoryID,
		BusinessCategoryName: category.BusinessCategory.BusinessCategoryName,
	}

	return resp, nil
}

func (r *queryResolver) GetBusinessCategories(ctx context.Context) ([]model.BusinessCategory, error) {
	categories, err := c.GetBusinessCategories(ctx)
	if err != nil {
		return nil, err
	}

	b, err := pkg.Serializer(categories)
	if err != nil {
		return nil, err
	}

	var resp []model.BusinessCategory
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *queryResolver) GetBusinessSubCategory(ctx context.Context, input model.BusinessSubCategoryRequest) (*model.BusinessSubCategory, error) {
	subCategory, err := sc.GetBusinessSubCategory(ctx, input.BusinessSubCategoryID)
	if err != nil {
		return nil, err
	}
	var resp = &model.BusinessSubCategory{
		BusinessSubCategoryID:   subCategory.BusinessSubCategory.GetBusinessCategoryID(),
		BusinessSubCategoryName: subCategory.BusinessSubCategory.GetBusinessSubCategoryName(),
		BusinessCategoryID:      subCategory.BusinessSubCategory.GetBusinessCategoryID(),
	}

	return resp, nil
}

func (r *queryResolver) GetBusinessSubCategories(ctx context.Context) (*model.BusinessSubCategories, error) {
	subCategories, err := sc.GetBusinessSubCategories(ctx)
	if err != nil {
		return nil, err
	}

	b, err := pkg.Serializer(subCategories)
	if err != nil {
		return nil, err
	}

	log.Println(string(b))
	var resp model.BusinessSubCategories
	err = json.Unmarshal(b, &resp)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *queryResolver) GetBusinessSubCategoriesUnderCategory(ctx context.Context, input *model.BusinessSubCategoriesUnderCategoryRequest) (*model.BusinessSubCategories, error) {
	subCategories, err := sc.GetBusinessSubCategoriesUnderCategory(ctx, input.BusinessCategoryID)
	if err != nil {
		return nil, err
	}

	b, err := pkg.Serializer(subCategories)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))

	var resp model.BusinessSubCategories
	err = json.Unmarshal(b, &resp)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *queryResolver) GetBusinessService(ctx context.Context, input model.GetBusinessServiceRequest) (*model.BusinessService, error) {
	service, err := bs.GetBusinessService(ctx, input.BusinessServiceID)
	if err != nil {
		return nil, err
	}
	var resp = &model.BusinessService{
		BusinessServiceID:   service.BusinessService.GetBusinessServiceID(),
		BusinessServiceName: service.BusinessService.GetBusinessServiceName(),
		SubCategories:       service.BusinessService.GetSubCategories(),
	}

	return resp, nil
}

func (r *queryResolver) GetBusinessServices(ctx context.Context) (*model.BusinessServices, error) {
	services, err := bs.GetBusinessServices(ctx)
	if err != nil {
		return nil, err
	}

	b, err := pkg.Serializer(services)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))

	var resp model.BusinessServices
	err = json.Unmarshal(b, &resp)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *queryResolver) GetBusinessServicesUnderSubCategory(ctx context.Context, input *model.GetBusinessServicesUnderSubCategoryRequest) (*model.BusinessServices, error) {
	services, err := bs.GetServicesUnderSubCategory(ctx, input.SubCategoryID)
	if err != nil {
		return nil, err
	}

	b, err := pkg.Serializer(services)
	if err != nil {
		return nil, err
	}
	log.Println(string(b))

	var resp model.BusinessServices
	err = json.Unmarshal(b, &resp)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *queryResolver) GetCompanyService(ctx context.Context, input model.GetCompanyServiceRequest) (*model.CompanyService, error) {
	service, err := cs.GetCompanyService(ctx, input.CompanyServiceID)
	if err != nil {
		return nil, err
	}
	var resp = &model.CompanyService{
		CompanyServiceID:       service.CompanyService.CompanyServiceID,
		CompanyServiceName:     service.CompanyService.CompanyServiceName,
		CompanyServiceDuration: service.CompanyService.CompanyServiceDuration,
		CompanyServicePrice:    service.CompanyService.CompanyServicePrice,
		BusinessServiceID:      &service.CompanyService.BusinessCompanyID,
		BusinessServiceName:    &service.CompanyService.BusinessServiceName,
		BusinessCompanyID:      &service.CompanyService.BusinessCompanyID,
		BusinessCompanyName:    &service.CompanyService.BusinessCompanyName,
	}

	return resp, nil
}

func (r *queryResolver) GetCompanyServices(ctx context.Context) (*model.CompanyServices, error) {
	services, err := cs.GetCompanyServices(ctx)
	if err != nil {
		return nil, err
	}

	b, err := pkg.Serializer(services)
	if err != nil {
		return nil, err
	}

	var resp model.CompanyServices
	err = json.Unmarshal(b, &resp)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *queryResolver) GetCompanyServicesUnderSubCategory(ctx context.Context, input model.GetCompanyServicesUnderSubCategoryRequest) (*model.CompanyServices, error) {
	services, err := cs.GetCompanyServicesUnderSubCategory(ctx, input.SubCategoryID)
	if err != nil {
		return nil, err
	}

	b, err := pkg.Serializer(services)
	if err != nil {
		return nil, err
	}

	var resp model.CompanyServices
	err = json.Unmarshal(b, &resp)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
