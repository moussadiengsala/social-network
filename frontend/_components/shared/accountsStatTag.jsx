import React from "react";

function AccountStatTag({ stat, statValue, icon, additionalAttributes }) {
  const themes = {
    follower: "bg-blue-100 text-blue-600",
    like: "bg-pink-100 text-pink-600",
    post: "bg-indigo-100 text-indigo-600",
  };
  return (
    <div
      {...additionalAttributes}
      className={`flex cursor-pointer ${
        themes[stat] ? themes[stat] : "bg-orange-100 text-orange-500"
      } items-center px-2 py-1 rounded-md`}
    >
      {icon}
      <p className="text-sm ml-2 font-bold">
        {statValue} {stat}(s)
      </p>
    </div>
  );
}

export default AccountStatTag;
