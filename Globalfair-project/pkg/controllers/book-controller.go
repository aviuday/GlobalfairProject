package controllers 

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/aviuday/Globalfair-project/pkg/utils"
	"github.com/aviuday/Globalfair-project/pkg/models"
)

var NewUser models.User

func UserDetails(w http.ResponseWriter, r *http.Request){
	newUsers := models.GetAllUsers()
	res, _ := json.Marshal(newUsers)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	userId := vars["userId"]
	Id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	userDetails, _ := models.GetUserByID(Id)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	createUser := &models.User{}
	utils.ParseBody(r, createUser)
	u := createUser.CreateUser()
	res, _ := json.Marshal(u)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	userId := vars["userId"]
	Id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	user := models.DeleteUser(Id)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	userId := vars["userId"]
	Id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	userDetails, db := models.GetUserByID(Id)
	if updateUser.Name != ""{
		userDetails.Name = updateUser.Name
	}
	if updateUser.MobileNo != ""{
		userDetails.MobileNo = updateUser.MobileNo
	}
	if updateUser.Salary != ""{
		userDetails.Salary = updateUser.Salary
	}
	if updateUser.UserID != ""{
		userDetails.UserID = updateUser.UserID
	}
	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


