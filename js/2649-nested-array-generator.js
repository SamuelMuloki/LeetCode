/**
 * @param {Array} arr
 * @return {Generator}
 */
var inorderTraversal = function* (arr) {
  function flat(arr, n) {
    return arr.reduce((output, element) => {
      if (Array.isArray(element)) {
        output.push(...flat(element, n - 1));
      } else {
        output.push(element);
      }

      return output;
    }, []);
  }

  arr = flat(arr, arr.length);
  for (const num of arr) {
    yield num;
  }
};

/**
 * const gen = inorderTraversal([1, [2, 3]]);
 * gen.next().value; // 1
 * gen.next().value; // 2
 * gen.next().value; // 3
 */

module.exports = { inorderTraversal };
