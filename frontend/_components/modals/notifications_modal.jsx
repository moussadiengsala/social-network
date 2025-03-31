import { Dialog, Transition } from "@headlessui/react";
import { Fragment } from "react";
import GroupsInvitationTab from "../section/tabs/eventTabs";
import FollowRequestTab from "../section/tabs/notifications/follow_request";
import GroupJoinRequestTab from "../section/tabs/notifications/groupJoinRequestTab";
import GroupEventTab from "../section/tabs/notifications/group_event_tab";
function NotificationsModal({ close, dataToRender }) {
  const tabs = {
    groups_invited: <GroupsInvitationTab />,
    groups_events: <GroupEventTab />,
    follow_request: <FollowRequestTab />,
    groups_requested: <GroupJoinRequestTab />,
  };
  const componentsToRender = tabs[dataToRender];
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
                  {componentsToRender}
                  <button
                    type="button"
                    onClick={close}
                    className="inline-block w-full shrink-0 rounded-md border border-pink-600 bg-transparent px-12 py-3 text-sm font-medium text-pink-600 transition hover:bg-pink-600 hover:text-white focus:outline-none focus:ring active:text-pink-500"
                  >
                    Close
                  </button>
                </div>
              </Dialog.Panel>
            </Transition.Child>
          </div>
        </div>
      </Dialog>
    </Transition.Root>
  );
}

export default NotificationsModal;
