# Weight Tracker Application

This is an application to help people track their weight over time.

It loosely follows the Controller-Service-Repo pattern found in most Java Spring Boot applications.

### To run locally

Install air

`go install github.com/air-verse/air@latest`

Run local Makefile command

`make local`

## Package/Folder Overview

`config`

Contains app config.

`controller`

Contains all the logic for handling/responding to requests.

`middleware`

Contains all the code for the middleware layer.

`service`

Contains most of the business logic.

`repo`

Contains logic for interacting with the storage layer.

`model`

Contains all data objects.

`web`

Contains html/css/js/images.

`sql`

Contains all SQL scripts.

## Local DB setup steps (MacOS)

install postgres 18

`brew install postgresql@18`

start postgres

`brew services start postgresql@18`

add postgres binaries to path

`echo 'export PATH="/opt/homebrew/opt/postgresql@18/bin:$PATH"' >> ~/.zshrc`

apply update

`source ~/.zshrc`

connect to the default db

`psql -d postgres`

run the scripts in the sql folder in this order:

`create_db_and_user.sql`

`create_tables.sql`

## Todos

- Add password hashing.

- If user enters a weight for a date that already exists, it needs to update.

- Add constraints to the weight entry forms to not allow floating numbers.

- Consider adding a grid view that lets you edit and update.

- Add a line to the graph that shows the target weight.

- Create the database.