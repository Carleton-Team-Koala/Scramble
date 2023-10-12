import React from "react";
import { Link } from "react-router-dom";
import "./Welcome.css";

export default function welcome() {

  const alertClick = () => {
    alert("This functionality is not supported yet!");
  };

  return (
    <div className="welcome-container">
      <Link to="/play/gameId"><button>New Game</button></Link>
      <button onClick={alertClick}>Load Game</button>
      <button onClick={alertClick}>Join Game</button>
    </div>
  );
};