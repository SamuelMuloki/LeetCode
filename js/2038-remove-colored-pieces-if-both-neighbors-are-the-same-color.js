/**
 * @param {string} colors
 * @return {boolean}
 */
var winnerOfGame = function (colors) {
  var aCount = 0,
    bCount = 0;
  for (var i = 1; i < colors.length - 1; i++) {
    if (
      colors.charCodeAt(i - 1) === colors.charCodeAt(i) &&
      colors.charCodeAt(i) === colors.charCodeAt(i + 1)
    ) {
      if (colors.charCodeAt(i) === "A".charCodeAt(0)) aCount++;
      else bCount++;
    }
  }

  return aCount > bCount;
};

module.exports = { winnerOfGame };
