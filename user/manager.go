package user

type User struct {
	UserName string `form:"username" json:"username"`
	PassWord string `form:"password" json:"password"`
}
