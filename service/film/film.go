package filmservice

import (
	"film-service/domain"
	"film-service/proxy/logger"
	"film-service/proxy/mysqlproxy"
	"log"
	"strconv"
)

//GetFilmsByID --> Fetches list of film by id
func GetFilmsByID(IDs []string) ([]domain.Film, error) {
	query := `SELECT
						film_id,
						title,
						description,
						release_year,
						f.language_id,
						l.name
					FROM
						film f
						INNER JOIN language l on f.language_id = l.language_id
					WHERE
						f.film_id in (:ids)`

	params := make([]int, len(IDs))
	for i, id := range IDs {
		params[i], _ = strconv.Atoi(id)
	}

	arg := map[string]interface{}{
		"ids": params,
	}

	rows, queryError := mysqlproxy.ExecuteNamedQuery(query, arg)

	if queryError != nil {
		log.Fatal("Failed to parse struct", queryError.Error())
		return nil, queryError
	}

	var films []domain.Film
	for rows.Next() {
		var film domain.Film
		err := rows.StructScan(&film)
		if err != nil {
			log.Fatal("Failed to parse struct", err)
		}
		films = append(films, film)
	}
	logger.Info("Retrieved films:", len(films))
	return films, nil
}

//UpsertFilms --> Creates/Updates list of films
func UpsertFilms(Films []domain.Film) {

}
