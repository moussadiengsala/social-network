import React from "react";

function AsideProfile() {
  return (
    <div className="flex items-center gap-x-4 px-6 py-3 text-sm font-semibold leading-6 text-gray-900 hover:bg-gray-50">
      <div className="h-8 w-8 rounded-full animate-pulse bg-gray-400" />
      <span className="sr-only">Your profile</span>
      <div className="flex gap-2 flex-col">
        <span
          className=" animate-pulse bg-gray-400 w-[120px] h-2"
          aria-hidden="true"
        ></span>
        <span
          className="font-light bg-gray-400 animate-pulse w-[100px] h-2 "
          aria-hidden="true"
        ></span>
      </div>
    </div>
  );
}

export default AsideProfile;
