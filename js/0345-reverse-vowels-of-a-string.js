/**
 * @param {string} s
 * @return {string}
 */
var reverseVowels = function (s) {
  var i = 0,
    j = s.length - 1;

  var isVowel = function (c) {
    return (
      "a" == c ||
      "e" == c ||
      "i" == c ||
      "o" == c ||
      "u" == c ||
      "A" == c ||
      "E" == c ||
      "I" == c ||
      "O" == c ||
      "U" == c
    );
  };

  while (i < j) {
    if (!isVowel(s[i])) {
      i++;
      continue;
    }

    if (!isVowel(s[j])) {
      j--;
      continue;
    }

    (sI = s[i]),
      (sJ = s[j]),
      (s = s.setCharAt(s, i, sJ)),
      (s = s.setCharAt(s, j, sI));
    i++;
    j--;
  }

  return s;
};

String.prototype.setCharAt = function (str, i, c) {
  if (i > str.length - 1) return str;
  return str.substring(0, i) + c + str.substring(i + 1);
};

module.exports = { reverseVowels };
