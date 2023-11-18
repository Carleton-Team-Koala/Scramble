import React, { useState, useEffect } from 'react';
import "../css/Welcome.css";
import Popup from "./Popup";

export const baseURL = "http://localhost:8080/";

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

function joinGame() {
  const player = sessionStorage.getItem('playerName'); // Get player name from sessionStorage
  const gameID = sessionStorage.getItem('gameId'); // Get game ID from sessionStorage
  const url = baseURL + "joingame/" + gameID + "/";

  if (!player || !gameID) {
    alert("Player name or game ID is missing.");
    return; // Exit the function if the necessary data is missing
  }

  fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ playerName: player })
  })
  .then(response => {
    if (response.ok) {
      // If the response status code is 200
      console.log("Joined game successfully!");
    } else {
      // If the response status code is not 200
      console.error(`Failed to join game: ${response.status}`);
      throw new Error(`HTTP error! status: ${response.status}`);
    }
  })
  .catch(error => {
    alert(error);
    console.error("Error: ", error);
  })
};

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