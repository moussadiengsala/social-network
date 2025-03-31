"use client";
import GroupeTeaser from "@/_components/cards/groupTeaser";
import GroupModal from "@/_components/modals/new_group";
import GroupFeed from "@/_components/section/GroupFeed";
import Alert from "@/_components/shared/alert";
import { useBelongedGroupStore } from "@/_store/zustand";
import { useWebSocket } from "@/context/realTimeContext";
import useFetch from "@/hooks/useFetch";
import useOpenModal from "@/hooks/useOpenModal";
import usePaginatedFetch from "@/hooks/usePaginatedFetch";
import useRealTimeAction from "@/hooks/useRealTimeAction";
import SiteConfig from "@/lib/site.config";
import { ArrowPathIcon, PlusCircleIcon } from "@heroicons/react/24/outline";

function Groups() {
  const { belongedGroups } = useWebSocket();
  console.log(belongedGroups);

  const { openModal, closeModal, ModalComponent } = useOpenModal();
  const fetcher = useFetch();
  const initialEndpoint = "groups";
  const {
    isLoading,
    uniqueArray: groupData,
    loadMore,
    errors,
    refetch,
    hasNextPage,
  } = usePaginatedFetch(
    fetcher,
    initialEndpoint,
    2,
    "groups-suggestions",
    false
  );

  return (
    <div className="flex">
      <ModalComponent>
        <GroupModal close={closeModal} />
      </ModalComponent>
      <nav className="h-full flex-1 overflow-y-auto" aria-label="Directory">
        <div className="flex items-center justify-between px-2">
          <h1 className="text-xl font-bold px-6 py-2">Groups</h1>
          <small
            onClick={openModal}
            className="bg-emerald-100 cursor-pointer px-2 flex gap-1 rounded-md py-1 text-emerald-500"
          >
            <PlusCircleIcon className="w-5 h-5" />
            New group
          </small>
        </div>
        <section>
          <div className="sticky top-0 z-10 border-y border-b-gray-200 border-t-gray-100 bg-gray-50 px-3 py-1.5 text-sm font-semibold leading-6 text-gray-900">
            <h3 className="px-6">Groups suggestions</h3>
          </div>
          {isLoading ? (
            <div className="divide-y px-6">
              <div className="flex cursor-pointer gap-x-4 py-4">
                <div className="h-12 w-12 flex-none rounded-full bg-gray-200 animate-pulse"></div>
                <div className="min-w-0">
                  <div className="h-4 w-[100px] bg-gray-200 rounded mb-2 animate-pulse"></div>
                  <div className="h-4 w-[160px] bg-gray-200 rounded animate-pulse"></div>
                </div>
              </div>
            </div>
          ) : (
            <>
              <div role="list" className="divide-y px-6 divide-gray-100">
                {groupData ? (
                  groupData.map((group, index) => (
                    <GroupeTeaser key={index} group={group} />
                  ))
                ) : (
                  <div className="flex flex-col items-center">
                    <Alert message={errors} status="error" />
                    <small
                      className="text-pink-600 bg-pink-50 px-2 py-1 rounded-md flex items-center gap-1 text-center mx-auto cursor-pointer hover:underline"
                      onClick={() => refetch()}
                    >
                      <ArrowPathIcon className="w-4 h-4" />
                      Try again
                    </small>
                  </div>
                )}
              </div>
              <div
                onClick={loadMore}
                className=" py-1 text-blue-600 hover:underline cursor-pointer mx-auto w-fit"
              >
                <small>
                  {hasNextPage ? "View more group" : "No more groups available"}
                </small>
              </div>
            </>
          )}
        </section>
        <section>
          <div className="sticky top-0 z-10 border-y border-b-gray-200 bg-gray-50 border-t-gray-100  px-3 py-1.5 text-sm font-semibold leading-6 text-gray-900">
            <h3 className="px-6">My groups</h3>
          </div>
          <ul role="list" className="divide-y px-6 divide-gray-100">
            {belongedGroups?.length != 0 ? (
              belongedGroups?.map((group, index) => (
                <div key={index}>
                  <GroupeTeaser key={index} group={group} />
                  {group.unread_message !== 0 && (
                    <div className="italic  justify-end text-sm flex items-center gap-2">
                      <div className="w-2 h-2 rounded-full bg-red-500"></div>
                      You missed {group.unread_message} message from this group
                    </div>
                  )}
                </div>
              ))
            ) : (
              <div className="flex text-gray-400 text-sm py-4 flex-col  items-center">
                You not part of any group yet
              </div>
            )}
          </ul>
        </section>
      </nav>
      <GroupFeed />
    </div>
  );
}

export default Groups;
