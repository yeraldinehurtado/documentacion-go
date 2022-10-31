package main

import "fmt"

func main() {
	t1 := Task{
		name:      "Curso de Go",
		desc:      "Completar curso de Go",
		completed: false,
	}

	t2 := Task{
		name:      "Curso de Html",
		desc:      "Completar curso de Html",
		completed: true,
	}

	t1.toPrint()
	t2.toPrint()

	List := TaskList{}
	List.appendTask(&t1)
}

// ? tarea
type Task struct {
	name      string
	desc      string
	completed bool
}

type TaskList struct {
	tasks []*Task
}

func (taskList *TaskList) appendTask(t *Task) {
	taskList.tasks = append(taskList.tasks, t) // agregar lo que trae t
}

func (taskList *TaskList) deleteTask(index int) { // metodo que elimina tarea por el indice
	taskList.tasks = append(taskList.tasks[:index], taskList.tasks[index+1:]...)
}

func (t *Task) markCompleted() { 
	t.completed = true
}

func (t *Task) toPrint() {
	fmt.Printf("Name: %s\nDescription %s\ncompleted %t", t.name, t.desc, t.completed)
}

/*
imprime ->

Name: Curso de Go
Description Completar curso de Go
completed falseName: Curso de Html
Description Completar curso de Html
completed true
*/