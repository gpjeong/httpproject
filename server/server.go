package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"httpproject/internal/config"
	"httpproject/internal/datastore"
	"httpproject/util/logger"
	"net/http"
)

func ApiTest() {
	r := gin.Default()
	tableStruct := &OjtInfo{}

	// GET 요청
	r.GET("/get", func(c *gin.Context) {

		nameParam := c.Query("name")

		if nameParam == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "input 파라미터가 필요합니다"})
		} else {
			query := make(map[string]interface{}, 0)
			query["name = ?"] = nameParam

			dataCount, err := datastore.DBService().GetCountCustomQuery(tableStruct, query)
			if err != nil {
				logger.Log.Error().Msgf("DB Get Name error :", err.Error())
			}

			if dataCount != 0 {
				c.JSON(200, gin.H{
					"message": fmt.Sprintf("%s 조회에 성공하였습니다", nameParam),
				})
			} else {
				c.JSON(500, gin.H{
					"message": fmt.Sprintf("%s 조회에 실패하였습니다", nameParam),
				})
			}
		}

	})

	// PUT 요청
	r.PUT("/put", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("PUT request for resource with ID %s", id),
		})
	})

	// DELETE 요청
	r.DELETE("/delete", func(c *gin.Context) {
		idParam := c.DefaultQuery("id", "")
		nameParam := c.DefaultQuery("name", "")

		if nameParam == "" || idParam == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "input 파라미터가 필요합니다"})
		} else {
			dataList := make([]OjtInfo, 0)
			data := OjtInfo{
				Id:   idParam,
				Name: nameParam,
			}
			dataList = append(dataList, data)

			query := make(map[string]interface{}, 0)
			query["name = ?"] = nameParam
			query["id = ?"] = idParam

			dataCount, err := datastore.DBService().GetCountCustomQuery(tableStruct, query)
			if err != nil {
				logger.Log.Error().Msgf("DB Get Name error :", err.Error())
			}

			if dataCount != 0 {
				_, err := datastore.DBService().DeleteCustomQuery(query, data)
				if err != nil {
					logger.Log.Error().Msgf("Delete Data Failed :", err.Error())
				}
				c.JSON(200, gin.H{
					"message": fmt.Sprintf("Id : %s, name : %s 삭제 성공하였습니다", idParam, nameParam),
				})
			} else {
				c.JSON(500, gin.H{
					"message": fmt.Sprintf("Id : %s, name : %s 삭제 실패하였습니다", idParam, nameParam),
				})
			}
		}

	})

	// PATCH 요청
	r.PATCH("/patch", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("PATCH request for resource with ID %s", id),
		})
	})

	// POST 요청
	r.POST("/post", func(c *gin.Context) {
		idParam := c.Query("id")
		nameParam := c.Query("name")

		if idParam == "" || nameParam == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "input 파라미터가 필요합니다"})
		} else {
			dataList := make([]OjtInfo, 0)
			data := OjtInfo{
				Id:   idParam,
				Name: nameParam,
			}
			dataList = append(dataList, data)

			query := make(map[string]interface{}, 0)
			query["id = ?"] = idParam

			dataCount, err := datastore.DBService().GetCountCustomQuery(tableStruct, query)
			if err != nil {
				logger.Log.Error().Msgf("DB Get Name error:", err.Error())
			}

			if dataCount == 0 {
				_, err := datastore.DBService().CreateData(dataList)
				if err != nil {
					logger.Log.Error().Msgf("Create Post Data Failed :", err.Error())
				}
				c.JSON(200, gin.H{
					"message": fmt.Sprintf("Id : %s, name : %s 생성에 성공하였습니다", idParam, nameParam),
				})
			} else {
				c.JSON(500, gin.H{
					"message": fmt.Sprintf("Id : %s, name : %s 생성에 실패하였습니다", idParam, nameParam),
				})
			}
		}

	})

	// 서버 시작
	r.Run(":" + config.ServerConfig.ApiInfo.ApiPort)
}
