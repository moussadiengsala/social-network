import { useWebsocketResponseStore, useWebsocketStore } from "@/_store/zustand";
import { useAuth } from "@/hooks/useAuth";
import { useEffect } from "react";

function SocketService({ children }) {
  const websocketService = useWebsocketStore((state) => state.websocketService);
  const websocketInitializer = useWebsocketStore((state) => state.setWebsocket);
  const setWebsocketResponse = useWebsocketResponseStore(
    (state) => state.setWebsocketResponse
  );
  const { JWT } = useAuth();
  useEffect(() => {
    if (!websocketService) {
      const ws = new WebSocket("ws://localhost:8000/api/ws");
      websocketInitializer(ws);
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
      });
    }
  }, [websocketService, JWT, websocketInitializer]);

  return <div>{children}</div>;
}

export default SocketService;
