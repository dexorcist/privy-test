package service

import (
	"context"
	"database/sql"
	"errors"
	"privy-test/enum"
	"privy-test/infra"
	"privy-test/integration/logging"
	"privy-test/model"
	"privy-test/param/cake"
	"privy-test/repository"
	"privy-test/utils"
)

type cakeService struct {
	config   infra.MergeConfig
	logger   logging.Logger
	cakeRepo repository.CakeRepository
}

func NewCakeService(
	config infra.MergeConfig,
	cakeRepo repository.CakeRepository,
	logger logging.Logger,
) CakeService {
	return &cakeService{
		config:   config,
		cakeRepo: cakeRepo,
		logger:   logger,
	}
}

func (c cakeService) GetDetail(ctx context.Context, cakeID int64) (*cake.HTTPGetDetailCakeResponse, error) {
	cakeModel, err := c.cakeRepo.GetDetail(ctx, cakeID)
	if err != nil {
		c.logger.ErrorWithContext(ctx, err, "Error Get Detail Cake %s", err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return nil, utils.DefaultMultiInternalError(enum.HTTPErrorCakeNotFound)
		}
		return nil, utils.DefaultMultiInternalError(enum.HTTPErrorInternalServerError)
	}
	return &cake.HTTPGetDetailCakeResponse{
		Data: cakeModel.ConvertResponse(),
	}, nil
}

func (c cakeService) Create(ctx context.Context, request *cake.CreateUpdateRequest) (*cake.HTTPGetDetailCakeResponse, error) {
	cakeModel := new(model.Cake)
	cakeModel, err := c.cakeRepo.Create(ctx, cakeModel.ParamToModel(request))
	if err != nil {
		c.logger.ErrorWithContext(ctx, err, "Error Create Cake %s", err.Error())
		return nil, utils.DefaultMultiInternalError(enum.HTTPErrorInternalServerError)
	}
	return &cake.HTTPGetDetailCakeResponse{
		Data: cakeModel.ConvertResponse(),
	}, nil
}

func (c cakeService) Delete(ctx context.Context, cakeID int64) error {
	_, err := c.cakeRepo.GetDetail(ctx, cakeID)
	if err != nil {
		c.logger.ErrorWithContext(ctx, err, "Error Get Detail Cake %s", err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return utils.DefaultMultiInternalError(enum.HTTPErrorCakeNotFound)
		}
		return utils.DefaultMultiInternalError(enum.HTTPErrorInternalServerError)
	}

	err = c.cakeRepo.Delete(ctx, cakeID)
	if err != nil {
		c.logger.ErrorWithContext(ctx, err, "Error Delete Cake %s", err.Error())
		return utils.DefaultMultiInternalError(enum.HTTPErrorInternalServerError)
	}
	return nil
}

func (c cakeService) Update(ctx context.Context, cakeID int64, request *cake.CreateUpdateRequest) (*cake.HTTPGetDetailCakeResponse, error) {
	cakeModel := new(model.Cake)

	_, err := c.cakeRepo.GetDetail(ctx, cakeID)
	if err != nil {
		c.logger.ErrorWithContext(ctx, err, "Error Get Detail Cake %s", err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return nil, utils.DefaultMultiInternalError(enum.HTTPErrorCakeNotFound)
		}
		return nil, utils.DefaultMultiInternalError(enum.HTTPErrorInternalServerError)
	}

	cakeModel, err = c.cakeRepo.Update(ctx, cakeID, cakeModel.ParamToModel(request))
	if err != nil {
		c.logger.ErrorWithContext(ctx, err, "Error Update Cake %s", err.Error())
		return nil, utils.DefaultMultiInternalError(enum.HTTPErrorInternalServerError)
	}
	return &cake.HTTPGetDetailCakeResponse{
		Data: cakeModel.ConvertResponse(),
	}, nil
}

func (c cakeService) GetList(ctx context.Context, request *cake.FindAllRequest) (*cake.HTTPGetListCakeResponse, error) {
	cakeList, err := c.cakeRepo.GetList(ctx, request)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &cake.HTTPGetListCakeResponse{
				Data: nil,
			}, nil
		}
		c.logger.ErrorWithContext(ctx, err, "Error Get List Cake %s", err.Error())
		return nil, utils.DefaultMultiInternalError(enum.HTTPErrorInternalServerError)
	}
	return &cake.HTTPGetListCakeResponse{
		Data: c.makeListResponse(cakeList),
	}, nil
}

func (c cakeService) makeListResponse(list []model.Cake) []cake.DetailCakeResponse {
	var cakeList []cake.DetailCakeResponse

	for idx := range list {
		cakeList = append(cakeList, list[idx].ConvertResponse())
	}

	return cakeList
}
