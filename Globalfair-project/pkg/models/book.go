package models 

import(
	"github.com/jinzhu/gorm"
	"github.com/aviuday/Globalfair-project/pkg/config"
)

var db *gorm.DB

type User struct{
	gorm.Model 
	Name string `gorm:""json:"name"`
	MobileNo string `json:"mobileno"`
	Salary string `json:"salary"`
	UserID string `json:"userid"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User{
	db.NewRecord(u)
	db.Create(&u)
	return u 
}

func GetAllUsers() []User{
	var Users []User 
	db.Find(&Users)
	return Users 
}

func GetUserByID(Id1 int64) (*User, *gorm.DB){
	var getUser User
	db := db.Where("ID=?", Id1).Find(&getUser)
	return &getUser, db
}

func DeleteUser(Id2 int64) User{
	var user User 
	db.Where("ID=?", Id2).Delete(user)
	return user 
}
