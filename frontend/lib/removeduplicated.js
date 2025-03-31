export function removeDuplicates(arr, prop) {
  const firstOccurrences = {};
  return arr.filter((obj, index) => {
    const key = obj[prop];
    if (key in firstOccurrences) {
      return false; // Already encountered, so filter it out
    } else {
      firstOccurrences[key] = index; // Record the index of the first occurrence
      return true;
    }
  });
}
