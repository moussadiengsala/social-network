"use client";
import CustomInput from "@/_components/inputs/input";
import Alert from "@/_components/shared/alert";
import { useGroupFeedStore } from "@/_store/zustand";
import { useWebSocket } from "@/context/realTimeContext";
import useApiRequest from "@/hooks/useApiRequest";
import useFetch from "@/hooks/useFetch";
import usePaginatedFetch from "@/hooks/usePaginatedFetch";
import { formatDate } from "@/lib/TimeFormatter";
import { MessageSchema } from "@/lib/validations/messageValidation";
import { ArrowLeftIcon, PaperAirplaneIcon } from "@heroicons/react/24/outline";
import Image from "next/image";
import Link from "next/link";
import { useEffect, useState } from "react";
import { useQuery } from "react-query";

function GroupChat({ params }) {
  const fetcher = useFetch();
  const { websocketResponse } = useWebSocket();
  const [messageID, setMessageID] = useState();
  const group = useGroupFeedStore((state) => state.group);
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
  const groupOPTIONS = {
    method: "GET",
    endpoint: `groups/${params.id}`,
  };

  const { isLoading: groupPreviewLoading, data: userProfile } = useQuery(
    ["group", groupOPTIONS],
    () => fetcher(groupOPTIONS)
  );

  const initialEndpoint = `groups/messages/${params.id}`;

  const {
    isLoading,
    uniqueArray: previousMessages,
    loadMore,

    hasNextPage,
  } = usePaginatedFetch(fetcher, initialEndpoint, 2, "groups-messages");
  const currentIds = new Set(newMessages.map((message) => message.id));

  useEffect(() => {
    if (websocketResponse?.action === "notification-group-message") {
      if (websocketResponse?.payload?.data?.Items) {
        const id = websocketResponse?.payload?.data?.Items.id;
        setSingleMessageOpt({
          method: "GET",
          endpoint: `groups/message/${id}`,
        });
        if (websocketResponse?.payload?.data?.Items.group_id == params.id) {
          console.log("dddddfhfgv");
          setMessageID(websocketResponse?.payload?.data?.Items.id);
        }

        console.log(
          websocketResponse?.payload?.data?.Items.group_id,
          params.id
        );
      }
    }
  }, [websocketResponse]);
  const OPTIONS = {
    method: "POST",
    endpoint: "group-message",
    format: "--formData",
    dataToRefresh: "groups-messages",
  };

  const { responseError, register, handleSubmit, errors, onSubmit } =
    useApiRequest(OPTIONS, MessageSchema);
  return (
    <div className={`px-6`}>
      <Link
        href="/app/groups"
        className="flex items-center gap-2 hover:text-blue-600"
      >
        <ArrowLeftIcon className="w-5 h-5" />
        Back to group
      </Link>
      {groupPreviewLoading ? (
        <div className="bg-gray-200 animate-pulse h-[120px] w-full rounded-lg"></div>
      ) : (
        <div
          style={{
            backgroundImage: `url(http://localhost:8000/api/static/uploads/${userProfile?.data?.cover})`,
          }}
          className={`border-gray-300 mt-2 rounded-lg h-[120px] flex flex-col justify-center overflow-hidden bg-center text-white relative bg-cover `}
        >
          <div className="absolute inset-0 bg-gray-900/75 sm:bg-transparent sm:from-gray-900/95 sm:to-gray-700/25 sm:bg-gradient-to-r "></div>
          <div className="z-50 relative px-3 py-4">
            <h1 className="text-xl  font-bold">{userProfile?.data?.title}</h1>
            <h1 className="text-sm">{userProfile?.data?.description}</h1>
            <div className="flex justify-between items-center">
              <h1 className="text-sm">
                Created since{" "}
                <span className="font-bold">
                  {formatDate(userProfile?.data?.created_at)}
                </span>{" "}
              </h1>
              <small className="">
                {userProfile?.data?.member_count} member(s)
              </small>
            </div>
          </div>
        </div>
      )}
      {isLoading ? (
        <div className="px-2 py-4 space-y-6">
          <p className="px-2 py-1 w-[200px] h-[30px]  relative text-sm rounded-md bg-gray-200 animate-pulse"></p>
          <p className="px-2 py-1 w-[200px] h-[30px]  relative text-sm rounded-md bg-gray-200 animate-pulse"></p>
          <p className="px-2 py-1 w-[200px] h-[30px]  relative text-sm rounded-md bg-gray-200 animate-pulse"></p>
        </div>
      ) : (
        <div className="py-4  overflow-scroll  space-y-6 max-h-[300px] mb-4">
          {userProfile?.data?.is_part_of_group ? (
            <>
              {hasNextPage ? (
                <div
                  onClick={loadMore}
                  className="py-1 text-blue-600 hover:underline cursor-pointer mx-auto w-fit"
                >
                  <small>See old messages</small>
                </div>
              ) : (
                <div className="mx-auto w-fit text-xs">No more message</div>
              )}
              {previousMessages && previousMessages.length ? (
                previousMessages
                  .filter((message) => !currentIds.has(message.id))
                  .reverse()
                  .map((prev) => (
                    <div className="flex items-center gap-4" key={prev.id}>
                      <Image
                        width={100}
                        height={100}
                        src={
                          prev.sender_avatar !== ""
                            ? `http://localhost:8000/api/static/uploads/${prev.sender_avatar}`
                            : "/unknown.jpg"
                        }
                        alt="user profile picture"
                        className="w-9 h-9 rounded-md mt-3"
                      />
                      <div className=" space-y-2 text-sm">
                        <p> @{prev.sender_username}</p>
                        <p
                          className={` px-2 py-1 bg-gray-100 text-gray-600 w-fit relative text-sm  rounded-md `}
                        >
                          {prev.content}
                        </p>
                      </div>
                    </div>
                  ))
              ) : (
                <div className="text-xs h-full text-gray-600  text-center italic">
                  <span className="font-bold">{group.title}</span> doesn&apos;t
                  have any conversation yet
                </div>
              )}
              {newMessages.map((message) => (
                <div className="flex items-center gap-4" key={message.id}>
                  <Image
                    width={100}
                    height={100}
                    src={
                      message.sender_avatar !== ""
                        ? `http://localhost:8000/api/static/uploads/${message.sender_avatar}`
                        : "/unknown.jpg"
                    }
                    alt="user profile picture"
                    className="w-9 h-9 rounded-md mt-3"
                  />
                  <div className=" space-y-2 text-sm">
                    <p> @{message.sender_username}</p>
                    <p
                      className={` px-2 py-1 bg-gray-100 text-gray-600 w-fit relative text-sm  rounded-md `}
                    >
                      {message.content}
                    </p>
                  </div>
                </div>
              ))}
              {responseError.length !== 0 && (
                <Alert message={responseError} status="error" />
              )}
            </>
          ) : (
            <h1 className="text-center text-gray-600">
              You not part of this group
            </h1>
          )}
        </div>
      )}
      {userProfile?.data?.is_part_of_group ? (
        <form onSubmit={handleSubmit(onSubmit)} action="">
          <div className="flex  items">
            <CustomInput
              type="hidden"
              value={params.id}
              register={register}
              name="group_id"
              className=""
            />
            <CustomInput
              type="text"
              register={register}
              label="Message"
              placeholder="write your message..."
              error={errors.content}
              name="content"
              className=" flex-1 mr-4"
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
      ) : null}
    </div>
  );
}

export default GroupChat;
