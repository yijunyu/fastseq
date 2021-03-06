// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fast_

type SrcmlKind = int32
const (
	SrcmlKindUNIT_KIND SrcmlKind = 0
	SrcmlKindDECL SrcmlKind = 1
	SrcmlKindDECL_STMT SrcmlKind = 2
	SrcmlKindINIT SrcmlKind = 3
	SrcmlKindEXPR SrcmlKind = 4
	SrcmlKindEXPR_STMT SrcmlKind = 5
	SrcmlKindCOMMENT SrcmlKind = 6
	SrcmlKindCALL SrcmlKind = 7
	SrcmlKindCONTROL SrcmlKind = 8
	SrcmlKindINCR SrcmlKind = 9
	SrcmlKindNONE SrcmlKind = 10
	SrcmlKindVARIABLE SrcmlKind = 11
	SrcmlKindFUNCTION SrcmlKind = 12
	SrcmlKindFUNCTION_DECL SrcmlKind = 13
	SrcmlKindCONSTRUCTOR SrcmlKind = 14
	SrcmlKindCONSTRUCTOR_DECL SrcmlKind = 15
	SrcmlKindDESTRUCTOR SrcmlKind = 16
	SrcmlKindDESTRUCTOR_DECL SrcmlKind = 17
	SrcmlKindMACRO SrcmlKind = 18
	SrcmlKindSINGLE_MACRO SrcmlKind = 19
	SrcmlKindNULLOPERATOR SrcmlKind = 20
	SrcmlKindENUM_DEFN SrcmlKind = 21
	SrcmlKindENUM_DECL SrcmlKind = 22
	SrcmlKindGLOBAL_ATTRIBUTE SrcmlKind = 23
	SrcmlKindPROPERTY_ACCESSOR SrcmlKind = 24
	SrcmlKindPROPERTY_ACCESSOR_DECL SrcmlKind = 25
	SrcmlKindEXPRESSION SrcmlKind = 26
	SrcmlKindCLASS_DEFN SrcmlKind = 27
	SrcmlKindCLASS_DECL SrcmlKind = 28
	SrcmlKindUNION_DEFN SrcmlKind = 29
	SrcmlKindUNION_DECL SrcmlKind = 30
	SrcmlKindSTRUCT_DEFN SrcmlKind = 31
	SrcmlKindSTRUCT_DECL SrcmlKind = 32
	SrcmlKindINTERFACE_DEFN SrcmlKind = 33
	SrcmlKindINTERFACE_DECL SrcmlKind = 34
	SrcmlKindACCESS_REGION SrcmlKind = 35
	SrcmlKindUSING SrcmlKind = 36
	SrcmlKindOPERATOR_FUNCTION SrcmlKind = 37
	SrcmlKindOPERATOR_FUNCTION_DECL SrcmlKind = 38
	SrcmlKindEVENT SrcmlKind = 39
	SrcmlKindPROPERTY SrcmlKind = 40
	SrcmlKindANNOTATION_DEFN SrcmlKind = 41
	SrcmlKindGLOBAL_TEMPLATE SrcmlKind = 42
	SrcmlKindUNIT SrcmlKind = 43
	SrcmlKindTART_ELEMENT_TOKEN SrcmlKind = 44
	SrcmlKindNOP SrcmlKind = 45
	SrcmlKindSTRING SrcmlKind = 46
	SrcmlKindCHAR SrcmlKind = 47
	SrcmlKindLITERAL SrcmlKind = 48
	SrcmlKindBOOLEAN SrcmlKind = 49
	SrcmlKindNULL2 SrcmlKind = 50
	SrcmlKindCOMPLEX SrcmlKind = 51
	SrcmlKindOPERATOR SrcmlKind = 52
	SrcmlKindMODIFIER SrcmlKind = 53
	SrcmlKindNAME SrcmlKind = 54
	SrcmlKindONAME SrcmlKind = 55
	SrcmlKindCNAME SrcmlKind = 56
	SrcmlKindTYPE SrcmlKind = 57
	SrcmlKindTYPEPREV SrcmlKind = 58
	SrcmlKindCONDITION SrcmlKind = 59
	SrcmlKindBLOCK SrcmlKind = 60
	SrcmlKindPSEUDO_BLOCK SrcmlKind = 61
	SrcmlKindINDEX SrcmlKind = 62
	SrcmlKindENUM SrcmlKind = 63
	SrcmlKindENUM_DECLARATION SrcmlKind = 64
	SrcmlKindIF_STATEMENT SrcmlKind = 65
	SrcmlKindTERNARY SrcmlKind = 66
	SrcmlKindTHEN SrcmlKind = 67
	SrcmlKindELSE SrcmlKind = 68
	SrcmlKindELSEIF SrcmlKind = 69
	SrcmlKindWHILE_STATEMENT SrcmlKind = 70
	SrcmlKindDO_STATEMENT SrcmlKind = 71
	SrcmlKindFOR_STATEMENT SrcmlKind = 72
	SrcmlKindFOREACH_STATEMENT SrcmlKind = 73
	SrcmlKindFOR_CONTROL SrcmlKind = 74
	SrcmlKindFOR_INITIALIZATION SrcmlKind = 75
	SrcmlKindFOR_CONDITION SrcmlKind = 76
	SrcmlKindFOR_INCREMENT SrcmlKind = 77
	SrcmlKindFOR_LIKE_CONTROL SrcmlKind = 78
	SrcmlKindEXPRESSION_STATEMENT SrcmlKind = 79
	SrcmlKindFUNCTION_CALL SrcmlKind = 80
	SrcmlKindDECLARATION_STATEMENT SrcmlKind = 81
	SrcmlKindDECLARATION SrcmlKind = 82
	SrcmlKindDECLARATION_INITIALIZATION SrcmlKind = 83
	SrcmlKindDECLARATION_RANGE SrcmlKind = 84
	SrcmlKindRANGE SrcmlKind = 85
	SrcmlKindGOTO_STATEMENT SrcmlKind = 86
	SrcmlKindCONTINUE_STATEMENT SrcmlKind = 87
	SrcmlKindBREAK_STATEMENT SrcmlKind = 88
	SrcmlKindLABEL_STATEMENT SrcmlKind = 89
	SrcmlKindLABEL SrcmlKind = 90
	SrcmlKindSWITCH SrcmlKind = 91
	SrcmlKindCASE SrcmlKind = 92
	SrcmlKindDEFAULT SrcmlKind = 93
	SrcmlKindFUNCTION_DEFINITION SrcmlKind = 94
	SrcmlKindFUNCTION_DECLARATION SrcmlKind = 95
	SrcmlKindLAMBDA SrcmlKind = 96
	SrcmlKindFUNCTION_LAMBDA SrcmlKind = 97
	SrcmlKindFUNCTION_SPECIFIER SrcmlKind = 98
	SrcmlKindRETURN_STATEMENT SrcmlKind = 99
	SrcmlKindPARAMETER_LIST SrcmlKind = 100
	SrcmlKindPARAMETER SrcmlKind = 101
	SrcmlKindKRPARAMETER_LIST SrcmlKind = 102
	SrcmlKindKRPARAMETER SrcmlKind = 103
	SrcmlKindARGUMENT_LIST SrcmlKind = 104
	SrcmlKindARGUMENT SrcmlKind = 105
	SrcmlKindPSEUDO_PARAMETER_LIST SrcmlKind = 106
	SrcmlKindINDEXER_PARAMETER_LIST SrcmlKind = 107
	SrcmlKindCLASS SrcmlKind = 108
	SrcmlKindCLASS_DECLARATION SrcmlKind = 109
	SrcmlKindSTRUCT SrcmlKind = 110
	SrcmlKindSTRUCT_DECLARATION SrcmlKind = 111
	SrcmlKindUNION SrcmlKind = 112
	SrcmlKindUNION_DECLARATION SrcmlKind = 113
	SrcmlKindDERIVATION_LIST SrcmlKind = 114
	SrcmlKindPUBLIC_ACCESS SrcmlKind = 115
	SrcmlKindPUBLIC_ACCESS_DEFAULT SrcmlKind = 116
	SrcmlKindPRIVATE_ACCESS SrcmlKind = 117
	SrcmlKindPRIVATE_ACCESS_DEFAULT SrcmlKind = 118
	SrcmlKindPROTECTED_ACCESS SrcmlKind = 119
	SrcmlKindPROTECTED_ACCESS_DEFAULT SrcmlKind = 120
	SrcmlKindMEMBER_INIT_LIST SrcmlKind = 121
	SrcmlKindMEMBER_INITIALIZATION_LIST SrcmlKind = 122
	SrcmlKindMEMBER_INITIALIZATION SrcmlKind = 123
	SrcmlKindCONSTRUCTOR_DEFINITION SrcmlKind = 124
	SrcmlKindCONSTRUCTOR_DECLARATION SrcmlKind = 125
	SrcmlKindDESTRUCTOR_DEFINITION SrcmlKind = 126
	SrcmlKindDESTRUCTOR_DECLARATION SrcmlKind = 127
	SrcmlKindFRIEND SrcmlKind = 128
	SrcmlKindCLASS_SPECIFIER SrcmlKind = 129
	SrcmlKindTRY_BLOCK SrcmlKind = 130
	SrcmlKindCATCH_BLOCK SrcmlKind = 131
	SrcmlKindFINALLY_BLOCK SrcmlKind = 132
	SrcmlKindTHROW_STATEMENT SrcmlKind = 133
	SrcmlKindTHROW_SPECIFIER SrcmlKind = 134
	SrcmlKindTHROW_SPECIFIER_JAVA SrcmlKind = 135
	SrcmlKindTEMPLATE SrcmlKind = 136
	SrcmlKindGENERIC_ARGUMENT SrcmlKind = 137
	SrcmlKindGENERIC_ARGUMENT_LIST SrcmlKind = 138
	SrcmlKindTEMPLATE_PARAMETER SrcmlKind = 139
	SrcmlKindTEMPLATE_PARAMETER_LIST SrcmlKind = 140
	SrcmlKindGENERIC_PARAMETER SrcmlKind = 141
	SrcmlKindGENERIC_PARAMETER_LIST SrcmlKind = 142
	SrcmlKindTYPEDEF SrcmlKind = 143
	SrcmlKindASM SrcmlKind = 144
	SrcmlKindMACRO_CALL SrcmlKind = 145
	SrcmlKindSIZEOF_CALL SrcmlKind = 146
	SrcmlKindEXTERN SrcmlKind = 147
	SrcmlKindNAMESPACE SrcmlKind = 148
	SrcmlKindUSING_DIRECTIVE SrcmlKind = 149
	SrcmlKindDIRECTIVE SrcmlKind = 150
	SrcmlKindATOMIC SrcmlKind = 151
	SrcmlKindSTATIC_ASSERT_STATEMENT SrcmlKind = 152
	SrcmlKindGENERIC_SELECTION SrcmlKind = 153
	SrcmlKindGENERIC_SELECTOR SrcmlKind = 154
	SrcmlKindGENERIC_ASSOCIATION_LIST SrcmlKind = 155
	SrcmlKindGENERIC_ASSOCIATION SrcmlKind = 156
	SrcmlKindALIGNAS SrcmlKind = 157
	SrcmlKindDECLTYPE SrcmlKind = 158
	SrcmlKindCAPTURE SrcmlKind = 159
	SrcmlKindLAMBDA_CAPTURE SrcmlKind = 160
	SrcmlKindNOEXCEPT SrcmlKind = 161
	SrcmlKindTYPENAME SrcmlKind = 162
	SrcmlKindALIGNOF SrcmlKind = 163
	SrcmlKindTYPEID SrcmlKind = 164
	SrcmlKindSIZEOF_PACK SrcmlKind = 165
	SrcmlKindENUM_CLASS SrcmlKind = 166
	SrcmlKindENUM_CLASS_DECLARATION SrcmlKind = 167
	SrcmlKindREF_QUALIFIER SrcmlKind = 168
	SrcmlKindSIGNAL_ACCESS SrcmlKind = 169
	SrcmlKindFOREVER_STATEMENT SrcmlKind = 170
	SrcmlKindEMIT_STATEMENT SrcmlKind = 171
	SrcmlKindCPP_DIRECTIVE SrcmlKind = 172
	SrcmlKindCPP_FILENAME SrcmlKind = 173
	SrcmlKindFILE SrcmlKind = 174
	SrcmlKindNUMBER SrcmlKind = 175
	SrcmlKindCPP_NUMBER SrcmlKind = 176
	SrcmlKindCPP_LITERAL SrcmlKind = 177
	SrcmlKindCPP_MACRO_DEFN SrcmlKind = 178
	SrcmlKindCPP_MACRO_VALUE SrcmlKind = 179
	SrcmlKindERROR SrcmlKind = 180
	SrcmlKindCPP_ERROR SrcmlKind = 181
	SrcmlKindCPP_WARNING SrcmlKind = 182
	SrcmlKindCPP_PRAGMA SrcmlKind = 183
	SrcmlKindCPP_INCLUDE SrcmlKind = 184
	SrcmlKindCPP_DEFINE SrcmlKind = 185
	SrcmlKindCPP_UNDEF SrcmlKind = 186
	SrcmlKindCPP_LINE SrcmlKind = 187
	SrcmlKindCPP_IF SrcmlKind = 188
	SrcmlKindCPP_IFDEF SrcmlKind = 189
	SrcmlKindCPP_IFNDEF SrcmlKind = 190
	SrcmlKindCPP_THEN SrcmlKind = 191
	SrcmlKindCPP_ELSE SrcmlKind = 192
	SrcmlKindCPP_ELIF SrcmlKind = 193
	SrcmlKindCPP_EMPTY SrcmlKind = 194
	SrcmlKindCPP_REGION SrcmlKind = 195
	SrcmlKindCPP_ENDREGION SrcmlKind = 196
	SrcmlKindUSING_STMT SrcmlKind = 197
	SrcmlKindESCAPE SrcmlKind = 198
	SrcmlKindVALUE SrcmlKind = 199
	SrcmlKindCPP_IMPORT SrcmlKind = 200
	SrcmlKindCPP_ENDIF SrcmlKind = 201
	SrcmlKindMARKER SrcmlKind = 202
	SrcmlKindERROR_PARSE SrcmlKind = 203
	SrcmlKindERROR_MODE SrcmlKind = 204
	SrcmlKindIMPLEMENTS SrcmlKind = 205
	SrcmlKindEXTENDS SrcmlKind = 206
	SrcmlKindIMPORT SrcmlKind = 207
	SrcmlKindPACKAGE SrcmlKind = 208
	SrcmlKindASSERT_STATEMENT SrcmlKind = 209
	SrcmlKindINTERFACE SrcmlKind = 210
	SrcmlKindINTERFACE_DECLARATION SrcmlKind = 211
	SrcmlKindSYNCHRONIZED_STATEMENT SrcmlKind = 212
	SrcmlKindANNOTATION SrcmlKind = 213
	SrcmlKindSTATIC_BLOCK SrcmlKind = 214
	SrcmlKindCHECKED_STATEMENT SrcmlKind = 215
	SrcmlKindUNCHECKED_STATEMENT SrcmlKind = 216
	SrcmlKindATTRIBUTE SrcmlKind = 217
	SrcmlKindTARGET SrcmlKind = 218
	SrcmlKindUNSAFE_STATEMENT SrcmlKind = 219
	SrcmlKindLOCK_STATEMENT SrcmlKind = 220
	SrcmlKindFIXED_STATEMENT SrcmlKind = 221
	SrcmlKindTYPEOF SrcmlKind = 222
	SrcmlKindUSING_STATEMENT SrcmlKind = 223
	SrcmlKindFUNCTION_DELEGATE SrcmlKind = 224
	SrcmlKindCONSTRAINT SrcmlKind = 225
	SrcmlKindLINQ SrcmlKind = 226
	SrcmlKindFROM SrcmlKind = 227
	SrcmlKindWHERE SrcmlKind = 228
	SrcmlKindSELECT SrcmlKind = 229
	SrcmlKindLET SrcmlKind = 230
	SrcmlKindORDERBY SrcmlKind = 231
	SrcmlKindJOIN SrcmlKind = 232
	SrcmlKindGROUP SrcmlKind = 233
	SrcmlKindIN SrcmlKind = 234
	SrcmlKindON SrcmlKind = 235
	SrcmlKindEQUALS SrcmlKind = 236
	SrcmlKindBY SrcmlKind = 237
	SrcmlKindINTO SrcmlKind = 238
	SrcmlKindEMPTY SrcmlKind = 239
	SrcmlKindEMPTY_STMT SrcmlKind = 240
	SrcmlKindRECEIVER SrcmlKind = 241
	SrcmlKindMESSAGE SrcmlKind = 242
	SrcmlKindSELECTOR SrcmlKind = 243
	SrcmlKindPROTOCOL_LIST SrcmlKind = 244
	SrcmlKindCATEGORY SrcmlKind = 245
	SrcmlKindPROTOCOL SrcmlKind = 246
	SrcmlKindREQUIRED_DEFAULT SrcmlKind = 247
	SrcmlKindREQUIRED SrcmlKind = 248
	SrcmlKindOPTIONAL SrcmlKind = 249
	SrcmlKindATTRIBUTE_LIST SrcmlKind = 250
	SrcmlKindSYNTHESIZE SrcmlKind = 251
	SrcmlKindDYNAMIC SrcmlKind = 252
	SrcmlKindENCODE SrcmlKind = 253
	SrcmlKindAUTORELEASEPOOL SrcmlKind = 254
	SrcmlKindCOMPATIBILITY_ALIAS SrcmlKind = 255
	SrcmlKindNIL SrcmlKind = 256
	SrcmlKindCLASS_INTERFACE SrcmlKind = 257
	SrcmlKindCLASS_IMPLEMENTATION SrcmlKind = 258
	SrcmlKindPROTOCOL_DECLARATION SrcmlKind = 259
	SrcmlKindCAST SrcmlKind = 260
	SrcmlKindCONST_CAST SrcmlKind = 261
	SrcmlKindDYNAMIC_CAST SrcmlKind = 262
	SrcmlKindREINTERPRET_CAST SrcmlKind = 263
	SrcmlKindSTATIC_CAST SrcmlKind = 264
	SrcmlKindPOSITION SrcmlKind = 265
	SrcmlKindCUDA_ARGUMENT_LIST SrcmlKind = 266
	SrcmlKindOMP_DIRECTIVE SrcmlKind = 267
	SrcmlKindOMP_NAME SrcmlKind = 268
	SrcmlKindOMP_CLAUSE SrcmlKind = 269
	SrcmlKindOMP_ARGUMENT_LIST SrcmlKind = 270
	SrcmlKindOMP_ARGUMENT SrcmlKind = 271
	SrcmlKindOMP_EXPRESSION SrcmlKind = 272
	SrcmlKindEND_ELEMENT_TOKEN SrcmlKind = 273
	SrcmlKindMAIN SrcmlKind = 274
	SrcmlKindBREAK SrcmlKind = 275
	SrcmlKindCONTINUE SrcmlKind = 276
	SrcmlKindWHILE SrcmlKind = 277
	SrcmlKindDO SrcmlKind = 278
	SrcmlKindFOR SrcmlKind = 279
	SrcmlKindIF SrcmlKind = 280
	SrcmlKindGOTO SrcmlKind = 281
	SrcmlKindVISUAL_CXX_ASM SrcmlKind = 282
	SrcmlKindSIZEOF SrcmlKind = 283
	SrcmlKindAUTO SrcmlKind = 284
	SrcmlKindREGISTER SrcmlKind = 285
	SrcmlKindRESTRICT SrcmlKind = 286
	SrcmlKindIMAGINARY SrcmlKind = 287
	SrcmlKindNORETURN SrcmlKind = 288
	SrcmlKindSTATIC_ASSERT SrcmlKind = 289
	SrcmlKindCRESTRICT SrcmlKind = 290
	SrcmlKindCXX_TRY SrcmlKind = 291
	SrcmlKindCXX_CATCH SrcmlKind = 292
	SrcmlKindCXX_CLASS SrcmlKind = 293
	SrcmlKindCONSTEXPR SrcmlKind = 294
	SrcmlKindTHREAD_LOCAL SrcmlKind = 295
	SrcmlKindNULLPTR SrcmlKind = 296
	SrcmlKindVOID SrcmlKind = 297
	SrcmlKindRETURN SrcmlKind = 298
	SrcmlKindINCLUDE SrcmlKind = 299
	SrcmlKindDEFINE SrcmlKind = 300
	SrcmlKindELIF SrcmlKind = 301
	SrcmlKindENDIF SrcmlKind = 302
	SrcmlKindERRORPREC SrcmlKind = 303
	SrcmlKindWARNING SrcmlKind = 304
	SrcmlKindIFDEF SrcmlKind = 305
	SrcmlKindIFNDEF SrcmlKind = 306
	SrcmlKindLINE SrcmlKind = 307
	SrcmlKindPRAGMA SrcmlKind = 308
	SrcmlKindUNDEF SrcmlKind = 309
	SrcmlKindINLINE SrcmlKind = 310
	SrcmlKindMACRO_TYPE_NAME SrcmlKind = 311
	SrcmlKindMACRO_CASE SrcmlKind = 312
	SrcmlKindMACRO_LABEL SrcmlKind = 313
	SrcmlKindSPECIFIER SrcmlKind = 314
	SrcmlKindTRY SrcmlKind = 315
	SrcmlKindCATCH SrcmlKind = 316
	SrcmlKindTHROW SrcmlKind = 317
	SrcmlKindTHROWS SrcmlKind = 318
	SrcmlKindPUBLIC SrcmlKind = 319
	SrcmlKindPRIVATE SrcmlKind = 320
	SrcmlKindPROTECTED SrcmlKind = 321
	SrcmlKindVIRTUAL SrcmlKind = 322
	SrcmlKindEXPLICIT SrcmlKind = 323
	SrcmlKindFOREVER SrcmlKind = 324
	SrcmlKindSIGNAL SrcmlKind = 325
	SrcmlKindEMIT SrcmlKind = 326
	SrcmlKindNEW SrcmlKind = 327
	SrcmlKindDELETE SrcmlKind = 328
	SrcmlKindSTATIC SrcmlKind = 329
	SrcmlKindCONST SrcmlKind = 330
	SrcmlKindMUTABLE SrcmlKind = 331
	SrcmlKindVOLATILE SrcmlKind = 332
	SrcmlKindTRANSIENT SrcmlKind = 333
	SrcmlKindFINALLY SrcmlKind = 334
	SrcmlKindFINAL SrcmlKind = 335
	SrcmlKindABSTRACT SrcmlKind = 336
	SrcmlKindSUPER SrcmlKind = 337
	SrcmlKindSYNCHRONIZED SrcmlKind = 338
	SrcmlKindNATIVE SrcmlKind = 339
	SrcmlKindSTRICTFP SrcmlKind = 340
	SrcmlKindNULLLITERAL SrcmlKind = 341
	SrcmlKindASSERT SrcmlKind = 342
	SrcmlKindFOREACH SrcmlKind = 343
	SrcmlKindREF SrcmlKind = 344
	SrcmlKindOUT SrcmlKind = 345
	SrcmlKindLOCK SrcmlKind = 346
	SrcmlKindIS SrcmlKind = 347
	SrcmlKindINTERNAL SrcmlKind = 348
	SrcmlKindSEALED SrcmlKind = 349
	SrcmlKindOVERRIDE SrcmlKind = 350
	SrcmlKindIMPLICIT SrcmlKind = 351
	SrcmlKindSTACKALLOC SrcmlKind = 352
	SrcmlKindAS SrcmlKind = 353
	SrcmlKindDELEGATE SrcmlKind = 354
	SrcmlKindFIXED SrcmlKind = 355
	SrcmlKindCHECKED SrcmlKind = 356
	SrcmlKindUNCHECKED SrcmlKind = 357
	SrcmlKindREGION SrcmlKind = 358
	SrcmlKindENDREGION SrcmlKind = 359
	SrcmlKindUNSAFE SrcmlKind = 360
	SrcmlKindREADONLY SrcmlKind = 361
	SrcmlKindGET SrcmlKind = 362
	SrcmlKindSET SrcmlKind = 363
	SrcmlKindADD SrcmlKind = 364
	SrcmlKindREMOVE SrcmlKind = 365
	SrcmlKindYIELD SrcmlKind = 366
	SrcmlKindPARTIAL SrcmlKind = 367
	SrcmlKindAWAIT SrcmlKind = 368
	SrcmlKindASYNC SrcmlKind = 369
	SrcmlKindTHIS SrcmlKind = 370
	SrcmlKindPARAMS SrcmlKind = 371
	SrcmlKindALIAS SrcmlKind = 372
	SrcmlKindASCENDING SrcmlKind = 373
	SrcmlKindDESCENDING SrcmlKind = 374
	SrcmlKindATINTERFACE SrcmlKind = 375
	SrcmlKindATIMPLEMENTATION SrcmlKind = 376
	SrcmlKindATEND SrcmlKind = 377
	SrcmlKindATPROTOCOL SrcmlKind = 378
	SrcmlKindATREQUIRED SrcmlKind = 379
	SrcmlKindATOPTIONAL SrcmlKind = 380
	SrcmlKindATCLASS SrcmlKind = 381
	SrcmlKindWEAK SrcmlKind = 382
	SrcmlKindSTRONG SrcmlKind = 383
	SrcmlKindOMP_OMP SrcmlKind = 384
	SrcmlKindSPECIAL_CHARS SrcmlKind = 385
	SrcmlKindSLICE_DEFINE SrcmlKind = 386
	SrcmlKindSLICE_USE SrcmlKind = 387
)

