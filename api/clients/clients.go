package reports

import (
	"encoding/json"
	"net/http"
	functions "travel-n-expenses/api/global-functions"
	Models "travel-n-expenses/api/models"
	DB "travel-n-expenses/config"

	"github.com/gorilla/mux"
)

func AddClientsHandler(router *mux.Router) {
	router.HandleFunc("/clients", getClients).Methods("GET")
	router.HandleFunc("/clients", newClient).Methods("POST")
}

// getClients : GET CLIENTS (returns clients records within a JSON)
func getClients(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")

	query := `SELECT idClient, name, address
			FROM clients
			WHERE isDeleted = 0
			ORDER BY name;`

	rows, err := DB.MySQL.Query(query)
	if functions.CheckError(http.StatusInternalServerError, err, &res) {
		return
	}

	//Slce that storages all the reports retrieved from the db
	var data []Models.Client

	for rows.Next() {
		var clientData Models.Client

		//Assignation of the values from the row to the client struct
		err = rows.Scan(&clientData.IDClient, &clientData.Name, &clientData.Address)

		if functions.CheckError(http.StatusInternalServerError, err, &res) {
			return
		}

		//The client is added to the slice
		data = append(data, clientData)
	}

	rows.Close()

	//Map to format the response json
	response := make(map[string]interface{})
	response["data"] = &data
	response["ok"] = true
	response["count"] = len(data)

	json.NewEncoder(res).Encode(response)
}

func newClient(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")

	//JSON request transformation to GO structure
	decoder := json.NewDecoder(req.Body)
	var client *Models.Client
	err := decoder.Decode(&client)

	if functions.CheckError(http.StatusInternalServerError, err, &res) {
		return
	}

	query := `INSERT INTO clients (name, address, isDeleted) VALUES (?,?,0)`

	stmt, err := DB.MySQL.Prepare(query)
	if functions.CheckError(http.StatusInternalServerError, err, &res) {
		return
	}

	result, err := stmt.Exec(client.Name, client.Address)
	if functions.CheckError(http.StatusInternalServerError, err, &res) {
		return
	}

	//Get the ID of the new record
	IDreport, err := result.LastInsertId()

	if functions.CheckError(http.StatusInternalServerError, err, &res) {
		return
	}

	client.IDClient = &IDreport

	//Map to format the response json
	response := make(map[string]interface{})
	response["client"] = &client
	response["ok"] = true

	//Return the new report data in JSON format .
	json.NewEncoder(res).Encode(response)

}
