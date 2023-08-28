/**
 * @param {any[]} arr
 * @param {number} depth
 * @return {any[]}
 */
var flat = function (arr, n) {
  while (n > 0) {
    for (i = 0; i < arr.length; ) {
      if (Array.isArray(arr[i])) {
        var iLen = arr[i].length;
        arr.splice(i, 1, ...arr[i]);
        i += iLen;
      } else {
        i++;
      }
    }
    n--;
  }

  return arr;
};
module.exports = { flat };
