module.exports = function(grunt) {

  var serverProcess;
  var serverProcessEventStart   = 'goserver.start';
  var serverProcessEventStarted = 'goserver.started';
  var serverProcessEventStop    = 'goserver.stop';
  var serverProcessEventStopped = 'goserver.stopped';
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

  grunt.event.on(serverProcessEventStart, function (done) {
    grunt.log.writeln(serverProcessEventStart);

    if (serverProcess) {
      return grunt.event.emit(serverProcessEventStop, done)
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
    
    // serverProcess is not the go process itself
    serverProcess.stdout.pipe(process.stdout);
    serverProcess.stderr.pipe(process.stderr);

    grunt.log.writeln('spawned ' + serverProcess.pid);

    setTimeout(function () {
      grunt.event.emit(serverProcessEventStarted, done);
    }, 250);
  });

  grunt.event.on(serverProcessEventStarted, function (done) {
    grunt.log.writeln(serverProcessEventStarted);
    done();
  });

  //done is the async callback coming from start
  grunt.event.on(serverProcessEventStop, function (done) {
    grunt.log.writeln(serverProcessEventStop);
    grunt.log.writeln('Sending signal to process ' + serverProcess.pid)

    // serverProcess.send('CLOSE');
    // return done();

    serverProcess.kill('SIGINT');

    grunt.log.writeln('Waiting for the Go process to die')

    setTimeout(function () {
      serverProcess = null;
      grunt.event.emit(serverProcessEventStopped, done);
    }, 1500);
  });

  grunt.event.on(serverProcessEventStopped, function (done) {
    grunt.log.writeln(serverProcessEventStopped);
    return grunt.event.emit(serverProcessEventStart, done);
  });

  grunt.registerTask('goserver', function () {
    var done = this.async();
    grunt.event.emit(serverProcessEventStart, done);
  });

  // kill the server process when grunt exits
  process.on('exit', function() {
    if (serverProcess) {
      serverProcess.kill('SIGINT');
    }
  });

  grunt.registerTask('start', function () {
    grunt.task.run('goserver');
    grunt.task.run('watch');
  });

  grunt.registerTask('default', ['start']);

};