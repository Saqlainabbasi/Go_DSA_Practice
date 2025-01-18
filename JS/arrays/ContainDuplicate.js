// ==> Problem
// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

// we will use the set to check for the duplicate
function CheckDuplicate(array) {
  let no_set = new Set();
  //loop on array
  for (const no of array) {
    if (no_set.has(no)) return true;
    no_set.add(no);
  }
  return false;
}

console.log(CheckDuplicate([1, 2, 3, 1.5]));
