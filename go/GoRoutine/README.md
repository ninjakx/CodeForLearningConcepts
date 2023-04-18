<h3> How to stop a goroutine from running endlessly </h3>


depending on the specific scenario. Here are a few common methods:

    Use a select statement with a channel for cancellation:
    You can use a select statement to wait for values on multiple channels simultaneously, 
    including a channel for cancellation. When a cancellation signal is received on the cancellation channel, 
    you can exit the goroutine. For example:

```go
func worker(cancel <-chan struct{}) {
    for {
        select {
        case <-cancel:
            return // exit the goroutine
        default:
            // do some work
        }
    }
}
```

In this example, the worker function waits for values on two channels: cancel for cancellation and the default case for work to do. If a cancellation signal is received on the cancel channel, the function exits the loop and returns, which terminates the goroutine.

    Use a context.Context for cancellation:
    You can also use a context.Context to manage the lifecycle of a goroutine and cancel it when necessary.
    You can create a context.Context with a cancel function using the context.WithCancel function, 
    and pass the context to the goroutine. When the parent context is cancelled using the cancel function, 
    all child contexts are also cancelled, which can be used to stop the goroutine. For example:


```go
func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            return // exit the goroutine
        default:
            // do some work
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel() // cancel the context when the main function returns
    go worker(ctx)
    // do some other work
    cancel() // cancel the context to stop the worker goroutine
}
```

In this example, the worker function waits for values on two channels: ctx.Done() for cancellation and the default case for work to do. If the context is cancelled by calling the cancel function,  the ctx.Done() channel receives a value, and the function exits the loop and returns, which terminates the goroutine.

    Use a shared boolean variable for cancellation:
    You can also use a shared boolean variable to signal to the goroutine that it should exit. 
    You can set the variable to true to signal that the goroutine should exit, and 
    check the variable periodically in the goroutine to determine if it should exit. For example:

```go
func worker(done <-chan struct{}) {
    for {
        select {
        case <-done:
            return // exit the goroutine
        default:
            // do some work
        }
    }
}

func main() {
    done := make(chan struct{})
    go worker(done)
    // do some other work
    close(done) // signal to the worker goroutine that it should exit
}
```

In this example, the worker function waits for values on two channels: done for cancellation and the default case for work to do. If a value is received on the done channel, the function exits the loop and returns, which terminates the goroutine. The main function signals to the goroutine that it should exit by closing the done channel.
