import React from "react";
import "../css/WaitingRoom.css";
import { baseURL, gameID, player1, player2 } from "./Welcome";

export default function Room({ setHand, setTilebag, setisgamestarted }) {

  const startGame = () => {
    let url = baseURL + "startgame/" + gameID + "/"
    fetch(url, {
      method: "GET",
      headers: {
        "Content-Type": "application/json"
      }
    })
      .then(response => response.json())
      .then(data => {
        if (data.valid) {
          setHand(data.gameState.Players[player1].hand);
          setTilebag(data.gameState.LetterDistribution);
          setisgamestarted(true); // Set game started state to true
        }
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
