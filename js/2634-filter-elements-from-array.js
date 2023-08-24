/**
 * @param {number[]} arr
 * @param {Function} fn
 * @return {number[]}
 */
var filter = function (arr, fn) {
  var filteredArr = [];
  for (let i = 0; i < arr.length; i++) {
    if (fn.call(this, arr[i], i)) {
      filteredArr.push(arr[i]);
    }
  }
  return filteredArr;
};
module.exports = { filter };
