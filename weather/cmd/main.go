package main

import (
	"example/weather/internal/server"

	"github.com/gin-gonic/gin"
)

func main() {

	//*
	router := gin.Default()
	//router.GET("/temp", getTemp)
	router.GET("/temperature/:city", server.GetCityLocation)
	//fmt.Printf("1:%s/n2:%s", server.CityLocation[0], server.CityLocation[1])
	//server.GetTemp()
	//router.GET("/temperature/:latitude/:longitude", server.GetTemp)
	//bson.D{{"city_ascii", "Mississauga"}}

	//http.HandleFunc("/view/", server.MakeHandler(server.ViewHandler))
	server.CreatePage()

	router.Run("localhost:8080")
	//*/
	//getTemp()
}
