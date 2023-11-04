package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"httpproject/internal/config"
	"httpproject/internal/datastore"
	"httpproject/util/logger"
	"net/http"
)

type RequestData struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Balance string `json:"balance"`
}

type PutRequestData struct {
	Name    string `json:"name"`
	Balance string `json:"balance"`
}

type DbData struct {
	Id      string
	Name    string
	balance string
	test    string
}

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
				c.JSON(200, gin.H{
					"message": fmt.Sprintf("%s 조회에 실패하였습니다", nameParam),
				})
			}
		}

	})

	// PUT 요청
	r.PUT("/put", func(c *gin.Context) {
		idParam := c.Query("id")
		var nameParam string
		var balanceParam string

		if idParam == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "input 파라미터의 입력 형식이 잘못되었거나 필드의 데이터가 없습니다"})
		} else {
			var putRequestData PutRequestData

			// ShouldBindJSON은 omitempty 태그의 영향을 받지 않음
			if err := c.ShouldBindJSON(&putRequestData); err == nil {
				if putRequestData.Name == "" {
					nameParam = ""
				}
				if putRequestData.Balance == "" {
					balanceParam = ""
				}
				if putRequestData.Name != "" {
					nameParam = putRequestData.Name
				}
				if putRequestData.Balance != "" {
					balanceParam = putRequestData.Balance
				}
				dataList := make([]OjtInfo, 0)
				data := OjtInfo{
					Id:      idParam,
					Name:    nameParam,
					Balance: balanceParam,
				}
				dataList = append(dataList, data)

				query := make(map[string]interface{}, 0)
				query["id = ?"] = idParam

				dataCount, err := datastore.DBService().GetCountCustomQuery(tableStruct, query)
				if err != nil {
					logger.Log.Error().Msgf("DB Get Name error :", err.Error())
				}
				if dataCount != 0 {
					_, err := datastore.DBService().UpsertData(dataList)
					if err != nil {
						logger.Log.Error().Msgf("Update Data Error :", err.Error())
					}
					c.JSON(200, gin.H{
						"message": fmt.Sprintf("Id : %s, name : %s 업데이트 성공하였습니다", idParam, nameParam),
					})
				} else {
					_, err := datastore.DBService().CreateData(dataList)
					if err != nil {
						logger.Log.Error().Msgf("DB Add Data error :", err.Error())
					}
					c.JSON(200, gin.H{
						"message": fmt.Sprintf("Id : %s, name : %s 데이터를 추가하였습니다", idParam, nameParam),
					})
				}

			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "input body 입력 형식이 잘못되었습니다"})
			}
		}
	})

	// DELETE 요청
	r.DELETE("/delete", func(c *gin.Context) {
		var requestData RequestData

		err := c.ShouldBindJSON(&requestData)
		if err != nil || requestData.Id == "" || requestData.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "input body의 입력 형식이 잘못되었거나 필드의 데이터가 없습니다"})
		} else {
			idParam := requestData.Id
			nameParam := requestData.Name
			balanceParam := requestData.Balance

			dataList := make([]OjtInfo, 0)
			data := OjtInfo{
				Id:      idParam,
				Name:    nameParam,
				Balance: balanceParam,
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
				c.JSON(200, gin.H{
					"message": fmt.Sprintf("Id : %s, name : %s 삭제 실패하였습니다", idParam, nameParam),
				})
			}
		}

	})

	// PATCH 요청
	r.PATCH("/patch", func(c *gin.Context) {
		idParam := c.Query("id")
		var nameParam string
		var balanceParam string

		if idParam == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "input 파라미터의 입력 형식이 잘못되었거나 필드의 데이터가 없습니다"})
		} else {
			var putRequestData PutRequestData

			// ShouldBindJSON은 omitempty 태그의 영향을 받지 않음
			if err := c.ShouldBindJSON(&putRequestData); err == nil {
				if putRequestData.Name == "" {
					nameParam = ""
				}
				if putRequestData.Balance == "" {
					balanceParam = ""
				}
				if putRequestData.Name != "" {
					nameParam = putRequestData.Name
				}
				if putRequestData.Balance != "" {
					balanceParam = putRequestData.Balance
				}
				dataList := make([]OjtInfo, 0)
				data := OjtInfo{
					Id:      idParam,
					Name:    nameParam,
					Balance: balanceParam,
				}
				dataList = append(dataList, data)

				query := make(map[string]interface{}, 0)
				query["id = ?"] = idParam

				dataCount, err := datastore.DBService().GetCountCustomQuery(tableStruct, query)
				if err != nil {
					logger.Log.Error().Msgf("DB Get Name error :", err.Error())
				}
				if dataCount != 0 {
					if nameParam == "" || balanceParam == "" {
						data, err := datastore.DBService().GetData(idParam, tableStruct)
						if err != nil {
							logger.Log.Error().Msgf("DB Get Data error :", err.Error())
						}
						if ptr, ok := data.(*OjtInfo); ok {
							if nameParam == "" {
								nameParam = ptr.Name
							}
							if balanceParam == "" {
								balanceParam = ptr.Balance
							}
							dataList := make([]OjtInfo, 0)
							data := OjtInfo{
								Id:      idParam,
								Name:    nameParam,
								Balance: balanceParam,
							}
							dataList = append(dataList, data)

							_, err := datastore.DBService().UpsertData(dataList)
							if err != nil {
								logger.Log.Error().Msgf("Update Data Error :", err.Error())
							}
							c.JSON(200, gin.H{
								"message": fmt.Sprintf("Id : %s 데이터 업데이트 성공했습니다", idParam),
							})

						}

					}

				} else {
					c.JSON(200, gin.H{
						"message": fmt.Sprintf("Id : %s 데이터 업데이트 실패했습니다", idParam),
					})
				}

			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "input body 입력 형식이 잘못되었습니다"})
			}
		}
	})

	// POST 요청
	r.POST("/post", func(c *gin.Context) {
		var requestData RequestData

		err := c.ShouldBindJSON(&requestData)
		if err != nil || requestData.Id == "" || requestData.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "input body의 입력 형식이 잘못되었거나 필드의 데이터가 없습니다"})
		} else {
			idParam := requestData.Id
			nameParam := requestData.Name
			balanceParam := requestData.Balance

			dataList := make([]OjtInfo, 0)
			data := OjtInfo{
				Id:      idParam,
				Name:    nameParam,
				Balance: balanceParam,
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
					logger.Log.Error().Msgf("Post Data Failed :", err.Error())
				}
				c.JSON(200, gin.H{
					"message": fmt.Sprintf("Id : %s, name : %s 생성에 성공하였습니다", idParam, nameParam),
				})
			} else {
				c.JSON(200, gin.H{
					"message": fmt.Sprintf("Id : %s, name : %s 생성에 실패하였습니다", idParam, nameParam),
				})
			}
		}

	})

	// 서버 시작
	r.Run(":" + config.ServerConfig.ApiInfo.ApiPort)
}
