import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter, Route, Routes, Navigate } from "react-router-dom";
import Welcome from './pages/Welcome.jsx';
import Room from './pages/Room.jsx';
import Game from './pages/Game.jsx';
import './css/index.css';

ReactDOM.createRoot(document.getElementById('root')).render(
  <BrowserRouter>
    <Routes>
      <Route path="/" element={<Navigate to="/home" />} />
      <Route index={true} path="/home" element={<Welcome />} />
      <Route path="/play/:id" element={<Game />} />
      <Route path="/room/:id" element={<Room />} />
    </Routes>
  </BrowserRouter>
)