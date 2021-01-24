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

import "fmt"

type state_ty int

// enum state_ty
const (
	state_uninitialized state_ty = iota
	state_C
	state_human
)

var state state_ty
var HAVE_SETLOCALE bool
var LC_ALL int

/*
 * NAME
 *      language_human - set for human conversation
 *
 * DESCRIPTION
 *      The language_human function must be called to change the general
 *      mode over to the default locale (usually dictated by the LANG
 *      environment variable, et al).
 *
 *      The language_human and language_C functions MUST bracket human
 *      interactions, otherwise the mostly-english C locale will be
 *      used.  The default locale through-out the program is otherwise
 *      assumed to be C.
 */

func language_human() {
	switch state {
	case state_uninitialized:
		fatal_raw("you must call language_init() in main (bug)")
	case state_human:
		fatal_raw("unbalanced language_human() call (bug)")
	}
	state = state_human
	if HAVE_SETLOCALE {
		if HAVE_GETTEXT {
			/*
			 * only need to flap the locale about like this
			 * if we are using the gettext function
			 */
			setlocale(LC_ALL, "")
		} /* HAVE_GETTEXT */
	} /* HAVE_SETLOCALE */
}

/*
 * NAME
 *      language_C - set for program conversation
 *
 * DESCRIPTION
 *      The language_C function must be called to restore the locale to
 *      C, so that all the non-human stuff will work.
 *
 *      The language_human and language_C functions MUST bracket human
 *      interactions, otherwise the mostly-english C locale will be
 *      used.  The default locale through-out the program is otherwise
 *      assumed to be C.
 */

func language_C() {
	switch state {
	case state_uninitialized:
		fatal_raw("you must call language_init() in main (bug)")
	case state_C:
		fatal_raw("unbalanced language_C() call (bug)")
	}
	state = state_C
	if HAVE_SETLOCALE {
		if HAVE_GETTEXT {
			/*
			 * only need to flap the locale about like this
			 * if we are using the gettext function
			 */
			setlocale(LC_ALL, "C")
		} /* HAVE_GETTEXT */
	} /* HAVE_SETLOCALE */
}

func setlocale(a int, b string) {
	panic(fmt.Sprintf("setlocale(LC_ALL, %q)", b))
}
