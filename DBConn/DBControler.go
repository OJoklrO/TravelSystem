package DBConn

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	DB *gorm.DB
	err error
)

const connstring = "root:1234@tcp(localhost:3307)/TravelSystem026?charset=utf8"

func init() {
	DB, err = gorm.Open(mysql.Open(connstring), &gorm.Config{})
	if err != nil {
		log.Fatal("can't connect to mysql db")
	}
}

type Flight struct {
	FlightNum string    `gorm:"type:varchar(100);primary_key" json:"flightNum"`
	Price     int       `gorm:"type:int" json:"price"`
	NumSeats  int       `gorm:"type:int" json:"numSeats"`
	NumAvail  int       `gorm:"type:int" json:"numAvail"`
	FromCity  string    `gorm:"type:varchar(100)" json:"fromCity"`
	ArivCity  string    `gorm:"type:varchar(100)" json:"arivCity"`
}

type Bus struct {
	Location    string  `gorm:"type:varchar(100);primary_key" json:"location"`
	Price       int     `gorm:"type:int" json:"price"`
	NumBus      int     `gorm:"type:int" json:"numBus"`
	NumAvail    int     `gorm:"type:int" json:"numAvail"`
}

type Hotel struct {
	Location    string  `gorm:"type:varchar(100);primary_key" json:"location"`
	Price       int     `gorm:"type:int" json:"price"`
	NumRoom     int     `gorm:"type:int" json:"numRoom"`
	NumAvail    int     `gorm:"type:int" json:"numAvail"`
}

type Customer struct {
	CustID      string  `gorm:"type:varchar(100);primary_key" json:"custID"`
	CustName    string  `gorm:"type:varchar(100)" json:"custName"`
}

type Reservation struct {
	CustID      string  `gorm:"type:varchar(100);primary_key" json:"custID"`
	ResvType    int     `gorm:"type:int;primary_key" json:"resvType"`
	ResvKey     string  `gorm:"type:varchar(100);primary_key" json:"resvKey"`
}

// flight

func SearchFlight(fliter *Flight) (fs []Flight) {
	DB.Model(&Flight{}).Where(fliter).Find(&fs)
	return
}

func InsertFlight(f *Flight) string {
	var temp Flight
	DB.Model(&Flight{}).Where(&Flight{
		FlightNum: f.FlightNum,
	}).Find(&temp)
	if temp.FlightNum != "" {
		return "0"
	}
	DB.Model(&Flight{}).Create(&f)
	return "1"
}

func DeleteFlight(fliter *Flight) {
	DB.Model(&Flight{}).Delete(fliter)
}

func ResvFlight(custId, flightNum string) (err error) {
	err = DB.Transaction(func(tx *gorm.DB) (err error) {
		err = errors.New("error")
		var (
			f Flight
			c Customer
			resv []Reservation
		)
		tx.Model(&Flight{}).Where(&Flight{
			FlightNum: flightNum,
		}).Find(&f)

		if f.FlightNum == "" || f.NumAvail <= 0 {
			return
		}

		tx.Model(&Customer{}).Where(&Customer{
			CustID: custId,
		}).Find(&c)

		if c.CustID == "" {
			return
		}

		tx.Model(&Reservation{}).Where(&Reservation{
			CustID: custId,
			ResvType: 1,
		}).Find(&resv)

		for _, r := range resv {
			if r.ResvKey == flightNum {
				return
			}
		}
		
		tx.Model(&f).Select("num_avail").Updates(map[string]interface{}{"num_avail": f.NumAvail - 1})
		tx.Model(&Reservation{}).Create(&Reservation{
			CustID:   c.CustID,
			ResvType: 1,
			ResvKey:  f.FlightNum,
		})

		return nil
	})

	return
}

// hotel

func SearchHotel(fliter *Hotel) (hs []Hotel) {
	DB.Model(&Hotel{}).Where(fliter).Find(&hs)
	return
}

