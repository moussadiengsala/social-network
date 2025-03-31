"use client";
import CustomInput from "@/_components/inputs/input";
import FollowerModal from "@/_components/modals/followers_modal";
import ProfileFeed from "@/_components/section/profileFeed";
import AccountStatTag from "@/_components/shared/accountsStatTag";
import Alert from "@/_components/shared/alert";
import Tabs from "@/_components/shared/tabs";
import useApiRequest from "@/hooks/useApiRequest";
import useFetch from "@/hooks/useFetch";
import useOpenModal from "@/hooks/useOpenModal";
import { FollowSchema } from "@/lib/validations/messageValidation";
import {
  ArrowPathIcon,
  CakeIcon,
  CreditCardIcon,
} from "@heroicons/react/24/outline";

import { queryClient } from "@/app/layout";
import { JWT as JWTService } from "@/lib/jwt";
import {
  ChatBubbleOvalLeftEllipsisIcon,
  HeartIcon,
  UserGroupIcon,
} from "@heroicons/react/24/solid";
import Cookies from "js-cookie";
import Image from "next/image";
import { useState } from "react";
import { useQuery } from "react-query";
import { formatDate, formatEventDate } from "@/lib/TimeFormatter";

function Profile({ params }) {
  const data = JWTService.decoder(Cookies.get("social-network-jwt"));

  const { openModal, closeModal, ModalComponent } = useOpenModal();

  const fetcher = useFetch();

  const OPTIONS = {
    method: "GET",
    endpoint: `users/${params.id}`,
  };
  const [entry, setEntry] = useState();
  const {
    isLoading,
    refetch,
    data: userProfile,
  } = useQuery(["profile", OPTIONS], () => fetcher(OPTIONS));

  const followingStatus = {
    "not-followed": "bg-blue-100 px-3 py-[.7]  text-blue-600 hover:bg-blue-200",
    accept: "hover:bg-pink-200 bg-pink-100 px-3 py-[.7]  text-pink-600",
    rejected: "",
    pending: "hover:bg-yellow-200 bg-yellow-100 px-3 py-[.7]  text-yellow-600",
  };

  const { responseError, register, handleSubmit, onSubmit } = useApiRequest(
    {
      method: "POST",
      endpoint: `follow?follower_id=${params.id}`,
      dataToRefresh: "profile",
    },
    FollowSchema
  );
  const options = {
    method: "Post",
    endpoint: `change-privacy-status`,
  };
  const { refetch: profileSwitchingRefetch, isLoading: profileSwitchLoading } =
    useQuery(["change-status", options], () => fetcher(options), {
      onSuccess: () => {
        queryClient.refetchQueries(["profile"], { active: true });
      },
      enabled: false,
    });
  console.log(userProfile);
  return (
    <main className="">
      <ModalComponent>
        <FollowerModal entry={entry} id={params.id} close={closeModal} />
      </ModalComponent>
      {isLoading ? (
        <div className="relative">
          <div className="bg-gray-300 animate-pulse w-full h-[250px]"></div>
          <div className="absolute break-all w-full py-6 space-y-4 border-b-[1.2px]  -bottom-64 lg:-bottom-52  px-4 sm:px-6 lg:px-8">
            <div
              className="w-20 h-20 lg:w-28 lg:h-28 bg-gray-200 shadow-xl object-cover rounded-full"
              alt="profile Image"
            />
            <div className="space-y-2">
              <h1 className="w-[240px] h-3 bg-gray-300"></h1>
              <p className="w-[200px] h-3 bg-gray-300"></p>
              <div className="space-y-1">
                <p className="w-[440px] h-3 bg-gray-300"></p>
                <p className="w-[420px] h-3 bg-gray-300"></p>
                <p className="w-[410px] h-3 bg-gray-300"></p>
              </div>
            </div>
          </div>
        </div>
      ) : userProfile?.data ? (
        <>
          <div className="relative ">
            <div className="bg-gradient-to-r from-blue-500 to-blue-600 w-full h-[250px]"></div>
            <div className="absolute break-all w-full py-6 space-y-4 border-b-[1.2px]  -bottom-64   px-4 sm:px-6 lg:px-8">
              <Image
                width={100}
                height={100}
                src={
                  userProfile.data.avatar != ""
                    ? `http://localhost:8000/api/static/uploads/${userProfile.data.avatar}`
                    : "/unknown.jpg"
                }
                className="w-20 h-20 lg:w-28 lg:h-28 shadow-xl object-cover rounded-full"
                alt="profile Image"
              />
              <div className="space-y-1">
                <div>
                  <div className="flex items-center gap-4">
                    <h1 className="text-xl font-bold">{`${userProfile.data.first_name} ${userProfile.data.last_name}`}</h1>
                    {userProfile.data.id != data.payload?.user?.id && (
                      <form onSubmit={handleSubmit(onSubmit)}>
                        <CustomInput
                          type="hidden"
                          value={params.id}
                          register={register}
                          name="follower_id"
                          className="col-span-6 sm:col-span-3 "
                        />
                        {userProfile?.data.following_status !== "reject" ? (
                          <button
                            type="submit"
                            className={`rounded-md ${
                              followingStatus[
                                userProfile?.data.following_status
                              ]
                            } `}
                          >
                            {userProfile?.data.following_status ===
                              "not-followed" && "Follow"}
                            {userProfile?.data.following_status === "accept" &&
                              "Unfollow"}
                            {userProfile?.data.following_status === "pending" &&
                              "Request sent"}
                          </button>
                        ) : (
                          <div className="bg-pink-100 rounded-md text-pink-600 px-2 py-1">
                            don&quot;t want to be friend with you
                          </div>
                        )}
                      </form>
                    )}
                    {userProfile.data.id == data.payload?.user?.id && (
                      <button
                        onClick={() => profileSwitchingRefetch()}
                        disabled={profileSwitchLoading}
                        className={`bg-orange-200 text-orange-500 disabled:opacity-25 px-2 py-1 rounded-md cursor-pointer`}
                      >
                        Switch to{" "}
                        {userProfile?.data.privacy === "private"
                          ? "public"
                          : "private"}{" "}
                        profile
                      </button>
                    )}
                  </div>
                  <p className="text-gray-500">{userProfile.data.username}</p>
                </div>
                <p className="text-gray-500 text-sm lg:text-base">
                  {userProfile.data.bio}
                </p>
                <p className="text-orange-500 text-sm lg:text-base">
                  {userProfile.data.email}
                </p>
                <p className="text-gray-500 flex items-center gap-1 text-sm lg:text-base">
                  <CakeIcon className="w-5 h-5" />
                  Born
                  <span className=" ">{`
                  ${formatEventDate(userProfile.data.date_of_birth).month} ${
                    formatEventDate(userProfile.data.date_of_birth).day
                  } `}</span>
                </p>
                {responseError.length != 0 && (
                  <Alert message={responseError} status="error" />
                )}
                <div className="flex gap-2"></div>
              </div>

              <div className="flex flex-wrap items-center gap-2">
                <div
                  onClick={() => {
                    setEntry("followers");
                    openModal();
                  }}
                >
                  <AccountStatTag
                    stat={"follower"}
                    statValue={userProfile.data.followers}
                    icon={<UserGroupIcon className="w-4 h-4 " />}
                  />
                </div>
                <div
                  onClick={() => {
                    setEntry("followings");
                    openModal();
                  }}
                >
                  <AccountStatTag
                    stat={"following"}
                    statValue={userProfile.data.following}
                    icon={<UserGroupIcon className="w-4 h-4 " />}
                  />
                </div>
                {/* likes_count */}
                <AccountStatTag
                  stat={"like"}
                  statValue={userProfile.data.likes_count}
                  icon={<HeartIcon className="w-4 h-4 " />}
                />
                <AccountStatTag
                  stat={"post"}
                  statValue={userProfile.data.post_count}
                  icon={<UserGroupIcon className="w-4 h-4 " />}
                />
              </div>
            </div>
          </div>
          {/* Main content */}
          <section className="py-20 lg:mt-52 space-y-6 px-4 sm:px-6 lg:px-8">
            {userProfile.data.id != data.payload?.user?.id &&
            userProfile?.data.following_status != "accept" &&
            userProfile?.data.privacy == "private" ? (
              <div className="py-6">
                <h1 className="text-2xl font-bold">This account is private</h1>
                <p className="text-gray-700 mt-2">
                  Only confirmed followers have access to this account. Click
                  the &quot;Follow&quot; button to send a follow request
                </p>
              </div>
            ) : (
              <>
                <Tabs
                  tabs={[
                    {
                      name: "Owned Posts",
                      href: "#",
                      icon: CreditCardIcon,
                      current: true,
                      componentToRender: (
                        <ProfileFeed
                          id={params.id}
                          key="owned_posts"
                          section="owned_posts"
                        />
                      ),
                    },
                    {
                      name: "Liked Posts",
                      href: "#",
                      icon: HeartIcon,
                      current: false,
                      componentToRender: (
                        <ProfileFeed
                          id={params.id}
                          key="liked_posts"
                          section="liked_posts"
                        />
                      ),
                    },
                    {
                      name: "Comment Posts",
                      href: "#",
                      icon: ChatBubbleOvalLeftEllipsisIcon,
                      current: false,
                      componentToRender: (
                        <ProfileFeed
                          id={params.id}
                          key="commented_posts"
                          section="commented_posts"
                        />
                      ),
                    },
                  ]}
                />
              </>
            )}
          </section>
        </>
      ) : (
        <div className="flex flex-col items-center">
          <Alert
            message={`
                ${userProfile?.message || "User not found"} ERROR_CODE : ${
              userProfile?.code || 404
            }`}
            status="error"
          />
          <small
            className="text-pink-600 bg-pink-50 px-2 py-1 rounded-md flex items-center gap-1 text-center mx-auto cursor-pointer hover:underline"
            onClick={() => refetch()}
          >
            <ArrowPathIcon className="w-4 h-4" />
            Try again
          </small>
        </div>
      )}
    </main>
  );
}

export default Profile;
