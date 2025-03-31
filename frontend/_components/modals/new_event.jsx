import {
  useGroupFeedStore,
  useWebsocketResponseStore,
  useWebsocketStore,
} from "@/_store/zustand";
import { queryClient } from "@/app/layout";
import useApiRequest from "@/hooks/useApiRequest";
import { eventSchema } from "@/lib/validations/eventValidation";
import { Dialog, Transition } from "@headlessui/react";
import { Fragment, useEffect } from "react";
import CustomInput from "../inputs/input";
import TextArea from "../inputs/textArea";
import Alert from "../shared/alert";

function NewEventModal({ close }) {
  const websocketResponse = useWebsocketResponseStore(
    (state) => state.websocketResponse
  );
  const group = useGroupFeedStore((state) => state.group);

  // useEffect(() => {
  //   if (
  //     websocketResponse.action === "groups-events" &&
  //     websocketResponse.payload.message === "ok"
  //   ) {
  //     close();
  //   }
  // }, [websocketResponse]);
  const OPTIONS = {
    method: "POST",
    endpoint: "groups-events",
    format: "--formData",
    dataToRefresh: "groups-events",
  };

  const { responseError, register, handleSubmit, errors, onSubmit } =
    useApiRequest(OPTIONS, eventSchema);

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
              <Dialog.Panel className="relative transform overflow-hidden rounded-lg bg-white px-4 pb-4 pt-5 text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-sm sm:p-6">
                <div className="space-y-3">
                  <h1 className="text-xl font-bold">New Event</h1>
                  <form
                    onSubmit={handleSubmit(onSubmit)}
                    className="space-y-4"
                    action=""
                    method="POST"
                  >
                    <CustomInput
                      type="hidden"
                      value={group.id}
                      register={register}
                      name="group_id"
                      className="col-span-6 sm:col-span-3 "
                    />
                    <CustomInput
                      type="text"
                      register={register}
                      label="title"
                      error={errors.title}
                      placeholder="Enter the event title"
                      name="title"
                    />
                    <TextArea
                      register={register}
                      error={errors.description}
                      additionalAttributes={{
                        className:
                          "w-full mt-2 w-full rounded-lg border-gray-100  p-3 border-2 align-top shadow-sm sm:text-sm",
                      }}
                      rows={4}
                      name="description"
                      placeholder="Event description"
                    />
                    <CustomInput
                      register={register}
                      type="datetime-local"
                      label="Date of the event"
                      error={errors.datetime}
                      name="datetime"
                      className="col-span-6 sm:col-span-3  "
                    />
                    {responseError.length !== 0 && (
                      <Alert message={responseError} status="error" />
                    )}
                    <div className="grid grid-cols-2 gap-2">
                      <button className="inline-block w-full disabled:opacity-25   shrink-0 rounded-md border border-blue-600 bg-blue-600 px-12 py-3 text-sm font-medium text-white transition hover:bg-blue-700 focus:outline-none focus:ring active:text-blue-500">
                        Create
                      </button>
                      <button
                        type="button"
                        onClick={close}
                        className="inline-block w-full shrink-0 rounded-md border border-pink-600 bg-transparent px-12 py-3 text-sm font-medium text-pink-600 transition hover:bg-pink-600 hover:text-white focus:outline-none focus:ring active:text-pink-500"
                      >
                        Cancel
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

export default NewEventModal;
