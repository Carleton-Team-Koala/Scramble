import React from "react";
import Draggable from "react-draggable";
import "./Tile.css";

export default function Tile({ letter }) {
    return (
        <Draggable>
            <div className="tile">
                {letter}
            </div>
        </Draggable>
    )
}