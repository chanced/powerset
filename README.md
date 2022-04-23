# powerset

Generic [powerset](https://en.wikipedia.org/wiki/Power_set) algorithm for go

## Install

```bash
go get github.com/chanced/powerset
```

## Usage

```go
package main

func main() {

    fmt.Println(powerset.Compute([]string{"a", "b", "c"}))
    // output: [[] [a] [b] [a b] [c] [a c] [b c] [a b c]]

    fmt.Println(powerset.Compute([]int{1, 2, 3}))
    // output: [[] [1] [2] [1 2] [3] [1 3] [2 3] [1 2 3]]
}
```

## Contributions

Contributions of any kind are always welcome. If there's a problem or there is a
way to improve this algorithm, please feel free to open an issue or pull
request.

## License

MIT
