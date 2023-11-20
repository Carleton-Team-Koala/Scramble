import React from 'react';
import '../css/Infoboard.css';

/**
 * Sums up all values inside the tilebag
 * 
 * @param {*} tile_dict 
 * @returns 
 */
function sumTiles(tile_dict) {
    return Object.values(tile_dict).reduce((sum, value) => {
        return sum + value;
    }, 0);
};

const Infoboard = ({ tilebag, p1_name, p2_name, p1_score, p2_score, currentPlayer }) => {

    /**
     * Displays the game information (scores, tiles left) to the user.
     */

    return (
        <div id="infoboard">
            <div className="scoreboard">
                <div className="scoreboard-cell">
                    {p1_name}: {p1_score}
                </div>
                <div className="scoreboard-cell">
                    {p2_name}: {p2_score}
                </div>
            </div>
            <div className="turnboard">
                <div className="turnboard-cell">
                    {currentPlayer}'s turn
                </div>
            </div>
            <div className="tilebag">
                <p>Tilebag: {sumTiles(tilebag)}</p>
                <p>
                    Ax{tilebag['A']} Bx{tilebag['B']} Cx{tilebag['C']} Dx{tilebag['D']}
                    Ex{tilebag['E']} Fx{tilebag['F']} Gx{tilebag['G']} Hx{tilebag['H']} Ix{tilebag['I']} Jx{tilebag['J']}
                    Kx{tilebag['K']} Lx{tilebag['L']} Mx{tilebag['M']} Nx{tilebag['N']} Ox{tilebag['O']} Px{tilebag['P']}
                    Qx{tilebag['Q']} Rx{tilebag['R']} Sx{tilebag['S']} Tx{tilebag['T']} Ux{tilebag['U']} Vx{tilebag['V']}
                    Wx{tilebag['W']} Xx{tilebag['X']} Yx{tilebag['Y']} Zx{tilebag['Z']}
                </p>
            </div>
        </div>
    );
}

export default Infoboard;