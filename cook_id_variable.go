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

type id_variable_ty struct {
	inherited id_ty
	value     string_list_ty
}

/*
 * NAME
 *      destructor
 *
 * SYNOPSIS
 *      void destructor(id_ty *);
 *
 * DESCRIPTION
 *      The destructor function is used to release the resources held by
 *      an ID instance.
 */

func destructor(idp *id_ty) {
	trace(fmt.Sprintf("id_variable::destructor(idp = %p)\n{\n", idp))
	this, ok := idp.(*id_variable_ty)
	assert(ok, "idp.(*id_variable_ty)")
	string_list_destructor(&this.value)
	trace("}\n")
}

/*
 * NAME
 *      evaluate
 *
 * SYNOPSIS
 *      int evaluate(id_ty *, const string_list_ty *, string_list_ty *);
 *
 * DESCRIPTION
 *      The evaluate function is used to evaluate an ID instance (there
 *      are several types).  The arguments to the evaluation are not to
 *      be changed, the results are only to be appended (not
 *      constructor'ed first).
 *
 * RETURNS
 *      int; 0 on success, -1 on error.
 */

func interpret(idp *id_ty, ocp *opcode_context_ty, pp *expr_position_ty) int {
	trace(fmt.Sprintf("id_variable::interpret(idp = %p)\n{\n", idp))
	this, ok := idp.(*id_variable_ty)
	assert(ok, "idp.(*id_variable_ty)")
	status := 0
	arg := opcode_context_string_list_pop(ocp)
	assert(len(arg.strings) >= 1, "len(arg.strings) >= 1")
	if len(arg.strings) >= 2 {
		scp := sub_context_new()
		sub_var_set_string(scp, "Name", arg.strings[0])
		error_with_position(pp, scp, i18n("$name: variable references no arguments"))
		sub_context_delete(scp)
		status = -1
	}
	arg = string_list_delete(arg)
	opcode_context_string_push_list(ocp, &this.value)
	trace("}\n")
	return status
}

/*
 * NAME
 *      method
 *
 * DESCRIPTION
 *      The method variable describes this ID class.
 *
 * CAVEAT
 *      This symbol is not to be exported from this file (its name is
 *      not unique).
 */

var method = id_method_ty{
	name:       "variable",
	size:       0, //todo: sizeof(id_variable_ty),
	destructor: destructor,
	interprets: interpret,
	script:     interpret, /* script */
}

/*
 * NAME
 *      id_variable_new
 *
 * SYNOPSIS
 *      void id_variable_new(void);
 *
 * DESCRIPTION
 *      The id_variable_new function is used to create a new instance of
 *      a variable ID's value.  The given value is copied.
 *
 * RETURNS
 *      id_ty *; a pointer to a ID instance is dynamic memory.
 *
 * CAVEAT
 *      Use id_instance_delete when you are done with it.
 */

func id_variable_new(slp *string_list_ty) *id_ty {
	trace("id_variable::new()\n{\n")
	idp := id_instance_new(&method)
	this, ok := idp.(*id_variable_ty)
	assert(ok, "idp.(*id_variable_ty)")
	string_list_copy_constructor(&this.value, slp)
	trace(fmt.Sprintf("return %p;\n", idp))
	trace("}\n")
	return idp
}
