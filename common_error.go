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
)

/*
 * NAME
 *      error - place a message on the error stream
 *
 * SYNOPSIS
 *      void error(char *s, ...);
 *
 * DESCRIPTION
 *      Error places a message on the error output stream.
 *      The first argument is a printf-like format string,
 *      optionally followed by other arguments.
 *      The message will be prefixed by the program name and a colon,
 *      and will be terminated with a newline, automatically.
 *
 * CAVEAT
 *      Things like "error(filename)" blow up if the filename
 *      contains a '%' character.
 */

func error_raw(format string, a ...interface{}) {
	wrap(fmt.Sprintf(format, a...))
}

/*
 * NAME
 *      fatal - place a message on the error stream and exit
 *
 * SYNOPSIS
 *      void fatal(char *s, ...);
 *
 * DESCRIPTION
 *      Fatal places a message on the error output stream and exits.
 *      The first argument is a printf-like format string,
 *      optionally followed by other arguments.
 *      The message will be prefixed by the program name and a colon,
 *      and will be terminated with a newline, automatically.
 *
 * CAVEAT
 *      Things like "error(filename)" blow up if the filename
 *      contains a '%' character.
 *
 *      This function does NOT return.
 */

func fatal_raw(format string, a ...interface{}) {
	wrap(fmt.Sprintf(format, a...))
	quit(1)
}

/*
 * NAME
 *      nfatal - place a system fault message on the error stream and exit
 *
 * SYNOPSIS
 *      void nfatal(char *s, ...);
 *
 * DESCRIPTION
 *      Nfatal places a message on the error output stream and exits.
 *      The first argument is a printf-like format string,
 *      optionally followed by other arguments.
 *      The message will be prefixed by the program name and a colon,
 *      and will be terminated with a text description of the error
 *      indicated by the 'errno' global variable, automatically.
 *
 * CAVEAT
 *      Things like "nfatal(filename)" blow up if the filename
 *      contains a '%' character.
 *
 *      This function does NOT return.
 */

func nfatal_raw(err error, format string, a ...interface{}) {
	error_raw("%s: %+v", fmt.Sprintf(format, a...), err)
	quit(1)
}

/*
 * NAME
 *      wrap - wrap s string over lines
 *
 * SYNOPSIS
 *      void wrap(char *);
 *
 * DESCRIPTION
 *      The wrap function is used to print error messages onto stderr
 *      wrapping ling lines.
 *
 * CAVEATS
 *      Line length is assumed to be 80 characters.
 */

func wrap(s string) {
	// todo: return to original functionality

	/*
	 * Flush stdout so that errors are in sync with the output.
	 * If you get an error doing this, whinge about it _after_ reporting
	 * the originating error.  Also, clear the error on stdout to
	 * avoid getting caught in an infinite loop.
	 */
	star_eoln()
	errStdout := fflush_slowly(os.Stdout)

	/*
	 * Ask the system how wide the terminal is.
	 * Don't use last column, many terminals are dumb.
	 */
	page_width = page_width_get()

	_, _ = fmt.Fprintf(os.Stderr, "%s: %s\n", progname_get(), s)

	if err := fflush_slowly(os.Stderr); err != nil {
		/* don't print why, there is no point! */
		quit(1)
	}

	if errStdout != nil {
		nfatal_raw(errStdout, "standard output")
	}
}
