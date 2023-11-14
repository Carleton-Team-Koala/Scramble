import React, { useState } from "react";
import { Link } from "react-router-dom";
import "../css/Welcome.css";
import Popup from "./Popup";

export let player1 = "";
export let player2 = "";
export let gameID = "";
export const baseURL = "http://localhost:8080/"
let frontendURL = "/play/";

export const createGame = () => {
  const url = baseURL + "newgame/";
  console.log(url);
  player1 = document.getElementById("username").value;
  console.log(player1);
  fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ playerName: player1 })
  })
    .then(response => response.json())
    .then(data => {
      if (data.valid) {
        gameID = data.gameID;
        frontendURL += gameID;
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

export default function Welcome() {
  const [popup, setPopup] = useState(false);

  const alertClick = () => {
    alert("This functionality is not supported yet!");
  };

  return (
    <div className="welcome-container">
      <button onClick={()=>setPopup(true)}>New Game</button>
      <button onClick={alertClick}>Load Game</button>
      <button onClick={alertClick}>Join Game</button>
      <Popup trigger={popup} setTrigger={setPopup} onSubmit={createGame}></Popup>
    </div>
  );
};