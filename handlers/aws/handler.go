package aws

import (
	"encoding/json"
	"net/http"

	. "github.com/mlabouardy/komiser/services/aws"
	. "github.com/mlabouardy/komiser/services/cache"
)

type AWSHandler struct {
	cache    Cache
	multiple bool
	aws      AWS
	services map[string]interface{}
}

func NewAWSHandler(cache Cache, multiple bool, regions []string, services map[string]interface{}) *AWSHandler {
	awsHandler := AWSHandler{
		cache:    cache,
		multiple: multiple,
		aws: AWS{
			Regions: regions,
		},
		services: services,
	}
	return &awsHandler
}

func (handler *AWSHandler) GetAWSHandler() AWS {
	return handler.aws
}

func (handler *AWSHandler) HasMultipleEnvs() bool {
	return handler.multiple
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
