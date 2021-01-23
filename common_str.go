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
 * Strings are the most heavily used resource in cook.  They are manipulated
 * inside the match functions, and hence are in the inside loop.  For this
 * reason they must be fast.
 *
 * A literal pool is maintained.  Each string has a reference count.  The
 * string stays in the literal pool for as long as it has a positive
 * reference count.  To determine if a string is already in the literal pool,
 * linear dynamic hashing is used to guarantee an O(1) search.  That all equal
 * strings are the same item in the literal pool means that string equality is
 * a pointer test, and thus very fast.
 */

/*
 * NAME
 *       str_initialize - start up string table
 *
 * SYNOPSIS
 *       void str_initialize(void);
 *
 * DESCRIPTION
 *       The str_initialize function is used to create the hash table and
 *       initialize it to empty.
 *
 * RETURNS
 *       void
 *
 * CAVEAT
 *       This function must be called before any other defined in this file.
 */

func str_initialize() {
	hash_table = make(map[string]*string_ty)
	str_true = str_from_c("1")
	str_false = str_from_c("")
}

func mem_alloc(i interface{}) **string_ty {
	panic("!")
}

var hash_table map[string]*string_ty

var str_true *string_ty
var str_false *string_ty

/*
 * NAME
 *      str_from_c - make string from C string
 *
 * SYNOPSIS
 *      string_ty *str_from_c(char*);
 *
 * DESCRIPTION
 *      The str_from_c function is used to make a string from a null terminated
 *      C string.
 *
 * RETURNS
 *      string_ty * - a pointer to a string in dynamic memory.
 *      Use str_free when finished with.
 *
 * CAVEAT
 *      The contents of the structure pointed to MUST NOT be altered.
 */

func str_from_c(s string) *string_ty {
	return str_n_from_slice([]byte(s), len(s))
}

/*
 * NAME
 *      str_n_from_c - make string
 *
 * SYNOPSIS
 *      string_ty *str_n_from_c(char *s, size_t n);
 *
 * DESCRIPTION
 *      The str_n_from_c function is used to make a string from an array of
 *      characters.  No null terminator is assumed.
 *
 * RETURNS
 *      string_ty * - a pointer to a string in dynamic memory.
 *      Use str_free when finished with.
 *
 * CAVEAT
 *      The contents of the structure pointed to MUST NOT be altered.
 */

func str_n_from_slice(s []byte, length int) *string_ty {
	if n, ok := hash_table[string(s)]; ok {
		return n
	}
	n := &string_ty{
		str_length: size_t(length),
		str_text:   make([]byte, length, length),
	}
	copy(n.str_text, s)
	n.str = string(n.str_text)
	return n
}

func str_from_slice(s []byte, length int) *string_ty {
	return str_n_from_slice(s, length)
}
