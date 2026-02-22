import axios from "axios"

const api = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL,
    headers: {"Content-Type":"application/json"},
    withCredentials:true
})

api.interceptors.response.use(
  res => res,
  err => {
    // Only redirect on 401 for non-auth endpoints to avoid auth check loops
    if (err.response?.status === 401 && !err.config?.url?.includes("/auth/me")) {
      window.location.href = "/login";
    }
    return Promise.reject(err);
  }
);

export default api;