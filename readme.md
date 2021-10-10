# How to use

Database structure is in
app/database/database.sql

Compile Binary:
```sh
cd backend
go build -o bin/app main.go
./bin/app
```
Then
Run docker
```sh
docker compose up -d
docker exec -it mysql-container bash
mysql -u alphasoft -p dbmusic  < database.sql
```

Or we already serve the compiled source in bin/main-linux-64 or bin/main-macos-bigsur

Run the vuejs file
```sh
cd frontend
npm install
npm run serve
```
Or copy the dist directory in your webserver

Note:
If you run in the webserver, and access the frontend url without port. Then you must config the nginx to
proxy pass /api to :8080 
