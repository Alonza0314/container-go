# stack

## How to use

1. NewStack(): it will return a stack struct.

    ```go
    stack := NewStack()
    ```

2. Len(): it will return how many items in the stack.

    ```go
    length := stack.Len()
    ```

3. IsEmpty(): it will return whether the stack is empty or not.

    ```go
    isEmpty := stack.IsEmpty()
    ```

4. Top(); it will return the top item in the stack, and also an error message if the stack is empty.

    ```go
    topItem, err := stack.Top()
    ```

5. Push(): it will push an item into the stack.

    ```go
    stack.Push(item)
    ```

6. Pop(): it will pop one item from the stack, and also return this item and the error message if necessary.

    ```go
    topItem, err := stack.Pop()
    ```
