# set

## How to use

1. NewSet(): it will return a set struct.

    ```go
    set := NewSet()
    ```

2. Len(): it will return how many items in the set.

    ```go
    length := set.Len()
    ```

3. IsEmpty(): it will return whether the set is empty or not.

    ```go
    isEmpty := set.IsEmpty()
    ```

4. Find(): it will return true if the item is in the set; otherwise, return false.

    ```go
    isFind := set.Find(item)
    ```

5. Push(): it will push an item into the set, and also return error if the item has already in the set.

    ```go
    err := set.Push(item)
    ```

6. Pop(): it will pop the destination item from the set and return error if the set doesn't contain this item.

    ```go
    err := set.Pop(item)
    ```
