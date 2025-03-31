import React from "react";

function Loading() {
  return (
    <div className="flex space-x-2 mt-6 justify-center items-center ">
      <span className="sr-only">Loading...</span>
      <div className="h-2 w-2 bg-emerald-500 rounded-full animate-bounce [animation-delay:-0.3s]"></div>
      <div className="h-2 w-2 bg-emerald-500 rounded-full animate-bounce [animation-delay:-0.15s]"></div>
      <div className="h-2 w-2 bg-emerald-500 rounded-full animate-bounce"></div>
    </div>
  );
}

export default Loading;
