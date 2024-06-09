import AppRouter from "./router/router";

import AOS from "aos";
import "aos/dist/aos.css";
AOS.init();

function App() {
  return <AppRouter />;
}

export default App;