var EnumNamesSrcmlKind = map[SrcmlKind]string{
	SrcmlKindUNIT_KIND:"UNIT_KIND",
	SrcmlKindDECL:"DECL",
	SrcmlKindDECL_STMT:"DECL_STMT",
	SrcmlKindINIT:"INIT",
	SrcmlKindEXPR:"EXPR",
	SrcmlKindEXPR_STMT:"EXPR_STMT",
	SrcmlKindCOMMENT:"COMMENT",
	SrcmlKindCALL:"CALL",
	SrcmlKindCONTROL:"CONTROL",
	SrcmlKindINCR:"INCR",
	SrcmlKindNONE:"NONE",
	SrcmlKindVARIABLE:"VARIABLE",
	SrcmlKindFUNCTION:"FUNCTION",
	SrcmlKindFUNCTION_DECL:"FUNCTION_DECL",
	SrcmlKindCONSTRUCTOR:"CONSTRUCTOR",
	SrcmlKindCONSTRUCTOR_DECL:"CONSTRUCTOR_DECL",
	SrcmlKindDESTRUCTOR:"DESTRUCTOR",
	SrcmlKindDESTRUCTOR_DECL:"DESTRUCTOR_DECL",
	SrcmlKindMACRO:"MACRO",
	SrcmlKindSINGLE_MACRO:"SINGLE_MACRO",
	SrcmlKindNULLOPERATOR:"NULLOPERATOR",
	SrcmlKindENUM_DEFN:"ENUM_DEFN",
	SrcmlKindENUM_DECL:"ENUM_DECL",
	SrcmlKindGLOBAL_ATTRIBUTE:"GLOBAL_ATTRIBUTE",
	SrcmlKindPROPERTY_ACCESSOR:"PROPERTY_ACCESSOR",
	SrcmlKindPROPERTY_ACCESSOR_DECL:"PROPERTY_ACCESSOR_DECL",
	SrcmlKindEXPRESSION:"EXPRESSION",
	SrcmlKindCLASS_DEFN:"CLASS_DEFN",
	SrcmlKindCLASS_DECL:"CLASS_DECL",
	SrcmlKindUNION_DEFN:"UNION_DEFN",
	SrcmlKindUNION_DECL:"UNION_DECL",
	SrcmlKindSTRUCT_DEFN:"STRUCT_DEFN",
	SrcmlKindSTRUCT_DECL:"STRUCT_DECL",
	SrcmlKindINTERFACE_DEFN:"INTERFACE_DEFN",
	SrcmlKindINTERFACE_DECL:"INTERFACE_DECL",
	SrcmlKindACCESS_REGION:"ACCESS_REGION",
	SrcmlKindUSING:"USING",
	SrcmlKindOPERATOR_FUNCTION:"OPERATOR_FUNCTION",
	SrcmlKindOPERATOR_FUNCTION_DECL:"OPERATOR_FUNCTION_DECL",
	SrcmlKindEVENT:"EVENT",
	SrcmlKindPROPERTY:"PROPERTY",
	SrcmlKindANNOTATION_DEFN:"ANNOTATION_DEFN",
	SrcmlKindGLOBAL_TEMPLATE:"GLOBAL_TEMPLATE",
	SrcmlKindUNIT:"UNIT",
	SrcmlKindTART_ELEMENT_TOKEN:"TART_ELEMENT_TOKEN",
	SrcmlKindNOP:"NOP",
	SrcmlKindSTRING:"STRING",
	SrcmlKindCHAR:"CHAR",
	SrcmlKindLITERAL:"LITERAL",
	SrcmlKindBOOLEAN:"BOOLEAN",
	SrcmlKindNULL2:"NULL2",
	SrcmlKindCOMPLEX:"COMPLEX",
	SrcmlKindOPERATOR:"OPERATOR",
	SrcmlKindMODIFIER:"MODIFIER",
	SrcmlKindNAME:"NAME",
	SrcmlKindONAME:"ONAME",
	SrcmlKindCNAME:"CNAME",
	SrcmlKindTYPE:"TYPE",
	SrcmlKindTYPEPREV:"TYPEPREV",
	SrcmlKindCONDITION:"CONDITION",
	SrcmlKindBLOCK:"BLOCK",
	SrcmlKindPSEUDO_BLOCK:"PSEUDO_BLOCK",
	SrcmlKindINDEX:"INDEX",
	SrcmlKindENUM:"ENUM",
	SrcmlKindENUM_DECLARATION:"ENUM_DECLARATION",
	SrcmlKindIF_STATEMENT:"IF_STATEMENT",
	SrcmlKindTERNARY:"TERNARY",
	SrcmlKindTHEN:"THEN",
	SrcmlKindELSE:"ELSE",
	SrcmlKindELSEIF:"ELSEIF",
	SrcmlKindWHILE_STATEMENT:"WHILE_STATEMENT",
	SrcmlKindDO_STATEMENT:"DO_STATEMENT",
	SrcmlKindFOR_STATEMENT:"FOR_STATEMENT",
	SrcmlKindFOREACH_STATEMENT:"FOREACH_STATEMENT",
	SrcmlKindFOR_CONTROL:"FOR_CONTROL",
	SrcmlKindFOR_INITIALIZATION:"FOR_INITIALIZATION",
	SrcmlKindFOR_CONDITION:"FOR_CONDITION",
	SrcmlKindFOR_INCREMENT:"FOR_INCREMENT",
	SrcmlKindFOR_LIKE_CONTROL:"FOR_LIKE_CONTROL",
	SrcmlKindEXPRESSION_STATEMENT:"EXPRESSION_STATEMENT",
	SrcmlKindFUNCTION_CALL:"FUNCTION_CALL",
	SrcmlKindDECLARATION_STATEMENT:"DECLARATION_STATEMENT",
	SrcmlKindDECLARATION:"DECLARATION",
	SrcmlKindDECLARATION_INITIALIZATION:"DECLARATION_INITIALIZATION",
	SrcmlKindDECLARATION_RANGE:"DECLARATION_RANGE",
	SrcmlKindRANGE:"RANGE",
	SrcmlKindGOTO_STATEMENT:"GOTO_STATEMENT",
	SrcmlKindCONTINUE_STATEMENT:"CONTINUE_STATEMENT",
	SrcmlKindBREAK_STATEMENT:"BREAK_STATEMENT",
	SrcmlKindLABEL_STATEMENT:"LABEL_STATEMENT",
	SrcmlKindLABEL:"LABEL",
	SrcmlKindSWITCH:"SWITCH",
	SrcmlKindCASE:"CASE",
	SrcmlKindDEFAULT:"DEFAULT",
	SrcmlKindFUNCTION_DEFINITION:"FUNCTION_DEFINITION",
	SrcmlKindFUNCTION_DECLARATION:"FUNCTION_DECLARATION",
	SrcmlKindLAMBDA:"LAMBDA",
	SrcmlKindFUNCTION_LAMBDA:"FUNCTION_LAMBDA",
	SrcmlKindFUNCTION_SPECIFIER:"FUNCTION_SPECIFIER",
	SrcmlKindRETURN_STATEMENT:"RETURN_STATEMENT",
	SrcmlKindPARAMETER_LIST:"PARAMETER_LIST",
	SrcmlKindPARAMETER:"PARAMETER",
	SrcmlKindKRPARAMETER_LIST:"KRPARAMETER_LIST",
	SrcmlKindKRPARAMETER:"KRPARAMETER",
	SrcmlKindARGUMENT_LIST:"ARGUMENT_LIST",
	SrcmlKindARGUMENT:"ARGUMENT",
	SrcmlKindPSEUDO_PARAMETER_LIST:"PSEUDO_PARAMETER_LIST",
	SrcmlKindINDEXER_PARAMETER_LIST:"INDEXER_PARAMETER_LIST",
	SrcmlKindCLASS:"CLASS",
	SrcmlKindCLASS_DECLARATION:"CLASS_DECLARATION",
	SrcmlKindSTRUCT:"STRUCT",
	SrcmlKindSTRUCT_DECLARATION:"STRUCT_DECLARATION",
	SrcmlKindUNION:"UNION",
	SrcmlKindUNION_DECLARATION:"UNION_DECLARATION",
	SrcmlKindDERIVATION_LIST:"DERIVATION_LIST",
	SrcmlKindPUBLIC_ACCESS:"PUBLIC_ACCESS",
	SrcmlKindPUBLIC_ACCESS_DEFAULT:"PUBLIC_ACCESS_DEFAULT",
	SrcmlKindPRIVATE_ACCESS:"PRIVATE_ACCESS",
	SrcmlKindPRIVATE_ACCESS_DEFAULT:"PRIVATE_ACCESS_DEFAULT",
	SrcmlKindPROTECTED_ACCESS:"PROTECTED_ACCESS",
	SrcmlKindPROTECTED_ACCESS_DEFAULT:"PROTECTED_ACCESS_DEFAULT",
	SrcmlKindMEMBER_INIT_LIST:"MEMBER_INIT_LIST",
	SrcmlKindMEMBER_INITIALIZATION_LIST:"MEMBER_INITIALIZATION_LIST",
	SrcmlKindMEMBER_INITIALIZATION:"MEMBER_INITIALIZATION",
	SrcmlKindCONSTRUCTOR_DEFINITION:"CONSTRUCTOR_DEFINITION",
	SrcmlKindCONSTRUCTOR_DECLARATION:"CONSTRUCTOR_DECLARATION",
	SrcmlKindDESTRUCTOR_DEFINITION:"DESTRUCTOR_DEFINITION",
	SrcmlKindDESTRUCTOR_DECLARATION:"DESTRUCTOR_DECLARATION",
	SrcmlKindFRIEND:"FRIEND",
	SrcmlKindCLASS_SPECIFIER:"CLASS_SPECIFIER",
	SrcmlKindTRY_BLOCK:"TRY_BLOCK",
	SrcmlKindCATCH_BLOCK:"CATCH_BLOCK",
	SrcmlKindFINALLY_BLOCK:"FINALLY_BLOCK",
	SrcmlKindTHROW_STATEMENT:"THROW_STATEMENT",
	SrcmlKindTHROW_SPECIFIER:"THROW_SPECIFIER",
	SrcmlKindTHROW_SPECIFIER_JAVA:"THROW_SPECIFIER_JAVA",
	SrcmlKindTEMPLATE:"TEMPLATE",
	SrcmlKindGENERIC_ARGUMENT:"GENERIC_ARGUMENT",
	SrcmlKindGENERIC_ARGUMENT_LIST:"GENERIC_ARGUMENT_LIST",
	SrcmlKindTEMPLATE_PARAMETER:"TEMPLATE_PARAMETER",
	SrcmlKindTEMPLATE_PARAMETER_LIST:"TEMPLATE_PARAMETER_LIST",
	SrcmlKindGENERIC_PARAMETER:"GENERIC_PARAMETER",
	SrcmlKindGENERIC_PARAMETER_LIST:"GENERIC_PARAMETER_LIST",
	SrcmlKindTYPEDEF:"TYPEDEF",
	SrcmlKindASM:"ASM",
	SrcmlKindMACRO_CALL:"MACRO_CALL",
	SrcmlKindSIZEOF_CALL:"SIZEOF_CALL",
	SrcmlKindEXTERN:"EXTERN",
	SrcmlKindNAMESPACE:"NAMESPACE",
	SrcmlKindUSING_DIRECTIVE:"USING_DIRECTIVE",
	SrcmlKindDIRECTIVE:"DIRECTIVE",
	SrcmlKindATOMIC:"ATOMIC",
	SrcmlKindSTATIC_ASSERT_STATEMENT:"STATIC_ASSERT_STATEMENT",
	SrcmlKindGENERIC_SELECTION:"GENERIC_SELECTION",
	SrcmlKindGENERIC_SELECTOR:"GENERIC_SELECTOR",
	SrcmlKindGENERIC_ASSOCIATION_LIST:"GENERIC_ASSOCIATION_LIST",
	SrcmlKindGENERIC_ASSOCIATION:"GENERIC_ASSOCIATION",
	SrcmlKindALIGNAS:"ALIGNAS",
	SrcmlKindDECLTYPE:"DECLTYPE",
	SrcmlKindCAPTURE:"CAPTURE",
	SrcmlKindLAMBDA_CAPTURE:"LAMBDA_CAPTURE",
	SrcmlKindNOEXCEPT:"NOEXCEPT",
	SrcmlKindTYPENAME:"TYPENAME",
	SrcmlKindALIGNOF:"ALIGNOF",
	SrcmlKindTYPEID:"TYPEID",
	SrcmlKindSIZEOF_PACK:"SIZEOF_PACK",
	SrcmlKindENUM_CLASS:"ENUM_CLASS",
	SrcmlKindENUM_CLASS_DECLARATION:"ENUM_CLASS_DECLARATION",
	SrcmlKindREF_QUALIFIER:"REF_QUALIFIER",
	SrcmlKindSIGNAL_ACCESS:"SIGNAL_ACCESS",
	SrcmlKindFOREVER_STATEMENT:"FOREVER_STATEMENT",
	SrcmlKindEMIT_STATEMENT:"EMIT_STATEMENT",
	SrcmlKindCPP_DIRECTIVE:"CPP_DIRECTIVE",
	SrcmlKindCPP_FILENAME:"CPP_FILENAME",
	SrcmlKindFILE:"FILE",
	SrcmlKindNUMBER:"NUMBER",
	SrcmlKindCPP_NUMBER:"CPP_NUMBER",
	SrcmlKindCPP_LITERAL:"CPP_LITERAL",
	SrcmlKindCPP_MACRO_DEFN:"CPP_MACRO_DEFN",
	SrcmlKindCPP_MACRO_VALUE:"CPP_MACRO_VALUE",
	SrcmlKindERROR:"ERROR",
	SrcmlKindCPP_ERROR:"CPP_ERROR",
	SrcmlKindCPP_WARNING:"CPP_WARNING",
	SrcmlKindCPP_PRAGMA:"CPP_PRAGMA",
	SrcmlKindCPP_INCLUDE:"CPP_INCLUDE",
	SrcmlKindCPP_DEFINE:"CPP_DEFINE",
	SrcmlKindCPP_UNDEF:"CPP_UNDEF",
	SrcmlKindCPP_LINE:"CPP_LINE",
	SrcmlKindCPP_IF:"CPP_IF",
	SrcmlKindCPP_IFDEF:"CPP_IFDEF",
	SrcmlKindCPP_IFNDEF:"CPP_IFNDEF",
	SrcmlKindCPP_THEN:"CPP_THEN",
	SrcmlKindCPP_ELSE:"CPP_ELSE",
	SrcmlKindCPP_ELIF:"CPP_ELIF",
	SrcmlKindCPP_EMPTY:"CPP_EMPTY",
	SrcmlKindCPP_REGION:"CPP_REGION",
	SrcmlKindCPP_ENDREGION:"CPP_ENDREGION",
	SrcmlKindUSING_STMT:"USING_STMT",
	SrcmlKindESCAPE:"ESCAPE",
	SrcmlKindVALUE:"VALUE",
	SrcmlKindCPP_IMPORT:"CPP_IMPORT",
	SrcmlKindCPP_ENDIF:"CPP_ENDIF",
	SrcmlKindMARKER:"MARKER",
	SrcmlKindERROR_PARSE:"ERROR_PARSE",
	SrcmlKindERROR_MODE:"ERROR_MODE",
	SrcmlKindIMPLEMENTS:"IMPLEMENTS",
	SrcmlKindEXTENDS:"EXTENDS",
	SrcmlKindIMPORT:"IMPORT",
	SrcmlKindPACKAGE:"PACKAGE",
	SrcmlKindASSERT_STATEMENT:"ASSERT_STATEMENT",
	SrcmlKindINTERFACE:"INTERFACE",
	SrcmlKindINTERFACE_DECLARATION:"INTERFACE_DECLARATION",
	SrcmlKindSYNCHRONIZED_STATEMENT:"SYNCHRONIZED_STATEMENT",
	SrcmlKindANNOTATION:"ANNOTATION",
	SrcmlKindSTATIC_BLOCK:"STATIC_BLOCK",
	SrcmlKindCHECKED_STATEMENT:"CHECKED_STATEMENT",
	SrcmlKindUNCHECKED_STATEMENT:"UNCHECKED_STATEMENT",
	SrcmlKindATTRIBUTE:"ATTRIBUTE",
	SrcmlKindTARGET:"TARGET",
	SrcmlKindUNSAFE_STATEMENT:"UNSAFE_STATEMENT",
	SrcmlKindLOCK_STATEMENT:"LOCK_STATEMENT",
	SrcmlKindFIXED_STATEMENT:"FIXED_STATEMENT",
	SrcmlKindTYPEOF:"TYPEOF",
	SrcmlKindUSING_STATEMENT:"USING_STATEMENT",
	SrcmlKindFUNCTION_DELEGATE:"FUNCTION_DELEGATE",
	SrcmlKindCONSTRAINT:"CONSTRAINT",
	SrcmlKindLINQ:"LINQ",
	SrcmlKindFROM:"FROM",
	SrcmlKindWHERE:"WHERE",
	SrcmlKindSELECT:"SELECT",
	SrcmlKindLET:"LET",
	SrcmlKindORDERBY:"ORDERBY",
	SrcmlKindJOIN:"JOIN",
	SrcmlKindGROUP:"GROUP",
	SrcmlKindIN:"IN",
	SrcmlKindON:"ON",
	SrcmlKindEQUALS:"EQUALS",
	SrcmlKindBY:"BY",
	SrcmlKindINTO:"INTO",
	SrcmlKindEMPTY:"EMPTY",
	SrcmlKindEMPTY_STMT:"EMPTY_STMT",
	SrcmlKindRECEIVER:"RECEIVER",
	SrcmlKindMESSAGE:"MESSAGE",
	SrcmlKindSELECTOR:"SELECTOR",
	SrcmlKindPROTOCOL_LIST:"PROTOCOL_LIST",
	SrcmlKindCATEGORY:"CATEGORY",
	SrcmlKindPROTOCOL:"PROTOCOL",
	SrcmlKindREQUIRED_DEFAULT:"REQUIRED_DEFAULT",
	SrcmlKindREQUIRED:"REQUIRED",
	SrcmlKindOPTIONAL:"OPTIONAL",
	SrcmlKindATTRIBUTE_LIST:"ATTRIBUTE_LIST",
	SrcmlKindSYNTHESIZE:"SYNTHESIZE",
	SrcmlKindDYNAMIC:"DYNAMIC",
	SrcmlKindENCODE:"ENCODE",
	SrcmlKindAUTORELEASEPOOL:"AUTORELEASEPOOL",
	SrcmlKindCOMPATIBILITY_ALIAS:"COMPATIBILITY_ALIAS",
	SrcmlKindNIL:"NIL",
	SrcmlKindCLASS_INTERFACE:"CLASS_INTERFACE",
	SrcmlKindCLASS_IMPLEMENTATION:"CLASS_IMPLEMENTATION",
	SrcmlKindPROTOCOL_DECLARATION:"PROTOCOL_DECLARATION",
	SrcmlKindCAST:"CAST",
	SrcmlKindCONST_CAST:"CONST_CAST",
	SrcmlKindDYNAMIC_CAST:"DYNAMIC_CAST",
	SrcmlKindREINTERPRET_CAST:"REINTERPRET_CAST",
	SrcmlKindSTATIC_CAST:"STATIC_CAST",
	SrcmlKindPOSITION:"POSITION",
	SrcmlKindCUDA_ARGUMENT_LIST:"CUDA_ARGUMENT_LIST",
	SrcmlKindOMP_DIRECTIVE:"OMP_DIRECTIVE",
	SrcmlKindOMP_NAME:"OMP_NAME",
	SrcmlKindOMP_CLAUSE:"OMP_CLAUSE",
	SrcmlKindOMP_ARGUMENT_LIST:"OMP_ARGUMENT_LIST",
	SrcmlKindOMP_ARGUMENT:"OMP_ARGUMENT",
	SrcmlKindOMP_EXPRESSION:"OMP_EXPRESSION",
	SrcmlKindEND_ELEMENT_TOKEN:"END_ELEMENT_TOKEN",
	SrcmlKindMAIN:"MAIN",
	SrcmlKindBREAK:"BREAK",
	SrcmlKindCONTINUE:"CONTINUE",
	SrcmlKindWHILE:"WHILE",
	SrcmlKindDO:"DO",
	SrcmlKindFOR:"FOR",
	SrcmlKindIF:"IF",
	SrcmlKindGOTO:"GOTO",
	SrcmlKindVISUAL_CXX_ASM:"VISUAL_CXX_ASM",
	SrcmlKindSIZEOF:"SIZEOF",
	SrcmlKindAUTO:"AUTO",
	SrcmlKindREGISTER:"REGISTER",
	SrcmlKindRESTRICT:"RESTRICT",
	SrcmlKindIMAGINARY:"IMAGINARY",
	SrcmlKindNORETURN:"NORETURN",
	SrcmlKindSTATIC_ASSERT:"STATIC_ASSERT",
	SrcmlKindCRESTRICT:"CRESTRICT",
	SrcmlKindCXX_TRY:"CXX_TRY",
	SrcmlKindCXX_CATCH:"CXX_CATCH",
	SrcmlKindCXX_CLASS:"CXX_CLASS",
	SrcmlKindCONSTEXPR:"CONSTEXPR",
	SrcmlKindTHREAD_LOCAL:"THREAD_LOCAL",
	SrcmlKindNULLPTR:"NULLPTR",
	SrcmlKindVOID:"VOID",
	SrcmlKindRETURN:"RETURN",
	SrcmlKindINCLUDE:"INCLUDE",
	SrcmlKindDEFINE:"DEFINE",
	SrcmlKindELIF:"ELIF",
	SrcmlKindENDIF:"ENDIF",
	SrcmlKindERRORPREC:"ERRORPREC",
	SrcmlKindWARNING:"WARNING",
	SrcmlKindIFDEF:"IFDEF",
	SrcmlKindIFNDEF:"IFNDEF",
	SrcmlKindLINE:"LINE",
	SrcmlKindPRAGMA:"PRAGMA",
	SrcmlKindUNDEF:"UNDEF",
	SrcmlKindINLINE:"INLINE",
	SrcmlKindMACRO_TYPE_NAME:"MACRO_TYPE_NAME",
	SrcmlKindMACRO_CASE:"MACRO_CASE",
	SrcmlKindMACRO_LABEL:"MACRO_LABEL",
	SrcmlKindSPECIFIER:"SPECIFIER",
	SrcmlKindTRY:"TRY",
	SrcmlKindCATCH:"CATCH",
	SrcmlKindTHROW:"THROW",
	SrcmlKindTHROWS:"THROWS",
	SrcmlKindPUBLIC:"PUBLIC",
	SrcmlKindPRIVATE:"PRIVATE",
	SrcmlKindPROTECTED:"PROTECTED",
	SrcmlKindVIRTUAL:"VIRTUAL",
	SrcmlKindEXPLICIT:"EXPLICIT",
	SrcmlKindFOREVER:"FOREVER",
	SrcmlKindSIGNAL:"SIGNAL",
	SrcmlKindEMIT:"EMIT",
	SrcmlKindNEW:"NEW",
	SrcmlKindDELETE:"DELETE",
	SrcmlKindSTATIC:"STATIC",
	SrcmlKindCONST:"CONST",
	SrcmlKindMUTABLE:"MUTABLE",
	SrcmlKindVOLATILE:"VOLATILE",
	SrcmlKindTRANSIENT:"TRANSIENT",
	SrcmlKindFINALLY:"FINALLY",
	SrcmlKindFINAL:"FINAL",
	SrcmlKindABSTRACT:"ABSTRACT",
	SrcmlKindSUPER:"SUPER",
	SrcmlKindSYNCHRONIZED:"SYNCHRONIZED",
	SrcmlKindNATIVE:"NATIVE",
	SrcmlKindSTRICTFP:"STRICTFP",
	SrcmlKindNULLLITERAL:"NULLLITERAL",
	SrcmlKindASSERT:"ASSERT",
	SrcmlKindFOREACH:"FOREACH",
	SrcmlKindREF:"REF",
	SrcmlKindOUT:"OUT",
	SrcmlKindLOCK:"LOCK",
	SrcmlKindIS:"IS",
	SrcmlKindINTERNAL:"INTERNAL",
	SrcmlKindSEALED:"SEALED",
	SrcmlKindOVERRIDE:"OVERRIDE",
	SrcmlKindIMPLICIT:"IMPLICIT",
	SrcmlKindSTACKALLOC:"STACKALLOC",
	SrcmlKindAS:"AS",
	SrcmlKindDELEGATE:"DELEGATE",
	SrcmlKindFIXED:"FIXED",
	SrcmlKindCHECKED:"CHECKED",
	SrcmlKindUNCHECKED:"UNCHECKED",
	SrcmlKindREGION:"REGION",
	SrcmlKindENDREGION:"ENDREGION",
	SrcmlKindUNSAFE:"UNSAFE",
	SrcmlKindREADONLY:"READONLY",
	SrcmlKindGET:"GET",
	SrcmlKindSET:"SET",
	SrcmlKindADD:"ADD",
	SrcmlKindREMOVE:"REMOVE",
	SrcmlKindYIELD:"YIELD",
	SrcmlKindPARTIAL:"PARTIAL",
	SrcmlKindAWAIT:"AWAIT",
	SrcmlKindASYNC:"ASYNC",
	SrcmlKindTHIS:"THIS",
	SrcmlKindPARAMS:"PARAMS",
	SrcmlKindALIAS:"ALIAS",
	SrcmlKindASCENDING:"ASCENDING",
	SrcmlKindDESCENDING:"DESCENDING",
	SrcmlKindATINTERFACE:"ATINTERFACE",
	SrcmlKindATIMPLEMENTATION:"ATIMPLEMENTATION",
	SrcmlKindATEND:"ATEND",
	SrcmlKindATPROTOCOL:"ATPROTOCOL",
	SrcmlKindATREQUIRED:"ATREQUIRED",
	SrcmlKindATOPTIONAL:"ATOPTIONAL",
	SrcmlKindATCLASS:"ATCLASS",
	SrcmlKindWEAK:"WEAK",
	SrcmlKindSTRONG:"STRONG",
	SrcmlKindOMP_OMP:"OMP_OMP",
	SrcmlKindSPECIAL_CHARS:"SPECIAL_CHARS",
	SrcmlKindSLICE_DEFINE:"SLICE_DEFINE",
	SrcmlKindSLICE_USE:"SLICE_USE",
}

