/**
 * @param {string} columnTitle
 * @return {number}
 */
var titleToNumber = function (columnTitle) {
  var num = 0,
    j = columnTitle.length - 1;
  for (var i = 0; i < columnTitle.length; i++) {
    num +=
      (columnTitle.charCodeAt(i) - "A".charCodeAt(0) + 1) * Math.pow(26, j);
    j--;
  }

  return num;
};

module.exports = { titleToNumber };
