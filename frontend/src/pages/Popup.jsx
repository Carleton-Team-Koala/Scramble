import { React } from 'react';
import { useNavigate } from "react-router-dom";
import '../css/Popup.css';

/**
 * Renders a popup component.
 * @param {Object} props - The component props.
 * @param {boolean} props.trigger - Determines whether the popup is triggered or not.
 * @param {string} props.type - The type of the popup ('joinGame' or other).
 * @param {Function} props.setTrigger - Function to set the trigger state.
 * @param {Function} props.onSubmit - Function to handle form submission.
 * @returns {JSX.Element} The rendered popup component.
 */
export default function Popup(props) {

    const navigate = useNavigate();

    const handleSubmit = async () => {
        const player = document.getElementById('playerName').value;
        sessionStorage.setItem('playerName', player);
    
        let gameId = '';
    
        if (props.type === 'joinGame') {
            gameId = document.getElementById('gameId').value;
            sessionStorage.setItem('gameId', gameId);
            gameId = await props.onSubmit();
        } else {
            // Await the createGame function to complete and get the gameId
            gameId = await props.onSubmit();
        }
        
        if (gameId != '' && gameId != null) {
            navigate(`/room/${gameId}`);
        }
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
                
                <button className='submit-btn' onClick={handleSubmit}>Submit</button>

            </div>
        </div>
    ) : "";
}
