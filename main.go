package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func handler(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error().Err(err).Msg("read body failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	path := c.FullPath()
	method := c.Request.Method
	header := c.Request.Header
	bodyStr := string(body)
	resp := gin.H{
		"path":   path,
		"method": method,
		"header": header,
		"body":   bodyStr,
	}
	log.Info().Str("body", bodyStr).
		Str("method", method).Str("path", path).Interface("header", header).Msg("request")
	c.JSON(http.StatusOK, resp)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", handler)
	r.POST("/", handler)
	if err := r.Run(); err != nil {
		panic(err)
	}
}
