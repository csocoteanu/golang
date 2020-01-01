package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"time"
)

func getSession(clusterIP string, clusterPort int, keyspace string) (*gocql.Session, error) {
	fmt.Printf("Creating CQL session...\n")

	cluster := gocql.NewCluster(clusterIP)
	cluster.Port = clusterPort
	cluster.Keyspace = keyspace
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = 10 * time.Second
	cluster.DisableInitialHostLookup = true
	cluster.Consistency = gocql.One

	return cluster.CreateSession()
}

func insertRecord(session *gocql.Session, tag, title string) (time.Time, string) {
	insertTime := time.Now()
	videoID := gocql.UUIDFromTime(insertTime).String()

	fmt.Printf("\nInserting row: | %s | %+v | %s | %s \n", tag, insertTime, videoID, title)

	err := session.Query("INSERT INTO killrvideo.video_by_tag(tag, added_date, video_id, title) values(?, ?, ?, ?)", tag, insertTime, videoID, title).Exec()
	if err != nil {
		log.Fatal(err)
	}

	return insertTime, videoID
}

func deleteRecord(session *gocql.Session, tag string, insertTime time.Time, videoID string) {
	fmt.Printf("\nRemoving row: | %s | %+v | %s | <title> \n", tag, insertTime, videoID)

	err := session.Query("DELETE from killrvideo.video_by_tag WHERE tag=? AND added_date=? AND video_id=?", tag, insertTime, videoID).Exec()
	if err != nil {
		log.Fatal(err)
	}
}

func getRecords(session *gocql.Session) {
	var tag, videoID, title string
	var addedDate time.Time

	fmt.Printf("\n----------------- Printing rows -----------------\n")
	defer fmt.Printf("-------------------------------------------\n")

	iter := session.Query("SELECT tag, added_date, video_id, title FROM video_by_tag where tag='cassandra' AND added_date > '2013-01-01'").Iter()
	for iter.Scan(&tag, &addedDate, &videoID, &title) {
		fmt.Printf("| %s | %+v | %s | %s \n", tag, addedDate, videoID, title)
	}

	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	session, err := getSession("127.0.0.1", 9999, "killrvideo")
	if err != nil {
		log.Fatal(err)
	}

	getRecords(session)
	insertTime, videoID := insertRecord(session, "cassandra", "This is my first row inserted!")
	getRecords(session)
	deleteRecord(session, "cassandra", insertTime, videoID)
	getRecords(session)

	session.Close()
}
