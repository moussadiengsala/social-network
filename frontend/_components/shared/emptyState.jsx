import { InformationCircleIcon } from "@heroicons/react/24/outline";

function EmptyState({ message }) {
  return (
    <div className="flex items-center gap-2 w-fit px-3">
      <div className="bg-blue-200 p-2 rounded-md text-blue-600 w-fit">
        <InformationCircleIcon className="w-5 h-5" />
      </div>
      <p className="text-sm text-blue-600">{message}</p>
    </div>
  );
}

export default EmptyState;
