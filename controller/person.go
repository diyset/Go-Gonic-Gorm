package controller

import (
	"my-rest/dto/response"
	"my-rest/structs"
	"my-rest/utility"
	"net/http"
	"time"

	"github.com/bearbin/go-age"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) AddAlamatNasabah(c *gin.Context) {
	var (
		nasabah     structs.Nasabah
		alamat      structs.Alamat
		responseDto response.ResponseData
	)

	id := c.PostForm("idNasabah")
	errNasabah := idb.DB.Where("id = ?", id).First(&nasabah).Error
	if errNasabah != nil {
		responseDto.Status = "error"
		responseDto.Message = errNasabah.Error()
		responseDto.Data = nil
		responseDto.Success = false
		c.JSON(http.StatusInternalServerError, responseDto)
		return
	}
	alamat.IdPerson = int(nasabah.ID)
	alamat.NamaJalan = c.PostForm("namaJalan")
	alamat.No = c.PostForm("no")
	alamat.Rw = c.PostForm("rw")
	alamat.Rt = c.PostForm("rt")
	alamat.Provinsi = c.PostForm("provinsi")

	errAlamat := idb.DB.Create(&alamat).Error
	if errAlamat != nil {
		responseDto.Status = "error"
		responseDto.Message = errAlamat.Error()
		responseDto.Data = nil
		responseDto.Success = false
		c.JSON(http.StatusInternalServerError, responseDto)
		return
	}
	responseDto.Status = "success"
	responseDto.Message = "Added Alamat Success"
	responseDto.Data = gin.H{
		"no":        alamat.No,
		"idAlamat":  alamat.IdAlamat,
		"idPerson":  alamat.IdPerson,
		"rt":        alamat.Rt,
		"rw":        alamat.Rw,
		"namaJalan": alamat.NamaJalan,
	}
	responseDto.Success = true
	c.JSON(http.StatusOK, responseDto)
}

func (idb *InDB) GetAlamatById(c *gin.Context) {
	var (
		alamat      structs.Alamat
		responseDto response.ResponseData
	)
	id := c.PostForm("id")
	err := idb.DB.Where("id = ?", id).First(&alamat).Error
	if err != nil {
		responseDto.Status = "success"
		responseDto.Message = "Not Found"
		responseDto.Data = gin.H{
			"mesasge": "Test Again",
			"stauts":  "okee",
		}
		c.JSON(http.StatusOK, responseDto)
	} else {
		responseDto.Status = "success"
		responseDto.Message = "Found!!"
		responseDto.Data = gin.H{
			"idAlamat": alamat.IdAlamat,
			"rt":       alamat.Rt,
			"rw":       alamat.Rw,
			"no":       alamat.No,
			"provinsi": alamat.Provinsi,
		}
		c.JSON(http.StatusOK, responseDto)
	}
}

func (idb *InDB) GetPerson(c *gin.Context) {
	var (
		person    structs.Nasabah
		personDto response.ResponseData
		result    gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&person).Error
	if err != nil {
		result = gin.H{
			"success": false,
			"message": gin.H{
				"result": err.Error(),
			},
			"data": nil,
		}
		c.JSON(http.StatusInternalServerError, result)
	} else {
		personDto.Status = "success"
		personDto.Message = "Found!"
		personDto.Data = gin.H{
			"idNasabah":    person.ID,
			"firstName":    person.First_Name,
			"lastName":     person.Last_Name,
			"email":        person.Email,
			"jenisKelamin": person.JenisKelamin,
			"tanggalLahir": person.TanggalLahir,
			"isAdult":      person.IsAdult,
		}
		c.JSON(http.StatusOK, personDto)
	}
}

