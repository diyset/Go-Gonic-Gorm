package controller

import (
	"github.com/gin-gonic/gin"
	"my-rest/dto/response"
	"my-rest/structs"
	"net/http"
	"strconv"
)

func (idb *InDB) CreateProduct(c *gin.Context) {
	var (
		responseDto response.ResponseData
		tblProduct  structs.Product
	)
	nameProduct := c.PostForm("nameProduct")
	idCategory := c.PostForm("idCategory")
	price := c.PostForm("price")
	stock := c.PostForm("stock")

	tblProduct.Price, _ = strconv.ParseFloat(price, 64)
	tblProduct.Stock, _ = strconv.Atoi(stock)
	tblProduct.NameProduct = nameProduct
	tblProduct.IdCategory, _ = strconv.Atoi(idCategory)

	errCreateProduct := idb.DB.Create(&tblProduct).Error
	if errCreateProduct != nil {
		responseDto.Data = nil
		responseDto.Status = "error"
		responseDto.Message = "error " + errCreateProduct.Error()
		responseDto.Success = false
		c.JSON(http.StatusInternalServerError, responseDto)
		return
	}
	responseDto.Status = "success"
	responseDto.Message = "Create Product Success"
	responseDto.Data = gin.H{
		"idProduct":   tblProduct.IdProduct,
		"idCategory":  tblProduct.IdCategory,
		"nameProduct": tblProduct.NameProduct,
		"stock":       tblProduct.Stock,
		"price":       tblProduct.Price,
	}
	responseDto.Success = true
	c.JSON(http.StatusOK, responseDto)
}

func (idb *InDB) GetAllProduct(c *gin.Context) {
	var (
		tblProducts []structs.Product
		responseDto response.ResponseDataList
		result      gin.H
		results     []gin.H
	)
	err := idb.DB.Find(&tblProducts).Error
	if err != nil {
		responseDto.Success = false
		responseDto.Message = "error : " + err.Error()
		responseDto.Status = "error"
		responseDto.Data = nil
		c.JSON(http.StatusInternalServerError, responseDto)
		return
	}
	if (len(tblProducts) > 0) {
		responseDto.Message = "success"
		responseDto.Status = "success"
		responseDto.Success = true
		for _, element := range tblProducts {
			result = gin.H{
				"idProduct":   element.IdProduct,
				"nameProduct": element.NameProduct,
				"stock":       element.Stock,
				"price":       element.Price,
				"idCategory":  element.IdCategory,
			}
			results = append(results, result)
		}
		responseDto.Data = results
		c.JSON(http.StatusOK, responseDto)
	}
}
