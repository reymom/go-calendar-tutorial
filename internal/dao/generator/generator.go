package generator

import (
	"math/rand"
	"time"

	"github.com/reymom/go-calendar-tutorial/internal/dao/mapping"
	"github.com/reymom/go-calendar-tutorial/pkg/model"
)

func GenerateRandomAddableTask() model.AddableTask {
	name, description := generateRandomName(), generateRandomDescription()
	timeRange := generateRandomTimeRange()
	priority, color := generateRandomPriority(), generateRandomColor()

	return model.AddableTask{
		Name:        name,
		Description: description,
		StartsAt:    timeRange[0].Round(time.Millisecond).UTC(),
		FinishesAt:  timeRange[1].Round(time.Millisecond).UTC(),
		Priority:    priority,
		Color:       color,
	}
}

func generateRandomName() string {
	return getRandomStringByLength(uint(rand.Intn(mapping.TaskNameLength)))
}

func generateRandomDescription() string {
	return getRandomStringByLength(uint(rand.Intn(mapping.TaskDescriptionLength)))
}

func generateRandomTimeRange() [2]time.Time {
	startsAt := time.Now().Add(-getRandomHourDuration(24))
	randomDuration := getRandomHourDuration(24)
	return [2]time.Time{startsAt, startsAt.Add(randomDuration)}
}

func generateRandomPriority() model.PriorityTypeId {
	return model.PriorityTypeId(getRandomUint8(int(model.LastPriorityTypeId)))
}

func generateRandomColor() model.ColorId {
	return model.ColorId(getRandomUint8(int(model.LastColorId)))
}
