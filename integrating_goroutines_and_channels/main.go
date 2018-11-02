package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	/////////////////////////////////
	/* Mutex Lock with Goroutines and Channels

	runtime.GOMAXPROCS(4)

	f, _ := os.Create("./log.txt")
	f.Close()

	logCh := make(chan string, 50)

	go func() {
		for {
			msg, ok := <-logCh
			if ok {
				f, _ := os.OpenFile("./log.txt", os.O_APPEND, os.ModeAppend)

				logTime := time.Now().Format(time.RFC3339)
				f.WriteString(logTime + " - " + msg)
				f.Close()
			} else {
				break
			}
		}
	}()

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			go func(i, j int) {
				msg := fmt.Sprintf("%d + %d = %d\n", i, j, i+j)
				logCh <- msg
				fmt.Print(msg)
			}(i, j)
		}
	}

	fmt.Scanln()

	*/

	//////////////////////////////////////////////
	/* Simulating Events

	btn := MakeButton()

	handlerOne := make(chan string)
	handlerTwo := make(chan string)

	btn.AddEventListener("click", handlerOne)
	btn.AddEventListener("click", handlerTwo)

	go func() {
		for {
			msg := <-handlerOne
			fmt.Println("Handler One: " + msg)
		}
	}()

	go func() {
		for {
			msg := <-handlerTwo
			fmt.Println("Handler Two: " + msg)
		}
	}()

	btn.TriggerEvent("click", "Button clicked!")
	btn.RemoveEventListener("click", handlerTwo)
	btn.TriggerEvent("click", "Button clicked again!")

	fmt.Scanln()

	*/

	////////////////////////////////////////////////////////
	/* Simulating Callbacks

	po := new(PurchaseOrder)
	po.Value = 42.27

	ch := make(chan *PurchaseOrder)

	go SavePO(po, ch)

	newPo := <-ch
	fmt.Printf("PO: %v", newPo)

	*/

	/*
		//////////////////////// Simulating Promises /////////////////////////
		po := new(PurchaseOrder)
		po.Value = 42.27
		SavePO(po, false).Then(func(obj interface{}) error {
			po := obj.(*PurchaseOrder)
			fmt.Printf("Purchase Order saved with ID: %d\n", po.Number)
			return nil
		}, func(err error) {
			fmt.Printf("Failed to save Purchase Order: " + err.Error() + "\n")
		})

		fmt.Scanln()

	*/

	/*
		//////////////////////////// The Pipe and Filter Pattern //////////////////////////

		runtime.GOMAXPROCS(4)
		ch := make(chan int)
		go generate(ch)
		for {
			prime := <-ch
			fmt.Println(prime)
			ch1 := make(chan int)
			go filter(ch, ch1, prime)
			ch = ch1
		}
	*/

	/*  */
	//////////////////////////////// Extract, Transform and Load (ETL) /////////////////////////////

	start := time.Now()

	extractChannel := make(chan *Order)
	transformChannel := make(chan *Order)
	doneChannel := make(chan bool)

	go extract(extractChannel)
	go transform(extractChannel, transformChannel)
	go load(transformChannel, doneChannel)

	<-doneChannel

	fmt.Println(time.Since(start))

}

//////////////////////////////// Extract, Transform and Load (ETL) /////////////////////////////
type Product struct {
	PartNumber string
	UnitCost   float64
	UnitPrice  float64
}

type Order struct {
	CustomerNumber int
	PartNumber     string
	Quantity       int

	UnitCost  float64
	UnitPrice float64
}

func extract(ch chan *Order) {

	f, _ := os.Open("./orders.txt")
	defer f.Close()
	r := csv.NewReader(f)

	for record, err := r.Read(); err == nil; record, err = r.Read() {
		order := new(Order)
		order.CustomerNumber, _ = strconv.Atoi(record[0])
		order.PartNumber = record[1]
		order.Quantity, _ = strconv.Atoi(record[2])
		ch <- order
	}

	close(ch)
}

