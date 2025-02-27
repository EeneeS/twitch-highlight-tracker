package tracker

import (
	"fmt"
	"math"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/eenees/twitch-highlight-tracker/internal/irc"
)

type Tracker struct {
	client         *irc.Client
	keywords       []string
	keywordCounter map[string]int
	keywordLock    sync.Mutex
	viewerCount    int
	viewerLock     sync.RWMutex
}

func NewTracker(client *irc.Client, keywords []string) *Tracker {
	return &Tracker{
		client:         client,
		keywords:       keywords,
		keywordCounter: make(map[string]int),
	}
}

func (t *Tracker) ReadIncomming() {
	for {
		raw, err := t.client.ReadMessage()
		if err != nil {
			fmt.Println("failed to read message")
			return
		}

		if strings.HasPrefix(raw, "PING") {
			t.client.SendData(strings.Replace(raw, "PING", "PONG", 1))
		}

		if strings.Contains(raw, "PRIVMSG") {
			message := irc.ParseMessage(raw)
			if slices.Contains(t.keywords, message.Text) {
				t.handleTrackedKeyword(message.Text)
			}
		}
	}
}

func (t *Tracker) handleTrackedKeyword(keyword string) {
	t.keywordLock.Lock()
	_, exists := t.keywordCounter[keyword]
	if !exists {
		t.keywordCounter[keyword] = 0
		go t.StartTimer(keyword)
	}
	t.keywordCounter[keyword]++
	t.keywordLock.Unlock()
}

func (t *Tracker) StartTimer(keyword string) {
	go func() {
		time.Sleep(time.Second * 15)

		t.keywordLock.Lock()
		count := t.keywordCounter[keyword]
		delete(t.keywordCounter, keyword)
		t.keywordLock.Unlock()

		fmt.Printf("%v appeared %v times in 15 seconds\n", keyword, count)

		viewerCount := t.GetViewerCount()
		if viewerCount < 1 {
			viewerCount = 1
		}

		messagesPerViewer := float64(count) / float64(viewerCount)

		baseThreshold := 5.0

		// Thanks Claude for the math
		scaledThreshold := baseThreshold
		if viewerCount > 10 {
			scaledThreshold = baseThreshold * (1.0 - 0.2*math.Log10(float64(viewerCount)/10.0))
			if scaledThreshold < 1.0 {
				scaledThreshold = 1.0
			}
		}

		// DEBUG
		fmt.Printf("Messages per viewer: %.2f, Threshold: %.2f\n", messagesPerViewer, scaledThreshold)

		if messagesPerViewer > scaledThreshold {
			fmt.Println("CLIP")
		}
	}()
}

func (t *Tracker) Run() {
	t.ReadIncomming()
}
