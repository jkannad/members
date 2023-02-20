package helper

import (
	"runtime/debug"
	"fmt"
	"net/http"
	"github.com/jkannad/spas/members/internal/config"
)

var appConfig *config.AppConfig

func New(a *config.AppConfig){
	appConfig = a
}

func ClientError(w http.ResponseWriter, status int){
	appConfig.InfoLogger.Println("Client error with status of ", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error){
	/*trace := fmt.Sprintf("%s\n", err.Error())
	appConfig.ErrorLogger.Println(trace) */
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func ServiceInfo(info string){
	appConfig.InfoLogger.Println(info)
}

func ServiceError(err error){
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	appConfig.ErrorLogger.Println(trace)
}