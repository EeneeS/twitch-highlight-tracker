package tracker

import (
	"fmt"
	"time"
)

func (t *Tracker) UpdateViewerCount() {
	for {
		viewerCount, err := t.fetchViewerCountFromAPI()
		if err != nil {
			fmt.Println("failed to update viewercount") // maybe fall back to the previous fetched value (if any or >0?)
			return
		}
		t.viewerLock.Lock()
		t.viewerCount = viewerCount
		t.viewerLock.Unlock()
		time.Sleep(time.Minute * 10)
	}
}

func (t *Tracker) fetchViewerCountFromAPI() (int, error) {
	return 1200, nil
}

func (t *Tracker) GetViewerCount() int {
	t.viewerLock.RLock()
	defer t.viewerLock.RUnlock()
	return t.viewerCount
}
