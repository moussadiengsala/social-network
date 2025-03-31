import { removeDuplicates } from "@/lib/removeduplicated";
import { useState } from "react";
import { useQuery } from "react-query";

const usePaginatedFetch = (
  fetcher,
  initialEndpoint,
  limit = 2,
  queryName = "default",
  fetchOnFocus = false
) => {
  const separator = initialEndpoint.includes("?") ? "&" : "?";
  const [offset, setOffset] = useState(0);
  const [endpoint, setEndpoint] = useState(initialEndpoint);
  const [hasNextPage, setHasNextPage] = useState(initialEndpoint);

  const [data, setData] = useState([]);
  const [errors, setErrors] = useState([]);
  const option = {
    method: "GET",
    endpoint: `${initialEndpoint}${separator}limit=${limit}&offset=${offset}`,
  };
  const { isLoading, refetch } = useQuery(
    [queryName, option],
    () => fetcher(option),
    {
      onSuccess: (successData) => {
        setHasNextPage(successData?.data?.length > 0);
        setData((prevData) => [...successData?.data, ...prevData]);
      },
      onError: (data) => setErrors(data.message.split("*")),
      refetchOnWindowFocus: fetchOnFocus,
    }
  );

  const loadMore = () => {
    const newOffset = offset + limit;
    setOffset(newOffset);
    setEndpoint(
      `${initialEndpoint}${separator}limit=${limit}&offset=${newOffset}`
    );
  };
  const uniqueArray = removeDuplicates(data, "id");

  return {
    isLoading,
    uniqueArray,
    loadMore,
    refetch,
    errors,
    hasNextPage,
  };
};

export default usePaginatedFetch;
