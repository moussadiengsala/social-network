import * as Yup from "yup";

Yup.addMethod(Yup.string, "stripEmptyString", function () {
  return this.transform((value) => (value === "" ? undefined : value));
});
const postSchema = Yup.object({
  privacy: Yup.string("this field must be string")
    .stripEmptyString()
    .default("public"),
  content: Yup.string("this field must be string")
    .max(255)
    .min(10)
    .required("the content is required"),
});
export { postSchema };
