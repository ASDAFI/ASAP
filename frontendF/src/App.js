import { BrowserRouter, Routes, Route } from "react-router-dom";
import Cookie from "universal-cookie";

import Login from "./Login";
import Main from "./Main";
import Intro from "./Intro";

import "./App.css";

function App() {
  const cookie = new Cookie();

  return (
    <div className="App">
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Intro />} />
          <Route path="/login" element={<Login />} />
          <Route path="/main" element={<Main />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
