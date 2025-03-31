import { create } from "zustand";
export const useChatStore = create((set) => ({
  chat: {},
  setChat: (target) => {
    set({ chat: target });
  },
}));
export const useGroupFeedStore = create((set) => ({
  group: {},
  setGroup: (target) => {
    set({ group: target });
  },
}));
export const useBelongedGroupStore = create((set) => ({
  belongedGroups: undefined,
  setBelongedGroups: (target) => {
    set({ belongedGroups: target });
  },
}));

export const usePostStore = create((set) => ({
  post: {},
  setPost: (target) => {
    set({ post: target });
  },
}));
export const useGroupPostStore = create((set) => ({
  groupPost: {},
  setGroupPost: (target) => {
    set({ groupPost: target });
  },
}));

export const useWebsocketStore = create((set, get) => ({
  websocketService: null,
  setWebsocket: (target) => {
    set({ websockeService: target });
  },
}));
export const useNotificationStore = create((set) => ({
  notifications: null,
  setNotification: (target) => {
    set({ notifications: target });
  },
}));
export const useFriendsStore = create((set) => ({
  friends: null,
  setFriends: (target) => {
    set({ friends: target });
  },
}));

export const useWebsocketResponseStore = create((set) => ({
  websocketResponse: null,
  setWebsocketResponse: (target) => {
    set({ websocketResponse: target });
  },
}));
