package controllers

import (
	"albumGinSwagger/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAlbums
// @Summary Get all albums
// @Description Get details of all albums
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Album
// @Router /v1/albums [get]
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Albums)
}

// PostAlbums godoc
// @Summary Create a new album record
// @Description Create a new album with the input payload
// @Accept  json
// @Produce  json
// @Param models.Album body models.Album true "Create Album"
// @Success 200 {object} models.Album
// @Router /v1/albums [post]
func PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	models.Albums = append(models.Albums, newAlbum)
	//c.IndentedJSON(http.StatusCreated, newAlbum) --old--
	//c.JSON(http.StatusOK, gin.H{"data": album}) --for-pro's--
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// GetAlbumById godoc
// @Summary Get details of an album
// @Description Get details of an album
// @Produce  json
// @Param id path string true "Album ID"
// @Success 200 {array} models.Album
// @Router /v1/albums/{id} [get]
func GetAlbumById(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
    for _, album := range models.Albums {
    	if album.ID == id {
    		c.JSON(http.StatusOK, album)
    		return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
