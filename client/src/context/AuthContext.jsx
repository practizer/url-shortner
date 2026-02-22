import { createContext, useEffect, useState } from "react";
import api from '../api/axios';

export const AuthContext = createContext();

export function AuthProvider({ children }) {
    const [isAuthenticated, setIsauthenticated] = useState(null);

    useEffect(() => {
        const checkAuth = async () => {
            try {
                await api.get("/auth/me");
                setIsauthenticated(true);
            } catch (err) {
                setIsauthenticated(false);
            }
        }
        checkAuth();
    },[]);

    return (
        <AuthContext.Provider value={{ isAuthenticated, setIsauthenticated }}>
            {children}
        </AuthContext.Provider>
    )
}