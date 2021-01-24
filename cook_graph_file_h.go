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

type graph_file_ty struct {
	reference_count    long
	filename           string_ty
	input              graph_recipe_list_nrc_ty
	output             *graph_recipe_list_nrc_ty
	pending            long
	previous_backtrack int
	previous_error     int
	mtime_oldest       long   /* used by graph_recipe_run */
	input_satisfied    size_t /* used by graph_walk */
	done               long   /* used by graph_walk */
	input_uptodate     size_t /* used by graph_walk */
	primary_target     int
}
