// problem
// Given an integer array nums and an integer k, return the kth smallest element in the array.un sorted array

// we will use the divide and conquer like quick sort
// solution when we have the unique unsorted array
function KthSmallestElement(array, k) {
  if (array.length === 0) return array;

  let left = [],
    right = [],
    pivot = array[array.length - 1];

  for (let no of array) {
    if (no < pivot) left.push(no); //smaller element to the left
    if (no > pivot) right.push(no); // larger to right
  }

  if (left.length === k - 1) return pivot; //pivot is the largest element
  if (left.length >= k) return KthSmallestElement(left, k);
  return KthSmallestElement(right, k - left.length - 1);
}
// solution when we have duplicate  in the unsoeted array
function _KthSmallestElement(array, k) {
  //   console.log(array, k);
  if (array.length === 0) return array;

  let left = [],
    right = [],
    equal = [],
    pivot = array[array.length - 1];

  for (let no of array) {
    if (no < pivot) left.push(no); //smaller element to the left
    if (no > pivot) right.push(no); // larger to the right
    if (no === pivot) equal.push(no); // equal to the equal array
  }
  if (k <= left.length) return _KthSmallestElement(left, k);
  if (k <= left.length + equal.lengt) return pivot;
  return _KthSmallestElement(right, k - left.length - equal.length);
}
console.log(_KthSmallestElement([33, 22, 33, 1, 51, 62, 41, 66], 4));
