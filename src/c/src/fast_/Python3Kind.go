// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fast_

type Python3Kind = int32
const (
	Python3Kindsingle_input Python3Kind = 0
	Python3Kindfile_input Python3Kind = 1
	Python3Kindeval_input Python3Kind = 2
	Python3Kinddecorator Python3Kind = 3
	Python3Kinddecorators Python3Kind = 4
	Python3Kinddecorated Python3Kind = 5
	Python3Kindasync_funcdef Python3Kind = 6
	Python3Kindfuncdef Python3Kind = 7
	Python3Kindparameters Python3Kind = 8
	Python3Kindtypedargslist Python3Kind = 9
	Python3Kindtfpdef Python3Kind = 10
	Python3Kindvarargslist Python3Kind = 11
	Python3Kindvfpdef Python3Kind = 12
	Python3Kindstmt Python3Kind = 13
	Python3Kindsimple_stmt Python3Kind = 14
	Python3Kindsmall_stmt Python3Kind = 15
	Python3Kindexpr_stmt Python3Kind = 16
	Python3Kindannassign Python3Kind = 17
	Python3Kindtestlist_star_expr Python3Kind = 18
	Python3Kindaugassign Python3Kind = 19
	Python3Kinddel_stmt Python3Kind = 20
	Python3Kindpass_stmt Python3Kind = 21
	Python3Kindflow_stmt Python3Kind = 22
	Python3Kindbreak_stmt Python3Kind = 23
	Python3Kindcontinue_stmt Python3Kind = 24
	Python3Kindreturn_stmt Python3Kind = 25
	Python3Kindyield_stmt Python3Kind = 26
	Python3Kindraise_stmt Python3Kind = 27
	Python3Kindimport_stmt Python3Kind = 28
	Python3Kindimport_name Python3Kind = 29
	Python3Kindimport_from Python3Kind = 30
	Python3Kindimport_as_name Python3Kind = 31
	Python3Kinddotted_as_name Python3Kind = 32
	Python3Kindimport_as_names Python3Kind = 33
	Python3Kinddotted_as_names Python3Kind = 34
	Python3Kinddotted_name Python3Kind = 35
	Python3Kindglobal_stmt Python3Kind = 36
	Python3Kindnonlocal_stmt Python3Kind = 37
	Python3Kindassert_stmt Python3Kind = 38
	Python3Kindcompound_stmt Python3Kind = 39
	Python3Kindasync_stmt Python3Kind = 40
	Python3Kindif_stmt Python3Kind = 41
	Python3Kindwhile_stmt Python3Kind = 42
	Python3Kindfor_stmt Python3Kind = 43
	Python3Kindtry_stmt Python3Kind = 44
	Python3Kindwith_stmt Python3Kind = 45
	Python3Kindwith_item Python3Kind = 46
	Python3Kindexcept_clause Python3Kind = 47
	Python3Kindsuite Python3Kind = 48
	Python3Kindtest Python3Kind = 49
	Python3Kindtest_nocond Python3Kind = 50
	Python3Kindlambdef Python3Kind = 51
	Python3Kindlambdef_nocond Python3Kind = 52
	Python3Kindor_test Python3Kind = 53
	Python3Kindand_test Python3Kind = 54
	Python3Kindnot_test Python3Kind = 55
	Python3Kindcomparison Python3Kind = 56
	Python3Kindcomp_op Python3Kind = 57
	Python3Kindstar_expr Python3Kind = 58
	Python3Kindexpr Python3Kind = 59
	Python3Kindxor_expr Python3Kind = 60
	Python3Kindand_expr Python3Kind = 61
	Python3Kindshift_expr Python3Kind = 62
	Python3Kindarith_expr Python3Kind = 63
	Python3Kindterm Python3Kind = 64
	Python3Kindfactor Python3Kind = 65
	Python3Kindpython_power Python3Kind = 66
	Python3Kindatom_expr Python3Kind = 67
	Python3Kindatom Python3Kind = 68
	Python3Kindtestlist_comp Python3Kind = 69
	Python3Kindtrailer Python3Kind = 70
	Python3Kindsubscriptlist Python3Kind = 71
	Python3Kindpython_subscript Python3Kind = 72
	Python3Kindsliceop Python3Kind = 73
	Python3Kindexprlist Python3Kind = 74
	Python3Kindtestlist Python3Kind = 75
	Python3Kinddictorsetmaker Python3Kind = 76
	Python3Kindclassdef Python3Kind = 77
	Python3Kindarglist Python3Kind = 78
	Python3Kindpython_argument Python3Kind = 79
	Python3Kindcomp_iter Python3Kind = 80
	Python3Kindcomp_for Python3Kind = 81
	Python3Kindcomp_if Python3Kind = 82
	Python3Kindencoding_decl Python3Kind = 83
	Python3Kindyield_expr Python3Kind = 84
	Python3Kindyield_arg Python3Kind = 85
)

