package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID int `json:"id"`
	Text string `json:"text"`
	Completed bool `json:"completed"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Completed_at string `json:"completed_at"`
}

func WriteToFile(task []Task, filename string) error {
	jsonData, err := json.Marshal(task)
	if err != nil {
		fmt.Println("Task could not be created because of: ", err.Error())
		return err
	}
	// We write to the file
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		fmt.Println("Task could not be created because of: ", err.Error())
		return err
	}
	return nil
}
func AddTask(text string, filename string) (Task, error) {
	var jsonData []Task
	filename = filename+".json"
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Task could not be added because of: ", err.Error())
		return Task{}, err
	}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("Task could not be added because of: ", err.Error())
		return Task{}, err
	}
	task := Task {
		ID: len(jsonData)+1,
		Text: text,
		Completed: false,
		Created_at: time.Now().String(),
		Updated_at: time.Now().String(),
		Completed_at: time.Now().String(),
	}
	jsonData = append(jsonData, task)
	err = WriteToFile(jsonData, filename)
	if err != nil {
		fmt.Println("Task could not be created because of: ", err.Error())
		return Task{}, err
	} 
	return task, nil
}
func CreateTask(text string, filename string) (Task, error) {
	filename = filename+".json"
	_, err := os.Create(filename)
	if err != nil {
		fmt.Println("Task could not be created because of: ", err.Error())
		return Task{}, err
	}
	task := Task {
		ID: 1,
		Text: text,
		Completed: false,
		Created_at: time.Now().String(),
		Updated_at: time.Now().String(),
		Completed_at: time.Now().String(),
	}
	err = WriteToFile([]Task{task}, filename)
	if err != nil {
		fmt.Println("Task could not be created because of: ", err.Error())
		return Task{}, err
	} 
	return task, nil
}