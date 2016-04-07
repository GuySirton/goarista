// Copyright (C) 2016  Arista Networks, Inc.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

package kafka

import (
	"flag"
	"os"
)

// Addresses is the flag for kafka's comma-separated addresses
var Addresses = flag.String("kafkaaddrs", "", "kafka's comma-separated addresses")

// Topic is the flag for kafka's topic
var Topic = flag.String("kafkatopic", os.Args[0], "kafka's topic")
