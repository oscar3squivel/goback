package reports

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"text/template"
	functions "travel-n-expenses/api/global-functions"
	Models "travel-n-expenses/api/models"
	F "travel-n-expenses/api/pdf"
	DB "travel-n-expenses/config"

	"github.com/gorilla/mux"
)

// ADDING ALL ROUTES FOR REPORTS MODULES
func AddReportsHandler(router *mux.Router) {
	router.HandleFunc("/report/create", insertNewReport).Methods("POST")
	router.HandleFunc("/report", getReports).Methods("GET")
	router.HandleFunc("/report/user/{IDReport:[0-9]+}", getReportByID).Methods("GET")
	router.HandleFunc("/report/employee/{IDEmployee:[0-9]+}", getReportsByEmployeeID).Methods("GET")
	router.HandleFunc("/report/{IDReport:[0-9]+}/pdf", pdfReport).Methods("GET")
	router.HandleFunc("/report/{IDReport:[0-9]+}", reportByID).Methods("GET")
	router.HandleFunc("/report/delete", deleteReport).Methods("PUT")
	router.HandleFunc("/report/recover", recoverReport).Methods("PUT")
}

//Function that return a report by IdEmployee and IdReport as parameters
func getReportByID(res http.ResponseWriter, req *http.Request) {
	//Getting url parameters
	params := mux.Vars(req)
	//var reports []Models.Report
	var report Models.Report
	//var errorHandler Models.ErrorHandler
	var employee Models.Employee
	var client Models.Client
	res.Header().Add("Content-type", "applicaction/json")
	//Creating query
	query := `SELECT r.*,e.*,c.name, c.address
			 FROM reports as r 
			 LEFT JOIN clients as c on r.idClient = c.idClient 
			 INNER JOIN employees as e on r.idEmployee = e.idEmployee 
			 WHERE r.idReport=?  and r.isDeleted=0`
	//create the sql statement with its parameters
	reportSQL, err := DB.MySQL.Prepare(query)
	if functions.CheckError(http.StatusInternalServerError, err, &res) != false {
		return
	}
	defer reportSQL.Close()
	//Retrieve data from SQL Statement
	err = reportSQL.QueryRow(params["IDReport"]).Scan(
		//Maping report data
		&report.IDReport, &report.OriginPlace, &report.DestinationPlace, &report.DepartureDate,
		&report.ReturnDate, &report.TeamName, &report.IDLeader, &report.Purpose, &report.IsSponsored,
		&report.ClientPurpose, &report.SpecialRequirements, &report.OperationDate,
		&report.OperationArrangement, &report.StatusChangeDate, &report.Comments,
		&report.ReportStatus, &report.IDClient, &report.IDEmployee, &report.IsInternational,
		&report.IsDeleted,
		//Maping employee data
		&employee.IDEmployee, &employee.Email, &employee.Name, &employee.PhoneNumber,
		&employee.BirthDate, &employee.Nationality, &employee.EmergencyNumber, &employee.Password,
		&employee.Role, &employee.IsDeleted,
		//Maping client data
		&client.Name, &client.Address,
	)
	report.Employee = &employee
	report.Client = &client
	//In case the SQL Instruction returns empty
	if functions.CheckError(int64(http.StatusBadRequest), err, &res) != false {
		return
	}
	json.NewEncoder(res).Encode(report)
}

//Functions for the html template
func Deref(status *string) string { return *status }
func DerefBool(val *bool) bool    { return *val }

