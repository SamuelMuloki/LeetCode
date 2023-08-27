/**
 * @param {Array} arr
 * @param {number} size
 * @return {Array[]}
 */
var chunk = function (arr, size) {
  let output = new Array(),
    chunkArr = [];
  for (let i = 0; i < arr.length; i++) {
    chunkArr.push(arr[i]);
    if (chunkArr.length === size) {
      output.push(chunkArr);
      chunkArr = [];
    }
  }

  if (chunkArr.length) {
    output.push(chunkArr)
  }

  return output;
};

module.exports = { chunk };
