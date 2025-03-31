import CustomInput from "@/_components/inputs/input";
import Alert from "@/_components/shared/alert";
import { useGroupFeedStore } from "@/_store/zustand";
import useApiRequest from "@/hooks/useApiRequest";
import { useAuth } from "@/hooks/useAuth";
import useFetch from "@/hooks/useFetch";
import usePaginatedFetch from "@/hooks/usePaginatedFetch";
import { MessageSchema } from "@/lib/validations/messageValidation";
import { PaperAirplaneIcon } from "@heroicons/react/24/outline";

function GroupMessageFeed() {
  const fetcher = useFetch();
  const group = useGroupFeedStore((state) => state.group);
  console.log("heeeere is the group id", group.id); 
  const initialEndpoint = `groups/messages/${group.id}`;

  const {
    isLoading,
    uniqueArray: previousMessages,
    loadMore,

    hasNextPage,
  } = usePaginatedFetch(fetcher, initialEndpoint, 2, "groups-messages");

  const { data } = useAuth();
  const OPTIONS = {
    method: "POST",
    endpoint: "group-message",
    format: "--formData",
    dataToRefresh: "groups-messages",
  };

  const { responseError, register, handleSubmit, errors, onSubmit } =
    useApiRequest(OPTIONS, MessageSchema);

  return (
    <div>
      {isLoading ? (
        <div className="px-2 py-4 space-y-6">
          <p className="px-2 py-1 w-[200px] h-[30px]  relative text-sm rounded-md bg-gray-200 animate-pulse"></p>
          <p className="px-2 py-1 w-[200px] h-[30px]  relative text-sm rounded-md bg-gray-200 animate-pulse"></p>
          <p className="px-2 py-1 w-[200px] h-[30px]  relative text-sm rounded-md bg-gray-200 animate-pulse"></p>
        </div>
      ) : (
        <div className="py-4 overflow-scroll px-2 space-y-6 max-h-[300px] mb-4">
          {hasNextPage ? (
            <div
              onClick={loadMore}
              className=" py-1 text-blue-600 hover:underline cursor-pointer mx-auto w-fit"
            >
              <small>See old messages</small>
            </div>
          ) : (
            <div className="mx-auto w-fit text-xs">No more message</div>
          )}
          {previousMessages && previousMessages.length ? (
            previousMessages.map((prev) => (
              <div>
                <p
                  key={prev.id}
                  className={`${
                    prev.sender_id === data?.user.id
                      ? "bg-blue-600 text-white"
                      : "bg-blue-200 text-blue-600"
                  } px-2 py-1 w-fit relative text-sm  rounded-md `}
                >
                  {prev.content}
                </p>
              </div>
            ))
          ) : (
            <div className="text-xs h-full text-gray-600  text-center italic">
              <span className="font-bold">{group.title}</span> doesn't have any
              conversation yet
            </div>
          )}
          {responseError.length !== 0 && (
            <Alert message={responseError} status="error" />
          )}
        </div>
      )}
      <form onSubmit={handleSubmit(onSubmit)} action="">
        <div className="flex gap-2 px-3 items">
          <CustomInput
            type="hidden"
            value={group.id}
            register={register}
            name="group_id"
            className="col-span-6  sm:col-span-3 "
          />
          <CustomInput
            type="text"
            register={register}
            label="Message"
            placeholder="write your message..."
            error={errors.content}
            name="content"
            className="col-span-6 sm:col-span-3 "
          />

          <button
            disabled={isLoading}
            className="bg-blue-600 w-[50px] flex items-center justify-center rounded-md text-white px-3 py-2"
            type="submit"
          >
            <PaperAirplaneIcon className="w-5 h-5" />
          </button>
        </div>
      </form>
    </div>
  );
}

export default GroupMessageFeed;
