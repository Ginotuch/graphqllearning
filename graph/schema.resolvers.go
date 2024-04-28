package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"
	"graphqllearning/graph/model"
	"sync"
)

// Reviews is the resolver for the reviews field.
func (r *authorResolver) Reviews(ctx context.Context, obj *model.Author) ([]*model.Review, error) {
	var reviews []*model.Review
	for _, review := range r.database.Reviews() {
		if review.AuthorID == obj.ID {
			reviews = append(reviews, review.ToGraphModel())
		}
	}
	return reviews, nil
}

// Series is the resolver for the series field.
func (r *gameResolver) Series(ctx context.Context, obj *model.Game) (*model.Series, error) {
	game, err := r.database.Game(obj.ID)
	if err != nil {
		return nil, err
	}
	if game.SeriesID == nil {
		return nil, nil
	}
	series, err := r.database.Series(*game.SeriesID)
	if err != nil {
		return nil, err
	}
	return series.ToGraphModel(), nil
}

// Platforms is the resolver for the platforms field.
func (r *gameResolver) Platforms(ctx context.Context, obj *model.Game) ([]*model.Platform, error) {
	game, err := r.database.Game(obj.ID)
	if err != nil {
		return nil, err
	}
	var platforms []*model.Platform
	for _, platformID := range game.PlatformIDs {
		platform, err := r.database.Platform(platformID)
		if err != nil {
			// for now, we just skip it. This probably means the database changed while querying, but
			// I'm not super keen on making it thread safe just yet, just want to get it working
			continue
		}
		platforms = append(platforms, platform.ToGraphModel())
	}
	return platforms, nil
}

// Reviews is the resolver for the reviews field.
func (r *gameResolver) Reviews(ctx context.Context, obj *model.Game) ([]*model.Review, error) {
	var reviews []*model.Review
	for _, review := range r.database.Reviews() {
		if review.GameID == obj.ID {
			reviews = append(reviews, review.ToGraphModel())
		}
	}
	return reviews, nil
}

// AddPlatform is the resolver for the addPlatform field.
func (r *mutationResolver) AddPlatform(ctx context.Context, platform model.PlatformInput) (*model.Platform, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.addPlatform(ctx, platform)
}

func (r *mutationResolver) addPlatform(ctx context.Context, platform model.PlatformInput) (*model.Platform, error) {
	if platform.ID != nil {
		if platform.Name != nil || platform.Company != nil {
			return nil, fmt.Errorf("cannot specify an ID while also specifying the Name or Company when creating a new platform")
		}
		foundPlatform, err := r.database.Platform(*platform.ID)
		if err != nil {
			return nil, err
		}
		return foundPlatform.ToGraphModel(), nil
	}

	if platform.Name == nil || platform.Company == nil {
		return nil, fmt.Errorf("platform name and company cannot be nil")
	}

	return r.database.AddPlatform(*platform.Name, *platform.Company).ToGraphModel(), nil
}

// AddSeries is the resolver for the addSeries field.
func (r *mutationResolver) AddSeries(ctx context.Context, series model.SeriesInput) (*model.Series, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.addSeries(ctx, series)
}

func (r *mutationResolver) addSeries(ctx context.Context, series model.SeriesInput) (*model.Series, error) {
	if series.ID != nil && series.Name != nil {
		return nil, fmt.Errorf("cannot supply both series ID and name for creation")
	}

	if series.ID != nil {
		// An ID was specified, first try to find it to use the existing one, otherwise error
		foundSeries, err := r.database.Series(*series.ID)
		if err != nil {
			return nil, err
		}
		return foundSeries.ToGraphModel(), nil
	}

	if series.Name == nil {
		return nil, fmt.Errorf("series name cannot be nil")
	}

	// No ID was specified, only the Name. So, create a new Series
	return r.database.AddSeries(*series.Name).ToGraphModel(), nil
}

