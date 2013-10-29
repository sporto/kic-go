module.exports = function(grunt) {

  // Project configuration.
  grunt.initConfig({

    pkg: grunt.file.readJSON('package.json'),

    watch: {
      go: {
        files: ['**/*.go'],
        tasks: ['goserver'],
        options: {
          nospawn: true,
        }
      }
    },

    goserver: {
      default: {
        srcPath: '/Users/Sebastian/GoDev/src/github.com/sporto/kic',
        srcFile: 'main',
        binPath: '/Users/Sebastian/GoDev/bin'
      }
    }

  });

  grunt.loadNpmTasks('grunt-contrib-watch');
  require('./grunt_go')(grunt);

  grunt.registerTask('start', function () {
    grunt.task.run('goserver');
    grunt.task.run('watch');
  });

  grunt.registerTask('default', ['start']);

};