import React from 'react';
import '../css/Rules.css';

/**
 * Renders the Rules component.
 * @param {Object} props - The component props.
 * @param {boolean} props.isRulesOpen - Indicates whether the rules are open or closed.
 * @param {function} props.setIsRulesOpen - Callback function to set the state of isRulesOpen.
 * @returns {JSX.Element} The rendered Rules component.
 */
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
