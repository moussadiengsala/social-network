import { AuthContext } from "@/context/authContext";
import { useContext } from "react";

export function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error("Context must be used within a Provider");
  }
  return context;
}
