import { useReducer, useEffect } from "react";

const initialState = {
  socket: null,
  isConnected: false,
  messages: [],
};

function reducer(state, action) {
  switch (action.type) {
    case "CONNECT":
      return {
        ...state,
        socket: new WebSocket(action.payload.url),
      };
    case "MESSAGE":
      return {
        ...state,
        messages: [...state.messages, action.payload.message],
      };
    case "CLOSE":
      return {
        ...state,
        socket: null,
        isConnected: false,
        messages: [],
      };
    default:
      return state;
  }
}

function useWebSocket(url) {
  const [state, dispatch] = useReducer(reducer, initialState);

  useEffect(() => {
    if (!state.socket) {
      dispatch({ type: "CONNECT", payload: { url } });
    }

    return () => {
      if (state.socket) {
        state.socket.close();
        dispatch({ type: "CLOSE" });
      }
    };
  }, []);

  useEffect(() => {
    if (state.socket) {
      state.socket.onopen = () => {
        dispatch({ type: "CONNECTED" });
      };

      state.socket.onmessage = (event) => {
        dispatch({ type: "MESSAGE", payload: { message: event.data } });
      };
    }
  }, [state.socket]);

  const sendMessage = (data) => {
    if (state.socket && state.socket.readyState === WebSocket.OPEN) {
      state.socket.send(JSON.stringify(data));
    }
  };

  return { messages: state.messages, sendMessage };
}

export default useWebSocket;
