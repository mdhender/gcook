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

func id_instance_delete(idp *id_ty) {
	assert(idp != nil, "idp != nil")
	assert(idp.method != nil, "idp.method != nil")
	assert(idp.method.destructor != nil, "idp.method.destructor")
	idp.method.destructor(idp)
	idp.method = nil /* paranoia */
	idp = nil        // mem_free(idp)
}

func id_instance_new(mp *id_method_ty) *id_ty {
	trace("id_new()\n{\n")
	assert(mp != nil, "mp != nil")
	trace(fmt.Sprintf("is a %q\n", mp.name))
	idp := &id_ty{} // mem_alloc(mp.size);
	idp.method = mp
	trace(fmt.Sprintf("return %p;\n", idp))
	trace("}\n")
	return idp
}
