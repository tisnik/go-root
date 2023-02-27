go test -bench=BenchmarkBuildStringUsingConcatenation             -benchtime=1000000x -cpuprofile concatenation.out
go test -bench=BenchmarkBuildStringUsingStringBuffer              -benchtime=1000000x -cpuprofile string_buffer.out
go test -bench=BenchmarkBuildStringUsingStringBuilder             -benchtime=1000000x -cpuprofile string_builder.out
go test -bench=BenchmarkBuildStringUsingPreallocatedStringBuffer  -benchtime=1000000x -cpuprofile preallocated_string_buffer.out
go test -bench=BenchmarkBuildStringUsingPreallocatedStringBuilder -benchtime=1000000x -cpuprofile preallocated_string_builder.out
