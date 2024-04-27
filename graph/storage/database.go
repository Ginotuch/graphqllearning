package storage

type Database interface {
	Games() []Game
	SeriesList() []Series
	Reviews() []Review
	Authors() []Author
	Platforms() []Platform

	Game(id string) (*Game, error)
	Series(id string) (*Series, error)
	Review(id string) (*Review, error)
	Author(id string) (*Author, error)
	Platform(id string) (*Platform, error)
}
