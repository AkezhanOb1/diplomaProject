package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	config "github.com/AkezhanOb1/diplomaProject/configs"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/AkezhanOb1/diplomaProject/api/graphQL/graph/generated"
	"github.com/AkezhanOb1/diplomaProject/api/graphQL/graph/model"
	pba "github.com/AkezhanOb1/diplomaProject/api/proto/auth"
	pbc "github.com/AkezhanOb1/diplomaProject/api/proto/customer"
	"github.com/AkezhanOb1/diplomaProject/pkg"
	db "github.com/AkezhanOb1/diplomaProject/repositories/business/company"
	c "github.com/AkezhanOb1/diplomaProject/services/client/business/category"
	bc "github.com/AkezhanOb1/diplomaProject/services/client/business/company"
	cs "github.com/AkezhanOb1/diplomaProject/services/client/business/companyService"
	bo "github.com/AkezhanOb1/diplomaProject/services/client/business/owner"
	bs "github.com/AkezhanOb1/diplomaProject/services/client/business/service"
	sc "github.com/AkezhanOb1/diplomaProject/services/client/business/subCategories"
	client "github.com/AkezhanOb1/diplomaProject/services/client/customer"
	bso "github.com/AkezhanOb1/diplomaProject/services/client/orders"
	t "github.com/AkezhanOb1/diplomaProject/services/client/token"
)

func (r *mutationResolver) SingleUpload(ctx context.Context, file graphql.Upload) (bool, error) {
	home, _ := os.Getwd()
	dir := home + "/images/5"
	os.Mkdir(dir, os.ModePerm)

	log.Println("HOME", home)
	log.Println("DIR", dir)
	path := filepath.Join(dir, filepath.Base(file.Filename))
	fileInfo, _ := os.Stat(path)
	if fileInfo != nil {
		return false, fmt.Errorf("file with such name already exists")
	}

	dst, err := os.Create(path)
	if err != nil {
		return false, err
	}

	if _, err = io.Copy(dst, file.File); err != nil {
		return false, err
	}

	actualPath := config.ImagePath + "/" + strconv.FormatInt(7, 10) + "/"+ file.Filename
	_, err = db.UploadBusinessCompanyImageRepository(ctx, actualPath, 5)
	if err != nil {
		return false, err
	}


	dst.Close()

	_, err = ioutil.ReadAll(file.File)
	if err != nil {
		return false, err
	}


	return true, nil
}

func (r *mutationResolver) BusinessCompanyImageUpload(ctx context.Context, input model.BusinessCompanyImageUploadRequest) (*model.File, error) {
	home, _ := os.Getwd()
	dir := home + "/images/" + strconv.FormatInt(input.BussinessCompanyID, 10)
	os.Mkdir(dir, os.ModePerm)

	file := input.File
	path := filepath.Join(dir, filepath.Base(file.Filename))
	fileInfo, _ := os.Stat(path)
	if fileInfo != nil {
		return nil, fmt.Errorf("file with such name already exists")
	}

	dst, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	if _, err = io.Copy(dst, file.File); err != nil {
		return nil, err
	}

	actualPath := config.ImagePath + "/" + strconv.FormatInt(7, 10) + "/"+ file.Filename
	id, err := db.UploadBusinessCompanyImageRepository(ctx, actualPath, input.BussinessCompanyID)
	if err != nil {
		return nil, err
	}

	dst.Close()

	content, err := ioutil.ReadAll(file.File)
	if err != nil {
		return nil, err
	}
	var f = model.File{
		ID:      *id,
		Name:    file.Filename,
		Content: string(content),
	}

	return &f, err
}

func (r *mutationResolver) BusinessCompanyImagesUpload(ctx context.Context, input model.BusinessCompanyImagesUploadRequest) ([]model.File, error) {
	home, _ := os.Getwd()
	dir := home + "/images/" + strconv.FormatInt(input.BussinessCompanyID, 10)
	os.Mkdir(dir, os.ModePerm)


	var resp []model.File

	for _, f := range input.Files {
		file := f.File
		path := filepath.Join(dir, filepath.Base(file.Filename))
		fileInfo, _ := os.Stat(path)
		if fileInfo != nil {
			log.Println("file with such name already exists")
			continue
		}

		dst, err := os.Create(path)
		if err != nil {
			log.Println(err)
			continue
		}

		if _, err = io.Copy(dst, file.File); err != nil {
			log.Println(err)
		}

		actualPath := config.ImagePath + "/" + strconv.FormatInt(7, 10) + "/"+ file.Filename
		id, err := db.UploadBusinessCompanyImageRepository(ctx, actualPath, input.BussinessCompanyID)
		if err != nil {
			log.Println(err)
			continue
		}

		dst.Close()

		content, err := ioutil.ReadAll(file.File)
		if err != nil {
			return nil, err
		}
		var f = model.File{
			ID:      *id,
			Name:    file.Filename,
			Content: string(content),
		}

		resp = append(resp, f)
	}

	return resp, nil
}

