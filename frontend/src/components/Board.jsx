import React from "react";
import "./Board.css";

const Board = () => {
    const rows = [];

    for (let i = 0; i < 15; i++) { // make a board
        const cells = [];

        for (let j = 0; j < 15; j++) {
            cells.push(
                <input
                    key={`${i}-${j}`}
                    type="text"
                    maxLength={1}
                    style={{ width: "30px", height: "30px", textAlign: "center" }}
                />
            );
        }

        rows.push(<div key={i}>{cells}</div>);
    }

    return (
        <div>
            {rows}
        </div>
    );
}

export default Board;