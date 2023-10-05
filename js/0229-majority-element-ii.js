/**
 * @param {number[]} nums
 * @return {number[]}
 */
var majorityElement = function (nums) {
  var sorted = nums.sort((a, b) => a - b);

  var count = 1,
    seen = Number.MIN_SAFE_INTEGER,
    res = [];
  for (var i = 0; i < sorted.length; i++) {
    if (i + 1 < sorted.length && sorted[i] === sorted[i + 1]) {
      count++;
    } else {
      count = 1;
    }

    if (seen !== sorted[i] && Math.floor(sorted.length / 3) < count) {
      seen = sorted[i];
      res.push(sorted[i]);
    }
  }

  return res;
};

module.exports = { majorityElement };
