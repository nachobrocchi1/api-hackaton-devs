# api-hackaton-devs

Retrieves hackatons with the best three devs.

## Dependencies

1. You need a Postgres database, I suggest to use Docker. Run at the command line 'docker pull postgres' to download the latest image.
2. Run the Postgres image with 'docker run --name hackatonsC -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=docker -p 5432:5432 -d postgres'.
   
## Build

Run at the command line 'go build'.

## Run

### On Linux
Run at the command line ./api-hackaton-devs

### On Windows
Run at the command line ./api-hackaton-devs.exe

## How to use it
The folder 'examples' contains a postman collection with a prepared request called GetHackaton.

## Tests
Run at the command line 'go test ./... -cover'