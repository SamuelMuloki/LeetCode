/**
 * @param {Array} arr1
 * @param {Array} arr2
 * @return {Array}
 */
var join = function (arr1, arr2) {
  var joinedArr = [...arr1, ...arr2];
  joinedArr.sort((a, b) => a.id - b.id);

  for (i = 0; i < joinedArr.length - 1; i++) {
    if (joinedArr[i].id === joinedArr[i + 1].id) {
      joinedArr[i + 1] = { ...joinedArr[i], ...joinedArr[i + 1] };
    }
  }

  return joinedArr.filter(
    (val, i, arr) =>
      (arr[i + 1] && val.id !== arr[i + 1].id) || (arr[i] && !arr[i + 1])
  );
};
module.exports = { join };
