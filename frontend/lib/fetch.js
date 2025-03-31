export const fetcher = async (options) => {
  const response = await fetch(
    `http://localhost:8000/api/${options.endpoint}`,
    {
      method: options.method || "GET",
      headers: {
        Authorization: options.authorizationHeader || "",
      },
      body:
        options.format === "--formData"
          ? options.data
          : JSON.stringify(options.data),
    }
  );

  return response.json();
};
