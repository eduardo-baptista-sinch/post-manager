## Requirements:

 - project organization
 - request validation
 - DB integration
 - Test env
 - Integration tests
 - Authentication
 - Authorization
 - API integration
 - Swagger Docs

## Features:

- Health check
    - Show
        - Method: GET
        - Responses:
            - 204: No Content
    
- User
    - Create
        - Method: POST
        - Responses:
            - 201: Created
    - List
        - Method: GET
        - Responses:
            - 200: Ok
    - Show
        - Method: GET
        - Responses:
            - 200: Ok
            - 404: Not found

- Posts

## Commands

Run tests:
 - go test ./... 
 - go test ./... --cover