import React, { useState } from "react";
import { Link } from "react-router-dom";
import "./Welcome.css";

export let player = "";
export let gameID = "";
export const baseURL = "http://localhost:8080/"
let frontendURL = "/play/";

export const createGame = () => {
  const url = baseURL + "newgame/";
  console.log(url);
  fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
  },
    body: JSON.stringify({playerName : "John"})
  })
    .then(response => response.json())
    .then(data => {
      console.log(data);
      gameID = data;
      frontendURL += gameID;
      console.log(gameID);
      console.log(frontendURL);
    })
    .catch(error => {
      alert(error);
      console.error("Error: ", error);
    })
}

export default function welcome() {
  const alertClick = () => {
    alert("This functionality is not supported yet!");
  };

  return (
    <div className="welcome-container">
      <Link to="/play"><button onClick={createGame}>New Game</button></Link>
      <button onClick={alertClick}>Load Game</button>
      <button onClick={alertClick}>Join Game</button>
    </div>
  );
};

const enterName = () => {
  return (
    <form>
      <label for="fname">Player Name:</label>
      <input type="text" id="fname"></input>
    </form>
  )
}

const enterGameID = () => {
  return (
    <form>
      
    </form>
  )
}