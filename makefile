OUTPUT_DIR = ./dist
OUTPUT_BIN = games-admin-server.out
LOG_FILE = out.log

.PHONY: build
build:
	go build -o $(OUTPUT_DIR)/$(OUTPUT_BIN)
	@chmod +x $(OUTPUT_DIR)/$(OUTPUT_BIN)

.PHONY: stop
stop:
	@if [ -e $(OUTPUT_DIR)/pid.txt ]; then \
	  	pid=`cat $(OUTPUT_DIR)/pid.txt` && \
		kill -15 $$pid && \
		rm $(OUTPUT_DIR)/pid.txt && \
		echo "Killed $$pid" ; \
	fi

.PHONY: start
start: stop build
	@nohup $(OUTPUT_DIR)/$(OUTPUT_BIN) -d > $(OUTPUT_DIR)/$(LOG_FILE) 2>&1 & echo $$! > $(OUTPUT_DIR)/pid.txt
	@echo "Server started, all file is in dist directory"

.PHONY: run
run: build
	@$(OUTPUT_DIR)/$(OUTPUT_BIN)

.PHONY: clean
clean: stop
	@rm -rf $(OUTPUT_DIR)