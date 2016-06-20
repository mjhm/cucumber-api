// var cucumberApi = require('cucumber-api');
var cucumberApi = require('../../../../../index.js');

module.exports = function () {

  var self = this;

  this.Before( function () {
    cucumberApi.config.call(this, {
      serverRoot: 'http://localhost:3000'
    });
    this.here = 'HERE3';
  });

};
