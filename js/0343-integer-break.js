/**
 * @param {number} n
 * @return {number}
 */
var integerBreak = function (n) {
  if (n <= 3) {
    return n - 1;
  }

  var dp = new Array(n + 1);
  var dfs = function (num) {
    if (num <= 3) {
      return num;
    }

    if (dp[num]) {
      return dp[num];
    }

    var res = num;
    for (var i = 2; i < num; i++) {
      res = Math.max(res, i * dfs(num - i));
    }

    return (dp[num] = res);
  };

  return dfs(n);
};

module.exports = { integerBreak };
