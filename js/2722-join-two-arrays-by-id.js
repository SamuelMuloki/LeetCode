/**
 * @param {Array} arr1
 * @param {Array} arr2
 * @return {Array}
 */
var join = function (arr1, arr2) {
  arr1.sort((a, b) => a.id - b.id);
  arr2.sort((a, b) => a.id - b.id);
  let i = 0,
    j = 0;

  const joinedArr = [];
  while (i < arr1.length && j < arr2.length) {
    if (arr1[i].id === arr2[j].id) {
      joinedArr.push({ ...arr1[i], ...arr2[j] });
      i++;
      j++;
      continue;
    }

    if (arr[i].id < arr[2].id) {
      joinedArr.push({ ...arr1[i] });
      i++;
      continue;
    }

    joinedArr.push({ ...arr2[j] });
    j++;
  }

  while (i < arr1.length) {
    joinedArr.push({ ...arr1[i] });
    i++;
  }

  while (j < arr2.length) {
    joinedArr.push({ ...arr2[j] });
    j++;
  }

  return joinedArr;
};
module.exports = { join };
