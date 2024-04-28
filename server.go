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

	db := storage.NewInMemoryDatabase()

	// Add authors
	author1 := db.AddAuthor("John Doe")
	author2 := db.AddAuthor("Jane Doe")
	author3 := db.AddAuthor("Bob Smith")
	author4 := db.AddAuthor("Alice Smith")
	author5 := db.AddAuthor("Charlie Brown")

	// Add platforms
	platform1 := db.AddPlatform("Nintendo Entertainment System", "Nintendo")
	platform2 := db.AddPlatform("Super Nintendo Entertainment System", "Nintendo")
	platform3 := db.AddPlatform("Nintendo Switch", "Nintendo")

	// Add series
	series1 := db.AddSeries("Super Mario")
	series2 := db.AddSeries("The Legend of Zelda")

	// Add games
	game1, _ := db.AddGame("Super Mario Bros", &series1.ID, []string{platform1.ID, platform2.ID})
	game2, _ := db.AddGame("Super Mario Bros 2", &series1.ID, []string{platform1.ID, platform2.ID})
	game3, _ := db.AddGame("The Legend of Zelda", &series2.ID, []string{platform1.ID})
	game4, _ := db.AddGame("The Legend of Zelda: Breath of the Wild", &series2.ID, []string{platform3.ID})
	game5, _ := db.AddGame("Minecraft", nil, []string{platform1.ID, platform2.ID, platform3.ID})

	// Add reviews
	_, _ = db.AddReview("Great game", "Lorem ipsum", 5, author1.ID, game1.ID)
	_, _ = db.AddReview("Awesome game", "Lorem ipsum", 4, author2.ID, game1.ID)
	_, _ = db.AddReview("Fantastic game", "Lorem ipsum", 5, author3.ID, game2.ID)
	_, _ = db.AddReview("Good game", "Lorem ipsum", 4, author4.ID, game2.ID)
	_, _ = db.AddReview("Amazing game", "Lorem ipsum", 5, author5.ID, game3.ID)
	_, _ = db.AddReview("Excellent game", "Lorem ipsum", 5, author1.ID, game3.ID)
	_, _ = db.AddReview("Superb game", "Lorem ipsum", 5, author2.ID, game4.ID)
	_, _ = db.AddReview("Wonderful game", "Lorem ipsum", 5, author3.ID, game4.ID)
	_, _ = db.AddReview("Incredible game", "Lorem ipsum", 5, author4.ID, game5.ID)
	_, _ = db.AddReview("Unbelievable game", "Lorem ipsum", 5, author5.ID, game5.ID)

	resolver := graph.NewResolver(db)
	config := graph.Config{Resolvers: resolver}
	executableSchema := graph.NewExecutableSchema(config)
	srv := handler.NewDefaultServer(executableSchema)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
