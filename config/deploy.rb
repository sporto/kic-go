require 'mina/git'
require "mina/rsync"

set :term_mode,   nil
# set :term_mode,   :pretty
set :domain,      '192.168.0.101'
set :user,        'sebastian'
set :repository,  'https://github.com/sporto/kic.git'
set :branch,      'master'
set :deploy_to,   '/usr/local/var/www'
set :rsync_options, %w[
	--recursive --delete --delete-excluded
	--exclude .git*
	--exclude /src/***
	--exclude /api/***
	--exclude *.go
	--exclude /config/***
	--exclude /tasks/***
	--exclude /docs/***
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

		# These are instructions to start the app after it's been prepared.
		to :launch do
			invoke :stop_api
			invoke :start_api
			# invoke :restart_nginx
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

# precompile client and api in dev machine before sending to server
# rsync_stage => tmp/deploy (in dev machine)
task :precompile do
	Dir.chdir settings.rsync_stage do
		system "npm install"
		system "grunt dist"
		system "go get"
		system "go build -o kic"
	end
end

task :stop_api do
	queue "launchctl unload ~/Library/LaunchAgents/com.sebastianporto.kic.plist"
end

task :start_api do
	# ENV variables must be set in the server
	queue "launchctl load ~/Library/LaunchAgents/com.sebastianporto.kic.plist"
end

task :restart_nginx do
	queue 'nginx -s stop'
	queue 'nginx'
end
