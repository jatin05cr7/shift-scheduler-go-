package main
import (
	"encoding/json"
	"net/http"
	"sync"
	"strconv"
)

type Employee struct{

	ID int `json:"json.id"`
	Name string `json:"name"`
}
type Shift struct{
	EmployeeID int `json:"employee_id"`
	Day string `json:"day"`
	Shift string `json:"shift"`
}
var(
employees []Employee
shifts []Shift
employeeMux sync.Mutex
shiftMux sync.Mutex
nextEmployeeID int=1
)
func main()  {
	http.HandleFunc("/",homeHandler)
    http.HandleFunc("/employees",employeeHandler)	
	http.HandleFunc("/shifts",shiftsHandler)
	http.HandleFunc("/schedule",scheduleHanlder)

	http.ListenAndServe(":8080",nil)
}
func employeeHandler(w http.ResponseWriter,r *http.Request){
	switch  r.Method {
	case http.MethodPost:
		name:=r.FormValue("name")
		if name == ""{
			http.Error(w,"Name is Required",http.StatusBadRequest)
			return

		}
		employeeMux.Lock()
		employee:=Employee{ID:nextEmployeeID,Name:name}
		nextEmployeeID++
		employees=append(employees,employee)
		employeeMux.Unlock()
		http.Redirect(w,r,"/",http.StatusSeeOther)
		

	case http.MethodGet:
		w.Header().Set("Content-Type","application/json")
		employeeMux.Lock()
		json.NewEncoder(w).Encode(employees)
		employeeMux.Unlock()
	default:
		http.Error(w,"Invalid request method",http.StatusMethodNotAllowed)
	}
 }
 func shiftsHandler(w http.ResponseWriter,r *http.Request) {
switch r.Method {
case http.MethodPost:
	employeeIDstr:=r.FormValue("employee_id")
		day:=r.FormValue("day")
		shift:= r.FormValue("shift")
		if employeeIDstr == "" || day == "" || shift ==""{
			http.Error(w,"All fields are required ",http.StatusBadRequest)
			return
		}
	employeeID,err:= strconv.Atoi(employeeIDstr)
	if err != nil{
		http.Error(w,"Invalid emloyee ID",http.StatusBadRequest)
	}
		shiftMux.Lock()
		shifts=append(shifts,Shift{EmployeeID:employeeID,Day:day,Shift:shift})
		shiftMux.Unlock()
		http.Redirect(w,r,"/",http.StatusSeeOther)
	
	case http.MethodGet:
		w.Header().Set("Content-Type","application/json")
		shiftMux.Lock()
		json.NewEncoder(w).Encode(shifts)
		shiftMux.Unlock()
	default:
		http.Error(w,"Invalid Request Method",http.StatusMethodNotAllowed)
}
}
 func scheduleHanlder(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	shiftMux.Lock()
	json.NewEncoder(w).Encode(shifts)
	shiftMux.Unlock()
 }

func homeHandler(w http.ResponseWriter,r *http.Request){
	w.Write([]byte(`
	<html>
            <body>
                <h1>Shift Scheduler</h1>
                <h2>Add Employee</h2>
                <form action="/employees" method="POST">
                    Name: <input type="text" name="name"><br>
                    <input type="submit" value="Add Employee">
                </form>
                <h2>Assign Shift</h2>
                <form action="/shifts" method="POST">
                    Employee ID: <input type="text" name="employee_id"><br>
                    Day: <input type="text" name="day"><br>
                    Shift: <input type="text" name="shift"><br>
                    <input type="submit" value="Assign Shift">
                </form>
                <h2>View Schedule</h2>
                <a href="/schedule">View Schedule</a>
            </body>
        </html>
	`))
}


