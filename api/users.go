package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"

	"github.com/julienschmidt/httprouter"
)

// User type in database
type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Hashpass string `json:"-"`
	Perms    int    `json:"perms"`
}

// Session type in database
type Session struct {
	User       User   `json:"user"`
	Token      string `json:"token"`
	Expiration int    `json:"expiration"`
}

// SessionResponse data response type
type SessionResponse struct {
	APIResponse
	Session
}

// RandStringRunes generates random string of n runes among `among` runes
func RandStringRunes(n int, among string) string {
	letterRunes := []rune(among)
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func newSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	login := r.FormValue("login")
	password := r.FormValue("password")

	w.Header().Set("Content-Type", "application/json")

	db, err := sql.Open("mysql", DSN)

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	_, err = db.Exec("DELETE FROM session WHERE expiration < ?", time.Now().Unix())

	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	row := db.QueryRow("SELECT id, login, hashpass, perms FROM users WHERE login = ?", login)

	var u User

	err = row.Scan(&u.ID, &u.Login, &u.Hashpass, &u.Perms)

	if err != nil {
		res := APIResponse{Status: -1, Message: "Identifiant invalide"}

		bres, err := json.Marshal(res)

		if err != nil {
			fmt.Println("Erreur :", err)
			return
		}

		w.Write(bres)
		return // No user named login
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Hashpass), []byte(password))

	if err != nil {
		res := APIResponse{Status: -1, Message: "Mot de passe invalide"}

		bres, err := json.Marshal(res)

		if err != nil {
			fmt.Println("Erreur :", err)
			return
		}

		w.Write(bres)
		return // Wrong password
	}

	token := RandStringRunes(128, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	expiration := int(time.Now().Add(2 * time.Hour).Unix())

	_, err = db.Exec("INSERT INTO session (id, user, token, expiration) VALUES (NULL, ?, ?, ?)", u.ID, token, expiration)

	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	db.Close()

	res := SessionResponse{}

	res.Status = 1
	res.Message = "ok"
	res.User = u
	res.Token = token
	res.Expiration = expiration

	bres, err := json.Marshal(res)

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	w.Write(bres)
}

func delSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	token := r.FormValue("token")

	w.Header().Set("Content-Type", "application/json")

	db, err := sql.Open("mysql", DSN)

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	_, err = db.Exec("DELETE FROM session WHERE expiration < ?", time.Now().Unix())

	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	_, err = db.Exec("DELETE FROM session WHERE token = ?", token)

	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	db.Close()

	res := APIResponse{Status: 1, Message: "ok"}

	bres, err := json.Marshal(res)

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	w.Write(bres)
}

func isLoggedIn(token string, login string) (bool, int) {
	db, err := sql.Open("mysql", DSN)

	if err != nil {
		fmt.Println("Erreur :", err)
		return false, -1
	}

	defer db.Close()

	_, err = db.Exec("DELETE FROM session WHERE expiration < ?", time.Now().Unix())

	if err != nil {
		fmt.Println("Erreur :", err)
		return false, -1
	}

	rows, err := db.Query("SELECT u.perms FROM session s, users u WHERE s.token = ? AND u.login = ? AND s.user = u.id", token, login)

	if err != nil {
		fmt.Println("Erreur :", err)
		return false, -1
	}

	var perms int
	i := 0

	for ; rows.Next(); i++ {
		rows.Scan(&perms)
	}

	if i == 1 {
		return true, perms
	}

	return false, -1
}

func verifData(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	loggedIn, perms := isLoggedIn(r.FormValue("token"), r.FormValue("login"))

	if !loggedIn {
		respond(w, -1, "not logged in")
		return
	}

	if perms == 1 {
		respond(w, 1, "Fondator")
		return
	} else if perms == 2 {
		respond(w, 2, "Admin")
		return
	} else if perms == 3 {
		respond(w, 3, "Modetator")
		return
	} else {
		respond(w, -2, "not correct perms")
		return
	}

}

func addUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	loggedIn, perms := isLoggedIn(r.FormValue("token"), r.FormValue("login"))

	if !loggedIn {
		respond(w, -1, "not logged in")
		return
	}

	if perms != 1 {
		respond(w, -2, "not allowed")
		return
	}

	db, err := sql.Open("mysql", DSN)

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	defer db.Close()

	newlogin := r.FormValue("newlogin")
	pass := r.FormValue("password")
	permUser := r.FormValue("perms")

	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	_, err = db.Exec("INSERT INTO users (id, login, hashPass, perms) VALUES (NULL, ?, ?, ?)", newlogin, string(hash), permUser)

	if err != nil {
		res := APIResponse{Status: -1, Message: "User already exist"}

		bres, err := json.Marshal(res)

		if err != nil {
			fmt.Println("Erreur :", err)
			return
		}

		w.Write(bres)
		return // No user named login
	}

	res := APIResponse{Status: 1, Message: "ok"}

	bres, err := json.Marshal(res)

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	w.Write(bres)
}

func delUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	loggedIn, perms := isLoggedIn(r.FormValue("token"), r.FormValue("login"))

	if !loggedIn {
		respond(w, -1, "not logged in")
		return
	}

	if perms != 1 {
		respond(w, -2, "not allowed")
		return
	}

	db, err := sql.Open("mysql", DSN)

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	defer db.Close()

	_, err = db.Exec("DELETE FROM users WHERE id = ?", r.FormValue("id"))

	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	res := APIResponse{Status: 1, Message: "ok"}

	bres, err := json.Marshal(res)

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	w.Write(bres)
}

// ListUsersResponse type user in database
type ListUsersResponse struct {
	APIResponse
	Users []User `json:"users"`
}

func listUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	loggedIn, perms := isLoggedIn(r.FormValue("token"), r.FormValue("login"))

	if !loggedIn {
		respond(w, -1, "not logged in")
		return
	}

	if perms != 1 {
		respond(w, -2, "not allowed")
		return
	}

	db, err := sql.Open("mysql", DSN)

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	rows, err := db.Query("SELECT u.id, u.login, u.perms FROM users u")

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	var u User
	var res ListUsersResponse

	for rows.Next() {
		rows.Scan(&u.ID, &u.Login, &u.Perms)
		res.Users = append(res.Users, u)
	}

	rows.Close()

	res.Status = 1
	res.Message = "ok"

	db.Close()

	bres, err := json.Marshal(res)

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	w.Write(bres)
}

func userByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	loggedIn, perms := isLoggedIn(r.FormValue("token"), r.FormValue("login"))

	if !loggedIn {
		respond(w, -1, "not logged in")
		return
	}

	if perms != 1 {
		respond(w, -2, "not allowed")
		return
	}

	db, err := sql.Open("mysql", DSN)

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	rows, err := db.Query("SELECT u.id, u.login, u.perms FROM users u WHERE u.id = ?", ps.ByName("idUser"))

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	var u User
	var res ListUsersResponse

	for rows.Next() {
		rows.Scan(&u.ID, &u.Login, &u.Perms)
		res.Users = append(res.Users, u)
	}

	rows.Close()

	res.Status = 1
	res.Message = "ok"

	db.Close()

	bres, err := json.Marshal(res)

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	w.Write(bres)
}

func changeLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	loggedIn, perms := isLoggedIn(r.FormValue("token"), r.FormValue("login"))

	if !loggedIn {
		respond(w, -1, "not logged in")
		return
	}

	if perms != 1 {
		respond(w, -2, "not allowed")
		return
	}

	db, err := sql.Open("mysql", DSN)
	defer db.Close()

	newLogin := r.FormValue("newLogin")
	userID := r.FormValue("userID")

	_, err = db.Exec("UPDATE users SET login = ? WHERE id = ?", newLogin, userID)

	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	db.Close()

	res := APIResponse{Status: 1, Message: "ok"}

	bres, err := json.Marshal(res)

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	w.Write(bres)
}

func changePass(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")

	loggedIn, perms := isLoggedIn(r.FormValue("token"), r.FormValue("login"))

	if !loggedIn {
		respond(w, -1, "not logged in")
		return
	}

	if perms != 1 {
		respond(w, -2, "not allowed")
		return
	}

	db, err := sql.Open("mysql", DSN)
	defer db.Close()

	oldPass := r.FormValue("oldPass")
	newPass := r.FormValue("newPass")
	userID := r.FormValue("userID")

	row := db.QueryRow("SELECT id, login, hashpass, perms FROM users WHERE id = ?", userID)

	var u User

	err = row.Scan(&u.ID, &u.Login, &u.Hashpass, &u.Perms)

	if err != nil {
		res := APIResponse{Status: -1, Message: "Invalid user ID"}

		bres, err := json.Marshal(res)

		if err != nil {
			fmt.Println("Erreur :", err)
			return
		}

		w.Write(bres)
		return // No user named login
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Hashpass), []byte(oldPass))

	if err != nil {
		res := APIResponse{Status: -1, Message: "Wrong current Password"}

		bres, err := json.Marshal(res)

		if err != nil {
			fmt.Println("Erreur :", err)
			return
		}

		w.Write(bres)
		return // Wrong password
	}

	password := []byte(newPass)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	_, err = db.Exec("UPDATE users SET hashPass = ? WHERE id = ?", hashedPassword, userID)

	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	db.Close()

	res := APIResponse{Status: 1, Message: "ok"}

	bres, err := json.Marshal(res)

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	w.Write(bres)
}

func changePerms(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	loggedIn, perms := isLoggedIn(r.FormValue("token"), r.FormValue("login"))

	if !loggedIn {
		respond(w, -1, "not logged in")
		return
	}

	if perms != 1 {
		respond(w, -2, "not allowed")
		return
	}

	db, err := sql.Open("mysql", DSN)
	defer db.Close()

	newPerms := r.FormValue("newperms")
	userID := r.FormValue("userID")

	_, err = db.Exec("UPDATE users SET perms = ? WHERE id = ?", newPerms, userID)

	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	db.Close()

	res := APIResponse{Status: 1, Message: "ok"}

	bres, err := json.Marshal(res)

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	w.Write(bres)
}
