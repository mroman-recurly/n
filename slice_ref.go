package n

import (
	"reflect"
)

// RefSlice implements the Slice interface providing a generic way to work with slice types
// including convenience methods on par with other rapid development languages. This type
// incurs the typical 10x reflection overhead costs. For high performance use the Slice
// implementation matching the type your working with or implement a new type that satisfies
// the Slice interface.
type RefSlice struct {
	k reflect.Kind
	v *reflect.Value
}

// // NewRefSlice uses reflection to encapsulate the given Go slice type inside a new *RefSlice.
// // Expects a Go slice type to be provided and will create an empty *RefSlice if nothing valid
// // is given.
// func NewRefSlice(slice interface{}) (new *RefSlice) {
// 	new = &RefSlice{}
// 	v := reflect.ValueOf(slice)

// 	k := v.Kind()
// 	x, interfaceSliceType := slice.([]interface{})
// 	switch {

// 	// Return the RefSlice.Nil
// 	case k == reflect.Invalid:

// 	// Iterate over array and append
// 	case k == reflect.Array:
// 		for i := 0; i < v.Len(); i++ {
// 			elem := v.Index(i).Interface()
// 			n.Append(item)
// 		}

// 	// Convert []interface to slice of elem type
// 	case interfaceSliceType:
// 		new = NewRefSliceV(x...)

// 	// Slice of distinct type can be used directly
// 	case k == reflect.Slice:
// 		new.o = obj
// 		new.len = v.Len()

// 	// Append single items
// 	default:
// 		n.Append(obj)
// 	}
// 	return
// }

// NewRefSliceV creates a new *RefSlice from the given variadic elements. Always returns
// at least a reference to an empty RefSlice.
func NewRefSliceV(elems ...interface{}) (new *RefSlice) {
	new = &RefSlice{}

	// Return RefSlice.Nil if nothing given
	if len(elems) == 0 {
		return
	}

	// Create new slice from the type of the first non Invalid element
	var slice *reflect.Value
	for i := 0; i < len(elems); i++ {

		// Create target slice from first Valid element
		if slice == nil && reflect.ValueOf(elems[i]).IsValid() {
			typ := reflect.SliceOf(reflect.TypeOf(elems[i]))
			v := reflect.MakeSlice(typ, 0, 10)
			slice = &v
		}

		// Append element to slice
		if slice != nil {
			elem := reflect.ValueOf(elems[i])
			*slice = reflect.Append(*slice, elem)
		}
	}
	if slice != nil {
		new.v = slice
		new.k = slice.Kind()
	}
	return
}

// // Any tests if this Slice is not empty or optionally if it contains
// // any of the given variadic elements. Incompatible types will return false.
// func (p *RefSlice) Any(elems ...interface{}) bool {
// 	if p == nil || len(*p) == 0 {
// 		return false
// 	}

// 	// Not looking for anything
// 	if len(elems) == 0 {
// 		return true
// 	}

// 	// Looking for something specific returns false if incompatible type
// 	for i := range elems {
// 		if x, ok := elems[i].(int); ok {
// 			for j := range *p {
// 				if (*p)[j] == x {
// 					return true
// 				}
// 			}
// 		}
// 	}
// 	return false
// }

// // AnyS tests if this Slice contains any of the given Slice's elements.
// // Incompatible types will return false.
// // Supports RefSlice, *RefSlice, []int or *[]int
// func (p *RefSlice) AnyS(slice interface{}) bool {
// 	if p == nil || len(*p) == 0 {
// 		return false
// 	}
// 	var elems []int
// 	switch x := slice.(type) {
// 	case []int:
// 		elems = x
// 	case *[]int:
// 		elems = *x
// 	case RefSlice:
// 		elems = x
// 	case *RefSlice:
// 		elems = (*x)
// 	}
// 	for i := range elems {
// 		for j := range *p {
// 			if (*p)[j] == elems[i] {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// // AnyW tests if this Slice contains any that match the lambda selector.
// func (p *RefSlice) AnyW(sel func(O) bool) bool {
// 	return p.CountW(sel) != 0
// }

