import { yupResolver } from "@hookform/resolvers/yup";
import { useForm } from "react-hook-form";
import { useState } from "react";
import { useMutation } from "react-query";
import useFetch from "./useFetch";
import { queryClient } from "@/app/layout";

const useApiRequest = (options, validationSchema, defaultValues) => {
  const fetcher = useFetch();
  const [responseError, setResponseError] = useState([]);
  const [response, setResponse] = useState();
  const {
    register,
    handleSubmit,
    watch,
    reset,
    formState: { errors },
  } = useForm({
    resolver: yupResolver(validationSchema),
    defaultValues,
  });
  const { mutate, isLoading, isSuccess } = useMutation(fetcher, {
    onError: (data) => setResponseError(data.message.split("*")),
    onSuccess: (data) => {
      reset();
      setResponse(data);
      setResponseError([]);
      // queryClient.refetchQueries([options.dataToRefresh], { active: true });
      queryClient.invalidateQueries({ queryKey: [options.dataToRefresh] });
    },
  });

  const onSubmit = (data, e) => {
    options.data =
      options.format === "--formData" ? new FormData(e.target) : data;

    for (const key in options.data) {
      if (key === "id" || key.includes("_id")) {
        options.data[key] = parseInt(options.data[key]);
      }
    }
    if (options.format === "--formData") {
      if (options.data?.getAll("selected_users")) {
        options.data.set(
          "selected_users",
          options.data.getAll("selected_users").join(",")
        );
      }
    }
    if (options.format != "--formData") {
      console.log(options.data);
      if (options.data.user) options.data.user = options.data.user.join(",");
    }
    console.log(options.data);
    mutate(options);
  };

  return {
    responseError,
    register,
    handleSubmit,
    watch,
    errors,
    isLoading,
    isSuccess,
    response,
    onSubmit,
  };
};

export default useApiRequest;
