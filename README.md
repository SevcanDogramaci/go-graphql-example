
# Go GraphQL Example

## Queries

Get all notes

```

curl --location -g --request GET 'http://localhost:3000/graphql?query={notes{id, title, description, author {name}}}'

```

Result:

```json

{
    "data": {
        "notes": [
            {
                "author": {
                    "name": "Jasmine"
                },
                "description": "Prioritize refactoring task",
                "id": "1",
                "title": "Meeting Note"
            },
            {
                "author": {
                    "name": "Jasmine"
                },
                "description": "Love yourself !",
                "id": "2",
                "title": "Self Note"
            },
            {
                "author": {
                    "name": "Jane"
                },
                "description": "Call me :)",
                "id": "3",
                "title": "Mom's Note"
            }
        ]
    }
}
```