package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../config"
	"../model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Org struct {
	Org_id   int    `json : "org_id"`
	Org_name string `json : "org_name"`
}

type Cms struct {
	Cms_id   int    `json : "cms_id"`
	Cms_name string `json : "cms_name"`
}

type Status struct {
	Stat_id   int    `json : "stat_id"`
	Stat_name string `json : "stat_name"`
}

type App struct {
	App_id   int    `json : "app_id"`
	App_name string `json : "app_name"`
}

type Intgs struct {
	I_id     int    `json : "i_id"`
	I_org    string `json : "i_org"`
	I_cms    string `json : "i_cms"`
	I_app    string `json : "i_app"`
	I_status string `json : "i_status"`
	App_url  string `json : "app_url"`
	Comment  string `json : "comment"`
}

//Sql connection

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:admin@123@tcp(127.0.0.1:3306)/intg_stat")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/org", GetAllOrg).Methods("GET")
	router.HandleFunc("/status", GetAllStat).Methods("GET")
	router.HandleFunc("/cms", GetAllCms).Methods("GET")
	router.HandleFunc("/app", GetAllApp).Methods("GET")
	router.HandleFunc("/intgs", GetAllIntgs).Methods("GET")
	router.HandleFunc("/intgs", PostAllIntgs).Methods("POST")

	fmt.Println("Server running on port : 8090")
	http.ListenAndServe(":8090", router)
}

// Access the Data
func GetAllOrg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := db.Query("Select * FROM org")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var org_id []Org
	for result.Next() {
		var org Org
		err := result.Scan(&org.Org_id, &org.Org_name)
		if err != nil {
			panic(err.Error())
		}
		org_id = append(org_id, org)

	}
	json.NewEncoder(w).Encode(org_id)
}

func GetAllStat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := db.Query("Select * FROM status")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var stat_id []Status
	for result.Next() {
		var status Status
		err := result.Scan(&status.Stat_id, &status.Stat_name)
		if err != nil {
			panic(err.Error())
		}
		stat_id = append(stat_id, status)

	}
	json.NewEncoder(w).Encode(stat_id)
}

func GetAllApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := db.Query("Select * FROM app")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var app_id []App
	for result.Next() {
		var app App
		err := result.Scan(&app.App_id, &app.App_name)
		if err != nil {
			panic(err.Error())
		}
		app_id = append(app_id, app)

	}
	json.NewEncoder(w).Encode(app_id)
}

func GetAllCms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := db.Query("Select * FROM cms")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var cms_id []Cms
	for result.Next() {
		var cms Cms
		err := result.Scan(&cms.Cms_id, &cms.Cms_name)
		if err != nil {
			panic(err.Error())
		}
		cms_id = append(cms_id, cms)

	}
	json.NewEncoder(w).Encode(cms_id)
}

func GetAllIntgs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := db.Query("Select * FROM intgs")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var i_id []Intgs
	for result.Next() {
		var intgs Intgs
		err := result.Scan(&intgs.I_id, &intgs.I_org, &intgs.I_cms, &intgs.I_app, &intgs.I_status, &intgs.App_url, &intgs.Comment)
		if err != nil {
			panic(err.Error())
		}
		i_id = append(i_id, intgs)

	}
	json.NewEncoder(w).Encode(i_id)
}

/*
	func PostAllIntgs(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		stmt, err := db.Prepare("INSERT INTO intgs(i_org) VALUE(?)")
		if err != nil {
			panic(err.Error())
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err.Error())
		}
		keyVal := make(map[string]string)
		json.Unmarshal(body, &keyVal)
		i_org := keyVal["i_org"]
		_, err = stmt.Exec(i_org)
		if err != nil {
			panic(err.Error())
		}

}
*/
func PostAllIntgs(w http.ResponseWriter, r *http.Request) {
	var response model.Response

	db := config.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}
	i_id := r.FormValue("i_id")
	i_name := r.FormValue("i_name")
	i_cms := r.FormValue("i_cms")
	i_app := r.FormValue("i_app")
	i_status := r.FormValue("i_status")
	app_url := r.FormValue("app_url")
	comment := r.FormValue("comment")
	_, err = db.Exec("INSERT INTO employee(name, city) VALUE(?, ?, ?, ?, ?, ?, ?, ?)", i_id, i_name, i_cms, i_app, i_status, app_url, comment)

	if err != nil {
		log.Print(err)
		return
	}
	response.Status = 200
	response.Message = "Insert data successfully"
	fmt.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

/*func PostAllIntgs(w http.ResponseWriter, r *http.Request) {
	i_id := r.FormValue("i_id")
	i_org := r.FormValue("i_org")
	i_cms := r.FormValue("i_cms")
	i_app := r.FormValue("i_app")
	i_status := r.FormValue("i_status")
	app_url := r.FormValue("app_url")
	comment := r.FormValue("comment")

	w.Header().Set("Content-Type", "application/json")
	result, err := db.Query("INSERT INTO INTGS(i_id,i_org,i_cms,i_app,i_status,app_url,comment) VALUE($1,$2,$3,$4,$5,$6,$7);", i_id, i_org, i_cms, i_app, i_status, app_url, comment)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var ii_id []Intgs
	for result.Next() {
		var intgs Intgs
		err := result.Scan(&intgs.I_id, &intgs.I_org, &intgs.I_cms, &intgs.I_app, &intgs.I_status, &intgs.App_url, &intgs.Comment)
		if err != nil {
			panic(err.Error())
		}
		ii_id = append(ii_id, intgs)

	}
	json.NewEncoder(w).Encode(ii_id)
}

func PostAllIntgs(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var intgs Intgs
	err = json.Unmarshal(b, &intgs)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(intgs)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
*/
