package main

import (
	"fmt"

	"github.com/SevcanDogramaci/go-graphql-example/internal/note"
	"github.com/gofiber/fiber/v2"
	"github.com/graphql-go/graphql"
)

func executeGrapQL(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	app := fiber.New()

	app.Get("/graphql", func(c *fiber.Ctx) error {
		query := c.Query("query")
		queryResult := executeGrapQL(query, note.NoteSchema)

		return c.JSON(queryResult)
	})

	app.Listen(":3000")
}
