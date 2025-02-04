package tests

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net"
	"regexp"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	simplefixgo "github.com/b2broker/simplefix-go"
	"github.com/b2broker/simplefix-go/fix"
	"github.com/b2broker/simplefix-go/session"
	"github.com/b2broker/simplefix-go/session/storages/memory"
	fixgen "github.com/b2broker/simplefix-go/tests/fix44"
	"github.com/b2broker/simplefix-go/utils"
)

func TestHeartbeat(t *testing.T) {
	const (
		countOfHeartbeats = 4
		heartBtInt        = 1
	)

	// close acceptor after work
	acceptor, addr := RunAcceptor(0, t, memory.NewStorage(100, 100))
	defer acceptor.Close()
	go func() {
		err := acceptor.ListenAndServe()
		if err != nil && !errors.Is(err, simplefixgo.ErrConnClosed) {
			panic(err)
		}
	}()

	initiatorSession, initiatorHandler := RunNewInitiator(addr, t, &session.LogonSettings{
		TargetCompID:  "Server",
		SenderCompID:  "Client",
		HeartBtInt:    heartBtInt,
		EncryptMethod: fixgen.EnumEncryptMethodNoneother,
	})

	waitHeartbeats := utils.TimedWaitGroup{}
	waitHeartbeats.Add(countOfHeartbeats)
	heartbeats := 4

	initiatorHandler.HandleIncoming(fixgen.MsgTypeHeartbeat, func(msg []byte) bool {
		if heartbeats <= 0 {
			return true
		}
		heartbeats--
		waitHeartbeats.Done()
		return true
	})

	initiatorSession.OnChangeState(utils.EventLogon, func() bool {
		t.Log("client connected to server")
		return true
	})

	err := waitHeartbeats.WaitWithTimeout(time.Second * countOfHeartbeats * heartBtInt * 2)
	if err != nil {
		t.Fatalf("wait heartbeats: %s", err)
	}
}

func TestGroup(t *testing.T) {
	const (
		heartBtInt = 1
	)
	var testInstrumentSymbols = map[string]struct{}{
		"BTC/USD": {},
		"ETH/GBP": {},
	}
	var done = make(chan struct{})

	// close acceptor after work
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("listen error: %s", err)
	}

	handlerFactory := simplefixgo.NewAcceptorHandlerFactory(fixgen.FieldMsgType, 10)

	acceptor := simplefixgo.NewAcceptor(listener, handlerFactory, func(handler simplefixgo.AcceptorHandler) {
		s, err := session.NewAcceptorSession(
			&pseudoGeneratedOpts,
			handler,
			&session.LogonSettings{
				LogonTimeout: time.Second * 30,
				HeartBtLimits: &session.IntLimits{
					Min: 5,
					Max: 60,
				},
			},
			func(request *session.LogonSettings) (err error) { return nil },
		)
		if err != nil {
			panic(err)
		}

		err = s.Run()
		if err != nil {
			t.Fatalf("run s: %s", err)
		}

		handler.HandleIncoming(fixgen.MsgTypeMarketDataRequest, func(msg []byte) bool {
			request, err := fixgen.ParseMarketDataRequest(msg)
			if err != nil {
				panic(err)
			}

			for _, relatedSym := range request.RelatedSymGrp().Entries() {
				symbol := relatedSym.Instrument().Symbol()
				if _, ok := testInstrumentSymbols[symbol]; !ok {
					t.Fatalf("unexpected symbol: %s", symbol)
				}
				delete(testInstrumentSymbols, symbol)
			}

			if len(testInstrumentSymbols) > 0 {
				t.Fatalf("some instruments remained at map: %v", testInstrumentSymbols)
			}

			close(done)
			return true
		})

		s.SetMessageStorage(memory.NewStorage(100, 100))
	})

	defer acceptor.Close()
	go func() {
		err := acceptor.ListenAndServe()
		if err != nil && !errors.Is(err, simplefixgo.ErrConnClosed) {
			panic(err)
		}
	}()

	initiatorSession, _ := RunNewInitiator(listener.Addr().String(), t, &session.LogonSettings{
		TargetCompID:  "Server",
		SenderCompID:  "Client",
		HeartBtInt:    heartBtInt,
		EncryptMethod: fixgen.EnumEncryptMethodNoneother,
	})

	initiatorSession.OnChangeState(utils.EventLogon, func() bool {
		relatedSymbols := fixgen.NewRelatedSymGrp()

		for symbol := range testInstrumentSymbols {
			relatedSymbols.AddEntry(fixgen.NewRelatedSymEntry().SetInstrument(fixgen.NewInstrument().SetSymbol(symbol)))
		}

		err := initiatorSession.Send(fixgen.NewMarketDataRequest(
			"test",
			fixgen.EnumSubscriptionRequestTypeSnapshot,
			20,
			fixgen.NewMDEntryTypesGrp(),
			relatedSymbols,
		))
		if err != nil {
			panic(err)
		}

		return true
	})

	initiatorSession.OnChangeState(utils.EventLogon, func() bool {
		t.Log("client connected to server")
		return true
	})

	select {
	case <-time.After(time.Second * 5):
		t.Fatalf("wait heartbeats: %s", err)
	case <-done:
	}
}

