/**
 * @param {number[]} nums
 * @return {number[]}
 */
var majorityElement = function (nums) {
  var candidate1 = 0,
    candidate2 = 1,
    count1 = 0,
    count2 = 0;

  for (var num of nums) {
    if (num == candidate1) {
      count1++;
    } else if (num == candidate2) {
      count2++;
    } else if (count1 == 0) {
      candidate1 = num;
      count1 = 1;
    } else if (count2 == 0) {
      candidate2 = num;
      count2 = 1;
    } else {
      count1--;
      count2--;
    }
  }

  var res = [];
  if (nums.filter((num) => num === candidate1).length > nums.length / 3) {
    res.push(candidate1);
  }
  if (nums.filter((num) => num === candidate2).length > nums.length / 3) {
    res.push(candidate2);
  }

  return res;
};

module.exports = { majorityElement };
