module.exports = function(grunt) {
  // Do grunt-related things in here

  // Project configuration.
  grunt.initConfig({

    pkg: grunt.file.readJSON('package.json'),


    forever: {
      options: {
        index: 'main.go',
        command: 'go run'
      }
    },

    watch: {
      go: {
        files: ['**/*.go'],
        tasks: ['forever:restart'],
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

  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-forever');

  grunt.registerTask('default', ['watch']);

};