import React from 'react';
import './ActionPanel.css';
import shuffleImage from '../assets/shuffle.jpg';
import refreshImage from '../assets/refresh.jpg';
import { baseURL, gameID, player } from "../Welcome"
import { boardTiles } from './Board';

const ActionPanel = ({ tilesAp, shuffle, logger }) => {

    const submit = () => {
        console.log(boardTiles);
        let data = []
        console.log(boardTiles);
        for (const [key, value] of Object.entries(boardTiles)) {
            data.push({letter: key, xLoc: value[0], yLoc: value[1]});
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
            body: JSON.stringify({ playerName : "John", updates : data})
        })
            .then(response => response.json())
            .then(data => {
                alert(data);
            })
            .catch(error => {
                alert(error);
                console.log("Error: ", error);
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
                    {tilesAp}
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
                <button className="button-ap" onClick={logger}>Skip</button>
                <button className="button-ap">Swap</button>
                <button className="button-ap submit-button" onClick={submit}>Submit</button>
            </div>
        </div>
    );
}

export default ActionPanel;
