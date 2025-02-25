//go:build windows
// +build windows

/*
 * Copyright 2019-2020 by Nedim Sabic Sabic
 * https://www.fibratus.io
 * All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package filetime

import (
	"syscall"
	"time"
)

// ToEpoch converts file timestamp to Unix time.
func ToEpoch(ts uint64) time.Time {
	ft := &syscall.Filetime{HighDateTime: uint32(ts >> 32), LowDateTime: uint32(ts)}
	return time.Unix(0, ft.Nanoseconds())
}
