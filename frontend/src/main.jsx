import React from 'react'
import ReactDOM from 'react-dom/client'
import { BrowserRouter, Route, Routes, Navigate } from "react-router-dom";
import Welcome from './Welcome.jsx'
import App from './App.jsx'
import WaitingRoom from './WaitingRoom.jsx'
import './index.css'

ReactDOM.createRoot(document.getElementById('root')).render(
  <BrowserRouter>
    <Routes>
      <Route path="/" element={<Navigate to="/home" />} />
      <Route index={true} path="/home" element={<Welcome />} />
      <Route index={true} path="/room" element={<WaitingRoom />} />
      <Route index={false} path="/play" element={<App />} />
    </Routes>
  </BrowserRouter>
)