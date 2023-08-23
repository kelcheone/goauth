import "./App.css";
// router
import { Route, Routes } from "react-router-dom";
import HomeView from "./HomeView";
import LoginView from "./LoginView";

function App() {
  return (
    <>
      <Routes>
        <Route path="/" element={<HomeView />} />
        <Route path="/login" element={<LoginView />} />
      </Routes>
    </>
  );
}

export default App;
