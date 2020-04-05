package controllers
import (
	"net/http"

	"github.com/YavuzSelimGugen/fullstack/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}