/**
 * @param {string} s
 * @return {string}
 */
var reverseVowels = function (s) {
  var vowel = s.match(/[aeiou]/gi);
  return s.replace(/[aeiou]/gi, (_) => vowel.pop());
};

module.exports = { reverseVowels };
