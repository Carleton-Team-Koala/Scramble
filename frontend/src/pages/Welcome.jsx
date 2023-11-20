import React, { useState, useEffect } from 'react';
import "../css/Welcome.css";
import Popup from "./Popup";

export const baseURL = "http://localhost:8080/";

/**
 * Creates a new game by sending a POST request to the server.
 * 
 * @returns {Promise<string|null>} A promise that resolves with the game ID if the game is created successfully,
 *                                or null if there is an error or the game cannot be started.
 */
function createGame() {
  const url = baseURL + "newgame/";
  const player = sessionStorage.getItem('playerName');

  // Return the fetch promise
  return fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ playerName: player })
  })
  .then(response => response.json())
  .then(data => {
    if (data.valid) {
      sessionStorage.setItem('gameId', data.gameID); // Store game ID in sessionStorage
      return data.gameID;  // Resolve with gameID
    } else {
      alert("The game could not be started at the moment!");
      return null;  // Resolve with null
    }
  })
  .catch(error => {
    alert(error);
    console.error("Error: ", error);
    return null;  // Resolve with null in case of error
  });
};

/**
 * Joins a game by sending a POST request to the server with the player name and game ID.
 * If the necessary data is missing, wrong gameID, player name has already been used, or the game has already begun, an alert is displayed and the function exits.
 * Any caught errors are alerted and logged to the console.
 */
function joinGame() {
  const player = sessionStorage.getItem('playerName'); // Get player name from sessionStorage
  const gameID = sessionStorage.getItem('gameId'); // Get game ID from sessionStorage
  const url = baseURL + "joingame/" + gameID + "/";

  if (!player || !gameID) {
    alert("Player name or game ID is missing.");
    return null; // Exit the function if the necessary data is missing
  }

  // Return the fetch promise
  return fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ playerName: player })
  })
  .then(response => response.json())
  .then(data => {
    if (data.valid) {
      return data.gameID;  // Resolve with gameID
    } else {
      alert(data.message);
      return null;  // Resolve with null
    }
  })
  .catch(error => {
    alert(error);
    console.error("Error: ", error);
    return null;  // Resolve with null in case of error
  });
};

export default function Welcome() {
  //state variables to control which Popup will appear when
  const [newGamePopup, setNewGamePopup] = useState(false);
  const [joinGamePopup, setJoinGamePopup] = useState(false);

  //render the HTML elements
  return (
    <div className="welcome-container">
      <button key="new-btn" onClick={() => setNewGamePopup(true)}>New Game</button>
      <button key="join-btn" onClick={() => setJoinGamePopup(true)}>Join Game</button>
      <Popup type='newGame' trigger={newGamePopup} setTrigger={setNewGamePopup} onSubmit={createGame}></Popup>
      <Popup type='joinGame' trigger={joinGamePopup} setTrigger={setJoinGamePopup} onSubmit={joinGame}></Popup>
    </div>
  );
};