package database

type Users struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Password string `json:"password"`
	IsDelete bool `json:"-"`
}