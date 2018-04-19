package main

import "fmt"

type ClientContext struct {
	Client string
}

//RequestContext contains request, opts, response parameters for handler execution
type RequestContext struct {
	RequestMsg string
}

type Handler interface {
	Handle(context *RequestContext, clientContext *ClientContext)
}

func getNext(next Handler) Handler {
	if next != nil {
		return next
	}
	return nil
}

//___________________________________________________________

func NewProcessor1Handler(next Handler) *Processor1Handler {
	return &Processor1Handler{next: getNext(next)}
}

type Processor1Handler struct {
	next Handler
}

func (h *Processor1Handler) Handle(requestContext *RequestContext, clientContext *ClientContext) {
	fmt.Println("Processor1Handler")
	//Delegate to next step if any
	if h.next != nil {
		h.next.Handle(requestContext, clientContext)
	}
}

//_________________________________________________

func NewProcessor2Handler(next Handler) *Processor2Handler {
	return &Processor2Handler{next: getNext(next)}
}

type Processor2Handler struct {
	next Handler
}

func (h *Processor2Handler) Handle(requestContext *RequestContext, clientContext *ClientContext) {
	fmt.Println("Processor2Handler")
	//Delegate to next step if any
	if h.next != nil {
		h.next.Handle(requestContext, clientContext)
	}
}

//_________________________________________________
func NewProcessor3Handler(next Handler) *Processor3Handler {
	return &Processor3Handler{next: getNext(next)}
}

type Processor3Handler struct {
	next Handler
}

func (h *Processor3Handler) Handle(requestContext *RequestContext, clientContext *ClientContext) {
	fmt.Println("Processor3Handler")
	//Delegate to next step if any
	if h.next != nil {
		h.next.Handle(requestContext, clientContext)
	}
}

//_________________________________________________

func NewExecuteHandler(next Handler) Handler {
	return NewProcessor1Handler(
		NewProcessor2Handler(
			NewProcessor3Handler(next),
			),
		)
}



func main() {
	requestContext := &RequestContext{"hello world!"}
	clientContext  := &ClientContext{"liuhy"}
	var handfunc  Handler
	handfunc = NewExecuteHandler(nil)
	handfunc.Handle(requestContext, clientContext)
}

/*
Processor1Handler
Processor2Handler
Processor3Handler
*/