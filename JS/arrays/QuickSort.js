//In Divide, first pick a pivot element. After that, partition or rearrange the array into two sub-arrays such that each element in the left sub-array is less than or equal to the pivot element and each element in the right sub-array is larger than the pivot element. Finally, recursively sort the sub-arrays.

// QuickSort is a divide and conquer algorithm. It works by selecting a 'pivot' element from the array and partitioning the other elements into two sub-arrays according to whether they are less than or greater than the pivot. The sub-arrays are then sorted recursively. This can be done in-place, requiring small additional amounts of memory to perform the sorting.

let array = [8, 3, 1, 7, 0, 10, 2];

function QuickSort(arr) {
  //check if the array is empty
  if (arr.length === 0) {
    return arr;
  }
  let left = [],
    right = [],
    pivot = arr[arr.length - 1];
  //loop on the and array
  for (let i = 0; i < arr.length - 1; i++) {
    let no = arr[i];
    //if the nno is less then pivot push to left array
    if (no < pivot) left.push(no);
    //if the nno is greter then pivot push to left array
    if (no > pivot) right.push(no);
  }

  return [...QuickSort(left), pivot, ...QuickSort(right)];
}

const sortedArray = QuickSort(array);

console.log("Unsorted Array:", array);
console.log("Sorted Array:", sortedArray);
