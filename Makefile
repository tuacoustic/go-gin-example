#AUTHOR - TUDINH
PROJECT_NAME 	     := "GO GIN EXAMPLE"

#PWD
PROJECT_PWD 	   := $(shell pwd)

#COLOR
RED          := $(shell tput -Txterm setaf 1)
GREEN        := $(shell tput -Txterm setaf 2)
YELLOW       := $(shell tput -Txterm setaf 3)
RESET 		 := $(shell tput -Txterm sgr0)
ALERT        := $(YELLOW)$(DATE)$(RESET)

# Package Reload
.PHONY: get-air
get-air: 
	@echo "[$(ALERT)] - make get-air -> $(GREEN)$(PROJECT_NAME)$(RESET)"
	@curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b ${PROJECT_PWD}/bin

.PHONY: hot-reload
hot-reload:
	@echo "[$(ALERT)] - make hot-reload -> $(GREEN)$(PROJECT_NAME)$(RESET)"
	@./bin/air

.PHONY: hot-reload-win
hot-reload-win:
	@echo "[$(ALERT)] - make hot-reload -> $(GREEN)$(PROJECT_NAME)$(RESET)"
	@./bin/air.exe