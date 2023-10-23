import React from "react";
import { Link } from "react-router-dom";
import "./Welcome.css";

export let gameID = "";
export let player = "";
export const baseURL = "http://localhost:8080"

export default function welcome() {
  const url = baseURL + "/newgame/";
  const alertClick = () => {
    alert("This functionality is not supported yet!");
  };

  const createGame = () => {
    console.log(url);
    fetch(url, {
      method: "GET"
    })
      .then(response => response.json())
      .then(data => {
        gameID = data.GameID;
        player = data.Players[0].name;
        console.log(gameID);
        console.log(player);
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