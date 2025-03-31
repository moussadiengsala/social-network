import { useGroupFeedStore } from "@/_store/zustand";
import { queryClient } from "@/app/layout";
import useFetch from "@/hooks/useFetch";
import { formatEventDate } from "@/lib/TimeFormatter";
import { SignalIcon, SignalSlashIcon } from "@heroicons/react/24/outline";
import Image from "next/image";
import { useState } from "react";
import { useQuery } from "react-query";

function EventCard({ event }) {
  const fetcher = useFetch();
  const { month, day, time } = formatEventDate(event.datetime);
  const [status, setStatus] = useState();
  const group = useGroupFeedStore((state) => state.group);

  const options = {
    method: "POST",
    endpoint: `groups-events-response?group_id=${group.id}&response=${status}&event_id=${event.id}`,
  };
  const { data, isLoading: followRequestLoading } = useQuery(
    ["group-events-response", options],
    () => fetcher(options),
    {
      onSuccess: () => {
        queryClient.refetchQueries(["groups-events"], { active: true });
      },
      enabled: !!status,
    }
  );

  return (
    <div className="relative space-y-2 overflow-hidden text-xs text-gray-800 border-2 border-gray-100 p-2 rounded-md shadow-md">
      <div className="flex items-center gap-1">
        {event.author_avatar && (
          <Image
            width={100}
            height={100}
            className="w-6 h-6 rounded-full"
            src={`http://localhost:8000/api/static/uploads/${event.author_avatar}`}
            alt="user profile picture"
          />
        )}
        {`${event.author_first_name} ${event.author_last_name}`}
      </div>
      <div className="gap-2 items-center  flex">
        <div className="h-[80px] border-gray-100 shadow-md border-2 flex-1 flex flex-col items-center justify-center  rounded-md ">
          <p className="text-2xl font-bold ">{day}</p>
          <p className="text-bas">{month}</p>
          <p className="text-bas">{time}</p>
        </div>
        <div className="w-[70%]">
          <p className="mb-2 text-blue-600 font-bold">{event.title}</p>
          <p className="">{event.description}</p>
          {event.response_status === "going" && (
            <div className="text-blue-600 text-xs w-fit bg-blue-200 px-1 rounded-sm my-2">
              You are interested
            </div>
          )}
          {event.response_status === "not going" && (
            <div className="text-pink-600 text-xs w-fit bg-pink-200 px-1 rounded-sm my-2">
              You are not interested
            </div>
          )}
          {event.response_status === "null" && (
            <div className="flex  items-center gap-x-2 mt-3 justify-end ">
              <small
                onClick={() => setStatus("going")}
                className="text-blue-600 cursor-pointer hover:font-semibold"
              >
                Going
              </small>
              <small
                onClick={() => setStatus("not_going")}
                className="text-pink-600 cursor-pointer hover:font-semibold"
              >
                Not going
              </small>
            </div>
          )}
          <span className="font-bold">{event.going_count}</span> will
          participate to this event
        </div>
      </div>
    </div>
  );
}

export default EventCard;
