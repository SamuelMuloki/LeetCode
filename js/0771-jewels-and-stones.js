/**
 * @param {string} jewels
 * @param {string} stones
 * @return {number}
 */
var numJewelsInStones = function (jewels, stones) {
  const jSet = new Set(jewels);
  return stones.split("").reduce((acc, curr) => acc + jSet.has(curr), 0);
};

module.exports = { numJewelsInStones };
