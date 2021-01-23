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
	"path"
	"path/filepath"
	"strings"
)

var progname string

// progname_fetch fetches the program name.
// It handles simple patterns like:
//      progname -args
//      ../bin/progname -args
//      /usr/local/progname -args
//      /usr//local///bin////progname///// -args
func progname_fetch() string {
	/*
	 * do NOT put tracing in this function
	 * do NOT put asserts in this function
	 *      they both depend on progname, which is not yet set
	 */

	s, err := os.Executable()
	if err != nil {
		s = "gcook"
	}
	if s = filepath.Base(s); s == "" || s == "." || s == "/" {
		s = "gcook"
	}

	// Nuke any suffix.
	if ext := path.Ext(s); ext != "" {
		s = strings.TrimSuffix(s, ext)
	}

	return s
}

// progname_get returns the global variable for the name of the program.
func progname_get() string {
	/* do NOT put tracing in this function */
	return progname
}

// progname_set sets the global variable for the name of the program.
func progname_set(s string) {
	/*
	 * do NOT put tracing in this function
	 * do NOT put asserts in this function
	 *      they both depend on progname, which is not yet set
	 */
	progname = s
}