// // Append an element to the end of this Slice and returns a reference to this Slice.
// func (p *RefSlice) Append(elem interface{}) Slice {
// 	if p == nil {
// 		p = NewRefSliceV()
// 	}
// 	if x, ok := elem.(int); ok {
// 		*p = append(*p, x)
// 	}
// 	return p
// }

// // AppendV appends the variadic elements to the end of this Slice and returns a reference to this Slice.
// func (p *RefSlice) AppendV(elems ...interface{}) Slice {
// 	if p == nil {
// 		p = NewRefSliceV()
// 	}
// 	for _, elem := range elems {
// 		p.Append(elem)
// 	}
// 	return p
// }

// // At returns the element at the given index location. Allows for negative notation.
// func (p *RefSlice) At(i int) (elem *Object) {
// 	elem = &Object{}
// 	if p == nil {
// 		return
// 	}
// 	if i = absIndex(len(*p), i); i == -1 {
// 		return
// 	}
// 	elem.o = (*p)[i]
// 	return
// }

// // Clear modifies this Slice to clear out all elements and returns a reference to this Slice.
// func (p *RefSlice) Clear() Slice {
// 	if p == nil {
// 		p = NewRefSliceV()
// 	} else {
// 		p.Drop()
// 	}
// 	return p
// }

// // Concat returns a new Slice by appending the given Slice to this Slice using variadic expansion.
// // Supports RefSlice, *RefSlice, []int or *[]int
// func (p *RefSlice) Concat(slice interface{}) (new Slice) {
// 	return p.Copy().ConcatM(slice)
// }

// // ConcatM modifies this Slice by appending the given Slice using variadic expansion and returns a reference to this Slice.
// // Supports RefSlice, *RefSlice, []int or *[]int
// func (p *RefSlice) ConcatM(slice interface{}) Slice {
// 	if p == nil {
// 		p = NewRefSliceV()
// 	}
// 	switch x := slice.(type) {
// 	case []int:
// 		*p = append(*p, x...)
// 	case *[]int:
// 		*p = append(*p, (*x)...)
// 	case RefSlice:
// 		*p = append(*p, x...)
// 	case *RefSlice:
// 		*p = append(*p, (*x)...)
// 	}
// 	return p
// }

// // Copy returns a new Slice with the indicated range of elements copied from this Slice.
// // Expects nothing, in which case everything is copied, or two indices i and j, in which
// // case positive and negative notation is supported and uses an inclusive behavior such
// // that Slice(0, -1) includes index -1 as opposed to Go's exclusive behavior. Out of
// // bounds indices will be moved within bounds.
// //
// // An empty Slice is returned if indicies are mutually exclusive or nothing can be returned.
// func (p *RefSlice) Copy(indices ...int) (new Slice) {
// 	if p == nil || len(*p) == 0 {
// 		return NewRefSliceV()
// 	}

// 	// Handle index manipulation
// 	i, j, err := absIndices(len(*p), indices...)
// 	if err != nil {
// 		return NewRefSliceV()
// 	}

// 	// Copy elements over to new Slice
// 	x := make([]int, j-i, j-i)
// 	copy(x, (*p)[i:j])
// 	return NewRefSlice(x)
// }

// // Count the number of elements in this Slice equal to the given element.
// func (p *RefSlice) Count(elem interface{}) (cnt int) {
// 	if y, ok := elem.(int); ok {
// 		cnt = p.CountW(func(x O) bool { return ExB(x.(int) == y) })
// 	}
// 	return
// }

// // CountW counts the number of elements in this Slice that match the lambda selector.
// func (p *RefSlice) CountW(sel func(O) bool) (cnt int) {
// 	if p == nil || len(*p) == 0 {
// 		return
// 	}
// 	for i := 0; i < len(*p); i++ {
// 		if sel((*p)[i]) {
// 			cnt++
// 		}
// 	}
// 	return
// }

