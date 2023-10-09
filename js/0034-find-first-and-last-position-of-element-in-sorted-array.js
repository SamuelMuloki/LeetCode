/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var searchRange = function (nums, target) {
  var l = 0,
    r = nums.length - 1;
  while (l <= r) {
    var mid = l + Math.floor((r - l) / 2);
    if (nums[mid] < target) {
      l = mid + 1;
    } else if (nums[mid] > target) {
      r = mid - 1;
    } else {
      var first = (last = mid);
      while (first > 0 && nums[first - 1] === target) {
        first--;
      }

      while (last < nums.length - 1 && nums[last + 1] === target) {
        last++;
      }

      return [first, last];
    }
  }

  return [-1, -1];
};

module.exports = { searchRange };
