import { useGroupPostStore } from "@/_store/zustand";
import useApiRequest from "@/hooks/useApiRequest";
import useFetch from "@/hooks/useFetch";
import useFilePreview from "@/hooks/useFilePreview";
import usePaginatedFetch from "@/hooks/usePaginatedFetch";
import { CommentSchema } from "@/lib/validations/commentValidation";
import { Dialog, Transition } from "@headlessui/react";
import { PaperAirplaneIcon, PaperClipIcon } from "@heroicons/react/24/outline";
import Image from "next/image";
import Link from "next/link";
import { Fragment } from "react";
import ReactTimeAgo from "react-time-ago";
import CustomInput from "../inputs/input";
import TextArea from "../inputs/textArea";
import Alert from "../shared/alert";

function GroupPostCommentModal({ close }) {
  const fetcher = useFetch();
  const groupPost = useGroupPostStore((state) => state.groupPost);
  const OPTIONS = {
    method: "POST",
    endpoint: "groups-comments",
    data: {},
    format: "--formData",
    dataToRefresh: "groups-post-comments",
  };

  const { register, handleSubmit, errors, watch, isLoading, onSubmit } =
    useApiRequest(OPTIONS, CommentSchema);
  const file = watch("file");
  const [filePreview] = useFilePreview(file);
  const initialEndpoint = `groups-comments?post_id=${groupPost.id}`;
  const {
    uniqueArray: comments,
    loadMore,
    errors: errorsComments,
    hasNextPage,
  } = usePaginatedFetch(fetcher, initialEndpoint, 2, "groups-post-comments");

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
              <Dialog.Panel className="relative transform overflow-hidden rounded-lg bg-white px-4 pb-4 pt-5 text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-2xl sm:p-6">
                <div className="space-y-3">
                  <h1 className="text-xl font-bold">
                    <span className="italic">
                      {groupPost.group_title}'s post
                    </span>
                  </h1>
                  <div className="flex items-center gap-1">
                    <div className="text-xs text-gray-500"></div>
                  </div>
                  <div className="gap-2 flex">
                    {groupPost.image && (
                      <Image
                        width={100}
                        height={100}
                        className="h-[150px] flex-1 rounded-md"
                        src={`http://localhost:8000/api/static/uploads/${groupPost.image}`}
                        alt="post image"
                      />
                    )}
                    <div className="w-[70%]">
                      <p>{groupPost.content}</p>
                      <div className="text-sm flex items-center gap-2 text-gray-600">
                        <Image
                          width={100}
                          height={100}
                          src={
                            groupPost.avatar !== ""
                              ? `http://localhost:8000/api/static/uploads/${groupPost.avatar}`
                              : "/unknown.jpg"
                          }
                          alt="user profile picture"
                          className="w-6 h-6 rounded-full"
                        />
                        by
                        <Link
                          href={`/app/profile/${groupPost.author_id}`}
                          className="font-bold hover:underline"
                        >
                          {groupPost.first_name} {groupPost.last_name}
                        </Link>{" "}
                        <ReactTimeAgo
                          date={groupPost.created_at}
                          locale="en-US"
                        />
                      </div>
                    </div>
                  </div>
                  {filePreview ? (
                    <Image
                      width={100}
                      height={100}
                      className="rounded-lg px-2  h-[80] w-[80]  object-cover"
                      src={filePreview}
                      alt="preview"
                    />
                  ) : null}
                  <form
                    method="POST"
                    onSubmit={handleSubmit(onSubmit)}
                    className="flex w-full items-center"
                  >
                    <div className="flex-1 relative overflow-hidden h-10">
                      <input
                        type="hidden"
                        name="post_id"
                        value={groupPost.id}
                      />
                      <input
                        type="hidden"
                        name="group_id"
                        value={groupPost.group_id}
                      />
                      <TextArea
                        register={register}
                        additionalAttributes={{
                          className:
                            "w-full resize-none !bg-transparent px-4 py-2 focus:!border-transparent focus:!ring-transparent",
                        }}
                        name="content"
                        error={errors.content}
                        rows={4}
                        placeholder="add your comment..."
                      />

                      <div
                        className="!top-2 pr-2 uk-drop"
                        uk-drop="pos: bottom-right; mode: click"
                      >
                        <div
                          className="flex items-center gap-2"
                          uk-scrollspy="target: > svg; cls: uk-animation-slide-right-small; delay: 100 ;repeat: true"
                        >
                          <svg
                            xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 24 24"
                            fill="currentColor"
                            className="w-6 h-6 fill-sky-600"
                            style={{ opacity: 0 }}
                          >
                            <path
                              fillRule="evenodd"
                              d="M1.5 6a2.25 2.25 0 012.25-2.25h16.5A2.25 2.25 0 0122.5 6v12a2.25 2.25 0 01-2.25 2.25H3.75A2.25 2.25 0 011.5 18V6zM3 16.06V18c0 .414.336.75.75.75h16.5A.75.75 0 0021 18v-1.94l-2.69-2.689a1.5 1.5 0 00-2.12 0l-.88.879.97.97a.75.75 0 11-1.06 1.06l-5.16-5.159a1.5 1.5 0 00-2.12 0L3 16.061zm10.125-7.81a1.125 1.125 0 112.25 0 1.125 1.125 0 01-2.25 0z"
                              clipRule="evenodd"
                            ></path>
                          </svg>
                          <svg
                            xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 20 20"
                            fill="currentColor"
                            className="w-5 h-5 fill-pink-600"
                            style={{ opacity: 0 }}
                          >
                            <path d="M3.25 4A2.25 2.25 0 001 6.25v7.5A2.25 2.25 0 003.25 16h7.5A2.25 2.25 0 0013 13.75v-7.5A2.25 2.25 0 0010.75 4h-7.5zM19 4.75a.75.75 0 00-1.28-.53l-3 3a.75.75 0 00-.22.53v4.5c0 .199.079.39.22.53l3 3a.75.75 0 001.28-.53V4.75z"></path>
                          </svg>
                        </div>
                      </div>
                    </div>
                    <div>
                      <div className="flex items-center text-blue-600 gap-2">
                        <label
                          htmlFor="file"
                          className="rounded-md  my-2 px-2.5 py-1.5 text-sm font-semibold cursor-pointer"
                        >
                          <PaperClipIcon className="w-5 h-5" />
                        </label>
                      </div>
                      <CustomInput
                        type="file"
                        register={register}
                        label="photo"
                        placeholder=""
                        name="file"
                        className="hidden"
                      />
                    </div>
                    <button
                      type="submit"
                      disabled={isLoading}
                      className="flex items-center gap-2 text-white rounded-md py-1.5 px-3.5 bg-blue-500"
                    >
                      Reply
                      <PaperAirplaneIcon className="w-4 h-4" />
                    </button>
                  </form>
                  <div className="space-y-2">
                    <h1 className="text-xl font-bold">Comments</h1>
                    {isLoading ? (
                      <div className="px-2 py-4 space-y-6">
                        <p className="px-2 py-1 w-[200px] h-[30px]  relative text-sm rounded-md bg-gray-200 animate-pulse"></p>
                        <p className="px-2 py-1 w-[200px] h-[30px]  relative text-sm rounded-md bg-gray-200 animate-pulse"></p>
                        <p className="px-2 py-1 w-[200px] h-[30px]  relative text-sm rounded-md bg-gray-200 animate-pulse"></p>
                      </div>
                    ) : (
                      <div className="py-4 overflow-scroll px-2 space-y-6 max-h-[300px] mb-4">
                        {comments && comments.length ? (
                          comments.map((comment,i) => (
                            <div key={i} className="relative border-2 border-gray-100 text-sm grid grid-cols-1 gap-4 p-4 mb-8  rounded-lg bg-white shadow-sm">
                              <div className="relative flex items-center mb-2 gap-4">
                                <Image
                                  width={100}
                                  height={100}
                                  src={
                                    comment.avatar != ""
                                      ? `http://localhost:8000/api/static/uploads/${comment.avatar}`
                                      : "/unknown.jpg"
                                  }
                                  className="relative rounded-full bg-white border w-9 h-8 object-cover"
                                  alt=""
                                  loading="lazy"
                                />
                                <div className="flex flex-col w-full">
                                  <div className="flex flex-row justify-between">
                                    <p className="relative  whitespace-nowrap truncate overflow-hidden">
                                      {comment.username}
                                    </p>
                                  </div>
                                  <div className="text-xs text-gray-500">
                                    <ReactTimeAgo
                                      date={comment.created_at}
                                      locale="en-US"
                                    />
                                  </div>
                                </div>
                              </div>
                              <div className="flex items-center gap-2">
                                {comment.image ? (
                                  <Image
                                    width={100}
                                    height={100}
                                    className="h-[80px] object-cover mb-4 rounded-lg"
                                    src={`http://localhost:8000/api/static/uploads/${comment.image}`}
                                    alt="comment image"
                                  />
                                ) : null}

                                <p className="-mt-4 text-gray-500 text-sm">
                                  {comment.content}
                                </p>
                              </div>
                            </div>
                          ))
                        ) : (
                          <div className="text-xs h-full text-gray-600  text-center italic">
                            No comment yet
                          </div>
                        )}
                        {errorsComments.length !== 0 && (
                          <Alert message={errorsComments} status="error" />
                        )}
                        {hasNextPage ? (
                          <div
                            onClick={loadMore}
                            className=" py-1 text-blue-600 hover:underline cursor-pointer mx-auto w-fit"
                          >
                            <small>See comments</small>
                          </div>
                        ) : (
                          <div className="mx-auto w-fit text-xs">
                            No more comments
                          </div>
                        )}
                      </div>
                    )}
                  </div>
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

export default GroupPostCommentModal;
