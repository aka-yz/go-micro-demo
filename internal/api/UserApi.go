package api

import (
	"context"
	"github.com/aka-yz/go-micro-core/providers/transport/http"
	"github.com/gin-gonic/gin"
	"go-micro-demo/internal/common"
	"go-micro-demo/internal/common/request"
	"go-micro-demo/internal/service"
)

type UserApi interface {
	GetUserList(c *gin.Context)
	AddUserInfo(c *gin.Context)
	DelUserInfo(c *gin.Context)
	EditUserInfo(c *gin.Context)
}

var _ UserApi = &UserApiImpl{}

type UserApiImpl struct {
	httpClient  http.HttpClient
	UserService service.UserInfoService `inject:""`
}

// GetUserList
// path: /user/manage/list
func (u *UserApiImpl) GetUserList(c *gin.Context) {
	var userListReq request.UserInfoPage
	c.ShouldBindJSON(&userListReq)
	err, result := u.UserService.GetUserInfoList(context.Background(), userListReq.Page, userListReq.PageSize)
	if err != nil {
		common.FailMessage(err.Error(), c)
		return
	}
	common.OkData(result, c)
	return
}

// AddUserInfo
// path: /user/manage/add
func (u *UserApiImpl) AddUserInfo(c *gin.Context) {
	var request request.UserInfoAdd
	c.ShouldBindJSON(&request)
	err := u.UserService.AddUserInfo(context.Background(), &request, 1)
	if err != nil {
		common.FailMessage(err.Error(), c)
		return
	}
	common.OkData(1, c)
	return
}

// DelUserInfo
// path: /user/manage/del
func (u *UserApiImpl) DelUserInfo(c *gin.Context) {
	var userDelReq request.UserInfoDel
	c.ShouldBindJSON(&userDelReq)
	err, rowNum := u.UserService.DelUserInfos(context.Background(), userDelReq.UserIds)
	if err != nil {
		common.FailMessage(err.Error(), c)
		return
	}
	common.OkData(rowNum, c)
	return
}

// EditUserInfo
// path: /user/manage/edit
func (u *UserApiImpl) EditUserInfo(c *gin.Context) {
	var userEdit request.UserInfoEdit
	c.ShouldBindJSON(&userEdit)
	err := u.UserService.EditUserInfo(context.Background(), &userEdit, 1)
	if err != nil {
		common.FailMessage(err.Error(), c)
		return
	}

}
