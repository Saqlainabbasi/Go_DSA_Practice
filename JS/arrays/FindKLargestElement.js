// problem
// Given an integer array nums and an integer k, return the kth largest element in the array.unsorted array

// solution
// we will use hashmap||object inn js and keep track of the frequency elments
// we will find the min and max of the array in same loop where we check frequency
// we are find min and max so that we can form max to min no

function KthLargestElement(array, k) {
  if (array.length === 0 || k === null) return null;

  let freq_map = {},
    max = -Infinity,
    min = Infinity;

  for (const no of array) {
    freq_map[no] = (freq_map[no] || 0) + 1;
    max = Math.max(max, no);
    min = Math.min(min, no);
  }

  for (let i = max; i >= min; i--) {
    if (freq_map[i]) {
      k -= freq_map[i];
      if (k <= 0) return i;
    }
  }
}
console.log(KthLargestElement([33, 22, 33, 11, 22, 43, 55, 51, 63], 4));
