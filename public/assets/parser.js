const fs = require('fs');
const csv = require('csv-parser');
const console = console;

const parser = require('./parser');
const config = require('../config');

class Parser {
  constructor(filePath) {
    this.filePath = filePath;
  }

  async parse() {
    return new Promise((resolve, reject) => {
      const results = [];
      fs.createReadStream(this.filePath)
        .pipe(csv())
        .on('data', (row) => {
          results.push(row);
        })
        .on('end', () => {
          resolve(results);
        })
        .on('error', (err) => {
          reject(err);
        });
    });
  }
}

module.exports = Parser;