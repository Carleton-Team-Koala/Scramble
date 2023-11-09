import React, { useState } from "react";
import Game from "./pages/Game";
import Room from "./pages/WaitingRoom"
import "./css/App.css";

function App() {

  const [initialHand, setInitialHand] = useState(['BLANK', 'B', 'C', 'D', 'E', 'A', 'G']);
  const [isGameStarted, setIsGameStarted] = useState(false);

  if (isGameStarted) {
    return (
      <div className="App" >
        <Game initialhand={initialHand} setInitialHand={setInitialHand} />
      </div>
    )
  } else {
    return (
      <div className="App">
        <Room initialhand={initialHand} setinitialhand={setInitialHand} setisgamestarted={setIsGameStarted} />
      </div>
    )
  }
};

export default App;
