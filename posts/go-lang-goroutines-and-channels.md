# go async() - goroutines in five minutes!

Well before I started working with Go, I've heard of goroutines. It was supposed to be some kind of magical solution to create asynchronous applications. They said that all you need to do is type "go" before function call. And you know what? They were right! It's so simple that it takes one short blog post to explain enough for you to be able to create your first async Go application.

### Why concurrency in Go?

Go is a perfect language for creating microservices, or at least small service-like applications. If you think about writing a pretty small application, you are most likely talking about doing one task, that is probably related to disk I/O, database connection of something like this. Generally chances are, that you are going to use some time-consuming resources. This is when you realize, that making this happen synchronously is a terrible thing to do. Let's take a look on an example where you want to ask database for some data, but at the same time you need to perform some other tasks:

    // main.go
    ...
    func main() {
        dbResult := database.MakeQuery() // can take eg. 5 seconds
        fmt.Println(dbResult)
        someRepetitiveTask(100)
    }

    func someRepetitiveTask(n int) {
        for i := 0; i < n; i++ {
            fmt.Printf("Hello, Bored Person #%d!\n", i)
            time.Sleep(200 * time.Millisecond)
        }
    }

And the output will be:

    ...
    Hello, Bored Peron #96
    Hello, Bored Peron #97
    Hello, Bored Peron #98
    Hello, Bored Peron #99
    --> Database query response <--

As you can see and imagine, we start by making a call to the database, which might take a couple seconds. For this long we cannot proceed with our code. Now, if our repetitive task does not rely directly on database response (and it doesn't), we could be doing something while waiting for the response.

### First goroutine

Writing goroutines in Go is as simple as prefixing your function call with `go`. And that's it - just do it and the call will be executed in the background! Let's try doing just that:

    // main.go
    ...
    func main() {
    fmt.Println("Hello world!")

    go func() {
    dbResult := database.MakeQuery()
    fmt.Println(dbResult)
    }()

    someRepetitiveTask(100)
    }
    ...

And the output will be:

    ...
    Hello, Bored Person #23!
    Hello, Bored Person #24!
    --> Database query response <--
    Hello, Bored Person #25!
    Hello, Bored Person #26!
    Hello, Bored Person #27!
    ...

As you can see, there is so little that we had to do, but it's enough! Our function is actually called in the background!

Source code for this example is available [on GitHub](https://github.com/mycodesmells/go-tutorial). Note that you should clone repository into $GOPATH/src/github.com/mycodesmells/go-tutorial directory.