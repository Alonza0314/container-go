# Container-Go

## Version

These containers are coded under golang version:

> go version go1.21.5 linux/amd64.

But I think they can be compiled under all versions.

## Import Containers

Here include these containers

|Type|How to import|Propoties|
|-|-|-|
|Stack|"github.com/Alonza0314/Container-Go/Stack"|Last In, First Out|

## Use Containers

Please read the README.md file under their paths.

## Something Important

These containers does not need to specify what kind of data type it contains. They are implemented by interface, so be free to use it with your own creativity.
Also, you can store different data type in the same container.

Like:

```go
    container := NewContainer()
    container.Push(0314)
    container.Push("Alonza")
    container.Push(true)
```

---

If there is any problem or confused or have a better way to construct container, feel free to contact me.

Although this is a simple module, but I hope this module can bring a convenient way to code for new "coder".
