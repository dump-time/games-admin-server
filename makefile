OUTPUT_DIR = ./dist
PID = $$(cat $(OUTPUT_DIR)/pid.txt)

build:
	go build -o $(OUTPUT_DIR)/games-admin-server.out
serve:
	@echo "Server start, all file is in dist file"
	@chmod +x $(OUTPUT_DIR)/games-admin-server.out
	@nohup $(OUTPUT_DIR)/games-admin-server.out >$(OUTPUT_DIR)/log.out 2>&1 &
stop:
	@echo "The pid is: "$(PID)
	@kill -15 $(PID)
	@echo "Killed"
restart:
	@echo "Stopping server ..."
	@make stop
	@echo "Starting server ..."
	@ make serve	
	
