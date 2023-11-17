import React from 'react';
import '../css/Rules.css';

export default function Rules(props) {
    return (props.isRulesOpen) ? (
        <div className='rules'>
            <div className='rules-container'>
                <button className='close-btn' onClick={() => props.setIsRulesOpen(false)}>Close</button>
                <p className='text-container'>
                    This is a rule. <br/> 
                    This is another rule. <br/> 
                    This is a 3rd rule. 
                </p>
            </div>
        </div>
    ) : <button className='open-btn' onClick={() => props.setIsRulesOpen(true)}>Open</button>
}