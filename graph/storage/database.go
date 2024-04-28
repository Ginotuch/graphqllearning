package storage

type Database interface {
	Games() []Game
	SeriesList() []Series
	Reviews() []Review
	Authors() []Author
	Platforms() []Platform

	Game(id string) (*Game, error)
	AddGame(name string, seriesId *string, platformIds []string) (*Game, error)
	Series(id string) (*Series, error)
	AddSeries(name string) Series
	Review(id string) (*Review, error)
	AddReview(title string, content string, rating int, authorId string, gameId string) (*Review, error)
	Author(id string) (*Author, error)
	AddAuthor(name string) Author
	Platform(id string) (*Platform, error)
	AddPlatform(name string, company string) Platform
}
