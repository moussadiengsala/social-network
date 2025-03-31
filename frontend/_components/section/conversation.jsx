import { useChatStore } from "@/_store/zustand";
import { InformationCircleIcon } from "@heroicons/react/24/outline";
import ChatCard from "../cards/chatCard";
function Conversation() {
  const a = useChatStore((state) => state.chat);
  return (
    <div className="w-[330px] hidden lg:block border-s px-3 py-4">
      {/* Empty state */}
      {Object.keys(a).length === 0 ? (
        <div className="flex items-center gap-2 w-fit">
          <div className="bg-blue-200 p-2 rounded-md text-blue-600 w-fit">
            <InformationCircleIcon className="w-5 h-5" />
          </div>
          <p className="text-sm text-blue-600">
            Click on an online user to start a conversation, or just keep
            swimming.
          </p>
        </div>
      ) : (
        <ChatCard key={a.id} />
      )}
    </div>
  );
}

export default Conversation;
