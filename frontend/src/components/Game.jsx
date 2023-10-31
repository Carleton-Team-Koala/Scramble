import React, { useState, useEffect } from 'react';
import Board from "./Board";
import ActionPanel from "./ActionPanel";
import Infoboard from "./Infoboard";
import Tile from './Tile';
import './Game.css';

export default function Game() {

  const [letterUpdates, setLetterUpdates] = useState({});
  const [tiles, setTiles] = useState(
    Array.from({ length: 7 }, (_, i) => ({ // hardcoding this data for now
      id: i,
      letter: 'A',
      position: 'ActionPanel' // initial position
    }))
  );

  function handleTileDrop(id, cellKey, letter) {
    id = Number(id);

    setLetterUpdates(prevState => ({
      ...prevState,
      [id]: [cellKey, letter]
    }));

    setTiles(prevTiles =>
      prevTiles.map(tile =>
        tile.id === id ? { ...tile, position: 'Board' } : tile
      )
    );
  };

  return (
    <div>
      <div className="board-score">
        <Board
          letterUpdates={letterUpdates}
          onTileDrop={handleTileDrop}
        />
        <Infoboard />
      </div>
      <ActionPanel
        tilesAp={tiles.map(tile => {
          if (tile.position === 'ActionPanel') {
            return <Tile key={tile.id} letter={tile.letter} id={tile.id} />;
          } else {
            return <div key={tile.id} className="tile-placeholder"></div>;
          }
        })}
      />
    </div>
  );
};