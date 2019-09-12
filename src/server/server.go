package server

import (
	"../config"
	"../task"
	"../utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Gitdata struct {
	Repository struct {
		Name      string `json:"name"`
		Full_name string `json:"full_name"`
	}
}

func StartService(address string, port string) error {
	http.HandleFunc("/", index)
	http.HandleFunc("/auto_build", autoBuild)

	utils.Log2file(fmt.Sprintf("service starting... %s:%s", address, port))
	return http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	utils.Log2file(string(r.URL.Host))
	fmt.Fprintln(w, "{\"code\":200, \"description\":\"service running...\"}")
}

func autoBuild(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("%#v\n", r.URL.RequestURI())
	// fmt.Printf("%#v\n", r.Header.Get("X-Hub-Signature"))

	if (r.Method == "post" || r.Method == "POST") && r.URL.RequestURI() == config.GetURL() {
		if r.Header.Get("x-github-event") == "push" {
			bodyContent, _ := ioutil.ReadAll(r.Body)
			r.Body.Close()

			var data Gitdata
			json.Unmarshal([]byte(bodyContent), &data)
			// fmt.Printf("3:%s\n", data.Repository.Full_name)
			serviceStr := "QQ"
			key := strings.ToLower(data.Repository.Name)
			switch key {
			case "ckiptagger":
				serviceStr = "hi"
			case "pitaya-backend-service-gateway":
				serviceStr = "gateway"
			case "pitaya-backend-service-member":
				serviceStr = "member"
			case "pitaya-backend-service-keycloak-admin":
				serviceStr = "keycloakadminservice"
			}
			signature := r.Header.Get("X-Hub-Signature")
			if utils.VerifySignature(signature, string(bodyContent), config.GetSecret()) {
				fmt.Fprintln(w, "{\"code\":200, \"description\":\"OK\"}")
				utils.Log2file("repo Signature success")
				task.AddNewTask(string(bodyContent), serviceStr)
			} else {
				utils.Log2file("repo Signature error")
				fmt.Fprintln(w, "{\"code\":200, \"error\":\"Signature error\"}")
			}
		} else {
			fmt.Fprintln(w, "{\"code\":200, \"error\":\"Unmatch x-github-event\"}")
		}
	} else {
		fmt.Fprintln(w, "{\"code\":200, \"error\":\"Error Method or unknow request url\"}")
	}
}
