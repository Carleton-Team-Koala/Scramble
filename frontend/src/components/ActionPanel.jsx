import React from 'react';
import '../css/ActionPanel.css';
import shuffleImage from '../assets/shuffle.jpg';
import refreshImage from '../assets/refresh.jpg';

const ActionPanel = ({ tilesAp, shuffle, submit, reset, refresh }) => {

    return (
        <div className="action-panel">
            <div className="hand-container">
                <button className="button-hand" onClick={shuffle}
                    style={{
                        backgroundImage: `url(${shuffleImage})`,
                    }}>
                </button>
                <div className='tile-hand'>
                    {tilesAp}
                </div>
                <button className="button-hand" onClick={reset}
                    style={{
                        backgroundImage: `url(${refreshImage})`,
                        backgroundSize: "30px",
                    }}>
                </button>
            </div>
            <div className="button-container">
                <button className="button-ap">Resign</button>
                <button className="button-ap">Skip</button>
                <button className="button-ap" onClick={refresh}>Refresh</button>
                <button className="button-ap submit-button" onClick={submit}>Submit</button>
            </div>
        </div>
    );
}

export default ActionPanel;
