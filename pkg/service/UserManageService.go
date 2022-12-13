package service

import (
	"context"
	"go-micro-demo/pkg/common"
	"go-micro-demo/pkg/common/request"
	"go-micro-demo/pkg/models"
)

type UserInfoService interface {
	AddUserInfo(ctx context.Context, userInfo *request.UserInfoAdd, createUser int64) (err error)
	EditUserInfo(ctx context.Context, userInfo *request.UserInfoEdit, editUser int64) (err error)
	GetUserInfoList(ctx context.Context, page uint64, pageSize uint64) (err error, result *common.PageResult)
}

var _ UserInfoService = &UserInfoServiceImpl{}

type UserInfoServiceImpl struct {
	UserInfoModel models.UserInfoModel `inject:""`
}

func (u *UserInfoServiceImpl) AddUserInfo(ctx context.Context, userInfo *request.UserInfoAdd, createUser int64) (err error) {
	err = u.UserInfoModel.AddUserInfo(ctx, userInfo, createUser)
	return
}

func (u *UserInfoServiceImpl) EditUserInfo(ctx context.Context, userInfo *request.UserInfoEdit, editUser int64) (err error) {
	err = u.UserInfoModel.UpdateUserInfo(ctx, userInfo, editUser)
	return
}

func (u *UserInfoServiceImpl) GetUserInfoList(ctx context.Context, page uint64, pageSize uint64) (err error, result *common.PageResult) {
	list, total, err := u.UserInfoModel.GetUserInfoList(ctx, page, pageSize)
	if err != nil {
		return
	}
	result = &common.PageResult{
		List:     list,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}
	return
}

func (u *UserInfoServiceImpl) DelUserInfos(ctx context.Context, userIds []int) (err error, row int64) {
	row, err = u.UserInfoModel.DelUserInfos(ctx, userIds)
	return
}
