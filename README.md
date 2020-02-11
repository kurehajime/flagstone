# flagstone

flagstone is a Go library to convert flags to web UI.

![flagstone](https://user-images.githubusercontent.com/4569916/74206771-2831b700-4cc0-11ea-87fc-e88c9b23261c.png) 

## install 

```
go get github.com/kurehajime
```

## usage

1. Use golang's standard flag package.


    ```
    package main

    import (
        "flag"
        "fmt"
    )

    var message *string

    func main() {
        message = flag.String("message", "world", "say hello to ...")
        flag.Parse()

        fmt.Println("hello " + *message + "!")
    }
    ```

1. Import flagstone and add one line.

    ```
    package main

    import (
        "flag"
        "flagstone" //★ here ★
        "fmt"
    )

    var message *string

    func main() {
        message = flag.String("message", "world", "say hello to ...")
        flag.Parse()

        flagstone.Launch("helloworld", "flagstone sample") //★ here ★

        fmt.Println("hello " + *message + "!")
    }

    ```

1. Finish!

    When you run the program, the browser starts.

    Pressing the "Run" button sets the flag and runs the program.


    ![screenshot](https://user-images.githubusercontent.com/4569916/74208613-ac3b6d00-4cc7-11ea-9f3c-e686874f2e38.png)

## options

#### SetPort

Specify the port number.

#### SetSubURL


Specify sub URL.
If not specified, it is determined randomly.


#### SetSort

Specify the order of the parameters.
If not specified, it will be in ABC order.

#### SetSilent

If set to true, flagstone will not produce any extra output.

#### SetUseNonFlagArgs

If true is set, non-flag arguments are added.
non-flag arguments can be obtained by the return value of the Lanch method.