/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isSubsequence = function(s, t) {
    var j = s.length-1;
    for (var i = t.length-1; i >= 0 && j >= 0; i--) {
        if (t.charCodeAt(i) === s.charCodeAt(j)) {
            j--;
        }
    }

    return j < 0;
};

module.exports = { isSubsequence }