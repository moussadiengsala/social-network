import useFetch from "@/hooks/useFetch";
import usePaginatedFetch from "@/hooks/usePaginatedFetch";
import React, { useEffect } from "react";
import Alert from "../shared/alert";
import PostSkeleton from "../shared/skeletons/postCardSkeleton";
import PostCard from "../cards/postCard";
import { ArrowPathIcon } from "@heroicons/react/24/outline";
import { queryClient } from "@/app/layout";

function ProfileFeed({ section, id, key }) {
  const fetcher = useFetch();
  const initialEndpoint = `feeds?section=${section}&user_id=${id}`;

  const {
    isLoading,
    uniqueArray: feed,
    loadMore,
    refetch,
    errors,
    hasNextPage,
  } = usePaginatedFetch(fetcher, initialEndpoint, 5, "profile-feed");
  useEffect(() => {
    queryClient.refetchQueries(["profile-feed"], { active: true });
  }, [section, key]);
  return (
    <div>
      <h1 className="text-xl mb-4 font-bold capitalize">{`${
        section.split("_")[0]
      } ${section.split("_")[1]}`}</h1>
      <div>
        {isLoading ? (
          <div>
            <PostSkeleton />
            <PostSkeleton />
            <PostSkeleton />
            <PostSkeleton />
            <PostSkeleton />
          </div>
        ) : (
          <section className="space-y-8">
            {feed && (
              <>
                {feed.length > 0 ? (
                  feed.map((post) => <PostCard key={post.id} post={post} />)
                ) : (
                  <div className="flex text-gray-400 text-sm py-4 flex-col items-center">
                    No posts yet in your timeline
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
                  <div className="text-gray-700 text-center mx-auto w-fit">
                    <small>No more data to display in your timeline</small>
                  </div>
                )}
              </>
            )}
            {errors.length > 0 && (
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
          </section>
        )}
      </div>
    </div>
  );
}
export default ProfileFeed;
