require 'mina/git'
require "mina/rsync"

set :term_mode,   nil
# set :term_mode,   :pretty
set :domain,      '192.168.0.101'
set :user,        'sebastian'
set :repository,  'https://github.com/sporto/kic.git'
set :branch,      'master'
set :deploy_to,   '/usr/local/var/www'
set :nginx_path,  '/path/to/bin/nginx'
set :rsync_options, %w[
	--recursive --delete --delete-excluded
	--exclude .git*
	--exclude /src/***
	--exclude /api/***
	--exclude *.go
	--exclude /config/***
	--exclude /tasks/***
	--exclude kic.sublime-project
	--exclude wercker.yml
	--exclude Gemfile
	--exclude Gemfile.lock
	--exclude package.json
	--exclude readme.md
	--exclude Gruntfile.js
	--exclude node_modules
]
set :rsync_copy, "/usr/local/bin/rsync --archive --acls --xattrs"

task :deploy do
	deploy do
		# rsync will copy all files to /usr/local/var/www/shared/deploy
		invoke "rsync:deploy"

		invoke :build_api

		# These are instructions to start the app after it's been prepared.
		to :launch do
			invoke :stop_api
			invoke :start_api
			invoke :restart_nginx
		end

    # This optional block defines how a broken release should be cleaned up.
		to :clean do
			queue 'log "failed deployment"'
		end

	end
end

# called by rsync:deploy
task "rsync:stage" do
	invoke :precompile
end

task :precompile do
	# this helps in the dev machine
	# rsync_stage = tmp/deploy
	Dir.chdir settings.rsync_stage do
		system "npm", "install"
		system "grunt", "dist"
		system "go", "get"
		system "go build -o kic"
	end
end

task :build_api do
	# queue 'go install main.go'
end

task :stop_api do
	# queue 'main -SIG'
end

task :start_api do
	# queue 'SOMETHING=ssks main'
end

task :restart_nginx do
	# /usr/local/bin/nginx
	# queue "#{settings.nginx_path!}/sbin/nginx restart"
end
