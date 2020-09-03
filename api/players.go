package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/julienschmidt/httprouter"
)

// Player type players in database
type Player struct {
	UID        int    `json:"uid"`
	Pid        string `json:"pid"`
	Name       string `json:"name"`
	Cash       int    `json:"cash"`
	Bank       int    `json:"bank"`
	CopLevel   int    `json:"copLevel"`
	MedicLevel int    `json:"medicLevel"`
	DonorLevel int    `json:"donorLevel"`
	AdminLevel int    `json:"adminLevel"`
}

// ListPlayersResponse response api data
type ListPlayersResponse struct {
	APIResponse
	Players []Player `json:"players"`
}

func listPlayers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	loggedIn, perms := isLoggedIn(r.FormValue("token"), r.FormValue("login"))

	if !loggedIn {
		respond(w, -1, "not logged in")
		return
	}

	if perms != 1 && perms != 2 {
		respond(w, -2, "not allowed")
		return
	}

	db, err := sql.Open("mysql", DSN)

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	rows, err := db.Query("SELECT uid, pid, name, cash, bankacc, coplevel, mediclevel, donorlevel, adminlevel FROM players")

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	var p Player
	var res ListPlayersResponse

	for rows.Next() {
		rows.Scan(&p.UID, &p.Pid, &p.Name, &p.Cash, &p.Bank, &p.CopLevel, &p.MedicLevel, &p.DonorLevel, &p.AdminLevel)
		res.Players = append(res.Players, p)
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

func listPlayersByName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	loggedIn, perms := isLoggedIn(r.FormValue("token"), r.FormValue("login"))

	if !loggedIn {
		respond(w, -1, "not logged in")
		return
	}

	if perms != 1 && perms != 2 {
		respond(w, -2, "not allowed")
		return
	}

	db, err := sql.Open("mysql", DSN)

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	rows, err := db.Query("SELECT uid, pid, name, cash, bankacc, coplevel, mediclevel, donorlevel, adminlevel FROM players WHERE name LIKE CONCAT('%', ?, '%')", r.FormValue("name"))

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	var p Player
	var res ListPlayersResponse

	for rows.Next() {
		rows.Scan(&p.UID, &p.Pid, &p.Name, &p.Cash, &p.Bank, &p.CopLevel, &p.MedicLevel, &p.DonorLevel, &p.AdminLevel)
		res.Players = append(res.Players, p)
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

func listPlayersByUID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	loggedIn, perms := isLoggedIn(r.FormValue("token"), r.FormValue("login"))

	if !loggedIn {
		respond(w, -1, "not logged in")
		return
	}

	if perms != 1 && perms != 2 {
		respond(w, -2, "not allowed")
		return
	}

	db, err := sql.Open("mysql", DSN)

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	rows, err := db.Query("SELECT uid, pid, name, cash, bankacc, coplevel, mediclevel, donorlevel, adminlevel FROM players WHERE pid = ?", r.FormValue("uid"))

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	var p Player
	var res ListPlayersResponse

	for rows.Next() {
		rows.Scan(&p.UID, &p.Pid, &p.Name, &p.Cash, &p.Bank, &p.CopLevel, &p.MedicLevel, &p.DonorLevel, &p.AdminLevel)
		res.Players = append(res.Players, p)
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

func playerByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	loggedIn, perms := isLoggedIn(r.FormValue("token"), r.FormValue("login"))

	if !loggedIn {
		respond(w, -1, "not logged in")
		return
	}

	if perms != 1 && perms != 2 {
		respond(w, -2, "not allowed")
		return
	}

	db, err := sql.Open("mysql", DSN)

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	rows, err := db.Query("SELECT uid, pid, name, cash, bankacc, coplevel, mediclevel, donorlevel, adminlevel FROM players u WHERE pid = ?", ps.ByName("idPlayer"))

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	var p Player
	var res ListPlayersResponse

	for rows.Next() {
		rows.Scan(&p.UID, &p.Pid, &p.Name, &p.Cash, &p.Bank, &p.CopLevel, &p.MedicLevel, &p.DonorLevel, &p.AdminLevel)
		res.Players = append(res.Players, p)
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

func changeCash(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	loggedIn, perms := isLoggedIn(r.FormValue("token"), r.FormValue("login"))

	if !loggedIn {
		respond(w, -1, "not logged in")
		return
	}

	if perms != 1 && perms != 2 {
		respond(w, -2, "not allowed")
		return
	}

	db, err := sql.Open("mysql", DSN)
	defer db.Close()

	playerID := r.FormValue("playerID")
	cash := r.FormValue("cash")

	_, err = db.Exec("UPDATE players SET cash = ? WHERE pid = ?", cash, playerID)

	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	action := "Changed cash to value : " + cash

	_, err = db.Exec("INSERT INTO logs (player, action, admin) VALUES (?, ?, ?)", playerID, action, r.FormValue("login"))

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

func changeBank(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	loggedIn, perms := isLoggedIn(r.FormValue("token"), r.FormValue("login"))

	if !loggedIn {
		respond(w, -1, "not logged in")
		return
	}

	if perms != 1 && perms != 2 {
		respond(w, -2, "not allowed")
		return
	}

	db, err := sql.Open("mysql", DSN)
	defer db.Close()

	playerID := r.FormValue("playerID")
	bank := r.FormValue("bank")

	_, err = db.Exec("UPDATE players SET bankacc = ? WHERE pid = ?", bank, playerID)

	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	action := "Changed bank to value : " + bank

	_, err = db.Exec("INSERT INTO logs (player, action, admin) VALUES (?, ?, ?)", playerID, action, r.FormValue("login"))

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

func changeCop(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	loggedIn, perms := isLoggedIn(r.FormValue("token"), r.FormValue("login"))

	if !loggedIn {
		respond(w, -1, "not logged in")
		return
	}

	if perms != 1 && perms != 2 {
		respond(w, -2, "not allowed")
		return
	}

	db, err := sql.Open("mysql", DSN)
	defer db.Close()

	playerID := r.FormValue("playerID")
	cop := r.FormValue("cop")

	_, err = db.Exec("UPDATE players SET coplevel = ? WHERE pid = ?", cop, playerID)

	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	action := "Changed copLevel to level : " + cop

	_, err = db.Exec("INSERT INTO logs (player, action, admin) VALUES (?, ?, ?)", playerID, action, r.FormValue("login"))

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

func changeMedic(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	loggedIn, perms := isLoggedIn(r.FormValue("token"), r.FormValue("login"))

	if !loggedIn {
		respond(w, -1, "not logged in")
		return
	}

	if perms != 1 && perms != 2 {
		respond(w, -2, "not allowed")
		return
	}

	db, err := sql.Open("mysql", DSN)
	defer db.Close()

	playerID := r.FormValue("playerID")
	medic := r.FormValue("medic")

	_, err = db.Exec("UPDATE players SET mediclevel = ? WHERE pid = ?", medic, playerID)

	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	action := "Changed medicLevel to level : " + medic

	_, err = db.Exec("INSERT INTO logs (player, action, admin) VALUES (?, ?, ?)", playerID, action, r.FormValue("login"))

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

func changeDonor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	loggedIn, perms := isLoggedIn(r.FormValue("token"), r.FormValue("login"))

	if !loggedIn {
		respond(w, -1, "not logged in")
		return
	}

	if perms != 1 && perms != 2 {
		respond(w, -2, "not allowed")
		return
	}

	db, err := sql.Open("mysql", DSN)
	defer db.Close()

	playerID := r.FormValue("playerID")
	donor := r.FormValue("donor")

	_, err = db.Exec("UPDATE players SET donorlevel = ? WHERE pid = ?", donor, playerID)

	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	action := "Changed donorLevel to level : " + donor

	_, err = db.Exec("INSERT INTO logs (player, action, admin) VALUES (?, ?, ?)", playerID, action, r.FormValue("login"))

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

func changeAdmin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	loggedIn, perms := isLoggedIn(r.FormValue("token"), r.FormValue("login"))

	if !loggedIn {
		respond(w, -1, "not logged in")
		return
	}

	if perms != 1 && perms != 2 {
		respond(w, -2, "not allowed")
		return
	}

	db, err := sql.Open("mysql", DSN)
	defer db.Close()

	playerID := r.FormValue("playerID")
	admin := r.FormValue("admin")

	_, err = db.Exec("UPDATE players SET adminlevel = ? WHERE pid = ?", admin, playerID)

	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	action := "Changed adminLevel to level : " + admin

	_, err = db.Exec("INSERT INTO logs (player, action, admin) VALUES (?, ?, ?)", playerID, action, r.FormValue("login"))

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
