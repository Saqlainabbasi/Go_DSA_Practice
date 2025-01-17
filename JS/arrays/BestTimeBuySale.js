//  Problem
// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

// prices = [7,1,5,3,6,4];

// will use the 2 pointer approch
function MaxProfit(prices) {
  let low = Infinity,
    max_profit = 0;

  for (let i = 0; i < prices.length; i++) {
    const price = prices[i];
    if (price < low) {
      low = price;
    }
    let profit = price - low;
    if (profit > max_profit) {
      max_profit = profit;
    }
  }
  return max_profit;
}
console.log(MaxProfit([7, 1, 5, 3, 6, 4]));

function Max_Profit(prices) {
  let left = 0,
    right = 0,
    max_profit = 0;

  while (left < prices.length) {
    if (prices[left] < prices[right]) {
      max_profit = Math.max(max_profit, prices[right] - prices[left]);
    } else {
      left = right;
    }
    right++;
  }
  return max_profit;
}
console.log("===>", Max_Profit([7, 1, 5, 3, 6, 4]));
