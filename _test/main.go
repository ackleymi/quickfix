package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"

	"github.com/quickfixgo/quickfix"
)

var router *quickfix.MessageRouter = quickfix.NewMessageRouter()

type (
	EchoApplication struct {
		log      *log.Logger
		OrderIds map[string]bool
	}
	//
	//
	// Override these impls from the field and tag pkgs for tests.
	//
	//
	// PossResendField is a BOOLEAN field.
	PossResendField struct{ quickfix.FIXBoolean }
	// ClOrdIDField is a STRING field.
	ClOrdIDField struct{ quickfix.FIXString }
	// MsgTypeField is a enum.MsgType field.
	MsgTypeField struct{ quickfix.FIXString }
)

// Tag returns tag.PossResend (97).
func (f PossResendField) Tag() quickfix.Tag { return 97 }

func (f PossResendField) Value() bool { return f.Bool() }

// Tag returns tag.ClOrdID (11).
func (f ClOrdIDField) Tag() quickfix.Tag { return 11 }

func (f ClOrdIDField) Value() string { return f.String() }

// Tag returns tag.MsgType (35).
func (f MsgTypeField) Tag() quickfix.Tag { return 35 }

func (f MsgTypeField) Value() string { return f.String() }

func (e EchoApplication) OnCreate(sessionID quickfix.SessionID) {
	e.log.Printf("OnCreate %v\n", sessionID.String())
}
func (e *EchoApplication) OnLogon(sessionID quickfix.SessionID) {
	e.log.Printf("OnLogon %v\n", sessionID.String())
	e.OrderIds = make(map[string]bool)
}
func (e *EchoApplication) OnLogout(sessionID quickfix.SessionID) {
	e.log.Printf("OnLogout %v\n", sessionID.String())
}
func (e EchoApplication) ToAdmin(msgBuilder *quickfix.Message, sessionID quickfix.SessionID) {
}

func (e EchoApplication) ToApp(msgBuilder *quickfix.Message, sessionID quickfix.SessionID) (err error) {
	return
}

func (e EchoApplication) FromAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) (err quickfix.MessageRejectError) {
	return
}

func (e *EchoApplication) FromApp(msg *quickfix.Message, sessionID quickfix.SessionID) (err quickfix.MessageRejectError) {
	e.log.Println("Got Message ", msg)
	return router.Route(msg, sessionID)
}

func (e *EchoApplication) processMsg(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
	var possResend PossResendField
	msg.Header.Get(&possResend)

	// Usually use tag.ClOrdID instead of literal 11 here.
	if msg.Body.Has(11) {
		var orderID ClOrdIDField

		if err := msg.Body.Get(&orderID); err != nil {
			return err
		}

		sessionOrderID := sessionID.String() + orderID.String()
		if possResend.FIXBoolean {
			if e.OrderIds[sessionOrderID] {
				return nil
			}
		}

		e.OrderIds[sessionOrderID] = true
	}

	reply := copyMessage(msg)
	if possResend.FIXBoolean {
		reply.Header.Set(possResend)
	}

	quickfix.SendToTarget(reply, sessionID)

	return nil
}

func copyMessage(msg *quickfix.Message) *quickfix.Message {
	msgType := new(MsgTypeField)
	msg.Header.Get(msgType)

	msg.Header.Clear()
	msg.Trailer.Clear()

	msg.Header.Set(msgType)

	return msg
}

func main() {
	app := &EchoApplication{}
	app.log = log.New(ioutil.Discard, "", log.LstdFlags)
	//app.log = log.New(os.Stdout, "", log.LstdFlags)

	router.AddRoute(quickfix.BeginStringFIX40, "D", app.processMsg)
	router.AddRoute(quickfix.BeginStringFIX41, "D", app.processMsg)
	router.AddRoute(quickfix.BeginStringFIX42, "D", app.processMsg)
	router.AddRoute(quickfix.BeginStringFIX43, "D", app.processMsg)
	router.AddRoute(quickfix.BeginStringFIX44, "D", app.processMsg)
	router.AddRoute(quickfix.ApplVerIDFIX50, "D", app.processMsg)
	router.AddRoute(quickfix.ApplVerIDFIX50SP1, "D", app.processMsg)
	router.AddRoute(quickfix.ApplVerIDFIX50SP2, "D", app.processMsg)

	router.AddRoute(quickfix.BeginStringFIX42, "d", app.processMsg)
	router.AddRoute(quickfix.BeginStringFIX43, "d", app.processMsg)
	router.AddRoute(quickfix.BeginStringFIX44, "d", app.processMsg)
	router.AddRoute(quickfix.ApplVerIDFIX50, "d", app.processMsg)
	router.AddRoute(quickfix.ApplVerIDFIX50SP1, "d", app.processMsg)
	router.AddRoute(quickfix.ApplVerIDFIX50SP2, "d", app.processMsg)

	cfg, err := os.Open(os.Args[1])
	if err != nil {
		log.Printf("Error opening %v, %v\n", os.Args[1], err)
		return
	}

	appSettings, err := quickfix.ParseSettings(cfg)
	if err != nil {
		log.Println("Error reading cfg:", err)
		return
	}

	fileLogFactory, err := quickfix.NewFileLogFactory(appSettings)

	if err != nil {
		log.Println("Error creating file log factory:", err)
		return
	}

	acceptor, err := quickfix.NewAcceptor(app, quickfix.NewMemoryStoreFactory(), appSettings, fileLogFactory)
	if err != nil {
		log.Println("Unable to create Acceptor: ", err)
		return
	}

	if err = acceptor.Start(); err != nil {
		log.Println("Unable to start Acceptor: ", err)
		return
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt)
	<-interrupt

	acceptor.Stop()
}