func transform(extractChannel, transformChannel chan *Order) {
	f, _ := os.Open("./productList.txt")
	defer f.Close()
	r := csv.NewReader(f)

	records, _ := r.ReadAll()
	productList := make(map[string]*Product)
	for _, record := range records {
		product := new(Product)
		product.PartNumber = record[0]
		product.UnitCost, _ = strconv.ParseFloat(record[1], 64)
		product.UnitPrice, _ = strconv.ParseFloat(record[2], 64)
		productList[product.PartNumber] = product
	}

	numMessage := 0

	for o := range extractChannel {
		numMessage++
		go func(o *Order) {
			time.Sleep(3 * time.Millisecond)
			o.UnitCost = productList[o.PartNumber].UnitCost
			o.UnitPrice = productList[o.PartNumber].UnitPrice
			transformChannel <- o
			numMessage--
		}(o)

		for numMessage > 0 {
			time.Sleep(1 * time.Millisecond)
		}

	}

	close(transformChannel)
}

func load(transformChannel chan *Order, doneChannel chan bool) {
	f, _ := os.Create("./dest.txt")
	defer f.Close()

	fmt.Fprintf(f, "%20s%15s%12s%12s%15s%15s\n",
		"Part Number", "Quantity", "Unit Cost", "UnitPrice", "Total Cost", "Total Price")

	numMessage := 0

	for o := range transformChannel {
		numMessage++
		go func(o *Order) {
			time.Sleep(1 * time.Millisecond)
			fmt.Fprintf(f, "%20s %15d %12.2f %12.2f %15.2f %15.2f\n",
				o.PartNumber, o.Quantity, o.UnitCost, o.UnitPrice, o.UnitCost*float64(o.Quantity), o.UnitPrice*float64(o.Quantity))
			numMessage--
		}(o)
	}

	doneChannel <- true
}

//////////////////////////// The Pipe and Filter Pattern //////////////////////////
func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(in, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

//////////////////////// Simulating Promises ///////////////////////////

type Promise struct {
	successChannel chan interface{}
	failureChannel chan error
}

func SavePO(po *PurchaseOrder, shouldFail bool) *Promise {
	result := new(Promise)

	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	go func() {
		if shouldFail {
			result.failureChannel <- errors.New("Failed to save purchase order")
		} else {
			po.Number = 1234
			result.successChannel <- po
		}
	}()

	return result
}

func (this *Promise) Then(success func(interface{}) error, failure func(error)) *Promise {
	result := new(Promise)

	result.successChannel = make(chan interface{}, 1)
	result.failureChannel = make(chan error, 1)

	go func() {
		select {
		case obj := <-this.successChannel:
			newErr := success(obj)
			if newErr != nil {
				result.successChannel <- obj
			} else {
				result.failureChannel <- newErr
			}
		case err := <-this.failureChannel:
			failure(err)
			result.failureChannel <- err
		}
	}()

	return result
}

//////////////////////// Simulating Callback ///////////////////////////

type PurchaseOrder struct {
	Number int
	Value  float64
}

// func SavePO(po *PurchaseOrder, callback chan *PurchaseOrder) {
// 	po.Number = 1234

// 	callback <- po
// }

/////////////////// Simulating Event /////////////////
type Button struct {
	eventListeners map[string][]chan string
}

func MakeButton() *Button {
	result := new(Button)
	result.eventListeners = make(map[string][]chan string)
	return result
}

func (this *Button) AddEventListener(event string, responseChannel chan string) {
	if _, present := this.eventListeners[event]; present {
		this.eventListeners[event] = append(this.eventListeners[event], responseChannel)
	} else {
		this.eventListeners[event] = []chan string{responseChannel}
	}
}

func (this *Button) RemoveEventListener(event string, listenerChannel chan string) {
	if _, present := this.eventListeners[event]; present {
		for idx, _ := range this.eventListeners[event] {
			if this.eventListeners[event][idx] == listenerChannel {
				this.eventListeners[event] = append(this.eventListeners[event][:idx],
					this.eventListeners[event][idx+1:]...)
				break
			}
		}
	}
}

func (this *Button) TriggerEvent(event string, response string) {
	if _, present := this.eventListeners[event]; present {
		for _, handler := range this.eventListeners[event] {
			go func(handler chan string) {
				handler <- response
			}(handler)
		}
	}
}
