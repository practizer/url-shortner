import React from "react";
import { BrowserRouter, Routes, Route,Navigate } from "react-router-dom";
import Login from "../Pages/Login";
import Home from "../Pages/Home";
import ProtectedRoute from "../components/ProtectedRoute";
import PublicRoute from "../components/PublicRoute";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Navigate to="/home" replace />} />
        <Route path="/login" element={<PublicRoute><Login /></PublicRoute>} />
        <Route path="/home" element={<ProtectedRoute><Home /></ProtectedRoute>} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;