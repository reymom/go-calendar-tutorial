//go:build integration_test

package integrationtest

import (
	"context"
	"testing"
	"time"

	"github.com/reymom/go-calendar-tutorial/internal/dao/generator"
	"github.com/reymom/go-calendar-tutorial/pkg/model"
	"github.com/stretchr/testify/assert"
)

func taskDaoTests(t *testing.T) {
	t.Run("Get Tasks", testGetTasks)
	t.Run("Create-Remove Task", testCreateTask)
	t.Run("Edit Task", testEditTask)
	t.Run("Complete Task", testCompleteTask)
}

func testGetTasks(t *testing.T) {
	a := assert.New(t)
	testEnv.cleanup()

	list, e := testEnv.ListTasks(context.Background(), model.NewMonthlyFilter(time.April, 2023))
	a.NoError(e)
	a.Len(list, 0)
	if a.NotNil(list) {
		a.Len(list, 0)
	}

	//add task, startsAt and finishesAt are set at max +-24h
	task, e := testEnv.CreateTask(context.Background(), generator.GenerateRandomAddableTask())
	a.NoError(e)

	//get 1 task
	list, e = testEnv.ListTasks(context.Background(), model.NewYearlyFilter(uint(time.Now().Year())))
	a.NoError(e)
	if a.NotNil(list) {
		if a.Len(list, 1) {
			a.Equal(task, &list[0])
		}
	}

	//get 5 tasks
	for i := 0; i <= 3; i++ {
		list, e = testEnv.ListTasks(context.Background(), model.NewYearlyFilter(uint(time.Now().Year())))
		a.NoError(e)
		a.Len(list, 5)
	}
}

func testCreateTask(t *testing.T) {
	a := assert.New(t)
	testEnv.cleanup()

	genTask := generator.GenerateRandomAddableTask()
	retTask, e := testEnv.CreateTask(context.Background(), genTask)
	a.NoError(e)
	a.NotNil(retTask)
	a.Equal(model.Task{
		TaskId:      retTask.TaskId,
		Completed:   false,
		AddableTask: genTask,
	}, retTask)

}

func testEditTask(t *testing.T) {
	a := assert.New(t)
	testEnv.cleanup()

	retTask, e := testEnv.CreateTask(context.Background(), generator.GenerateRandomAddableTask())
	a.NoError(e)

	task := generator.GenerateRandomAddableTask()
	updatedTask, e := testEnv.EditTask(context.Background(), retTask.TaskId, task)
	a.NoError(e)
	a.NotNil(e)
	a.NotEqual(retTask, updatedTask)
	a.Equal(model.Task{
		TaskId:      retTask.TaskId,
		Completed:   false,
		AddableTask: task,
	}, updatedTask)
}

func testCompleteTask(t *testing.T) {
	a := assert.New(t)
	testEnv.cleanup()

	retTask, e := testEnv.CreateTask(context.Background(), generator.GenerateRandomAddableTask())
	a.NoError(e)
	a.False(retTask.Completed)

	e = testEnv.SetCompleted(context.Background(), retTask.TaskId, true)
	a.NoError(e)

	list, e := testEnv.ListTasks(context.Background(), model.NewYearlyFilter(uint(time.Now().Year())))
	a.NoError(e)
	a.True(list[0].Completed)
}
