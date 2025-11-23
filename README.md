# Weight Tracker Application

This is an application to help people track their weight over time.

It loosely follows the Controller-Service-Repo pattern found in most spring boot applications.

### To run locally

Install air

`go install github.com/air-verse/air@latest`

Run air

`air`

## Package/Folder Overview

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


## Todos

- If user enters a weight for a date that already exists, it needs to update.

- Add constraints to the weight entry forms to not allow floating numbers.

- Consider adding a grid view that lets you edit and update.

- Add a line to the graph that shows the target weight.