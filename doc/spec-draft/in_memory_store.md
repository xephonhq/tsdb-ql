# In Memory Store Draft

In order to use the QL, there must be a backend store, the easiest way is to have
an in memory store, and this might be merged into [xephon-k](https://github.com/xephonhq/xephon-k)
very soon.

Do NOT provide

- RESTful API
- Disk persistent
  - though disk is a very important problem when it comes to storage, we just ignore it for now

Problems to consider

- Golang's garbage collection, though it is also possible to have a C++ library, but it would be too much overhead

## Data structure

I think I could follow the Gorilla paper with a lot more simplification

- no compression (i.e. no delta of delta)
- keep length of the bucket
  - so there could be optimization for count without criteria

Meta map

- `map[string]int[]`
  - [ ] TODO: how to declare int array in go
  - the integers point to the position in the data array

Data array, one for time one for value

- time stamp `int[][]`
- value `int[][]`
