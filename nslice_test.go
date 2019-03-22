package n

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Slice function
//--------------------------------------------------------------------------------------------------

func ExampleSlice() {
	slice := Slice([]int{1, 2, 3})
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestNSlice_Slice(t *testing.T) {

	// arrays
	var array [2]string
	array[0] = "1"
	array[1] = "2"
	assert.Equal(t, []string{"1", "2"}, Slice(array).O())

	// empty
	assert.Equal(t, nil, Slice(nil).O())
	assert.Equal(t, &NSlice{}, Slice(nil))
	assert.Equal(t, []int{}, Slice([]int{}).O())
	assert.Equal(t, []bool{}, Slice([]bool{}).O())
	assert.Equal(t, []string{}, Slice([]string{}).O())
	assert.Equal(t, []NObj{}, Slice([]NObj{}).O())
	assert.Equal(t, nil, Slice([]interface{}{}).O())

	// pointers
	var obj *NObj
	assert.Equal(t, []*NObj{nil}, Slice(obj).O())
	assert.Equal(t, []*NObj{&(NObj{"bob"})}, Slice(&(NObj{"bob"})).O())
	assert.Equal(t, []*NObj{&(NObj{"1"}), &(NObj{"2"})}, Slice([]*NObj{&(NObj{"1"}), &(NObj{"2"})}).O())

	// interface
	assert.Equal(t, nil, Slice([]interface{}{nil}).O())
	assert.Equal(t, []string{""}, Slice([]interface{}{nil, ""}).O())
	assert.Equal(t, []bool{true}, Slice([]interface{}{true}).O())
	assert.Equal(t, []int{1}, Slice([]interface{}{1}).O())
	assert.Equal(t, []string{""}, Slice([]interface{}{""}).O())
	assert.Equal(t, []string{"bob"}, Slice([]interface{}{"bob"}).O())
	assert.Equal(t, []NObj{{nil}}, Slice([]interface{}{NObj{}}).O())

	// singles
	assert.Equal(t, []int{1}, Slice(1).O())
	assert.Equal(t, []bool{true}, Slice(true).O())
	assert.Equal(t, []string{""}, Slice("").O())
	assert.Equal(t, []string{"1"}, Slice("1").O())
	assert.Equal(t, []NObj{{1}}, Slice(NObj{1}).O())
	assert.Equal(t, []NObj{NObj{"bob"}}, Slice(NObj{"bob"}).O())
	assert.Equal(t, []map[string]string{{"1": "one"}}, Slice(map[string]string{"1": "one"}).O())

	// slices
	assert.Equal(t, []int{1, 2}, Slice([]int{1, 2}).O())
	assert.Equal(t, []bool{true}, Slice([]bool{true}).O())
	assert.Equal(t, []NObj{{"bob"}}, Slice([]NObj{{"bob"}}).O())
	assert.Equal(t, []string{"1", "2"}, Slice([]string{"1", "2"}).O())
	assert.Equal(t, [][]string{{"1"}}, Slice([]interface{}{[]string{"1"}}).O())
	assert.Equal(t, []map[string]string{{"1": "one"}}, Slice([]interface{}{map[string]string{"1": "one"}}).O())
}

// SliceV function
//--------------------------------------------------------------------------------------------------

func ExampleSliceV_empty() {
	slice := SliceV()
	fmt.Println(slice.O())
	// Output: <nil>
}

func ExampleSliceV_variadic() {
	slice := SliceV(1, 2, 3)
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestNSlice_SliceV(t *testing.T) {
	var obj *NObj

	// Arrays
	var array [2]string
	array[0] = "1"
	array[1] = "2"
	assert.Equal(t, [][2]string{array}, SliceV(array).O())

	// Test empty values
	assert.True(t, !SliceV().Any())
	assert.Equal(t, 0, SliceV().Len())
	assert.Equal(t, nil, SliceV().O())
	assert.Equal(t, nil, SliceV(nil).O())
	assert.Equal(t, &NSlice{}, SliceV(nil))
	assert.Equal(t, []string{""}, SliceV(nil, "").O())
	assert.Equal(t, []*NObj{nil}, SliceV(nil, obj).O())

	// Test pointers
	assert.Equal(t, []*NObj{nil}, SliceV(obj).O())
	assert.Equal(t, []*NObj{&(NObj{"bob"})}, SliceV(&(NObj{"bob"})).O())
	assert.Equal(t, []*NObj{nil}, SliceV(obj).O())
	assert.Equal(t, []*NObj{&(NObj{"bob"})}, SliceV(&(NObj{"bob"})).O())
	assert.Equal(t, [][]*NObj{{&(NObj{"1"}), &(NObj{"2"})}}, SliceV([]*NObj{&(NObj{"1"}), &(NObj{"2"})}).O())

	// Singles
	assert.Equal(t, []int{1}, SliceV(1).O())
	assert.Equal(t, []string{"1"}, SliceV("1").O())
	assert.Equal(t, []NObj{NObj{"bob"}}, SliceV(NObj{"bob"}).O())
	assert.Equal(t, []map[string]string{{"1": "one"}}, SliceV(map[string]string{"1": "one"}).O())

	// Multiples
	assert.Equal(t, []int{1, 2}, SliceV(1, 2).O())
	assert.Equal(t, []string{"1", "2"}, SliceV("1", "2").O())
	assert.Equal(t, []NObj{NObj{1}, NObj{2}}, SliceV(NObj{1}, NObj{2}).O())

	// Test slices
	assert.Equal(t, [][]int{{1, 2}}, SliceV([]int{1, 2}).O())
	assert.Equal(t, [][]string{{"1"}}, SliceV([]string{"1"}).O())
}

func TestNSlice_newEmptySlice(t *testing.T) {

	// Array
	var array [2]string
	array[0] = "1"
	assert.Equal(t, []string{}, newEmptySlice(array).O())

	// Singles
	assert.Equal(t, []int{}, newEmptySlice(1).O())
	assert.Equal(t, []bool{}, newEmptySlice(true).O())
	assert.Equal(t, []string{}, newEmptySlice("").O())
	assert.Equal(t, []string{}, newEmptySlice("bob").O())
	assert.Equal(t, []NObj{}, newEmptySlice(NObj{1}).O())

	// Slices
	assert.Equal(t, []int{}, newEmptySlice([]int{1, 2}).O())
	assert.Equal(t, []bool{}, newEmptySlice([]bool{true}).O())
	assert.Equal(t, []string{}, newEmptySlice([]string{"bob"}).O())
	assert.Equal(t, []NObj{}, newEmptySlice([]NObj{{"bob"}}).O())
	assert.Equal(t, [][]string{}, newEmptySlice([]interface{}{[]string{"1"}}).O())
	assert.Equal(t, []map[string]string{}, newEmptySlice([]interface{}{map[string]string{"1": "one"}}).O())

	// Empty slices
	assert.Equal(t, []int{}, newEmptySlice([]int{}).O())
	assert.Equal(t, []bool{}, newEmptySlice([]bool{}).O())
	assert.Equal(t, []string{}, newEmptySlice([]string{}).O())
	assert.Equal(t, []NObj{}, newEmptySlice([]NObj{}).O())

	// Interface types
	assert.Equal(t, []interface{}{}, newEmptySlice(nil).O())
	assert.Equal(t, []interface{}{}, newEmptySlice([]interface{}{nil}).O())
	assert.Equal(t, []int{}, newEmptySlice([]interface{}{1}).O())
	assert.Equal(t, []int{}, newEmptySlice([]interface{}{interface{}(1)}).O())
	assert.Equal(t, []string{}, newEmptySlice([]interface{}{""}).O())
	assert.Equal(t, []NObj{}, newEmptySlice([]interface{}{NObj{}}).O())
}

// Numerably interface methods
//--------------------------------------------------------------------------------------------------
func TestNSlice_O(t *testing.T) {
	assert.Nil(t, SliceV().O())
	assert.Len(t, SliceV().Append("2").O(), 1)
}

func TestNSlice_Any(t *testing.T) {
	assert.False(t, SliceV().Any())
	assert.True(t, SliceV().Append("2").Any())
}

func TestNSlice_Len(t *testing.T) {
	assert.Equal(t, 0, SliceV().Len())
	assert.Equal(t, 1, SliceV().Append("2").Len())
}

func TestNSlice_Nil(t *testing.T) {
	assert.True(t, SliceV().Nil())
	var q *NSlice
	assert.True(t, q.Nil())
	assert.False(t, SliceV().Append("2").Nil())
}

// Append
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_Append_Normal(t *testing.B) {
	ints := []int{}
	for _, i := range Range(0, nines6) {
		ints = append(ints, i)
	}
}

func BenchmarkNSlice_Append_Optimized(t *testing.B) {
	n := &NSlice{o: []int{}}
	for _, i := range Range(0, nines6) {
		n.Append(i)
	}
}

func BenchmarkNSlice_Append_Reflect(t *testing.B) {
	n := &NSlice{o: []NObj{}}
	for _, i := range Range(0, nines6) {
		n.Append(NObj{i})
	}
}

func ExampleNSlice_Append() {
	slice := SliceV(1).Append(2).Append(3)
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestNSlice_Append_Reflect(t *testing.T) {

	// Use a custom type to invoke reflection
	n := SliceV(NObj{"1"})
	assert.Equal(t, 1, n.Len())
	assert.Equal(t, false, n.Nil())
	assert.Equal(t, []NObj{{"1"}}, n.O())

	// Append another to it
	n.Append(NObj{"2"})
	assert.Equal(t, 2, n.Len())
	assert.Equal(t, []NObj{{"1"}, {"2"}}, n.O())

	// Given an invalid type which will abort the function so put at end
	defer func() {
		err := recover()
		assert.Equal(t, "reflect.Set: value of type int is not assignable to type n.NObj", err)
	}()
	n.Append(2)
}

func TestNSlice_Append(t *testing.T) {

	// Append one back to back
	{
		n := SliceV()
		assert.Equal(t, 0, n.Len())
		assert.Equal(t, true, n.Nil())

		// First append invokes 10x reflect overhead because the slice is nil
		n.Append("1")
		assert.Equal(t, 1, n.Len())
		assert.Equal(t, []string{"1"}, n.O())

		// Second append another which will be 2x at most
		n.Append("2")
		assert.Equal(t, 2, n.Len())
		assert.Equal(t, []string{"1", "2"}, n.O())
	}

	// Start with just appending without chaining
	{
		n := SliceV()
		assert.Equal(t, 0, n.Len())
		n.Append(1)
		assert.Equal(t, []int{1}, n.O())
		n.Append(2)
		assert.Equal(t, []int{1, 2}, n.O())
	}

	// Start with nil not chained
	{
		n := SliceV()
		assert.Equal(t, 0, n.Len())
		n.Append(1).Append(2).Append(3)
		assert.Equal(t, 3, n.Len())
		assert.Equal(t, []int{1, 2, 3}, n.O())
	}

	// Start with nil chained
	{
		n := SliceV().Append(1).Append(2)
		assert.Equal(t, 2, n.Len())
		assert.Equal(t, []int{1, 2}, n.O())
	}

	// Start with non nil
	{
		n := SliceV(1).Append(2).Append(3)
		assert.Equal(t, 3, n.Len())
		assert.Equal(t, []int{1, 2, 3}, n.O())
	}

	// Use append result directly
	{
		n := SliceV(1)
		assert.Equal(t, 1, n.Len())
		assert.Equal(t, []int{1, 2}, n.Append(2).O())
	}

	// Test all supported types
	{
		// bool
		{
			n := SliceV(true)
			assert.Equal(t, []bool{true, false}, n.Append(false).O())
			assert.Equal(t, 2, n.Len())
		}

		// int
		{
			n := SliceV(0)
			assert.Equal(t, []int{0, 1}, n.Append(1).O())
			assert.Equal(t, 2, n.Len())
		}

		// string
		{
			n := SliceV("0")
			assert.Equal(t, []string{"0", "1"}, n.Append("1").O())
			assert.Equal(t, 2, n.Len())
		}

		// Append to a slice of custom type i.e. reflection
		{
			n := Slice([]NObj{{"3"}})
			assert.Equal(t, []NObj{{"3"}, {"1"}}, n.Append(NObj{"1"}).O())
			assert.Equal(t, 2, n.Len())
		}
	}

	// Invalid type which will abort the function so put at end
	{
		n := SliceV("1")
		defer func() {
			err := recover()
			assert.Equal(t, "can't insert type 'int' into '[]string'", err)
		}()
		n.Append(2)
	}
}

// AppendV
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_AppendV_Normal(t *testing.B) {
	ints := []int{}
	ints = append(ints, Range(0, nines6)...)
}

func BenchmarkNSlice_AppendV_Optimized(t *testing.B) {
	n := &NSlice{o: []int{}}
	new := rangeO(0, nines6)
	n.AppendV(new...)
}

func BenchmarkNSlice_AppendV_Reflect(t *testing.B) {
	n := &NSlice{o: []NObj{}}
	new := rangeNObjO(0, nines6)
	n.AppendV(new...)
}

func ExampleNSlice_AppendV() {
	slice := SliceV(1).AppendV(2, 3)
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestNSlice_AppendV(t *testing.T) {

	// Append many ints
	{
		n := SliceV(1)
		assert.Equal(t, []int{1, 2, 3}, n.AppendV(2, 3).O())
	}

	// Append many strings
	{
		{
			n := SliceV()
			assert.Equal(t, 0, n.Len())
			assert.Equal(t, []string{"1", "2", "3"}, n.AppendV("1", "2", "3").O())
			assert.Equal(t, 3, n.Len())
		}
		{
			n := Slice([]string{"1"})
			assert.Equal(t, 1, n.Len())
			assert.Equal(t, []string{"1", "2", "3"}, n.AppendV("2", "3").O())
			assert.Equal(t, 3, n.Len())
		}
	}

	// Append to a slice of custom type
	{
		n := Slice([]NObj{{"3"}})
		assert.Equal(t, []NObj{{"3"}, {"1"}}, n.AppendV(NObj{"1"}).O())
		assert.Equal(t, []NObj{{"3"}, {"1"}, {"2"}, {"4"}}, n.AppendV(NObj{"2"}, NObj{"4"}).O())
	}

	// Test all supported types
	{
		// bool
		{
			n := SliceV(true)
			assert.Equal(t, []bool{true, false}, n.AppendV(false).O())
			assert.Equal(t, 2, n.Len())
		}

		// int
		{
			n := SliceV(0)
			assert.Equal(t, []int{0, 1}, n.AppendV(1).O())
			assert.Equal(t, 2, n.Len())
		}

		// string
		{
			n := SliceV("0")
			assert.Equal(t, []string{"0", "1"}, n.AppendV("1").O())
			assert.Equal(t, 2, n.Len())
		}

		// Append to a slice of custom type i.e. reflection
		{
			n := Slice([]NObj{{"3"}})
			assert.Equal(t, []NObj{{"3"}, {"1"}}, n.AppendV(NObj{"1"}).O())
			assert.Equal(t, 2, n.Len())
		}
	}

	// Append to a slice of map
	{
		n := SliceV(map[string]string{"1": "one"})
		expected := []map[string]string{
			{"1": "one"},
			{"2": "two"},
		}
		assert.Equal(t, expected, n.AppendV(map[string]string{"2": "two"}).O())
	}
}

// AppendS
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_AppendS_Normal10(t *testing.B) {
	dest := []int{}
	src := Range(0, nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest = append(dest, (src[j:i])...)
		j = i
	}
}

func BenchmarkNSlice_AppendS_Normal100(t *testing.B) {
	dest := []int{}
	src := Range(0, nines6)
	j := 0
	for i := 100; i < len(src); i += 100 {
		dest = append(dest, (src[j:i])...)
		j = i
	}
}

func BenchmarkNSlice_AppendS_Optimized19(t *testing.B) {
	dest := &NSlice{o: []int{}}
	src := Range(0, nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest.AppendS(src[j:i])
		j = i
	}
}

func BenchmarkNSlice_AppendS_Optimized100(t *testing.B) {
	dest := &NSlice{o: []int{}}
	src := Range(0, nines6)
	j := 0
	for i := 100; i < len(src); i += 100 {
		dest.AppendS(src[j:i])
		j = i
	}
}

func BenchmarkNSlice_AppendS_Reflect10(t *testing.B) {
	dest := &NSlice{o: []NObj{}}
	src := rangeNObj(0, nines6)
	j := 0
	for i := 10; i < len(src); i += 10 {
		dest.AppendS(src[j:i])
		j = i
	}
}

func BenchmarkNSlice_AppendS_Reflect100(t *testing.B) {
	dest := &NSlice{o: []NObj{}}
	src := rangeNObj(0, nines6)
	j := 0
	for i := 100; i < len(src); i += 100 {
		dest.AppendS(src[j:i])
		j = i
	}
}

func ExampleNSlice_AppendS() {
	slice := SliceV(1).AppendS([]int{2, 3})
	fmt.Println(slice.O())
	// Output: [1 2 3]
}

func TestNSlice_AppendS(t *testing.T) {

	// Append many ints
	{
		n := SliceV(1)
		assert.Equal(t, []int{1, 2, 3}, n.AppendS([]int{2, 3}).O())
	}

	// Append many strings
	{
		{
			n := SliceV()
			assert.Equal(t, 0, n.Len())
			assert.Equal(t, []string{"1", "2", "3"}, n.AppendS([]string{"1", "2", "3"}).O())
			assert.Equal(t, 3, n.Len())
		}
		{
			n := Slice([]string{"1"})
			assert.Equal(t, 1, n.Len())
			assert.Equal(t, []string{"1", "2", "3"}, n.AppendS([]string{"2", "3"}).O())
			assert.Equal(t, 3, n.Len())
		}
	}

	// Append to a slice of custom type
	{
		n := Slice([]NObj{{"3"}})
		assert.Equal(t, []NObj{{"3"}, {"1"}}, n.AppendS([]NObj{{"1"}}).O())
		assert.Equal(t, []NObj{{"3"}, {"1"}, {"2"}, {"4"}}, n.AppendS([]NObj{{"2"}, {"4"}}).O())
	}

	// Append to a slice of map
	{
		n := SliceV(map[string]string{"1": "one"})
		expected := []map[string]string{
			{"1": "one"},
			{"2": "two"},
		}
		assert.Equal(t, expected, n.AppendS([]map[string]string{{"2": "two"}}).O())
	}

	// Test all supported types
	{
		// bool
		{
			n := SliceV(true)
			assert.Equal(t, []bool{true, false}, n.AppendS([]bool{false}).O())
			assert.Equal(t, 2, n.Len())
		}

		// int
		{
			n := SliceV(0)
			assert.Equal(t, []int{0, 1}, n.AppendS([]int{1}).O())
			assert.Equal(t, 2, n.Len())
		}

		// string
		{
			n := SliceV("0")
			assert.Equal(t, []string{"0", "1"}, n.AppendS([]string{"1"}).O())
			assert.Equal(t, 2, n.Len())
		}

		// Append to a slice of custom type i.e. reflection
		{
			n := Slice([]NObj{{"3"}})
			assert.Equal(t, []NObj{{"3"}, {"1"}}, n.AppendS([]NObj{{"1"}}).O())
			assert.Equal(t, 2, n.Len())
		}
	}
}

// At
//--------------------------------------------------------------------------------------------------
func BenchmarkNSlice_At_Normal(t *testing.B) {
	ints := Range(0, nines6)
	for i := range ints {
		assert.IsType(t, 0, ints[i])
	}
}

func BenchmarkNSlice_At_Optimized(t *testing.B) {
	src := Range(0, nines6)
	slice := Slice(src)
	for _, i := range src {
		_, ok := (slice.At(i).O()).(int)
		assert.True(t, ok)
	}
}

func BenchmarkNSlice_At_Reflect(t *testing.B) {
	src := rangeNObj(0, nines6)
	slice := Slice(src)
	for i := range src {
		_, ok := (slice.At(i).O()).(NObj)
		assert.True(t, ok)
	}
}

func ExampleNSlice_At() {
	slice := SliceV(1).AppendS([]int{2, 3})
	fmt.Println(slice.At(2).O())
	// Output: 3
}

func TestNSlice_At(t *testing.T) {

	// strings
	{
		slice := SliceV("1", "2", "3", "4")
		assert.Equal(t, "4", slice.At(-1).O())
		assert.Equal(t, "3", slice.At(-2).O())
		assert.Equal(t, "2", slice.At(-3).O())
		assert.Equal(t, "1", slice.At(0).O())
		assert.Equal(t, "2", slice.At(1).O())
		assert.Equal(t, "3", slice.At(2).O())
		assert.Equal(t, "4", slice.At(3).O())
	}
	{
		slice := SliceV("1")
		assert.Equal(t, "1", slice.At(-1).O())
	}
}

// Clear
//--------------------------------------------------------------------------------------------------

func TestQSlice_Clear(t *testing.T) {
	slice := SliceV("1", "2", "3", "4")
	assert.Equal(t, 4, slice.Len())
	slice.Clear()
	assert.Equal(t, 0, slice.Len())
	slice.Clear()
	assert.Equal(t, 0, slice.Len())
}

// // func TestStrSliceAnyContain(t *testing.T) {
// // 	assert.True(t, S("one", "two", "three").AnyContain("thr"))
// // 	assert.False(t, S("one", "two", "three").AnyContain("2"))
// // }

// // func TestStrSliceContains(t *testing.T) {
// // 	assert.True(t, S("1", "2", "3").Contains("2"))
// // 	assert.False(t, S("1", "2", "3").Contains("4"))
// // }

// // func TestStrSliceContainsAny(t *testing.T) {
// // 	assert.True(t, S("1", "2", "3").ContainsAny([]string{"2"}))
// // 	assert.False(t, S("1", "2", "3").ContainsAny([]string{"4"}))
// // }

// // func TestStrSliceDel(t *testing.T) {
// // 	{
// // 		// Pos: delete invalid
// // 		slice := S("0", "1", "2")
// // 		ok := slice.Del(3)
// // 		assert.False(t, ok)
// // 		assert.Equal(t, []string{"0", "1", "2"}, slice.S())
// // 	}
// // 	{
// // 		// Pos: delete last
// // 		slice := S("0", "1", "2")
// // 		ok := slice.Del(2)
// // 		assert.True(t, ok)
// // 		assert.Equal(t, []string{"0", "1"}, slice.S())
// // 	}
// // 	{
// // 		// Pos: delete middle
// // 		slice := S("0", "1", "2")
// // 		ok := slice.Del(1)
// // 		assert.True(t, ok)
// // 		assert.Equal(t, []string{"0", "2"}, slice.S())
// // 	}
// // 	{
// // 		// delete first
// // 		slice := S("0", "1", "2")
// // 		ok := slice.Del(0)
// // 		assert.True(t, ok)
// // 		assert.Equal(t, []string{"1", "2"}, slice.S())
// // 	}
// // 	{
// // 		// Neg: delete invalid
// // 		slice := S("0", "1", "2")
// // 		ok := slice.Del(-4)
// // 		assert.False(t, ok)
// // 		assert.Equal(t, []string{"0", "1", "2"}, slice.S())
// // 	}
// // 	{
// // 		// Neg: delete last
// // 		slice := S("0", "1", "2")
// // 		ok := slice.Del(-1)
// // 		assert.True(t, ok)
// // 		assert.Equal(t, []string{"0", "1"}, slice.S())
// // 	}
// // 	{
// // 		// Neg: delete middle
// // 		slice := S("0", "1", "2")
// // 		ok := slice.Del(-2)
// // 		assert.True(t, ok)
// // 		assert.Equal(t, []string{"0", "2"}, slice.S())
// // 	}
// // }

// // func TestStrSliceDrop(t *testing.T) {
// // 	{
// // 		slice := S().Append("1", "2", "3").Drop(3)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // 	{
// // 		slice := S().Append("1", "2", "3").Drop(5)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // 	{
// // 		slice := S().Drop(3)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // 	{
// // 		slice := S().Append("1", "2", "3").Drop(1)
// // 		assert.Equal(t, []string{"2", "3"}, slice.S())
// // 	}
// // 	{
// // 		slice := S().Append("1", "2", "3").Drop(2)
// // 		assert.Equal(t, []string{"3"}, slice.S())
// // 	}
// // 	{
// // 		slice := S().Append("1", "2", "3").Drop(0)
// // 		assert.Equal(t, []string{"1", "2", "3"}, slice.S())
// // 	}
// // }

// // func TestStrSliceEquals(t *testing.T) {
// // 	{
// // 		slice := S().Append("1", "2", "3")
// // 		target := S().Append("1", "2", "3")
// // 		assert.True(t, slice.Equals(target))
// // 	}
// // 	{
// // 		slice := S().Append("1", "2", "4")
// // 		target := S().Append("1", "2", "3")
// // 		assert.False(t, slice.Equals(target))
// // 	}
// // 	{
// // 		slice := S().Append("1", "2", "3", "4")
// // 		target := S().Append("1", "2", "3")
// // 		assert.False(t, slice.Equals(target))
// // 	}
// // }

// // func TestStrSliceFirst(t *testing.T) {
// // 	assert.Equal(t, A(""), S().First())
// // 	assert.Equal(t, A("1"), S("1").First())
// // 	assert.Equal(t, A("1"), S("1", "2").First())
// // 	assert.Equal(t, "foo", A("foo::").Split("::").First().A())
// // 	{
// // 		// Test that the original slice wasn't modified
// // 		q := S("1")
// // 		assert.Equal(t, []string{"1"}, q.S())
// // 		assert.Equal(t, A("1"), q.First())
// // 		assert.Equal(t, []string{"1"}, q.S())
// // 	}
// // }

// // func TestStrSliceJoin(t *testing.T) {
// // 	assert.Equal(t, "", S().Join(".").A())
// // 	assert.Equal(t, "1", S("1").Join(".").A())
// // 	assert.Equal(t, "1.2", S("1", "2").Join(".").A())
// // }

// // func TestStrSliceLen(t *testing.T) {
// // 	assert.Equal(t, 0, S().Len())
// // 	assert.Equal(t, 1, S("1").Len())
// // 	assert.Equal(t, 2, S("1", "2").Len())
// // }

// // func TestStrSliceLast(t *testing.T) {
// // 	assert.Equal(t, A(""), S().Last())
// // 	assert.Equal(t, A("1"), S("1").Last())
// // 	assert.Equal(t, A("2"), S("1", "2").Last())
// // 	assert.Equal(t, "foo", A("::foo").Split("::").Last().A())
// // 	{
// // 		// Test that the original slice wasn't modified
// // 		q := S("1")
// // 		assert.Equal(t, []string{"1"}, q.S())
// // 		assert.Equal(t, A("1"), q.Last())
// // 		assert.Equal(t, []string{"1"}, q.S())
// // 	}
// // }

// // func TestStrSlicePrepend(t *testing.T) {
// // 	slice := S().Prepend("1")
// // 	assert.Equal(t, "1", slice.At(0))

// // 	slice.Prepend("2", "3")
// // 	assert.Equal(t, "2", slice.At(0))
// // 	assert.Equal(t, []string{"2", "3", "1"}, slice.S())
// // }

// // func TestStrSliceSort(t *testing.T) {
// // 	slice := S().Append("b", "d", "a")
// // 	assert.Equal(t, []string{"a", "b", "d"}, slice.Sort().S())
// // }

// // func TestStrSliceSlice(t *testing.T) {
// // 	assert.Equal(t, S(), S().Slice(0, -1))
// // 	assert.Equal(t, S(""), S("").Slice(0, -1))
// // 	assert.Equal(t, S("1", "2", "3"), S("1", "2", "3").Slice(0, -1))
// // 	assert.Equal(t, S("1", "2"), S("1", "2", "3").Slice(0, -2))
// // 	assert.Equal(t, S("1"), S("1", "2", "3").Slice(0, -3))
// // 	assert.Equal(t, S(), S("1", "2", "3").Slice(0, -4))
// // 	assert.Equal(t, S("2", "3"), S("1", "2", "3").Slice(1, -1))
// // 	assert.Equal(t, S("3"), S("1", "2", "3").Slice(2, -1))
// // 	assert.Equal(t, S(), S("1", "2", "3").Slice(3, -1))
// // 	assert.Equal(t, S(), S("1", "2", "3").Slice(5, -1))
// // 	assert.Equal(t, S("2", "3"), S("1", "2", "3").Slice(1, 2))
// // 	assert.Equal(t, S(), S("1", "2", "3").Slice(3, 2))
// // 	{
// // 		// old FirstCnt ops
// // 		assert.Equal(t, S(), S().Slice(0, 2))
// // 		assert.Equal(t, S("1"), S("1").Slice(0, 2))
// // 		assert.Equal(t, S("1", "2"), S("1", "2").Slice(0, 2))
// // 		assert.Equal(t, S("1", "2", "3"), S("1", "2", "3").Slice(0, 2))
// // 		assert.Equal(t, S("", "foo", "bar"), A("/foo/bar/one").Split("/").Slice(0, 2))
// // 		assert.Equal(t, A("/foo/bar"), A("/foo/bar/one").Split("/").Slice(0, 2).Join("/"))
// // 		{
// // 			// Test that the original slice wasn't modified
// // 			q := S("1")
// // 			assert.Equal(t, []string{"1"}, q.S())
// // 			assert.Equal(t, S("1"), q.Slice(0, 1))
// // 			assert.Equal(t, []string{"1"}, q.S())
// // 		}
// // 	}
// // 	{
// // 		// old LastCnt(2) tests
// // 		assert.Equal(t, S(), S().Slice(-3, -1))
// // 		assert.Equal(t, S("1"), S("1").Slice(-2, -1))
// // 		assert.Equal(t, S("1", "2"), S("1", "2").Slice(-2, -1))
// // 		assert.Equal(t, S("2", "3"), S("1", "2", "3").Slice(-2, -1))
// // 		assert.Equal(t, S("bar", "one"), A("/foo/bar/one").Split("/").Slice(-2, -1))
// // 		assert.Equal(t, A("bar/one"), A("/foo/bar/one").Split("/").Slice(-2, -1).Join("/"))
// // 		{
// // 			// Test that the original slice wasn't modified
// // 			q := S("1")
// // 			assert.Equal(t, []string{"1"}, q.S())
// // 			assert.Equal(t, S("1"), q.Slice(-2, -1))
// // 			assert.Equal(t, []string{"1"}, q.S())
// // 		}
// // 	}
// // }

// // func TestStrSliceTakeFirst(t *testing.T) {
// // 	{
// // 		slice := S("0", "1", "2")
// // 		results := []string{}
// // 		expected := []string{"0", "1", "2"}
// // 		for item, ok := slice.TakeFirst(); ok; item, ok = slice.TakeFirst() {
// // 			results = append(results, item)
// // 		}
// // 		assert.Equal(t, expected, results)
// // 	}
// // 	{
// // 		slice := S("0", "1", "2")
// // 		item, ok := slice.TakeFirst()
// // 		assert.True(t, ok)
// // 		assert.Equal(t, "0", item)
// // 		assert.Equal(t, []string{"1", "2"}, slice.S())
// // 	}
// // 	{
// // 		slice := S("0")
// // 		item, ok := slice.TakeFirst()
// // 		assert.True(t, ok)
// // 		assert.Equal(t, "0", item)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // 	{
// // 		slice := S()
// // 		item, ok := slice.TakeFirst()
// // 		assert.False(t, ok)
// // 		assert.Equal(t, "", item)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // }

// // func TestStrSliceTakeFirstCnt(t *testing.T) {
// // 	{
// // 		slice := S("0", "1", "2")
// // 		items := slice.TakeFirstCnt(2).S()
// // 		assert.Equal(t, []string{"0", "1"}, items)
// // 		assert.Equal(t, []string{"2"}, slice.S())
// // 	}
// // 	{
// // 		slice := S("0", "1", "2")
// // 		items := slice.TakeFirstCnt(3).S()
// // 		assert.Equal(t, []string{"0", "1", "2"}, items)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // 	{
// // 		slice := S("0", "1", "2")
// // 		items := slice.TakeFirstCnt(4).S()
// // 		assert.Equal(t, []string{"0", "1", "2"}, items)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // }

// // func TestStrSliceTakeLast(t *testing.T) {
// // 	{
// // 		slice := S("0", "1", "2")
// // 		results := []string{}
// // 		expected := []string{"2", "1", "0"}
// // 		for item, ok := slice.TakeLast(); ok; item, ok = slice.TakeLast() {
// // 			results = append(results, item)
// // 		}
// // 		assert.Equal(t, expected, results)
// // 	}
// // 	{
// // 		slice := S("0", "1", "2")
// // 		item, ok := slice.TakeLast()
// // 		assert.True(t, ok)
// // 		assert.Equal(t, "2", item)
// // 		assert.Equal(t, []string{"0", "1"}, slice.S())
// // 	}
// // 	{
// // 		slice := S("0")
// // 		item, ok := slice.TakeLast()
// // 		assert.True(t, ok)
// // 		assert.Equal(t, "0", item)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // 	{
// // 		slice := S()
// // 		item, ok := slice.TakeLast()
// // 		assert.False(t, ok)
// // 		assert.Equal(t, "", item)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // }
// // func TestStrSliceTakeLastCnt(t *testing.T) {
// // 	{
// // 		slice := S("0", "1", "2")
// // 		items := slice.TakeLastCnt(2).S()
// // 		assert.Equal(t, []string{"1", "2"}, items)
// // 		assert.Equal(t, []string{"0"}, slice.S())
// // 	}
// // 	{
// // 		slice := S("0", "1", "2")
// // 		items := slice.TakeLastCnt(3).S()
// // 		assert.Equal(t, []string{"0", "1", "2"}, items)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // 	{
// // 		slice := S("0", "1", "2")
// // 		items := slice.TakeLastCnt(4).S()
// // 		assert.Equal(t, []string{"0", "1", "2"}, items)
// // 		assert.Equal(t, []string{}, slice.S())
// // 	}
// // }

// // func TestStrSliceUniq(t *testing.T) {
// // 	{
// // 		data := S().Uniq().S()
// // 		expected := []string{}
// // 		assert.Equal(t, expected, data)
// // 	}
// // 	{
// // 		data := S("1", "2", "3").Uniq().S()
// // 		expected := []string{"1", "2", "3"}
// // 		assert.Equal(t, expected, data)
// // 	}
// // 	{
// // 		data := S("1", "2", "2", "3").Uniq().S()
// // 		expected := []string{"1", "2", "3"}
// // 		assert.Equal(t, expected, data)
// // 	}
// // }

// // func TestYamlPair(t *testing.T) {
// // 	{
// // 		k, v := A("foo=bar").Split("=").YamlPair()
// // 		assert.Equal(t, "foo", k)
// // 		assert.Equal(t, "bar", v)
// // 	}
// // 	{
// // 		k, v := A("=bar").Split("=").YamlPair()
// // 		assert.Equal(t, "", k)
// // 		assert.Equal(t, "bar", v)
// // 	}
// // 	{
// // 		k, v := A("bar=").Split("=").YamlPair()
// // 		assert.Equal(t, "bar", k)
// // 		assert.Equal(t, "", v)
// // 	}
// // 	{
// // 		k, v := A("").Split("=").YamlPair()
// // 		assert.Equal(t, "", k)
// // 		assert.Equal(t, nil, v)
// // 	}
// // }
// // func TestYamlKeyVal(t *testing.T) {
// // 	{
// // 		pair := A("foo=bar").Split("=").YamlKeyVal()
// // 		assert.Equal(t, "foo", pair.Key)
// // 		assert.Equal(t, "bar", pair.Val)
// // 	}
// // 	{
// // 		pair := A("=bar").Split("=").YamlKeyVal()
// // 		assert.Equal(t, "", pair.Key)
// // 		assert.Equal(t, "bar", pair.Val)
// // 	}
// // 	{
// // 		pair := A("bar=").Split("=").YamlKeyVal()
// // 		assert.Equal(t, "bar", pair.Key)
// // 		assert.Equal(t, "", pair.Val)
// // 	}
// // 	{
// // 		pair := A("").Split("=").YamlKeyVal()
// // 		assert.Equal(t, "", pair.Key)
// // 		assert.Equal(t, "", pair.Val)
// // 	}
// // }