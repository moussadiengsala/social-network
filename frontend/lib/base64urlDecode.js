export default function base64urlDecode(base64url) {
  // Convert base64url to base64 by replacing characters and padding
  const base64 = base64url.replace(/-/g, "+").replace(/_/g, "/");
  const padded = base64 + "=".repeat((4 - (base64.length % 4)) % 4);
  // Decode base64 string
  return atob(padded);
}
