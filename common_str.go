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

import "fmt"

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

// #define MAX_HASH_LEN 20
const MAX_HASH_LEN = 20

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
	str_true = str_from_string("1")
	str_false = str_from_string("")
}

/*
 * NAME
 *      str_copy - make a copy of a string
 *
 * SYNOPSIS
 *      string_ty *str_copy(string_ty *s);
 *
 * DESCRIPTION
 *      The str_copy function is used to make a copy of a string.
 *
 * RETURNS
 *      string_ty * - a pointer to a string in dynamic memory.
 *      Use str_free when finished with.
 *
 * CAVEAT
 *      The contents of the structure pointed to MUST NOT be altered.
 */

func str_copy(s *string_ty) *string_ty {
	s.str_references++
	return s
}

/*
 * NAME
 *      str_equal - test equality of strings
 *
 * SYNOPSIS
 *      int str_equal(string_ty *, string_ty *);
 *
 * DESCRIPTION
 *      The str_equal function is used to test if two strings are equal.
 *
 * RETURNS
 *      int; zero if the strings are not equal, nonzero if the strings are
 *      equal.
 *
 * CAVEAT
 *      This function is implemented as a macro in strings.h
 */

func str_equal(s1, s2 *string_ty) bool {
	return s1 == s2
}

/*
 * NAME
 *      str_free - release a string
 *
 * SYNOPSIS
 *      void str_free(string_ty *s);
 *
 * DESCRIPTION
 *      The str_free function is used to indicate that a string hash been
 *      finished with.
 *
 * RETURNS
 *      void
 *
 * CAVEAT
 *      This is the only way to release strings DO NOT use the free function.
 */

func str_free(s *string_ty) *string_ty {
	assert(str_valid(s), "str_valid(s)")
	if s.str_references = s.str_references - 1; s.str_references > 0 {
		return nil
	}
	// remove the string from the map
	delete(hash_table, s.str)
	return nil
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

func str_from_c(s []byte) *string_ty {
	return str_n_from_c(s, len(s))
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

func str_n_from_c(s []byte, length int) *string_ty {
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

func str_from_string(s string) *string_ty {
	return str_n_from_c([]byte(s), len(s))
}

/*
 * NAME
 *      str_valid - test a string
 *
 * SYNOPSIS
 *      int str_valid(string_ty *s);
 *
 * DESCRIPTION
 *      The str_valid function is used to test if a pointer points to a valid
 *      string.
 *
 * RETURNS
 *      int: zero if the string is not valid, nonzero if the string is valid.
 *
 * CAVEAT
 *      This function is only available then the DEBUG symbol is #define'd.
 */

func str_valid(s *string_ty) bool {
	return s != nil && s.str_references > 0 && strlen(s.str_text) == s.str_length && s.str_hash == hash_generate(s.str_text, s.str_length)
}

/*
 * NAME
 *      str_format - analog of sprintf
 *
 * SYNOPSIS
 *      string_ty *str_format(char *, ...);
 *
 * DESCRIPTION
 *      The str_format function is used to create new strings
 *      using a format specification similar to printf(3).
 *
 * RETURNS
 *      string_ty * - a pointer to a string in dynamic memory.
 *      Use str_free when finished with.
 */

func str_format(format string, a ...interface{}) *string_ty {
	return str_from_string(fmt.Sprintf(format, a...))
}

func str_vformat(format string, a ...interface{}) *string_ty {
	return str_from_string(fmt.Sprintf(format, a...))
}

/*
 * NAME
 *      hash_generate - hash string to number
 *
 * SYNOPSIS
 *      str_hash_ty hash_generate(char *s, size_t n);
 *
 * DESCRIPTION
 *      The hash_generate function is used to make a number from a string.
 *
 * RETURNS
 *      str_hash_ty - the magic number
 *
 * CAVEAT
 *      Only the last MAX_HASH_LEN characters are used.
 *      It is important that str_hash_ty be unsigned (int or long).
 */

func hash_generate(b []byte, n size_t) (hashval str_hash_ty) {
	if n > MAX_HASH_LEN {
		b = b[n-MAX_HASH_LEN:]
		n = MAX_HASH_LEN
	}
	for i := 0; i < int(n); i++ {
		hashval = (hashval + (hashval << 1)) ^ uint64(b[i])
	}
	return hashval
}
