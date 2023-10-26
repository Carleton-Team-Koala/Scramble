import React, { useState } from "react";
import "./Tile.css";

// assign x-y coordinates to each tile and update the tilePositions object

export default function Tile({ letter }) {

    const [visible, setVisible] = useState(true);

    const handleDragStart = (e) => {
        e.dataTransfer.setData("letter", letter);
    };

    const handleDragEnd = () => {
        setVisible(false);
    };
    
    return (
        <div 
            className="tile"
            draggable="true"
            onDragStart={handleDragStart}
            onDragEnd={handleDragEnd}
            style={{ visibility: visible ? "visible" : "hidden" }}
        >
            {letter}
        </div>
    );
}