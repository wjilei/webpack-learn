require("./main.css");
const logo = require("./logo.png");

document.querySelector('#btn').addEventListener('click', function(e) {
    alert('Hello webpack');
}, false)

document.querySelector('#logo').src=logo;