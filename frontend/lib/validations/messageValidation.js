import * as Yup from "yup";

const MessageSchema = Yup.object({
  content: Yup.string("this field must be string").required(
    "Cannot leave this blank"
  ),
});
const GroupJoinSchema = Yup.object({
  group_id: Yup.string("this field must be string").required(
    "Cannot leave this blank"
  ),
});
const FollowSchema = Yup.object({
  follower_id: Yup.string("this field must be string").required(
    "Cannot leave this blank"
  ),
});
const InvitationSchema = Yup.object({
  group_id: Yup.string("this field must be string").required(
    "Cannot leave this blank"
  ),
  user: Yup.array()
    .required("Select at least one user to invite")
    .min(1, "Select at least one user to invite"),
});

export { MessageSchema, GroupJoinSchema, FollowSchema, InvitationSchema };
