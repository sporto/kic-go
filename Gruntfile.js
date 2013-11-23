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

		clean: [".tmp", "dist"],

		copy: {
			main: {
				src: 'src/index.html',
				dest: 'dist/index.html',
			},
		},

		less: {
			dev: {
				files: {
					"src/public/css/app/main.css": "src/public/css/app/main.less"
				}
			}
		},

		jshint: {
			all: ['src/public/app/**/*.js']
		},

		concat: {
			// add templates.js into app.js
			// this needs to run after concat:generated (created by usemin)
			templates: {
				src: ['.tmp/concat/js/app.js', '.tmp/templates.js'],
				dest: '.tmp/concat/js/app.js'
			}
		},

		ngmin: {
			all: {
				src: ['.tmp/concat/js/app.js'],
				dest: '.tmp/concat/js/app.js'
			}
		},

		// compiles angular tempaltes
		// into a js file
		ngtemplates: {
			dist: {
				src: 'src/public/views/**/*.html',
				dest: '.tmp/templates.js',
				options: {
					// add the generated templates into concat:templates task
					concat: 'templates'
				}
			}
		},

		useminPrepare: {
			html: ['src/index.html']
		},

		usemin: {
			html: ['dist/index.html']
		}


	});

	// load tasks
	grunt.loadNpmTasks('grunt-contrib-clean');
	grunt.loadNpmTasks('grunt-contrib-copy');
	grunt.loadNpmTasks('grunt-contrib-jshint');
	grunt.loadNpmTasks('grunt-contrib-concat');
	grunt.loadNpmTasks('grunt-contrib-uglify');
	grunt.loadNpmTasks('grunt-contrib-cssmin');
	grunt.loadNpmTasks('grunt-contrib-less');
	grunt.loadNpmTasks('grunt-usemin');
	grunt.loadNpmTasks('grunt-ngmin');
	grunt.loadNpmTasks('grunt-angular-templates');

	grunt.loadNpmTasks('grunt-contrib-watch');
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

	// grunt.registerTask('dist', ['lint', 'clean', 'concat', 'ngmin', 'jsmin', 'cssmin', 'less:dist']);
	// clean dist and .tmp
	// compile less
	// generate usemin config from index.html
	// concat all files and copy to .tmp
	// uglify
	// modify html (usemin)
	grunt.registerTask('dist', ['lint', 'clean', 'copy', 'less', 'useminPrepare', 'ngtemplates', 'concat:generated', 'concat:templates', 'ngmin', 'uglify', 'cssmin', 'usemin']);

};