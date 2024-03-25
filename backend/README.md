# PostgreSQL Versioning

Currently have a postgreSQL server running locally.

Using golang-migrate I am carrying out version control to avoid ending up with my db in a bad state.

Key scripts to know:

- to create a new migration file: `create -ext sql -dir ./migrations/ -seq init_mg`
- to run migration upwards and persist new changes: `migrate -path ./migrations -database "postgres connection string" -verbose up` with the connection string looking something like:
  postgresql://username:password@host:port/database
  (you can also swap up with down to revert changes to database)
