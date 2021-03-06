// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fast_

type SolidityKind = int32
const (
	SolidityKindsourceunit SolidityKind = 0
	SolidityKindpragmadirective SolidityKind = 1
	SolidityKindpragmaname SolidityKind = 2
	SolidityKindpragmavalue SolidityKind = 3
	SolidityKindversion SolidityKind = 4
	SolidityKindversionoperator SolidityKind = 5
	SolidityKindversionconstraint SolidityKind = 6
	SolidityKindimportdeclaration SolidityKind = 7
	SolidityKindimportdirective SolidityKind = 8
	SolidityKindcontractdefinition SolidityKind = 9
	SolidityKindinheritancespecifier SolidityKind = 10
	SolidityKindcontractpart SolidityKind = 11
	SolidityKindstatevariabledeclaration SolidityKind = 12
	SolidityKindusingfordeclaration SolidityKind = 13
	SolidityKindstructdefinition SolidityKind = 14
	SolidityKindmodifierdefinition SolidityKind = 15
	SolidityKindmodifierinvocation SolidityKind = 16
	SolidityKindfunctiondefinition SolidityKind = 17
	SolidityKindreturnparameters SolidityKind = 18
	SolidityKindmodifierlist SolidityKind = 19
	SolidityKindeventdefinition SolidityKind = 20
	SolidityKindenumvalue SolidityKind = 21
	SolidityKindenumdefinition SolidityKind = 22
	SolidityKindindexedparameterlist SolidityKind = 23
	SolidityKindindexedparameter SolidityKind = 24
	SolidityKindparameterlist SolidityKind = 25
	SolidityKindparameter SolidityKind = 26
	SolidityKindtypenamelist SolidityKind = 27
	SolidityKindunnamedparameter SolidityKind = 28
	SolidityKindvariabledeclaration SolidityKind = 29
	SolidityKindtypename SolidityKind = 30
	SolidityKinduserdefinedtypename SolidityKind = 31
	SolidityKindmapping SolidityKind = 32
	SolidityKindfunctiontypename SolidityKind = 33
	SolidityKindstoragelocation SolidityKind = 34
	SolidityKindstatemutability SolidityKind = 35
	SolidityKindblock SolidityKind = 36
	SolidityKindsolidity_statement SolidityKind = 37
	SolidityKindexpressionstatement SolidityKind = 38
	SolidityKindsolidity_ifstatement SolidityKind = 39
	SolidityKindwhilestatement SolidityKind = 40
	SolidityKindsimplestatement SolidityKind = 41
	SolidityKindforstatement SolidityKind = 42
	SolidityKindinlineassemblystatement SolidityKind = 43
	SolidityKinddowhilestatement SolidityKind = 44
	SolidityKindsolidity_continuestatement SolidityKind = 45
	SolidityKindbreakstatement SolidityKind = 46
	SolidityKindsolidity_returnstatement SolidityKind = 47
	SolidityKindthrowstatement SolidityKind = 48
	SolidityKindvariabledeclarationstatement SolidityKind = 49
	SolidityKindidentifierlist SolidityKind = 50
	SolidityKindelementarytypename SolidityKind = 51
	SolidityKindexpression SolidityKind = 52
	SolidityKindprimaryexpression SolidityKind = 53
	SolidityKindexpressionlist SolidityKind = 54
	SolidityKindnamevaluelist SolidityKind = 55
	SolidityKindnamevalue SolidityKind = 56
	SolidityKindfunctioncallarguments SolidityKind = 57
	SolidityKindassemblyblock SolidityKind = 58
	SolidityKindassemblyitem SolidityKind = 59
	SolidityKindassemblyexpression SolidityKind = 60
	SolidityKindassemblycall SolidityKind = 61
	SolidityKindassemblylocaldefinition SolidityKind = 62
	SolidityKindassemblyassignment SolidityKind = 63
	SolidityKindassemblyidentifierorlist SolidityKind = 64
	SolidityKindassemblyidentifierlist SolidityKind = 65
	SolidityKindassemblyrightassignment SolidityKind = 66
	SolidityKindlabeldefinition SolidityKind = 67
	SolidityKindassemblyswitch SolidityKind = 68
	SolidityKindassemblycase SolidityKind = 69
	SolidityKindassemblyfunctiondefinition SolidityKind = 70
	SolidityKindassemblyfor SolidityKind = 71
	SolidityKindassemblyliteral SolidityKind = 72
	SolidityKindsubassembly SolidityKind = 73
	SolidityKinddatasize SolidityKind = 74
	SolidityKindlinkersymbol SolidityKind = 75
	SolidityKindtupleexpression SolidityKind = 76
	SolidityKindelementarytypenameexpression SolidityKind = 77
	SolidityKindnumberliteral SolidityKind = 78
	SolidityKindsolidity_identifier SolidityKind = 79
)

