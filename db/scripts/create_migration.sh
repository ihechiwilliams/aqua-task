migrate \
		-database "${DATABASE_URL}" \
		create -ext sql -dir db/migrations "$1"
