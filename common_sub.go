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

type fp func(*sub_context_ty, *wstring_list_ty) *wstring_ty

type table_ty struct {
	name             string
	fp               fp
	resubstitute     bool
	value            *wstring_ty
	must_be_used     bool
	append_if_unused bool
	override         bool
}

type diversion_ty struct {
	pos          size_t
	text         *wstring_ty
	prev         *diversion_ty
	resubstitute int
}

type collect_ty struct {
	pos  size_t
	size size_t
	buf  *wchar_t
}

type sub_context_ty struct {
	diversion       *diversion_ty
	sub_var_list    []*table_ty
	sub_var_size    size_t
	sub_var_pos     size_t
	suberr          *char
	errno_sequester int
}

func sub_context_new() *sub_context_ty {
	scp := &sub_context_ty{}
	sub_context_constructor(scp)
	return scp
}

func sub_context_constructor(scp *sub_context_ty) {
	trace("sub_context_constructor()\n{\n")
	scp.diversion = nil
	scp.sub_var_list = nil
	scp.sub_var_size = 0
	scp.sub_var_pos = 0
	scp.suberr = nil
	scp.errno_sequester = 0
	trace("}\n")
}

func sub_context_delete(scp *sub_context_ty) *sub_context_ty {
	sub_context_destructor(scp)
	return nil // mem_free(scp);
}

func sub_context_destructor(scp *sub_context_ty) {
	for _, o := range scp.sub_var_list {
		wstr_free(o.value)
	}
	scp.sub_var_list = nil // mem_free(scp.sub_var_list)

	scp.diversion = nil
	scp.sub_var_list = nil
	scp.sub_var_size = 0
	scp.sub_var_pos = 0
	scp.suberr = nil
	scp.errno_sequester = 0
}

/*
 * NAME
 *      sub_var_set
 *
 * SYNOPSIS
 *      void sub_var_set(char *name, char *fmt, ...);
 *
 * DESCRIPTION
 *      The sub_var_set function is used to set the value of a
 *      substitution variable.  These variables are command specific,
 *      as opposed to the functions which are always present.
 *      The user documentation does NOT make this distinction by
 *      using the names "variable" and "function", they are always referred
 *      to as "substitutions".
 *
 * ARGUMENTS
 *      name    - the name of the variable
 *      fmt,... - a format string and arguments to construct the value.
 *                Handed to str_vformat to make a (string_ty *) out of it.
 *
 * CAVEAT
 *      remains in scope until the next invokation of sub_var_clear,
 *      or until the end of the next invokation of substitute.
 */

func sub_var_set(scp *sub_context_ty, name string, format string, a ...interface{}) {
	trace(fmt.Sprintf("sub_var_set(scp = %p, name = %q)\n{\n", scp, name))
	s := str_vformat(format, a...)
	sub_var_set_string(scp, name, s)
	str_free(s)
	trace("}\n")
}

func sub_var_set_long(scp *sub_context_ty, name string, value long) {
	trace(fmt.Sprintf("sub_var_set_long(scp = %p, name = \"%s\", value = %ld)\n{\n", scp, name, value))
	sub_var_set(scp, name, "%ld", value)
	trace("}\n")
}

func sub_var_set_string(scp *sub_context_ty, name string, value *string_ty) {
	trace(fmt.Sprintf("sub_var_set_string(scp = %p, name = %q, value = %q)\n{\n", scp, name, value))

	scp.sub_var_pos++

	svp := &table_ty{}
	svp.name = name
	svp.fp = nil
	svp.value = str_to_wstr(value)
	svp.must_be_used = true
	svp.append_if_unused = false
	svp.override = false
	svp.resubstitute = !svp.must_be_used
	trace("}\n")
}

func subst(scp *sub_context_ty, s *wstring_ty) *wstring_ty {
	panic("!implemented")
}

func subst_intl(scp *sub_context_ty, s string) *string_ty {
	trace(fmt.Sprintf("subst_intl(scp = %p, s = %q)\n{\n", scp, s))
	result_wide := subst_intl_wide(scp, s)
	result := wstr_to_str(result_wide)
	wstr_free(result_wide)
	trace(fmt.Sprintf("return %q;\n", result.str_text))
	trace("}\n")
	return result
}

func subst_intl_wide(scp *sub_context_ty, msg string) *wstring_ty {
	trace(fmt.Sprintf("subst_intl_wide(scp = %p, msg = %q)\n{\n", scp, msg))
	language_human()
	tmp := gettext(msg)
	language_C()
	s := wstr_from_string(tmp)
	result := subst(scp, s)
	wstr_free(s)
	trace(fmt.Sprintf("return %p;\n", result))
	trace("}\n")
	return result
}
