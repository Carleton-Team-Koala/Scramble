import React from "react";
import Tile from "./Tile.jsx"
import "./Hand.css";

export default function Hand() {
    let tiles = [];
    for (let i = 0; i < 7; i++){
        tiles.push(<Tile key={i} />);
    }
    return (
        <div className="hand">
            {tiles}
        </div>
    );   
}