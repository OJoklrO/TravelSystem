package main

import (
	"github.com/OJoklrO/dbServer/DBConn"
	//"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

const (
	flight = "flight"
	hotel = "hotel"
	bus = "bus"
	customer = "customer"
	resv = "resv"
)

func main() {
	r := gin.Default()

	//r.Use(cors.Default())
	// solve web page serve
	r.Use(static.Serve("/", static.LocalFile("./source/", false)))

	r.GET("/", GetStatic)

	r.POST("/search", Search)
	r.POST("/insert", InsertRow)
	r.POST("/delete", DeleteRow)
	r.POST("/resv", Reservate)

	log.Fatal(r.Run(":80"))
}

func GetStatic(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Search(c *gin.Context) {
	t := c.PostForm("type")
	switch t {
	case flight:
		fs := DBConn.SearchFlight(getFlightParam(c))
		c.JSON(200, fs)
		break
	case hotel:
		hs := DBConn.SearchHotel(getHotelParam(c))
		c.JSON(200, hs)
		break
	case bus:
		bs := DBConn.SearchBus(getBusParam(c))
		c.JSON(200, bs)
		break
	case customer:
		cs := DBConn.SearchCustomer(getCustomerParam(c))
		c.JSON(200, cs)
		break
	case resv:
		rs := DBConn.SearchResv(getResvParam(c))
		c.JSON(200, rs)
		break
	}
}

func InsertRow(c *gin.Context) {
	var res string
	t := c.PostForm("type")
	switch t {
	case flight:
		res = DBConn.InsertFlight(getFlightParam(c))
		break
	case hotel:
		res = DBConn.InsertHotel(getHotelParam(c))
		break
	case bus:
		res = DBConn.InsertBus(getBusParam(c))
		break
	case customer:
		res = DBConn.InsertCustomer(getCustomerParam(c))
		break
	}
	c.String(200, res)
}

func DeleteRow(c *gin.Context) {
	t := c.PostForm("type")
	switch t {
	case flight:
		DBConn.DeleteFlight(getFlightParam(c))
		break
	case hotel:
		DBConn.DeleteHotel(getHotelParam(c))
		break
	case bus:
		DBConn.DeleteBus(getBusParam(c))
		break
	case customer:
		DBConn.DeleteCustomer(getCustomerParam(c))
		break
	case resv:
		DBConn.DeleteResv(getResvParam(c))
		break
	}
}

func Reservate(c *gin.Context) {
	var err error
	t := c.PostForm("resvType")
	switch t {
	case "1":
		err = DBConn.ResvFlight(getResvCall(c))
		break
	case "2":
		err = DBConn.ResvHotel(getResvCall(c))
		break
	case "3":
		err = DBConn.ResvBus(getResvCall(c))
		break
	}
	if err != nil {
		c.String(200, "0")
	}
	c.String(200, "1")
}

func getFlightParam(c *gin.Context) *DBConn.Flight {
	price, _ := strconv.Atoi(c.PostForm("price"))
	numSeats, _ := strconv.Atoi(c.PostForm("numSeats"))
	numAvail, _ := strconv.Atoi(c.PostForm("numAvail"))

	return &DBConn.Flight{
		FlightNum: c.PostForm("flightNum"),
		Price: price    ,
		NumSeats:  numSeats,
		NumAvail:  numAvail,
		FromCity:  c.PostForm("fromCity"),
		ArivCity:  c.PostForm("arivCity"),
	}
}

func getHotelParam(c *gin.Context) *DBConn.Hotel {
	price, _ := strconv.Atoi(c.PostForm("price"))
	numRoom, _ := strconv.Atoi(c.PostForm("numRoom"))
	numAvail, _ := strconv.Atoi(c.PostForm("numAvail"))

	return &DBConn.Hotel{
		Location: c.PostForm("location"),
		Price:    price,
		NumRoom:  numRoom,
		NumAvail: numAvail,
	}
}

func getBusParam(c *gin.Context) *DBConn.Bus {
	price, _ := strconv.Atoi(c.PostForm("price"))
	numBus, _ := strconv.Atoi(c.PostForm("numBus"))
	numAvail, _ := strconv.Atoi(c.PostForm("numAvail"))

	return &DBConn.Bus{
		Location: c.PostForm("location"),
		Price:    price,
		NumBus:  numBus,
		NumAvail: numAvail,
	}
}

func getCustomerParam(c *gin.Context) *DBConn.Customer {
	return &DBConn.Customer{
		CustID: c.PostForm("custID"),
		CustName: c.PostForm("custName"),
	}
}

func getResvParam(c *gin.Context) *DBConn.Reservation {
	resvType, _ := strconv.Atoi(c.PostForm("resvType"))
	return &DBConn.Reservation{
		CustID:   c.PostForm("custID"),
		ResvType: resvType,
		ResvKey:  c.PostForm("resvKey"),
	}
}

func getResvCall(c *gin.Context) (string, string) {
	return c.PostForm("custID"), c.PostForm("resvKey")
}
