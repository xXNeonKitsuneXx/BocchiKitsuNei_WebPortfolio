import { BrowserRouter, Routes, Route } from "react-router-dom";
import { Home } from "../pages/home";
import { Project } from "../pages/project";
import { Error } from "../pages/error";

const App = () => {
  return (
    <BrowserRouter>
      <div className="bg-white">
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/project" element={<Project />} />
          <Route path="/*" element={<Error />} />
        </Routes>
      </div>
    </BrowserRouter>
  );
};

export default App;
