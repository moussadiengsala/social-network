import { useGroupFeedStore } from "@/_store/zustand";
import useApiRequest from "@/hooks/useApiRequest";
import { GroupJoinSchema } from "@/lib/validations/messageValidation";
import { UsersIcon } from "@heroicons/react/24/outline";
import Image from "next/image";
import CustomInput from "../inputs/input";
import Alert from "../shared/alert";

function GroupeTeaser({ group, additionalAttributes, section = "page" }) {
  const setGroup = useGroupFeedStore((state) => state.setGroup);

  const OPTIONS = {
    method: "POST",
    endpoint: "groups-join",
    dataToRefresh: "groups-suggestions",
  };
  const { responseError, register, handleSubmit, errors, onSubmit } =
    useApiRequest(OPTIONS, GroupJoinSchema);
  return (
    <div className="flex flex-wrap items-center justify-between">
      <div
        onClick={() => {
          if (group.is_part_of_group) setGroup(group);
        }}
        key={group.email}
        {...additionalAttributes}
        className={`flex items-center  ${
          !group.is_part_of_group ? "opacity-55" : ""
        }  cursor-pointer gap-x-4 py-4`}
      >
        {group.cover ? (
          <Image
            width={100}
            height={100}
            className="h-12 w-12 shadow-md flex-none object-cover rounded-full bg-gray-50"
            src={`http://localhost:8000/api/static/uploads/${group.cover}`}
            alt="group cover"
          />
        ) : (
          <UsersIcon className="w-10 h-10 text bg-gray-200 text-gray-500 rounded-full shadow-md border-2 border-gray-100" />
        )}
        <div className="min-w-0">
          <p className="text-sm font-semibold leading-6 text-gray-900">
            {group.title}
          </p>
          {section === "page" && (
            <p className="mt-1 truncate bg-blue-100 w-fit  text-blue-600 px-2 py-1 rounded-md text-xs leading-5 ">
              {group.member_count} member(s)
            </p>
          )}
        </div>
      </div>
      <div className="flex gap-2 items-center">
        {group.status === "requested"}
        {!group.is_part_of_group && (
          <form onSubmit={handleSubmit(onSubmit)}>
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
              value="requested"
              register={register}
              name="status"
              className="col-span-6 sm:col-span-3 "
            />
            <button
              disabled={
                group.status === "rejected" ||
                group.status === "requested" ||
                group.status === "invited"
              }
              className="bg-blue-600 disabled:bg-yellow-200 disabled:text-yellow-600  text-sm  text-white rounded-md px-2 py-1"
            >
              Join group
              {group.status === "requested" && " requested sent"}
              {group.status === "invited" && " invited"}
              {group.status === "rejected" && " rejected"}
            </button>
          </form>
        )}
      </div>
      {responseError.length !== 0 && (
        <Alert message={responseError} status="error" />
      )}
    </div>
  );
}

export default GroupeTeaser;
