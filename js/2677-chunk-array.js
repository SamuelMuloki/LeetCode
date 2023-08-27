/**
 * @param {Array} arr
 * @param {number} size
 * @return {Array[]}
 */
var chunk = function (arr, size) {
  let output = [];
  for (let i = 0, len = arr.length; i < len; )
    output.push(arr.slice(i, (i += size)));

  return output;
};

module.exports = { chunk };
