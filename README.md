## Description

This Go repo contains one model: bank. The bank account model has the following attributes:

* id: unique id of the bank account
* name: name of the bank account
* iban: accounts iban

The excerise aimed to enable an end user to call the api with the post / get / destroy / put methods. Learnings included:

1. Chi http Handler
2. PGX v5 databasedriver

## Navigation

* [Requirements](#requirements)
* [Set-Up](#set-up)
* [Run Service on Local Development](#run-service-on-local-development)
* [Improvements](#improvements)

## Requirements

This repo currently works with:

* Go 1.19
* PostgreSQL

## Set-Up

- Pull the project from this public repository using clone repo `gh repo clone Wigwamwam/CRUD_app` and enter the repo `cd CRUD_app`.
- Run `go mod download` in your terminal to install the necessary gems and dependencies.
- Run `go build` in your terminal to build the Go code. 
- Run `go run main.go` in your terminal host the app on http://localhost:3000/

## Run Service on Local Development

To call the API service, go to Postman [https://www.postman.com/] and sign up. Then, follow these instructions [https://learning.postman.com/docs/getting-started/sending-the-first-request/] with the following route paths to access the service:

```
GET         http://localhost:3000/v1/bank_accounts
GET         http://localhost:3000/v1/bank_accounts/:id
POST        http://localhost:3000/v1/bank_accounts
DELETE      http://localhost:3000/v1/bank_accounts/:id
PUT         http://localhost:3000/v1/bank_accounts/:id
```

## Improvements

Here are several aspects to like to learn more about and add to this repo:

* Testing
* Dockerise the repo
* DB validations
* Added authentication
* Set up continuous intergration on Github
* Deploy on GCP


