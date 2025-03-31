"use client";
import ProfileCard from "@/_components/cards/profileCard";
import Aside from "@/_components/layout/aside";
import PostModal from "@/_components/modals/new_post";
import FriendSuggestions from "@/_components/section/friendSuggestions";
import SocketService from "@/_components/shared/SocketService";
import { useWebSocket } from "@/context/realTimeContext";
import useOpenModal from "@/hooks/useOpenModal";
import useRealTimeAction from "@/hooks/useRealTimeAction";
import { PlusIcon } from "@heroicons/react/24/outline";
import { useEffect } from "react";
export default function AppLayout({ children }) {
  const { openModal, closeModal, ModalComponent } = useOpenModal();

  return (
    <div>
      <div className="flex h-screen bg-white">
        <div className="flex-1 flex flex-col overflow-hidden">
          <div className="flex h-full">
            <Aside />
            <main className="flex flex-col py-6 relative w-full bg-white overflow-x-hidden overflow-y-auto mb-14">
              {children}
              <div
                onClick={openModal}
                className="fixed z-[60] h-fit   right-1/4 top-4  rounded-full p-2 bg-blue-500/80 text-white cursor-pointer"
              >
                <PlusIcon className="w-6 h-6" />
              </div>
              <ModalComponent>
                <PostModal close={closeModal} />
              </ModalComponent>
            </main>
            <nav className="lg:flex flex-col gap-4 px-4 overflow-x-hidden overflow-y-auto py-6 w-[450px] hidden  h-full border-l border-gray-200 ">
              <ProfileCard />
              <FriendSuggestions />
            </nav>
          </div>
        </div>
      </div>
    </div>
  );
}
