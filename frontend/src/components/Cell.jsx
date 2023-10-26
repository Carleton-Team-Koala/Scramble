import React, { useState } from "react";
import Tile from "./Tile";
import "./Cell.css";

export default function Cell({ i, j, cellStyle, children, onTileDrop }) {

    const [currentChild, setCurrentChild] = useState(null);
    const key = `${i}-${j}`; // key for the tilePositions object

    const handleDrop = (e) => {
        e.preventDefault();
        const droppedLetter = e.dataTransfer.getData("letter");
        console.log("dropped: letter:", droppedLetter, "key:", key);
        setCurrentChild(<Tile letter={droppedLetter} />);
        onTileDrop(key, droppedLetter);
    };

    const handleDragOver = (e) => {
        e.preventDefault();  // This is important to allow dropping
    };

    return (
        <div
            className={`cell ${cellStyle}`}
            onDrop={handleDrop}
            onDragOver={handleDragOver}
        >
            <div className="cell-content">
                {currentChild ? currentChild : children}
            </div>
        </div>
    );
};