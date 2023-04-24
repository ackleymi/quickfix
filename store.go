// Copyright (c) quickfixengine.org  All rights reserved.
//
// This file may be distributed under the terms of the quickfixengine.org
// license as defined by quickfixengine.org and appearing in the file
// LICENSE included in the packaging of this file.
//
// This file is provided AS IS with NO WARRANTY OF ANY KIND, INCLUDING
// THE WARRANTY OF DESIGN, MERCHANTABILITY AND FITNESS FOR A
// PARTICULAR PURPOSE.
//
// See http://www.quickfixengine.org/LICENSE for licensing information.
//
// Contact ask@quickfixengine.org if any conditions of this licensing
// are not clear to you.

package quickfix

import (
	"time"

	"github.com/pkg/errors"
)

// The MessageStore interface provides methods to record and retrieve messages for resend purposes.
type MessageStore interface {
	NextSenderMsgSeqNum() uint
	NextTargetMsgSeqNum() uint

	IncrNextSenderMsgSeqNum() error
	IncrNextTargetMsgSeqNum() error

	SetNextSenderMsgSeqNum(next uint) error
	SetNextTargetMsgSeqNum(next uint) error

	CreationTime() time.Time

	SaveMessage(seqNum uint, msg []byte) error
	SaveMessageAndIncrNextSenderMsgSeqNum(seqNum uint, msg []byte) error
	GetMessages(beginSeqNum, endSeqNum uint) ([][]byte, error)

	Refresh() error
	Reset() error

	Close() error
}

// The MessageStoreFactory interface is used by session to create a session specific message store.
type MessageStoreFactory interface {
	Create(sessionID SessionID) (MessageStore, error)
}

type memoryStore struct {
	senderMsgSeqNum, targetMsgSeqNum uint
	creationTime                     time.Time
	messageMap                       map[uint][]byte
}

func (store *memoryStore) NextSenderMsgSeqNum() uint {
	return store.senderMsgSeqNum + 1
}

func (store *memoryStore) NextTargetMsgSeqNum() uint {
	return store.targetMsgSeqNum + 1
}

func (store *memoryStore) IncrNextSenderMsgSeqNum() error {
	store.senderMsgSeqNum++
	return nil
}

func (store *memoryStore) IncrNextTargetMsgSeqNum() error {
	store.targetMsgSeqNum++
	return nil
}

func (store *memoryStore) SetNextSenderMsgSeqNum(nextSeqNum uint) error {
	store.senderMsgSeqNum = nextSeqNum - 1
	return nil
}
func (store *memoryStore) SetNextTargetMsgSeqNum(nextSeqNum uint) error {
	store.targetMsgSeqNum = nextSeqNum - 1
	return nil
}

func (store *memoryStore) CreationTime() time.Time {
	return store.creationTime
}

func (store *memoryStore) Reset() error {
	store.senderMsgSeqNum = 0
	store.targetMsgSeqNum = 0
	store.creationTime = time.Now()
	store.messageMap = nil
	return nil
}

func (store *memoryStore) Refresh() error {
	// NOP, nothing to refresh.
	return nil
}

func (store *memoryStore) Close() error {
	// NOP, nothing to close.
	return nil
}

func (store *memoryStore) SaveMessage(seqNum uint, msg []byte) error {
	if store.messageMap == nil {
		store.messageMap = make(map[uint][]byte)
	}

	store.messageMap[seqNum] = msg
	return nil
}

func (store *memoryStore) SaveMessageAndIncrNextSenderMsgSeqNum(seqNum uint, msg []byte) error {
	err := store.SaveMessage(seqNum, msg)
	if err != nil {
		return err
	}
	return store.IncrNextSenderMsgSeqNum()
}

func (store *memoryStore) GetMessages(beginSeqNum, endSeqNum uint) ([][]byte, error) {
	var msgs [][]byte
	for seqNum := beginSeqNum; seqNum <= endSeqNum; seqNum++ {
		if m, ok := store.messageMap[seqNum]; ok {
			msgs = append(msgs, m)
		}
	}
	return msgs, nil
}

type memoryStoreFactory struct{}

func (f memoryStoreFactory) Create(sessionID SessionID) (MessageStore, error) {
	m := new(memoryStore)
	if err := m.Reset(); err != nil {
		return m, errors.Wrap(err, "reset")
	}
	return m, nil
}

// NewMemoryStoreFactory returns a MessageStoreFactory instance that created in-memory MessageStores.
func NewMemoryStoreFactory() MessageStoreFactory { return memoryStoreFactory{} }
