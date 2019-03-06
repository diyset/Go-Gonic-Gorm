package controller

import (
	"github.com/gin-gonic/gin"
	"my-rest/dto/response"
	"my-rest/structs"
	"my-rest/utility"
	"net/http"
	"strconv"
)

func (idb *InDB) DropProductInWarehouse(c *gin.Context) {
	var (
		responseDto  response.ResponseData
		tblWarehouse structs.Warehouse
		tblproduct   structs.Product
	)
	idProduct := c.PostForm("idProduct")
	quantity := c.PostForm("quantity")
	tanggalKirim := c.PostForm("tanggalKirim")


	errProduct := idb.DB.Where("id_product= ?", idProduct).First(&tblproduct).Error
	if errProduct != nil {
		panic("Id Product Not Found!")
		return
	}
	quantityOld := tblproduct.Stock
	quantityNew, _ := strconv.Atoi(quantity)
	quantityClean := quantityOld + quantityNew

	tblWarehouse.IdProduct, _ = strconv.Atoi(idProduct)
	tblWarehouse.Quantity, _ = strconv.Atoi(quantity)
	tblWarehouse.TanggalKirim, _ = utility.DateFormatMyApp(tanggalKirim)
	errWarehouse := idb.DB.Create(&tblWarehouse).Error
	if errWarehouse != nil {
		panic("Error Warehouse : " + errWarehouse.Error())
		return
	}
	tblproduct.Stock = quantityClean
	// Update TblProduct
	errProduct = idb.DB.Save(&tblproduct).Error
	if errProduct != nil {
		panic("Error TblProduct : " + errProduct.Error())
		return
	}
	responseDto.Data = gin.H{
		"idWarehouse":  tblWarehouse.IdWarehouse,
		"idProduct":    tblWarehouse.IdProduct,
		"quantity":     tblWarehouse.Quantity,
		"tanggalKirim": utility.DateFormatMyLayout(tblWarehouse.TanggalKirim),
		"stockTemp":    tblproduct.Stock,
	}
	responseDto.Success = true
	responseDto.Status = "success"
	responseDto.Message = "success barang masuk Product : " + tblproduct.NameProduct
	c.JSON(http.StatusOK, responseDto)
}

func (idb *InDB) GetAllWarehouse(c *gin.Context) {
	var (
		responseDto  response.ResponseDataList
		tblWarehouse []structs.Warehouse
		result       gin.H
		results      []gin.H
	)

	err := idb.DB.Find(&tblWarehouse).Error
	if err != nil {
		panic("Error : " + err.Error())
		return
	}

	if len(tblWarehouse) > 0 {
		responseDto.Message = "success"
		responseDto.Success = true
		responseDto.Status = "success"
		for _, element := range tblWarehouse {
			result = gin.H{
				"idWarehouse":  element.IdWarehouse,
				"idProduct":    element.IdProduct,
				"quantity":     element.Quantity,
				"tanggalKirim": utility.DateFormatMyLayout(element.TanggalKirim),
			}
			results = append(results, result)
		}
		responseDto.Data = results
		c.JSON(http.StatusOK, responseDto)

	}
}
