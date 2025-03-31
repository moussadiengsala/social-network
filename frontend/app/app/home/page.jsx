"use client";
import PostCard from "@/_components/cards/postCard";
import Alert from "@/_components/shared/alert";
import PostSkeleton from "@/_components/shared/skeletons/postCardSkeleton";
import useFetch from "@/hooks/useFetch";
import usePaginatedFetch from "@/hooks/usePaginatedFetch";
import { ArrowPathIcon } from "@heroicons/react/24/outline";

function Home() {
  const fetcher = useFetch();
  const initialEndpoint = "posts";
  const {
    isLoading,
    uniqueArray: data,
    loadMore,
    errors,
    refetch,
    hasNextPage,
  } = usePaginatedFetch(fetcher, initialEndpoint, 2, "posts");

  return (
    <div className="px-4 sm:px-6 lg:px-8">
      {isLoading ? (
        <div>
          <PostSkeleton />
          <PostSkeleton />
          <PostSkeleton />
          <PostSkeleton />
          <PostSkeleton />
        </div>
      ) : (
        <>
          <section className="space-y-8">
            {errors.length ? (
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
            ) : (
              <>
                {data && (
                  <>
                    {data.length ? (
                      data.map((post, index) => (
                        <PostCard key={index} post={post} />
                      ))
                    ) : (
                      <div className="flex text-gray-400 text-sm py-4 flex-col  items-center">
                        No post yet in your timeline
                      </div>
                    )}
                    {hasNextPage ? (
                      <div
                        onClick={loadMore}
                        className="bg-blue-500 text-center text-white w-fit px-3 rounded-md shadow-md py-2 mx-auto cursor-pointer"
                      >
                        Load More Posts
                      </div>
                    ) : (
                      <div className="text-gray-700 text-center  mx-auto w-fit">
                        <small>No more data to display in your timeline</small>
                      </div>
                    )}
                  </>
                )}
              </>
            )}
          </section>
        </>
      )}
    </div>
  );
}

export default Home;
