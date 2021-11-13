# urlShortener

Steps to execute:
1. Clone the project
2. Run go mod download
3. Run go mod vendor
4. Run mongodb in your local
5. Set mongo creds in config/config.go file [MongoServer, MongoDbName, MongoUser, MongoPassword]
6. Run the program using the command: go run main/*.go


Creating key cUrl:

    curl --location --request POST 'http://localhost:9000/shorten-url' \
    --header 'Content-Type: application/json' \
    --data-raw '{
    "url": "www.google.com"
    }'

Querying original url cUrl:

    curl --location --request GET 'http://localhost:9000/lxmKrOEfM'