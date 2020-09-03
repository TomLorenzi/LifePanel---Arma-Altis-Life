package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Database connections infos
const (
	//MODIFY HERE
	DBHOST string = "<ip adress of your DB server>"
	DBPORT string = "3306" // Default port
	DBUSER string = "<User>"
	DBPWRD string = "<Password>"
	DBNAME string = "altislife" //Should be correct

	//DO NOT TOUCH
	DSN string = DBUSER + ":" + DBPWRD + "@tcp(" + DBHOST + ":" + DBPORT + ")/" + DBNAME + "?parseTime=true"
)

// APIResponse basic api respnse
type APIResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func main() {
	router := httprouter.New()

	router.GET("/users/sessions/new", newSession)
	router.GET("/users/sessions/del", delSession) // Protected par auth
	router.GET("/users/add", addUser)             // Protected by auth
	router.GET("/users/del", delUser)             // Protected by auth
	router.GET("/users/id/:idUser", userByID)     // Protected by auth
	router.GET("/users/newlogin", changeLogin)    // Protected by auth
	router.GET("/users/newpass", changePass)      // Protected by auth
	router.GET("/users/newperm", changePerms)     // Protected by auth
	router.GET("/users", listUser)                // Protected by auth

	router.GET("/verif", verifData) // Protected by auth

	router.GET("/players/name", listPlayersByName)  //Protected by auth
	router.GET("/players/uid", listPlayersByUID)    //Protected by auth
	router.GET("/players/id/:idPlayer", playerByID) //Protected by auth
	router.GET("/players/changecash", changeCash)   //Protected by auth
	router.GET("/players/changebank", changeBank)   //Protected by auth
	router.GET("/players/changecop", changeCop)     //Protected by auth
	router.GET("/players/changemedic", changeMedic) //Protected by auth
	router.GET("/players/changedonor", changeDonor) //Protected by auth
	router.GET("/players/changeadmin", changeAdmin) //Protected by auth

	router.GET("/logs", listLogs) //Protected by auth

	http.ListenAndServe(":8080", router)
}

func respond(w http.ResponseWriter, status int, message string) {
	res := APIResponse{Status: status, Message: message}

	bres, err := json.Marshal(res)

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	w.Write(bres)
}
