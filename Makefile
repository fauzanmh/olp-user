# DATABASE
DB_USER=postgres
DB_PASSWORD=*r00t123*
DB_HOST=127.0.0.1
DB_PORT=5432
DB_NAME=online_store
DB_SSL=disable

# INSTALL PROGRAM
install:
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest && \
	go get -u github.com/swaggo/swag/cmd/swag && go get -u github.com/cosmtrek/air && \
	go mod vendor && swag init   

# RUN PROGRAM
run:
	air -c config/.air.toml