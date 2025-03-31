const PostSkeleton = () => {
  return (
    <div className="bg-white border-2 border-gray-100 text-gray-700 rounded-xl shadow-sm text-sm font-medium animate-pulse">
      <div className="flex gap-3 sm:p-4 p-2.5 text-sm font-medium">
        <div className="w-9 h-9 rounded-full bg-gray-200"></div>
        <div className="flex-1">
          <div className="w-1/2 h-4 bg-gray-200 rounded mb-2 animate-pulse"></div>
          <div className="w-1/4 h-4 bg-gray-200 rounded animate-pulse"></div>
        </div>
        <div className="-mr-1">
          <div className="w-6 h-6 bg-gray-200 rounded-full animate-pulse"></div>
        </div>
      </div>
      <div className="sm:px-4 space-y-2 p-2.5 pt-0">
        <div className="w-full h-[150px] sm:h-[250px] lg:h-[330px] bg-gray-200 mb-4 rounded-lg"></div>
        <div className="w-3/4 h-4 bg-gray-200 rounded animate-pulse"></div>
        <div className="w-5/6 h-4 bg-gray-200 rounded animate-pulse"></div>
        <div className="w-2/4 h-4 bg-gray-200 rounded animate-pulse"></div>
      </div>
      <div className="sm:p-4 p-2.5 flex items-center gap-4 text-xs font-semibold">
        <div className="flex items-center gap-2.5">
          <div className="w-5 h-6 bg-gray-200 rounded-full animate-pulse"></div>
          <div className="w-8 h-4 bg-gray-200 rounded animate-pulse"></div>
        </div>
        <div className="flex items-center gap-3">
          <div className="w-5 h-6 bg-gray-200 rounded-full animate-pulse"></div>
          <div className="w-4 h-4 bg-gray-200 rounded animate-pulse"></div>
        </div>
      </div>
      <div className="sm:p-4 p-2.5 border-t border-gray-100 font-normal space-y-3 relative">
        <div className="flex items-start gap-3 relative">
          <div className="w-6 h-6 rounded-full bg-gray-200"></div>
          <div className="flex-1 space-y-2">
            <div className="w-1/3 h-4 bg-gray-200 rounded animate-pulse"></div>
            <div className="w-4/5 h-4 bg-gray-200 rounded animate-pulse"></div>
          </div>
        </div>
      </div>
      <div className="sm:px-4 sm:py-3 p-2.5 border-t border-gray-100 flex items-center gap-1">
        <div className="w-6 h-6 rounded-full bg-gray-200"></div>
        <div className="flex-1 relative overflow-hidden h-10">
          <div className="w-full h-full bg-gray-200 rounded-md animate-pulse"></div>
        </div>
      </div>
    </div>
  );
};

export default PostSkeleton;
