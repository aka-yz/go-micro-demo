package entity

type UserInfo struct {
	Id        int64  `db:"id"`
	Role      int    `db:"role"`
	Name      string `db:"name"`
	Email     string `db:"email"`
	Avatar    string `db:"avatar"`
	NickName  string `db:"nick_name"`
	Password  string `db:"password"`
	LoginTime int    `db:"login_time"`
	CreatedAt int    `db:"created_at"`
	UpdatedAt int    `db:"updated_at"`
}

func UserInfoTableName() string {
	return "user_info_tab"
}

var ColumnsUserInfoFields = []string{
	"id",
	"role",
	"name",
	"email",
	"avatar",
	"nick_name",
	"password",
	"login_time",
	"created_at",
	"updated_at",
}

type columnsUserInfoType struct {
	ID        string
	Role      string
	Name      string
	Email     string
	Avatar    string
	NickName  string
	Password  string
	LoginTime string
	CreatedAt string
	UpdatedAt string
}

var ColumnsUserInfo = columnsUserInfoType{
	ID:        "id",
	Role:      "role",
	Name:      "name",
	Email:     "email",
	Avatar:    "avatar",
	NickName:  "nick_name",
	Password:  "password",
	LoginTime: "login_time",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}
