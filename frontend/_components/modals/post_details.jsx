import { usePostStore } from "@/_store/zustand";
import { Dialog, Transition } from "@headlessui/react";
import { Fragment } from "react";
import PostCard from "../cards/postCard";
import useFetch from "@/hooks/useFetch";
import { useQuery } from "react-query";
import Comment from "../cards/commentCard";
import CommentSkeleton from "../shared/skeletons/commentSkeleton";
import Alert from "../shared/alert";
import usePaginatedFetch from "@/hooks/usePaginatedFetch";
function PostDetailsModal({ close }) {
  const post = usePostStore((state) => state.post);

  const fetcher = useFetch();

  const initialEndpoint = `comments?post_id=${post.id}`;
  const {
    isLoading,
    uniqueArray: comments,
    loadMore,
    errors,
    refetch,
    hasNextPage,
  } = usePaginatedFetch(fetcher, initialEndpoint, 2, "comments");

  return (
    <Transition.Root show={true} as={Fragment}>
      <Dialog as="div" className="relative z-1" onClose={close}>
        <Transition.Child
          as={Fragment}
          enter="ease-out duration-300"
          enterFrom="opacity-0"
          enterTo="opacity-100"
          leave="ease-in duration-200"
          leaveFrom="opacity-100"
          leaveTo="opacity-0"
        >
          <div className="fixed inset-0  z-[60] bg-gray-500 bg-opacity-75 transition-opacity" />
        </Transition.Child>

        <div className="fixed inset-0 z-[99] w-screen overflow-y-auto">
          <div className="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
            <Transition.Child
              as={Fragment}
              enter="ease-out duration-300"
              enterFrom="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
              enterTo="opacity-100 translate-y-0 sm:scale-100"
              leave="ease-in duration-200"
              leaveFrom="opacity-100 translate-y-0 sm:scale-100"
              leaveTo="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
            >
              <Dialog.Panel
                className="relative transform overflow-hidden rounded-lg bg-white px-4 pb-4 pt-5 text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-xl
               sm:p-6"
              >
                <div className="space-y-8">
                  <PostCard post={post} isCommentInputDisplayed={false} />
                  <div className="space-y-2">
                    <h1 className="text-xl font-bold">Comments</h1>
                    {isLoading ? (
                      <CommentSkeleton />
                    ) : comments ? (
                      comments.map((comment, i) => (
                        <Comment key={i} comment={comment} />
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
                  <button
                    type="button"
                    onClick={close}
                    className="inline-block w-full shrink-0 rounded-md border border-pink-600 bg-transparent px-12 py-3 text-sm font-medium text-pink-600 transition hover:bg-pink-600 hover:text-white focus:outline-none focus:ring active:text-pink-500"
                  >
                    Close
                  </button>
                </div>
              </Dialog.Panel>
            </Transition.Child>
          </div>
        </div>
      </Dialog>
    </Transition.Root>
  );
}

export default PostDetailsModal;
