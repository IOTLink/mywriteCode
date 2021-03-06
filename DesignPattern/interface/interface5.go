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

func getNext(next []Handler) Handler {
	if len(next) > 0 {
		return next[0]
	}
	return nil
}

//___________________________________________________________

func NewProcessor1Handler(next ...Handler) *Processor1Handler {
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



func NewExecuteHandler(next ...Handler) Handler {
	return NewProcessor1Handler(next...)
}

func main() {
	requestContext := &RequestContext{"hello world!"}
	clientContext  := &ClientContext{"liuhy"}
	var handfunc  Handler
	handfunc = NewExecuteHandler()
	handfunc.Handle(requestContext, clientContext)
}

/*
Processor1Handler
Processor2Handler
Processor3Handler
*/