func reportByID(res http.ResponseWriter, req *http.Request) {
	//Getting url parameters
	params := mux.Vars(req)
	//var reports []Models.Report
	var report Models.Report
	//var errorHandler Models.ErrorHandler
	var employee Models.Employee
	var client Models.Client //Creating query
	query := `SELECT r.*,e.*,c.name, c.address FROM reports as r 
		 	LEFT JOIN clients as c on r.idClient = c.idClient 
			INNER JOIN employees as e on r.idEmployee = e.idEmployee 
			WHERE r.IDReport=?`
	//create the sql statement with its parameters
	reportSQL, err := DB.MySQL.Prepare(query)
	if functions.CheckError(http.StatusInternalServerError, err, &res) != false {
		return
	}
	defer reportSQL.Close()
	//Retrieve data from SQL Statement
	err = reportSQL.QueryRow(params["IDReport"]).Scan(
		//Maping report data
		&report.IDReport, &report.OriginPlace, &report.DestinationPlace, &report.DepartureDate,
		&report.ReturnDate, &report.TeamName, &report.IDLeader, &report.Purpose, &report.IsSponsored,
		&report.ClientPurpose, &report.SpecialRequirements, &report.OperationDate,
		&report.OperationArrangement, &report.StatusChangeDate, &report.Comments,
		&report.ReportStatus, &report.IDClient, &report.IDEmployee, &report.IsInternational,
		&report.IsDeleted,
		//Maping employee data
		&employee.IDEmployee, &employee.Email, &employee.Name, &employee.PhoneNumber,
		&employee.BirthDate, &employee.Nationality, &employee.EmergencyNumber, &employee.Password,
		&employee.Role, &employee.IsDeleted,
		//Maping client data
		&client.Name, &client.Address,
	)
	report.Employee = &employee
	report.Client = &client
	//In case the SQL Instruction returns empty
	if functions.CheckError(int64(http.StatusBadRequest), err, &res) != false {
		return
	}

	//Show the info on the template
	//t, err := template.ParseFiles("api/pdf/form.html")
	t, err := template.New("form.html").Funcs(template.FuncMap{
		"Deref":     Deref,
		"DerefBool": DerefBool,
	}).ParseFiles("api/pdf/form.html")

	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	t.Execute(res, report)

}
func pdfReport(res http.ResponseWriter, req *http.Request) {
	//Get the value of id
	vars := mux.Vars(req)
	rID, _ := (vars["IDReport"])
	/*if err != nil { panic(err.Error()) }*/
	F.PDF(rID)
	fmt.Fprint(res, "PDF Generated")
}

// insertNewReport : POST function. Creates a new report in the database
func insertNewReport(res http.ResponseWriter, req *http.Request) {
	//JSON request transformation to GO structure
	decoder := json.NewDecoder(req.Body)
	var report *Models.Report
	err := decoder.Decode(&report)
	if functions.CheckError(http.StatusInternalServerError, err, &res) {
		return
	}

	query := `INSERT INTO reports 
			( originPlace, destinationPlace, departureDate, returnDate, teamName, idLeader, 
			purpose, isSponsored, clientPurpose, specialRequirements, statusChangeDate, 
			reportStatus, idClient, idEmployee, isInternational, isDeleted) 
			VALUES (?,?,?,?,?,?,?,?,?,?,NOW(),'P',?,?,?,0)`
	stmt, err := DB.MySQL.Prepare(query)
	if functions.CheckError(http.StatusInternalServerError, err, &res) {
		return
	}
	rest, err := stmt.Exec(report.OriginPlace, report.DestinationPlace, report.DepartureDate,
		report.ReturnDate, report.TeamName, report.IDLeader, report.Purpose, report.IsSponsored,
		report.ClientPurpose, report.SpecialRequirements, report.IDClient, report.IDEmployee,
		report.IsInternational)

	if functions.CheckError(http.StatusInternalServerError, err, &res) {
		return
	}
	//Get the ID of the new record
	IDreport, err := rest.LastInsertId()
	if functions.CheckError(http.StatusInternalServerError, err, &res) {
		return
	}
	report.IDReport = &IDreport
	//Return the new report data in JSON format
	json.NewEncoder(res).Encode(report)
}

// getReports : GET REQUEST (returns reports records within a JSON)
func getReports(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	query := `SELECT r.idReport, r.idEmployee, r.idLeader, r.originPlace, 
	r.destinationPlace, r.departureDate, r.returnDate, IFNULL(r.teamName, '-') teamName, 
	r.isSponsored, r.reportStatus, e.idEmployee, e.name FROM reports r 
	INNER JOIN employees e ON r.idEmployee = e.idEmployee WHERE r.isDeleted = 0 
	GROUP BY r.idReport ORDER BY r.idReport;`
	rows, err := DB.MySQL.Query(query)
	if functions.CheckError(http.StatusInternalServerError, err, &w) {
		return
	}
	//Slce that storages all the reports retrieved from the db
	var data []Models.Report
	for rows.Next() {
		var reportData Models.Report
		var employeeData Models.Employee
		//Assignation of the values from the row to the structures (report and employee) values
		err = rows.Scan(&reportData.IDReport, &reportData.IDEmployee, &reportData.IDLeader,
			&reportData.OriginPlace, &reportData.DestinationPlace, &reportData.DepartureDate,
			&reportData.ReturnDate, &reportData.TeamName, &reportData.IsSponsored,
			&reportData.ReportStatus, &employeeData.IDEmployee, &employeeData.Name)
		if functions.CheckError(http.StatusInternalServerError, err, &w) {
			return
		}
		reportData.Employee = &employeeData
		//The report is added to the slice
		data = append(data, reportData)
	}
	rows.Close()
	//Map to format the response json
	response := make(map[string]interface{})
	response["data"] = &data
	response["ok"] = true
	response["count"] = len(data)
	json.NewEncoder(w).Encode(response)
}

