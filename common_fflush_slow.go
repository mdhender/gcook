/*
 * cook - file construction tool
 * Copyright (C) 2021 Michael D Henderson
 * Copyright (C) 1993-2010 Peter Miller
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"os"
	"time"
)

/*
 * If the fflush call reports and error, we allow this many seconds for
 * recovery.  After that, we report the error as it happened.
 */
const MAX_FLUSH_TRY = 10

var fflush_retry_count = 0

func fflush_slowly(fp *os.File) (err error) {
	for attempts := 0; attempts < MAX_FLUSH_TRY; attempts++ {
		if err = fp.Sync(); err == nil {
			/*
			 * No problem - quit trying.
			 * Report success.
			 */
			return nil
		}

		fflush_retry_count++

		/*
		 * Perhaps a little rest will clear the problem.
		 */
		time.Sleep(time.Second)
	}

	/* ran out of tries - not temporary */
	return err
}
