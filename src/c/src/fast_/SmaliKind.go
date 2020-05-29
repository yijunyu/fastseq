// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fast_

type SmaliKind = int32
const (
	SmaliKindsmali_file SmaliKind = 0
	SmaliKindclass_spec SmaliKind = 1
	SmaliKindsuper_spec SmaliKind = 2
	SmaliKindimplements_spec SmaliKind = 3
	SmaliKindsource_spec SmaliKind = 4
	SmaliKindaccess_list SmaliKind = 5
	SmaliKindfield SmaliKind = 6
	SmaliKindmethod SmaliKind = 7
	SmaliKindstatements_and_directives SmaliKind = 8
	SmaliKindordered_method_item SmaliKind = 9
	SmaliKindregisters_directive SmaliKind = 10
	SmaliKindparam_list_or_id SmaliKind = 11
	SmaliKindsimple_name SmaliKind = 12
	SmaliKindmember_name SmaliKind = 13
	SmaliKindmethod_prototype SmaliKind = 14
	SmaliKindparam_list_or_id_primitive_type SmaliKind = 15
	SmaliKindparam_list SmaliKind = 16
	SmaliKindarray_descriptor SmaliKind = 17
	SmaliKindtype_descriptor SmaliKind = 18
	SmaliKindnonvoid_type_descriptor SmaliKind = 19
	SmaliKindreference_type_descriptor SmaliKind = 20
	SmaliKindinteger_literal SmaliKind = 21
	SmaliKindfloat_literal SmaliKind = 22
	SmaliKinddouble_literal SmaliKind = 23
	SmaliKindsmali_literal SmaliKind = 24
	SmaliKindparsed_integer_literal SmaliKind = 25
	SmaliKindintegral_literal SmaliKind = 26
	SmaliKindfixed_32bit_literal SmaliKind = 27
	SmaliKindfixed_literal SmaliKind = 28
	SmaliKindarray_literal SmaliKind = 29
	SmaliKindannotation_element SmaliKind = 30
	SmaliKindannotation SmaliKind = 31
	SmaliKindsubannotation SmaliKind = 32
	SmaliKindenum_literal SmaliKind = 33
	SmaliKindtype_field_method_literal SmaliKind = 34
	SmaliKindmethod_reference SmaliKind = 35
	SmaliKindfield_reference SmaliKind = 36
	SmaliKindlabel SmaliKind = 37
	SmaliKindlabel_ref SmaliKind = 38
	SmaliKindregister_list SmaliKind = 39
	SmaliKindregister_range SmaliKind = 40
	SmaliKindverification_error_reference SmaliKind = 41
	SmaliKindcatch_directive SmaliKind = 42
	SmaliKindcatchall_directive SmaliKind = 43
	SmaliKindparameter_directive SmaliKind = 44
	SmaliKinddebug_directive SmaliKind = 45
	SmaliKindline_directive SmaliKind = 46
	SmaliKindlocal_directive SmaliKind = 47
	SmaliKindend_local_directive SmaliKind = 48
	SmaliKindrestart_local_directive SmaliKind = 49
	SmaliKindprologue_directive SmaliKind = 50
	SmaliKindepilogue_directive SmaliKind = 51
	SmaliKindsource_directive SmaliKind = 52
	SmaliKindinstruction_format12x SmaliKind = 53
	SmaliKindinstruction_format22s SmaliKind = 54
	SmaliKindinstruction_format31i SmaliKind = 55
	SmaliKindinstruction SmaliKind = 56
	SmaliKindinsn_format10t SmaliKind = 57
	SmaliKindinsn_format10x SmaliKind = 58
	SmaliKindinsn_format10x_odex SmaliKind = 59
	SmaliKindinsn_format11n SmaliKind = 60
	SmaliKindinsn_format11x SmaliKind = 61
	SmaliKindinsn_format12x SmaliKind = 62
	SmaliKindinsn_format20bc SmaliKind = 63
	SmaliKindinsn_format20t SmaliKind = 64
	SmaliKindinsn_format21c_field SmaliKind = 65
	SmaliKindinsn_format21c_field_odex SmaliKind = 66
	SmaliKindinsn_format21c_string SmaliKind = 67
	SmaliKindinsn_format21c_type SmaliKind = 68
	SmaliKindinsn_format21ih SmaliKind = 69
	SmaliKindinsn_format21lh SmaliKind = 70
	SmaliKindinsn_format21s SmaliKind = 71
	SmaliKindinsn_format21t SmaliKind = 72
	SmaliKindinsn_format22b SmaliKind = 73
	SmaliKindinsn_format22c_field SmaliKind = 74
	SmaliKindinsn_format22c_field_odex SmaliKind = 75
	SmaliKindinsn_format22c_type SmaliKind = 76
	SmaliKindinsn_format22cs_field SmaliKind = 77
	SmaliKindinsn_format22s SmaliKind = 78
	SmaliKindinsn_format22t SmaliKind = 79
	SmaliKindinsn_format22x SmaliKind = 80
	SmaliKindinsn_format23x SmaliKind = 81
	SmaliKindinsn_format30t SmaliKind = 82
	SmaliKindinsn_format31c SmaliKind = 83
	SmaliKindinsn_format31i SmaliKind = 84
	SmaliKindinsn_format31t SmaliKind = 85
	SmaliKindinsn_format32x SmaliKind = 86
	SmaliKindinsn_format35c_method SmaliKind = 87
	SmaliKindinsn_format35c_type SmaliKind = 88
	SmaliKindinsn_format35c_method_odex SmaliKind = 89
	SmaliKindinsn_format35mi_method SmaliKind = 90
	SmaliKindinsn_format35ms_method SmaliKind = 91
	SmaliKindinsn_format3rc_method SmaliKind = 92
	SmaliKindinsn_format3rc_method_odex SmaliKind = 93
	SmaliKindinsn_format3rc_type SmaliKind = 94
	SmaliKindinsn_format3rmi_method SmaliKind = 95
	SmaliKindinsn_format3rms_method SmaliKind = 96
	SmaliKindinsn_format45cc_method SmaliKind = 97
	SmaliKindinsn_format4rcc_method SmaliKind = 98
	SmaliKindinsn_format51l SmaliKind = 99
	SmaliKindinsn_array_data_directive SmaliKind = 100
	SmaliKindinsn_packed_switch_directive SmaliKind = 101
	SmaliKindinsn_sparse_switch_directive SmaliKind = 102
)

