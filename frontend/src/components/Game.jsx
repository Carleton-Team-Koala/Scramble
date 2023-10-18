import React, { useState } from 'react';
import Board from "./Board";
import ActionPanel from "./ActionPanel";
import Infoboard from "./Infoboard";
import Tile from './Tile';
import './Game.css';

function Game() {

  let tiles = []; // hardcoding this data for now
  for (let i = 0; i < 7; i++) {
    tiles.push(<Tile key={i} />); // will be passed by the server in the future
  }

  let tilePositions = [[4, 10, "I"], [6, 7, "J"]]; // hardcoding the cell positions for now as well

  return (
    <div>
      <div className="board-score">
        <Board tiles={tiles} tilePositions={tilePositions} />
        <Infoboard />
      </div>
      <ActionPanel tiles={tiles} />
    </div>
  );
};

export default Game;