import React from 'react';
import { Link } from "react-router-dom";
import '../css/Popup.css';

export default function Popup(props) {
    return (props.trigger) ? (
        <div className='popup'>
            <button className='close-btn' onClick={() => props.setTrigger(false)}>Close</button>
            <div className='inner-container'>
                <h1>Enter Username:</h1>
                <input type='text' id='username'></input>
                <Link to="/play"><button className='submit-btn' onClick={props.onSubmit}>Submit</button></Link>
            </div>
        </div>
    ) : ""
}