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
 *      mem_alloc - allocate and clear memory
 *
 * SYNOPSIS
 *      char *mem_alloc(size_t n);
 *
 * DESCRIPTION
 *      Mem_alloc uses malloc to allocate the required sized chunk of memory.
 *      If any error is returned from malloc() an fatal diagnostic is issued.
 *      The memory is zeroed befor it is returned.
 *
 * CAVEAT
 *      It is the responsibility of the caller to ensure that the space is
 *      freed when finished with, by a call to free().
 */

func mem_alloc(i interface{}) **string_ty {
	panic("!")
}

func mem_free(p interface{}) {
	// noop
}
