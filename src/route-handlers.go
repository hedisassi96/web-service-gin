package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouteHandler struct {
	dataStore DataStore
}

// for GET /albums
func (routeHandler RouteHandler) getAlbums(c *gin.Context) {
	albums, err := routeHandler.dataStore.getAllAlbums()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, albums)
	}
	c.IndentedJSON(http.StatusOK, albums)
}

// for GET /albums/:id
func (routeHandler RouteHandler) getAlbum(c *gin.Context) {
	id := c.Param("id")
	album, err := routeHandler.dataStore.getAlbumById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, album)
	}
	c.IndentedJSON(http.StatusOK, album)
}

func (routeHandler RouteHandler) postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Printf("Error during unmarshalling: %v", err)
		c.IndentedJSON(http.StatusBadRequest, newAlbum)
	}

	err := routeHandler.dataStore.addAlbum(newAlbum)
	if err != nil {
		fmt.Printf("Error adding album: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, newAlbum)
	}
}
