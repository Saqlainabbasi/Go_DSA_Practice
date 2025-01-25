//Merge sort is the sorting technique that follows the divide and conquer approach.Merge sort is similar to the quick sort algorithm as it uses the divide and conquer approach to sort the elements. It is one of the most popular and efficient sorting algorithm. It divides the given list into two equal halves, calls itself for the two halves and then merges the two sorted halves

//The sub-lists are divided again and again into halves until the list cannot be divided further. Then we combine the pair of one element lists into two-element lists, sorting them in the process. The sorted two-element pairs is merged into the four-element lists, and so on until we get the sorted list.

function MergeSort(array) {
  if (array.length <= 1) return array; //array is already sorted...

  //find the mid
  let mid = Math.floor(array.length / 2);
  let left = array.slice(0, mid);
  let right = array.slice(mid);

  return merge(MergeSort(left), MergeSort(right));
}

function merge(left, right) {
  let sortedArray = [];
  let i = 0,
    j = 0;

  while (i < left.length && j < right.length) {
    if (left[i] < right[j]) {
      sortedArray.push(left[i]);
      i++;
    } else {
      sortedArray.push(right[j]);
      j++;
    }
  }

  while (i < left.length) {
    sortedArray.push(left[i]);
    i++;
  }
  while (j < right.length) {
    sortedArray.push(right[j]);
    j++;
  }
  return sortedArray;
}
console.log(MergeSort([38, 27, 43, 3, 9, 82, 10]));
