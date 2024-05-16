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
|Queue|"github.com/Alonza0314/Container-Go/Queue"|First In, First Out|
|Heap|"github.com/Alonza0314/Container-Go/Heap"|Priority Queue|
|Set|"github.com/Alonza0314/Container-Go/Set"|Unique Element|

## Use Containers

Please read the README.md file under their paths.

## Something Important

These containers does not need to specify what kind of data type it contains. They are implemented by interface, so be free to use it with your own creativity.
Also, you can store different data type in the same container.

Like:

```go
    container := NewContainer()
    container.Push(314)
    container.Push("Alonza")
    container.Push(true)
```

But be careful that the container-Heap, which can only store one data type since you need to specify the compare function to define which kind of heap you will use (i.e. minHeap or maxHeap).

---

If there is any problem or confused or have a better way to construct container, feel free to contact me.

Although this is a simple module, but I hope this module can bring a convenient way to code for new "coder".