func TestTestRequest(t *testing.T) {
	const (
		heartBtInt = 5
		testReqID  = "aloha"
	)

	// close acceptor after work
	acceptor, addr := RunAcceptor(0, t, memory.NewStorage(100, 100))
	defer acceptor.Close()
	go func() {
		err := acceptor.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	initiatorSession, initiatorHandler := RunNewInitiator(addr, t, &session.LogonSettings{
		TargetCompID:  "Server",
		SenderCompID:  "Client",
		HeartBtInt:    heartBtInt,
		EncryptMethod: fixgen.EnumEncryptMethodNoneother,
	})

	waitHeartbeats := utils.TimedWaitGroup{}
	waitHeartbeats.Add(1)

	initiatorHandler.HandleIncoming(fixgen.MsgTypeHeartbeat, func(msg []byte) bool {
		heartbeatMsg, err := fixgen.ParseHeartbeat(msg)
		if err != nil {
			t.Fatalf("parse heartbeat: %s", err)
		}

		if heartbeatMsg.TestReqID() == testReqID {
			waitHeartbeats.Done()
		}

		return true
	})

	initiatorSession.OnChangeState(utils.EventLogon, func() bool {
		t.Log("client connected to server")
		t.Log("send test request")

		testRequestMsg := fixgen.TestRequest{}.New()
		testRequestMsg.SetFieldTestReqID(testReqID)

		err := initiatorSession.Send(testRequestMsg)
		if err != nil {
			panic(err)
		}

		return true
	})

	err := waitHeartbeats.WaitWithTimeout(time.Second * heartBtInt * 2)
	if err != nil {
		t.Fatalf("wait heartbeats: %s", err)
	}
}

func TestResendSequence(t *testing.T) {
	const (
		waitingResend       = time.Second * 6
		beforeResendRequest = time.Second * 3
		resendBegin         = 1
		resendEnd           = 3
	)

	var countOfResending = resendEnd - resendBegin + 1 // including

	// close acceptor after work
	acceptor, addr := RunAcceptor(0, t, memory.NewStorage(100, 100))
	defer acceptor.Close()
	go func() {
		err := acceptor.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	initiatorSession, initiatorHandler := RunNewInitiator(addr, t, &session.LogonSettings{
		TargetCompID:  "Server",
		SenderCompID:  "Client",
		HeartBtInt:    1,
		EncryptMethod: fixgen.EnumEncryptMethodNoneother,
	})

	waitRepeats := utils.TimedWaitGroup{}
	waitRepeats.Add(countOfResending)
	messages := new(sync.Map)

	initiatorHandler.HandleIncoming(simplefixgo.AllMsgTypes, func(msg []byte) bool {
		msgSeqNumB, err := fix.ValueByTag(msg, fixgen.FieldMsgSeqNum)
		if err != nil {
			t.Fatalf("message sequence num parsing: %s", err)
		}

		msgSeqNum := string(msgSeqNumB)

		old, ok := messages.Load(msgSeqNum)
		if ok {
			if !bytes.Equal(old.([]byte), msg) {
				t.Log("> incoming", string(msg))
				t.Log("> saved", string(old.([]byte)))
				t.Fatalf("> different messages with same sequence number")
			} else {
				defer waitRepeats.Done()
			}
		} else {
			messages.Store(msgSeqNum, msg)
		}

		return true
	})

	initiatorSession.OnChangeState(utils.EventLogon, func() bool {
		t.Log("client connected to server")
		return true
	})

	time.Sleep(beforeResendRequest)
	err := initiatorSession.Send(fixgen.ResendRequest{}.New().SetFieldBeginSeqNo(resendBegin).SetFieldEndSeqNo(resendEnd))
	if err != nil {
		panic(err)
	}

	defer acceptor.Close()
	err = waitRepeats.WaitWithTimeout(waitingResend)
	if err != nil {
		t.Fatalf("wait heartbeats: %s", err)
	}
}

func TestCloseInitiatorConn(t *testing.T) {
	// close acceptor after work
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("listen error: %s", err)
	}

	waitClientClosed := make(chan struct{})
	handlerFactory := simplefixgo.NewAcceptorHandlerFactory(fixgen.FieldMsgType, 10)
	server := simplefixgo.NewAcceptor(listener, handlerFactory, func(handler simplefixgo.AcceptorHandler) {
		s, err := session.NewAcceptorSession(
			&pseudoGeneratedOpts,
			handler,
			&session.LogonSettings{HeartBtLimits: &session.IntLimits{
				Min: 1,
				Max: 30,
			}, LogonTimeout: time.Second * 30},
			func(request *session.LogonSettings) (err error) { return nil },
		)
		if err != nil {
			panic(err)
		}

		err = s.Run()
		if err != nil {
			t.Fatalf("run s: %s", err)
		}

		handler.OnDisconnect(func() bool {
			t.Log("client disconnected")
			waitClientClosed <- struct{}{}
			return true
		})
	})

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatalf("could not dial: %s", err)
	}

	handler := simplefixgo.NewInitiatorHandler(context.Background(), fixgen.FieldMsgType, 10)
	client := simplefixgo.NewInitiator(conn, handler, 10)

	s, err := session.NewInitiatorSession(
		handler,
		&pseudoGeneratedOpts,
		&session.LogonSettings{
			TargetCompID:  "Server",
			SenderCompID:  "Client",
			HeartBtInt:    1,
			EncryptMethod: fixgen.EnumEncryptMethodNoneother,
		},
	)
	if err != nil {
		panic(err)
	}

	go func() {
		err := client.Serve()
		if err != nil && !errors.Is(err, simplefixgo.ErrConnClosed) {
			panic(fmt.Errorf("serve client: %s", err))
		}
	}()

	err = s.Run()
	if err != nil {
		t.Fatalf("run session: %s", err)
	}

	client.Close()

	select {
	case <-waitClientClosed:
	case <-time.After(time.Second * 3):
		t.Fatalf("too long time waiting close")
	}
}