func InsertHotel(h *Hotel) string {
	var temp Hotel
	DB.Model(&Hotel{}).Where(&Hotel{
		Location: h.Location,
	}).Find(&temp)
	if temp.Location != "" {
		return "0"
	}
	DB.Model(&Hotel{}).Create(&h)
	return "1"
}

func DeleteHotel(fliter *Hotel) {
	DB.Model(&Hotel{}).Delete(fliter)
}

func ResvHotel(custId, location string) (err error) {
	err = DB.Transaction(func(tx *gorm.DB) (err error) {
		err = errors.New("error")
		var (
			h Hotel
			c Customer
			resv []Reservation
		)
		tx.Model(&Hotel{}).Where(&Hotel{
			Location: location,
		}).Find(&h)

		if h.Location == "" || h.NumAvail <= 0 {
			return
		}

		tx.Model(&Customer{}).Where(&Customer{
			CustID: custId,
		}).Find(&c)

		if c.CustID == "" {
			return
		}

		tx.Model(&Reservation{}).Where(&Reservation{
			CustID: custId,
			ResvType: 2,
		}).Find(&resv)

		for _, r := range resv {
			if r.ResvKey == location {
				return
			}
		}

		tx.Model(&h).Select("num_avail").Updates(map[string]interface{}{"num_avail": h.NumAvail - 1})
		tx.Model(&Reservation{}).Create(&Reservation{
			CustID:   c.CustID,
			ResvType: 2,
			ResvKey:  h.Location,
		})

		return nil
	})

	return
}

// bus

func SearchBus(fliter *Bus) (bs []Bus) {
	DB.Model(&Bus{}).Where(fliter).Find(&bs)
	return
}

func InsertBus(b *Bus) string {
	var temp Bus
	DB.Model(&Bus{}).Where(&Bus{
		Location: b.Location,
	}).Find(&temp)
	if temp.Location != "" {
		return "0"
	}
	DB.Model(&Bus{}).Create(&b)
	return "1"
}

func DeleteBus(fliter *Bus) {
	DB.Model(&Bus{}).Delete(fliter)
}

func ResvBus(custId, location string) (err error) {
	err = DB.Transaction(func(tx *gorm.DB) (err error) {
		err = errors.New("error")
		var (
			b Bus
			c Customer
			resv []Reservation
		)
		tx.Model(&Bus{}).Where(&Bus{
			Location: location,
		}).Find(&b)

		if b.Location == "" || b.NumAvail <= 0 {
			return
		}

		tx.Model(&Customer{}).Where(&Customer{
			CustID: custId,
		}).Find(&c)

		if c.CustID == "" {
			return
		}

		tx.Model(&Reservation{}).Where(&Reservation{
			CustID: custId,
			ResvType: 3,
		}).Find(&resv)

		for _, r := range resv {
			if r.ResvKey == location {
				return
			}
		}

		tx.Model(&b).Select("num_avail").Updates(map[string]interface{}{"num_avail": b.NumAvail - 1})
		tx.Model(&Reservation{}).Create(&Reservation{
			CustID:   c.CustID,
			ResvType: 3,
			ResvKey:  b.Location,
		})

		return nil
	})

	return
}

// reservation

func SearchResv(fliter *Reservation) (rs []Reservation) {
	DB.Model(&Reservation{}).Where(fliter).Find(&rs)
	return
}

func DeleteResv(resv *Reservation) {
	DB.Model(&Reservation{}).Delete(&resv)
}

// customer

func SearchCustomer(fliter *Customer) (cs []Customer) {
	DB.Model(&Customer{}).Where(fliter).Find(&cs)
	return
}

func InsertCustomer(c *Customer) string {
	var temp Customer
	DB.Model(&Customer{}).Where(&Customer{
		CustID: c.CustID,
	}).Find(&temp)
	if temp.CustID != "" {
		return "0"
	}
	DB.Model(&Customer{}).Create(&c)
	return "1"
}

func DeleteCustomer(fliter *Customer) {
	DB.Model(&Customer{}).Delete(fliter)
}
