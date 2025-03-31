import useFetch from "@/hooks/useFetch";
import React from "react";
import { useQuery } from "react-query";
import UserCard from "../cards/userMessageTeaser";
import Link from "next/link";
import usePaginatedFetch from "@/hooks/usePaginatedFetch";

function FriendSuggestions() {
  const fetcher = useFetch();

  const initialEndpoint = "users";
  const {
    isLoading,
    uniqueArray: friends,
    loadMore,
    errors,
    refetch,
    hasNextPage,
  } = usePaginatedFetch(fetcher, initialEndpoint);

  return (
    <section className="space-y-3">
      <h1 className="text-lg font-bold">Friends suggestions</h1>
      {isLoading ? (
        <div className="divide-y">
          <div className="flex cursor-pointer gap-x-4 py-4">
            <div className="h-12 w-12 flex-none rounded-full bg-gray-200 animate-pulse"></div>
            <div className="min-w-0">
              <div className="h-4 w-[100px] bg-gray-200 rounded mb-2 animate-pulse"></div>
              <div className="h-4 w-[160px] bg-gray-200 rounded animate-pulse"></div>
            </div>
          </div>
        </div>
      ) : (
        <div className="shadow-md px-2 divide-y-2 border-2 border-gray-100 rounded-md">
          {friends &&
            friends?.map((friend, i) => (
              <Link key={i} href={`/app/profile/${friend.id}`}>
                <UserCard person={friend} />
              </Link>
            ))}
          {hasNextPage ? (
            <div
              onClick={loadMore}
              className=" py-1 text-blue-600 hover:underline cursor-pointer mx-auto w-fit"
            >
              <small>View more friends</small>
            </div>
          ) : (
            <div className="mx-auto w-fit text-xs">
              No more friends avaliable
            </div>
          )}
        </div>
      )}
    </section>
  );
}

export default FriendSuggestions;
