/**
 * https://leetcode.com/problems/apply-transform-over-each-element-in-array/
 * @param {number[]} arr
 * @param {Function} fn
 * @return {number[]}
 */
var map = function (arr, fn) {
  for (let i = 0; i < arr.length; i++) {
    arr[i] = fn.call(this, arr[i], i);
  }
  return arr;
};

module.exports = { map };
