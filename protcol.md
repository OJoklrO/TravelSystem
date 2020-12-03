# post params:

`application/x-www-form-urlencoded`



### 相关类型

``` 
flight {
	flightNum string		key
	price int
	numSeats int
	numAvail int
	fromCity string
	arivCity string
}
```

``` 
bus {
	location string			key
	price int
	numBus int
	numAvail int
}
```

``` 
hotel {
	location string			key
	price int
	numRoom int 
	numAvail int
}
```

``` 
customer {
	custID string			key
	custName string
}
```

``` 
reservation {
	custID string										key
	resvType int		(1 flight, 2 hotel, 3 bus)		key
	resvKey string										key
}
```



### /search

|   type   | flight, hotel, bus, customer, resv |
| :------: | :--------------------------------: |
| 过滤内容 |          参考结构参数添加          |
|          |                                    |



### /insert

|     type     | flight, hotel, bus, customer, resv |
| :----------: | :--------------------------------: |
| 参数参考结构 |                                    |
|              |                                    |
|              |                                    |



### /delete

| type                      | flight, hotel, bus, customer, resv |
| ------------------------- | ---------------------------------- |
| 参数参考结构(禁止全空!!!) |                                    |
|                           |                                    |
|                           |                                    |



### /resv

| custID   | string |
| -------- | ------ |
| resvType | 1,2,3  |
| resvKey  | string |
|          |        |