func TestCloseAcceptorConn(t *testing.T) {
	// close acceptor after work
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("listen error: %s", err)
	}

	waitServerDisconnect := make(chan struct{})
	handlerFactory := simplefixgo.NewAcceptorHandlerFactory(fixgen.FieldMsgType, 10)
	server := simplefixgo.NewAcceptor(listener, handlerFactory, func(handler simplefixgo.AcceptorHandler) {
		s, err := session.NewAcceptorSession(
			&pseudoGeneratedOpts,
			handler,
			&session.LogonSettings{
				HeartBtLimits: &session.IntLimits{
					Min: 5,
					Max: 60,
				}, LogonTimeout: time.Second * 30},
			func(request *session.LogonSettings) (err error) { return nil },
		)
		if err != nil {
			panic(err)
		}

		err = s.Run()
		if err != nil {
			t.Fatalf("run s: %s", err)
		}

		handler.OnConnect(func() bool {
			t.Log("server: client connected")
			return true
		})
	})

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatalf("could not dial: %s", err)
	}

	initiatorHandler := simplefixgo.NewInitiatorHandler(context.Background(), fixgen.FieldMsgType, 10)
	client := simplefixgo.NewInitiator(conn, initiatorHandler, 10)

	s, err := session.NewInitiatorSession(
		initiatorHandler,
		&pseudoGeneratedOpts,
		&session.LogonSettings{
			TargetCompID:  "Server",
			SenderCompID:  "Client",
			HeartBtInt:    1,
			EncryptMethod: fixgen.EnumEncryptMethodNoneother,
		},
	)
	if err != nil {
		panic(err)
	}

	initiatorHandler.OnConnect(func() bool {
		t.Log("client: connected to server")
		server.Close()

		return true
	})

	initiatorHandler.OnDisconnect(func() bool {
		t.Log("server disconnected")
		waitServerDisconnect <- struct{}{}
		return true
	})

	go func() {
		err := client.Serve()
		if !errors.Is(err, simplefixgo.ErrConnClosed) {
			panic(fmt.Errorf("serve client: %s", err))
		}
	}()

	err = s.Run()
	if err != nil {
		t.Fatalf("run session: %s", err)
	}

	select {
	case <-waitServerDisconnect:
	case <-time.After(time.Second * 3):
		t.Fatalf("too long time waiting close")
	}
}

