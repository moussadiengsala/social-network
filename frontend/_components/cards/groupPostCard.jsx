import React from "react";

import Image from "next/image";
import Link from "next/link";
import ReactTimeAgo from "react-time-ago";
import { ChatBubbleOvalLeftEllipsisIcon } from "@heroicons/react/24/outline";
import { useGroupPostStore } from "@/_store/zustand";
import GroupPostCommentModal from "../modals/group_post_comment_modal";
import useOpenModal from "@/hooks/useOpenModal";

const GroupPostCard = ({ groupPost }) => {
  const setGroupPost = useGroupPostStore((state) => state.setGroupPost);
  const { openModal, closeModal, ModalComponent } = useOpenModal();
  const viewMore = () => {
    openModal();
    setGroupPost(groupPost);
  };

  return (
    <div className="relative space-y-2 overflow-hidden text-xs text-gray-800 border-2 border-gray-100 p-2 rounded-md shadow-md">
      <ModalComponent>
        <GroupPostCommentModal close={closeModal} />
      </ModalComponent>
      <div className="flex items-center gap-1">
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
        <Link href={`/app/profile/${groupPost.author_id}`}>
          <h4 className="text-black">
            {groupPost.first_name} {groupPost.last_name}
          </h4>
        </Link>
        <div className="text-xs text-gray-500">
          <ReactTimeAgo date={groupPost.created_at} locale="en-US" />
        </div>
      </div>
      <div className="gap-2 flex">
        {groupPost.image && (
          <Image
            width={100}
            height={100}
            className="h-[80px] flex-1 rounded-md"
            src={`http://localhost:8000/api/static/uploads/${groupPost.image}`}
            alt="post image"
          />
        )}
        <div className="w-[70%]">
          <p>{groupPost.content}</p>
          <div
            onClick={viewMore}
            className="flex text-xs items-center text-blue-600 cursor-pointer hover:underline absolute bottom-0 right-2 gap-2 mt-2"
          >
            Comments <ChatBubbleOvalLeftEllipsisIcon className="w-5 h-6" />
          </div>
        </div>
      </div>
    </div>
  );
};

export default GroupPostCard;
