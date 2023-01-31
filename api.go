package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/gin-gonic/gin"
	"github.com/speps/go-hashids"
)


func apipost(c *gin.Context){
	db, err := gorm.Open("mysql", "root:123@tcp(localhost:3306)/urls?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	type Request struct {
		URL string `json:"url"`
	}

	var request Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Access request data
	url := request.URL
	ip := c.ClientIP()
	limiter := rateLimiter[ip]
	if limiter >= 5 {
		c.JSON(http.StatusOK, gin.H{"limit": "You have reached the maximum limit, please try again later"})
		return
	}
	rateLimiter[ip]++
	rand.Seed(time.Now().UnixNano())
	var code string
	for {
		hd := hashids.NewData()
		h, _ := hashids.NewWithData(hd)
		id := rand.Intn(100000)
		code, _ = h.Encode([]int{int(id)})
		var count int
		db.Model(&URL{}).Where("code = ?", code).Count(&count)
		if count == 0 {
			break
		}
	}
	db.Create(&URL{Target: url, Code: code})
	c.JSON(200, gin.H{"Newurl": "http://localhost:8080/" + code})
	
}

func apiget(c *gin.Context){
		c.JSON(200, gin.H{"Message": "Check out API documentation here - https://github.com/iseiaki/go-url-changer-public/blob/maina/README.md#api-and-how-to-use-it"})
}