func TestLookAtClosingOfInitiator(t *testing.T) {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("listen error: %s", err)
	}

	waitClientDisconnect := make(chan struct{})
	handlerFactory := simplefixgo.NewAcceptorHandlerFactory(fixgen.FieldMsgType, 10)
	server := simplefixgo.NewAcceptor(listener, handlerFactory, func(handler simplefixgo.AcceptorHandler) {
		acceptorSession, err := session.NewAcceptorSession(
			&pseudoGeneratedOpts,
			handler,
			&session.LogonSettings{
				HeartBtLimits: &session.IntLimits{
					Min: 5,
					Max: 60,
				}, LogonTimeout: time.Second * 30},
			func(request *session.LogonSettings) (err error) { return nil },
		)
		if err != nil {
			panic(err)
		}

		err = acceptorSession.Run()
		if err != nil {
			t.Fatalf("run s: %s", err)
		}

		handler.OnConnect(func() bool {
			t.Log("start some message stream")
			go func() {
				for {
					select {
					case <-acceptorSession.Context().Done():
						waitClientDisconnect <- struct{}{}
						return
					case <-time.After(time.Second):
						err := acceptorSession.Send(fixgen.NewMarketDataIncrementalRefresh(fixgen.NewMDEntriesGrp()))
						if err != nil {
							panic(err)
						}
					}
				}
			}()

			return true
		})
	})

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatalf("could not dial: %s", err)
	}

	initiatorHandler := simplefixgo.NewInitiatorHandler(context.Background(), fixgen.FieldMsgType, 10)
	client := simplefixgo.NewInitiator(conn, initiatorHandler, 10)

	initiatorSession, err := session.NewInitiatorSession(
		initiatorHandler,
		&pseudoGeneratedOpts,
		&session.LogonSettings{
			TargetCompID:  "Server",
			SenderCompID:  "Client",
			HeartBtInt:    1,
			EncryptMethod: fixgen.EnumEncryptMethodNoneother,
		},
	)
	if err != nil {
		panic(err)
	}

	initiatorHandler.OnConnect(func() bool {
		t.Log("client: connected to server")
		client.Close()

		return true
	})

	go func() {
		err := client.Serve()
		if err != nil && !errors.Is(err, simplefixgo.ErrConnClosed) {
			panic(fmt.Errorf("serve client: %s", err))
		}
	}()

	err = initiatorSession.Run()
	if err != nil {
		t.Fatalf("run session: %s", err)
	}

	select {
	case <-waitClientDisconnect:
	case <-time.After(time.Second * 3):
		t.Fatalf("too long time waiting close")
	}
}