// // Drop modifies this Slice to delete the indicated range of elements and returns a referece to this Slice.
// // Expects nothing, in which case everything is dropped, or two indices i and j, in which case positive and
// // negative notation is supported and uses an inclusive behavior such that DropAt(0, -1) includes index -1
// // as opposed to Go's exclusive behavior. Out of bounds indices will be moved within bounds.
// func (p *RefSlice) Drop(indices ...int) Slice {
// 	if p == nil || len(*p) == 0 {
// 		return p
// 	}

// 	// Handle index manipulation
// 	i, j, err := absIndices(len(*p), indices...)
// 	if err != nil {
// 		return p
// 	}

// 	// Execute
// 	n := j - i
// 	if i+n < len(*p) {
// 		*p = append((*p)[:i], (*p)[i+n:]...)
// 	} else {
// 		*p = (*p)[:i]
// 	}
// 	return p
// }

// // DropAt modifies this Slice to delete the element at the given index location. Allows for negative notation.
// // Returns a reference to this Slice.
// func (p *RefSlice) DropAt(i int) Slice {
// 	return p.Drop(i, i)
// }

// // DropFirst modifies this Slice to delete the first element and returns a reference to this Slice.
// func (p *RefSlice) DropFirst() Slice {
// 	return p.Drop(0, 0)
// }

// // DropFirstN modifies this Slice to delete the first n elements and returns a reference to this Slice.
// func (p *RefSlice) DropFirstN(n int) Slice {
// 	if n == 0 {
// 		return p
// 	}
// 	return p.Drop(0, abs(n)-1)
// }

// // DropLast modifies this Slice to delete the last element and returns a reference to this Slice.
// func (p *RefSlice) DropLast() Slice {
// 	return p.Drop(-1, -1)
// }

// // DropLastN modifies thi Slice to delete the last n elements and returns a reference to this Slice.
// func (p *RefSlice) DropLastN(n int) Slice {
// 	if n == 0 {
// 		return p
// 	}
// 	return p.Drop(absNeg(n), -1)
// }

// // DropW modifies this Slice to delete the elements that match the lambda selector and returns a reference to this Slice.
// // The slice is updated instantly when lambda expression is evaluated not after DropW completes.
// func (p *RefSlice) DropW(sel func(O) bool) Slice {
// 	if p == nil || len(*p) == 0 {
// 		return p
// 	}
// 	l := len(*p)
// 	for i := 0; i < l; i++ {
// 		if sel((*p)[i]) {
// 			p.DropAt(i)
// 			l--
// 			i--
// 		}
// 	}
// 	return p
// }

// // Each calls the given lambda once for each element in this Slice, passing in that element
// // as a parameter. Returns a reference to this Slice
// func (p *RefSlice) Each(action func(O)) Slice {
// 	if p == nil {
// 		return p
// 	}
// 	for i := 0; i < len(*p); i++ {
// 		action((*p)[i])
// 	}
// 	return p
// }

// // EachE calls the given lambda once for each element in this Slice, passing in that element
// // as a parameter. Returns a reference to this Slice and any error from the lambda.
// func (p *RefSlice) EachE(action func(O) error) (Slice, error) {
// 	var err error
// 	if p == nil {
// 		return p, err
// 	}
// 	for i := 0; i < len(*p); i++ {
// 		if err = action((*p)[i]); err != nil {
// 			return p, err
// 		}
// 	}
// 	return p, err
// }

// // EachI calls the given lambda once for each element in this Slice, passing in the index and element
// // as a parameter. Returns a reference to this Slice
// func (p *RefSlice) EachI(action func(int, O)) Slice {
// 	if p == nil {
// 		return p
// 	}
// 	for i := 0; i < len(*p); i++ {
// 		action(i, (*p)[i])
// 	}
// 	return p
// }

// // EachIE calls the given lambda once for each element in this Slice, passing in the index and element
// // as a parameter. Returns a reference to this Slice and any error from the lambda.
// func (p *RefSlice) EachIE(action func(int, O) error) (Slice, error) {
// 	var err error
// 	if p == nil {
// 		return p, err
// 	}
// 	for i := 0; i < len(*p); i++ {
// 		if err = action(i, (*p)[i]); err != nil {
// 			return p, err
// 		}
// 	}
// 	return p, err
// }

