"use client";
import GroupMessageFeed from "@/_components/section/groups/groupMessages";
import { useWebSocket } from "@/context/realTimeContext";
import { ArrowLeftIcon, BackwardIcon } from "@heroicons/react/24/outline";
import Link from "next/link";
import React from "react";

function GroupChat() {
  return (
    <div>
      <Link
        className="flex gap-2 items-center px-4 hover:text-blue-600"
        href="/app/groups"
      >
        <ArrowLeftIcon className="w-5 h-5" />
        <p>Back to group</p>
      </Link>
      <GroupMessageFeed />
    </div>
  );
}

export default GroupChat;
