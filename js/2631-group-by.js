/**
 * @param {Function} fn
 * @return {Array}
 */
Array.prototype.groupBy = function (fn) {
  var output = {};
  this.forEach((element) => {
    var res = fn(element);

    if (res in output) {
      var curr = output[res];
      curr.push(element);
    } else {
      output[res] = [element];
    }
  });

  return output;
};

/**
 * [1,2,3].groupBy(String) // {"1":[1],"2":[2],"3":[3]}
 */
