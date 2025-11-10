package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {

	path := os.Args[1];

	dataStore := NewFileDataStore(path);

	routeHandler := RouteHandler{
		dataStore: dataStore,
	};

	router := gin.Default()
    router.GET("/albums", routeHandler.getAlbums)
    router.GET("/albums/:id", routeHandler.getAlbum)
    router.POST("/albums", routeHandler.postAlbums)	

    router.Run("localhost:8080")
}