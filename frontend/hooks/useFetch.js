import Cookies from "js-cookie";

const useFetch = () => {
  const fetcher = async (options) => {
    const jwt = async () => Cookies.get("social-network-jwt");

        const response = await fetch(`${"http://localhost:8000"}/api/${options.endpoint}`, {
      method: options.method || "GET",

      headers: {
        Authorization: await jwt(),
      },
      body:
        options.format === "--formData"
          ? options.data
          : JSON.stringify(options.data),
    });

    if (!response.ok) {
      const errorResponse = await response.json(); // await here
      const message = errorResponse.message; // Access message property

      throw new Error(`${message}`);
    }

    return response.json();
  };

  return fetcher;
};

export default useFetch;
