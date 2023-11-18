# Scramble

## Abstract
This is Scramble, a word game based off of Scrabble. 

### Contributers
|Name|Email|  
|----|-----|  
|Roo Case|caser@carleton.edu|
|Sunny Kim|kims3@carleton.edu|
|John Win|winj@carleton.edu|
|Artem Yushko|yushkoa@carleton.edu|

## Contents
- [Description](#description)
- [Instructions](#instructions)
    - [Front End](#frontend)
    - [Back End](#backend)
- [Tests](#tests)
- [Credits](#credits)

## Description

## Instructions

### App
To run the whole program, go to the `Scramble` directory and use the following commands on the terminal:

```terminal
> docker compose build
> docker compose up
```

Wait until you see the following command:

```terminal
scramble-game-1       | Server is running on :8080
```

This ensures that the game is running. 

The app will be available on http://localhost:3000/

Once you go into the url, you will be able to access 



### Frontend
To run the skeleton of our frontend app, go to the `Scramble` directory and use the following commands on the terminal:
```terminal
> cd frontend
> npm install
> npm run dev
```

### Backend
To run the skeleton of our backend app on the terminal, go to the `Scramble` directory and use the following commands on the terminal:

```terminal
> cd app/backend/cmd/main
> go run main.go
```

The app will be available on http://localhost:8080/


### Languages
The languages section runs separately from the backend part of the game. 

For now, the language section only works through Docker. 

~~To run the skeleton of the language package on the terminal, go to the `Scramble` directory and use the following commands on the terminal:~~

 
<!-- ```terminal
> cd app/languages/cmd/main
> go run main.go
``` -->


~~The app will be available on http://localhost:8000/~~




## Player Leaderboard
Starter Code for player leaderboard
Add columns for player1 score and player2 score
Run through a game and get score for each player 
Order by playerTotal Score
Frontend
leaderboard.html: displays a static leaderboard with same background and styles
as main project. 

## Tests

## Credits