func (r *mutationResolver) UpdateBusinessServiceOrder(ctx context.Context, input model.UpdateBusinessServiceOrderRequest) (*model.UpdateBusinessServiceOrderResponse, error) {
	order, err := bso.UpdateBusinessServiceOrder(ctx, &input)
	if err != nil {
		return nil, err
	}

	b, err := pkg.Serializer(order)
	if err != nil {
		return nil, err
	}

	var resp model.UpdateBusinessServiceOrderResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *mutationResolver) DeleteBusinessServiceOrder(ctx context.Context, input model.DeleteBusinessServiceOrderRequest) (*model.DeleteBusinessServiceOrderResponse, error) {
	order, err := bso.DeleteBusinessServiceOrder(ctx, &input)
	if err != nil {
		return nil, err
	}

	b, err := pkg.Serializer(order)
	if err != nil {
		return nil, err
	}

	var resp model.DeleteBusinessServiceOrderResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *mutationResolver) CreateCustomer(ctx context.Context, input model.CreateCustomerRequest) (*model.CreateCustomerResponse, error) {
	var req = &pbc.CreateCustomerRequest{
		CustomerFirstName:         input.CustomerFirstName,
		CustomerSecondName:        input.CustomerSecondName,
		CustomerEmail:             input.CustomerEmail,
		CustomerPhoneNumberPrefix: input.CustomerPhoneNumberPrefix,
		CustomerPhoneNumber:       input.CustomerPhoneNumber,
		CustomerPassword:          input.CustomerPassword,
	}

	customer, err := client.CreateCustomer(ctx, req)

	if err != nil {
		return nil, nil
	}

	b, err := pkg.Serializer(customer)
	if err != nil {
		return nil, err
	}

	var response *model.CreateCustomerResponse
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	var tokenRequest = pba.GenerateCustomerTokenRequest{
		Credentials: &pba.ClientCredentials{
			Email:    customer.Customer.GetCustomerEmail(),
			Password: "",
		},
	}
	token, err := client.GenerateToken(context.Background(), &tokenRequest)
	if err != nil {
		return nil, err
	}

	response.Token = &model.Token{
		AccessToken:  token.Token.GetAccessToken(),
		RefreshToken: token.Token.GetRefreshToken(),
		ExpiresIn:    token.Token.GetExpiresIn(),
		TokenType:    token.Token.GetTokenType(),
	}

	return response, nil
}

func (r *mutationResolver) CreateCustomerToken(ctx context.Context, input model.CreateCustomerTokenRequest) (*model.CreateCustomerTokenResponse, error) {
	var checkPasswordRequest = pbc.CheckCustomerPasswordRequest{
		Email:    input.Email,
		Password: input.Password,
	}
	checkResult, err := client.CheckOwnerPassword(context.Background(), &checkPasswordRequest)
	if err != nil {
		return nil, err
	}

	var tokenRequest = pba.GenerateCustomerTokenRequest{
		Credentials: &pba.ClientCredentials{
			Email:    input.Email,
			Password: "",
		},
	}

	token, err := client.GenerateToken(context.Background(), &tokenRequest)
	if err != nil {
		return nil, err
	}

	if checkResult.GetValid() != true {
		return nil, fmt.Errorf("email or password is not correct")
	}

	return &model.CreateCustomerTokenResponse{
		Token: &model.Token{
			AccessToken:  token.Token.GetAccessToken(),
			RefreshToken: token.Token.GetRefreshToken(),
			ExpiresIn:    token.Token.GetExpiresIn(),
			TokenType:    token.Token.GetTokenType(),
		},
	}, nil
}

func (r *mutationResolver) CreateBusinessServiceOrder(ctx context.Context, input model.CreateBusinessServiceOrderRequest) (*model.CreateBusinessServiceOrderResponse, error) {
	businessServiceOrder, err := bso.CreateBusinessServiceOrder(ctx, input)
	if err != nil {
		return nil, err
	}

	var order = businessServiceOrder.GetBusinessServiceOrder()
	var resp = &model.CreateBusinessServiceOrderResponse{
		BusinessServiceOrder: &model.BusinessServiceOrder{
			BusinessServiceOrderID:  order.BusinessServiceOrderID,
			ClientID:                order.ClientID,
			BusinessServiceID:       order.BusinessServiceID,
			StartAt:                 order.StartAt,
			EndAt:                   order.EndAt,
			CreatedAt:               order.CreatedAt,
			PrePaid:                 order.PrePaid,
			ClientFirstName:         order.ClientFirstName,
			ClientPhoneNumber:       order.ClientPhoneNumber,
			ClientPhoneNumberPrefix: order.ClientPhoneNumberPrefix,
			ClientCommentary:        order.ClientCommentary,
		},
	}

	return resp, nil
}

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

