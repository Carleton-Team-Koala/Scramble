import React from "react";
import "../css/Cell.css";

export default function Cell({ i, j, cellStyle, children, tile, onTileDrop }) {

    const key = `${i}-${j}`; // key for the tilePositions object

    const handleDrop = (e) => {
        e.preventDefault();
        const droppedLetter = e.dataTransfer.getData("letter");
        const droppedId = e.dataTransfer.getData("id");
        onTileDrop(droppedId, key, droppedLetter);
    };

    const handleDragOver = (e) => {
        e.preventDefault();  // This is important to allow dropping
    };

    return (
        <div
            className={`cell ${cellStyle}`}
            onDrop={tile ? undefined : handleDrop}
            onDragOver={tile ? undefined : handleDragOver}
        >
            <div className="cell-content">
                {tile ? tile : children}
            </div>
        </div>
    );
};