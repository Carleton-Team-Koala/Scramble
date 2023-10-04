import React, { useState } from 'react';
import Board from "./Board";
import ActionPanel from "./ActionPanel";
import Scoreboard from "./Scoreboard";
import Hand from './Hand';

function Game() {

  return (
    <div>
      <Board /> 
      <Hand/>
      <ActionPanel />
      <Scoreboard />
    </div>
  );
};

export default Game;