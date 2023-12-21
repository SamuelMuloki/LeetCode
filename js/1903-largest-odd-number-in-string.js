/**
 * @param {string} num
 * @return {string}
 */
var largestOddNumber = function (num) {
  for (var i = num.length - 1; i >= 0; i--) {
    if (num[i] % 2 === 1) return num.substring(0, i + 1);
  }
  return "";
};

module.exports = { largestOddNumber };
