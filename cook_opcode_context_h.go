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

type opcode_frame_ty struct {
	olp *opcode_list_ty
	pc  size_t
	sp  *symtab_ty
}

type long = int64

type opcode_context_ty struct {
	call_stack_length   size_t
	call_stack_maximum  size_t
	call_stack          opcode_frame_ty
	value_stack_length  size_t
	value_stack_maximum size_t
	value_stack         []*string_list_ty
	thread_id           long

	thread_stp *symtab_ty
	msp        *match_stack_ty

	pid         int         /* used by opcode_command */
	exit_status int         /* used by opcode_command */
	meter_p     meter_ty    /* used by opcode_command */
	wlp         interface{} /* used by opcode_command */ // was void *
	need_age    int         /* used by graph_run */

	/* for suspend/resume */
	flags        interface{} // was void *
	mp           *match_ty
	host_binding *string_ty

	/* for information about the graph */
	gp *graph_ty
}
