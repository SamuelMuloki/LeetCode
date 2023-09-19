/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function (x) {
  let reversed = 0,
    temp = x;
  while (temp > 0) {
    reversed = reversed * 10 + (temp % 10);
    temp = Math.floor(temp / 10);
  }

  return reversed == x;
};
module.exports = { isPalindrome }