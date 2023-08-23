/**
 * https://leetcode.com/problems/to-be-or-not-to-be
 * @param {string} val
 * @return {Object}
 */
var expect = function (val) {
  return {
    toBe: function (val2) {
      if (val !== val2) throw new Error("Not Equal");
      return true;
    },
    notToBe: function (val3) {
      if (val === val3) throw new Error("Equal");
      return true;
    },
  };
};

/**
 * expect(5).toBe(5); // true
 * expect(5).notToBe(5); // throws "Equal"
 */

module.exports = { expect };
