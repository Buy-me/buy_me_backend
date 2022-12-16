package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	router.Use(CORS())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run()
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, Origin, Cache-Control, X-Requested-With")
		// c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		// c.Writer.Header().Set("Access-Control-Allow-Methods", "PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")

		fmt.Println(c.Request.Method)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// import (
// 	"encoding/json"
// 	"fmt"
// 	"time"
// )

// type Model struct {
// 	Name   string `json:"name"`
// 	Millis int64  `json:"lastModified"`
// }

// func (m Model) Lastmodified() time.Time {
// 	return time.Unix(0, m.Millis*int64(time.Millisecond))
// }

// func main() {
// 	modelVar := Model{}
// 	err := json.Unmarshal([]byte(`{ "name" : "hello", "lastModified" : 564483600000 }`), &modelVar)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(modelVar.Lastmodified())
// }
