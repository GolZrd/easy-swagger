package main

import (
	_ "easy-swagger/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title easy-Swagger example API
// @version 1.0
// @description easy-swagger example of API server

// @host localhost:8080
// @basePath /

type Note struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var Notes = []Note{
	{ID: "1", Title: "title1", Text: "text1"},
	{ID: "2", Title: "title2", Text: "text2"},
	{ID: "3", Title: "title3", Text: "text3"},
}

func main() {

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "welcome home")
	})

	r.GET("/notes", GetNotes)
	r.GET("/notes/:id", GetNoteByID)
	r.POST("/notes", PostNote)

	r.Run(":8080")
}

// getNotes godoc
// @Summary Get notes array
// @Description Responds with an array of notes as JSON
// @Tags notes
// @Accept json
// @Produce json
// @Success 200 {array} Note
// @Router /notes [get]
func GetNotes(c *gin.Context) {
	c.IndentedJSON(200, Notes)
}

// PostNote godoc
// @Summary Create note
// @Description Add a new note to the array. Returns the new note as JSON.
// @Tags notes
// @Accept json
// @Produce json
// @Param note body Note true "New note"
// @Success 201 {object} Note
// @Router /notes [post]
func PostNote(c *gin.Context) {
	newNote := Note{}

	if err := c.BindJSON(&newNote); err != nil {
		return
	}

	Notes = append(Notes, newNote)
	c.IndentedJSON(201, newNote)
}

// getNoteByID godoc
// @Summary Get note by ID
// @Description Responds with a note as JSON
// @Tags notes
// @Accept json
// @Produce json
// @Param id path string true "Note ID"
// @Success 200 {object} Note
// @Router /notes/{id} [get]
func GetNoteByID(c *gin.Context) {
	id := c.Param("id")

	for _, note := range Notes {
		if note.ID == id {
			c.IndentedJSON(200, note)
			return
		}
	}

	c.IndentedJSON(404, gin.H{"message": "note not found"})
}
