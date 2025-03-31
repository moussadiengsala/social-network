import { useNotificationStore } from "@/_store/zustand";
import useOpenModal from "@/hooks/useOpenModal";
import {
  ArrowsRightLeftIcon,
  BellIcon,
  EnvelopeOpenIcon,
  UserGroupIcon,
} from "@heroicons/react/24/outline";
import { Accordion, AccordionItem } from "@nextui-org/accordion";
import NotificationsModal from "../modals/notifications_modal";
import { useState } from "react";
import { CalendarDaysIcon } from "@heroicons/react/24/solid";
import { useWebSocket } from "@/context/realTimeContext";

const NotificationAccordion = () => {
  const { notifications: notificationCount, friends } = useWebSocket();

  const [dataToRender, setDataToRender] = useState("follow_request");
  const { openModal, closeModal, ModalComponent } = useOpenModal();
  return (
    <>
      <ModalComponent>
        <NotificationsModal dataToRender={dataToRender} close={closeModal} />
      </ModalComponent>
      <Accordion className=" px-0 py-2">
        <AccordionItem
          key="1"
          className="space-y-2"
          aria-label="Notifications"
          title={
            <div className="flex px-0 gap-x-3">
              <div className="relative flex  ">
                <BellIcon className="h-6 w-6 shrink-0 hover:text-blue-600 text-gray-400" />
                {notificationCount > 0 && (
                  <div className="bg-pink-600 text-[8px] absolute top-0 right-0 text-white p-1 text-center flex items-center justify-center rounded-full h-3 w-3">
                    <p>{notificationCount}</p>
                  </div>
                )}
              </div>
              Notifications
            </div>
          }
        >
          <div
            onClick={() => {
              setDataToRender("follow_request");
              openModal();
            }}
            className="hover:text-blue-600 text-gray-700 flex items-center gap-1  px-4"
          >
            <ArrowsRightLeftIcon className="w-5 h-5" />
            Follow Request
          </div>
          <div
            onClick={() => {
              setDataToRender("groups_requested");
              openModal();
            }}
            className="hover:text-blue-600 flex items-center gap-1 text-gray-700 my-2 px-4"
          >
            <UserGroupIcon className="w-5 h-5" />
            Group Join Request
          </div>

          <div
            onClick={() => {
              setDataToRender("groups_events");
              openModal();
            }}
            className="hover:text-blue-600 text-gray-700 flex items-center gap-1 my-2 px-4"
          >
            <CalendarDaysIcon className="w-5 h-5" />
            Events
          </div>
          <div
            onClick={() => {
              setDataToRender("groups_invited");
              openModal();
            }}
            className="hover:text-blue-600 text-gray-700 flex items-center gap-1 my-2 px-4"
          >
            <EnvelopeOpenIcon className="w-5 h-5" />
            Group Invitation
          </div>
        </AccordionItem>
      </Accordion>
    </>
  );
};

export default NotificationAccordion;
