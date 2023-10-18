import React, { useState } from 'react';
import './ActionPanel.css';
import shuffleImage from '../assets/shuffle.jpg';
import refreshImage from '../assets/refresh.jpg';

const ActionPanel = ({ tiles }) => {

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
                <button className="button-ap submit-button">Submit</button>
            </div>
        </div>
    );
}

export default ActionPanel;