// // EachR calls the given lambda once for each element in this Slice in reverse, passing in that element
// // as a parameter. Returns a reference to this Slice
// func (p *RefSlice) EachR(action func(O)) Slice {
// 	if p == nil {
// 		return p
// 	}
// 	for i := len(*p) - 1; i >= 0; i-- {
// 		action((*p)[i])
// 	}
// 	return p
// }

// // EachRE calls the given lambda once for each element in this Slice in reverse, passing in that element
// // as a parameter. Returns a reference to this Slice and any error from the lambda.
// func (p *RefSlice) EachRE(action func(O) error) (Slice, error) {
// 	var err error
// 	if p == nil {
// 		return p, err
// 	}
// 	for i := len(*p) - 1; i >= 0; i-- {
// 		if err = action((*p)[i]); err != nil {
// 			return p, err
// 		}
// 	}
// 	return p, err
// }

// // EachRI calls the given lambda once for each element in this Slice in reverse, passing in that element
// // as a parameter. Returns a reference to this Slice
// func (p *RefSlice) EachRI(action func(int, O)) Slice {
// 	if p == nil {
// 		return p
// 	}
// 	for i := len(*p) - 1; i >= 0; i-- {
// 		action(i, (*p)[i])
// 	}
// 	return p
// }

// // EachRIE calls the given lambda once for each element in this Slice in reverse, passing in that element
// // as a parameter. Returns a reference to this Slice and any error from the lambda.
// func (p *RefSlice) EachRIE(action func(int, O) error) (Slice, error) {
// 	var err error
// 	if p == nil {
// 		return p, err
// 	}
// 	for i := len(*p) - 1; i >= 0; i-- {
// 		if err = action(i, (*p)[i]); err != nil {
// 			return p, err
// 		}
// 	}
// 	return p, err
// }

// // Empty tests if this Slice is empty.
// func (p *RefSlice) Empty() bool {
// 	if p == nil || len(*p) == 0 {
// 		return true
// 	}
// 	return false
// }

// // First returns the first element in this Slice as Object.
// // Object.Nil() == true will be returned when there are no elements in the slice.
// func (p *RefSlice) First() (elem *Object) {
// 	elem = p.At(0)
// 	return
// }

// // FirstN returns the first n elements in this slice as a Slice reference to the original.
// // Best effort is used such that as many as can be will be returned up until the request is satisfied.
// func (p *RefSlice) FirstN(n int) Slice {
// 	return p.Slice(0, abs(n)-1)
// }

// // Index returns the index of the first element in this Slice where element == elem
// // Returns a -1 if the element was not not found.
// func (p *RefSlice) Index(elem interface{}) (loc int) {
// 	loc = -1
// 	if p == nil || len(*p) == 0 {
// 		return
// 	}
// 	for i := 0; i < len(*p); i++ {
// 		if elem == (*p)[i] {
// 			return i
// 		}
// 	}
// 	return
// }

// // Insert modifies this Slice to insert the given element before the element with the given index.
// // Negative indices count backwards from the end of the slice, where -1 is the last element. If a
// // negative index is used, the given element will be inserted after that element, so using an index
// // of -1 will insert the element at the end of the slice. Slice is returned for chaining. Invalid
// // index locations will not change the slice.
// func (p *RefSlice) Insert(i int, elem interface{}) Slice {
// 	if p == nil || len(*p) == 0 {
// 		return p.Append(elem)
// 	}
// 	j := i
// 	if j = absIndex(len(*p), j); j == -1 {
// 		return p
// 	}
// 	if i < 0 {
// 		j++
// 	}

// 	// Insert the item before j if pos and after j if neg
// 	if x, ok := elem.(int); ok {
// 		if j == 0 {
// 			*p = append([]int{x}, (*p)...)
// 		} else if j < len(*p) {
// 			*p = append(*p, x)
// 			copy((*p)[j+1:], (*p)[j:])
// 			(*p)[j] = x
// 		} else {
// 			*p = append(*p, x)
// 		}
// 	}
// 	return p
// }

