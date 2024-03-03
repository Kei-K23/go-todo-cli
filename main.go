package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
}

func main() {

	// get description value from command
	add := flag.String("add", "", "Add a new task")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Mark a task as complete")
	remove := flag.Int("remove", 0, "Remove a task for the list")

	flag.Parse()

	if *add != "" {
		// initial tasks
		var tasks []Task

		_, err := os.Stat("task.json")
		// check if task.json already exists if not exist then create new task.json file and add initial task
		if os.IsNotExist(err) {
			// create file
			file, err := os.Create("task.json")
			if err != nil {
				log.Fatal(err)
			}
			// close file
			defer file.Close()

			newTask := Task{ID: 1, Description: *add, Complete: false}

			tasks = append(tasks, newTask)

			// create json
			jsonTask, err := json.MarshalIndent(tasks, "", "\t")

			if err != nil {
				log.Fatal(err)
			}

			err = os.WriteFile("task.json", jsonTask, 0644)

			if err != nil {
				log.Fatal(err)
			}

			 fmt.Println("Task added successfully")
			return
		}

		// read file
		data, err := os.ReadFile("task.json")

		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(data, &tasks)

		if err != nil {
			log.Fatal(err)
		}

		newTask := Task{ID: len(tasks) + 1, Description: *add, Complete: false}
		tasks = append(tasks, newTask)

		updatedTask, err := json.MarshalIndent(tasks, "", "\t")
		
		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile("task.json", updatedTask, 0644)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Task added successfully")
		return
	}

	if *list { 
		// read file
		var tasks []Task

        data, err := os.ReadFile("task.json")

        if err!= nil {
            log.Fatal(err)
        }

        err = json.Unmarshal(data, &tasks)

        if err!= nil {
            log.Fatal(err)
        }
		fmt.Println("ID\tTask\t\tCompleted")
		fmt.Println("--------------------------------")
        for _, task := range tasks {
            fmt.Printf("%d\t%s\t%t\n", task.ID, task.Description, task.Complete)
        }
		fmt.Println("--------------------------------")
        return
	}

	if *complete!= 0 { 
		// read file
        var tasks []Task

        data, err := os.ReadFile("task.json")

        if err!= nil {
            log.Fatal(err)
        }

        err = json.Unmarshal(data, &tasks)

        if err!= nil {
            log.Fatal(err)
        }

		// signal to track not found task
		taskFound := false

        for i, task := range tasks {
            if task.ID == *complete {
                tasks[i].Complete = true
				taskFound = true
                break
            }
        }

		if !taskFound { 
			log.Fatal("Task not found to make as complete")
		}

        updatedTask, err := json.MarshalIndent(tasks, "", "\t")

        if err!= nil {
            log.Fatal(err)
        }

        err = os.WriteFile("task.json", updatedTask, 0644)

        if err!= nil {
            log.Fatal(err)
        }
        fmt.Printf("Task Id %d completed successfully", *complete)
	}

	if *remove!= 0 {
		var tasks []Task

		data , err := os.ReadFile("task.json")

		if err!= nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(data , &tasks)
		
		if err!= nil {
            log.Fatal(err)
        }

		taskFound := false

		for i , task := range tasks { 
			if task.ID == *remove {
				tasks = append(tasks[:i], tasks[i+1:]...)
                break
			}
		}

		if !taskFound { 
			log.Fatal("Task not found to make as complete")
		}

        updatedTask, err := json.MarshalIndent(tasks, "", "\t")

		if err!= nil {
            log.Fatal(err)
        }

		err = os.WriteFile("task.json", updatedTask, 0644)

		if err!= nil {
            log.Fatal(err)
        }

		fmt.Printf("Task Id %d removed successfully", *remove)
	}
}
