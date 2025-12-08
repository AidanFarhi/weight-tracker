# run app localy
local:
	python ./scripts/run_local.py

# sets Heroku app config and deploys
deploy-with-config:
	python ./scripts/set_heroku_config.py
	git push heroku main

# deploy to Heroku
deploy:
	git push heroku main