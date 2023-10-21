import React from "react";
import { Link } from "react-router-dom";
import "./Welcome.css";

export const gameID = "";
export const player = "";

export default function welcome() {

  const alertClick = () => {
    alert("This functionality is not supported yet!");
  };

  const createGame = () => {
    fetch("/newgame/", {
      method: "GET"
    })
      .then(response => response.json())
      .then(data => {
        gameID = data.gameID;
        player = data.playerList[0];
      })
      .catch(error => {
        alert(error);
        console.error("Error: ", error);
      })
  }

  return (
    <div className="welcome-container">
      <Link to="/play/gameId"><button onClick={createGame}>New Game</button></Link>
      <button onClick={alertClick}>Load Game</button>
      <button onClick={alertClick}>Join Game</button>
    </div>
  );
};