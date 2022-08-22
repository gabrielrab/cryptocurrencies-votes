# Cryptocurrencies Votes

<img alt="Logo" align="right" src="https://emojipedia-us.s3.dualstack.us-west-1.amazonaws.com/thumbs/240/apple/325/coin_1fa99.png" width="128" />

### Description
API for the purpose of receiving votes from referrals created with cryptocurrencies.

### Files Structure
This API was created following the principles of Clean Architecture and contains the following folder and file structure.
```
├── Dockerfile
├── README.md
├── controllers
│   ├── coin_controller.go
│   └── vote_controller.go
├── database
│   ├── database.go
│   ├── migrations.go
│   └── repositories
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── models
│   ├── coin.go
│   └── vote.go
└── server
    ├── routes
    │   └── routes.go
    ├── server.go
    └── socket
        └── socket.go
```
### How to run locally
> Before starting the application it is necessary that you have installed the [Docker](https://docs.docker.com/engine/install/) and [Docker Compose](https://docs.docker.com/compose/install/).

- On your computer, clone this repository
    `` git clone https://github.com/gabrielrab/cryptocurrencies-votes.git ``
- Create a file `.env` at the root of the repository and fill in the variables as per the example file `.env.example`
- Run this build command
    `` docker-compose up --build -d  ``

### Endpoints
List and description from endpoints from this API.
```
GET    /                         --> Healthcheck 
GET    /coins                    --> List of all coins
POST   /coin                     --> Create a new coin
GET    /votes                    --> List all coins
GET    /votes/calculate          --> Calculate votes from coins
POST   /vote/:coin/:value        --> Make a new vote
GET    /ws                       --> Websocket route from this api
```
> If you prefer you can import the routes file from [Insomnia](https://insomnia.rest/download) from this API [click here](https://github.com/gabrielrab/cryptocurrencies-votes/tree/main/assets/cryptocurrencies-insomnia.json)

### Websockets
For real-time communication, this API uses the Websockets protocol. Every time a valid vote is made, the API sends a message to the socket containing the information about the vote performed.

To test this functionality you can copy the code below in your browser's console tab:
```
var ws = new WebSocket("ws://159.223.104.144/ws")
ws.onmessage = (content) => {
    console.log(content.data)
}
```
> If running locally, you can try replacing the socket url following this pattern `ws://localhost:{your_api_port}/ws`