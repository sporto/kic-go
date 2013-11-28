require 'mina/git'
require "mina/rsync"

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
	--exclude kic.sublime-project
	--exclude wercker.yml
	--exclude Gemfile
	--exclude Gemfile.lock
	--exclude package.json
	--exclude readme.md
	--exclude Gruntfile.js
	--exclude node_modules
]

task :deploy do
	deploy do
		# invoke :clone
		# rsync will copy all files to /usr/local/var/www/shared/deploy
		invoke "rsync:deploy"
		# invoke :build_fe
		invoke :build_api

		to :launch do
			invoke :stop_api
			invoke :start_api
			invoke :restart_nginx
		end

	end
end

# task :clone do
# 	invoke :'git:clone'
# end

# called by rsync:deploy
task "rsync:stage" do
	invoke "precompile"
end

task :precompile do
	# rsync_stage => tmp/deploy
	Dir.chdir settings.rsync_stage do
		system "npm", "install"
		system "grunt", "dist"
	end
end

task :build_fe do
	# queue %[cd #{deploy_to}/current && npm install]
	# queue %[cd #{deploy_to}/current && grunt dist]
	queue 'npm install'
	# queue 'grunt dist'
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
