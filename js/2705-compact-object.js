/**
 * @param {Object} obj
 * @return {Object}
 */
var compactObject = function (obj) {
  function dfs(obj) {
    if (!obj) return false;
    if (typeof obj !== "object") return obj;

    if (Array.isArray(obj)) {
      const newArr = [];
      for (let i = 0; i < obj.length; i++) {
        const subRes = dfs(obj[i]);

        if (subRes) {
          newArr.push(subRes);
        }
      }

      return newArr;
    }

    const newObj = {};
    for (const key in obj) {
      const subRes = dfs(obj[key]);

      if (subRes) {
        newObj[key] = subRes;
      }
    }

    return newObj;
  }

  return dfs(obj);
};
module.exports = { compactObject };