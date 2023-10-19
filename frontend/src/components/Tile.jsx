import React, { useState } from "react";
import Draggable from "react-draggable";
import "./Tile.css";

export default function Tile({ letter }) {

    const [position, setPosition] = useState({ x: 0, y: 0 });

    const handleStop = (event, { x, y }) => { // to be fixed, these coordinates are not perfect
        const cellX = Math.round(x / 25) * 25;
        const cellY = Math.round(y / 25) * 25;
        setPosition({ x: cellX, y: cellY });
    };

    const handleClick = event => { // without this we won't be able to drop the tile
        handleStop(event, position);
    };

    return (
        <Draggable onStop={handleStop} position={position}>
            <div className="tile" onClick={handleClick}>
                {letter}
            </div>
        </Draggable>
    )
}