var { containsDuplicate } = require("./0217-contains-duplicate");
var { isAnagram } = require("./0242-valid-anagram");
var { convertToTitle } = require("./0168-excel-sheet-column-title");
var { createHelloWorld } = require("./2667-create-hello-world-function");
var { createCounter } = require("./2620-counter");
const { expect } = require("./2704-to-be-or-not-to-be");
const { createCounter2 } = require("./2665-counter-ii");
const { map } = require("./2635-apply-transform-over-each-element-in-array");
const { filter } = require("./2634-filter-elements-from-array");
const { reduce } = require("./2626-array-reduce-transformation");
const { compose } = require("./2629-function-composition");
const { once } = require("./2666-allow-one-function-call");

console.log(containsDuplicate([1, 2, 3, 1]));
console.log(isAnagram("rat", "car"));
console.log(convertToTitle(701));
console.log(createHelloWorld()());
const counter = createCounter(10);
console.log(counter());
console.log(expect(5).toBe(5));
const counter2 = createCounter2(5);
console.log(counter2.increment());
console.log(counter2.reset());
console.log(counter2.decrement());
console.log(
  map([1, 2, 3], function plusI(n, i) {
    return n + i;
  })
);
console.log(
  filter([-2, -1, 0, 1, 2], function plusOne(n) {
    return n + 1;
  })
);
console.log(
  reduce(
    [],
    function sum(accum, curr) {
      return 0;
    },
    25
  )
);
const composeFn = compose([]);
console.log(composeFn(42));

let fn = (a, b, c) => a + b + c;
let onceFn = once(fn);
console.log(onceFn(1, 2, 3));
console.log(onceFn(2, 3, 6));
