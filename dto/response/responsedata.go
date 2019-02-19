package response

import (
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Status  string
	Message string
	Data    gin.H
	Success bool
}

type ResponseDataList struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Data    []gin.H `json:"data"`
	Success bool    `json:"success"`
}
