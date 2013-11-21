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
		},

		jshint: {
			all: ['Gruntfile.js', 'src/public/app/**/*.js']
		},

		// concat JS files
		concat: {
			options: {
				separator: ';',
			},
			distJSApp: {
				src: [
					'src/public/js/app/app.js',
					'src/public/js/app/controllers/*.js',
					'src/public/js/app/services/*.js'
				],
				dest: 'dist/public/js/app.js',
			},
			distJSLib: {
				src: ['src/public/js/lib/*.js'],
				dest: 'dist/public/js/lib.js'
			}
		},

		// minify CSS
		cssmin: {
			distLib: {
				files: {
					'dist/public/css/lib/lib.min.css': ['src/public/css/lib/*.css']
				}
			},
			distApp: {
				files: {
					'dist/public/css/app/app.min.css': ['src/public/css/app/*.css']
				}
			}
		},

		// minify JS files
		uglify: {
			dist: {
				files: [{
					src: 'dist/public/js/app.js',
					dest: 'dist/public/js/app.min.js'
				}, {
					src: 'dist/public/js/lib.js',
					dest: 'dist/public/js/lib.min.js'
				}]
			}
		}

	});

	// load tasks
	grunt.loadNpmTasks('grunt-contrib-watch');
	grunt.loadNpmTasks('grunt-contrib-jshint');
	grunt.loadNpmTasks('grunt-contrib-concat');
	grunt.loadNpmTasks('grunt-contrib-uglify');
	grunt.loadNpmTasks('grunt-contrib-cssmin');
	grunt.loadNpmTasks('grunt-goserver');

	// custom tasks
	grunt.registerTask('start', function() {
		grunt.task.run('goserver');
		grunt.task.run('watch');
	});

	// tasks aliases
	grunt.registerTask('default', 'start');
	grunt.registerTask('jsmin', 'uglify');
	grunt.registerTask('lint', 'jshint');
	grunt.registerTask('build', ['lint', 'concat', 'jsmin', 'cssmin']);

};