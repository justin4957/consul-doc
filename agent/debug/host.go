// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/debug/host.go
Host module for agent layer

## Tags
agent

## Exports
CollectHostInfo, HostInfo

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/debug/host.go> a code:Module ;
    code:name "agent/debug/host.go" ;
    code:description "Host module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :CollectHostInfo, :HostInfo ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package debug

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

const (
	// DiskUsagePath is the path to check usage of the disk.
	// Must be a filesystem path such as "/", not device file path like "/dev/vda1"
	DiskUsagePath = "/"
)

// HostInfo includes information about resources on the host as well as
// collection time and
type HostInfo struct {
	Memory         *mem.VirtualMemoryStat
	CPU            []cpu.InfoStat
	Host           *host.InfoStat
	Disk           *disk.UsageStat
	CollectionTime int64
	Errors         []error
}

// CollectHostInfo queries the host system and returns HostInfo. Any
// errors encountered will be returned in HostInfo.Errors
func CollectHostInfo() *HostInfo {
	info := &HostInfo{CollectionTime: time.Now().UTC().UnixNano()}

	if h, err := host.Info(); err != nil {
		info.Errors = append(info.Errors, err)
	} else {
		info.Host = h
	}

	if v, err := mem.VirtualMemory(); err != nil {
		info.Errors = append(info.Errors, err)
	} else {
		info.Memory = v
	}

	if d, err := disk.Usage(DiskUsagePath); err != nil {
		info.Errors = append(info.Errors, err)
	} else {
		info.Disk = d
	}

	if c, err := cpu.Info(); err != nil {
		info.Errors = append(info.Errors, err)
	} else {
		info.CPU = c
	}

	return info
}
