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

var changed int64

var wstr_hash_table map[string]*wstring_ty

/*
 * NAME
 *      wstr_initialize - start up string table
 *
 * SYNOPSIS
 *      void wstr_initialize(void);
 *
 * DESCRIPTION
 *      The wstr_initialize function is used to create the hash table and
 *      initialize it to empty.
 *
 * RETURNS
 *      void
 *
 * CAVEAT
 *      This function must be called before any other defined in this file.
 */

func wstr_initialize() {
	wstr_hash_table = make(map[string]*wstring_ty)
}

/*
 * NAME
 *      wstr_equal - test equality of strings
 *
 * SYNOPSIS
 *      int wstr_equal(wstring_ty *, wstring_ty *);
 *
 * DESCRIPTION
 *      The wstr_equal function is used to test if two strings are equal.
 *
 * RETURNS
 *      int; zero if the strings are not equal, nonzero if the strings are
 *      equal.
 *
 * CAVEAT
 *      This function is implemented as a macro in strings.h
 */

func wstr_equal(ws1, ws2 *wstring_ty) bool {
	return ws1 == ws2
}

/*
 * NAME
 *      wstr_free - release a string
 *
 * SYNOPSIS
 *      void wstr_free(wstring_ty *s);
 *
 * DESCRIPTION
 *      The wstr_free function is used to indicate that a string hash been
 *      finished with.
 *
 * RETURNS
 *      void
 *
 * CAVEAT
 *      This is the only way to release strings DO NOT use the free function.
 */

func wstr_free(ws *wstring_ty) *wstring_ty {
	if ws == nil {
		return nil
	}
	if ws.wstr_references = ws.wstr_references - 1; ws.wstr_references > 0 {
		return nil
	}
	changed++

	// remove the wstring from the map
	delete(wstr_hash_table, ws.String())
	return nil
}

/*
 * NAME
 *      wstr_from_c - make string from C string
 *
 * SYNOPSIS
 *      wstring_ty *wstr_from_c(char *);
 *
 * DESCRIPTION
 *      The wstr_from_c function is used to make a string from a NUL
 *      terminated C string.  The conversion from multi-byte to wide
 *      characters is done in the current locale.
 *
 * RETURNS
 *      wstring_ty* - a pointer to a string in dynamic memory.  Use
 *      wstr_free when finished with.
 *
 * CAVEAT
 *      The contents of the structure pointed to MUST NOT be altered.
 */

func wstr_from_c(s []byte, length size_t) *wstring_ty {
	return wstr_n_from_slice(s, length)
}

func wstr_from_string(s string) *wstring_ty {
	return wstr_n_from_slice([]byte(s), size_t(len(s)))
}

/*
 * NAME
 *      wstr_n_from_c - make string
 *
 * SYNOPSIS
 *      wstring_ty *wstr_n_from_c(char *s, size_t n);
 *
 * DESCRIPTION
 *      The wstr_n_from_c function is used to make a string from an
 *      array of characters.  No NUL terminator is assumed.  The
 *      conversion from muti-byte to wide characters is done in the
 *      current locale.
 *
 * RETURNS
 *      wstring_ty* - a pointer to a string in dynamic memory.  Use
 *      wstr_free when finished with.
 *
 * CAVEAT
 *      The contents of the structure pointed to MUST NOT be altered.
 */

func wstr_n_from_slice(s []byte, length size_t) *wstring_ty {
	if ws, ok := wstr_hash_table[string(s)]; ok {
		return ws
	}
	ws := &wstring_ty{
		wstr_length: size_t(length),
		wstr_text:   make([]byte, length, length),
	}
	copy(ws.wstr_text, s)
	ws.str = string(ws.wstr_text)
	return ws
}

func str_to_wstr(s *string_ty) *wstring_ty {
	return wstr_n_from_slice(s.str_text, s.str_length)
}

func wstr_to_str(ws *wstring_ty) *string_ty {
	var text *char
	var length size_t

	wstr_to_mbs(ws, &text, &length)
	return str_from_string(ws.String())
}

/*
 * NAME
 *      wstr_to_mbs - wide string to multi-byte C string
 *
 * SYNOPSIS
 *      void wstr_to_mbs(wstring_ty *s, char **rslt, size_t *rslt_len);
 *
 * DESCRIPTION
 *      The wstr_to_mbs function convers a wide character string into a
 *      multi-byte C string.  The conversion is done in the current
 *      locale.  The result is NUL terminated, however the result length
 *      does not include the NUL.
 *
 * CAVEAT
 *      DO NOT free the result.  The result will change between calls,
 *      so copy it if you need to keep it.
 */

func wstr_to_mbs(ws *wstring_ty, result_p **char, result_length_p *size_t) string {
	panic("!implemented")
}

func (ws *wstring_ty) String() string {
	return ws.str
}
