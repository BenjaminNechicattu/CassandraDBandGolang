package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

var emp_id, emp_phone, emp_sal int
var emp_city, emp_name string
var option string

func main() {
	// connect to the cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "testone"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()

	for option != "exit" {
		fmt.Println("\nall \t\t:- print all in table")
		fmt.Println("newdata \t:- add new data to table")
		fmt.Println("exit \t\t:- exit ")
		fmt.Println("Waiting for command,. ")
		fmt.Scan(&option)
		if option == "all" {
			/*if err := session.Query(`SELECT * FROM emp`).Consistency(gocql.One).Scan(&emp_id, &emp_city, &emp_name, &emp_phone, &emp_sal); err != nil {
				log.Fatal(err)
			}
			fmt.Println("all", emp_id, emp_city, emp_name, emp_phone, emp_sal)*/
			fmt.Println("\nall Entries")
			iter := session.Query(`SELECT * FROM emp`).Consistency(gocql.One).Iter()
			for iter.Scan(&emp_id, &emp_city, &emp_name, &emp_phone, &emp_sal) {
				fmt.Println(emp_id, "\t", emp_city, "\t", emp_name, "\t", emp_phone, "\t", emp_sal)
			}
			if err := iter.Close(); err != nil {
				log.Fatal(err)
			}
		} else if option == "newdata" {
			fmt.Println("enter employee id :")
			fmt.Scanln(&emp_id)
			fmt.Println("enter employee city :")
			fmt.Scanln(&emp_city)
			fmt.Println("enter employee name :")
			fmt.Scanln(&emp_name)
			fmt.Println("enter employee phone no :")
			fmt.Scanln(&emp_phone)
			fmt.Println("enter employee sal :")
			fmt.Scanln(&emp_sal)
			if err := session.Query(`INSERT INTO emp (emp_id,emp_city,emp_name,emp_phone,emp_sal) VALUES (?,?,?,?,?)`, emp_id, emp_city, emp_name, emp_phone, emp_sal).Consistency(gocql.One).Exec(); err != nil {
				log.Fatal(err)
			}
		} else if option == "exit" {
			option = "exit"
		} else {
			fmt.Println("invalid")
		}

	}

	if err := session.Query(`SELECT * FROM emp`).Consistency(gocql.One).Scan(&emp_id, &emp_city, &emp_name, &emp_phone, &emp_sal); err != nil {
		log.Fatal(err)
	}
	fmt.Println("all", emp_id, emp_city, emp_name, emp_phone, emp_sal)
}
