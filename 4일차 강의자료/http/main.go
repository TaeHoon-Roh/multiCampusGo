package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	WebServer()
}

type MenData struct {
	Test string
}

func WebServer() {
	g := gin.Default()
	g.LoadHTMLGlob("c:\\workspace_go\\templates\\html\\*.html")
	g.GET("/", indexPage)
	g.POST("/", indexPostPage)
	g.GET("/test", testPage)
	g.Run(":8082")

}

func indexPage(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}
func indexPostPage(c *gin.Context) {
	// decoder := json.NewDecoder(c.Request.Body)
	// jsonBytes, err := json.Marshal(decoder)
	// if err != nil {
	// 	panic(err)
	// }

	// jsonString := string(jsonBytes)
	// fmt.Println(c.Request.Body)
	// fmt.Println(jsonString)

	// fmt.Println("Post Call")
	myCity := make([]int, 10)
	c.JSON(http.StatusOK, gin.H{
		"City": myCity,
		"Data": "check_wd",
	})
}
func testPage(c *gin.Context) {
	c.HTML(200, "test.html", gin.H{})
}
