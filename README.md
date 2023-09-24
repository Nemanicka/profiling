# Usage

To run tests just execute the following:

```
  go run bbst_prof.go > bench.csv
```

and then:

```
  python visualize.py
```

That's it! 

# Some explanations...

1. The granularity level of data frame is hardcoded withing go program.
2. What profiler is used? It's insane (I'm not professional golang dev as of now), but I didn't find **simple** golang profilr. The ffprof seems like to much of work to me - it makes it possible to track the performance and other stats only if you write additional Benchmark functions as wrappers, put them is a specific test files and so on... Actually I was thinking about using some binary analyzers, but the optimized code structure might not reflect what I want to see. So the really easiset way was to track time and memory myself, which is really very elegant.
