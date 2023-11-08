import React from "react";
import "../css/Tile.css";

export default function Tile({ letter, id }) {

    const handleDragStart = (e) => {
        e.dataTransfer.setData("letter", letter);
        e.dataTransfer.setData("id", id);
    };

    return (
        <div
            className="tile"
            draggable="true"
            onDragStart={handleDragStart}
        >
            {letter}
        </div>
    );
}