package main

import(
	_"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

func HelloPage(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message":"Welcome to bgops, please visit https://xxbandy.github.io!",
	})
}

func main(){
	r:=gin.Default()
	v1:=r.Group("/v1")
	{
		v1.GET("/hello", HelloPage)
		v1.GET("/hello/:name", func(c *gin.Context){
			name:=c.Param("name")
			c.String(http.StatusOK, "Hello %s", name)
		})
		// 匹配的url格式:  /hello2?firstname=Jane&lastname=Doe
		// curl "localhost:8000/v1/hello2?firstname=Jane&lastname=Doe"
		v1.GET("/hello2", func(c *gin.Context){
			// firstname := c.DefaultQuery("firstname", "Guest")
			firstname:=c.Query("firstname")
        	lastname:= c.Query("lastname") // 是 c.Request.URL.Query().Get("lastname") 的简写

        	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
		})
		v1.GET("/hello3/:name", func(c *gin.Context){
			name:=c.Param("name")
			firstname:=c.Query("firstname")
        	lastname:= c.Query("lastname") // 是 c.Request.URL.Query().Get("lastname") 的简写

        	c.String(http.StatusOK, "Hello %s %s %s", name, firstname, lastname)
		})
		
		v1.GET("/line", func(c *gin.Context){
			// 注意:在前后端分离过程中，需要注意跨域问题，因此需要设置请求头
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			legendData:=[]string{"周一", "周二", "周三", "周四", "周五", "周六", "周日"}
			xAxisData:=[]int{120, 240, rand.Intn(500), rand.Intn(500), 150, 230, 180}
			c.JSON(200, gin.H{
				"legend_data":legendData,
				"xAxisData":xAxisData,
			})
		})
		//curl -X POST localhost:8000/v1/hello4 -H "Content-Type:application/x-www-form-urlencoded" -d "message=hello&nick=rsj217" | python -m json.tool % Total % Received % Xferd 
		v1.POST("/hello4", func(c *gin.Context){
			message := c.PostForm("message") 
        	nick := c.DefaultPostForm("nick", "anonymous") 
        	c.JSON(http.StatusOK, gin.H{ 
            	"status": gin.H{ 
                "status_code": http.StatusOK, 
                "status": "ok", 
            	}, 
            	"message": message, 
            	"nick": nick,         
        	}) 
		})

		//定义默认路由
		r.NoRoute(func(c *gin.Context){
			c.JSON(http.StatusNotFound, gin.H{
				"status":404,
				"error": "404, page not exits!",
			})
		})
		r.Run(":8000")


	}
}
