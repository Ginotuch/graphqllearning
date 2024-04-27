package storage

import "graphqllearning/graph/model"

type Author struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Game struct {
	ID          string   `json:"id"`
	SeriesID    string   `json:"seriesID"`
	Name        string   `json:"name"`
	PlatformIDs []string `json:"platformIDs"`
}

type Platform struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Company string `json:"company"`
}

type Review struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Rating   int    `json:"rating"`
	AuthorID string `json:"authorID"`
	GameID   string `json:"gameID"`
}

type Series struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (author Author) ToGraphModel() *model.Author {
	return &model.Author{
		ID:      author.ID,
		Name:    author.Name,
		Reviews: nil,
	}
}

func (game Game) ToGraphModel() *model.Game {
	return &model.Game{
		ID:        game.ID,
		Series:    nil,
		Name:      game.Name,
		Platforms: nil,
		Reviews:   nil,
	}
}

func (platform Platform) ToGraphModel() *model.Platform {
	return &model.Platform{
		ID:      platform.ID,
		Name:    platform.Name,
		Company: platform.Company,
		Games:   nil,
	}
}

func (review Review) ToGraphModel() *model.Review {
	return &model.Review{
		ID:      review.ID,
		Title:   review.Title,
		Content: review.Content,
		Rating:  review.Rating,
		Author:  nil,
		Game:    nil,
	}
}

func (series Series) ToGraphModel() *model.Series {
	return &model.Series{
		ID:    series.ID,
		Name:  series.Name,
		Games: nil,
	}
}
