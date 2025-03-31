import React from "react";

export default function Alert({ status, message }) {
  const statusColor = {
    error: "bg-pink-50 text-pink-500",
    success: "bg-emerald-50 text-emerald-500",
    warning: "bg-yellow-50 text-yellow-500",
  };

  return (
    <div className={`rounded-md my-4 w-fit ${statusColor[status]} p-4`}>
      <div className="flex">
        <div className="ml-3">
          {typeof message === "string"  ? (
            <h3 className="text-sm font-medium">{message}</h3>
          ) : (
            <React.Fragment>
            <h3 className="text-sm font-medium">{message.title}</h3>
              <ul className="list-disc">
                {/* Map over responseError and render each item as a list item */}
                {message.map((error, index) => (
                  <li className="text-sm" key={index}>{error}</li>
                ))}
              </ul>
            </React.Fragment>
          )}
        </div>
      </div>
    </div>
  );
}
