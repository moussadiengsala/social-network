import useFetch from "@/hooks/useFetch";
import { useState } from "react";

import usePaginatedFetch from "@/hooks/usePaginatedFetch";
import Image from "next/image";
import { useQuery } from "react-query";
function FollowRequestTab() {
  const fetcher = useFetch();
  const initialEndpoint = `notifications?notification=follow_request`;
  const [requestID, setRequestId] = useState();
  const [status, setStatus] = useState();
  const {
    isLoading,
    uniqueArray: notifications,
    loadMore,
    hasNextPage,
  } = usePaginatedFetch(
    fetcher,
    initialEndpoint,
    5,
    "follow-request-notifications"
  );

  const options = {
    method: "GET",
    endpoint: `accept-follow?user_id=${requestID}&status=${status}`,
  };
  const { data, isLoading: followRequestLoading } = useQuery(
    ["accept-follow", options],
    () => fetcher(options),
    {
      enabled: !!requestID,
    }
  );

  return (
    <section>
      <h1 className="text-xl font-bold capitalize mb-2">Follow Request</h1>
      {isLoading ? (
        <>Loading</>
      ) : (
        <>
          {notifications?.map((notification,i) => (
            <div key={i} className="flex items-center divide-y-1 bg-ellow-500 justify-between gap-1">
              <div className="flex items-center gap-2">
                <Image
                  width={100}
                  height={100}
                  src={
                    notification.sender_avatar != ""
                      ? `http://localhost:8000/api/static/uploads/${notification.sender_avatar}`
                      : "/unknown.jpg"
                  }
                  alt="user profile picture"
                  className="w-9 h-9 rounded-full"
                />
                <div>
                  <span className="italic font-bold">{`${notification.sender_first_name} ${notification.sender_last_name} `}</span>
                  have request sent you a friend request
                </div>
              </div>
              <div className="space-x-2">
                <button
                  disabled={followRequestLoading}
                  onClick={() => {
                    setRequestId(notification.sender_id);
                    setStatus("accept");
                  }}
                  className="bg-blue-600 disabled:opacity-35 text-sm rounded-md px-3 py-2 text-white"
                >
                  Accept
                </button>
                <button
                  disabled={followRequestLoading}
                  onClick={() => {
                    setRequestId(notification.sender_id);
                    setStatus("reject");
                  }}
                  className="bg-pink-600 disabled:opacity-35 text-sm rounded-md px-3 py-2 text-white"
                >
                  Decline
                </button>
              </div>
            </div>
          ))}
          {hasNextPage ? (
            <div
              onClick={loadMore}
              className="text-center text-sm text-blue-600 w-fit hover:underline mx-auto cursor-pointer"
            >
              See more notifications
            </div>
          ) : (
            <div className="text-gray-700 text-center  mx-auto w-fit">
              <small>No more notifications</small>
            </div>
          )}
        </>
      )}
    </section>
  );
}

export default FollowRequestTab;
