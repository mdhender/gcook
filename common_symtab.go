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
 * It is important to preserve the order of the links because
 * they can be push-down stacks, and to simply add them to the
 * head of the list will reverse the order of the stack!
 */

func symtab_alloc(size int) *symtab_ty {
	trace(fmt.Sprintf("symtab_alloc(size = %d)\n{\n", size))
	stp := &symtab_ty{
		hash_table: make(map[string][]*symtab_row_ty),
	}
	trace(fmt.Sprintf("return %p;\n", stp))
	trace("}\n")
	return stp
}

/*
 * NAME
 *      symtab_assign - assign a variable
 *
 * SYNOPSIS
 *      void symtab_assign(symtab_ty *, string_ty *key, void *data);
 *
 * DESCRIPTION
 *      The symtab_assign function is used to assign
 *      a value to a given variable.
 *
 * CAVEAT
 *      The name is copied, the data is not.
 */

func symtab_assign(stp *symtab_ty, key *string_ty, data interface{}) {
	trace(fmt.Sprintf("symtab_assign(stp = %p, key = %q, data = %p)\n{\n", stp, key.str_text, data))

	modifyExisting := false
	for _, row := range stp.hash_table[key.String()] {
		if str_equal(key, row.key) {
			trace("modify existing entry\n")
			if stp.reap != nil {
				stp.reap(row.data)
			}
			row.data = data
			modifyExisting = true
			break
		}
	}

	if !modifyExisting {
		trace("new entry\n")
		p := &symtab_row_ty{} // mem_alloc(sizeof(symtab_row_ty));
		p.key = str_copy(key)
		p.data = data
		stp.hash_table[p.key.String()] = append(stp.hash_table[p.key.String()], p)
	}

	trace("}\n")
}

func symtab_free(stp *symtab_ty) *symtab_ty {
	trace(fmt.Sprintf("symtab_free(stp = %p)\n{\n", stp))
	stp.hash_table = nil // mem_free(stp.hash_table);
	trace("}\n")
	return nil
}
