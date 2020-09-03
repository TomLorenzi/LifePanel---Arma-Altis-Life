package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/julienschmidt/httprouter"
)

// Log type logs in database
type Log struct {
	ID     int       `json:"id"`
	Player string    `json:"player"`
	Action string    `json:"action"`
	Admin  string    `json:"admin"`
	Date   time.Time `json:"date"`
}

// ListLogsResponse response api data
type ListLogsResponse struct {
	APIResponse
	Logs []Log `json:"logs"`
}

func listLogs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	rows, err := db.Query("SELECT id, player, action, admin, date FROM logs ORDER BY id DESC LIMIT 50")

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	var l Log
	var res ListLogsResponse

	for rows.Next() {
		rows.Scan(&l.ID, &l.Player, &l.Action, &l.Admin, &l.Date)
		res.Logs = append(res.Logs, l)
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

func listLogsByPlayer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	rows, err := db.Query("SELECT id, player, action, admin, date FROM logs WHERE player = ?", r.FormValue("player"))

	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	var l Log
	var res ListLogsResponse

	for rows.Next() {
		rows.Scan(&l.ID, &l.Player, &l.Action, &l.Admin, &l.Date)
		res.Logs = append(res.Logs, l)
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
