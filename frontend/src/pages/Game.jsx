import React, { useState, useEffect } from 'react';
import Board from "../components/Board";
import ActionPanel from "../components/ActionPanel";
import Infoboard from "../components/Infoboard";
import Tile from '../components/Tile';
import '../css/Game.css';
import { baseURL, gameID, player } from "./Welcome"

export default function Game({ initialhand }) {
  const [letterUpdates, setLetterUpdates] = useState({});
  const [tiles, setTiles] = useState(
    Array.from({ length: 7 }, (_, i) => ({ // hardcoding this data for now
      id: i,
      letter: initialhand[i] === 'BLANK' ? '' : initialhand[i],
      position: 'ActionPanel' // initial position
    })
    ));

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
    console.log(letterUpdates);
  };

  const shuffle = () => {
    let indices = [0, 1, 2, 3, 4, 5, 6]
    let tilesCopy = [...tiles];
    for (let i = 0; i < 7; i++) {
      let loc = Math.floor(Math.random() * indices.length);
      console.log(indices[loc]);
      tilesCopy[indices[loc]] = tiles[i];
      // indices.splice(loc, 1); 
      indices = indices.filter(value => value !== loc);
      console.log(indices);
    }
    setTiles(tilesCopy);
  }

  const submit = () => {
    let data = []
    console.log(letterUpdates);
    for (const [key, value] of Object.entries(letterUpdates)) {
      let locs = value[0].split("-");
      data.push({ letter: value[1], xLoc: Number(locs[1]), yLoc: Number(locs[0]) });
      console.log(locs[0]);
      console.log(locs[1]);
    }
    console.log(data);

    const url = baseURL + gameID + "/updategame/"
    console.log(url);
    console.log(data);
    // const data = JSON.stringify({ playerName: player, updates: tilePositions })
    fetch(url, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({ playerName: "John", updates: data })
    })
      .then(response => response.json())
      .then(data => {
        console.log(data);
      })
      .catch(error => {
        alert(error);
        console.log("Error: ", error);
      })
  }

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
        // shuffle={shuffle}
        submit={submit}
      />
    </div>
  );
};