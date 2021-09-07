package Cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/spf13/viper"
	"snap/Database/Uuid"
	"strings"
)

type Connection struct {
	Cluster *gocql.ClusterConfig
	Session *gocql.Session
}

type TableMetaData struct {
	Table      string
	Columns    map[string]struct{}
	Pk         map[string]struct{}
	Ck         map[string]struct{}
	Keyspace   string
	DependsOn  TableDependencies
	Maps       map[string]interface{}
	Connection *gocql.Session
}

type TableDependencies []TableDependency
type TableDependency func(map[string]interface{}, *gocql.Batch) bool

var connections = make(map[string]Connection)
var Configs *viper.Viper

var ConnectionManager = struct {
	GetSession func(name string) *gocql.Session
	AddSession func(keyspace string, connection Connection)
}{
	GetSession: func(name string) *gocql.Session {
		return connections[name].Session
	},
	AddSession: func(keyspace string, connection Connection) {
		fmt.Println("Connection " + keyspace + " just got set")
		connections[keyspace] = connection
	},
}

func FilterData(data map[string]interface{}, metaData TableMetaData) map[string]interface{} {
	values := make(map[string]interface{})
	for column, _ := range metaData.Columns {
		value, isset := data[column]
		switch isset {
		case true:
			values[column] = value
		}
	}
	return values
}

func GenerateEmptyInputs(count int) string {
	var inputs []string
	for i := 0; i < count; i++ {
		inputs = append(inputs, "?")
	}
	return strings.Join(inputs, ",")
}

func BindArgs(data map[string]interface{}) ([]interface{}, []string) {
	Args := []interface{}{}
	fields := make([]string, 0)
	for field, value := range data {
		Args = append(Args, value)
		fields = append(fields, field)
	}
	return Args, fields
}

func AddId(values *map[string]interface{}, idName interface{}) {
	var id string
	switch idName == nil {
	case true:
		id = GenerateUuidv4()
		break
	default:
		id = Uuid.GenerateV5(idName.(string))
	}
	_, isset := (*values)["id"]
	switch isset {
	case false:
		(*values)["id"] = id
	}
}

func GenerateUuidv4() string {
	id, _ := gocql.RandomUUID()
	return id.String()
}

func (metaData *TableMetaData) NewRecord(values map[string]interface{}, batch *gocql.Batch) bool {
	switch CheckData(&values, *metaData) {
	case false:
		return false
	}

	Args, fields := BindArgs(values)
	batch.Entries = append(batch.Entries, gocql.BatchEntry{
		Stmt:       "INSERT INTO " + metaData.Table + " (" + strings.Join(fields, ",") + ") VALUES (" + GenerateEmptyInputs(len(values)) + ")",
		Args:       Args,
		Idempotent: false,
	})
	return true
}

func (metaData *TableMetaData) GetSelectStatement(conditions map[string]interface{}, selectedFields []string) (statement *gocql.Query) {
	switch CheckData(&conditions, *metaData) {
	case false:
		return
	}

	session := metaData.Connection
	Args, fields := BindArgs(conditions)
	statement = bindValues(session.Query("SELECT "+strings.Join(selectedFields, ",")+" FROM "+metaData.Table+" WHERE "+GenerateWhereConditions(fields)), Args)

	return
}

func (metaData *TableMetaData) GetRecord(conditions map[string]interface{}, selectedFields []string) (data map[string]interface{}) {
	data = make(map[string]interface{})

	statement := metaData.GetSelectStatement(conditions,selectedFields)
	switch statement == nil {
	case true:
		return
	}

	statement.MapScan(data)

	return
}

func (metaData *TableMetaData) UpdateRecord(conditions map[string]interface{}, values map[string]interface{}, batch *gocql.Batch) bool {
	switch CheckData(&conditions, *metaData) {
	case false:
		return false
	}

	Args, fields := BindArgs(values)
	whereArgs, whereFields := BindArgs(conditions)
	batch.Entries = append(batch.Entries, gocql.BatchEntry{
		Stmt:       "UPDATE " + metaData.Table + " SET " + GenerateWhereConditions(fields) + " WHERE " + GenerateWhereConditions(whereFields),
		Args:       MergeArgs(Args, whereArgs),
		Idempotent: false,
	})

	return true
}

func MergeArgs(args ...[]interface{}) []interface{} {
	result := make([]interface{}, 0)
	for _, arg := range args {
		for _, element := range arg {
			result = append(result, element)
		}
	}

	return result
}

func bindValues(statement *gocql.Query, args []interface{}) *gocql.Query {
	for _, value := range args {
		statement = statement.Bind(value)
	}
	return statement
}

func GenerateWhereConditions(fields []string) string {
	for index, value := range fields {
		fields[index] = value + " = ?"
	}

	return strings.Join(fields, " AND ")
}

func CheckData(values *map[string]interface{}, metaData TableMetaData) bool {
	data := FilterData(*values, metaData)
	switch len(data) == 0 {
	case true:
		return false
	}
	switch CheckPK(metaData, &data) {
	case false:
		return false
	}

	values = &data
	return true
}

func AddDependencies(dependencies TableDependencies, values map[string]interface{}, statement *gocql.Batch) bool {
	isSuccessful := true
	channel := make(chan bool)
	taskNum := 0

	for _, dependency := range dependencies {
		taskNum++
		go addDependency(channel, dependency, values, statement)
	}

	for isSucceed := range channel {
		fmt.Println("succeed:", isSucceed)
		isSuccessful = isSuccessful && isSucceed
		taskNum--

		if taskNum == 0 {
			fmt.Println("getting out")
			break
		}
	}
	fmt.Println("finished")

	return isSuccessful
}

func addDependency(channel chan bool, dependency TableDependency, values map[string]interface{}, statement *gocql.Batch) {
	channel <- dependency(values, statement)
}

func CheckPK(metaData TableMetaData, data *map[string]interface{}) bool {
	for field := range metaData.Pk {
		switch _, isSet := (*data)[field]; isSet {
		case false:
			return false
		}
	}
	return true
}
