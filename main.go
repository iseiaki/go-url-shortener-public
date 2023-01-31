package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/speps/go-hashids"
)

type URL struct {
	ID        uint   `gorm:"primary_key"`
	Code   string `gorm:"unique;not null"`
	Target string `gorm:"not null"`
}

type Visit struct {
	ID      int `gorm:"primary_key"`
	Visits  int
}

var rateLimiter = make(map[string]int)


func main() {
	
	db, err := gorm.Open("mysql", "root:123@tcp(localhost:3306)/urls?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	db.AutoMigrate(&URL{})
	router := gin.Default()

	//take care of api
	router.POST("/api", apipost)
	router.GET("/api", apiget)

	router.LoadHTMLGlob("templates/*")

	router.GET("/:code", func(c *gin.Context) {
		db.AutoMigrate(&Visit{})
		var visit Visit
		db.Where("id = ?", 1).FirstOrCreate(&visit)
		visit.Visits++
		db.Save(&visit)	
		var url URL
		db.Where("code = ?", c.Param("code")).First(&url)
		c.Redirect(302, url.Target)
	})

	router.GET("/", func(c *gin.Context) {
		var visit Visit
		db.Where("id = ?", 1).FirstOrCreate(&visit)
		var count int 
		db.Table("urls").Count(&count) 
		c.HTML(200, "index.html", gin.H{"alltime": count, "visits": visit.Visits})
	})	
	
	
	router.POST("/shorten", func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := rateLimiter[ip]
		if limiter >= 5 {
			c.HTML(429, "exceeded.html", gin.H{})
			return
		}
		rateLimiter[ip]++
		target := c.PostForm("target")
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

		db.Create(&URL{Target: target, Code: code})

		c.HTML(200, "shorten.html", gin.H{"code": code})

		go func() {
			time.Sleep(time.Minute)
			rateLimiter[ip]--
		}()
	})

    router.Run(":8080")
}
