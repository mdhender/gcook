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

type char = byte

type match_method_ty struct {
	name            *char
	size            int
	destructor      func(*match_ty)
	constructor     func(*match_ty)
	compile         func(*match_ty, *string_ty, *expr_position_ty) int
	execute         func(*match_ty, *string_ty, *expr_position_ty) int
	reconstruct_lhs func(*match_ty, *string_ty, *expr_position_ty) *string_ty
	reconstruct_rhs func(*match_ty, *string_ty, *expr_position_ty) *string_ty
	usage_mask      func(*match_ty, *string_ty, *expr_position_ty) int
}
