package filmcontroller

import (
	"encoding/json"
	"film-service/domain"
	"film-service/proxy/logger"
	filmservice "film-service/service/film"
	"film-service/utils"
	"film-service/utils/response"
	"net/http"
	"strings"
)

// CreateFilms --> Create films from the request body
func CreateFilms(res http.ResponseWriter, req *http.Request) {
	var body struct {
		Films []domain.Film `json:"films"`
	}
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		response.Error(res, http.StatusBadRequest, map[string]string{"message": "Mandatory params missing"})
		return
	}
	response.Success(res, 200, map[string]bool{"ok": true})
}

// GetFilms --> Fetches list of films from the database
func GetFilms(res http.ResponseWriter, req *http.Request) {
	queryParams := utils.ParseQueryParams(req.URL.Query())
	logger.Info("Retriving film information for params:", queryParams)
	if _, fetchByID := queryParams["id"]; fetchByID {
		ids := strings.Split(queryParams["id"], ",")
		films, err := filmservice.GetFilmsByID(ids)
		if err != nil {
			logger.Error("Failed to get films", err.Error())
			response.Error(res, 500, err.Error())
		}
		response.Success(res, 200, films)
		return
	}
	response.Error(res, http.StatusBadRequest, map[string]string{"message": "Mandatory params missing"})
}
