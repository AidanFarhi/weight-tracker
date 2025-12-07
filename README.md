# Weight Tracker Application

This is an application to help people track their weight over time.

It loosely follows the Controller-Service-Repo pattern found in most Java Spring Boot applications.

App URL:

https://afarhidev-weight-tracker-ed183f93bb4c.herokuapp.com/

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

`scripts`

Contains all scripts (build, local stuff, ect...)

`model`

Contains all data objects.

`web`

Contains html/css/js/images.

`sql`

Contains all SQL scripts.

## Local DB setup steps (MacOS)

Install postgres 18

`brew install postgresql@18`

Start postgres

`brew services start postgresql@18`

Add postgres binaries to path

`echo 'export PATH="/opt/homebrew/opt/postgresql@18/bin:$PATH"' >> ~/.zshrc`

Apply update

`source ~/.zshrc`

Create the user and database

`psql -d postgres -f ./sql/create_db_and_user.sql`

Create the tables

`psql -U wt_user -d wt_db -f ./sql/create_tables.sql`

## To run app locally

Install air

`go install github.com/air-verse/air@latest`

Set environment variable to run locally

`export WT_APP_ENV=local`

Run local Makefile command

`make local`

## Todos

- Unit testing.

- Add constraints to the weight entry forms to not allow floating numbers.

- Add a logout feature.

- Grid view that lets you edit and update.

- Add placeholder values for missing weight data.

- Add a prompt when the user first logs in and there is no target weight set.