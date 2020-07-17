package controllers
import (
	"net/http"
	"crypto/md5"
	"encoding/hex"
	"../../database"
	"../models"
	"github.com/labstack/echo"
)

func GetUsersController(c echo.Context) error {
	db := database.CreateCon()
	defer db.Close()

	rows, err := db.Raw("SELECT id, first_name, last_name, alamat, zip_code FROM tb_user").Rows()
	if err != nil {
		logs.Println(err)
		internal_server_error(c)
		return nil
	}
	defer rows.Close()

	each   := models.MasterUser{}
	result := []models.MasterUser{}

	for rows.Next() {

		var id, first_name, last_name, alamat, zip_code[]byte
		err  := rows.Scan(&id, &first_name, &last_name, &alamat, &zip_code)
		if err != nil {
			logs.Println(err)
			internal_server_error(c)
			return nil
		}

		var str string = string(id)
		hasher := md5.New()
		hasher.Write([]byte(str))
		converId := hex.EncodeToString(hasher.Sum(nil))

		each.Id         = string(id)
		each.First_name = string(first_name)
		each.Last_name  = string(last_name)
		each.Alamat     = string(alamat)
		each.Zip_code   =string (zip_code)
		each.Additional = converId

		result          = append(result, each)
	}
	response     := response_json{
		"data"   :   result,
		"status" :   status_200,
	}
	return c.JSON(http	.StatusOK, response)
}

func AddUserController(c echo.Context) error{
	db := database.CreateCon()
	defer db.Close()

	first_name     := c.FormValue("first_name")
	last_name 	  := c.FormValue("last_name")
	alamat      := c.FormValue("alamat")
	zip_code   := c.FormValue("zip_code")
	//zip_code,_     := strconv.Atoi(c.FormValue("zip_code"))


	user := models.Users{
		First_name  : first_name,
		Last_name   : last_name,	
		Alamat      : alamat,
		Zip_code    : zip_code,
	}
	db.NewRecord(user)

	if error_insert := db.Create(&user); error_insert.Error != nil {
		logs.Println(error_insert)
		internal_server_error(c)
		return nil
	}
	db.NewRecord(user)

	response := response_json{
		"status"		: status_200,
	}
	return c.JSON(200, response)

}

func EditUsersController(c echo.Context) error{
	
	db := database.CreateCon()
	defer db.Close()


	// Define_fo

	//zip_code,_     := strconv.Atoi(c.FormValue("zip_code"))

	id              := c.Param("id")
	first_name      := c.FormValue("first_name")
	last_name 	    := c.FormValue("last_name")
	alamat          := c.FormValue("alamat")
	zip_code        :=c.FormValue("zip_code")

	// insert
	var user models.Users  
	data := db.Model(&user).Where("id = ?", id). Updates(map[string]interface{}{
      "first_name"  : first_name,
      "last_name"   : last_name,
      "alamat"      : alamat,
      "zip_code"    : zip_code,
     
	}) 
 

       // cek apakah data bisa di ubah atau tidak
	if data.Error != nil {
		logs.Println(data.Error)
		internal_server_error(c)
		return nil
	}else if data.RowsAffected  == 0 {
		not_modified(c)
		return nil
	}
	response   := response_json{
		"status"   : status_200,
	}
	return c.JSON(200, response)
}

func DeleteUsersController(c echo.Context) error{
	db := database.CreateCon()
	defer db.Close()


	id       := c.Param("id")
	var user models.Users
	data := db.Unscoped().Where("id = ?", id).Delete(&user)
	if data.Error != nil{
		logs.Println(data.Error)
		internal_server_error(c)
		return nil
	}else if data.RowsAffected == 0{
		not_modified(c)
		return nil
	}

	response := response_json{
		"status"		: status_200,
	}

	return c.JSON(http.StatusOK, response)
}








