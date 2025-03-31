"use client";
import useApiRequest from "@/hooks/useApiRequest";
import useFilePreview from "@/hooks/useFilePreview";
import { signupSchema } from "@/lib/validations/signupValidation";
import { ShieldExclamationIcon } from "@heroicons/react/24/outline";
import {
  AtSymbolIcon,
  CakeIcon,
  EnvelopeIcon,
  KeyIcon,
  UserCircleIcon,
} from "@heroicons/react/24/solid";
import Link from "next/link";
import { IoMaleFemaleOutline } from "react-icons/io5";
import CustomInput from "../inputs/input";
import CustomSelect from "../inputs/select";
import TextArea from "../inputs/textArea";
import Alert from "../shared/alert";
import Image from "next/image";

function RegisterForm() {
  const OPTIONS = {
    method: "POST",
    endpoint: "auth/signup",
    data: {},
    format: "--formData",
  };
  const {
    responseError,
    register,
    handleSubmit,
    watch,
    errors,
    response,
    isLoading,
    onSubmit,
  } = useApiRequest(OPTIONS, signupSchema);
  const file = watch("avatar");
  const [filePreview] = useFilePreview(file);

  return (
    <form
      method="POST"
      enctype="multipart/form-data"
      className="mt-8 bg-white shadow-lg border-2 px-6 py-10 rounded-md border-gray-100  grid grid-cols-6 gap-6"
      onSubmit={handleSubmit(onSubmit)}
    >
      <div className="col-span-6 space-y-2">
        <h1 className="text-2xl font-bold">
          Welcome Zone01 moment sharing platform
        </h1>
        <p className="text-sm text-gray-600">
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Sed, vero
          placeat! Officiis ad voluptates nihil eius earum, libero voluptas
          fugit!
        </p>
      </div>

      {/* Avatar */}
      <div className="col-span-6 sm:col-span-3">
        <div className="mt-2 flex items-center gap-x-3">
          {filePreview ? (
            <Image
              width={100}
              height={100}
              className="w-10 rounded-full h-10  object-cover"
              src={filePreview}
              alt="preview"
            />
          ) : (
            <UserCircleIcon className="h-12 w-12 text-gray-300" />
          )}

          <CustomInput
            register={register}
            type="file"
            label="avatar"
            placeholder=""
            name="avatar"
            icon={<UserCircleIcon className="w-5 h-5 text-gray-400" />}
            className="col-span-6 sm:col-span-3 hidden"
          />
          <label
            htmlFor="avatar"
            className="rounded-md bg-white px-2.5 py-1.5 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50"
          >
            Choose
          </label>
        </div>
      </div>
      <CustomSelect
        register={register}
        label="Gender"
        options={["male", "female"]}
        error={errors.gender}
        name="gender"
        icon={<IoMaleFemaleOutline className="w-5 h-5 text-gray-400" />}
        additionalAttributes={{ className: "col-span-6 sm:col-span-3 " }}
      />
      <CustomInput
        type="text"
        register={register}
        label="First Name"
        placeholder="Enter your first name"
        error={errors.first_name}
        name="first_name"
        icon={<UserCircleIcon className="w-5 h-5 text-gray-400" />}
        className="col-span-6 sm:col-span-3 "
      />
      {/* LN */}
      <CustomInput
        register={register}
        type="text"
        label="Last Name"
        error={errors.last_name}
        placeholder="Enter your last name"
        name="last_name"
        icon={<UserCircleIcon className="w-5 h-5 text-gray-400" />}
        className="col-span-6 sm:col-span-3  "
      />
      {/* Date of Birth */}
      <CustomInput
        register={register}
        type="date"
        label="Date of Birth"
        error={errors.date_of_birth}
        placeholder="Enter your date of birth"
        name="date_of_birth"
        icon={<CakeIcon className="w-5 h-5 text-gray-400" />}
        className="col-span-6 sm:col-span-3  "
      />
      {/* Username */}
      <CustomInput
        register={register}
        type="text"
        label="Nickname/Username"
        placeholder="Enter your username"
        error={errors.username}
        name="username"
        icon={<AtSymbolIcon className="w-5 h-5 text-gray-400" />}
        className="col-span-6 sm:col-span-3  "
      />
      {/* Email */}
      <CustomInput
        register={register}
        type="email"
        label="Email"
        placeholder="Enter your email"
        error={errors.email}
        name="email"
        icon={<EnvelopeIcon className="w-5 h-5 text-gray-400" />}
        className="col-span-6"
      />
      {/* Email */}
      {/* bio */}
      <div className="col-span-6">
        <label
          htmlFor="bio"
          className="block text-sm font-medium text-gray-700"
        >
          {" "}
          Bio{" "}
        </label>

        <TextArea
          register={register}
          name="bio"
          id="bio"
          rows={4}
          placeholder="Tell us about you..."
        />
      </div>
      <CustomInput
        register={register}
        type="password"
        label="Password"
        placeholder="Enter your password"
        name="password"
        error={errors.password}
        icon={<KeyIcon className="w-4 h-4 text-gray-400" />}
        className="col-span-6 sm:col-span-3"
      />
      <CustomSelect
        register={register}
        label="privacy"
        options={["Choose your account type", "public", "private"]}
        error={errors.gender}
        name="privacy"
        icon={<ShieldExclamationIcon className="w-5 h-5 text-gray-400" />}
        additionalAttributes={{ className: "col-span-6 sm:col-span-3 " }}
      />
      <div className="col-span-6">
        <p className="text-sm text-gray-500">
          By creating an account, you agree to our terms and conditions and
          privacy policy .
        </p>
        <div className="col-span-6">
          {responseError.length !== 0 && (
            <Alert message={responseError} status="error" />
          )}
          {responseError.length === 0 && response?.message === "ok" && (
            <Alert
              message="Your account has been successfully created. We're thrilled to have you join our community!"
              status="success"
            />
          )}
        </div>
      </div>
      <div className="col-span-6 sm:flex sm:items-center sm:gap-4">
        <button
          disabled={isLoading}
          className="inline-block disabled:opacity-20 shrink-0 rounded-md border border-blue-600 bg-blue-600 px-12 py-3 text-sm font-medium text-white transition hover:bg-transparent hover:text-blue-600 focus:outline-none focus:ring active:text-blue-500"
        >
          Create an account
        </button>
        <p className="mt-4 text-sm text-gray-500 sm:mt-0">
          Already have an account?{" "}
          <Link href="/login" className="hover:underline hover:text-blue-600">
            Log in
          </Link>{" "}
          .
        </p>
      </div>
    </form>
  );
}

export default RegisterForm;
