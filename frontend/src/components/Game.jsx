import React, { useState } from 'react';
import Board from "./Board";
import ActionPanel from "./ActionPanel";
import Infoboard from "./Infoboard";
import Tile from './Tile';
import './Game.css';

let tilePositions = [
  { letter: "I", xLoc: 4, yLoc: 10 },
  { letter: "J", xLoc: 6, yLoc: 7 }
];

let tiles_coded = []; // hardcoding this data for now
for (let i = 0; i < 7; i++) {
  tiles_coded.push(
  <Tile
    key={i}
    letter='A'
  />); // will be passed by the server in the future
}

export default function Game() {

  const [boardState, setBoardState] = useState({});
  const [tiles, setTiles] = useState(tiles_coded);

  function handleTileDrop(cellKey, letter) {
    setBoardState(prevState => ({
      ...prevState,
      [cellKey]: letter
    }));
    console.log(boardState);
  };

  function handleTileDrag(tile) {
    setTiles(prevTiles => prevTiles.filter(t => t !== tile));
  };

  return (
    <div>
      <div className="board-score">
        <Board
          tilePositions={tilePositions}
          onTileDrop={handleTileDrop}
        />
        <Infoboard />
      </div>
      <ActionPanel
        tiles={tiles}
        tilePositions={tilePositions}
        onTileDrag={handleTileDrag}
      />
    </div>
  );
};