package gbus

import (
	"database/sql"
	"github.com/streadway/amqp"
	"reflect"
	"runtime"
	"strings"
)

//MessageHandler signature for all command handlers
type MessageHandler func(invocation Invocation, message *BusMessage) error

//DeadLetterMessageHandler signature for dead letter handler
type DeadLetterMessageHandler func(tx *sql.Tx, poison amqp.Delivery) error

//Name is a helper function returning the runtime name of the function bound to an instance of the MessageHandler type
func (mg MessageHandler) Name() string {
	return nameFromFunc(mg)
}

//Name is a helper function returning the runtime name of the function bound to an instance of the DeadLetterMessageHandler type
func (dlmg DeadLetterMessageHandler) Name() string {
	return nameFromFunc(dlmg)
}

func nameFromFunc(function interface{}) string {
	funName := runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name()
	splits := strings.Split(funName, ".")
	fn := strings.Replace(splits[len(splits)-1], "-fm", "", -1)
	return fn
}