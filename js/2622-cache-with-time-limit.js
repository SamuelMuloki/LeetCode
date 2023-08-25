// https://leetcode.com/problems/cache-with-time-limit
var TimeLimitedCache = function () {
  this.cache = new Map();
  this.timeouts = new Map();
};

/**
 * @param {number} key
 * @param {number} value
 * @param {number} duration time until expiration in ms
 * @return {boolean} if un-expired key already existed
 */
TimeLimitedCache.prototype.set = function (key, value, duration) {
  var exists = this.cache.has(key);

  if (this.timeouts.has(key)) {
    clearTimeout(this.timeouts.get(key));
  }

  this.cache.set(key, value);
  this.timeouts.set(
    key,
    setTimeout(() => {
      this.cache.delete(key);
    }, duration)
  );

  return exists;
};

/**
 * @param {number} key
 * @return {number} value associated with key
 */
TimeLimitedCache.prototype.get = function (key) {
  return this.cache.get(key) || -1;
};

/**
 * @return {number} count of non-expired keys
 */
TimeLimitedCache.prototype.count = function () {
  return this.cache.size;
};

/**
 * Your TimeLimitedCache object will be instantiated and called as such:
 * var obj = new TimeLimitedCache()
 * obj.set(1, 42, 1000); // false
 * obj.get(1) // 42
 * obj.count() // 1
 */
module.exports = { TimeLimitedCache };