func (r *mutationResolver) CreateBusinessOwner(ctx context.Context, input model.CreateBusinessOwnerRequest) (*model.CreateBusinessOwnerResponse, error) {
	newBusinessOwner, err := bo.CreateBusinessOwner(ctx, input)
	if err != nil {
		return nil, err
	}

	var owner = &model.BusinessOwner{
		BusinessOwnerID:                newBusinessOwner.BusinessOwner.BusinessOwnerID,
		BusinessOwnerName:              input.BusinessOwnerName,
		BusinessOwnerEmail:             input.BusinessOwnerEmail,
		BusinessOwnerPhoneNumberPrefix: input.BusinessOwnerPhoneNumberPrefix,
		BusinessOwnerPhoneNumber:       input.BusinessOwnerPhoneNumber,
	}

	newToken, err := t.GenerateToken(input.BusinessOwnerEmail, input.BusinessOwnerPassword)
	if err != nil {
		return nil, err
	}

	var token = &model.Token{
		AccessToken:  newToken.GetAccessToken(),
		RefreshToken: newToken.GetRefreshToken(),
		ExpiresIn:    newToken.GetExpiresIn(),
		TokenType:    newToken.GetTokenType(),
	}

	return &model.CreateBusinessOwnerResponse{
		BusinessOwner: owner,
		Token:         token,
	}, nil
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

func (r *mutationResolver) CreateBusinessCompanyServiceOperationHours(ctx context.Context, input model.CreateBusinessCompanyServiceOperationHoursRequest) (*model.CreateBusinessCompanyServiceOperationHoursResponse, error) {
	newOperationHours, err := cs.CreateBusinessCompanyServiceOperationHour(ctx, input)
	if err != nil {
		return nil, err
	}

	var operationHours = newOperationHours.GetBusinessCompanyServiceOperationHour()
	var resp = &model.CreateBusinessCompanyServiceOperationHoursResponse{
		BusinessCompanyServiceOperationHour: &model.BusinessCompanyServiceOperationHour{
			ServiceOperationHourID: operationHours.GetServiceOperationHourID(),
			BusinessCompanyID:      operationHours.GetBusinessCompanyID(),
			BusinessServiceID:      operationHours.GetBusinessServiceID(),
			DayOfWeek:              operationHours.GetDayOfWeek(),
			OpenTime:               operationHours.GetOpenTime(),
			CloseTime:              operationHours.GetCloseTime(),
		},
	}

	return resp, nil
}

func (r *mutationResolver) UpdateBusinessCompanyServiceOperationHours(ctx context.Context, input model.UpdateBusinessCompanyServiceOperationHoursRequest) (*model.UpdateBusinessCompanyServiceOperationHoursResponse, error) {
	updatesOperationHour, err := cs.UpdateBusinessCompanyServiceOperationHour(ctx, input)
	if err != nil {
		return nil, err
	}

	var operationHours = updatesOperationHour.GetBusinessCompanyServiceOperationHour()
	var resp = &model.UpdateBusinessCompanyServiceOperationHoursResponse{
		BusinessCompanyServiceOperationHour: &model.BusinessCompanyServiceOperationHour{
			ServiceOperationHourID: operationHours.GetServiceOperationHourID(),
			BusinessCompanyID:      operationHours.GetBusinessCompanyID(),
			BusinessServiceID:      operationHours.GetBusinessServiceID(),
			DayOfWeek:              operationHours.GetDayOfWeek(),
			OpenTime:               operationHours.GetOpenTime(),
			CloseTime:              operationHours.GetCloseTime(),
		},
	}

	return resp, nil
}

func (r *mutationResolver) DeleteBusinessCompanyServiceOperationHours(ctx context.Context, input model.DeleteBusinessCompanyServiceOperationHoursRequest) (*model.DeleteBusinessCompanyServiceOperationHoursResponse, error) {
	deletedOperationHour, err := cs.DeleteBusinessCompanyServiceOperationHour(ctx, input)
	if err != nil {
		return nil, err
	}

	var operationHours = deletedOperationHour.GetBusinessCompanyServiceOperationHour()
	var resp = &model.DeleteBusinessCompanyServiceOperationHoursResponse{
		BusinessCompanyServiceOperationHour: &model.BusinessCompanyServiceOperationHour{
			ServiceOperationHourID: operationHours.GetServiceOperationHourID(),
			BusinessCompanyID:      operationHours.GetBusinessCompanyID(),
			BusinessServiceID:      operationHours.GetBusinessServiceID(),
			DayOfWeek:              operationHours.GetDayOfWeek(),
			OpenTime:               operationHours.GetOpenTime(),
			CloseTime:              operationHours.GetCloseTime(),
		},
	}

	return resp, nil
}

func (r *mutationResolver) GenerateToken(ctx context.Context, input model.GenerateTokenRequest) (*model.Token, error) {
	validatePassword, err := bo.CheckOwnerPassword(ctx, input.Email, input.Password)
	if err != nil {
		return nil, err
	}

	if validatePassword.GetValid() == false {
		return nil, fmt.Errorf("password or email is invalid")
	}

	token, err := t.GenerateToken(input.Email, input.Password)
	if err != nil {
		return nil, err
	}

	return &model.Token{
		AccessToken:  token.GetAccessToken(),
		RefreshToken: token.GetRefreshToken(),
		ExpiresIn:    token.GetExpiresIn(),
		TokenType:    token.GetTokenType(),
	}, nil
}

func (r *queryResolver) GetBusinessServiceOrder(ctx context.Context, input model.GetBusinessServiceOrderRequest) (*model.GetBusinessServiceOrderResponse, error) {
	businessServiceOrder, err := bso.GetBusinessServiceOrder(ctx, input.BusinessServiceOrderID)
	if err != nil {
		return nil, err
	}

	var order = businessServiceOrder.GetBusinessServiceOrder()
	var resp = &model.GetBusinessServiceOrderResponse{
		BusinessServiceOrder: &model.BusinessServiceOrder{
			BusinessServiceOrderID:  order.BusinessServiceOrderID,
			ClientID:                order.ClientID,
			BusinessServiceID:       order.BusinessServiceID,
			StartAt:                 order.StartAt,
			EndAt:                   order.EndAt,
			CreatedAt:               order.CreatedAt,
			PrePaid:                 order.PrePaid,
			ClientFirstName:         order.ClientFirstName,
			ClientPhoneNumber:       order.ClientPhoneNumber,
			ClientPhoneNumberPrefix: order.ClientPhoneNumberPrefix,
			ClientCommentary:        order.ClientCommentary,
		},
	}

	return resp, nil
}

func (r *queryResolver) GetBusinessServiceOrders(ctx context.Context, input model.GetBusinessServiceOrdersRequest) (*model.GetBusinessServiceOrdersResponse, error) {
	orders, err := bso.GetBusinessServiceOrders(ctx, input.BusinessServiceID)
	if err != nil {
		return nil, err
	}

	b, err := pkg.Serializer(orders)
	if err != nil {
		return nil, err
	}

	var resp model.GetBusinessServiceOrdersResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *queryResolver) GetCompanyAvailableHoursByDate(ctx context.Context, input model.GetCompanyAvailableHoursByDateRequest) (*model.GetCompanyAvailableHoursByDateResponse, error) {
	availableHours, err := bso.GetBusinessServiceAvailableHours(ctx, input.BusinessServiceID, input.Date)
	if err != nil {
		return nil, err
	}

	b, err := pkg.Serializer(availableHours)
	if err != nil {
		return nil, err
	}

	var resp model.GetCompanyAvailableHoursByDateResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *queryResolver) GetBusinessServiceOrderByDate(ctx context.Context, input model.GetBusinessServiceOrderByDateRequest) (*model.GetBusinessServiceOrderByDateResponse, error) {
	orders, err := bso.GetBusinessServiceOrderByDate(ctx, input.BusinessServiceID, input.Date)
	if err != nil {
		return nil, err
	}

	b, err := pkg.Serializer(orders)
	if err != nil {
		return nil, err
	}

	var resp model.GetBusinessServiceOrderByDateResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *queryResolver) GetBusinessServiceOrdersByEmail(ctx context.Context, input model.GetBusinessServiceOrdersByEmailRequest) (*model.GetBusinessServiceOrdersByEmailResponse, error) {
	orders, err := bso.GetBusinessServiceOrdersByEmail(ctx, &input)
	if err != nil {
		return nil, err
	}

	b, err := pkg.Serializer(orders)
	if err != nil {
		return nil, err
	}

	var resp model.GetBusinessServiceOrdersByEmailResponse
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
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

func (r *queryResolver) GetBusinessCompaniesUnderCategory(ctx context.Context, input model.GetBusinessCompaniesUnderCategoryRequest) (*model.BusinessCompanies, error) {
	companies, err := bc.GetBusinessCompaniesUnderCategory(ctx, input.CategoryID)
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

func (r *queryResolver) GetBusinessOwnerCompanies(ctx context.Context, input *model.GetBusinessOwnerCompaniesRequest) (*model.GetBusinessOwnerCompaniesResponse, error) {
	operationHours, err := bo.GetBusinessOwnerCompanies(ctx, input.Email)
	if err != nil {
		return nil, err
	}

	b, err := pkg.Serializer(operationHours)
	if err != nil {
		return nil, err
	}

	var resp model.GetBusinessOwnerCompaniesResponse
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

func (r *queryResolver) GetBusinessCompanyServiceOperationHourByDay(ctx context.Context, input *model.GetGetBusinessCompanyServiceOperationHourByDayRequest) (*model.BusinessCompanyServiceOperationHourResponse, error) {
	operationHours, err := cs.GetBusinessCompanyServiceOperationHourByDay(ctx, input.ServiceID, input.DayOfWeek)
	if err != nil {
		return nil, err
	}

	return &model.BusinessCompanyServiceOperationHourResponse{
		BusinessCompanyServiceOperationHour: &model.BusinessCompanyServiceOperationHour{
			ServiceOperationHourID: operationHours.BusinessCompanyServiceOperationHour.GetServiceOperationHourID(),
			BusinessCompanyID:      operationHours.BusinessCompanyServiceOperationHour.GetBusinessCompanyID(),
			BusinessServiceID:      operationHours.BusinessCompanyServiceOperationHour.GetBusinessServiceID(),
			DayOfWeek:              operationHours.BusinessCompanyServiceOperationHour.GetDayOfWeek(),
			OpenTime:               operationHours.BusinessCompanyServiceOperationHour.GetOpenTime(),
			CloseTime:              operationHours.BusinessCompanyServiceOperationHour.GetCloseTime(),
		}}, nil
}

func (r *queryResolver) GetBusinessCompanyServiceOperationHours(ctx context.Context, input *model.GetBusinessCompanyServiceOperationHoursRequest) (*model.BusinessCompanyServiceOperationHours, error) {
	operationHours, err := cs.GetBusinessCompanyServiceOperationHours(ctx, input.ServiceID)
	if err != nil {
		return nil, err
	}

	b, err := pkg.Serializer(operationHours)
	if err != nil {
		return nil, err
	}

	var resp model.BusinessCompanyServiceOperationHours
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

func (r *queryResolver) GetCompanyServicesUnderCategory(ctx context.Context, input *model.GetCompanyServicesUnderCategoryRequest) (*model.CompanyServices, error) {
	services, err := cs.GetCompanyServicesUnderCategory(ctx, input.CategoryID)
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

func (r *queryResolver) RetrieveTokenInfo(ctx context.Context, input model.RetrieveTokenInfoRequst) (*model.RetrieveTokenInfoResponse, error) {
	token, err := t.RetrieveTokenInformation(input.AccessToken)
	if err != nil {
		return nil, err
	}

	return &model.RetrieveTokenInfoResponse{
		Email:     token.GetEmail(),
		ExpiresAt: token.GetExpiresAt(),
		IssuedAt:  token.GetIssuedAt(),
	}, nil
}

func (r *queryResolver) GetCustomerByEmail(ctx context.Context, input model.GetCustomerByEmailRequest) (*model.GetCustomerByEmailResponse, error) {
	var req = &pbc.GetCustomerByEmailRequest{
		Email: input.Email,
	}

	customer, err := client.GetCustomerByEmail(ctx, req)

	if err != nil {
		return nil, nil
	}

	b, err := pkg.Serializer(customer)
	if err != nil {
		return nil, err
	}

	var response *model.GetCustomerByEmailResponse
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *queryResolver) GetCustomerTokenInfo(ctx context.Context, input model.GetCustomerTokenInfoRequest) (*model.GetCustomerTokenInfoResponse, error) {
	var tokenInfoRequest = pba.RetrieveCustomerTokenInformationRequest{
		AccessToken: input.AccessToken,
	}
	info, err := client.RetrieveTokenInformation(context.Background(), &tokenInfoRequest)
	if err != nil {
		return nil, err
	}

	return &model.GetCustomerTokenInfoResponse{
		Email:     info.GetEmail(),
		IssuedAt:  info.GetIssuedAt(),
		ExpiresAt: info.GetExpiresAt(),
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
