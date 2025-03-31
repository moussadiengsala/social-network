export default function CommentSkeleton() {
  return (
    <div className="relative border-2 border-gray-100 text-sm grid grid-cols-1 gap-4 p-4 mb-8 rounded-lg bg-white shadow-sm animate-pulse">
      <div className="relative flex items-center mb-2 gap-4">
        <div className="relative rounded-full bg-gray-200 border h-10 w-12 animate-pulse"></div>
        <div className="flex flex-col w-full">
          <div className="flex flex-row justify-between">
            <div className="relative whitespace-nowrap truncate overflow-hidden bg-gray-200 w-3/4 animate-pulse"></div>
          </div>
          <div className="text-gray-400 text-sm bg-gray-200 w-1/2 animate-pulse"></div>
        </div>
      </div>
      <div className="-mt-4 text-gray-500 text-sm bg-gray-200 h-3 mb-2 w-3/4 animate-pulse"></div>
      <div className="-mt-4 text-gray-500 text-sm bg-gray-200 h-3 w-3/4 animate-pulse"></div>
    </div>
  );
}
