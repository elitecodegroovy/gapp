package sqlstore

import (
	"github.com/elitecodegroovy/gnetwork/pkg/setting"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type sqlStoreTest struct {
	name          string
	dbType        string
	dbHost        string
	connStrValues []string
}

var sqlStoreTestCases = []sqlStoreTest{
	{
		name:          "MySQL IPv4",
		dbType:        "mysql",
		dbHost:        "1.2.3.4:5678",
		connStrValues: []string{"tcp(1.2.3.4:5678)"},
	},
	{
		name:          "MySQL IPv4 (Default Port)",
		dbType:        "mysql",
		dbHost:        "1.2.3.4",
		connStrValues: []string{"tcp(1.2.3.4)"},
	},
	{
		name:          "MySQL IPv6",
		dbType:        "mysql",
		dbHost:        "[fe80::24e8:31b2:91df:b177]:1234",
		connStrValues: []string{"tcp([fe80::24e8:31b2:91df:b177]:1234)"},
	},
	{
		name:          "MySQL IPv6 (Default Port)",
		dbType:        "mysql",
		dbHost:        "::1",
		connStrValues: []string{"tcp(::1)"},
	},
}

func TestSqlConnectionString(t *testing.T) {
	Convey("Testing SQL Connection Strings", t, func() {
		t.Helper()

		for _, testCase := range sqlStoreTestCases {
			Convey(testCase.name, func() {
				sqlstore := &SqlStore{}
				sqlstore.Cfg = makeSqlStoreTestConfig(testCase.dbType, testCase.dbHost)
				sqlstore.readConfig()

				connStr, err := sqlstore.buildConnectionString()

				So(err, ShouldBeNil)
				for _, connSubStr := range testCase.connStrValues {
					So(connStr, ShouldContainSubstring, connSubStr)
				}
			})
		}
	})
}

func makeSqlStoreTestConfig(dbType string, host string) *setting.Cfg {
	cfg := setting.NewCfg()

	sec, _ := cfg.Raw.NewSection("database")
	sec.NewKey("type", dbType)
	sec.NewKey("host", host)
	sec.NewKey("user", "user")
	sec.NewKey("name", "test_db")
	sec.NewKey("password", "pass")

	return cfg
}