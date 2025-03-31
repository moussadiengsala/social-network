import * as Yup from "yup";

const signupSchema = Yup.object({
  email: Yup.string("this field must be string").required("email is required"),
  first_name: Yup.string("this field must be string").required(
    "first name is required"
  ),
  last_name: Yup.string("this field must be string").required(
    "last name is required"
  ),
  date_of_birth: Yup.date()
    .max(new Date("2024-01-01"), "You do not have the required age")
    .min(new Date("1970-01-01"), "You are too old for this "),
  username: Yup.string("this field must be string").required(
    "username is required"
  ),

  gender: Yup.string("this field must be string").required(
    "gender is required"
  ),
  privacy: Yup.string("this field must be string").required(
    "gender is required"
  ),
  password: Yup.string("this field must be string").required(
    "Password is required"
  ),
});
export { signupSchema };
