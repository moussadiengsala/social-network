import { usePostStore } from "@/_store/zustand";
import useApiRequest from "@/hooks/useApiRequest";
import useOpenModal from "@/hooks/useOpenModal";
import { CommentSchema } from "@/lib/validations/commentValidation";
import {
  EllipsisVerticalIcon,
  PaperAirplaneIcon,
  PaperClipIcon,
} from "@heroicons/react/24/outline";

import useFilePreview from "@/hooks/useFilePreview";
import { LikeSchema } from "@/lib/validations/likeValidation";
import {
  ChatBubbleOvalLeftEllipsisIcon,
  HeartIcon,
} from "@heroicons/react/24/solid";
import Image from "next/image";
import Link from "next/link";
import { useEffect, useState } from "react";
import ReactTimeAgo from "react-time-ago";
import CustomInput from "../inputs/input";
import TextArea from "../inputs/textArea";
import PostDetailsModal from "../modals/post_details";
import Alert from "../shared/alert";

function PostCard({ post, isCommentInputDisplayed = true }) {
  const setPost = usePostStore((state) => state.setPost);

  const { openModal, closeModal, ModalComponent } = useOpenModal();

  const OPTIONS = {
    method: "POST",
    endpoint: "comments",
    data: {},
    format: "--formData",
    dataToRefresh: "posts",
  };

  const {
    responseError,
    register,
    handleSubmit,
    errors,
    watch,
    isLoading,
    isSuccess,
    onSubmit,
  } = useApiRequest(OPTIONS, CommentSchema);

  const {
    responseError: likeErr,
    register: registerLike,
    handleSubmit: handleSubmitLike,
    errors: errorsLike,
    isLoading: isLoadingLike,

    onSubmit: onSubmitLike,
  } = useApiRequest(
    {
      method: "POST",
      endpoint: "reactions",
      dataToRefresh: "posts",
    },
    LikeSchema
  );

  const file = watch("file");

  const [filePreview] = useFilePreview(file);

  const viewMore = () => {
    openModal();
    setPost(post);
  };
  useEffect(() => {
    if (isSuccess) {
      viewMore();
    }
  }, [isSuccess]);

  return (
    <div className="bg-white border-2 border-gray-100 text-gray-700 rounded-xl  shadow-sm text-sm font-medium">
      <ModalComponent>
        <PostDetailsModal close={closeModal} />
      </ModalComponent>
      <div className="flex gap-3 sm:p-4 p-2.5 text-sm font-medium">
        <div>
          <Image
            width={100}
            height={100}
            src={
              post.avatar != ""
                ? `http://localhost:8000/api/static/uploads/${post.avatar}`
                : "/unknown.jpg"
            }
            alt="user profile picture"
            className="w-9 h-9 rounded-full"
          />
        </div>
        <div className="flex-1">
          <Link
            className="hover:underline"
            href={`/app/profile/${post.author_id}`}
          >
            <h4 className="text-black">
              {post.first_name} {post.last_name}
            </h4>
          </Link>
          <div className="text-xs text-gray-500">
            <ReactTimeAgo date={post.created_at} locale="en-US" />
          </div>
        </div>

        <div className="-mr-1">
          <EllipsisVerticalIcon className="w-6 h-6" />
        </div>
      </div>
      <div className="sm:px-4 p-2.5 pt-0">
        {post.image && (
          <Image
            width={100}
            height={100}
            className="w-full h-[150px] sm:h-[250px] lg:h-[330px] object-cover mb-4 rounded-lg"
            src={`http://localhost:8000/api/static/uploads/${post.image}`}
            alt="post image"
          />
        )}

        <p className="font-normal">{post.content}</p>
      </div>
      {/* rEAC */}
      <div className="sm:p-4 p-2.5 flex items-center gap-4 text-xs font-semibold">
        <div className="flex items-center gap-2">{/* LIKE */}</div>

        <div className="flex items-center gap-3">
          <button onClick={viewMore} type="button" className="button-icon">
            {" "}
            <ChatBubbleOvalLeftEllipsisIcon className="w-5 h-6" />
          </button>
          <span>{post.comments}</span>
        </div>
      </div>
      <div className="sm:p-4 p-2.5 border-t border-gray-100 font-normal space-y-3 relative"></div>
      {filePreview ? (
        <Image
          width={100}
          height={100}
          className="rounded-lg px-2  h-[80] w-[80]  object-cover"
          src={filePreview}
          alt="preview"
        />
      ) : null}
      {isCommentInputDisplayed && (
        <div className="sm:px-4  p-2.5  flex items-center gap-1">
          <Image
            width={100}
            height={100}
            src={
              post.avatar != ""
                ? `http://localhost:8000/api/static/uploads/${post.avatar}`
                : "/unknown.jpg"
            }
            alt="user profile picture"
            className="w-9 h-9 rounded-full"
          />
          {/* Comment  */}
          <form
            method="POST"
            onSubmit={handleSubmit(onSubmit)}
            className="flex w-full items-center"
          >
            <div className="flex-1 relative overflow-hidden h-10">
              <input type="hidden" name="entrie_id" value={post.id} />

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
        </div>
      )}
      <div className="px-3">
        {responseError.length !== 0 && (
          <Alert message={responseError} status="error" />
        )}
      </div>
    </div>
  );
}

export default PostCard;
