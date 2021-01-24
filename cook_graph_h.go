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

type graph_ty struct {
	/*
	 * The try_list is a list of files that were not used, and
	 * subsequently backtracked.  This list is printed out in the
	 * error message when cook doesn't know how.
	 */
	try_list *string_list_ty

	/*
	 * Collect statistics about building the graph.  They were
	 * originally for debugging, but they look as impressive as
	 * hell, so I kept them.
	 */
	statistic struct {
		backtrack_bad_path                  long
		backtrack_by_ingredient             long
		backtrack_cache                     long
		error_by_ingredient                 long
		error_cache                         long
		error_in_expr                       long
		explicit_applicable                 long
		explicit_ingredients_applicable     long
		explicit_ingredients_not_applicable long
		explicit_not_applicable             long
		implicit_applicable                 long
		implicit_ingredients_applicable     long
		implicit_ingredients_not_applicable long
		implicit_not_applicable             long
		infinite_loop                       long
		inhibit_self_recursion              long
		leaf_error                          long
		leaf_backtrack                      long
		leaf_exists                         long
		pattern_match_query                 long
		phony                               long
		precondition_rejection              long
		success                             long
		success_reuse                       long
	}

	/*
	 * Symbol table of files already considered.
	 */
	already *symtab_ty

	/*
	 * The list of recipe instances used in this graph.
	 */
	already_recipe *graph_recipe_list_ty

	/*
	 * Used to remember file pairs when checking for essential
	 * information residing only in dependency files.
	 */
	file_pair *graph_file_pair_ty
}
