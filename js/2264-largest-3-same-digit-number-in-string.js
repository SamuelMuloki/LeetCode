/**
 * @param {string} num
 * @return {string}
 */
var largestGoodInteger = function (num) {
  var [first, second, ans] = [num[0], num[1], -1];
  [...num].slice(2).forEach((val) => {
    if (val === first && val === second) {
      ans = Math.max(ans, val);
    }
    [first, second] = [second, val];
  });
  return ans === -1 ? "" : (ans + "").repeat(3);
};

module.exports = { largestGoodInteger };