var EnumNamesPython3Kind = map[Python3Kind]string{
	Python3Kindsingle_input:"single_input",
	Python3Kindfile_input:"file_input",
	Python3Kindeval_input:"eval_input",
	Python3Kinddecorator:"decorator",
	Python3Kinddecorators:"decorators",
	Python3Kinddecorated:"decorated",
	Python3Kindasync_funcdef:"async_funcdef",
	Python3Kindfuncdef:"funcdef",
	Python3Kindparameters:"parameters",
	Python3Kindtypedargslist:"typedargslist",
	Python3Kindtfpdef:"tfpdef",
	Python3Kindvarargslist:"varargslist",
	Python3Kindvfpdef:"vfpdef",
	Python3Kindstmt:"stmt",
	Python3Kindsimple_stmt:"simple_stmt",
	Python3Kindsmall_stmt:"small_stmt",
	Python3Kindexpr_stmt:"expr_stmt",
	Python3Kindannassign:"annassign",
	Python3Kindtestlist_star_expr:"testlist_star_expr",
	Python3Kindaugassign:"augassign",
	Python3Kinddel_stmt:"del_stmt",
	Python3Kindpass_stmt:"pass_stmt",
	Python3Kindflow_stmt:"flow_stmt",
	Python3Kindbreak_stmt:"break_stmt",
	Python3Kindcontinue_stmt:"continue_stmt",
	Python3Kindreturn_stmt:"return_stmt",
	Python3Kindyield_stmt:"yield_stmt",
	Python3Kindraise_stmt:"raise_stmt",
	Python3Kindimport_stmt:"import_stmt",
	Python3Kindimport_name:"import_name",
	Python3Kindimport_from:"import_from",
	Python3Kindimport_as_name:"import_as_name",
	Python3Kinddotted_as_name:"dotted_as_name",
	Python3Kindimport_as_names:"import_as_names",
	Python3Kinddotted_as_names:"dotted_as_names",
	Python3Kinddotted_name:"dotted_name",
	Python3Kindglobal_stmt:"global_stmt",
	Python3Kindnonlocal_stmt:"nonlocal_stmt",
	Python3Kindassert_stmt:"assert_stmt",
	Python3Kindcompound_stmt:"compound_stmt",
	Python3Kindasync_stmt:"async_stmt",
	Python3Kindif_stmt:"if_stmt",
	Python3Kindwhile_stmt:"while_stmt",
	Python3Kindfor_stmt:"for_stmt",
	Python3Kindtry_stmt:"try_stmt",
	Python3Kindwith_stmt:"with_stmt",
	Python3Kindwith_item:"with_item",
	Python3Kindexcept_clause:"except_clause",
	Python3Kindsuite:"suite",
	Python3Kindtest:"test",
	Python3Kindtest_nocond:"test_nocond",
	Python3Kindlambdef:"lambdef",
	Python3Kindlambdef_nocond:"lambdef_nocond",
	Python3Kindor_test:"or_test",
	Python3Kindand_test:"and_test",
	Python3Kindnot_test:"not_test",
	Python3Kindcomparison:"comparison",
	Python3Kindcomp_op:"comp_op",
	Python3Kindstar_expr:"star_expr",
	Python3Kindexpr:"expr",
	Python3Kindxor_expr:"xor_expr",
	Python3Kindand_expr:"and_expr",
	Python3Kindshift_expr:"shift_expr",
	Python3Kindarith_expr:"arith_expr",
	Python3Kindterm:"term",
	Python3Kindfactor:"factor",
	Python3Kindpython_power:"python_power",
	Python3Kindatom_expr:"atom_expr",
	Python3Kindatom:"atom",
	Python3Kindtestlist_comp:"testlist_comp",
	Python3Kindtrailer:"trailer",
	Python3Kindsubscriptlist:"subscriptlist",
	Python3Kindpython_subscript:"python_subscript",
	Python3Kindsliceop:"sliceop",
	Python3Kindexprlist:"exprlist",
	Python3Kindtestlist:"testlist",
	Python3Kinddictorsetmaker:"dictorsetmaker",
	Python3Kindclassdef:"classdef",
	Python3Kindarglist:"arglist",
	Python3Kindpython_argument:"python_argument",
	Python3Kindcomp_iter:"comp_iter",
	Python3Kindcomp_for:"comp_for",
	Python3Kindcomp_if:"comp_if",
	Python3Kindencoding_decl:"encoding_decl",
	Python3Kindyield_expr:"yield_expr",
	Python3Kindyield_arg:"yield_arg",
}

