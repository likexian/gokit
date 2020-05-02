# GoKit - xlump

Lump kits for Golang development.

## Installation

    go get -u github.com/likexian/gokit

## Importing

    import (
        "github.com/likexian/gokit/xlump"
    )

## How it work

    //             .---------.                   .---------.                    .---------.
    // Task.Add -> |         |                   |         |                    |         |
    // Task.Add -> |         |    Worker.Work    |         |                    |         |
    // Task.Add -> | TaskIn  | -> Worker.Work -> | TaskOut | -> Merger.Merge -> | TaskSum |
    // Task.Add -> |         |    Worker.Work    |         |                    |         |
    // Task.Add -> |         |                   |         |                    |         |
    //             '---------'                   '---------'                    '---------'

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/gokit/xlump)

## Example

### Plus 1 to every number and sum to total

```go
// Worker: Plus 1 to every number
mathPlus := func(t xlump.Task) xlump.Task {
    return t.(int) + 1
}

// Merger: Sum to total
mathSum := func(r xlump.Task, t xlump.Task) xlump.Task {
    return r.(int) + t.(int)
}

// New a work queue
wq := xlump.New(100)
// Set Worker func
wq.SetWorker(mathPlus, 10)
// Set Merger func
wq.SetMerger(mathSum, 0)

// Add number to queue
for i := 0; i < 1000; i++ {
    wq.Add(i)
}

// Wait for result and print
result := wq.Wait()
fmt.Println("sum is:", result)
```

## License

Copyright 2012-2020 [Li Kexian](https://www.likexian.com/)

Licensed under the Apache License 2.0

## Donation

If this project is helpful, please share it with friends.

If you want to thank me, you can [give me a cup of coffee](https://www.likexian.com/donate/).
