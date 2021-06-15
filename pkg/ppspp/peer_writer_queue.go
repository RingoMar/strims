package ppspp

import (
	"sync"
)

var peerWriterQueueItemPool = sync.Pool{
	New: func() interface{} {
		return &peerWriterQueueItem{}
	},
}

func putPeerWriterQueueItemPool(i *peerWriterQueueItem) {
	i.next = nil
	i.c = nil
	peerWriterQueueItemPool.Put(i)
}

type peerWriterQueueItem struct {
	next *peerWriterQueueItem
	c    PeerWriter
}

type PeerWriterQueueTicket struct {
	v uint64
}

func newPeerWriterQueue() peerWriterQueue {
	return peerWriterQueue{
		v: 1,
	}
}

type peerWriterQueue struct {
	head, tail *peerWriterQueueItem
	v          uint64
}

func (q *peerWriterQueue) Push(t *PeerWriterQueueTicket, c PeerWriter) bool {
	if q.v == t.v {
		return false
	}
	t.v = q.v

	i := peerWriterQueueItemPool.Get().(*peerWriterQueueItem)
	i.next = nil
	i.c = c

	if q.head == nil {
		q.head = i
	}

	if q.tail != nil {
		q.tail.next = i
	}
	q.tail = i

	return true
}

func (q *peerWriterQueue) Remove(c PeerWriter) {
	var prev *peerWriterQueueItem
	for i := q.head; i != nil; i = i.next {
		if i.c == c {
			if prev != nil {
				prev.next = i.next
			}
			if q.head == i {
				q.head = i.next
			}
			if q.tail == i {
				q.tail = prev
			}

			putPeerWriterQueueItemPool(i)
			return
		}

		prev = i
	}
}

func (q *peerWriterQueue) Detach() detachedPeerWriterQueue {
	c := detachedPeerWriterQueue{
		head: q.head,
	}

	q.head = nil
	q.tail = nil
	q.v++

	return c
}

func (q *peerWriterQueue) Empty() bool {
	return q.head == nil
}

type detachedPeerWriterQueue struct {
	head *peerWriterQueueItem
}

func (q *detachedPeerWriterQueue) Pop() (PeerWriter, bool) {
	if q.head == nil {
		return nil, false
	}

	h := q.head

	q.head = h.next

	defer putPeerWriterQueueItemPool(h)
	return h.c, true
}