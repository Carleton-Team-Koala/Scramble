import React from 'react';
import { Link } from "react-router-dom";
import '../css/Popup.css';

export default function Popup(props) {

    const handleSubmit = () => {
        const username = document.getElementById('playerName').value;
        sessionStorage.setItem('playerName', username);

        if (props.type === 'joinGame') {
            const gameID = document.getElementById('gameId').value;
            sessionStorage.setItem('gameId', gameID);
        }

        props.onSubmit();
    };

    return (props.trigger) ? (
        <div className='popup'>
            <button className='close-btn' onClick={() => props.setTrigger(false)}>Close</button>
            <div className='inner-container'>
                <h1>Enter Username:</h1>
                <input type='text' id='playerName'></input>

                {props.type === 'joinGame' && (
                    <div>
                        <h1>Enter Game ID:</h1>
                        <input type='text' id='gameId'></input>
                    </div>
                )}

                <Link to="/play">
                    <button className='submit-btn' onClick={handleSubmit}>Submit</button>
                </Link>
            </div>
        </div>
    ) : "";
}
