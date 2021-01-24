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
	"fmt"
	"os"
	"strconv"
)

const DEFAULT_PAGE_LENGTH = 24
const MIN_PAGE_LENGTH = 10
const MAX_PAGE_LENGTH = 30000

const DEFAULT_PAGE_WIDTH = 80
const MIN_PAGE_WIDTH = 40

// todo: /* MAX_PAGE_WIDTH is defined in common/page.h */

var page_length int
var page_width int

func default_page_sizes() {
	if page_length == 0 {
		if lines := os.Getenv("LINES"); lines != "" {
			if n, err := strconv.Atoi(lines); err == nil {
				if n < MIN_PAGE_LENGTH {
					n = MIN_PAGE_LENGTH
				}
				if n > MAX_PAGE_LENGTH {
					n = MAX_PAGE_LENGTH
				}
				page_length = n
			}
		}
	}

	if page_width == 0 {
		if cols := os.Getenv("COLS"); cols != "" {
			if n, err := strconv.Atoi(cols); err == nil {
				if n < MIN_PAGE_WIDTH {
					n = MIN_PAGE_WIDTH
				}
				if n > MAX_PAGE_WIDTH {
					n = MAX_PAGE_WIDTH
				}
				page_width = n
			}
		}
	}

	// todo: #ifdef TIOCGWINSZ
	if page_length == 0 || page_width == 0 {
		fmt.Printf("todo: #ifdef TIOCGWINSZ\n")
	}

	if page_length == 0 {
		page_length = DEFAULT_PAGE_LENGTH
	}

	if page_width == 0 {
		page_width = DEFAULT_PAGE_WIDTH
	}
}

func page_width_get() int {
	/*
	 * must not generate a fatal error in this function,
	 * as it is used by 'error.c' when reporting fatal errors.
	 *
	 * must not put tracing in this function,
	 * because 'trace.c' uses it to determine the width.
	 */
	if page_width == 0 {
		default_page_sizes()
	}
	return page_width
}
