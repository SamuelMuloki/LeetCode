/**
 * https://leetcode.com/problems/valid-anagram/
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isAnagram = function (s, t) {
  var sCount = new Map();
  if (s.length != t.length) {
    return false;
  }

  for (const char of s) {
    const plusCount = (sCount.get(char) || 0) + 1;

    sCount.set(char, plusCount);
  }

  for (const char of t) {
    if (!sCount.has(char)) continue;
    const subCount = sCount.get(char) - 1;

    sCount.set(char, subCount);
  }

  for (const [char, count] of sCount) {
    if (count !== 0) return false;
  }

  return true;
};

module.exports = { isAnagram };
