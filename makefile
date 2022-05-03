OUTPUT_DIR = ./dist
PID = $$(cat $(OUTPUT_DIR)/pid.txt)

build:
	go build -o $(OUTPUT_DIR)/games-admin-server.out
stop:
	@echo "The pid is: "$(PID)
	@kill -15 $(PID)
	@rm $(OUTPUT_DIR)/pid.txt
	@echo "Killed"
serve:
	@if [ -e $(OUTPUT_DIR)/pid.txt ]; then make stop; fi
	
	@make build

	@chmod +x $(OUTPUT_DIR)/games-admin-server.out
	@nohup $(OUTPUT_DIR)/games-admin-server.out >$(OUTPUT_DIR)/log.out 2>&1 &
	
	@echo "Server start, all file is in dist file"