var EnumNamesSolidityKind = map[SolidityKind]string{
	SolidityKindsourceunit:"sourceunit",
	SolidityKindpragmadirective:"pragmadirective",
	SolidityKindpragmaname:"pragmaname",
	SolidityKindpragmavalue:"pragmavalue",
	SolidityKindversion:"version",
	SolidityKindversionoperator:"versionoperator",
	SolidityKindversionconstraint:"versionconstraint",
	SolidityKindimportdeclaration:"importdeclaration",
	SolidityKindimportdirective:"importdirective",
	SolidityKindcontractdefinition:"contractdefinition",
	SolidityKindinheritancespecifier:"inheritancespecifier",
	SolidityKindcontractpart:"contractpart",
	SolidityKindstatevariabledeclaration:"statevariabledeclaration",
	SolidityKindusingfordeclaration:"usingfordeclaration",
	SolidityKindstructdefinition:"structdefinition",
	SolidityKindmodifierdefinition:"modifierdefinition",
	SolidityKindmodifierinvocation:"modifierinvocation",
	SolidityKindfunctiondefinition:"functiondefinition",
	SolidityKindreturnparameters:"returnparameters",
	SolidityKindmodifierlist:"modifierlist",
	SolidityKindeventdefinition:"eventdefinition",
	SolidityKindenumvalue:"enumvalue",
	SolidityKindenumdefinition:"enumdefinition",
	SolidityKindindexedparameterlist:"indexedparameterlist",
	SolidityKindindexedparameter:"indexedparameter",
	SolidityKindparameterlist:"parameterlist",
	SolidityKindparameter:"parameter",
	SolidityKindtypenamelist:"typenamelist",
	SolidityKindunnamedparameter:"unnamedparameter",
	SolidityKindvariabledeclaration:"variabledeclaration",
	SolidityKindtypename:"typename",
	SolidityKinduserdefinedtypename:"userdefinedtypename",
	SolidityKindmapping:"mapping",
	SolidityKindfunctiontypename:"functiontypename",
	SolidityKindstoragelocation:"storagelocation",
	SolidityKindstatemutability:"statemutability",
	SolidityKindblock:"block",
	SolidityKindsolidity_statement:"solidity_statement",
	SolidityKindexpressionstatement:"expressionstatement",
	SolidityKindsolidity_ifstatement:"solidity_ifstatement",
	SolidityKindwhilestatement:"whilestatement",
	SolidityKindsimplestatement:"simplestatement",
	SolidityKindforstatement:"forstatement",
	SolidityKindinlineassemblystatement:"inlineassemblystatement",
	SolidityKinddowhilestatement:"dowhilestatement",
	SolidityKindsolidity_continuestatement:"solidity_continuestatement",
	SolidityKindbreakstatement:"breakstatement",
	SolidityKindsolidity_returnstatement:"solidity_returnstatement",
	SolidityKindthrowstatement:"throwstatement",
	SolidityKindvariabledeclarationstatement:"variabledeclarationstatement",
	SolidityKindidentifierlist:"identifierlist",
	SolidityKindelementarytypename:"elementarytypename",
	SolidityKindexpression:"expression",
	SolidityKindprimaryexpression:"primaryexpression",
	SolidityKindexpressionlist:"expressionlist",
	SolidityKindnamevaluelist:"namevaluelist",
	SolidityKindnamevalue:"namevalue",
	SolidityKindfunctioncallarguments:"functioncallarguments",
	SolidityKindassemblyblock:"assemblyblock",
	SolidityKindassemblyitem:"assemblyitem",
	SolidityKindassemblyexpression:"assemblyexpression",
	SolidityKindassemblycall:"assemblycall",
	SolidityKindassemblylocaldefinition:"assemblylocaldefinition",
	SolidityKindassemblyassignment:"assemblyassignment",
	SolidityKindassemblyidentifierorlist:"assemblyidentifierorlist",
	SolidityKindassemblyidentifierlist:"assemblyidentifierlist",
	SolidityKindassemblyrightassignment:"assemblyrightassignment",
	SolidityKindlabeldefinition:"labeldefinition",
	SolidityKindassemblyswitch:"assemblyswitch",
	SolidityKindassemblycase:"assemblycase",
	SolidityKindassemblyfunctiondefinition:"assemblyfunctiondefinition",
	SolidityKindassemblyfor:"assemblyfor",
	SolidityKindassemblyliteral:"assemblyliteral",
	SolidityKindsubassembly:"subassembly",
	SolidityKinddatasize:"datasize",
	SolidityKindlinkersymbol:"linkersymbol",
	SolidityKindtupleexpression:"tupleexpression",
	SolidityKindelementarytypenameexpression:"elementarytypenameexpression",
	SolidityKindnumberliteral:"numberliteral",
	SolidityKindsolidity_identifier:"solidity_identifier",
}

