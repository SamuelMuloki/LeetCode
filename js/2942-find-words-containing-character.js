/**
 * @param {string[]} words
 * @param {character} x
 * @return {number[]}
 */
var findWordsContaining = function (words, x) {
  return words.reduce(
    (acc, curr, i) => (curr.indexOf(x) > -1 ? [...acc, i] : acc),
    []
  );
};

module.exports = { findWordsContaining };
