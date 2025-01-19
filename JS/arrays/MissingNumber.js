// ===> Problem
// Given an array nums containing n distinct numbers in the range [0, n],
// return the only number in the range that is missing from the array.

// we will the math formula to calculate the total sum of the array
// then calculte the sum by looping onn the array
// at the end we will subtract the formula sum fom the cal sum
function MissingNumber(array) {
  let formula_sum = 0;
  let cal_sum = 0;
  len = array.length;
  formula_sum = (len * (len + 1)) / 2;

  //loop to cal the sum
  for (const element of array) {
    cal_sum += element;
  }

  return formula_sum - cal_sum;
}
console.log(MissingNumber([9, 6, 4, 2, 3, 5, 7, 0, 1]));
