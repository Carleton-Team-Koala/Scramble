import React from "react";
import { useNavigate } from "react-router-dom";
import "../css/WaitingRoom.css";
import "../css/App.css";
import { baseURL } from "./Welcome";

export default function Room() {

  const navigate = useNavigate();

  const startGame = () => {
    console.log("Start game button clicked!");
    let gameID = sessionStorage.getItem('gameId');
    let url = baseURL + "startgame/" + gameID + "/"
    fetch(url, {
      method: "GET",
      headers: {
        "Content-Type": "application/json"
      }
    })
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.valid) {
          console.log("Great success!");
          navigate(`/play/${gameID}`);
        }
      })
      .catch(error => {
        alert(error);
        console.error("Error: ", error);
      })
  };

  return (
    <div className="App">
      <div className="room">
        <button onClick={startGame}>
          Start Game
        </button>
      </div>
    </div>
  )
}
