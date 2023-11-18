import React from "react";
import "../css/WaitingRoom.css";
import { baseURL } from "./Welcome";
import { Link } from "react-router-dom";

export default function Room({ setHand, setTilebag, setisgamestarted }) {
  const [isGameStarted, setIsGameStarted] = useState(false);

  useEffect(() => {
    const interval = setInterval(() => {
      // Your code logic here, this will run every 2 seconds
      console.log('This runs every 2 seconds');

      let playerName = sessionStorage.getItem('playerName');
      let gameID = sessionStorage.getItem('gameId');

      const url = baseURL + "/getgamestate/" + gameID + "/";
      fetch(url, {
        method: "GET",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({ playerName: playerName })
      })
        .then(response => response.json())
        .then(data => {
          setIsGameStarted(data.gameState.gameStarted);
        })
        .catch(error => {
          alert(error);
          console.error("Error: ", error);
        })
    }, 2000); // 2000 milliseconds = 2 seconds

    if(isGameStarted) {
      return (
        <Link>
        
        </Link>
      );
    }

    return () => clearInterval(interval);
  }, []); // Empty dependency array means this effect runs once after the component mounts


  const startGame = () => {
    let player1 = sessionStorage.getItem('playerName');
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
