import useApiRequest from "@/hooks/useApiRequest";
import { LikeSchema } from "@/lib/validations/likeValidation";
import Image from "next/image";
import ReactTimeAgo from "react-time-ago";
import CustomInput from "../inputs/input";
import { HeartIcon } from "@heroicons/react/24/solid";

export default function Comment({ comment }) {
  const { responseError, register, handleSubmit, errors, isLoading, onSubmit } =
    useApiRequest(
      {
        method: "POST",
        endpoint: "reactions",
        dataToRefresh: "comments",
      },
      LikeSchema
    );

  return (
    <div className="relative border-2 border-gray-100 text-sm grid grid-cols-1 gap-4 p-4 mb-8  rounded-lg bg-white shadow-sm">
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
            <ReactTimeAgo date={comment.created_at} locale="en-US" />
          </div>
        </div>
      </div>
      {comment.image ? (
        <Image
          width={100}
          height={100}
          className="w-full h-[150px] sm:h-[250px] lg:h-[330px] object-cover mb-4 rounded-lg"
          src={`http://localhost:8000/api/static/uploads/${comment.image}`}
          alt="comment image"
        />
      ) : null}

      <p className="-mt-4 text-gray-500 text-sm">{comment.content}</p>
      <div className="flex items-center text-gray-700 gap-3 text-xs font-semibold">
        <form method="POST" onSubmit={handleSubmit(onSubmit)}>
          <CustomInput
            type="hidden"
            value={comment.comment_id}
            register={register}
            name="entries_id"
            error={errors.entries_id}
            className="col-span-6 sm:col-span-3 "
          />
          <CustomInput
            type="hidden"
            value="comment_like"
            register={register}
            error={errors.action}
            name="action"
            className="col-span-6 sm:col-span-3 "
          />

          {/* <button
            type="submit"
            disabled={isLoading}
            className={`button-icon w-5 h-5 disabled:opacity-25  ${
              comment.like_status &&
              "bg-blue-200 rounded-full  flex items-center justify-center  p-[1px]"
            } text-blue-500`}
          >
            {" "}
            <HeartIcon
              className={`w-4 h-4 ${
                comment.like_status ? "" : "hover:scale-125"
              } `}
            />
          </button> */}
        </form>
        {/* <span href="#">{comment.like_count}</span> */}
      </div>
    </div>
  );
}
