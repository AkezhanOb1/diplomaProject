package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"


	"github.com/AkezhanOb1/diplomaProject/api/graphQL/graph/generated"
	"github.com/AkezhanOb1/diplomaProject/api/graphQL/graph/model"
	"github.com/AkezhanOb1/diplomaProject/pkg"
	c "github.com/AkezhanOb1/diplomaProject/services/client/business/category"
	bc "github.com/AkezhanOb1/diplomaProject/services/client/business/company"
	bo "github.com/AkezhanOb1/diplomaProject/services/client/business/owner"
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
