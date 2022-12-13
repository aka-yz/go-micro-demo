package response

// UserCommentInfo : 活动详情需要
type UserCommentInfo struct {
	Id       int    `json:"id" gorm:"column:commentId"`
	Name     string `json:"name" gorm:"column:name"`
	NickName string `json:"nickName" gorm:"column:nickName"`
	Avatar   string `json:"avatar" gorm:"column:avatar"`
	Comment  string `json:"comment" gorm:"column:comment"`
}

type UserDetail struct {
	Name      string `json:"name" gorm:"not null;unique"`
	NickName  string `json:"nickName"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	Role      int    `json:"role"` // 0-普通用户|1-运营
	LoginTime int    `json:"loginTime"`
	CreatedAt int    `json:"createdAt"`
}

type UserLogInfo struct {
	Id       int    `json:"id" gorm:"column:id"`
	Name     string `json:"name" gorm:"column:name"`
	Password string `json:"password" gorm:"column:password"`
}

type UserList struct {
	Id        int    `json:"id"`
	Name      string `json:"name" gorm:"not null;unique"`
	NickName  string `json:"nickName"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	Role      int    `json:"role"`
	LoginTime int    `json:"loginTime"`
	CreatedAt int    `json:"createdAt"`
}
