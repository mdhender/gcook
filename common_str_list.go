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
 * NAME
 *      string_list_append - append to a word list
 *
 * SYNOPSIS
 *      void string_list_append(string_list_ty *wlp, string_ty *wp);
 *
 * DESCRIPTION
 *      Wl_append is used to append to a word list.
 *
 * CAVEAT
 *      The word being appended IS copied.
 */

func string_list_append(wlp *string_list_ty, w *string_ty) {
	assert(wlp != nil, "wlp != nil")
	assert(w != nil, "w != nil")
	wlp.strings = append(wlp.strings, str_copy(w))
}

/*
 * NAME
 *      string_list_append_list
 *
 * SYNOPSIS
 *      void string_list_append_list(string_list_ty *to, string_list_ty *from);
 *
 * DESCRIPTION
 *      The string_list_append_list function is used to append one
 *      string list (from) onto the end of another (to).
 */

func string_list_append_list(to *string_list_ty, from *string_list_ty) {
	for _, str := range from.strings {
		string_list_append(to, str)
	}
}

/*
 * NAME
 *      string_list_constructor
 *
 * SYNOPSIS
 *      void string_list_constructor(string_list_ty *);
 *
 * DESCRIPTION
 *      The string_list_constructor function is used to prepare a string
 *      list for use.  It will be empty.
 *
 * CAVEAT
 *      This must be called on the string list before any other action
 *      is taken.  Use string_list_destructor when you are done.
 */

func string_list_constructor(wlp *string_list_ty) {
	// todo: should this force a str_free of wlp.strings?
	wlp.strings = nil
}

/*
 * NAME
 *      string_list_copy_constructor - copy a word list
 *
 * SYNOPSIS
 *      void string_list_copy_constructor(string_list_ty *to,
 *              string_list_ty *from);
 *
 * DESCRIPTION
 *      Wl_copy is used to copy word lists.
 *
 * RETURNS
 *      A copy of the 'to' word list is placed in 'from'.
 *
 * CAVEAT
 *      It is the responsibility of the caller to ensure that the
 *      new word list is freed when finished with, by a call to
 *      string_list_destructor().
 */

func string_list_copy_constructor(to, from *string_list_ty) {
	string_list_constructor(to)
	for _, str := range from.strings {
		string_list_append(to, str_copy(str))
	}
}

/*
 * NAME
 *      string_list_delete
 *
 * SYNOPSIS
 *      void string_list_delete(string_list_ty *);
 *
 * DESCRIPTION
 *      The string_list_delete function is used to release the resources
 *      held by a string list in dynamic memory.
 */

func string_list_delete(slp *string_list_ty) *string_list_ty {
	string_list_destructor(slp)
	return nil
}

/*
 * NAME
 *      string_list_destructor - free a word list
 *
 * SYNOPSIS
 *      void string_list_destructor(string_list_ty *wlp);
 *
 * DESCRIPTION
 *      Wl_free is used to free the contents of a word list
 *      when it is finished with.
 *
 * CAVEAT
 *      It is assumed that the contents of the word list were all
 *      created using strdup() or similar, and grown using string_list_append().
 */

func string_list_destructor(wlp *string_list_ty) {
	// call str_free against all of the strings to keep the references accurate
	for _, s := range wlp.strings {
		str_free(s)
	}
	wlp.strings = nil
}