var EnumNamesSmaliKind = map[SmaliKind]string{
	SmaliKindsmali_file:"smali_file",
	SmaliKindclass_spec:"class_spec",
	SmaliKindsuper_spec:"super_spec",
	SmaliKindimplements_spec:"implements_spec",
	SmaliKindsource_spec:"source_spec",
	SmaliKindaccess_list:"access_list",
	SmaliKindfield:"field",
	SmaliKindmethod:"method",
	SmaliKindstatements_and_directives:"statements_and_directives",
	SmaliKindordered_method_item:"ordered_method_item",
	SmaliKindregisters_directive:"registers_directive",
	SmaliKindparam_list_or_id:"param_list_or_id",
	SmaliKindsimple_name:"simple_name",
	SmaliKindmember_name:"member_name",
	SmaliKindmethod_prototype:"method_prototype",
	SmaliKindparam_list_or_id_primitive_type:"param_list_or_id_primitive_type",
	SmaliKindparam_list:"param_list",
	SmaliKindarray_descriptor:"array_descriptor",
	SmaliKindtype_descriptor:"type_descriptor",
	SmaliKindnonvoid_type_descriptor:"nonvoid_type_descriptor",
	SmaliKindreference_type_descriptor:"reference_type_descriptor",
	SmaliKindinteger_literal:"integer_literal",
	SmaliKindfloat_literal:"float_literal",
	SmaliKinddouble_literal:"double_literal",
	SmaliKindsmali_literal:"smali_literal",
	SmaliKindparsed_integer_literal:"parsed_integer_literal",
	SmaliKindintegral_literal:"integral_literal",
	SmaliKindfixed_32bit_literal:"fixed_32bit_literal",
	SmaliKindfixed_literal:"fixed_literal",
	SmaliKindarray_literal:"array_literal",
	SmaliKindannotation_element:"annotation_element",
	SmaliKindannotation:"annotation",
	SmaliKindsubannotation:"subannotation",
	SmaliKindenum_literal:"enum_literal",
	SmaliKindtype_field_method_literal:"type_field_method_literal",
	SmaliKindmethod_reference:"method_reference",
	SmaliKindfield_reference:"field_reference",
	SmaliKindlabel:"label",
	SmaliKindlabel_ref:"label_ref",
	SmaliKindregister_list:"register_list",
	SmaliKindregister_range:"register_range",
	SmaliKindverification_error_reference:"verification_error_reference",
	SmaliKindcatch_directive:"catch_directive",
	SmaliKindcatchall_directive:"catchall_directive",
	SmaliKindparameter_directive:"parameter_directive",
	SmaliKinddebug_directive:"debug_directive",
	SmaliKindline_directive:"line_directive",
	SmaliKindlocal_directive:"local_directive",
	SmaliKindend_local_directive:"end_local_directive",
	SmaliKindrestart_local_directive:"restart_local_directive",
	SmaliKindprologue_directive:"prologue_directive",
	SmaliKindepilogue_directive:"epilogue_directive",
	SmaliKindsource_directive:"source_directive",
	SmaliKindinstruction_format12x:"instruction_format12x",
	SmaliKindinstruction_format22s:"instruction_format22s",
	SmaliKindinstruction_format31i:"instruction_format31i",
	SmaliKindinstruction:"instruction",
	SmaliKindinsn_format10t:"insn_format10t",
	SmaliKindinsn_format10x:"insn_format10x",
	SmaliKindinsn_format10x_odex:"insn_format10x_odex",
	SmaliKindinsn_format11n:"insn_format11n",
	SmaliKindinsn_format11x:"insn_format11x",
	SmaliKindinsn_format12x:"insn_format12x",
	SmaliKindinsn_format20bc:"insn_format20bc",
	SmaliKindinsn_format20t:"insn_format20t",
	SmaliKindinsn_format21c_field:"insn_format21c_field",
	SmaliKindinsn_format21c_field_odex:"insn_format21c_field_odex",
	SmaliKindinsn_format21c_string:"insn_format21c_string",
	SmaliKindinsn_format21c_type:"insn_format21c_type",
	SmaliKindinsn_format21ih:"insn_format21ih",
	SmaliKindinsn_format21lh:"insn_format21lh",
	SmaliKindinsn_format21s:"insn_format21s",
	SmaliKindinsn_format21t:"insn_format21t",
	SmaliKindinsn_format22b:"insn_format22b",
	SmaliKindinsn_format22c_field:"insn_format22c_field",
	SmaliKindinsn_format22c_field_odex:"insn_format22c_field_odex",
	SmaliKindinsn_format22c_type:"insn_format22c_type",
	SmaliKindinsn_format22cs_field:"insn_format22cs_field",
	SmaliKindinsn_format22s:"insn_format22s",
	SmaliKindinsn_format22t:"insn_format22t",
	SmaliKindinsn_format22x:"insn_format22x",
	SmaliKindinsn_format23x:"insn_format23x",
	SmaliKindinsn_format30t:"insn_format30t",
	SmaliKindinsn_format31c:"insn_format31c",
	SmaliKindinsn_format31i:"insn_format31i",
	SmaliKindinsn_format31t:"insn_format31t",
	SmaliKindinsn_format32x:"insn_format32x",
	SmaliKindinsn_format35c_method:"insn_format35c_method",
	SmaliKindinsn_format35c_type:"insn_format35c_type",
	SmaliKindinsn_format35c_method_odex:"insn_format35c_method_odex",
	SmaliKindinsn_format35mi_method:"insn_format35mi_method",
	SmaliKindinsn_format35ms_method:"insn_format35ms_method",
	SmaliKindinsn_format3rc_method:"insn_format3rc_method",
	SmaliKindinsn_format3rc_method_odex:"insn_format3rc_method_odex",
	SmaliKindinsn_format3rc_type:"insn_format3rc_type",
	SmaliKindinsn_format3rmi_method:"insn_format3rmi_method",
	SmaliKindinsn_format3rms_method:"insn_format3rms_method",
	SmaliKindinsn_format45cc_method:"insn_format45cc_method",
	SmaliKindinsn_format4rcc_method:"insn_format4rcc_method",
	SmaliKindinsn_format51l:"insn_format51l",
	SmaliKindinsn_array_data_directive:"insn_array_data_directive",
	SmaliKindinsn_packed_switch_directive:"insn_packed_switch_directive",
	SmaliKindinsn_sparse_switch_directive:"insn_sparse_switch_directive",
}
