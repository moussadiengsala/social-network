"use client";
import { useChatStore } from "@/_store/zustand";
import Image from "next/image";
import React from "react";

const UserCard = ({ person, additionalAttributes }) => {
  const setChat = useChatStore((state) => state.setChat);

  return (
    <li
      key={person.id}
      {...additionalAttributes}
      className="flex cursor-pointer gap-x-4  py-4"
      onClick={() => {
        if (!person.followers) setChat(person);
      }}
    >
      <div className="relative">
        <Image
          width={100}
          height={100}
          className="h-12 w-12 flex-none rounded-full bg-gray-50"
          src={
            person.avatar != ""
              ? `http://localhost:8000/api/static/uploads/${person.avatar}`
              : "/unknown.jpg"
          }
          alt=""
        />
        {person.online ? (
          <div className="flex-none absolute bottom-2 right-0 rounded-full bg-emerald-500/20 p-1">
            <div className="h-1.5 w-1.5 rounded-full bg-emerald-500" />
          </div>
        ) : (
          <></>
        )}
      </div>
      <div className="min-w-0">
        <p
          className={`text-sm ${
            person.unread_message &&
            person.unread_message !== 0 &&
            "font-semibold"
          } truncate  leading-6 text-gray-900`}
        >
          {`${person.first_name} ${person.last_name}`}{" "}
          {person.unread_message && person.unread_message !== 0 ? (
            <span className=" italic">
              have sent you {person.unread_message} messages
            </span>
          ) : null}
        </p>
        <p className="mt-1 truncate w-fit bg-blue-100  text-blue-600 px-2 py-1 rounded-md text-xs leading-5 ">
          {person.email}
        </p>
      </div>
    </li>
  );
};

export default UserCard;