// // Join converts each element into a string then joins them together using the given separator or comma by default.
// func (p *RefSlice) Join(separator ...string) (str *Object) {
// 	if p == nil || len(*p) == 0 {
// 		str = &Object{""}
// 		return
// 	}
// 	sep := ","
// 	if len(separator) > 0 {
// 		sep = separator[0]
// 	}

// 	var builder strings.Builder
// 	for i := 0; i < len(*p); i++ {
// 		builder.WriteString((&Object{(*p)[i]}).ToString())
// 		if i+1 < len(*p) {
// 			builder.WriteString(sep)
// 		}
// 	}
// 	str = &Object{builder.String()}
// 	return
// }

// // Last returns the last element in this Slice as an Object.
// // Object.Nil() == true will be returned if there are no elements in the slice.
// func (p *RefSlice) Last() (elem *Object) {
// 	elem = p.At(-1)
// 	return
// }

// // LastN returns the last n elements in this Slice as a Slice reference to the original.
// // Best effort is used such that as many as can be will be returned up until the request is satisfied.
// func (p *RefSlice) LastN(n int) Slice {
// 	return p.Slice(absNeg(n), -1)
// }

// // Len returns the number of elements in this Slice
// func (p *RefSlice) Len() int {
// 	if p == nil {
// 		return 0
// 	}
// 	return len(*p)
// }

// // Less returns true if the element indexed by i is less than the element indexed by j.
// func (p *RefSlice) Less(i, j int) bool {
// 	if p == nil || len(*p) < 2 || i < 0 || j < 0 || i >= len(*p) || j >= len(*p) {
// 		return false
// 	}
// 	return (*p)[i] < (*p)[j]
// }

// Nil tests if this Slice is nil
func (p *RefSlice) Nil() bool {
	if p == nil || p.v == nil {
		return true
	}
	return false
}

// O returns the underlying data structure as is
func (p *RefSlice) O() interface{} {
	if p.Nil() {
		return nil
	}
	return p.v.Interface()
}

// // Pair simply returns the first and second Slice elements as Objects
// func (p *RefSlice) Pair() (first, second *Object) {
// 	first, second = &Object{}, &Object{}
// 	if len(*p) > 0 {
// 		first = p.At(0)
// 	}
// 	if len(*p) > 1 {
// 		second = p.At(1)
// 	}
// 	return
// }

// // Pop modifies this Slice to remove the last element and returns the removed element as an Object.
// func (p *RefSlice) Pop() (elem *Object) {
// 	elem = p.Last()
// 	p.DropLast()
// 	return
// }

// // PopN modifies this Slice to remove the last n elements and returns the removed elements as a new Slice.
// func (p *RefSlice) PopN(n int) (new Slice) {
// 	if n == 0 {
// 		return NewRefSliceV()
// 	}
// 	new = p.Copy(absNeg(n), -1)
// 	p.DropLastN(n)
// 	return
// }

// // Prepend modifies this Slice to add the given element at the begining and returns a reference to this Slice.
// func (p *RefSlice) Prepend(elem interface{}) Slice {
// 	return p.Insert(0, elem)
// }

// // Reverse returns a new Slice with the order of the elements reversed.
// func (p *RefSlice) Reverse() (new Slice) {
// 	if p == nil || len(*p) < 2 {
// 		return p.Copy()
// 	}
// 	return p.Copy().ReverseM()
// }

// // ReverseM modifies this Slice reversing the order of the elements and returns a reference to this Slice.
// func (p *RefSlice) ReverseM() Slice {
// 	if p == nil || len(*p) == 0 {
// 		return p
// 	}
// 	for i, j := 0, len(*p)-1; i < j; i, j = i+1, j-1 {
// 		p.Swap(i, j)
// 	}
// 	return p
// }

