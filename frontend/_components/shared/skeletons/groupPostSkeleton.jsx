import React from "react";

function GroupPostSkeleton() {
  return (
    <div className="animate-pulse bg-gray-100 rounded-md p-2">
      {/* User profile skeleton */}
      <div className="flex items-center gap-1 mb-2">
        <div className="w-6 h-6 rounded-full bg-gray-200"></div>
        <div className="h-4 w-20 bg-gray-200 rounded"></div>
      </div>
      {/* Post image skeleton */}
      <div className="flex gap-2">
        <div className="h-[80px] w-1/3  bg-gray-200 rounded-md"></div>
        <div className="w-2/3">
          <div className="space-y-2">
            <div className="h-4 bg-gray-200 rounded"></div>
            <div className="h-4 bg-gray-200 rounded"></div>
            <div className="h-4 bg-gray-200 rounded"></div>
          </div>
          {/* Reaction skeleton */}
          <div className="flex text-xs gap-2 mt-2">
            <div className="h-4 w-14 bg-gray-200 rounded"></div>
            <div className="h-4 w-14 bg-gray-200 rounded"></div>
            <div className="h-4 w-14 bg-gray-200 rounded"></div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default GroupPostSkeleton;
