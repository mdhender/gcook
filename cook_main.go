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
	"github.com/mdhender/gcook/internal/signals"
	"os"
)

/*
 * NAME
 *      main - initial entry point for cook
 *
 * SYNOPSIS
 *      void main(int argc, char **argv);
 *
 * DESCRIPTION
 *      Main is the initial entry point for cook.
 *
 * RETURNS
 *      Exit is always through exit().
 *      The exit code will be 0 for success, or 1 for some error.
 */

func main() {
	var retval int

	/*
	 * Some versions of cron(8) and at(1) set SIGCHLD to SIG_IGN.
	 * This is kinda dumb, because it breaks assumptions made in
	 * libc (like pclose, for instance).  It also blows away most
	 * of Cook's process handling.  We explicitly set the SIGCHLD
	 * signal handling to SIG_DFL to make sure this signal does what
	 * we expect no matter how we are invoked.
	 */
	signals.Signal("SIGCHLD", "SIG_DFL")

	/*
	 * initialize things
	 * (order is critical here)
	 */
	progname_set(progname_fetch())
	fmt.Println(progname_get())
	str_initialize()
	id_initialize()

	os.Exit(retval)
}
