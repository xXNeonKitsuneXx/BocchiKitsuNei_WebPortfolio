import { BrowserRouter, Routes, Route } from "react-router-dom";
import { HomePage } from "../pages/homepage";
import { ProjectPage } from "../pages/projectpage";
import { ErrorPage } from "../pages/errorpage";

const App = () => {
  return (
    <BrowserRouter>
      <div className="bg-white">
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/project" element={<ProjectPage />} />
          <Route path="/*" element={<ErrorPage />} />
        </Routes>
      </div>
    </BrowserRouter>
  );
};

export default App;
