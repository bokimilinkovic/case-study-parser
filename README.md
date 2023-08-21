# case-study-parser
Simple parser of csv file that is running periodically to store informations in postgres database.

In case of million records in CSV file, we can scale number of workers that will be running each in separate go routine and reading lines from a single unbuffered channel.

HTTP handler is exposed to fetch promotions based on ID
`curl localhost:8080/promotions/769b000e-1d48-4716-92a3-4285dd6cc1e8`

To start everything, just run: ```docker-compose up```
