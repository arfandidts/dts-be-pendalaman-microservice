package database

type Auth struct {
	ID       string `json:"-" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
