module.exports = function(grunt) {
  // Do grunt-related things in here

  // Project configuration.
  grunt.initConfig({

    pkg: grunt.file.readJSON('package.json'),


    // forever: {
    //   options: {
    //     index: 'main.go',
    //     command: 'go run'
    //   }
    // },

    watch: {
      go: {
        files: ['**/*.go'],
        tasks: ['start'],
        options: {
          spawn: false,
        },
      },
    },

    uglify: {
      options: {
        banner: '/*! <%= pkg.name %> <%= grunt.template.today("yyyy-mm-dd") %> */\n'
      },
      build: {
        src: 'src/<%= pkg.name %>.js',
        dest: 'build/<%= pkg.name %>.min.js'
      }
    }
  });

  // exit is working
  // but watch on changes is not
  // ALSO show GO errors

  grunt.loadNpmTasks('grunt-contrib-watch');
  // grunt.loadNpmTasks('grunt-forever');

  var serverProcess;

  grunt.registerTask('start', function () {
    
    grunt.log.writeln('start...');

    serverProcess = grunt.util.spawn({
        cmd: 'go', 
        args: ['run','main.go']
      });

    grunt.task.run('watch');
  });

  grunt.registerTask('killServer', function () {
    grunt.log.writeln('killServer');

    if (serverProcess) serverProcess.kill();
  });

  grunt.registerTask('default', ['watch']);

  process.on('exit', function () {
    grunt.log.writeln('exit...');
    grunt.task.run('killServer');
  });

};