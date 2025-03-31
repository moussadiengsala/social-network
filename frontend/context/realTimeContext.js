import React, { useState, useEffect, useContext } from "react";
import { useAuth } from "@/hooks/useAuth";
import Cookies from "js-cookie";

const WebSocketContext = React.createContext();

function useWebSocket() {
  return useContext(WebSocketContext);
}

function SocketServiceProvider({ children }) {
  const [websocketService, setWebsocketService] = useState(null);
  const [websocketResponse, setWebsocketResponse] = useState(null);
  const [notifications, setNotifications] = useState();
  const [friends, setFriends] = useState([]);
  const [belongedGroups, setBelongedGroups] = useState([]);
  const JWT = Cookies.get("social-network-jwt");

  useEffect(() => {
    if (!websocketService) {
      const ws = new WebSocket("ws://localhost:8000/api/ws");
      setWebsocketService(ws);
      ws.addEventListener("open", () => {
        console.log("connection established");

        ws.send(
          JSON.stringify({
            action: "join",
            payload: {
              token: JWT,
            },
          })
        );
      });
      ws.addEventListener("message", (e) => {
        console.log(JSON.parse(e.data));
        setWebsocketResponse(JSON.parse(e.data));
        const data = JSON.parse(e.data);
        if (
          data?.payload.data.unread_notifications ||
          data?.payload.data.unread_notifications === 0
        )
          setNotifications(data?.payload?.data?.unread_notifications);
        if (data?.payload?.data?.users) setFriends(data?.payload?.data?.users);
        if (data?.payload?.data?.groups)
          setBelongedGroups(data?.payload?.data?.groups);
      });
    }
  }, [websocketService, JWT]);

  const contextValue = {
    websocketService,
    notifications,
    friends,
    websocketResponse,
    belongedGroups,
  };

  return (
    <WebSocketContext.Provider value={contextValue}>
      {children}
    </WebSocketContext.Provider>
  );
}

export { SocketServiceProvider, useWebSocket };
