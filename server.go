package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"graphqllearning/graph"
	"graphqllearning/graph/storage"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	games := map[string]storage.Game{
		"1": {ID: "1", SeriesID: "1", Name: "Super Mario Bros", PlatformIDs: []string{"1", "2"}},
		"2": {ID: "2", SeriesID: "1", Name: "Super Mario Bros 2", PlatformIDs: []string{"1", "2"}},
		"3": {ID: "3", SeriesID: "2", Name: "The Legend of Zelda", PlatformIDs: []string{"1"}},
		"4": {ID: "4", SeriesID: "2", Name: "The Legend of Zelda: Breath of the Wild", PlatformIDs: []string{"3"}},
		"5": {ID: "5", Name: "Minecraft", PlatformIDs: []string{"1", "2", "3"}},
	}

	series := map[string]storage.Series{
		"1": {ID: "1", Name: "Super Mario"},
		"2": {ID: "2", Name: "The Legend of Zelda"},
	}

	reviews := map[string]storage.Review{
		"1":  {ID: "1", Title: "Great game", Content: "Lorem ipsum", Rating: 5, AuthorID: "1", GameID: "1"},
		"2":  {ID: "2", Title: "Awesome game", Content: "Lorem ipsum", Rating: 4, AuthorID: "2", GameID: "1"},
		"3":  {ID: "3", Title: "Fantastic game", Content: "Lorem ipsum", Rating: 5, AuthorID: "3", GameID: "2"},
		"4":  {ID: "4", Title: "Good game", Content: "Lorem ipsum", Rating: 4, AuthorID: "4", GameID: "2"},
		"5":  {ID: "5", Title: "Amazing game", Content: "Lorem ipsum", Rating: 5, AuthorID: "5", GameID: "3"},
		"6":  {ID: "6", Title: "Excellent game", Content: "Lorem ipsum", Rating: 5, AuthorID: "1", GameID: "3"},
		"7":  {ID: "7", Title: "Superb game", Content: "Lorem ipsum", Rating: 5, AuthorID: "2", GameID: "4"},
		"8":  {ID: "8", Title: "Wonderful game", Content: "Lorem ipsum", Rating: 5, AuthorID: "3", GameID: "4"},
		"9":  {ID: "9", Title: "Incredible game", Content: "Lorem ipsum", Rating: 5, AuthorID: "4", GameID: "5"},
		"10": {ID: "10", Title: "Unbelievable game", Content: "Lorem ipsum", Rating: 5, AuthorID: "5", GameID: "5"},
	}

	authors := map[string]storage.Author{
		"1": {ID: "1", Name: "John Doe"},
		"2": {ID: "2", Name: "Jane Doe"},
		"3": {ID: "3", Name: "Bob Smith"},
		"4": {ID: "4", Name: "Alice Smith"},
		"5": {ID: "5", Name: "Charlie Brown"},
		"6": {ID: "6", Name: "Lucy Van Pelt"},
	}

	platforms := map[string]storage.Platform{
		"1": {ID: "1", Name: "Nintendo Entertainment System", Company: "Nintendo"},
		"2": {ID: "2", Name: "Super Nintendo Entertainment System", Company: "Nintendo"},
		"3": {ID: "3", Name: "Nintendo Switch", Company: "Nintendo"},
	}

	db, err := storage.NewInMemoryDatabase(
		storage.WithGames(games),
		storage.WithSeries(series),
		storage.WithReviews(reviews),
		storage.WithAuthors(authors),
		storage.WithPlatforms(platforms),
	)
	if err != nil {
		panic(err)
	}
	resolver := graph.NewResolver(db)
	config := graph.Config{Resolvers: resolver}
	executableSchema := graph.NewExecutableSchema(config)
	srv := handler.NewDefaultServer(executableSchema)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
