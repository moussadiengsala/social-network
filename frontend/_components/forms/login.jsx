"use client";
import useApiRequest from "@/hooks/useApiRequest";
import { signinSchema } from "@/lib/validations/signinValidation";
import { AtSymbolIcon } from "@heroicons/react/24/outline";
import { EyeIcon } from "@heroicons/react/24/solid";

import Link from "next/link";
import { redirect } from "next/navigation";
import { useEffect } from "react";
import CustomInput from "../inputs/input";
import Alert from "../shared/alert";

import Cookies from "js-cookie";

const SignInForm = () => {
  const OPTIONS = {
    method: "POST",
    endpoint: "auth/signin",
    data: {},
  };
  const {
    responseError,
    register,
    handleSubmit,
    errors,
    isLoading,
    response,
    onSubmit,
  } = useApiRequest(OPTIONS, signinSchema);
  useEffect(() => {
    if (response?.data) {
      Cookies.set("social-network-jwt", response.data.jwt);
      redirect("/app/home");
    }
  }, [response]);
  return (
    <div className="mx-auto max-w-screen-xl px-4 py-16 sm:px-6 lg:px-8">
      <div className="mx-auto max-w-lg text-center">
        <h1 className="text-2xl font-bold sm:text-3xl">Get started today!</h1>
        <p className="mt-4 text-gray-500">
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Et libero
          nulla eaque error neque ipsa culpa autem, at itaque nostrum!
        </p>
      </div>
      <form
        action="POST"
        onSubmit={handleSubmit(onSubmit)}
        className="mx-auto mb-0 mt-8 max-w-md space-y-4"
      >
        <CustomInput
          name="identifiers"
          type="username"
          error={errors.identifiers}
          register={register}
          label="Email Adress"
          placeholder="Enter email or username/nickname"
          icon={<AtSymbolIcon className="w-5 h-5 text-gray-500" />}
        />
        <CustomInput
          name="password"
          type="password"
          error={errors.password}
          register={register}
          label="Password"
          placeholder="Enter password"
          icon={<EyeIcon />}
        />
        {responseError.length !== 0 && (
          <Alert message={responseError} status="error" />
        )}
        <div className="flex items-center justify-between">
          <p className="text-sm text-gray-500">
            No account?{" "}
            <Link className="underline hover:text-blue-500" href="/">
              Sign up
            </Link>
          </p>

          <button
            disabled={isLoading}
            type="submit"
            className="inline-block rounded-lg disabled:opacity-20 bg-blue-500 px-5 py-3 text-sm font-medium text-white"
          >
            Sign in
          </button>
        </div>
      </form>
    </div>
  );
};

export default SignInForm;
