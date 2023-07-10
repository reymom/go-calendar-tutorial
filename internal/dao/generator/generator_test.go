//go:build unit_test

package generator

import (
	"testing"
	"time"

	"github.com/reymom/go-calendar-tutorial/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestGenerator(t *testing.T) {
	t.Run("Random Task", getRandomTask)
}

func getRandomTask(t *testing.T) {
	a := assert.New(t)
	randomString := getRandomStringByLength(101)
	a.NotEmpty(randomString)
	randomUint8 := getRandomUint8(int(model.LastPriorityTypeId))
	a.Less(randomUint8, (model.LastPriorityTypeId))
	randomHourDuration := getRandomHourDuration(10)
	a.Less(randomHourDuration.Hours(), 10*time.Hour.Hours())
}
