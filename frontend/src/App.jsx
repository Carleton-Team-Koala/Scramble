import React, { useState } from "react";
import Game from "./pages/Game";
import Room from "./pages/WaitingRoom"
import "./css/App.css";

function App() {

  const [hand, setHand] = useState(['BLANK', 'B', 'C', 'D', 'E', 'A', 'G']); // hand, gets rendered in the action panel
  const [tilebag, setTilebag] = useState({
    'A': 0, 'B': 0, 'C': 0, 'D': 0, 'E': 0, 'F': 0, 'G': 0,
    'H': 0, 'I': 0, 'J': 0, 'K': 0, 'L': 0, 'M': 0, 'N': 0,
    'O': 0, 'P': 0, 'Q': 0, 'R': 0, 'S': 0, 'T': 0, 'U': 0,
    'V': 0, 'W': 0, 'X': 0, 'Y': 0, 'Z': 0, 'BLANK': 0
  }); // tilebag, gets rendered on the infoboard
  const [isGameStarted, setIsGameStarted] = useState(false);

  if (isGameStarted) {
    return (
      <div className="App" >
        <Game hand={hand} setHand={setHand} tilebag={tilebag} setTilebag={setTilebag} />
      </div>
    )
  } else {
    return (
      <div className="App">
        <Room hand={hand} setHand={setHand} setisgamestarted={setIsGameStarted} />
      </div>
    )
  }
};

export default App;
