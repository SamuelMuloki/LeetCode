/**
 * @param {number[]} nums
 * @param {Function} fn
 * @param {number} init
 * @return {number}
 */
var reduce = function (nums, fn, init) {
  for (var i = 0; i < nums.length; i++) {
    init = fn.call(this, init, nums[i]);
  }

  return init;
};
module.exports = { reduce };
