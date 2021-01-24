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

type graph_file_and_type_ty struct {
	file      *graph_file_ty
	edge_type edge_type_ty
}

type graph_file_list_ty struct {
	nfiles     size_t
	nfiles_max size_t
	item       *graph_file_and_type_ty
}

/*
 * again, this time without touching the reference counts...
 */
type graph_file_list_nrc_ty struct {
	nfiles     size_t
	nfiles_max size_t
	item       *graph_file_and_type_ty
}
