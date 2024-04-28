package storage

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
)

type InMemoryDatabase struct {
	currentID atomic.Uint64

	mu sync.Mutex

	games     map[string]Game
	series    map[string]Series
	reviews   map[string]Review
	authors   map[string]Author
	platforms map[string]Platform
}

func (i *InMemoryDatabase) createNewID() string {
	return strconv.FormatUint(i.currentID.Add(1), 10)
}

func (i *InMemoryDatabase) AddGame(name string, seriesId *string, platformIds []string) (*Game, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	for _, platformId := range platformIds {
		if _, ok := i.platforms[platformId]; !ok {
			return nil, fmt.Errorf("platform ID %s not found", platformId)
		}
	}
	if seriesId != nil {
		if _, ok := i.series[*seriesId]; !ok {
			return nil, fmt.Errorf("seires ID %s not found", *seriesId)
		}
	}
	game := Game{
		ID:          i.createNewID(),
		SeriesID:    seriesId,
		Name:        name,
		PlatformIDs: platformIds,
	}
	i.games[game.ID] = game
	return &game, nil

}

func (i *InMemoryDatabase) AddSeries(name string) Series {
	i.mu.Lock()
	defer i.mu.Unlock()

	series := Series{
		ID:   i.createNewID(),
		Name: name,
	}
	i.series[series.ID] = series
	return series
}

func (i *InMemoryDatabase) AddReview(title string, content string, rating int, authorId string, gameId string) (*Review, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	if _, ok := i.authors[authorId]; !ok {
		return nil, fmt.Errorf("author ID %s not found", authorId)
	}
	if _, ok := i.games[gameId]; !ok {
		return nil, fmt.Errorf("game ID %s not found", gameId)
	}
	review := Review{
		ID:       i.createNewID(),
		Title:    title,
		Content:  content,
		Rating:   rating,
		AuthorID: authorId,
		GameID:   gameId,
	}
	i.reviews[review.ID] = review
	return &review, nil
}

func (i *InMemoryDatabase) AddAuthor(name string) Author {
	i.mu.Lock()
	defer i.mu.Unlock()

	author := Author{
		ID:   i.createNewID(),
		Name: name,
	}
	i.authors[author.ID] = author
	return author

}

func (i *InMemoryDatabase) AddPlatform(name string, company string) Platform {
	i.mu.Lock()
	defer i.mu.Unlock()

	platform := Platform{
		ID:      i.createNewID(),
		Name:    name,
		Company: company,
	}
	i.platforms[platform.ID] = platform
	return platform
}

func (i *InMemoryDatabase) Games() []Game {
	var games []Game
	for _, game := range i.games {
		games = append(games, game)
	}
	return games
}

func (i *InMemoryDatabase) SeriesList() []Series {
	var seriesList []Series
	for _, series := range i.series {
		seriesList = append(seriesList, series)
	}
	return seriesList
}

func (i *InMemoryDatabase) Reviews() []Review {
	var reviews []Review
	for _, review := range i.reviews {
		reviews = append(reviews, review)
	}
	return reviews
}

func (i *InMemoryDatabase) Authors() []Author {
	var authors []Author
	for _, author := range i.authors {
		authors = append(authors, author)
	}
	return authors
}

func (i *InMemoryDatabase) Platforms() []Platform {
	var platforms []Platform
	for _, platform := range i.platforms {
		platforms = append(platforms, platform)
	}
	return platforms
}

func (i *InMemoryDatabase) Game(id string) (*Game, error) {
	game, exists := i.games[id]
	if !exists {
		return nil, fmt.Errorf("game with ID %s not found", id)
	}
	return &game, nil
}

func (i *InMemoryDatabase) Series(id string) (*Series, error) {
	series, exists := i.series[id]
	if !exists {
		return nil, fmt.Errorf("series with ID %s not found", id)
	}
	return &series, nil
}

func (i *InMemoryDatabase) Review(id string) (*Review, error) {
	review, exists := i.reviews[id]
	if !exists {
		return nil, fmt.Errorf("review with ID %s not found", id)
	}
	return &review, nil
}

func (i *InMemoryDatabase) Author(id string) (*Author, error) {
	author, exists := i.authors[id]
	if !exists {
		return nil, fmt.Errorf("author with ID %s not found", id)
	}
	return &author, nil
}

func (i *InMemoryDatabase) Platform(id string) (*Platform, error) {
	platform, exists := i.platforms[id]
	if !exists {
		return nil, fmt.Errorf("platform with ID %s not found", id)
	}
	return &platform, nil
}

func NewInMemoryDatabase() *InMemoryDatabase {
	db := InMemoryDatabase{
		games:     make(map[string]Game),
		series:    make(map[string]Series),
		reviews:   make(map[string]Review),
		authors:   make(map[string]Author),
		platforms: make(map[string]Platform),
	}
	return &db
}
