require 'mina/git'

set :domain,      '192.168.0.101'
set :user,        'sebastian'
set :repository,  'https://github.com/sporto/kic.git'
set :branch,      'master'
set :deploy_to,   '/usr/local/var/www'
set :nginx_path,  '/path/to/bin/nginx'

task :deploy do
	deploy do
		
		invoke :clone
		invoke :build_fe
		invoke :build_api

		to :launch do
			invoke :stop_api
			invoke :start_api
			invoke :restart_nginx
		end

	end
end

task :clone do
	invoke :'git:clone'
end

task :build_fe do
	queue 'npm install'
	queue 'grunt dist'
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
