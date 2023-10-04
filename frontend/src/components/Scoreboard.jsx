import React, { useState } from 'react';
import './Scoreboard.css';

const Scoreboard = (props) => {
    return (
        <div className="scoreboard">
            <h2>Scoreboard</h2>
            <div className="scoreboard-cell">
                Player1: 0
            </div>
            <div className="scoreboard-cell">
                Player2: 0
            </div>
        </div>
    );
}

export default Scoreboard;