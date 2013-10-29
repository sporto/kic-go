var child_process  = require('child_process');
var util           = require('util');
var path           = require('path');
var fs             = require('fs');
var childProcesses = {};

function killProcessesSync(grunt) {
  grunt.log.writeln('killProcessesSync');

  Object.keys(childProcesses).forEach(function (key) {
    killProcessSync(key, grunt);
  });
}

function killProcess(key, grunt, callback) {
  grunt.log.writeln('killProcess ' + key);
  killProcessSync(key, grunt);
  callback(null);
}

function killProcessSync(key, grunt) {
  grunt.log.writeln('killProcessSync ' + key);
  var p = childProcesses[key];
  if (p && p.pid) {
    grunt.log.writeln('Killing process ' + p.pid);
    p.kill('SIGINT');
  } else {
    grunt.log.warn('Process not found');
  }
}

function buildAndLaunch(data, grunt, callback) {
  grunt.log.writeln('buildAndLaunch');
  grunt.log.writeln('srcPath ' + data.srcPath);
  grunt.log.writeln('srcFile ' + data.srcFile);
  grunt.log.writeln('binPath ' + data.binPath);

  killProcess(data.key, grunt, function (err) {
    buildBinary(data, grunt, function (err) {
      if (err) return callback(err);

      launchBinary(data, grunt, function (err, childProcess) {
        if (err) return callback(err);
        childProcesses[data.key] = childProcess;
        callback(null, childProcess);
      });

    });
  });

}

function buildBinary(data, grunt, callback) {
  grunt.log.writeln('buildBinary');
  // check that src path exists
  // check that srcFile exists
  var srcPath = data.srcPath;
  var srcFile = data.srcFile + '.go';
  var fullPath = path.join(srcPath, srcFile);

  grunt.log.writeln('fullPath is ' + fullPath);

  fs.exists(fullPath, function (exists) {
    if (exists) {
      grunt.log.writeln('path ' + fullPath + ' found');
      return child_process.exec('go install', {
        cwd: srcPath
      }, function (err, stdout, stderr) {
        if (err) {
          grunt.log.warn(err)
          return callback(err);
        }
        grunt.log.writeln('Binary built');
        callback(null);
      });
    } else {
      var msg = 'Source file ' + fullPath + ' not found';
      grunt.log.writeln(msg);
      callback(new Error(msg));
    }
  });

}

function launchBinary(data, grunt, callback) {
  grunt.log.writeln('launchBinary');

  var binaryFile = data.srcPath.split(path.sep).pop();
  var fullPath = path.join(data.binPath, binaryFile);

  grunt.log.writeln('Binary path is ' + fullPath);

  // check that binary exists
  return fs.exists(fullPath, function (exists) {
    if (exists) {
      grunt.log.writeln('Binary file found');
      return launchBinaryFromPath(fullPath, data, grunt, callback);
    } else {
      var msg = 'Binary file ' + fullPath + ' not found';
      grunt.log.warn(msg);
      callback(new Error(msg));
    }
  });
}

function launchBinaryFromPath(fullPath, data, grunt, callback) {

  // launch the binary
  var childProcess =  grunt.util.spawn({
    cmd: fullPath,
    args: [],
    options: {
      // cwd: path
    }
  }, function (err, result, code) {
    grunt.log.warn(err);
  })
  .on('exit', function (code, signal) {
    if (signal !== null) {
      grunt.log.warn(util.format('application exited with signal %s', signal));
    } else {
      grunt.log.warn(util.format('application exited with code %s', code));
    }
  });

  childProcess.stdout.pipe(process.stdout);
  childProcess.stderr.pipe(process.stderr);

  grunt.log.writeln('spawned ' + childProcess.pid);

  callback(null, childProcess);
}


module.exports = function(grunt) {
  // var cp = require('child_process');

  grunt.registerMultiTask('goserver', 'Spawn a Go Server', function() {
    var done = this.async();
    // clean data
    var data = {
      key: this.target,
      srcPath: this.data.srcPath,
      srcFile: this.data.srcFile.split('.')[0],
      binPath: this.data.binPath
    }
    buildAndLaunch(data, grunt, function (err) {
      if (err) grunt.log.warn(err);
      done();
    });
  });

  // grunt.registerTask('goserver:reload', function () {
  //   grunt.task.run('goserver:stop');
  // });

  grunt.registerTask('goserver:stop', 'Stops the Go server watcher', function () {
    killProcessesSync(grunt);
  });

  process.on('exit', function() {
    killProcessesSync(grunt);
  });

};