package entity

// TODO: 这里可以用脚本进行生成
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

	UpdateInfo map[string]interface{}
}

func UserInfoTableName() string {
	return "user_info_tab"
}

func (e *UserInfo) set(k string, v interface{}) {
	if e.UpdateInfo == nil {
		e.UpdateInfo = map[string]interface{}{}
	}

	e.UpdateInfo[k] = v
}

func (e *UserInfo) SetID(v interface{}) *UserInfo {
	e.set("id", v)
	return e
}

func (e *UserInfo) SetRole(v interface{}) *UserInfo {
	e.set("role", v)
	return e
}

func (e *UserInfo) SetName(v interface{}) *UserInfo {
	e.set("name", v)
	return e
}

func (e *UserInfo) SetEmail(v interface{}) *UserInfo {
	e.set("email", v)
	return e
}

func (e *UserInfo) SetAvatar(v interface{}) *UserInfo {
	e.set("avatar", v)
	return e
}

func (e *UserInfo) SetNickName(v interface{}) *UserInfo {
	e.set("nick_name", v)
	return e
}

func (e *UserInfo) SetPassword(v interface{}) *UserInfo {
	e.set("name", v)
	return e
}

func (e *UserInfo) SetLoginTime(v interface{}) *UserInfo {
	e.set("login_time", v)
	return e
}

func (e *UserInfo) SetCreatedAt(v interface{}) *UserInfo {
	e.set("created_at", v)
	return e
}

func (e *UserInfo) SetUpdatedAt(v interface{}) *UserInfo {
	e.set("updated_at", v)
	return e
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
