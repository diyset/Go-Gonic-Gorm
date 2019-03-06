package controller

import (
	"github.com/gin-gonic/gin"
	"my-rest/config"
	"my-rest/dto/request"
	"my-rest/structs"
	"my-rest/utility"
	"net/http"
	"strconv"
	"time"
)

func (idb *InDB) PostOrder(c *gin.Context) {
	config.DBinit()
	tx := idb.DB.Begin()
	var (
		listTestStruct    []request.DataProduct
		testStruct        request.DataProduct
		tblOrder          structs.Order
		tblOrderDetail    structs.OrderDetail
		tblAlamat         structs.Alamat
		tblNasabah        structs.Nasabah
		tblProduct        structs.Product
		messageValidation string
	)

	bodyRequest := new(request.DataListOrder)
	errCatching := c.Bind(bodyRequest)
	if errCatching != nil {
		tx.Close()
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": errCatching.Error(),
		})
		return
	}

	responseBody := gin.H{}
	var responseBodys = []gin.H{}
	//fmt.Println("idGeneratorOrder : " + utility.GeneratorIdOrder(bodyRequest.MetodePengiriman, bodyRequest.IdNasabah))

	tblOrder.IdOrder = utility.GeneratorIdOrder(bodyRequest.MetodePengiriman, bodyRequest.IdNasabah)
	errCatching = idb.DB.Where("id = ?", bodyRequest.IdNasabah).First(&tblNasabah).Error

	if errCatching != nil {
		tx.Close()
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": errCatching.Error(),
		})
		return
	}

	errCatching = idb.DB.Where("id_alamat = ?", bodyRequest.IdAlamat).First(&tblAlamat).Error
	if errCatching != nil {
		tx.Close()
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": errCatching.Error(),
		})
		return
	}

	if (tblAlamat != structs.Alamat{} || tblNasabah != structs.Nasabah{}) && int(tblNasabah.ID) != tblAlamat.IdPerson {
		tx.Close()
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Id Alamat is not match for idPerson",
		})
		return
	}

	var hargaTotalOrder float64 = 0.0
	var quantityAllItem int16 = 0
	indexIdOrderDet := int(time.Time.Unix(time.Now()))
	if len(bodyRequest.Products) > 0 {
		for _, element := range bodyRequest.Products {
			testStruct.IdProduct = element.IdProduct
			testStruct.Quantity = element.Quantity
			responseBodys = append(responseBodys, responseBody)
			listTestStruct = append(listTestStruct, testStruct)
			errCatching = idb.DB.Where("id_product = ?", element.IdProduct).First(&tblProduct).Error
			if errCatching != nil {
				tx.Close()
				c.JSON(http.StatusBadRequest, gin.H{
					"message": errCatching.Error(),
					"success": false,
				})
				return
			}

			quantityOld := tblProduct.Stock
			quantityClean := quantityOld - element.Quantity
			if quantityOld > element.Quantity && quantityOld != 0 {
				tblProduct.Stock = quantityClean
				errCatching = tx.Save(&tblProduct).Error
			} else {
				tx.Close()
				c.JSON(http.StatusBadRequest, gin.H{
					"success": false,
					"message": "idProduct `" + element.IdProduct + "` is not enough stock",
				})
				return
			}

			parsingIntIdProduct, errParsing := strconv.Atoi(element.IdProduct)
			if errParsing != nil {
				messageValidation = "Error Parsing " + errParsing.Error()
				tx.Close()
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Error Parsing " + errParsing.Error(),
					"success": false,
				})
				return
			}

			tblOrderDetail.IdOrderDetail = strconv.Itoa(indexIdOrderDet)
			quantityAllItem = quantityAllItem + int16(element.Quantity)
			tblOrderDetail.IdProduct = parsingIntIdProduct
			tblOrderDetail.Quantity = element.Quantity
			totalHargaPerItem := tblProduct.Price * float64(element.Quantity)
			tblOrderDetail.TotaHargaPerProduct = totalHargaPerItem
			tblOrderDetail.IdOrder = tblOrder.IdOrder
			hargaTotalOrder = hargaTotalOrder + tblOrderDetail.TotaHargaPerProduct

			errCatching = tx.Create(&tblOrderDetail).Error
			indexIdOrderDet = indexIdOrderDet + 1
			if errCatching != nil {
				tx.Rollback()
				tx.Close()
				c.JSON(http.StatusBadRequest, gin.H{
					"success": false,
					"message": errCatching.Error(),
				})
				return
			}

			tblProduct = structs.Product{}
			tblOrderDetail = structs.OrderDetail{}
		}
	}

	tblOrder.TanggalOrder = time.Now()
	tblOrder.TotalHarga = hargaTotalOrder
	tblOrder.IdAlamat = bodyRequest.IdAlamat
	tblOrder.IdNasabah = bodyRequest.IdNasabah
	tblOrder.MetodePengiriman = bodyRequest.MetodePengiriman
	tblOrder.TotalItem = int(quantityAllItem)
	errCatching = tx.Save(&tblOrder).Error

	if errCatching != nil {
		tx.Rollback()
		tx.Close()
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   errCatching.Error(),
			"message": messageValidation,
		})
		return
	}
	//tx.Commit()
	c.JSON(http.StatusOK, gin.H{
		"success":         true,
		"message":         "Pembelian successfully! your idOrder " + tblOrder.IdOrder,
		"totalHargaOrder": hargaTotalOrder,
		"totalItem":       tblOrder.TotalItem,
		"namaNasabah":     structs.GetFullNameNasabah(tblNasabah),
		"namaAlamat":      tblAlamat.NamaJalan,
	})
	return
}
