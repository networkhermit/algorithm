'use strict';

const readline = require('readline');

const input = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

if (module === require.main) {
  input.on('line', (line) => {
    const pair = line.replace(/(^\s*)|(\s*$)/g, '').split(/ +/);
    const a = Number(pair[0]);
    const b = Number(pair[1]);
    console.log(a + b);
  });
}
