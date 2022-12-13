package models

import (
	"context"
	"errors"
	"fmt"
	"github.com/aka-yz/go-micro-core/providers/config/log"
	"github.com/aka-yz/go-micro-core/providers/config/middleware/db"
	"github.com/gocraft/dbr/v2"
	"go-micro-demo/pkg/common/request"
	"go-micro-demo/pkg/common/response"
	"go-micro-demo/pkg/common/util"
	"go-micro-demo/pkg/data/entity"
)

type UserInfoModel interface {
	DelUserInfos(ctx context.Context, userIds []int) (rows int64, err error)
	GetUserByName(ctx context.Context, name string) (userInfo *entity.UserInfo, err error)
	GetUserDetail(ctx context.Context, userId int) (userDetail *entity.UserInfo, err error)
	AddUserInfo(ctx context.Context, userInfo *request.UserInfoAdd, createUser int64) (err error)
	UpdateUserInfo(ctx context.Context, userInfo *request.UserInfoEdit, editUser int64) (err error)
	GetUserInfoList(ctx context.Context, page uint64, pageSize uint64) (
		userListRes []*response.UserList, total uint64, err error)
}

var _ UserInfoModel = &UserInfoModelImpl{}

type UserInfoModelImpl struct {
	Connection *db.Connection `inject:"db.default"`
}

func (u *UserInfoModelImpl) DelUserInfos(ctx context.Context, userIds []int) (rows int64, err error) {
	log.Infof(ctx, "[DelUserInfos] param in:%v\n", userIds)
	if len(userIds) <= 0 {
		err = errors.New("用户id集合为空！")
		log.Warnf(ctx, err.Error())
		return
	}
	session := u.Connection.NewSession()
	rs, err := session.DeleteFrom(entity.UserInfoTableName()).
		Where(dbr.Eq(entity.ColumnsUserInfo.ID, userIds)).
		Exec()
	if err != nil {
		log.Error(ctx, err.Error())
		return 0, err
	}
	rows, err = rs.RowsAffected()
	if err != nil {
		log.Error(ctx, err.Error())
	}
	return
}

func (u *UserInfoModelImpl) GetUserByName(ctx context.Context, name string) (userInfo *entity.UserInfo, err error) {
	session := u.Connection.NewSession()
	_, err = session.Select(entity.ColumnsUserInfo.ID, entity.ColumnsUserInfo.Password).
		Where(dbr.Eq(entity.ColumnsUserInfo.Name, name)).Load(userInfo)
	if errors.Is(err, dbr.ErrNotFound) {
		err = errors.New("当前用户名对应数据不存在！")
		log.Error(ctx, err.Error())
		return
	}
	if err != nil {
		log.Error(ctx, err.Error())
		return
	}
	return
}

func (u *UserInfoModelImpl) GetUserDetail(ctx context.Context, userId int) (userDetail *entity.UserInfo, err error) {
	log.Infof(ctx, "[GetUserDetail] param in:%v\n", userId)
	session := u.Connection.NewSession()
	_, err = session.Select("*").
		Where(dbr.Eq(entity.ColumnsUserInfo.ID, userId)).
		Load(userDetail)
	if err != nil {
		err = errors.New(fmt.Sprintf("当前用户不存在, err=%v", err))
		log.Error(ctx, err.Error())
		return
	}
	return
}

func (u *UserInfoModelImpl) AddUserInfo(ctx context.Context, userInfo *request.UserInfoAdd, createUser int64) (err error) {
	log.Infof(ctx, "[AddUserInfo] param in:%v, %v\n", userInfo, createUser)
	var userEntity entity.UserInfo
	util.CopyProperties(userEntity, &userInfo)
	userEntity.CreatedAt = util.CurrentSecond()
	// 查询 username 是否重复
	session := u.Connection.NewSession()
	num, err := session.Select(entity.ColumnsUserInfo.Name).
		Where(dbr.Eq(entity.ColumnsUserInfo.Name, userInfo.Name)).
		Limit(1).Load(&entity.UserInfo{})
	if err != nil {
		log.Error(ctx, err.Error())
		return
	}
	if num > 0 {
		err = errors.New("用户名重复！")
		log.Error(ctx, err.Error())
		return
	}
	rs, err := session.InsertInto(entity.UserInfoTableName()).Exec()
	if err != nil {
		log.Error(ctx, err.Error())
		return
	}
	log.Infof(ctx, "[AddUserInfo] 新增用户数：%v\n", rs.RowsAffected)
	return
}

func (u *UserInfoModelImpl) UpdateUserInfo(ctx context.Context, userInfo *request.UserInfoEdit, editUser int64) (err error) {
	log.Infof(ctx, "[UpdateUserInfo] param in:%v, %v", userInfo, editUser)
	var user entity.UserInfo
	session := u.Connection.NewSession()
	_, err = session.Select("*").
		Where(dbr.Eq(entity.ColumnsUserInfo.ID, userInfo.Id)).
		Limit(1).Load(&user)
	if err != nil {
		log.Error(ctx, err.Error())
		return
	}
	var userEntity entity.UserInfo
	util.CopyProperties(&userEntity, userInfo)
	userEntity.UpdatedAt = util.CurrentSecond()
	rs, err := session.Update(entity.UserInfoTableName()).Exec()
	if err != nil {
		log.Error(ctx, err.Error())
		return
	}
	log.Infof(ctx, "[UpdateUserInfo] 更新行数：%v\n", rs.RowsAffected)
	return
}

func (u *UserInfoModelImpl) GetUserInfoList(ctx context.Context, page uint64, pageSize uint64) (
	userListRes []*response.UserList, total uint64, err error) {
	log.Infof(ctx, "[GetUserInfoList] param in:%v, %v\n", page, pageSize)
	session := u.Connection.NewSession()

	num := 0
	_, err = session.Select("count(*)").Load(num)
	if err != nil {
		log.Error(ctx, err.Error())
		return
	}

	var userList []*entity.UserInfo
	offset := (page - 1) * pageSize
	_, err = session.Select("*").Offset(offset).Limit(pageSize).Load(userList)
	if err != nil {
		log.Error(ctx, err.Error())
		return
	}
	// 对象集合拷贝
	for _, entity := range userList {
		var user response.UserList
		err = util.CopyProperties(user, entity)
		if err != nil {
			return
		}
		userListRes = append(userListRes, &user)
	}
	return
}
