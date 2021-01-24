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

type graph_recipe_ty struct {
	reference_count long
	id              int
	rp              *recipe_ty
	mp              *match_ty
	input           *graph_file_list_nrc_ty
	output          *graph_file_list_nrc_ty
	input_satisfied size_t             /* used by graph_walk */
	input_uptodate  long               /* used by graph_walk */
	ocp             *opcode_context_ty /* used by graph_run */
	single_thread   *string_list_ty
	host_binding    *string_list_ty
	multi_forced    int /* used by graph_walk */
}
