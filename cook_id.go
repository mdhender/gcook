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

var id_need *string_ty
var id_younger *string_ty
var id_target *string_ty
var id_targets *string_ty
var id_search_list *string_ty

/*
 * NAME
 *      id_initialize - start up symbol table
 *
 * SYNOPSIS
 *      void id_initialize(void);
 *
 * DESCRIPTION
 *      The id_initialize function is used to create the hash table.
 *
 * RETURNS
 *      void
 *
 * CAVEAT
 *      Assumes the str_initialize function has been called already.
 */

func id_initialize() {
	trace("init\n")

	id_need = str_from_c("need")
	id_younger = str_from_c("younger")
	id_target = str_from_c("target")
	id_targets = str_from_c("targets")
	id_search_list = str_from_c("search_list")

	id_reset()
}

func id_reset() {
	id_global_reset()

	/*
	 * set the "version" predefined variable
	 */
	var wl string_list_ty
	string_list_constructor(&wl)
	s := str_from_c(version_stamp())
	string_list_append(&wl, s)
	s = str_free(s)
	s = str_from_c("version")
	symtab_assign(id_global_stp(), s, id_variable_new(&wl))
	s = str_free(s)
	string_list_destructor(&wl)

	panic("!")
}
