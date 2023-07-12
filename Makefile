.PHONY: app 

# Go parameters
GOCMD=go


# Runs piilot app
app:
	@echo "===== RUNNING SERVICES ====="
	@export GO111MODULE=auto
	@ISLOCAL=1 ENV=development $(GOCMD) run cmd/*.go
