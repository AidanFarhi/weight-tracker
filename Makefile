# run app localy
local:
	python ./scripts/run_local.py

# run tests with coverage
cover:
	go test -cover ./service

# run tests
test:
	go test ./service

# deploy to Heroku
deploy: test
	git push heroku main

# sets Heroku app config
update-heroku-config:
	python ./scripts/set_heroku_config.py

# sets Heroku app config and deploys
deploy-with-config: update-heroku-config
	$(MAKE) deploy
