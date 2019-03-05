package login

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

const clientID = "51a2e850-dbcf-42d8-a10c-5e88fda828a9"
const clientSecret = "vsKTY59]$tgesfJAUC807?*"

const dire = "user.read"

//http%3A%2F%2Flocalhost%3A3001%2Foauth%2Fredirect
const ruri = "http://localhost:3001/Login"

// ADDING ROUTE FOR LOGIN MODULE
func AddLoginHandler(router *mux.Router) {
	router.HandleFunc("/Login", login)
}

func login(w http.ResponseWriter, r *http.Request) {
	//fs := http.FileServer(http.Dir("../../public/travel-n-expenses-client/src/components/login/login.js"))
	//http.Handle("/", fs)

	//crear httpcliente para hacer un http request externo despues
	httpClient := http.Client{}

	//Sacar el valor de code del query del url
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(os.Stdout, "no se pudo hacer el parse del query: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	//code := r.FormValue("code")

	//hacer un rquest http para llamar al oauth endpoint y tener el token de acceso
	//"https://login.microsoftonline.com/common/oauth2/token?grant_type=authorization_code&client_id=%s&code=%s&redirect_uri=http%3A%2F%2Flocalhost%3A3001%2Foauth%2Fredirect&client_secret=%s"
	//"https://login.microsoftonline.com/common/oauth2/token?grant_type=%s&client_id=%s&code=%s&redirect_uri=http%3A%2F%2Flocalhost%3A3001%2Foauth%2Fredirect&client_secret=%s
	//https://login.microsoftonline.com/common/oauth2/token?grant_type=%s&client_id=%s&code=%s&redirect_uri=http://localhost:3001/oauth/redirect&client_secret=%s
	//https://login.microsoftonline.com/common/oauth2/token?grant_type=authorization_code&client_id=%s&code=%s&redirect_uri=%s&client_secret=%s

	form := url.Values{}
	form.Set("client_id", clientID)
	form.Set("scope", dire)
	form.Set("grant_type", "authorization_code")
	form.Set("code", r.FormValue("code"))
	form.Set("redirect_uri", ruri)
	//form.Set("resource", "https://graph.windows.net")
	//form.Set("resource", "https://graph.windows.net")
	form.Set("client_secret", "vsKTY59]$tgesfJAUC807?*")
	//Imprimir codigo de autorizacion
	println(r.FormValue("code"))

	//reqURL := fmt.Sprintf("https://login.microsoftonline.com/e436ac59-887c-458a-b2d1-7d6759a4ec75/oauth2/token?grant_type=authorization_code&client_id=%s&code=%s&redirect_uri=%s&client_secret=%s", clientID, code, ruri, clientSecret)

	//sin version 2
	//https://login.microsoftonline.com/e436ac59-887c-458a-b2d1-7d6759a4ec75/oauth2/token
	reqURL := "https://login.microsoftonline.com/e436ac59-887c-458a-b2d1-7d6759a4ec75/oauth2/v2.0/token"
	//tokenReq, err := http.NewRequest(http.MethodPost, reqURL, strings.NewReader(form.Encode()))

	//post a url
	//req, err := http.NewRequest(http.MethodPost, reqURL, nil)

	req, err := http.NewRequest(http.MethodPost, reqURL, strings.NewReader(form.Encode()))
	//println(req)

	if err != nil {
		fmt.Fprintf(os.Stdout, "no se pudo crear un HTTP request: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	//header con una respuesta com json
	//post: "accept","application/json"
	//application/x-www-form-urlencoded
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	//mandar la http request
	res, err := httpClient.Do(req)

	if err != nil {
		fmt.Fprintf(os.Stdout, "no HTTP request: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer res.Body.Close()

	//json.Unmarshal(buf, &spockets)

	//parsear el body request en la estructura 0authresponse
	var t OAuthAccessResponse

	b, err := ioutil.ReadAll(res.Body)
	println(string(b))

	if err := json.Unmarshal(b, &t); err != nil {
		fmt.Fprintf(os.Stdout, "no se pudo parsear el JSON response: %v", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	/*
		if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
			fmt.Fprintf(os.Stdout, "no se pudo parsear el JSON response: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}*/

	//redireccionar a welcom con el token de acceso
	///welcome.html?access_token=
	//w.Header().Set("Location", "../../public/travel-n-expenses-client/src/components/login/welcome.html?access_token="+t.AccessToken)
	//w.Header().Set("Location", "/welcome.html?access_token="+t.AccessToken)
	// w.Header().Set("http://localhost:3000", "/Reports")
	// w.WriteHeader(http.StatusFound)
	http.Redirect(w, r, "http://localhost:3000/Reports", http.StatusSeeOther)
}

//OAuthAccessResponse es la represemntacoon de un oauth
type OAuthAccessResponse struct {
	//`json:"access_token"`
	AccessToken string `json:"access_token"`
}
