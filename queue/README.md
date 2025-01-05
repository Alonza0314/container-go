# queue

## How to use

1. NewQueue(): it will return a queue struct.

    ```go
    queue := NewQueue()
    ```

2. Len(): it will return how many items in the stack.

    ```go
    length := queue.Len()
    ```

3. IsEmpty(): it will return whether the stack is empty or not.

    ```go
    isEmpty := queue.IsEmpty()
    ```

4. Front(): it will return the first item (i.e. the frontest item) in the queue, and also an error message if the stack is empty.

    ```go
    frontItem, err := queue.Front()
    ```

5. Push(): it will push an item into the stack.

    ```go
    queue.Push(item)
    ```

6. Pop(): it will pop one item from the front of the queue, and also return this item and the error message if necessary.

    ```go
    frontItem, err := queue.Pop()
    ```
