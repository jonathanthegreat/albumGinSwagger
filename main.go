package main

import (
	config "albumGinSwagger/configs"
	_ "albumGinSwagger/docs" // docs file is generated by Swag CLI, you have to import it.
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jerry", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// @title Album Store API
// @version 1.0
// @description Managing tool for album records
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email arian.khanjani@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
func main() {
	config.Connect()
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumById)

	// Swagger
	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	err := router.Run(":4080")
	if err != nil {
		return 
	}
}

// getAlbums responds with the list of all albums as JSON.

// GetAlbums godoc
// @Summary Get details of all albums
// @Description Get details of all albums
// @Tags albums
// @Accept  json
// @Produce  json
// @Success 200 {array} album
// @Router /albums [get]
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.

// PostAlbums godoc
// @Summary Create a new album record
// @Description Create a new album with the input payload
// @Tags albums
// @Accept  json
// @Produce  json
// @Param album body album true "Create Album"
// @Success 200 {object} album
// @Router /albums [post]
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// GetAlbumById godoc
// @Summary Get details of an album
// @Description Get details of an album
// @Tags albums/:id
// @Produce  json
// @Param id path string true "Album ID"
// @Success 200 {array} album
// @Router /albums/{id} [get]
func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
