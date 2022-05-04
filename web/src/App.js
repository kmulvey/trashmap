import "./App.css";
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import Login from "./Pages/Login";
import Map from "./Pages/Map";


function App() {
  return (
    <Router>
      <nav>
        <Link to="/"> Home </Link>
      </nav>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/map" element={<Map />} />
        <Route path="*" element={<Login />} />
      </Routes>
    </Router>
  );
}

export default App;