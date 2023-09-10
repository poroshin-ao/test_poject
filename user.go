package todo

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"reguired"`
	Username string `json:"username" binding:"reguired"`
	Password string `json:"password" binding:"reguired"`
}
