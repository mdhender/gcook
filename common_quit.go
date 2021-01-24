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

import "os"

var quit_list_prio []quit_ty
var quit_list []quit_ty

/*
 * NAME
 *      quit - leave program
 *
 * SYNOPSIS
 *      void quit(int status);
 *
 * DESCRIPTION
 *      The quit function causes normal program termination to occur.
 *
 *      First, all functions registered by the quit_handler function are
 *      called, in the reverse order of their registration.
 *
 *      Next, the program is terminated using the exit() function.
 *
 * CAVEAT
 *      The quit function never returns to its caller.
 */

func quit(n int) {
	star_eoln()

	length := len(quit_list_prio)
	for length > 0 {
		length--
		quit_list_prio[length]()
	}

	length = len(quit_list)
	for length > 0 {
		length--
		quit_list[length]()
	}

	os.Exit(n)
}

/*
 * NAME
 *      quit_handler
 *
 * SYNOPSIS
 *      int quit_handler(quit_ty);
 *
 * DESCRIPTION
 *      The quit_handler function registers the function pointed to by func,
 *      to be called without arguments at normal program termination.
 */

func quit_handler(fn quit_ty) {
	quit_list = append(quit_list, fn)
}

func quit_handler_prio(fn quit_ty) {
	quit_list_prio = append(quit_list_prio, fn)
}
