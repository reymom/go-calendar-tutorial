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
	t.Run("Filter Tasks", testFilterTasks)
}

func testGetTasks(t *testing.T) {
	a := assert.New(t)
	testEnv.cleanup()

	list, e := testEnv.ListTasks(context.Background(), model.NewYearlyFilter(uint(time.Now().Year())))
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
		task, e = testEnv.CreateTask(context.Background(), generator.GenerateRandomAddableTask())
		a.NoError(e)
	}
	list, e = testEnv.ListTasks(context.Background(), model.NewYearlyFilter(uint(time.Now().Year())))
	a.NoError(e)
	a.Len(list, 5)
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
	}, *retTask)

}

func testEditTask(t *testing.T) {
	a := assert.New(t)
	testEnv.cleanup()

	retTask, e := testEnv.CreateTask(context.Background(), generator.GenerateRandomAddableTask())
	a.NoError(e)

	task := generator.GenerateRandomAddableTask()
	updatedTask, e := testEnv.EditTask(context.Background(), retTask.TaskId, task)
	a.NoError(e)
	a.NotNil(updatedTask)
	a.NotEqual(retTask, updatedTask)
	a.Equal(model.Task{
		TaskId:      retTask.TaskId,
		Completed:   false,
		AddableTask: task,
	}, *updatedTask)
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

func testFilterTasks(t *testing.T) {
	a := assert.New(t)
	testEnv.cleanup()
	currentDate := time.Now()

	genTask := generator.GenerateRandomAddableTask()
	genTask.StartsAt = currentDate.AddDate(-1, 0, 0)
	_, e := testEnv.CreateTask(context.Background(), genTask)
	a.NoError(e)

	list, e := testEnv.ListTasks(context.Background(), model.NewYearlyFilter(uint(currentDate.Year())))
	a.NoError(e)
	a.Len(list, 0)
	list, e = testEnv.ListTasks(context.Background(), model.NewYearlyFilter(uint(currentDate.Year()-1)))
	a.NoError(e)
	a.Len(list, 1)

	testEnv.cleanup()
	genTask = generator.GenerateRandomAddableTask()
	genTask.StartsAt = currentDate.AddDate(0, -2, 0)
	_, e = testEnv.CreateTask(context.Background(), genTask)
	a.NoError(e)

	list, e = testEnv.ListTasks(context.Background(), model.NewMonthlyFilter(currentDate.Month(), uint(currentDate.Year())))
	a.NoError(e)
	a.Len(list, 0)
	list, e = testEnv.ListTasks(context.Background(), model.NewMonthlyFilter(currentDate.Month()-2, uint(currentDate.Year())))
	a.NoError(e)
	a.Len(list, 1)

	testEnv.cleanup()
	genTask = generator.GenerateRandomAddableTask()
	if currentDate.Weekday() < time.Thursday {
		genTask.StartsAt = currentDate.AddDate(0, 0, 2)
	} else {
		genTask.StartsAt = currentDate.AddDate(0, 0, -2)
	}
	_, e = testEnv.CreateTask(context.Background(), genTask)
	a.NoError(e)
	genTask = generator.GenerateRandomAddableTask()
	if currentDate.Weekday() < time.Thursday {
		genTask.StartsAt = currentDate.AddDate(0, 0, 3)
	} else {
		genTask.StartsAt = currentDate.AddDate(0, 0, -3)
	}
	_, e = testEnv.CreateTask(context.Background(), genTask)
	a.NoError(e)
	genTask = generator.GenerateRandomAddableTask()
	var (
		addWeek int
		day     uint
	)
	if currentDate.Day() > 15 {
		addWeek = -1
		day = uint(currentDate.Day()) - 7
		genTask.StartsAt = currentDate.AddDate(0, 0, -7)
	} else {
		addWeek = 1
		day = uint(currentDate.Day()) + 7
		genTask.StartsAt = currentDate.AddDate(0, 0, 7)
	}
	_, e = testEnv.CreateTask(context.Background(), genTask)
	a.NoError(e)

	year, week := currentDate.ISOWeek()
	list, e = testEnv.ListTasks(context.Background(), model.NewWeeklyFilter(uint(week), uint(year)))
	a.NoError(e)
	a.Len(list, 2)
	list, e = testEnv.ListTasks(context.Background(), model.NewWeeklyFilter(uint(week+addWeek), uint(year)))
	a.NoError(e)
	a.Len(list, 1)

	list, e = testEnv.ListTasks(context.Background(), model.NewDaylyFilter(uint(currentDate.Day()), currentDate.Month(), uint(currentDate.Year())))
	a.NoError(e)
	a.Len(list, 0)

	list, e = testEnv.ListTasks(context.Background(), model.NewDaylyFilter(day, currentDate.Month(), uint(currentDate.Year())))
	a.NoError(e)
	a.Len(list, 1)
}
