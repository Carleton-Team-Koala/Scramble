import React, { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import "../css/Room.css";
import "../css/App.css";
import { baseURL } from "./Welcome";

/**
 * Renders the Room component.
 * 
 * @returns {JSX.Element} The rendered Room component.
 */
export default function Room() {
  const gameID = sessionStorage.getItem('gameId'); //in case of refreshing the page, store the gameID in a sessionStorage
  const playerName = sessionStorage.getItem('playerName');
  const navigate = useNavigate();

  //if game has started, send update to the server and navigate to the game page
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
        if (data.GameStarted) {
          navigate(`/play/${gameID}`);
        }
      })
      .catch(error => {
        console.error("Error: ", error);
      });
  };

  //check if game has started
  useEffect(() => {
    const intervalId = setInterval(() => {
      checkGameStarted();
    }, 2000); // Check every 2 seconds

    return () => clearInterval(intervalId); // Clean up on unmount
  }, []);

  /**
   * Notify the server to start the game. 
   * If there is only one player, then return an alert. 
   */
  const startGame = () => {
    let url = baseURL + "startgame/" + gameID + "/"
    fetch(url, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({ playerName: playerName })
    })
      .then(response => response.json())
      .then(data => {
        if (data.valid) {
          navigate(`/play/${gameID}`);
        }
        else {
          console.log(data);
          alert(data.message);
        }
      })
      .catch(error => {
        alert(error);
        console.error("Error: ", error);
      })
  };

  //render the HTML elements
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
