import { useChatStore } from "@/_store/zustand";
import { useWebSocket } from "@/context/realTimeContext";
import useApiRequest from "@/hooks/useApiRequest";
import useFetch from "@/hooks/useFetch";
import usePaginatedFetch from "@/hooks/usePaginatedFetch";
import { MessageSchema } from "@/lib/validations/messageValidation";
import { PaperAirplaneIcon } from "@heroicons/react/24/outline";
import Image from "next/image";
import { useEffect, useState } from "react";
import { useQuery } from "react-query";
import CustomInput from "../inputs/input";
import Alert from "../shared/alert";

function ChatCard({ key }) {
  const { websocketResponse } = useWebSocket();
  const fetcher = useFetch();
  const chat = useChatStore((state) => state.chat);
  const [messageID, setMessageID] = useState();
  const initialEndpoint = `messages/${chat.id}`;
  const {
    isLoading,
    uniqueArray: previousMessages,
    loadMore,
    hasNextPage,
    refetch,
  } = usePaginatedFetch(fetcher, initialEndpoint, 10, "private-message");
  const [singlemessageOpt, setSingleMessageOpt] = useState({
    method: "GET",
    endpoint: `/message/${messageID}`,
  });
  const [newMessages, setNewMessages] = useState([]);
  useQuery(["message", singlemessageOpt], () => fetcher(singlemessageOpt), {
    onSuccess: (data) => {
      console.log(data?.data);
      setNewMessages((prevData) => [...prevData, data?.data]);
    },

    enabled: !!messageID,
    refetchOnWindowFocus: false,
  });
  useEffect(() => {
    console.log(websocketResponse?.action);
    if (websocketResponse?.action === "notification-private-message") {
      if (websocketResponse?.payload?.data?.Items) {
        const id = websocketResponse?.payload?.data?.Items.id;
        setSingleMessageOpt({
          method: "GET",
          endpoint: `message/${id}`,
        });
        if (websocketResponse?.payload?.data?.Items.sender_id === chat.id) {
          console.log("Go fetch this");
          setMessageID(websocketResponse?.payload?.data?.Items.id);
        }

        console.log(websocketResponse?.payload?.data?.Items.sender_id, chat.id);
      }
    }
  }, [websocketResponse, key]);

  const OPTIONS = {
    method: "POST",
    endpoint: "private-message",
    format: "--formData",
    dataToRefresh: "private-message",
  };

  const { responseError, register, handleSubmit, errors, onSubmit } =
    useApiRequest(OPTIONS, MessageSchema);
const currentIds = new Set(newMessages.map(message => message.id));
  return (
    <section>
      {/* Chat card Header */}
      <div className="flex items-center gap-2 border-b-[1px] pb-3">
        <div className="relative">
          <Image
            width={100}
            height={100}
            className="h-12 w-12 flex-none border-2 shadow-xl rounded-full bg-gray-50"
            src={
              chat.avatar != ""
                ? `http://localhost:8000/api/static/uploads/${chat.avatar}`
                : "/unknown.jpg"
            }
            alt="user avatar"
          />

          <div
            className={`flex-none absolute bottom-0 right-0 rounded-full ${
              chat.online ? "bg-emerald-500/20" : "bg-pink-500/20"
            }  p-1`}
          >
            <div
              className={`h-1.5 w-1.5 rounded-full ${
                chat.online ? "bg-emerald-500" : "bg-pink-500"
              } `}
            />
          </div>
        </div>
        <div className="min-w-0">
          <p className="text-sm font-semibold leading-6 text-gray-900">
            {`${chat.first_name} ${chat.last_name}`}
          </p>
          <p className="truncate text-gray-500 rounded-md text-xs leading-5 ">
            {chat.email}
          </p>
        </div>
      </div>
      {/* Chat card Content */}
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
            previousMessages.filter(message => !currentIds.has(message.id)).reverse().map((prev) => (
              <p
                key={prev.id}
                className={`${
                  prev.receiver_id
                } px-2 py-1 w-fit relative text-sm  rounded-md ${chat.id} ${
                  prev.receiver_id == chat.id
                    ? "bg-blue-600 text-white text-right"
                    : "bg-blue-100 text-blue-600"
                }`}
              >
                {prev.content}
              </p>
            ))
          ) : (
            <div className="text-xs h-full text-gray-600  text-center italic">
              You don't have conversation with {chat.first_name}
            </div>
          )}
          {newMessages?.map((message, index) => (
            <p
              key={index}
              className={` px-2  py-1 w-fit relative text-sm  rounded-md ${
                message.receiver_id == chat.id
                  ? "bg-blue-600 text-white text-right"
                  : "bg-blue-100 text-blue-600"
              } `}
            >
              {message?.content}
            </p>
          ))}

          {responseError.length !== 0 && (
            <Alert message={responseError} status="error" />
          )}
        </div>
      )}

      <form onSubmit={handleSubmit(onSubmit)} action="">
        <CustomInput
          type="hidden"
          value={chat.id}
          register={register}
          name="receiver_id"
          className="col-span-6 sm:col-span-3 "
        />
        <div className="flex gap-2 items">
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
            className="bg-blue-600 disabled:opacity-25 w-[50px] flex items-center justify-center rounded-md text-white px-3 py-2"
            type="submit"
          >
            <PaperAirplaneIcon className="w-5 h-5" />
          </button>
        </div>
      </form>
    </section>
  );
}

export default ChatCard;
