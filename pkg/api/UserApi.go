package api

import (
	"entry-task/pkg/common"
	"entry-task/pkg/common/request"
	. "entry-task/pkg/common/util"
	"entry-task/pkg/service"
	"github.com/gin-gonic/gin"
)

// GetUserList
// path: /user/manage/list
func GetUserList(c *gin.Context) {
	var userListReq request.UserInfoPage
	c.ShouldBindJSON(&userListReq)
	err, result := service.GetUserInfoList(userListReq.Page, userListReq.PageSize)
	if err != nil {
		common.FailMessage(err.Error(), c)
		return
	}
	common.OkData(result, c)
	return
}

// AddUserInfo
// path: /user/manage/add
func AddUserInfo(c *gin.Context) {
	createUser := GetLoginUserId(c)

	var request request.UserInfoAdd
	c.ShouldBindJSON(&request)
	err := service.AddUserInfo(request, createUser)
	if err != nil {
		common.FailMessage(err.Error(), c)
		return
	}
	common.OkData(1, c)
	return
}

// DelUserInfo
// path: /user/manage/del
func DelUserInfo(c *gin.Context) {
	var userDelReq request.UserInfoDel
	c.ShouldBindJSON(&userDelReq)
	err, rowNum := service.DelUserInfos(userDelReq.UserIds)
	if err != nil {
		common.FailMessage(err.Error(), c)
		return
	}
	common.OkData(rowNum, c)
	return
}

// EditUserInfo
// path: /user/manage/edit
func EditUserInfo(c *gin.Context) {
	currentUserId := GetLoginUserId(c)
	var userEdit request.UserInfoEdit
	err := service.EditUserInfo(userEdit, currentUserId)
	if err != nil {
		common.FailMessage(err.Error(), c)
		return
	}

}
