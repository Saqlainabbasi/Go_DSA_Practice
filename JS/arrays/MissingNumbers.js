// Probblem
// Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.

function FindAllMissingNumbers(numbers) {
  let no_obj = {},
    result = [];
  for (const no of numbers) {
    no_obj[no] = (no_obj[no] || 0) + 1;
  }
  console.log(no_obj);
  for (let i = 1; i <= numbers.length; i++) {
    if (!no_obj[i]) {
      result.push(i);
    }
  }
  return result;
}
console.log(FindAllMissingNumbers([1, 1]));
