import useFetch from "@/hooks/useFetch";
import usePaginatedFetch from "@/hooks/usePaginatedFetch";
import Image from "next/image";

function GroupEventTab() {
  const fetcher = useFetch();
  const initialEndpoint = `notifications?notification=groups_events`;

  const {
    isLoading,
    uniqueArray: notifications,
    loadMore,
    hasNextPage,
  } = usePaginatedFetch(
    fetcher,
    initialEndpoint,
    3,
    "groups-event-notifications"
  );
  console.log(notifications);
  return (
    <section>
      <h1 className="text-xl font-bold capitalize mb-2">Group Event</h1>
      {isLoading ? (
        <>Loading</>
      ) : (
        notifications?.map((notification, i) => (
          <div key={i} className="flex items-center my-4 divide-y-1  gap-1">
            <div className="flex italic font-semibold items-center gap-2">
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
              {`${notification?.sender_first_name} ${notification?.sender_last_name} `}
            </div>
            create a new event on
            <span className="italic font-semibold">
              {notification?.group_title}
            </span>
          </div>
        ))
      )}
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
    </section>
  );
}

export default GroupEventTab;
