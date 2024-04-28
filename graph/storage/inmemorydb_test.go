package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInMemoryDatabase_AddGame(t *testing.T) {
	db := NewInMemoryDatabase()

	// Adding a new game with no series and no platforms
	game, err := db.AddGame("New Game", nil, []string{})
	assert.Nil(t, err)
	assert.NotNil(t, game)
	assert.Equal(t, "New Game", game.Name)
	assert.Nil(t, game.SeriesID)
	assert.Empty(t, game.PlatformIDs)

	// Adding a game with non-existing platform ID
	game, err = db.AddGame("Game 2", nil, []string{"1"})
	assert.NotNil(t, err)
	assert.Nil(t, game)

	// Valid platform and game addition
	platform := db.AddPlatform("Platform One", "Company One")
	game, err = db.AddGame("Game With Platform", nil, []string{platform.ID})
	assert.Nil(t, err)
	assert.NotNil(t, game)
	assert.Contains(t, game.PlatformIDs, platform.ID)

	// Adding a game with non-existing series ID
	nonExistingSeriesId := "non-existing"
	game, err = db.AddGame("Game With Series", &nonExistingSeriesId, []string{})
	assert.NotNil(t, err)
	assert.Nil(t, game)

	// Valid series and game addition
	series := db.AddSeries("New Series")
	game, err = db.AddGame("Game In Series", &series.ID, []string{})
	assert.Nil(t, err)
	assert.NotNil(t, game)
	assert.Equal(t, series.ID, *game.SeriesID)
}

func TestInMemoryDatabase_AddSeries(t *testing.T) {
	db := NewInMemoryDatabase()
	series := db.AddSeries("Series 1")
	assert.NotEmpty(t, series.ID)
	assert.Equal(t, "Series 1", series.Name)
}

func TestInMemoryDatabase_AddReview(t *testing.T) {
	db := NewInMemoryDatabase()
	author := db.AddAuthor("Author One")
	game, err := db.AddGame("Game One", nil, nil)

	// Adding a review with valid author and game
	review, err := db.AddReview("Review Title", "Review Content", 5, author.ID, game.ID)
	assert.Nil(t, err)
	assert.NotNil(t, review)
	assert.Equal(t, "Review Title", review.Title)
	assert.Equal(t, 5, review.Rating)

	// Adding a review with non-existing author
	review, err = db.AddReview("Review 2", "Content 2", 3, "non-existing", game.ID)
	assert.NotNil(t, err)
	assert.Nil(t, review)

	// Adding a review with non-existing game
	review, err = db.AddReview("Review 3", "Content 3", 4, author.ID, "non-existing")
	assert.NotNil(t, err)
	assert.Nil(t, review)
}

func TestInMemoryDatabase_AddAuthor(t *testing.T) {
	db := NewInMemoryDatabase()
	author := db.AddAuthor("Author Two")
	assert.NotEmpty(t, author.ID)
	assert.Equal(t, "Author Two", author.Name)
}

func TestInMemoryDatabase_AddPlatform(t *testing.T) {
	db := NewInMemoryDatabase()
	platform := db.AddPlatform("Platform Two", "Company Two")
	assert.NotEmpty(t, platform.ID)
	assert.Equal(t, "Platform Two", platform.Name)
	assert.Equal(t, "Company Two", platform.Company)
}

func TestInMemoryDatabase_Getters(t *testing.T) {
	db := NewInMemoryDatabase()
	author := db.AddAuthor("Some Author")
	series := db.AddSeries("Some Series")
	platform := db.AddPlatform("Some Platform", "some Company")

	// Test fetching lists
	assert.Len(t, db.Authors(), 1)
	assert.Len(t, db.SeriesList(), 1)
	assert.Len(t, db.Platforms(), 1)

	// Test individual fetches
	gotAuthor, _ := db.Author(author.ID)
	assert.Equal(t, author.Name, gotAuthor.Name)

	gotSeries, _ := db.Series(series.ID)
	assert.Equal(t, series.Name, gotSeries.Name)

	gotPlatform, _ := db.Platform(platform.ID)
	assert.Equal(t, platform.Name, gotPlatform.Name)
}

func TestInMemoryDatabase_Errors(t *testing.T) {
	db := NewInMemoryDatabase()

	// Test non-existent entries
	_, err := db.Game("non-existent")
	assert.NotNil(t, err)

	_, err = db.Series("non-existent")
	assert.NotNil(t, err)

	_, err = db.Review("non-existent")
	assert.NotNil(t, err)

	_, err = db.Author("non-existent")
	assert.NotNil(t, err)

	_, err = db.Platform("non-existent")
	assert.NotNil(t, err)
}
