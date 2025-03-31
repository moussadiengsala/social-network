import EventCard from "@/_components/cards/eventCard";
import NewEventModal from "@/_components/modals/new_event";
import GroupPostSkeleton from "@/_components/shared/skeletons/groupPostSkeleton";
import { useGroupFeedStore } from "@/_store/zustand";
import useFetch from "@/hooks/useFetch";
import useOpenModal from "@/hooks/useOpenModal";
import usePaginatedFetch from "@/hooks/usePaginatedFetch";
import { PlusCircleIcon } from "@heroicons/react/24/outline";

function GroupEventFeed() {
  const { openModal, closeModal, ModalComponent } = useOpenModal();
  const group = useGroupFeedStore((state) => state.group);
  const fetcher = useFetch(); // Replace useFetch with your actual fetch function
  const initialEndpoint = `groups-events?group_id=${group.id}`;
  const {
    isLoading,
    uniqueArray: groupEvents,
    loadMore,
    errors,
    hasNextPage,
  } = usePaginatedFetch(fetcher, initialEndpoint, 2, "groups-events");
  console.log(groupEvents);
  return (
    <div className="px-4">
      <ModalComponent>
        <NewEventModal close={closeModal} />
      </ModalComponent>
      <div className="flex items-center mb-2 justify-between gap-1">
        <h1 className="text-xl font-bold mb-2">Events</h1>
        <div
          onClick={openModal}
          className="text-sm flex bg-blue-100 text-blue-600 gap-1 px-2 cursor-pointer py-1 rounded-md"
        >
          <PlusCircleIcon className="w-5 h-5" />
          <small>New Event</small>
        </div>
      </div>
      <div className="space-y-8 h-[600px]">
        {isLoading ? (
          <>
            <GroupPostSkeleton />
            <GroupPostSkeleton />
            <GroupPostSkeleton />
            <GroupPostSkeleton />
          </>
        ) : (
          <>
            {groupEvents.length ? (
              <>
                {groupEvents.map((event, i) => (
                  <EventCard key={i} event={event} />
                ))}
                {hasNextPage ? (
                  <div
                    onClick={loadMore}
                    className=" py-1 text-blue-600 hover:underline cursor-pointer mx-auto w-fit"
                  >
                    <small>View more event</small>
                  </div>
                ) : (
                  <div className="mx-auto w-fit text-xs">No more events</div>
                )}
              </>
            ) : (
              <div className="text-xs h-full text-gray-600  text-center italic">
                <span className="font-bold">{group.title}</span> doesn't have
                any post yet
              </div>
            )}
          </>
        )}
        {errors.length ? <Alert message={errors} status="error" /> : null}
      </div>
    </div>
  );
}

export default GroupEventFeed;
