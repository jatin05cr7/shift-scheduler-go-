package main
import (
	"encoding/json"
	"net/http"
	"sync"
)

type Employee struct{

	ID int `json:"json.id"`
	Name string `json:"name"`
}
type Shift struct{
	EmployeeId int `json:"employee_id"`
	Day string `json:"day"`
	Shift string `json:"shift"`
}
var(
employees []Employee
shifts []shift
employeeMux sync.Mutex
NextEmployeeID int=1
)
func main()  {
	http.HandlerFunc("/",homeHandler)
    http.HandlerFunc("/employees",employeesHandler)	
	http.HandlerFunc("/shifts",shiftsHandler)
	http.HandlerFunc("/schedule",scheduleHanlder)

	http.ListenAndServe(":8080",nil)
}

func homeHandler(w http.ResponseWriter,r *http.Request){
	w.Write([]byte(`
	<html>
	<body>
	<h1>Shift Scheduler</h1>
	<form>
	
	</form>

	</body>
	
	</html>
	
	
	
	
	
	`))
}