import * as Yup from "yup";
export const CommentSchema = Yup.object({
  content: Yup.string("this field must be string")
    .max(255)
    .min(10)
    .required("the content is required"),
});
