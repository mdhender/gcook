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

/*
 * NAME
 *      vmprintfes - build a formatted string in dynamic memory
 *
 * SYNOPSIS
 *      char *vmprintfes(char *fmt, va_list ap);
 *
 * DESCRIPTION
 *      The vmprintfes function is used to build a formatted string in memory.
 *      It understands all of the ANSI standard sprintf formatting directives.
 *
 * ARGUMENTS
 *      fmt     - string spefiifying formatting to perform
 *      ap      - arguments of types as indicated by the format string
 *
 * RETURNS
 *      string_ty *; string containing formatted string
 *
 * CAVEATS
 *      On error, prints a fatal error message and exists; does not return.
 *
 *      It is the resposnsibility of the caller to invoke str_free to release
 *      the results when finished with.
 */

func vmprintfes(format string, a ...interface{}) *string_ty {
	panic("!")
}
