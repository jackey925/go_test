package main

import (
	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

//goroutine middleware
func goroutineMiddleWare()  {
	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {
		// create copy to be used inside the goroutine
		cCp := c.Copy()
		go func() {
			// simulate a long task with time.Sleep(). 5 seconds
			time.Sleep(5 * time.Second)

			// note that you are using the copied context "cCp", IMPORTANT
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	r.GET("/long_sync", func(c *gin.Context) {
		// simulate a long task with time.Sleep(). 5 seconds
		time.Sleep(5 * time.Second)

		// since we are NOT using a goroutine, we do not have to copy the context
		log.Println("Done! in path " + c.Request.URL.Path)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run("127.0.0.1:8080")
}

func initRedis()  {
	
}

func main() {
	//goroutineMiddleWare()
	//r := gin.New()
	//r.Use(func(context *gin.Context) {
	//	fmt.Println(context.Request.Method)
	//})
	//gin.SetMode("debug")
	////r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	name,_ := c.GetQuery("name")
	//	fmt.Println(name)
	//	c.String(200,name+"\n")
	//	//c.JSON(200, gin.H{
	//	//	"message": "pong",
	//	//})
	//})
	//router := gin.Default()
	//
	//// Example for binding JSON ({"user": "manu", "password": "123"})
	//router.POST("/loginJSON", func(c *gin.Context) {
	//	var json Login
	//	if err := c.ShouldBindJSON(&json); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//		return
	//	}
	//
	//	if json.User != "manu" || json.Password != "manu" {
	//		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	//		return
	//	}
	//
	//	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	//})
	//
	//
	//router.Run("127.0.0.1:9000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}





