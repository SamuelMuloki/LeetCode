/**
 * @param {string} s
 * @return {string}
 */
var removeStars = function (s) {
  var st = [];
  for (var i = 0; i < s.length; i++) {
    if (s.charCodeAt(i) === "*".charCodeAt(0)) {
      st.pop();
    } else {
      st.push(s[i]);
    }
  }

  return st.join("");
};

module.exports = { removeStars };
