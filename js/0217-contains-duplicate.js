/**
 * https://leetcode.com/problems/contains-duplicate/
 * @param {number[]} nums
 * @return {boolean}
 */
var containsDuplicate = function (nums) {
  var numsSet = new Set(nums);
  const isEqual = numsSet.size === nums.length;

  return !isEqual;
};

module.exports = { containsDuplicate };
