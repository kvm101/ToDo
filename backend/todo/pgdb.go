package todo

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 2345
	login    = "postgres"
	password = "admin"
	dbname   = "postgres"
)

type user struct {
	ID       uint32 `json:"user_id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type task struct {
	Task_id     int16  `json:"task_id"`
	Head        string `json:"head"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	Complexity  byte   `json:"complexity"`
	Importance  byte   `json:"importance"`
	Date        string `json:"date"`
}

func getDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, login, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	return db
}

func addTask(id uint32, head, description string, complexity, importance byte) error {
	if complexity > 3 {
		return fmt.Errorf("have(complexity = %d) \nwant (complexity <= 3)", complexity)
	}

	if importance > 3 {
		return fmt.Errorf("have(importance = %d) \nwant (importance <= 3)", importance)
	}

	db := getDB()
	defer db.Close()

	listQuery := `CREATE TABLE IF NOT EXISTS list(
		task_id SERIAL PRIMARY KEY,
		head VARCHAR(50) NOT NULL,
		description VARCHAR(255),
		done BOOLEAN DEFAULT FALSE,
		complexity NUMERIC(1,0) DEFAULT 0,
		importance NUMERIC(1,0) DEFAULT 0,
		date DATE NOT NULL DEFAULT CURRENT_DATE
	)`

	insertTask := `insert into "list"("head", "description", "complexity", "importance", "user_id") values($1, $2, $3, $4, $5)`

	_, err := db.Exec(listQuery)
	if err != nil {
		log.Println(err)
	}

	_, err = db.Exec(insertTask, head, description, complexity, importance, id)
	if err != nil {
		log.Println(err)
	}

	return nil
}

func deleteTask(id uint32, task_id int16) {
	db := getDB()
	defer db.Close()

	deleteTask := `delete from "list" where task_id=$1 and "user_id"=$2`

	_, err := db.Exec(deleteTask, task_id, id)
	if err != nil {
		log.Println(err)
	}
}

func editTask(id uint32, task_id int16, head, description string, complexity, importance byte) error {
	if complexity <= 3 {
		return fmt.Errorf("have(complexity = %d) \nwant (complexity <= 3)", complexity)
	}

	if importance <= 3 {
		return fmt.Errorf("have(importance = %d) \nwant (importance <= 3)", importance)
	}

	db := getDB()
	defer db.Close()
	updateTask := `update "list" set "head"=$1, "description"=$2, "complexity"=$3, "importance"=$4 where "task_id"=$5 and "user_id"=$6`

	_, err := db.Exec(updateTask, head, description, complexity, importance, task_id, id)
	if err != nil {
		log.Println(err)
	}

	return nil
}

func doneTask(id uint32, task_id int16, done bool) {
	db := getDB()
	defer db.Close()

	doneTask := `update "list" set "done"=$1 where "task_id"=$2 and "user_id"=$3`

	_, err := db.Exec(doneTask, done, task_id, id)
	if err != nil {
		log.Println(err)
	}
}

func readTasks(id uint32, section string, sortf string) ([]task, error) {
	var err error
	var tasks []task = make([]task, 0)
	db := getDB()
	defer db.Close()

	var section_tasks string

	if section == "" {
		section = "all"
	}

	s_id := fmt.Sprintf(`"user_id"=%d`, id)

	switch section {
	case "all":
		section_tasks = `select * from "list" where ` + s_id
	case "done":
		section_tasks = `select * from "list" where "done"=TRUE and ` + s_id

	case "undone":
		section_tasks = `select * from "list" where "done"=FALSE and ` + s_id

	default:
		err = fmt.Errorf("not correct section in function readTasks()")
		return nil, nil
	}

	if sortf != "" {
		section_tasks += ` ORDER BY ` + sortf
	}

	rows, err := db.Query(section_tasks)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var task_id int16
		var head string
		var description string
		var done bool
		var complexity byte
		var importance byte
		var date string
		rows.Scan(&task_id, &head, &description, &done, &complexity, &importance, &date)
		tasks = append(tasks, task{task_id, head, description, done, complexity, importance, date})
	}

	if err != nil {
		log.Println(err)
	}

	return tasks, err
}

func authentification(login string, password string) bool {
	db := getDB()
	defer db.Close()

	var count1, count2, count3 string

	err := db.QueryRow(`SELECT * FROM users WHERE "login"=$1 AND "password"=$2`, login, password).Scan(&count1, &count2, &count3)
	if err != nil {
		log.Println(err)
	}

	if count1 > "" || count2 > "" || count3 > "" {
		fmt.Println(count1, count2, count3)
		return true
	} else {
		return false
	}
}
