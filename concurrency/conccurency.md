# Go Routine Concurrency Demonstration

This program demonstrates **concurrency** using goroutines in Go.

## 🔍 What It Does

- `say(label string)` prints a message 5 times, once every 100ms.
- The function is invoked once using a goroutine (`go say("world")`) and once directly (`say("hello")`).
- We log the **iteration number** and the **timestamp** of each print to observe the scheduling pattern.

## 🧠 Key Concepts

### Goroutines

- Lightweight threads managed by the Go runtime.
- Spawned using `go <function>()`.
- Run **concurrently** (not necessarily in parallel).

### WaitGroup

- Ensures the main program waits for all goroutines to finish.
- `Add(1)` registers a new goroutine.
- `Done()` marks a goroutine as finished.
- `Wait()` blocks until all goroutines are done.

### Concurrency vs Parallelism

- **Concurrency**: Multiple tasks in progress, overlapping in execution.
- **Parallelism**: Multiple tasks running **simultaneously** on separate CPU cores.
- This code runs concurrently, and on multi-core machines, it **may** also execute in parallel.

## 🧪 Example Output

```txt
hello [0] at 12:00:00.123
hello [1] at 12:00:00.223
hello [2] at 12:00:00.323
hello [3] at 12:00:00.423
hello [4] at 12:00:00.523
world [0] at 12:00:00.125
world [1] at 12:00:00.225
...
```
