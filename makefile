OUTPUT_DIR = ./dist
OUTPUT_BIN = games-admin-server.out
LOG_FILE = $(OUTPUT_DIR)/out.log
PID_FILE = /tmp/game_admin_server.pid

.PHONY: build
build:
	go build -o $(OUTPUT_DIR)/$(OUTPUT_BIN)
	@chmod +x $(OUTPUT_DIR)/$(OUTPUT_BIN)

.PHONY: stop
stop:
	@if [ -e $(PID_FILE) ]; then \
	  	pid=`cat $(PID_FILE)` && \
		((kill -15 $$pid && echo "Killed $$pid") || echo "No such process, ignored") && \
		rm $(PID_FILE) ; \
	fi

.PHONY: start
start: stop build
	@nohup $(OUTPUT_DIR)/$(OUTPUT_BIN) -d > $(LOG_FILE) 2>&1 & echo $$! > $(PID_FILE)
	@echo "Server started, all file is in dist directory"

.PHONY: run
run: build
	@$(OUTPUT_DIR)/$(OUTPUT_BIN)

.PHONY: clean
clean: stop
	@rm -rf $(OUTPUT_DIR)