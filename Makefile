VERSION      = $(shell git rev-parse --short --verify HEAD)
ACCESS_TOKEN = $(shell gcloud auth print-access-token)

test:
	@go test ./gae/...

deploy: PROJECT_ID
	@appcfg.py update --oauth2_access_token $(ACCESS_TOKEN) \
		-A $(PROJECT_ID) \
		-V $(VERSION) \
		gae/app.yaml
	@gcloud app services set-traffic default --project $(PROJECT_ID) --splits $(VERSION)=1 --quiet

PROJECT_ID:
ifndef PROJECT_ID
	@echo environment variable PROJECT_ID not set
	exit 1
endif

.PHONY: test deploy PROJECT_ID
