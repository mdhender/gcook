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
 * NAME
 *      opcode_context_string_list_pop
 *
 * SYNOPSIS
 *      string_list_ty *opcode_context_string_list_pop(opcode_context_ty *);
 *
 * DESCRIPTION
 *      The opcode_context_string_list_pop function is used to obtain
 *      the top-most string list from the value stack (it is removed
 *      from the stack).  This is the normal mechanism for obtaining
 *      argument lists.  Use string_list_delete when you are done with it.
 *
 * CAVEAT
 *      To be used only internally to the interpratation by individial
 *      opcodes.
 */

func opcode_context_string_list_pop(ocp *opcode_context_ty) *string_list_ty {
	trace(fmt.Sprintf("opcode_context_string_list_pop(ocp = %p)\n{\n", ocp))
	assert(ocp != nil, "ocp != nil")
	assert(ocp.value_stack_length > 0, "ocp.value_stack_length > 0")
	ocp.value_stack_length--
	slp := ocp.value_stack[ocp.value_stack_length]
	trace(fmt.Sprintf("return %p;\n", slp))
	trace("}\n")
	return slp
}

func opcode_context_string_push_list(ocp *opcode_context_ty, i *string_list_ty) {
	trace(fmt.Sprintf("opcode_context_string_push_list(ocp = %p)\n{\n", ocp))
	assert(ocp != nil, "ocp != nil")
	assert(ocp.value_stack_length > 0, "ocp.value_stack_length > 0")
	slp := ocp.value_stack[ocp.value_stack_length-1]
	string_list_append_list(slp, i)
	trace("}\n")
}