// // Select creates a new slice with the elements that match the lambda selector.
// func (p *RefSlice) Select(sel func(O) bool) (new Slice) {
// 	slice := NewRefSliceV()
// 	if p == nil || len(*p) == 0 {
// 		return slice
// 	}
// 	for i := 0; i < len(*p); i++ {
// 		if sel((*p)[i]) {
// 			*slice = append(*slice, (*p)[i])
// 		}
// 	}
// 	return slice
// }

// // Set the element at the given index location to the given element. Allows for negative notation.
// // Returns a reference to this Slice and swallows any errors.
// func (p *RefSlice) Set(i int, elem interface{}) Slice {
// 	slice, _ := p.SetE(i, elem)
// 	return slice
// }

// // SetE the element at the given index location to the given element. Allows for negative notation.
// // Returns a referenc to this Slice and an error if out of bounds or elem is the wrong type.
// func (p *RefSlice) SetE(i int, elem interface{}) (Slice, error) {
// 	var err error
// 	if p == nil {
// 		return p, err
// 	}
// 	if i = absIndex(len(*p), i); i == -1 {
// 		err = errors.Errorf("slice assignment is out of bounds")
// 		return p, err
// 	}

// 	if x, ok := elem.(int); ok {
// 		(*p)[i] = x
// 	} else {
// 		err = errors.Errorf("can't set type '%T' in '%T'", elem, p)
// 	}
// 	return p, err
// }

// // Shift modifies this Slice to remove the first element and returns the removed element as an Object.
// func (p *RefSlice) Shift() (elem *Object) {
// 	elem = p.First()
// 	p.DropFirst()
// 	return
// }

// // ShiftN modifies this Slice to remove the first n elements and returns the removed elements as a new Slice.
// func (p *RefSlice) ShiftN(n int) (new Slice) {
// 	if n == 0 {
// 		return NewRefSliceV()
// 	}
// 	new = p.Copy(0, abs(n)-1)
// 	p.DropFirstN(n)
// 	return
// }

// // Single reports true if there is only one element in this Slice.
// func (p *RefSlice) Single() bool {
// 	return len(*p) == 1
// }

// // Slice returns a range of elements from this Slice as a Slice reference to the original. Allows for negative notation.
// // Expects nothing, in which case everything is included, or two indices i and j, in which case an inclusive behavior
// // is used such that Slice(0, -1) includes index -1 as opposed to Go's exclusive behavior. Out of bounds indices will
// // be moved within bounds.
// //
// // An empty Slice is returned if indicies are mutually exclusive or nothing can be returned.
// //
// // e.g. NewRefSliceV(1,2,3).Slice(0, -1) == [1,2,3] && NewRefSliceV(1,2,3).Slice(1,2) == [2,3]
// func (p *RefSlice) Slice(indices ...int) Slice {
// 	if p == nil || len(*p) == 0 {
// 		return NewRefSliceV()
// 	}

// 	// Handle index manipulation
// 	i, j, err := absIndices(len(*p), indices...)
// 	if err != nil {
// 		return NewRefSliceV()
// 	}

// 	return NewRefSlice((*p)[i:j])
// }

// // Sort returns a new Slice with sorted elements.
// func (p *RefSlice) Sort() (new Slice) {
// 	if p == nil || len(*p) < 2 {
// 		return p.Copy()
// 	}
// 	return p.Copy().SortM()
// }

// // SortM modifies this Slice sorting the elements and returns a reference to this Slice.
// func (p *RefSlice) SortM() Slice {
// 	if p == nil || len(*p) < 2 {
// 		return p
// 	}
// 	sort.Sort(p)
// 	return p
// }

// // SortReverse returns a new Slice sorting the elements in reverse.
// func (p *RefSlice) SortReverse() (new Slice) {
// 	if p == nil || len(*p) < 2 {
// 		return p.Copy()
// 	}
// 	return p.Copy().SortReverseM()
// }

// // SortReverseM modifies this Slice sorting the elements in reverse and returns a reference to this Slice.
// func (p *RefSlice) SortReverseM() Slice {
// 	if p == nil || len(*p) < 2 {
// 		return p
// 	}
// 	sort.Sort(sort.Reverse(p))
// 	return p
// }

