install:
	@go get -t -v ./...

start:
	@dev_appserver.py 'appengine/app.yaml'

deploy:
	@gcloud app deploy appengine/app.yaml