// getReports : GET REQUEST (returns reports records within a JSON)
func getReportsByEmployeeID(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")

	//Getting url parameters
	params := mux.Vars(req)
	idEmployee := params["IDEmployee"]

	query := `SELECT r.idReport, r.idEmployee, r.idLeader, r.originPlace, 
	r.destinationPlace, r.departureDate, r.returnDate, purpose, IFNULL(r.teamName, '-') teamName, 
	r.isSponsored, r.reportStatus, e.idEmployee, e.name FROM reports r 
	INNER JOIN employees e ON r.idEmployee = e.idEmployee WHERE r.idEmployee = ? AND r.isDeleted = 0 
	GROUP BY r.idReport ORDER BY r.idReport;`
	rows, err := DB.MySQL.Query(query, idEmployee)
	if functions.CheckError(http.StatusInternalServerError, err, &res) {
		return
	}
	//Slce that storages all the reports retrieved from the db
	var data []Models.Report
	var employeeData Models.Employee
	for rows.Next() {
		var reportData Models.Report
		//Assignation of the values from the row to the structures (report and employee) values
		err = rows.Scan(&reportData.IDReport, &reportData.IDEmployee, &reportData.IDLeader,
			&reportData.OriginPlace, &reportData.DestinationPlace, &reportData.DepartureDate,
			&reportData.ReturnDate, &reportData.Purpose, &reportData.TeamName, &reportData.IsSponsored,
			&reportData.ReportStatus, &employeeData.IDEmployee, &employeeData.Name)
		if functions.CheckError(http.StatusInternalServerError, err, &res) {
			return
		}
		//The report is added to the slice
		data = append(data, reportData)
	}
	employeeData.Reports = &data
	rows.Close()
	//Map to format the response json
	response := make(map[string]interface{})
	response["data"] = &employeeData
	response["ok"] = true
	json.NewEncoder(res).Encode(response)
}

func deleteReport(res http.ResponseWriter, req *http.Request) {
	changeReportDeletedStatus(res, req, true)
}

func recoverReport(res http.ResponseWriter, req *http.Request) {
	changeReportDeletedStatus(res, req, false)
}

func changeReportDeletedStatus(res http.ResponseWriter, req *http.Request, delete bool) {
	res.Header().Add("Content-Type", "application/json")
	type ReqStruct struct {
		Ids []int64 `json:"ids"`
	}
	decoder := json.NewDecoder(req.Body)
	var reqS ReqStruct
	err := decoder.Decode(&reqS)
	if functions.CheckError(http.StatusInternalServerError, err, &res) {
		return
	}

	if len(reqS.Ids) == 0 {
		//Return the new report data in JSON format
		response := make(map[string]interface{})
		response["ok"] = false
		response["message"] = "Not a single ID was received"
		json.NewEncoder(res).Encode(response)
		return
	}

	//Interface where the ids will be saved to make only when query per request
	args := make([]interface{}, len(reqS.Ids))
	for i, id := range reqS.Ids {
		args[i] = id
	}

	//Query declaration
	query := ""

	//If delete
	if delete {
		query = `UPDATE reports SET isDeleted = 1 WHERE idReport IN (?` + strings.Repeat(",?", len(args)-1) + `)`
	} else { //if recover
		query = `UPDATE reports SET isDeleted = 0 WHERE idReport IN (?` + strings.Repeat(",?", len(args)-1) + `)`
	}

	stmt, err := DB.MySQL.Prepare(query)
	if functions.CheckError(http.StatusInternalServerError, err, &res) {
		return
	}
	rest, err := stmt.Exec(args...)

	if functions.CheckError(http.StatusInternalServerError, err, &res) {
		return
	}

	//Get the num of rows affected by the query
	rowsAffected, err := rest.RowsAffected()
	if functions.CheckError(http.StatusInternalServerError, err, &res) {
		return
	}

	//Return the new report data in JSON format
	response := make(map[string]interface{})
	response["ok"] = true
	response["rowsAffected"] = rowsAffected
	json.NewEncoder(res).Encode(response)
}
