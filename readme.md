# Alura Store Go project
## Set up
Setup the database running the docker compose 
```bash
docker compose up -d
```
Execute the migrations 
```bash
docker run -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "postgres://hello:hello@localhost:5432/hello?sslmode=disable" up 
```

## Tear down
To stop the database container run the docker compose down command
```bash
docker compose down
```

If you want to undo the database migrations run the following command
```bash
 docker run -v ./migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "postgres://hello:hello@localhost:5432/hello?sslmode=disable" down -all
```