// AddGame is the resolver for the addGame field.
func (r *mutationResolver) AddGame(ctx context.Context, game model.GameInput) (*model.Game, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.addGame(ctx, game)
}

func (r *mutationResolver) addGame(ctx context.Context, game model.GameInput) (*model.Game, error) {
	if game.ID != nil {
		if game.Name != nil || game.Series != nil || len(game.Platforms) != 0 {
			return nil, fmt.Errorf("cannot create new game with ID and other details specified together")
		}
		foundGame, err := r.database.Game(*game.ID)
		if err != nil {
			return nil, err
		}
		return foundGame.ToGraphModel(), nil
	}

	if game.Name == nil {
		return nil, fmt.Errorf("name cannot be nil")
	}

	var seriesId *string

	if game.Series != nil {
		series, err := r.addSeries(ctx, *game.Series)
		if err != nil {
			return nil, err
		}
		seriesId = &series.ID
	}

	var platformIDs []string
	for _, platformInput := range game.Platforms {
		platform, err := r.addPlatform(ctx, *platformInput)
		if err != nil {
			return nil, err
		}
		platformIDs = append(platformIDs, platform.ID)
	}

	newGame, err := r.database.AddGame(*game.Name, seriesId, platformIDs)
	if err != nil {
		return nil, err
	}
	return newGame.ToGraphModel(), nil
}

// AddAuthor is the resolver for the addAuthor field.
func (r *mutationResolver) AddAuthor(ctx context.Context, author model.AuthorInput) (*model.Author, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.addAuthor(ctx, author)
}

func (r *mutationResolver) addAuthor(ctx context.Context, author model.AuthorInput) (*model.Author, error) {
	if (author.ID == nil && author.Name == nil) || (author.ID != nil && author.Name != nil) {
		return nil, fmt.Errorf("must specify either the ID or the name (exclusive)")
	}
	if author.ID != nil {
		foundAuthor, err := r.database.Author(*author.ID)
		if err != nil {
			return nil, err
		}
		return foundAuthor.ToGraphModel(), nil
	}

	return r.database.AddAuthor(*author.Name).ToGraphModel(), nil
}

// AddReview is the resolver for the addReview field.
func (r *mutationResolver) AddReview(ctx context.Context, review model.ReviewInput) (*model.Review, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if review.Author == nil || review.Game == nil {
		return nil, fmt.Errorf("both Author and Game must not be nil")
	}

	author, err := r.addAuthor(ctx, *review.Author)
	if err != nil {
		return nil, err
	}
	game, err := r.addGame(ctx, *review.Game)
	if err != nil {
		return nil, err
	}
	newReview, err := r.database.AddReview(review.Title, review.Content, review.Rating, author.ID, game.ID)
	if err != nil {
		return nil, err
	}
	return newReview.ToGraphModel(), nil
}

// Games is the resolver for the games field.
func (r *platformResolver) Games(ctx context.Context, obj *model.Platform) ([]*model.Game, error) {
	var games []*model.Game
	for _, game := range r.database.Games() {
		for _, platformID := range game.PlatformIDs {
			if platformID == obj.ID {
				games = append(games, game.ToGraphModel())
			}
		}
	}
	return games, nil
}

// Reviews is the resolver for the reviews field.
func (r *queryResolver) Reviews(ctx context.Context) ([]*model.Review, error) {
	var reviews []*model.Review
	for _, review := range r.database.Reviews() {
		reviews = append(reviews, review.ToGraphModel())
	}
	return reviews, nil
}

// Review is the resolver for the review field.
func (r *queryResolver) Review(ctx context.Context, id string) (*model.Review, error) {
	review, err := r.database.Review(id)
	if err != nil {
		return nil, err
	}
	return review.ToGraphModel(), nil
}

// Games is the resolver for the games field.
func (r *queryResolver) Games(ctx context.Context) ([]*model.Game, error) {
	var games []*model.Game
	for _, game := range r.database.Games() {
		games = append(games, game.ToGraphModel())
	}
	return games, nil
}

