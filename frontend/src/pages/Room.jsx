import React, { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import "../css/WaitingRoom.css";
import "../css/App.css";
import { baseURL } from "./Welcome";

export default function Room() {

  const gameID = sessionStorage.getItem('gameId');
  const navigate = useNavigate();

  const checkGameStarted = () => {
    const url = baseURL + "/getgamestate/" + gameID + "/";

    fetch(url, {
      method: "GET",
      headers: {
        "Content-Type": "application/json"
      }
    })
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.GameStarted) {
          console.log("Game started!");
          navigate(`/play/${gameID}`);
        }
      })
      .catch(error => {
        console.error("Error: ", error);
      });
  };

  useEffect(() => {
    const intervalId = setInterval(() => {
      checkGameStarted();
    }, 2000); // Check every 2 seconds

    return () => clearInterval(intervalId); // Clean up on unmount
  }, []);

  const startGame = () => {
    console.log("Start game button clicked!");
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
