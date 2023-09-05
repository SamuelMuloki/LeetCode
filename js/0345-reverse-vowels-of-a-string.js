/**
 * @param {string} s
 * @return {string}
 */
var reverseVowels = function (s) {
  var i = 0,
    j = s.length - 1;

  var isVowel = function (c) {
    return (
      c == 'a' ||
      c == 'e' ||
      c == 'i' ||
      c == 'o' ||
      c == 'u' ||
      c == 'A' ||
      c == 'E' ||
      c == 'I' ||
      c == 'O' ||
      c == 'U'
    );
  };

  s = s.split("");
  while (i < j) {
    if (!isVowel(s[i])) {
      i++;
      continue;
    }

    if (!isVowel(s[j])) {
      j--;
      continue;
    }

    var temp = s[i];
    s[i] = s[j];
    s[j] = temp;

    i++;
    j--;
  }

  return s.join("");
};

module.exports = { reverseVowels };
