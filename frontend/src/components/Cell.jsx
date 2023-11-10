import React from "react";
import "../css/Cell.css";

export default function Cell({ i, j, cellStyle, children, tile, onTileDrop }) {

    /**
     * Separate component for each cell on the board.
     * A tile can be passed as a child to cover it up.
     * Each cell has a key that is used to identify it in the letterupdates object.
     * Each cell can hold only one tile.
     */

    const key = `${i}-${j}`; // key for the tilePositions object

    const handleDrop = (e) => {
        /**
         * Native HTML drag and drop API to transfer data.
         */
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