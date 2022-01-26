// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

package resourcesmisc

import (
	"fmt"
	"strings"

	"github.com/cppforlife/color"
	ctlres "github.com/k14s/kapp/pkg/kapp/resources"
)

type Deleting struct {
	resource ctlres.Resource
}

func NewDeleting(resource ctlres.Resource) *Deleting {
	if resource.IsDeleting() {
		return &Deleting{resource}
	}
	return nil
}

func (s Deleting) IsDoneApplying() DoneApplyState {
	return DoneApplyState{Done: false, Message: "Deleting"}
}

var (
	uiWaitMsgPrefix = color.New(color.Faint).Sprintf(" ^ ")
)

func (s Deleting) BuildDescMsg(doneApplying bool) []string {
	if !doneApplying && len(s.resource.Finalizers()) > 0 {
		return []string{uiWaitMsgPrefix + fmt.Sprintf("Waiting on finalizers: %s",
			strings.Join(s.resource.Finalizers(), ", "))}
	}
	return []string{}
}
