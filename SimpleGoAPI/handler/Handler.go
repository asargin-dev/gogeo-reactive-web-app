package handler

import (
	"TarlaIO/users"
	"TarlaIO/utils"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"text/template"
)

// Users The Data models to adapt json type API's response to Go struct
var Users []users.User
var Posts []users.Post

// Index Page - Main Handler Function
func Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	// Collect all html templates to serve
	view, err := template.ParseFiles(utils.Include()...)
	if err != nil {
		fmt.Println(err)
	}

	//Getting User API Data
	responseUsers, _ := http.Get("https://jsonplaceholder.typicode.com/users")
	userData, _ := ioutil.ReadAll(responseUsers.Body)
	//Set API Data to Go structs
	err = json.Unmarshal(userData, &Users)
	if err != nil {
		fmt.Println(err)
	}

	//Getting Posts API Data
	responsePosts, _ := http.Get("https://jsonplaceholder.typicode.com/posts")
	postData, _ := ioutil.ReadAll(responsePosts.Body)
	//Set API Data to Go structs
	err = json.Unmarshal(postData, &Posts)

	//Serve HTML and Send Data to HTML
	err = view.ExecuteTemplate(w, "index", Users)

}

// SelectedItemProcess This handler function creates map and pie charts with selected table items
func SelectedItemProcess(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	//Getting values from html form that is posted
	formValue := r.FormValue("id")
	//Converting string to integer array to send multiple choice reference
	request := utils.StringToIntArray(formValue)

	//Create map for selected parameters
	utils.GeoChartCreator{}.CreateMap(Users, request)

	//This block calculates post count of each user
	counter := 0
	for i, user := range Users {
		for _, posts := range Posts {
			if user.Id == posts.UserID {
				counter++
			}
		}
		Users[i].PostNumber = counter
		counter = 0
	}
	//

	//Create pie chart for selected parameters
	utils.PieChartGenerator{}.CreatePieChart(Users, request)
}

