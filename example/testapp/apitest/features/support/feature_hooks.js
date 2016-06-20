
var cucumberApi = require('cucumber-api');
var Promise = require('bluebird');
var childProcess = Promise.promisifyAll(require('child_process'), {multiArgs: true});

module.exports = function () {
  var self = this;

  this.registerHandler('BeforeFeatures', function (event, done) {
    this.serverPID = childProcess.exec("../src/serv")
    /* Start application and database here */
    done()
  });


  this.registerHandler('AfterFeatures', function (event, done) {
    /* Stop application and db teardown here */
    this.serverPID.kill()
    done()
  });
};
