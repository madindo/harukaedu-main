package errors
import (
	"fmt"
	"github.com/madindo/harukaedu-main/logs"
	"net/http"
)

func ErrorConnection(err error) {
	if err != nil {
		fmt.Println(err.Error())
		logs.Logging("ERROR", "Failed to open")
		panic("Failed to open");
	}
}

func LogNPrint(response http.ResponseWriter, msg string, msgType string) {
	fmt.Fprintf(response, msg)
	logs.Logging(msgType, msg)
}
