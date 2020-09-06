package domain

// Film struct properties
type Film struct {
	FilmID      int64  `json:"id" db:"film_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	ReleaseYear string `json:"releaseYear" db:"release_year"`
	LanguageID  int64  `json:"languageId" db:"language_id"`
	Language    string `json:"language" db:"name"`
}
