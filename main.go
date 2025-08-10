package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

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

// getListenPaths returns the list of paths to listen on
// Default is "/" if no environment variable is set
func getListenPaths() []string {
	paths := os.Getenv("LISTEN_PATHS")
	if paths == "" {
		return []string{"/"}
	}

	// Split by comma and trim whitespace
	pathList := strings.Split(paths, ",")
	for i, path := range pathList {
		pathList[i] = strings.TrimSpace(path)
	}

	// Filter out empty paths
	var result []string
	for _, path := range pathList {
		if path != "" {
			result = append(result, path)
		}
	}

	// If no valid paths found, return default
	if len(result) == 0 {
		return []string{"/"}
	}

	return result
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Get configured paths
	listenPaths := getListenPaths()

	// Register handlers for each path
	for _, path := range listenPaths {
		r.GET(path, handler)
		r.POST(path, handler)
		log.Info().Str("path", path).Msg("registered handler")
	}

	log.Info().Interface("paths", listenPaths).Msg("server starting with configured paths")
	if err := r.Run(); err != nil {
		panic(err)
	}
}
