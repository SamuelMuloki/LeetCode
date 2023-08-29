/**
 * https://leetcode.com/problems/function-composition/
 * @param {Function[]} functions
 * @return {Function}
 */
var compose = function (functions) {
  return function (x) {
    if (functions.length == 0) {
      return x;
    }

    x = functions[functions.length - 1].call(this, x);
    functions.pop();
    return compose(functions)(x);
  };
};

/**
 * const fn = compose([x => x + 1, x => 2 * x])
 * fn(4) // 9
 */

module.exports = { compose };
