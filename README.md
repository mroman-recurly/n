# n
***n*** is a collection of missing Go convenience functions reminiscent of Ruby/C#. I love the
elegance of Ruby's short named plethera of chainable methods while C# has some awesome deferred
execution.  Why not stay with Ruby or C# then, well I like Go's ability to generate a single
statically linked binary, Go's concurrancy model, Go's performance, Go's package ecosystem and Go's
tool chain.  This package was created to reduce the friction I had adopting Go as my primary
language of choice.  ***n*** is intended to reduce the coding verbosity required by Go via
convenience functions and the Queryable type.

## Table of Contents
* [Queryable](#queryable)
  * [Iterator Pattern](#iterator-pattern)
  * [Deferred Execution](#deferred-execution)
  * [Types](#types)
  * [Functions](#functions)
  * [Methods](#methods)
  * [Exports](#exports)

# Queryable <a name="queryable"></a>
***Queryable*** provides a way to generically handle Go collection types and provides a plethera
of Ruby like methods to make life a little easier. Since I'm using Reflection to accomplish this it
obviously comes at a cost, which in some cases isn't worth it. However as found in many cases, the
actual work being done far out ways the bookkeeping overhead incurred with the use of reflection.
Other times the speed and convenience of not having to reimplement a Delete or Contains function for
a Slice for the millionth time far out weighs the performance cost.

## Iterator Pattern <a name="iterator-pattern"></a>
Since Queryable is fundamentally based on the notion of iterables, iterating over collections, that
was the first challenge to solve. How do you generically iterate over all primitive Go types.

I implemented the iterator pattern based off of the iterator closure pattern disscussed by blog
https://ewencp.org/blog/golang-iterators/index.htm mainly for syntactic style.  Some
[sources](https://stackoverflow.com/questions/14000534/what-is-most-idiomatic-way-to-create-an-iterator-in-go)
indicate that the closure style iterator is about 3x slower than normal. However my own benchmarking
was much closer coming in at only 33% hit. Even at 3x slower I'm willing to take that hit for
convenience in most case.

Changing the order in which my benchmarks run seems to affect the time (caching?).
At any rate on average I'm seeing only about a 33.33% performance hit. 33% in nested large
data sets may impact performance in some cases but I'm betting in most cases performance
will be dominated by actual work and not looping overhead.

```bash
# 36% slower to use Each function
BenchmarkEach-16               	       1	1732989848 ns/op
BenchmarkArrayIterator-16      	       1	1111445479 ns/op
BenchmarkClosureIterator-16    	       1	1565197326 ns/op

# 25% slower to use Each function
BenchmarkArrayIterator-16      	       1	1210185264 ns/op
BenchmarkClosureIterator-16    	       1	1662226578 ns/op
BenchmarkEach-16               	       1	1622667779 ns/op

# 30% slower to use Each function
BenchmarkClosureIterator-16    	       1	1695826796 ns/op
BenchmarkArrayIterator-16      	       1	1178876072 ns/op
BenchmarkEach-16               	       1	1686159938 ns/op
```

## Deferred Execution <a name="deferred-execution"></a>
I haven't got around to it yet but the intent is there

## Types <a name="types"></a>
***n*** provides a number of types to assis in working with collection types.

| Type         | Description                                                                 |
| ------------ | --------------------------------------------------------------------------- |
| O            | O is an alias for interface{} to reduce verbosity                           |
| Queryable    | Chainable execution and is the heart of algorithm abstration layer          |
| Iterator     | Closure interator pattern implementation                                    |
| KeyVal       | Simple key value pair structure for iterating over map types                |

## Functions <a name="functions"></a>
***n*** provides a number of functions to assist in working with collection types.

| Function     | Description                                     | Slice | Map | Str | Cust |
| ------------ | ----------------------------------------------- | ----- | ----| --- | ---- |
| N            | Creates queryable encapsulating nil             | 1     |     |     |      |
| Q            | Creates queryable encapsulating the given TYPE  | 1     | 1   | 1   | 1    |

## Methods <a name="methods"></a>
Some methods only apply to particular underlying collection types as called out in the table.

**Key: '1' = Implemented, '0' = Not Implemented, 'blank' = Unsupported, Bench nx = slowness**

| Function     | Description                                     | Slice | Map | Str | Cust | Bench |
| ------------ | ----------------------------------------------- | ----- | ----| --- | ---- | ----- |
| Any          | Check if the queryable is not nil and not empty | 1     | 1   | 1   | 1    | 1x    |
| AnyWhere     | Check if any match the given lambda             | 1     | 1   | 1   | 1    | 3x    |
| Append       | Add items to the end of the collection          | 1     |     | 1   | 1    | 10x   |
| At           | Return item at the given neg/pos index notation | 1     |     | 1   | 1    | 1x    |
| Clear        | Clear out the underlying collection             | 1     | 1   | 1   | 1    | 1x    |
| Contains     | Check that all given items are found            | 1     | 1   | 1   | 1    |       |
| ContainsAny  | Check that any given items are found            | 1     | 1   | 1   | 1    |       |
| Copy         | Copy the given obj into this queryable          | 1     | 1   | 1   | 1    | 1x    |
| DeleteAt     | Deletes the item at the given index location    | 1     | 1   | 1   | 1    | 1x    |
| Each         | Iterate over the queryable and execute actions  | 1     | 1   | 1   | 1    | 1.10x |
| Join         | Join slice items as string with given delimiter | 1     |     |     |      |       |
| Len          | Get the length of the collection                | 1     | 1   | 1   | 1    |       |
| Load         | Load Yaml/JSON from file into queryable         |       | 1   |     |      |       |
| Map          | Manipulate the queryable data into a new form   | 1     | 1   | 1   | 1    |       |
| Merge        | Merge other queryables in priority order        | 0     | 0   | 0   | 0    |       |
| Set          | Set the queryable's encapsulated object         | 1     | 1   | 1   | 1    |       |
| TakeFirst    | Remove and return the first item                | 1     |     | 1   | 1    |       |
| TakeFirstCnt | Remove and return the first cnt items           | 0     | 0   | 0   | 0    |       |
| TakeLast     | Remove and return the last item                 | 0     | 0   | 0   | 0    |       |
| TakeLastCnt  | Remove and return the last cnt items            | 0     | 0   | 0   | 0    |       |
| TypeIter     | Is queryable iterable                           | 1     | 1   | 1   | 1    |       |
| TypeMap      | Is queryable reflect.Map                        | 1     | 1   | 1   | 1    |       |
| TypeStr      | Is queryable encapsualting a string             | 1     | 1   | 1   | 1    |       |
| TypeSlice    | Is queryable reflect.Array or reflect.Map       | 1     | 1   | 1   | 1    |       |
| TypeSingle   | Is queryable encapsualting a non-collection     | 1     | 1   | 1   | 1    |       |

## Exports <a name="exports"></a>
Exports process deferred execution and convert the result to a usable external type

| Function     | Description                                     | Return Type               |
| ------------ | ----------------------------------------------- | ------------------------- |
| A            | Export queryable into a string                  | `string`                  |
| I            | Export queryable into an int                    | `int`                     |
| M            | Export queryable to map                         | `map[string]interface{}`  |
| O            | Export queryable to interface{}                 | `interface{}`             |
| S            | Export queryable to a slice of interface{}      | `[]interface{}`           |
| AAMap        | Export queryable to string to string map        | `map[string]string`       |
| Ints         | Export queryable into an int slice              | `[]int`                   |
| Strs         | Export queryable into a string slice            | `[]string`                |

## String Functions <a name="string-functions"></a>
| Function     | Description                                     | Done  |
| ------------ | ----------------------------------------------- | ----- |
| Split        | Split the string into a slice on delimiter      |       |

## Slice Functions
| Function     | Description                                     | Slice | IntSlice | StrSlice | StrMapSlice |
| ------------ | ----------------------------------------------- | ----- | -------- | -------- | ----------- |
| NewTYPE      | Creates a new nub encapsulating the TYPE        | 1     | 1        | 1        | 1           |
| Any          | Check if the slice has anything in it           | 1     | 1        | 1        | 1           |
| AnyWhere     | Match slice items against given lambda          | 0     | 0        | 0        | 0           |
| Append       | Add items to the end of the slice               | 1     | 1        | 1        | 1           |
| At           | Get item using neg/pos index notation           | 0     | 1        | 1        | 1           |
| Clear        | Clear out the underlying slice                  | 0     | 1        | 1        | 1           |
| Contains     | Check if the slice contains the given item      | 0     | 1        | 1        | 1           |
| ContainsAny  | Check if the slice contains any given items     | 0     | 1        | 1        | 1           |
| Count        | Count items that match lambda result            | 0     | 0        | 0        | 0           |
| Del          | Delete item using neg/pos index notation        | 0     | 1        | 1        | 1           |
| DelWhere     | Delete the items that match the given lambda    | 0     | 0        | 0        | 0           |
| Each         | Execute given lambda for each item in slice     | 0     | 0        | 0        | 0           |
| Equals       | Check if the given slice is equal to this slice | 0     | 1        | 1        | 1           |
| Index        | Get the index of the item matchin the given     | 0     | 0        | 0        | 0           |
| Insert       | Insert an item into the underlying slice        | 0     | 0        | 0        | 0           |
| Join         | Join slice items as string with given delimiter | 0     | 1        | 1        | 0           |
| Len          | Get the length of the slice                     | 0     | 1        | 1        | 1           |
| E            | Exports object invoking deferred execution      | 0     | 1        | 1        | 1           |
| Prepend      | Add items to the begining of the slice          | 0     | 1        | 1        | 1           |
| Reverse      | Reverse the items                               | 0     | 0        | 0        | 0           |
| Sort         | Sort the items                                  | 0     | 1        | 1        | 0           |
| Uniq         | Ensure only uniq items exist in the slice       | 0     | 1        | 1        | 0           |
| Where        | Select the items that match the given lambda    | 0     | 0        | 0        | 0           |

## Map Functions
| Function     | Description                                     | IntMap | StrMap | ? |
| ------------ | ----------------------------------------------- | -------- | -------- | ----------- |
| NewTYPE      | Creates a new nub encapsulating the TYPE        | 0        | 1        | 0           |
| Load         | Load Yaml/JSON from file                        | 0        | 1        | 0           |
| Add          | Add a new item to the underlying map            | 0        | 1        | 0           |
| Any          | Check if the map has anything in it             | 0        | 1        | 0           |
| Equals       | Check if the given map is equal to this map     | 0        | 1        | 0           |
| Len          | Get the length of the map                       | 0        | 1        | 0           |
| M            | Exports object invoking deferred execution | 0        | 1        | 0           |
| Merge        | Merge other maps in, in priority order          | 0        | 1        | 0           |
| MergeNub     | Merge other nub maps in, in priority order      | 0        | 1        | 0           |
| Slice        | Get the slice indicated by the multi-key        | 0        | 1        | 0           |
| Str          | Get the str indicated by the multi-key          | 0        | 1        | 0           |
| StrMap       | Get the str map indicated by the multi-key      | 0        | 1        | 0           |
| StrMapByName | Get the str map by name and multi-key           | 0        | 1        | 0           |
| StrMapSlice  | Get the str map slice by the multi-key          | 0        | 1        | 0           |
| StrSlice     | Get the str slice by the multi-key              | 0        | 1        | 0           |
