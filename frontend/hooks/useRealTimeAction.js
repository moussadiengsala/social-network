import {
  useBelongedGroupStore,
  useFriendsStore,
  useNotificationStore,
  useWebsocketResponseStore,
} from "@/_store/zustand";

const useRealTimeAction = () => {
  const websocketResponse = useWebsocketResponseStore(
    (state) => state.websocketResponse
  );
  const setFriends = useFriendsStore((state) => state.setFriends);
  const setNotification = useNotificationStore(
    (state) => state.setNotification
  );

  const setBelongedGroups = useBelongedGroupStore(
    (state) => state.setBelongedGroups
  );

  const update = () => {
    console.log("hello");
    if (websocketResponse) {
      if (websocketResponse?.payload?.data?.users) {
        setFriends(websocketResponse?.payload?.data?.users);
      }
      if (websocketResponse?.payload?.data?.groups) {
        setBelongedGroups(websocketResponse?.payload?.data?.groups);
      }
      if (
        websocketResponse?.payload?.data?.unread_notifications ||
        websocketResponse?.payload?.data?.unread_notifications === 0
      ) {
        setNotification(websocketResponse?.payload?.data?.unread_notifications);
      }
    }
  };
  return update;
};

export default useRealTimeAction;