func TestInterruptHandling(t *testing.T) {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("listen error: %s", err)
	}

	waitClientDisconnect := make(chan struct{})
	handlerFactory := simplefixgo.NewAcceptorHandlerFactory(fixgen.FieldMsgType, 10)
	server := simplefixgo.NewAcceptor(listener, handlerFactory, func(handler simplefixgo.AcceptorHandler) {
		acceptorSession, err := session.NewAcceptorSession(
			&pseudoGeneratedOpts,
			handler,
			&session.LogonSettings{
				HeartBtLimits: &session.IntLimits{
					Min: 5,
					Max: 60,
				}, LogonTimeout: time.Second * 30},
			func(request *session.LogonSettings) (err error) { return nil },
		)
		if err != nil {
			t.Fatalf("new session: %s", err)
		}

		err = acceptorSession.Run()
		if err != nil {
			t.Fatalf("run acceptor session: %s", err)
		}

		handler.OnConnect(func() bool {
			t.Log("start some message stream")
			go func() {
				select {
				case <-acceptorSession.Context().Done():
					waitClientDisconnect <- struct{}{}
					return
				case <-time.After(time.Second):
					err := acceptorSession.Send(fixgen.NewMarketDataIncrementalRefresh(fixgen.NewMDEntriesGrp()))
					if err != nil {
						panic(err)
					}
				}
			}()

			return true
		})
	})

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatalf("could not dial: %s", err)
	}

	initiatorHandler := simplefixgo.NewInitiatorHandler(context.Background(), fixgen.FieldMsgType, 10)
	client := simplefixgo.NewInitiator(conn, initiatorHandler, 10)

	initiatorSession, err := session.NewInitiatorSession(
		initiatorHandler,
		&pseudoGeneratedOpts,
		&session.LogonSettings{
			TargetCompID:  "Server",
			SenderCompID:  "Client",
			HeartBtInt:    1,
			EncryptMethod: fixgen.EnumEncryptMethodNoneother,
		},
	)
	if err != nil {
		panic(err)
	}

	initiatorHandler.OnConnect(func() bool {
		t.Log("client: connected to server")
		client.Close()

		return true
	})

	go func() {
		err := client.Serve()
		if err != nil && errors.Is(err, simplefixgo.ErrConnClosed) {
			panic(fmt.Errorf("serve client: %s", err))
		}
	}()

	err = initiatorSession.Run()
	if err != nil {
		t.Fatalf("run session: %s", err)
	}

	select {
	case <-waitClientDisconnect:
	case <-time.After(time.Second * 3):
		t.Fatalf("too long time waiting close")
	}
}

