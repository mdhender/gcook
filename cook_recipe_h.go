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

type recipe_ty struct {
	reference_count long
	target          *string_list_ty
	need1           *opcode_list_ty
	need2           *opcode_list_ty
	flags           *flag_ty
	multiple        int
	precondition    *opcode_list_ty
	single_thread   *opcode_list_ty
	host_binding    *opcode_list_ty
	out_of_date     *opcode_list_ty
	up_to_date      *opcode_list_ty
	pos             expr_position_ty /* for tracing and debugging */
	implicit        int
	inhibit         int /* for graph generation */
}
