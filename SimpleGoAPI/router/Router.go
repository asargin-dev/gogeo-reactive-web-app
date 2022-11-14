package router

import (
	"TarlaIO/handler"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Routes() *httprouter.Router{

	//The main router in app. It handles GET and POST methods

	r:= httprouter.New()

	//Router for homepage
	r.GET("/",handler.Index)
	//Router for Get Chart action
	r.POST("/",handler.Index)
	//Router for selected row that will be created their charts
	r.POST("/check",handler.SelectedItemProcess)

	//Allow server to use source files like favicon,images etc.
	r.ServeFiles("/view/resources/*filepath",http.Dir("view/resources"))

	return r
}
