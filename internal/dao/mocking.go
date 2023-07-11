//go:build mocking_app

package dao

import (
	"context"
	"math/rand"
	"time"

	"github.com/reymom/go-calendar-tutorial/internal/dao/generator"
	"github.com/rs/zerolog/log"
)

func (d *PsqlDao) MockingCleanUp() {
	const query = "TRUNCATE Tasks " +
		"RESTART IDENTITY CASCADE"

	_, e := d.writePool.Exec(context.Background(), query)
	if e != nil {
		log.Err(e).Msg("")
		return
	}
	return
}

func (d *PsqlDao) MockTask(daysInRange uint) error {
	task := generator.GenerateRandomAddableTask()
	days := rand.Intn(int(daysInRange))
	hours := time.Hour * time.Duration(rand.Intn(24))
	task.StartsAt = time.Now().AddDate(0, 0, int(days)).Add(hours)
	task.FinishesAt = task.StartsAt.Add(time.Hour * time.Duration(rand.Intn(12)))

	_, e := d.CreateTask(context.Background(), task)
	return e
}
