import React from 'react';
import '../css/ActionPanel.css';
import shuffleImage from '../assets/shuffle.jpg';
import refreshImage from '../assets/refresh.jpg';

/**
 * ActionPanel component.
 *
 * @component
 * @param {Object} props - The component props.
 * @param {Array} props.tilesAp - The tiles to display in the hand container.
 * @param {Function} props.shuffle - The function to shuffle the tiles.
 * @param {Function} props.submit - The function to submit the tiles.
 * @param {Function} props.reset - The function to reset the tiles.
 * @param {Function} props.refresh - The function to refresh the tiles.
 * @returns {JSX.Element} The ActionPanel component.
 */
const ActionPanel = ({ tilesAp, shuffle, submit, reset, refresh, skip, resign }) => {

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
                <button className="button-ap" onClick={skip}>Skip</button>
                <button className="button-ap" onClick={refresh}>Refresh</button>
                <button className="button-ap submit-button" onClick={submit}>Submit</button>
            </div>
        </div>
    );
}

export default ActionPanel;
