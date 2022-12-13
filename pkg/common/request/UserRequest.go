package request

type UserInfoAdd struct {
	Name     string `json:"name"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     int    `json:"role"` // 0-普通用户|1-运营人员
}

type UserInfoEdit struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     int    `json:"role"` // 0-普通用户|1-运营人员
}

type UserInfoDel struct {
	UserIds []int `json:"UserIds" validated:"required"`
}

type UserInfoPage struct {
	Page     int `json:"page" validated:"required"`
	PageSize int `json:"pageSize" validated:"required"`
}
