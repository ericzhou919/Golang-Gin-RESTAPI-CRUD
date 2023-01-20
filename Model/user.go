package models

import (
	"strings"
	"github.com/google/uuid"
	"fmt"
	"encoding/json"
)

var Map = make(map[string]User)
type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
}

func GetAll() string{
	jsonString, _ := json.Marshal(Map)
	return string(jsonString)
}

func GetOne(id string) string{
	if s, found := Map[id]; found {
		jsonString, _ := json.Marshal(s)
		return string(jsonString)
	}else{
		return "not found"
	}
}

func Insert(name string) {
	id := strings.Replace(uuid.New().String(), "-", "", -1)
	data := User{id, name}
	fmt.Print(id)
	fmt.Print(data)
	Map[id]=data
}

func Update(id string,name string) {
	if data, found := Map[id]; found {
		data.Username=name
	}
}

func Delete(id string) {
	if x, found := Map[id]; found {
		fmt.Println(x)  
		delete(Map, id)
	}   
}
