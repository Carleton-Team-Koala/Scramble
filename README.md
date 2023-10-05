This is Scramble, a word game. Any situations and similarities are coincidental.

To run the skeleton of our frontend app, type

`cd frontend`

`npm install`

`npm run dev`

To run the skeleton of our backend app on the terminal, go to the Scramble/backend/cmd/main directory and use the following commands:

`go run main.go`

The app will be available on http://localhost:8080/

If you want to run it with Docker, go to the Scramble directory and use the following commands:

`docker build --rm -t scramble .`

`docker run -p 8080:8080 scramble`