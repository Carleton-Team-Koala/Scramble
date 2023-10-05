import React, { useState } from 'react';
import Board from "./Board";
import ActionPanel from "./ActionPanel";
import Infoboard from "./Infoboard";
import './Game.css';

function Game() {

  return (
    <div>
      <div className="board-score">
        <Board /> 
        <Infoboard />
      </div>
      <ActionPanel />
    </div>
  );
};

export default Game;