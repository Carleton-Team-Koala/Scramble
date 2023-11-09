import React from "react";
import "../css/Tile.css";

export default function Tile({ letter, id, draggable = true }) {

    const handleDragStart = (e) => {
        e.dataTransfer.setData("letter", letter);
        e.dataTransfer.setData("id", id);
    };

    return (
        <div
            className="tile"
            {...(draggable ? { draggable: true, onDragStart: handleDragStart } : {})}
        >
            {letter}
        </div>
    );
}