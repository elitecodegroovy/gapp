package sqlstore

import (
	"fmt"
	"github.com/elitecodegroovy/gapp/pkg/bus"
	m "github.com/elitecodegroovy/gapp/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetDBHealthQuery)
	fmt.Println("Initialized sqlstore DB health....")
}

func GetDBHealthQuery(query *m.GetDBHealthQuery) error {
	return x.Ping()
}
