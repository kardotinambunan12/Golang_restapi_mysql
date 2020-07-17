package models


type Users struct {
	Id           int     `gorm:"AUTO_INCREMENT"`
	First_name   string   `gorm: "type:varchar(150)"`
	Last_name    string  `gorm: "type:varchar(150)"`
	Alamat       string   `gorm: "type:varchar(150)"`
	Zip_code     string    `gorm: "type: varchar(10)"`
}
func(Users) TableName() string{
	return "tb_user"
}

type MasterUser struct {
	Id        	  		   string `json:"id"`
	First_name        	   string `json:"firt_name"`
	Last_name        	  	string `json:"last_name"`
	Alamat        	        string `json:"alamat"`
	Zip_code               string  `json : "zip_code`
	Additional           	string `json:"additional"`
}
type UsersModels struct{
	UsersModels []MasterUser `json:"users"`
}