// Game is the resolver for the game field.
func (r *queryResolver) Game(ctx context.Context, id string) (*model.Game, error) {
	game, err := r.database.Game(id)
	if err != nil {
		return nil, err
	}
	return game.ToGraphModel(), nil
}

// SeriesList is the resolver for the seriesList field.
func (r *queryResolver) SeriesList(ctx context.Context) ([]*model.Series, error) {
	var seriesList []*model.Series
	for _, series := range r.database.SeriesList() {
		seriesList = append(seriesList, series.ToGraphModel())
	}
	return seriesList, nil
}

// Series is the resolver for the series field.
func (r *queryResolver) Series(ctx context.Context, id string) (*model.Series, error) {
	series, err := r.database.Series(id)
	if err != nil {
		return nil, err
	}
	return series.ToGraphModel(), nil
}

// Authors is the resolver for the authors field.
func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	var authors []*model.Author
	for _, author := range r.database.Authors() {
		authors = append(authors, author.ToGraphModel())
	}
	return authors, nil
}

// Author is the resolver for the author field.
func (r *queryResolver) Author(ctx context.Context, id string) (*model.Author, error) {
	author, err := r.database.Author(id)
	if err != nil {
		return nil, err
	}
	return author.ToGraphModel(), nil
}

// Platforms is the resolver for the platforms field.
func (r *queryResolver) Platforms(ctx context.Context) ([]*model.Platform, error) {
	var platforms []*model.Platform
	for _, platform := range r.database.Platforms() {
		platforms = append(platforms, platform.ToGraphModel())
	}
	return platforms, nil
}

// Platform is the resolver for the platform field.
func (r *queryResolver) Platform(ctx context.Context, id string) (*model.Platform, error) {
	platform, err := r.database.Platform(id)
	if err != nil {
		return nil, err
	}
	return platform.ToGraphModel(), nil
}

// Author is the resolver for the author field.
func (r *reviewResolver) Author(ctx context.Context, obj *model.Review) (*model.Author, error) {
	review, err := r.database.Review(obj.ID)
	if err != nil {
		return nil, err
	}
	author, err := r.database.Author(review.AuthorID)
	if err != nil {
		return nil, err
	}
	return author.ToGraphModel(), nil
}

// Game is the resolver for the game field.
func (r *reviewResolver) Game(ctx context.Context, obj *model.Review) (*model.Game, error) {
	review, err := r.database.Review(obj.ID)
	if err != nil {
		return nil, err
	}
	game, err := r.database.Game(review.GameID)
	if err != nil {
		return nil, err
	}
	return game.ToGraphModel(), nil
}

// Games is the resolver for the games field.
func (r *seriesResolver) Games(ctx context.Context, obj *model.Series) ([]*model.Game, error) {
	var games []*model.Game
	for _, game := range r.database.Games() {
		if *game.SeriesID == obj.ID {
			games = append(games, game.ToGraphModel())
		}
	}
	return games, nil
}

// Author returns AuthorResolver implementation.
func (r *Resolver) Author() AuthorResolver { return &authorResolver{r} }

// Game returns GameResolver implementation.
func (r *Resolver) Game() GameResolver { return &gameResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{Resolver: r} }

// Platform returns PlatformResolver implementation.
func (r *Resolver) Platform() PlatformResolver { return &platformResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Review returns ReviewResolver implementation.
func (r *Resolver) Review() ReviewResolver { return &reviewResolver{r} }

// Series returns SeriesResolver implementation.
func (r *Resolver) Series() SeriesResolver { return &seriesResolver{r} }

type authorResolver struct{ *Resolver }
type gameResolver struct{ *Resolver }
type mutationResolver struct {
	*Resolver
	mu sync.Mutex
}
type platformResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type reviewResolver struct{ *Resolver }
type seriesResolver struct{ *Resolver }
