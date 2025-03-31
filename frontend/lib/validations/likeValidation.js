import * as Yup from "yup";
export const LikeSchema = Yup.object({
  action: Yup.string("this field must be string").required(
    "action is required"
  ),
  entries_id: Yup.string("this field must be string").required(
    "entries is required"
  ),
});
