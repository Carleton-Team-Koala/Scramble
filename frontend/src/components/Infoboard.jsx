import React, { useState } from 'react';
import './Infoboard.css';

const Infoboard = (props) => {
    return (
        <div id="infoboard">
            <div className="scoreboard">
                <div className="scoreboard-cell">
                    Player1: 0
                </div>
                <div className="scoreboard-cell">
                    Player2: 0
                </div>
            </div>
            <div className="turnboard">
                <div className="turnboard-cell">
                    Player1's turn
                </div>
            </div>
            <div className="tilebag">
                <p>Tilebag: 86</p>
                <p>
                    ?x2 Ax9 Bx2 Cx2 Dx4 Ex8 Fx2 Gx3 Hx2 Ix9 J K
                     Lx4 Mx2 Nx3 Ox7 Px2 Q Rx6 Sx4 Tx4 Ux3 Vx2
                      W X Yx2 Z
                </p>
            </div>
        </div>
    );
}

export default Infoboard;