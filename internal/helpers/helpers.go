package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/VatJittiprasert/goBooking.git/internal/config"
)

var app *config.AppConfig

func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, stasus int) {
	app.InfoLog.Println("Client error with statur of", stasus)
	http.Error(w, http.StatusText(stasus), stasus)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
