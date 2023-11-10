import React from "react";
import "../css/Tile.css";

export default function Tile({ letter, id, draggable = true }) {

    /**
     * Separate component for each tile.
     * ID is used for letterupdates, draggable determines whether we can drag the tile across the board.
     * @param {*} e 
     */

    const handleDragStart = (e) => {
        /**
         * Native HTML drag and drop API to transfer data.
         */
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