// +build !windows,!plan9,!aix

package badger

import "golang.org/x/sys/unix"

/*
 * Copyright 2017 Dgraph Labs, Inc. and Contributors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

func flock(fd uintptr, readOnly bool) error {
	opts := unix.LOCK_EX | unix.LOCK_NB
	if readOnly {
		opts = unix.LOCK_SH | unix.LOCK_NB
	}

	return unix.Flock(int(fd), opts)
}
