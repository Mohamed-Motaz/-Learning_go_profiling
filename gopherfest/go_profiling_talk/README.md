# Go Profiling and Optimization

This code is example code used in the talk [Go Profiling and Optimization](https://docs.google.com/presentation/d/1n6bse0JifemG7yve0Bb0ZAC-IWhTQjCNAclblnn2ANY/edit#slide=id.g3a3e2af65_029).

It shows how pprof and [go-torch](https://github.com/uber/go-torch) can be
used to identify performance bottlenecks, and optimize them.


# Commands

Load testing: 
- go-wrk -d 5 http://localhost:9090/simple

Create a pprof file without benchmarking:
- go tool pprof --seconds 5 localhost:9090/debug/pprof/profile

View full UI:
- go tool pprof -http=":8000" {pprof_location}

Benchmark:
- go test -bench . -benchmem -cpuprofile prof.cpu -memprofile prof.mem -blockprofile prof.block

Profile the mem (num of objects):
- go tool pprof -alloc_objects stats.test prof.mem

Profile the mem (size of objects):
- go tool pprof -alloc_space stats.test prof.mem