package mapping

import (
	"math"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/reymom/go-calendar-tutorial/pkg/model"
)

func uuidFromPGUuid(pgUuid pgtype.UUID) (string, error) {
	if pgUuid.Valid == false {
		return "", model.ErrUuidNotPresent
	}
	parsedUuid, e := uuid.FromBytes(pgUuid.Bytes[:])
	if e != nil {
		return "", e
	}
	return parsedUuid.String(), nil
}

func uint8FromPgInt2[T interface{ ~uint8 }](i pgtype.Int2) (T, error) {
	if i.Valid == false {
		return 0, model.ErrInt2NotPresent
	}
	if i.Int16 < 0 || i.Int16 > math.MaxUint8 {
		return 0, model.ErrInt2NotCastableToUint8
	}
	return T(i.Int16), nil
}

func toPgInt2[T interface {
	~uint8 | ~int16 | int8 | uint16
}](i T) (pgtype.Int2, error) {
	if int64(i) > math.MaxInt16 {
		return pgtype.Int2{}, model.ErrOutOfBoundsOfPGInt2
	}
	return pgtype.Int2{
		Int16: int16(i),
		Valid: true,
	}, nil
}

func trimSpacesOfString(s string) string {
	return strings.Trim(s, " ")
}

func stringFromPGVarchar(pgText pgtype.Text) (string, error) {
	if pgText.Valid == false {
		return "", model.ErrTextNotPresent
	}
	return pgText.String, nil
}

func pgTextFromString(s string, length uint) (pgtype.Text, error) {
	adaptedString := trimSpacesOfString(s)
	if len(adaptedString) > int(length) {
		return pgtype.Text{}, model.ErrTextToLong
	}
	return pgtype.Text{
		String: adaptedString,
		Valid:  true,
	}, nil
}

func timeFromPgTimestamptz(pgT pgtype.Timestamptz) (time.Time, error) {
	if pgT.Valid == false {
		return time.Time{}, model.ErrTimestamptzNotPresent
	}
	return pgT.Time.UTC(), nil
}

func pgTimestamptzFromTime(t time.Time) (pgtype.Timestamptz, error) {
	return pgtype.Timestamptz{
		Time:             t.UTC(),
		Valid:            true,
		InfinityModifier: pgtype.Finite,
	}, nil
}

func boolFromPgBool(pgB pgtype.Bool) (bool, error) {
	if pgB.Valid == false {
		return false, model.ErrBoolNotPresent
	}
	return pgB.Bool, nil
}
