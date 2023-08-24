/**
 * https://leetcode.com/problems/promise-time-limit
 * @param {Function} fn
 * @param {number} t
 * @return {Function}
 */
var timeLimit = function (fn, t) {
  return async function (...args) {
    return Promise.race([
      new Promise((_, reject) => setTimeout(reject, t, "Time Limit Exceeded")),
      fn.apply(null, args),
    ]);
  };
};

/**
 * const limited = timeLimit((t) => new Promise(res => setTimeout(res, t)), 100);
 * limited(150).catch(console.log) // "Time Limit Exceeded" at t=100ms
 */

module.exports = { timeLimit };
