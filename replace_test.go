package main

import (
	"testing"
	"strconv"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
)

func newTestReplace(t testing.TB, config ReplaceConfig, ) *Replace {

	c, err := common.NewConfigFrom(map[string]interface{}{
		"field": config.Field,
		"target": config.Target,
		"find": config.Find,
		"replace": config.Replace,
		"regex": strconv.FormatBool(config.Regex),
	})
	if err != nil {
		t.Fatal(err)
	}

	r, err := New(c)
	if err != nil {
		t.Fatal(err)
	}
	return r.(*Replace)
}

func TestReplaceHashes(t *testing.T) {
	var tests = []ReplaceConfig {
		{Find: "FIND_ME", Replace: "REPLACED"},
		{Field: "message", Target: "message", Find: "FIND_ME", Replace: "REPLACED"},
		{Field: "message1", Target: "message1", Find: "FIND_ME", Replace: "REPLACED"},
		{Field: "message1", Target: "message1", Find: "[^\\w]", Replace: " ", Regex: true},
	}
	var testEvents = []struct {
		fieldName       string
		fieldContents   string
	}{
		{"message", "some message with -> FIND_ME <-"},
		{"message", "some message with -> FIND_ME <-"},
		{"message1", "some message with -> FIND_ME <-"},
		{"message1", "some message with -> FIND_ME <-"},
	}
	var results = []struct {
		targetFieldName       string
		targetFieldContents   string
	}{
		{"message", "some message with -> REPLACED <-"},
		{"message", "some message with -> REPLACED <-"},
		{"message1", "some message with -> REPLACED <-"},
		{"message1", "some message with    FIND_ME   "},
	}
	for i := range tests {
		f := newTestReplace(t, tests[i])
		event := &beat.Event{Fields: common.MapStr{testEvents[i].fieldName: testEvents[i].fieldContents}}
		event, err := f.Run(event)
		if assert.NoError(t, err) {
			t.Logf("Running test case %d", i)
			assert.Equal(t, results[i].targetFieldContents, event.Fields[results[i].targetFieldName])
		}
	}
}
