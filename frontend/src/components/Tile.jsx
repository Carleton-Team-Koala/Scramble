import React, { useState } from "react";
import Draggable from "react-draggable";
import "./Tile.css";

export default function Tile({ letter, tilePositions, updateTilePositions }) {

    const [position, setPosition] = useState({ x: 0, y: 0 });

    const handleStop = (event, { x, y }) => {
        const cellX = Math.round(x / 25) * 25;
        const cellY = Math.round(y / 25) * 25;
        setPosition({ x: cellX, y: cellY });

        const cellI = Math.round(cellY / 25);
        const cellJ = Math.round(cellX / 25);
        const newTilePositions = { ...tilePositions, [`${cellI}-${cellJ}`]: letter };
        updateTilePositions(newTilePositions);
        console.log(tilePositions);
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