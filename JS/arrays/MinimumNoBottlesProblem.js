// problem a water bottle company have 4 different type of bottles 10,7,2,1 liter. They want to complete an order of n liter using the minimum no of bottles

//  solution ===> o solve the problem of finding the minimum number of bottles required to fulfill an order, you can use a greedy algorithm. The algorithm selects the largest bottle size available that doesn't exceed the remaining required quantity and continues until the order is complete.

let bottle_type = [10, 5, 2, 1];

function Min_Bottles(order) {
  //sort the bbottle type inn decending order if its nnot,

  let count = 0;
  let remaining_order = order;
  // loop on the bottles..
  for (const bottle of bottle_type) {
    if (remaining_order === 0) {
      break;
    }
    // if (remaining_order < bottle) {
    //   continue;
    // }
    // as the answer is in points we are using the floor function to round it of
    let bottles_used = Math.floor(remaining_order / bottle);
    console.log(bottles_used, bottle, remaining_order);
    count += bottles_used;
    remaining_order -= bottles_used * bottle;
  }
  if (remaining_order > 0) {
    return -1; //order cannot be completed
  }
  return count;
}

console.log(Min_Bottles(37));
