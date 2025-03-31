import { queryClient } from "@/app/layout";
import useFetch from "@/hooks/useFetch";
import usePaginatedFetch from "@/hooks/usePaginatedFetch";
import Image from "next/image";
import { useEffect } from "react";
import { useQuery } from "react-query";

function NotificationsTeaser({ notificationType, key }) {
  const fetcher = useFetch();

  const initialEndpoint = `notifications?notification=${notificationType}`;

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
 
  return (
    <section>
      <h1 className="text-xl font-bold capitalize mb-2">
        {notificationType.split("_")[0]} {notificationType.split("_")[1]}
      </h1>
    </section>
  );
}

export default NotificationsTeaser;
