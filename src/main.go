package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	// "github.com/mattn/go-sqlite3"
	// "database/sql"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type anime struct {
	ID               string `json:"id`
	Name             string `json:"name"`
	YearOfRelease    string `json:"year_of_release"`
	NumberOfEpisodes string `json:"number_of_episodes"`
	Rating           string `json:"rating"`
}

type data struct {
	Anime anime `json:"anime"`
	Album album `json:"album"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

var animes = []anime{
	{ID: "1", Name: "Naruto", YearOfRelease: "2002", NumberOfEpisodes: "220", Rating: "8.3"},
	{ID: "2", Name: "One Piece", YearOfRelease: "1999", NumberOfEpisodes: "1000+", Rating: "8.7"},
	{ID: "3", Name: "Attack on Titan", YearOfRelease: "2013", NumberOfEpisodes: "75", Rating: "9.0"},
}

var dataList = []data{
	{Anime: animes[0], Album: albums[0]},
	{Anime: animes[1], Album: albums[1]},
	{Anime: animes[2], Album: albums[2]},
}

func main() {

	// const database_path string = "database.db"
	// db, _ := sql.Open("sqlite3", database_path)
	// query, _ := db.Query("SELECT * FROM users")
	// fmt.Println(query)

	router := gin.Default()
	router.GET("/data", getData)
	router.GET("/data/:id", getDataById)
	getDataByKey("h")
	router.Run("https://rest-api-0g8j.onrender.com")
}

func getData(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, dataList)
}

func getDataById(c *gin.Context) {
	id := c.Param("id")

	for _, data := range dataList {
		if data.Anime.ID == id && data.Album.ID == id {
			c.IndentedJSON(http.StatusOK, data)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Data not found. Please check id param in URL"})
}

// ! This function will probably be needed to make it easier to get data, but not needed now
func getDataByKey(key string) {
	// temp_key := c.Param(string(key))
	for _, data := range dataList {
		fmt.Println(data)
	}
}

// func postData(c *gin.Context) {
// 	var new_album album

// 	if err := c.BindJSON(&new_album); err != nil {
// 		return
// 	}

// 	albums = append(albums, new_album)
// 	c.IndentedJSON(http.StatusCreated, new_album)
// }
