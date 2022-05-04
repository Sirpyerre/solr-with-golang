# API Search engine with #GO and Apache #Solr

## Run docker compose for Apache solr
```bash
docker-compose up
``` 

## Upload data to solr
```bash
curl 'http://localhost:8983/solr/gettingstarted/update?commit=true' --data-binary @jobs.json -H 'Content-type:application/json'
```

## Delete data
```bash
curl -X POST -H 'Content-Type: application/json' \
    'http://localhost:8983/solr/gettingstarted/update?commit=true' \
    -d '{ "delete": {"query":"*:*"} }'
```

## Run project
```bash
./run.sh
```

## Get result by title, description, category, salary and location
For example in web browser type:
`http://localhost/?title=desarrollador`

## Get facets by field, fields available: title, description, category, salary and location
`http://localhost:3000/facet?field=category&title=master`

## Run test:
```bash
go test internal/handlers/* -v
``` 
