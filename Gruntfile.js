module.exports = function(grunt) {

  var serverProcess;
  var serverRunning = false;
  var util = require('util');

  // Project configuration.
  grunt.initConfig({

    pkg: grunt.file.readJSON('package.json'),

    watch: {
      go: {
        files: ['**/*.go'],
        tasks: ['goserver'],
        options: {
          nospawn: true,
        },
      },
    },

  });

  grunt.loadNpmTasks('grunt-contrib-watch');

  grunt.event.on('goserver.start', function (done) {
    grunt.log.writeln('goserver.start');

    if (serverProcess) {
      return grunt.event.emit('goserver.stop', done)
    }

    grunt.log.writeln('spawning new process');
    serverProcess = grunt.util.spawn({
      cmd: 'go', 
      args: ['run', 'main.go']
    }, function (error, result, code) {
      // grunt.log.writeln(error);
    })
    .on('exit', function (code, signal) {
      if (signal !== null) {
        grunt.log.warn(util.format('application exited with signal %s', signal));
      } else {
        grunt.log.warn(util.format('application exited with code %s', code));
      }
    });
    
    serverProcess.stdout.pipe(process.stdout);
    serverProcess.stderr.pipe(process.stderr);

    grunt.log.writeln('spawned ' + serverProcess.pid);

    setTimeout(function () {
      grunt.event.emit('goserver.started', done);
    }, 250);
  });

  grunt.event.on('goserver.started', function (done) {
    grunt.log.writeln('goserver.started');
    done();
  });

  //done is the async callback coming from start
  grunt.event.on('goserver.stop', function (done) {
    grunt.log.writeln('goserver.stop');

    serverProcess.kill('SIGINT');

    grunt.log.writeln('Waiting for the Go process to die')

    setTimeout(function () {
      serverProcess = null;
      grunt.event.emit('goserver.stopped', done);
    }, 1500);
  });

  grunt.event.on('goserver.stopped', function (done) {
    grunt.log.writeln('goserver.stopped');
    return grunt.event.emit('goserver.start', done);
  });

  grunt.registerTask('goserver', function () {
    var done = this.async();
    grunt.event.emit('goserver.start', done);
  });

  // kill the server process when grunt exits
  process.on('exit', function() {
    if (serverRunning) {
      serverProcess.kill('SIGINT');
    }
  });

  grunt.registerTask('start', function () {
    grunt.task.run('goserver');
    grunt.task.run('watch');
  });

  grunt.registerTask('default', ['start']);

};