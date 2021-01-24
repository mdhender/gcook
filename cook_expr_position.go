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
 *      error_with_position
 *
 * SYNOPSIS
 *      void error_with_position(expr_position_ty *, char *);
 *
 * DESCRIPTION
 *      The error_with_position function is used to report an error at a
 *      given location.  The arguments should be set using sub_var_set,
 *      as the string will be passed through the internationalized error
 *      functions.
 */

func error_with_position(pp *expr_position_ty, scp *sub_context_ty, fmt string) {
	need_to_delete := scp != nil

	if scp == nil {
		scp = sub_context_new()
	}

	s := subst_intl(scp, fmt)

	/* re-use the substitution context */
	if pp != nil && pp.pos_name != nil && pp.pos_line != 0 {
		sub_var_set_string(scp, "File_Name", pp.pos_name)
		sub_var_set_long(scp, "Number", pp.pos_line)
		sub_var_set_string(scp, "MeSsaGe", s)
		error_intl(scp, i18n("$filename: $number: $message"))
		str_free(s)
	} else {
		sub_var_set_string(scp, "MeSsaGe", s)
		error_intl(scp, i18n("$message"))
		str_free(s)
	}

	if need_to_delete {
		sub_context_delete(scp)
	}
}
