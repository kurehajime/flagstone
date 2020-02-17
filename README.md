# flagstone

flagstone is a Go library to convert flags to web UI.

![flagstone](https://user-images.githubusercontent.com/4569916/74228957-ce9fab80-4d04-11ea-8eb5-2970496e75c5.png) 

## install 

```
go get github.com/kurehajime/flagstone
```

## usage

1. Use golang's standard flag package.


    ```
    package main

    import (
        "flag"
        "fmt"
    )

    var who *string

    func main() {
        who = flag.String("who", "world", "say hello to ...")
        flag.Parse()

        fmt.Println("hello " + *who + "!")
    }
    ```

1. Import flagstone and add one line.

    ```
    package main

    import (
        "flag"
        "github.com/kurehajime/flagstone" //★ here ★
        "fmt"
    )

    var who *string

    func main() {
        who = flag.String("who", "world", "say hello to ...")
        flag.Parse()

        flagstone.Launch("helloworld", "flagstone sample") //★ here ★

        fmt.Println("hello " + *who + "!")
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

#### SetCSS

Specify CSS.
