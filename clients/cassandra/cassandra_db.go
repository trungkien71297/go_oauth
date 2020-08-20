package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
	err     error
)

func init() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}

}

func GetSession() *gocql.Session {
	return session
}
