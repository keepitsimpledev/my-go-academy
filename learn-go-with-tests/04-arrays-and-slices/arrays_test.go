package arraysandslices

import "reflect"
import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("multiple slices", func(t *testing.T) {
		numbersSumIsFifteen := []int{1, 2, 3, 4, 5}
		numbersSumIsEighteen := []int{3, 6, 9}

		got := SumAll(numbersSumIsFifteen, numbersSumIsEighteen)
		want := []int{15, 18}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v. wanted %v. inputs %v, %v", got, want, numbersSumIsFifteen, numbersSumIsEighteen)
		}
	})
}

func TestSumAllTailsMine(t *testing.T) {
	t.Run("multiple slices", func(t *testing.T) {
		tailsSumIsFourteen := []int{1, 2, 3, 4, 5}
		tailsSumIsFifteen := []int{3, 6, 9}
		emptySlice := []int{}

		got := SumAllTails(tailsSumIsFourteen, tailsSumIsFifteen, emptySlice)
		want := []int{14, 15, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v. wanted %v. inputs %v, %v, %v", got, want, tailsSumIsFourteen, tailsSumIsFifteen, emptySlice)
		}
	})
}

func TestSumAllTailsTheirs(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
