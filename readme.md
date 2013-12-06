KIC
===

[![wercker status](https://app.wercker.com/status/a32c249547feeb153cd6481fff4ce782/m "wercker status")](https://app.wercker.com/project/bykey/a32c249547feeb153cd6481fff4ce782)

Kids Investment Company


Development
----

In development Go takes care of serving all assets. Run the go application using:

  go run main.go

  or

  grunt dev

Production
----------

In production all assets are precompiled using Grunt as well as the go application. The mina deploy script takes care of this.
To test the assets compilation do `grunt dist`

In the server nginx (or similar) must be running. Nginx should server all the static assets and proxy all requests to \api to the go process running on port 9000.

Running Tasks
----------

Install Gofer

  go get -u github.com/chuckpreslar/gofer/gofer

Then

  gofer accounts:create


Testing
-------

Install Ginkgo
	
	go get github.com/onsi/ginkgo
	go get github.com/onsi/gomega
	go install github.com/onsi/ginkgo/ginkgo

	ginkgo -r

Deployment
-----------

To deploy

  mina deploy

This is the current deployment process:

  - npm init
  - grunt dist
  - build main.go
  - rsync files to server
  - restart go process in server
  - restart nginx?

Server configuration
---------------------

Server must be running rethinkdb in port as specified by the env variable KIC_PROD_DB_HOST.
The go application must be launched using a daemon controller e.g. Upstart

At the moment I am using a Mac server, so I am using launchd, example launchd plist

  <?xml version="1.0" encoding="UTF-8"?>
  <!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
  <plist version="1.0">
  <dict>
      <key>Label</key>
      <string>com.sebastianporto.kic</string>
      <key>ProgramArguments</key>
      <array>
          <string>/usr/local/var/www/current/kic</string>
      </array>
      <key>RunAtLoad</key>
          <true/>
      <key>KeepAlive</key>
      <true/>
          <key>StandardOutPath</key>
          <string>/usr/local/var/www/kic_process.log</string>
          <key>StandardErrorPath</key>
          <string>/usr/local/var/www/kic_process.log</string>
          <key>EnvironmentVariables</key>
          <dict>
                  <key>ENV</key>
                  <string>prod</string>
                  <key>KIC_PROD_DB_HOST</key>
                  <string>localhost:28015</string>
                  <key>KIC_PROD_DB_NAME</key>
                  <string>kic</string>
          </dict>
  </dict>
  </plist>