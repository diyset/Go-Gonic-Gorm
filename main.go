package main

import (
	"fmt"
	"my-rest/config"
	"my-rest/controller"
	myAuth "my-rest/controller/auth"
	"my-rest/dto/request"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var PortG = ""

func main() {
	envResource := godotenv.Load()
	if envResource != nil {
		fmt.Println(envResource)
	}
	PORT := os.Getenv("PORT_MACHINE")
	db := config.DBinit()
	PortG = PORT
	inDb := &controller.InDB{DB: db}
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/login", myAuth.LoginHandler)
		api.POST("/test", indexHandler)
		api.GET("/hello/:name", helloHandler)
		api.GET("/person/:id", myAuth.Auth, inDb.GetPerson)
		api.GET("/persons", inDb.GetPersons)
		api.POST("/person", inDb.CreateNasabah)
		api.POST("/get/alamat", inDb.GetAlamatById)
		api.POST("/add/alamat", inDb.AddAlamatNasabah)

		api.POST("/add/product", inDb.CreateProduct)
		api.GET("/get/products", inDb.GetAllProduct)

		api.POST("/add/warehouse/product", inDb.DropProductInWarehouse)
		api.GET("/get/warehouses", inDb.GetAllWarehouse)

		api.POST("/post/order", inDb.PostOrder)
	}

	router.Run(":" + PORT)
}

func indexHandler(c *gin.Context) {
	bodyRequest := new(request.DataList)
	err := c.Bind(bodyRequest)
	if err != nil {
		panic("Error " + err.Error())
	}
	//testJson := c.PostForm("list")
	fmt.Println(len(bodyRequest.List))
	responseBody := gin.H{}
	var responseBodys = []gin.H{}
	for _, element := range bodyRequest.List {
		responseBody = gin.H{
			"test": element.Test,
		}
		responseBodys = append(responseBodys, responseBody)
	}
	c.JSON(http.StatusOK, gin.H{
		"title":      "Tutorial Golang and GORM",
		"timestamp":  time.Now().UnixNano() / int64(time.Millisecond),
		"authorized": "Dian Setiyadi",
		"port":       PortG,
		"testJson":   responseBodys,
	})
}

func helloHandler(c *gin.Context) {
	paramPreffix := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"yourname": paramPreffix,
	})
}
