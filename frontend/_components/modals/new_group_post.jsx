import { useGroupFeedStore } from "@/_store/zustand";
import useApiRequest from "@/hooks/useApiRequest";
import useFilePreview from "@/hooks/useFilePreview";
import { postSchema } from "@/lib/validations/postValidation";
import { Dialog, Transition } from "@headlessui/react";
import { PhotoIcon } from "@heroicons/react/24/solid";
import Image from "next/image";
import { Fragment, useEffect } from "react";
import CustomInput from "../inputs/input";
import TextArea from "../inputs/textArea";
import Alert from "../shared/alert";

export default function GroupPostModal({ close }) {
  const OPTIONS = {
    method: "POST",
    endpoint: "groups-posts",
    data: {},
    format: "--formData",
    dataToRefresh: "groups-posts",
  };
  const group = useGroupFeedStore((state) => state.group);

  const {
    responseError,
    register,
    handleSubmit,
    watch,
    errors,
    isSuccess,
    onSubmit,
  } = useApiRequest(OPTIONS, postSchema);
  const file = watch("file");
  const [filePreview] = useFilePreview(file);

  useEffect(() => {
    if (isSuccess) {
      close();
    }
  }, [isSuccess, close]);

  useEffect(() => {}, [responseError]);

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
          <div className="flex  items-end justify-center p-4 text-center sm:items-center sm:p-0">
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
                  <h1 className="text-xl font-bold">What's on your mind ?</h1>
                  <form
                    onSubmit={handleSubmit(onSubmit)}
                    className="space-y-4"
                    action=""
                    method="POST"
                    encType=""
                  >
                    <CustomInput
                      type="hidden"
                      value={group.id}
                      register={register}
                      name="group_id"
                      className="col-span-6 sm:col-span-3 "
                    />
                    <TextArea
                      register={register}
                      error={errors.content}
                      additionalAttributes={{
                        className:
                          "w-full mt-2 w-full rounded-lg border-gray-100  p-3 border-2 align-top shadow-sm sm:text-sm",
                      }}
                      rows={4}
                      name="content"
                      placeholder="Express yourself....😀"
                    />
                    <div className="flex items-center text-blue-600 gap-2">
                      <label
                        htmlFor="file"
                        className="rounded-md bg-blue-200 my-2 px-2.5 py-1.5 text-sm font-semibold shadow-sm ring-1 ring-inset ring-blue-300 hover:bg-blue-50"
                      >
                        <PhotoIcon className="w-5 h-5" />
                      </label>
                      <small>Only png, jpg, gif are supported</small>
                    </div>
                    {filePreview ? (
                      <Image
                        width={100}
                        height={100}
                        className="rounded-lg h-[150px] w-full object-cover"
                        src={filePreview}
                        alt="preview"
                      />
                    ) : null}
                    <CustomInput
                      type="file"
                      register={register}
                      label="photo"
                      placeholder=""
                      name="file"
                      className="hidden"
                    />

                    <div className="grid grid-cols-2 gap-2">
                      <button className="inline-block w-full   shrink-0 rounded-md border border-blue-600 bg-blue-600 px-12 py-3 text-sm font-medium text-white transition hover:bg-blue-700 focus:outline-none focus:ring active:text-blue-500">
                        Publish
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
                  {responseError.length !== 0 && (
                    <Alert message={responseError} status="error" />
                  )}
                </div>
              </Dialog.Panel>
            </Transition.Child>
          </div>
        </div>
      </Dialog>
    </Transition.Root>
  );
}
