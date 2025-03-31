import { jwtVerify } from "jose";
import base64urlDecode from "./base64urlDecode";

export const JWT = {
  decoder: (token) => {
    if (!token) {
      return {
        header: null,
        payload: null,
        signature: null,
      };
    }
    const [header, payload, signature] = token?.split(".");

    // Decode header and payload
    const decodedHeader = JSON.parse(base64urlDecode(header));
    const decodedPayload = JSON.parse(base64urlDecode(payload));

    return {
      header: decodedHeader,
      payload: decodedPayload,
      signature: signature,
    };
  },
};
