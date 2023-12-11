/**
 * @param {number[]} nums
 * @return {number}
 */
var minMoves2 = function (nums) {
  nums.sort((a, b) => a - b);
  var mid = Math.floor(nums.length / 2);
  return nums
    .slice(0, mid)
    .reduce(
      (acc, curr, i) =>
        acc + (nums[mid] - curr) + (nums[nums.length - i - 1] - nums[mid]),
      0
    );
};

module.exports = { minMoves2 };