// // Returns a string representation of this Slice, implements the Stringer interface
// func (p *RefSlice) String() string {
// 	var builder strings.Builder
// 	builder.WriteString("[")
// 	for i := 0; i < len(*p); i++ {
// 		builder.WriteString(fmt.Sprintf("%d", (*p)[i]))
// 		if i+1 < len(*p) {
// 			builder.WriteString(" ")
// 		}
// 	}
// 	builder.WriteString("]")
// 	return builder.String()
// }

// // Swap modifies this Slice swapping the indicated elements.
// func (p *RefSlice) Swap(i, j int) {
// 	if p == nil || len(*p) < 2 || i < 0 || j < 0 || i >= len(*p) || j >= len(*p) {
// 		return
// 	}
// 	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
// }

// // Take modifies this Slice removing the indicated range of elements from this Slice and returning them as a new Slice.
// // Expects nothing, in which case everything is taken, or two indices i and j, in which case positive and negative
// // notation is supported and uses an inclusive behavior such that Take(0, -1) includes index -1 as opposed to Go's
// // exclusive behavior. Out of bounds indices will be moved within bounds.
// func (p *RefSlice) Take(indices ...int) (new Slice) {
// 	new = p.Copy(indices...)
// 	p.Drop(indices...)
// 	return
// }

// // TakeAt modifies this Slice removing the elemement at the given index location and returns the removed element as an Object.
// // Allows for negative notation.
// func (p *RefSlice) TakeAt(i int) (elem *Object) {
// 	elem = p.At(i)
// 	p.DropAt(i)
// 	return
// }

// // TakeW modifies this Slice removing the elements that match the lambda selector and returns them as a new Slice.
// func (p *RefSlice) TakeW(sel func(O) bool) (new Slice) {
// 	slice := NewRefSliceV()
// 	if p == nil || len(*p) == 0 {
// 		return slice
// 	}
// 	l := len(*p)
// 	for i := 0; i < l; i++ {
// 		if sel((*p)[i]) {
// 			*slice = append(*slice, (*p)[i])
// 			p.DropAt(i)
// 			l--
// 			i--
// 		}
// 	}
// 	return slice
// }

// // Union returns a new Slice by joining uniq elements from this Slice with uniq elements from the given Slice while preserving order.
// // Supports RefSlice, *RefSlice, []int or *[]int
// func (p *RefSlice) Union(slice interface{}) (new Slice) {
// 	return p.Copy().UnionM(slice)
// }

// // UnionM modifies this Slice by joining uniq elements from this Slice with uniq elements from the given Slice while preserving order.
// // Supports RefSlice, *RefSlice, []int or *[]int
// func (p *RefSlice) UnionM(slice interface{}) Slice {
// 	return p.ConcatM(slice).UniqM()
// }

// // Uniq returns a new Slice with all non uniq elements removed while preserving element order.
// // Cost for this call vs the UniqM is roughly the same, this one is appending that one dropping.
// func (p *RefSlice) Uniq() (new Slice) {
// 	if p == nil || len(*p) < 2 {
// 		return p.Copy()
// 	}
// 	m := NewIntMapBool()
// 	slice := NewRefSliceV()
// 	for i := 0; i < len(*p); i++ {
// 		if ok := m.Set((*p)[i], true); ok {
// 			slice.Append((*p)[i])
// 		}
// 	}
// 	return slice
// }

// // UniqM modifies this Slice to remove all non uniq elements while preserving element order.
// // Cost for this call vs the Uniq is roughly the same, this one is dropping that one appending.
// func (p *RefSlice) UniqM() Slice {
// 	if p == nil || len(*p) < 2 {
// 		return p
// 	}
// 	m := NewIntMapBool()
// 	l := len(*p)
// 	for i := 0; i < l; i++ {
// 		if ok := m.Set((*p)[i], true); !ok {
// 			p.DropAt(i)
// 			l--
// 			i--
// 		}
// 	}
// 	return p
// }