/**
 * @param {number} n
 * @return {number}
 */
var minimumOneBitOperations = function (n) {
  var ans = n;
  ans ^= ans >> 16;
  ans ^= ans >> 8;
  ans ^= ans >> 4;
  ans ^= ans >> 2;
  ans ^= ans >> 1;

  return ans;
};

module.exports = { minimumOneBitOperations };
