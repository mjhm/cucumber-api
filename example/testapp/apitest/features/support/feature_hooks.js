
var cucumberApi = require('cucumber-api');
var Promise = require('bluebird');
var childProcess = Promise.promisifyAll(require('child_process'), {multiArgs: true});
var portUsed = require('tcp-port-used');

waitForStartup = function (done) {
  portUsed.check(3000, "localhost")
  .then(function(inUse) {
    if (inUse) return done();
    setTimeout(function () {
      waitForStartup(done);
    }, 50);
  })
  .catch( function (err) {
    console.log("Cannot connect to server. Please make sure that your ports are configured properly");
    done(err);
  });
}


module.exports = function () {
  var self = this;

  this.registerHandler('BeforeFeatures', function (event, done) {
    self.serverPID = childProcess.exec("../src/serv")
    /* Start application and database here */
    waitForStartup(done);
  });


  this.registerHandler('AfterFeatures', function (event, done) {
    /* Stop application and db teardown here */
    self.serverPID.kill()
    done()
  });
};
