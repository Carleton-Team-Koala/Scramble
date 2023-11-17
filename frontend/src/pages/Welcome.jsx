import React, { useState, useEffect } from 'react';
import "../css/Welcome.css";
import Popup from "./Popup";

export const baseURL = "http://localhost:8080/"
let frontendURL = "/play/";

function createGame() {
  /**
   * Sets up game ID in sessionStorage, making it tab-unique.
   * Sends a POST request to the backend to create a new game.
   */
  const url = baseURL + "newgame/";
  const player = sessionStorage.getItem('playerName');
  fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ playerName: player })
  })
    .then(response => response.json())
    .then(data => {
      if (data.valid) {
        sessionStorage.setItem('gameId', data.gameID);
        frontendURL += data.gameID;
        console.log(data.gameID);
      }
      else {
        alert("The game could not be started at the moment!");
      }
    })
    .catch(error => {
      alert(error);
      console.error("Error: ", error);
    })
}

function joinGame() {
  /**
   * Uses sessionStorage to get player name and the ID of the game they would like to join.
   * Sends a POST request to the backend to join a game.
   */
  const url = baseURL + "joingame/";
  const player = sessionStorage.getItem('playerName'); // Get player name from sessionStorage
  const gameID = sessionStorage.getItem('gameId'); // Get game ID from sessionStorage

  if (!player || !gameID) {
    alert("Player name or game ID is missing.");
    return; // Exit the function if the necessary data is missing
  }

  frontendURL += gameID;
  fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ playerName: player })
  })
    .then(response => response.json())
    .then(data => {
      if (data.valid) {
        console.log("Great success!");
      }
      else {
        alert("This gameID doesn't exist or this game could not be joined at the moment!");
      }
    })
    .catch(error => {
      alert(error);
      console.error("Error: ", error);
    })
}

export default function Welcome() {
  const [newGamePopup, setNewGamePopup] = useState(false);
  const [joinGamePopup, setJoinGamePopup] = useState(false);

  const alertClick = () => {
    alert("This functionality is not supported yet!");
  };

  return (
    <div className="welcome-container">
      <button onClick={() => setNewGamePopup(true)}>New Game</button>
      <button onClick={alertClick}>Load Game</button>
      <button onClick={() => setJoinGamePopup(true)}>Join Game</button>
      <Popup type='newGame' trigger={newGamePopup} setTrigger={setNewGamePopup} onSubmit={createGame}></Popup>
      <Popup type='joinGame' trigger={joinGamePopup} setTrigger={setJoinGamePopup} onSubmit={joinGame}></Popup>
    </div>
  );
};