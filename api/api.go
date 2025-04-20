package api

import (
	"errors"

	"github.com/baldurstod/go-steam-assetclassinfo/config"
	"github.com/gin-gonic/gin"
)

type ApiRequest struct {
	Action  string         `json:"action" binding:"required"`
	Version int            `json:"version" binding:"required"`
	Params  map[string]any `json:"params"`
}

type apiFunc func(apiKey string, params map[string]any) (map[string]any, error)
type apiKey struct {
	name    string
	version int
}

var api = map[apiKey]apiFunc{}

func declareApi(name string, version int, f apiFunc) {
	api[apiKey{name: name, version: version}] = f
}

func getApi(name string, version int) apiFunc {
	return api[apiKey{name: name, version: version}]
}

var apiConfig config.Api

func SetConfig(conf config.Api) {
	apiConfig = conf
}

func ApiHandler(c *gin.Context) {
	var request ApiRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		jsonError(c, errors.New("bad request: body is not json"))
		return
	}

	f := getApi(request.Action, request.Version)
	if f == nil {
		jsonError(c, NotFoundError{})
		return
	}

	result, err := f(apiConfig.SteamApiKey, request.Params)

	if err != nil {
		jsonError(c, err)
	} else {
		jsonSuccess(c, result)
	}
}
