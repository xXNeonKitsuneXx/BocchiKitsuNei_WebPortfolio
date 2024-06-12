import axios from "axios";

const Axios = axios.create({
  baseURL: import.meta.env.VITE_BACKEND_URL ?? "http://localhost:8888",
});

export { Axios };
