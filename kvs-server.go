package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xxorde/kvs"
	"net/http"
	"strconv"
	"time"
)

var (
	store kvs.Kvs
)

func main() {
	store = *kvs.NewKvs()
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", Index)
	r.GET("/get/:key", Get)
	r.POST("/put", Put)
	r.POST("/delete", Delete)
	r.GET("/json", JSON)
	r.GET("/yaml", Yaml)
	r.Run(":8080")
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "kvs-server"})
}

func Get(c *gin.Context) {
	key := c.Params.ByName("key")
	value := store.Get(key)
	content := gin.H{"key": key, "value": value}
	c.JSON(http.StatusOK, content)
}

func Put(c *gin.Context) {
	key := c.PostForm("key")
	value := c.PostForm("value")
	ttl, err := strconv.ParseInt(c.PostForm("ttl"), 10, 64)
	if ttl == 0 || err != nil {
		store.Put(key, value)
	} else {
		store.PutTTL(key, value, time.Unix(ttl, 0))
	}
	content := gin.H{"put": "ok", "ttl": ttl}
	c.JSON(http.StatusOK, content)
}

func Delete(c *gin.Context) {
	key := c.PostForm("key")
	store.Delete(key)
	content := gin.H{"delete": "ok"}
	c.JSON(http.StatusOK, content)
}

func JSON(c *gin.Context) {
	c.JSON(http.StatusOK, store.JSON())
}

func Yaml(c *gin.Context) {
	c.JSON(http.StatusOK, store.Yaml())
}
