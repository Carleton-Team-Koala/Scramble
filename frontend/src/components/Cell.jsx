import React from "react";
import "./Cell.css";

export default function Cell({ i, j, cellStyle, cellValue, children }) {
    return (
        <div
            key={`${i}-${j}`}
            className={`cell ${cellStyle}`}
        >
            <div className="cell-content">
                {children}
            </div>
        </div>
    );
};