## Users

```
Database used - MySQL
Table Name - user
CREATE TABLE user (
    userid varchar(255) NOT NULL,
    fname varchar(255) NOT NULL,
    lname varchar(255),
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    PRIMARY KEY (email)
);
```

## Menu

```
Database used - MongoDB
Database Name - counterburger
Collection Name - menu
ItemId: string 	
ItemName: string 
Price: int	   
Description: string
ItemType: string
```

## Payments

```
Database used - MongoDB
Database Name - payments
Collection Name - payment_details
PaymentID   string
OrderID     string
UserID      string
TotalPrice  float32
OrderStatus bool
PaymentDate string
```