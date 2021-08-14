# mytheresa-api

This is a simple products and discounts api


## Requirements
- Clone this repo
- Install [go 1.16.3+](https://golang.org/doc/install).
- **IMPORTANT**: if you don't have ```go mod``` enabled, see this [article](https://lets-go.alexedwards.net/sample/02.02-project-setup-and-enabling-modules.html)

## How to run the tests
- Open the terminal, go to the root folder of this app and execute ```go test ./...```

## How to run the app
- Open the terminal, go to the root folder of this app and execute ```go run cmd/http/main.go```. This will run on port ```:8080```


## Decisions made

Golang was used as technology because it is simple, powerful and fast to build rest api. Simple applications up to complex levels can also be performed.

A hexagonal architecture was implemented because it is scalable and very flexible to business changes. Also, it can be implemented very well with Golang.

Considering that this is a PoC, I moved on with an in-memory database implementation in order to quickly wrap up the example. Due to the implementation of the architecture, you can create a new implementation of the repository interface and use the technology you want.

In the case of discounts, I thought of modeling them as records in a table itself (in memory) so that later I could obtain the discount through a request object and thus obtain the same.
They are modeled in a simple way, but the idea is that they have a priority, which must be assigned when creating it, to indicate what percentage is greater than another. On the other hand, the ideal would be that there are no repeated and active records, this could add one more property called "active" in order to have a history of them. There are several approaches that can be taken in this case.

In the case of products in memory, they are ordered when the app is started. If it is taken to a database, an index could be placed on the price column/document to be able to make queries and replicate the same thing that is done in this example.
Another possibility is to be able to have a database and in each instance that the server is up, generate an object in memory as this app does and thus improve the response. You could run a cron that is responsible for taking news every n amount of time. This decision was also taking into account that the maximum assumption of products could be 20k, which is relatively little for technologies like golang.
The idea of ​​having a map in memory is precisely to improve the response time of the api.

In the case of tests, they are to cover the main cases.
