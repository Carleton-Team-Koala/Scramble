import React, { useState } from "react";
import "./Tile.css";

export default function Tile({ letter, id }) {

    const [visible, setVisible] = useState(true);

    const handleDragStart = (e) => {
        e.dataTransfer.setData("letter", letter);
        e.dataTransfer.setData("id", id);
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