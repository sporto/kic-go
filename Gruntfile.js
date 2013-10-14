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
        tasks: ['goserver:reload'],
        options: {
          nospawn: true,
        },
      },
    },

  });

  grunt.loadNpmTasks('grunt-contrib-watch');

  grunt.registerTask('goserver:start', function () {
    grunt.log.writeln('goserver.start');

    if (serverProcess) {
      grunt.task.run('goserver:stop');
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
    grunt.event.emit('goserver.started');
  });

  grunt.registerTask('goserver:stop', function (done) {
    grunt.log.warn('goserver:stop');
    // Tell grunt this task is asynchronous.
    var done = this.async();
    
    serverProcess.kill('SIGINT');
    serverProcess = null;
    setTimeout(done, 1000);
    // grunt.event.emit('goserver.stopped');
  });

  grunt.registerTask('goserver:reload', function () {
    grunt.task.run('goserver:start');
  });

  // kill the server process when grunt exits
  process.on('exit', function() {
    if (serverRunning) {
      serverProcess.kill('SIGINT');
    }
  });

  grunt.registerTask('start', function () {
    grunt.task.run('goserver:start');
    grunt.task.run('watch');
  });

  grunt.registerTask('default', ['start']);

};