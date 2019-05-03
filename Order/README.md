### Order Module APIs

### APIs to manage Orders

1. Ping the API endpoint

    **Request**
    ```
    GET /orders/ping
    Content-Type: application/json
    ```
    **Response**
    ```
    {
    "vCloud9.0 Burger Order API version 1.0 alive!"
    }
    
2. Get all the orders

3. Get order by the provided Order ID:

    **Request**
    
    ```
    GET /order/{orderId}
    Content-Type: application/json
    ```
    |Parameter	|Type |	Description|
    |-----|-----|------|
    |orderId	|String|Order Id of the placed order|

    **Response Body**

    (Status code: 200)

    |Parameter	|Type	|Description  |
    |----|----|----|
    |orderId |String | Order ID of the order placed |
    |userId | String | Id of user who has placed the order |
    |orderStatus | String  | Status of the order(Active/Placed)|
    |items |Struct | itemId,itemName,itemType,price,description |
    |totalAmount	| float32 | Total amount of the order placed by user|
   

    (Status code: 404)

    |Parameter	|Type |	Description|
    |-----|-----|------|
    |messsage	|String|Error string(Order with provided orderId not found)|
    
4.  Get order by the provided User ID:

    **Request**
    
    ```
    GET /orders/{userId}
    Content-Type: application/json
    ```
    |Parameter	|Type |	Description|
    |-----|-----|------|
    |userId	|String| Id of the user who has placed the order |

    **Response Body**

    (Status code: 200)

    |Parameter	|Type	|Description  |
    |----|----|----|
    |orderId |String | Order ID of the order placed |
    |userId | String | Id of user who has placed the order |
    |orderStatus | String  | Status of the order(Active/Placed)|
    |items |Struct | itemId,itemName,itemType,price,description |
    |totalAmount	| float32 | Total amount of the order placed by user|
   

    (Status code: 404)

    |Parameter	|Type |	Description|
    |-----|-----|------|
    |messsage	|String|Error string(Order with provided userId not found)|
