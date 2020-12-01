# CassandraDB_and_Golang
This Repository contains a Golang Program for adding data into a table employee in Cassandra DB and View data in the table, all from the Console.

# Before You Go
This program inserts values to a table (table name 'emp') by cassandra database. Make sure a table is created. 
in this program,
`create table example.tweet(timeline text, id UUID, text text, PRIMARY KEY(id));`
now run the program
`main.go`

make sure 
line ~16
`cluster := gocql.NewCluster("127.0.0.1")` ip corresponds to your system local host

Reference :
  https://github.com/gocql/gocql
