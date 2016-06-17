// var cucumberApi = require('cucumber-api');
var Promise = require('bluebird');
var childProcess = Promise.promisifyAll(require('child_process'), {multiArgs: true});
var cucumberApi = require('../../../../../index.js');

resetDatabase = function() {
  childProcess.exec("../src/serv --reset-db")
}

module.exports = function () {
  cucumberApi.steps.shell.call(this);
  cucumberApi.steps.request.call(this);

  var self = this;
  self.Given(/^testa$/, function () {
    console.log('testa self', self);
    console.log('testa this', this);
  });

  self.Then(/^the database is reset$/, resetDatabase);


  self.Given(/^testb$/, function () {
  });

};
