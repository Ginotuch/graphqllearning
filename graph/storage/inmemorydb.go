package storage

import (
	"fmt"
)

type InMemoryDatabase struct {
	games     map[string]Game
	series    map[string]Series
	reviews   map[string]Review
	authors   map[string]Author
	platforms map[string]Platform
}

func (i InMemoryDatabase) Games() []Game {
	var games []Game
	for _, game := range i.games {
		games = append(games, game)
	}
	return games
}

func (i InMemoryDatabase) SeriesList() []Series {
	var seriesList []Series
	for _, series := range i.series {
		seriesList = append(seriesList, series)
	}
	return seriesList
}

func (i InMemoryDatabase) Reviews() []Review {
	var reviews []Review
	for _, review := range i.reviews {
		reviews = append(reviews, review)
	}
	return reviews
}

func (i InMemoryDatabase) Authors() []Author {
	var authors []Author
	for _, author := range i.authors {
		authors = append(authors, author)
	}
	return authors
}

func (i InMemoryDatabase) Platforms() []Platform {
	var platforms []Platform
	for _, platform := range i.platforms {
		platforms = append(platforms, platform)
	}
	return platforms
}

func (i InMemoryDatabase) Game(id string) (*Game, error) {
	game, exists := i.games[id]
	if !exists {
		return nil, fmt.Errorf("game with ID %s not found", id)
	}
	return &game, nil
}

func (i InMemoryDatabase) Series(id string) (*Series, error) {
	series, exists := i.series[id]
	if !exists {
		return nil, fmt.Errorf("series with ID %s not found", id)
	}
	return &series, nil
}

func (i InMemoryDatabase) Review(id string) (*Review, error) {
	review, exists := i.reviews[id]
	if !exists {
		return nil, fmt.Errorf("review with ID %s not found", id)
	}
	return &review, nil
}

func (i InMemoryDatabase) Author(id string) (*Author, error) {
	author, exists := i.authors[id]
	if !exists {
		return nil, fmt.Errorf("author with ID %s not found", id)
	}
	return &author, nil
}

func (i InMemoryDatabase) Platform(id string) (*Platform, error) {
	platform, exists := i.platforms[id]
	if !exists {
		return nil, fmt.Errorf("platform with ID %s not found", id)
	}
	return &platform, nil
}

type InMemoryDatabaseOption = func(database *InMemoryDatabase) error

func WithGames(games map[string]Game) InMemoryDatabaseOption {
	return func(db *InMemoryDatabase) error {
		db.games = games
		return nil
	}
}
func WithSeries(series map[string]Series) InMemoryDatabaseOption {
	return func(db *InMemoryDatabase) error {
		db.series = series
		return nil
	}
}
func WithReviews(review map[string]Review) InMemoryDatabaseOption {
	return func(db *InMemoryDatabase) error {
		db.reviews = review
		return nil
	}
}
func WithAuthors(authors map[string]Author) InMemoryDatabaseOption {
	return func(db *InMemoryDatabase) error {
		db.authors = authors
		return nil
	}
}
func WithPlatforms(platforms map[string]Platform) InMemoryDatabaseOption {
	return func(db *InMemoryDatabase) error {
		db.platforms = platforms
		return nil
	}
}

func NewInMemoryDatabase(opts ...InMemoryDatabaseOption) (*InMemoryDatabase, error) {
	db := InMemoryDatabase{
		games:     make(map[string]Game),
		series:    make(map[string]Series),
		reviews:   make(map[string]Review),
		authors:   make(map[string]Author),
		platforms: make(map[string]Platform),
	}

	for _, opt := range opts {
		err := opt(&db)
		if err != nil {
			return nil, err
		}
	}

	return &db, nil
}
