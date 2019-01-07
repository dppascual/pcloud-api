package main

import "time"

// DeviceStats represents device metrics
type DeviceStats struct {
	Read    time.Time `json:"read"`
	Process `json:"procs"`
	Memory  `json:"mem"`
	Swap    `json:"swap"`
	IO      `json:"io"`
	System  `json:"system"`
	CPU     `json:"cpu"`
}

// Process represents process metrics
type Process struct {
	// Running represents the number of processes waiting for run time
	Running int `json:"r"`
	// Waiting represents the number of processes in uninterruptible sleep
	Waiting int `json:"b"`
}

// Memory represents memory metrics
type Memory struct {
	// Virtual represents the amount of virtual memory used
	Virtual int `json:"swpd"`
	// Free represents the amount of idle memory
	Free int `json:"free"`
	// Buffer represents the amount of memory used as buffers
	Buffer int `json:"buff"`
	// Cache represents the amount of memory used as cache
	Cache int `json:"cache"`
}

// Swap represents swap metrics
type Swap struct {
	// SwapIn represents the amount of memory swapped in from disk
	SwapIn int `json:"si"`
	// SwapOut represents the amount of memory swapped to disk
	SwapOut int `json:"so"`
}

// IO represents io metrics
type IO struct {
	// BlocksIn represents blocks received from a block device
	BlocksIn int `json:"bi"`
	// BlocksOut represents blocks sent to a block device
	BlocksOut int `json:"bo"`
}

// System represents system metrics
type System struct {
	// Interrupts represents the number of interrupts per second
	Interrupts int `json:"in"`
	// ContextSwitch represents the number of context switches per second
	ContextSwitch int `json:"cs"`
}

// CPU represents cpu metrics
type CPU struct {
	// Idle represents the percentage of time spent idle
	Idle int `json:"id"`
	// User represents the percentage of time spent running non-kernel code (user time)
	User int `json:"us"`
	// System represents the percentage of time spent running kernel code
	System int `json:"sy"`
	// Wait represents the percentage of time spent waiting for IO
	Wait int `json:"wa"`
	// Stolen represents the percentage of time stolen from a virtual machine
	Stolen int `json:"st"`
}
