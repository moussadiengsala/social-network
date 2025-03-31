import useFetch from "@/hooks/useFetch";
import usePaginatedFetch from "@/hooks/usePaginatedFetch";
import { EnvelopeIcon, PhoneIcon } from "@heroicons/react/24/solid";
import Image from "next/image";
import React from "react";

function ProfileCard() {
  const fetcher = useFetch();

  const initialEndpoint = "groups";
  const {
    isLoading,
    uniqueArray: groupData,
    loadMore,
    errors,
    refetch,
    hasNextPage,
  } = usePaginatedFetch(fetcher, initialEndpoint, 3, "groups-suggestions");

  return (
    <li className="col-span-1 flex flex-col divide-y divide-gray-200 rounded-lg bg-white text-center border-2 border-gray-100 shadow-md">
      <div className="flex flex-1 flex-col px-4 py-3">
        <Image
          width={100}
          height={100}
          className="mx-auto h-28 w-28 flex-shrink-0 rounded-full"
          src="https://source.unsplash.com/300x400"
          alt=""
        />
        <h3 className="mt-6 text-sm font-medium text-gray-900">
          Abdou Aziz Ndiaye
        </h3>
        <dl className="mt-1 flex flex-grow flex-col justify-between">
          <dt className="sr-only">Title</dt>
          <dd className="text-sm text-gray-500">@darkcode</dd>
          <dt className="sr-only">Account stats</dt>
          <div className="flex gap-x-2 gap-y-1 flex-wrap items-center">
            <dd className="mt-3">
              <span className="inline-flex items-center rounded-full bg-green-50 px-2 py-1 text-xs font-medium text-green-700 ring-1 ring-inset ring-green-600/20">
                2 post(s)
              </span>
            </dd>
            <dd className="mt-3">
              <span className="inline-flex items-center rounded-full bg-pink-50 px-2 py-1 text-xs font-medium text-pink-700 ring-1 ring-inset ring-pink-600/20">
                1K follower(s)
              </span>
            </dd>
            <dd className="mt-3">
              <span className="inline-flex items-center rounded-full bg-blue-50 px-2 py-1 text-xs font-medium text-blue-700 ring-1 ring-inset ring-blue-600/20">
                1 group(s)
              </span>
            </dd>
          </div>
          <small className="font-light mt-6 text-gray-500">
            Member since June 2023
          </small>
        </dl>
      </div>
    </li>
  );
}

export default ProfileCard;