func TestHighload(t *testing.T) {
	const heartBtInt = 5

	triesNum := 100
	threadsNum := 5

	// close acceptor after work
	acceptor, addr := RunAcceptor(0, t, memory.NewStorage(100, 100))
	defer acceptor.Close()
	go func() {
		err := acceptor.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	initiatorSession, initiatorHandler := RunNewInitiator(addr, t, &session.LogonSettings{
		TargetCompID:  "Server",
		SenderCompID:  "Client",
		HeartBtInt:    heartBtInt,
		EncryptMethod: fixgen.EnumEncryptMethodNoneother,
	})

	waitSnapshots := utils.TimedWaitGroup{}
	waitSnapshots.Add(triesNum * threadsNum)

	seqCount := int64(2)
	seqRegexp := regexp.MustCompile(`34=(\d+)`)

	initiatorHandler.HandleOutgoing(fixgen.MsgTypeMarketDataSnapshotFullRefresh, func(msg simplefixgo.SendingMessage) bool {
		waitSnapshots.Done()
		return true
	})

	initiatorHandler.HandleOutgoing(simplefixgo.AllMsgTypes, func(msg simplefixgo.SendingMessage) bool {
		data, err := msg.ToBytes()
		if err != nil {
			t.Fatalf("parse snapshot: %s", err)
		}

		str := string(bytes.ReplaceAll(data, fix.Delimiter, []byte("|")))

		seq := seqRegexp.FindStringSubmatch(str)
		if len(seq) > 0 {
			seqInt, _ := strconv.Atoi(seq[1])
			if int64(seqInt) != seqCount {
				t.Fatalf("broken sequence: %d, reference %d", seqInt, seqCount)
			}
			atomic.AddInt64(&seqCount, 1)
		}

		return true
	})

	initiatorSession.OnChangeState(utils.EventLogon, func() bool {
		t.Log("client connected to server")
		t.Log("send test request")

		symbol := fixgen.NewInstrument().SetSymbol("XXX/YYY")
		group := fixgen.NewMDEntriesGrp()
		for i := 0; i < 30; i++ {
			entry := fixgen.NewMDEntriesEntry().
				SetQuoteEntryID("73b7bb2e-2cf2-445f-ad54-c3ba5e443eda").SetMDEntryTime(time.Now().Format(time.RFC3339)).
				SetMDEntryType(fixgen.EnumMDEntryTypeBid).
				SetMDEntryPx(1234.5).
				SetMDEntrySize(0.98)
			group.AddEntry(entry)
		}

		testMsg := fixgen.NewMarketDataSnapshotFullRefresh(symbol, group)

		for j := 0; j < threadsNum; j++ {
			go func() {
				for i := 0; i < triesNum; i++ {
					err := initiatorSession.Send(testMsg)
					if err != nil {
						panic(err)
					}
				}
			}()
		}

		return true
	})

	err := waitSnapshots.WaitWithTimeout(time.Second * heartBtInt)
	if err != nil {
		t.Fatalf("wait snapshots: %s", err)
	}
}

func TestSessionClosing(t *testing.T) {
	const (
		heartBtInt = 5
		testReqID  = "aloha"
	)

	// close acceptor after work
	acceptor, addr := RunAcceptor(0, t, memory.NewStorage(100, 100))
	defer acceptor.Close()
	go func() {
		err := acceptor.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	initiatorSession, initiatorHandler := RunNewInitiator(addr, t, &session.LogonSettings{
		TargetCompID:  "Server",
		SenderCompID:  "Client",
		HeartBtInt:    heartBtInt,
		EncryptMethod: fixgen.EnumEncryptMethodNoneother,
	})

	triesNum := 5

	waitHeartbeats := utils.TimedWaitGroup{}
	waitHeartbeats.Add(triesNum)

	initiatorHandler.HandleIncoming(fixgen.MsgTypeHeartbeat, func(msg []byte) bool {
		heartbeatMsg, err := fixgen.ParseHeartbeat(msg)
		if err != nil {
			t.Fatalf("parse heartbeat: %s", err)
		}

		if heartbeatMsg.TestReqID() == testReqID {
			waitHeartbeats.Done()
		}

		return true
	})

	initiatorSession.OnChangeState(utils.EventLogon, func() bool {
		t.Log("client connected to server")
		t.Log("send test request")

		testRequestMsg := fixgen.TestRequest{}.New()
		testRequestMsg.SetFieldTestReqID(testReqID)

		for i := 0; i < triesNum; i++ {
			err := initiatorSession.Send(testRequestMsg)
			if err != nil {
				panic(err)
			}
		}

		return true
	})

	err := waitHeartbeats.WaitWithTimeout(time.Second * heartBtInt)
	if err != nil {
		t.Fatalf("wait heartbeats: %s", err)
	}

	if err := initiatorSession.Stop(); err != nil {
		t.Fatalf("unexpected behaviour, got error: %v", err)
	}

	time.Sleep(10 * time.Millisecond)

	if ctxErr := initiatorSession.Context().Err(); ctxErr == nil {
		t.Fatalf("context should be already cancelled")
	}

}
