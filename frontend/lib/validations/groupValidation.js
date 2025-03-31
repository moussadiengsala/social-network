import * as Yup from "yup";

const GroupSchema = Yup.object({
  title: Yup.string("this title must be string")
    .max(100)
    .min(10)
    .required("the title is required"),
  description: Yup.string("this field must be string")
    .max(255)
    .min(10)
    .required("the description is required"),
});
export { GroupSchema };
