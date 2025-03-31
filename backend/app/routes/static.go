package routes

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	core "learn.zone01dakar.sn/forum-rest-api/app/internals/core/app"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"
)

type Static struct{}

func (g Static) Route(app *core.App) {
	app.GET("/static/uploads/{item:string}", g.Getstatic)
}

func (g Static) Getstatic(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	// Serve static files (CSS, JavaScript, images, etc.)
	// http.Handle("/static/uploads/", http.StripPrefix("/static/uploads/", http.FileServer(http.Dir("/static/uploads/"))))

	imagePath := strings.TrimPrefix(r.URL.Path, "/api/")

	// Open the image file
	file, err := os.Open(imagePath)
	if err != nil {
		lib.ErrorWriter(&response, "Image not found", http.StatusNotFound)
		lib.ResponseFormatter(w, response)
		return
	}
	defer file.Close()

	// Set the appropriate content type for the image (e.g., "image/jpeg" for JPEG)
	w.Header().Set("Content-Type", "image/jpeg")

	// Copy the image data to the response writer
	_, err = io.Copy(w, file)
	if err != nil {
		fmt.Println("Error serving image:", err)
		lib.ErrorWriter(&response, "Internal Server Error", http.StatusInternalServerError)
		lib.ResponseFormatter(w, response)
		return
	}
}
