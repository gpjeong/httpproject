package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"httpproject/internal/datastore"
	"httpproject/util/logger"
	"net/http"
)

func ApiTest() {
	r := gin.Default()
	tableStruct := &OjtInfo{}

	// GET 요청
	r.GET("/get", func(c *gin.Context) {

		inputParam := c.Query("name")

		if inputParam == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "input 파라미터가 필요합니다"})
		} else {
			query := make(map[string]interface{}, 0)
			query["name = ?"] = inputParam

			dataCount, err := datastore.DBService().GetCountCustomQuery(tableStruct, query)
			if err != nil {
				logger.Log.Error().Msgf("DB Get Name error: %s", err.Error())
			}

			if dataCount != 0 {
				c.JSON(200, gin.H{
					"message": fmt.Sprintf("%s 조회에 성공하였습니다", inputParam),
				})
			} else {
				c.JSON(500, gin.H{
					"message": fmt.Sprintf("%s 조회에 실패하였습니다", inputParam),
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
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("DELETE request for resource with ID %s", id),
		})
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
		inputParam := c.Query("input")

		if inputParam == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "input 파라미터가 필요합니다"})
		} else {
			query := make(map[string]interface{}, 0)
			query["name = ?"] = inputParam

			dataCount, err := datastore.DBService().GetCountCustomQuery(tableStruct, query)
			if err != nil {
				logger.Log.Error().Msgf("Vrops Alerts get id, provider error: %s", err.Error())
			}

			if dataCount != 0 {
				c.JSON(200, gin.H{
					"message": fmt.Sprintf("%s 조회에 성공하였습니다", inputParam),
				})
			} else {
				c.JSON(500, gin.H{
					"message": fmt.Sprintf("%s 조회에 실패하였습니다", inputParam),
				})
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "POST request for resource",
		})
	})

	// 서버 시작
	r.Run(":8080")
}
