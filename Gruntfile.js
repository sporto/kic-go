module.exports = function(grunt) {

  var rerunProcess;
  var util = require('util');

  // Project configuration.
  grunt.initConfig({

    pkg: grunt.file.readJSON('package.json'),

    watch: {
      // go: {
      //   files: ['**/*.go'],
      //   tasks: ['goserver'],
      //   options: {
      //     nospawn: true,
      //   },
      // },
    },

    goserver: {
      default: {
        package: 'github.com/sporto/kic',
        cwd: '/Users/sebastian/GoDev/src'
      }
    }

  });

  grunt.loadNpmTasks('grunt-contrib-watch');

  grunt.registerMultiTask('goserver', 'Starts and reloads Go server', function () {
    rerunProcess =  grunt.util.spawn({
      cmd: 'rerun',
      args: [this.data.package],
      options: {
        cwd: this.data.cwd
      }
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
    
    rerunProcess.stdout.pipe(process.stdout);
    rerunProcess.stderr.pipe(process.stderr);

    grunt.log.writeln('spawned ' + rerunProcess.pid);
  });

  grunt.registerTask('goserver:stop', 'Stops the Go server watcher', function () {
    if (rerunProcess) {
      rerunProcess.kill('SIGINT');
    }
  });

  process.on('exit', function() {
    grunt.task.run('goserver:stop');
  });

  grunt.registerTask('start', function () {
    grunt.task.run('goserver');
    grunt.task.run('watch');
  });

  grunt.registerTask('default', ['start']);

};