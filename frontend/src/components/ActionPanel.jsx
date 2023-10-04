import React, { useState } from 'react';
import './ActionPanel.css';

const ActionPanel = (props) => {

    return (
        <div className="button-container">
            <button>Resign</button>
            <button>Replace</button>
            <button>Submit</button>
            <button>Rearrange</button>
            <button>Remaining</button>
        </div>
      );
    }

export default ActionPanel;