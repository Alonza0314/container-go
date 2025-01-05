# heap

## How to use

1. NewHeap(): it will return a heap struct, but there is a compare function you need to declare before. It is according to the data type you will store into the heap and the min or max heap you want.
This example is using int as the data type to set up a minHeap (all the child node should be larger than the parent node).

    ```go
    cmpFunc := func(child, parent interface{}) bool {
        // TODO
        return child.(int) > parent.(int)
        // end TODO
    }
    heap := NewHeap(cmpFunc)
    ```

2. Len(): it will return how many items in the heap.

    ```go
    length := heap.Len()
    ```

3. IsEmpty(): it will return whether the heap is empty or not.

    ```go
    isEmpty := heap.IsEmpty()
    ```

4. Top(): it will return the min or the max item from the heap, and also an error message if the heap is empty.

    ```go
    topItem, err := heap.Top()
    ```

5. Push(): it will push an item into the heap.

    ```go
    heap.Push(item)
    ```

6. Pop(): it will pop the minimum item or the maximum item from the heap, and also return this item and the error message if necessary.

    ```go
    topItem, err := heap.Pop()
    ```
