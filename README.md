# id3

## Installation
```
go get github.com/florentsorel/id3
```

## Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/florentsorel/id3"
)

func main() {
	tag, err := id3.Open("id3v1_0.mp3")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	if tag.HasID3v1() {
		fmt.Printf("%+v\n", tag.ID3v1)
	}
}
```

### Read specific version
```go
package main

import (
	"fmt"
	"log"

	"github.com/florentsorel/id3/id3v1"
)

func main() {
	id3v1, err := id3v1.Open("id3v1_0.mp3")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	fmt.Printf("%+v\n", id3v1)
}
```
