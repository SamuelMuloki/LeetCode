/**
 * @param {any[]} arr
 * @param {number} depth
 * @return {any[]}
 */
var flat = function (arr, n) {
  if (n === 0) {
    return arr;
  }

  return arr.reduce((output, element) => {
    if (Array.isArray(element)) {
      output.push(...flat(element, n - 1));
    } else {
      output.push(element);
    }

    return output;
  }, []);
};
module.exports = { flat };
