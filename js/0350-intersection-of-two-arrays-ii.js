/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 */
var intersect = function (nums1, nums2) {
  var minArr = nums1,
    maxArr = nums2;

  if (nums1.length > nums2.length) {
    (maxArr = nums1), (minArr = nums2);
  }

  return minArr.filter((val) => {
    var res = maxArr.includes(val);

    if (res) {
      maxArr.splice(maxArr.indexOf(val), 1);
    }

    return res;
  });
};

module.exports = { intersect };
