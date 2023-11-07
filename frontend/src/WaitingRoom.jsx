import React, { useState } from "react";
import { Link } from "react-router-dom";
import "./WaitingRoom.css";
import { baseURL, gameID } from "./Welcome";

export let initialHand = ['BLANK', 'B', 'C', 'D', 'E', 'A', 'G'];

const startGame = () => {
  console.log(gameID);
  let url = baseURL + "startgame/" + gameID + "/"
  console.log(url);
  fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json"
  }})
    .then(response => response.json())
    .then(data => {
      initialHand = data.Players.John.hand;
      console.log(initialHand);
    })
    .catch(error => {
      alert(error);
      console.error("Error: ", error);
    })
}

export default function Room() {

    return (
        <div className="room">
          <Link to="/play">
            <button onClick={startGame}>
              Start Game
            </button>
          </Link>
        </div>
    )
}