package server

import "testing"

func TestCollectRenderAnalyze(t *testing.T) {
	pReply, err := CollectRenderAnalyze()
	if err != nil {
		t.Errorf("CollectRenderAnalyze error: %v", err)
	} else {
		t.Logf("CollectRenderAnalyze reply: %v", *pReply)
	}
}
