package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/redis.v4"
)

var (
	params struct {
		Host  string
		Moods []string
	}
	rHost string
	rc    *redis.Client
)

func init() {
	host, err := os.Hostname()
	if err == nil {
		params.Host = host
	} else {
		params.Host = "<unknown>"
	}
	params.Moods = []string{"happy", "meh", "angry"}

	rHost = os.Getenv("HAPPY_REDIS")
	if len(rHost) == 0 {
		rHost = "localhost"
	}
	rc = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:6379", rHost),
	})
	_, err = rc.Ping().Result()
	if err != nil {
		log.Fatalf("error: can't connect to redis on %s - %s", rHost, err)
	}

	logDir := os.Getenv("HAPPY_LOG_DIR")
	if len(logDir) == 0 {
		logDir = "logs"
	}
	if err = os.MkdirAll(logDir, 0740); err != nil {
		log.Fatalf("error: can't create log dir - %s", err)
	}

	logFile := fmt.Sprintf("%s/happy_access.log", logDir)
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalf("error: can't open log file - %s", err)
	}
	gin.DefaultWriter = io.MultiWriter(gin.DefaultWriter, file)
}

func homeHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", params)
}

func getCounterHandler(ctx *gin.Context) {
	counter := ctx.Params.ByName("name")
	sval, err := rc.Get(counter).Result()
	if err == redis.Nil {
		sval = "0"
	} else if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	value, err := strconv.Atoi(sval)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"counter": counter,
		"value":   value,
	})
}

func postCounterHandler(ctx *gin.Context) {
	counter := ctx.Params.ByName("name")
	value, err := rc.Incr(counter).Result()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"counter": counter,
		"value":   value,
	})
}

func main() {
	rtr := gin.Default()
	rtr.SetHTMLTemplate(template.Must(template.ParseFiles("index.html")))
	rtr.GET("/", homeHandler)
	rtr.GET("/counter/:name", getCounterHandler)
	rtr.POST("/counter/:name", postCounterHandler)
	rtr.Static("/static", "static")
	rtr.Run(":8080")
}
