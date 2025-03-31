import * as Yup from "yup";

const signinSchema = Yup.object({
  identifiers: Yup.string("this field must be string").required(
    "Cannot leave this blank"
  ),
  password: Yup.string("this field must be string").required(
    "Password is required"
  ),
});
export { signinSchema };
