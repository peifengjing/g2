// Copyright 2011 Xing Xing <mikespook@gmail.com> All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*
This module is a Gearman API for the Go Programming Language.
The protocols were written in pure Go. It contains two sub-packages:

The client package is used for sending jobs to the Gearman job server,
and getting responses from the server.

	import "github.com/quantcast/g2/client"

The worker package will help developers to develop Gearman's worker
in an easy way.

	import "github.com/quantcast/g2/worker"
*/
package g2

import (
	_ "github.com/quantcast/g2/client"
	_ "github.com/quantcast/g2/worker"
)
