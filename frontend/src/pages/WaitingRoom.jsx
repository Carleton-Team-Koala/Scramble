import React, { useState } from "react";
import { Link } from "react-router-dom";
import "../css/WaitingRoom.css";
import { baseURL, gameID } from "./Welcome";

export default function Room({ initialhand, setinitialhand, setisgamestarted }) {

  const startGame = () => {
    console.log(gameID);
    let url = baseURL + "startgame/" + gameID + "/"
    fetch(url, {
      method: "GET",
      headers: {
        "Content-Type": "application/json"
      }
    })
      .then(response => response.json())
      .then(data => {
        setinitialhand(data.Players.John.hand);
        setisgamestarted(true); // Set game started state to true
      })
      .catch(error => {
        alert(error);
        console.error("Error: ", error);
      })
  }

  return (
    <div className="room">
      <button onClick={startGame}>
        Start Game
      </button>
    </div>
  )
}
