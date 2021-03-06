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

type str_hash_ty = uint64
type size_t = int64

type string_ty struct {
	str_hash       str_hash_ty
	str_next       *string_ty
	str_references int64
	str_length     size_t
	str_text       []byte
	str            string
}

func (s *string_ty) String() string {
	return s.str
}
