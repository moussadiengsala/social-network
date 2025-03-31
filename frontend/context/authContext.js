import { JWT as JWTService } from "@/lib/jwt";
import { getCookie } from "cookies-next";
import Cookies from "js-cookie";
import { createContext, useEffect, useState } from "react";
const AuthContext = createContext();

function AuthContextProvider({ children }) {
    // Initialize state
    const [data, setData] = useState();
    const [JWT, setJWT] = useState();

    // Fetch data
    useEffect(() => {
        setJWT(Cookies.get("social-network-jwt"));
        const tmp = JWTService.decoder(Cookies.get("social-network-jwt"));
        const { payload } = tmp;
        setData(payload);
    }, []);
    return (
        <AuthContext.Provider value={{ data, JWT }}>
            {children}
        </AuthContext.Provider>
    );
}

export { AuthContext, AuthContextProvider };