func (idb *InDB) GetPersons(c *gin.Context) {
	var (
		alamat        []structs.Alamat
		persons       []structs.Nasabah
		result        gin.H
		results       []gin.H
		resultAlamats []gin.H
		getAlamat     gin.H
		responseDto   response.ResponseDataList
	)

	err := idb.DB.Find(&persons).Error
	if err != nil {
		result = gin.H{
			utility.ResponseStatus():  "error",
			utility.ResponseMessage(): err.Error(),
			utility.ResponseSuccess(): false,
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	if len(persons) <= 0 {
		result = gin.H{
			"success": false,
			"message": gin.H{
				"result": nil,
				"count":  0,
			},
			"data": nil,
		}
	} else {
		responseDto.Message = "get All Nasabah and Alamat"
		responseDto.Success = true
		responseDto.Status = "sucess"
		for _, element := range persons {
			errAlamat := idb.DB.Where("id_person = ?", element.ID).Find(&alamat).Error
			if errAlamat != nil {
				panic("error alamat find by id")
				return
			}
			if len(alamat) > 0 {
				for _, element := range alamat {
					getAlamat = gin.H{
						"idAlamat":  element.IdAlamat,
						"namaJalan": element.NamaJalan,
						"provinsi":  element.Provinsi,
						"no":        element.No,
						"rt":        element.Rt,
						"rw":        element.Rw,
					}
					resultAlamats = append(resultAlamats, getAlamat)
				}
			}
			result = gin.H{
				"idNasabah":    element.ID,
				"fullName":     element.First_Name + " " + element.Last_Name,
				"firstName":    element.First_Name,
				"lastName":     element.Last_Name,
				"email":        element.Email,
				"jenisKelamin": element.JenisKelamin,
				"tanggalLahir": element.TanggalLahir.Format("02-01-2006"),
				"isAdult":      element.IsAdult,
				"listAlamat":   resultAlamats,
			}
			results = append(results, result)
			resultAlamats = nil
		}
	}
	responseDto.Data = results
	c.JSON(http.StatusOK, responseDto)
}

func (idb *InDB) CreateNasabah(c *gin.Context) {
	var (
		responseDto response.ResponseData
		nasabah     structs.Nasabah
	)
	firstName := c.PostForm("firstName")
	lastName := c.PostForm("lastName")
	email := c.PostForm("email")
	jenisKelamin := c.PostForm("jenisKelamin")
	tanggalLahir := c.PostForm("tanggalLahir")

	if !utility.ValidateEmailFormat(email) {
		c.JSON(http.StatusOK, gin.H{
			utility.ResponseStatus():  "error",
			utility.ResponseMessage(): "Error Format Email",
			utility.ResponseSuccess(): false,
			utility.ResponseData():    nil,
		})
		return
	}
	if !utility.ValidateDateFormat(tanggalLahir) {
		c.JSON(http.StatusOK, gin.H{
			utility.ResponseStatus():  "error	",
			utility.ResponseMessage(): "Error Format TanggalLahir (dd-MM-yyyy)",
			utility.ResponseSuccess(): false,
			utility.ResponseData():    nil,
		})
		return
	}
	isAdult := false
	nasabah.First_Name = firstName
	nasabah.Last_Name = lastName
	nasabah.Email = email
	nasabah.JenisKelamin = jenisKelamin
	parsingTime, err := time.Parse("02-01-2006", tanggalLahir)
	getAge := age.AgeAt(parsingTime, time.Now())
	if getAge > 18 {
		isAdult = true
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			utility.ResponseStatus():  "error",
			utility.ResponseMessage(): err.Error(),
			utility.ResponseData():    nil,
		})
		panic("error parsing" + err.Error())
		return
	}
	nasabah.IsAdult = isAdult
	nasabah.TanggalLahir = parsingTime
	errCreateNasabah := idb.DB.Create(&nasabah).Error
	if errCreateNasabah != nil {
		responseDto.Data = nil
		responseDto.Status = "success"
		responseDto.Message = "error : " + errCreateNasabah.Error()
		c.JSON(http.StatusOK, responseDto)
		panic("errCreateNasabah : " + errCreateNasabah.Error())
		return
	}
	responseDto.Status = "success"
	responseDto.Message = "Success Create Nasabah"
	responseDto.Data = gin.H{
		"idNasabah":       nasabah.ID,
		"fullNameNasabah": nasabah.First_Name + " " + nasabah.Last_Name,
		"yourAge":         age.AgeAt(parsingTime, time.Now()),
	}
	c.JSON(http.StatusOK, responseDto)
}
