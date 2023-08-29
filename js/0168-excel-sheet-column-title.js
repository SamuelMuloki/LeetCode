/**
 * https://leetcode.com/problems/excel-sheet-column-title/
 * @param {number} columnNumber
 * @return {string}
 */
var convertToTitle = function (columnNumber) {
  if (!columnNumber) {
    return "";
  }

  columnNumber--;
  var charCode = "A".charCodeAt(0) + (columnNumber % 26);

  return (
    convertToTitle(Math.floor(columnNumber / 26)) +
    String.fromCharCode(charCode)
  );
};

module.exports = { convertToTitle };
