export function formatDate(dateString) {
  const date = new Date(dateString);
  const options = { month: "long", year: "numeric" };
  return date.toLocaleDateString("en-US", options);
}
export function formatEventDate(dateString) {
  const date = new Date(dateString);

  // Format month
  const monthOptions = { month: "long" };
  const month = date.toLocaleDateString("en-US", monthOptions); // Outputs something like "December"

  // Format day
  const dayOptions = { day: "numeric" };
  const day = date.toLocaleDateString("en-US", dayOptions); // Outputs something like "12"

  // Format time
  const timeOptions = { hour: "numeric", minute: "numeric", hour12: true };
  const time = date.toLocaleTimeString("en-US", timeOptions); // Outputs something like "12:00 PM"

  return { month, day, time };
}
