package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

type User struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

func getUsers(){
	url := "http://localhost:8080"	
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url+"/users/", nil)
	res, _ := client.Do(req)
	if res.StatusCode == 200 {
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
	}
	defer res.Body.Close()
}

func createUser(){
	url := "http://localhost:8080"
	user := User{"Jane","Doe"}
	userJSON, _ := json.Marshal(user)
	userBytes := bytes.NewBuffer(userJSON)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url+"/users", userBytes)
	res, _ := client.Do(req)
	if res.StatusCode == 200 {
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
	}
	defer res.Body.Close()
}

func updateUser(){
	url := "http://localhost:8080"
	user := User{"John","Doe"}
	userJSON, _ := json.Marshal(user)
	userBytes := bytes.NewBuffer(userJSON)

	client := &http.Client{}
	req, _ := http.NewRequest("PUT", url+"/users/1", userBytes)
	res, _ := client.Do(req)	
	if res.StatusCode == 200 {
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
	}
	defer res.Body.Close()
}

func getSingleUser(){
	url := "http://localhost:8080"
	res, _ := http.Get(url+"/users/1")
	if res.StatusCode == 200 {
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
	}
	defer res.Body.Close()
}

func deleteUser(){
	url := "http://localhost:8080"
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", url+"/users/1", nil)
	res, _ := client.Do(req)
	if res.StatusCode == 200 {
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
	}
	defer res.Body.Close()
}

func main(){
	createUser()
	getUsers()
	getSingleUser()
	updateUser()
	deleteUser()
}