import React, { useState } from 'react';
import './ActionPanel.css';
import shuffleImage from '../assets/shuffle.jpg';
import refreshImage from '../assets/refresh.jpg';
import tilePositions from "./Game"

const ActionPanel = ({ tiles }) => {

    const submit = () => {
        const baseURL = "http://languages:8000/{gameID}/updategame/"
        const url = baseURL + "/"
        const data = JSON.stringify(tilePositions)
        fetch(url, {
          method: "POST",
          Header: {
            "Content-Type": "application/json"
          },
          body: JSON.stringify({
            playerName : playerName,
            updates: [data]
          })
        })
      }

    return (
        <div className="action-panel">
            <div className="hand-container">
                <button className="button-hand"
                    style={{
                        backgroundImage: `url(${shuffleImage})`,
                    }}>
                </button>
                <div className='tile-hand'>
                    {tiles}
                </div>
                <button className="button-hand"
                    style={{
                        backgroundImage: `url(${refreshImage})`,
                        backgroundSize: "30px",
                    }}>
                </button>
            </div>
            <div className="button-container">
                <button className="button-ap">Resign</button>
                <button className="button-ap">Skip</button>
                <button className="button-ap">Swap</button>
                <button className="button-ap submit-button" onClick={submit}>Submit</button>
            </div>
        </div>
    );
}

export default ActionPanel;