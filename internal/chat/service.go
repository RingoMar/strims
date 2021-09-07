package chat

import (
	"context"
	"errors"
	"sync"
	"time"

	parser "github.com/MemeLabs/chat-parser"
	"github.com/MemeLabs/go-ppspp/internal/dialer"
	chatv1 "github.com/MemeLabs/go-ppspp/pkg/apis/chat/v1"
	networkv1 "github.com/MemeLabs/go-ppspp/pkg/apis/network/v1"
	"github.com/MemeLabs/go-ppspp/pkg/apis/type/certificate"
	"github.com/MemeLabs/go-ppspp/pkg/protoutil"
	"github.com/MemeLabs/go-ppspp/pkg/timeutil"
	"go.uber.org/zap"
)

// errors
var (
	ErrListingNotFound = errors.New("listing not found")
	ErrSessionNotFound = errors.New("session not found")
	ErrUserNotFound    = errors.New("user not found")
)

const broadcastInterval = 15 * time.Second

func newChatService(logger *zap.Logger, ew *protoutil.ChunkStreamWriter) *chatService {
	return &chatService{
		logger:          logger,
		done:            make(chan struct{}),
		broadcastTicker: time.NewTicker(broadcastInterval),
		eventWriter:     ew,
		entityExtractor: newEntityExtractor(),
		combos:          newComboTransformer(),
	}
}

type chatService struct {
	logger            *zap.Logger
	closeOnce         sync.Once
	done              chan struct{}
	broadcastTicker   *time.Ticker
	lastBroadcastTime timeutil.Time
	eventWriter       *protoutil.ChunkStreamWriter
	lock              sync.Mutex
	certificate       *certificate.Certificate
	entityExtractor   *entityExtractor
	combos            *comboTransformer
}

func (d *chatService) Run(ctx context.Context) error {
	defer d.Close()

	for {
		select {
		case now := <-d.broadcastTicker.C:
			if err := d.broadcast(timeutil.NewFromTime(now)); err != nil {
				return err
			}
		case <-d.done:
			return errors.New("closed")
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (d *chatService) Sync(config *chatv1.Server, emotes []*chatv1.Emote) error {
	emoteNames := [][]rune{}
	for _, emote := range emotes {
		emoteNames = append(emoteNames, []rune(emote.Name))
	}

	d.entityExtractor.parserCtx.Emotes.Replace(emoteNames)
	d.entityExtractor.parserCtx.EmoteModifiers.Replace(parser.RunesFromStrings(config.Room.Modifiers))
	d.entityExtractor.parserCtx.Tags.Replace(parser.RunesFromStrings(config.Room.Tags))

	return nil
}

func (d *chatService) Close() {
	d.closeOnce.Do(func() {
		d.broadcastTicker.Stop()
		close(d.done)
	})
}

func (d *chatService) broadcast(now timeutil.Time) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	var events []*networkv1.DirectoryEvent

	if events != nil {
		err := d.eventWriter.Write(&networkv1.DirectoryEventBroadcast{
			Events: events,
		})
		if err != nil {
			return err
		}
	}

	d.lastBroadcastTime = now
	return nil
}

func (d *chatService) SendMessage(ctx context.Context, req *chatv1.SendMessageRequest) (*chatv1.SendMessageResponse, error) {
	hostCert := dialer.VPNCertificate(ctx)

	m := &chatv1.Message{
		ServerTime: time.Now().UnixNano() / int64(time.Millisecond),
		HostId:     hostCert.Key,
		Nick:       hostCert.GetParent().Subject,
		Body:       req.Body,
		Entities:   d.entityExtractor.Extract(req.Body),
	}

	if err := d.combos.Transform(m); err != nil {
		return nil, err
	}

	err := d.eventWriter.Write(&chatv1.ServerEvent{
		Body: &chatv1.ServerEvent_Message{
			Message: m,
		},
	})
	if err != nil {
		return nil, err
	}

	return &chatv1.SendMessageResponse{}, nil
}