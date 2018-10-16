package main

import (
	"fmt"
	"strings"
	"regexp"
	"reflect"

	"github.com/pkg/errors"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/processors"
)

type Replace struct {
	config ReplaceConfig
}

func (r Replace) String() string {
	return fmt.Sprintf("config => %#v", r.config)
}

func New(c *common.Config) (processors.Processor, error) {
	rc := defaultReplaceConfig
	err := c.Unpack(&rc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unpack replace config")
	}

	return &Replace{
		config: rc,
	}, nil
}

func (r *Replace) Run(event *beat.Event) (*beat.Event, error) {

	var findInField = r.config.Field
	if findInField == "" {
		findInField = defaultReplaceConfig.Field
	}
	fieldObject, _ := event.GetValue(findInField)
	val := reflect.ValueOf(fieldObject)

	// Follow the pointer.
	if val.Kind() == reflect.Ptr && !val.IsNil() {
		val = reflect.ValueOf(val.Elem().Interface())
	}
	var fieldText = ""
	if val.IsValid() {
		fieldText = fmt.Sprintf("%v", val.Interface())
	} else {
		return event, nil
	}
	var target = r.config.Target

	if target == "" {
		target = findInField
	}

	if fieldText == "" || r.config.Find == "" {
		event.PutValue(target, fieldText)
	} else if r.config.Regex {
		replacer := regexp.MustCompile(r.config.Find)
		result := replacer.ReplaceAllString(fieldText, r.config.Replace)
		event.PutValue(target, result)
	} else {
		replacer := strings.NewReplacer(r.config.Find, r.config.Replace)
		result := replacer.Replace(fieldText)
		event.PutValue(target, result)
	}


	return event, nil
}
