package note

import (
	"github.com/graphql-go/graphql"
)

type (
	Note struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Author      Author `json:"author"`
	}

	Author struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
)

var (
	TypeAuthor = graphql.NewObject(graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.String,
				Description: "Id of author",
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "Name of author",
			},
			"age": &graphql.Field{
				Type:        graphql.Int,
				Description: "Age of author",
			},
		},
	})

	TypeNote = graphql.NewObject(graphql.ObjectConfig{
		Name: "Note",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.String,
				Description: "Id of note",
			},
			"title": &graphql.Field{
				Type:        graphql.String,
				Description: "Title of note",
			},
			"description": &graphql.Field{
				Type:        graphql.String,
				Description: "Description of note",
			},
			"author": &graphql.Field{
				Type:        TypeAuthor,
				Description: "Author of the note",
			},
		},
	})

	NotesQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"notes": &graphql.Field{
				Type:        graphql.NewList(TypeNote),
				Description: "Get all notes",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return NOTES, nil
				},
			},
		},
	})

	NoteSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: NotesQuery,
	})
)
