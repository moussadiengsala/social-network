import { useGroupFeedStore } from "@/_store/zustand";
import useOpenModal from "@/hooks/useOpenModal";
import { formatDate } from "@/lib/TimeFormatter";
import {
  CalendarDaysIcon,
  ChatBubbleLeftRightIcon,
  CreditCardIcon,
} from "@heroicons/react/24/outline";
import GroupInviteModal from "../modals/group_invite_modal";
import EmptyState from "../shared/emptyState";
import Tabs from "../shared/tabs";
import GroupEventFeed from "./groups/groupEventFeed";
import GroupMessageFeed from "./groups/groupMessages";
import GroupPostFeed from "./groups/groupPostFeed";
import { useQuery } from "react-query";
import useFetch from "@/hooks/useFetch";
import Link from "next/link";

function GroupFeed() {
  const fetcher = useFetch();
  const group = useGroupFeedStore((state) => state.group);
  const { openModal, closeModal, ModalComponent } = useOpenModal();
  // const OPTIONS = {
  //   method: "GET",
  //   endpoint: `groups`,
  // };
  // const {
  //   isLoading,
  //   refetch,
  //   data: userProfile,
  // } = useQuery(["group", OPTIONS], () => fetcher(OPTIONS));
  return (
    <div className="w-[330px] hidden lg:block border-s">
      <ModalComponent>
        <GroupInviteModal close={closeModal} />
      </ModalComponent>
      {Object.keys(group).length !== 0 ? (
        <section>
          <div
            style={{
              backgroundImage: `url(http://localhost:8000/api/static/uploads/${group.cover})`,
            }}
            className={`border-gray-300 bg-center text-white relative bg-cover `}
          >
            <div className="absolute inset-0 bg-gray-900/75 sm:bg-transparent sm:from-gray-900/95 sm:to-gray-700/25 sm:bg-gradient-to-r "></div>
            <div className="z-50 relative px-3 py-4">
              <h1 className="text-xl  font-bold">{group.title}</h1>
              <h1 className="text-sm">{group.description}</h1>
              <div className="flex justify-between items-center">
                <h1 className="text-sm">
                  Created since{" "}
                  <span className="font-bold">
                    {formatDate(group.created_at)}
                  </span>{" "}
                </h1>
                <small className="">{group.member_count} member(s)</small>
                <p
                  onClick={openModal}
                  className="hover:underline cursor-pointer"
                >
                  Invite
                </p>
              </div>
            </div>
          </div>
          <div className="px-3 w-full text-white text-sm  bottom-0 py-2 bg-gray-900/95">
            {" "}
            <Link
              className="hover:underline"
              href={`/app/group-chats/${group.id}`}
            >
              Go to chat
            </Link>
          </div>

          <Tabs
            tabs={[
              {
                name: "Posts",
                href: "#",
                icon: CreditCardIcon,
                current: true,
                componentToRender: <GroupPostFeed />,
              },

              {
                name: "Events",
                href: "#",
                icon: CalendarDaysIcon,
                current: false,
                componentToRender: <GroupEventFeed />,
              },
            ]}
          />
        </section>
      ) : (
        <EmptyState
          message="Click to a group to start a conversation, create event or make a post or just keep
       swimming."
        />
      )}
    </div>
  );
}
export default GroupFeed;
