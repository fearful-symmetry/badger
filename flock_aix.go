// +build aix

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
	opts := unix.F_WRLCK | unix.F_RDLCK
	if readOnly {
		opts = unix.F_RDLCK
	}
	for {
		err := unix.FcntlFlock(fd, int(unix.F_SETLK), &unix.Flock_t{
			Type:   int16(opts),
			Whence: io.SeekStart,
			Start:  0,
			Len:    0, // All bytes.
		})
		if err != unix.EINTR {
			return err
		}
	}
}
