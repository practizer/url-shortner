import { createContext, useEffect, useState, useRef } from "react";
import api from '../api/axios';

export const AuthContext = createContext();

export function AuthProvider({ children }) {
    const [isAuthenticated, setIsauthenticated] = useState(null);
    const hasChecked = useRef(false);

    useEffect(() => {
        if (hasChecked.current) return;
        hasChecked.current = true;

        const checkAuth = async () => {
            try {
                await api.get("/auth/me");
                setIsauthenticated(true);
            } catch (err) {
                setIsauthenticated(false);
            }
        }
        checkAuth();
    }, []);

    return (
        <AuthContext.Provider value={{ isAuthenticated, setIsauthenticated }}>
            {children}
        </AuthContext.Provider>
    )
}