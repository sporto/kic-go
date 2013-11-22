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
			},
			less: {
				files: ['**/*.less'],
				tasks: ['less:dev']
			}
		},

		goserver: {
			default: {
				srcPath: '/Users/Sebastian/GoDev/src/github.com/sporto/kic',
				srcFile: 'main',
				binPath: '/Users/Sebastian/GoDev/bin'
			}
		},

		less: {
			dev: {
				options: {
					// paths: ["assets/css"]
				},
				files: {
					"src/public/css/app/main.css": "src/public/css/app/main.less"
				}
			},
			dist: {
				options: {
					// paths: ["assets/css"],
					cleancss: true
				},
				files: {
					"dist/public/css/app/main.min.css": "src/public/css/app/*.less"
				}
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
				dest: 'tmp/public/js/app.concat.js',
			},
			distJSLib: {
				src: ['src/public/js/lib/*.js'],
				dest: 'tmp/public/js/lib.concat.js'
			}
		},

		ngmin: {
			all: {
				src: ['tmp/public/js/app.concat.js'],
				dest: 'tmp/public/js/app.ngmin.js'
			}
		},

		// minify CSS
		// only lib files
		// app files are using less
		cssmin: {
			distLib: {
				files: {
					'dist/public/css/lib/lib.min.css': ['src/public/css/lib/*.css']
				}
			}
		},

		// minify JS files
		uglify: {
			dist: {
				files: [{
					src: 'tmp/public/js/app.ngmin.js',
					dest: 'dist/public/js/app.js'
				}, {
					src: 'tmp/public/js/lib.concat.js',
					dest: 'dist/public/js/lib.js'
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
	grunt.loadNpmTasks('grunt-contrib-less');
	grunt.loadNpmTasks('grunt-ngmin');
	grunt.loadNpmTasks('grunt-goserver');

	// custom tasks
	grunt.registerTask('dev', function() {
		grunt.task.run('goserver');
		grunt.task.run('watch');
	});

	// tasks aliases
	grunt.registerTask('default', 'dev');

	grunt.registerTask('jsmin', 'uglify');
	grunt.registerTask('lint', 'jshint');

	// TODO need to concat, then ngmin then minify
	grunt.registerTask('dist', ['lint', 'concat', 'ngmin', 'jsmin', 'cssmin', 'less:dist']);

};