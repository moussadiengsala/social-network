import { useFriendsStore, useGroupFeedStore } from "@/_store/zustand";
import { Fragment, useState } from "react";
import { Dialog, Transition } from "@headlessui/react";
import UserCard from "../cards/userMessageTeaser";
import CustomInput from "../inputs/input";
import { yupResolver } from "@hookform/resolvers/yup";
import { useForm } from "react-hook-form";
import { InvitationSchema } from "@/lib/validations/messageValidation";
import useApiRequest from "@/hooks/useApiRequest";
import Alert from "../shared/alert";
import { useWebSocket } from "@/context/realTimeContext";
function GroupInviteModal({ close }) {
  const { friends } = useWebSocket();

  const group = useGroupFeedStore((state) => state.group);

  const OPTIONS = {
    method: "POST",
    endpoint: "groups-join",
  };

  const { responseError, register, handleSubmit, isSuccess, errors, onSubmit } =
    useApiRequest(OPTIONS, InvitationSchema, { user: [] });

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
                    Invite friends to
                    <span className="italic"> {group.title}</span>
                  </h1>
                  <form onSubmit={handleSubmit(onSubmit)} action="">
                    <ul role="list" className="divide-y  divide-gray-100">
                      {!friends ? (
                        <div className="divide-y">
                          <div className="flex cursor-pointer gap-x-4 py-4">
                            <div className="h-12 w-12 flex-none rounded-full bg-gray-200 animate-pulse"></div>
                            <div className="min-w-0">
                              <div className="h-4 w-[100px] bg-gray-200 rounded mb-2 animate-pulse"></div>
                              <div className="h-4 w-[160px] bg-gray-200 rounded animate-pulse"></div>
                            </div>
                          </div>
                        </div>
                      ) : friends?.length != 0 ? (
                        friends.map((friend, index) => (
                          <div
                            key={index}
                            className="flex items-center justify-between overflow-y-auto"
                          >
                            <label htmlFor={`follower_${friend.id}`}>
                              <UserCard person={friend} />
                            </label>
                            <CustomInput
                              register={register}
                              type="checkbox"
                              placeholder=""
                              error={errors.user}
                              name={`user`}
                              value={friend.id}
                              id={`follower_${friend.id}`}
                              className="col-span-6 sm:col-span-3"
                            />
                          </div>
                        ))
                      ) : (
                        <div className="flex text-gray-400 text-sm py-4 flex-col  items-center">
                          You not have any friend yet
                        </div>
                      )}
                    </ul>
                    <CustomInput
                      type="hidden"
                      value={group.id}
                      register={register}
                      name="group_id"
                      a={{ valueAsNumber: true }}
                      className="col-span-6 sm:col-span-3 "
                    />
                    <CustomInput
                      type="hidden"
                      value="invited"
                      register={register}
                      name="status"
                      className="col-span-6 sm:col-span-3 "
                    />

                    {responseError.length !== 0 && (
                      <Alert message={responseError} status="error" />
                    )}
                    {isSuccess && (
                      <Alert
                        message="The group invitation was successfully sent to your friends"
                        status="success"
                      />
                    )}
                    <div className="grid grid-cols-2 gap-4">
                      <button
                        type="submit"
                        className="inline-block w-full   shrink-0 rounded-md border border-blue-600 bg-blue-600 px-12 py-3 text-sm font-medium text-white transition hover:bg-blue-700 focus:outline-none focus:ring active:text-blue-500"
                      >
                        Invite
                      </button>
                      <button
                        type="button"
                        onClick={close}
                        className="inline-block w-full shrink-0 rounded-md border border-pink-600 bg-transparent px-12 py-3 text-sm font-medium text-pink-600 transition hover:bg-pink-600 hover:text-white focus:outline-none focus:ring active:text-pink-500"
                      >
                        Close
                      </button>
                    </div>
                  </form>
                </div>
              </Dialog.Panel>
            </Transition.Child>
          </div>
        </div>
      </Dialog>
    </Transition.Root>
  );
}

export default GroupInviteModal;
