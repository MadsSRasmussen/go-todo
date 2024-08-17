package manager

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/MadsSRasmussen/go-todo/internal/tasks"
)

// Manager stores tasks and enables operation on the tasks-list
type Manager struct {
	tasks    []*tasks.Task
	filename string
}

// New creates an instance of Manager and loads data from a csv-file
func New(filename string) *Manager {
	manager := &Manager{
		tasks:    []*tasks.Task{},
		filename: filename,
	}
	manager.LoadFromFile()

	return manager
}

// LoadFromFile loads data from a csv-file
func (m *Manager) LoadFromFile() {

	file, err := os.Open(m.filename)
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(data); i++ {

		row := data[i]
		id, err := strconv.Atoi(row[0])
		if err != nil {
			panic(err)
		}

		task := tasks.New(id, row[1])
		m.tasks = append(m.tasks, task)
	}

}

// WriteToFile writes the data to a csv-file
func (m *Manager) WriteToFile() {

	file, err := os.Create(m.filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < len(m.tasks); i++ {
		task := m.tasks[i]
		record := []string{strconv.Itoa(task.ID), task.Description}
		writer.Write(record)
	}

}

// AddTask adds a task
func (m *Manager) AddTask(desc string) {
	task := tasks.Task{ID: len(m.tasks) + 1, Description: desc}
	m.tasks = append(m.tasks, &task)
}

// RemoveTask removes a task
func (m *Manager) RemoveTask(id int) error {
	for i := 0; i < len(m.tasks); i++ {
		if m.tasks[i].ID == id {
			m.tasks = append(m.tasks[:i], m.tasks[i+1:]...)
			m.AssignIds()
			return nil
		}
	}

	return fmt.Errorf("task with ID %d was not found", id)
}

func (m *Manager) AssignIds() {
	for i := 0; i < len(m.tasks); i++ {
		m.tasks[i].ID = i + 1
	}
}

func (m *Manager) PrintTasks() {
	for i := 0; i < len(m.tasks); i++ {
		fmt.Printf("%s\n", m.tasks[i].String())
	}
}
