/**
 * @param {number[]} nums
 * @return {number}
 */
var numIdenticalPairs = function (nums) {
  var set = {},
    count = 0;
  nums.forEach((num) => (set[num] = (set[num] || 0) + 1));
  Object.values(set).forEach(
    (val) => (count += Math.floor((val * (val - 1)) / 2))
  );
  return count;
};

module.exports = { numIdenticalPairs };
