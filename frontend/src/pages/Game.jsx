import React, { useState, useEffect } from 'react';
import Board from "../components/Board";
import ActionPanel from "../components/ActionPanel";
import Infoboard from "../components/Infoboard";
import Tile from '../components/Tile';
import '../css/Game.css';
import { baseURL, gameID, player1, player2 } from "./Welcome";

const defaultTilebag = {
  'A': 0, 'B': 0, 'C': 0, 'D': 0, 'E': 0, 'F': 0, 'G': 0,
  'H': 0, 'I': 0, 'J': 0, 'K': 0, 'L': 0, 'M': 0, 'N': 0,
  'O': 0, 'P': 0, 'Q': 0, 'R': 0, 'S': 0, 'T': 0, 'U': 0,
  'V': 0, 'W': 0, 'X': 0, 'Y': 0, 'Z': 0, 'BLANK': 0
};

function initializeTiles(initialHand) {
  return Array.from({ length: initialHand.length }, (_, i) => ({
    id: i,
    letter: initialHand[i] === 'BLANK' ? '' : initialHand[i],
    position: 'ActionPanel', // initial position
  }));
};

export default function Game({ initialhand, setInitialHand }) {

  const [scoredLetters, setScoredLetters] = useState({}); // {cellKey: letter}, letters returned by server go here
  const [letterUpdates, setLetterUpdates] = useState({}); // {id: [cellKey, letter]}, gets sent to server on submit
  const [tiles, setTiles] = useState(initializeTiles(initialhand)); // array of tiles, gets rendered on the board and hand
  const [tilebag, setTilebag] = useState(defaultTilebag); // tilebag, gets rendered on the infoboard
  const [p1_score, setp1_score] = useState(0); // scores for both players
  const [p2_score, setp2_score] = useState(0);

  useEffect(() => {
    setTiles(initializeTiles(initialhand));
  }, [initialhand]);

  /**
    * Handles the event when a tile is dropped onto the board.
    * Sets letterUpdates and tiles states.
    *
    * @param {number} id - The ID of the tile that was dropped.
    * @param {string} cellKey - The key of the cell where the tile was dropped.
    * @param {string} letter - The letter on the tile that was dropped.
    */
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

  /** 
   * Parses the board returned by the server.
  */
  function parseBoard(board) {

    let newpos = {};

    for (let i = 0; i < 15; i++) {
      for (let j = 0; j < 15; j++) {
        if (board[i][j] !== '') {
          let cellKey = `${i}-${j}`;
          newpos[cellKey] = board[i][j];
        }
      }
    }

    setScoredLetters(prevScoredLetters => ({
      ...prevScoredLetters,
      ...newpos
    }));

  };

  /**
   * Parses all the updates returned by the server
   */
  function parseUpdates(updates) {

    parseBoard(updates.Board);
    setInitialHand(updates.Players.John.hand);
    setTilebag(updates.LetterDistribution);
    setp1_score(updates.Players.John.score);
    // set score for player 2 here

  };

  /**
   * Submits the current state of the game board.
   *
   * This function iterates over the `letterUpdates` object, which contains the updates to the letters on the board.
   * For each update, it extracts the location and the letter and adds them to the `data` array.
   * The location is split into x and y coordinates, which are converted to numbers.
   * The `data` array is then ready to be sent to the server or processed further.
   */
  const submit = () => {
    let data = []
    for (const [key, value] of Object.entries(letterUpdates)) {
      let locs = value[0].split("-");
      data.push({ letter: value[1], xLoc: Number(locs[0]), yLoc: Number(locs[1]) });
    };
    setLetterUpdates({});
    const url = baseURL + gameID + "/updategame/";
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
        // processing the server response
        console.log(data);
        parseUpdates(data);
      })
      .catch(error => {
        alert(error);
        console.log("Error: ", error);
      })
  };

  return (
    <div>
      <div className="board-score">
        <Board
          letterUpdates={letterUpdates}
          onTileDrop={handleTileDrop}
          scoredLetters={scoredLetters}
        />
        <Infoboard
          tilebag={tilebag}
          p1_score={p1_score}
          p2_score={p2_score}
        />
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