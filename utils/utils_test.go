package utils

import (
	"testing"
	"time"
)

func TestAverageMinutes(t *testing.T) {
	averageMinutesTests := []struct {
		name      string
		want      int
		timeSlice []string
	}{
		{name: "Example TestCase 1", want: 241, timeSlice: []string{"12:00 PM, DAY 1", "12:01 PM, DAY 1"}},
		{name: "Example TestCase 2", want: 960, timeSlice: []string{"12:00 AM, DAY 2"}},
		{name: "Example TestCase 3", want: 27239, timeSlice: []string{"02:00 PM, DAY 19", "02:00 PM, DAY 20", "01:58 PM, DAY 20"}},
	}

	for _, tt := range averageMinutesTests {
		t.Run(tt.name, func(t *testing.T) {
			got := AverageMinutes(tt.timeSlice)
			assertIntEquals(t, got, tt.want)
		})
	}
}

func TestDeconstructTimeString(t *testing.T) {
	deconstuctTests := []struct {
		name            string
		want            deconstructedTime
		constructedTime string
	}{
		{name: "Before Noon", want: deconstructedTime{2, 0, "AM", 4}, constructedTime: "02:00 AM, DAY 4"},
		{name: "Noon", want: deconstructedTime{12, 0, "PM", 1}, constructedTime: "12:00 PM, DAY 1"},
		{name: "After Noon", want: deconstructedTime{16, 38, "PM", 12}, constructedTime: "04:38 PM, DAY 12"},
		{name: "Midnight", want: deconstructedTime{0, 0, "AM", 12}, constructedTime: "12:00 AM, DAY 12"},
	}

	for _, tt := range deconstuctTests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := DeconstructTimeString(tt.constructedTime)
			assertStructEquals(t, got, tt.want)
		})
	}

	deconstuctErrorTests := []struct {
		name            string
		want            string
		constructedTime string
	}{
		{name: "Leading Zero Day", want: "invalid timestring expected format: hh:mm xM, DAY n", constructedTime: "02:00 AM, DAY 04"},
		{name: "Single Digit Hour", want: "invalid timestring expected format: hh:mm xM, DAY n", constructedTime: "2:00 PM, DAY 1"},
		{name: "Missing Comma", want: "invalid timestring expected format: hh:mm xM, DAY n", constructedTime: "04:38 PM DAY 12"},
	}

	for _, tt := range deconstuctErrorTests {
		t.Run(tt.name, func(t *testing.T) {
			_, got := DeconstructTimeString(tt.constructedTime)
			assertError(t, got, tt.want)
		})
	}
}

func TestFindDifferenceInMinutes(t *testing.T) {
	got := FindDifferenceInMinutes(time.Date(0, 0, 1, 8, 0, 0, 0, time.UTC), time.Date(0, 0, 1, 12, 0, 0, 0, time.UTC))
	want := 240.00
	assertFloatEquals(t, got, want)
}

func assertFloatEquals(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func assertIntEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertStructEquals(t testing.TB, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v want %+v", got, want)
	}
}

func assertError(t testing.TB, got error, want string) {
	t.Helper()
	if got == nil {
		t.Error("wanted an error but didn't get one")
	}

	if got.Error() != want {
		t.Errorf("got %q want %q", got, want)
	}
}
