import React, { useState } from 'react';
import Board from "./Board";
import ActionPanel from "./ActionPanel";
import Infoboard from "./Infoboard";
import Tile from './Tile';
import './Game.css';

function Game() {

  // const [tilePositions, setTilePositions] = useState({}); // function for placing the tiles onto the board

  // function updateTilePositions(change) { // wrapper to be passed to the action panel
  //   setTilePositions(change);
  // };

  let tiles = []; // hardcoding this data for now
  for (let i = 0; i < 7; i++) {
    tiles.push(<Tile
      key={i}
      letter='A'
      tilePositions={tilePositions}
      updateTilePositions={updateTilePositions}
    />); // will be passed by the server in the future
  }

  let tilePositions = [
    {letter: "I", xLoc: 4, yLoc: 10},
    {letter: "J", xLoc: 6, yLoc: 7}
   ]; // hardcoding the cell positions for now as well

  return (
    <div>
      <div className="board-score">
        <Board
          tilePositions={tilePositions}
        />
        <Infoboard />
      </div>
      <ActionPanel
        tiles={tiles}
      />
    </div>
  );
};

export default Game;