# tech-challenge-fase-1

Tech Challenge Fase 1...

Let's go tech challenge!!!

Links:

<https://miro.com/app/board/uXjVKQtHwOA=/>

## Run project

To run the application it is necessary to execute the command `make start`

### Aplication

### Migration

All migrations are executed as soon as the `make start` or `make build` command is executed

#### Create

To create a migration, you need to run the `make migrate/create` command passing the file name

example:

```bash
make migrate/create name=add_user
```

to create a migration to add a user

### Swagger

URL to access Swagger is <http://localhost:3003/api/v1/swagger/index.html>

#### Testing

To test using swagger, just site this down `api/v1/swagger/index.html` to give it a go.

#### Rebuilding

Instead the needing to rebuild the swagger documentation, install the swaggo package ant type `swag init` then documentation rebuilt.
