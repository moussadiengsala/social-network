import useOpenModal from "@/hooks/useOpenModal";
import { JWT as JWTService } from "@/lib/jwt";

import { useFriendsStore, useWebsocketStore } from "@/_store/zustand";
import { Dialog, Transition } from "@headlessui/react";
import {
  Bars3Icon,
  ChatBubbleLeftRightIcon,
  FaceFrownIcon,
  SparklesIcon,
  UserIcon,
  UsersIcon,
  XMarkIcon,
} from "@heroicons/react/24/outline";
import { deleteCookie } from "cookies-next";
import Cookies from "js-cookie";
import Image from "next/image";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { Fragment, useEffect, useState } from "react";
import NotificationsModal from "../modals/notifications_modal";
import NotificationAccordion from "../shared/notification_accordion";
import AsideProfile from "../shared/skeletons/asideProfile";
import { useWebSocket } from "@/context/realTimeContext";
function classNames(...classes) {
  return classes.filter(Boolean).join(" ");
}

export default function Aside() {
  const { friends, belongedGroups, websocketService } = useWebSocket();
  const totalUnreadGroupMessages = belongedGroups.reduce(
    (accumulator, currentValue) => {
      return accumulator + currentValue.unread_message;
    },
    0
  );
  const [data, setData] = useState();
  const totalUnreadMessages = friends?.reduce((total, friend) => {
    return total + friend?.unread_message;
  }, 0);
  const router = useRouter();

  useEffect(() => {
    setData(JWTService.decoder(Cookies.get("social-network-jwt")));
  }, []);

  const navigation = [
    { name: "Feed", href: "/app/home", icon: SparklesIcon, current: true },
    {
      name: "Profile",
      href: `/app/profile/${data?.payload?.user?.id}`,
      icon: UserIcon,
      current: false,
    },
    {
      name: "Messages",
      href: "/app/messages",
      icon: ChatBubbleLeftRightIcon,
      current: false,
    },

    { name: "Groups", href: "/app/groups", icon: UsersIcon, current: false },
  ];
  const [sidebarOpen, setSidebarOpen] = useState(false);

  const { closeModal, ModalComponent } = useOpenModal();
  return (
    <>
      <ModalComponent>
        <NotificationsModal close={closeModal} />
      </ModalComponent>
      <div>
        <Transition.Root show={sidebarOpen} as={Fragment}>
          <Dialog
            as="div"
            className="relative z-50 lg:hidden"
            onClose={setSidebarOpen}
          >
            <Transition.Child
              as={Fragment}
              enter="transition-opacity ease-linear duration-300"
              enterFrom="opacity-0"
              enterTo="opacity-100"
              leave="transition-opacity ease-linear duration-300"
              leaveFrom="opacity-100"
              leaveTo="opacity-0"
            >
              <div className="fixed inset-0 bg-gray-900/80" />
            </Transition.Child>

            <div className="fixed inset-0 flex">
              <Transition.Child
                as={Fragment}
                enter="transition ease-in-out duration-300 transform"
                enterFrom="-translate-x-full"
                enterTo="translate-x-0"
                leave="transition ease-in-out duration-300 transform"
                leaveFrom="translate-x-0"
                leaveTo="-translate-x-full"
              >
                <Dialog.Panel className="relative mr-16 flex w-full max-w-xs flex-1">
                  <Transition.Child
                    as={Fragment}
                    enter="ease-in-out duration-300"
                    enterFrom="opacity-0"
                    enterTo="opacity-100"
                    leave="ease-in-out duration-300"
                    leaveFrom="opacity-100"
                    leaveTo="opacity-0"
                  >
                    <div className="absolute left-full top-0 flex w-16 justify-center pt-5">
                      <button
                        type="button"
                        className="-m-2.5 p-2.5"
                        onClick={() => setSidebarOpen(false)}
                      >
                        <span className="sr-only">Close sidebar</span>
                        <XMarkIcon
                          className="h-6 w-6 text-white"
                          aria-hidden="true"
                        />
                      </button>
                    </div>
                  </Transition.Child>
                  {/* Sidebar component, swap this element with another sidebar if you like */}
                  <div className="flex grow flex-col gap-y-5 overflow-y-auto bg-white px-6 pb-2">
                    <div className="flex h-16 shrink-0 items-center">
                      <Image
                        width={100}
                        height={100}
                        className="h-8 w-auto"
                        src="/image3.png"
                        alt="Your Company"
                      />
                    </div>
                    <nav className="flex flex-1 flex-col">
                      <ul role="list" className="flex flex-1 flex-col gap-y-7">
                        <li>
                          <ul role="list" className="-mx-2 space-y-1">
                            {navigation.map((item, i) => (
                              <li key={i}>
                                <Link
                                  href={item.href}
                                  className={classNames(
                                    item.current
                                      ? "bg-gray-50 text-blue-600"
                                      : "hover:text-blue-600 text-gray-700  hover:bg-gray-50",
                                    "group flex gap-x-3 rounded-md p-2 text-sm leading-6 font-semibold"
                                  )}
                                >
                                  <item.icon
                                    className={classNames(
                                      item.current
                                        ? "text-blue-600"
                                        : "text-gray-400 group-hover:text-blue-600",
                                      "h-6 w-6 shrink-0"
                                    )}
                                    aria-hidden="true"
                                  />
                                  {item.name}
                                </Link>
                              </li>
                            ))}
                          </ul>
                        </li>
                      </ul>
                    </nav>
                  </div>
                </Dialog.Panel>
              </Transition.Child>
            </div>
          </Dialog>
        </Transition.Root>

        {/* Static sidebar for desktop */}
        <div className="hidden  lg:flex lg:w-72 h-full lg:flex-col">
          {/* Sidebar component, swap this element with another sidebar if you like */}
          <div className="flex grow flex-col gap-y-5 overflow-y-auto border-r border-gray-200 bg-white px-6">
            <div className="flex h-16 shrink-0 items-center">
              <Image
                width={100}
                height={100}
                className="h-8 w-auto"
                src="/image3.png"
                alt="Your Company"
              />
            </div>
            <nav className="flex flex-1 flex-col">
              <ul role="list" className="flex flex-1 flex-col gap-y-7">
                <li>
                  <ul role="list" className="-mx-2 space-y-1">
                    {navigation.map((item, i) => (
                      <li key={i}>
                        <Link
                          href={item.href}
                          className={classNames(
                            item.current
                              ? "bg-gray-50 text-blue-600"
                              : ` ${"hover:text-blue-600 text-gray-700  hover:bg-gray-50"} `,
                            "group flex gap-x-3 rounded-md p-2 text-sm leading-6 font-semibold"
                          )}
                        >
                          <div className="relative">
                            {item.name === "Messages" &&
                              totalUnreadMessages > 0 && (
                                <div className="bg-pink-600 text-[8px] absolute top-0 right-0 text-white p-1 text-center flex items-center justify-center rounded-full h-3 w-3">
                                  <p>{totalUnreadMessages}</p>
                                </div>
                              )}

                            {item.name === "Groups" &&
                              totalUnreadGroupMessages > 0 && (
                                <div className="bg-pink-600 text-[8px] absolute top-0 right-0 text-white p-1 text-center flex items-center justify-center rounded-full h-3 w-3">
                                  <p>{totalUnreadGroupMessages}</p>
                                </div>
                              )}
                            <item.icon
                              className={classNames(
                                item.current
                                  ? "text-blue-600"
                                  : `text-gray-400 ${"group-hover:text-blue-600 text-gray-400"}`,
                                "h-6 w-6 shrink-0"
                              )}
                              aria-hidden="true"
                            />
                          </div>
                          {item.name}
                        </Link>
                      </li>
                    ))}
                    <li className="group cursor-pointer  text-gray-700 flex gap-x-3 rounded-md  text-sm leading-6 font-semibold">
                      <NotificationAccordion />
                    </li>
                    <li
                      onClick={() => {
                        deleteCookie("social-network-jwt");
                        router.push("/login");
                        websocketService.close();
                      }}
                      className="group cursor-pointer bg-pink-100 flex gap-x-3 rounded-md p-2 text-sm leading-6 font-semibold text-pink-500 hover:text-pink-600  hover:bg-pink-200"
                    >
                      <FaceFrownIcon className="h-6 w-6 shrink-0" />
                      Log out
                    </li>
                  </ul>
                </li>

                <li className="-mx-6 mt-auto">
                  {data?.payload ? (
                    <div className="flex items-center gap-x-4 px-6 py-3 text-sm font-semibold leading-6 text-gray-900 hover:bg-gray-50">
                      {data?.payload?.user && (
                        <Image
                          width={100}
                          height={100}
                          className="h-8 w-8 rounded-full object-cover bg-gray-50"
                          src={
                            data?.payload?.user?.avatar.length != 0
                              ? `http://localhost:8000/api/static/uploads/${data?.payload?.user?.avatar}`
                              : "/unknown.jpg"
                          }
                          alt=""
                        />
                      )}

                      <span className="sr-only">Your profile</span>
                      <div className="flex flex-col">
                        <span aria-hidden="true">
                          {`${data?.payload?.user.first_name} ${data?.payload?.user.last_name}`}
                        </span>
                        <span className="font-light " aria-hidden="true">
                          @{data?.payload?.user.username}
                        </span>
                      </div>
                    </div>
                  ) : (
                    <AsideProfile />
                  )}
                </li>
              </ul>
            </nav>
          </div>
        </div>

        <div className="sticky top-0 z-40 flex items-center gap-x-6 bg-white px-4 py-4 shadow-sm sm:px-6 lg:hidden">
          <button
            type="button"
            className="-m-2.5 p-2.5 text-gray-700 lg:hidden"
            onClick={() => setSidebarOpen(true)}
          >
            <span className="sr-only">Open sidebar</span>
            <Bars3Icon className="h-6 w-6" aria-hidden="true" />
          </button>

          <span className="sr-only">Your profile</span>
          {data && (
            <Image
              width={100}
              height={100}
              className="h-8 w-8 rounded-full object-cover bg-gray-50"
              src={
                data?.payload?.user?.avatar != ""
                  ? `http://localhost:8000/api/static/uploads/${data?.payload?.use?.avatar}`
                  : "/unknown.jpg"
              }
              alt=""
            />
          )}
        </div>
      </div>
    </>
  );
}
