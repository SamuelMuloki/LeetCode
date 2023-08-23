var { containsDuplicate } = require("./0217-contains-duplicate");
var { isAnagram } = require("./0242-valid-anagram");
var { convertToTitle } = require("./0168-excel-sheet-column-title");
var { createHelloWorld } = require("./2667-create-hello-world-function")
var { createCounter } = require("./2620-counter")
const { expect } = require("./2704-to-be-or-not-to-be");

console.log(containsDuplicate([1, 2, 3, 1]));
console.log(isAnagram("rat", "car"));
console.log(convertToTitle(701));
console.log(createHelloWorld()())
const counter = createCounter(10)
console.log(counter())
console.log(expect(5).toBe(5))
