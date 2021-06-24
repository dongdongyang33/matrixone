package client

//tuple (collation id, collation name)
type collationIDName struct {
	collationID int
	collationName string
}

//the map: charset --> default (collation id, collation name)
//the relation between the charset with the collation: one charset can correspond to multiple collations.
//so there is a default collation for the charset.
//Run the SQL below in Mysql 8.0.23 to get the map.
//the SQL:select concat('"',A.character_set_name,'":\t\t{',B.ID,',\t"',A.default_collate_name,'"},') from INFORMATION_SCHEMA.CHARACTER_SETS AS A join INFORMATION_SCHEMA.COLLATIONS AS B on (A.default_collate_name = B.collation_name) order by A.character_set_name ;
var charset2Collation=map[string]collationIDName{
	"armscii8":		{32,	"armscii8_general_ci"},
	"ascii":		{11,	"ascii_general_ci"},
	"big5":		{1,	"big5_chinese_ci"},
	"binary":		{63,	"binary"},
	"cp1250":		{26,	"cp1250_general_ci"},
	"cp1251":		{51,	"cp1251_general_ci"},
	"cp1256":		{57,	"cp1256_general_ci"},
	"cp1257":		{59,	"cp1257_general_ci"},
	"cp850":		{4,	"cp850_general_ci"},
	"cp852":		{40,	"cp852_general_ci"},
	"cp866":		{36,	"cp866_general_ci"},
	"cp932":		{95,	"cp932_japanese_ci"},
	"dec8":		{3,	"dec8_swedish_ci"},
	"eucjpms":		{97,	"eucjpms_japanese_ci"},
	"euckr":		{19,	"euckr_korean_ci"},
	"gb18030":		{248,	"gb18030_chinese_ci"},
	"gb2312":		{24,	"gb2312_chinese_ci"},
	"gbk":		{28,	"gbk_chinese_ci"},
	"geostd8":		{92,	"geostd8_general_ci"},
	"greek":		{25,	"greek_general_ci"},
	"hebrew":		{16,	"hebrew_general_ci"},
	"hp8":		{6,	"hp8_english_ci"},
	"keybcs2":		{37,	"keybcs2_general_ci"},
	"koi8r":		{7,	"koi8r_general_ci"},
	"koi8u":		{22,	"koi8u_general_ci"},
	"latin1":		{8,	"latin1_swedish_ci"},
	"latin2":		{9,	"latin2_general_ci"},
	"latin5":		{30,	"latin5_turkish_ci"},
	"latin7":		{41,	"latin7_general_ci"},
	"macce":		{38,	"macce_general_ci"},
	"macroman":		{39,	"macroman_general_ci"},
	"sjis":		{13,	"sjis_japanese_ci"},
	"swe7":		{10,	"swe7_swedish_ci"},
	"tis620":		{18,	"tis620_thai_ci"},
	"ucs2":		{35,	"ucs2_general_ci"},
	"ujis":		{12,	"ujis_japanese_ci"},
	"utf16":		{54,	"utf16_general_ci"},
	"utf16le":		{56,	"utf16le_general_ci"},
	"utf32":		{60,	"utf32_general_ci"},
	"utf8":		{33,	"utf8_general_ci"},
	"utf8mb4":		{255,	"utf8mb4_0900_ai_ci"},
}

//tuple (collation name, charset)
type charsetCollationName struct {
	charset string
	collationName string
}

//the map: collation id --> (charset, collation name)
//Run the SQL below in Mysql 8.0.23 to get the map.
//the SQL: select concat(ID,':\t\t{"',CHARACTER_SET_NAME,'",\t"',collation_name,'"},') from INFORMATION_SCHEMA.COLLATIONS order by id;
var collationID2CharsetAndName = map[int]charsetCollationName{
	1:		{"big5",	"big5_chinese_ci"},
	2:		{"latin2",	"latin2_czech_cs"},
	3:		{"dec8",	"dec8_swedish_ci"},
	4:		{"cp850",	"cp850_general_ci"},
	5:		{"latin1",	"latin1_german1_ci"},
	6:		{"hp8",	"hp8_english_ci"},
	7:		{"koi8r",	"koi8r_general_ci"},
	8:		{"latin1",	"latin1_swedish_ci"},
	9:		{"latin2",	"latin2_general_ci"},
	10:		{"swe7",	"swe7_swedish_ci"},
	11:		{"ascii",	"ascii_general_ci"},
	12:		{"ujis",	"ujis_japanese_ci"},
	13:		{"sjis",	"sjis_japanese_ci"},
	14:		{"cp1251",	"cp1251_bulgarian_ci"},
	15:		{"latin1",	"latin1_danish_ci"},
	16:		{"hebrew",	"hebrew_general_ci"},
	18:		{"tis620",	"tis620_thai_ci"},
	19:		{"euckr",	"euckr_korean_ci"},
	20:		{"latin7",	"latin7_estonian_cs"},
	21:		{"latin2",	"latin2_hungarian_ci"},
	22:		{"koi8u",	"koi8u_general_ci"},
	23:		{"cp1251",	"cp1251_ukrainian_ci"},
	24:		{"gb2312",	"gb2312_chinese_ci"},
	25:		{"greek",	"greek_general_ci"},
	26:		{"cp1250",	"cp1250_general_ci"},
	27:		{"latin2",	"latin2_croatian_ci"},
	28:		{"gbk",	"gbk_chinese_ci"},
	29:		{"cp1257",	"cp1257_lithuanian_ci"},
	30:		{"latin5",	"latin5_turkish_ci"},
	31:		{"latin1",	"latin1_german2_ci"},
	32:		{"armscii8",	"armscii8_general_ci"},
	33:		{"utf8",	"utf8_general_ci"},
	34:		{"cp1250",	"cp1250_czech_cs"},
	35:		{"ucs2",	"ucs2_general_ci"},
	36:		{"cp866",	"cp866_general_ci"},
	37:		{"keybcs2",	"keybcs2_general_ci"},
	38:		{"macce",	"macce_general_ci"},
	39:		{"macroman",	"macroman_general_ci"},
	40:		{"cp852",	"cp852_general_ci"},
	41:		{"latin7",	"latin7_general_ci"},
	42:		{"latin7",	"latin7_general_cs"},
	43:		{"macce",	"macce_bin"},
	44:		{"cp1250",	"cp1250_croatian_ci"},
	45:		{"utf8mb4",	"utf8mb4_general_ci"},
	46:		{"utf8mb4",	"utf8mb4_bin"},
	47:		{"latin1",	"latin1_bin"},
	48:		{"latin1",	"latin1_general_ci"},
	49:		{"latin1",	"latin1_general_cs"},
	50:		{"cp1251",	"cp1251_bin"},
	51:		{"cp1251",	"cp1251_general_ci"},
	52:		{"cp1251",	"cp1251_general_cs"},
	53:		{"macroman",	"macroman_bin"},
	54:		{"utf16",	"utf16_general_ci"},
	55:		{"utf16",	"utf16_bin"},
	56:		{"utf16le",	"utf16le_general_ci"},
	57:		{"cp1256",	"cp1256_general_ci"},
	58:		{"cp1257",	"cp1257_bin"},
	59:		{"cp1257",	"cp1257_general_ci"},
	60:		{"utf32",	"utf32_general_ci"},
	61:		{"utf32",	"utf32_bin"},
	62:		{"utf16le",	"utf16le_bin"},
	63:		{"binary",	"binary"},
	64:		{"armscii8",	"armscii8_bin"},
	65:		{"ascii",	"ascii_bin"},
	66:		{"cp1250",	"cp1250_bin"},
	67:		{"cp1256",	"cp1256_bin"},
	68:		{"cp866",	"cp866_bin"},
	69:		{"dec8",	"dec8_bin"},
	70:		{"greek",	"greek_bin"},
	71:		{"hebrew",	"hebrew_bin"},
	72:		{"hp8",	"hp8_bin"},
	73:		{"keybcs2",	"keybcs2_bin"},
	74:		{"koi8r",	"koi8r_bin"},
	75:		{"koi8u",	"koi8u_bin"},
	76:		{"utf8",	"utf8_tolower_ci"},
	77:		{"latin2",	"latin2_bin"},
	78:		{"latin5",	"latin5_bin"},
	79:		{"latin7",	"latin7_bin"},
	80:		{"cp850",	"cp850_bin"},
	81:		{"cp852",	"cp852_bin"},
	82:		{"swe7",	"swe7_bin"},
	83:		{"utf8",	"utf8_bin"},
	84:		{"big5",	"big5_bin"},
	85:		{"euckr",	"euckr_bin"},
	86:		{"gb2312",	"gb2312_bin"},
	87:		{"gbk",	"gbk_bin"},
	88:		{"sjis",	"sjis_bin"},
	89:		{"tis620",	"tis620_bin"},
	90:		{"ucs2",	"ucs2_bin"},
	91:		{"ujis",	"ujis_bin"},
	92:		{"geostd8",	"geostd8_general_ci"},
	93:		{"geostd8",	"geostd8_bin"},
	94:		{"latin1",	"latin1_spanish_ci"},
	95:		{"cp932",	"cp932_japanese_ci"},
	96:		{"cp932",	"cp932_bin"},
	97:		{"eucjpms",	"eucjpms_japanese_ci"},
	98:		{"eucjpms",	"eucjpms_bin"},
	99:		{"cp1250",	"cp1250_polish_ci"},
	101:		{"utf16",	"utf16_unicode_ci"},
	102:		{"utf16",	"utf16_icelandic_ci"},
	103:		{"utf16",	"utf16_latvian_ci"},
	104:		{"utf16",	"utf16_romanian_ci"},
	105:		{"utf16",	"utf16_slovenian_ci"},
	106:		{"utf16",	"utf16_polish_ci"},
	107:		{"utf16",	"utf16_estonian_ci"},
	108:		{"utf16",	"utf16_spanish_ci"},
	109:		{"utf16",	"utf16_swedish_ci"},
	110:		{"utf16",	"utf16_turkish_ci"},
	111:		{"utf16",	"utf16_czech_ci"},
	112:		{"utf16",	"utf16_danish_ci"},
	113:		{"utf16",	"utf16_lithuanian_ci"},
	114:		{"utf16",	"utf16_slovak_ci"},
	115:		{"utf16",	"utf16_spanish2_ci"},
	116:		{"utf16",	"utf16_roman_ci"},
	117:		{"utf16",	"utf16_persian_ci"},
	118:		{"utf16",	"utf16_esperanto_ci"},
	119:		{"utf16",	"utf16_hungarian_ci"},
	120:		{"utf16",	"utf16_sinhala_ci"},
	121:		{"utf16",	"utf16_german2_ci"},
	122:		{"utf16",	"utf16_croatian_ci"},
	123:		{"utf16",	"utf16_unicode_520_ci"},
	124:		{"utf16",	"utf16_vietnamese_ci"},
	128:		{"ucs2",	"ucs2_unicode_ci"},
	129:		{"ucs2",	"ucs2_icelandic_ci"},
	130:		{"ucs2",	"ucs2_latvian_ci"},
	131:		{"ucs2",	"ucs2_romanian_ci"},
	132:		{"ucs2",	"ucs2_slovenian_ci"},
	133:		{"ucs2",	"ucs2_polish_ci"},
	134:		{"ucs2",	"ucs2_estonian_ci"},
	135:		{"ucs2",	"ucs2_spanish_ci"},
	136:		{"ucs2",	"ucs2_swedish_ci"},
	137:		{"ucs2",	"ucs2_turkish_ci"},
	138:		{"ucs2",	"ucs2_czech_ci"},
	139:		{"ucs2",	"ucs2_danish_ci"},
	140:		{"ucs2",	"ucs2_lithuanian_ci"},
	141:		{"ucs2",	"ucs2_slovak_ci"},
	142:		{"ucs2",	"ucs2_spanish2_ci"},
	143:		{"ucs2",	"ucs2_roman_ci"},
	144:		{"ucs2",	"ucs2_persian_ci"},
	145:		{"ucs2",	"ucs2_esperanto_ci"},
	146:		{"ucs2",	"ucs2_hungarian_ci"},
	147:		{"ucs2",	"ucs2_sinhala_ci"},
	148:		{"ucs2",	"ucs2_german2_ci"},
	149:		{"ucs2",	"ucs2_croatian_ci"},
	150:		{"ucs2",	"ucs2_unicode_520_ci"},
	151:		{"ucs2",	"ucs2_vietnamese_ci"},
	159:		{"ucs2",	"ucs2_general_mysql500_ci"},
	160:		{"utf32",	"utf32_unicode_ci"},
	161:		{"utf32",	"utf32_icelandic_ci"},
	162:		{"utf32",	"utf32_latvian_ci"},
	163:		{"utf32",	"utf32_romanian_ci"},
	164:		{"utf32",	"utf32_slovenian_ci"},
	165:		{"utf32",	"utf32_polish_ci"},
	166:		{"utf32",	"utf32_estonian_ci"},
	167:		{"utf32",	"utf32_spanish_ci"},
	168:		{"utf32",	"utf32_swedish_ci"},
	169:		{"utf32",	"utf32_turkish_ci"},
	170:		{"utf32",	"utf32_czech_ci"},
	171:		{"utf32",	"utf32_danish_ci"},
	172:		{"utf32",	"utf32_lithuanian_ci"},
	173:		{"utf32",	"utf32_slovak_ci"},
	174:		{"utf32",	"utf32_spanish2_ci"},
	175:		{"utf32",	"utf32_roman_ci"},
	176:		{"utf32",	"utf32_persian_ci"},
	177:		{"utf32",	"utf32_esperanto_ci"},
	178:		{"utf32",	"utf32_hungarian_ci"},
	179:		{"utf32",	"utf32_sinhala_ci"},
	180:		{"utf32",	"utf32_german2_ci"},
	181:		{"utf32",	"utf32_croatian_ci"},
	182:		{"utf32",	"utf32_unicode_520_ci"},
	183:		{"utf32",	"utf32_vietnamese_ci"},
	192:		{"utf8",	"utf8_unicode_ci"},
	193:		{"utf8",	"utf8_icelandic_ci"},
	194:		{"utf8",	"utf8_latvian_ci"},
	195:		{"utf8",	"utf8_romanian_ci"},
	196:		{"utf8",	"utf8_slovenian_ci"},
	197:		{"utf8",	"utf8_polish_ci"},
	198:		{"utf8",	"utf8_estonian_ci"},
	199:		{"utf8",	"utf8_spanish_ci"},
	200:		{"utf8",	"utf8_swedish_ci"},
	201:		{"utf8",	"utf8_turkish_ci"},
	202:		{"utf8",	"utf8_czech_ci"},
	203:		{"utf8",	"utf8_danish_ci"},
	204:		{"utf8",	"utf8_lithuanian_ci"},
	205:		{"utf8",	"utf8_slovak_ci"},
	206:		{"utf8",	"utf8_spanish2_ci"},
	207:		{"utf8",	"utf8_roman_ci"},
	208:		{"utf8",	"utf8_persian_ci"},
	209:		{"utf8",	"utf8_esperanto_ci"},
	210:		{"utf8",	"utf8_hungarian_ci"},
	211:		{"utf8",	"utf8_sinhala_ci"},
	212:		{"utf8",	"utf8_german2_ci"},
	213:		{"utf8",	"utf8_croatian_ci"},
	214:		{"utf8",	"utf8_unicode_520_ci"},
	215:		{"utf8",	"utf8_vietnamese_ci"},
	223:		{"utf8",	"utf8_general_mysql500_ci"},
	224:		{"utf8mb4",	"utf8mb4_unicode_ci"},
	225:		{"utf8mb4",	"utf8mb4_icelandic_ci"},
	226:		{"utf8mb4",	"utf8mb4_latvian_ci"},
	227:		{"utf8mb4",	"utf8mb4_romanian_ci"},
	228:		{"utf8mb4",	"utf8mb4_slovenian_ci"},
	229:		{"utf8mb4",	"utf8mb4_polish_ci"},
	230:		{"utf8mb4",	"utf8mb4_estonian_ci"},
	231:		{"utf8mb4",	"utf8mb4_spanish_ci"},
	232:		{"utf8mb4",	"utf8mb4_swedish_ci"},
	233:		{"utf8mb4",	"utf8mb4_turkish_ci"},
	234:		{"utf8mb4",	"utf8mb4_czech_ci"},
	235:		{"utf8mb4",	"utf8mb4_danish_ci"},
	236:		{"utf8mb4",	"utf8mb4_lithuanian_ci"},
	237:		{"utf8mb4",	"utf8mb4_slovak_ci"},
	238:		{"utf8mb4",	"utf8mb4_spanish2_ci"},
	239:		{"utf8mb4",	"utf8mb4_roman_ci"},
	240:		{"utf8mb4",	"utf8mb4_persian_ci"},
	241:		{"utf8mb4",	"utf8mb4_esperanto_ci"},
	242:		{"utf8mb4",	"utf8mb4_hungarian_ci"},
	243:		{"utf8mb4",	"utf8mb4_sinhala_ci"},
	244:		{"utf8mb4",	"utf8mb4_german2_ci"},
	245:		{"utf8mb4",	"utf8mb4_croatian_ci"},
	246:		{"utf8mb4",	"utf8mb4_unicode_520_ci"},
	247:		{"utf8mb4",	"utf8mb4_vietnamese_ci"},
	248:		{"gb18030",	"gb18030_chinese_ci"},
	249:		{"gb18030",	"gb18030_bin"},
	250:		{"gb18030",	"gb18030_unicode_520_ci"},
	255:		{"utf8mb4",	"utf8mb4_0900_ai_ci"},
	256:		{"utf8mb4",	"utf8mb4_de_pb_0900_ai_ci"},
	257:		{"utf8mb4",	"utf8mb4_is_0900_ai_ci"},
	258:		{"utf8mb4",	"utf8mb4_lv_0900_ai_ci"},
	259:		{"utf8mb4",	"utf8mb4_ro_0900_ai_ci"},
	260:		{"utf8mb4",	"utf8mb4_sl_0900_ai_ci"},
	261:		{"utf8mb4",	"utf8mb4_pl_0900_ai_ci"},
	262:		{"utf8mb4",	"utf8mb4_et_0900_ai_ci"},
	263:		{"utf8mb4",	"utf8mb4_es_0900_ai_ci"},
	264:		{"utf8mb4",	"utf8mb4_sv_0900_ai_ci"},
	265:		{"utf8mb4",	"utf8mb4_tr_0900_ai_ci"},
	266:		{"utf8mb4",	"utf8mb4_cs_0900_ai_ci"},
	267:		{"utf8mb4",	"utf8mb4_da_0900_ai_ci"},
	268:		{"utf8mb4",	"utf8mb4_lt_0900_ai_ci"},
	269:		{"utf8mb4",	"utf8mb4_sk_0900_ai_ci"},
	270:		{"utf8mb4",	"utf8mb4_es_trad_0900_ai_ci"},
	271:		{"utf8mb4",	"utf8mb4_la_0900_ai_ci"},
	273:		{"utf8mb4",	"utf8mb4_eo_0900_ai_ci"},
	274:		{"utf8mb4",	"utf8mb4_hu_0900_ai_ci"},
	275:		{"utf8mb4",	"utf8mb4_hr_0900_ai_ci"},
	277:		{"utf8mb4",	"utf8mb4_vi_0900_ai_ci"},
	278:		{"utf8mb4",	"utf8mb4_0900_as_cs"},
	279:		{"utf8mb4",	"utf8mb4_de_pb_0900_as_cs"},
	280:		{"utf8mb4",	"utf8mb4_is_0900_as_cs"},
	281:		{"utf8mb4",	"utf8mb4_lv_0900_as_cs"},
	282:		{"utf8mb4",	"utf8mb4_ro_0900_as_cs"},
	283:		{"utf8mb4",	"utf8mb4_sl_0900_as_cs"},
	284:		{"utf8mb4",	"utf8mb4_pl_0900_as_cs"},
	285:		{"utf8mb4",	"utf8mb4_et_0900_as_cs"},
	286:		{"utf8mb4",	"utf8mb4_es_0900_as_cs"},
	287:		{"utf8mb4",	"utf8mb4_sv_0900_as_cs"},
	288:		{"utf8mb4",	"utf8mb4_tr_0900_as_cs"},
	289:		{"utf8mb4",	"utf8mb4_cs_0900_as_cs"},
	290:		{"utf8mb4",	"utf8mb4_da_0900_as_cs"},
	291:		{"utf8mb4",	"utf8mb4_lt_0900_as_cs"},
	292:		{"utf8mb4",	"utf8mb4_sk_0900_as_cs"},
	293:		{"utf8mb4",	"utf8mb4_es_trad_0900_as_cs"},
	294:		{"utf8mb4",	"utf8mb4_la_0900_as_cs"},
	296:		{"utf8mb4",	"utf8mb4_eo_0900_as_cs"},
	297:		{"utf8mb4",	"utf8mb4_hu_0900_as_cs"},
	298:		{"utf8mb4",	"utf8mb4_hr_0900_as_cs"},
	300:		{"utf8mb4",	"utf8mb4_vi_0900_as_cs"},
	303:		{"utf8mb4",	"utf8mb4_ja_0900_as_cs"},
	304:		{"utf8mb4",	"utf8mb4_ja_0900_as_cs_ks"},
	305:		{"utf8mb4",	"utf8mb4_0900_as_ci"},
	306:		{"utf8mb4",	"utf8mb4_ru_0900_ai_ci"},
	307:		{"utf8mb4",	"utf8mb4_ru_0900_as_cs"},
	308:		{"utf8mb4",	"utf8mb4_zh_0900_as_cs"},
	309:		{"utf8mb4",	"utf8mb4_0900_bin"},
}

//tuple (charset,collation id)
type charsetCollationID struct {
	charset string
	collationID int
}
//the map: collation name --> (charset,collationId)
//Run the SQL below in Mysql 8.0.23 to get the map.
//SQL: select concat('"',collation_name,'":\t\t{"',CHARACTER_SET_NAME,'",',ID,'},') from INFORMATION_SCHEMA.COLLATIONS order by collation_name;
var collationName2CharsetAndID = map[string]charsetCollationID{
	"armscii8_bin":		{"armscii8",64},
	"armscii8_general_ci":		{"armscii8",32},
	"ascii_bin":		{"ascii",65},
	"ascii_general_ci":		{"ascii",11},
	"big5_bin":		{"big5",84},
	"big5_chinese_ci":		{"big5",1},
	"binary":		{"binary",63},
	"cp1250_bin":		{"cp1250",66},
	"cp1250_croatian_ci":		{"cp1250",44},
	"cp1250_czech_cs":		{"cp1250",34},
	"cp1250_general_ci":		{"cp1250",26},
	"cp1250_polish_ci":		{"cp1250",99},
	"cp1251_bin":		{"cp1251",50},
	"cp1251_bulgarian_ci":		{"cp1251",14},
	"cp1251_general_ci":		{"cp1251",51},
	"cp1251_general_cs":		{"cp1251",52},
	"cp1251_ukrainian_ci":		{"cp1251",23},
	"cp1256_bin":		{"cp1256",67},
	"cp1256_general_ci":		{"cp1256",57},
	"cp1257_bin":		{"cp1257",58},
	"cp1257_general_ci":		{"cp1257",59},
	"cp1257_lithuanian_ci":		{"cp1257",29},
	"cp850_bin":		{"cp850",80},
	"cp850_general_ci":		{"cp850",4},
	"cp852_bin":		{"cp852",81},
	"cp852_general_ci":		{"cp852",40},
	"cp866_bin":		{"cp866",68},
	"cp866_general_ci":		{"cp866",36},
	"cp932_bin":		{"cp932",96},
	"cp932_japanese_ci":		{"cp932",95},
	"dec8_bin":		{"dec8",69},
	"dec8_swedish_ci":		{"dec8",3},
	"eucjpms_bin":		{"eucjpms",98},
	"eucjpms_japanese_ci":		{"eucjpms",97},
	"euckr_bin":		{"euckr",85},
	"euckr_korean_ci":		{"euckr",19},
	"gb18030_bin":		{"gb18030",249},
	"gb18030_chinese_ci":		{"gb18030",248},
	"gb18030_unicode_520_ci":		{"gb18030",250},
	"gb2312_bin":		{"gb2312",86},
	"gb2312_chinese_ci":		{"gb2312",24},
	"gbk_bin":		{"gbk",87},
	"gbk_chinese_ci":		{"gbk",28},
	"geostd8_bin":		{"geostd8",93},
	"geostd8_general_ci":		{"geostd8",92},
	"greek_bin":		{"greek",70},
	"greek_general_ci":		{"greek",25},
	"hebrew_bin":		{"hebrew",71},
	"hebrew_general_ci":		{"hebrew",16},
	"hp8_bin":		{"hp8",72},
	"hp8_english_ci":		{"hp8",6},
	"keybcs2_bin":		{"keybcs2",73},
	"keybcs2_general_ci":		{"keybcs2",37},
	"koi8r_bin":		{"koi8r",74},
	"koi8r_general_ci":		{"koi8r",7},
	"koi8u_bin":		{"koi8u",75},
	"koi8u_general_ci":		{"koi8u",22},
	"latin1_bin":		{"latin1",47},
	"latin1_danish_ci":		{"latin1",15},
	"latin1_general_ci":		{"latin1",48},
	"latin1_general_cs":		{"latin1",49},
	"latin1_german1_ci":		{"latin1",5},
	"latin1_german2_ci":		{"latin1",31},
	"latin1_spanish_ci":		{"latin1",94},
	"latin1_swedish_ci":		{"latin1",8},
	"latin2_bin":		{"latin2",77},
	"latin2_croatian_ci":		{"latin2",27},
	"latin2_czech_cs":		{"latin2",2},
	"latin2_general_ci":		{"latin2",9},
	"latin2_hungarian_ci":		{"latin2",21},
	"latin5_bin":		{"latin5",78},
	"latin5_turkish_ci":		{"latin5",30},
	"latin7_bin":		{"latin7",79},
	"latin7_estonian_cs":		{"latin7",20},
	"latin7_general_ci":		{"latin7",41},
	"latin7_general_cs":		{"latin7",42},
	"macce_bin":		{"macce",43},
	"macce_general_ci":		{"macce",38},
	"macroman_bin":		{"macroman",53},
	"macroman_general_ci":		{"macroman",39},
	"sjis_bin":		{"sjis",88},
	"sjis_japanese_ci":		{"sjis",13},
	"swe7_bin":		{"swe7",82},
	"swe7_swedish_ci":		{"swe7",10},
	"tis620_bin":		{"tis620",89},
	"tis620_thai_ci":		{"tis620",18},
	"ucs2_bin":		{"ucs2",90},
	"ucs2_croatian_ci":		{"ucs2",149},
	"ucs2_czech_ci":		{"ucs2",138},
	"ucs2_danish_ci":		{"ucs2",139},
	"ucs2_esperanto_ci":		{"ucs2",145},
	"ucs2_estonian_ci":		{"ucs2",134},
	"ucs2_general_ci":		{"ucs2",35},
	"ucs2_general_mysql500_ci":		{"ucs2",159},
	"ucs2_german2_ci":		{"ucs2",148},
	"ucs2_hungarian_ci":		{"ucs2",146},
	"ucs2_icelandic_ci":		{"ucs2",129},
	"ucs2_latvian_ci":		{"ucs2",130},
	"ucs2_lithuanian_ci":		{"ucs2",140},
	"ucs2_persian_ci":		{"ucs2",144},
	"ucs2_polish_ci":		{"ucs2",133},
	"ucs2_romanian_ci":		{"ucs2",131},
	"ucs2_roman_ci":		{"ucs2",143},
	"ucs2_sinhala_ci":		{"ucs2",147},
	"ucs2_slovak_ci":		{"ucs2",141},
	"ucs2_slovenian_ci":		{"ucs2",132},
	"ucs2_spanish2_ci":		{"ucs2",142},
	"ucs2_spanish_ci":		{"ucs2",135},
	"ucs2_swedish_ci":		{"ucs2",136},
	"ucs2_turkish_ci":		{"ucs2",137},
	"ucs2_unicode_520_ci":		{"ucs2",150},
	"ucs2_unicode_ci":		{"ucs2",128},
	"ucs2_vietnamese_ci":		{"ucs2",151},
	"ujis_bin":		{"ujis",91},
	"ujis_japanese_ci":		{"ujis",12},
	"utf16le_bin":		{"utf16le",62},
	"utf16le_general_ci":		{"utf16le",56},
	"utf16_bin":		{"utf16",55},
	"utf16_croatian_ci":		{"utf16",122},
	"utf16_czech_ci":		{"utf16",111},
	"utf16_danish_ci":		{"utf16",112},
	"utf16_esperanto_ci":		{"utf16",118},
	"utf16_estonian_ci":		{"utf16",107},
	"utf16_general_ci":		{"utf16",54},
	"utf16_german2_ci":		{"utf16",121},
	"utf16_hungarian_ci":		{"utf16",119},
	"utf16_icelandic_ci":		{"utf16",102},
	"utf16_latvian_ci":		{"utf16",103},
	"utf16_lithuanian_ci":		{"utf16",113},
	"utf16_persian_ci":		{"utf16",117},
	"utf16_polish_ci":		{"utf16",106},
	"utf16_romanian_ci":		{"utf16",104},
	"utf16_roman_ci":		{"utf16",116},
	"utf16_sinhala_ci":		{"utf16",120},
	"utf16_slovak_ci":		{"utf16",114},
	"utf16_slovenian_ci":		{"utf16",105},
	"utf16_spanish2_ci":		{"utf16",115},
	"utf16_spanish_ci":		{"utf16",108},
	"utf16_swedish_ci":		{"utf16",109},
	"utf16_turkish_ci":		{"utf16",110},
	"utf16_unicode_520_ci":		{"utf16",123},
	"utf16_unicode_ci":		{"utf16",101},
	"utf16_vietnamese_ci":		{"utf16",124},
	"utf32_bin":		{"utf32",61},
	"utf32_croatian_ci":		{"utf32",181},
	"utf32_czech_ci":		{"utf32",170},
	"utf32_danish_ci":		{"utf32",171},
	"utf32_esperanto_ci":		{"utf32",177},
	"utf32_estonian_ci":		{"utf32",166},
	"utf32_general_ci":		{"utf32",60},
	"utf32_german2_ci":		{"utf32",180},
	"utf32_hungarian_ci":		{"utf32",178},
	"utf32_icelandic_ci":		{"utf32",161},
	"utf32_latvian_ci":		{"utf32",162},
	"utf32_lithuanian_ci":		{"utf32",172},
	"utf32_persian_ci":		{"utf32",176},
	"utf32_polish_ci":		{"utf32",165},
	"utf32_romanian_ci":		{"utf32",163},
	"utf32_roman_ci":		{"utf32",175},
	"utf32_sinhala_ci":		{"utf32",179},
	"utf32_slovak_ci":		{"utf32",173},
	"utf32_slovenian_ci":		{"utf32",164},
	"utf32_spanish2_ci":		{"utf32",174},
	"utf32_spanish_ci":		{"utf32",167},
	"utf32_swedish_ci":		{"utf32",168},
	"utf32_turkish_ci":		{"utf32",169},
	"utf32_unicode_520_ci":		{"utf32",182},
	"utf32_unicode_ci":		{"utf32",160},
	"utf32_vietnamese_ci":		{"utf32",183},
	"utf8mb4_0900_ai_ci":		{"utf8mb4",255},
	"utf8mb4_0900_as_ci":		{"utf8mb4",305},
	"utf8mb4_0900_as_cs":		{"utf8mb4",278},
	"utf8mb4_0900_bin":		{"utf8mb4",309},
	"utf8mb4_bin":		{"utf8mb4",46},
	"utf8mb4_croatian_ci":		{"utf8mb4",245},
	"utf8mb4_cs_0900_ai_ci":		{"utf8mb4",266},
	"utf8mb4_cs_0900_as_cs":		{"utf8mb4",289},
	"utf8mb4_czech_ci":		{"utf8mb4",234},
	"utf8mb4_danish_ci":		{"utf8mb4",235},
	"utf8mb4_da_0900_ai_ci":		{"utf8mb4",267},
	"utf8mb4_da_0900_as_cs":		{"utf8mb4",290},
	"utf8mb4_de_pb_0900_ai_ci":		{"utf8mb4",256},
	"utf8mb4_de_pb_0900_as_cs":		{"utf8mb4",279},
	"utf8mb4_eo_0900_ai_ci":		{"utf8mb4",273},
	"utf8mb4_eo_0900_as_cs":		{"utf8mb4",296},
	"utf8mb4_esperanto_ci":		{"utf8mb4",241},
	"utf8mb4_estonian_ci":		{"utf8mb4",230},
	"utf8mb4_es_0900_ai_ci":		{"utf8mb4",263},
	"utf8mb4_es_0900_as_cs":		{"utf8mb4",286},
	"utf8mb4_es_trad_0900_ai_ci":		{"utf8mb4",270},
	"utf8mb4_es_trad_0900_as_cs":		{"utf8mb4",293},
	"utf8mb4_et_0900_ai_ci":		{"utf8mb4",262},
	"utf8mb4_et_0900_as_cs":		{"utf8mb4",285},
	"utf8mb4_general_ci":		{"utf8mb4",45},
	"utf8mb4_german2_ci":		{"utf8mb4",244},
	"utf8mb4_hr_0900_ai_ci":		{"utf8mb4",275},
	"utf8mb4_hr_0900_as_cs":		{"utf8mb4",298},
	"utf8mb4_hungarian_ci":		{"utf8mb4",242},
	"utf8mb4_hu_0900_ai_ci":		{"utf8mb4",274},
	"utf8mb4_hu_0900_as_cs":		{"utf8mb4",297},
	"utf8mb4_icelandic_ci":		{"utf8mb4",225},
	"utf8mb4_is_0900_ai_ci":		{"utf8mb4",257},
	"utf8mb4_is_0900_as_cs":		{"utf8mb4",280},
	"utf8mb4_ja_0900_as_cs":		{"utf8mb4",303},
	"utf8mb4_ja_0900_as_cs_ks":		{"utf8mb4",304},
	"utf8mb4_latvian_ci":		{"utf8mb4",226},
	"utf8mb4_la_0900_ai_ci":		{"utf8mb4",271},
	"utf8mb4_la_0900_as_cs":		{"utf8mb4",294},
	"utf8mb4_lithuanian_ci":		{"utf8mb4",236},
	"utf8mb4_lt_0900_ai_ci":		{"utf8mb4",268},
	"utf8mb4_lt_0900_as_cs":		{"utf8mb4",291},
	"utf8mb4_lv_0900_ai_ci":		{"utf8mb4",258},
	"utf8mb4_lv_0900_as_cs":		{"utf8mb4",281},
	"utf8mb4_persian_ci":		{"utf8mb4",240},
	"utf8mb4_pl_0900_ai_ci":		{"utf8mb4",261},
	"utf8mb4_pl_0900_as_cs":		{"utf8mb4",284},
	"utf8mb4_polish_ci":		{"utf8mb4",229},
	"utf8mb4_romanian_ci":		{"utf8mb4",227},
	"utf8mb4_roman_ci":		{"utf8mb4",239},
	"utf8mb4_ro_0900_ai_ci":		{"utf8mb4",259},
	"utf8mb4_ro_0900_as_cs":		{"utf8mb4",282},
	"utf8mb4_ru_0900_ai_ci":		{"utf8mb4",306},
	"utf8mb4_ru_0900_as_cs":		{"utf8mb4",307},
	"utf8mb4_sinhala_ci":		{"utf8mb4",243},
	"utf8mb4_sk_0900_ai_ci":		{"utf8mb4",269},
	"utf8mb4_sk_0900_as_cs":		{"utf8mb4",292},
	"utf8mb4_slovak_ci":		{"utf8mb4",237},
	"utf8mb4_slovenian_ci":		{"utf8mb4",228},
	"utf8mb4_sl_0900_ai_ci":		{"utf8mb4",260},
	"utf8mb4_sl_0900_as_cs":		{"utf8mb4",283},
	"utf8mb4_spanish2_ci":		{"utf8mb4",238},
	"utf8mb4_spanish_ci":		{"utf8mb4",231},
	"utf8mb4_sv_0900_ai_ci":		{"utf8mb4",264},
	"utf8mb4_sv_0900_as_cs":		{"utf8mb4",287},
	"utf8mb4_swedish_ci":		{"utf8mb4",232},
	"utf8mb4_tr_0900_ai_ci":		{"utf8mb4",265},
	"utf8mb4_tr_0900_as_cs":		{"utf8mb4",288},
	"utf8mb4_turkish_ci":		{"utf8mb4",233},
	"utf8mb4_unicode_520_ci":		{"utf8mb4",246},
	"utf8mb4_unicode_ci":		{"utf8mb4",224},
	"utf8mb4_vietnamese_ci":		{"utf8mb4",247},
	"utf8mb4_vi_0900_ai_ci":		{"utf8mb4",277},
	"utf8mb4_vi_0900_as_cs":		{"utf8mb4",300},
	"utf8mb4_zh_0900_as_cs":		{"utf8mb4",308},
	"utf8_bin":		{"utf8",83},
	"utf8_croatian_ci":		{"utf8",213},
	"utf8_czech_ci":		{"utf8",202},
	"utf8_danish_ci":		{"utf8",203},
	"utf8_esperanto_ci":		{"utf8",209},
	"utf8_estonian_ci":		{"utf8",198},
	"utf8_general_ci":		{"utf8",33},
	"utf8_general_mysql500_ci":		{"utf8",223},
	"utf8_german2_ci":		{"utf8",212},
	"utf8_hungarian_ci":		{"utf8",210},
	"utf8_icelandic_ci":		{"utf8",193},
	"utf8_latvian_ci":		{"utf8",194},
	"utf8_lithuanian_ci":		{"utf8",204},
	"utf8_persian_ci":		{"utf8",208},
	"utf8_polish_ci":		{"utf8",197},
	"utf8_romanian_ci":		{"utf8",195},
	"utf8_roman_ci":		{"utf8",207},
	"utf8_sinhala_ci":		{"utf8",211},
	"utf8_slovak_ci":		{"utf8",205},
	"utf8_slovenian_ci":		{"utf8",196},
	"utf8_spanish2_ci":		{"utf8",206},
	"utf8_spanish_ci":		{"utf8",199},
	"utf8_swedish_ci":		{"utf8",200},
	"utf8_tolower_ci":		{"utf8",76},
	"utf8_turkish_ci":		{"utf8",201},
	"utf8_unicode_520_ci":		{"utf8",214},
	"utf8_unicode_ci":		{"utf8",192},
	"utf8_vietnamese_ci":		{"utf8",215},
}

//mysql client capability
const (
	CLIENT_LONG_PASSWORD 					uint32 = 0x00000001
	CLIENT_FOUND_ROWS	 					uint32 = 0x00000002
	CLIENT_LONG_FLAG	 					uint32 = 0x00000004
	CLIENT_CONNECT_WITH_DB 					uint32 = 0x00000008
	CLIENT_NO_SCHEMA						uint32 = 0x00000010
	CLIENT_COMPRESS     					uint32 = 0x00000020
	CLIENT_LOCAL_FILES 						uint32 = 0x00000080
	CLIENT_IGNORE_SPACE 					uint32 = 0x00000100
	CLIENT_PROTOCOL_41 						uint32 = 0x00000200
	CLIENT_INTERACTIVE 						uint32 = 0x00000400
	CLIENT_SSL 								uint32 = 0x00000800
	CLIENT_IGNORE_SIGPIPE 					uint32 = 0x00001000
	CLIENT_TRANSACTIONS 					uint32 = 0x00002000
	CLIENT_RESERVED 						uint32 = 0x00004000
	CLIENT_SECURE_CONNECTION 				uint32 = 0x00008000
	CLIENT_MULTI_STATEMENTS 				uint32 = 0x00010000
	CLIENT_MULTI_RESULTS 					uint32 = 0x00020000
	CLIENT_PS_MULTI_RESULTS 				uint32 = 0x00040000
	CLIENT_PLUGIN_AUTH 						uint32 = 0x00080000
	CLIENT_CONNECT_ATTRS 					uint32 = 0x00100000
	CLIENT_PLUGIN_AUTH_LENENC_CLIENT_DATA 	uint32 = 0x00200000
	CLIENT_CAN_HANDLE_EXPIRED_PASSWORDS 	uint32 = 0x00400000
	CLIENT_SESSION_TRACK 					uint32 = 0x00800000
	CLIENT_DEPRECATE_EOF 					uint32 = 0x01000000
)

//server status
const (
	SERVER_STATUS_IN_TRANS				uint16 = 0x0001
	SERVER_STATUS_AUTOCOMMIT			uint16 = 0x0002
	SERVER_MORE_RESULTS_EXISTS			uint16 = 0x0008
	SERVER_STATUS_NO_GOOD_INDEX_USED	uint16 = 0x0010
	SERVER_STATUS_NO_INDEX_USED			uint16 = 0x0020
	SERVER_STATUS_CURSOR_EXISTS			uint16 = 0x0040
	SERVER_STATUS_LAST_ROW_SENT			uint16 = 0x0080
	SERVER_STATUS_DB_DROPPED			uint16 = 0x0100
	SERVER_STATUS_NO_BACKSLASH_ESCAPES	uint16 = 0x0200
	SERVER_STATUS_METADATA_CHANGED		uint16 = 0x0400
	SERVER_QUERY_WAS_SLOW				uint16 = 0x0800
	SERVER_PS_OUT_PARAMS				uint16 = 0x1000
	SERVER_STATUS_IN_TRANS_READONLY		uint16 = 0x2000
	SERVER_SESSION_STATE_CHANGED		uint16 = 0x4000
)

//text protocol in mysql client protocol
//iteration command
const(
	COM_SLEEP 				uint8 = 0x00
	COM_QUIT 				uint8 = 0x01
	COM_INIT_DB				uint8 = 0x02
	COM_QUERY				uint8 = 0x03
	COM_FIELD_LIST			uint8 = 0x04
	COM_CREATE_DB			uint8 = 0x05
	COM_DROP_DB				uint8 = 0x06
	COM_REFRESH				uint8 = 0x07
	COM_SHUTDOWN			uint8 = 0x08
	COM_STATISTICS			uint8 = 0x09
	COM_PROCESS_INFO		uint8 = 0x0a
	COM_CONNECT				uint8 = 0x0b
	COM_PROCESS_KILL		uint8 = 0x0c
	COM_DEBUG				uint8 = 0x0d
	COM_PING				uint8 = 0x0e
	COM_TIME				uint8 = 0x0f
	COM_DELAYED_INSERT		uint8 = 0x10
	COM_CHANGE_USER			uint8 = 0x11
	COM_STMT_PREPARE		uint8 = 0x16
	COM_STMT_EXECUTE		uint8 = 0x17
	COM_STMT_SEND_LONG_DATA	uint8 = 0x18
	COM_STMT_CLOSE			uint8 = 0x19
	COM_STMT_RESET			uint8 = 0x1a
	COM_SET_OPTION			uint8 = 0x1b
	COM_STMT_FETCH			uint8 = 0x1c
	COM_DAEMON				uint8 = 0x1d
	COM_RESET_CONNECTION	uint8 = 0x1f
)

/*
Mysql Error Code
information from https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html
mysql version 8.0.23
usually it is in the directory /usr/local/include/mysql/mysqld_error.h

Error information includes several elements: an error code, SQLSTATE value, and message string.
	Error code: This value is numeric. It is MySQL-specific and is not portable to other database systems.
	SQLSTATE value: This value is a five-character string (for example, '42S02'). SQLSTATE values are taken from ANSI SQL and ODBC and are more standardized than the numeric error codes.
	Message string: This string provides a textual description of the error.
*/
const (
	//1 to 999: Global error codes. This error code range is called “global” because it is a shared range that is used by the server as well as by clients.
	//When an error in this range originates on the server side, the server writes it to the error log, padding the error code with leading zeros to six digits and adding a prefix of MY-.
	//When an error in this range originates on the client side, the client library makes it available to the client program with no zero-padding or prefix.

	// No such code here.

	//1,000 to 1,999: Server error codes reserved for messages sent to clients.
	//OBSOLETE_ER_HASHCHK 	uint16 = 1000
	//OBSOLETE_ER_NISAMCHK 	uint16 = 1001
	ER_NO 					uint16 = 1002
	ER_YES 					uint16 = 1003
	ER_CANT_CREATE_FILE 	uint16 = 1004
	ER_CANT_CREATE_TABLE 	uint16 = 1005
	ER_CANT_CREATE_DB 		uint16 = 1006
	ER_DB_CREATE_EXISTS 	uint16 = 1007
	ER_DB_DROP_EXISTS uint16 = 1008
	//OBSOLETE_ER_DB_DROP_DELETE uint16 = 1009
	ER_DB_DROP_RMDIR uint16 = 1010
	//OBSOLETE_ER_CANT_DELETE_FILE uint16 = 1011
	ER_CANT_FIND_SYSTEM_REC uint16 = 1012
	ER_CANT_GET_STAT uint16 = 1013
	//OBSOLETE_ER_CANT_GET_WD uint16 = 1014
	ER_CANT_LOCK uint16 = 1015
	ER_CANT_OPEN_FILE uint16 = 1016
	ER_FILE_NOT_FOUND uint16 = 1017
	ER_CANT_READ_DIR uint16 = 1018
	//OBSOLETE_ER_CANT_SET_WD uint16 = 1019
	ER_CHECKREAD uint16 = 1020
	//OBSOLETE_ER_DISK_FULL uint16 = 1021
	ER_DUP_KEY uint16 = 1022
	//OBSOLETE_ER_ERROR_ON_CLOSE uint16 = 1023
	ER_ERROR_ON_READ uint16 = 1024
	ER_ERROR_ON_RENAME uint16 = 1025
	ER_ERROR_ON_WRITE uint16 = 1026
	ER_FILE_USED uint16 = 1027
	//OBSOLETE_ER_FILSORT_ABORT uint16 = 1028
	//OBSOLETE_ER_FORM_NOT_FOUND uint16 = 1029
	ER_GET_ERRNO uint16 = 1030
	ER_ILLEGAL_HA uint16 = 1031
	ER_KEY_NOT_FOUND uint16 = 1032
	ER_NOT_FORM_FILE uint16 = 1033
	ER_NOT_KEYFILE uint16 = 1034
	ER_OLD_KEYFILE uint16 = 1035
	ER_OPEN_AS_READONLY uint16 = 1036
	ER_OUTOFMEMORY uint16 = 1037
	ER_OUT_OF_SORTMEMORY uint16 = 1038
	//OBSOLETE_ER_UNEXPECTED_EOF uint16 = 1039
	ER_CON_COUNT_ERROR uint16 = 1040
	ER_OUT_OF_RESOURCES uint16 = 1041
	ER_BAD_HOST_ERROR uint16 = 1042
	ER_HANDSHAKE_ERROR uint16 = 1043
	ER_DBACCESS_DENIED_ERROR uint16 = 1044
	ER_ACCESS_DENIED_ERROR uint16 = 1045
	ER_NO_DB_ERROR uint16 = 1046
	ER_UNKNOWN_COM_ERROR uint16 = 1047
	ER_BAD_NULL_ERROR uint16 = 1048
	ER_BAD_DB_ERROR uint16 = 1049
	ER_TABLE_EXISTS_ERROR uint16 = 1050
	ER_BAD_TABLE_ERROR uint16 = 1051
	ER_NON_UNIQ_ERROR uint16 = 1052
	ER_SERVER_SHUTDOWN uint16 = 1053
	ER_BAD_FIELD_ERROR uint16 = 1054
	ER_WRONG_FIELD_WITH_GROUP uint16 = 1055
	ER_WRONG_GROUP_FIELD uint16 = 1056
	ER_WRONG_SUM_SELECT uint16 = 1057
	ER_WRONG_VALUE_COUNT uint16 = 1058
	ER_TOO_LONG_IDENT uint16 = 1059
	ER_DUP_FIELDNAME uint16 = 1060
	ER_DUP_KEYNAME uint16 = 1061
	ER_DUP_ENTRY uint16 = 1062
	ER_WRONG_FIELD_SPEC uint16 = 1063
	ER_PARSE_ERROR uint16 = 1064
	ER_EMPTY_QUERY uint16 = 1065
	ER_NONUNIQ_TABLE uint16 = 1066
	ER_INVALID_DEFAULT uint16 = 1067
	ER_MULTIPLE_PRI_KEY uint16 = 1068
	ER_TOO_MANY_KEYS uint16 = 1069
	ER_TOO_MANY_KEY_PARTS uint16 = 1070
	ER_TOO_LONG_KEY uint16 = 1071
	ER_KEY_COLUMN_DOES_NOT_EXITS uint16 = 1072
	ER_BLOB_USED_AS_KEY uint16 = 1073
	ER_TOO_BIG_FIELDLENGTH uint16 = 1074
	ER_WRONG_AUTO_KEY uint16 = 1075
	ER_READY uint16 = 1076
	//OBSOLETE_ER_NORMAL_SHUTDOWN uint16 = 1077
	//OBSOLETE_ER_GOT_SIGNAL uint16 = 1078
	ER_SHUTDOWN_COMPLETE uint16 = 1079
	ER_FORCING_CLOSE uint16 = 1080
	ER_IPSOCK_ERROR uint16 = 1081
	ER_NO_SUCH_INDEX uint16 = 1082
	ER_WRONG_FIELD_TERMINATORS uint16 = 1083
	ER_BLOBS_AND_NO_TERMINATED uint16 = 1084
	ER_TEXTFILE_NOT_READABLE uint16 = 1085
	ER_FILE_EXISTS_ERROR uint16 = 1086
	ER_LOAD_INFO uint16 = 1087
	ER_ALTER_INFO uint16 = 1088
	ER_WRONG_SUB_KEY uint16 = 1089
	ER_CANT_REMOVE_ALL_FIELDS uint16 = 1090
	ER_CANT_DROP_FIELD_OR_KEY uint16 = 1091
	ER_INSERT_INFO uint16 = 1092
	ER_UPDATE_TABLE_USED uint16 = 1093
	ER_NO_SUCH_THREAD uint16 = 1094
	ER_KILL_DENIED_ERROR uint16 = 1095
	ER_NO_TABLES_USED uint16 = 1096
	ER_TOO_BIG_SET uint16 = 1097
	ER_NO_UNIQUE_LOGFILE uint16 = 1098
	ER_TABLE_NOT_LOCKED_FOR_WRITE uint16 = 1099
	ER_TABLE_NOT_LOCKED uint16 = 1100
	ER_BLOB_CANT_HAVE_DEFAULT uint16 = 1101
	ER_WRONG_DB_NAME uint16 = 1102
	ER_WRONG_TABLE_NAME uint16 = 1103
	ER_TOO_BIG_SELECT uint16 = 1104
	ER_UNKNOWN_ERROR uint16 = 1105
	ER_UNKNOWN_PROCEDURE uint16 = 1106
	ER_WRONG_PARAMCOUNT_TO_PROCEDURE uint16 = 1107
	ER_WRONG_PARAMETERS_TO_PROCEDURE uint16 = 1108
	ER_UNKNOWN_TABLE uint16 = 1109
	ER_FIELD_SPECIFIED_TWICE uint16 = 1110
	ER_INVALID_GROUP_FUNC_USE uint16 = 1111
	ER_UNSUPPORTED_EXTENSION uint16 = 1112
	ER_TABLE_MUST_HAVE_COLUMNS uint16 = 1113
	ER_RECORD_FILE_FULL uint16 = 1114
	ER_UNKNOWN_CHARACTER_SET uint16 = 1115
	ER_TOO_MANY_TABLES uint16 = 1116
	ER_TOO_MANY_FIELDS uint16 = 1117
	ER_TOO_BIG_ROWSIZE uint16 = 1118
	ER_STACK_OVERRUN uint16 = 1119
	ER_WRONG_OUTER_JOIN_UNUSED uint16 = 1120
	ER_NULL_COLUMN_IN_INDEX uint16 = 1121
	ER_CANT_FIND_UDF uint16 = 1122
	ER_CANT_INITIALIZE_UDF uint16 = 1123
	ER_UDF_NO_PATHS uint16 = 1124
	ER_UDF_EXISTS uint16 = 1125
	ER_CANT_OPEN_LIBRARY uint16 = 1126
	ER_CANT_FIND_DL_ENTRY uint16 = 1127
	ER_FUNCTION_NOT_DEFINED uint16 = 1128
	ER_HOST_IS_BLOCKED uint16 = 1129
	ER_HOST_NOT_PRIVILEGED uint16 = 1130
	ER_PASSWORD_ANONYMOUS_USER uint16 = 1131
	ER_PASSWORD_NOT_ALLOWED uint16 = 1132
	ER_PASSWORD_NO_MATCH uint16 = 1133
	ER_UPDATE_INFO uint16 = 1134
	ER_CANT_CREATE_THREAD uint16 = 1135
	ER_WRONG_VALUE_COUNT_ON_ROW uint16 = 1136
	ER_CANT_REOPEN_TABLE uint16 = 1137
	ER_INVALID_USE_OF_NULL uint16 = 1138
	ER_REGEXP_ERROR uint16 = 1139
	ER_MIX_OF_GROUP_FUNC_AND_FIELDS uint16 = 1140
	ER_NONEXISTING_GRANT uint16 = 1141
	ER_TABLEACCESS_DENIED_ERROR uint16 = 1142
	ER_COLUMNACCESS_DENIED_ERROR uint16 = 1143
	ER_ILLEGAL_GRANT_FOR_TABLE uint16 = 1144
	ER_GRANT_WRONG_HOST_OR_USER uint16 = 1145
	ER_NO_SUCH_TABLE uint16 = 1146
	ER_NONEXISTING_TABLE_GRANT uint16 = 1147
	ER_NOT_ALLOWED_COMMAND uint16 = 1148
	ER_SYNTAX_ERROR uint16 = 1149
	//OBSOLETE_ER_UNUSED1 uint16 = 1150
	//OBSOLETE_ER_UNUSED2 uint16 = 1151
	ER_ABORTING_CONNECTION uint16 = 1152
	ER_NET_PACKET_TOO_LARGE uint16 = 1153
	ER_NET_READ_ERROR_FROM_PIPE uint16 = 1154
	ER_NET_FCNTL_ERROR uint16 = 1155
	ER_NET_PACKETS_OUT_OF_ORDER uint16 = 1156
	ER_NET_UNCOMPRESS_ERROR uint16 = 1157
	ER_NET_READ_ERROR uint16 = 1158
	ER_NET_READ_INTERRUPTED uint16 = 1159
	ER_NET_ERROR_ON_WRITE uint16 = 1160
	ER_NET_WRITE_INTERRUPTED uint16 = 1161
	ER_TOO_LONG_STRING uint16 = 1162
	ER_TABLE_CANT_HANDLE_BLOB uint16 = 1163
	ER_TABLE_CANT_HANDLE_AUTO_INCREMENT uint16 = 1164
	//OBSOLETE_ER_UNUSED3 uint16 = 1165
	ER_WRONG_COLUMN_NAME uint16 = 1166
	ER_WRONG_KEY_COLUMN uint16 = 1167
	ER_WRONG_MRG_TABLE uint16 = 1168
	ER_DUP_UNIQUE uint16 = 1169
	ER_BLOB_KEY_WITHOUT_LENGTH uint16 = 1170
	ER_PRIMARY_CANT_HAVE_NULL uint16 = 1171
	ER_TOO_MANY_ROWS uint16 = 1172
	ER_REQUIRES_PRIMARY_KEY uint16 = 1173
	//OBSOLETE_ER_NO_RAID_COMPILED uint16 = 1174
	ER_UPDATE_WITHOUT_KEY_IN_SAFE_MODE uint16 = 1175
	ER_KEY_DOES_NOT_EXITS uint16 = 1176
	ER_CHECK_NO_SUCH_TABLE uint16 = 1177
	ER_CHECK_NOT_IMPLEMENTED uint16 = 1178
	ER_CANT_DO_THIS_DURING_AN_TRANSACTION uint16 = 1179
	ER_ERROR_DURING_COMMIT uint16 = 1180
	ER_ERROR_DURING_ROLLBACK uint16 = 1181
	ER_ERROR_DURING_FLUSH_LOGS uint16 = 1182
	//OBSOLETE_ER_ERROR_DURING_CHECKPOINT uint16 = 1183
	ER_NEW_ABORTING_CONNECTION uint16 = 1184
	//OBSOLETE_ER_DUMP_NOT_IMPLEMENTED uint16 = 1185
	//OBSOLETE_ER_FLUSH_MASTER_BINLOG_CLOSED uint16 = 1186
	//OBSOLETE_ER_INDEX_REBUILD uint16 = 1187
	ER_MASTER uint16 = 1188
	ER_MASTER_NET_READ uint16 = 1189
	ER_MASTER_NET_WRITE uint16 = 1190
	ER_FT_MATCHING_KEY_NOT_FOUND uint16 = 1191
	ER_LOCK_OR_ACTIVE_TRANSACTION uint16 = 1192
	ER_UNKNOWN_SYSTEM_VARIABLE uint16 = 1193
	ER_CRASHED_ON_USAGE uint16 = 1194
	ER_CRASHED_ON_REPAIR uint16 = 1195
	ER_WARNING_NOT_COMPLETE_ROLLBACK uint16 = 1196
	ER_TRANS_CACHE_FULL uint16 = 1197
	//OBSOLETE_ER_SLAVE_MUST_STOP uint16 = 1198
	ER_SLAVE_NOT_RUNNING uint16 = 1199
	ER_BAD_SLAVE uint16 = 1200
	ER_MASTER_INFO uint16 = 1201
	ER_SLAVE_THREAD uint16 = 1202
	ER_TOO_MANY_USER_CONNECTIONS uint16 = 1203
	ER_SET_CONSTANTS_ONLY uint16 = 1204
	ER_LOCK_WAIT_TIMEOUT uint16 = 1205
	ER_LOCK_TABLE_FULL uint16 = 1206
	ER_READ_ONLY_TRANSACTION uint16 = 1207
	//OBSOLETE_ER_DROP_DB_WITH_READ_LOCK uint16 = 1208
	//OBSOLETE_ER_CREATE_DB_WITH_READ_LOCK uint16 = 1209
	ER_WRONG_ARGUMENTS uint16 = 1210
	ER_NO_PERMISSION_TO_CREATE_USER uint16 = 1211
	//OBSOLETE_ER_UNION_TABLES_IN_DIFFERENT_DIR uint16 = 1212
	ER_LOCK_DEADLOCK uint16 = 1213
	ER_TABLE_CANT_HANDLE_FT uint16 = 1214
	ER_CANNOT_ADD_FOREIGN uint16 = 1215
	ER_NO_REFERENCED_ROW uint16 = 1216
	ER_ROW_IS_REFERENCED uint16 = 1217
	ER_CONNECT_TO_MASTER uint16 = 1218
	//OBSOLETE_ER_QUERY_ON_MASTER uint16 = 1219
	ER_ERROR_WHEN_EXECUTING_COMMAND uint16 = 1220
	ER_WRONG_USAGE uint16 = 1221
	ER_WRONG_NUMBER_OF_COLUMNS_IN_SELECT uint16 = 1222
	ER_CANT_UPDATE_WITH_READLOCK uint16 = 1223
	ER_MIXING_NOT_ALLOWED uint16 = 1224
	ER_DUP_ARGUMENT uint16 = 1225
	ER_USER_LIMIT_REACHED uint16 = 1226
	ER_SPECIFIC_ACCESS_DENIED_ERROR uint16 = 1227
	ER_LOCAL_VARIABLE uint16 = 1228
	ER_GLOBAL_VARIABLE uint16 = 1229
	ER_NO_DEFAULT uint16 = 1230
	ER_WRONG_VALUE_FOR_VAR uint16 = 1231
	ER_WRONG_TYPE_FOR_VAR uint16 = 1232
	ER_VAR_CANT_BE_READ uint16 = 1233
	ER_CANT_USE_OPTION_HERE uint16 = 1234
	ER_NOT_SUPPORTED_YET uint16 = 1235
	ER_MASTER_FATAL_ERROR_READING_BINLOG uint16 = 1236
	ER_SLAVE_IGNORED_TABLE uint16 = 1237
	ER_INCORRECT_GLOBAL_LOCAL_VAR uint16 = 1238
	ER_WRONG_FK_DEF uint16 = 1239
	ER_KEY_REF_DO_NOT_MATCH_TABLE_REF uint16 = 1240
	ER_OPERAND_COLUMNS uint16 = 1241
	ER_SUBQUERY_NO_1_ROW uint16 = 1242
	ER_UNKNOWN_STMT_HANDLER uint16 = 1243
	ER_CORRUPT_HELP_DB uint16 = 1244
	//OBSOLETE_ER_CYCLIC_REFERENCE uint16 = 1245
	ER_AUTO_CONVERT uint16 = 1246
	ER_ILLEGAL_REFERENCE uint16 = 1247
	ER_DERIVED_MUST_HAVE_ALIAS uint16 = 1248
	ER_SELECT_REDUCED uint16 = 1249
	ER_TABLENAME_NOT_ALLOWED_HERE uint16 = 1250
	ER_NOT_SUPPORTED_AUTH_MODE uint16 = 1251
	ER_SPATIAL_CANT_HAVE_NULL uint16 = 1252
	ER_COLLATION_CHARSET_MISMATCH uint16 = 1253
	//OBSOLETE_ER_SLAVE_WAS_RUNNING uint16 = 1254
	//OBSOLETE_ER_SLAVE_WAS_NOT_RUNNING uint16 = 1255
	ER_TOO_BIG_FOR_UNCOMPRESS uint16 = 1256
	ER_ZLIB_Z_MEM_ERROR uint16 = 1257
	ER_ZLIB_Z_BUF_ERROR uint16 = 1258
	ER_ZLIB_Z_DATA_ERROR uint16 = 1259
	ER_CUT_VALUE_GROUP_CONCAT uint16 = 1260
	ER_WARN_TOO_FEW_RECORDS uint16 = 1261
	ER_WARN_TOO_MANY_RECORDS uint16 = 1262
	ER_WARN_NULL_TO_NOTNULL uint16 = 1263
	ER_WARN_DATA_OUT_OF_RANGE uint16 = 1264
	WARN_DATA_TRUNCATED uint16 = 1265
	ER_WARN_USING_OTHER_HANDLER uint16 = 1266
	ER_CANT_AGGREGATE_2COLLATIONS uint16 = 1267
	//OBSOLETE_ER_DROP_USER uint16 = 1268
	ER_REVOKE_GRANTS uint16 = 1269
	ER_CANT_AGGREGATE_3COLLATIONS uint16 = 1270
	ER_CANT_AGGREGATE_NCOLLATIONS uint16 = 1271
	ER_VARIABLE_IS_NOT_STRUCT uint16 = 1272
	ER_UNKNOWN_COLLATION uint16 = 1273
	ER_SLAVE_IGNORED_SSL_PARAMS uint16 = 1274
	//OBSOLETE_ER_SERVER_IS_IN_SECURE_AUTH_MODE uint16 = 1275
	ER_WARN_FIELD_RESOLVED uint16 = 1276
	ER_BAD_SLAVE_UNTIL_COND uint16 = 1277
	ER_MISSING_SKIP_SLAVE uint16 = 1278
	ER_UNTIL_COND_IGNORED uint16 = 1279
	ER_WRONG_NAME_FOR_INDEX uint16 = 1280
	ER_WRONG_NAME_FOR_CATALOG uint16 = 1281
	//OBSOLETE_ER_WARN_QC_RESIZE uint16 = 1282
	ER_BAD_FT_COLUMN uint16 = 1283
	ER_UNKNOWN_KEY_CACHE uint16 = 1284
	ER_WARN_HOSTNAME_WONT_WORK uint16 = 1285
	ER_UNKNOWN_STORAGE_ENGINE uint16 = 1286
	ER_WARN_DEPRECATED_SYNTAX uint16 = 1287
	ER_NON_UPDATABLE_TABLE uint16 = 1288
	ER_FEATURE_DISABLED uint16 = 1289
	ER_OPTION_PREVENTS_STATEMENT uint16 = 1290
	ER_DUPLICATED_VALUE_IN_TYPE uint16 = 1291
	ER_TRUNCATED_WRONG_VALUE uint16 = 1292
	//OBSOLETE_ER_TOO_MUCH_AUTO_TIMESTAMP_COLS uint16 = 1293
	ER_INVALID_ON_UPDATE uint16 = 1294
	ER_UNSUPPORTED_PS uint16 = 1295
	ER_GET_ERRMSG uint16 = 1296
	ER_GET_TEMPORARY_ERRMSG uint16 = 1297
	ER_UNKNOWN_TIME_ZONE uint16 = 1298
	ER_WARN_INVALID_TIMESTAMP uint16 = 1299
	ER_INVALID_CHARACTER_STRING uint16 = 1300
	ER_WARN_ALLOWED_PACKET_OVERFLOWED uint16 = 1301
	ER_CONFLICTING_DECLARATIONS uint16 = 1302
	ER_SP_NO_RECURSIVE_CREATE uint16 = 1303
	ER_SP_ALREADY_EXISTS uint16 = 1304
	ER_SP_DOES_NOT_EXIST uint16 = 1305
	ER_SP_DROP_FAILED uint16 = 1306
	ER_SP_STORE_FAILED uint16 = 1307
	ER_SP_LILABEL_MISMATCH uint16 = 1308
	ER_SP_LABEL_REDEFINE uint16 = 1309
	ER_SP_LABEL_MISMATCH uint16 = 1310
	ER_SP_UNINIT_VAR uint16 = 1311
	ER_SP_BADSELECT uint16 = 1312
	ER_SP_BADRETURN uint16 = 1313
	ER_SP_BADSTATEMENT uint16 = 1314
	ER_UPDATE_LOG_DEPRECATED_IGNORED uint16 = 1315
	ER_UPDATE_LOG_DEPRECATED_TRANSLATED uint16 = 1316
	ER_QUERY_INTERRUPTED uint16 = 1317
	ER_SP_WRONG_NO_OF_ARGS uint16 = 1318
	ER_SP_COND_MISMATCH uint16 = 1319
	ER_SP_NORETURN uint16 = 1320
	ER_SP_NORETURNEND uint16 = 1321
	ER_SP_BAD_CURSOR_QUERY uint16 = 1322
	ER_SP_BAD_CURSOR_SELECT uint16 = 1323
	ER_SP_CURSOR_MISMATCH uint16 = 1324
	ER_SP_CURSOR_ALREADY_OPEN uint16 = 1325
	ER_SP_CURSOR_NOT_OPEN uint16 = 1326
	ER_SP_UNDECLARED_VAR uint16 = 1327
	ER_SP_WRONG_NO_OF_FETCH_ARGS uint16 = 1328
	ER_SP_FETCH_NO_DATA uint16 = 1329
	ER_SP_DUP_PARAM uint16 = 1330
	ER_SP_DUP_VAR uint16 = 1331
	ER_SP_DUP_COND uint16 = 1332
	ER_SP_DUP_CURS uint16 = 1333
	ER_SP_CANT_ALTER uint16 = 1334
	ER_SP_SUBSELECT_NYI uint16 = 1335
	ER_STMT_NOT_ALLOWED_IN_SF_OR_TRG uint16 = 1336
	ER_SP_VARCOND_AFTER_CURSHNDLR uint16 = 1337
	ER_SP_CURSOR_AFTER_HANDLER uint16 = 1338
	ER_SP_CASE_NOT_FOUND uint16 = 1339
	ER_FPARSER_TOO_BIG_FILE uint16 = 1340
	ER_FPARSER_BAD_HEADER uint16 = 1341
	ER_FPARSER_EOF_IN_COMMENT uint16 = 1342
	ER_FPARSER_ERROR_IN_PARAMETER uint16 = 1343
	ER_FPARSER_EOF_IN_UNKNOWN_PARAMETER uint16 = 1344
	ER_VIEW_NO_EXPLAIN uint16 = 1345
	//OBSOLETE_ER_FRM_UNKNOWN_TYPE uint16 = 1346
	ER_WRONG_OBJECT uint16 = 1347
	ER_NONUPDATEABLE_COLUMN uint16 = 1348
	//OBSOLETE_ER_VIEW_SELECT_DERIVED_UNUSED uint16 = 1349
	ER_VIEW_SELECT_CLAUSE uint16 = 1350
	ER_VIEW_SELECT_VARIABLE uint16 = 1351
	ER_VIEW_SELECT_TMPTABLE uint16 = 1352
	ER_VIEW_WRONG_LIST uint16 = 1353
	ER_WARN_VIEW_MERGE uint16 = 1354
	ER_WARN_VIEW_WITHOUT_KEY uint16 = 1355
	ER_VIEW_INVALID uint16 = 1356
	ER_SP_NO_DROP_SP uint16 = 1357
	//OBSOLETE_ER_SP_GOTO_IN_HNDLR uint16 = 1358
	ER_TRG_ALREADY_EXISTS uint16 = 1359
	ER_TRG_DOES_NOT_EXIST uint16 = 1360
	ER_TRG_ON_VIEW_OR_TEMP_TABLE uint16 = 1361
	ER_TRG_CANT_CHANGE_ROW uint16 = 1362
	ER_TRG_NO_SUCH_ROW_IN_TRG uint16 = 1363
	ER_NO_DEFAULT_FOR_FIELD uint16 = 1364
	ER_DIVISION_BY_ZERO uint16 = 1365
	ER_TRUNCATED_WRONG_VALUE_FOR_FIELD uint16 = 1366
	ER_ILLEGAL_VALUE_FOR_TYPE uint16 = 1367
	ER_VIEW_NONUPD_CHECK uint16 = 1368
	ER_VIEW_CHECK_FAILED uint16 = 1369
	ER_PROCACCESS_DENIED_ERROR uint16 = 1370
	ER_RELAY_LOG_FAIL uint16 = 1371
	//OBSOLETE_ER_PASSWD_LENGTH uint16 = 1372
	ER_UNKNOWN_TARGET_BINLOG uint16 = 1373
	ER_IO_ERR_LOG_INDEX_READ uint16 = 1374
	ER_BINLOG_PURGE_PROHIBITED uint16 = 1375
	ER_FSEEK_FAIL uint16 = 1376
	ER_BINLOG_PURGE_FATAL_ERR uint16 = 1377
	ER_LOG_IN_USE uint16 = 1378
	ER_LOG_PURGE_UNKNOWN_ERR uint16 = 1379
	ER_RELAY_LOG_INIT uint16 = 1380
	ER_NO_BINARY_LOGGING uint16 = 1381
	ER_RESERVED_SYNTAX uint16 = 1382
	//OBSOLETE_ER_WSAS_FAILED uint16 = 1383
	//OBSOLETE_ER_DIFF_GROUPS_PROC uint16 = 1384
	//OBSOLETE_ER_NO_GROUP_FOR_PROC uint16 = 1385
	//OBSOLETE_ER_ORDER_WITH_PROC uint16 = 1386
	//OBSOLETE_ER_LOGGING_PROHIBIT_CHANGING_OF uint16 = 1387
	//OBSOLETE_ER_NO_FILE_MAPPING uint16 = 1388
	//OBSOLETE_ER_WRONG_MAGIC uint16 = 1389
	ER_PS_MANY_PARAM uint16 = 1390
	ER_KEY_PART_0 uint16 = 1391
	ER_VIEW_CHECKSUM uint16 = 1392
	ER_VIEW_MULTIUPDATE uint16 = 1393
	ER_VIEW_NO_INSERT_FIELD_LIST uint16 = 1394
	ER_VIEW_DELETE_MERGE_VIEW uint16 = 1395
	ER_CANNOT_USER uint16 = 1396
	ER_XAER_NOTA uint16 = 1397
	ER_XAER_INVAL uint16 = 1398
	ER_XAER_RMFAIL uint16 = 1399
	ER_XAER_OUTSIDE uint16 = 1400
	ER_XAER_RMERR uint16 = 1401
	ER_XA_RBROLLBACK uint16 = 1402
	ER_NONEXISTING_PROC_GRANT uint16 = 1403
	ER_PROC_AUTO_GRANT_FAIL uint16 = 1404
	ER_PROC_AUTO_REVOKE_FAIL uint16 = 1405
	ER_DATA_TOO_LONG uint16 = 1406
	ER_SP_BAD_SQLSTATE uint16 = 1407
	ER_STARTUP uint16 = 1408
	ER_LOAD_FROM_FIXED_SIZE_ROWS_TO_VAR uint16 = 1409
	ER_CANT_CREATE_USER_WITH_GRANT uint16 = 1410
	ER_WRONG_VALUE_FOR_TYPE uint16 = 1411
	ER_TABLE_DEF_CHANGED uint16 = 1412
	ER_SP_DUP_HANDLER uint16 = 1413
	ER_SP_NOT_VAR_ARG uint16 = 1414
	ER_SP_NO_RETSET uint16 = 1415
	ER_CANT_CREATE_GEOMETRY_OBJECT uint16 = 1416
	//OBSOLETE_ER_FAILED_ROUTINE_BREAK_BINLOG uint16 = 1417
	ER_BINLOG_UNSAFE_ROUTINE uint16 = 1418
	ER_BINLOG_CREATE_ROUTINE_NEED_SUPER uint16 = 1419
	//OBSOLETE_ER_EXEC_STMT_WITH_OPEN_CURSOR uint16 = 1420
	ER_STMT_HAS_NO_OPEN_CURSOR uint16 = 1421
	ER_COMMIT_NOT_ALLOWED_IN_SF_OR_TRG uint16 = 1422
	ER_NO_DEFAULT_FOR_VIEW_FIELD uint16 = 1423
	ER_SP_NO_RECURSION uint16 = 1424
	ER_TOO_BIG_SCALE uint16 = 1425
	ER_TOO_BIG_PRECISION uint16 = 1426
	ER_M_BIGGER_THAN_D uint16 = 1427
	ER_WRONG_LOCK_OF_SYSTEM_TABLE uint16 = 1428
	ER_CONNECT_TO_FOREIGN_DATA_SOURCE uint16 = 1429
	ER_QUERY_ON_FOREIGN_DATA_SOURCE uint16 = 1430
	ER_FOREIGN_DATA_SOURCE_DOESNT_EXIST uint16 = 1431
	ER_FOREIGN_DATA_STRING_INVALID_CANT_CREATE uint16 = 1432
	ER_FOREIGN_DATA_STRING_INVALID uint16 = 1433
	//OBSOLETE_ER_CANT_CREATE_FEDERATED_TABLE uint16 = 1434
	ER_TRG_IN_WRONG_SCHEMA uint16 = 1435
	ER_STACK_OVERRUN_NEED_MORE uint16 = 1436
	ER_TOO_LONG_BODY uint16 = 1437
	ER_WARN_CANT_DROP_DEFAULT_KEYCACHE uint16 = 1438
	ER_TOO_BIG_DISPLAYWIDTH uint16 = 1439
	ER_XAER_DUPID uint16 = 1440
	ER_DATETIME_FUNCTION_OVERFLOW uint16 = 1441
	ER_CANT_UPDATE_USED_TABLE_IN_SF_OR_TRG uint16 = 1442
	ER_VIEW_PREVENT_UPDATE uint16 = 1443
	ER_PS_NO_RECURSION uint16 = 1444
	ER_SP_CANT_SET_AUTOCOMMIT uint16 = 1445
	//OBSOLETE_ER_MALFORMED_DEFINER uint16 = 1446
	ER_VIEW_FRM_NO_USER uint16 = 1447
	ER_VIEW_OTHER_USER uint16 = 1448
	ER_NO_SUCH_USER uint16 = 1449
	ER_FORBID_SCHEMA_CHANGE uint16 = 1450
	ER_ROW_IS_REFERENCED_2 uint16 = 1451
	ER_NO_REFERENCED_ROW_2 uint16 = 1452
	ER_SP_BAD_VAR_SHADOW uint16 = 1453
	ER_TRG_NO_DEFINER uint16 = 1454
	ER_OLD_FILE_FORMAT uint16 = 1455
	ER_SP_RECURSION_LIMIT uint16 = 1456
	//OBSOLETE_ER_SP_PROC_TABLE_CORRUPT uint16 = 1457
	ER_SP_WRONG_NAME uint16 = 1458
	ER_TABLE_NEEDS_UPGRADE uint16 = 1459
	ER_SP_NO_AGGREGATE uint16 = 1460
	ER_MAX_PREPARED_STMT_COUNT_REACHED uint16 = 1461
	ER_VIEW_RECURSIVE uint16 = 1462
	ER_NON_GROUPING_FIELD_USED uint16 = 1463
	ER_TABLE_CANT_HANDLE_SPKEYS uint16 = 1464
	ER_NO_TRIGGERS_ON_SYSTEM_SCHEMA uint16 = 1465
	ER_REMOVED_SPACES uint16 = 1466
	ER_AUTOINC_READ_FAILED uint16 = 1467
	ER_USERNAME uint16 = 1468
	ER_HOSTNAME uint16 = 1469
	ER_WRONG_STRING_LENGTH uint16 = 1470
	ER_NON_INSERTABLE_TABLE uint16 = 1471
	ER_ADMIN_WRONG_MRG_TABLE uint16 = 1472
	ER_TOO_HIGH_LEVEL_OF_NESTING_FOR_SELECT uint16 = 1473
	ER_NAME_BECOMES_EMPTY uint16 = 1474
	ER_AMBIGUOUS_FIELD_TERM uint16 = 1475
	ER_FOREIGN_SERVER_EXISTS uint16 = 1476
	ER_FOREIGN_SERVER_DOESNT_EXIST uint16 = 1477
	ER_ILLEGAL_HA_CREATE_OPTION uint16 = 1478
	ER_PARTITION_REQUIRES_VALUES_ERROR uint16 = 1479
	ER_PARTITION_WRONG_VALUES_ERROR uint16 = 1480
	ER_PARTITION_MAXVALUE_ERROR uint16 = 1481
	//OBSOLETE_ER_PARTITION_SUBPARTITION_ERROR uint16 = 1482
	//OBSOLETE_ER_PARTITION_SUBPART_MIX_ERROR uint16 = 1483
	ER_PARTITION_WRONG_NO_PART_ERROR uint16 = 1484
	ER_PARTITION_WRONG_NO_SUBPART_ERROR uint16 = 1485
	ER_WRONG_EXPR_IN_PARTITION_FUNC_ERROR uint16 = 1486
	//OBSOLETE_ER_NO_CONST_EXPR_IN_RANGE_OR_LIST_ERROR uint16 = 1487
	ER_FIELD_NOT_FOUND_PART_ERROR uint16 = 1488
	//OBSOLETE_ER_LIST_OF_FIELDS_ONLY_IN_HASH_ERROR uint16 = 1489
	ER_INCONSISTENT_PARTITION_INFO_ERROR uint16 = 1490
	ER_PARTITION_FUNC_NOT_ALLOWED_ERROR uint16 = 1491
	ER_PARTITIONS_MUST_BE_DEFINED_ERROR uint16 = 1492
	ER_RANGE_NOT_INCREASING_ERROR uint16 = 1493
	ER_INCONSISTENT_TYPE_OF_FUNCTIONS_ERROR uint16 = 1494
	ER_MULTIPLE_DEF_CONST_IN_LIST_PART_ERROR uint16 = 1495
	ER_PARTITION_ENTRY_ERROR uint16 = 1496
	ER_MIX_HANDLER_ERROR uint16 = 1497
	ER_PARTITION_NOT_DEFINED_ERROR uint16 = 1498
	ER_TOO_MANY_PARTITIONS_ERROR uint16 = 1499
	ER_SUBPARTITION_ERROR uint16 = 1500
	ER_CANT_CREATE_HANDLER_FILE uint16 = 1501
	ER_BLOB_FIELD_IN_PART_FUNC_ERROR uint16 = 1502
	ER_UNIQUE_KEY_NEED_ALL_FIELDS_IN_PF uint16 = 1503
	ER_NO_PARTS_ERROR uint16 = 1504
	ER_PARTITION_MGMT_ON_NONPARTITIONED uint16 = 1505
	ER_FOREIGN_KEY_ON_PARTITIONED uint16 = 1506
	ER_DROP_PARTITION_NON_EXISTENT uint16 = 1507
	ER_DROP_LAST_PARTITION uint16 = 1508
	ER_COALESCE_ONLY_ON_HASH_PARTITION uint16 = 1509
	ER_REORG_HASH_ONLY_ON_SAME_NO uint16 = 1510
	ER_REORG_NO_PARAM_ERROR uint16 = 1511
	ER_ONLY_ON_RANGE_LIST_PARTITION uint16 = 1512
	ER_ADD_PARTITION_SUBPART_ERROR uint16 = 1513
	ER_ADD_PARTITION_NO_NEW_PARTITION uint16 = 1514
	ER_COALESCE_PARTITION_NO_PARTITION uint16 = 1515
	ER_REORG_PARTITION_NOT_EXIST uint16 = 1516
	ER_SAME_NAME_PARTITION uint16 = 1517
	ER_NO_BINLOG_ERROR uint16 = 1518
	ER_CONSECUTIVE_REORG_PARTITIONS uint16 = 1519
	ER_REORG_OUTSIDE_RANGE uint16 = 1520
	ER_PARTITION_FUNCTION_FAILURE uint16 = 1521
	//OBSOLETE_ER_PART_STATE_ERROR uint16 = 1522
	ER_LIMITED_PART_RANGE uint16 = 1523
	ER_PLUGIN_IS_NOT_LOADED uint16 = 1524
	ER_WRONG_VALUE uint16 = 1525
	ER_NO_PARTITION_FOR_GIVEN_VALUE uint16 = 1526
	ER_FILEGROUP_OPTION_ONLY_ONCE uint16 = 1527
	ER_CREATE_FILEGROUP_FAILED uint16 = 1528
	ER_DROP_FILEGROUP_FAILED uint16 = 1529
	ER_TABLESPACE_AUTO_EXTEND_ERROR uint16 = 1530
	ER_WRONG_SIZE_NUMBER uint16 = 1531
	ER_SIZE_OVERFLOW_ERROR uint16 = 1532
	ER_ALTER_FILEGROUP_FAILED uint16 = 1533
	ER_BINLOG_ROW_LOGGING_FAILED uint16 = 1534
	//OBSOLETE_ER_BINLOG_ROW_WRONG_TABLE_DEF uint16 = 1535
	//OBSOLETE_ER_BINLOG_ROW_RBR_TO_SBR uint16 = 1536
	ER_EVENT_ALREADY_EXISTS uint16 = 1537
	//OBSOLETE_ER_EVENT_STORE_FAILED uint16 = 1538
	ER_EVENT_DOES_NOT_EXIST uint16 = 1539
	//OBSOLETE_ER_EVENT_CANT_ALTER uint16 = 1540
	//OBSOLETE_ER_EVENT_DROP_FAILED uint16 = 1541
	ER_EVENT_INTERVAL_NOT_POSITIVE_OR_TOO_BIG uint16 = 1542
	ER_EVENT_ENDS_BEFORE_STARTS uint16 = 1543
	ER_EVENT_EXEC_TIME_IN_THE_PAST uint16 = 1544
	//OBSOLETE_ER_EVENT_OPEN_TABLE_FAILED uint16 = 1545
	//OBSOLETE_ER_EVENT_NEITHER_M_EXPR_NOR_M_AT uint16 = 1546
	//OBSOLETE_ER_COL_COUNT_DOESNT_MATCH_CORRUPTED uint16 = 1547
	//OBSOLETE_ER_CANNOT_LOAD_FROM_TABLE uint16 = 1548
	//OBSOLETE_ER_EVENT_CANNOT_DELETE uint16 = 1549
	//OBSOLETE_ER_EVENT_COMPILE_ERROR uint16 = 1550
	ER_EVENT_SAME_NAME uint16 = 1551
	//OBSOLETE_ER_EVENT_DATA_TOO_LONG uint16 = 1552
	ER_DROP_INDEX_FK uint16 = 1553
	ER_WARN_DEPRECATED_SYNTAX_WITH_VER uint16 = 1554
	//OBSOLETE_ER_CANT_WRITE_LOCK_LOG_TABLE uint16 = 1555
	ER_CANT_LOCK_LOG_TABLE uint16 = 1556
	ER_FOREIGN_DUPLICATE_KEY_OLD_UNUSED uint16 = 1557
	ER_COL_COUNT_DOESNT_MATCH_PLEASE_UPDATE uint16 = 1558
	//OBSOLETE_ER_TEMP_TABLE_PREVENTS_SWITCH_OUT_OF_RBR uint16 = 1559
	ER_STORED_FUNCTION_PREVENTS_SWITCH_BINLOG_FORMAT uint16 = 1560
	//OBSOLETE_ER_NDB_CANT_SWITCH_BINLOG_FORMAT uint16 = 1561
	ER_PARTITION_NO_TEMPORARY uint16 = 1562
	ER_PARTITION_CONST_DOMAIN_ERROR uint16 = 1563
	ER_PARTITION_FUNCTION_IS_NOT_ALLOWED uint16 = 1564
	//OBSOLETE_ER_DDL_LOG_ERROR_UNUSED uint16 = 1565
	ER_NULL_IN_VALUES_LESS_THAN uint16 = 1566
	ER_WRONG_PARTITION_NAME uint16 = 1567
	ER_CANT_CHANGE_TX_CHARACTERISTICS uint16 = 1568
	ER_DUP_ENTRY_AUTOINCREMENT_CASE uint16 = 1569
	//OBSOLETE_ER_EVENT_MODIFY_QUEUE_ERROR uint16 = 1570
	ER_EVENT_SET_VAR_ERROR uint16 = 1571
	ER_PARTITION_MERGE_ERROR uint16 = 1572
	//OBSOLETE_ER_CANT_ACTIVATE_LOG uint16 = 1573
	//OBSOLETE_ER_RBR_NOT_AVAILABLE uint16 = 1574
	ER_BASE64_DECODE_ERROR uint16 = 1575
	ER_EVENT_RECURSION_FORBIDDEN uint16 = 1576
	//OBSOLETE_ER_EVENTS_DB_ERROR uint16 = 1577
	ER_ONLY_INTEGERS_ALLOWED uint16 = 1578
	ER_UNSUPORTED_LOG_ENGINE uint16 = 1579
	ER_BAD_LOG_STATEMENT uint16 = 1580
	ER_CANT_RENAME_LOG_TABLE uint16 = 1581
	ER_WRONG_PARAMCOUNT_TO_NATIVE_FCT uint16 = 1582
	ER_WRONG_PARAMETERS_TO_NATIVE_FCT uint16 = 1583
	ER_WRONG_PARAMETERS_TO_STORED_FCT uint16 = 1584
	ER_NATIVE_FCT_NAME_COLLISION uint16 = 1585
	ER_DUP_ENTRY_WITH_KEY_NAME uint16 = 1586
	ER_BINLOG_PURGE_EMFILE uint16 = 1587
	ER_EVENT_CANNOT_CREATE_IN_THE_PAST uint16 = 1588
	ER_EVENT_CANNOT_ALTER_IN_THE_PAST uint16 = 1589
	//OBSOLETE_ER_SLAVE_INCIDENT uint16 = 1590
	ER_NO_PARTITION_FOR_GIVEN_VALUE_SILENT uint16 = 1591
	ER_BINLOG_UNSAFE_STATEMENT uint16 = 1592
	ER_BINLOG_FATAL_ERROR uint16 = 1593
	//OBSOLETE_ER_SLAVE_RELAY_LOG_READ_FAILURE uint16 = 1594
	//OBSOLETE_ER_SLAVE_RELAY_LOG_WRITE_FAILURE uint16 = 1595
	//OBSOLETE_ER_SLAVE_CREATE_EVENT_FAILURE uint16 = 1596
	//OBSOLETE_ER_SLAVE_MASTER_COM_FAILURE uint16 = 1597
	ER_BINLOG_LOGGING_IMPOSSIBLE uint16 = 1598
	ER_VIEW_NO_CREATION_CTX uint16 = 1599
	ER_VIEW_INVALID_CREATION_CTX uint16 = 1600
	//OBSOLETE_ER_SR_INVALID_CREATION_CTX uint16 = 1601
	ER_TRG_CORRUPTED_FILE uint16 = 1602
	ER_TRG_NO_CREATION_CTX uint16 = 1603
	ER_TRG_INVALID_CREATION_CTX uint16 = 1604
	ER_EVENT_INVALID_CREATION_CTX uint16 = 1605
	ER_TRG_CANT_OPEN_TABLE uint16 = 1606
	//OBSOLETE_ER_CANT_CREATE_SROUTINE uint16 = 1607
	//OBSOLETE_ER_NEVER_USED uint16 = 1608
	ER_NO_FORMAT_DESCRIPTION_EVENT_BEFORE_BINLOG_STATEMENT uint16 = 1609
	ER_SLAVE_CORRUPT_EVENT uint16 = 1610
	//OBSOLETE_ER_LOAD_DATA_INVALID_COLUMN_UNUSED uint16 = 1611
	ER_LOG_PURGE_NO_FILE uint16 = 1612
	ER_XA_RBTIMEOUT uint16 = 1613
	ER_XA_RBDEADLOCK uint16 = 1614
	ER_NEED_REPREPARE uint16 = 1615
	//OBSOLETE_ER_DELAYED_NOT_SUPPORTED uint16 = 1616
	WARN_NO_MASTER_INFO uint16 = 1617
	WARN_OPTION_IGNORED uint16 = 1618
	ER_PLUGIN_DELETE_BUILTIN uint16 = 1619
	WARN_PLUGIN_BUSY uint16 = 1620
	ER_VARIABLE_IS_READONLY uint16 = 1621
	ER_WARN_ENGINE_TRANSACTION_ROLLBACK uint16 = 1622
	//OBSOLETE_ER_SLAVE_HEARTBEAT_FAILURE uint16 = 1623
	ER_SLAVE_HEARTBEAT_VALUE_OUT_OF_RANGE uint16 = 1624
	ER_NDB_REPLICATION_SCHEMA_ERROR uint16 = 1625
	ER_CONFLICT_FN_PARSE_ERROR uint16 = 1626
	ER_EXCEPTIONS_WRITE_ERROR uint16 = 1627
	ER_TOO_LONG_TABLE_COMMENT uint16 = 1628
	ER_TOO_LONG_FIELD_COMMENT uint16 = 1629
	ER_FUNC_INEXISTENT_NAME_COLLISION uint16 = 1630
	ER_DATABASE_NAME uint16 = 1631
	ER_TABLE_NAME uint16 = 1632
	ER_PARTITION_NAME uint16 = 1633
	ER_SUBPARTITION_NAME uint16 = 1634
	ER_TEMPORARY_NAME uint16 = 1635
	ER_RENAMED_NAME uint16 = 1636
	ER_TOO_MANY_CONCURRENT_TRXS uint16 = 1637
	WARN_NON_ASCII_SEPARATOR_NOT_IMPLEMENTED uint16 = 1638
	ER_DEBUG_SYNC_TIMEOUT uint16 = 1639
	ER_DEBUG_SYNC_HIT_LIMIT uint16 = 1640
	ER_DUP_SIGNAL_SET uint16 = 1641
	ER_SIGNAL_WARN uint16 = 1642
	ER_SIGNAL_NOT_FOUND uint16 = 1643
	ER_SIGNAL_EXCEPTION uint16 = 1644
	ER_RESIGNAL_WITHOUT_ACTIVE_HANDLER uint16 = 1645
	ER_SIGNAL_BAD_CONDITION_TYPE uint16 = 1646
	WARN_COND_ITEM_TRUNCATED uint16 = 1647
	ER_COND_ITEM_TOO_LONG uint16 = 1648
	ER_UNKNOWN_LOCALE uint16 = 1649
	ER_SLAVE_IGNORE_SERVER_IDS uint16 = 1650
	//OBSOLETE_ER_QUERY_CACHE_DISABLED uint16 = 1651
	ER_SAME_NAME_PARTITION_FIELD uint16 = 1652
	ER_PARTITION_COLUMN_LIST_ERROR uint16 = 1653
	ER_WRONG_TYPE_COLUMN_VALUE_ERROR uint16 = 1654
	ER_TOO_MANY_PARTITION_FUNC_FIELDS_ERROR uint16 = 1655
	ER_MAXVALUE_IN_VALUES_IN uint16 = 1656
	ER_TOO_MANY_VALUES_ERROR uint16 = 1657
	ER_ROW_SINGLE_PARTITION_FIELD_ERROR uint16 = 1658
	ER_FIELD_TYPE_NOT_ALLOWED_AS_PARTITION_FIELD uint16 = 1659
	ER_PARTITION_FIELDS_TOO_LONG uint16 = 1660
	ER_BINLOG_ROW_ENGINE_AND_STMT_ENGINE uint16 = 1661
	ER_BINLOG_ROW_MODE_AND_STMT_ENGINE uint16 = 1662
	ER_BINLOG_UNSAFE_AND_STMT_ENGINE uint16 = 1663
	ER_BINLOG_ROW_INJECTION_AND_STMT_ENGINE uint16 = 1664
	ER_BINLOG_STMT_MODE_AND_ROW_ENGINE uint16 = 1665
	ER_BINLOG_ROW_INJECTION_AND_STMT_MODE uint16 = 1666
	ER_BINLOG_MULTIPLE_ENGINES_AND_SELF_LOGGING_ENGINE uint16 = 1667
	ER_BINLOG_UNSAFE_LIMIT uint16 = 1668
	//OBSOLETE_ER_UNUSED4 uint16 = 1669
	ER_BINLOG_UNSAFE_SYSTEM_TABLE uint16 = 1670
	ER_BINLOG_UNSAFE_AUTOINC_COLUMNS uint16 = 1671
	ER_BINLOG_UNSAFE_UDF uint16 = 1672
	ER_BINLOG_UNSAFE_SYSTEM_VARIABLE uint16 = 1673
	ER_BINLOG_UNSAFE_SYSTEM_FUNCTION uint16 = 1674
	ER_BINLOG_UNSAFE_NONTRANS_AFTER_TRANS uint16 = 1675
	ER_MESSAGE_AND_STATEMENT uint16 = 1676
	//OBSOLETE_ER_SLAVE_CONVERSION_FAILED uint16 = 1677
	ER_SLAVE_CANT_CREATE_CONVERSION uint16 = 1678
	ER_INSIDE_TRANSACTION_PREVENTS_SWITCH_BINLOG_FORMAT uint16 = 1679
	ER_PATH_LENGTH uint16 = 1680
	ER_WARN_DEPRECATED_SYNTAX_NO_REPLACEMENT uint16 = 1681
	ER_WRONG_NATIVE_TABLE_STRUCTURE uint16 = 1682
	ER_WRONG_PERFSCHEMA_USAGE uint16 = 1683
	ER_WARN_I_S_SKIPPED_TABLE uint16 = 1684
	ER_INSIDE_TRANSACTION_PREVENTS_SWITCH_BINLOG_DIRECT uint16 = 1685
	ER_STORED_FUNCTION_PREVENTS_SWITCH_BINLOG_DIRECT uint16 = 1686
	ER_SPATIAL_MUST_HAVE_GEOM_COL uint16 = 1687
	ER_TOO_LONG_INDEX_COMMENT uint16 = 1688
	ER_LOCK_ABORTED uint16 = 1689
	ER_DATA_OUT_OF_RANGE uint16 = 1690
	//OBSOLETE_ER_WRONG_SPVAR_TYPE_IN_LIMIT uint16 = 1691
	ER_BINLOG_UNSAFE_MULTIPLE_ENGINES_AND_SELF_LOGGING_ENGINE uint16 = 1692
	ER_BINLOG_UNSAFE_MIXED_STATEMENT uint16 = 1693
	ER_INSIDE_TRANSACTION_PREVENTS_SWITCH_SQL_LOG_BIN uint16 = 1694
	ER_STORED_FUNCTION_PREVENTS_SWITCH_SQL_LOG_BIN uint16 = 1695
	ER_FAILED_READ_FROM_PAR_FILE uint16 = 1696
	ER_VALUES_IS_NOT_INT_TYPE_ERROR uint16 = 1697
	ER_ACCESS_DENIED_NO_PASSWORD_ERROR uint16 = 1698
	ER_SET_PASSWORD_AUTH_PLUGIN uint16 = 1699
	//OBSOLETE_ER_GRANT_PLUGIN_USER_EXISTS uint16 = 1700
	ER_TRUNCATE_ILLEGAL_FK uint16 = 1701
	ER_PLUGIN_IS_PERMANENT uint16 = 1702
	ER_SLAVE_HEARTBEAT_VALUE_OUT_OF_RANGE_MIN uint16 = 1703
	ER_SLAVE_HEARTBEAT_VALUE_OUT_OF_RANGE_MAX uint16 = 1704
	ER_STMT_CACHE_FULL uint16 = 1705
	ER_MULTI_UPDATE_KEY_CONFLICT uint16 = 1706
	ER_TABLE_NEEDS_REBUILD uint16 = 1707
	WARN_OPTION_BELOW_LIMIT uint16 = 1708
	ER_INDEX_COLUMN_TOO_LONG uint16 = 1709
	ER_ERROR_IN_TRIGGER_BODY uint16 = 1710
	ER_ERROR_IN_UNKNOWN_TRIGGER_BODY uint16 = 1711
	ER_INDEX_CORRUPT uint16 = 1712
	ER_UNDO_RECORD_TOO_BIG uint16 = 1713
	ER_BINLOG_UNSAFE_INSERT_IGNORE_SELECT uint16 = 1714
	ER_BINLOG_UNSAFE_INSERT_SELECT_UPDATE uint16 = 1715
	ER_BINLOG_UNSAFE_REPLACE_SELECT uint16 = 1716
	ER_BINLOG_UNSAFE_CREATE_IGNORE_SELECT uint16 = 1717
	ER_BINLOG_UNSAFE_CREATE_REPLACE_SELECT uint16 = 1718
	ER_BINLOG_UNSAFE_UPDATE_IGNORE uint16 = 1719
	ER_PLUGIN_NO_UNINSTALL uint16 = 1720
	ER_PLUGIN_NO_INSTALL uint16 = 1721
	ER_BINLOG_UNSAFE_WRITE_AUTOINC_SELECT uint16 = 1722
	ER_BINLOG_UNSAFE_CREATE_SELECT_AUTOINC uint16 = 1723
	ER_BINLOG_UNSAFE_INSERT_TWO_KEYS uint16 = 1724
	ER_TABLE_IN_FK_CHECK uint16 = 1725
	ER_UNSUPPORTED_ENGINE uint16 = 1726
	ER_BINLOG_UNSAFE_AUTOINC_NOT_FIRST uint16 = 1727
	ER_CANNOT_LOAD_FROM_TABLE_V2 uint16 = 1728
	ER_MASTER_DELAY_VALUE_OUT_OF_RANGE uint16 = 1729
	ER_ONLY_FD_AND_RBR_EVENTS_ALLOWED_IN_BINLOG_STATEMENT uint16 = 1730
	ER_PARTITION_EXCHANGE_DIFFERENT_OPTION uint16 = 1731
	ER_PARTITION_EXCHANGE_PART_TABLE uint16 = 1732
	ER_PARTITION_EXCHANGE_TEMP_TABLE uint16 = 1733
	ER_PARTITION_INSTEAD_OF_SUBPARTITION uint16 = 1734
	ER_UNKNOWN_PARTITION uint16 = 1735
	ER_TABLES_DIFFERENT_METADATA uint16 = 1736
	ER_ROW_DOES_NOT_MATCH_PARTITION uint16 = 1737
	ER_BINLOG_CACHE_SIZE_GREATER_THAN_MAX uint16 = 1738
	ER_WARN_INDEX_NOT_APPLICABLE uint16 = 1739
	ER_PARTITION_EXCHANGE_FOREIGN_KEY uint16 = 1740
	//OBSOLETE_ER_NO_SUCH_KEY_VALUE uint16 = 1741
	ER_RPL_INFO_DATA_TOO_LONG uint16 = 1742
	//OBSOLETE_ER_NETWORK_READ_EVENT_CHECKSUM_FAILURE uint16 = 1743
	//OBSOLETE_ER_BINLOG_READ_EVENT_CHECKSUM_FAILURE uint16 = 1744
	ER_BINLOG_STMT_CACHE_SIZE_GREATER_THAN_MAX uint16 = 1745
	ER_CANT_UPDATE_TABLE_IN_CREATE_TABLE_SELECT uint16 = 1746
	ER_PARTITION_CLAUSE_ON_NONPARTITIONED uint16 = 1747
	ER_ROW_DOES_NOT_MATCH_GIVEN_PARTITION_SET uint16 = 1748
	//OBSOLETE_ER_NO_SUCH_PARTITION__UNUSED uint16 = 1749
	ER_CHANGE_RPL_INFO_REPOSITORY_FAILURE uint16 = 1750
	ER_WARNING_NOT_COMPLETE_ROLLBACK_WITH_CREATED_TEMP_TABLE uint16 = 1751
	ER_WARNING_NOT_COMPLETE_ROLLBACK_WITH_DROPPED_TEMP_TABLE uint16 = 1752
	ER_MTS_FEATURE_IS_NOT_SUPPORTED uint16 = 1753
	ER_MTS_UPDATED_DBS_GREATER_MAX uint16 = 1754
	ER_MTS_CANT_PARALLEL uint16 = 1755
	ER_MTS_INCONSISTENT_DATA uint16 = 1756
	ER_FULLTEXT_NOT_SUPPORTED_WITH_PARTITIONING uint16 = 1757
	ER_DA_INVALID_CONDITION_NUMBER uint16 = 1758
	ER_INSECURE_PLAIN_TEXT uint16 = 1759
	ER_INSECURE_CHANGE_MASTER uint16 = 1760
	ER_FOREIGN_DUPLICATE_KEY_WITH_CHILD_INFO uint16 = 1761
	ER_FOREIGN_DUPLICATE_KEY_WITHOUT_CHILD_INFO uint16 = 1762
	ER_SQLTHREAD_WITH_SECURE_SLAVE uint16 = 1763
	ER_TABLE_HAS_NO_FT uint16 = 1764
	ER_VARIABLE_NOT_SETTABLE_IN_SF_OR_TRIGGER uint16 = 1765
	ER_VARIABLE_NOT_SETTABLE_IN_TRANSACTION uint16 = 1766
	//OBSOLETE_ER_GTID_NEXT_IS_NOT_IN_GTID_NEXT_LIST uint16 = 1767
	//OBSOLETE_ER_CANT_CHANGE_GTID_NEXT_IN_TRANSACTION uint16 = 1768
	ER_SET_STATEMENT_CANNOT_INVOKE_FUNCTION uint16 = 1769
	ER_GTID_NEXT_CANT_BE_AUTOMATIC_IF_GTID_NEXT_LIST_IS_NON_NULL uint16 = 1770
	//OBSOLETE_ER_SKIPPING_LOGGED_TRANSACTION uint16 = 1771
	ER_MALFORMED_GTID_SET_SPECIFICATION uint16 = 1772
	ER_MALFORMED_GTID_SET_ENCODING uint16 = 1773
	ER_MALFORMED_GTID_SPECIFICATION uint16 = 1774
	ER_GNO_EXHAUSTED uint16 = 1775
	ER_BAD_SLAVE_AUTO_POSITION uint16 = 1776
	ER_AUTO_POSITION_REQUIRES_GTID_MODE_NOT_OFF uint16 = 1777
	ER_CANT_DO_IMPLICIT_COMMIT_IN_TRX_WHEN_GTID_NEXT_IS_SET uint16 = 1778
	ER_GTID_MODE_ON_REQUIRES_ENFORCE_GTID_CONSISTENCY_ON uint16 = 1779
	//OBSOLETE_ER_GTID_MODE_REQUIRES_BINLOG uint16 = 1780
	ER_CANT_SET_GTID_NEXT_TO_GTID_WHEN_GTID_MODE_IS_OFF uint16 = 1781
	ER_CANT_SET_GTID_NEXT_TO_ANONYMOUS_WHEN_GTID_MODE_IS_ON uint16 = 1782
	ER_CANT_SET_GTID_NEXT_LIST_TO_NON_NULL_WHEN_GTID_MODE_IS_OFF uint16 = 1783
	//OBSOLETE_ER_FOUND_GTID_EVENT_WHEN_GTID_MODE_IS_OFF__UNUSED uint16 = 1784
	ER_GTID_UNSAFE_NON_TRANSACTIONAL_TABLE uint16 = 1785
	ER_GTID_UNSAFE_CREATE_SELECT uint16 = 1786
	//OBSOLETE_ER_GTID_UNSAFE_CREATE_DROP_TEMP_TABLE_IN_TRANSACTION uint16 = 1787
	ER_GTID_MODE_CAN_ONLY_CHANGE_ONE_STEP_AT_A_TIME uint16 = 1788
	ER_MASTER_HAS_PURGED_REQUIRED_GTIDS uint16 = 1789
	ER_CANT_SET_GTID_NEXT_WHEN_OWNING_GTID uint16 = 1790
	ER_UNKNOWN_EXPLAIN_FORMAT uint16 = 1791
	ER_CANT_EXECUTE_IN_READ_ONLY_TRANSACTION uint16 = 1792
	ER_TOO_LONG_TABLE_PARTITION_COMMENT uint16 = 1793
	ER_SLAVE_CONFIGURATION uint16 = 1794
	ER_INNODB_FT_LIMIT uint16 = 1795
	ER_INNODB_NO_FT_TEMP_TABLE uint16 = 1796
	ER_INNODB_FT_WRONG_DOCID_COLUMN uint16 = 1797
	ER_INNODB_FT_WRONG_DOCID_INDEX uint16 = 1798
	ER_INNODB_ONLINE_LOG_TOO_BIG uint16 = 1799
	ER_UNKNOWN_ALTER_ALGORITHM uint16 = 1800
	ER_UNKNOWN_ALTER_LOCK uint16 = 1801
	ER_MTS_CHANGE_MASTER_CANT_RUN_WITH_GAPS uint16 = 1802
	ER_MTS_RECOVERY_FAILURE uint16 = 1803
	ER_MTS_RESET_WORKERS uint16 = 1804
	ER_COL_COUNT_DOESNT_MATCH_CORRUPTED_V2 uint16 = 1805
	ER_SLAVE_SILENT_RETRY_TRANSACTION uint16 = 1806
	ER_DISCARD_FK_CHECKS_RUNNING uint16 = 1807
	ER_TABLE_SCHEMA_MISMATCH uint16 = 1808
	ER_TABLE_IN_SYSTEM_TABLESPACE uint16 = 1809
	ER_IO_READ_ERROR uint16 = 1810
	ER_IO_WRITE_ERROR uint16 = 1811
	ER_TABLESPACE_MISSING uint16 = 1812
	ER_TABLESPACE_EXISTS uint16 = 1813
	ER_TABLESPACE_DISCARDED uint16 = 1814
	ER_INTERNAL_ERROR uint16 = 1815
	ER_INNODB_IMPORT_ERROR uint16 = 1816
	ER_INNODB_INDEX_CORRUPT uint16 = 1817
	ER_INVALID_YEAR_COLUMN_LENGTH uint16 = 1818
	ER_NOT_VALID_PASSWORD uint16 = 1819
	ER_MUST_CHANGE_PASSWORD uint16 = 1820
	ER_FK_NO_INDEX_CHILD uint16 = 1821
	ER_FK_NO_INDEX_PARENT uint16 = 1822
	ER_FK_FAIL_ADD_SYSTEM uint16 = 1823
	ER_FK_CANNOT_OPEN_PARENT uint16 = 1824
	ER_FK_INCORRECT_OPTION uint16 = 1825
	ER_FK_DUP_NAME uint16 = 1826
	ER_PASSWORD_FORMAT uint16 = 1827
	ER_FK_COLUMN_CANNOT_DROP uint16 = 1828
	ER_FK_COLUMN_CANNOT_DROP_CHILD uint16 = 1829
	ER_FK_COLUMN_NOT_NULL uint16 = 1830
	ER_DUP_INDEX uint16 = 1831
	ER_FK_COLUMN_CANNOT_CHANGE uint16 = 1832
	ER_FK_COLUMN_CANNOT_CHANGE_CHILD uint16 = 1833
	//OBSOLETE_ER_UNUSED5 uint16 = 1834
	ER_MALFORMED_PACKET uint16 = 1835
	ER_READ_ONLY_MODE uint16 = 1836
	ER_GTID_NEXT_TYPE_UNDEFINED_GTID uint16 = 1837
	ER_VARIABLE_NOT_SETTABLE_IN_SP uint16 = 1838
	//OBSOLETE_ER_CANT_SET_GTID_PURGED_WHEN_GTID_MODE_IS_OFF uint16 = 1839
	ER_CANT_SET_GTID_PURGED_WHEN_GTID_EXECUTED_IS_NOT_EMPTY uint16 = 1840
	ER_CANT_SET_GTID_PURGED_WHEN_OWNED_GTIDS_IS_NOT_EMPTY uint16 = 1841
	ER_GTID_PURGED_WAS_CHANGED uint16 = 1842
	ER_GTID_EXECUTED_WAS_CHANGED uint16 = 1843
	ER_BINLOG_STMT_MODE_AND_NO_REPL_TABLES uint16 = 1844
	ER_ALTER_OPERATION_NOT_SUPPORTED uint16 = 1845
	ER_ALTER_OPERATION_NOT_SUPPORTED_REASON uint16 = 1846
	ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_COPY uint16 = 1847
	ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_PARTITION uint16 = 1848
	ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FK_RENAME uint16 = 1849
	ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_COLUMN_TYPE uint16 = 1850
	ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FK_CHECK uint16 = 1851
	//OBSOLETE_ER_UNUSED6 uint16 = 1852
	ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_NOPK uint16 = 1853
	ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_AUTOINC uint16 = 1854
	ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_HIDDEN_FTS uint16 = 1855
	ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_CHANGE_FTS uint16 = 1856
	ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FTS uint16 = 1857
	//OBSOLETE_ER_SQL_SLAVE_SKIP_COUNTER_NOT_SETTABLE_IN_GTID_MODE uint16 = 1858
	ER_DUP_UNKNOWN_IN_INDEX uint16 = 1859
	ER_IDENT_CAUSES_TOO_LONG_PATH uint16 = 1860
	ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_NOT_NULL uint16 = 1861
	ER_MUST_CHANGE_PASSWORD_LOGIN uint16 = 1862
	ER_ROW_IN_WRONG_PARTITION uint16 = 1863
	ER_MTS_EVENT_BIGGER_PENDING_JOBS_SIZE_MAX uint16 = 1864
	//OBSOLETE_ER_INNODB_NO_FT_USES_PARSER uint16 = 1865
	ER_BINLOG_LOGICAL_CORRUPTION uint16 = 1866
	ER_WARN_PURGE_LOG_IN_USE uint16 = 1867
	ER_WARN_PURGE_LOG_IS_ACTIVE uint16 = 1868
	ER_AUTO_INCREMENT_CONFLICT uint16 = 1869
	WARN_ON_BLOCKHOLE_IN_RBR uint16 = 1870
	ER_SLAVE_MI_INIT_REPOSITORY uint16 = 1871
	ER_SLAVE_RLI_INIT_REPOSITORY uint16 = 1872
	ER_ACCESS_DENIED_CHANGE_USER_ERROR uint16 = 1873
	ER_INNODB_READ_ONLY uint16 = 1874
	ER_STOP_SLAVE_SQL_THREAD_TIMEOUT uint16 = 1875
	ER_STOP_SLAVE_IO_THREAD_TIMEOUT uint16 = 1876
	ER_TABLE_CORRUPT uint16 = 1877
	ER_TEMP_FILE_WRITE_FAILURE uint16 = 1878
	ER_INNODB_FT_AUX_NOT_HEX_ID uint16 = 1879
	ER_OLD_TEMPORALS_UPGRADED uint16 = 1880
	ER_INNODB_FORCED_RECOVERY uint16 = 1881
	ER_AES_INVALID_IV uint16 = 1882
	ER_PLUGIN_CANNOT_BE_UNINSTALLED uint16 = 1883
	ER_GTID_UNSAFE_BINLOG_SPLITTABLE_STATEMENT_AND_ASSIGNED_GTID uint16 = 1884
	ER_SLAVE_HAS_MORE_GTIDS_THAN_MASTER uint16 = 1885
	ER_MISSING_KEY uint16 = 1886
	WARN_NAMED_PIPE_ACCESS_EVERYONE uint16 = 1887

	//2,000 to 2,999: Client error codes reserved for use by the client library.
	// no such code here

	//3,000 to 4,999: Server error codes reserved for messages sent to clients.

	ER_FILE_CORRUPT uint16 = 3000
	ER_ERROR_ON_MASTER uint16 = 3001
	//OBSOLETE_ER_INCONSISTENT_ERROR uint16 = 3002
	ER_STORAGE_ENGINE_NOT_LOADED uint16 = 3003
	ER_GET_STACKED_DA_WITHOUT_ACTIVE_HANDLER uint16 = 3004
	ER_WARN_LEGACY_SYNTAX_CONVERTED uint16 = 3005
	ER_BINLOG_UNSAFE_FULLTEXT_PLUGIN uint16 = 3006
	ER_CANNOT_DISCARD_TEMPORARY_TABLE uint16 = 3007
	ER_FK_DEPTH_EXCEEDED uint16 = 3008
	ER_COL_COUNT_DOESNT_MATCH_PLEASE_UPDATE_V2 uint16 = 3009
	ER_WARN_TRIGGER_DOESNT_HAVE_CREATED uint16 = 3010
	ER_REFERENCED_TRG_DOES_NOT_EXIST uint16 = 3011
	ER_EXPLAIN_NOT_SUPPORTED uint16 = 3012
	ER_INVALID_FIELD_SIZE uint16 = 3013
	ER_MISSING_HA_CREATE_OPTION uint16 = 3014
	ER_ENGINE_OUT_OF_MEMORY uint16 = 3015
	ER_PASSWORD_EXPIRE_ANONYMOUS_USER uint16 = 3016
	ER_SLAVE_SQL_THREAD_MUST_STOP uint16 = 3017
	ER_NO_FT_MATERIALIZED_SUBQUERY uint16 = 3018
	ER_INNODB_UNDO_LOG_FULL uint16 = 3019
	ER_INVALID_ARGUMENT_FOR_LOGARITHM uint16 = 3020
	ER_SLAVE_CHANNEL_IO_THREAD_MUST_STOP uint16 = 3021
	ER_WARN_OPEN_TEMP_TABLES_MUST_BE_ZERO uint16 = 3022
	ER_WARN_ONLY_MASTER_LOG_FILE_NO_POS uint16 = 3023
	ER_QUERY_TIMEOUT uint16 = 3024
	ER_NON_RO_SELECT_DISABLE_TIMER uint16 = 3025
	ER_DUP_LIST_ENTRY uint16 = 3026
	//OBSOLETE_ER_SQL_MODE_NO_EFFECT uint16 = 3027
	ER_AGGREGATE_ORDER_FOR_UNION uint16 = 3028
	ER_AGGREGATE_ORDER_NON_AGG_QUERY uint16 = 3029
	ER_SLAVE_WORKER_STOPPED_PREVIOUS_THD_ERROR uint16 = 3030
	ER_DONT_SUPPORT_SLAVE_PRESERVE_COMMIT_ORDER uint16 = 3031
	ER_SERVER_OFFLINE_MODE uint16 = 3032
	ER_GIS_DIFFERENT_SRIDS uint16 = 3033
	ER_GIS_UNSUPPORTED_ARGUMENT uint16 = 3034
	ER_GIS_UNKNOWN_ERROR uint16 = 3035
	ER_GIS_UNKNOWN_EXCEPTION uint16 = 3036
	ER_GIS_INVALID_DATA uint16 = 3037
	ER_BOOST_GEOMETRY_EMPTY_INPUT_EXCEPTION uint16 = 3038
	ER_BOOST_GEOMETRY_CENTROID_EXCEPTION uint16 = 3039
	ER_BOOST_GEOMETRY_OVERLAY_INVALID_INPUT_EXCEPTION uint16 = 3040
	ER_BOOST_GEOMETRY_TURN_INFO_EXCEPTION uint16 = 3041
	ER_BOOST_GEOMETRY_SELF_INTERSECTION_POINT_EXCEPTION uint16 = 3042
	ER_BOOST_GEOMETRY_UNKNOWN_EXCEPTION uint16 = 3043
	ER_STD_BAD_ALLOC_ERROR uint16 = 3044
	ER_STD_DOMAIN_ERROR uint16 = 3045
	ER_STD_LENGTH_ERROR uint16 = 3046
	ER_STD_INVALID_ARGUMENT uint16 = 3047
	ER_STD_OUT_OF_RANGE_ERROR uint16 = 3048
	ER_STD_OVERFLOW_ERROR uint16 = 3049
	ER_STD_RANGE_ERROR uint16 = 3050
	ER_STD_UNDERFLOW_ERROR uint16 = 3051
	ER_STD_LOGIC_ERROR uint16 = 3052
	ER_STD_RUNTIME_ERROR uint16 = 3053
	ER_STD_UNKNOWN_EXCEPTION uint16 = 3054
	ER_GIS_DATA_WRONG_ENDIANESS uint16 = 3055
	ER_CHANGE_MASTER_PASSWORD_LENGTH uint16 = 3056
	ER_USER_LOCK_WRONG_NAME uint16 = 3057
	ER_USER_LOCK_DEADLOCK uint16 = 3058
	ER_REPLACE_INACCESSIBLE_ROWS uint16 = 3059
	ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_GIS uint16 = 3060
	ER_ILLEGAL_USER_VAR uint16 = 3061
	ER_GTID_MODE_OFF uint16 = 3062
	//OBSOLETE_ER_UNSUPPORTED_BY_REPLICATION_THREAD uint16 = 3063
	ER_INCORRECT_TYPE uint16 = 3064
	ER_FIELD_IN_ORDER_NOT_SELECT uint16 = 3065
	ER_AGGREGATE_IN_ORDER_NOT_SELECT uint16 = 3066
	ER_INVALID_RPL_WILD_TABLE_FILTER_PATTERN uint16 = 3067
	ER_NET_OK_PACKET_TOO_LARGE uint16 = 3068
	ER_INVALID_JSON_DATA uint16 = 3069
	ER_INVALID_GEOJSON_MISSING_MEMBER uint16 = 3070
	ER_INVALID_GEOJSON_WRONG_TYPE uint16 = 3071
	ER_INVALID_GEOJSON_UNSPECIFIED uint16 = 3072
	ER_DIMENSION_UNSUPPORTED uint16 = 3073
	ER_SLAVE_CHANNEL_DOES_NOT_EXIST uint16 = 3074
	//OBSOLETE_ER_SLAVE_MULTIPLE_CHANNELS_HOST_PORT uint16 = 3075
	ER_SLAVE_CHANNEL_NAME_INVALID_OR_TOO_LONG uint16 = 3076
	ER_SLAVE_NEW_CHANNEL_WRONG_REPOSITORY uint16 = 3077
	//OBSOLETE_ER_SLAVE_CHANNEL_DELETE uint16 = 3078
	ER_SLAVE_MULTIPLE_CHANNELS_CMD uint16 = 3079
	ER_SLAVE_MAX_CHANNELS_EXCEEDED uint16 = 3080
	ER_SLAVE_CHANNEL_MUST_STOP uint16 = 3081
	ER_SLAVE_CHANNEL_NOT_RUNNING uint16 = 3082
	ER_SLAVE_CHANNEL_WAS_RUNNING uint16 = 3083
	ER_SLAVE_CHANNEL_WAS_NOT_RUNNING uint16 = 3084
	ER_SLAVE_CHANNEL_SQL_THREAD_MUST_STOP uint16 = 3085
	ER_SLAVE_CHANNEL_SQL_SKIP_COUNTER uint16 = 3086
	ER_WRONG_FIELD_WITH_GROUP_V2 uint16 = 3087
	ER_MIX_OF_GROUP_FUNC_AND_FIELDS_V2 uint16 = 3088
	ER_WARN_DEPRECATED_SYSVAR_UPDATE uint16 = 3089
	ER_WARN_DEPRECATED_SQLMODE uint16 = 3090
	ER_CANNOT_LOG_PARTIAL_DROP_DATABASE_WITH_GTID uint16 = 3091
	ER_GROUP_REPLICATION_CONFIGURATION uint16 = 3092
	ER_GROUP_REPLICATION_RUNNING uint16 = 3093
	ER_GROUP_REPLICATION_APPLIER_INIT_ERROR uint16 = 3094
	ER_GROUP_REPLICATION_STOP_APPLIER_THREAD_TIMEOUT uint16 = 3095
	ER_GROUP_REPLICATION_COMMUNICATION_LAYER_SESSION_ERROR uint16 = 3096
	ER_GROUP_REPLICATION_COMMUNICATION_LAYER_JOIN_ERROR uint16 = 3097
	ER_BEFORE_DML_VALIDATION_ERROR uint16 = 3098
	ER_PREVENTS_VARIABLE_WITHOUT_RBR uint16 = 3099
	ER_RUN_HOOK_ERROR uint16 = 3100
	ER_TRANSACTION_ROLLBACK_DURING_COMMIT uint16 = 3101
	ER_GENERATED_COLUMN_FUNCTION_IS_NOT_ALLOWED uint16 = 3102
	ER_UNSUPPORTED_ALTER_INPLACE_ON_VIRTUAL_COLUMN uint16 = 3103
	ER_WRONG_FK_OPTION_FOR_GENERATED_COLUMN uint16 = 3104
	ER_NON_DEFAULT_VALUE_FOR_GENERATED_COLUMN uint16 = 3105
	ER_UNSUPPORTED_ACTION_ON_GENERATED_COLUMN uint16 = 3106
	ER_GENERATED_COLUMN_NON_PRIOR uint16 = 3107
	ER_DEPENDENT_BY_GENERATED_COLUMN uint16 = 3108
	ER_GENERATED_COLUMN_REF_AUTO_INC uint16 = 3109
	ER_FEATURE_NOT_AVAILABLE uint16 = 3110
	ER_CANT_SET_GTID_MODE uint16 = 3111
	ER_CANT_USE_AUTO_POSITION_WITH_GTID_MODE_OFF uint16 = 3112
	//OBSOLETE_ER_CANT_REPLICATE_ANONYMOUS_WITH_AUTO_POSITION uint16 = 3113
	//OBSOLETE_ER_CANT_REPLICATE_ANONYMOUS_WITH_GTID_MODE_ON uint16 = 3114
	//OBSOLETE_ER_CANT_REPLICATE_GTID_WITH_GTID_MODE_OFF uint16 = 3115
	ER_CANT_ENFORCE_GTID_CONSISTENCY_WITH_ONGOING_GTID_VIOLATING_TX uint16 = 3116
	ER_ENFORCE_GTID_CONSISTENCY_WARN_WITH_ONGOING_GTID_VIOLATING_TX uint16 = 3117
	ER_ACCOUNT_HAS_BEEN_LOCKED uint16 = 3118
	ER_WRONG_TABLESPACE_NAME uint16 = 3119
	ER_TABLESPACE_IS_NOT_EMPTY uint16 = 3120
	ER_WRONG_FILE_NAME uint16 = 3121
	ER_BOOST_GEOMETRY_INCONSISTENT_TURNS_EXCEPTION uint16 = 3122
	ER_WARN_OPTIMIZER_HINT_SYNTAX_ERROR uint16 = 3123
	ER_WARN_BAD_MAX_EXECUTION_TIME uint16 = 3124
	ER_WARN_UNSUPPORTED_MAX_EXECUTION_TIME uint16 = 3125
	ER_WARN_CONFLICTING_HINT uint16 = 3126
	ER_WARN_UNKNOWN_QB_NAME uint16 = 3127
	ER_UNRESOLVED_HINT_NAME uint16 = 3128
	ER_WARN_ON_MODIFYING_GTID_EXECUTED_TABLE uint16 = 3129
	ER_PLUGGABLE_PROTOCOL_COMMAND_NOT_SUPPORTED uint16 = 3130
	ER_LOCKING_SERVICE_WRONG_NAME uint16 = 3131
	ER_LOCKING_SERVICE_DEADLOCK uint16 = 3132
	ER_LOCKING_SERVICE_TIMEOUT uint16 = 3133
	ER_GIS_MAX_POINTS_IN_GEOMETRY_OVERFLOWED uint16 = 3134
	ER_SQL_MODE_MERGED uint16 = 3135
	ER_VTOKEN_PLUGIN_TOKEN_MISMATCH uint16 = 3136
	ER_VTOKEN_PLUGIN_TOKEN_NOT_FOUND uint16 = 3137
	ER_CANT_SET_VARIABLE_WHEN_OWNING_GTID uint16 = 3138
	ER_SLAVE_CHANNEL_OPERATION_NOT_ALLOWED uint16 = 3139
	ER_INVALID_JSON_TEXT uint16 = 3140
	ER_INVALID_JSON_TEXT_IN_PARAM uint16 = 3141
	ER_INVALID_JSON_BINARY_DATA uint16 = 3142
	ER_INVALID_JSON_PATH uint16 = 3143
	ER_INVALID_JSON_CHARSET uint16 = 3144
	ER_INVALID_JSON_CHARSET_IN_FUNCTION uint16 = 3145
	ER_INVALID_TYPE_FOR_JSON uint16 = 3146
	ER_INVALID_CAST_TO_JSON uint16 = 3147
	ER_INVALID_JSON_PATH_CHARSET uint16 = 3148
	ER_INVALID_JSON_PATH_WILDCARD uint16 = 3149
	ER_JSON_VALUE_TOO_BIG uint16 = 3150
	ER_JSON_KEY_TOO_BIG uint16 = 3151
	ER_JSON_USED_AS_KEY uint16 = 3152
	ER_JSON_VACUOUS_PATH uint16 = 3153
	ER_JSON_BAD_ONE_OR_ALL_ARG uint16 = 3154
	ER_NUMERIC_JSON_VALUE_OUT_OF_RANGE uint16 = 3155
	ER_INVALID_JSON_VALUE_FOR_CAST uint16 = 3156
	ER_JSON_DOCUMENT_TOO_DEEP uint16 = 3157
	ER_JSON_DOCUMENT_NULL_KEY uint16 = 3158
	ER_SECURE_TRANSPORT_REQUIRED uint16 = 3159
	ER_NO_SECURE_TRANSPORTS_CONFIGURED uint16 = 3160
	ER_DISABLED_STORAGE_ENGINE uint16 = 3161
	ER_USER_DOES_NOT_EXIST uint16 = 3162
	ER_USER_ALREADY_EXISTS uint16 = 3163
	ER_AUDIT_API_ABORT uint16 = 3164
	ER_INVALID_JSON_PATH_ARRAY_CELL uint16 = 3165
	ER_BUFPOOL_RESIZE_INPROGRESS uint16 = 3166
	ER_FEATURE_DISABLED_SEE_DOC uint16 = 3167
	ER_SERVER_ISNT_AVAILABLE uint16 = 3168
	ER_SESSION_WAS_KILLED uint16 = 3169
	ER_CAPACITY_EXCEEDED uint16 = 3170
	ER_CAPACITY_EXCEEDED_IN_RANGE_OPTIMIZER uint16 = 3171
	//OBSOLETE_ER_TABLE_NEEDS_UPG_PART uint16 = 3172
	ER_CANT_WAIT_FOR_EXECUTED_GTID_SET_WHILE_OWNING_A_GTID uint16 = 3173
	ER_CANNOT_ADD_FOREIGN_BASE_COL_VIRTUAL uint16 = 3174
	ER_CANNOT_CREATE_VIRTUAL_INDEX_CONSTRAINT uint16 = 3175
	ER_ERROR_ON_MODIFYING_GTID_EXECUTED_TABLE uint16 = 3176
	ER_LOCK_REFUSED_BY_ENGINE uint16 = 3177
	ER_UNSUPPORTED_ALTER_ONLINE_ON_VIRTUAL_COLUMN uint16 = 3178
	ER_MASTER_KEY_ROTATION_NOT_SUPPORTED_BY_SE uint16 = 3179
	//OBSOLETE_ER_MASTER_KEY_ROTATION_ERROR_BY_SE uint16 = 3180
	ER_MASTER_KEY_ROTATION_BINLOG_FAILED uint16 = 3181
	ER_MASTER_KEY_ROTATION_SE_UNAVAILABLE uint16 = 3182
	ER_TABLESPACE_CANNOT_ENCRYPT uint16 = 3183
	ER_INVALID_ENCRYPTION_OPTION uint16 = 3184
	ER_CANNOT_FIND_KEY_IN_KEYRING uint16 = 3185
	ER_CAPACITY_EXCEEDED_IN_PARSER uint16 = 3186
	ER_UNSUPPORTED_ALTER_ENCRYPTION_INPLACE uint16 = 3187
	ER_KEYRING_UDF_KEYRING_SERVICE_ERROR uint16 = 3188
	ER_USER_COLUMN_OLD_LENGTH uint16 = 3189
	ER_CANT_RESET_MASTER uint16 = 3190
	ER_GROUP_REPLICATION_MAX_GROUP_SIZE uint16 = 3191
	ER_CANNOT_ADD_FOREIGN_BASE_COL_STORED uint16 = 3192
	ER_TABLE_REFERENCED uint16 = 3193
	//OBSOLETE_ER_PARTITION_ENGINE_DEPRECATED_FOR_TABLE uint16 = 3194
	//OBSOLETE_ER_WARN_USING_GEOMFROMWKB_TO_SET_SRID_ZERO uint16 = 3195
	//OBSOLETE_ER_WARN_USING_GEOMFROMWKB_TO_SET_SRID uint16 = 3196
	ER_XA_RETRY uint16 = 3197
	ER_KEYRING_AWS_UDF_AWS_KMS_ERROR uint16 = 3198
	ER_BINLOG_UNSAFE_XA uint16 = 3199
	ER_UDF_ERROR uint16 = 3200
	ER_KEYRING_MIGRATION_FAILURE uint16 = 3201
	ER_KEYRING_ACCESS_DENIED_ERROR uint16 = 3202
	ER_KEYRING_MIGRATION_STATUS uint16 = 3203
	//OBSOLETE_ER_PLUGIN_FAILED_TO_OPEN_TABLES uint16 = 3204
	//OBSOLETE_ER_PLUGIN_FAILED_TO_OPEN_TABLE uint16 = 3205
	//OBSOLETE_ER_AUDIT_LOG_NO_KEYRING_PLUGIN_INSTALLED uint16 = 3206
	//OBSOLETE_ER_AUDIT_LOG_ENCRYPTION_PASSWORD_HAS_NOT_BEEN_SET uint16 = 3207
	//OBSOLETE_ER_AUDIT_LOG_COULD_NOT_CREATE_AES_KEY uint16 = 3208
	//OBSOLETE_ER_AUDIT_LOG_ENCRYPTION_PASSWORD_CANNOT_BE_FETCHED uint16 = 3209
	//OBSOLETE_ER_AUDIT_LOG_JSON_FILTERING_NOT_ENABLED uint16 = 3210
	//OBSOLETE_ER_AUDIT_LOG_UDF_INSUFFICIENT_PRIVILEGE uint16 = 3211
	//OBSOLETE_ER_AUDIT_LOG_SUPER_PRIVILEGE_REQUIRED uint16 = 3212
	//OBSOLETE_ER_COULD_NOT_REINITIALIZE_AUDIT_LOG_FILTERS uint16 = 3213
	//OBSOLETE_ER_AUDIT_LOG_UDF_INVALID_ARGUMENT_TYPE uint16 = 3214
	//OBSOLETE_ER_AUDIT_LOG_UDF_INVALID_ARGUMENT_COUNT uint16 = 3215
	//OBSOLETE_ER_AUDIT_LOG_HAS_NOT_BEEN_INSTALLED uint16 = 3216
	//OBSOLETE_ER_AUDIT_LOG_UDF_READ_INVALID_MAX_ARRAY_LENGTH_ARG_TYPE uint16 = 3217
	ER_AUDIT_LOG_UDF_READ_INVALID_MAX_ARRAY_LENGTH_ARG_VALUE uint16 = 3218
	//OBSOLETE_ER_AUDIT_LOG_JSON_FILTER_PARSING_ERROR uint16 = 3219
	//OBSOLETE_ER_AUDIT_LOG_JSON_FILTER_NAME_CANNOT_BE_EMPTY uint16 = 3220
	//OBSOLETE_ER_AUDIT_LOG_JSON_USER_NAME_CANNOT_BE_EMPTY uint16 = 3221
	//OBSOLETE_ER_AUDIT_LOG_JSON_FILTER_DOES_NOT_EXISTS uint16 = 3222
	//OBSOLETE_ER_AUDIT_LOG_USER_FIRST_CHARACTER_MUST_BE_ALPHANUMERIC uint16 = 3223
	//OBSOLETE_ER_AUDIT_LOG_USER_NAME_INVALID_CHARACTER uint16 = 3224
	//OBSOLETE_ER_AUDIT_LOG_HOST_NAME_INVALID_CHARACTER uint16 = 3225
	OBSOLETE_WARN_DEPRECATED_MAXDB_SQL_MODE_FOR_TIMESTAMP uint16 = 3226
	//OBSOLETE_ER_XA_REPLICATION_FILTERS uint16 = 3227
	//OBSOLETE_ER_CANT_OPEN_ERROR_LOG uint16 = 3228
	//OBSOLETE_ER_GROUPING_ON_TIMESTAMP_IN_DST uint16 = 3229
	//OBSOLETE_ER_CANT_START_SERVER_NAMED_PIPE uint16 = 3230
	ER_WRITE_SET_EXCEEDS_LIMIT uint16 = 3231
	ER_UNSUPPORT_COMPRESSED_TEMPORARY_TABLE uint16 = 3500
	ER_ACL_OPERATION_FAILED uint16 = 3501
	ER_UNSUPPORTED_INDEX_ALGORITHM uint16 = 3502
	ER_NO_SUCH_DB uint16 = 3503
	ER_TOO_BIG_ENUM uint16 = 3504
	ER_TOO_LONG_SET_ENUM_VALUE uint16 = 3505
	ER_INVALID_DD_OBJECT uint16 = 3506
	ER_UPDATING_DD_TABLE uint16 = 3507
	ER_INVALID_DD_OBJECT_ID uint16 = 3508
	ER_INVALID_DD_OBJECT_NAME uint16 = 3509
	ER_TABLESPACE_MISSING_WITH_NAME uint16 = 3510
	ER_TOO_LONG_ROUTINE_COMMENT uint16 = 3511
	ER_SP_LOAD_FAILED uint16 = 3512
	ER_INVALID_BITWISE_OPERANDS_SIZE uint16 = 3513
	ER_INVALID_BITWISE_AGGREGATE_OPERANDS_SIZE uint16 = 3514
	ER_WARN_UNSUPPORTED_HINT uint16 = 3515
	ER_UNEXPECTED_GEOMETRY_TYPE uint16 = 3516
	ER_SRS_PARSE_ERROR uint16 = 3517
	ER_SRS_PROJ_PARAMETER_MISSING uint16 = 3518
	ER_WARN_SRS_NOT_FOUND uint16 = 3519
	ER_SRS_NOT_CARTESIAN uint16 = 3520
	ER_SRS_NOT_CARTESIAN_UNDEFINED uint16 = 3521
	ER_PK_INDEX_CANT_BE_INVISIBLE uint16 = 3522
	ER_UNKNOWN_AUTHID uint16 = 3523
	ER_FAILED_ROLE_GRANT uint16 = 3524
	ER_OPEN_ROLE_TABLES uint16 = 3525
	ER_FAILED_DEFAULT_ROLES uint16 = 3526
	ER_COMPONENTS_NO_SCHEME uint16 = 3527
	ER_COMPONENTS_NO_SCHEME_SERVICE uint16 = 3528
	ER_COMPONENTS_CANT_LOAD uint16 = 3529
	ER_ROLE_NOT_GRANTED uint16 = 3530
	ER_FAILED_REVOKE_ROLE uint16 = 3531
	ER_RENAME_ROLE uint16 = 3532
	ER_COMPONENTS_CANT_ACQUIRE_SERVICE_IMPLEMENTATION uint16 = 3533
	ER_COMPONENTS_CANT_SATISFY_DEPENDENCY uint16 = 3534
	ER_COMPONENTS_LOAD_CANT_REGISTER_SERVICE_IMPLEMENTATION uint16 = 3535
	ER_COMPONENTS_LOAD_CANT_INITIALIZE uint16 = 3536
	ER_COMPONENTS_UNLOAD_NOT_LOADED uint16 = 3537
	ER_COMPONENTS_UNLOAD_CANT_DEINITIALIZE uint16 = 3538
	ER_COMPONENTS_CANT_RELEASE_SERVICE uint16 = 3539
	ER_COMPONENTS_UNLOAD_CANT_UNREGISTER_SERVICE uint16 = 3540
	ER_COMPONENTS_CANT_UNLOAD uint16 = 3541
	ER_WARN_UNLOAD_THE_NOT_PERSISTED uint16 = 3542
	ER_COMPONENT_TABLE_INCORRECT uint16 = 3543
	ER_COMPONENT_MANIPULATE_ROW_FAILED uint16 = 3544
	ER_COMPONENTS_UNLOAD_DUPLICATE_IN_GROUP uint16 = 3545
	ER_CANT_SET_GTID_PURGED_DUE_SETS_CONSTRAINTS uint16 = 3546
	ER_CANNOT_LOCK_USER_MANAGEMENT_CACHES uint16 = 3547
	ER_SRS_NOT_FOUND uint16 = 3548
	ER_VARIABLE_NOT_PERSISTED uint16 = 3549
	ER_IS_QUERY_INVALID_CLAUSE uint16 = 3550
	ER_UNABLE_TO_STORE_STATISTICS uint16 = 3551
	ER_NO_SYSTEM_SCHEMA_ACCESS uint16 = 3552
	ER_NO_SYSTEM_TABLESPACE_ACCESS uint16 = 3553
	ER_NO_SYSTEM_TABLE_ACCESS uint16 = 3554
	ER_NO_SYSTEM_TABLE_ACCESS_FOR_DICTIONARY_TABLE uint16 = 3555
	ER_NO_SYSTEM_TABLE_ACCESS_FOR_SYSTEM_TABLE uint16 = 3556
	ER_NO_SYSTEM_TABLE_ACCESS_FOR_TABLE uint16 = 3557
	ER_INVALID_OPTION_KEY uint16 = 3558
	ER_INVALID_OPTION_VALUE uint16 = 3559
	ER_INVALID_OPTION_KEY_VALUE_PAIR uint16 = 3560
	ER_INVALID_OPTION_START_CHARACTER uint16 = 3561
	ER_INVALID_OPTION_END_CHARACTER uint16 = 3562
	ER_INVALID_OPTION_CHARACTERS uint16 = 3563
	ER_DUPLICATE_OPTION_KEY uint16 = 3564
	ER_WARN_SRS_NOT_FOUND_AXIS_ORDER uint16 = 3565
	ER_NO_ACCESS_TO_NATIVE_FCT uint16 = 3566
	ER_RESET_MASTER_TO_VALUE_OUT_OF_RANGE uint16 = 3567
	ER_UNRESOLVED_TABLE_LOCK uint16 = 3568
	ER_DUPLICATE_TABLE_LOCK uint16 = 3569
	ER_BINLOG_UNSAFE_SKIP_LOCKED uint16 = 3570
	ER_BINLOG_UNSAFE_NOWAIT uint16 = 3571
	ER_LOCK_NOWAIT uint16 = 3572
	ER_CTE_RECURSIVE_REQUIRES_UNION uint16 = 3573
	ER_CTE_RECURSIVE_REQUIRES_NONRECURSIVE_FIRST uint16 = 3574
	ER_CTE_RECURSIVE_FORBIDS_AGGREGATION uint16 = 3575
	ER_CTE_RECURSIVE_FORBIDDEN_JOIN_ORDER uint16 = 3576
	ER_CTE_RECURSIVE_REQUIRES_SINGLE_REFERENCE uint16 = 3577
	ER_SWITCH_TMP_ENGINE uint16 = 3578
	ER_WINDOW_NO_SUCH_WINDOW uint16 = 3579
	ER_WINDOW_CIRCULARITY_IN_WINDOW_GRAPH uint16 = 3580
	ER_WINDOW_NO_CHILD_PARTITIONING uint16 = 3581
	ER_WINDOW_NO_INHERIT_FRAME uint16 = 3582
	ER_WINDOW_NO_REDEFINE_ORDER_BY uint16 = 3583
	ER_WINDOW_FRAME_START_ILLEGAL uint16 = 3584
	ER_WINDOW_FRAME_END_ILLEGAL uint16 = 3585
	ER_WINDOW_FRAME_ILLEGAL uint16 = 3586
	ER_WINDOW_RANGE_FRAME_ORDER_TYPE uint16 = 3587
	ER_WINDOW_RANGE_FRAME_TEMPORAL_TYPE uint16 = 3588
	ER_WINDOW_RANGE_FRAME_NUMERIC_TYPE uint16 = 3589
	ER_WINDOW_RANGE_BOUND_NOT_CONSTANT uint16 = 3590
	ER_WINDOW_DUPLICATE_NAME uint16 = 3591
	ER_WINDOW_ILLEGAL_ORDER_BY uint16 = 3592
	ER_WINDOW_INVALID_WINDOW_FUNC_USE uint16 = 3593
	ER_WINDOW_INVALID_WINDOW_FUNC_ALIAS_USE uint16 = 3594
	ER_WINDOW_NESTED_WINDOW_FUNC_USE_IN_WINDOW_SPEC uint16 = 3595
	ER_WINDOW_ROWS_INTERVAL_USE uint16 = 3596
	ER_WINDOW_NO_GROUP_ORDER_UNUSED uint16 = 3597
	ER_WINDOW_EXPLAIN_JSON uint16 = 3598
	ER_WINDOW_FUNCTION_IGNORES_FRAME uint16 = 3599
	ER_WL9236_NOW_UNUSED uint16 = 3600
	ER_INVALID_NO_OF_ARGS uint16 = 3601
	ER_FIELD_IN_GROUPING_NOT_GROUP_BY uint16 = 3602
	ER_TOO_LONG_TABLESPACE_COMMENT uint16 = 3603
	ER_ENGINE_CANT_DROP_TABLE uint16 = 3604
	ER_ENGINE_CANT_DROP_MISSING_TABLE uint16 = 3605
	ER_TABLESPACE_DUP_FILENAME uint16 = 3606
	ER_DB_DROP_RMDIR2 uint16 = 3607
	ER_IMP_NO_FILES_MATCHED uint16 = 3608
	ER_IMP_SCHEMA_DOES_NOT_EXIST uint16 = 3609
	ER_IMP_TABLE_ALREADY_EXISTS uint16 = 3610
	ER_IMP_INCOMPATIBLE_MYSQLD_VERSION uint16 = 3611
	ER_IMP_INCOMPATIBLE_DD_VERSION uint16 = 3612
	ER_IMP_INCOMPATIBLE_SDI_VERSION uint16 = 3613
	ER_WARN_INVALID_HINT uint16 = 3614
	ER_VAR_DOES_NOT_EXIST uint16 = 3615
	ER_LONGITUDE_OUT_OF_RANGE uint16 = 3616
	ER_LATITUDE_OUT_OF_RANGE uint16 = 3617
	ER_NOT_IMPLEMENTED_FOR_GEOGRAPHIC_SRS uint16 = 3618
	ER_ILLEGAL_PRIVILEGE_LEVEL uint16 = 3619
	ER_NO_SYSTEM_VIEW_ACCESS uint16 = 3620
	ER_COMPONENT_FILTER_FLABBERGASTED uint16 = 3621
	ER_PART_EXPR_TOO_LONG uint16 = 3622
	ER_UDF_DROP_DYNAMICALLY_REGISTERED uint16 = 3623
	ER_UNABLE_TO_STORE_COLUMN_STATISTICS uint16 = 3624
	ER_UNABLE_TO_UPDATE_COLUMN_STATISTICS uint16 = 3625
	ER_UNABLE_TO_DROP_COLUMN_STATISTICS uint16 = 3626
	ER_UNABLE_TO_BUILD_HISTOGRAM uint16 = 3627
	ER_MANDATORY_ROLE uint16 = 3628
	ER_MISSING_TABLESPACE_FILE uint16 = 3629
	ER_PERSIST_ONLY_ACCESS_DENIED_ERROR uint16 = 3630
	ER_CMD_NEED_SUPER uint16 = 3631
	ER_PATH_IN_DATADIR uint16 = 3632
	ER_CLONE_DDL_IN_PROGRESS uint16 = 3633
	ER_CLONE_TOO_MANY_CONCURRENT_CLONES uint16 = 3634
	ER_APPLIER_LOG_EVENT_VALIDATION_ERROR uint16 = 3635
	ER_CTE_MAX_RECURSION_DEPTH uint16 = 3636
	ER_NOT_HINT_UPDATABLE_VARIABLE uint16 = 3637
	ER_CREDENTIALS_CONTRADICT_TO_HISTORY uint16 = 3638
	ER_WARNING_PASSWORD_HISTORY_CLAUSES_VOID uint16 = 3639
	ER_CLIENT_DOES_NOT_SUPPORT uint16 = 3640
	ER_I_S_SKIPPED_TABLESPACE uint16 = 3641
	ER_TABLESPACE_ENGINE_MISMATCH uint16 = 3642
	ER_WRONG_SRID_FOR_COLUMN uint16 = 3643
	ER_CANNOT_ALTER_SRID_DUE_TO_INDEX uint16 = 3644
	ER_WARN_BINLOG_PARTIAL_UPDATES_DISABLED uint16 = 3645
	ER_WARN_BINLOG_V1_ROW_EVENTS_DISABLED uint16 = 3646
	ER_WARN_BINLOG_PARTIAL_UPDATES_SUGGESTS_PARTIAL_IMAGES uint16 = 3647
	ER_COULD_NOT_APPLY_JSON_DIFF uint16 = 3648
	ER_CORRUPTED_JSON_DIFF uint16 = 3649
	ER_RESOURCE_GROUP_EXISTS uint16 = 3650
	ER_RESOURCE_GROUP_NOT_EXISTS uint16 = 3651
	ER_INVALID_VCPU_ID uint16 = 3652
	ER_INVALID_VCPU_RANGE uint16 = 3653
	ER_INVALID_THREAD_PRIORITY uint16 = 3654
	ER_DISALLOWED_OPERATION uint16 = 3655
	ER_RESOURCE_GROUP_BUSY uint16 = 3656
	ER_RESOURCE_GROUP_DISABLED uint16 = 3657
	ER_FEATURE_UNSUPPORTED uint16 = 3658
	ER_ATTRIBUTE_IGNORED uint16 = 3659
	ER_INVALID_THREAD_ID uint16 = 3660
	ER_RESOURCE_GROUP_BIND_FAILED uint16 = 3661
	ER_INVALID_USE_OF_FORCE_OPTION uint16 = 3662
	ER_GROUP_REPLICATION_COMMAND_FAILURE uint16 = 3663
	ER_SDI_OPERATION_FAILED uint16 = 3664
	ER_MISSING_JSON_TABLE_VALUE uint16 = 3665
	ER_WRONG_JSON_TABLE_VALUE uint16 = 3666
	ER_TF_MUST_HAVE_ALIAS uint16 = 3667
	ER_TF_FORBIDDEN_JOIN_TYPE uint16 = 3668
	ER_JT_VALUE_OUT_OF_RANGE uint16 = 3669
	ER_JT_MAX_NESTED_PATH uint16 = 3670
	ER_PASSWORD_EXPIRATION_NOT_SUPPORTED_BY_AUTH_METHOD uint16 = 3671
	ER_INVALID_GEOJSON_CRS_NOT_TOP_LEVEL uint16 = 3672
	ER_BAD_NULL_ERROR_NOT_IGNORED uint16 = 3673
	WARN_USELESS_SPATIAL_INDEX uint16 = 3674
	ER_DISK_FULL_NOWAIT uint16 = 3675
	ER_PARSE_ERROR_IN_DIGEST_FN uint16 = 3676
	ER_UNDISCLOSED_PARSE_ERROR_IN_DIGEST_FN uint16 = 3677
	ER_SCHEMA_DIR_EXISTS uint16 = 3678
	ER_SCHEMA_DIR_MISSING uint16 = 3679
	ER_SCHEMA_DIR_CREATE_FAILED uint16 = 3680
	ER_SCHEMA_DIR_UNKNOWN uint16 = 3681
	ER_ONLY_IMPLEMENTED_FOR_SRID_0_AND_4326 uint16 = 3682
	ER_BINLOG_EXPIRE_LOG_DAYS_AND_SECS_USED_TOGETHER uint16 = 3683
	ER_REGEXP_BUFFER_OVERFLOW uint16 = 3684
	ER_REGEXP_ILLEGAL_ARGUMENT uint16 = 3685
	ER_REGEXP_INDEX_OUTOFBOUNDS_ERROR uint16 = 3686
	ER_REGEXP_INTERNAL_ERROR uint16 = 3687
	ER_REGEXP_RULE_SYNTAX uint16 = 3688
	ER_REGEXP_BAD_ESCAPE_SEQUENCE uint16 = 3689
	ER_REGEXP_UNIMPLEMENTED uint16 = 3690
	ER_REGEXP_MISMATCHED_PAREN uint16 = 3691
	ER_REGEXP_BAD_INTERVAL uint16 = 3692
	ER_REGEXP_MAX_LT_MIN uint16 = 3693
	ER_REGEXP_INVALID_BACK_REF uint16 = 3694
	ER_REGEXP_LOOK_BEHIND_LIMIT uint16 = 3695
	ER_REGEXP_MISSING_CLOSE_BRACKET uint16 = 3696
	ER_REGEXP_INVALID_RANGE uint16 = 3697
	ER_REGEXP_STACK_OVERFLOW uint16 = 3698
	ER_REGEXP_TIME_OUT uint16 = 3699
	ER_REGEXP_PATTERN_TOO_BIG uint16 = 3700
	ER_CANT_SET_ERROR_LOG_SERVICE uint16 = 3701
	ER_EMPTY_PIPELINE_FOR_ERROR_LOG_SERVICE uint16 = 3702
	ER_COMPONENT_FILTER_DIAGNOSTICS uint16 = 3703
	ER_NOT_IMPLEMENTED_FOR_CARTESIAN_SRS uint16 = 3704
	ER_NOT_IMPLEMENTED_FOR_PROJECTED_SRS uint16 = 3705
	ER_NONPOSITIVE_RADIUS uint16 = 3706
	ER_RESTART_SERVER_FAILED uint16 = 3707
	ER_SRS_MISSING_MANDATORY_ATTRIBUTE uint16 = 3708
	ER_SRS_MULTIPLE_ATTRIBUTE_DEFINITIONS uint16 = 3709
	ER_SRS_NAME_CANT_BE_EMPTY_OR_WHITESPACE uint16 = 3710
	ER_SRS_ORGANIZATION_CANT_BE_EMPTY_OR_WHITESPACE uint16 = 3711
	ER_SRS_ID_ALREADY_EXISTS uint16 = 3712
	ER_WARN_SRS_ID_ALREADY_EXISTS uint16 = 3713
	ER_CANT_MODIFY_SRID_0 uint16 = 3714
	ER_WARN_RESERVED_SRID_RANGE uint16 = 3715
	ER_CANT_MODIFY_SRS_USED_BY_COLUMN uint16 = 3716
	ER_SRS_INVALID_CHARACTER_IN_ATTRIBUTE uint16 = 3717
	ER_SRS_ATTRIBUTE_STRING_TOO_LONG uint16 = 3718
	ER_DEPRECATED_UTF8_ALIAS uint16 = 3719
	ER_DEPRECATED_NATIONAL uint16 = 3720
	ER_INVALID_DEFAULT_UTF8MB4_COLLATION uint16 = 3721
	ER_UNABLE_TO_COLLECT_LOG_STATUS uint16 = 3722
	ER_RESERVED_TABLESPACE_NAME uint16 = 3723
	ER_UNABLE_TO_SET_OPTION uint16 = 3724
	ER_SLAVE_POSSIBLY_DIVERGED_AFTER_DDL uint16 = 3725
	ER_SRS_NOT_GEOGRAPHIC uint16 = 3726
	ER_POLYGON_TOO_LARGE uint16 = 3727
	ER_SPATIAL_UNIQUE_INDEX uint16 = 3728
	ER_INDEX_TYPE_NOT_SUPPORTED_FOR_SPATIAL_INDEX uint16 = 3729
	ER_FK_CANNOT_DROP_PARENT uint16 = 3730
	ER_GEOMETRY_PARAM_LONGITUDE_OUT_OF_RANGE uint16 = 3731
	ER_GEOMETRY_PARAM_LATITUDE_OUT_OF_RANGE uint16 = 3732
	ER_FK_CANNOT_USE_VIRTUAL_COLUMN uint16 = 3733
	ER_FK_NO_COLUMN_PARENT uint16 = 3734
	ER_CANT_SET_ERROR_SUPPRESSION_LIST uint16 = 3735
	ER_SRS_GEOGCS_INVALID_AXES uint16 = 3736
	ER_SRS_INVALID_SEMI_MAJOR_AXIS uint16 = 3737
	ER_SRS_INVALID_INVERSE_FLATTENING uint16 = 3738
	ER_SRS_INVALID_ANGULAR_UNIT uint16 = 3739
	ER_SRS_INVALID_PRIME_MERIDIAN uint16 = 3740
	ER_TRANSFORM_SOURCE_SRS_NOT_SUPPORTED uint16 = 3741
	ER_TRANSFORM_TARGET_SRS_NOT_SUPPORTED uint16 = 3742
	ER_TRANSFORM_SOURCE_SRS_MISSING_TOWGS84 uint16 = 3743
	ER_TRANSFORM_TARGET_SRS_MISSING_TOWGS84 uint16 = 3744
	ER_TEMP_TABLE_PREVENTS_SWITCH_SESSION_BINLOG_FORMAT uint16 = 3745
	ER_TEMP_TABLE_PREVENTS_SWITCH_GLOBAL_BINLOG_FORMAT uint16 = 3746
	ER_RUNNING_APPLIER_PREVENTS_SWITCH_GLOBAL_BINLOG_FORMAT uint16 = 3747
	ER_CLIENT_GTID_UNSAFE_CREATE_DROP_TEMP_TABLE_IN_TRX_IN_SBR uint16 = 3748
	//OBSOLETE_ER_XA_CANT_CREATE_MDL_BACKUP uint16 = 3749
	ER_TABLE_WITHOUT_PK uint16 = 3750
	ER_WARN_DATA_TRUNCATED_FUNCTIONAL_INDEX uint16 = 3751
	ER_WARN_DATA_OUT_OF_RANGE_FUNCTIONAL_INDEX uint16 = 3752
	ER_FUNCTIONAL_INDEX_ON_JSON_OR_GEOMETRY_FUNCTION uint16 = 3753
	ER_FUNCTIONAL_INDEX_REF_AUTO_INCREMENT uint16 = 3754
	ER_CANNOT_DROP_COLUMN_FUNCTIONAL_INDEX uint16 = 3755
	ER_FUNCTIONAL_INDEX_PRIMARY_KEY uint16 = 3756
	ER_FUNCTIONAL_INDEX_ON_LOB uint16 = 3757
	ER_FUNCTIONAL_INDEX_FUNCTION_IS_NOT_ALLOWED uint16 = 3758
	ER_FULLTEXT_FUNCTIONAL_INDEX uint16 = 3759
	ER_SPATIAL_FUNCTIONAL_INDEX uint16 = 3760
	ER_WRONG_KEY_COLUMN_FUNCTIONAL_INDEX uint16 = 3761
	ER_FUNCTIONAL_INDEX_ON_FIELD uint16 = 3762
	ER_GENERATED_COLUMN_NAMED_FUNCTION_IS_NOT_ALLOWED uint16 = 3763
	ER_GENERATED_COLUMN_ROW_VALUE uint16 = 3764
	ER_GENERATED_COLUMN_VARIABLES uint16 = 3765
	ER_DEPENDENT_BY_DEFAULT_GENERATED_VALUE uint16 = 3766
	ER_DEFAULT_VAL_GENERATED_NON_PRIOR uint16 = 3767
	ER_DEFAULT_VAL_GENERATED_REF_AUTO_INC uint16 = 3768
	ER_DEFAULT_VAL_GENERATED_FUNCTION_IS_NOT_ALLOWED uint16 = 3769
	ER_DEFAULT_VAL_GENERATED_NAMED_FUNCTION_IS_NOT_ALLOWED uint16 = 3770
	ER_DEFAULT_VAL_GENERATED_ROW_VALUE uint16 = 3771
	ER_DEFAULT_VAL_GENERATED_VARIABLES uint16 = 3772
	ER_DEFAULT_AS_VAL_GENERATED uint16 = 3773
	ER_UNSUPPORTED_ACTION_ON_DEFAULT_VAL_GENERATED uint16 = 3774
	ER_GTID_UNSAFE_ALTER_ADD_COL_WITH_DEFAULT_EXPRESSION uint16 = 3775
	ER_FK_CANNOT_CHANGE_ENGINE uint16 = 3776
	ER_WARN_DEPRECATED_USER_SET_EXPR uint16 = 3777
	ER_WARN_DEPRECATED_UTF8MB3_COLLATION uint16 = 3778
	ER_WARN_DEPRECATED_NESTED_COMMENT_SYNTAX uint16 = 3779
	ER_FK_INCOMPATIBLE_COLUMNS uint16 = 3780
	ER_GR_HOLD_WAIT_TIMEOUT uint16 = 3781
	ER_GR_HOLD_KILLED uint16 = 3782
	ER_GR_HOLD_MEMBER_STATUS_ERROR uint16 = 3783
	ER_RPL_ENCRYPTION_FAILED_TO_FETCH_KEY uint16 = 3784
	ER_RPL_ENCRYPTION_KEY_NOT_FOUND uint16 = 3785
	ER_RPL_ENCRYPTION_KEYRING_INVALID_KEY uint16 = 3786
	ER_RPL_ENCRYPTION_HEADER_ERROR uint16 = 3787
	ER_RPL_ENCRYPTION_FAILED_TO_ROTATE_LOGS uint16 = 3788
	ER_RPL_ENCRYPTION_KEY_EXISTS_UNEXPECTED uint16 = 3789
	ER_RPL_ENCRYPTION_FAILED_TO_GENERATE_KEY uint16 = 3790
	ER_RPL_ENCRYPTION_FAILED_TO_STORE_KEY uint16 = 3791
	ER_RPL_ENCRYPTION_FAILED_TO_REMOVE_KEY uint16 = 3792
	ER_RPL_ENCRYPTION_UNABLE_TO_CHANGE_OPTION uint16 = 3793
	ER_RPL_ENCRYPTION_MASTER_KEY_RECOVERY_FAILED uint16 = 3794
	ER_SLOW_LOG_MODE_IGNORED_WHEN_NOT_LOGGING_TO_FILE uint16 = 3795
	ER_GRP_TRX_CONSISTENCY_NOT_ALLOWED uint16 = 3796
	ER_GRP_TRX_CONSISTENCY_BEFORE uint16 = 3797
	ER_GRP_TRX_CONSISTENCY_AFTER_ON_TRX_BEGIN uint16 = 3798
	ER_GRP_TRX_CONSISTENCY_BEGIN_NOT_ALLOWED uint16 = 3799
	ER_FUNCTIONAL_INDEX_ROW_VALUE_IS_NOT_ALLOWED uint16 = 3800
	ER_RPL_ENCRYPTION_FAILED_TO_ENCRYPT uint16 = 3801
	ER_PAGE_TRACKING_NOT_STARTED uint16 = 3802
	ER_PAGE_TRACKING_RANGE_NOT_TRACKED uint16 = 3803
	ER_PAGE_TRACKING_CANNOT_PURGE uint16 = 3804
	ER_RPL_ENCRYPTION_CANNOT_ROTATE_BINLOG_MASTER_KEY uint16 = 3805
	ER_BINLOG_MASTER_KEY_RECOVERY_OUT_OF_COMBINATION uint16 = 3806
	ER_BINLOG_MASTER_KEY_ROTATION_FAIL_TO_OPERATE_KEY uint16 = 3807
	ER_BINLOG_MASTER_KEY_ROTATION_FAIL_TO_ROTATE_LOGS uint16 = 3808
	ER_BINLOG_MASTER_KEY_ROTATION_FAIL_TO_REENCRYPT_LOG uint16 = 3809
	ER_BINLOG_MASTER_KEY_ROTATION_FAIL_TO_CLEANUP_UNUSED_KEYS uint16 = 3810
	ER_BINLOG_MASTER_KEY_ROTATION_FAIL_TO_CLEANUP_AUX_KEY uint16 = 3811
	ER_NON_BOOLEAN_EXPR_FOR_CHECK_CONSTRAINT uint16 = 3812
	ER_COLUMN_CHECK_CONSTRAINT_REFERENCES_OTHER_COLUMN uint16 = 3813
	ER_CHECK_CONSTRAINT_NAMED_FUNCTION_IS_NOT_ALLOWED uint16 = 3814
	ER_CHECK_CONSTRAINT_FUNCTION_IS_NOT_ALLOWED uint16 = 3815
	ER_CHECK_CONSTRAINT_VARIABLES uint16 = 3816
	ER_CHECK_CONSTRAINT_ROW_VALUE uint16 = 3817
	ER_CHECK_CONSTRAINT_REFERS_AUTO_INCREMENT_COLUMN uint16 = 3818
	ER_CHECK_CONSTRAINT_VIOLATED uint16 = 3819
	ER_CHECK_CONSTRAINT_REFERS_UNKNOWN_COLUMN uint16 = 3820
	ER_CHECK_CONSTRAINT_NOT_FOUND uint16 = 3821
	ER_CHECK_CONSTRAINT_DUP_NAME uint16 = 3822
	ER_CHECK_CONSTRAINT_CLAUSE_USING_FK_REFER_ACTION_COLUMN uint16 = 3823
	WARN_UNENCRYPTED_TABLE_IN_ENCRYPTED_DB uint16 = 3824
	ER_INVALID_ENCRYPTION_REQUEST uint16 = 3825
	ER_CANNOT_SET_TABLE_ENCRYPTION uint16 = 3826
	ER_CANNOT_SET_DATABASE_ENCRYPTION uint16 = 3827
	ER_CANNOT_SET_TABLESPACE_ENCRYPTION uint16 = 3828
	ER_TABLESPACE_CANNOT_BE_ENCRYPTED uint16 = 3829
	ER_TABLESPACE_CANNOT_BE_DECRYPTED uint16 = 3830
	ER_TABLESPACE_TYPE_UNKNOWN uint16 = 3831
	ER_TARGET_TABLESPACE_UNENCRYPTED uint16 = 3832
	ER_CANNOT_USE_ENCRYPTION_CLAUSE uint16 = 3833
	ER_INVALID_MULTIPLE_CLAUSES uint16 = 3834
	ER_UNSUPPORTED_USE_OF_GRANT_AS uint16 = 3835
	ER_UKNOWN_AUTH_ID_OR_ACCESS_DENIED_FOR_GRANT_AS uint16 = 3836
	ER_DEPENDENT_BY_FUNCTIONAL_INDEX uint16 = 3837
	ER_PLUGIN_NOT_EARLY uint16 = 3838
	ER_INNODB_REDO_LOG_ARCHIVE_START_SUBDIR_PATH uint16 = 3839
	ER_INNODB_REDO_LOG_ARCHIVE_START_TIMEOUT uint16 = 3840
	ER_INNODB_REDO_LOG_ARCHIVE_DIRS_INVALID uint16 = 3841
	ER_INNODB_REDO_LOG_ARCHIVE_LABEL_NOT_FOUND uint16 = 3842
	ER_INNODB_REDO_LOG_ARCHIVE_DIR_EMPTY uint16 = 3843
	ER_INNODB_REDO_LOG_ARCHIVE_NO_SUCH_DIR uint16 = 3844
	ER_INNODB_REDO_LOG_ARCHIVE_DIR_CLASH uint16 = 3845
	ER_INNODB_REDO_LOG_ARCHIVE_DIR_PERMISSIONS uint16 = 3846
	ER_INNODB_REDO_LOG_ARCHIVE_FILE_CREATE uint16 = 3847
	ER_INNODB_REDO_LOG_ARCHIVE_ACTIVE uint16 = 3848
	ER_INNODB_REDO_LOG_ARCHIVE_INACTIVE uint16 = 3849
	ER_INNODB_REDO_LOG_ARCHIVE_FAILED uint16 = 3850
	ER_INNODB_REDO_LOG_ARCHIVE_SESSION uint16 = 3851
	ER_STD_REGEX_ERROR uint16 = 3852
	ER_INVALID_JSON_TYPE uint16 = 3853
	ER_CANNOT_CONVERT_STRING uint16 = 3854
	ER_DEPENDENT_BY_PARTITION_FUNC uint16 = 3855
	ER_WARN_DEPRECATED_FLOAT_AUTO_INCREMENT uint16 = 3856
	ER_RPL_CANT_STOP_SLAVE_WHILE_LOCKED_BACKUP uint16 = 3857
	ER_WARN_DEPRECATED_FLOAT_DIGITS uint16 = 3858
	ER_WARN_DEPRECATED_FLOAT_UNSIGNED uint16 = 3859
	ER_WARN_DEPRECATED_INTEGER_DISPLAY_WIDTH uint16 = 3860
	ER_WARN_DEPRECATED_ZEROFILL uint16 = 3861
	ER_CLONE_DONOR uint16 = 3862
	ER_CLONE_PROTOCOL uint16 = 3863
	ER_CLONE_DONOR_VERSION uint16 = 3864
	ER_CLONE_OS uint16 = 3865
	ER_CLONE_PLATFORM uint16 = 3866
	ER_CLONE_CHARSET uint16 = 3867
	ER_CLONE_CONFIG uint16 = 3868
	ER_CLONE_SYS_CONFIG uint16 = 3869
	ER_CLONE_PLUGIN_MATCH uint16 = 3870
	ER_CLONE_LOOPBACK uint16 = 3871
	ER_CLONE_ENCRYPTION uint16 = 3872
	ER_CLONE_DISK_SPACE uint16 = 3873
	ER_CLONE_IN_PROGRESS uint16 = 3874
	ER_CLONE_DISALLOWED uint16 = 3875
	ER_CANNOT_GRANT_ROLES_TO_ANONYMOUS_USER uint16 = 3876
	ER_SECONDARY_ENGINE_PLUGIN uint16 = 3877
	ER_SECOND_PASSWORD_CANNOT_BE_EMPTY uint16 = 3878
	ER_DB_ACCESS_DENIED uint16 = 3879
	ER_DA_AUTH_ID_WITH_SYSTEM_USER_PRIV_IN_MANDATORY_ROLES uint16 = 3880
	ER_DA_RPL_GTID_TABLE_CANNOT_OPEN uint16 = 3881
	ER_GEOMETRY_IN_UNKNOWN_LENGTH_UNIT uint16 = 3882
	ER_DA_PLUGIN_INSTALL_ERROR uint16 = 3883
	ER_NO_SESSION_TEMP uint16 = 3884
	ER_DA_UNKNOWN_ERROR_NUMBER uint16 = 3885
	ER_COLUMN_CHANGE_SIZE uint16 = 3886
	ER_REGEXP_INVALID_CAPTURE_GROUP_NAME uint16 = 3887
	ER_DA_SSL_LIBRARY_ERROR uint16 = 3888
	ER_SECONDARY_ENGINE uint16 = 3889
	ER_SECONDARY_ENGINE_DDL uint16 = 3890
	ER_INCORRECT_CURRENT_PASSWORD uint16 = 3891
	ER_MISSING_CURRENT_PASSWORD uint16 = 3892
	ER_CURRENT_PASSWORD_NOT_REQUIRED uint16 = 3893
	ER_PASSWORD_CANNOT_BE_RETAINED_ON_PLUGIN_CHANGE uint16 = 3894
	ER_CURRENT_PASSWORD_CANNOT_BE_RETAINED uint16 = 3895
	ER_PARTIAL_REVOKES_EXIST uint16 = 3896
	ER_CANNOT_GRANT_SYSTEM_PRIV_TO_MANDATORY_ROLE uint16 = 3897
	ER_XA_REPLICATION_FILTERS uint16 = 3898
	ER_UNSUPPORTED_SQL_MODE uint16 = 3899
	ER_REGEXP_INVALID_FLAG uint16 = 3900
	ER_PARTIAL_REVOKE_AND_DB_GRANT_BOTH_EXISTS uint16 = 3901
	ER_UNIT_NOT_FOUND uint16 = 3902
	ER_INVALID_JSON_VALUE_FOR_FUNC_INDEX uint16 = 3903
	ER_JSON_VALUE_OUT_OF_RANGE_FOR_FUNC_INDEX uint16 = 3904
	ER_EXCEEDED_MV_KEYS_NUM uint16 = 3905
	ER_EXCEEDED_MV_KEYS_SPACE uint16 = 3906
	ER_FUNCTIONAL_INDEX_DATA_IS_TOO_LONG uint16 = 3907
	ER_WRONG_MVI_VALUE uint16 = 3908
	ER_WARN_FUNC_INDEX_NOT_APPLICABLE uint16 = 3909
	ER_GRP_RPL_UDF_ERROR uint16 = 3910
	ER_UPDATE_GTID_PURGED_WITH_GR uint16 = 3911
	ER_GROUPING_ON_TIMESTAMP_IN_DST uint16 = 3912
	ER_TABLE_NAME_CAUSES_TOO_LONG_PATH uint16 = 3913
	ER_AUDIT_LOG_INSUFFICIENT_PRIVILEGE uint16 = 3914
	//OBSOLETE_ER_AUDIT_LOG_PASSWORD_HAS_BEEN_COPIED uint16 = 3915
	ER_DA_GRP_RPL_STARTED_AUTO_REJOIN uint16 = 3916
	ER_SYSVAR_CHANGE_DURING_QUERY uint16 = 3917
	ER_GLOBSTAT_CHANGE_DURING_QUERY uint16 = 3918
	ER_GRP_RPL_MESSAGE_SERVICE_INIT_FAILURE uint16 = 3919
	ER_CHANGE_MASTER_WRONG_COMPRESSION_ALGORITHM_CLIENT uint16 = 3920
	ER_CHANGE_MASTER_WRONG_COMPRESSION_LEVEL_CLIENT uint16 = 3921
	ER_WRONG_COMPRESSION_ALGORITHM_CLIENT uint16 = 3922
	ER_WRONG_COMPRESSION_LEVEL_CLIENT uint16 = 3923
	ER_CHANGE_MASTER_WRONG_COMPRESSION_ALGORITHM_LIST_CLIENT uint16 = 3924
	ER_CLIENT_PRIVILEGE_CHECKS_USER_CANNOT_BE_ANONYMOUS uint16 = 3925
	ER_CLIENT_PRIVILEGE_CHECKS_USER_DOES_NOT_EXIST uint16 = 3926
	ER_CLIENT_PRIVILEGE_CHECKS_USER_CORRUPT uint16 = 3927
	ER_CLIENT_PRIVILEGE_CHECKS_USER_NEEDS_RPL_APPLIER_PRIV uint16 = 3928
	ER_WARN_DA_PRIVILEGE_NOT_REGISTERED uint16 = 3929
	ER_CLIENT_KEYRING_UDF_KEY_INVALID uint16 = 3930
	ER_CLIENT_KEYRING_UDF_KEY_TYPE_INVALID uint16 = 3931
	ER_CLIENT_KEYRING_UDF_KEY_TOO_LONG uint16 = 3932
	ER_CLIENT_KEYRING_UDF_KEY_TYPE_TOO_LONG uint16 = 3933
	ER_JSON_SCHEMA_VALIDATION_ERROR_WITH_DETAILED_REPORT uint16 = 3934
	ER_DA_UDF_INVALID_CHARSET_SPECIFIED uint16 = 3935
	ER_DA_UDF_INVALID_CHARSET uint16 = 3936
	ER_DA_UDF_INVALID_COLLATION uint16 = 3937
	ER_DA_UDF_INVALID_EXTENSION_ARGUMENT_TYPE uint16 = 3938
	ER_MULTIPLE_CONSTRAINTS_WITH_SAME_NAME uint16 = 3939
	ER_CONSTRAINT_NOT_FOUND uint16 = 3940
	ER_ALTER_CONSTRAINT_ENFORCEMENT_NOT_SUPPORTED uint16 = 3941
	ER_TABLE_VALUE_CONSTRUCTOR_MUST_HAVE_COLUMNS uint16 = 3942
	ER_TABLE_VALUE_CONSTRUCTOR_CANNOT_HAVE_DEFAULT uint16 = 3943
	ER_CLIENT_QUERY_FAILURE_INVALID_NON_ROW_FORMAT uint16 = 3944
	ER_REQUIRE_ROW_FORMAT_INVALID_VALUE uint16 = 3945
	ER_FAILED_TO_DETERMINE_IF_ROLE_IS_MANDATORY uint16 = 3946
	ER_FAILED_TO_FETCH_MANDATORY_ROLE_LIST uint16 = 3947
	ER_CLIENT_LOCAL_FILES_DISABLED uint16 = 3948
	ER_IMP_INCOMPATIBLE_CFG_VERSION uint16 = 3949
	ER_DA_OOM uint16 = 3950
	ER_DA_UDF_INVALID_ARGUMENT_TO_SET_CHARSET uint16 = 3951
	ER_DA_UDF_INVALID_RETURN_TYPE_TO_SET_CHARSET uint16 = 3952
	ER_MULTIPLE_INTO_CLAUSES uint16 = 3953
	ER_MISPLACED_INTO uint16 = 3954
	ER_USER_ACCESS_DENIED_FOR_USER_ACCOUNT_BLOCKED_BY_PASSWORD_LOCK uint16 = 3955
	ER_WARN_DEPRECATED_YEAR_UNSIGNED uint16 = 3956
	ER_CLONE_NETWORK_PACKET uint16 = 3957
	ER_SDI_OPERATION_FAILED_MISSING_RECORD uint16 = 3958
	ER_DEPENDENT_BY_CHECK_CONSTRAINT uint16 = 3959
	ER_GRP_OPERATION_NOT_ALLOWED_GR_MUST_STOP uint16 = 3960
	ER_WARN_DEPRECATED_JSON_TABLE_ON_ERROR_ON_EMPTY uint16 = 3961
	ER_WARN_DEPRECATED_INNER_INTO uint16 = 3962
	ER_WARN_DEPRECATED_VALUES_FUNCTION_ALWAYS_NULL uint16 = 3963
	ER_WARN_DEPRECATED_SQL_CALC_FOUND_ROWS uint16 = 3964
	ER_WARN_DEPRECATED_FOUND_ROWS uint16 = 3965
	ER_MISSING_JSON_VALUE uint16 = 3966
	ER_MULTIPLE_JSON_VALUES uint16 = 3967
	ER_HOSTNAME_TOO_LONG uint16 = 3968
	ER_WARN_CLIENT_DEPRECATED_PARTITION_PREFIX_KEY uint16 = 3969
	ER_GROUP_REPLICATION_USER_EMPTY_MSG uint16 = 3970
	ER_GROUP_REPLICATION_USER_MANDATORY_MSG uint16 = 3971
	ER_GROUP_REPLICATION_PASSWORD_LENGTH uint16 = 3972
	ER_SUBQUERY_TRANSFORM_REJECTED uint16 = 3973
	ER_DA_GRP_RPL_RECOVERY_ENDPOINT_FORMAT uint16 = 3974
	ER_DA_GRP_RPL_RECOVERY_ENDPOINT_INVALID uint16 = 3975
	ER_WRONG_VALUE_FOR_VAR_PLUS_ACTIONABLE_PART uint16 = 3976
	ER_STATEMENT_NOT_ALLOWED_AFTER_START_TRANSACTION uint16 = 3977
	ER_FOREIGN_KEY_WITH_ATOMIC_CREATE_SELECT uint16 = 3978
	ER_NOT_ALLOWED_WITH_START_TRANSACTION uint16 = 3979
	ER_INVALID_JSON_ATTRIBUTE uint16 = 3980
	ER_ENGINE_ATTRIBUTE_NOT_SUPPORTED uint16 = 3981
	ER_INVALID_USER_ATTRIBUTE_JSON uint16 = 3982
	ER_INNODB_REDO_DISABLED uint16 = 3983
	ER_INNODB_REDO_ARCHIVING_ENABLED uint16 = 3984
	ER_MDL_OUT_OF_RESOURCES uint16 = 3985
	ER_IMPLICIT_COMPARISON_FOR_JSON uint16 = 3986
	ER_FUNCTION_DOES_NOT_SUPPORT_CHARACTER_SET uint16 = 3987
	ER_IMPOSSIBLE_STRING_CONVERSION uint16 = 3988
	ER_SCHEMA_READ_ONLY uint16 = 3989
	ER_RPL_ASYNC_RECONNECT_GTID_MODE_OFF uint16 = 3990
	ER_RPL_ASYNC_RECONNECT_AUTO_POSITION_OFF uint16 = 3991
	ER_DISABLE_GTID_MODE_REQUIRES_ASYNC_RECONNECT_OFF uint16 = 3992
	ER_DISABLE_AUTO_POSITION_REQUIRES_ASYNC_RECONNECT_OFF uint16 = 3993
	ER_INVALID_PARAMETER_USE uint16 = 3994
	ER_CHARACTER_SET_MISMATCH uint16 = 3995
	ER_WARN_VAR_VALUE_CHANGE_NOT_SUPPORTED uint16 = 3996
	ER_INVALID_TIME_ZONE_INTERVAL uint16 = 3997
	ER_INVALID_CAST uint16 = 3998
	ER_HYPERGRAPH_NOT_SUPPORTED_YET uint16 = 3999
	ER_WARN_HYPERGRAPH_EXPERIMENTAL uint16 = 4000
	ER_DA_NO_ERROR_LOG_PARSER_CONFIGURED uint16 = 4001
	ER_DA_ERROR_LOG_TABLE_DISABLED uint16 = 4002
	ER_DA_ERROR_LOG_MULTIPLE_FILTERS uint16 = 4003
	ER_DA_CANT_OPEN_ERROR_LOG uint16 = 4004
	ER_USER_REFERENCED_AS_DEFINER uint16 = 4005
	ER_CANNOT_USER_REFERENCED_AS_DEFINER uint16 = 4006
	ER_REGEX_NUMBER_TOO_BIG uint16 = 4007
	ER_SPVAR_NONINTEGER_TYPE uint16 = 4008
	WARN_UNSUPPORTED_ACL_TABLES_READ uint16 = 4009
	ER_BINLOG_UNSAFE_ACL_TABLE_READ_IN_DML_DDL uint16 = 4010
	ER_STOP_REPLICA_MONITOR_IO_THREAD_TIMEOUT uint16 = 4011
	ER_STARTING_REPLICA_MONITOR_IO_THREAD uint16 = 4012
	ER_CANT_USE_ANONYMOUS_TO_GTID_WITH_GTID_MODE_NOT_ON uint16 = 4013
	ER_CANT_COMBINE_ANONYMOUS_TO_GTID_AND_AUTOPOSITION uint16 = 4014
	ER_ASSIGN_GTIDS_TO_ANONYMOUS_TRANSACTIONS_REQUIRES_GTID_MODE_ON uint16 = 4015
	ER_SQL_SLAVE_SKIP_COUNTER_USED_WITH_GTID_MODE_ON uint16 = 4016
	ER_USING_ASSIGN_GTIDS_TO_ANONYMOUS_TRANSACTIONS_AS_LOCAL_OR_UUID uint16 = 4017
	ER_CANT_SET_ANONYMOUS_TO_GTID_AND_WAIT_UNTIL_SQL_THD_AFTER_GTIDS uint16 = 4018
	ER_CANT_SET_SQL_AFTER_OR_BEFORE_GTIDS_WITH_ANONYMOUS_TO_GTID uint16 = 4019
	ER_ANONYMOUS_TO_GTID_UUID_SAME_AS_GROUP_NAME uint16 = 4020
	ER_CANT_USE_SAME_UUID_AS_GROUP_NAME uint16 = 4021
	ER_GRP_RPL_RECOVERY_CHANNEL_STILL_RUNNING uint16 = 4022
	ER_INNODB_INVALID_AUTOEXTEND_SIZE_VALUE uint16 = 4023
	ER_INNODB_INCOMPATIBLE_WITH_TABLESPACE uint16 = 4024
	ER_INNODB_AUTOEXTEND_SIZE_OUT_OF_RANGE uint16 = 4025
	ER_CANNOT_USE_AUTOEXTEND_SIZE_CLAUSE uint16 = 4026
	ER_ROLE_GRANTED_TO_ITSELF uint16 = 4027
	ER_TABLE_MUST_HAVE_A_VISIBLE_COLUMN uint16 = 4028
	ER_INNODB_COMPRESSION_FAILURE uint16 = 4029
	ER_WARN_ASYNC_CONN_FAILOVER_NETWORK_NAMESPACE uint16 = 4030

	//5,000 to 5,999: Error codes reserved for use by X Plugin for messages sent to clients.

	// no such code

	//10,000 to 49,999: Server error codes reserved for messages to be written to the error log (not sent to clients).

	ER_PARSER_TRACE uint16 = 10000
	ER_BOOTSTRAP_CANT_THREAD uint16 = 10001
	ER_TRIGGER_INVALID_VALUE uint16 = 10002
	ER_OPT_WRONG_TREE uint16 = 10003
	ER_DD_FAILSAFE uint16 = 10004
	ER_DD_NO_WRITES_NO_REPOPULATION uint16 = 10005
	ER_DD_VERSION_FOUND uint16 = 10006
	ER_DD_VERSION_INSTALLED uint16 = 10007
	ER_DD_VERSION_UNSUPPORTED uint16 = 10008
	//OBSOLETE_ER_LOG_SYSLOG_FACILITY_FAIL uint16 = 10009
	ER_LOG_SYSLOG_CANNOT_OPEN uint16 = 10010
	ER_LOG_SLOW_CANNOT_OPEN uint16 = 10011
	ER_LOG_GENERAL_CANNOT_OPEN uint16 = 10012
	ER_LOG_CANNOT_WRITE uint16 = 10013
	ER_RPL_ZOMBIE_ENCOUNTERED uint16 = 10014
	ER_RPL_GTID_TABLE_CANNOT_OPEN uint16 = 10015
	ER_SYSTEM_SCHEMA_NOT_FOUND uint16 = 10016
	ER_DD_INIT_UPGRADE_FAILED uint16 = 10017
	ER_VIEW_UNKNOWN_CHARSET_OR_COLLATION uint16 = 10018
	ER_DD_VIEW_CANT_ALLOC_CHARSET uint16 = 10019
	ER_DD_INIT_FAILED uint16 = 10020
	ER_DD_UPDATING_PLUGIN_MD_FAILED uint16 = 10021
	ER_DD_POPULATING_TABLES_FAILED uint16 = 10022
	ER_DD_VIEW_CANT_CREATE uint16 = 10023
	ER_DD_METADATA_NOT_FOUND uint16 = 10024
	ER_DD_CACHE_NOT_EMPTY_AT_SHUTDOWN uint16 = 10025
	ER_DD_OBJECT_REMAINS uint16 = 10026
	ER_DD_OBJECT_REMAINS_IN_RELEASER uint16 = 10027
	ER_DD_OBJECT_RELEASER_REMAINS uint16 = 10028
	ER_DD_CANT_GET_OBJECT_KEY uint16 = 10029
	ER_DD_CANT_CREATE_OBJECT_KEY uint16 = 10030
	ER_CANT_CREATE_HANDLE_MGR_THREAD uint16 = 10031
	ER_RPL_REPO_HAS_GAPS uint16 = 10032
	ER_INVALID_VALUE_FOR_ENFORCE_GTID_CONSISTENCY uint16 = 10033
	ER_CHANGED_ENFORCE_GTID_CONSISTENCY uint16 = 10034
	ER_CHANGED_GTID_MODE uint16 = 10035
	ER_DISABLED_STORAGE_ENGINE_AS_DEFAULT uint16 = 10036
	ER_DEBUG_SYNC_HIT uint16 = 10037
	ER_DEBUG_SYNC_EXECUTED uint16 = 10038
	ER_DEBUG_SYNC_THREAD_MAX uint16 = 10039
	ER_DEBUG_SYNC_OOM uint16 = 10040
	ER_CANT_INIT_TC_LOG uint16 = 10041
	ER_EVENT_CANT_INIT_QUEUE uint16 = 10042
	ER_EVENT_PURGING_QUEUE uint16 = 10043
	ER_EVENT_LAST_EXECUTION uint16 = 10044
	ER_EVENT_MESSAGE_STACK uint16 = 10045
	ER_EVENT_EXECUTION_FAILED uint16 = 10046
	ER_CANT_INIT_SCHEDULER_THREAD uint16 = 10047
	ER_SCHEDULER_STOPPED uint16 = 10048
	ER_CANT_CREATE_SCHEDULER_THREAD uint16 = 10049
	ER_SCHEDULER_WAITING uint16 = 10050
	ER_SCHEDULER_STARTED uint16 = 10051
	ER_SCHEDULER_STOPPING_FAILED_TO_GET_EVENT uint16 = 10052
	ER_SCHEDULER_STOPPING_FAILED_TO_CREATE_WORKER uint16 = 10053
	ER_SCHEDULER_KILLING uint16 = 10054
	ER_UNABLE_TO_RESOLVE_IP uint16 = 10055
	ER_UNABLE_TO_RESOLVE_HOSTNAME uint16 = 10056
	ER_HOSTNAME_RESEMBLES_IPV4 uint16 = 10057
	ER_HOSTNAME_DOESNT_RESOLVE_TO uint16 = 10058
	ER_ADDRESSES_FOR_HOSTNAME_HEADER uint16 = 10059
	ER_ADDRESSES_FOR_HOSTNAME_LIST_ITEM uint16 = 10060
	ER_TRG_WITHOUT_DEFINER uint16 = 10061
	ER_TRG_NO_CLIENT_CHARSET uint16 = 10062
	ER_PARSING_VIEW uint16 = 10063
	ER_COMPONENTS_INFRASTRUCTURE_BOOTSTRAP uint16 = 10064
	ER_COMPONENTS_INFRASTRUCTURE_SHUTDOWN uint16 = 10065
	ER_COMPONENTS_PERSIST_LOADER_BOOTSTRAP uint16 = 10066
	ER_DEPART_WITH_GRACE uint16 = 10067
	ER_CA_SELF_SIGNED uint16 = 10068
	ER_SSL_LIBRARY_ERROR uint16 = 10069
	ER_NO_THD_NO_UUID uint16 = 10070
	ER_UUID_SALT uint16 = 10071
	ER_UUID_IS uint16 = 10072
	ER_UUID_INVALID uint16 = 10073
	ER_UUID_SCRUB uint16 = 10074
	ER_CREATING_NEW_UUID uint16 = 10075
	ER_CANT_CREATE_UUID uint16 = 10076
	ER_UNKNOWN_UNSUPPORTED_STORAGE_ENGINE uint16 = 10077
	ER_SECURE_AUTH_VALUE_UNSUPPORTED uint16 = 10078
	ER_INVALID_INSTRUMENT uint16 = 10079
	ER_INNODB_MANDATORY uint16 = 10080
	//OBSOLETE_ER_INNODB_CANNOT_BE_IGNORED uint16 = 10081
	//OBSOLETE_ER_OLD_PASSWORDS_NO_MIDDLE_GROUND uint16 = 10082
	ER_VERBOSE_REQUIRES_HELP uint16 = 10083
	ER_POINTLESS_WITHOUT_SLOWLOG uint16 = 10084
	ER_WASTEFUL_NET_BUFFER_SIZE uint16 = 10085
	ER_DEPRECATED_TIMESTAMP_IMPLICIT_DEFAULTS uint16 = 10086
	ER_FT_BOOL_SYNTAX_INVALID uint16 = 10087
	ER_CREDENTIALLESS_AUTO_USER_BAD uint16 = 10088
	ER_CONNECTION_HANDLING_OOM uint16 = 10089
	ER_THREAD_HANDLING_OOM uint16 = 10090
	ER_CANT_CREATE_TEST_FILE uint16 = 10091
	ER_CANT_CREATE_PID_FILE uint16 = 10092
	ER_CANT_REMOVE_PID_FILE uint16 = 10093
	ER_CANT_CREATE_SHUTDOWN_THREAD uint16 = 10094
	ER_SEC_FILE_PRIV_CANT_ACCESS_DIR uint16 = 10095
	ER_SEC_FILE_PRIV_IGNORED uint16 = 10096
	ER_SEC_FILE_PRIV_EMPTY uint16 = 10097
	ER_SEC_FILE_PRIV_NULL uint16 = 10098
	ER_SEC_FILE_PRIV_DIRECTORY_INSECURE uint16 = 10099
	ER_SEC_FILE_PRIV_CANT_STAT uint16 = 10100
	ER_SEC_FILE_PRIV_DIRECTORY_PERMISSIONS uint16 = 10101
	ER_SEC_FILE_PRIV_ARGUMENT_TOO_LONG uint16 = 10102
	ER_CANT_CREATE_NAMED_PIPES_THREAD uint16 = 10103
	ER_CANT_CREATE_TCPIP_THREAD uint16 = 10104
	ER_CANT_CREATE_SHM_THREAD uint16 = 10105
	ER_CANT_CREATE_INTERRUPT_THREAD uint16 = 10106
	ER_WRITABLE_CONFIG_REMOVED uint16 = 10107
	ER_CORE_VALUES uint16 = 10108
	ER_WRONG_DATETIME_SPEC uint16 = 10109
	ER_RPL_BINLOG_FILTERS_OOM uint16 = 10110
	ER_KEYCACHE_OOM uint16 = 10111
	ER_CONFIRMING_THE_FUTURE uint16 = 10112
	ER_BACK_IN_TIME uint16 = 10113
	ER_FUTURE_DATE uint16 = 10114
	ER_UNSUPPORTED_DATE uint16 = 10115
	ER_STARTING_AS uint16 = 10116
	ER_SHUTTING_DOWN_SLAVE_THREADS uint16 = 10117
	ER_DISCONNECTING_REMAINING_CLIENTS uint16 = 10118
	ER_ABORTING uint16 = 10119
	ER_BINLOG_END uint16 = 10120
	ER_CALL_ME_LOCALHOST uint16 = 10121
	ER_USER_REQUIRES_ROOT uint16 = 10122
	ER_REALLY_RUN_AS_ROOT uint16 = 10123
	ER_USER_WHAT_USER uint16 = 10124
	ER_TRANSPORTS_WHAT_TRANSPORTS uint16 = 10125
	ER_FAIL_SETGID uint16 = 10126
	ER_FAIL_SETUID uint16 = 10127
	ER_FAIL_SETREGID uint16 = 10128
	ER_FAIL_SETREUID uint16 = 10129
	ER_FAIL_CHROOT uint16 = 10130
	ER_WIN_LISTEN_BUT_HOW uint16 = 10131
	ER_NOT_RIGHT_NOW uint16 = 10132
	ER_FIXING_CLIENT_CHARSET uint16 = 10133
	ER_OOM uint16 = 10134
	ER_FAILED_TO_LOCK_MEM uint16 = 10135
	ER_MYINIT_FAILED uint16 = 10136
	ER_BEG_INITFILE uint16 = 10137
	ER_END_INITFILE uint16 = 10138
	ER_CHANGED_MAX_OPEN_FILES uint16 = 10139
	ER_CANT_INCREASE_MAX_OPEN_FILES uint16 = 10140
	ER_CHANGED_MAX_CONNECTIONS uint16 = 10141
	ER_CHANGED_TABLE_OPEN_CACHE uint16 = 10142
	ER_THE_USER_ABIDES uint16 = 10143
	ER_RPL_CANT_ADD_DO_TABLE uint16 = 10144
	ER_RPL_CANT_ADD_IGNORE_TABLE uint16 = 10145
	ER_TRACK_VARIABLES_BOGUS uint16 = 10146
	ER_EXCESS_ARGUMENTS uint16 = 10147
	ER_VERBOSE_HINT uint16 = 10148
	ER_CANT_READ_ERRMSGS uint16 = 10149
	ER_CANT_INIT_DBS uint16 = 10150
	ER_LOG_OUTPUT_CONTRADICTORY uint16 = 10151
	ER_NO_CSV_NO_LOG_TABLES uint16 = 10152
	ER_RPL_REWRITEDB_MISSING_ARROW uint16 = 10153
	ER_RPL_REWRITEDB_EMPTY_FROM uint16 = 10154
	ER_RPL_REWRITEDB_EMPTY_TO uint16 = 10155
	ER_LOG_FILES_GIVEN_LOG_OUTPUT_IS_TABLE uint16 = 10156
	ER_LOG_FILE_INVALID uint16 = 10157
	ER_LOWER_CASE_TABLE_NAMES_CS_DD_ON_CI_FS_UNSUPPORTED uint16 = 10158
	ER_LOWER_CASE_TABLE_NAMES_USING_2 uint16 = 10159
	ER_LOWER_CASE_TABLE_NAMES_USING_0 uint16 = 10160
	ER_NEED_LOG_BIN uint16 = 10161
	ER_NEED_FILE_INSTEAD_OF_DIR uint16 = 10162
	ER_LOG_BIN_BETTER_WITH_NAME uint16 = 10163
	ER_BINLOG_NEEDS_SERVERID uint16 = 10164
	ER_RPL_CANT_MAKE_PATHS uint16 = 10165
	ER_CANT_INITIALIZE_GTID uint16 = 10166
	ER_CANT_INITIALIZE_EARLY_PLUGINS uint16 = 10167
	ER_CANT_INITIALIZE_BUILTIN_PLUGINS uint16 = 10168
	ER_CANT_INITIALIZE_DYNAMIC_PLUGINS uint16 = 10169
	ER_PERFSCHEMA_INIT_FAILED uint16 = 10170
	ER_STACKSIZE_UNEXPECTED uint16 = 10171
	//OBSOLETE_ER_CANT_SET_DATADIR uint16 = 10172
	ER_CANT_STAT_DATADIR uint16 = 10173
	ER_CANT_CHOWN_DATADIR uint16 = 10174
	ER_CANT_SET_UP_PERSISTED_VALUES uint16 = 10175
	ER_CANT_SAVE_GTIDS uint16 = 10176
	ER_AUTH_CANT_SET_DEFAULT_PLUGIN uint16 = 10177
	ER_CANT_JOIN_SHUTDOWN_THREAD uint16 = 10178
	ER_CANT_HASH_DO_AND_IGNORE_RULES uint16 = 10179
	ER_CANT_OPEN_CA uint16 = 10180
	ER_CANT_ACCESS_CAPATH uint16 = 10181
	ER_SSL_TRYING_DATADIR_DEFAULTS uint16 = 10182
	ER_AUTO_OPTIONS_FAILED uint16 = 10183
	ER_CANT_INIT_TIMER uint16 = 10184
	ER_SERVERID_TOO_LARGE uint16 = 10185
	ER_DEFAULT_SE_UNAVAILABLE uint16 = 10186
	ER_CANT_OPEN_ERROR_LOG uint16 = 10187
	ER_INVALID_ERROR_LOG_NAME uint16 = 10188
	ER_RPL_INFINITY_DENIED uint16 = 10189
	ER_RPL_INFINITY_IGNORED uint16 = 10190
	//OBSOLETE_ER_NDB_TABLES_NOT_READY uint16 = 10191
	ER_TABLE_CHECK_INTACT uint16 = 10192
	ER_DD_TABLESPACE_NOT_FOUND uint16 = 10193
	ER_DD_TRG_CONNECTION_COLLATION_MISSING uint16 = 10194
	ER_DD_TRG_DB_COLLATION_MISSING uint16 = 10195
	ER_DD_TRG_DEFINER_OOM uint16 = 10196
	ER_DD_TRG_FILE_UNREADABLE uint16 = 10197
	ER_TRG_CANT_PARSE uint16 = 10198
	ER_DD_TRG_CANT_ADD uint16 = 10199
	ER_DD_CANT_RESOLVE_VIEW uint16 = 10200
	ER_DD_VIEW_WITHOUT_DEFINER uint16 = 10201
	ER_PLUGIN_INIT_FAILED uint16 = 10202
	ER_RPL_TRX_DELEGATES_INIT_FAILED uint16 = 10203
	ER_RPL_BINLOG_STORAGE_DELEGATES_INIT_FAILED uint16 = 10204
	ER_RPL_BINLOG_TRANSMIT_DELEGATES_INIT_FAILED uint16 = 10205
	ER_RPL_BINLOG_RELAY_DELEGATES_INIT_FAILED uint16 = 10206
	ER_RPL_PLUGIN_FUNCTION_FAILED uint16 = 10207
	ER_SQL_HA_READ_FAILED uint16 = 10208
	ER_SR_BOGUS_VALUE uint16 = 10209
	ER_SR_INVALID_CONTEXT uint16 = 10210
	ER_READING_TABLE_FAILED uint16 = 10211
	ER_DES_FILE_WRONG_KEY uint16 = 10212
	ER_CANT_SET_PERSISTED uint16 = 10213
	ER_JSON_PARSE_ERROR uint16 = 10214
	ER_CONFIG_OPTION_WITHOUT_GROUP uint16 = 10215
	ER_VALGRIND_DO_QUICK_LEAK_CHECK uint16 = 10216
	ER_VALGRIND_COUNT_LEAKS uint16 = 10217
	ER_LOAD_DATA_INFILE_FAILED_IN_UNEXPECTED_WAY uint16 = 10218
	ER_UNKNOWN_ERROR_NUMBER uint16 = 10219
	ER_UDF_CANT_ALLOC_FOR_STRUCTURES uint16 = 10220
	ER_UDF_CANT_ALLOC_FOR_FUNCTION uint16 = 10221
	ER_UDF_INVALID_ROW_IN_FUNCTION_TABLE uint16 = 10222
	ER_UDF_CANT_OPEN_FUNCTION_TABLE uint16 = 10223
	ER_XA_RECOVER_FOUND_TRX_IN_SE uint16 = 10224
	ER_XA_RECOVER_FOUND_XA_TRX uint16 = 10225
	ER_XA_IGNORING_XID uint16 = 10226
	ER_XA_COMMITTING_XID uint16 = 10227
	ER_XA_ROLLING_BACK_XID uint16 = 10228
	ER_XA_STARTING_RECOVERY uint16 = 10229
	ER_XA_NO_MULTI_2PC_HEURISTIC_RECOVER uint16 = 10230
	ER_XA_RECOVER_EXPLANATION uint16 = 10231
	ER_XA_RECOVERY_DONE uint16 = 10232
	ER_TRX_GTID_COLLECT_REJECT uint16 = 10233
	ER_SQL_AUTHOR_DEFAULT_ROLES_FAIL uint16 = 10234
	ER_SQL_USER_TABLE_CREATE_WARNING uint16 = 10235
	ER_SQL_USER_TABLE_ALTER_WARNING uint16 = 10236
	ER_ROW_IN_WRONG_PARTITION_PLEASE_REPAIR uint16 = 10237
	ER_MYISAM_CRASHED_ERROR_IN_THREAD uint16 = 10238
	ER_MYISAM_CRASHED_ERROR_IN uint16 = 10239
	ER_TOO_MANY_STORAGE_ENGINES uint16 = 10240
	ER_SE_TYPECODE_CONFLICT uint16 = 10241
	ER_TRX_WRITE_SET_OOM uint16 = 10242
	ER_HANDLERTON_OOM uint16 = 10243
	ER_CONN_SHM_LISTENER uint16 = 10244
	ER_CONN_SHM_CANT_CREATE_SERVICE uint16 = 10245
	ER_CONN_SHM_CANT_CREATE_CONNECTION uint16 = 10246
	ER_CONN_PIP_CANT_CREATE_EVENT uint16 = 10247
	ER_CONN_PIP_CANT_CREATE_PIPE uint16 = 10248
	ER_CONN_PER_THREAD_NO_THREAD uint16 = 10249
	ER_CONN_TCP_NO_SOCKET uint16 = 10250
	ER_CONN_TCP_CREATED uint16 = 10251
	ER_CONN_TCP_ADDRESS uint16 = 10252
	ER_CONN_TCP_IPV6_AVAILABLE uint16 = 10253
	ER_CONN_TCP_IPV6_UNAVAILABLE uint16 = 10254
	ER_CONN_TCP_ERROR_WITH_STRERROR uint16 = 10255
	ER_CONN_TCP_CANT_RESOLVE_HOSTNAME uint16 = 10256
	ER_CONN_TCP_IS_THERE_ANOTHER_USING_PORT uint16 = 10257
	ER_CONN_UNIX_IS_THERE_ANOTHER_USING_SOCKET uint16 = 10258
	ER_CONN_UNIX_PID_CLAIMED_SOCKET_FILE uint16 = 10259
	ER_CONN_TCP_CANT_RESET_V6ONLY uint16 = 10260
	ER_CONN_TCP_BIND_RETRY uint16 = 10261
	ER_CONN_TCP_BIND_FAIL uint16 = 10262
	ER_CONN_TCP_IP_NOT_LOGGED uint16 = 10263
	ER_CONN_TCP_RESOLVE_INFO uint16 = 10264
	ER_CONN_TCP_START_FAIL uint16 = 10265
	ER_CONN_TCP_LISTEN_FAIL uint16 = 10266
	ER_CONN_UNIX_PATH_TOO_LONG uint16 = 10267
	ER_CONN_UNIX_LOCK_FILE_FAIL uint16 = 10268
	ER_CONN_UNIX_NO_FD uint16 = 10269
	ER_CONN_UNIX_NO_BIND_NO_START uint16 = 10270
	ER_CONN_UNIX_LISTEN_FAILED uint16 = 10271
	ER_CONN_UNIX_LOCK_FILE_GIVING_UP uint16 = 10272
	ER_CONN_UNIX_LOCK_FILE_CANT_CREATE uint16 = 10273
	ER_CONN_UNIX_LOCK_FILE_CANT_OPEN uint16 = 10274
	ER_CONN_UNIX_LOCK_FILE_CANT_READ uint16 = 10275
	ER_CONN_UNIX_LOCK_FILE_EMPTY uint16 = 10276
	ER_CONN_UNIX_LOCK_FILE_PIDLESS uint16 = 10277
	ER_CONN_UNIX_LOCK_FILE_CANT_WRITE uint16 = 10278
	ER_CONN_UNIX_LOCK_FILE_CANT_DELETE uint16 = 10279
	ER_CONN_UNIX_LOCK_FILE_CANT_SYNC uint16 = 10280
	ER_CONN_UNIX_LOCK_FILE_CANT_CLOSE uint16 = 10281
	ER_CONN_SOCKET_SELECT_FAILED uint16 = 10282
	ER_CONN_SOCKET_ACCEPT_FAILED uint16 = 10283
	ER_AUTH_RSA_CANT_FIND uint16 = 10284
	ER_AUTH_RSA_CANT_PARSE uint16 = 10285
	ER_AUTH_RSA_CANT_READ uint16 = 10286
	ER_AUTH_RSA_FILES_NOT_FOUND uint16 = 10287
	ER_CONN_ATTR_TRUNCATED uint16 = 10288
	ER_X509_CIPHERS_MISMATCH uint16 = 10289
	ER_X509_ISSUER_MISMATCH uint16 = 10290
	ER_X509_SUBJECT_MISMATCH uint16 = 10291
	ER_AUTH_CANT_ACTIVATE_ROLE uint16 = 10292
	ER_X509_NEEDS_RSA_PRIVKEY uint16 = 10293
	ER_X509_CANT_WRITE_KEY uint16 = 10294
	ER_X509_CANT_CHMOD_KEY uint16 = 10295
	ER_X509_CANT_READ_CA_KEY uint16 = 10296
	ER_X509_CANT_READ_CA_CERT uint16 = 10297
	ER_X509_CANT_CREATE_CERT uint16 = 10298
	ER_X509_CANT_WRITE_CERT uint16 = 10299
	ER_AUTH_CANT_CREATE_RSA_PAIR uint16 = 10300
	ER_AUTH_CANT_WRITE_PRIVKEY uint16 = 10301
	ER_AUTH_CANT_WRITE_PUBKEY uint16 = 10302
	ER_AUTH_SSL_CONF_PREVENTS_CERT_GENERATION uint16 = 10303
	ER_AUTH_USING_EXISTING_CERTS uint16 = 10304
	ER_AUTH_CERTS_SAVED_TO_DATADIR uint16 = 10305
	ER_AUTH_CERT_GENERATION_DISABLED uint16 = 10306
	ER_AUTH_RSA_CONF_PREVENTS_KEY_GENERATION uint16 = 10307
	ER_AUTH_KEY_GENERATION_SKIPPED_PAIR_PRESENT uint16 = 10308
	ER_AUTH_KEYS_SAVED_TO_DATADIR uint16 = 10309
	ER_AUTH_KEY_GENERATION_DISABLED uint16 = 10310
	ER_AUTHCACHE_PROXIES_PRIV_SKIPPED_NEEDS_RESOLVE uint16 = 10311
	ER_AUTHCACHE_PLUGIN_MISSING uint16 = 10312
	ER_AUTHCACHE_PLUGIN_CONFIG uint16 = 10313
	//OBSOLETE_ER_AUTHCACHE_ROLE_TABLES_DODGY uint16 = 10314
	ER_AUTHCACHE_USER_SKIPPED_NEEDS_RESOLVE uint16 = 10315
	ER_AUTHCACHE_USER_TABLE_DODGY uint16 = 10316
	ER_AUTHCACHE_USER_IGNORED_DEPRECATED_PASSWORD uint16 = 10317
	ER_AUTHCACHE_USER_IGNORED_NEEDS_PLUGIN uint16 = 10318
	ER_AUTHCACHE_USER_IGNORED_INVALID_PASSWORD uint16 = 10319
	ER_AUTHCACHE_EXPIRED_PASSWORD_UNSUPPORTED uint16 = 10320
	ER_NO_SUPER_WITHOUT_USER_PLUGIN uint16 = 10321
	ER_AUTHCACHE_DB_IGNORED_EMPTY_NAME uint16 = 10322
	ER_AUTHCACHE_DB_SKIPPED_NEEDS_RESOLVE uint16 = 10323
	ER_AUTHCACHE_DB_ENTRY_LOWERCASED_REVOKE_WILL_FAIL uint16 = 10324
	ER_AUTHCACHE_TABLE_PROXIES_PRIV_MISSING uint16 = 10325
	ER_AUTHCACHE_CANT_OPEN_AND_LOCK_PRIVILEGE_TABLES uint16 = 10326
	ER_AUTHCACHE_CANT_INIT_GRANT_SUBSYSTEM uint16 = 10327
	ER_AUTHCACHE_PROCS_PRIV_SKIPPED_NEEDS_RESOLVE uint16 = 10328
	ER_AUTHCACHE_PROCS_PRIV_ENTRY_IGNORED_BAD_ROUTINE_TYPE uint16 = 10329
	ER_AUTHCACHE_TABLES_PRIV_SKIPPED_NEEDS_RESOLVE uint16 = 10330
	ER_USER_NOT_IN_EXTRA_USERS_BINLOG_POSSIBLY_INCOMPLETE uint16 = 10331
	ER_DD_SCHEMA_NOT_FOUND uint16 = 10332
	ER_DD_TABLE_NOT_FOUND uint16 = 10333
	ER_DD_SE_INIT_FAILED uint16 = 10334
	ER_DD_ABORTING_PARTIAL_UPGRADE uint16 = 10335
	ER_DD_FRM_EXISTS_FOR_TABLE uint16 = 10336
	ER_DD_CREATED_FOR_UPGRADE uint16 = 10337
	ER_ERRMSG_CANT_FIND_FILE uint16 = 10338
	ER_ERRMSG_LOADING_55_STYLE uint16 = 10339
	ER_ERRMSG_MISSING_IN_FILE uint16 = 10340
	ER_ERRMSG_OOM uint16 = 10341
	ER_ERRMSG_CANT_READ uint16 = 10342
	ER_TABLE_INCOMPATIBLE_DECIMAL_FIELD uint16 = 10343
	ER_TABLE_INCOMPATIBLE_YEAR_FIELD uint16 = 10344
	ER_INVALID_CHARSET_AND_DEFAULT_IS_MB uint16 = 10345
	ER_TABLE_WRONG_KEY_DEFINITION uint16 = 10346
	ER_CANT_OPEN_FRM_FILE uint16 = 10347
	ER_CANT_READ_FRM_FILE uint16 = 10348
	ER_TABLE_CREATED_WITH_DIFFERENT_VERSION uint16 = 10349
	ER_VIEW_UNPARSABLE uint16 = 10350
	ER_FILE_TYPE_UNKNOWN uint16 = 10351
	ER_INVALID_INFO_IN_FRM uint16 = 10352
	ER_CANT_OPEN_AND_LOCK_PRIVILEGE_TABLES uint16 = 10353
	ER_AUDIT_PLUGIN_DOES_NOT_SUPPORT_AUDIT_AUTH_EVENTS uint16 = 10354
	ER_AUDIT_PLUGIN_HAS_INVALID_DATA uint16 = 10355
	ER_TZ_OOM_INITIALIZING_TIME_ZONES uint16 = 10356
	ER_TZ_CANT_OPEN_AND_LOCK_TIME_ZONE_TABLE uint16 = 10357
	ER_TZ_OOM_LOADING_LEAP_SECOND_TABLE uint16 = 10358
	ER_TZ_TOO_MANY_LEAPS_IN_LEAP_SECOND_TABLE uint16 = 10359
	ER_TZ_ERROR_LOADING_LEAP_SECOND_TABLE uint16 = 10360
	ER_TZ_UNKNOWN_OR_ILLEGAL_DEFAULT_TIME_ZONE uint16 = 10361
	ER_TZ_CANT_FIND_DESCRIPTION_FOR_TIME_ZONE uint16 = 10362
	ER_TZ_CANT_FIND_DESCRIPTION_FOR_TIME_ZONE_ID uint16 = 10363
	ER_TZ_TRANSITION_TYPE_TABLE_TYPE_TOO_LARGE uint16 = 10364
	ER_TZ_TRANSITION_TYPE_TABLE_ABBREVIATIONS_EXCEED_SPACE uint16 = 10365
	ER_TZ_TRANSITION_TYPE_TABLE_LOAD_ERROR uint16 = 10366
	ER_TZ_TRANSITION_TABLE_TOO_MANY_TRANSITIONS uint16 = 10367
	ER_TZ_TRANSITION_TABLE_BAD_TRANSITION_TYPE uint16 = 10368
	ER_TZ_TRANSITION_TABLE_LOAD_ERROR uint16 = 10369
	ER_TZ_NO_TRANSITION_TYPES_IN_TIME_ZONE uint16 = 10370
	ER_TZ_OOM_LOADING_TIME_ZONE_DESCRIPTION uint16 = 10371
	ER_TZ_CANT_BUILD_MKTIME_MAP uint16 = 10372
	ER_TZ_OOM_WHILE_LOADING_TIME_ZONE uint16 = 10373
	ER_TZ_OOM_WHILE_SETTING_TIME_ZONE uint16 = 10374
	ER_SLAVE_SQL_THREAD_STOPPED_UNTIL_CONDITION_BAD uint16 = 10375
	ER_SLAVE_SQL_THREAD_STOPPED_UNTIL_POSITION_REACHED uint16 = 10376
	ER_SLAVE_SQL_THREAD_STOPPED_BEFORE_GTIDS_ALREADY_APPLIED uint16 = 10377
	ER_SLAVE_SQL_THREAD_STOPPED_BEFORE_GTIDS_REACHED uint16 = 10378
	ER_SLAVE_SQL_THREAD_STOPPED_AFTER_GTIDS_REACHED uint16 = 10379
	ER_SLAVE_SQL_THREAD_STOPPED_GAP_TRX_PROCESSED uint16 = 10380
	ER_GROUP_REPLICATION_PLUGIN_NOT_INSTALLED uint16 = 10381
	ER_GTID_ALREADY_ADDED_BY_USER uint16 = 10382
	ER_FAILED_TO_DELETE_FROM_GTID_EXECUTED_TABLE uint16 = 10383
	ER_FAILED_TO_COMPRESS_GTID_EXECUTED_TABLE uint16 = 10384
	ER_FAILED_TO_COMPRESS_GTID_EXECUTED_TABLE_OOM uint16 = 10385
	ER_FAILED_TO_INIT_THREAD_ATTR_FOR_GTID_TABLE_COMPRESSION uint16 = 10386
	ER_FAILED_TO_CREATE_GTID_TABLE_COMPRESSION_THREAD uint16 = 10387
	ER_FAILED_TO_JOIN_GTID_TABLE_COMPRESSION_THREAD uint16 = 10388
	ER_NPIPE_FAILED_TO_INIT_SECURITY_DESCRIPTOR uint16 = 10389
	ER_NPIPE_FAILED_TO_SET_SECURITY_DESCRIPTOR uint16 = 10390
	ER_NPIPE_PIPE_ALREADY_IN_USE uint16 = 10391
	//OBSOLETE_ER_NDB_SLAVE_SAW_EPOCH_LOWER_THAN_PREVIOUS_ON_START uint16 = 10392
	//OBSOLETE_ER_NDB_SLAVE_SAW_EPOCH_LOWER_THAN_PREVIOUS uint16 = 10393
	//OBSOLETE_ER_NDB_SLAVE_SAW_ALREADY_COMMITTED_EPOCH uint16 = 10394
	//OBSOLETE_ER_NDB_SLAVE_PREVIOUS_EPOCH_NOT_COMMITTED uint16 = 10395
	//OBSOLETE_ER_NDB_SLAVE_MISSING_DATA_FOR_TIMESTAMP_COLUMN uint16 = 10396
	//OBSOLETE_ER_NDB_SLAVE_LOGGING_EXCEPTIONS_TO uint16 = 10397
	//OBSOLETE_ER_NDB_SLAVE_LOW_EPOCH_RESOLUTION uint16 = 10398
	//OBSOLETE_ER_NDB_INFO_FOUND_UNEXPECTED_FIELD_TYPE uint16 = 10399
	//OBSOLETE_ER_NDB_INFO_FAILED_TO_CREATE_NDBINFO uint16 = 10400
	//OBSOLETE_ER_NDB_INFO_FAILED_TO_INIT_NDBINFO uint16 = 10401
	//OBSOLETE_ER_NDB_CLUSTER_WRONG_NUMBER_OF_FUNCTION_ARGUMENTS uint16 = 10402
	//OBSOLETE_ER_NDB_CLUSTER_SCHEMA_INFO uint16 = 10403
	//OBSOLETE_ER_NDB_CLUSTER_GENERIC_MESSAGE uint16 = 10404
	ER_RPL_CANT_OPEN_INFO_TABLE uint16 = 10405
	ER_RPL_CANT_SCAN_INFO_TABLE uint16 = 10406
	ER_RPL_CORRUPTED_INFO_TABLE uint16 = 10407
	ER_RPL_CORRUPTED_KEYS_IN_INFO_TABLE uint16 = 10408
	ER_RPL_WORKER_ID_IS uint16 = 10409
	ER_RPL_INCONSISTENT_TIMESTAMPS_IN_TRX uint16 = 10410
	ER_RPL_INCONSISTENT_SEQUENCE_NO_IN_TRX uint16 = 10411
	ER_RPL_CHANNELS_REQUIRE_TABLES_AS_INFO_REPOSITORIES uint16 = 10412
	ER_RPL_CHANNELS_REQUIRE_NON_ZERO_SERVER_ID uint16 = 10413
	ER_RPL_REPO_SHOULD_BE_TABLE uint16 = 10414
	ER_RPL_ERROR_CREATING_MASTER_INFO uint16 = 10415
	ER_RPL_ERROR_CHANGING_MASTER_INFO_REPO_TYPE uint16 = 10416
	ER_RPL_CHANGING_RELAY_LOG_INFO_REPO_TYPE_FAILED_DUE_TO_GAPS uint16 = 10417
	ER_RPL_ERROR_CREATING_RELAY_LOG_INFO uint16 = 10418
	ER_RPL_ERROR_CHANGING_RELAY_LOG_INFO_REPO_TYPE uint16 = 10419
	ER_RPL_FAILED_TO_DELETE_FROM_SLAVE_WORKERS_INFO_REPOSITORY uint16 = 10420
	ER_RPL_FAILED_TO_RESET_STATE_IN_SLAVE_INFO_REPOSITORY uint16 = 10421
	ER_RPL_ERROR_CHECKING_REPOSITORY uint16 = 10422
	ER_RPL_SLAVE_GENERIC_MESSAGE uint16 = 10423
	ER_RPL_SLAVE_COULD_NOT_CREATE_CHANNEL_LIST uint16 = 10424
	ER_RPL_MULTISOURCE_REQUIRES_TABLE_TYPE_REPOSITORIES uint16 = 10425
	ER_RPL_SLAVE_FAILED_TO_INIT_A_MASTER_INFO_STRUCTURE uint16 = 10426
	ER_RPL_SLAVE_FAILED_TO_INIT_MASTER_INFO_STRUCTURE uint16 = 10427
	ER_RPL_SLAVE_FAILED_TO_CREATE_CHANNEL_FROM_MASTER_INFO uint16 = 10428
	ER_RPL_FAILED_TO_CREATE_NEW_INFO_FILE uint16 = 10429
	ER_RPL_FAILED_TO_CREATE_CACHE_FOR_INFO_FILE uint16 = 10430
	ER_RPL_FAILED_TO_OPEN_INFO_FILE uint16 = 10431
	ER_RPL_GTID_MEMORY_FINALLY_AVAILABLE uint16 = 10432
	ER_SERVER_COST_UNKNOWN_COST_CONSTANT uint16 = 10433
	ER_SERVER_COST_INVALID_COST_CONSTANT uint16 = 10434
	ER_ENGINE_COST_UNKNOWN_COST_CONSTANT uint16 = 10435
	ER_ENGINE_COST_UNKNOWN_STORAGE_ENGINE uint16 = 10436
	ER_ENGINE_COST_INVALID_DEVICE_TYPE_FOR_SE uint16 = 10437
	ER_ENGINE_COST_INVALID_CONST_CONSTANT_FOR_SE_AND_DEVICE uint16 = 10438
	ER_SERVER_COST_FAILED_TO_READ uint16 = 10439
	ER_ENGINE_COST_FAILED_TO_READ uint16 = 10440
	ER_FAILED_TO_OPEN_COST_CONSTANT_TABLES uint16 = 10441
	ER_RPL_UNSUPPORTED_UNIGNORABLE_EVENT_IN_STREAM uint16 = 10442
	ER_RPL_GTID_LOG_EVENT_IN_STREAM uint16 = 10443
	ER_RPL_UNEXPECTED_BEGIN_IN_STREAM uint16 = 10444
	ER_RPL_UNEXPECTED_COMMIT_ROLLBACK_OR_XID_LOG_EVENT_IN_STREAM uint16 = 10445
	ER_RPL_UNEXPECTED_XA_ROLLBACK_IN_STREAM uint16 = 10446
	ER_EVENT_EXECUTION_FAILED_CANT_AUTHENTICATE_USER uint16 = 10447
	ER_EVENT_EXECUTION_FAILED_USER_LOST_EVEN_PRIVILEGE uint16 = 10448
	ER_EVENT_ERROR_DURING_COMPILATION uint16 = 10449
	ER_EVENT_DROPPING uint16 = 10450
	//OBSOLETE_ER_NDB_SCHEMA_GENERIC_MESSAGE uint16 = 10451
	ER_RPL_INCOMPATIBLE_DECIMAL_IN_RBR uint16 = 10452
	ER_INIT_ROOT_WITHOUT_PASSWORD uint16 = 10453
	ER_INIT_GENERATING_TEMP_PASSWORD_FOR_ROOT uint16 = 10454
	ER_INIT_CANT_OPEN_BOOTSTRAP_FILE uint16 = 10455
	ER_INIT_BOOTSTRAP_COMPLETE uint16 = 10456
	ER_INIT_DATADIR_NOT_EMPTY_WONT_INITIALIZE uint16 = 10457
	ER_INIT_DATADIR_EXISTS_WONT_INITIALIZE uint16 = 10458
	ER_INIT_DATADIR_EXISTS_AND_PATH_TOO_LONG_WONT_INITIALIZE uint16 = 10459
	ER_INIT_DATADIR_EXISTS_AND_NOT_WRITABLE_WONT_INITIALIZE uint16 = 10460
	ER_INIT_CREATING_DD uint16 = 10461
	ER_RPL_BINLOG_STARTING_DUMP uint16 = 10462
	ER_RPL_BINLOG_MASTER_SENDS_HEARTBEAT uint16 = 10463
	ER_RPL_BINLOG_SKIPPING_REMAINING_HEARTBEAT_INFO uint16 = 10464
	ER_RPL_BINLOG_MASTER_USES_CHECKSUM_AND_SLAVE_CANT uint16 = 10465
	//OBSOLETE_ER_NDB_QUERY_FAILED uint16 = 10466
	ER_KILLING_THREAD uint16 = 10467
	ER_DETACHING_SESSION_LEFT_BY_PLUGIN uint16 = 10468
	ER_CANT_DETACH_SESSION_LEFT_BY_PLUGIN uint16 = 10469
	ER_DETACHED_SESSIONS_LEFT_BY_PLUGIN uint16 = 10470
	ER_FAILED_TO_DECREMENT_NUMBER_OF_THREADS uint16 = 10471
	ER_PLUGIN_DID_NOT_DEINITIALIZE_THREADS uint16 = 10472
	ER_KILLED_THREADS_OF_PLUGIN uint16 = 10473
	//OBSOLETE_ER_NDB_SLAVE_MAX_REPLICATED_EPOCH_UNKNOWN uint16 = 10474
	//OBSOLETE_ER_NDB_SLAVE_MAX_REPLICATED_EPOCH_SET_TO uint16 = 10475
	//OBSOLETE_ER_NDB_NODE_ID_AND_MANAGEMENT_SERVER_INFO uint16 = 10476
	//OBSOLETE_ER_NDB_DISCONNECT_INFO uint16 = 10477
	//OBSOLETE_ER_NDB_COLUMN_DEFAULTS_DIFFER uint16 = 10478
	//OBSOLETE_ER_NDB_COLUMN_SHOULD_NOT_HAVE_NATIVE_DEFAULT uint16 = 10479
	//OBSOLETE_ER_NDB_FIELD_INFO uint16 = 10480
	//OBSOLETE_ER_NDB_COLUMN_INFO uint16 = 10481
	//OBSOLETE_ER_NDB_OOM_IN_FIX_UNIQUE_INDEX_ATTR_ORDER uint16 = 10482
	//OBSOLETE_ER_NDB_SLAVE_MALFORMED_EVENT_RECEIVED_ON_TABLE uint16 = 10483
	//OBSOLETE_ER_NDB_SLAVE_CONFLICT_FUNCTION_REQUIRES_ROLE uint16 = 10484
	//OBSOLETE_ER_NDB_SLAVE_CONFLICT_TRANSACTION_IDS uint16 = 10485
	//OBSOLETE_ER_NDB_SLAVE_BINLOG_MISSING_INFO_FOR_CONFLICT_DETECTION uint16 = 10486
	//OBSOLETE_ER_NDB_ERROR_IN_READAUTOINCREMENTVALUE uint16 = 10487
	//OBSOLETE_ER_NDB_FOUND_UNCOMMITTED_AUTOCOMMIT uint16 = 10488
	//OBSOLETE_ER_NDB_SLAVE_TOO_MANY_RETRIES uint16 = 10489
	//OBSOLETE_ER_NDB_SLAVE_ERROR_IN_UPDATE_CREATE_INFO uint16 = 10490
	//OBSOLETE_ER_NDB_SLAVE_CANT_ALLOCATE_TABLE_SHARE uint16 = 10491
	//OBSOLETE_ER_NDB_BINLOG_ERROR_INFO_FROM_DA uint16 = 10492
	//OBSOLETE_ER_NDB_BINLOG_CREATE_TABLE_EVENT uint16 = 10493
	//OBSOLETE_ER_NDB_BINLOG_FAILED_CREATE_TABLE_EVENT_OPERATIONS uint16 = 10494
	//OBSOLETE_ER_NDB_BINLOG_RENAME_EVENT uint16 = 10495
	//OBSOLETE_ER_NDB_BINLOG_FAILED_CREATE_DURING_RENAME uint16 = 10496
	//OBSOLETE_ER_NDB_UNEXPECTED_RENAME_TYPE uint16 = 10497
	//OBSOLETE_ER_NDB_ERROR_IN_GET_AUTO_INCREMENT uint16 = 10498
	//OBSOLETE_ER_NDB_CREATING_SHARE_IN_OPEN uint16 = 10499
	//OBSOLETE_ER_NDB_TABLE_OPENED_READ_ONLY uint16 = 10500
	//OBSOLETE_ER_NDB_INITIALIZE_GIVEN_CLUSTER_PLUGIN_DISABLED uint16 = 10501
	//OBSOLETE_ER_NDB_BINLOG_FORMAT_CHANGED_FROM_STMT_TO_MIXED uint16 = 10502
	//OBSOLETE_ER_NDB_TRAILING_SHARE_RELEASED_BY_CLOSE_CACHED_TABLES uint16 = 10503
	//OBSOLETE_ER_NDB_SHARE_ALREADY_EXISTS uint16 = 10504
	//OBSOLETE_ER_NDB_HANDLE_TRAILING_SHARE_INFO uint16 = 10505
	//OBSOLETE_ER_NDB_CLUSTER_GET_SHARE_INFO uint16 = 10506
	//OBSOLETE_ER_NDB_CLUSTER_REAL_FREE_SHARE_INFO uint16 = 10507
	//OBSOLETE_ER_NDB_CLUSTER_REAL_FREE_SHARE_DROP_FAILED uint16 = 10508
	//OBSOLETE_ER_NDB_CLUSTER_FREE_SHARE_INFO uint16 = 10509
	//OBSOLETE_ER_NDB_CLUSTER_MARK_SHARE_DROPPED_INFO uint16 = 10510
	//OBSOLETE_ER_NDB_CLUSTER_MARK_SHARE_DROPPED_DESTROYING_SHARE uint16 = 10511
	//OBSOLETE_ER_NDB_CLUSTER_OOM_THD_NDB uint16 = 10512
	//OBSOLETE_ER_NDB_BINLOG_NDB_TABLES_INITIALLY_READ_ONLY uint16 = 10513
	//OBSOLETE_ER_NDB_UTIL_THREAD_OOM uint16 = 10514
	//OBSOLETE_ER_NDB_ILLEGAL_VALUE_FOR_NDB_RECV_THREAD_CPU_MASK uint16 = 10515
	//OBSOLETE_ER_NDB_TOO_MANY_CPUS_IN_NDB_RECV_THREAD_CPU_MASK uint16 = 10516
	ER_DBUG_CHECK_SHARES_OPEN uint16 = 10517
	ER_DBUG_CHECK_SHARES_INFO uint16 = 10518
	ER_DBUG_CHECK_SHARES_DROPPED uint16 = 10519
	ER_INVALID_OR_OLD_TABLE_OR_DB_NAME uint16 = 10520
	ER_TC_RECOVERING_AFTER_CRASH_USING uint16 = 10521
	ER_TC_CANT_AUTO_RECOVER_WITH_TC_HEURISTIC_RECOVER uint16 = 10522
	ER_TC_BAD_MAGIC_IN_TC_LOG uint16 = 10523
	ER_TC_NEED_N_SE_SUPPORTING_2PC_FOR_RECOVERY uint16 = 10524
	ER_TC_RECOVERY_FAILED_THESE_ARE_YOUR_OPTIONS uint16 = 10525
	ER_TC_HEURISTIC_RECOVERY_MODE uint16 = 10526
	ER_TC_HEURISTIC_RECOVERY_FAILED uint16 = 10527
	ER_TC_RESTART_WITHOUT_TC_HEURISTIC_RECOVER uint16 = 10528
	ER_RPL_SLAVE_FAILED_TO_CREATE_OR_RECOVER_INFO_REPOSITORIES uint16 = 10529
	ER_RPL_SLAVE_AUTO_POSITION_IS_1_AND_GTID_MODE_IS_OFF uint16 = 10530
	ER_RPL_SLAVE_CANT_START_SLAVE_FOR_CHANNEL uint16 = 10531
	ER_RPL_SLAVE_CANT_STOP_SLAVE_FOR_CHANNEL uint16 = 10532
	ER_RPL_RECOVERY_NO_ROTATE_EVENT_FROM_MASTER uint16 = 10533
	ER_RPL_RECOVERY_ERROR_READ_RELAY_LOG uint16 = 10534
	//OBSOLETE_ER_RPL_RECOVERY_ERROR_FREEING_IO_CACHE uint16 = 10535
	ER_RPL_RECOVERY_SKIPPED_GROUP_REPLICATION_CHANNEL uint16 = 10536
	ER_RPL_RECOVERY_ERROR uint16 = 10537
	ER_RPL_RECOVERY_IO_ERROR_READING_RELAY_LOG_INDEX uint16 = 10538
	ER_RPL_RECOVERY_FILE_MASTER_POS_INFO uint16 = 10539
	ER_RPL_RECOVERY_REPLICATE_SAME_SERVER_ID_REQUIRES_POSITION uint16 = 10540
	ER_RPL_MTS_RECOVERY_STARTING_COORDINATOR uint16 = 10541
	ER_RPL_MTS_RECOVERY_FAILED_TO_START_COORDINATOR uint16 = 10542
	ER_RPL_MTS_AUTOMATIC_RECOVERY_FAILED uint16 = 10543
	ER_RPL_MTS_RECOVERY_CANT_OPEN_RELAY_LOG uint16 = 10544
	ER_RPL_MTS_RECOVERY_SUCCESSFUL uint16 = 10545
	ER_RPL_SERVER_ID_MISSING uint16 = 10546
	ER_RPL_CANT_CREATE_SLAVE_THREAD uint16 = 10547
	ER_RPL_SLAVE_IO_THREAD_WAS_KILLED uint16 = 10548
	ER_RPL_SLAVE_MASTER_UUID_HAS_CHANGED uint16 = 10549
	ER_RPL_SLAVE_USES_CHECKSUM_AND_MASTER_PRE_50 uint16 = 10550
	ER_RPL_SLAVE_SECONDS_BEHIND_MASTER_DUBIOUS uint16 = 10551
	ER_RPL_SLAVE_CANT_FLUSH_MASTER_INFO_FILE uint16 = 10552
	ER_RPL_SLAVE_REPORT_HOST_TOO_LONG uint16 = 10553
	ER_RPL_SLAVE_REPORT_USER_TOO_LONG uint16 = 10554
	ER_RPL_SLAVE_REPORT_PASSWORD_TOO_LONG uint16 = 10555
	ER_RPL_SLAVE_ERROR_RETRYING uint16 = 10556
	ER_RPL_SLAVE_ERROR_READING_FROM_SERVER uint16 = 10557
	ER_RPL_SLAVE_DUMP_THREAD_KILLED_BY_MASTER uint16 = 10558
	ER_RPL_MTS_STATISTICS uint16 = 10559
	ER_RPL_MTS_RECOVERY_COMPLETE uint16 = 10560
	ER_RPL_SLAVE_CANT_INIT_RELAY_LOG_POSITION uint16 = 10561
	ER_RPL_SLAVE_CONNECTED_TO_MASTER_REPLICATION_STARTED uint16 = 10562
	ER_RPL_SLAVE_IO_THREAD_KILLED uint16 = 10563
	ER_RPL_SLAVE_IO_THREAD_CANT_REGISTER_ON_MASTER uint16 = 10564
	ER_RPL_SLAVE_FORCING_TO_RECONNECT_IO_THREAD uint16 = 10565
	ER_RPL_SLAVE_ERROR_REQUESTING_BINLOG_DUMP uint16 = 10566
	ER_RPL_LOG_ENTRY_EXCEEDS_SLAVE_MAX_ALLOWED_PACKET uint16 = 10567
	ER_RPL_SLAVE_STOPPING_AS_MASTER_OOM uint16 = 10568
	ER_RPL_SLAVE_IO_THREAD_ABORTED_WAITING_FOR_RELAY_LOG_SPACE uint16 = 10569
	ER_RPL_SLAVE_IO_THREAD_EXITING uint16 = 10570
	ER_RPL_SLAVE_CANT_INITIALIZE_SLAVE_WORKER uint16 = 10571
	ER_RPL_MTS_GROUP_RECOVERY_RELAY_LOG_INFO_FOR_WORKER uint16 = 10572
	ER_RPL_ERROR_LOOKING_FOR_LOG uint16 = 10573
	ER_RPL_MTS_GROUP_RECOVERY_RELAY_LOG_INFO uint16 = 10574
	ER_RPL_CANT_FIND_FOLLOWUP_FILE uint16 = 10575
	ER_RPL_MTS_CHECKPOINT_PERIOD_DIFFERS_FROM_CNT uint16 = 10576
	ER_RPL_SLAVE_WORKER_THREAD_CREATION_FAILED uint16 = 10577
	ER_RPL_SLAVE_WORKER_THREAD_CREATION_FAILED_WITH_ERRNO uint16 = 10578
	ER_RPL_SLAVE_FAILED_TO_INIT_PARTITIONS_HASH uint16 = 10579
	//OBSOLETE_ER_RPL_SLAVE_NDB_TABLES_NOT_AVAILABLE uint16 = 10580
	ER_RPL_SLAVE_SQL_THREAD_STARTING uint16 = 10581
	ER_RPL_SLAVE_SKIP_COUNTER_EXECUTED uint16 = 10582
	ER_RPL_SLAVE_ADDITIONAL_ERROR_INFO_FROM_DA uint16 = 10583
	ER_RPL_SLAVE_ERROR_INFO_FROM_DA uint16 = 10584
	ER_RPL_SLAVE_ERROR_LOADING_USER_DEFINED_LIBRARY uint16 = 10585
	ER_RPL_SLAVE_ERROR_RUNNING_QUERY uint16 = 10586
	ER_RPL_SLAVE_SQL_THREAD_EXITING uint16 = 10587
	ER_RPL_SLAVE_READ_INVALID_EVENT_FROM_MASTER uint16 = 10588
	ER_RPL_SLAVE_QUEUE_EVENT_FAILED_INVALID_CONFIGURATION uint16 = 10589
	ER_RPL_SLAVE_IO_THREAD_DETECTED_UNEXPECTED_EVENT_SEQUENCE uint16 = 10590
	ER_RPL_SLAVE_CANT_USE_CHARSET uint16 = 10591
	ER_RPL_SLAVE_CONNECTED_TO_MASTER_REPLICATION_RESUMED uint16 = 10592
	ER_RPL_SLAVE_NEXT_LOG_IS_ACTIVE uint16 = 10593
	ER_RPL_SLAVE_NEXT_LOG_IS_INACTIVE uint16 = 10594
	ER_RPL_SLAVE_SQL_THREAD_IO_ERROR_READING_EVENT uint16 = 10595
	ER_RPL_SLAVE_ERROR_READING_RELAY_LOG_EVENTS uint16 = 10596
	ER_SLAVE_CHANGE_MASTER_TO_EXECUTED uint16 = 10597
	ER_RPL_SLAVE_NEW_MASTER_INFO_NEEDS_REPOS_TYPE_OTHER_THAN_FILE uint16 = 10598
	ER_RPL_FAILED_TO_STAT_LOG_IN_INDEX uint16 = 10599
	ER_RPL_LOG_NOT_FOUND_WHILE_COUNTING_RELAY_LOG_SPACE uint16 = 10600
	ER_SLAVE_CANT_USE_TEMPDIR uint16 = 10601
	ER_RPL_RELAY_LOG_NEEDS_FILE_NOT_DIRECTORY uint16 = 10602
	ER_RPL_RELAY_LOG_INDEX_NEEDS_FILE_NOT_DIRECTORY uint16 = 10603
	ER_RPL_PLEASE_USE_OPTION_RELAY_LOG uint16 = 10604
	ER_RPL_OPEN_INDEX_FILE_FAILED uint16 = 10605
	ER_RPL_CANT_INITIALIZE_GTID_SETS_IN_RLI_INIT_INFO uint16 = 10606
	ER_RPL_CANT_OPEN_LOG_IN_RLI_INIT_INFO uint16 = 10607
	ER_RPL_ERROR_WRITING_RELAY_LOG_CONFIGURATION uint16 = 10608
	//OBSOLETE_ER_NDB_OOM_GET_NDB_BLOBS_VALUE uint16 = 10609
	//OBSOLETE_ER_NDB_THREAD_TIMED_OUT uint16 = 10610
	//OBSOLETE_ER_NDB_TABLE_IS_NOT_DISTRIBUTED uint16 = 10611
	//OBSOLETE_ER_NDB_CREATING_TABLE uint16 = 10612
	//OBSOLETE_ER_NDB_FLUSHING_TABLE_INFO uint16 = 10613
	//OBSOLETE_ER_NDB_CLEANING_STRAY_TABLES uint16 = 10614
	//OBSOLETE_ER_NDB_DISCOVERED_MISSING_DB uint16 = 10615
	//OBSOLETE_ER_NDB_DISCOVERED_REMAINING_DB uint16 = 10616
	//OBSOLETE_ER_NDB_CLUSTER_FIND_ALL_DBS_RETRY uint16 = 10617
	//OBSOLETE_ER_NDB_CLUSTER_FIND_ALL_DBS_FAIL uint16 = 10618
	//OBSOLETE_ER_NDB_SKIPPING_SETUP_TABLE uint16 = 10619
	//OBSOLETE_ER_NDB_FAILED_TO_SET_UP_TABLE uint16 = 10620
	//OBSOLETE_ER_NDB_MISSING_FRM_DISCOVERING uint16 = 10621
	//OBSOLETE_ER_NDB_MISMATCH_IN_FRM_DISCOVERING uint16 = 10622
	//OBSOLETE_ER_NDB_BINLOG_CLEANING_UP_SETUP_LEFTOVERS uint16 = 10623
	//OBSOLETE_ER_NDB_WAITING_INFO uint16 = 10624
	//OBSOLETE_ER_NDB_WAITING_INFO_WITH_MAP uint16 = 10625
	//OBSOLETE_ER_NDB_TIMEOUT_WHILE_DISTRIBUTING uint16 = 10626
	//OBSOLETE_ER_NDB_NOT_WAITING_FOR_DISTRIBUTING uint16 = 10627
	//OBSOLETE_ER_NDB_DISTRIBUTED_INFO uint16 = 10628
	//OBSOLETE_ER_NDB_DISTRIBUTION_COMPLETE uint16 = 10629
	//OBSOLETE_ER_NDB_SCHEMA_DISTRIBUTION_FAILED uint16 = 10630
	//OBSOLETE_ER_NDB_SCHEMA_DISTRIBUTION_REPORTS_SUBSCRIBE uint16 = 10631
	//OBSOLETE_ER_NDB_SCHEMA_DISTRIBUTION_REPORTS_UNSUBSCRIBE uint16 = 10632
	//OBSOLETE_ER_NDB_BINLOG_CANT_DISCOVER_TABLE_FROM_SCHEMA_EVENT uint16 = 10633
	//OBSOLETE_ER_NDB_BINLOG_SIGNALLING_UNKNOWN_VALUE uint16 = 10634
	//OBSOLETE_ER_NDB_BINLOG_REPLY_TO uint16 = 10635
	//OBSOLETE_ER_NDB_BINLOG_CANT_RELEASE_SLOCK uint16 = 10636
	//OBSOLETE_ER_NDB_CANT_FIND_TABLE uint16 = 10637
	//OBSOLETE_ER_NDB_DISCARDING_EVENT_NO_OBJ uint16 = 10638
	//OBSOLETE_ER_NDB_DISCARDING_EVENT_ID_VERSION_MISMATCH uint16 = 10639
	//OBSOLETE_ER_NDB_CLEAR_SLOCK_INFO uint16 = 10640
	//OBSOLETE_ER_NDB_BINLOG_SKIPPING_LOCAL_TABLE uint16 = 10641
	//OBSOLETE_ER_NDB_BINLOG_ONLINE_ALTER_RENAME uint16 = 10642
	//OBSOLETE_ER_NDB_BINLOG_CANT_REOPEN_SHADOW_TABLE uint16 = 10643
	//OBSOLETE_ER_NDB_BINLOG_ONLINE_ALTER_RENAME_COMPLETE uint16 = 10644
	//OBSOLETE_ER_NDB_BINLOG_SKIPPING_DROP_OF_LOCAL_TABLE uint16 = 10645
	//OBSOLETE_ER_NDB_BINLOG_SKIPPING_RENAME_OF_LOCAL_TABLE uint16 = 10646
	//OBSOLETE_ER_NDB_BINLOG_SKIPPING_DROP_OF_TABLES uint16 = 10647
	//OBSOLETE_ER_NDB_BINLOG_GOT_DIST_PRIV_EVENT_FLUSHING_PRIVILEGES uint16 = 10648
	//OBSOLETE_ER_NDB_BINLOG_GOT_SCHEMA_EVENT uint16 = 10649
	//OBSOLETE_ER_NDB_BINLOG_SKIPPING_OLD_SCHEMA_OPERATION uint16 = 10650
	//OBSOLETE_ER_NDB_CLUSTER_FAILURE uint16 = 10651
	//OBSOLETE_ER_NDB_TABLES_INITIALLY_READ_ONLY_ON_RECONNECT uint16 = 10652
	//OBSOLETE_ER_NDB_IGNORING_UNKNOWN_EVENT uint16 = 10653
	//OBSOLETE_ER_NDB_BINLOG_OPENING_INDEX uint16 = 10654
	//OBSOLETE_ER_NDB_BINLOG_CANT_LOCK_NDB_BINLOG_INDEX uint16 = 10655
	//OBSOLETE_ER_NDB_BINLOG_INJECTING_RANDOM_WRITE_FAILURE uint16 = 10656
	//OBSOLETE_ER_NDB_BINLOG_CANT_WRITE_TO_NDB_BINLOG_INDEX uint16 = 10657
	//OBSOLETE_ER_NDB_BINLOG_WRITING_TO_NDB_BINLOG_INDEX uint16 = 10658
	//OBSOLETE_ER_NDB_BINLOG_CANT_COMMIT_TO_NDB_BINLOG_INDEX uint16 = 10659
	//OBSOLETE_ER_NDB_BINLOG_WRITE_INDEX_FAILED_AFTER_KILL uint16 = 10660
	//OBSOLETE_ER_NDB_BINLOG_USING_SERVER_ID_0_SLAVES_WILL_NOT uint16 = 10661
	//OBSOLETE_ER_NDB_SERVER_ID_RESERVED_OR_TOO_LARGE uint16 = 10662
	//OBSOLETE_ER_NDB_BINLOG_REQUIRES_V2_ROW_EVENTS uint16 = 10663
	//OBSOLETE_ER_NDB_BINLOG_STATUS_FORCING_FULL_USE_WRITE uint16 = 10664
	//OBSOLETE_ER_NDB_BINLOG_GENERIC_MESSAGE uint16 = 10665
	//OBSOLETE_ER_NDB_CONFLICT_GENERIC_MESSAGE uint16 = 10666
	//OBSOLETE_ER_NDB_TRANS_DEPENDENCY_TRACKER_ERROR uint16 = 10667
	//OBSOLETE_ER_NDB_CONFLICT_FN_PARSE_ERROR uint16 = 10668
	//OBSOLETE_ER_NDB_CONFLICT_FN_SETUP_ERROR uint16 = 10669
	//OBSOLETE_ER_NDB_BINLOG_FAILED_TO_GET_TABLE uint16 = 10670
	//OBSOLETE_ER_NDB_BINLOG_NOT_LOGGING uint16 = 10671
	//OBSOLETE_ER_NDB_BINLOG_CREATE_TABLE_EVENT_FAILED uint16 = 10672
	//OBSOLETE_ER_NDB_BINLOG_CREATE_TABLE_EVENT_INFO uint16 = 10673
	//OBSOLETE_ER_NDB_BINLOG_DISCOVER_TABLE_EVENT_INFO uint16 = 10674
	//OBSOLETE_ER_NDB_BINLOG_BLOB_REQUIRES_PK uint16 = 10675
	//OBSOLETE_ER_NDB_BINLOG_CANT_CREATE_EVENT_IN_DB uint16 = 10676
	//OBSOLETE_ER_NDB_BINLOG_CANT_CREATE_EVENT_IN_DB_AND_CANT_DROP uint16 = 10677
	//OBSOLETE_ER_NDB_BINLOG_CANT_CREATE_EVENT_IN_DB_DROPPED uint16 = 10678
	//OBSOLETE_ER_NDB_BINLOG_DISCOVER_REUSING_OLD_EVENT_OPS uint16 = 10679
	//OBSOLETE_ER_NDB_BINLOG_CREATING_NDBEVENTOPERATION_FAILED uint16 = 10680
	//OBSOLETE_ER_NDB_BINLOG_CANT_CREATE_BLOB uint16 = 10681
	//OBSOLETE_ER_NDB_BINLOG_NDBEVENT_EXECUTE_FAILED uint16 = 10682
	//OBSOLETE_ER_NDB_CREATE_EVENT_OPS_LOGGING_INFO uint16 = 10683
	//OBSOLETE_ER_NDB_BINLOG_CANT_DROP_EVENT_FROM_DB uint16 = 10684
	//OBSOLETE_ER_NDB_TIMED_OUT_IN_DROP_TABLE uint16 = 10685
	//OBSOLETE_ER_NDB_BINLOG_UNHANDLED_ERROR_FOR_TABLE uint16 = 10686
	//OBSOLETE_ER_NDB_BINLOG_CLUSTER_FAILURE uint16 = 10687
	//OBSOLETE_ER_NDB_BINLOG_UNKNOWN_NON_DATA_EVENT uint16 = 10688
	//OBSOLETE_ER_NDB_BINLOG_INJECTOR_DISCARDING_ROW_EVENT_METADATA uint16 = 10689
	//OBSOLETE_ER_NDB_REMAINING_OPEN_TABLES uint16 = 10690
	//OBSOLETE_ER_NDB_REMAINING_OPEN_TABLE_INFO uint16 = 10691
	//OBSOLETE_ER_NDB_COULD_NOT_GET_APPLY_STATUS_SHARE uint16 = 10692
	//OBSOLETE_ER_NDB_BINLOG_SERVER_SHUTDOWN_DURING_NDB_CLUSTER_START uint16 = 10693
	//OBSOLETE_ER_NDB_BINLOG_CLUSTER_RESTARTED_RESET_MASTER_SUGGESTED uint16 = 10694
	//OBSOLETE_ER_NDB_BINLOG_CLUSTER_HAS_RECONNECTED uint16 = 10695
	//OBSOLETE_ER_NDB_BINLOG_STARTING_LOG_AT_EPOCH uint16 = 10696
	//OBSOLETE_ER_NDB_BINLOG_NDB_TABLES_WRITABLE uint16 = 10697
	//OBSOLETE_ER_NDB_BINLOG_SHUTDOWN_DETECTED uint16 = 10698
	//OBSOLETE_ER_NDB_BINLOG_LOST_SCHEMA_CONNECTION_WAITING uint16 = 10699
	//OBSOLETE_ER_NDB_BINLOG_LOST_SCHEMA_CONNECTION_CONTINUING uint16 = 10700
	//OBSOLETE_ER_NDB_BINLOG_ERROR_HANDLING_SCHEMA_EVENT uint16 = 10701
	//OBSOLETE_ER_NDB_BINLOG_CANT_INJECT_APPLY_STATUS_WRITE_ROW uint16 = 10702
	//OBSOLETE_ER_NDB_BINLOG_ERROR_DURING_GCI_ROLLBACK uint16 = 10703
	//OBSOLETE_ER_NDB_BINLOG_ERROR_DURING_GCI_COMMIT uint16 = 10704
	//OBSOLETE_ER_NDB_BINLOG_LATEST_TRX_IN_EPOCH_NOT_IN_BINLOG uint16 = 10705
	//OBSOLETE_ER_NDB_BINLOG_RELEASING_EXTRA_SHARE_REFERENCES uint16 = 10706
	//OBSOLETE_ER_NDB_BINLOG_REMAINING_OPEN_TABLES uint16 = 10707
	//OBSOLETE_ER_NDB_BINLOG_REMAINING_OPEN_TABLE_INFO uint16 = 10708
	ER_TREE_CORRUPT_PARENT_SHOULD_POINT_AT_PARENT uint16 = 10709
	ER_TREE_CORRUPT_ROOT_SHOULD_BE_BLACK uint16 = 10710
	ER_TREE_CORRUPT_2_CONSECUTIVE_REDS uint16 = 10711
	ER_TREE_CORRUPT_RIGHT_IS_LEFT uint16 = 10712
	ER_TREE_CORRUPT_INCORRECT_BLACK_COUNT uint16 = 10713
	ER_WRONG_COUNT_FOR_ORIGIN uint16 = 10714
	ER_WRONG_COUNT_FOR_KEY uint16 = 10715
	ER_WRONG_COUNT_OF_ELEMENTS uint16 = 10716
	ER_RPL_ERROR_READING_SLAVE_WORKER_CONFIGURATION uint16 = 10717
	//OBSOLETE_ER_RPL_ERROR_WRITING_SLAVE_WORKER_CONFIGURATION uint16 = 10718
	ER_RPL_FAILED_TO_OPEN_RELAY_LOG uint16 = 10719
	ER_RPL_WORKER_CANT_READ_RELAY_LOG uint16 = 10720
	ER_RPL_WORKER_CANT_FIND_NEXT_RELAY_LOG uint16 = 10721
	ER_RPL_MTS_SLAVE_COORDINATOR_HAS_WAITED uint16 = 10722
	ER_BINLOG_FAILED_TO_WRITE_DROP_FOR_TEMP_TABLES uint16 = 10723
	ER_BINLOG_OOM_WRITING_DELETE_WHILE_OPENING_HEAP_TABLE uint16 = 10724
	ER_FAILED_TO_REPAIR_TABLE uint16 = 10725
	ER_FAILED_TO_REMOVE_TEMP_TABLE uint16 = 10726
	ER_SYSTEM_TABLE_NOT_TRANSACTIONAL uint16 = 10727
	ER_RPL_ERROR_WRITING_MASTER_CONFIGURATION uint16 = 10728
	ER_RPL_ERROR_READING_MASTER_CONFIGURATION uint16 = 10729
	ER_RPL_SSL_INFO_IN_MASTER_INFO_IGNORED uint16 = 10730
	ER_PLUGIN_FAILED_DEINITIALIZATION uint16 = 10731
	ER_PLUGIN_HAS_NONZERO_REFCOUNT_AFTER_DEINITIALIZATION uint16 = 10732
	ER_PLUGIN_SHUTTING_DOWN_PLUGIN uint16 = 10733
	ER_PLUGIN_REGISTRATION_FAILED uint16 = 10734
	ER_PLUGIN_CANT_OPEN_PLUGIN_TABLE uint16 = 10735
	ER_PLUGIN_CANT_LOAD uint16 = 10736
	ER_PLUGIN_LOAD_PARAMETER_TOO_LONG uint16 = 10737
	ER_PLUGIN_FORCING_SHUTDOWN uint16 = 10738
	ER_PLUGIN_HAS_NONZERO_REFCOUNT_AFTER_SHUTDOWN uint16 = 10739
	ER_PLUGIN_UNKNOWN_VARIABLE_TYPE uint16 = 10740
	ER_PLUGIN_VARIABLE_SET_READ_ONLY uint16 = 10741
	ER_PLUGIN_VARIABLE_MISSING_NAME uint16 = 10742
	ER_PLUGIN_VARIABLE_NOT_ALLOCATED_THREAD_LOCAL uint16 = 10743
	ER_PLUGIN_OOM uint16 = 10744
	ER_PLUGIN_BAD_OPTIONS uint16 = 10745
	ER_PLUGIN_PARSING_OPTIONS_FAILED uint16 = 10746
	ER_PLUGIN_DISABLED uint16 = 10747
	ER_PLUGIN_HAS_CONFLICTING_SYSTEM_VARIABLES uint16 = 10748
	ER_PLUGIN_CANT_SET_PERSISTENT_OPTIONS uint16 = 10749
	ER_MY_NET_WRITE_FAILED_FALLING_BACK_ON_STDERR uint16 = 10750
	ER_RETRYING_REPAIR_WITHOUT_QUICK uint16 = 10751
	ER_RETRYING_REPAIR_WITH_KEYCACHE uint16 = 10752
	ER_FOUND_ROWS_WHILE_REPAIRING uint16 = 10753
	ER_ERROR_DURING_OPTIMIZE_TABLE uint16 = 10754
	ER_ERROR_ENABLING_KEYS uint16 = 10755
	ER_CHECKING_TABLE uint16 = 10756
	ER_RECOVERING_TABLE uint16 = 10757
	ER_CANT_CREATE_TABLE_SHARE_FROM_FRM uint16 = 10758
	ER_CANT_LOCK_TABLE uint16 = 10759
	ER_CANT_ALLOC_TABLE_OBJECT uint16 = 10760
	ER_CANT_CREATE_HANDLER_OBJECT_FOR_TABLE uint16 = 10761
	ER_CANT_SET_HANDLER_REFERENCE_FOR_TABLE uint16 = 10762
	ER_CANT_LOCK_TABLESPACE uint16 = 10763
	ER_CANT_UPGRADE_GENERATED_COLUMNS_TO_DD uint16 = 10764
	ER_DD_ERROR_CREATING_ENTRY uint16 = 10765
	ER_DD_CANT_FETCH_TABLE_DATA uint16 = 10766
	ER_DD_CANT_FIX_SE_DATA uint16 = 10767
	ER_DD_CANT_CREATE_SP uint16 = 10768
	ER_CANT_OPEN_DB_OPT_USING_DEFAULT_CHARSET uint16 = 10769
	ER_CANT_CREATE_CACHE_FOR_DB_OPT uint16 = 10770
	ER_CANT_IDENTIFY_CHARSET_USING_DEFAULT uint16 = 10771
	ER_DB_OPT_NOT_FOUND_USING_DEFAULT_CHARSET uint16 = 10772
	ER_EVENT_CANT_GET_TIMEZONE_FROM_FIELD uint16 = 10773
	ER_EVENT_CANT_FIND_TIMEZONE uint16 = 10774
	ER_EVENT_CANT_GET_CHARSET uint16 = 10775
	ER_EVENT_CANT_GET_COLLATION uint16 = 10776
	ER_EVENT_CANT_OPEN_TABLE_MYSQL_EVENT uint16 = 10777
	ER_CANT_PARSE_STORED_ROUTINE_BODY uint16 = 10778
	ER_CANT_OPEN_TABLE_MYSQL_PROC uint16 = 10779
	ER_CANT_READ_TABLE_MYSQL_PROC uint16 = 10780
	ER_FILE_EXISTS_DURING_UPGRADE uint16 = 10781
	ER_CANT_OPEN_DATADIR_AFTER_UPGRADE_FAILURE uint16 = 10782
	ER_CANT_SET_PATH_FOR uint16 = 10783
	ER_CANT_OPEN_DIR uint16 = 10784
	//OBSOLETE_ER_NDB_CLUSTER_CONNECTION_POOL_NODEIDS uint16 = 10785
	//OBSOLETE_ER_NDB_CANT_PARSE_NDB_CLUSTER_CONNECTION_POOL_NODEIDS uint16 = 10786
	//OBSOLETE_ER_NDB_INVALID_CLUSTER_CONNECTION_POOL_NODEIDS uint16 = 10787
	//OBSOLETE_ER_NDB_DUPLICATE_CLUSTER_CONNECTION_POOL_NODEIDS uint16 = 10788
	//OBSOLETE_ER_NDB_POOL_SIZE_CLUSTER_CONNECTION_POOL_NODEIDS uint16 = 10789
	//OBSOLETE_ER_NDB_NODEID_NOT_FIRST_CONNECTION_POOL_NODEIDS uint16 = 10790
	//OBSOLETE_ER_NDB_USING_NODEID uint16 = 10791
	//OBSOLETE_ER_NDB_CANT_ALLOC_GLOBAL_NDB_CLUSTER_CONNECTION uint16 = 10792
	//OBSOLETE_ER_NDB_CANT_ALLOC_GLOBAL_NDB_OBJECT uint16 = 10793
	//OBSOLETE_ER_NDB_USING_NODEID_LIST uint16 = 10794
	//OBSOLETE_ER_NDB_CANT_ALLOC_NDB_CLUSTER_CONNECTION uint16 = 10795
	//OBSOLETE_ER_NDB_STARTING_CONNECT_THREAD uint16 = 10796
	//OBSOLETE_ER_NDB_NODE_INFO uint16 = 10797
	//OBSOLETE_ER_NDB_CANT_START_CONNECT_THREAD uint16 = 10798
	//OBSOLETE_ER_NDB_GENERIC_ERROR uint16 = 10799
	//OBSOLETE_ER_NDB_CPU_MASK_TOO_SHORT uint16 = 10800
	ER_EVENT_ERROR_CREATING_QUERY_TO_WRITE_TO_BINLOG uint16 = 10801
	ER_EVENT_SCHEDULER_ERROR_LOADING_FROM_DB uint16 = 10802
	ER_EVENT_SCHEDULER_ERROR_GETTING_EVENT_OBJECT uint16 = 10803
	ER_EVENT_SCHEDULER_GOT_BAD_DATA_FROM_TABLE uint16 = 10804
	ER_EVENT_CANT_GET_LOCK_FOR_DROPPING_EVENT uint16 = 10805
	ER_EVENT_UNABLE_TO_DROP_EVENT uint16 = 10806
	//OBSOLETE_ER_BINLOG_ATTACHING_THREAD_MEMORY_FINALLY_AVAILABLE uint16 = 10807
	ER_BINLOG_CANT_RESIZE_CACHE uint16 = 10808
	ER_BINLOG_FILE_BEING_READ_NOT_PURGED uint16 = 10809
	ER_BINLOG_IO_ERROR_READING_HEADER uint16 = 10810
	//OBSOLETE_ER_BINLOG_CANT_OPEN_LOG uint16 = 10811
	//OBSOLETE_ER_BINLOG_CANT_CREATE_CACHE_FOR_LOG uint16 = 10812
	ER_BINLOG_FILE_EXTENSION_NUMBER_EXHAUSTED uint16 = 10813
	ER_BINLOG_FILE_NAME_TOO_LONG uint16 = 10814
	ER_BINLOG_FILE_EXTENSION_NUMBER_RUNNING_LOW uint16 = 10815
	ER_BINLOG_CANT_OPEN_FOR_LOGGING uint16 = 10816
	ER_BINLOG_FAILED_TO_SYNC_INDEX_FILE uint16 = 10817
	ER_BINLOG_ERROR_READING_GTIDS_FROM_RELAY_LOG uint16 = 10818
	ER_BINLOG_EVENTS_READ_FROM_RELAY_LOG_INFO uint16 = 10819
	ER_BINLOG_ERROR_READING_GTIDS_FROM_BINARY_LOG uint16 = 10820
	ER_BINLOG_EVENTS_READ_FROM_BINLOG_INFO uint16 = 10821
	ER_BINLOG_CANT_GENERATE_NEW_FILE_NAME uint16 = 10822
	ER_BINLOG_FAILED_TO_SYNC_INDEX_FILE_IN_OPEN uint16 = 10823
	ER_BINLOG_CANT_USE_FOR_LOGGING uint16 = 10824
	ER_BINLOG_FAILED_TO_CLOSE_INDEX_FILE_WHILE_REBUILDING uint16 = 10825
	ER_BINLOG_FAILED_TO_DELETE_INDEX_FILE_WHILE_REBUILDING uint16 = 10826
	ER_BINLOG_FAILED_TO_RENAME_INDEX_FILE_WHILE_REBUILDING uint16 = 10827
	ER_BINLOG_FAILED_TO_OPEN_INDEX_FILE_AFTER_REBUILDING uint16 = 10828
	ER_BINLOG_CANT_APPEND_LOG_TO_TMP_INDEX uint16 = 10829
	ER_BINLOG_CANT_LOCATE_OLD_BINLOG_OR_RELAY_LOG_FILES uint16 = 10830
	ER_BINLOG_CANT_DELETE_FILE uint16 = 10831
	ER_BINLOG_CANT_SET_TMP_INDEX_NAME uint16 = 10832
	ER_BINLOG_FAILED_TO_OPEN_TEMPORARY_INDEX_FILE uint16 = 10833
	//OBSOLETE_ER_BINLOG_ERROR_GETTING_NEXT_LOG_FROM_INDEX uint16 = 10834
	ER_BINLOG_CANT_OPEN_TMP_INDEX uint16 = 10835
	ER_BINLOG_CANT_COPY_INDEX_TO_TMP uint16 = 10836
	ER_BINLOG_CANT_CLOSE_TMP_INDEX uint16 = 10837
	ER_BINLOG_CANT_MOVE_TMP_TO_INDEX uint16 = 10838
	ER_BINLOG_PURGE_LOGS_CALLED_WITH_FILE_NOT_IN_INDEX uint16 = 10839
	ER_BINLOG_PURGE_LOGS_CANT_SYNC_INDEX_FILE uint16 = 10840
	ER_BINLOG_PURGE_LOGS_CANT_COPY_TO_REGISTER_FILE uint16 = 10841
	ER_BINLOG_PURGE_LOGS_CANT_FLUSH_REGISTER_FILE uint16 = 10842
	ER_BINLOG_PURGE_LOGS_CANT_UPDATE_INDEX_FILE uint16 = 10843
	ER_BINLOG_PURGE_LOGS_FAILED_TO_PURGE_LOG uint16 = 10844
	ER_BINLOG_FAILED_TO_SET_PURGE_INDEX_FILE_NAME uint16 = 10845
	ER_BINLOG_FAILED_TO_OPEN_REGISTER_FILE uint16 = 10846
	ER_BINLOG_FAILED_TO_REINIT_REGISTER_FILE uint16 = 10847
	ER_BINLOG_FAILED_TO_READ_REGISTER_FILE uint16 = 10848
	ER_CANT_STAT_FILE uint16 = 10849
	ER_BINLOG_CANT_DELETE_LOG_FILE_DOES_INDEX_MATCH_FILES uint16 = 10850
	ER_BINLOG_CANT_DELETE_FILE_AND_READ_BINLOG_INDEX uint16 = 10851
	ER_BINLOG_FAILED_TO_DELETE_LOG_FILE uint16 = 10852
	ER_BINLOG_LOGGING_INCIDENT_TO_STOP_SLAVES uint16 = 10853
	ER_BINLOG_CANT_FIND_LOG_IN_INDEX uint16 = 10854
	ER_BINLOG_RECOVERING_AFTER_CRASH_USING uint16 = 10855
	ER_BINLOG_CANT_OPEN_CRASHED_BINLOG uint16 = 10856
	ER_BINLOG_CANT_TRIM_CRASHED_BINLOG uint16 = 10857
	ER_BINLOG_CRASHED_BINLOG_TRIMMED uint16 = 10858
	ER_BINLOG_CANT_CLEAR_IN_USE_FLAG_FOR_CRASHED_BINLOG uint16 = 10859
	ER_BINLOG_FAILED_TO_RUN_AFTER_SYNC_HOOK uint16 = 10860
	ER_TURNING_LOGGING_OFF_FOR_THE_DURATION uint16 = 10861
	ER_BINLOG_FAILED_TO_RUN_AFTER_FLUSH_HOOK uint16 = 10862
	ER_BINLOG_CRASH_RECOVERY_FAILED uint16 = 10863
	ER_BINLOG_WARNING_SUPPRESSED uint16 = 10864
	ER_NDB_LOG_ENTRY uint16 = 10865
	ER_NDB_LOG_ENTRY_WITH_PREFIX uint16 = 10866
	//OBSOLETE_ER_NDB_BINLOG_CANT_CREATE_PURGE_THD uint16 = 10867
	ER_INNODB_UNKNOWN_COLLATION uint16 = 10868
	ER_INNODB_INVALID_LOG_GROUP_HOME_DIR uint16 = 10869
	ER_INNODB_INVALID_INNODB_UNDO_DIRECTORY uint16 = 10870
	ER_INNODB_ILLEGAL_COLON_IN_POOL uint16 = 10871
	ER_INNODB_INVALID_PAGE_SIZE uint16 = 10872
	ER_INNODB_DIRTY_WATER_MARK_NOT_LOW uint16 = 10873
	ER_INNODB_IO_CAPACITY_EXCEEDS_MAX uint16 = 10874
	ER_INNODB_FILES_SAME uint16 = 10875
	ER_INNODB_UNREGISTERED_TRX_ACTIVE uint16 = 10876
	ER_INNODB_CLOSING_CONNECTION_ROLLS_BACK uint16 = 10877
	ER_INNODB_TRX_XLATION_TABLE_OOM uint16 = 10878
	ER_INNODB_CANT_FIND_INDEX_IN_INNODB_DD uint16 = 10879
	ER_INNODB_INDEX_COLUMN_INFO_UNLIKE_MYSQLS uint16 = 10880
	//OBSOLETE_ER_INNODB_CANT_OPEN_TABLE uint16 = 10881
	ER_INNODB_CANT_BUILD_INDEX_XLATION_TABLE_FOR uint16 = 10882
	ER_INNODB_PK_NOT_IN_MYSQL uint16 = 10883
	ER_INNODB_PK_ONLY_IN_MYSQL uint16 = 10884
	ER_INNODB_CLUSTERED_INDEX_PRIVATE uint16 = 10885
	//OBSOLETE_ER_INNODB_PARTITION_TABLE_LOWERCASED uint16 = 10886
	ER_ERRMSG_REPLACEMENT_DODGY uint16 = 10887
	ER_ERRMSG_REPLACEMENTS_FAILED uint16 = 10888
	ER_NPIPE_CANT_CREATE uint16 = 10889
	ER_PARTITION_MOVE_CREATED_DUPLICATE_ROW_PLEASE_FIX uint16 = 10890
	ER_AUDIT_CANT_ABORT_COMMAND uint16 = 10891
	ER_AUDIT_CANT_ABORT_EVENT uint16 = 10892
	ER_AUDIT_WARNING uint16 = 10893
	//OBSOLETE_ER_NDB_NUMBER_OF_CHANNELS uint16 = 10894
	//OBSOLETE_ER_NDB_SLAVE_PARALLEL_WORKERS uint16 = 10895
	//OBSOLETE_ER_NDB_DISTRIBUTING_ERR uint16 = 10896
	ER_RPL_SLAVE_INSECURE_CHANGE_MASTER uint16 = 10897
	//OBSOLETE_ER_RPL_SLAVE_FLUSH_RELAY_LOGS_NOT_ALLOWED uint16 = 10898
	ER_RPL_SLAVE_INCORRECT_CHANNEL uint16 = 10899
	ER_FAILED_TO_FIND_DL_ENTRY uint16 = 10900
	ER_FAILED_TO_OPEN_SHARED_LIBRARY uint16 = 10901
	ER_THREAD_PRIORITY_IGNORED uint16 = 10902
	ER_BINLOG_CACHE_SIZE_TOO_LARGE uint16 = 10903
	ER_BINLOG_STMT_CACHE_SIZE_TOO_LARGE uint16 = 10904
	ER_FAILED_TO_GENERATE_UNIQUE_LOGFILE uint16 = 10905
	ER_FAILED_TO_READ_FILE uint16 = 10906
	ER_FAILED_TO_WRITE_TO_FILE uint16 = 10907
	ER_BINLOG_UNSAFE_MESSAGE_AND_STATEMENT uint16 = 10908
	ER_FORCE_CLOSE_THREAD uint16 = 10909
	ER_SERVER_SHUTDOWN_COMPLETE uint16 = 10910
	ER_RPL_CANT_HAVE_SAME_BASENAME uint16 = 10911
	ER_RPL_GTID_MODE_REQUIRES_ENFORCE_GTID_CONSISTENCY_ON uint16 = 10912
	ER_WARN_NO_SERVERID_SPECIFIED uint16 = 10913
	ER_ABORTING_USER_CONNECTION uint16 = 10914
	ER_SQL_MODE_MERGED_WITH_STRICT_MODE uint16 = 10915
	ER_GTID_PURGED_WAS_UPDATED uint16 = 10916
	ER_GTID_EXECUTED_WAS_UPDATED uint16 = 10917
	ER_DEPRECATE_MSG_WITH_REPLACEMENT uint16 = 10918
	ER_TRG_CREATION_CTX_NOT_SET uint16 = 10919
	ER_FILE_HAS_OLD_FORMAT uint16 = 10920
	ER_VIEW_CREATION_CTX_NOT_SET uint16 = 10921
	//OBSOLETE_ER_TABLE_NAME_CAUSES_TOO_LONG_PATH uint16 = 10922
	ER_TABLE_UPGRADE_REQUIRED uint16 = 10923
	ER_GET_ERRNO_FROM_STORAGE_ENGINE uint16 = 10924
	ER_ACCESS_DENIED_ERROR_WITHOUT_PASSWORD uint16 = 10925
	ER_ACCESS_DENIED_ERROR_WITH_PASSWORD uint16 = 10926
	ER_ACCESS_DENIED_FOR_USER_ACCOUNT_LOCKED uint16 = 10927
	ER_MUST_CHANGE_EXPIRED_PASSWORD uint16 = 10928
	ER_SYSTEM_TABLES_NOT_SUPPORTED_BY_STORAGE_ENGINE uint16 = 10929
	//OBSOLETE_ER_FILESORT_TERMINATED uint16 = 10930
	ER_SERVER_STARTUP_MSG uint16 = 10931
	ER_FAILED_TO_FIND_LOCALE_NAME uint16 = 10932
	ER_FAILED_TO_FIND_COLLATION_NAME uint16 = 10933
	ER_SERVER_OUT_OF_RESOURCES uint16 = 10934
	ER_SERVER_OUTOFMEMORY uint16 = 10935
	ER_INVALID_COLLATION_FOR_CHARSET uint16 = 10936
	ER_CANT_START_ERROR_LOG_SERVICE uint16 = 10937
	ER_CREATING_NEW_UUID_FIRST_START uint16 = 10938
	ER_FAILED_TO_GET_ABSOLUTE_PATH uint16 = 10939
	ER_PERFSCHEMA_COMPONENTS_INFRASTRUCTURE_BOOTSTRAP uint16 = 10940
	ER_PERFSCHEMA_COMPONENTS_INFRASTRUCTURE_SHUTDOWN uint16 = 10941
	ER_DUP_FD_OPEN_FAILED uint16 = 10942
	ER_SYSTEM_VIEW_INIT_FAILED uint16 = 10943
	ER_RESOURCE_GROUP_POST_INIT_FAILED uint16 = 10944
	ER_RESOURCE_GROUP_SUBSYSTEM_INIT_FAILED uint16 = 10945
	ER_FAILED_START_MYSQLD_DAEMON uint16 = 10946
	ER_CANNOT_CHANGE_TO_ROOT_DIR uint16 = 10947
	ER_PERSISTENT_PRIVILEGES_BOOTSTRAP uint16 = 10948
	ER_BASEDIR_SET_TO uint16 = 10949
	ER_RPL_FILTER_ADD_WILD_DO_TABLE_FAILED uint16 = 10950
	ER_RPL_FILTER_ADD_WILD_IGNORE_TABLE_FAILED uint16 = 10951
	ER_PRIVILEGE_SYSTEM_INIT_FAILED uint16 = 10952
	ER_CANNOT_SET_LOG_ERROR_SERVICES uint16 = 10953
	ER_PERFSCHEMA_TABLES_INIT_FAILED uint16 = 10954
	ER_TX_EXTRACTION_ALGORITHM_FOR_BINLOG_TX_DEPEDENCY_TRACKING uint16 = 10955
	ER_INVALID_REPLICATION_TIMESTAMPS uint16 = 10956
	ER_RPL_TIMESTAMPS_RETURNED_TO_NORMAL uint16 = 10957
	ER_BINLOG_FILE_OPEN_FAILED uint16 = 10958
	ER_BINLOG_EVENT_WRITE_TO_STMT_CACHE_FAILED uint16 = 10959
	ER_SLAVE_RELAY_LOG_TRUNCATE_INFO uint16 = 10960
	ER_SLAVE_RELAY_LOG_PURGE_FAILED uint16 = 10961
	ER_RPL_SLAVE_FILTER_CREATE_FAILED uint16 = 10962
	ER_RPL_SLAVE_GLOBAL_FILTERS_COPY_FAILED uint16 = 10963
	ER_RPL_SLAVE_RESET_FILTER_OPTIONS uint16 = 10964
	ER_MISSING_GRANT_SYSTEM_TABLE uint16 = 10965
	ER_MISSING_ACL_SYSTEM_TABLE uint16 = 10966
	ER_ANONYMOUS_AUTH_ID_NOT_ALLOWED_IN_MANDATORY_ROLES uint16 = 10967
	ER_UNKNOWN_AUTH_ID_IN_MANDATORY_ROLE uint16 = 10968
	ER_WRITE_ROW_TO_PARTITION_FAILED uint16 = 10969
	ER_RESOURCE_GROUP_METADATA_UPDATE_SKIPPED uint16 = 10970
	ER_FAILED_TO_PERSIST_RESOURCE_GROUP_METADATA uint16 = 10971
	ER_FAILED_TO_DESERIALIZE_RESOURCE_GROUP uint16 = 10972
	ER_FAILED_TO_UPDATE_RESOURCE_GROUP uint16 = 10973
	ER_RESOURCE_GROUP_VALIDATION_FAILED uint16 = 10974
	ER_FAILED_TO_ALLOCATE_MEMORY_FOR_RESOURCE_GROUP uint16 = 10975
	ER_FAILED_TO_ALLOCATE_MEMORY_FOR_RESOURCE_GROUP_HASH uint16 = 10976
	ER_FAILED_TO_ADD_RESOURCE_GROUP_TO_MAP uint16 = 10977
	ER_RESOURCE_GROUP_IS_DISABLED uint16 = 10978
	ER_FAILED_TO_APPLY_RESOURCE_GROUP_CONTROLLER uint16 = 10979
	ER_FAILED_TO_ACQUIRE_LOCK_ON_RESOURCE_GROUP uint16 = 10980
	ER_PFS_NOTIFICATION_FUNCTION_REGISTER_FAILED uint16 = 10981
	ER_RES_GRP_SET_THR_AFFINITY_FAILED uint16 = 10982
	ER_RES_GRP_SET_THR_AFFINITY_TO_CPUS_FAILED uint16 = 10983
	ER_RES_GRP_THD_UNBIND_FROM_CPU_FAILED uint16 = 10984
	ER_RES_GRP_SET_THREAD_PRIORITY_FAILED uint16 = 10985
	ER_RES_GRP_FAILED_TO_DETERMINE_NICE_CAPABILITY uint16 = 10986
	ER_RES_GRP_FAILED_TO_GET_THREAD_HANDLE uint16 = 10987
	ER_RES_GRP_GET_THREAD_PRIO_NOT_SUPPORTED uint16 = 10988
	ER_RES_GRP_FAILED_DETERMINE_CPU_COUNT uint16 = 10989
	ER_RES_GRP_FEATURE_NOT_AVAILABLE uint16 = 10990
	ER_RES_GRP_INVALID_THREAD_PRIORITY uint16 = 10991
	ER_RES_GRP_SOLARIS_PROCESSOR_BIND_TO_CPUID_FAILED uint16 = 10992
	ER_RES_GRP_SOLARIS_PROCESSOR_BIND_TO_THREAD_FAILED uint16 = 10993
	ER_RES_GRP_SOLARIS_PROCESSOR_AFFINITY_FAILED uint16 = 10994
	ER_DD_UPGRADE_RENAME_IDX_STATS_FILE_FAILED uint16 = 10995
	ER_DD_UPGRADE_DD_OPEN_FAILED uint16 = 10996
	ER_DD_UPGRADE_FAILED_TO_FETCH_TABLESPACES uint16 = 10997
	ER_DD_UPGRADE_FAILED_TO_ACQUIRE_TABLESPACE uint16 = 10998
	ER_DD_UPGRADE_FAILED_TO_RESOLVE_TABLESPACE_ENGINE uint16 = 10999
	ER_FAILED_TO_CREATE_SDI_FOR_TABLESPACE uint16 = 11000
	ER_FAILED_TO_STORE_SDI_FOR_TABLESPACE uint16 = 11001
	ER_DD_UPGRADE_FAILED_TO_FETCH_TABLES uint16 = 11002
	ER_DD_UPGRADE_DD_POPULATED uint16 = 11003
	ER_DD_UPGRADE_INFO_FILE_OPEN_FAILED uint16 = 11004
	ER_DD_UPGRADE_INFO_FILE_CLOSE_FAILED uint16 = 11005
	ER_DD_UPGRADE_TABLESPACE_MIGRATION_FAILED uint16 = 11006
	ER_DD_UPGRADE_FAILED_TO_CREATE_TABLE_STATS uint16 = 11007
	ER_DD_UPGRADE_TABLE_STATS_MIGRATE_COMPLETED uint16 = 11008
	ER_DD_UPGRADE_FAILED_TO_CREATE_INDEX_STATS uint16 = 11009
	ER_DD_UPGRADE_INDEX_STATS_MIGRATE_COMPLETED uint16 = 11010
	ER_DD_UPGRADE_FAILED_FIND_VALID_DATA_DIR uint16 = 11011
	ER_DD_UPGRADE_START uint16 = 11012
	ER_DD_UPGRADE_FAILED_INIT_DD_SE uint16 = 11013
	ER_DD_UPGRADE_FOUND_PARTIALLY_UPGRADED_DD_ABORT uint16 = 11014
	ER_DD_UPGRADE_FOUND_PARTIALLY_UPGRADED_DD_CONTINUE uint16 = 11015
	ER_DD_UPGRADE_SE_LOGS_FAILED uint16 = 11016
	ER_DD_UPGRADE_SDI_INFO_UPDATE_FAILED uint16 = 11017
	ER_SKIP_UPDATING_METADATA_IN_SE_RO_MODE uint16 = 11018
	ER_CREATED_SYSTEM_WITH_VERSION uint16 = 11019
	ER_UNKNOWN_ERROR_DETECTED_IN_SE uint16 = 11020
	ER_READ_LOG_EVENT_FAILED uint16 = 11021
	ER_ROW_DATA_TOO_BIG_TO_WRITE_IN_BINLOG uint16 = 11022
	ER_FAILED_TO_CONSTRUCT_DROP_EVENT_QUERY uint16 = 11023
	ER_FAILED_TO_BINLOG_DROP_EVENT uint16 = 11024
	ER_FAILED_TO_START_SLAVE_THREAD uint16 = 11025
	ER_RPL_IO_THREAD_KILLED uint16 = 11026
	ER_SLAVE_RECONNECT_FAILED uint16 = 11027
	ER_SLAVE_KILLED_AFTER_RECONNECT uint16 = 11028
	ER_SLAVE_NOT_STARTED_ON_SOME_CHANNELS uint16 = 11029
	ER_FAILED_TO_ADD_RPL_FILTER uint16 = 11030
	ER_PER_CHANNEL_RPL_FILTER_CONF_FOR_GRP_RPL uint16 = 11031
	ER_RPL_FILTERS_NOT_ATTACHED_TO_CHANNEL uint16 = 11032
	ER_FAILED_TO_BUILD_DO_AND_IGNORE_TABLE_HASHES uint16 = 11033
	ER_CLONE_PLUGIN_NOT_LOADED_TRACE uint16 = 11034
	ER_CLONE_HANDLER_EXIST_TRACE uint16 = 11035
	ER_CLONE_CREATE_HANDLER_FAIL_TRACE uint16 = 11036
	ER_CYCLE_TIMER_IS_NOT_AVAILABLE uint16 = 11037
	ER_NANOSECOND_TIMER_IS_NOT_AVAILABLE uint16 = 11038
	ER_MICROSECOND_TIMER_IS_NOT_AVAILABLE uint16 = 11039
	ER_PFS_MALLOC_ARRAY_OVERFLOW uint16 = 11040
	ER_PFS_MALLOC_ARRAY_OOM uint16 = 11041
	ER_INNODB_FAILED_TO_FIND_IDX_WITH_KEY_NO uint16 = 11042
	ER_INNODB_FAILED_TO_FIND_IDX uint16 = 11043
	ER_INNODB_FAILED_TO_FIND_IDX_FROM_DICT_CACHE uint16 = 11044
	ER_INNODB_ACTIVE_INDEX_CHANGE_FAILED uint16 = 11045
	ER_INNODB_DIFF_IN_REF_LEN uint16 = 11046
	ER_WRONG_TYPE_FOR_COLUMN_PREFIX_IDX_FLD uint16 = 11047
	ER_INNODB_CANNOT_CREATE_TABLE uint16 = 11048
	ER_INNODB_INTERNAL_INDEX uint16 = 11049
	ER_INNODB_IDX_CNT_MORE_THAN_DEFINED_IN_MYSQL uint16 = 11050
	ER_INNODB_IDX_CNT_FEWER_THAN_DEFINED_IN_MYSQL uint16 = 11051
	ER_INNODB_IDX_COLUMN_CNT_DIFF uint16 = 11052
	ER_INNODB_USE_MONITOR_GROUP_NAME uint16 = 11053
	ER_INNODB_MONITOR_DEFAULT_VALUE_NOT_DEFINED uint16 = 11054
	ER_INNODB_MONITOR_IS_ENABLED uint16 = 11055
	ER_INNODB_INVALID_MONITOR_COUNTER_NAME uint16 = 11056
	ER_WIN_LOAD_LIBRARY_FAILED uint16 = 11057
	ER_PARTITION_HANDLER_ADMIN_MSG uint16 = 11058
	ER_RPL_RLI_INIT_INFO_MSG uint16 = 11059
	ER_DD_UPGRADE_TABLE_INTACT_ERROR uint16 = 11060
	ER_SERVER_INIT_COMPILED_IN_COMMANDS uint16 = 11061
	ER_MYISAM_CHECK_METHOD_ERROR uint16 = 11062
	ER_MYISAM_CRASHED_ERROR uint16 = 11063
	ER_WAITPID_FAILED uint16 = 11064
	ER_FAILED_TO_FIND_MYSQLD_STATUS uint16 = 11065
	ER_INNODB_ERROR_LOGGER_MSG uint16 = 11066
	ER_INNODB_ERROR_LOGGER_FATAL_MSG uint16 = 11067
	ER_DEPRECATED_SYNTAX_WITH_REPLACEMENT uint16 = 11068
	ER_DEPRECATED_SYNTAX_NO_REPLACEMENT uint16 = 11069
	ER_DEPRECATE_MSG_NO_REPLACEMENT uint16 = 11070
	ER_LOG_PRINTF_MSG uint16 = 11071
	ER_BINLOG_LOGGING_NOT_POSSIBLE uint16 = 11072
	ER_FAILED_TO_SET_PERSISTED_OPTIONS uint16 = 11073
	ER_COMPONENTS_FAILED_TO_ACQUIRE_SERVICE_IMPLEMENTATION uint16 = 11074
	ER_RES_GRP_INVALID_VCPU_RANGE uint16 = 11075
	ER_RES_GRP_INVALID_VCPU_ID uint16 = 11076
	ER_ERROR_DURING_FLUSH_LOG_COMMIT_PHASE uint16 = 11077
	ER_DROP_DATABASE_FAILED_RMDIR_MANUALLY uint16 = 11078
	ER_EXPIRE_LOGS_DAYS_IGNORED uint16 = 11079
	ER_BINLOG_MALFORMED_OR_OLD_RELAY_LOG uint16 = 11080
	ER_DD_UPGRADE_VIEW_COLUMN_NAME_TOO_LONG uint16 = 11081
	ER_TABLE_NEEDS_DUMP_UPGRADE uint16 = 11082
	ER_DD_UPGRADE_FAILED_TO_UPDATE_VER_NO_IN_TABLESPACE uint16 = 11083
	ER_KEYRING_MIGRATION_FAILED uint16 = 11084
	ER_KEYRING_MIGRATION_SUCCESSFUL uint16 = 11085
	ER_RESTART_RECEIVED_INFO uint16 = 11086
	ER_LCTN_CHANGED uint16 = 11087
	ER_DD_INITIALIZE uint16 = 11088
	ER_DD_RESTART uint16 = 11089
	ER_DD_UPGRADE uint16 = 11090
	ER_DD_UPGRADE_OFF uint16 = 11091
	ER_DD_UPGRADE_VERSION_NOT_SUPPORTED uint16 = 11092
	ER_DD_UPGRADE_SCHEMA_UNAVAILABLE uint16 = 11093
	ER_DD_MINOR_DOWNGRADE uint16 = 11094
	ER_DD_MINOR_DOWNGRADE_VERSION_NOT_SUPPORTED uint16 = 11095
	ER_DD_NO_VERSION_FOUND uint16 = 11096
	ER_THREAD_POOL_NOT_SUPPORTED_ON_PLATFORM uint16 = 11097
	ER_THREAD_POOL_SIZE_TOO_LOW uint16 = 11098
	ER_THREAD_POOL_SIZE_TOO_HIGH uint16 = 11099
	ER_THREAD_POOL_ALGORITHM_INVALID uint16 = 11100
	ER_THREAD_POOL_INVALID_STALL_LIMIT uint16 = 11101
	ER_THREAD_POOL_INVALID_PRIO_KICKUP_TIMER uint16 = 11102
	ER_THREAD_POOL_MAX_UNUSED_THREADS_INVALID uint16 = 11103
	ER_THREAD_POOL_CON_HANDLER_INIT_FAILED uint16 = 11104
	ER_THREAD_POOL_INIT_FAILED uint16 = 11105
	ER_THREAD_POOL_PLUGIN_STARTED uint16 = 11106
	ER_THREAD_POOL_CANNOT_SET_THREAD_SPECIFIC_DATA uint16 = 11107
	ER_THREAD_POOL_FAILED_TO_CREATE_CONNECT_HANDLER_THD uint16 = 11108
	ER_THREAD_POOL_FAILED_TO_CREATE_THD_AND_AUTH_CONN uint16 = 11109
	ER_THREAD_POOL_FAILED_PROCESS_CONNECT_EVENT uint16 = 11110
	ER_THREAD_POOL_FAILED_TO_CREATE_POOL uint16 = 11111
	ER_THREAD_POOL_RATE_LIMITED_ERROR_MSGS uint16 = 11112
	ER_TRHEAD_POOL_LOW_LEVEL_INIT_FAILED uint16 = 11113
	ER_THREAD_POOL_LOW_LEVEL_REARM_FAILED uint16 = 11114
	ER_THREAD_POOL_BUFFER_TOO_SMALL uint16 = 11115
	ER_MECAB_NOT_SUPPORTED uint16 = 11116
	ER_MECAB_NOT_VERIFIED uint16 = 11117
	ER_MECAB_CREATING_MODEL uint16 = 11118
	ER_MECAB_FAILED_TO_CREATE_MODEL uint16 = 11119
	ER_MECAB_FAILED_TO_CREATE_TRIGGER uint16 = 11120
	ER_MECAB_UNSUPPORTED_CHARSET uint16 = 11121
	ER_MECAB_CHARSET_LOADED uint16 = 11122
	ER_MECAB_PARSE_FAILED uint16 = 11123
	ER_MECAB_OOM_WHILE_PARSING_TEXT uint16 = 11124
	ER_MECAB_CREATE_LATTICE_FAILED uint16 = 11125
	ER_SEMISYNC_TRACE_ENTER_FUNC uint16 = 11126
	ER_SEMISYNC_TRACE_EXIT_WITH_INT_EXIT_CODE uint16 = 11127
	ER_SEMISYNC_TRACE_EXIT_WITH_BOOL_EXIT_CODE uint16 = 11128
	ER_SEMISYNC_TRACE_EXIT uint16 = 11129
	ER_SEMISYNC_RPL_INIT_FOR_TRX uint16 = 11130
	ER_SEMISYNC_FAILED_TO_ALLOCATE_TRX_NODE uint16 = 11131
	ER_SEMISYNC_BINLOG_WRITE_OUT_OF_ORDER uint16 = 11132
	ER_SEMISYNC_INSERT_LOG_INFO_IN_ENTRY uint16 = 11133
	ER_SEMISYNC_PROBE_LOG_INFO_IN_ENTRY uint16 = 11134
	ER_SEMISYNC_CLEARED_ALL_ACTIVE_TRANSACTION_NODES uint16 = 11135
	ER_SEMISYNC_CLEARED_ACTIVE_TRANSACTION_TILL_POS uint16 = 11136
	ER_SEMISYNC_REPLY_MAGIC_NO_ERROR uint16 = 11137
	ER_SEMISYNC_REPLY_PKT_LENGTH_TOO_SMALL uint16 = 11138
	ER_SEMISYNC_REPLY_BINLOG_FILE_TOO_LARGE uint16 = 11139
	ER_SEMISYNC_SERVER_REPLY uint16 = 11140
	ER_SEMISYNC_FUNCTION_CALLED_TWICE uint16 = 11141
	ER_SEMISYNC_RPL_ENABLED_ON_MASTER uint16 = 11142
	ER_SEMISYNC_MASTER_OOM uint16 = 11143
	ER_SEMISYNC_DISABLED_ON_MASTER uint16 = 11144
	ER_SEMISYNC_FORCED_SHUTDOWN uint16 = 11145
	ER_SEMISYNC_MASTER_GOT_REPLY_AT_POS uint16 = 11146
	ER_SEMISYNC_MASTER_SIGNAL_ALL_WAITING_THREADS uint16 = 11147
	ER_SEMISYNC_MASTER_TRX_WAIT_POS uint16 = 11148
	ER_SEMISYNC_BINLOG_REPLY_IS_AHEAD uint16 = 11149
	ER_SEMISYNC_MOVE_BACK_WAIT_POS uint16 = 11150
	ER_SEMISYNC_INIT_WAIT_POS uint16 = 11151
	ER_SEMISYNC_WAIT_TIME_FOR_BINLOG_SENT uint16 = 11152
	ER_SEMISYNC_WAIT_FOR_BINLOG_TIMEDOUT uint16 = 11153
	ER_SEMISYNC_WAIT_TIME_ASSESSMENT_FOR_COMMIT_TRX_FAILED uint16 = 11154
	ER_SEMISYNC_RPL_SWITCHED_OFF uint16 = 11155
	ER_SEMISYNC_RPL_SWITCHED_ON uint16 = 11156
	ER_SEMISYNC_NO_SPACE_IN_THE_PKT uint16 = 11157
	ER_SEMISYNC_SYNC_HEADER_UPDATE_INFO uint16 = 11158
	ER_SEMISYNC_FAILED_TO_INSERT_TRX_NODE uint16 = 11159
	ER_SEMISYNC_TRX_SKIPPED_AT_POS uint16 = 11160
	ER_SEMISYNC_MASTER_FAILED_ON_NET_FLUSH uint16 = 11161
	ER_SEMISYNC_RECEIVED_ACK_IS_SMALLER uint16 = 11162
	ER_SEMISYNC_ADD_ACK_TO_SLOT uint16 = 11163
	ER_SEMISYNC_UPDATE_EXISTING_SLAVE_ACK uint16 = 11164
	ER_SEMISYNC_FAILED_TO_START_ACK_RECEIVER_THD uint16 = 11165
	ER_SEMISYNC_STARTING_ACK_RECEIVER_THD uint16 = 11166
	ER_SEMISYNC_FAILED_TO_WAIT_ON_DUMP_SOCKET uint16 = 11167
	ER_SEMISYNC_STOPPING_ACK_RECEIVER_THREAD uint16 = 11168
	ER_SEMISYNC_FAILED_REGISTER_SLAVE_TO_RECEIVER uint16 = 11169
	ER_SEMISYNC_START_BINLOG_DUMP_TO_SLAVE uint16 = 11170
	ER_SEMISYNC_STOP_BINLOG_DUMP_TO_SLAVE uint16 = 11171
	ER_SEMISYNC_UNREGISTER_TRX_OBSERVER_FAILED uint16 = 11172
	ER_SEMISYNC_UNREGISTER_BINLOG_STORAGE_OBSERVER_FAILED uint16 = 11173
	ER_SEMISYNC_UNREGISTER_BINLOG_TRANSMIT_OBSERVER_FAILED uint16 = 11174
	ER_SEMISYNC_UNREGISTERED_REPLICATOR uint16 = 11175
	ER_SEMISYNC_SOCKET_FD_TOO_LARGE uint16 = 11176
	ER_SEMISYNC_SLAVE_REPLY uint16 = 11177
	ER_SEMISYNC_MISSING_MAGIC_NO_FOR_SEMISYNC_PKT uint16 = 11178
	ER_SEMISYNC_SLAVE_START uint16 = 11179
	ER_SEMISYNC_SLAVE_REPLY_WITH_BINLOG_INFO uint16 = 11180
	ER_SEMISYNC_SLAVE_NET_FLUSH_REPLY_FAILED uint16 = 11181
	ER_SEMISYNC_SLAVE_SEND_REPLY_FAILED uint16 = 11182
	ER_SEMISYNC_EXECUTION_FAILED_ON_MASTER uint16 = 11183
	ER_SEMISYNC_NOT_SUPPORTED_BY_MASTER uint16 = 11184
	ER_SEMISYNC_SLAVE_SET_FAILED uint16 = 11185
	ER_SEMISYNC_FAILED_TO_STOP_ACK_RECEIVER_THD uint16 = 11186
	ER_FIREWALL_FAILED_TO_READ_FIREWALL_TABLES uint16 = 11187
	ER_FIREWALL_FAILED_TO_REG_DYNAMIC_PRIVILEGES uint16 = 11188
	ER_FIREWALL_RECORDING_STMT_WAS_TRUNCATED uint16 = 11189
	ER_FIREWALL_RECORDING_STMT_WITHOUT_TEXT uint16 = 11190
	ER_FIREWALL_SUSPICIOUS_STMT uint16 = 11191
	ER_FIREWALL_ACCESS_DENIED uint16 = 11192
	ER_FIREWALL_SKIPPED_UNKNOWN_USER_MODE uint16 = 11193
	ER_FIREWALL_RELOADING_CACHE uint16 = 11194
	ER_FIREWALL_RESET_FOR_USER uint16 = 11195
	ER_FIREWALL_STATUS_FLUSHED uint16 = 11196
	ER_KEYRING_LOGGER_ERROR_MSG uint16 = 11197
	ER_AUDIT_LOG_FILTER_IS_NOT_INSTALLED uint16 = 11198
	ER_AUDIT_LOG_SWITCHING_TO_INCLUDE_LIST uint16 = 11199
	ER_AUDIT_LOG_CANNOT_SET_LOG_POLICY_WITH_OTHER_POLICIES uint16 = 11200
	ER_AUDIT_LOG_ONLY_INCLUDE_LIST_USED uint16 = 11201
	ER_AUDIT_LOG_INDEX_MAP_CANNOT_ACCESS_DIR uint16 = 11202
	ER_AUDIT_LOG_WRITER_RENAME_FILE_FAILED uint16 = 11203
	ER_AUDIT_LOG_WRITER_DEST_FILE_ALREADY_EXISTS uint16 = 11204
	ER_AUDIT_LOG_WRITER_RENAME_FILE_FAILED_REMOVE_FILE_MANUALLY uint16 = 11205
	ER_AUDIT_LOG_WRITER_INCOMPLETE_FILE_RENAMED uint16 = 11206
	ER_AUDIT_LOG_WRITER_FAILED_TO_WRITE_TO_FILE uint16 = 11207
	ER_AUDIT_LOG_EC_WRITER_FAILED_TO_INIT_ENCRYPTION uint16 = 11208
	ER_AUDIT_LOG_EC_WRITER_FAILED_TO_INIT_COMPRESSION uint16 = 11209
	ER_AUDIT_LOG_EC_WRITER_FAILED_TO_CREATE_FILE uint16 = 11210
	ER_AUDIT_LOG_RENAME_LOG_FILE_BEFORE_FLUSH uint16 = 11211
	ER_AUDIT_LOG_FILTER_RESULT_MSG uint16 = 11212
	ER_AUDIT_LOG_JSON_READER_FAILED_TO_PARSE uint16 = 11213
	ER_AUDIT_LOG_JSON_READER_BUF_TOO_SMALL uint16 = 11214
	ER_AUDIT_LOG_JSON_READER_FAILED_TO_OPEN_FILE uint16 = 11215
	ER_AUDIT_LOG_JSON_READER_FILE_PARSING_ERROR uint16 = 11216
	ER_AUDIT_LOG_FILTER_INVALID_COLUMN_COUNT uint16 = 11217
	ER_AUDIT_LOG_FILTER_INVALID_COLUMN_DEFINITION uint16 = 11218
	ER_AUDIT_LOG_FILTER_FAILED_TO_STORE_TABLE_FLDS uint16 = 11219
	ER_AUDIT_LOG_FILTER_FAILED_TO_UPDATE_TABLE uint16 = 11220
	ER_AUDIT_LOG_FILTER_FAILED_TO_INSERT_INTO_TABLE uint16 = 11221
	ER_AUDIT_LOG_FILTER_FAILED_TO_DELETE_FROM_TABLE uint16 = 11222
	ER_AUDIT_LOG_FILTER_FAILED_TO_INIT_TABLE_FOR_READ uint16 = 11223
	ER_AUDIT_LOG_FILTER_FAILED_TO_READ_TABLE uint16 = 11224
	ER_AUDIT_LOG_FILTER_FAILED_TO_CLOSE_TABLE_AFTER_READING uint16 = 11225
	ER_AUDIT_LOG_FILTER_USER_AND_HOST_CANNOT_BE_EMPTY uint16 = 11226
	ER_AUDIT_LOG_FILTER_FLD_FILTERNAME_CANNOT_BE_EMPTY uint16 = 11227
	ER_VALIDATE_PWD_DICT_FILE_NOT_SPECIFIED uint16 = 11228
	ER_VALIDATE_PWD_DICT_FILE_NOT_LOADED uint16 = 11229
	ER_VALIDATE_PWD_DICT_FILE_TOO_BIG uint16 = 11230
	ER_VALIDATE_PWD_FAILED_TO_READ_DICT_FILE uint16 = 11231
	ER_VALIDATE_PWD_FAILED_TO_GET_FLD_FROM_SECURITY_CTX uint16 = 11232
	ER_VALIDATE_PWD_FAILED_TO_GET_SECURITY_CTX uint16 = 11233
	ER_VALIDATE_PWD_LENGTH_CHANGED uint16 = 11234
	ER_REWRITER_QUERY_ERROR_MSG uint16 = 11235
	ER_REWRITER_QUERY_FAILED uint16 = 11236
	ER_XPLUGIN_STARTUP_FAILED uint16 = 11237
	//OBSOLETE_ER_XPLUGIN_SERVER_EXITING uint16 = 11238
	//OBSOLETE_ER_XPLUGIN_SERVER_EXITED uint16 = 11239
	ER_XPLUGIN_USING_SSL_CONF_FROM_SERVER uint16 = 11240
	ER_XPLUGIN_USING_SSL_CONF_FROM_MYSQLX uint16 = 11241
	ER_XPLUGIN_FAILED_TO_USE_SSL_CONF uint16 = 11242
	ER_XPLUGIN_USING_SSL_FOR_TLS_CONNECTION uint16 = 11243
	ER_XPLUGIN_REFERENCE_TO_SECURE_CONN_WITH_XPLUGIN uint16 = 11244
	ER_XPLUGIN_ERROR_MSG uint16 = 11245
	ER_SHA_PWD_FAILED_TO_PARSE_AUTH_STRING uint16 = 11246
	ER_SHA_PWD_FAILED_TO_GENERATE_MULTI_ROUND_HASH uint16 = 11247
	ER_SHA_PWD_AUTH_REQUIRES_RSA_OR_SSL uint16 = 11248
	ER_SHA_PWD_RSA_KEY_TOO_LONG uint16 = 11249
	ER_PLUGIN_COMMON_FAILED_TO_OPEN_FILTER_TABLES uint16 = 11250
	ER_PLUGIN_COMMON_FAILED_TO_OPEN_TABLE uint16 = 11251
	ER_AUTH_LDAP_ERROR_LOGGER_ERROR_MSG uint16 = 11252
	ER_CONN_CONTROL_ERROR_MSG uint16 = 11253
	ER_GRP_RPL_ERROR_MSG uint16 = 11254
	ER_SHA_PWD_SALT_FOR_USER_CORRUPT uint16 = 11255
	ER_SYS_VAR_COMPONENT_OOM uint16 = 11256
	ER_SYS_VAR_COMPONENT_VARIABLE_SET_READ_ONLY uint16 = 11257
	ER_SYS_VAR_COMPONENT_UNKNOWN_VARIABLE_TYPE uint16 = 11258
	ER_SYS_VAR_COMPONENT_FAILED_TO_PARSE_VARIABLE_OPTIONS uint16 = 11259
	ER_SYS_VAR_COMPONENT_FAILED_TO_MAKE_VARIABLE_PERSISTENT uint16 = 11260
	ER_COMPONENT_FILTER_CONFUSED uint16 = 11261
	ER_STOP_SLAVE_IO_THREAD_DISK_SPACE uint16 = 11262
	ER_LOG_FILE_CANNOT_OPEN uint16 = 11263
	//OBSOLETE_ER_UNABLE_TO_COLLECT_LOG_STATUS uint16 = 11264
	//OBSOLETE_ER_DEPRECATED_UTF8_ALIAS uint16 = 11265
	//OBSOLETE_ER_DEPRECATED_NATIONAL uint16 = 11266
	//OBSOLETE_ER_SLAVE_POSSIBLY_DIVERGED_AFTER_DDL uint16 = 11267
	ER_PERSIST_OPTION_STATUS uint16 = 11268
	ER_NOT_IMPLEMENTED_GET_TABLESPACE_STATISTICS uint16 = 11269
	//OBSOLETE_ER_UNABLE_TO_SET_OPTION uint16 = 11270
	//OBSOLETE_ER_RESERVED_TABLESPACE_NAME uint16 = 11271
	ER_SSL_FIPS_MODE_ERROR uint16 = 11272
	ER_CONN_INIT_CONNECT_IGNORED uint16 = 11273
	//OBSOLETE_ER_UNSUPPORTED_SQL_MODE uint16 = 11274
	ER_REWRITER_OOM uint16 = 11275
	ER_REWRITER_TABLE_MALFORMED_ERROR uint16 = 11276
	ER_REWRITER_LOAD_FAILED uint16 = 11277
	ER_REWRITER_READ_FAILED uint16 = 11278
	ER_CONN_CONTROL_EVENT_COORDINATOR_INIT_FAILED uint16 = 11279
	ER_CONN_CONTROL_STAT_CONN_DELAY_TRIGGERED_UPDATE_FAILED uint16 = 11280
	ER_CONN_CONTROL_STAT_CONN_DELAY_TRIGGERED_RESET_FAILED uint16 = 11281
	ER_CONN_CONTROL_INVALID_CONN_DELAY_TYPE uint16 = 11282
	ER_CONN_CONTROL_DELAY_ACTION_INIT_FAILED uint16 = 11283
	ER_CONN_CONTROL_FAILED_TO_SET_CONN_DELAY uint16 = 11284
	ER_CONN_CONTROL_FAILED_TO_UPDATE_CONN_DELAY_HASH uint16 = 11285
	ER_XPLUGIN_FORCE_STOP_CLIENT uint16 = 11286
	ER_XPLUGIN_MAX_AUTH_ATTEMPTS_REACHED uint16 = 11287
	ER_XPLUGIN_BUFFER_PAGE_ALLOC_FAILED uint16 = 11288
	ER_XPLUGIN_DETECTED_HANGING_CLIENTS uint16 = 11289
	ER_XPLUGIN_FAILED_TO_ACCEPT_CLIENT uint16 = 11290
	ER_XPLUGIN_FAILED_TO_SCHEDULE_CLIENT uint16 = 11291
	ER_XPLUGIN_FAILED_TO_PREPARE_IO_INTERFACES uint16 = 11292
	ER_XPLUGIN_SRV_SESSION_INIT_THREAD_FAILED uint16 = 11293
	ER_XPLUGIN_UNABLE_TO_USE_USER_SESSION_ACCOUNT uint16 = 11294
	ER_XPLUGIN_REFERENCE_TO_USER_ACCOUNT_DOC_SECTION uint16 = 11295
	ER_XPLUGIN_UNEXPECTED_EXCEPTION_DISPATCHING_CMD uint16 = 11296
	ER_XPLUGIN_EXCEPTION_IN_TASK_SCHEDULER uint16 = 11297
	ER_XPLUGIN_TASK_SCHEDULING_FAILED uint16 = 11298
	ER_XPLUGIN_EXCEPTION_IN_EVENT_LOOP uint16 = 11299
	ER_XPLUGIN_LISTENER_SETUP_FAILED uint16 = 11300
	ER_XPLUING_NET_STARTUP_FAILED uint16 = 11301
	ER_XPLUGIN_FAILED_AT_SSL_CONF uint16 = 11302
	//OBSOLETE_ER_XPLUGIN_CLIENT_SSL_HANDSHAKE_FAILED uint16 = 11303
	//OBSOLETE_ER_XPLUGIN_SSL_HANDSHAKE_WITH_SERVER_FAILED uint16 = 11304
	ER_XPLUGIN_FAILED_TO_CREATE_SESSION_FOR_CONN uint16 = 11305
	ER_XPLUGIN_FAILED_TO_INITIALIZE_SESSION uint16 = 11306
	ER_XPLUGIN_MESSAGE_TOO_LONG uint16 = 11307
	ER_XPLUGIN_UNINITIALIZED_MESSAGE uint16 = 11308
	ER_XPLUGIN_FAILED_TO_SET_MIN_NUMBER_OF_WORKERS uint16 = 11309
	ER_XPLUGIN_UNABLE_TO_ACCEPT_CONNECTION uint16 = 11310
	ER_XPLUGIN_ALL_IO_INTERFACES_DISABLED uint16 = 11311
	//OBSOLETE_ER_XPLUGIN_INVALID_MSG_DURING_CLIENT_INIT uint16 = 11312
	//OBSOLETE_ER_XPLUGIN_CLOSING_CLIENTS_ON_SHUTDOWN uint16 = 11313
	ER_XPLUGIN_ERROR_READING_SOCKET uint16 = 11314
	ER_XPLUGIN_PEER_DISCONNECTED_WHILE_READING_MSG_BODY uint16 = 11315
	ER_XPLUGIN_READ_FAILED_CLOSING_CONNECTION uint16 = 11316
	//OBSOLETE_ER_XPLUGIN_INVALID_AUTH_METHOD uint16 = 11317
	//OBSOLETE_ER_XPLUGIN_UNEXPECTED_MSG_DURING_AUTHENTICATION uint16 = 11318
	//OBSOLETE_ER_XPLUGIN_ERROR_WRITING_TO_CLIENT uint16 = 11319
	//OBSOLETE_ER_XPLUGIN_SCHEDULER_STARTED uint16 = 11320
	//OBSOLETE_ER_XPLUGIN_SCHEDULER_STOPPED uint16 = 11321
	ER_XPLUGIN_LISTENER_SYS_VARIABLE_ERROR uint16 = 11322
	ER_XPLUGIN_LISTENER_STATUS_MSG uint16 = 11323
	ER_XPLUGIN_RETRYING_BIND_ON_PORT uint16 = 11324
	//OBSOLETE_ER_XPLUGIN_SHUTDOWN_TRIGGERED uint16 = 11325
	//OBSOLETE_ER_XPLUGIN_USER_ACCOUNT_WITH_ALL_PERMISSIONS uint16 = 11326
	ER_XPLUGIN_EXISTING_USER_ACCOUNT_WITH_INCOMPLETE_GRANTS uint16 = 11327
	//OBSOLETE_ER_XPLUGIN_SERVER_STARTS_HANDLING_CONNECTIONS uint16 = 11328
	//OBSOLETE_ER_XPLUGIN_SERVER_STOPPED_HANDLING_CONNECTIONS uint16 = 11329
	//OBSOLETE_ER_XPLUGIN_FAILED_TO_INTERRUPT_SESSION uint16 = 11330
	//OBSOLETE_ER_XPLUGIN_CLIENT_RELEASE_TRIGGERED uint16 = 11331
	ER_XPLUGIN_IPv6_AVAILABLE uint16 = 11332
	//OBSOLETE_ER_XPLUGIN_UNIX_SOCKET_NOT_CONFIGURED uint16 = 11333
	ER_XPLUGIN_CLIENT_KILL_MSG uint16 = 11334
	ER_XPLUGIN_FAILED_TO_GET_SECURITY_CTX uint16 = 11335
	//OBSOLETE_ER_XPLUGIN_FAILED_TO_SWITCH_SECURITY_CTX_TO_ROOT uint16 = 11336
	ER_XPLUGIN_FAILED_TO_CLOSE_SQL_SESSION uint16 = 11337
	ER_XPLUGIN_FAILED_TO_EXECUTE_ADMIN_CMD uint16 = 11338
	ER_XPLUGIN_EMPTY_ADMIN_CMD uint16 = 11339
	ER_XPLUGIN_FAILED_TO_GET_SYS_VAR uint16 = 11340
	ER_XPLUGIN_FAILED_TO_GET_CREATION_STMT uint16 = 11341
	ER_XPLUGIN_FAILED_TO_GET_ENGINE_INFO uint16 = 11342
	//OBSOLETE_ER_XPLUGIN_FAIL_TO_GET_RESULT_DATA uint16 = 11343
	//OBSOLETE_ER_XPLUGIN_CAPABILITY_EXPIRED_PASSWORD uint16 = 11344
	ER_XPLUGIN_FAILED_TO_SET_SO_REUSEADDR_FLAG uint16 = 11345
	ER_XPLUGIN_FAILED_TO_OPEN_INTERNAL_SESSION uint16 = 11346
	ER_XPLUGIN_FAILED_TO_SWITCH_CONTEXT uint16 = 11347
	ER_XPLUGIN_FAILED_TO_UNREGISTER_UDF uint16 = 11348
	//OBSOLETE_ER_XPLUGIN_GET_PEER_ADDRESS_FAILED uint16 = 11349
	//OBSOLETE_ER_XPLUGIN_CAPABILITY_CLIENT_INTERACTIVE_FAILED uint16 = 11350
	ER_XPLUGIN_FAILED_TO_RESET_IPV6_V6ONLY_FLAG uint16 = 11351
	ER_KEYRING_INVALID_KEY_TYPE uint16 = 11352
	ER_KEYRING_INVALID_KEY_LENGTH uint16 = 11353
	ER_KEYRING_FAILED_TO_CREATE_KEYRING_DIR uint16 = 11354
	ER_KEYRING_FILE_INIT_FAILED uint16 = 11355
	ER_KEYRING_INTERNAL_EXCEPTION_FAILED_FILE_INIT uint16 = 11356
	ER_KEYRING_FAILED_TO_GENERATE_KEY uint16 = 11357
	ER_KEYRING_CHECK_KEY_FAILED_DUE_TO_INVALID_KEY uint16 = 11358
	ER_KEYRING_CHECK_KEY_FAILED_DUE_TO_EMPTY_KEY_ID uint16 = 11359
	ER_KEYRING_OPERATION_FAILED_DUE_TO_INTERNAL_ERROR uint16 = 11360
	ER_KEYRING_INCORRECT_FILE uint16 = 11361
	ER_KEYRING_FOUND_MALFORMED_BACKUP_FILE uint16 = 11362
	ER_KEYRING_FAILED_TO_RESTORE_FROM_BACKUP_FILE uint16 = 11363
	ER_KEYRING_FAILED_TO_FLUSH_KEYRING_TO_FILE uint16 = 11364
	ER_KEYRING_FAILED_TO_GET_FILE_STAT uint16 = 11365
	ER_KEYRING_FAILED_TO_REMOVE_FILE uint16 = 11366
	ER_KEYRING_FAILED_TO_TRUNCATE_FILE uint16 = 11367
	ER_KEYRING_UNKNOWN_ERROR uint16 = 11368
	ER_KEYRING_FAILED_TO_SET_KEYRING_FILE_DATA uint16 = 11369
	ER_KEYRING_FILE_IO_ERROR uint16 = 11370
	ER_KEYRING_FAILED_TO_LOAD_KEYRING_CONTENT uint16 = 11371
	ER_KEYRING_FAILED_TO_FLUSH_KEYS_TO_KEYRING uint16 = 11372
	ER_KEYRING_FAILED_TO_FLUSH_KEYS_TO_KEYRING_BACKUP uint16 = 11373
	ER_KEYRING_KEY_FETCH_FAILED_DUE_TO_EMPTY_KEY_ID uint16 = 11374
	ER_KEYRING_FAILED_TO_REMOVE_KEY_DUE_TO_EMPTY_ID uint16 = 11375
	ER_KEYRING_OKV_INCORRECT_KEY_VAULT_CONFIGURED uint16 = 11376
	ER_KEYRING_OKV_INIT_FAILED_DUE_TO_INCORRECT_CONF uint16 = 11377
	ER_KEYRING_OKV_INIT_FAILED_DUE_TO_INTERNAL_ERROR uint16 = 11378
	ER_KEYRING_OKV_INVALID_KEY_TYPE uint16 = 11379
	ER_KEYRING_OKV_INVALID_KEY_LENGTH_FOR_CIPHER uint16 = 11380
	ER_KEYRING_OKV_FAILED_TO_GENERATE_KEY_DUE_TO_INTERNAL_ERROR uint16 = 11381
	ER_KEYRING_OKV_FAILED_TO_FIND_SERVER_ENTRY uint16 = 11382
	ER_KEYRING_OKV_FAILED_TO_FIND_STANDBY_SERVER_ENTRY uint16 = 11383
	ER_KEYRING_OKV_FAILED_TO_PARSE_CONF_FILE uint16 = 11384
	ER_KEYRING_OKV_FAILED_TO_LOAD_KEY_UID uint16 = 11385
	ER_KEYRING_OKV_FAILED_TO_INIT_SSL_LAYER uint16 = 11386
	ER_KEYRING_OKV_FAILED_TO_INIT_CLIENT uint16 = 11387
	ER_KEYRING_OKV_CONNECTION_TO_SERVER_FAILED uint16 = 11388
	ER_KEYRING_OKV_FAILED_TO_REMOVE_KEY uint16 = 11389
	ER_KEYRING_OKV_FAILED_TO_ADD_ATTRIBUTE uint16 = 11390
	ER_KEYRING_OKV_FAILED_TO_GENERATE_KEY uint16 = 11391
	ER_KEYRING_OKV_FAILED_TO_STORE_KEY uint16 = 11392
	ER_KEYRING_OKV_FAILED_TO_ACTIVATE_KEYS uint16 = 11393
	ER_KEYRING_OKV_FAILED_TO_FETCH_KEY uint16 = 11394
	ER_KEYRING_OKV_FAILED_TO_STORE_OR_GENERATE_KEY uint16 = 11395
	ER_KEYRING_OKV_FAILED_TO_RETRIEVE_KEY_SIGNATURE uint16 = 11396
	ER_KEYRING_OKV_FAILED_TO_RETRIEVE_KEY uint16 = 11397
	ER_KEYRING_OKV_FAILED_TO_LOAD_SSL_TRUST_STORE uint16 = 11398
	ER_KEYRING_OKV_FAILED_TO_SET_CERTIFICATE_FILE uint16 = 11399
	ER_KEYRING_OKV_FAILED_TO_SET_KEY_FILE uint16 = 11400
	ER_KEYRING_OKV_KEY_MISMATCH uint16 = 11401
	ER_KEYRING_ENCRYPTED_FILE_INCORRECT_KEYRING_FILE uint16 = 11402
	ER_KEYRING_ENCRYPTED_FILE_DECRYPTION_FAILED uint16 = 11403
	ER_KEYRING_ENCRYPTED_FILE_FOUND_MALFORMED_BACKUP_FILE uint16 = 11404
	ER_KEYRING_ENCRYPTED_FILE_FAILED_TO_RESTORE_KEYRING uint16 = 11405
	ER_KEYRING_ENCRYPTED_FILE_FAILED_TO_FLUSH_KEYRING uint16 = 11406
	ER_KEYRING_ENCRYPTED_FILE_ENCRYPTION_FAILED uint16 = 11407
	ER_KEYRING_ENCRYPTED_FILE_INVALID_KEYRING_DIR uint16 = 11408
	ER_KEYRING_ENCRYPTED_FILE_FAILED_TO_CREATE_KEYRING_DIR uint16 = 11409
	ER_KEYRING_ENCRYPTED_FILE_PASSWORD_IS_INVALID uint16 = 11410
	ER_KEYRING_ENCRYPTED_FILE_PASSWORD_IS_TOO_LONG uint16 = 11411
	ER_KEYRING_ENCRYPTED_FILE_INIT_FAILURE uint16 = 11412
	ER_KEYRING_ENCRYPTED_FILE_INIT_FAILED_DUE_TO_INTERNAL_ERROR uint16 = 11413
	ER_KEYRING_ENCRYPTED_FILE_GEN_KEY_FAILED_DUE_TO_INTERNAL_ERROR uint16 = 11414
	ER_KEYRING_AWS_FAILED_TO_SET_CMK_ID uint16 = 11415
	ER_KEYRING_AWS_FAILED_TO_SET_REGION uint16 = 11416
	ER_KEYRING_AWS_FAILED_TO_OPEN_CONF_FILE uint16 = 11417
	ER_KEYRING_AWS_FAILED_TO_ACCESS_KEY_ID_FROM_CONF_FILE uint16 = 11418
	ER_KEYRING_AWS_FAILED_TO_ACCESS_KEY_FROM_CONF_FILE uint16 = 11419
	ER_KEYRING_AWS_INVALID_CONF_FILE_PATH uint16 = 11420
	ER_KEYRING_AWS_INVALID_DATA_FILE_PATH uint16 = 11421
	ER_KEYRING_AWS_FAILED_TO_ACCESS_OR_CREATE_KEYRING_DIR uint16 = 11422
	ER_KEYRING_AWS_FAILED_TO_ACCESS_OR_CREATE_KEYRING_DATA_FILE uint16 = 11423
	ER_KEYRING_AWS_FAILED_TO_INIT_DUE_TO_INTERNAL_ERROR uint16 = 11424
	ER_KEYRING_AWS_FAILED_TO_ACCESS_DATA_FILE uint16 = 11425
	ER_KEYRING_AWS_CMK_ID_NOT_SET uint16 = 11426
	ER_KEYRING_AWS_FAILED_TO_GET_KMS_CREDENTIAL_FROM_CONF_FILE uint16 = 11427
	ER_KEYRING_AWS_INIT_FAILURE uint16 = 11428
	ER_KEYRING_AWS_FAILED_TO_INIT_DUE_TO_PLUGIN_INTERNAL_ERROR uint16 = 11429
	ER_KEYRING_AWS_INVALID_KEY_LENGTH_FOR_CIPHER uint16 = 11430
	ER_KEYRING_AWS_FAILED_TO_GENERATE_KEY_DUE_TO_INTERNAL_ERROR uint16 = 11431
	ER_KEYRING_AWS_INCORRECT_FILE uint16 = 11432
	ER_KEYRING_AWS_FOUND_MALFORMED_BACKUP_FILE uint16 = 11433
	ER_KEYRING_AWS_FAILED_TO_RESTORE_FROM_BACKUP_FILE uint16 = 11434
	ER_KEYRING_AWS_FAILED_TO_FLUSH_KEYRING_TO_FILE uint16 = 11435
	ER_KEYRING_AWS_INCORRECT_REGION uint16 = 11436
	ER_KEYRING_AWS_FAILED_TO_CONNECT_KMS uint16 = 11437
	ER_KEYRING_AWS_FAILED_TO_GENERATE_NEW_KEY uint16 = 11438
	ER_KEYRING_AWS_FAILED_TO_ENCRYPT_KEY uint16 = 11439
	ER_KEYRING_AWS_FAILED_TO_RE_ENCRYPT_KEY uint16 = 11440
	ER_KEYRING_AWS_FAILED_TO_DECRYPT_KEY uint16 = 11441
	ER_KEYRING_AWS_FAILED_TO_ROTATE_CMK uint16 = 11442
	ER_GRP_RPL_GTID_ALREADY_USED uint16 = 11443
	ER_GRP_RPL_APPLIER_THD_KILLED uint16 = 11444
	ER_GRP_RPL_EVENT_HANDLING_ERROR uint16 = 11445
	ER_GRP_RPL_ERROR_GTID_EXECUTION_INFO uint16 = 11446
	ER_GRP_RPL_CERTIFICATE_SIZE_ERROR uint16 = 11447
	ER_GRP_RPL_CREATE_APPLIER_CACHE_ERROR uint16 = 11448
	ER_GRP_RPL_UNBLOCK_WAITING_THD uint16 = 11449
	ER_GRP_RPL_APPLIER_PIPELINE_NOT_DISPOSED uint16 = 11450
	ER_GRP_RPL_APPLIER_THD_EXECUTION_ABORTED uint16 = 11451
	ER_GRP_RPL_APPLIER_EXECUTION_FATAL_ERROR uint16 = 11452
	ER_GRP_RPL_ERROR_STOPPING_CHANNELS uint16 = 11453
	ER_GRP_RPL_ERROR_SENDING_SINGLE_PRIMARY_MSSG uint16 = 11454
	ER_GRP_RPL_UPDATE_TRANS_SNAPSHOT_VER_ERROR uint16 = 11455
	ER_GRP_RPL_SIDNO_FETCH_ERROR uint16 = 11456
	ER_GRP_RPL_BROADCAST_COMMIT_TRANS_MSSG_FAILED uint16 = 11457
	ER_GRP_RPL_GROUP_NAME_PARSE_ERROR uint16 = 11458
	ER_GRP_RPL_ADD_GRPSID_TO_GRPGTIDSID_MAP_ERROR uint16 = 11459
	ER_GRP_RPL_UPDATE_GRPGTID_EXECUTED_ERROR uint16 = 11460
	ER_GRP_RPL_DONOR_TRANS_INFO_ERROR uint16 = 11461
	ER_GRP_RPL_SERVER_CONN_ERROR uint16 = 11462
	ER_GRP_RPL_ERROR_FETCHING_GTID_EXECUTED_SET uint16 = 11463
	ER_GRP_RPL_ADD_GTID_TO_GRPGTID_EXECUTED_ERROR uint16 = 11464
	ER_GRP_RPL_ERROR_FETCHING_GTID_SET uint16 = 11465
	ER_GRP_RPL_ADD_RETRIEVED_SET_TO_GRP_GTID_EXECUTED_ERROR uint16 = 11466
	ER_GRP_RPL_CERTIFICATION_INITIALIZATION_FAILURE uint16 = 11467
	ER_GRP_RPL_UPDATE_LAST_CONFLICT_FREE_TRANS_ERROR uint16 = 11468
	ER_GRP_RPL_UPDATE_TRANS_SNAPSHOT_REF_VER_ERROR uint16 = 11469
	ER_GRP_RPL_FETCH_TRANS_SIDNO_ERROR uint16 = 11470
	ER_GRP_RPL_ERROR_VERIFYING_SIDNO uint16 = 11471
	ER_GRP_RPL_CANT_GENERATE_GTID uint16 = 11472
	ER_GRP_RPL_INVALID_GTID_SET uint16 = 11473
	ER_GRP_RPL_UPDATE_GTID_SET_ERROR uint16 = 11474
	ER_GRP_RPL_RECEIVED_SET_MISSING_GTIDS uint16 = 11475
	ER_GRP_RPL_SKIP_COMPUTATION_TRANS_COMMITTED uint16 = 11476
	ER_GRP_RPL_NULL_PACKET uint16 = 11477
	ER_GRP_RPL_CANT_READ_GTID uint16 = 11478
	ER_GRP_RPL_PROCESS_GTID_SET_ERROR uint16 = 11479
	ER_GRP_RPL_PROCESS_INTERSECTION_GTID_SET_ERROR uint16 = 11480
	ER_GRP_RPL_SET_STABLE_TRANS_ERROR uint16 = 11481
	ER_GRP_RPL_CANT_READ_GRP_GTID_EXTRACTED uint16 = 11482
	ER_GRP_RPL_CANT_READ_WRITE_SET_ITEM uint16 = 11483
	ER_GRP_RPL_INIT_CERTIFICATION_INFO_FAILURE uint16 = 11484
	ER_GRP_RPL_CONFLICT_DETECTION_DISABLED uint16 = 11485
	ER_GRP_RPL_MSG_DISCARDED uint16 = 11486
	ER_GRP_RPL_MISSING_GRP_RPL_APPLIER uint16 = 11487
	ER_GRP_RPL_CERTIFIER_MSSG_PROCESS_ERROR uint16 = 11488
	ER_GRP_RPL_SRV_NOT_ONLINE uint16 = 11489
	ER_GRP_RPL_SRV_ONLINE uint16 = 11490
	ER_GRP_RPL_DISABLE_SRV_READ_MODE_RESTRICTED uint16 = 11491
	ER_GRP_RPL_MEM_ONLINE uint16 = 11492
	ER_GRP_RPL_MEM_UNREACHABLE uint16 = 11493
	ER_GRP_RPL_MEM_REACHABLE uint16 = 11494
	ER_GRP_RPL_SRV_BLOCKED uint16 = 11495
	ER_GRP_RPL_SRV_BLOCKED_FOR_SECS uint16 = 11496
	ER_GRP_RPL_CHANGE_GRP_MEM_NOT_PROCESSED uint16 = 11497
	ER_GRP_RPL_MEMBER_CONTACT_RESTORED uint16 = 11498
	ER_GRP_RPL_MEMBER_REMOVED uint16 = 11499
	ER_GRP_RPL_PRIMARY_MEMBER_LEFT_GRP uint16 = 11500
	ER_GRP_RPL_MEMBER_ADDED uint16 = 11501
	ER_GRP_RPL_MEMBER_EXIT_PLUGIN_ERROR uint16 = 11502
	ER_GRP_RPL_MEMBER_CHANGE uint16 = 11503
	ER_GRP_RPL_MEMBER_LEFT_GRP uint16 = 11504
	ER_GRP_RPL_MEMBER_EXPELLED uint16 = 11505
	ER_GRP_RPL_SESSION_OPEN_FAILED uint16 = 11506
	ER_GRP_RPL_NEW_PRIMARY_ELECTED uint16 = 11507
	ER_GRP_RPL_DISABLE_READ_ONLY_FAILED uint16 = 11508
	ER_GRP_RPL_ENABLE_READ_ONLY_FAILED uint16 = 11509
	ER_GRP_RPL_SRV_PRIMARY_MEM uint16 = 11510
	ER_GRP_RPL_SRV_SECONDARY_MEM uint16 = 11511
	ER_GRP_RPL_NO_SUITABLE_PRIMARY_MEM uint16 = 11512
	ER_GRP_RPL_SUPER_READ_ONLY_ACTIVATE_ERROR uint16 = 11513
	ER_GRP_RPL_EXCEEDS_AUTO_INC_VALUE uint16 = 11514
	ER_GRP_RPL_DATA_NOT_PROVIDED_BY_MEM uint16 = 11515
	ER_GRP_RPL_MEMBER_ALREADY_EXISTS uint16 = 11516
	ER_GRP_RPL_GRP_CHANGE_INFO_EXTRACT_ERROR uint16 = 11517
	ER_GRP_RPL_GTID_EXECUTED_EXTRACT_ERROR uint16 = 11518
	ER_GRP_RPL_GTID_SET_EXTRACT_ERROR uint16 = 11519
	ER_GRP_RPL_START_FAILED uint16 = 11520
	ER_GRP_RPL_MEMBER_VER_INCOMPATIBLE uint16 = 11521
	ER_GRP_RPL_TRANS_NOT_PRESENT_IN_GRP uint16 = 11522
	ER_GRP_RPL_TRANS_GREATER_THAN_GRP uint16 = 11523
	ER_GRP_RPL_MEMBER_VERSION_LOWER_THAN_GRP uint16 = 11524
	ER_GRP_RPL_LOCAL_GTID_SETS_PROCESS_ERROR uint16 = 11525
	ER_GRP_RPL_MEMBER_TRANS_GREATER_THAN_GRP uint16 = 11526
	ER_GRP_RPL_BLOCK_SIZE_DIFF_FROM_GRP uint16 = 11527
	ER_GRP_RPL_TRANS_WRITE_SET_EXTRACT_DIFF_FROM_GRP uint16 = 11528
	ER_GRP_RPL_MEMBER_CFG_INCOMPATIBLE_WITH_GRP_CFG uint16 = 11529
	ER_GRP_RPL_MEMBER_STOP_RPL_CHANNELS_ERROR uint16 = 11530
	ER_GRP_RPL_PURGE_APPLIER_LOGS uint16 = 11531
	ER_GRP_RPL_RESET_APPLIER_MODULE_LOGS_ERROR uint16 = 11532
	ER_GRP_RPL_APPLIER_THD_SETUP_ERROR uint16 = 11533
	ER_GRP_RPL_APPLIER_THD_START_ERROR uint16 = 11534
	ER_GRP_RPL_APPLIER_THD_STOP_ERROR uint16 = 11535
	ER_GRP_RPL_FETCH_TRANS_DATA_FAILED uint16 = 11536
	ER_GRP_RPL_SLAVE_IO_THD_PRIMARY_UNKNOWN uint16 = 11537
	ER_GRP_RPL_SALVE_IO_THD_ON_SECONDARY_MEMBER uint16 = 11538
	ER_GRP_RPL_SLAVE_SQL_THD_PRIMARY_UNKNOWN uint16 = 11539
	ER_GRP_RPL_SLAVE_SQL_THD_ON_SECONDARY_MEMBER uint16 = 11540
	ER_GRP_RPL_NEEDS_INNODB_TABLE uint16 = 11541
	ER_GRP_RPL_PRIMARY_KEY_NOT_DEFINED uint16 = 11542
	ER_GRP_RPL_FK_WITH_CASCADE_UNSUPPORTED uint16 = 11543
	ER_GRP_RPL_AUTO_INC_RESET uint16 = 11544
	ER_GRP_RPL_AUTO_INC_OFFSET_RESET uint16 = 11545
	ER_GRP_RPL_AUTO_INC_SET uint16 = 11546
	ER_GRP_RPL_AUTO_INC_OFFSET_SET uint16 = 11547
	ER_GRP_RPL_FETCH_TRANS_CONTEXT_FAILED uint16 = 11548
	ER_GRP_RPL_FETCH_FORMAT_DESC_LOG_EVENT_FAILED uint16 = 11549
	ER_GRP_RPL_FETCH_TRANS_CONTEXT_LOG_EVENT_FAILED uint16 = 11550
	ER_GRP_RPL_FETCH_SNAPSHOT_VERSION_FAILED uint16 = 11551
	ER_GRP_RPL_FETCH_GTID_LOG_EVENT_FAILED uint16 = 11552
	ER_GRP_RPL_UPDATE_SERV_CERTIFICATE_FAILED uint16 = 11553
	ER_GRP_RPL_ADD_GTID_INFO_WITH_LOCAL_GTID_FAILED uint16 = 11554
	ER_GRP_RPL_ADD_GTID_INFO_WITHOUT_LOCAL_GTID_FAILED uint16 = 11555
	ER_GRP_RPL_NOTIFY_CERTIFICATION_OUTCOME_FAILED uint16 = 11556
	ER_GRP_RPL_ADD_GTID_INFO_WITH_REMOTE_GTID_FAILED uint16 = 11557
	ER_GRP_RPL_ADD_GTID_INFO_WITHOUT_REMOTE_GTID_FAILED uint16 = 11558
	ER_GRP_RPL_FETCH_VIEW_CHANGE_LOG_EVENT_FAILED uint16 = 11559
	ER_GRP_RPL_CONTACT_WITH_SRV_FAILED uint16 = 11560
	ER_GRP_RPL_SRV_WAIT_TIME_OUT uint16 = 11561
	ER_GRP_RPL_FETCH_LOG_EVENT_FAILED uint16 = 11562
	ER_GRP_RPL_START_GRP_RPL_FAILED uint16 = 11563
	ER_GRP_RPL_CONN_INTERNAL_PLUGIN_FAIL uint16 = 11564
	ER_GRP_RPL_SUPER_READ_ON uint16 = 11565
	ER_GRP_RPL_SUPER_READ_OFF uint16 = 11566
	ER_GRP_RPL_KILLED_SESSION_ID uint16 = 11567
	ER_GRP_RPL_KILLED_FAILED_ID uint16 = 11568
	ER_GRP_RPL_INTERNAL_QUERY uint16 = 11569
	ER_GRP_RPL_COPY_FROM_EMPTY_STRING uint16 = 11570
	ER_GRP_RPL_QUERY_FAIL uint16 = 11571
	ER_GRP_RPL_CREATE_SESSION_UNABLE uint16 = 11572
	ER_GRP_RPL_MEMBER_NOT_FOUND uint16 = 11573
	ER_GRP_RPL_MAXIMUM_CONNECTION_RETRIES_REACHED uint16 = 11574
	ER_GRP_RPL_ALL_DONORS_LEFT_ABORT_RECOVERY uint16 = 11575
	ER_GRP_RPL_ESTABLISH_RECOVERY_WITH_DONOR uint16 = 11576
	ER_GRP_RPL_ESTABLISH_RECOVERY_WITH_ANOTHER_DONOR uint16 = 11577
	ER_GRP_RPL_NO_VALID_DONOR uint16 = 11578
	ER_GRP_RPL_CONFIG_RECOVERY uint16 = 11579
	ER_GRP_RPL_ESTABLISHING_CONN_GRP_REC_DONOR uint16 = 11580
	ER_GRP_RPL_CREATE_GRP_RPL_REC_CHANNEL uint16 = 11581
	ER_GRP_RPL_DONOR_SERVER_CONN uint16 = 11582
	ER_GRP_RPL_CHECK_STATUS_TABLE uint16 = 11583
	ER_GRP_RPL_STARTING_GRP_REC uint16 = 11584
	ER_GRP_RPL_DONOR_CONN_TERMINATION uint16 = 11585
	ER_GRP_RPL_STOPPING_GRP_REC uint16 = 11586
	ER_GRP_RPL_PURGE_REC uint16 = 11587
	ER_GRP_RPL_UNABLE_TO_KILL_CONN_REC_DONOR_APPLIER uint16 = 11588
	ER_GRP_RPL_UNABLE_TO_KILL_CONN_REC_DONOR_FAILOVER uint16 = 11589
	ER_GRP_RPL_FAILED_TO_NOTIFY_GRP_MEMBERSHIP_EVENT uint16 = 11590
	ER_GRP_RPL_FAILED_TO_BROADCAST_GRP_MEMBERSHIP_NOTIFICATION uint16 = 11591
	ER_GRP_RPL_FAILED_TO_BROADCAST_MEMBER_STATUS_NOTIFICATION uint16 = 11592
	ER_GRP_RPL_OOM_FAILED_TO_GENERATE_IDENTIFICATION_HASH uint16 = 11593
	ER_GRP_RPL_WRITE_IDENT_HASH_BASE64_ENCODING_FAILED uint16 = 11594
	ER_GRP_RPL_INVALID_BINLOG_FORMAT uint16 = 11595
	//OBSOLETE_ER_GRP_RPL_BINLOG_CHECKSUM_SET uint16 = 11596
	ER_GRP_RPL_TRANS_WRITE_SET_EXTRACTION_NOT_SET uint16 = 11597
	ER_GRP_RPL_UNSUPPORTED_TRANS_ISOLATION uint16 = 11598
	ER_GRP_RPL_CANNOT_EXECUTE_TRANS_WHILE_STOPPING uint16 = 11599
	ER_GRP_RPL_CANNOT_EXECUTE_TRANS_WHILE_RECOVERING uint16 = 11600
	ER_GRP_RPL_CANNOT_EXECUTE_TRANS_IN_ERROR_STATE uint16 = 11601
	ER_GRP_RPL_CANNOT_EXECUTE_TRANS_IN_OFFLINE_MODE uint16 = 11602
	ER_GRP_RPL_MULTIPLE_CACHE_TYPE_NOT_SUPPORTED_FOR_SESSION uint16 = 11603
	ER_GRP_RPL_FAILED_TO_REINIT_BINLOG_CACHE_FOR_READ uint16 = 11604
	ER_GRP_RPL_FAILED_TO_CREATE_TRANS_CONTEXT uint16 = 11605
	ER_GRP_RPL_FAILED_TO_EXTRACT_TRANS_WRITE_SET uint16 = 11606
	ER_GRP_RPL_FAILED_TO_GATHER_TRANS_WRITE_SET uint16 = 11607
	ER_GRP_RPL_TRANS_SIZE_EXCEEDS_LIMIT uint16 = 11608
	//OBSOLETE_ER_GRP_RPL_REINIT_OF_INTERNAL_CACHE_FOR_READ_FAILED uint16 = 11609
	//OBSOLETE_ER_GRP_RPL_APPENDING_DATA_TO_INTERNAL_CACHE_FAILED uint16 = 11610
	ER_GRP_RPL_WRITE_TO_TRANSACTION_MESSAGE_FAILED uint16 = 11611
	ER_GRP_RPL_FAILED_TO_REGISTER_TRANS_OUTCOME_NOTIFICTION uint16 = 11612
	ER_GRP_RPL_MSG_TOO_LONG_BROADCASTING_TRANS_FAILED uint16 = 11613
	ER_GRP_RPL_BROADCASTING_TRANS_TO_GRP_FAILED uint16 = 11614
	ER_GRP_RPL_ERROR_WHILE_WAITING_FOR_CONFLICT_DETECTION uint16 = 11615
	//OBSOLETE_ER_GRP_RPL_REINIT_OF_INTERNAL_CACHE_FOR_WRITE_FAILED uint16 = 11616
	//OBSOLETE_ER_GRP_RPL_FAILED_TO_CREATE_COMMIT_CACHE uint16 = 11617
	//OBSOLETE_ER_GRP_RPL_REINIT_OF_COMMIT_CACHE_FOR_WRITE_FAILED uint16 = 11618
	//OBSOLETE_ER_GRP_RPL_PREV_REC_SESSION_RUNNING uint16 = 11619
	ER_GRP_RPL_FATAL_REC_PROCESS uint16 = 11620
	//OBSOLETE_ER_GRP_RPL_WHILE_STOPPING_REP_CHANNEL uint16 = 11621
	ER_GRP_RPL_UNABLE_TO_EVALUATE_APPLIER_STATUS uint16 = 11622
	ER_GRP_RPL_ONLY_ONE_SERVER_ALIVE uint16 = 11623
	ER_GRP_RPL_CERTIFICATION_REC_PROCESS uint16 = 11624
	ER_GRP_RPL_UNABLE_TO_ENSURE_EXECUTION_REC uint16 = 11625
	ER_GRP_RPL_WHILE_SENDING_MSG_REC uint16 = 11626
	ER_GRP_RPL_READ_UNABLE_FOR_SUPER_READ_ONLY uint16 = 11627
	ER_GRP_RPL_READ_UNABLE_FOR_READ_ONLY_SUPER_READ_ONLY uint16 = 11628
	ER_GRP_RPL_UNABLE_TO_RESET_SERVER_READ_MODE uint16 = 11629
	ER_GRP_RPL_UNABLE_TO_CERTIFY_PLUGIN_TRANS uint16 = 11630
	ER_GRP_RPL_UNBLOCK_CERTIFIED_TRANS uint16 = 11631
	//OBSOLETE_ER_GRP_RPL_SERVER_WORKING_AS_SECONDARY uint16 = 11632
	ER_GRP_RPL_FAILED_TO_START_WITH_INVALID_SERVER_ID uint16 = 11633
	ER_GRP_RPL_FORCE_MEMBERS_MUST_BE_EMPTY uint16 = 11634
	ER_GRP_RPL_PLUGIN_STRUCT_INIT_NOT_POSSIBLE_ON_SERVER_START uint16 = 11635
	ER_GRP_RPL_FAILED_TO_ENABLE_SUPER_READ_ONLY_MODE uint16 = 11636
	ER_GRP_RPL_FAILED_TO_INIT_COMMUNICATION_ENGINE uint16 = 11637
	ER_GRP_RPL_FAILED_TO_START_ON_SECONDARY_WITH_ASYNC_CHANNELS uint16 = 11638
	ER_GRP_RPL_FAILED_TO_START_COMMUNICATION_ENGINE uint16 = 11639
	ER_GRP_RPL_TIMEOUT_ON_VIEW_AFTER_JOINING_GRP uint16 = 11640
	ER_GRP_RPL_FAILED_TO_CALL_GRP_COMMUNICATION_INTERFACE uint16 = 11641
	ER_GRP_RPL_MEMBER_SERVER_UUID_IS_INCOMPATIBLE_WITH_GRP uint16 = 11642
	ER_GRP_RPL_MEMBER_CONF_INFO uint16 = 11643
	ER_GRP_RPL_FAILED_TO_CONFIRM_IF_SERVER_LEFT_GRP uint16 = 11644
	ER_GRP_RPL_SERVER_IS_ALREADY_LEAVING uint16 = 11645
	ER_GRP_RPL_SERVER_ALREADY_LEFT uint16 = 11646
	ER_GRP_RPL_WAITING_FOR_VIEW_UPDATE uint16 = 11647
	ER_GRP_RPL_TIMEOUT_RECEIVING_VIEW_CHANGE_ON_SHUTDOWN uint16 = 11648
	ER_GRP_RPL_REQUESTING_NON_MEMBER_SERVER_TO_LEAVE uint16 = 11649
	ER_GRP_RPL_IS_STOPPING uint16 = 11650
	ER_GRP_RPL_IS_STOPPED uint16 = 11651
	ER_GRP_RPL_FAILED_TO_ENABLE_READ_ONLY_MODE_ON_SHUTDOWN uint16 = 11652
	ER_GRP_RPL_RECOVERY_MODULE_TERMINATION_TIMED_OUT_ON_SHUTDOWN uint16 = 11653
	ER_GRP_RPL_APPLIER_TERMINATION_TIMED_OUT_ON_SHUTDOWN uint16 = 11654
	ER_GRP_RPL_FAILED_TO_SHUTDOWN_REGISTRY_MODULE uint16 = 11655
	ER_GRP_RPL_FAILED_TO_INIT_HANDLER uint16 = 11656
	ER_GRP_RPL_FAILED_TO_REGISTER_SERVER_STATE_OBSERVER uint16 = 11657
	ER_GRP_RPL_FAILED_TO_REGISTER_TRANS_STATE_OBSERVER uint16 = 11658
	ER_GRP_RPL_FAILED_TO_REGISTER_BINLOG_STATE_OBSERVER uint16 = 11659
	ER_GRP_RPL_FAILED_TO_START_ON_BOOT uint16 = 11660
	ER_GRP_RPL_FAILED_TO_STOP_ON_PLUGIN_UNINSTALL uint16 = 11661
	ER_GRP_RPL_FAILED_TO_UNREGISTER_SERVER_STATE_OBSERVER uint16 = 11662
	ER_GRP_RPL_FAILED_TO_UNREGISTER_TRANS_STATE_OBSERVER uint16 = 11663
	ER_GRP_RPL_FAILED_TO_UNREGISTER_BINLOG_STATE_OBSERVER uint16 = 11664
	ER_GRP_RPL_ALL_OBSERVERS_UNREGISTERED uint16 = 11665
	ER_GRP_RPL_FAILED_TO_PARSE_THE_GRP_NAME uint16 = 11666
	ER_GRP_RPL_FAILED_TO_GENERATE_SIDNO_FOR_GRP uint16 = 11667
	ER_GRP_RPL_APPLIER_NOT_STARTED_DUE_TO_RUNNING_PREV_SHUTDOWN uint16 = 11668
	ER_GRP_RPL_FAILED_TO_INIT_APPLIER_MODULE uint16 = 11669
	ER_GRP_RPL_APPLIER_INITIALIZED uint16 = 11670
	ER_GRP_RPL_COMMUNICATION_SSL_CONF_INFO uint16 = 11671
	ER_GRP_RPL_ABORTS_AS_SSL_NOT_SUPPORTED_BY_MYSQLD uint16 = 11672
	ER_GRP_RPL_SSL_DISABLED uint16 = 11673
	ER_GRP_RPL_UNABLE_TO_INIT_COMMUNICATION_ENGINE uint16 = 11674
	ER_GRP_RPL_BINLOG_DISABLED uint16 = 11675
	ER_GRP_RPL_GTID_MODE_OFF uint16 = 11676
	ER_GRP_RPL_LOG_SLAVE_UPDATES_NOT_SET uint16 = 11677
	ER_GRP_RPL_INVALID_TRANS_WRITE_SET_EXTRACTION_VALUE uint16 = 11678
	ER_GRP_RPL_RELAY_LOG_INFO_REPO_MUST_BE_TABLE uint16 = 11679
	ER_GRP_RPL_MASTER_INFO_REPO_MUST_BE_TABLE uint16 = 11680
	ER_GRP_RPL_INCORRECT_TYPE_SET_FOR_PARALLEL_APPLIER uint16 = 11681
	ER_GRP_RPL_SLAVE_PRESERVE_COMMIT_ORDER_NOT_SET uint16 = 11682
	ER_GRP_RPL_SINGLE_PRIM_MODE_NOT_ALLOWED_WITH_UPDATE_EVERYWHERE uint16 = 11683
	ER_GRP_RPL_MODULE_TERMINATE_ERROR uint16 = 11684
	ER_GRP_RPL_GRP_NAME_OPTION_MANDATORY uint16 = 11685
	ER_GRP_RPL_GRP_NAME_IS_TOO_LONG uint16 = 11686
	ER_GRP_RPL_GRP_NAME_IS_NOT_VALID_UUID uint16 = 11687
	ER_GRP_RPL_FLOW_CTRL_MIN_QUOTA_GREATER_THAN_MAX_QUOTA uint16 = 11688
	ER_GRP_RPL_FLOW_CTRL_MIN_RECOVERY_QUOTA_GREATER_THAN_MAX_QUOTA uint16 = 11689
	ER_GRP_RPL_FLOW_CTRL_MAX_QUOTA_SMALLER_THAN_MIN_QUOTAS uint16 = 11690
	ER_GRP_RPL_INVALID_SSL_RECOVERY_STRING uint16 = 11691
	ER_GRP_RPL_SUPPORTS_ONLY_ONE_FORCE_MEMBERS_SET uint16 = 11692
	ER_GRP_RPL_FORCE_MEMBERS_SET_UPDATE_NOT_ALLOWED uint16 = 11693
	ER_GRP_RPL_GRP_COMMUNICATION_INIT_WITH_CONF uint16 = 11694
	ER_GRP_RPL_UNKNOWN_GRP_RPL_APPLIER_PIPELINE_REQUESTED uint16 = 11695
	ER_GRP_RPL_FAILED_TO_BOOTSTRAP_EVENT_HANDLING_INFRASTRUCTURE uint16 = 11696
	ER_GRP_RPL_APPLIER_HANDLER_NOT_INITIALIZED uint16 = 11697
	ER_GRP_RPL_APPLIER_HANDLER_IS_IN_USE uint16 = 11698
	ER_GRP_RPL_APPLIER_HANDLER_ROLE_IS_IN_USE uint16 = 11699
	ER_GRP_RPL_FAILED_TO_INIT_APPLIER_HANDLER uint16 = 11700
	ER_GRP_RPL_SQL_SERVICE_FAILED_TO_INIT_SESSION_THREAD uint16 = 11701
	ER_GRP_RPL_SQL_SERVICE_COMM_SESSION_NOT_INITIALIZED uint16 = 11702
	ER_GRP_RPL_SQL_SERVICE_SERVER_SESSION_KILLED uint16 = 11703
	ER_GRP_RPL_SQL_SERVICE_FAILED_TO_RUN_SQL_QUERY uint16 = 11704
	ER_GRP_RPL_SQL_SERVICE_SERVER_INTERNAL_FAILURE uint16 = 11705
	ER_GRP_RPL_SQL_SERVICE_RETRIES_EXCEEDED_ON_SESSION_STATE uint16 = 11706
	ER_GRP_RPL_SQL_SERVICE_FAILED_TO_FETCH_SECURITY_CTX uint16 = 11707
	ER_GRP_RPL_SQL_SERVICE_SERVER_ACCESS_DENIED_FOR_USER uint16 = 11708
	ER_GRP_RPL_SQL_SERVICE_MAX_CONN_ERROR_FROM_SERVER uint16 = 11709
	ER_GRP_RPL_SQL_SERVICE_SERVER_ERROR_ON_CONN uint16 = 11710
	ER_GRP_RPL_UNREACHABLE_MAJORITY_TIMEOUT_FOR_MEMBER uint16 = 11711
	ER_GRP_RPL_SERVER_SET_TO_READ_ONLY_DUE_TO_ERRORS uint16 = 11712
	ER_GRP_RPL_GMS_LISTENER_FAILED_TO_LOG_NOTIFICATION uint16 = 11713
	ER_GRP_RPL_GRP_COMMUNICATION_ENG_INIT_FAILED uint16 = 11714
	ER_GRP_RPL_SET_GRP_COMMUNICATION_ENG_LOGGER_FAILED uint16 = 11715
	ER_GRP_RPL_DEBUG_OPTIONS uint16 = 11716
	ER_GRP_RPL_INVALID_DEBUG_OPTIONS uint16 = 11717
	ER_GRP_RPL_EXIT_GRP_GCS_ERROR uint16 = 11718
	ER_GRP_RPL_GRP_MEMBER_OFFLINE uint16 = 11719
	ER_GRP_RPL_GCS_INTERFACE_ERROR uint16 = 11720
	ER_GRP_RPL_FORCE_MEMBER_VALUE_SET_ERROR uint16 = 11721
	ER_GRP_RPL_FORCE_MEMBER_VALUE_SET uint16 = 11722
	ER_GRP_RPL_FORCE_MEMBER_VALUE_TIME_OUT uint16 = 11723
	ER_GRP_RPL_BROADCAST_COMMIT_MSSG_TOO_BIG uint16 = 11724
	ER_GRP_RPL_SEND_STATS_ERROR uint16 = 11725
	ER_GRP_RPL_MEMBER_STATS_INFO uint16 = 11726
	ER_GRP_RPL_FLOW_CONTROL_STATS uint16 = 11727
	ER_GRP_RPL_UNABLE_TO_CONVERT_PACKET_TO_EVENT uint16 = 11728
	ER_GRP_RPL_PIPELINE_CREATE_FAILED uint16 = 11729
	ER_GRP_RPL_PIPELINE_REINIT_FAILED_WRITE uint16 = 11730
	ER_GRP_RPL_UNABLE_TO_CONVERT_EVENT_TO_PACKET uint16 = 11731
	ER_GRP_RPL_PIPELINE_FLUSH_FAIL uint16 = 11732
	ER_GRP_RPL_PIPELINE_REINIT_FAILED_READ uint16 = 11733
	//OBSOLETE_ER_GRP_RPL_STOP_REP_CHANNEL uint16 = 11734
	ER_GRP_RPL_GCS_GR_ERROR_MSG uint16 = 11735
	ER_GRP_RPL_SLAVE_IO_THREAD_UNBLOCKED uint16 = 11736
	ER_GRP_RPL_SLAVE_IO_THREAD_ERROR_OUT uint16 = 11737
	ER_GRP_RPL_SLAVE_APPLIER_THREAD_UNBLOCKED uint16 = 11738
	ER_GRP_RPL_SLAVE_APPLIER_THREAD_ERROR_OUT uint16 = 11739
	ER_LDAP_AUTH_FAILED_TO_CREATE_OR_GET_CONNECTION uint16 = 11740
	ER_LDAP_AUTH_DEINIT_FAILED uint16 = 11741
	ER_LDAP_AUTH_SKIPPING_USER_GROUP_SEARCH uint16 = 11742
	ER_LDAP_AUTH_POOL_DISABLE_MAX_SIZE_ZERO uint16 = 11743
	ER_LDAP_AUTH_FAILED_TO_CREATE_LDAP_OBJECT_CREATOR uint16 = 11744
	ER_LDAP_AUTH_FAILED_TO_CREATE_LDAP_OBJECT uint16 = 11745
	ER_LDAP_AUTH_TLS_CONF uint16 = 11746
	ER_LDAP_AUTH_TLS_CONNECTION uint16 = 11747
	ER_LDAP_AUTH_CONN_POOL_NOT_CREATED uint16 = 11748
	ER_LDAP_AUTH_CONN_POOL_INITIALIZING uint16 = 11749
	ER_LDAP_AUTH_CONN_POOL_DEINITIALIZING uint16 = 11750
	ER_LDAP_AUTH_ZERO_MAX_POOL_SIZE_UNCHANGED uint16 = 11751
	ER_LDAP_AUTH_POOL_REINITIALIZING uint16 = 11752
	ER_LDAP_AUTH_FAILED_TO_WRITE_PACKET uint16 = 11753
	ER_LDAP_AUTH_SETTING_USERNAME uint16 = 11754
	ER_LDAP_AUTH_USER_AUTH_DATA uint16 = 11755
	ER_LDAP_AUTH_INFO_FOR_USER uint16 = 11756
	ER_LDAP_AUTH_USER_GROUP_SEARCH_INFO uint16 = 11757
	ER_LDAP_AUTH_GRP_SEARCH_SPECIAL_HDL uint16 = 11758
	ER_LDAP_AUTH_GRP_IS_FULL_DN uint16 = 11759
	ER_LDAP_AUTH_USER_NOT_FOUND_IN_ANY_GRP uint16 = 11760
	ER_LDAP_AUTH_USER_FOUND_IN_MANY_GRPS uint16 = 11761
	ER_LDAP_AUTH_USER_HAS_MULTIPLE_GRP_NAMES uint16 = 11762
	ER_LDAP_AUTH_SEARCHED_USER_GRP_NAME uint16 = 11763
	ER_LDAP_AUTH_OBJECT_CREATE_TIMESTAMP uint16 = 11764
	ER_LDAP_AUTH_CERTIFICATE_NAME uint16 = 11765
	ER_LDAP_AUTH_FAILED_TO_POOL_DEINIT uint16 = 11766
	ER_LDAP_AUTH_FAILED_TO_INITIALIZE_POOL_IN_RECONSTRUCTING uint16 = 11767
	ER_LDAP_AUTH_FAILED_TO_INITIALIZE_POOL_IN_INIT_STATE uint16 = 11768
	ER_LDAP_AUTH_FAILED_TO_INITIALIZE_POOL_IN_DEINIT_STATE uint16 = 11769
	ER_LDAP_AUTH_FAILED_TO_DEINITIALIZE_POOL_IN_RECONSTRUCT_STATE uint16 = 11770
	ER_LDAP_AUTH_FAILED_TO_DEINITIALIZE_NOT_READY_POOL uint16 = 11771
	ER_LDAP_AUTH_FAILED_TO_GET_CONNECTION_AS_PLUGIN_NOT_READY uint16 = 11772
	ER_LDAP_AUTH_CONNECTION_POOL_INIT_FAILED uint16 = 11773
	ER_LDAP_AUTH_MAX_ALLOWED_CONNECTION_LIMIT_HIT uint16 = 11774
	ER_LDAP_AUTH_MAX_POOL_SIZE_SET_FAILED uint16 = 11775
	ER_LDAP_AUTH_PLUGIN_FAILED_TO_READ_PACKET uint16 = 11776
	ER_LDAP_AUTH_CREATING_LDAP_CONNECTION uint16 = 11777
	ER_LDAP_AUTH_GETTING_CONNECTION_FROM_POOL uint16 = 11778
	ER_LDAP_AUTH_RETURNING_CONNECTION_TO_POOL uint16 = 11779
	ER_LDAP_AUTH_SEARCH_USER_GROUP_ATTR_NOT_FOUND uint16 = 11780
	ER_LDAP_AUTH_LDAP_INFO_NULL uint16 = 11781
	ER_LDAP_AUTH_FREEING_CONNECTION uint16 = 11782
	ER_LDAP_AUTH_CONNECTION_PUSHED_TO_POOL uint16 = 11783
	ER_LDAP_AUTH_CONNECTION_CREATOR_ENTER uint16 = 11784
	ER_LDAP_AUTH_STARTING_TLS uint16 = 11785
	ER_LDAP_AUTH_CONNECTION_GET_LDAP_INFO_NULL uint16 = 11786
	ER_LDAP_AUTH_DELETING_CONNECTION_KEY uint16 = 11787
	ER_LDAP_AUTH_POOLED_CONNECTION_KEY uint16 = 11788
	ER_LDAP_AUTH_CREATE_CONNECTION_KEY uint16 = 11789
	ER_LDAP_AUTH_COMMUNICATION_HOST_INFO uint16 = 11790
	ER_LDAP_AUTH_METHOD_TO_CLIENT uint16 = 11791
	ER_LDAP_AUTH_SASL_REQUEST_FROM_CLIENT uint16 = 11792
	ER_LDAP_AUTH_SASL_PROCESS_SASL uint16 = 11793
	ER_LDAP_AUTH_SASL_BIND_SUCCESS_INFO uint16 = 11794
	ER_LDAP_AUTH_STARTED_FOR_USER uint16 = 11795
	ER_LDAP_AUTH_DISTINGUISHED_NAME uint16 = 11796
	ER_LDAP_AUTH_INIT_FAILED uint16 = 11797
	ER_LDAP_AUTH_OR_GROUP_RETRIEVAL_FAILED uint16 = 11798
	ER_LDAP_AUTH_USER_GROUP_SEARCH_FAILED uint16 = 11799
	ER_LDAP_AUTH_USER_BIND_FAILED uint16 = 11800
	ER_LDAP_AUTH_POOL_GET_FAILED_TO_CREATE_CONNECTION uint16 = 11801
	ER_LDAP_AUTH_FAILED_TO_CREATE_LDAP_CONNECTION uint16 = 11802
	ER_LDAP_AUTH_FAILED_TO_ESTABLISH_TLS_CONNECTION uint16 = 11803
	ER_LDAP_AUTH_FAILED_TO_SEARCH_DN uint16 = 11804
	ER_LDAP_AUTH_CONNECTION_POOL_REINIT_ENTER uint16 = 11805
	ER_SYSTEMD_NOTIFY_PATH_TOO_LONG uint16 = 11806
	ER_SYSTEMD_NOTIFY_CONNECT_FAILED uint16 = 11807
	ER_SYSTEMD_NOTIFY_WRITE_FAILED uint16 = 11808
	ER_FOUND_MISSING_GTIDS uint16 = 11809
	ER_PID_FILE_PRIV_DIRECTORY_INSECURE uint16 = 11810
	ER_CANT_CHECK_PID_PATH uint16 = 11811
	ER_VALIDATE_PWD_STATUS_VAR_REGISTRATION_FAILED uint16 = 11812
	ER_VALIDATE_PWD_STATUS_VAR_UNREGISTRATION_FAILED uint16 = 11813
	ER_VALIDATE_PWD_DICT_FILE_OPEN_FAILED uint16 = 11814
	ER_VALIDATE_PWD_COULD_BE_NULL uint16 = 11815
	ER_VALIDATE_PWD_STRING_CONV_TO_LOWERCASE_FAILED uint16 = 11816
	ER_VALIDATE_PWD_STRING_CONV_TO_BUFFER_FAILED uint16 = 11817
	ER_VALIDATE_PWD_STRING_HANDLER_MEM_ALLOCATION_FAILED uint16 = 11818
	ER_VALIDATE_PWD_STRONG_POLICY_DICT_FILE_UNSPECIFIED uint16 = 11819
	ER_VALIDATE_PWD_CONVERT_TO_BUFFER_FAILED uint16 = 11820
	ER_VALIDATE_PWD_VARIABLE_REGISTRATION_FAILED uint16 = 11821
	ER_VALIDATE_PWD_VARIABLE_UNREGISTRATION_FAILED uint16 = 11822
	ER_KEYRING_MIGRATION_EXTRA_OPTIONS uint16 = 11823
	//OBSOLETE_ER_INVALID_DEFAULT_UTF8MB4_COLLATION uint16 = 11824
	ER_IB_MSG_0 uint16 = 11825
	ER_IB_MSG_1 uint16 = 11826
	ER_IB_MSG_2 uint16 = 11827
	ER_IB_MSG_3 uint16 = 11828
	ER_IB_MSG_4 uint16 = 11829
	ER_IB_MSG_5 uint16 = 11830
	ER_IB_MSG_6 uint16 = 11831
	ER_IB_MSG_7 uint16 = 11832
	ER_IB_MSG_8 uint16 = 11833
	ER_IB_MSG_9 uint16 = 11834
	ER_IB_MSG_10 uint16 = 11835
	ER_IB_MSG_11 uint16 = 11836
	ER_IB_MSG_12 uint16 = 11837
	ER_IB_MSG_13 uint16 = 11838
	ER_IB_MSG_14 uint16 = 11839
	ER_IB_MSG_15 uint16 = 11840
	ER_IB_MSG_16 uint16 = 11841
	ER_IB_MSG_17 uint16 = 11842
	ER_IB_MSG_18 uint16 = 11843
	ER_IB_MSG_19 uint16 = 11844
	ER_IB_MSG_20 uint16 = 11845
	ER_IB_MSG_21 uint16 = 11846
	ER_IB_MSG_22 uint16 = 11847
	ER_IB_MSG_23 uint16 = 11848
	ER_IB_MSG_24 uint16 = 11849
	ER_IB_MSG_25 uint16 = 11850
	ER_IB_MSG_26 uint16 = 11851
	ER_IB_MSG_27 uint16 = 11852
	ER_IB_MSG_28 uint16 = 11853
	ER_IB_MSG_29 uint16 = 11854
	ER_IB_MSG_30 uint16 = 11855
	ER_IB_MSG_31 uint16 = 11856
	ER_IB_MSG_32 uint16 = 11857
	ER_IB_MSG_33 uint16 = 11858
	ER_IB_MSG_34 uint16 = 11859
	ER_IB_MSG_35 uint16 = 11860
	ER_IB_MSG_36 uint16 = 11861
	ER_IB_MSG_37 uint16 = 11862
	ER_IB_MSG_38 uint16 = 11863
	ER_IB_MSG_39 uint16 = 11864
	ER_IB_MSG_40 uint16 = 11865
	ER_IB_MSG_41 uint16 = 11866
	ER_IB_MSG_42 uint16 = 11867
	ER_IB_MSG_43 uint16 = 11868
	ER_IB_MSG_44 uint16 = 11869
	ER_IB_MSG_45 uint16 = 11870
	ER_IB_MSG_46 uint16 = 11871
	ER_IB_MSG_47 uint16 = 11872
	ER_IB_MSG_48 uint16 = 11873
	ER_IB_MSG_49 uint16 = 11874
	ER_IB_MSG_50 uint16 = 11875
	ER_IB_MSG_51 uint16 = 11876
	ER_IB_MSG_52 uint16 = 11877
	ER_IB_MSG_53 uint16 = 11878
	ER_IB_MSG_54 uint16 = 11879
	ER_IB_MSG_55 uint16 = 11880
	ER_IB_MSG_56 uint16 = 11881
	ER_IB_MSG_57 uint16 = 11882
	ER_IB_MSG_58 uint16 = 11883
	ER_IB_MSG_59 uint16 = 11884
	ER_IB_MSG_60 uint16 = 11885
	ER_IB_MSG_61 uint16 = 11886
	ER_IB_MSG_62 uint16 = 11887
	ER_IB_MSG_63 uint16 = 11888
	ER_IB_MSG_64 uint16 = 11889
	ER_IB_MSG_65 uint16 = 11890
	ER_IB_MSG_66 uint16 = 11891
	ER_IB_MSG_67 uint16 = 11892
	ER_IB_MSG_68 uint16 = 11893
	ER_IB_MSG_69 uint16 = 11894
	ER_IB_MSG_70 uint16 = 11895
	ER_IB_MSG_71 uint16 = 11896
	ER_IB_MSG_72 uint16 = 11897
	ER_IB_MSG_73 uint16 = 11898
	ER_IB_MSG_74 uint16 = 11899
	ER_IB_MSG_75 uint16 = 11900
	ER_IB_MSG_76 uint16 = 11901
	ER_IB_MSG_77 uint16 = 11902
	ER_IB_MSG_78 uint16 = 11903
	ER_IB_MSG_79 uint16 = 11904
	ER_IB_MSG_80 uint16 = 11905
	ER_IB_MSG_81 uint16 = 11906
	ER_IB_MSG_82 uint16 = 11907
	ER_IB_MSG_83 uint16 = 11908
	ER_IB_MSG_84 uint16 = 11909
	ER_IB_MSG_85 uint16 = 11910
	ER_IB_MSG_86 uint16 = 11911
	//OBSOLETE_ER_IB_MSG_87 uint16 = 11912
	//OBSOLETE_ER_IB_MSG_88 uint16 = 11913
	//OBSOLETE_ER_IB_MSG_89 uint16 = 11914
	//OBSOLETE_ER_IB_MSG_90 uint16 = 11915
	//OBSOLETE_ER_IB_MSG_91 uint16 = 11916
	//OBSOLETE_ER_IB_MSG_92 uint16 = 11917
	//OBSOLETE_ER_IB_MSG_93 uint16 = 11918
	//OBSOLETE_ER_IB_MSG_94 uint16 = 11919
	ER_IB_MSG_95 uint16 = 11920
	ER_IB_MSG_96 uint16 = 11921
	ER_IB_MSG_97 uint16 = 11922
	ER_IB_MSG_98 uint16 = 11923
	ER_IB_MSG_99 uint16 = 11924
	ER_IB_MSG_100 uint16 = 11925
	ER_IB_MSG_101 uint16 = 11926
	ER_IB_MSG_102 uint16 = 11927
	ER_IB_MSG_103 uint16 = 11928
	ER_IB_MSG_104 uint16 = 11929
	ER_IB_MSG_105 uint16 = 11930
	ER_IB_MSG_106 uint16 = 11931
	ER_IB_MSG_107 uint16 = 11932
	ER_IB_MSG_108 uint16 = 11933
	ER_IB_MSG_109 uint16 = 11934
	ER_IB_MSG_110 uint16 = 11935
	ER_IB_MSG_111 uint16 = 11936
	ER_IB_MSG_112 uint16 = 11937
	//OBSOLETE_ER_IB_MSG_113 uint16 = 11938
	//OBSOLETE_ER_IB_MSG_114 uint16 = 11939
	//OBSOLETE_ER_IB_MSG_115 uint16 = 11940
	//OBSOLETE_ER_IB_MSG_116 uint16 = 11941
	//OBSOLETE_ER_IB_MSG_117 uint16 = 11942
	//OBSOLETE_ER_IB_MSG_118 uint16 = 11943
	ER_IB_MSG_119 uint16 = 11944
	ER_IB_MSG_120 uint16 = 11945
	ER_IB_MSG_121 uint16 = 11946
	ER_IB_MSG_122 uint16 = 11947
	ER_IB_MSG_123 uint16 = 11948
	ER_IB_MSG_124 uint16 = 11949
	ER_IB_MSG_125 uint16 = 11950
	ER_IB_MSG_126 uint16 = 11951
	ER_IB_MSG_127 uint16 = 11952
	ER_IB_MSG_128 uint16 = 11953
	ER_IB_MSG_129 uint16 = 11954
	ER_IB_MSG_130 uint16 = 11955
	ER_IB_MSG_131 uint16 = 11956
	ER_IB_MSG_132 uint16 = 11957
	ER_IB_MSG_133 uint16 = 11958
	ER_IB_MSG_134 uint16 = 11959
	ER_IB_MSG_135 uint16 = 11960
	ER_IB_MSG_136 uint16 = 11961
	ER_IB_MSG_137 uint16 = 11962
	ER_IB_MSG_138 uint16 = 11963
	ER_IB_MSG_139 uint16 = 11964
	ER_IB_MSG_140 uint16 = 11965
	ER_IB_MSG_141 uint16 = 11966
	ER_IB_MSG_142 uint16 = 11967
	ER_IB_MSG_143 uint16 = 11968
	ER_IB_MSG_144 uint16 = 11969
	ER_IB_MSG_145 uint16 = 11970
	ER_IB_MSG_146 uint16 = 11971
	ER_IB_MSG_147 uint16 = 11972
	ER_IB_MSG_148 uint16 = 11973
	ER_IB_CLONE_INTERNAL uint16 = 11974
	ER_IB_CLONE_TIMEOUT uint16 = 11975
	ER_IB_CLONE_STATUS_FILE uint16 = 11976
	ER_IB_CLONE_SQL uint16 = 11977
	ER_IB_CLONE_VALIDATE uint16 = 11978
	ER_IB_CLONE_PUNCH_HOLE uint16 = 11979
	ER_IB_CLONE_GTID_PERSIST uint16 = 11980
	ER_IB_MSG_156 uint16 = 11981
	ER_IB_MSG_157 uint16 = 11982
	ER_IB_MSG_158 uint16 = 11983
	ER_IB_MSG_159 uint16 = 11984
	ER_IB_MSG_160 uint16 = 11985
	ER_IB_MSG_161 uint16 = 11986
	ER_IB_MSG_162 uint16 = 11987
	ER_IB_MSG_163 uint16 = 11988
	ER_IB_MSG_164 uint16 = 11989
	ER_IB_MSG_165 uint16 = 11990
	ER_IB_MSG_166 uint16 = 11991
	ER_IB_MSG_167 uint16 = 11992
	ER_IB_MSG_168 uint16 = 11993
	ER_IB_MSG_169 uint16 = 11994
	ER_IB_MSG_170 uint16 = 11995
	ER_IB_MSG_171 uint16 = 11996
	ER_IB_MSG_172 uint16 = 11997
	ER_IB_MSG_173 uint16 = 11998
	ER_IB_MSG_174 uint16 = 11999
	ER_IB_MSG_175 uint16 = 12000
	ER_IB_MSG_176 uint16 = 12001
	ER_IB_MSG_177 uint16 = 12002
	ER_IB_MSG_178 uint16 = 12003
	ER_IB_MSG_179 uint16 = 12004
	ER_IB_MSG_180 uint16 = 12005
	ER_IB_MSG_181 uint16 = 12006
	ER_IB_MSG_182 uint16 = 12007
	ER_IB_MSG_183 uint16 = 12008
	ER_IB_MSG_184 uint16 = 12009
	ER_IB_MSG_185 uint16 = 12010
	ER_IB_MSG_186 uint16 = 12011
	ER_IB_MSG_187 uint16 = 12012
	ER_IB_MSG_188 uint16 = 12013
	ER_IB_MSG_189 uint16 = 12014
	ER_IB_MSG_190 uint16 = 12015
	ER_IB_MSG_191 uint16 = 12016
	ER_IB_MSG_192 uint16 = 12017
	ER_IB_MSG_193 uint16 = 12018
	ER_IB_MSG_194 uint16 = 12019
	ER_IB_MSG_195 uint16 = 12020
	ER_IB_MSG_196 uint16 = 12021
	ER_IB_MSG_197 uint16 = 12022
	ER_IB_MSG_198 uint16 = 12023
	ER_IB_MSG_199 uint16 = 12024
	ER_IB_MSG_200 uint16 = 12025
	ER_IB_MSG_201 uint16 = 12026
	ER_IB_MSG_202 uint16 = 12027
	ER_IB_MSG_203 uint16 = 12028
	ER_IB_MSG_204 uint16 = 12029
	ER_IB_MSG_205 uint16 = 12030
	ER_IB_MSG_206 uint16 = 12031
	ER_IB_MSG_207 uint16 = 12032
	ER_IB_MSG_208 uint16 = 12033
	ER_IB_MSG_209 uint16 = 12034
	ER_IB_MSG_210 uint16 = 12035
	ER_IB_MSG_211 uint16 = 12036
	ER_IB_MSG_212 uint16 = 12037
	ER_IB_MSG_213 uint16 = 12038
	ER_IB_MSG_214 uint16 = 12039
	ER_IB_MSG_215 uint16 = 12040
	ER_IB_MSG_216 uint16 = 12041
	ER_IB_MSG_217 uint16 = 12042
	ER_IB_MSG_218 uint16 = 12043
	ER_IB_MSG_219 uint16 = 12044
	ER_IB_MSG_220 uint16 = 12045
	ER_IB_MSG_221 uint16 = 12046
	ER_IB_MSG_222 uint16 = 12047
	ER_IB_MSG_223 uint16 = 12048
	ER_IB_MSG_224 uint16 = 12049
	ER_IB_MSG_225 uint16 = 12050
	ER_IB_MSG_226 uint16 = 12051
	ER_IB_MSG_227 uint16 = 12052
	ER_IB_MSG_228 uint16 = 12053
	ER_IB_MSG_229 uint16 = 12054
	ER_IB_MSG_230 uint16 = 12055
	ER_IB_MSG_231 uint16 = 12056
	ER_IB_MSG_232 uint16 = 12057
	ER_IB_MSG_233 uint16 = 12058
	ER_IB_MSG_234 uint16 = 12059
	ER_IB_MSG_235 uint16 = 12060
	ER_IB_MSG_236 uint16 = 12061
	ER_IB_MSG_237 uint16 = 12062
	ER_IB_MSG_238 uint16 = 12063
	ER_IB_MSG_239 uint16 = 12064
	ER_IB_MSG_240 uint16 = 12065
	ER_IB_MSG_241 uint16 = 12066
	ER_IB_MSG_242 uint16 = 12067
	ER_IB_MSG_243 uint16 = 12068
	ER_IB_MSG_244 uint16 = 12069
	ER_IB_MSG_245 uint16 = 12070
	ER_IB_MSG_246 uint16 = 12071
	ER_IB_MSG_247 uint16 = 12072
	ER_IB_MSG_248 uint16 = 12073
	ER_IB_MSG_249 uint16 = 12074
	ER_IB_MSG_250 uint16 = 12075
	ER_IB_MSG_251 uint16 = 12076
	ER_IB_MSG_252 uint16 = 12077
	ER_IB_MSG_253 uint16 = 12078
	ER_IB_MSG_254 uint16 = 12079
	ER_IB_MSG_255 uint16 = 12080
	ER_IB_MSG_256 uint16 = 12081
	ER_IB_MSG_257 uint16 = 12082
	ER_IB_MSG_258 uint16 = 12083
	ER_IB_MSG_259 uint16 = 12084
	ER_IB_MSG_260 uint16 = 12085
	ER_IB_MSG_261 uint16 = 12086
	ER_IB_MSG_262 uint16 = 12087
	ER_IB_MSG_263 uint16 = 12088
	ER_IB_MSG_264 uint16 = 12089
	ER_IB_MSG_265 uint16 = 12090
	ER_IB_MSG_266 uint16 = 12091
	ER_IB_MSG_267 uint16 = 12092
	ER_IB_MSG_268 uint16 = 12093
	ER_IB_MSG_269 uint16 = 12094
	ER_IB_MSG_270 uint16 = 12095
	ER_IB_MSG_271 uint16 = 12096
	ER_IB_MSG_272 uint16 = 12097
	ER_IB_MSG_273 uint16 = 12098
	ER_IB_MSG_274 uint16 = 12099
	ER_IB_MSG_275 uint16 = 12100
	ER_IB_MSG_276 uint16 = 12101
	ER_IB_MSG_277 uint16 = 12102
	ER_IB_MSG_278 uint16 = 12103
	ER_IB_MSG_279 uint16 = 12104
	ER_IB_MSG_280 uint16 = 12105
	ER_IB_MSG_281 uint16 = 12106
	ER_IB_MSG_282 uint16 = 12107
	ER_IB_MSG_283 uint16 = 12108
	ER_IB_MSG_284 uint16 = 12109
	ER_IB_MSG_285 uint16 = 12110
	ER_IB_MSG_286 uint16 = 12111
	ER_IB_MSG_287 uint16 = 12112
	ER_IB_MSG_288 uint16 = 12113
	ER_IB_MSG_289 uint16 = 12114
	//OBSOLETE_ER_IB_MSG_290 uint16 = 12115
	ER_IB_MSG_291 uint16 = 12116
	ER_IB_MSG_292 uint16 = 12117
	ER_IB_MSG_293 uint16 = 12118
	ER_IB_MSG_294 uint16 = 12119
	ER_IB_MSG_295 uint16 = 12120
	ER_IB_MSG_296 uint16 = 12121
	ER_IB_MSG_297 uint16 = 12122
	ER_IB_MSG_298 uint16 = 12123
	ER_IB_MSG_299 uint16 = 12124
	ER_IB_MSG_300 uint16 = 12125
	ER_IB_MSG_301 uint16 = 12126
	ER_IB_MSG_UNEXPECTED_FILE_EXISTS uint16 = 12127
	ER_IB_MSG_303 uint16 = 12128
	ER_IB_MSG_304 uint16 = 12129
	ER_IB_MSG_305 uint16 = 12130
	ER_IB_MSG_306 uint16 = 12131
	ER_IB_MSG_307 uint16 = 12132
	ER_IB_MSG_308 uint16 = 12133
	ER_IB_MSG_309 uint16 = 12134
	ER_IB_MSG_310 uint16 = 12135
	ER_IB_MSG_311 uint16 = 12136
	ER_IB_MSG_312 uint16 = 12137
	ER_IB_MSG_313 uint16 = 12138
	ER_IB_MSG_314 uint16 = 12139
	ER_IB_MSG_315 uint16 = 12140
	ER_IB_MSG_316 uint16 = 12141
	ER_IB_MSG_317 uint16 = 12142
	ER_IB_MSG_318 uint16 = 12143
	ER_IB_MSG_319 uint16 = 12144
	ER_IB_MSG_320 uint16 = 12145
	ER_IB_MSG_321 uint16 = 12146
	ER_IB_MSG_322 uint16 = 12147
	ER_IB_MSG_323 uint16 = 12148
	ER_IB_MSG_324 uint16 = 12149
	ER_IB_MSG_325 uint16 = 12150
	ER_IB_MSG_326 uint16 = 12151
	ER_IB_MSG_327 uint16 = 12152
	ER_IB_MSG_328 uint16 = 12153
	ER_IB_MSG_329 uint16 = 12154
	ER_IB_MSG_330 uint16 = 12155
	ER_IB_MSG_331 uint16 = 12156
	ER_IB_MSG_332 uint16 = 12157
	ER_IB_MSG_333 uint16 = 12158
	ER_IB_MSG_334 uint16 = 12159
	ER_IB_MSG_335 uint16 = 12160
	ER_IB_MSG_336 uint16 = 12161
	ER_IB_MSG_337 uint16 = 12162
	ER_IB_MSG_338 uint16 = 12163
	ER_IB_MSG_339 uint16 = 12164
	ER_IB_MSG_340 uint16 = 12165
	ER_IB_MSG_341 uint16 = 12166
	ER_IB_MSG_342 uint16 = 12167
	ER_IB_MSG_343 uint16 = 12168
	ER_IB_MSG_344 uint16 = 12169
	ER_IB_MSG_345 uint16 = 12170
	ER_IB_MSG_346 uint16 = 12171
	ER_IB_MSG_347 uint16 = 12172
	ER_IB_MSG_348 uint16 = 12173
	ER_IB_MSG_349 uint16 = 12174
	ER_IB_MSG_350 uint16 = 12175
	//OBSOLETE_ER_IB_MSG_351 uint16 = 12176
	ER_IB_MSG_UNPROTECTED_LOCATION_ALLOWED uint16 = 12177
	//OBSOLETE_ER_IB_MSG_353 uint16 = 12178
	ER_IB_MSG_354 uint16 = 12179
	ER_IB_MSG_355 uint16 = 12180
	ER_IB_MSG_356 uint16 = 12181
	ER_IB_MSG_357 uint16 = 12182
	ER_IB_MSG_358 uint16 = 12183
	ER_IB_MSG_359 uint16 = 12184
	ER_IB_MSG_360 uint16 = 12185
	ER_IB_MSG_361 uint16 = 12186
	ER_IB_MSG_362 uint16 = 12187
	//OBSOLETE_ER_IB_MSG_363 uint16 = 12188
	ER_IB_MSG_364 uint16 = 12189
	ER_IB_MSG_365 uint16 = 12190
	ER_IB_MSG_IGNORE_SCAN_PATH uint16 = 12191
	ER_IB_MSG_367 uint16 = 12192
	ER_IB_MSG_368 uint16 = 12193
	ER_IB_MSG_369 uint16 = 12194
	ER_IB_MSG_370 uint16 = 12195
	ER_IB_MSG_371 uint16 = 12196
	ER_IB_MSG_372 uint16 = 12197
	ER_IB_MSG_373 uint16 = 12198
	ER_IB_MSG_374 uint16 = 12199
	ER_IB_MSG_375 uint16 = 12200
	ER_IB_MSG_376 uint16 = 12201
	ER_IB_MSG_377 uint16 = 12202
	ER_IB_MSG_378 uint16 = 12203
	ER_IB_MSG_379 uint16 = 12204
	ER_IB_MSG_380 uint16 = 12205
	ER_IB_MSG_381 uint16 = 12206
	ER_IB_MSG_382 uint16 = 12207
	ER_IB_MSG_383 uint16 = 12208
	ER_IB_MSG_384 uint16 = 12209
	ER_IB_MSG_385 uint16 = 12210
	ER_IB_MSG_386 uint16 = 12211
	ER_IB_MSG_387 uint16 = 12212
	ER_IB_MSG_GENERAL_TABLESPACE_UNDER_DATADIR uint16 = 12213
	ER_IB_MSG_IMPLICIT_TABLESPACE_IN_DATADIR uint16 = 12214
	ER_IB_MSG_390 uint16 = 12215
	ER_IB_MSG_391 uint16 = 12216
	ER_IB_MSG_392 uint16 = 12217
	ER_IB_MSG_393 uint16 = 12218
	ER_IB_MSG_394 uint16 = 12219
	ER_IB_MSG_395 uint16 = 12220
	ER_IB_MSG_396 uint16 = 12221
	ER_IB_MSG_397 uint16 = 12222
	ER_IB_MSG_398 uint16 = 12223
	ER_IB_MSG_399 uint16 = 12224
	//OBSOLETE_ER_IB_MSG_400 uint16 = 12225
	ER_IB_MSG_401 uint16 = 12226
	ER_IB_MSG_402 uint16 = 12227
	ER_IB_MSG_403 uint16 = 12228
	ER_IB_MSG_404 uint16 = 12229
	ER_IB_MSG_405 uint16 = 12230
	ER_IB_MSG_406 uint16 = 12231
	ER_IB_MSG_407 uint16 = 12232
	ER_IB_MSG_408 uint16 = 12233
	ER_IB_MSG_409 uint16 = 12234
	ER_IB_MSG_410 uint16 = 12235
	ER_IB_MSG_411 uint16 = 12236
	ER_IB_MSG_412 uint16 = 12237
	ER_IB_MSG_413 uint16 = 12238
	ER_IB_MSG_414 uint16 = 12239
	ER_IB_MSG_415 uint16 = 12240
	ER_IB_MSG_416 uint16 = 12241
	ER_IB_MSG_417 uint16 = 12242
	ER_IB_MSG_418 uint16 = 12243
	ER_IB_MSG_419 uint16 = 12244
	ER_IB_MSG_420 uint16 = 12245
	ER_IB_MSG_421 uint16 = 12246
	ER_IB_MSG_422 uint16 = 12247
	ER_IB_MSG_423 uint16 = 12248
	ER_IB_MSG_424 uint16 = 12249
	ER_IB_MSG_425 uint16 = 12250
	ER_IB_MSG_426 uint16 = 12251
	ER_IB_MSG_427 uint16 = 12252
	ER_IB_MSG_428 uint16 = 12253
	ER_IB_MSG_429 uint16 = 12254
	ER_IB_MSG_430 uint16 = 12255
	ER_IB_MSG_431 uint16 = 12256
	ER_IB_MSG_432 uint16 = 12257
	ER_IB_MSG_433 uint16 = 12258
	ER_IB_MSG_434 uint16 = 12259
	ER_IB_MSG_435 uint16 = 12260
	ER_IB_MSG_436 uint16 = 12261
	ER_IB_MSG_437 uint16 = 12262
	ER_IB_MSG_438 uint16 = 12263
	ER_IB_MSG_439 uint16 = 12264
	ER_IB_MSG_440 uint16 = 12265
	ER_IB_MSG_441 uint16 = 12266
	ER_IB_MSG_442 uint16 = 12267
	ER_IB_MSG_443 uint16 = 12268
	ER_IB_MSG_444 uint16 = 12269
	ER_IB_MSG_445 uint16 = 12270
	ER_IB_MSG_446 uint16 = 12271
	ER_IB_MSG_447 uint16 = 12272
	ER_IB_MSG_448 uint16 = 12273
	ER_IB_MSG_449 uint16 = 12274
	ER_IB_MSG_450 uint16 = 12275
	ER_IB_MSG_451 uint16 = 12276
	ER_IB_MSG_452 uint16 = 12277
	ER_IB_MSG_453 uint16 = 12278
	ER_IB_MSG_454 uint16 = 12279
	ER_IB_MSG_455 uint16 = 12280
	ER_IB_MSG_456 uint16 = 12281
	ER_IB_MSG_457 uint16 = 12282
	ER_IB_MSG_458 uint16 = 12283
	ER_IB_MSG_459 uint16 = 12284
	ER_IB_MSG_460 uint16 = 12285
	ER_IB_MSG_461 uint16 = 12286
	ER_IB_MSG_462 uint16 = 12287
	ER_IB_MSG_463 uint16 = 12288
	ER_IB_MSG_464 uint16 = 12289
	ER_IB_MSG_465 uint16 = 12290
	ER_IB_MSG_466 uint16 = 12291
	ER_IB_MSG_467 uint16 = 12292
	ER_IB_MSG_468 uint16 = 12293
	ER_IB_MSG_469 uint16 = 12294
	ER_IB_MSG_470 uint16 = 12295
	ER_IB_MSG_471 uint16 = 12296
	ER_IB_MSG_472 uint16 = 12297
	ER_IB_MSG_473 uint16 = 12298
	ER_IB_MSG_474 uint16 = 12299
	ER_IB_MSG_475 uint16 = 12300
	ER_IB_MSG_476 uint16 = 12301
	ER_IB_MSG_477 uint16 = 12302
	ER_IB_MSG_478 uint16 = 12303
	ER_IB_MSG_479 uint16 = 12304
	ER_IB_MSG_480 uint16 = 12305
	ER_IB_MSG_481 uint16 = 12306
	ER_IB_MSG_482 uint16 = 12307
	ER_IB_MSG_483 uint16 = 12308
	ER_IB_MSG_484 uint16 = 12309
	ER_IB_MSG_485 uint16 = 12310
	ER_IB_MSG_486 uint16 = 12311
	ER_IB_MSG_487 uint16 = 12312
	ER_IB_MSG_488 uint16 = 12313
	ER_IB_MSG_489 uint16 = 12314
	ER_IB_MSG_490 uint16 = 12315
	ER_IB_MSG_491 uint16 = 12316
	ER_IB_MSG_492 uint16 = 12317
	ER_IB_MSG_493 uint16 = 12318
	ER_IB_MSG_494 uint16 = 12319
	ER_IB_MSG_495 uint16 = 12320
	ER_IB_MSG_496 uint16 = 12321
	ER_IB_MSG_497 uint16 = 12322
	ER_IB_MSG_498 uint16 = 12323
	ER_IB_MSG_499 uint16 = 12324
	ER_IB_MSG_500 uint16 = 12325
	ER_IB_MSG_501 uint16 = 12326
	ER_IB_MSG_502 uint16 = 12327
	ER_IB_MSG_503 uint16 = 12328
	ER_IB_MSG_504 uint16 = 12329
	ER_IB_MSG_505 uint16 = 12330
	ER_IB_MSG_506 uint16 = 12331
	ER_IB_MSG_507 uint16 = 12332
	ER_IB_MSG_508 uint16 = 12333
	ER_IB_MSG_509 uint16 = 12334
	ER_IB_MSG_510 uint16 = 12335
	ER_IB_MSG_511 uint16 = 12336
	ER_IB_MSG_512 uint16 = 12337
	ER_IB_MSG_513 uint16 = 12338
	ER_IB_MSG_514 uint16 = 12339
	ER_IB_MSG_515 uint16 = 12340
	ER_IB_MSG_516 uint16 = 12341
	ER_IB_MSG_517 uint16 = 12342
	ER_IB_MSG_518 uint16 = 12343
	ER_IB_MSG_519 uint16 = 12344
	ER_IB_MSG_520 uint16 = 12345
	ER_IB_MSG_521 uint16 = 12346
	ER_IB_MSG_522 uint16 = 12347
	ER_IB_MSG_523 uint16 = 12348
	ER_IB_MSG_524 uint16 = 12349
	ER_IB_MSG_525 uint16 = 12350
	ER_IB_MSG_526 uint16 = 12351
	ER_IB_MSG_527 uint16 = 12352
	//OBSOLETE_ER_IB_MSG_528 uint16 = 12353
	//OBSOLETE_ER_IB_MSG_529 uint16 = 12354
	ER_IB_MSG_530 uint16 = 12355
	ER_IB_MSG_531 uint16 = 12356
	ER_IB_MSG_532 uint16 = 12357
	ER_IB_MSG_533 uint16 = 12358
	ER_IB_MSG_534 uint16 = 12359
	ER_IB_MSG_535 uint16 = 12360
	ER_IB_MSG_536 uint16 = 12361
	ER_IB_MSG_537 uint16 = 12362
	ER_IB_MSG_538 uint16 = 12363
	ER_IB_MSG_539 uint16 = 12364
	ER_IB_MSG_540 uint16 = 12365
	ER_IB_MSG_541 uint16 = 12366
	ER_IB_MSG_542 uint16 = 12367
	ER_IB_MSG_543 uint16 = 12368
	ER_IB_MSG_544 uint16 = 12369
	ER_IB_MSG_545 uint16 = 12370
	ER_IB_MSG_546 uint16 = 12371
	ER_IB_MSG_547 uint16 = 12372
	ER_IB_MSG_548 uint16 = 12373
	ER_IB_MSG_549 uint16 = 12374
	ER_IB_MSG_550 uint16 = 12375
	ER_IB_MSG_551 uint16 = 12376
	ER_IB_MSG_552 uint16 = 12377
	ER_IB_MSG_553 uint16 = 12378
	ER_IB_MSG_554 uint16 = 12379
	ER_IB_MSG_555 uint16 = 12380
	ER_IB_MSG_556 uint16 = 12381
	ER_IB_MSG_557 uint16 = 12382
	ER_IB_MSG_558 uint16 = 12383
	ER_IB_MSG_559 uint16 = 12384
	ER_IB_MSG_560 uint16 = 12385
	ER_IB_MSG_561 uint16 = 12386
	ER_IB_MSG_562 uint16 = 12387
	ER_IB_MSG_563 uint16 = 12388
	ER_IB_MSG_564 uint16 = 12389
	ER_IB_MSG_INVALID_LOCATION_FOR_TABLE uint16 = 12390
	ER_IB_MSG_566 uint16 = 12391
	ER_IB_MSG_567 uint16 = 12392
	ER_IB_MSG_568 uint16 = 12393
	ER_IB_MSG_569 uint16 = 12394
	ER_IB_MSG_570 uint16 = 12395
	ER_IB_MSG_571 uint16 = 12396
	//OBSOLETE_ER_IB_MSG_572 uint16 = 12397
	ER_IB_MSG_573 uint16 = 12398
	ER_IB_MSG_574 uint16 = 12399
	//OBSOLETE_ER_IB_MSG_575 uint16 = 12400
	//OBSOLETE_ER_IB_MSG_576 uint16 = 12401
	//OBSOLETE_ER_IB_MSG_577 uint16 = 12402
	ER_IB_MSG_578 uint16 = 12403
	ER_IB_MSG_579 uint16 = 12404
	ER_IB_MSG_580 uint16 = 12405
	ER_IB_MSG_581 uint16 = 12406
	ER_IB_MSG_582 uint16 = 12407
	ER_IB_MSG_583 uint16 = 12408
	ER_IB_MSG_584 uint16 = 12409
	ER_IB_MSG_585 uint16 = 12410
	ER_IB_MSG_586 uint16 = 12411
	ER_IB_MSG_587 uint16 = 12412
	ER_IB_MSG_588 uint16 = 12413
	ER_IB_MSG_589 uint16 = 12414
	ER_IB_MSG_590 uint16 = 12415
	ER_IB_MSG_591 uint16 = 12416
	ER_IB_MSG_592 uint16 = 12417
	ER_IB_MSG_593 uint16 = 12418
	ER_IB_MSG_594 uint16 = 12419
	ER_IB_MSG_595 uint16 = 12420
	ER_IB_MSG_596 uint16 = 12421
	ER_IB_MSG_597 uint16 = 12422
	ER_IB_MSG_598 uint16 = 12423
	ER_IB_MSG_599 uint16 = 12424
	ER_IB_MSG_600 uint16 = 12425
	ER_IB_MSG_601 uint16 = 12426
	ER_IB_MSG_602 uint16 = 12427
	ER_IB_MSG_603 uint16 = 12428
	ER_IB_MSG_604 uint16 = 12429
	ER_IB_MSG_605 uint16 = 12430
	ER_IB_MSG_606 uint16 = 12431
	ER_IB_MSG_607 uint16 = 12432
	ER_IB_MSG_608 uint16 = 12433
	ER_IB_MSG_609 uint16 = 12434
	ER_IB_MSG_610 uint16 = 12435
	ER_IB_MSG_611 uint16 = 12436
	ER_IB_MSG_612 uint16 = 12437
	ER_IB_MSG_613 uint16 = 12438
	ER_IB_MSG_614 uint16 = 12439
	ER_IB_MSG_615 uint16 = 12440
	ER_IB_MSG_616 uint16 = 12441
	ER_IB_MSG_617 uint16 = 12442
	ER_IB_MSG_618 uint16 = 12443
	ER_IB_MSG_619 uint16 = 12444
	ER_IB_MSG_620 uint16 = 12445
	ER_IB_MSG_621 uint16 = 12446
	ER_IB_MSG_622 uint16 = 12447
	ER_IB_MSG_623 uint16 = 12448
	ER_IB_MSG_624 uint16 = 12449
	ER_IB_MSG_625 uint16 = 12450
	ER_IB_MSG_626 uint16 = 12451
	ER_IB_MSG_627 uint16 = 12452
	ER_IB_MSG_628 uint16 = 12453
	ER_IB_MSG_629 uint16 = 12454
	ER_IB_MSG_630 uint16 = 12455
	ER_IB_MSG_631 uint16 = 12456
	ER_IB_MSG_632 uint16 = 12457
	ER_IB_MSG_633 uint16 = 12458
	ER_IB_MSG_634 uint16 = 12459
	ER_IB_MSG_635 uint16 = 12460
	ER_IB_MSG_636 uint16 = 12461
	ER_IB_MSG_637 uint16 = 12462
	ER_IB_MSG_638 uint16 = 12463
	ER_IB_MSG_639 uint16 = 12464
	//OBSOLETE_ER_IB_MSG_640 uint16 = 12465
	//OBSOLETE_ER_IB_MSG_641 uint16 = 12466
	ER_IB_MSG_642 uint16 = 12467
	ER_IB_MSG_643 uint16 = 12468
	ER_IB_MSG_644 uint16 = 12469
	ER_IB_MSG_645 uint16 = 12470
	ER_IB_MSG_646 uint16 = 12471
	ER_IB_MSG_647 uint16 = 12472
	ER_IB_MSG_648 uint16 = 12473
	ER_IB_MSG_649 uint16 = 12474
	ER_IB_MSG_650 uint16 = 12475
	ER_IB_MSG_651 uint16 = 12476
	ER_IB_MSG_652 uint16 = 12477
	ER_IB_MSG_DDL_LOG_DELETE_BY_ID_OK uint16 = 12478
	ER_IB_MSG_654 uint16 = 12479
	ER_IB_MSG_655 uint16 = 12480
	ER_IB_MSG_656 uint16 = 12481
	ER_IB_MSG_657 uint16 = 12482
	ER_IB_MSG_658 uint16 = 12483
	ER_IB_MSG_659 uint16 = 12484
	ER_IB_MSG_660 uint16 = 12485
	ER_IB_MSG_661 uint16 = 12486
	ER_IB_MSG_662 uint16 = 12487
	ER_IB_MSG_663 uint16 = 12488
	//OBSOLETE_ER_IB_MSG_664 uint16 = 12489
	//OBSOLETE_ER_IB_MSG_665 uint16 = 12490
	//OBSOLETE_ER_IB_MSG_666 uint16 = 12491
	//OBSOLETE_ER_IB_MSG_667 uint16 = 12492
	//OBSOLETE_ER_IB_MSG_668 uint16 = 12493
	//OBSOLETE_ER_IB_MSG_669 uint16 = 12494
	//OBSOLETE_ER_IB_MSG_670 uint16 = 12495
	//OBSOLETE_ER_IB_MSG_671 uint16 = 12496
	//OBSOLETE_ER_IB_MSG_672 uint16 = 12497
	//OBSOLETE_ER_IB_MSG_673 uint16 = 12498
	//OBSOLETE_ER_IB_MSG_674 uint16 = 12499
	//OBSOLETE_ER_IB_MSG_675 uint16 = 12500
	//OBSOLETE_ER_IB_MSG_676 uint16 = 12501
	//OBSOLETE_ER_IB_MSG_677 uint16 = 12502
	//OBSOLETE_ER_IB_MSG_678 uint16 = 12503
	//OBSOLETE_ER_IB_MSG_679 uint16 = 12504
	//OBSOLETE_ER_IB_MSG_680 uint16 = 12505
	//OBSOLETE_ER_IB_MSG_681 uint16 = 12506
	//OBSOLETE_ER_IB_MSG_682 uint16 = 12507
	//OBSOLETE_ER_IB_MSG_683 uint16 = 12508
	//OBSOLETE_ER_IB_MSG_684 uint16 = 12509
	//OBSOLETE_ER_IB_MSG_685 uint16 = 12510
	//OBSOLETE_ER_IB_MSG_686 uint16 = 12511
	//OBSOLETE_ER_IB_MSG_687 uint16 = 12512
	//OBSOLETE_ER_IB_MSG_688 uint16 = 12513
	//OBSOLETE_ER_IB_MSG_689 uint16 = 12514
	//OBSOLETE_ER_IB_MSG_690 uint16 = 12515
	//OBSOLETE_ER_IB_MSG_691 uint16 = 12516
	//OBSOLETE_ER_IB_MSG_692 uint16 = 12517
	//OBSOLETE_ER_IB_MSG_693 uint16 = 12518
	ER_IB_MSG_694 uint16 = 12519
	ER_IB_MSG_695 uint16 = 12520
	ER_IB_MSG_696 uint16 = 12521
	ER_IB_MSG_697 uint16 = 12522
	ER_IB_MSG_698 uint16 = 12523
	ER_IB_MSG_699 uint16 = 12524
	ER_IB_MSG_700 uint16 = 12525
	ER_IB_MSG_701 uint16 = 12526
	//OBSOLETE_ER_IB_MSG_702 uint16 = 12527
	//OBSOLETE_ER_IB_MSG_703 uint16 = 12528
	ER_IB_MSG_704 uint16 = 12529
	ER_IB_MSG_705 uint16 = 12530
	ER_IB_MSG_706 uint16 = 12531
	ER_IB_MSG_707 uint16 = 12532
	ER_IB_MSG_708 uint16 = 12533
	ER_IB_MSG_709 uint16 = 12534
	ER_IB_MSG_710 uint16 = 12535
	ER_IB_MSG_711 uint16 = 12536
	ER_IB_MSG_712 uint16 = 12537
	ER_IB_MSG_713 uint16 = 12538
	ER_IB_MSG_714 uint16 = 12539
	ER_IB_MSG_715 uint16 = 12540
	ER_IB_MSG_716 uint16 = 12541
	ER_IB_MSG_717 uint16 = 12542
	ER_IB_MSG_718 uint16 = 12543
	ER_IB_MSG_719 uint16 = 12544
	ER_IB_MSG_720 uint16 = 12545
	ER_IB_MSG_721 uint16 = 12546
	ER_IB_MSG_722 uint16 = 12547
	ER_IB_MSG_723 uint16 = 12548
	ER_IB_MSG_724 uint16 = 12549
	ER_IB_MSG_725 uint16 = 12550
	ER_IB_MSG_726 uint16 = 12551
	ER_IB_MSG_727 uint16 = 12552
	ER_IB_MSG_728 uint16 = 12553
	ER_IB_MSG_729 uint16 = 12554
	ER_IB_MSG_730 uint16 = 12555
	ER_IB_MSG_731 uint16 = 12556
	ER_IB_MSG_732 uint16 = 12557
	ER_IB_MSG_733 uint16 = 12558
	ER_IB_MSG_734 uint16 = 12559
	ER_IB_MSG_735 uint16 = 12560
	ER_IB_MSG_736 uint16 = 12561
	ER_IB_MSG_737 uint16 = 12562
	ER_IB_MSG_738 uint16 = 12563
	ER_IB_MSG_739 uint16 = 12564
	ER_IB_MSG_740 uint16 = 12565
	ER_IB_MSG_741 uint16 = 12566
	ER_IB_MSG_742 uint16 = 12567
	ER_IB_MSG_743 uint16 = 12568
	ER_IB_MSG_744 uint16 = 12569
	ER_IB_MSG_745 uint16 = 12570
	ER_IB_MSG_746 uint16 = 12571
	ER_IB_MSG_747 uint16 = 12572
	ER_IB_MSG_748 uint16 = 12573
	ER_IB_MSG_749 uint16 = 12574
	ER_IB_MSG_750 uint16 = 12575
	ER_IB_MSG_751 uint16 = 12576
	ER_IB_MSG_752 uint16 = 12577
	ER_IB_MSG_753 uint16 = 12578
	ER_IB_MSG_754 uint16 = 12579
	ER_IB_MSG_755 uint16 = 12580
	ER_IB_MSG_756 uint16 = 12581
	ER_IB_MSG_757 uint16 = 12582
	ER_IB_MSG_758 uint16 = 12583
	ER_IB_MSG_759 uint16 = 12584
	ER_IB_MSG_760 uint16 = 12585
	ER_IB_MSG_761 uint16 = 12586
	ER_IB_MSG_762 uint16 = 12587
	ER_IB_MSG_763 uint16 = 12588
	ER_IB_MSG_764 uint16 = 12589
	ER_IB_MSG_765 uint16 = 12590
	ER_IB_MSG_766 uint16 = 12591
	ER_IB_MSG_767 uint16 = 12592
	ER_IB_MSG_768 uint16 = 12593
	ER_IB_MSG_769 uint16 = 12594
	ER_IB_MSG_770 uint16 = 12595
	ER_IB_MSG_771 uint16 = 12596
	ER_IB_MSG_772 uint16 = 12597
	ER_IB_MSG_773 uint16 = 12598
	ER_IB_MSG_774 uint16 = 12599
	ER_IB_MSG_775 uint16 = 12600
	ER_IB_MSG_776 uint16 = 12601
	ER_IB_MSG_777 uint16 = 12602
	ER_IB_MSG_778 uint16 = 12603
	ER_IB_MSG_779 uint16 = 12604
	ER_IB_MSG_780 uint16 = 12605
	ER_IB_MSG_781 uint16 = 12606
	ER_IB_MSG_782 uint16 = 12607
	ER_IB_MSG_783 uint16 = 12608
	ER_IB_MSG_784 uint16 = 12609
	ER_IB_MSG_785 uint16 = 12610
	ER_IB_MSG_786 uint16 = 12611
	ER_IB_MSG_787 uint16 = 12612
	ER_IB_MSG_788 uint16 = 12613
	ER_IB_MSG_789 uint16 = 12614
	ER_IB_MSG_790 uint16 = 12615
	ER_IB_MSG_791 uint16 = 12616
	ER_IB_MSG_792 uint16 = 12617
	ER_IB_MSG_793 uint16 = 12618
	ER_IB_MSG_794 uint16 = 12619
	ER_IB_MSG_795 uint16 = 12620
	ER_IB_MSG_796 uint16 = 12621
	ER_IB_MSG_797 uint16 = 12622
	ER_IB_MSG_798 uint16 = 12623
	ER_IB_MSG_799 uint16 = 12624
	ER_IB_MSG_800 uint16 = 12625
	ER_IB_MSG_801 uint16 = 12626
	ER_IB_MSG_802 uint16 = 12627
	ER_IB_MSG_803 uint16 = 12628
	ER_IB_MSG_804 uint16 = 12629
	ER_IB_MSG_805 uint16 = 12630
	ER_IB_MSG_806 uint16 = 12631
	ER_IB_MSG_807 uint16 = 12632
	ER_IB_MSG_808 uint16 = 12633
	ER_IB_MSG_809 uint16 = 12634
	ER_IB_MSG_810 uint16 = 12635
	ER_IB_MSG_811 uint16 = 12636
	ER_IB_MSG_812 uint16 = 12637
	ER_IB_MSG_813 uint16 = 12638
	ER_IB_MSG_814 uint16 = 12639
	ER_IB_MSG_815 uint16 = 12640
	ER_IB_MSG_816 uint16 = 12641
	ER_IB_MSG_817 uint16 = 12642
	ER_IB_MSG_818 uint16 = 12643
	ER_IB_MSG_819 uint16 = 12644
	ER_IB_MSG_820 uint16 = 12645
	ER_IB_MSG_821 uint16 = 12646
	ER_IB_MSG_822 uint16 = 12647
	ER_IB_MSG_823 uint16 = 12648
	ER_IB_MSG_824 uint16 = 12649
	ER_IB_MSG_825 uint16 = 12650
	ER_IB_MSG_826 uint16 = 12651
	ER_IB_MSG_827 uint16 = 12652
	ER_IB_MSG_828 uint16 = 12653
	ER_IB_MSG_829 uint16 = 12654
	ER_IB_MSG_830 uint16 = 12655
	ER_IB_MSG_831 uint16 = 12656
	ER_IB_MSG_832 uint16 = 12657
	ER_IB_MSG_833 uint16 = 12658
	ER_IB_MSG_834 uint16 = 12659
	ER_IB_MSG_835 uint16 = 12660
	ER_IB_MSG_836 uint16 = 12661
	ER_IB_MSG_837 uint16 = 12662
	ER_IB_MSG_838 uint16 = 12663
	ER_IB_MSG_839 uint16 = 12664
	ER_IB_MSG_840 uint16 = 12665
	ER_IB_MSG_841 uint16 = 12666
	ER_IB_MSG_842 uint16 = 12667
	ER_IB_MSG_843 uint16 = 12668
	ER_IB_MSG_844 uint16 = 12669
	ER_IB_MSG_845 uint16 = 12670
	ER_IB_MSG_846 uint16 = 12671
	ER_IB_MSG_847 uint16 = 12672
	ER_IB_MSG_848 uint16 = 12673
	ER_IB_MSG_849 uint16 = 12674
	ER_IB_MSG_850 uint16 = 12675
	ER_IB_MSG_851 uint16 = 12676
	ER_IB_MSG_852 uint16 = 12677
	ER_IB_MSG_853 uint16 = 12678
	ER_IB_MSG_854 uint16 = 12679
	ER_IB_MSG_855 uint16 = 12680
	ER_IB_MSG_856 uint16 = 12681
	ER_IB_MSG_857 uint16 = 12682
	ER_IB_MSG_858 uint16 = 12683
	ER_IB_MSG_859 uint16 = 12684
	ER_IB_MSG_860 uint16 = 12685
	ER_IB_MSG_861 uint16 = 12686
	ER_IB_MSG_862 uint16 = 12687
	ER_IB_MSG_863 uint16 = 12688
	ER_IB_MSG_864 uint16 = 12689
	ER_IB_MSG_865 uint16 = 12690
	ER_IB_MSG_866 uint16 = 12691
	ER_IB_MSG_867 uint16 = 12692
	ER_IB_MSG_868 uint16 = 12693
	ER_IB_MSG_869 uint16 = 12694
	ER_IB_MSG_870 uint16 = 12695
	ER_IB_MSG_871 uint16 = 12696
	ER_IB_MSG_872 uint16 = 12697
	ER_IB_MSG_873 uint16 = 12698
	ER_IB_MSG_874 uint16 = 12699
	ER_IB_MSG_875 uint16 = 12700
	ER_IB_MSG_876 uint16 = 12701
	ER_IB_MSG_877 uint16 = 12702
	ER_IB_MSG_878 uint16 = 12703
	ER_IB_MSG_879 uint16 = 12704
	ER_IB_MSG_880 uint16 = 12705
	ER_IB_MSG_881 uint16 = 12706
	ER_IB_MSG_882 uint16 = 12707
	ER_IB_MSG_883 uint16 = 12708
	ER_IB_MSG_884 uint16 = 12709
	ER_IB_MSG_885 uint16 = 12710
	ER_IB_MSG_886 uint16 = 12711
	ER_IB_MSG_887 uint16 = 12712
	ER_IB_MSG_888 uint16 = 12713
	ER_IB_MSG_889 uint16 = 12714
	ER_IB_MSG_890 uint16 = 12715
	ER_IB_MSG_891 uint16 = 12716
	ER_IB_MSG_892 uint16 = 12717
	ER_IB_MSG_893 uint16 = 12718
	ER_IB_MSG_894 uint16 = 12719
	ER_IB_MSG_895 uint16 = 12720
	ER_IB_MSG_896 uint16 = 12721
	ER_IB_MSG_897 uint16 = 12722
	ER_IB_MSG_898 uint16 = 12723
	ER_IB_MSG_899 uint16 = 12724
	ER_IB_MSG_900 uint16 = 12725
	ER_IB_MSG_901 uint16 = 12726
	ER_IB_MSG_902 uint16 = 12727
	ER_IB_MSG_903 uint16 = 12728
	ER_IB_MSG_904 uint16 = 12729
	ER_IB_MSG_905 uint16 = 12730
	ER_IB_MSG_906 uint16 = 12731
	ER_IB_MSG_907 uint16 = 12732
	ER_IB_MSG_908 uint16 = 12733
	ER_IB_MSG_909 uint16 = 12734
	ER_IB_MSG_910 uint16 = 12735
	ER_IB_MSG_911 uint16 = 12736
	ER_IB_MSG_912 uint16 = 12737
	ER_IB_MSG_913 uint16 = 12738
	ER_IB_MSG_914 uint16 = 12739
	ER_IB_MSG_915 uint16 = 12740
	ER_IB_MSG_916 uint16 = 12741
	ER_IB_MSG_917 uint16 = 12742
	ER_IB_MSG_918 uint16 = 12743
	ER_IB_MSG_919 uint16 = 12744
	ER_IB_MSG_920 uint16 = 12745
	ER_IB_MSG_921 uint16 = 12746
	ER_IB_MSG_922 uint16 = 12747
	ER_IB_MSG_923 uint16 = 12748
	ER_IB_MSG_924 uint16 = 12749
	ER_IB_MSG_925 uint16 = 12750
	ER_IB_MSG_926 uint16 = 12751
	ER_IB_MSG_927 uint16 = 12752
	ER_IB_MSG_928 uint16 = 12753
	ER_IB_MSG_929 uint16 = 12754
	ER_IB_MSG_930 uint16 = 12755
	ER_IB_MSG_931 uint16 = 12756
	ER_IB_MSG_932 uint16 = 12757
	ER_IB_MSG_933 uint16 = 12758
	ER_IB_MSG_934 uint16 = 12759
	ER_IB_MSG_935 uint16 = 12760
	ER_IB_MSG_936 uint16 = 12761
	ER_IB_MSG_937 uint16 = 12762
	ER_IB_MSG_938 uint16 = 12763
	ER_IB_MSG_939 uint16 = 12764
	ER_IB_MSG_940 uint16 = 12765
	ER_IB_MSG_941 uint16 = 12766
	ER_IB_MSG_942 uint16 = 12767
	ER_IB_MSG_943 uint16 = 12768
	ER_IB_MSG_944 uint16 = 12769
	ER_IB_MSG_945 uint16 = 12770
	ER_IB_MSG_946 uint16 = 12771
	ER_IB_MSG_947 uint16 = 12772
	ER_IB_MSG_948 uint16 = 12773
	ER_IB_MSG_949 uint16 = 12774
	ER_IB_MSG_950 uint16 = 12775
	ER_IB_MSG_951 uint16 = 12776
	ER_IB_MSG_952 uint16 = 12777
	ER_IB_MSG_953 uint16 = 12778
	ER_IB_MSG_954 uint16 = 12779
	ER_IB_MSG_955 uint16 = 12780
	ER_IB_MSG_956 uint16 = 12781
	ER_IB_MSG_957 uint16 = 12782
	ER_IB_MSG_958 uint16 = 12783
	ER_IB_MSG_959 uint16 = 12784
	ER_IB_MSG_960 uint16 = 12785
	ER_IB_MSG_961 uint16 = 12786
	ER_IB_MSG_962 uint16 = 12787
	ER_IB_MSG_963 uint16 = 12788
	ER_IB_MSG_964 uint16 = 12789
	ER_IB_MSG_965 uint16 = 12790
	ER_IB_MSG_966 uint16 = 12791
	ER_IB_MSG_967 uint16 = 12792
	ER_IB_MSG_968 uint16 = 12793
	ER_IB_MSG_969 uint16 = 12794
	ER_IB_MSG_970 uint16 = 12795
	ER_IB_MSG_971 uint16 = 12796
	ER_IB_MSG_972 uint16 = 12797
	ER_IB_MSG_973 uint16 = 12798
	ER_IB_MSG_974 uint16 = 12799
	ER_IB_MSG_975 uint16 = 12800
	ER_IB_MSG_976 uint16 = 12801
	ER_IB_MSG_977 uint16 = 12802
	ER_IB_MSG_978 uint16 = 12803
	ER_IB_MSG_979 uint16 = 12804
	ER_IB_MSG_980 uint16 = 12805
	ER_IB_MSG_981 uint16 = 12806
	ER_IB_MSG_982 uint16 = 12807
	ER_IB_MSG_983 uint16 = 12808
	ER_IB_MSG_984 uint16 = 12809
	ER_IB_MSG_985 uint16 = 12810
	ER_IB_MSG_986 uint16 = 12811
	ER_IB_MSG_987 uint16 = 12812
	ER_IB_MSG_988 uint16 = 12813
	ER_IB_MSG_989 uint16 = 12814
	ER_IB_MSG_990 uint16 = 12815
	ER_IB_MSG_991 uint16 = 12816
	ER_IB_MSG_992 uint16 = 12817
	ER_IB_MSG_993 uint16 = 12818
	ER_IB_MSG_994 uint16 = 12819
	ER_IB_MSG_995 uint16 = 12820
	ER_IB_MSG_996 uint16 = 12821
	ER_IB_MSG_997 uint16 = 12822
	ER_IB_MSG_998 uint16 = 12823
	ER_IB_MSG_999 uint16 = 12824
	ER_IB_MSG_1000 uint16 = 12825
	ER_IB_MSG_1001 uint16 = 12826
	ER_IB_MSG_1002 uint16 = 12827
	ER_IB_MSG_1003 uint16 = 12828
	ER_IB_MSG_1004 uint16 = 12829
	ER_IB_MSG_1005 uint16 = 12830
	ER_IB_MSG_1006 uint16 = 12831
	ER_IB_MSG_1007 uint16 = 12832
	ER_IB_MSG_1008 uint16 = 12833
	ER_IB_MSG_1009 uint16 = 12834
	ER_IB_MSG_1010 uint16 = 12835
	ER_IB_MSG_1011 uint16 = 12836
	ER_IB_MSG_1012 uint16 = 12837
	ER_IB_MSG_1013 uint16 = 12838
	ER_IB_MSG_1014 uint16 = 12839
	ER_IB_MSG_1015 uint16 = 12840
	ER_IB_MSG_1016 uint16 = 12841
	ER_IB_MSG_1017 uint16 = 12842
	ER_IB_MSG_1018 uint16 = 12843
	ER_IB_MSG_1019 uint16 = 12844
	ER_IB_MSG_1020 uint16 = 12845
	ER_IB_MSG_1021 uint16 = 12846
	ER_IB_MSG_1022 uint16 = 12847
	ER_IB_MSG_1023 uint16 = 12848
	ER_IB_MSG_1024 uint16 = 12849
	ER_IB_MSG_1025 uint16 = 12850
	ER_IB_MSG_1026 uint16 = 12851
	ER_IB_MSG_1027 uint16 = 12852
	ER_IB_MSG_1028 uint16 = 12853
	ER_IB_MSG_1029 uint16 = 12854
	ER_IB_MSG_1030 uint16 = 12855
	ER_IB_MSG_1031 uint16 = 12856
	ER_IB_MSG_1032 uint16 = 12857
	ER_IB_MSG_1033 uint16 = 12858
	ER_IB_MSG_1034 uint16 = 12859
	ER_IB_MSG_1035 uint16 = 12860
	ER_IB_MSG_1036 uint16 = 12861
	ER_IB_MSG_1037 uint16 = 12862
	ER_IB_MSG_1038 uint16 = 12863
	ER_IB_MSG_1039 uint16 = 12864
	ER_IB_MSG_1040 uint16 = 12865
	ER_IB_MSG_1041 uint16 = 12866
	ER_IB_MSG_1042 uint16 = 12867
	ER_IB_MSG_1043 uint16 = 12868
	ER_IB_MSG_1044 uint16 = 12869
	ER_IB_MSG_1045 uint16 = 12870
	ER_IB_MSG_1046 uint16 = 12871
	ER_IB_MSG_1047 uint16 = 12872
	ER_IB_MSG_1048 uint16 = 12873
	ER_IB_MSG_1049 uint16 = 12874
	//OBSOLETE_ER_IB_MSG_1050 uint16 = 12875
	ER_IB_MSG_1051 uint16 = 12876
	ER_IB_MSG_1052 uint16 = 12877
	ER_IB_MSG_1053 uint16 = 12878
	ER_IB_MSG_1054 uint16 = 12879
	ER_IB_MSG_1055 uint16 = 12880
	ER_IB_MSG_1056 uint16 = 12881
	ER_IB_MSG_1057 uint16 = 12882
	ER_IB_MSG_1058 uint16 = 12883
	ER_IB_MSG_1059 uint16 = 12884
	ER_IB_MSG_1060 uint16 = 12885
	ER_IB_MSG_1061 uint16 = 12886
	ER_IB_MSG_1062 uint16 = 12887
	ER_IB_MSG_1063 uint16 = 12888
	ER_IB_MSG_1064 uint16 = 12889
	ER_IB_MSG_1065 uint16 = 12890
	ER_IB_MSG_1066 uint16 = 12891
	ER_IB_MSG_1067 uint16 = 12892
	ER_IB_MSG_1068 uint16 = 12893
	ER_IB_MSG_1069 uint16 = 12894
	ER_IB_MSG_1070 uint16 = 12895
	ER_IB_MSG_1071 uint16 = 12896
	ER_IB_MSG_1072 uint16 = 12897
	ER_IB_MSG_1073 uint16 = 12898
	ER_IB_MSG_1074 uint16 = 12899
	ER_IB_MSG_1075 uint16 = 12900
	ER_IB_MSG_1076 uint16 = 12901
	ER_IB_MSG_1077 uint16 = 12902
	ER_IB_MSG_1078 uint16 = 12903
	ER_IB_MSG_1079 uint16 = 12904
	ER_IB_MSG_1080 uint16 = 12905
	ER_IB_MSG_1081 uint16 = 12906
	ER_IB_MSG_1082 uint16 = 12907
	ER_IB_MSG_1083 uint16 = 12908
	ER_IB_MSG_CANNOT_OPEN_57_UNDO uint16 = 12909
	ER_IB_MSG_1085 uint16 = 12910
	ER_IB_MSG_1086 uint16 = 12911
	ER_IB_MSG_1087 uint16 = 12912
	ER_IB_MSG_1088 uint16 = 12913
	ER_IB_MSG_1089 uint16 = 12914
	ER_IB_MSG_1090 uint16 = 12915
	ER_IB_MSG_1091 uint16 = 12916
	ER_IB_MSG_1092 uint16 = 12917
	ER_IB_MSG_1093 uint16 = 12918
	ER_IB_MSG_1094 uint16 = 12919
	ER_IB_MSG_1095 uint16 = 12920
	ER_IB_MSG_1096 uint16 = 12921
	ER_IB_MSG_1097 uint16 = 12922
	ER_IB_MSG_1098 uint16 = 12923
	ER_IB_MSG_1099 uint16 = 12924
	ER_IB_MSG_1100 uint16 = 12925
	ER_IB_MSG_1101 uint16 = 12926
	ER_IB_MSG_1102 uint16 = 12927
	ER_IB_MSG_1103 uint16 = 12928
	ER_IB_MSG_1104 uint16 = 12929
	ER_IB_MSG_1105 uint16 = 12930
	ER_IB_MSG_1106 uint16 = 12931
	ER_IB_MSG_1107 uint16 = 12932
	ER_IB_MSG_1108 uint16 = 12933
	ER_IB_MSG_1109 uint16 = 12934
	ER_IB_MSG_1110 uint16 = 12935
	ER_IB_MSG_1111 uint16 = 12936
	ER_IB_MSG_1112 uint16 = 12937
	ER_IB_MSG_1113 uint16 = 12938
	ER_IB_MSG_1114 uint16 = 12939
	ER_IB_MSG_1115 uint16 = 12940
	ER_IB_MSG_1116 uint16 = 12941
	ER_IB_MSG_1117 uint16 = 12942
	//OBSOLETE_ER_IB_MSG_1118 uint16 = 12943
	ER_IB_MSG_1119 uint16 = 12944
	ER_IB_MSG_1120 uint16 = 12945
	ER_IB_MSG_1121 uint16 = 12946
	ER_IB_MSG_1122 uint16 = 12947
	ER_IB_MSG_1123 uint16 = 12948
	ER_IB_MSG_1124 uint16 = 12949
	ER_IB_MSG_1125 uint16 = 12950
	ER_IB_MSG_1126 uint16 = 12951
	ER_IB_MSG_1127 uint16 = 12952
	ER_IB_MSG_1128 uint16 = 12953
	ER_IB_MSG_1129 uint16 = 12954
	ER_IB_MSG_1130 uint16 = 12955
	ER_IB_MSG_1131 uint16 = 12956
	ER_IB_MSG_1132 uint16 = 12957
	ER_IB_MSG_1133 uint16 = 12958
	ER_IB_MSG_1134 uint16 = 12959
	ER_IB_MSG_1135 uint16 = 12960
	ER_IB_MSG_1136 uint16 = 12961
	ER_IB_MSG_1137 uint16 = 12962
	ER_IB_MSG_1138 uint16 = 12963
	ER_IB_MSG_1139 uint16 = 12964
	ER_IB_MSG_1140 uint16 = 12965
	ER_IB_MSG_1141 uint16 = 12966
	ER_IB_MSG_1142 uint16 = 12967
	ER_IB_MSG_1143 uint16 = 12968
	ER_IB_MSG_1144 uint16 = 12969
	ER_IB_MSG_1145 uint16 = 12970
	ER_IB_MSG_1146 uint16 = 12971
	ER_IB_MSG_1147 uint16 = 12972
	ER_IB_MSG_1148 uint16 = 12973
	ER_IB_MSG_1149 uint16 = 12974
	ER_IB_MSG_1150 uint16 = 12975
	ER_IB_MSG_1151 uint16 = 12976
	ER_IB_MSG_1152 uint16 = 12977
	//OBSOLETE_ER_IB_MSG_1153 uint16 = 12978
	ER_IB_MSG_1154 uint16 = 12979
	ER_IB_MSG_1155 uint16 = 12980
	ER_IB_MSG_1156 uint16 = 12981
	ER_IB_MSG_1157 uint16 = 12982
	ER_IB_MSG_1158 uint16 = 12983
	ER_IB_MSG_1159 uint16 = 12984
	ER_IB_MSG_1160 uint16 = 12985
	ER_IB_MSG_1161 uint16 = 12986
	ER_IB_MSG_1162 uint16 = 12987
	ER_IB_MSG_1163 uint16 = 12988
	ER_IB_MSG_1164 uint16 = 12989
	ER_IB_MSG_1165 uint16 = 12990
	ER_IB_MSG_UNDO_TRUNCATE_FAIL_TO_READ_LOG_FILE uint16 = 12991
	ER_IB_MSG_UNDO_MARKED_FOR_TRUNCATE uint16 = 12992
	//OBSOLETE_ER_IB_MSG_UNDO_INJECT_BEFORE_MDL uint16 = 12993
	ER_IB_MSG_UNDO_TRUNCATE_START uint16 = 12994
	//OBSOLETE_ER_IB_MSG_UNDO_INJECT_BEFORE_DDL_LOG_START uint16 = 12995
	ER_IB_MSG_UNDO_TRUNCATE_DELAY_BY_LOG_CREATE uint16 = 12996
	//OBSOLETE_ER_IB_MSG_UNDO_INJECT_BEFORE_TRUNCATE uint16 = 12997
	ER_IB_MSG_UNDO_TRUNCATE_DELAY_BY_FAILURE uint16 = 12998
	//OBSOLETE_ER_IB_MSG_UNDO_INJECT_BEFORE_STATE_UPDATE uint16 = 12999
	ER_IB_MSG_UNDO_TRUNCATE_COMPLETE uint16 = 13000
	//OBSOLETE_ER_IB_MSG_UNDO_INJECT_TRUNCATE_DONE uint16 = 13001
	ER_IB_MSG_1177 uint16 = 13002
	ER_IB_MSG_1178 uint16 = 13003
	ER_IB_MSG_1179 uint16 = 13004
	ER_IB_MSG_1180 uint16 = 13005
	ER_IB_MSG_1181 uint16 = 13006
	ER_IB_MSG_1182 uint16 = 13007
	ER_IB_MSG_1183 uint16 = 13008
	ER_IB_MSG_1184 uint16 = 13009
	ER_IB_MSG_1185 uint16 = 13010
	ER_IB_MSG_1186 uint16 = 13011
	ER_IB_MSG_1187 uint16 = 13012
	ER_IB_MSG_1188 uint16 = 13013
	ER_IB_MSG_1189 uint16 = 13014
	ER_IB_MSG_TRX_RECOVERY_ROLLBACK_COMPLETED uint16 = 13015
	ER_IB_MSG_1191 uint16 = 13016
	ER_IB_MSG_1192 uint16 = 13017
	ER_IB_MSG_1193 uint16 = 13018
	ER_IB_MSG_1194 uint16 = 13019
	ER_IB_MSG_1195 uint16 = 13020
	ER_IB_MSG_1196 uint16 = 13021
	ER_IB_MSG_1197 uint16 = 13022
	ER_IB_MSG_1198 uint16 = 13023
	ER_IB_MSG_1199 uint16 = 13024
	ER_IB_MSG_1200 uint16 = 13025
	ER_IB_MSG_1201 uint16 = 13026
	ER_IB_MSG_1202 uint16 = 13027
	ER_IB_MSG_1203 uint16 = 13028
	ER_IB_MSG_1204 uint16 = 13029
	ER_IB_MSG_1205 uint16 = 13030
	ER_IB_MSG_1206 uint16 = 13031
	ER_IB_MSG_1207 uint16 = 13032
	ER_IB_MSG_1208 uint16 = 13033
	ER_IB_MSG_1209 uint16 = 13034
	ER_IB_MSG_1210 uint16 = 13035
	ER_IB_MSG_1211 uint16 = 13036
	ER_IB_MSG_1212 uint16 = 13037
	ER_IB_MSG_1213 uint16 = 13038
	ER_IB_MSG_1214 uint16 = 13039
	ER_IB_MSG_1215 uint16 = 13040
	ER_IB_MSG_1216 uint16 = 13041
	ER_IB_MSG_1217 uint16 = 13042
	ER_IB_MSG_1218 uint16 = 13043
	ER_IB_MSG_1219 uint16 = 13044
	ER_IB_MSG_1220 uint16 = 13045
	ER_IB_MSG_1221 uint16 = 13046
	ER_IB_MSG_1222 uint16 = 13047
	ER_IB_MSG_1223 uint16 = 13048
	ER_IB_MSG_1224 uint16 = 13049
	ER_IB_MSG_1225 uint16 = 13050
	ER_IB_MSG_1226 uint16 = 13051
	ER_IB_MSG_1227 uint16 = 13052
	ER_IB_MSG_1228 uint16 = 13053
	ER_IB_MSG_1229 uint16 = 13054
	//OBSOLETE_ER_IB_MSG_1230 uint16 = 13055
	ER_IB_MSG_1231 uint16 = 13056
	ER_IB_MSG_1232 uint16 = 13057
	ER_IB_MSG_1233 uint16 = 13058
	ER_IB_MSG_1234 uint16 = 13059
	ER_IB_MSG_1235 uint16 = 13060
	ER_IB_MSG_1236 uint16 = 13061
	ER_IB_MSG_1237 uint16 = 13062
	ER_IB_MSG_1238 uint16 = 13063
	ER_IB_MSG_1239 uint16 = 13064
	ER_IB_MSG_1240 uint16 = 13065
	ER_IB_MSG_1241 uint16 = 13066
	ER_IB_MSG_1242 uint16 = 13067
	ER_IB_MSG_1243 uint16 = 13068
	ER_IB_MSG_1244 uint16 = 13069
	ER_IB_MSG_1245 uint16 = 13070
	ER_IB_MSG_1246 uint16 = 13071
	ER_IB_MSG_1247 uint16 = 13072
	ER_IB_MSG_1248 uint16 = 13073
	ER_IB_MSG_1249 uint16 = 13074
	ER_IB_MSG_1250 uint16 = 13075
	ER_IB_MSG_1251 uint16 = 13076
	ER_IB_MSG_1252 uint16 = 13077
	ER_IB_MSG_1253 uint16 = 13078
	//OBSOLETE_ER_IB_MSG_1254 uint16 = 13079
	ER_IB_MSG_1255 uint16 = 13080
	ER_IB_MSG_1256 uint16 = 13081
	ER_IB_MSG_1257 uint16 = 13082
	ER_IB_MSG_1258 uint16 = 13083
	ER_IB_MSG_1259 uint16 = 13084
	ER_IB_MSG_1260 uint16 = 13085
	ER_IB_MSG_1261 uint16 = 13086
	ER_IB_MSG_1262 uint16 = 13087
	ER_IB_MSG_1263 uint16 = 13088
	ER_IB_MSG_1264 uint16 = 13089
	ER_IB_MSG_1265 uint16 = 13090
	ER_IB_MSG_1266 uint16 = 13091
	ER_IB_MSG_1267 uint16 = 13092
	ER_IB_MSG_1268 uint16 = 13093
	ER_IB_MSG_1269 uint16 = 13094
	ER_IB_MSG_1270 uint16 = 13095
	ER_RPL_SLAVE_SQL_THREAD_STOP_CMD_EXEC_TIMEOUT uint16 = 13096
	ER_RPL_SLAVE_IO_THREAD_STOP_CMD_EXEC_TIMEOUT uint16 = 13097
	ER_RPL_GTID_UNSAFE_STMT_ON_NON_TRANS_TABLE uint16 = 13098
	ER_RPL_GTID_UNSAFE_STMT_CREATE_SELECT uint16 = 13099
	//OBSOLETE_ER_RPL_GTID_UNSAFE_STMT_ON_TEMPORARY_TABLE uint16 = 13100
	ER_BINLOG_ROW_VALUE_OPTION_IGNORED uint16 = 13101
	ER_BINLOG_USE_V1_ROW_EVENTS_IGNORED uint16 = 13102
	ER_BINLOG_ROW_VALUE_OPTION_USED_ONLY_FOR_AFTER_IMAGES uint16 = 13103
	ER_CONNECTION_ABORTED uint16 = 13104
	ER_NORMAL_SERVER_SHUTDOWN uint16 = 13105
	ER_KEYRING_MIGRATE_FAILED uint16 = 13106
	ER_GRP_RPL_LOWER_CASE_TABLE_NAMES_DIFF_FROM_GRP uint16 = 13107
	ER_OOM_SAVE_GTIDS uint16 = 13108
	ER_LCTN_NOT_FOUND uint16 = 13109
	//OBSOLETE_ER_REGEXP_INVALID_CAPTURE_GROUP_NAME uint16 = 13110
	ER_COMPONENT_FILTER_WRONG_VALUE uint16 = 13111
	ER_XPLUGIN_FAILED_TO_STOP_SERVICES uint16 = 13112
	ER_INCONSISTENT_ERROR uint16 = 13113
	ER_SERVER_MASTER_FATAL_ERROR_READING_BINLOG uint16 = 13114
	ER_NETWORK_READ_EVENT_CHECKSUM_FAILURE uint16 = 13115
	ER_SLAVE_CREATE_EVENT_FAILURE uint16 = 13116
	ER_SLAVE_FATAL_ERROR uint16 = 13117
	ER_SLAVE_HEARTBEAT_FAILURE uint16 = 13118
	ER_SLAVE_INCIDENT uint16 = 13119
	ER_SLAVE_MASTER_COM_FAILURE uint16 = 13120
	ER_SLAVE_RELAY_LOG_READ_FAILURE uint16 = 13121
	ER_SLAVE_RELAY_LOG_WRITE_FAILURE uint16 = 13122
	ER_SERVER_SLAVE_MI_INIT_REPOSITORY uint16 = 13123
	ER_SERVER_SLAVE_RLI_INIT_REPOSITORY uint16 = 13124
	ER_SERVER_NET_PACKET_TOO_LARGE uint16 = 13125
	ER_SERVER_NO_SYSTEM_TABLE_ACCESS uint16 = 13126
	ER_SERVER_UNKNOWN_ERROR uint16 = 13127
	ER_SERVER_UNKNOWN_SYSTEM_VARIABLE uint16 = 13128
	ER_SERVER_NO_SESSION_TO_SEND_TO uint16 = 13129
	ER_SERVER_NEW_ABORTING_CONNECTION uint16 = 13130
	ER_SERVER_OUT_OF_SORTMEMORY uint16 = 13131
	ER_SERVER_RECORD_FILE_FULL uint16 = 13132
	ER_SERVER_DISK_FULL_NOWAIT uint16 = 13133
	ER_SERVER_HANDLER_ERROR uint16 = 13134
	ER_SERVER_NOT_FORM_FILE uint16 = 13135
	ER_SERVER_CANT_OPEN_FILE uint16 = 13136
	ER_SERVER_FILE_NOT_FOUND uint16 = 13137
	ER_SERVER_FILE_USED uint16 = 13138
	ER_SERVER_CANNOT_LOAD_FROM_TABLE_V2 uint16 = 13139
	ER_ERROR_INFO_FROM_DA uint16 = 13140
	ER_SERVER_TABLE_CHECK_FAILED uint16 = 13141
	ER_SERVER_COL_COUNT_DOESNT_MATCH_PLEASE_UPDATE_V2 uint16 = 13142
	ER_SERVER_COL_COUNT_DOESNT_MATCH_CORRUPTED_V2 uint16 = 13143
	ER_SERVER_ACL_TABLE_ERROR uint16 = 13144
	ER_SERVER_SLAVE_INIT_QUERY_FAILED uint16 = 13145
	ER_SERVER_SLAVE_CONVERSION_FAILED uint16 = 13146
	ER_SERVER_SLAVE_IGNORED_TABLE uint16 = 13147
	ER_CANT_REPLICATE_ANONYMOUS_WITH_AUTO_POSITION uint16 = 13148
	ER_CANT_REPLICATE_ANONYMOUS_WITH_GTID_MODE_ON uint16 = 13149
	ER_CANT_REPLICATE_GTID_WITH_GTID_MODE_OFF uint16 = 13150
	ER_SERVER_TEST_MESSAGE uint16 = 13151
	ER_AUDIT_LOG_JSON_FILTER_PARSING_ERROR uint16 = 13152
	ER_AUDIT_LOG_JSON_FILTERING_NOT_ENABLED uint16 = 13153
	ER_PLUGIN_FAILED_TO_OPEN_TABLES uint16 = 13154
	ER_PLUGIN_FAILED_TO_OPEN_TABLE uint16 = 13155
	ER_AUDIT_LOG_JSON_FILTER_NAME_CANNOT_BE_EMPTY uint16 = 13156
	ER_AUDIT_LOG_USER_NAME_INVALID_CHARACTER uint16 = 13157
	ER_AUDIT_LOG_UDF_INSUFFICIENT_PRIVILEGE uint16 = 13158
	ER_AUDIT_LOG_NO_KEYRING_PLUGIN_INSTALLED uint16 = 13159
	ER_AUDIT_LOG_HOST_NAME_INVALID_CHARACTER uint16 = 13160
	ER_AUDIT_LOG_ENCRYPTION_PASSWORD_HAS_NOT_BEEN_SET uint16 = 13161
	ER_AUDIT_LOG_COULD_NOT_CREATE_AES_KEY uint16 = 13162
	ER_AUDIT_LOG_ENCRYPTION_PASSWORD_CANNOT_BE_FETCHED uint16 = 13163
	ER_COULD_NOT_REINITIALIZE_AUDIT_LOG_FILTERS uint16 = 13164
	ER_AUDIT_LOG_JSON_USER_NAME_CANNOT_BE_EMPTY uint16 = 13165
	ER_AUDIT_LOG_USER_FIRST_CHARACTER_MUST_BE_ALPHANUMERIC uint16 = 13166
	ER_AUDIT_LOG_JSON_FILTER_DOES_NOT_EXIST uint16 = 13167
	ER_IB_MSG_1271 uint16 = 13168
	ER_STARTING_INIT uint16 = 13169
	ER_ENDING_INIT uint16 = 13170
	ER_IB_MSG_1272 uint16 = 13171
	ER_SERVER_SHUTDOWN_INFO uint16 = 13172
	ER_GRP_RPL_PLUGIN_ABORT uint16 = 13173
	//OBSOLETE_ER_REGEXP_INVALID_FLAG uint16 = 13174
	//OBSOLETE_ER_XA_REPLICATION_FILTERS uint16 = 13175
	//OBSOLETE_ER_UPDATE_GTID_PURGED_WITH_GR uint16 = 13176
	ER_AUDIT_LOG_TABLE_DEFINITION_NOT_UPDATED uint16 = 13177
	ER_DD_INITIALIZE_SQL_ERROR uint16 = 13178
	ER_NO_PATH_FOR_SHARED_LIBRARY uint16 = 13179
	ER_UDF_ALREADY_EXISTS uint16 = 13180
	ER_SET_EVENT_FAILED uint16 = 13181
	ER_FAILED_TO_ALLOCATE_SSL_BIO uint16 = 13182
	ER_IB_MSG_1273 uint16 = 13183
	ER_PID_FILEPATH_LOCATIONS_INACCESSIBLE uint16 = 13184
	ER_UNKNOWN_VARIABLE_IN_PERSISTED_CONFIG_FILE uint16 = 13185
	ER_FAILED_TO_HANDLE_DEFAULTS_FILE uint16 = 13186
	ER_DUPLICATE_SYS_VAR uint16 = 13187
	ER_FAILED_TO_INIT_SYS_VAR uint16 = 13188
	ER_SYS_VAR_NOT_FOUND uint16 = 13189
	ER_IB_MSG_1274 uint16 = 13190
	ER_IB_MSG_1275 uint16 = 13191
	//OBSOLETE_ER_TARGET_TS_UNENCRYPTED uint16 = 13192
	ER_IB_MSG_WAIT_FOR_ENCRYPT_THREAD uint16 = 13193
	ER_IB_MSG_1277 uint16 = 13194
	ER_IB_MSG_NO_ENCRYPT_PROGRESS_FOUND uint16 = 13195
	ER_IB_MSG_RESUME_OP_FOR_SPACE uint16 = 13196
	ER_IB_MSG_1280 uint16 = 13197
	ER_IB_MSG_1281 uint16 = 13198
	ER_IB_MSG_1282 uint16 = 13199
	ER_IB_MSG_1283 uint16 = 13200
	ER_IB_MSG_1284 uint16 = 13201
	ER_CANT_SET_ERROR_SUPPRESSION_LIST_FROM_COMMAND_LINE uint16 = 13202
	ER_INVALID_VALUE_OF_BIND_ADDRESSES uint16 = 13203
	ER_RELAY_LOG_SPACE_LIMIT_DISABLED uint16 = 13204
	ER_GRP_RPL_ERROR_GTID_SET_EXTRACTION uint16 = 13205
	ER_GRP_RPL_MISSING_GRP_RPL_ACTION_COORDINATOR uint16 = 13206
	ER_GRP_RPL_JOIN_WHEN_GROUP_ACTION_RUNNING uint16 = 13207
	ER_GRP_RPL_JOINER_EXIT_WHEN_GROUP_ACTION_RUNNING uint16 = 13208
	ER_GRP_RPL_CHANNEL_THREAD_WHEN_GROUP_ACTION_RUNNING uint16 = 13209
	ER_GRP_RPL_APPOINTED_PRIMARY_NOT_PRESENT uint16 = 13210
	ER_GRP_RPL_ERROR_ON_MESSAGE_SENDING uint16 = 13211
	ER_GRP_RPL_CONFIGURATION_ACTION_ERROR uint16 = 13212
	ER_GRP_RPL_CONFIGURATION_ACTION_LOCAL_TERMINATION uint16 = 13213
	ER_GRP_RPL_CONFIGURATION_ACTION_START uint16 = 13214
	ER_GRP_RPL_CONFIGURATION_ACTION_END uint16 = 13215
	ER_GRP_RPL_CONFIGURATION_ACTION_KILLED_ERROR uint16 = 13216
	ER_GRP_RPL_PRIMARY_ELECTION_PROCESS_ERROR uint16 = 13217
	ER_GRP_RPL_PRIMARY_ELECTION_STOP_ERROR uint16 = 13218
	ER_GRP_RPL_NO_STAGE_SERVICE uint16 = 13219
	ER_GRP_RPL_UDF_REGISTER_ERROR uint16 = 13220
	ER_GRP_RPL_UDF_UNREGISTER_ERROR uint16 = 13221
	ER_GRP_RPL_UDF_REGISTER_SERVICE_ERROR uint16 = 13222
	ER_GRP_RPL_SERVER_UDF_ERROR uint16 = 13223
	//OBSOLETE_ER_CURRENT_PASSWORD_NOT_REQUIRED uint16 = 13224
	//OBSOLETE_ER_INCORRECT_CURRENT_PASSWORD uint16 = 13225
	//OBSOLETE_ER_MISSING_CURRENT_PASSWORD uint16 = 13226
	ER_SERVER_WRONG_VALUE_FOR_VAR uint16 = 13227
	ER_COULD_NOT_CREATE_WINDOWS_REGISTRY_KEY uint16 = 13228
	ER_SERVER_GTID_UNSAFE_CREATE_DROP_TEMP_TABLE_IN_TRX_IN_SBR uint16 = 13229
	//OBSOLETE_ER_SECONDARY_ENGINE uint16 = 13230
	//OBSOLETE_ER_SECONDARY_ENGINE_DDL uint16 = 13231
	//OBSOLETE_ER_NO_SESSION_TEMP uint16 = 13232
	ER_XPLUGIN_FAILED_TO_SWITCH_SECURITY_CTX uint16 = 13233
	ER_RPL_GTID_UNSAFE_ALTER_ADD_COL_WITH_DEFAULT_EXPRESSION uint16 = 13234
	ER_UPGRADE_PARSE_ERROR uint16 = 13235
	ER_DATA_DIRECTORY_UNUSABLE uint16 = 13236
	ER_LDAP_AUTH_USER_GROUP_SEARCH_ROOT_BIND uint16 = 13237
	ER_PLUGIN_INSTALL_ERROR uint16 = 13238
	ER_PLUGIN_UNINSTALL_ERROR uint16 = 13239
	ER_SHARED_TABLESPACE_USED_BY_PARTITIONED_TABLE uint16 = 13240
	ER_UNKNOWN_TABLESPACE_TYPE uint16 = 13241
	ER_WARN_DEPRECATED_UTF8_ALIAS_OPTION uint16 = 13242
	ER_WARN_DEPRECATED_UTF8MB3_CHARSET_OPTION uint16 = 13243
	ER_WARN_DEPRECATED_UTF8MB3_COLLATION_OPTION uint16 = 13244
	ER_SSL_MEMORY_INSTRUMENTATION_INIT_FAILED uint16 = 13245
	ER_IB_MSG_MADV_DONTDUMP_UNSUPPORTED uint16 = 13246
	ER_IB_MSG_MADVISE_FAILED uint16 = 13247
	//OBSOLETE_ER_COLUMN_CHANGE_SIZE uint16 = 13248
	ER_WARN_REMOVED_SQL_MODE uint16 = 13249
	ER_IB_MSG_FAILED_TO_ALLOCATE_WAIT uint16 = 13250
	ER_IB_MSG_NUM_POOLS uint16 = 13251
	ER_IB_MSG_USING_UNDO_SPACE uint16 = 13252
	ER_IB_MSG_FAIL_TO_SAVE_SPACE_STATE uint16 = 13253
	ER_IB_MSG_MAX_UNDO_SPACES_REACHED uint16 = 13254
	ER_IB_MSG_ERROR_OPENING_NEW_UNDO_SPACE uint16 = 13255
	ER_IB_MSG_FAILED_SDI_Z_BUF_ERROR uint16 = 13256
	ER_IB_MSG_FAILED_SDI_Z_MEM_ERROR uint16 = 13257
	ER_IB_MSG_SDI_Z_STREAM_ERROR uint16 = 13258
	ER_IB_MSG_SDI_Z_UNKNOWN_ERROR uint16 = 13259
	ER_IB_MSG_FOUND_WRONG_UNDO_SPACE uint16 = 13260
	ER_IB_MSG_NOT_END_WITH_IBU uint16 = 13261
	//OBSOLETE_ER_IB_MSG_UNDO_TRUNCATE_EMPTY_FILE uint16 = 13262
	//OBSOLETE_ER_IB_MSG_UNDO_INJECT_BEFORE_DD_UPDATE uint16 = 13263
	//OBSOLETE_ER_IB_MSG_UNDO_INJECT_BEFORE_UNDO_LOGGING uint16 = 13264
	//OBSOLETE_ER_IB_MSG_UNDO_INJECT_BEFORE_RSEG uint16 = 13265
	ER_IB_MSG_FAILED_TO_FINISH_TRUNCATE uint16 = 13266
	ER_IB_MSG_DEPRECATED_INNODB_UNDO_TABLESPACES uint16 = 13267
	ER_IB_MSG_WRONG_TABLESPACE_DIR uint16 = 13268
	ER_IB_MSG_LOCK_FREE_HASH_USAGE_STATS uint16 = 13269
	ER_CLONE_DONOR_TRACE uint16 = 13270
	ER_CLONE_PROTOCOL_TRACE uint16 = 13271
	ER_CLONE_CLIENT_TRACE uint16 = 13272
	ER_CLONE_SERVER_TRACE uint16 = 13273
	ER_THREAD_POOL_PFS_TABLES_INIT_FAILED uint16 = 13274
	ER_THREAD_POOL_PFS_TABLES_ADD_FAILED uint16 = 13275
	ER_CANT_SET_DATA_DIR uint16 = 13276
	ER_INNODB_INVALID_INNODB_UNDO_DIRECTORY_LOCATION uint16 = 13277
	ER_SERVER_RPL_ENCRYPTION_FAILED_TO_FETCH_KEY uint16 = 13278
	ER_SERVER_RPL_ENCRYPTION_KEY_NOT_FOUND uint16 = 13279
	ER_SERVER_RPL_ENCRYPTION_KEYRING_INVALID_KEY uint16 = 13280
	ER_SERVER_RPL_ENCRYPTION_HEADER_ERROR uint16 = 13281
	ER_SERVER_RPL_ENCRYPTION_FAILED_TO_ROTATE_LOGS uint16 = 13282
	ER_SERVER_RPL_ENCRYPTION_KEY_EXISTS_UNEXPECTED uint16 = 13283
	ER_SERVER_RPL_ENCRYPTION_FAILED_TO_GENERATE_KEY uint16 = 13284
	ER_SERVER_RPL_ENCRYPTION_FAILED_TO_STORE_KEY uint16 = 13285
	ER_SERVER_RPL_ENCRYPTION_FAILED_TO_REMOVE_KEY uint16 = 13286
	ER_SERVER_RPL_ENCRYPTION_MASTER_KEY_RECOVERY_FAILED uint16 = 13287
	ER_SERVER_RPL_ENCRYPTION_UNABLE_TO_INITIALIZE uint16 = 13288
	ER_SERVER_RPL_ENCRYPTION_UNABLE_TO_ROTATE_MASTER_KEY_AT_STARTUP uint16 = 13289
	ER_SERVER_RPL_ENCRYPTION_IGNORE_ROTATE_MASTER_KEY_AT_STARTUP uint16 = 13290
	ER_INVALID_ADMIN_ADDRESS uint16 = 13291
	ER_SERVER_STARTUP_ADMIN_INTERFACE uint16 = 13292
	ER_CANT_CREATE_ADMIN_THREAD uint16 = 13293
	ER_WARNING_RETAIN_CURRENT_PASSWORD_CLAUSE_VOID uint16 = 13294
	ER_WARNING_DISCARD_OLD_PASSWORD_CLAUSE_VOID uint16 = 13295
	//OBSOLETE_ER_SECOND_PASSWORD_CANNOT_BE_EMPTY uint16 = 13296
	//OBSOLETE_ER_PASSWORD_CANNOT_BE_RETAINED_ON_PLUGIN_CHANGE uint16 = 13297
	//OBSOLETE_ER_CURRENT_PASSWORD_CANNOT_BE_RETAINED uint16 = 13298
	ER_WARNING_AUTHCACHE_INVALID_USER_ATTRIBUTES uint16 = 13299
	ER_MYSQL_NATIVE_PASSWORD_SECOND_PASSWORD_USED_INFORMATION uint16 = 13300
	ER_SHA256_PASSWORD_SECOND_PASSWORD_USED_INFORMATION uint16 = 13301
	ER_CACHING_SHA2_PASSWORD_SECOND_PASSWORD_USED_INFORMATION uint16 = 13302
	ER_GRP_RPL_SEND_TRX_PREPARED_MESSAGE_FAILED uint16 = 13303
	ER_GRP_RPL_RELEASE_COMMIT_AFTER_GROUP_PREPARE_FAILED uint16 = 13304
	ER_GRP_RPL_TRX_ALREADY_EXISTS_ON_TCM_ON_AFTER_CERTIFICATION uint16 = 13305
	ER_GRP_RPL_FAILED_TO_INSERT_TRX_ON_TCM_ON_AFTER_CERTIFICATION uint16 = 13306
	ER_GRP_RPL_REGISTER_TRX_TO_WAIT_FOR_GROUP_PREPARE_FAILED uint16 = 13307
	ER_GRP_RPL_TRX_WAIT_FOR_GROUP_PREPARE_FAILED uint16 = 13308
	ER_GRP_RPL_TRX_DOES_NOT_EXIST_ON_TCM_ON_HANDLE_REMOTE_PREPARE uint16 = 13309
	ER_GRP_RPL_RELEASE_BEGIN_TRX_AFTER_DEPENDENCIES_COMMIT_FAILED uint16 = 13310
	ER_GRP_RPL_REGISTER_TRX_TO_WAIT_FOR_DEPENDENCIES_FAILED uint16 = 13311
	ER_GRP_RPL_WAIT_FOR_DEPENDENCIES_FAILED uint16 = 13312
	ER_GRP_RPL_REGISTER_TRX_TO_WAIT_FOR_SYNC_BEFORE_EXECUTION_FAILED uint16 = 13313
	ER_GRP_RPL_SEND_TRX_SYNC_BEFORE_EXECUTION_FAILED uint16 = 13314
	ER_GRP_RPL_TRX_WAIT_FOR_SYNC_BEFORE_EXECUTION_FAILED uint16 = 13315
	ER_GRP_RPL_RELEASE_BEGIN_TRX_AFTER_WAIT_FOR_SYNC_BEFORE_EXEC uint16 = 13316
	ER_GRP_RPL_TRX_WAIT_FOR_GROUP_GTID_EXECUTED uint16 = 13317
	//OBSOLETE_ER_UNIT_NOT_FOUND uint16 = 13318
	//OBSOLETE_ER_GEOMETRY_IN_UNKNOWN_LENGTH_UNIT uint16 = 13319
	ER_WARN_PROPERTY_STRING_PARSE_FAILED uint16 = 13320
	ER_INVALID_PROPERTY_KEY uint16 = 13321
	ER_GRP_RPL_GTID_SET_EXTRACT_ERROR_DURING_RECOVERY uint16 = 13322
	ER_SERVER_RPL_ENCRYPTION_FAILED_TO_ENCRYPT uint16 = 13323
	ER_CANNOT_GET_SERVER_VERSION_FROM_TABLESPACE_HEADER uint16 = 13324
	ER_CANNOT_SET_SERVER_VERSION_IN_TABLESPACE_HEADER uint16 = 13325
	ER_SERVER_UPGRADE_VERSION_NOT_SUPPORTED uint16 = 13326
	ER_SERVER_UPGRADE_FROM_VERSION uint16 = 13327
	ER_GRP_RPL_ERROR_ON_CERT_DB_INSTALL uint16 = 13328
	ER_GRP_RPL_FORCE_MEMBERS_WHEN_LEAVING uint16 = 13329
	ER_TRG_WRONG_ORDER uint16 = 13330
	//OBSOLETE_ER_SECONDARY_ENGINE_PLUGIN uint16 = 13331
	ER_LDAP_AUTH_GRP_SEARCH_NOT_SPECIAL_HDL uint16 = 13332
	ER_LDAP_AUTH_GRP_USER_OBJECT_HAS_GROUP_INFO uint16 = 13333
	ER_LDAP_AUTH_GRP_INFO_FOUND_IN_MANY_OBJECTS uint16 = 13334
	ER_LDAP_AUTH_GRP_INCORRECT_ATTRIBUTE uint16 = 13335
	ER_LDAP_AUTH_GRP_NULL_ATTRIBUTE_VALUE uint16 = 13336
	ER_LDAP_AUTH_GRP_DN_PARSING_FAILED uint16 = 13337
	ER_LDAP_AUTH_GRP_OBJECT_HAS_USER_INFO uint16 = 13338
	ER_LDAP_AUTH_LDAPS uint16 = 13339
	ER_LDAP_MAPPING_GET_USER_PROXY uint16 = 13340
	ER_LDAP_MAPPING_USER_DONT_BELONG_GROUP uint16 = 13341
	ER_LDAP_MAPPING_INFO uint16 = 13342
	ER_LDAP_MAPPING_EMPTY_MAPPING uint16 = 13343
	ER_LDAP_MAPPING_PROCESS_MAPPING uint16 = 13344
	ER_LDAP_MAPPING_CHECK_DELIMI_QUOTE uint16 = 13345
	ER_LDAP_MAPPING_PROCESS_DELIMITER uint16 = 13346
	ER_LDAP_MAPPING_PROCESS_DELIMITER_EQUAL_NOT_FOUND uint16 = 13347
	ER_LDAP_MAPPING_PROCESS_DELIMITER_TRY_COMMA uint16 = 13348
	ER_LDAP_MAPPING_PROCESS_DELIMITER_COMMA_NOT_FOUND uint16 = 13349
	ER_LDAP_MAPPING_NO_SEPEARATOR_END_OF_GROUP uint16 = 13350
	ER_LDAP_MAPPING_GETTING_NEXT_MAPPING uint16 = 13351
	ER_LDAP_MAPPING_PARSING_CURRENT_STATE uint16 = 13352
	ER_LDAP_MAPPING_PARSING_MAPPING_INFO uint16 = 13353
	ER_LDAP_MAPPING_PARSING_ERROR uint16 = 13354
	ER_LDAP_MAPPING_TRIMMING_SPACES uint16 = 13355
	ER_LDAP_MAPPING_IS_QUOTE uint16 = 13356
	ER_LDAP_MAPPING_NON_DESIRED_STATE uint16 = 13357
	ER_INVALID_NAMED_PIPE_FULL_ACCESS_GROUP uint16 = 13358
	ER_PREPARE_FOR_SECONDARY_ENGINE uint16 = 13359
	ER_SERVER_WARN_DEPRECATED uint16 = 13360
	ER_AUTH_ID_WITH_SYSTEM_USER_PRIV_IN_MANDATORY_ROLES uint16 = 13361
	ER_SERVER_BINLOG_MASTER_KEY_RECOVERY_OUT_OF_COMBINATION uint16 = 13362
	ER_SERVER_BINLOG_MASTER_KEY_ROTATION_FAIL_TO_CLEANUP_AUX_KEY uint16 = 13363
	//OBSOLETE_ER_CANNOT_GRANT_SYSTEM_PRIV_TO_MANDATORY_ROLE uint16 = 13364
	//OBSOLETE_ER_PARTIAL_REVOKE_AND_DB_GRANT_BOTH_EXISTS uint16 = 13365
	//OBSOLETE_ER_DB_ACCESS_DENIED uint16 = 13366
	//OBSOLETE_ER_PARTIAL_REVOKES_EXIST uint16 = 13367
	ER_TURNING_ON_PARTIAL_REVOKES uint16 = 13368
	ER_WARN_PARTIAL_REVOKE_AND_DB_GRANT uint16 = 13369
	ER_WARN_INCORRECT_PRIVILEGE_FOR_DB_RESTRICTIONS uint16 = 13370
	ER_WARN_INVALID_DB_RESTRICTIONS uint16 = 13371
	ER_GRP_RPL_INVALID_COMMUNICATION_PROTOCOL uint16 = 13372
	ER_GRP_RPL_STARTED_AUTO_REJOIN uint16 = 13373
	ER_GRP_RPL_TIMEOUT_RECEIVED_VC_ON_REJOIN uint16 = 13374
	ER_GRP_RPL_FINISHED_AUTO_REJOIN uint16 = 13375
	ER_GRP_RPL_DEFAULT_TABLE_ENCRYPTION_DIFF_FROM_GRP uint16 = 13376
	ER_SERVER_UPGRADE_OFF uint16 = 13377
	ER_SERVER_UPGRADE_SKIP uint16 = 13378
	ER_SERVER_UPGRADE_PENDING uint16 = 13379
	ER_SERVER_UPGRADE_FAILED uint16 = 13380
	ER_SERVER_UPGRADE_STATUS uint16 = 13381
	ER_SERVER_UPGRADE_REPAIR_REQUIRED uint16 = 13382
	ER_SERVER_UPGRADE_REPAIR_STATUS uint16 = 13383
	ER_SERVER_UPGRADE_INFO_FILE uint16 = 13384
	ER_SERVER_UPGRADE_SYS_SCHEMA uint16 = 13385
	ER_SERVER_UPGRADE_MYSQL_TABLES uint16 = 13386
	ER_SERVER_UPGRADE_SYSTEM_TABLES uint16 = 13387
	ER_SERVER_UPGRADE_EMPTY_SYS uint16 = 13388
	ER_SERVER_UPGRADE_NO_SYS_VERSION uint16 = 13389
	ER_SERVER_UPGRADE_SYS_VERSION_EMPTY uint16 = 13390
	ER_SERVER_UPGRADE_SYS_SCHEMA_OUTDATED uint16 = 13391
	ER_SERVER_UPGRADE_SYS_SCHEMA_UP_TO_DATE uint16 = 13392
	ER_SERVER_UPGRADE_SYS_SCHEMA_OBJECT_COUNT uint16 = 13393
	ER_SERVER_UPGRADE_CHECKING_DB uint16 = 13394
	ER_IB_MSG_DDL_LOG_DELETE_BY_ID_TMCT uint16 = 13395
	ER_IB_MSG_POST_RECOVER_DDL_LOG_RECOVER uint16 = 13396
	ER_IB_MSG_POST_RECOVER_POST_TS_ENCRYPT uint16 = 13397
	ER_IB_MSG_DDL_LOG_FAIL_POST_DDL uint16 = 13398
	ER_SERVER_BINLOG_UNSAFE_SYSTEM_FUNCTION uint16 = 13399
	ER_SERVER_UPGRADE_HELP_TABLE_STATUS uint16 = 13400
	ER_GRP_RPL_SRV_GTID_WAIT_ERROR uint16 = 13401
	ER_GRP_DELAYED_VCLE_LOGGING uint16 = 13402
	//OBSOLETE_ER_CANNOT_GRANT_ROLES_TO_ANONYMOUS_USER uint16 = 13403
	ER_BINLOG_UNABLE_TO_ROTATE_GTID_TABLE_READONLY uint16 = 13404
	ER_NETWORK_NAMESPACES_NOT_SUPPORTED uint16 = 13405
	ER_UNKNOWN_NETWORK_NAMESPACE uint16 = 13406
	ER_NETWORK_NAMESPACE_NOT_ALLOWED_FOR_WILDCARD_ADDRESS uint16 = 13407
	ER_SETNS_FAILED uint16 = 13408
	ER_WILDCARD_NOT_ALLOWED_FOR_MULTIADDRESS_BIND uint16 = 13409
	ER_NETWORK_NAMESPACE_FILE_PATH_TOO_LONG uint16 = 13410
	ER_IB_MSG_TOO_LONG_PATH uint16 = 13411
	ER_IB_RECV_FIRST_REC_GROUP_INVALID uint16 = 13412
	ER_DD_UPGRADE_COMPLETED uint16 = 13413
	ER_SSL_SERVER_CERT_VERIFY_FAILED uint16 = 13414
	ER_PERSIST_OPTION_USER_TRUNCATED uint16 = 13415
	ER_PERSIST_OPTION_HOST_TRUNCATED uint16 = 13416
	ER_NET_WAIT_ERROR uint16 = 13417
	ER_IB_MSG_1285 uint16 = 13418
	ER_IB_MSG_CLOCK_MONOTONIC_UNSUPPORTED uint16 = 13419
	ER_IB_MSG_CLOCK_GETTIME_FAILED uint16 = 13420
	ER_PLUGIN_NOT_EARLY_DUP uint16 = 13421
	ER_PLUGIN_NO_INSTALL_DUP uint16 = 13422
	//OBSOLETE_ER_WARN_DEPRECATED_SQL_CALC_FOUND_ROWS uint16 = 13423
	//OBSOLETE_ER_WARN_DEPRECATED_FOUND_ROWS uint16 = 13424
	ER_BINLOG_UNSAFE_DEFAULT_EXPRESSION_IN_SUBSTATEMENT uint16 = 13425
	ER_GRP_RPL_MEMBER_VER_READ_COMPATIBLE uint16 = 13426
	ER_LOCK_ORDER_INIT_FAILED uint16 = 13427
	ER_AUDIT_LOG_KEYRING_ID_TIMESTAMP_VALUE_IS_INVALID uint16 = 13428
	ER_AUDIT_LOG_FILE_NAME_TIMESTAMP_VALUE_IS_MISSING_OR_INVALID uint16 = 13429
	ER_AUDIT_LOG_FILE_NAME_DOES_NOT_HAVE_REQUIRED_FORMAT uint16 = 13430
	ER_AUDIT_LOG_FILE_NAME_KEYRING_ID_VALUE_IS_MISSING uint16 = 13431
	ER_AUDIT_LOG_FILE_HAS_BEEN_SUCCESSFULLY_PROCESSED uint16 = 13432
	ER_AUDIT_LOG_COULD_NOT_OPEN_FILE_FOR_READING uint16 = 13433
	ER_AUDIT_LOG_INVALID_FILE_CONTENT uint16 = 13434
	ER_AUDIT_LOG_CANNOT_READ_PASSWORD uint16 = 13435
	ER_AUDIT_LOG_CANNOT_STORE_PASSWORD uint16 = 13436
	ER_AUDIT_LOG_CANNOT_REMOVE_PASSWORD uint16 = 13437
	ER_AUDIT_LOG_PASSWORD_HAS_BEEN_COPIED uint16 = 13438
	//OBSOLETE_ER_AUDIT_LOG_INSUFFICIENT_PRIVILEGE uint16 = 13439
	//OBSOLETE_ER_WRONG_MVI_VALUE uint16 = 13440
	//OBSOLETE_ER_WARN_FUNC_INDEX_NOT_APPLICABLE uint16 = 13441
	//OBSOLETE_ER_EXCEEDED_MV_KEYS_NUM uint16 = 13442
	//OBSOLETE_ER_EXCEEDED_MV_KEYS_SPACE uint16 = 13443
	//OBSOLETE_ER_FUNCTIONAL_INDEX_DATA_IS_TOO_LONG uint16 = 13444
	//OBSOLETE_ER_INVALID_JSON_VALUE_FOR_FUNC_INDEX uint16 = 13445
	//OBSOLETE_ER_JSON_VALUE_OUT_OF_RANGE_FOR_FUNC_INDEX uint16 = 13446
	ER_LDAP_EMPTY_USERDN_PASSWORD uint16 = 13447
	//OBSOLETE_ER_GROUPING_ON_TIMESTAMP_IN_DST uint16 = 13448
	ER_ACL_WRONG_OR_MISSING_ACL_TABLES_LOG uint16 = 13449
	ER_LOCK_ORDER_FAILED_WRITE_FILE uint16 = 13450
	ER_LOCK_ORDER_FAILED_READ_FILE uint16 = 13451
	ER_LOCK_ORDER_MESSAGE uint16 = 13452
	ER_LOCK_ORDER_DEPENDENCIES_SYNTAX uint16 = 13453
	ER_LOCK_ORDER_SCANNER_SYNTAX uint16 = 13454
	ER_DATA_DIRECTORY_UNUSABLE_DELETABLE uint16 = 13455
	ER_IB_MSG_BTREE_LEVEL_LIMIT_EXCEEDED uint16 = 13456
	ER_IB_CLONE_START_STOP uint16 = 13457
	ER_IB_CLONE_OPERATION uint16 = 13458
	ER_IB_CLONE_RESTART uint16 = 13459
	ER_IB_CLONE_USER_DATA uint16 = 13460
	ER_IB_CLONE_NON_INNODB_TABLE uint16 = 13461
	ER_CLONE_SHUTDOWN_TRACE uint16 = 13462
	ER_GRP_RPL_GTID_PURGED_EXTRACT_ERROR uint16 = 13463
	ER_GRP_RPL_CLONE_PROCESS_PREPARE_ERROR uint16 = 13464
	ER_GRP_RPL_CLONE_PROCESS_EXEC_ERROR uint16 = 13465
	ER_GRP_RPL_RECOVERY_EVAL_ERROR uint16 = 13466
	ER_GRP_RPL_NO_POSSIBLE_RECOVERY uint16 = 13467
	ER_GRP_RPL_CANT_KILL_THREAD uint16 = 13468
	ER_GRP_RPL_RECOVERY_STRAT_CLONE_THRESHOLD uint16 = 13469
	ER_GRP_RPL_RECOVERY_STRAT_CLONE_PURGED uint16 = 13470
	ER_GRP_RPL_RECOVERY_STRAT_CHOICE uint16 = 13471
	ER_GRP_RPL_RECOVERY_STRAT_FALLBACK uint16 = 13472
	ER_GRP_RPL_RECOVERY_STRAT_NO_FALLBACK uint16 = 13473
	ER_GRP_RPL_SLAVE_THREAD_ERROR_ON_CLONE uint16 = 13474
	ER_UNKNOWN_TABLE_IN_UPGRADE uint16 = 13475
	ER_IDENT_CAUSES_TOO_LONG_PATH_IN_UPGRADE uint16 = 13476
	ER_XA_CANT_CREATE_MDL_BACKUP uint16 = 13477
	ER_AUDIT_LOG_SUPER_PRIVILEGE_REQUIRED uint16 = 13478
	ER_AUDIT_LOG_UDF_INVALID_ARGUMENT_TYPE uint16 = 13479
	ER_AUDIT_LOG_UDF_INVALID_ARGUMENT_COUNT uint16 = 13480
	ER_AUDIT_LOG_HAS_NOT_BEEN_INSTALLED uint16 = 13481
	ER_AUDIT_LOG_UDF_READ_INVALID_MAX_ARRAY_LENGTH_ARG_TYPE uint16 = 13482
	ER_LOG_CANNOT_WRITE_EXTENDED uint16 = 13483
	//OBSOLETE_ER_UPGRADE_WITH_PARTITIONED_TABLES_REJECTED uint16 = 13484
	ER_KEYRING_AWS_INCORRECT_PROXY uint16 = 13485
	ER_GRP_RPL_SERVER_SET_TO_OFFLINE_MODE_DUE_TO_ERRORS uint16 = 13486
	ER_GRP_RPL_MESSAGE_SERVICE_FATAL_ERROR uint16 = 13487
	ER_WARN_WRONG_COMPRESSION_ALGORITHM_LOG uint16 = 13488
	ER_WARN_WRONG_COMPRESSION_LEVEL_LOG uint16 = 13489
	ER_PROTOCOL_COMPRESSION_RESET_LOG uint16 = 13490
	ER_XPLUGIN_COMPRESSION_ERROR uint16 = 13491
	ER_MYSQLBACKUP_MSG uint16 = 13492
	ER_WARN_UNKNOWN_KEYRING_AWS_REGION uint16 = 13493
	ER_WARN_LOG_PRIVILEGE_CHECKS_USER_DOES_NOT_EXIST uint16 = 13494
	ER_WARN_LOG_PRIVILEGE_CHECKS_USER_CORRUPT uint16 = 13495
	ER_WARN_LOG_PRIVILEGE_CHECKS_USER_NEEDS_RPL_APPLIER_PRIV uint16 = 13496
	ER_FILE_PRIVILEGE_FOR_REPLICATION_CHECKS uint16 = 13497
	ER_RPL_SLAVE_SQL_THREAD_STARTING_WITH_PRIVILEGE_CHECKS uint16 = 13498
	ER_AUDIT_LOG_CANNOT_GENERATE_PASSWORD uint16 = 13499
	ER_INIT_FAILED_TO_GENERATE_ROOT_PASSWORD uint16 = 13500
	ER_PLUGIN_LOAD_OPTIONS_IGNORED uint16 = 13501
	ER_WARN_AUTH_ID_WITH_SYSTEM_USER_PRIV_IN_MANDATORY_ROLES uint16 = 13502
	ER_IB_MSG_SKIP_HIDDEN_DIR uint16 = 13503
	ER_WARN_RPL_RECOVERY_NO_ROTATE_EVENT_FROM_MASTER_EOF uint16 = 13504
	ER_IB_LOB_ROLLBACK_INDEX_LEN uint16 = 13505
	ER_CANT_PROCESS_EXPRESSION_FOR_GENERATED_COLUMN_TO_DD uint16 = 13506
	ER_RPL_SLAVE_QUEUE_EVENT_FAILED_INVALID_NON_ROW_FORMAT uint16 = 13507
	ER_RPL_SLAVE_APPLY_LOG_EVENT_FAILED_INVALID_NON_ROW_FORMAT uint16 = 13508
	ER_LOG_PRIV_CHECKS_REQUIRE_ROW_FORMAT_NOT_SET uint16 = 13509
	ER_RPL_SLAVE_SQL_THREAD_DETECTED_UNEXPECTED_EVENT_SEQUENCE uint16 = 13510
	ER_IB_MSG_UPGRADE_PARTITION_FILE uint16 = 13511
	ER_IB_MSG_DOWNGRADE_PARTITION_FILE uint16 = 13512
	ER_IB_MSG_UPGRADE_PARTITION_FILE_IMPORT uint16 = 13513
	ER_IB_WARN_OPEN_PARTITION_FILE uint16 = 13514
	ER_IB_MSG_FIL_STATE_MOVED_CORRECTED uint16 = 13515
	ER_IB_MSG_FIL_STATE_MOVED_CHANGED_PATH uint16 = 13516
	ER_IB_MSG_FIL_STATE_MOVED_CHANGED_NAME uint16 = 13517
	ER_IB_MSG_FIL_STATE_MOVED_TOO_MANY uint16 = 13518
	ER_GR_ELECTED_PRIMARY_GTID_INFORMATION uint16 = 13519
	ER_SCHEMA_NAME_IN_UPPER_CASE_NOT_ALLOWED uint16 = 13520
	ER_TABLE_NAME_IN_UPPER_CASE_NOT_ALLOWED uint16 = 13521
	ER_SCHEMA_NAME_IN_UPPER_CASE_NOT_ALLOWED_FOR_FK uint16 = 13522
	ER_TABLE_NAME_IN_UPPER_CASE_NOT_ALLOWED_FOR_FK uint16 = 13523
	ER_IB_MSG_DICT_PARTITION_NOT_FOUND uint16 = 13524
	ER_ACCESS_DENIED_FOR_USER_ACCOUNT_BLOCKED_BY_PASSWORD_LOCK uint16 = 13525
	ER_INNODB_OUT_OF_RESOURCES uint16 = 13526
	ER_DD_UPGRADE_FOUND_PREPARED_XA_TRANSACTION uint16 = 13527
	ER_MIGRATE_TABLE_TO_DD_OOM uint16 = 13528
	ER_RPL_RELAY_LOG_RECOVERY_INFO_AFTER_CLONE uint16 = 13529
	ER_IB_MSG_57_UNDO_SPACE_DELETE_FAIL uint16 = 13530
	ER_IB_MSG_DBLWR_1285 uint16 = 13531
	ER_IB_MSG_DBLWR_1286 uint16 = 13532
	ER_IB_MSG_DBLWR_1287 uint16 = 13533
	ER_IB_MSG_DBLWR_1288 uint16 = 13534
	ER_IB_MSG_DBLWR_1290 uint16 = 13535
	ER_IB_MSG_BAD_DBLWR_FILE_NAME uint16 = 13536
	//OBSOLETE_ER_IB_MSG_DBLWR_1292 uint16 = 13537
	ER_IB_MSG_DBLWR_1293 uint16 = 13538
	ER_IB_MSG_DBLWR_1294 uint16 = 13539
	ER_IB_MSG_DBLWR_1295 uint16 = 13540
	ER_IB_MSG_DBLWR_1296 uint16 = 13541
	ER_IB_MSG_DBLWR_1297 uint16 = 13542
	ER_IB_MSG_DBLWR_1298 uint16 = 13543
	ER_IB_MSG_DBLWR_1300 uint16 = 13544
	ER_IB_MSG_DBLWR_1301 uint16 = 13545
	ER_IB_MSG_DBLWR_1304 uint16 = 13546
	ER_IB_MSG_DBLWR_1305 uint16 = 13547
	ER_IB_MSG_DBLWR_1306 uint16 = 13548
	ER_IB_MSG_DBLWR_1307 uint16 = 13549
	ER_IB_MSG_DBLWR_1308 uint16 = 13550
	ER_IB_MSG_DBLWR_1309 uint16 = 13551
	ER_IB_MSG_DBLWR_1310 uint16 = 13552
	ER_IB_MSG_DBLWR_1311 uint16 = 13553
	ER_IB_MSG_DBLWR_1312 uint16 = 13554
	ER_IB_MSG_DBLWR_1313 uint16 = 13555
	ER_IB_MSG_DBLWR_1314 uint16 = 13556
	ER_IB_MSG_DBLWR_1315 uint16 = 13557
	ER_IB_MSG_DBLWR_1316 uint16 = 13558
	ER_IB_MSG_DBLWR_1317 uint16 = 13559
	ER_IB_MSG_DBLWR_1318 uint16 = 13560
	ER_IB_MSG_DBLWR_1319 uint16 = 13561
	ER_IB_MSG_DBLWR_1320 uint16 = 13562
	ER_IB_MSG_DBLWR_1321 uint16 = 13563
	ER_IB_MSG_DBLWR_1322 uint16 = 13564
	ER_IB_MSG_DBLWR_1323 uint16 = 13565
	ER_IB_MSG_DBLWR_1324 uint16 = 13566
	ER_IB_MSG_DBLWR_1325 uint16 = 13567
	ER_IB_MSG_DBLWR_1326 uint16 = 13568
	ER_IB_MSG_DBLWR_1327 uint16 = 13569
	ER_IB_MSG_GTID_FLUSH_AT_SHUTDOWN uint16 = 13570
	ER_IB_MSG_57_STAT_SPACE_DELETE_FAIL uint16 = 13571
	ER_NDBINFO_UPGRADING_SCHEMA uint16 = 13572
	ER_NDBINFO_NOT_UPGRADING_SCHEMA uint16 = 13573
	ER_NDBINFO_UPGRADING_SCHEMA_FAIL uint16 = 13574
	ER_IB_MSG_CREATE_LOG_FILE uint16 = 13575
	ER_IB_MSG_INNODB_START_INITIALIZE uint16 = 13576
	ER_IB_MSG_INNODB_END_INITIALIZE uint16 = 13577
	ER_IB_MSG_PAGE_ARCH_NO_RESET_POINTS uint16 = 13578
	ER_IB_WRN_PAGE_ARCH_FLUSH_DATA uint16 = 13579
	ER_IB_ERR_PAGE_ARCH_INVALID_DOUBLE_WRITE_BUF uint16 = 13580
	ER_IB_ERR_PAGE_ARCH_RECOVERY_FAILED uint16 = 13581
	ER_IB_ERR_PAGE_ARCH_INVALID_FORMAT uint16 = 13582
	ER_INVALID_XPLUGIN_SOCKET_SAME_AS_SERVER uint16 = 13583
	ER_INNODB_UNABLE_TO_ACQUIRE_DD_OBJECT uint16 = 13584
	ER_WARN_LOG_DEPRECATED_PARTITION_PREFIX_KEY uint16 = 13585
	ER_IB_MSG_UNDO_TRUNCATE_TOO_OFTEN uint16 = 13586
	ER_GRP_RPL_IS_STARTING uint16 = 13587
	ER_IB_MSG_INVALID_LOCATION_FOR_TABLESPACE uint16 = 13588
	ER_IB_MSG_INVALID_LOCATION_WRONG_DB uint16 = 13589
	ER_IB_MSG_CANNOT_FIND_DD_UNDO_SPACE uint16 = 13590
	ER_GRP_RPL_RECOVERY_ENDPOINT_FORMAT uint16 = 13591
	ER_GRP_RPL_RECOVERY_ENDPOINT_INVALID uint16 = 13592
	ER_GRP_RPL_RECOVERY_ENDPOINT_INVALID_DONOR_ENDPOINT uint16 = 13593
	ER_GRP_RPL_RECOVERY_ENDPOINT_INTERFACES_IPS uint16 = 13594
	ER_WARN_TLS_CHANNEL_INITIALIZATION_ERROR uint16 = 13595
	ER_XPLUGIN_FAILED_TO_VALIDATE_ADDRESS uint16 = 13596
	ER_XPLUGIN_FAILED_TO_BIND_INTERFACE_ADDRESS uint16 = 13597
	ER_IB_ERR_RECOVERY_REDO_DISABLED uint16 = 13598
	ER_IB_WRN_FAST_SHUTDOWN_REDO_DISABLED uint16 = 13599
	ER_IB_WRN_REDO_DISABLED uint16 = 13600
	ER_IB_WRN_REDO_ENABLED uint16 = 13601
	ER_TLS_CONFIGURED_FOR_CHANNEL uint16 = 13602
	ER_TLS_CONFIGURATION_REUSED uint16 = 13603
	ER_IB_TABLESPACE_PATH_VALIDATION_SKIPPED uint16 = 13604
	ER_IB_CANNOT_UPGRADE_WITH_DISCARDED_TABLESPACES uint16 = 13605
	ER_USERNAME_TRUNKATED uint16 = 13606
	ER_HOSTNAME_TRUNKATED uint16 = 13607
	ER_IB_MSG_TRX_RECOVERY_ROLLBACK_NOT_COMPLETED uint16 = 13608
	ER_AUTHCACHE_ROLE_EDGES_IGNORED_EMPTY_NAME uint16 = 13609
	ER_AUTHCACHE_ROLE_EDGES_UNKNOWN_AUTHORIZATION_ID uint16 = 13610
	ER_AUTHCACHE_DEFAULT_ROLES_IGNORED_EMPTY_NAME uint16 = 13611
	ER_AUTHCACHE_DEFAULT_ROLES_UNKNOWN_AUTHORIZATION_ID uint16 = 13612
	ER_IB_ERR_DDL_LOG_INSERT_FAILURE uint16 = 13613
	ER_IB_LOCK_VALIDATE_LATCH_ORDER_VIOLATION uint16 = 13614
	ER_IB_RELOCK_LATCH_ORDER_VIOLATION uint16 = 13615
	//OBSOLETE_ER_IB_MSG_1352 uint16 = 13616
	//OBSOLETE_ER_IB_MSG_1353 uint16 = 13617
	//OBSOLETE_ER_IB_MSG_1354 uint16 = 13618
	//OBSOLETE_ER_IB_MSG_1355 uint16 = 13619
	//OBSOLETE_ER_IB_MSG_1356 uint16 = 13620
	ER_IB_MSG_1357 uint16 = 13621
	ER_IB_MSG_1358 uint16 = 13622
	ER_IB_MSG_1359 uint16 = 13623
	ER_IB_FAILED_TO_DELETE_TABLESPACE_FILE uint16 = 13624
	ER_IB_UNABLE_TO_EXPAND_TEMPORARY_TABLESPACE_POOL uint16 = 13625
	ER_IB_TMP_TABLESPACE_CANNOT_CREATE_DIRECTORY uint16 = 13626
	ER_IB_MSG_SCANNING_TEMP_TABLESPACE_DIR uint16 = 13627
	ER_IB_ERR_TEMP_TABLESPACE_DIR_DOESNT_EXIST uint16 = 13628
	ER_IB_ERR_TEMP_TABLESPACE_DIR_EMPTY uint16 = 13629
	ER_IB_ERR_TEMP_TABLESPACE_DIR_CONTAINS_SEMICOLON uint16 = 13630
	ER_IB_ERR_TEMP_TABLESPACE_DIR_SUBDIR_OF_DATADIR uint16 = 13631
	ER_IB_ERR_SCHED_SETAFFNINITY_FAILED uint16 = 13632
	ER_IB_ERR_UNKNOWN_PAGE_FETCH_MODE uint16 = 13633
	ER_IB_ERR_LOG_PARSING_BUFFER_OVERFLOW uint16 = 13634
	ER_IB_ERR_NOT_ENOUGH_MEMORY_FOR_PARSE_BUFFER uint16 = 13635
	ER_IB_MSG_1372 uint16 = 13636
	ER_IB_MSG_1373 uint16 = 13637
	ER_IB_MSG_1374 uint16 = 13638
	ER_IB_MSG_1375 uint16 = 13639
	ER_IB_ERR_ZLIB_UNCOMPRESS_FAILED uint16 = 13640
	ER_IB_ERR_ZLIB_BUF_ERROR uint16 = 13641
	ER_IB_ERR_ZLIB_MEM_ERROR uint16 = 13642
	ER_IB_ERR_ZLIB_DATA_ERROR uint16 = 13643
	ER_IB_ERR_ZLIB_UNKNOWN_ERROR uint16 = 13644
	ER_IB_MSG_1381 uint16 = 13645
	ER_IB_ERR_INDEX_RECORDS_WRONG_ORDER uint16 = 13646
	ER_IB_ERR_INDEX_DUPLICATE_KEY uint16 = 13647
	ER_IB_ERR_FOUND_N_DUPLICATE_KEYS uint16 = 13648
	ER_IB_ERR_FOUND_N_RECORDS_WRONG_ORDER uint16 = 13649
	ER_IB_ERR_PARALLEL_READ_OOM uint16 = 13650
	ER_IB_MSG_UNDO_MARKED_ACTIVE uint16 = 13651
	ER_IB_MSG_UNDO_ALTERED_ACTIVE uint16 = 13652
	ER_IB_MSG_UNDO_ALTERED_INACTIVE uint16 = 13653
	ER_IB_MSG_UNDO_MARKED_EMPTY uint16 = 13654
	ER_IB_MSG_UNDO_TRUNCATE_DELAY_BY_CLONE uint16 = 13655
	ER_IB_MSG_UNDO_TRUNCATE_DELAY_BY_MDL uint16 = 13656
	ER_IB_MSG_INJECT_CRASH uint16 = 13657
	ER_IB_MSG_INJECT_FAILURE uint16 = 13658
	ER_GRP_RPL_TIMEOUT_RECEIVED_VC_LEAVE_ON_REJOIN uint16 = 13659
	ER_RPL_ASYNC_RECONNECT_FAIL_NO_SOURCE uint16 = 13660
	ER_UDF_REGISTER_SERVICE_ERROR uint16 = 13661
	ER_UDF_REGISTER_ERROR uint16 = 13662
	ER_UDF_UNREGISTER_ERROR uint16 = 13663
	ER_EMPTY_PRIVILEGE_NAME_IGNORED uint16 = 13664
	ER_IB_MSG_INCORRECT_SIZE uint16 = 13665
	ER_TMPDIR_PATH_TOO_LONG uint16 = 13666
	ER_ERROR_LOG_DESTINATION_NOT_A_FILE uint16 = 13667
	ER_NO_ERROR_LOG_PARSER_CONFIGURED uint16 = 13668
	ER_UPGRADE_NONEXISTENT_SCHEMA uint16 = 13669
	ER_IB_MSG_CREATED_UNDO_SPACE uint16 = 13670
	ER_IB_MSG_DROPPED_UNDO_SPACE uint16 = 13671
	ER_IB_MSG_MASTER_KEY_ROTATED uint16 = 13672
	ER_IB_DBLWR_DECOMPRESS_FAILED uint16 = 13673
	ER_IB_DBLWR_DECRYPT_FAILED uint16 = 13674
	ER_IB_DBLWR_KEY_MISSING uint16 = 13675
	ER_INNODB_IO_WRITE_ERROR_RETRYING uint16 = 13676
	ER_INNODB_IO_WRITE_FAILED uint16 = 13677
	ER_LOG_COMPONENT_CANNOT_INIT uint16 = 13678
	ER_RPL_ASYNC_CHANNEL_CANT_CONNECT uint16 = 13679
	ER_RPL_ASYNC_SENDER_ADDED uint16 = 13680
	ER_RPL_ASYNC_SENDER_REMOVED uint16 = 13681
	ER_RPL_ASYNC_CHANNEL_STOPPED_QUORUM_LOST uint16 = 13682
	ER_RPL_ASYNC_CHANNEL_CANT_CONNECT_NO_QUORUM uint16 = 13683
	ER_RPL_ASYNC_EXECUTING_QUERY uint16 = 13684
	ER_RPL_REPLICA_MONITOR_IO_THREAD_EXITING uint16 = 13685
	ER_RPL_ASYNC_MANAGED_NAME_REMOVED uint16 = 13686
	ER_RPL_ASYNC_MANAGED_NAME_ADDED uint16 = 13687
	ER_RPL_ASYNC_READ_FAILOVER_TABLE uint16 = 13688
	ER_RPL_REPLICA_MONITOR_IO_THREAD_RECONNECT_CHANNEL uint16 = 13689
	ER_SLAVE_ANONYMOUS_TO_GTID_IS_LOCAL_OR_UUID_AND_GTID_MODE_NOT_ON uint16 = 13690
	ER_REPLICA_ANONYMOUS_TO_GTID_UUID_SAME_AS_GROUP_NAME uint16 = 13691
	ER_GRP_RPL_GRP_NAME_IS_SAME_AS_ANONYMOUS_TO_GTID_UUID uint16 = 13692
	ER_WARN_GTID_THRESHOLD_BREACH uint16 = 13693
	ER_HEALTH_INFO uint16 = 13694
	ER_HEALTH_WARNING uint16 = 13695
	ER_HEALTH_ERROR uint16 = 13696
	ER_HEALTH_WARNING_DISK_USAGE_LEVEL_1 uint16 = 13697
	ER_HEALTH_WARNING_DISK_USAGE_LEVEL_2 uint16 = 13698
	ER_HEALTH_WARNING_DISK_USAGE_LEVEL_3 uint16 = 13699
	ER_IB_INNODB_TBSP_OUT_OF_SPACE uint16 = 13700
	ER_GRP_RPL_APPLIER_CHANNEL_STILL_RUNNING uint16 = 13701
	ER_RPL_ASYNC_RECONNECT_GTID_MODE_OFF_CHANNEL uint16 = 13702
	ER_FIREWALL_SERVICES_NOT_ACQUIRED uint16 = 13703
	ER_FIREWALL_UDF_REGISTER_FAILED uint16 = 13704
	ER_FIREWALL_PFS_TABLE_REGISTER_FAILED uint16 = 13705
	ER_IB_MSG_STATS_SAMPLING_TOO_LARGE uint16 = 13706

	//50,000 to 51,999: Error codes reserved for use by third parties.
	//no such code
)

/*
SQLSTATE
information from https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html
SQLSTATE value: This value is a five-character string (for example, '42S02'). SQLSTATE values are taken from ANSI SQL and ODBC and are more standardized than the numeric error codes. The first two characters of an SQLSTATE value indicate the error class:
Class = '00' indicates success.
Class = '01' indicates a warning.
Class = '02' indicates “not found.” This is relevant within the context of cursors and is used to control what happens when a cursor reaches the end of a data set. This condition also occurs for SELECT ... INTO var_list statements that retrieve no rows.
Class > '02' indicates an exception.

For server-side errors, not all MySQL error numbers have corresponding SQLSTATE values. In these cases, 'HY000' (general error) is used.
For client-side errors, the SQLSTATE value is always 'HY000' (general error), so it is not meaningful for distinguishing one client error from another.
*/

//error message components
type errorMsgItem struct{
	errorCode uint16
	sqlStates []string
	errorMsgOrFormat string
}

/*
error code -> [error code integer ; [sqlstate1,...] ; error message or format]

information merged from:
	https://dev.mysql.com/doc/mysql-errors/8.0/en/server-error-reference.html

	under mysql 8.0.23 installation directory(like /usr/local/opt/mysql@8.0)?:
		/share/mysql/messages_to_clients.txt
		/share/mysql/messages_to_error_log.txt
	
*/
var errorMsgRefer=map[uint16]errorMsgItem{
	//OBSOLETE_ER_HASHCHK : {0000,[]string{""},"hashchk"},
	//OBSOLETE_ER_NISAMCHK : {0000,[]string{""},"isamchk"},
	ER_NO : {1002,[]string{"HY000"},"NO"},
	ER_YES : {1003,[]string{"HY000"},"YES"},
	ER_CANT_CREATE_FILE : {1004,[]string{"HY000"},"Can't create file '%-.200s' (errno: %d - %s)"},
	ER_CANT_CREATE_TABLE : {1005,[]string{"HY000"},"Can't create table '%-.200s' (errno: %d - %s)"},
	ER_CANT_CREATE_DB : {1006,[]string{"HY000"},"Can't create database '%-.192s' (errno: %d - %s)"},
	ER_DB_CREATE_EXISTS : {1007,[]string{"HY000"},"Can't create database '%-.192s'; database exists"},
	ER_DB_DROP_EXISTS : {1008,[]string{"HY000"},"Can't drop database '%-.192s'; database doesn't exist"},
	//OBSOLETE_ER_DB_DROP_DELETE : {0000,[]string{""},"Error dropping database (can't delete '%-.192s', errno: %d - %s)"},
	ER_DB_DROP_RMDIR : {1010,[]string{"HY000"},"Error dropping database (can't rmdir '%-.192s', errno: %d - %s)"},
	//OBSOLETE_ER_CANT_DELETE_FILE : {0000,[]string{""},"Error on delete of '%-.192s' (errno: %d - %s)"},
	ER_CANT_FIND_SYSTEM_REC : {1012,[]string{"HY000"},"Can't read record in system table"},
	ER_CANT_GET_STAT : {1013,[]string{"HY000"},"Can't get status of '%-.200s' (errno: %d - %s)"},
	//OBSOLETE_ER_CANT_GET_WD : {0000,[]string{""},"Can't get working directory (errno: %d - %s)"},
	ER_CANT_LOCK : {1015,[]string{"HY000"},"Can't lock file (errno: %d - %s)"},
	ER_CANT_OPEN_FILE : {1016,[]string{"HY000"},"Can't open file: '%-.200s' (errno: %d - %s)"},
	ER_FILE_NOT_FOUND : {1017,[]string{"HY000"},"Can't find file: '%-.200s' (errno: %d - %s)"},
	ER_CANT_READ_DIR : {1018,[]string{"HY000"},"Can't read dir of '%-.192s' (errno: %d - %s)"},
	//OBSOLETE_ER_CANT_SET_WD : {0000,[]string{""},"Can't change dir to '%-.192s' (errno: %d - %s)"},
	ER_CHECKREAD : {1020,[]string{"HY000"},"Record has changed since last read in table '%-.192s'"},
	//OBSOLETE_ER_DISK_FULL : {0000,[]string{""},"Disk full (%s); waiting for someone to free some space... (errno: %d - %s)"},
	ER_DUP_KEY : {1022,[]string{"23000"},"Can't write; duplicate key in table '%-.192s'"},
	//OBSOLETE_ER_ERROR_ON_CLOSE : {0000,[]string{""},"Error on close of '%-.192s' (errno: %d - %s)"},
	ER_ERROR_ON_READ : {1024,[]string{"HY000"},"Error reading file '%-.200s' (errno: %d - %s)"},
	ER_ERROR_ON_RENAME : {1025,[]string{"HY000"},"Error on rename of '%-.210s' to '%-.210s' (errno: %d - %s)"},
	ER_ERROR_ON_WRITE : {1026,[]string{"HY000"},"Error writing file '%-.200s' (errno: %d - %s)"},
	ER_FILE_USED : {1027,[]string{"HY000"},"'%-.192s' is locked against change"},
	//OBSOLETE_ER_FILSORT_ABORT : {1028,[]string{"HY000"},"Sort aborted"},
	//OBSOLETE_ER_FORM_NOT_FOUND : {0000,[]string{""},"View '%-.192s' doesn't exist for '%-.192s'"},
	ER_GET_ERRNO : {1030,[]string{"HY000"},"Got error %d - '%-.192s' from storage engine"},
	ER_ILLEGAL_HA : {1031,[]string{"HY000"},"Table storage engine for '%-.192s' doesn't have this option"},
	ER_KEY_NOT_FOUND : {1032,[]string{"HY000"},"Can't find record in '%-.192s'"},
	ER_NOT_FORM_FILE : {1033,[]string{"HY000"},"Incorrect information in file: '%-.200s'"},
	ER_NOT_KEYFILE : {1034,[]string{"HY000"},"Incorrect key file for table '%-.200s'; try to repair it"},
	ER_OLD_KEYFILE : {1035,[]string{"HY000"},"Old key file for table '%-.192s'; repair it!"},
	ER_OPEN_AS_READONLY : {1036,[]string{"HY000"},"Table '%-.192s' is read only"},
	ER_OUTOFMEMORY : {1037,[]string{"HY001","S1001"},"Out of memory; restart server and try again (needed %d bytes)"},
	ER_OUT_OF_SORTMEMORY : {1038,[]string{"HY001","S1001"},"Out of sort memory, consider increasing server sort buffer size"},
	//OBSOLETE_ER_UNEXPECTED_EOF : {0000,[]string{""},"Unexpected EOF found when reading file '%-.192s' (errno: %d - %s)"},
	ER_CON_COUNT_ERROR : {1040,[]string{"08004"},"Too many connections"},
	ER_OUT_OF_RESOURCES : {1041,[]string{"HY000"},"Out of memory; check if mysqld or some other process uses all available memory; if not, you may have to use 'ulimit' to allow mysqld to use more memory or you can add more swap space"},
	ER_BAD_HOST_ERROR : {1042,[]string{"08S01"},"Can't get hostname for your address"},
	ER_HANDSHAKE_ERROR : {1043,[]string{"08S01"},"Bad handshake"},
	ER_DBACCESS_DENIED_ERROR : {1044,[]string{"42000"},"Access denied for user '%-.48s'@'%-.64s' to database '%-.192s'"},
	ER_ACCESS_DENIED_ERROR : {1045,[]string{"28000"},"Access denied for user '%-.48s'@'%-.64s' (using password: %s)"},
	ER_NO_DB_ERROR : {1046,[]string{"3D000"},"No database selected"},
	ER_UNKNOWN_COM_ERROR : {1047,[]string{"08S01"},"Unknown command"},
	ER_BAD_NULL_ERROR : {1048,[]string{"23000"},"Column '%-.192s' cannot be null"},
	ER_BAD_DB_ERROR : {1049,[]string{"42000"},"Unknown database '%-.192s'"},
	ER_TABLE_EXISTS_ERROR : {1050,[]string{"42S01"},"Table '%-.192s' already exists"},
	ER_BAD_TABLE_ERROR : {1051,[]string{"42S02"},"Unknown table '%-.129s'"},
	ER_NON_UNIQ_ERROR : {1052,[]string{"23000"},"Column '%-.192s' in %-.192s is ambiguous"},
	ER_SERVER_SHUTDOWN : {1053,[]string{"08S01"},"Server shutdown in progress"},
	ER_BAD_FIELD_ERROR : {1054,[]string{"42S22","S0022"},"Unknown column '%-.192s' in '%-.192s'"},
	ER_WRONG_FIELD_WITH_GROUP : {1055,[]string{"42000","S1009"},"'%-.192s' isn't in GROUP BY"},
	ER_WRONG_GROUP_FIELD : {1056,[]string{"42000","S1009"},"Can't group on '%-.192s'"},
	ER_WRONG_SUM_SELECT : {1057,[]string{"42000","S1009"},"Statement has sum functions and columns in same statement"},
	ER_WRONG_VALUE_COUNT : {1058,[]string{"21S01"},"Column count doesn't match value count"},
	ER_TOO_LONG_IDENT : {1059,[]string{"42000","S1009"},"Identifier name '%-.100s' is too long"},
	ER_DUP_FIELDNAME : {1060,[]string{"42S21","S1009"},"Duplicate column name '%-.192s'"},
	ER_DUP_KEYNAME : {1061,[]string{"42000","S1009"},"Duplicate key name '%-.192s'"},
	ER_DUP_ENTRY : {1062,[]string{"23000","S1009"},"Duplicate entry '%-.192s' for key %d"},
	ER_WRONG_FIELD_SPEC : {1063,[]string{"42000","S1009"},"Incorrect column specifier for column '%-.192s'"},
	ER_PARSE_ERROR : {1064,[]string{"42000","s1009"},"%s near '%-.80s' at line %d"},
	ER_EMPTY_QUERY : {1065,[]string{"42000"},"Query was empty"},
	ER_NONUNIQ_TABLE : {1066,[]string{"42000","S1009"},"Not unique table/alias: '%-.192s'"},
	ER_INVALID_DEFAULT : {1067,[]string{"42000","S1009"},"Invalid default value for '%-.192s'"},
	ER_MULTIPLE_PRI_KEY : {1068,[]string{"42000","S1009"},"Multiple primary key defined"},
	ER_TOO_MANY_KEYS : {1069,[]string{"42000","S1009"},"Too many keys specified; max %d keys allowed"},
	ER_TOO_MANY_KEY_PARTS : {1070,[]string{"42000","S1009"},"Too many key parts specified; max %d parts allowed"},
	ER_TOO_LONG_KEY : {1071,[]string{"42000","S1009"},"Specified key was too long; max key length is %d bytes"},
	ER_KEY_COLUMN_DOES_NOT_EXITS : {1072,[]string{"42000","S1009"},"Key column '%-.192s' doesn't exist in table"},
	ER_BLOB_USED_AS_KEY : {1073,[]string{"42000","S1009"},"BLOB column '%-.192s' can't be used in key specification with the used table type"},
	ER_TOO_BIG_FIELDLENGTH : {1074,[]string{"42000","S1009"},"Column length too big for column '%-.192s' (max = %lu); use BLOB or TEXT instead"},
	ER_WRONG_AUTO_KEY : {1075,[]string{"42000","S1009"},"Incorrect table definition; there can be only one auto column and it must be defined as a key"},
	ER_READY : {1076,[]string{"HY000"},"%s: ready for connections.\nVersion: '%s'  socket: '%s'  port: %d"},
	//OBSOLETE_ER_NORMAL_SHUTDOWN : {1077,[]string{"HY000"},"%s: Normal shutdown\n"},
	//OBSOLETE_ER_GOT_SIGNAL : {0000,[]string{""},"%s: Got signal %d. Aborting!\n"},
	ER_SHUTDOWN_COMPLETE : {1079,[]string{"HY000"},"%s: Shutdown complete\n"},
	ER_FORCING_CLOSE : {1080,[]string{"08S01"},"%s: Forcing close of thread %ld  user: '%-.48s'\n"},
	ER_IPSOCK_ERROR : {1081,[]string{"08S01"},"Can't create IP socket"},
	ER_NO_SUCH_INDEX : {1082,[]string{"42S12","S1009"},"Table '%-.192s' has no index like the one used in CREATE INDEX; recreate the table"},
	ER_WRONG_FIELD_TERMINATORS : {1083,[]string{"42000","S1009"},"Field separator argument is not what is expected; check the manual"},
	ER_BLOBS_AND_NO_TERMINATED : {1084,[]string{"42000","S1009"},"You can't use fixed rowlength with BLOBs; please use 'fields terminated by'"},
	ER_TEXTFILE_NOT_READABLE : {1085,[]string{"HY000"},"The file '%-.128s' must be in the database directory or be readable by all"},
	ER_FILE_EXISTS_ERROR : {1086,[]string{"HY000"},"File '%-.200s' already exists"},
	ER_LOAD_INFO : {1087,[]string{"HY000"},"Records: %ld  Deleted: %ld  Skipped: %ld  Warnings: %ld"},
	ER_ALTER_INFO : {1088,[]string{"HY000"},"Records: %ld  Duplicates: %ld"},
	ER_WRONG_SUB_KEY : {1089,[]string{"HY000"},"Incorrect prefix key; the used key part isn't a string, the used length is longer than the key part, or the storage engine doesn't support unique prefix keys"},
	ER_CANT_REMOVE_ALL_FIELDS : {1090,[]string{"42000"},"You can't delete all columns with ALTER TABLE; use DROP TABLE instead"},
	ER_CANT_DROP_FIELD_OR_KEY : {1091,[]string{"42000"},"Can't DROP '%-.192s'; check that column/key exists"},
	ER_INSERT_INFO : {1092,[]string{"HY000"},"Records: %ld  Duplicates: %ld  Warnings: %ld"},
	ER_UPDATE_TABLE_USED : {1093,[]string{"HY000"},"You can't specify target table '%-.192s' for update in FROM clause"},
	ER_NO_SUCH_THREAD : {1094,[]string{"HY000"},"Unknown thread id: %lu"},
	ER_KILL_DENIED_ERROR : {1095,[]string{"HY000"},"You are not owner of thread %lu"},
	ER_NO_TABLES_USED : {1096,[]string{"HY000"},"No tables used"},
	ER_TOO_BIG_SET : {1097,[]string{"HY000"},"Too many strings for column %-.192s and SET"},
	ER_NO_UNIQUE_LOGFILE : {1098,[]string{"HY000"},"Can't generate a unique log-filename %-.200s.(1-999)\n"},
	ER_TABLE_NOT_LOCKED_FOR_WRITE : {1099,[]string{"HY000"},"Table '%-.192s' was locked with a READ lock and can't be updated"},
	ER_TABLE_NOT_LOCKED : {1100,[]string{"HY000"},"Table '%-.192s' was not locked with LOCK TABLES"},
	ER_BLOB_CANT_HAVE_DEFAULT : {1101,[]string{"42000"},"BLOB, TEXT, GEOMETRY or JSON column '%-.192s' can't have a default value"},
	ER_WRONG_DB_NAME : {1102,[]string{"42000"},"Incorrect database name '%-.100s'"},
	ER_WRONG_TABLE_NAME : {1103,[]string{"42000"},"Incorrect table name '%-.100s'"},
	ER_TOO_BIG_SELECT : {1104,[]string{"42000"},"The SELECT would examine more than MAX_JOIN_SIZE rows; check your WHERE and use SET SQL_BIG_SELECTS=1 or SET MAX_JOIN_SIZE=# if the SELECT is okay"},
	ER_UNKNOWN_ERROR : {1105,[]string{"HY000"},"Unknown error"},
	ER_UNKNOWN_PROCEDURE : {1106,[]string{"42000"},"Unknown procedure '%-.192s'"},
	ER_WRONG_PARAMCOUNT_TO_PROCEDURE : {1107,[]string{"42000"},"Incorrect parameter count to procedure '%-.192s'"},
	ER_WRONG_PARAMETERS_TO_PROCEDURE : {1108,[]string{"HY000"},"Incorrect parameters to procedure '%-.192s'"},
	ER_UNKNOWN_TABLE : {1109,[]string{"42S02"},"Unknown table '%-.192s' in %-.32s"},
	ER_FIELD_SPECIFIED_TWICE : {1110,[]string{"42000"},"Column '%-.192s' specified twice"},
	ER_INVALID_GROUP_FUNC_USE : {1111,[]string{"HY000"},"Invalid use of group function"},
	ER_UNSUPPORTED_EXTENSION : {1112,[]string{"42000"},"Table '%-.192s' uses an extension that doesn't exist in this MySQL version"},
	ER_TABLE_MUST_HAVE_COLUMNS : {1113,[]string{"42000"},"A table must have at least 1 column"},
	ER_RECORD_FILE_FULL : {1114,[]string{"HY000"},"The table '%-.192s' is full"},
	ER_UNKNOWN_CHARACTER_SET : {1115,[]string{"42000"},"Unknown character set: '%-.64s'"},
	ER_TOO_MANY_TABLES : {1116,[]string{"HY000"},"Too many tables; MySQL can only use %d tables in a join"},
	ER_TOO_MANY_FIELDS : {1117,[]string{"HY000"},"Too many columns"},
	ER_TOO_BIG_ROWSIZE : {1118,[]string{"42000"},"Row size too large. The maximum row size for the used table type, not counting BLOBs, is %ld. This includes storage overhead, check the manual. You have to change some columns to TEXT or BLOBs"},
	ER_STACK_OVERRUN : {1119,[]string{"HY000"},"Thread stack overrun:  Used: %ld of a %ld stack.  Use 'mysqld --thread_stack=#' to specify a bigger stack if needed"},
	ER_WRONG_OUTER_JOIN_UNUSED : {1120,[]string{"42000"},"Cross dependency found in OUTER JOIN; examine your ON conditions"},
	ER_NULL_COLUMN_IN_INDEX : {1121,[]string{"42000"},"Table handler doesn't support NULL in given index. Please change column '%-.192s' to be NOT NULL or use another handler"},
	ER_CANT_FIND_UDF : {1122,[]string{"HY000"},"Can't load function '%-.192s'"},
	ER_CANT_INITIALIZE_UDF : {1123,[]string{"HY000"},"Can't initialize function '%-.192s'; %-.80s"},
	ER_UDF_NO_PATHS : {1124,[]string{"HY000"},"No paths allowed for shared library"},
	ER_UDF_EXISTS : {1125,[]string{"HY000"},"Function '%-.192s' already exists"},
	ER_CANT_OPEN_LIBRARY : {1126,[]string{"HY000"},"Can't open shared library '%-.192s' (errno: %d %-.128s)"},
	ER_CANT_FIND_DL_ENTRY : {1127,[]string{"HY000"},"Can't find symbol '%-.128s' in library"},
	ER_FUNCTION_NOT_DEFINED : {1128,[]string{"HY000"},"Function '%-.192s' is not defined"},
	ER_HOST_IS_BLOCKED : {1129,[]string{"HY000"},"Host '%-.255s' is blocked because of many connection errors; unblock with 'mysqladmin flush-hosts'"},
	ER_HOST_NOT_PRIVILEGED : {1130,[]string{"HY000"},"Host '%-.255s' is not allowed to connect to this MySQL server"},
	ER_PASSWORD_ANONYMOUS_USER : {1131,[]string{"42000"},"You are using MySQL as an anonymous user and anonymous users are not allowed to change passwords"},
	ER_PASSWORD_NOT_ALLOWED : {1132,[]string{"42000"},"You must have privileges to update tables in the mysql database to be able to change passwords for others"},
	ER_PASSWORD_NO_MATCH : {1133,[]string{"42000"},"Can't find any matching row in the user table"},
	ER_UPDATE_INFO : {1134,[]string{"HY000"},"Rows matched: %ld  Changed: %ld  Warnings: %ld"},
	ER_CANT_CREATE_THREAD : {1135,[]string{"HY000"},"Can't create a new thread (errno %d); if you are not out of available memory, you can consult the manual for a possible OS-dependent bug"},
	ER_WRONG_VALUE_COUNT_ON_ROW : {1136,[]string{"21S01"},"Column count doesn't match value count at row %ld"},
	ER_CANT_REOPEN_TABLE : {1137,[]string{"HY000"},"Can't reopen table: '%-.192s'"},
	ER_INVALID_USE_OF_NULL : {1138,[]string{"22004"},"Invalid use of NULL value"},
	ER_REGEXP_ERROR : {1139,[]string{"42000"},"Got error '%-.64s' from regexp"},
	ER_MIX_OF_GROUP_FUNC_AND_FIELDS : {1140,[]string{"42000"},"Mixing of GROUP columns (MIN(),MAX(),COUNT(),...) with no GROUP columns is illegal if there is no GROUP BY clause"},
	ER_NONEXISTING_GRANT : {1141,[]string{"42000"},"There is no such grant defined for user '%-.48s' on host '%-.255s'"},
	ER_TABLEACCESS_DENIED_ERROR : {1142,[]string{"42000"},"%-.128s command denied to user '%-.48s'@'%-.64s' for table '%-.64s'"},
	ER_COLUMNACCESS_DENIED_ERROR : {1143,[]string{"42000"},"%-.16s command denied to user '%-.48s'@'%-.64s' for column '%-.192s' in table '%-.192s'"},
	ER_ILLEGAL_GRANT_FOR_TABLE : {1144,[]string{"42000"},"Illegal GRANT/REVOKE command; please consult the manual to see which privileges can be used"},
	ER_GRANT_WRONG_HOST_OR_USER : {1145,[]string{"42000"},"The host or user argument to GRANT is too long"},
	ER_NO_SUCH_TABLE : {1146,[]string{"42S02"},"Table '%-.192s.%-.192s' doesn't exist"},
	ER_NONEXISTING_TABLE_GRANT : {1147,[]string{"42000"},"There is no such grant defined for user '%-.48s' on host '%-.255s' on table '%-.192s'"},
	ER_NOT_ALLOWED_COMMAND : {1148,[]string{"42000"},"The used command is not allowed with this MySQL version"},
	ER_SYNTAX_ERROR : {1149,[]string{"42000"},"You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use"},
	//OBSOLETE_ER_UNUSED1 : {0000,[]string{""},"Delayed insert thread couldn't get requested lock for table %-.192s"},
	//OBSOLETE_ER_UNUSED2 : {0000,[]string{""},"Too many delayed threads in use"},
	ER_ABORTING_CONNECTION : {1152,[]string{"08S01"},"Aborted connection %ld to db: '%-.192s' user: '%-.48s' (%-.64s)"},
	ER_NET_PACKET_TOO_LARGE : {1153,[]string{"08S01"},"Got a packet bigger than 'max_allowed_packet' bytes"},
	ER_NET_READ_ERROR_FROM_PIPE : {1154,[]string{"08S01"},"Got a read error from the connection pipe"},
	ER_NET_FCNTL_ERROR : {1155,[]string{"08S01"},"Got an error from fcntl()"},
	ER_NET_PACKETS_OUT_OF_ORDER : {1156,[]string{"08S01"},"Got packets out of order"},
	ER_NET_UNCOMPRESS_ERROR : {1157,[]string{"08S01"},"Couldn't uncompress communication packet"},
	ER_NET_READ_ERROR : {1158,[]string{"08S01"},"Got an error reading communication packets"},
	ER_NET_READ_INTERRUPTED : {1159,[]string{"08S01"},"Got timeout reading communication packets"},
	ER_NET_ERROR_ON_WRITE : {1160,[]string{"08S01"},"Got an error writing communication packets"},
	ER_NET_WRITE_INTERRUPTED : {1161,[]string{"08S01"},"Got timeout writing communication packets"},
	ER_TOO_LONG_STRING : {1162,[]string{"42000"},"Result string is longer than 'max_allowed_packet' bytes"},
	ER_TABLE_CANT_HANDLE_BLOB : {1163,[]string{"42000"},"The used table type doesn't support BLOB/TEXT columns"},
	ER_TABLE_CANT_HANDLE_AUTO_INCREMENT : {1164,[]string{"42000"},"The used table type doesn't support AUTO_INCREMENT columns"},
	//OBSOLETE_ER_UNUSED3 : {0000,[]string{""},"INSERT DELAYED can't be used with table '%-.192s' because it is locked with LOCK TABLES"},
	ER_WRONG_COLUMN_NAME : {1166,[]string{"42000"},"Incorrect column name '%-.100s'"},
	ER_WRONG_KEY_COLUMN : {1167,[]string{"42000"},"The used storage engine can't index column '%-.192s'"},
	ER_WRONG_MRG_TABLE : {1168,[]string{"HY000"},"Unable to open underlying table which is differently defined or of non-MyISAM type or doesn't exist"},
	ER_DUP_UNIQUE : {1169,[]string{"23000"},"Can't write, because of unique constraint, to table '%-.192s'"},
	ER_BLOB_KEY_WITHOUT_LENGTH : {1170,[]string{"42000"},"BLOB/TEXT column '%-.192s' used in key specification without a key length"},
	ER_PRIMARY_CANT_HAVE_NULL : {1171,[]string{"42000"},"All parts of a PRIMARY KEY must be NOT NULL; if you need NULL in a key, use UNIQUE instead"},
	ER_TOO_MANY_ROWS : {1172,[]string{"42000"},"Result consisted of more than one row"},
	ER_REQUIRES_PRIMARY_KEY : {1173,[]string{"42000"},"This table type requires a primary key"},
	//OBSOLETE_ER_NO_RAID_COMPILED : {0000,[]string{""},"This version of MySQL is not compiled with RAID support"},
	ER_UPDATE_WITHOUT_KEY_IN_SAFE_MODE : {1175,[]string{"HY000"},"You are using safe update mode and you tried to update a table without a WHERE that uses a KEY column. %s"},
	ER_KEY_DOES_NOT_EXITS : {1176,[]string{"42000","S1009"},"Key '%-.192s' doesn't exist in table '%-.192s'"},
	ER_CHECK_NO_SUCH_TABLE : {1177,[]string{"42000"},"Can't open table"},
	ER_CHECK_NOT_IMPLEMENTED : {1178,[]string{"42000"},"The storage engine for the table doesn't support %s"},
	ER_CANT_DO_THIS_DURING_AN_TRANSACTION : {1179,[]string{"25000"},"You are not allowed to execute this command in a transaction"},
	ER_ERROR_DURING_COMMIT : {1180,[]string{"HY000"},"Got error %d - '%-.192s' during COMMIT"},
	ER_ERROR_DURING_ROLLBACK : {1181,[]string{"HY000"},"Got error %d - '%-.192s' during ROLLBACK"},
	ER_ERROR_DURING_FLUSH_LOGS : {1182,[]string{"HY000"},"Got error %d during FLUSH_LOGS"},
	//OBSOLETE_ER_ERROR_DURING_CHECKPOINT : {0000,[]string{""},"Got error %d during CHECKPOINT"},
	ER_NEW_ABORTING_CONNECTION : {1184,[]string{"08S01"},"Aborted connection %u to db: '%-.192s' user: '%-.48s' host: '%-.255s' (%-.64s)"},
	//OBSOLETE_ER_DUMP_NOT_IMPLEMENTED : {0000,[]string{""},"The storage engine for the table does not support binary table dump"},
	//OBSOLETE_ER_FLUSH_MASTER_BINLOG_CLOSED : {0000,[]string{""},"Binlog closed, cannot RESET MASTER"},
	//OBSOLETE_ER_INDEX_REBUILD : {0000,[]string{""},"Failed rebuilding the index of  dumped table '%-.192s'"},
	ER_MASTER : {1188,[]string{"HY000"},"Error from master: '%-.64s'"},
	ER_MASTER_NET_READ : {1189,[]string{"08S01"},"Net error reading from master"},
	ER_MASTER_NET_WRITE : {1190,[]string{"08S01"},"Net error writing to master"},
	ER_FT_MATCHING_KEY_NOT_FOUND : {1191,[]string{"HY000"},"Can't find FULLTEXT index matching the column list"},
	ER_LOCK_OR_ACTIVE_TRANSACTION : {1192,[]string{"HY000"},"Can't execute the given command because you have active locked tables or an active transaction"},
	ER_UNKNOWN_SYSTEM_VARIABLE : {1193,[]string{"HY000"},"Unknown system variable '%-.64s'"},
	ER_CRASHED_ON_USAGE : {1194,[]string{"HY000"},"Table '%-.192s' is marked as crashed and should be repaired"},
	ER_CRASHED_ON_REPAIR : {1195,[]string{"HY000"},"Table '%-.192s' is marked as crashed and last (automatic?) repair failed"},
	ER_WARNING_NOT_COMPLETE_ROLLBACK : {1196,[]string{"HY000"},"Some non-transactional changed tables couldn't be rolled back"},
	ER_TRANS_CACHE_FULL : {1197,[]string{"HY000"},"Multi-statement transaction required more than 'max_binlog_cache_size' bytes of storage; increase this mysqld variable and try again"},
	//OBSOLETE_ER_SLAVE_MUST_STOP : {0000,[]string{""},"This operation cannot be performed with a running slave; run STOP SLAVE first"},
	ER_SLAVE_NOT_RUNNING : {1199,[]string{"HY000"},"This operation requires a running slave; configure slave and do START SLAVE"},
	ER_BAD_SLAVE : {1200,[]string{"HY000"},"The server is not configured as slave; fix in config file or with CHANGE MASTER TO"},
	ER_MASTER_INFO : {1201,[]string{"HY000"},"Could not initialize master info structure; more error messages can be found in the MySQL error log"},
	ER_SLAVE_THREAD : {1202,[]string{"HY000"},"Could not create slave thread; check system resources"},
	ER_TOO_MANY_USER_CONNECTIONS : {1203,[]string{"42000"},"User %-.64s already has more than 'max_user_connections' active connections"},
	ER_SET_CONSTANTS_ONLY : {1204,[]string{"HY000"},"You may only use constant expressions with SET"},
	ER_LOCK_WAIT_TIMEOUT : {1205,[]string{"HY000"},"Lock wait timeout exceeded; try restarting transaction"},
	ER_LOCK_TABLE_FULL : {1206,[]string{"HY000"},"The total number of locks exceeds the lock table size"},
	ER_READ_ONLY_TRANSACTION : {1207,[]string{"25000"},"Update locks cannot be acquired during a READ UNCOMMITTED transaction"},
	//OBSOLETE_ER_DROP_DB_WITH_READ_LOCK : {0000,[]string{""},"DROP DATABASE not allowed while thread is holding global read lock"},
	//OBSOLETE_ER_CREATE_DB_WITH_READ_LOCK : {0000,[]string{""},"CREATE DATABASE not allowed while thread is holding global read lock"},
	ER_WRONG_ARGUMENTS : {1210,[]string{"HY000"},"Incorrect arguments to %s"},
	ER_NO_PERMISSION_TO_CREATE_USER : {1211,[]string{"42000"},"'%-.48s'@'%-.64s' is not allowed to create new users"},
	//OBSOLETE_ER_UNION_TABLES_IN_DIFFERENT_DIR : {0000,[]string{""},"Incorrect table definition; all MERGE tables must be in the same database"},
	ER_LOCK_DEADLOCK : {1213,[]string{"40001"},"Deadlock found when trying to get lock; try restarting transaction"},
	ER_TABLE_CANT_HANDLE_FT : {1214,[]string{"HY000"},"The used table type doesn't support FULLTEXT indexes"},
	ER_CANNOT_ADD_FOREIGN : {1215,[]string{"HY000"},"Cannot add foreign key constraint"},
	ER_NO_REFERENCED_ROW : {1216,[]string{"23000"},"Cannot add or update a child row: a foreign key constraint fails"},
	ER_ROW_IS_REFERENCED : {1217,[]string{"23000"},"Cannot delete or update a parent row: a foreign key constraint fails"},
	ER_CONNECT_TO_MASTER : {1218,[]string{"08S01"},"Error connecting to master: %-.128s"},
	//OBSOLETE_ER_QUERY_ON_MASTER : {0000,[]string{""},"Error running query on master: %-.128s"},
	ER_ERROR_WHEN_EXECUTING_COMMAND : {1220,[]string{"HY000"},"Error when executing command %s: %-.128s"},
	ER_WRONG_USAGE : {1221,[]string{"HY000"},"Incorrect usage of %s and %s"},
	ER_WRONG_NUMBER_OF_COLUMNS_IN_SELECT : {1222,[]string{"21000"},"The used SELECT statements have a different number of columns"},
	ER_CANT_UPDATE_WITH_READLOCK : {1223,[]string{"HY000"},"Can't execute the query because you have a conflicting read lock"},
	ER_MIXING_NOT_ALLOWED : {1224,[]string{"HY000"},"Mixing of transactional and non-transactional tables is disabled"},
	ER_DUP_ARGUMENT : {1225,[]string{"HY000"},"Option '%s' used twice in statement"},
	ER_USER_LIMIT_REACHED : {1226,[]string{"42000"},"User '%-.64s' has exceeded the '%s' resource (current value: %ld)"},
	ER_SPECIFIC_ACCESS_DENIED_ERROR : {1227,[]string{"42000"},"Access denied; you need (at least one of) the %-.128s privilege(s) for this operation"},
	ER_LOCAL_VARIABLE : {1228,[]string{"HY000"},"Variable '%-.64s' is a SESSION variable and can't be used with SET GLOBAL"},
	ER_GLOBAL_VARIABLE : {1229,[]string{"HY000"},"Variable '%-.64s' is a GLOBAL variable and should be set with SET GLOBAL"},
	ER_NO_DEFAULT : {1230,[]string{"42000"},"Variable '%-.64s' doesn't have a default value"},
	ER_WRONG_VALUE_FOR_VAR : {1231,[]string{"42000"},"Variable '%-.64s' can't be set to the value of '%-.200s'"},
	ER_WRONG_TYPE_FOR_VAR : {1232,[]string{"42000"},"Incorrect argument type to variable '%-.64s'"},
	ER_VAR_CANT_BE_READ : {1233,[]string{"HY000"},"Variable '%-.64s' can only be set, not read"},
	ER_CANT_USE_OPTION_HERE : {1234,[]string{"42000"},"Incorrect usage/placement of '%s'"},
	ER_NOT_SUPPORTED_YET : {1235,[]string{"42000"},"This version of MySQL doesn't yet support '%s'"},
	ER_MASTER_FATAL_ERROR_READING_BINLOG : {1236,[]string{"HY000"},"Got fatal error %d from master when reading data from binary log: '%-.512s'"},
	ER_SLAVE_IGNORED_TABLE : {1237,[]string{"HY000"},"Slave SQL thread ignored the query because of replicate-*-table rules"},
	ER_INCORRECT_GLOBAL_LOCAL_VAR : {1238,[]string{"HY000"},"Variable '%-.192s' is a %s variable"},
	ER_WRONG_FK_DEF : {1239,[]string{"42000"},"Incorrect foreign key definition for '%-.192s': %s"},
	ER_KEY_REF_DO_NOT_MATCH_TABLE_REF : {1240,[]string{"HY000"},"Key reference and table reference don't match"},
	ER_OPERAND_COLUMNS : {1241,[]string{"21000"},"Operand should contain %d column(s)"},
	ER_SUBQUERY_NO_1_ROW : {1242,[]string{"21000"},"Subquery returns more than 1 row"},
	ER_UNKNOWN_STMT_HANDLER : {1243,[]string{"HY000"},"Unknown prepared statement handler (%.*s) given to %s"},
	ER_CORRUPT_HELP_DB : {1244,[]string{"HY000"},"Help database is corrupt or does not exist"},
	//OBSOLETE_ER_CYCLIC_REFERENCE : {0000,[]string{""},"Cyclic reference on subqueries"},
	ER_AUTO_CONVERT : {1246,[]string{"HY000"},"Converting column '%s' from %s to %s"},
	ER_ILLEGAL_REFERENCE : {1247,[]string{"42S22"},"Reference '%-.64s' not supported (%s)"},
	ER_DERIVED_MUST_HAVE_ALIAS : {1248,[]string{"42000"},"Every derived table must have its own alias"},
	ER_SELECT_REDUCED : {1249,[]string{"01000"},"Select %u was reduced during optimization"},
	ER_TABLENAME_NOT_ALLOWED_HERE : {1250,[]string{"42000"},"Table '%-.192s' from one of the SELECTs cannot be used in %-.32s"},
	ER_NOT_SUPPORTED_AUTH_MODE : {1251,[]string{"08004"},"Client does not support authentication protocol requested by server; consider upgrading MySQL client"},
	ER_SPATIAL_CANT_HAVE_NULL : {1252,[]string{"42000"},"All parts of a SPATIAL index must be NOT NULL"},
	ER_COLLATION_CHARSET_MISMATCH : {1253,[]string{"42000"},"COLLATION '%s' is not valid for CHARACTER SET '%s'"},
	//OBSOLETE_ER_SLAVE_WAS_RUNNING : {0000,[]string{""},"Slave is already running"},
	//OBSOLETE_ER_SLAVE_WAS_NOT_RUNNING : {0000,[]string{""},"Slave already has been stopped"},
	ER_TOO_BIG_FOR_UNCOMPRESS : {1256,[]string{"HY000"},"Uncompressed data size too large; the maximum size is %d (probably, length of uncompressed data was corrupted)"},
	ER_ZLIB_Z_MEM_ERROR : {1257,[]string{"HY000"},"ZLIB: Not enough memory"},
	ER_ZLIB_Z_BUF_ERROR : {1258,[]string{"HY000"},"ZLIB: Not enough room in the output buffer (probably, length of uncompressed data was corrupted)"},
	ER_ZLIB_Z_DATA_ERROR : {1259,[]string{"HY000"},"ZLIB: Input data corrupted"},
	ER_CUT_VALUE_GROUP_CONCAT : {1260,[]string{"HY000"},"Row %u was cut by GROUP_CONCAT()"},
	ER_WARN_TOO_FEW_RECORDS : {1261,[]string{"01000"},"Row %ld doesn't contain data for all columns"},
	ER_WARN_TOO_MANY_RECORDS : {1262,[]string{"01000"},"Row %ld was truncated; it contained more data than there were input columns"},
	ER_WARN_NULL_TO_NOTNULL : {1263,[]string{"22004"},"Column set to default value; NULL supplied to NOT NULL column '%s' at row %ld"},
	ER_WARN_DATA_OUT_OF_RANGE : {1264,[]string{"22003"},"Out of range value for column '%s' at row %ld"},
	WARN_DATA_TRUNCATED : {1265,[]string{"01000"},"Data truncated for column '%s' at row %ld"},
	ER_WARN_USING_OTHER_HANDLER : {1266,[]string{"HY000"},"Using storage engine %s for table '%s'"},
	ER_CANT_AGGREGATE_2COLLATIONS : {1267,[]string{"HY000"},"Illegal mix of collations (%s,%s) and (%s,%s) for operation '%s'"},
	//OBSOLETE_ER_DROP_USER : {0000,[]string{""},"Cannot drop one or more of the requested users"},
	ER_REVOKE_GRANTS : {1269,[]string{"HY000"},"Can't revoke all privileges for one or more of the requested users"},
	ER_CANT_AGGREGATE_3COLLATIONS : {1270,[]string{"HY000"},"Illegal mix of collations (%s,%s), (%s,%s), (%s,%s) for operation '%s'"},
	ER_CANT_AGGREGATE_NCOLLATIONS : {1271,[]string{"HY000"},"Illegal mix of collations for operation '%s'"},
	ER_VARIABLE_IS_NOT_STRUCT : {1272,[]string{"HY000"},"Variable '%-.64s' is not a variable component (can't be used as XXXX.variable_name)"},
	ER_UNKNOWN_COLLATION : {1273,[]string{"HY000"},"Unknown collation: '%-.64s'"},
	ER_SLAVE_IGNORED_SSL_PARAMS : {1274,[]string{"HY000"},"SSL parameters in CHANGE MASTER are ignored because this MySQL slave was compiled without SSL support; they can be used later if MySQL slave with SSL is started"},
	//OBSOLETE_ER_SERVER_IS_IN_SECURE_AUTH_MODE : {1275,[]string{"HY000"},"Server is running in --secure-auth mode, but '%s'@'%s' has a password in the old format; please change the password to the new format"},
	ER_WARN_FIELD_RESOLVED : {1276,[]string{"HY000"},"Field or reference '%-.192s%s%-.192s%s%-.192s' of SELECT #%d was resolved in SELECT #%d"},
	ER_BAD_SLAVE_UNTIL_COND : {1277,[]string{"HY000"},"Incorrect parameter or combination of parameters for START SLAVE UNTIL"},
	ER_MISSING_SKIP_SLAVE : {1278,[]string{"HY000"},"It is recommended to use --skip-slave-start when doing step-by-step replication with START SLAVE UNTIL; otherwise, you will get problems if you get an unexpected slave's mysqld restart"},
	ER_UNTIL_COND_IGNORED : {1279,[]string{"HY000"},"SQL thread is not to be started so UNTIL options are ignored"},
	ER_WRONG_NAME_FOR_INDEX : {1280,[]string{"42000"},"Incorrect index name '%-.100s'"},
	ER_WRONG_NAME_FOR_CATALOG : {1281,[]string{"42000"},"Incorrect catalog name '%-.100s'"},
	//OBSOLETE_ER_WARN_QC_RESIZE : {1282,[]string{"HY000"},"Query cache failed to set size %lu; new query cache size is %lu"},
	ER_BAD_FT_COLUMN : {1283,[]string{"HY000"},"Column '%-.192s' cannot be part of FULLTEXT index"},
	ER_UNKNOWN_KEY_CACHE : {1284,[]string{"HY000"},"Unknown key cache '%-.100s'"},
	ER_WARN_HOSTNAME_WONT_WORK : {1285,[]string{"HY000"},"MySQL is started in --skip-name-resolve mode; you must restart it without this switch for this grant to work"},
	ER_UNKNOWN_STORAGE_ENGINE : {1286,[]string{"42000"},"Unknown storage engine '%s'"},
	ER_WARN_DEPRECATED_SYNTAX : {1287,[]string{"HY000"},"'%s' is deprecated and will be removed in a future release. Please use %s instead"},
	ER_NON_UPDATABLE_TABLE : {1288,[]string{"HY000"},"The target table %-.100s of the %s is not updatable"},
	ER_FEATURE_DISABLED : {1289,[]string{"HY000"},"The '%s' feature is disabled; you need MySQL built with '%s' to have it working"},
	ER_OPTION_PREVENTS_STATEMENT : {1290,[]string{"HY000"},"The MySQL server is running with the %s option so it cannot execute this statement"},
	ER_DUPLICATED_VALUE_IN_TYPE : {1291,[]string{"HY000"},"Column '%-.100s' has duplicated value '%-.64s' in %s"},
	ER_TRUNCATED_WRONG_VALUE : {1292,[]string{"22007"},"Truncated incorrect %-.64s value: '%-.128s'"},
	//OBSOLETE_ER_TOO_MUCH_AUTO_TIMESTAMP_COLS : {0000,[]string{""},"Incorrect table definition; there can be only one TIMESTAMP column with CURRENT_TIMESTAMP in DEFAULT or ON UPDATE clause"},
	ER_INVALID_ON_UPDATE : {1294,[]string{"HY000"},"Invalid ON UPDATE clause for '%-.192s' column"},
	ER_UNSUPPORTED_PS : {1295,[]string{"HY000"},"This command is not supported in the prepared statement protocol yet"},
	ER_GET_ERRMSG : {1296,[]string{"HY000"},"Got error %d '%-.100s' from %s"},
	ER_GET_TEMPORARY_ERRMSG : {1297,[]string{"HY000"},"Got temporary error %d '%-.100s' from %s"},
	ER_UNKNOWN_TIME_ZONE : {1298,[]string{"HY000"},"Unknown or incorrect time zone: '%-.64s'"},
	ER_WARN_INVALID_TIMESTAMP : {1299,[]string{"HY000"},"Invalid TIMESTAMP value in column '%s' at row %ld"},
	ER_INVALID_CHARACTER_STRING : {1300,[]string{"HY000"},"Invalid %s character string: '%.64s'"},
	ER_WARN_ALLOWED_PACKET_OVERFLOWED : {1301,[]string{"HY000"},"Result of %s() was larger than max_allowed_packet (%ld) - truncated"},
	ER_CONFLICTING_DECLARATIONS : {1302,[]string{"HY000"},"Conflicting declarations: '%s%s' and '%s%s'"},
	ER_SP_NO_RECURSIVE_CREATE : {1303,[]string{"2F003"},"Can't create a %s from within another stored routine"},
	ER_SP_ALREADY_EXISTS : {1304,[]string{"42000"},"%s %s already exists"},
	ER_SP_DOES_NOT_EXIST : {1305,[]string{"42000"},"%s %s does not exist"},
	ER_SP_DROP_FAILED : {1306,[]string{"HY000"},"Failed to DROP %s %s"},
	ER_SP_STORE_FAILED : {1307,[]string{"HY000"},"Failed to CREATE %s %s"},
	ER_SP_LILABEL_MISMATCH : {1308,[]string{"42000"},"%s with no matching label: %s"},
	ER_SP_LABEL_REDEFINE : {1309,[]string{"42000"},"Redefining label %s"},
	ER_SP_LABEL_MISMATCH : {1310,[]string{"42000"},"End-label %s without match"},
	ER_SP_UNINIT_VAR : {1311,[]string{"01000"},"Referring to uninitialized variable %s"},
	ER_SP_BADSELECT : {1312,[]string{"0A000"},"PROCEDURE %s can't return a result set in the given context"},
	ER_SP_BADRETURN : {1313,[]string{"42000"},"RETURN is only allowed in a FUNCTION"},
	ER_SP_BADSTATEMENT : {1314,[]string{"0A000"},"%s is not allowed in stored procedures"},
	ER_UPDATE_LOG_DEPRECATED_IGNORED : {1315,[]string{"42000"},"The update log is deprecated and replaced by the binary log; SET SQL_LOG_UPDATE has been ignored."},
	ER_UPDATE_LOG_DEPRECATED_TRANSLATED : {1316,[]string{"42000"},"The update log is deprecated and replaced by the binary log; SET SQL_LOG_UPDATE has been translated to SET SQL_LOG_BIN."},
	ER_QUERY_INTERRUPTED : {1317,[]string{"70100"},"Query execution was interrupted"},
	ER_SP_WRONG_NO_OF_ARGS : {1318,[]string{"42000"},"Incorrect number of arguments for %s %s; expected %u, got %u"},
	ER_SP_COND_MISMATCH : {1319,[]string{"42000"},"Undefined CONDITION: %s"},
	ER_SP_NORETURN : {1320,[]string{"42000"},"No RETURN found in FUNCTION %s"},
	ER_SP_NORETURNEND : {1321,[]string{"2F005"},"FUNCTION %s ended without RETURN"},
	ER_SP_BAD_CURSOR_QUERY : {1322,[]string{"42000"},"Cursor statement must be a SELECT"},
	ER_SP_BAD_CURSOR_SELECT : {1323,[]string{"42000"},"Cursor SELECT must not have INTO"},
	ER_SP_CURSOR_MISMATCH : {1324,[]string{"42000"},"Undefined CURSOR: %s"},
	ER_SP_CURSOR_ALREADY_OPEN : {1325,[]string{"24000"},"Cursor is already open"},
	ER_SP_CURSOR_NOT_OPEN : {1326,[]string{"24000"},"Cursor is not open"},
	ER_SP_UNDECLARED_VAR : {1327,[]string{"42000"},"Undeclared variable: %s"},
	ER_SP_WRONG_NO_OF_FETCH_ARGS : {1328,[]string{"HY000"},"Incorrect number of FETCH variables"},
	ER_SP_FETCH_NO_DATA : {1329,[]string{"02000"},"No data - zero rows fetched, selected, or processed"},
	ER_SP_DUP_PARAM : {1330,[]string{"42000"},"Duplicate parameter: %s"},
	ER_SP_DUP_VAR : {1331,[]string{"42000"},"Duplicate variable: %s"},
	ER_SP_DUP_COND : {1332,[]string{"42000"},"Duplicate condition: %s"},
	ER_SP_DUP_CURS : {1333,[]string{"42000"},"Duplicate cursor: %s"},
	ER_SP_CANT_ALTER : {1334,[]string{"HY000"},"Failed to ALTER %s %s"},
	ER_SP_SUBSELECT_NYI : {1335,[]string{"0A000"},"Subquery value not supported"},
	ER_STMT_NOT_ALLOWED_IN_SF_OR_TRG : {1336,[]string{"0A000"},"%s is not allowed in stored function or trigger"},
	ER_SP_VARCOND_AFTER_CURSHNDLR : {1337,[]string{"42000"},"Variable or condition declaration after cursor or handler declaration"},
	ER_SP_CURSOR_AFTER_HANDLER : {1338,[]string{"42000"},"Cursor declaration after handler declaration"},
	ER_SP_CASE_NOT_FOUND : {1339,[]string{"20000"},"Case not found for CASE statement"},
	ER_FPARSER_TOO_BIG_FILE : {1340,[]string{"HY000"},"Configuration file '%-.192s' is too big"},
	ER_FPARSER_BAD_HEADER : {1341,[]string{"HY000"},"Malformed file type header in file '%-.192s'"},
	ER_FPARSER_EOF_IN_COMMENT : {1342,[]string{"HY000"},"Unexpected end of file while parsing comment '%-.200s'"},
	ER_FPARSER_ERROR_IN_PARAMETER : {1343,[]string{"HY000"},"Error while parsing parameter '%-.192s' (line: '%-.192s')"},
	ER_FPARSER_EOF_IN_UNKNOWN_PARAMETER : {1344,[]string{"HY000"},"Unexpected end of file while skipping unknown parameter '%-.192s'"},
	ER_VIEW_NO_EXPLAIN : {1345,[]string{"HY000"},"EXPLAIN/SHOW can not be issued; lacking privileges for underlying table"},
	//OBSOLETE_ER_FRM_UNKNOWN_TYPE : {0000,[]string{""},"File '%-.192s' has unknown type '%-.64s' in its header"},
	ER_WRONG_OBJECT : {1347,[]string{"HY000"},"'%-.192s.%-.192s' is not %s"},
	ER_NONUPDATEABLE_COLUMN : {1348,[]string{"HY000"},"Column '%-.192s' is not updatable"},
	//OBSOLETE_ER_VIEW_SELECT_DERIVED_UNUSED : {0000,[]string{""},"View's SELECT contains a subquery in the FROM clause"},
	ER_VIEW_SELECT_CLAUSE : {1350,[]string{"HY000"},"View's SELECT contains a '%s' clause"},
	ER_VIEW_SELECT_VARIABLE : {1351,[]string{"HY000"},"View's SELECT contains a variable or parameter"},
	ER_VIEW_SELECT_TMPTABLE : {1352,[]string{"HY000"},"View's SELECT refers to a temporary table '%-.192s'"},
	ER_VIEW_WRONG_LIST : {1353,[]string{"HY000"},"In definition of view, derived table or common table expression, SELECT list and column names list have different column counts"},
	ER_WARN_VIEW_MERGE : {1354,[]string{"HY000"},"View merge algorithm can't be used here for now (assumed undefined algorithm)"},
	ER_WARN_VIEW_WITHOUT_KEY : {1355,[]string{"HY000"},"View being updated does not have complete key of underlying table in it"},
	ER_VIEW_INVALID : {1356,[]string{"HY000"},"View '%-.192s.%-.192s' references invalid table(s) or column(s) or function(s) or definer/invoker of view lack rights to use them"},
	ER_SP_NO_DROP_SP : {1357,[]string{"HY000"},"Can't drop or alter a %s from within another stored routine"},
	//OBSOLETE_ER_SP_GOTO_IN_HNDLR : {0000,[]string{""},"GOTO is not allowed in a stored procedure handler"},
	ER_TRG_ALREADY_EXISTS : {1359,[]string{"HY000"},"Trigger already exists"},
	ER_TRG_DOES_NOT_EXIST : {1360,[]string{"HY000"},"Trigger does not exist"},
	ER_TRG_ON_VIEW_OR_TEMP_TABLE : {1361,[]string{"HY000"},"Trigger's '%-.192s' is view or temporary table"},
	ER_TRG_CANT_CHANGE_ROW : {1362,[]string{"HY000"},"Updating of %s row is not allowed in %strigger"},
	ER_TRG_NO_SUCH_ROW_IN_TRG : {1363,[]string{"HY000"},"There is no %s row in %s trigger"},
	ER_NO_DEFAULT_FOR_FIELD : {1364,[]string{"HY000"},"Field '%-.192s' doesn't have a default value"},
	ER_DIVISION_BY_ZERO : {1365,[]string{"22012"},"Division by 0"},
	ER_TRUNCATED_WRONG_VALUE_FOR_FIELD : {1366,[]string{"HY000"},"Incorrect %-.32s value: '%-.128s' for column '%.192s' at row %ld"},
	ER_ILLEGAL_VALUE_FOR_TYPE : {1367,[]string{"22007"},"Illegal %s '%-.192s' value found during parsing"},
	ER_VIEW_NONUPD_CHECK : {1368,[]string{"HY000"},"CHECK OPTION on non-updatable view '%-.192s.%-.192s'"},
	ER_VIEW_CHECK_FAILED : {1369,[]string{"HY000"},"CHECK OPTION failed '%-.192s.%-.192s'"},
	ER_PROCACCESS_DENIED_ERROR : {1370,[]string{"42000"},"%-.16s command denied to user '%-.48s'@'%-.64s' for routine '%-.192s'"},
	ER_RELAY_LOG_FAIL : {1371,[]string{"HY000"},"Failed purging old relay logs: %s"},
	//OBSOLETE_ER_PASSWD_LENGTH : {0000,[]string{""},"Password hash should be a %d-digit hexadecimal number"},
	ER_UNKNOWN_TARGET_BINLOG : {1373,[]string{"HY000"},"Target log not found in binlog index"},
	ER_IO_ERR_LOG_INDEX_READ : {1374,[]string{"HY000"},"I/O error reading log index file"},
	ER_BINLOG_PURGE_PROHIBITED : {1375,[]string{"HY000"},"Server configuration does not permit binlog purge"},
	ER_FSEEK_FAIL : {1376,[]string{"HY000"},"Failed on fseek()"},
	ER_BINLOG_PURGE_FATAL_ERR : {1377,[]string{"HY000"},"Fatal error during log purge"},
	ER_LOG_IN_USE : {1378,[]string{"HY000"},"A purgeable log is in use, will not purge"},
	ER_LOG_PURGE_UNKNOWN_ERR : {1379,[]string{"HY000"},"Unknown error during log purge"},
	ER_RELAY_LOG_INIT : {1380,[]string{"HY000"},"Failed initializing relay log position: %s"},
	ER_NO_BINARY_LOGGING : {1381,[]string{"HY000"},"You are not using binary logging"},
	ER_RESERVED_SYNTAX : {1382,[]string{"HY000"},"The '%-.64s' syntax is reserved for purposes internal to the MySQL server"},
	//OBSOLETE_ER_WSAS_FAILED : {0000,[]string{""},"WSAStartup Failed"},
	//OBSOLETE_ER_DIFF_GROUPS_PROC : {0000,[]string{""},"Can't handle procedures with different groups yet"},
	//OBSOLETE_ER_NO_GROUP_FOR_PROC : {0000,[]string{""},"Select must have a group with this procedure"},
	//OBSOLETE_ER_ORDER_WITH_PROC : {0000,[]string{""},"Can't use ORDER clause with this procedure"},
	//OBSOLETE_ER_LOGGING_PROHIBIT_CHANGING_OF : {0000,[]string{""},"Binary logging and replication forbid changing the global server %s"},
	//OBSOLETE_ER_NO_FILE_MAPPING : {0000,[]string{""},"Can't map file: %-.200s, errno: %d"},
	//OBSOLETE_ER_WRONG_MAGIC : {0000,[]string{""},"Wrong magic in %-.64s"},
	ER_PS_MANY_PARAM : {1390,[]string{"HY000"},"Prepared statement contains too many placeholders"},
	ER_KEY_PART_0 : {1391,[]string{"HY000"},"Key part '%-.192s' length cannot be 0"},
	ER_VIEW_CHECKSUM : {1392,[]string{"HY000"},"View text checksum failed"},
	ER_VIEW_MULTIUPDATE : {1393,[]string{"HY000"},"Can not modify more than one base table through a join view '%-.192s.%-.192s'"},
	ER_VIEW_NO_INSERT_FIELD_LIST : {1394,[]string{"HY000"},"Can not insert into join view '%-.192s.%-.192s' without fields list"},
	ER_VIEW_DELETE_MERGE_VIEW : {1395,[]string{"HY000"},"Can not delete from join view '%-.192s.%-.192s'"},
	ER_CANNOT_USER : {1396,[]string{"HY000"},"Operation %s failed for %.256s"},
	ER_XAER_NOTA : {1397,[]string{"XAE04"},"XAER_NOTA: Unknown XID"},
	ER_XAER_INVAL : {1398,[]string{"XAE05"},"XAER_INVAL: Invalid arguments (or unsupported command)"},
	ER_XAER_RMFAIL : {1399,[]string{"XAE07"},"XAER_RMFAIL: The command cannot be executed when global transaction is in the  %.64s state"},
	ER_XAER_OUTSIDE : {1400,[]string{"XAE09"},"XAER_OUTSIDE: Some work is done outside global transaction"},
	ER_XAER_RMERR : {1401,[]string{"XAE03"},"XAER_RMERR: Fatal error occurred in the transaction branch - check your data for consistency"},
	ER_XA_RBROLLBACK : {1402,[]string{"XA100"},"XA_RBROLLBACK: Transaction branch was rolled back"},
	ER_NONEXISTING_PROC_GRANT : {1403,[]string{"42000"},"There is no such grant defined for user '%-.48s' on host '%-.255s' on routine '%-.192s'"},
	ER_PROC_AUTO_GRANT_FAIL : {1404,[]string{"HY000"},"Failed to grant EXECUTE and ALTER ROUTINE privileges"},
	ER_PROC_AUTO_REVOKE_FAIL : {1405,[]string{"HY000"},"Failed to revoke all privileges to dropped routine"},
	ER_DATA_TOO_LONG : {1406,[]string{"22001"},"Data too long for column '%s' at row %ld"},
	ER_SP_BAD_SQLSTATE : {1407,[]string{"42000"},"Bad SQLSTATE: '%s'"},
	ER_STARTUP : {1408,[]string{"HY000"},"%s: ready for connections. Version: '%s'  socket: '%s'  port: %d  %s"},
	ER_LOAD_FROM_FIXED_SIZE_ROWS_TO_VAR : {1409,[]string{"HY000"},"Can't load value from file with fixed size rows to variable"},
	ER_CANT_CREATE_USER_WITH_GRANT : {1410,[]string{"42000"},"You are not allowed to create a user with GRANT"},
	ER_WRONG_VALUE_FOR_TYPE : {1411,[]string{"HY000"},"Incorrect %-.32s value: '%-.128s' for function %-.32s"},
	ER_TABLE_DEF_CHANGED : {1412,[]string{"HY000"},"Table definition has changed, please retry transaction"},
	ER_SP_DUP_HANDLER : {1413,[]string{"42000"},"Duplicate handler declared in the same block"},
	ER_SP_NOT_VAR_ARG : {1414,[]string{"42000"},"OUT or INOUT argument %d for routine %s is not a variable or NEW pseudo-variable in BEFORE trigger"},
	ER_SP_NO_RETSET : {1415,[]string{"0A000"},"Not allowed to return a result set from a %s"},
	ER_CANT_CREATE_GEOMETRY_OBJECT : {1416,[]string{"22003"},"Cannot get geometry object from data you send to the GEOMETRY field"},
	//OBSOLETE_ER_FAILED_ROUTINE_BREAK_BINLOG : {0000,[]string{""},"A routine failed and has neither NO SQL nor READS SQL DATA in its declaration and binary logging is enabled; if non-transactional tables were updated, the binary log will miss their changes"},
	ER_BINLOG_UNSAFE_ROUTINE : {1418,[]string{"HY000"},"This function has none of DETERMINISTIC, NO SQL, or READS SQL DATA in its declaration and binary logging is enabled (you *might* want to use the less safe log_bin_trust_function_creators variable)"},
	ER_BINLOG_CREATE_ROUTINE_NEED_SUPER : {1419,[]string{"HY000"},"You do not have the SUPER privilege and binary logging is enabled (you *might* want to use the less safe log_bin_trust_function_creators variable)"},
	//OBSOLETE_ER_EXEC_STMT_WITH_OPEN_CURSOR : {0000,[]string{""},"You can't execute a prepared statement which has an open cursor associated with it. Reset the statement to re-execute it."},
	ER_STMT_HAS_NO_OPEN_CURSOR : {1421,[]string{"HY000"},"The statement (%lu) has no open cursor."},
	ER_COMMIT_NOT_ALLOWED_IN_SF_OR_TRG : {1422,[]string{"HY000"},"Explicit or implicit commit is not allowed in stored function or trigger."},
	ER_NO_DEFAULT_FOR_VIEW_FIELD : {1423,[]string{"HY000"},"Field of view '%-.192s.%-.192s' underlying table doesn't have a default value"},
	ER_SP_NO_RECURSION : {1424,[]string{"HY000"},"Recursive stored functions and triggers are not allowed."},
	ER_TOO_BIG_SCALE : {1425,[]string{"42000","S1009"},"Too big scale %d specified for column '%-.192s'. Maximum is %lu."},
	ER_TOO_BIG_PRECISION : {1426,[]string{"42000","S1009"},"Too-big precision %d specified for '%-.192s'. Maximum is %lu."},
	ER_M_BIGGER_THAN_D : {1427,[]string{"42000","S1009"},"For float(M,D), double(M,D) or decimal(M,D), M must be >= D (column '%-.192s')."},
	ER_WRONG_LOCK_OF_SYSTEM_TABLE : {1428,[]string{"HY000"},"You can't combine write-locking of system tables with other tables or lock types"},
	ER_CONNECT_TO_FOREIGN_DATA_SOURCE : {1429,[]string{"HY000"},"Unable to connect to foreign data source: %.64s"},
	ER_QUERY_ON_FOREIGN_DATA_SOURCE : {1430,[]string{"HY000"},"There was a problem processing the query on the foreign data source. Data source error: %-.64s"},
	ER_FOREIGN_DATA_SOURCE_DOESNT_EXIST : {1431,[]string{"HY000"},"The foreign data source you are trying to reference does not exist. Data source error:  %-.64s"},
	ER_FOREIGN_DATA_STRING_INVALID_CANT_CREATE : {1432,[]string{"HY000"},"Can't create federated table. The data source connection string '%-.64s' is not in the correct format"},
	ER_FOREIGN_DATA_STRING_INVALID : {1433,[]string{"HY000"},"The data source connection string '%-.64s' is not in the correct format"},
	//OBSOLETE_ER_CANT_CREATE_FEDERATED_TABLE : {0000,[]string{""},"Can't create federated table. Foreign data src error:  %-.64s"},
	ER_TRG_IN_WRONG_SCHEMA : {1435,[]string{"HY000"},"Trigger in wrong schema"},
	ER_STACK_OVERRUN_NEED_MORE : {1436,[]string{"HY000"},"Thread stack overrun:  %ld bytes used of a %ld byte stack, and %ld bytes needed.  Use 'mysqld --thread_stack=#' to specify a bigger stack."},
	ER_TOO_LONG_BODY : {1437,[]string{"42000","S1009"},"Routine body for '%-.100s' is too long"},
	ER_WARN_CANT_DROP_DEFAULT_KEYCACHE : {1438,[]string{"HY000"},"Cannot drop default keycache"},
	ER_TOO_BIG_DISPLAYWIDTH : {1439,[]string{"42000","S1009"},"Display width out of range for column '%-.192s' (max = %lu)"},
	ER_XAER_DUPID : {1440,[]string{"XAE08"},"XAER_DUPID: The XID already exists"},
	ER_DATETIME_FUNCTION_OVERFLOW : {1441,[]string{"22008"},"Datetime function: %-.32s field overflow"},
	ER_CANT_UPDATE_USED_TABLE_IN_SF_OR_TRG : {1442,[]string{"HY000"},"Can't update table '%-.192s' in stored function/trigger because it is already used by statement which invoked this stored function/trigger."},
	ER_VIEW_PREVENT_UPDATE : {1443,[]string{"HY000"},"The definition of table '%-.192s' prevents operation %.192s on table '%-.192s'."},
	ER_PS_NO_RECURSION : {1444,[]string{"HY000"},"The prepared statement contains a stored routine call that refers to that same statement. It's not allowed to execute a prepared statement in such a recursive manner"},
	ER_SP_CANT_SET_AUTOCOMMIT : {1445,[]string{"HY000"},"Not allowed to set autocommit from a stored function or trigger"},
	//OBSOLETE_ER_MALFORMED_DEFINER : {0000,[]string{""},"Definer is not fully qualified"},
	ER_VIEW_FRM_NO_USER : {1447,[]string{"HY000"},"View '%-.192s'.'%-.192s' has no definer information (old table format). Current user is used as definer. Please recreate the view!"},
	ER_VIEW_OTHER_USER : {1448,[]string{"HY000"},"You need the SUPER privilege for creation view with '%-.192s'@'%-.192s' definer"},
	ER_NO_SUCH_USER : {1449,[]string{"HY000"},"The user specified as a definer ('%-.64s'@'%-.64s') does not exist"},
	ER_FORBID_SCHEMA_CHANGE : {1450,[]string{"HY000"},"Changing schema from '%-.192s' to '%-.192s' is not allowed."},
	ER_ROW_IS_REFERENCED_2 : {1451,[]string{"23000"},"Cannot delete or update a parent row: a foreign key constraint fails (%.192s)"},
	ER_NO_REFERENCED_ROW_2 : {1452,[]string{"23000"},"Cannot add or update a child row: a foreign key constraint fails (%.192s)"},
	ER_SP_BAD_VAR_SHADOW : {1453,[]string{"42000"},"Variable '%-.64s' must be quoted with `...`, or renamed"},
	ER_TRG_NO_DEFINER : {1454,[]string{"HY000"},"No definer attribute for trigger '%-.192s'.'%-.192s'. It's disallowed to create trigger without definer."},
	ER_OLD_FILE_FORMAT : {1455,[]string{"HY000"},"'%-.192s' has an old format, you should re-create the '%s' object(s)"},
	ER_SP_RECURSION_LIMIT : {1456,[]string{"HY000"},"Recursive limit %d (as set by the max_sp_recursion_depth variable) was exceeded for routine %.192s"},
	//OBSOLETE_ER_SP_PROC_TABLE_CORRUPT : {0000,[]string{""},"Failed to load routine %-.192s. The data dictionary table missing, corrupt, or contains bad data (internal code %d)"},
	ER_SP_WRONG_NAME : {1458,[]string{"42000"},"Incorrect routine name '%-.192s'"},
	ER_TABLE_NEEDS_UPGRADE : {1459,[]string{"HY000"},"Table upgrade required. Please do \"REPAIR TABLE `%-.64s`\" or dump/reload to fix it!"},
	ER_SP_NO_AGGREGATE : {1460,[]string{"42000"},"AGGREGATE is not supported for stored functions"},
	ER_MAX_PREPARED_STMT_COUNT_REACHED : {1461,[]string{"42000"},"Can't create more than max_prepared_stmt_count statements (current value: %lu)"},
	ER_VIEW_RECURSIVE : {1462,[]string{"HY000"},"`%-.192s`.`%-.192s` contains view recursion"},
	ER_NON_GROUPING_FIELD_USED : {1463,[]string{"42000"},"Non-grouping field '%-.192s' is used in %-.64s clause"},
	ER_TABLE_CANT_HANDLE_SPKEYS : {1464,[]string{"HY000"},"The used table type doesn't support SPATIAL indexes"},
	ER_NO_TRIGGERS_ON_SYSTEM_SCHEMA : {1465,[]string{"HY000"},"Triggers can not be created on system tables"},
	ER_REMOVED_SPACES : {1466,[]string{"HY000"},"Leading spaces are removed from name '%s'"},
	ER_AUTOINC_READ_FAILED : {1467,[]string{"HY000"},"Failed to read auto-increment value from storage engine"},
	ER_USERNAME : {1468,[]string{"HY000"},"user name"},
	ER_HOSTNAME : {1469,[]string{"HY000"},"host name"},
	ER_WRONG_STRING_LENGTH : {1470,[]string{"HY000"},"String '%-.70s' is too long for %s (should be no longer than %d)"},
	ER_NON_INSERTABLE_TABLE : {1471,[]string{"HY000"},"The target table %-.100s of the %s is not insertable-into"},
	ER_ADMIN_WRONG_MRG_TABLE : {1472,[]string{"HY000"},"Table '%-.64s' is differently defined or of non-MyISAM type or doesn't exist"},
	ER_TOO_HIGH_LEVEL_OF_NESTING_FOR_SELECT : {1473,[]string{"HY000"},"Too high level of nesting for select"},
	ER_NAME_BECOMES_EMPTY : {1474,[]string{"HY000"},"Name '%-.64s' has become ''"},
	ER_AMBIGUOUS_FIELD_TERM : {1475,[]string{"HY000"},"First character of the FIELDS TERMINATED string is ambiguous; please use non-optional and non-empty FIELDS ENCLOSED BY"},
	ER_FOREIGN_SERVER_EXISTS : {1476,[]string{"HY000"},"The foreign server, %s, you are trying to create already exists."},
	ER_FOREIGN_SERVER_DOESNT_EXIST : {1477,[]string{"HY000"},"The foreign server name you are trying to reference does not exist. Data source error:  %-.64s"},
	ER_ILLEGAL_HA_CREATE_OPTION : {1478,[]string{"HY000"},"Table storage engine '%-.64s' does not support the create option '%.64s'"},
	ER_PARTITION_REQUIRES_VALUES_ERROR : {1479,[]string{"HY000"},"Syntax error: %-.64s PARTITIONING requires definition of VALUES %-.64s for each partition"},
	ER_PARTITION_WRONG_VALUES_ERROR : {1480,[]string{"HY000"},"Only %-.64s PARTITIONING can use VALUES %-.64s in partition definition"},
	ER_PARTITION_MAXVALUE_ERROR : {1481,[]string{"HY000"},"MAXVALUE can only be used in last partition definition"},
	//OBSOLETE_ER_PARTITION_SUBPARTITION_ERROR : {0000,[]string{""},"Subpartitions can only be hash partitions and by key"},
	//OBSOLETE_ER_PARTITION_SUBPART_MIX_ERROR : {0000,[]string{""},"Must define subpartitions on all partitions if on one partition"},
	ER_PARTITION_WRONG_NO_PART_ERROR : {1484,[]string{"HY000"},"Wrong number of partitions defined, mismatch with previous setting"},
	ER_PARTITION_WRONG_NO_SUBPART_ERROR : {1485,[]string{"HY000"},"Wrong number of subpartitions defined, mismatch with previous setting"},
	ER_WRONG_EXPR_IN_PARTITION_FUNC_ERROR : {1486,[]string{"HY000"},"Constant, random or timezone-dependent expressions in (sub)partitioning function are not allowed"},
	//OBSOLETE_ER_NO_CONST_EXPR_IN_RANGE_OR_LIST_ERROR : {0000,[]string{""},"Expression in RANGE/LIST VALUES must be constant"},
	ER_FIELD_NOT_FOUND_PART_ERROR : {1488,[]string{"HY000"},"Field in list of fields for partition function not found in table"},
	//OBSOLETE_ER_LIST_OF_FIELDS_ONLY_IN_HASH_ERROR : {0000,[]string{""},"List of fields is only allowed in KEY partitions"},
	ER_INCONSISTENT_PARTITION_INFO_ERROR : {1490,[]string{"HY000"},"The partition info in the frm file is not consistent with what can be written into the frm file"},
	ER_PARTITION_FUNC_NOT_ALLOWED_ERROR : {1491,[]string{"HY000"},"The %-.192s function returns the wrong type"},
	ER_PARTITIONS_MUST_BE_DEFINED_ERROR : {1492,[]string{"HY000"},"For %-.64s partitions each partition must be defined"},
	ER_RANGE_NOT_INCREASING_ERROR : {1493,[]string{"HY000"},"VALUES LESS THAN value must be strictly increasing for each partition"},
	ER_INCONSISTENT_TYPE_OF_FUNCTIONS_ERROR : {1494,[]string{"HY000"},"VALUES value must be of same type as partition function"},
	ER_MULTIPLE_DEF_CONST_IN_LIST_PART_ERROR : {1495,[]string{"HY000"},"Multiple definition of same constant in list partitioning"},
	ER_PARTITION_ENTRY_ERROR : {1496,[]string{"HY000"},"Partitioning can not be used stand-alone in query"},
	ER_MIX_HANDLER_ERROR : {1497,[]string{"HY000"},"The mix of handlers in the partitions is not allowed in this version of MySQL"},
	ER_PARTITION_NOT_DEFINED_ERROR : {1498,[]string{"HY000"},"For the partitioned engine it is necessary to define all %-.64s"},
	ER_TOO_MANY_PARTITIONS_ERROR : {1499,[]string{"HY000"},"Too many partitions (including subpartitions) were defined"},
	ER_SUBPARTITION_ERROR : {1500,[]string{"HY000"},"It is only possible to mix RANGE/LIST partitioning with HASH/KEY partitioning for subpartitioning"},
	ER_CANT_CREATE_HANDLER_FILE : {1501,[]string{"HY000"},"Failed to create specific handler file"},
	ER_BLOB_FIELD_IN_PART_FUNC_ERROR : {1502,[]string{"HY000"},"A BLOB field is not allowed in partition function"},
	ER_UNIQUE_KEY_NEED_ALL_FIELDS_IN_PF : {1503,[]string{"HY000"},"A %-.192s must include all columns in the table's partitioning function (prefixed columns are not considered)."},
	ER_NO_PARTS_ERROR : {1504,[]string{"HY000"},"Number of %-.64s = 0 is not an allowed value"},
	ER_PARTITION_MGMT_ON_NONPARTITIONED : {1505,[]string{"HY000"},"Partition management on a not partitioned table is not possible"},
	ER_FOREIGN_KEY_ON_PARTITIONED : {1506,[]string{"HY000"},"Foreign keys are not yet supported in conjunction with partitioning"},
	ER_DROP_PARTITION_NON_EXISTENT : {1507,[]string{"HY000"},"Error in list of partitions to %-.64s"},
	ER_DROP_LAST_PARTITION : {1508,[]string{"HY000"},"Cannot remove all partitions, use DROP TABLE instead"},
	ER_COALESCE_ONLY_ON_HASH_PARTITION : {1509,[]string{"HY000"},"COALESCE PARTITION can only be used on HASH/KEY partitions"},
	ER_REORG_HASH_ONLY_ON_SAME_NO : {1510,[]string{"HY000"},"REORGANIZE PARTITION can only be used to reorganize partitions not to change their numbers"},
	ER_REORG_NO_PARAM_ERROR : {1511,[]string{"HY000"},"REORGANIZE PARTITION without parameters can only be used on auto-partitioned tables using HASH PARTITIONs"},
	ER_ONLY_ON_RANGE_LIST_PARTITION : {1512,[]string{"HY000"},"%-.64s PARTITION can only be used on RANGE/LIST partitions"},
	ER_ADD_PARTITION_SUBPART_ERROR : {1513,[]string{"HY000"},"Trying to Add partition(s) with wrong number of subpartitions"},
	ER_ADD_PARTITION_NO_NEW_PARTITION : {1514,[]string{"HY000"},"At least one partition must be added"},
	ER_COALESCE_PARTITION_NO_PARTITION : {1515,[]string{"HY000"},"At least one partition must be coalesced"},
	ER_REORG_PARTITION_NOT_EXIST : {1516,[]string{"HY000"},"More partitions to reorganize than there are partitions"},
	ER_SAME_NAME_PARTITION : {1517,[]string{"HY000"},"Duplicate partition name %-.192s"},
	ER_NO_BINLOG_ERROR : {1518,[]string{"HY000"},"It is not allowed to shut off binlog on this command"},
	ER_CONSECUTIVE_REORG_PARTITIONS : {1519,[]string{"HY000"},"When reorganizing a set of partitions they must be in consecutive order"},
	ER_REORG_OUTSIDE_RANGE : {1520,[]string{"HY000"},"Reorganize of range partitions cannot change total ranges except for last partition where it can extend the range"},
	ER_PARTITION_FUNCTION_FAILURE : {1521,[]string{"HY000"},"Partition function not supported in this version for this handler"},
	//OBSOLETE_ER_PART_STATE_ERROR : {0000,[]string{""},"Partition state cannot be defined from CREATE/ALTER TABLE"},
	ER_LIMITED_PART_RANGE : {1523,[]string{"HY000"},"The %-.64s handler only supports 32 bit integers in VALUES"},
	ER_PLUGIN_IS_NOT_LOADED : {1524,[]string{"HY000"},"Plugin '%-.192s' is not loaded"},
	ER_WRONG_VALUE : {1525,[]string{"HY000"},"Incorrect %-.32s value: '%-.128s'"},
	ER_NO_PARTITION_FOR_GIVEN_VALUE : {1526,[]string{"HY000"},"Table has no partition for value %-.64s"},
	ER_FILEGROUP_OPTION_ONLY_ONCE : {1527,[]string{"HY000"},"It is not allowed to specify %s more than once"},
	ER_CREATE_FILEGROUP_FAILED : {1528,[]string{"HY000"},"Failed to create %s"},
	ER_DROP_FILEGROUP_FAILED : {1529,[]string{"HY000"},"Failed to drop %s"},
	ER_TABLESPACE_AUTO_EXTEND_ERROR : {1530,[]string{"HY000"},"The handler doesn't support autoextend of tablespaces"},
	ER_WRONG_SIZE_NUMBER : {1531,[]string{"HY000"},"A size parameter was incorrectly specified, either number or on the form 10M"},
	ER_SIZE_OVERFLOW_ERROR : {1532,[]string{"HY000"},"The size number was correct but we don't allow the digit part to be more than 2 billion"},
	ER_ALTER_FILEGROUP_FAILED : {1533,[]string{"HY000"},"Failed to alter: %s"},
	ER_BINLOG_ROW_LOGGING_FAILED : {1534,[]string{"HY000"},"Writing one row to the row-based binary log failed"},
	//OBSOLETE_ER_BINLOG_ROW_WRONG_TABLE_DEF : {0000,[]string{""},"Table definition on master and slave does not match: %s"},
	//OBSOLETE_ER_BINLOG_ROW_RBR_TO_SBR : {0000,[]string{""},"Slave running with --log-slave-updates must use row-based binary logging to be able to replicate row-based binary log events"},
	ER_EVENT_ALREADY_EXISTS : {1537,[]string{"HY000"},"Event '%-.192s' already exists"},
	//OBSOLETE_ER_EVENT_STORE_FAILED : {0000,[]string{""},"Failed to store event %s. Error code %d from storage engine."},
	ER_EVENT_DOES_NOT_EXIST : {1539,[]string{"HY000"},"Unknown event '%-.192s'"},
	//OBSOLETE_ER_EVENT_CANT_ALTER : {0000,[]string{""},"Failed to alter event '%-.192s'"},
	//OBSOLETE_ER_EVENT_DROP_FAILED : {0000,[]string{""},"Failed to drop %s"},
	ER_EVENT_INTERVAL_NOT_POSITIVE_OR_TOO_BIG : {1542,[]string{"HY000"},"INTERVAL is either not positive or too big"},
	ER_EVENT_ENDS_BEFORE_STARTS : {1543,[]string{"HY000"},"ENDS is either invalid or before STARTS"},
	ER_EVENT_EXEC_TIME_IN_THE_PAST : {1544,[]string{"HY000"},"Event execution time is in the past. Event has been disabled"},
	//OBSOLETE_ER_EVENT_OPEN_TABLE_FAILED : {0000,[]string{""},"Failed to open mysql.event"},
	//OBSOLETE_ER_EVENT_NEITHER_M_EXPR_NOR_M_AT : {0000,[]string{""},"No datetime expression provided"},
	//OBSOLETE_ER_COL_COUNT_DOESNT_MATCH_CORRUPTED : {0000,[]string{""},"Column count of mysql.%s is wrong. Expected %d, found %d. The table is probably corrupted"},
	//OBSOLETE_ER_CANNOT_LOAD_FROM_TABLE : {0000,[]string{""},"Cannot load from mysql.%s. The table is probably corrupted"},
	//OBSOLETE_ER_EVENT_CANNOT_DELETE : {0000,[]string{""},"Failed to delete the event from mysql.event"},
	//OBSOLETE_ER_EVENT_COMPILE_ERROR : {0000,[]string{""},"Error during compilation of event's body"},
	ER_EVENT_SAME_NAME : {1551,[]string{"HY000"},"Same old and new event name"},
	//OBSOLETE_ER_EVENT_DATA_TOO_LONG : {0000,[]string{""},"Data for column '%s' too long"},
	ER_DROP_INDEX_FK : {1553,[]string{"HY000"},"Cannot drop index '%-.192s': needed in a foreign key constraint"},
	ER_WARN_DEPRECATED_SYNTAX_WITH_VER : {1554,[]string{"HY000"},"The syntax '%s' is deprecated and will be removed in MySQL %s. Please use %s instead"},
	//OBSOLETE_ER_CANT_WRITE_LOCK_LOG_TABLE : {0000,[]string{""},"You can't write-lock a log table. Only read access is possible"},
	ER_CANT_LOCK_LOG_TABLE : {1556,[]string{"HY000"},"You can't use locks with log tables."},
	ER_FOREIGN_DUPLICATE_KEY_OLD_UNUSED : {1557,[]string{"23000","S1009"},"Upholding foreign key constraints for table '%.192s', entry '%-.192s', key %d would lead to a duplicate entry"},
	ER_COL_COUNT_DOESNT_MATCH_PLEASE_UPDATE : {1558,[]string{"HY000"},"The column count of mysql.%s is wrong. Expected %d, found %d. Created with MySQL %d, now running %d. Please perform the MySQL upgrade procedure."},
	//OBSOLETE_ER_TEMP_TABLE_PREVENTS_SWITCH_OUT_OF_RBR : {1559,[]string{"HY000"},"Cannot switch out of the row-based binary log format when the session has open temporary tables"},
	ER_STORED_FUNCTION_PREVENTS_SWITCH_BINLOG_FORMAT : {1560,[]string{"HY000"},"Cannot change the binary logging format inside a stored function or trigger"},
	//OBSOLETE_ER_NDB_CANT_SWITCH_BINLOG_FORMAT : {0000,[]string{""},"The NDB cluster engine does not support changing the binlog format on the fly yet"},
	ER_PARTITION_NO_TEMPORARY : {1562,[]string{"HY000"},"Cannot create temporary table with partitions"},
	ER_PARTITION_CONST_DOMAIN_ERROR : {1563,[]string{"HY000"},"Partition constant is out of partition function domain"},
	ER_PARTITION_FUNCTION_IS_NOT_ALLOWED : {1564,[]string{"HY000"},"This partition function is not allowed"},
	//OBSOLETE_ER_DDL_LOG_ERROR_UNUSED : {0000,[]string{""},"Error in DDL log"},
	ER_NULL_IN_VALUES_LESS_THAN : {1566,[]string{"HY000"},"Not allowed to use NULL value in VALUES LESS THAN"},
	ER_WRONG_PARTITION_NAME : {1567,[]string{"HY000"},"Incorrect partition name"},
	ER_CANT_CHANGE_TX_CHARACTERISTICS : {1568,[]string{"25001"},"Transaction characteristics can't be changed while a transaction is in progress"},
	ER_DUP_ENTRY_AUTOINCREMENT_CASE : {1569,[]string{"HY000"},"ALTER TABLE causes auto_increment resequencing, resulting in duplicate entry '%-.192s' for key '%-.192s'"},
	//OBSOLETE_ER_EVENT_MODIFY_QUEUE_ERROR : {0000,[]string{""},"Internal scheduler error %d"},
	ER_EVENT_SET_VAR_ERROR : {1571,[]string{"HY000"},"Error during starting/stopping of the scheduler. Error code %u"},
	ER_PARTITION_MERGE_ERROR : {1572,[]string{"HY000"},"Engine cannot be used in partitioned tables"},
	//OBSOLETE_ER_CANT_ACTIVATE_LOG : {0000,[]string{""},"Cannot activate '%-.64s' log"},
	//OBSOLETE_ER_RBR_NOT_AVAILABLE : {0000,[]string{""},"The server was not built with row-based replication"},
	ER_BASE64_DECODE_ERROR : {1575,[]string{"HY000"},"Decoding of base64 string failed"},
	ER_EVENT_RECURSION_FORBIDDEN : {1576,[]string{"HY000"},"Recursion of EVENT DDL statements is forbidden when body is present"},
	//OBSOLETE_ER_EVENTS_DB_ERROR : {0000,[]string{""},"Cannot proceed because system tables used by Event Scheduler were found damaged at server start"},
	ER_ONLY_INTEGERS_ALLOWED : {1578,[]string{"HY000"},"Only integers allowed as number here"},
	ER_UNSUPORTED_LOG_ENGINE : {1579,[]string{"HY000"},"This storage engine cannot be used for log tables"},
	ER_BAD_LOG_STATEMENT : {1580,[]string{"HY000"},"You cannot '%s' a log table if logging is enabled"},
	ER_CANT_RENAME_LOG_TABLE : {1581,[]string{"HY000"},"Cannot rename '%s'. When logging enabled, rename to/from log table must rename two tables: the log table to an archive table and another table back to '%s'"},
	ER_WRONG_PARAMCOUNT_TO_NATIVE_FCT : {1582,[]string{"42000"},"Incorrect parameter count in the call to native function '%-.192s'"},
	ER_WRONG_PARAMETERS_TO_NATIVE_FCT : {1583,[]string{"42000"},"Incorrect parameters in the call to native function '%-.192s'"},
	ER_WRONG_PARAMETERS_TO_STORED_FCT : {1584,[]string{"42000"},"Incorrect parameters in the call to stored function %-.192s"},
	ER_NATIVE_FCT_NAME_COLLISION : {1585,[]string{"HY000"},"This function '%-.192s' has the same name as a native function"},
	ER_DUP_ENTRY_WITH_KEY_NAME : {1586,[]string{"23000","S1009"},"Duplicate entry '%-.64s' for key '%-.385s'"},
	ER_BINLOG_PURGE_EMFILE : {1587,[]string{"HY000"},"Too many files opened, please execute the command again"},
	ER_EVENT_CANNOT_CREATE_IN_THE_PAST : {1588,[]string{"HY000"},"Event execution time is in the past and ON COMPLETION NOT PRESERVE is set. The event was dropped immediately after creation."},
	ER_EVENT_CANNOT_ALTER_IN_THE_PAST : {1589,[]string{"HY000"},"Event execution time is in the past and ON COMPLETION NOT PRESERVE is set. The event was not changed. Specify a time in the future."},
	//OBSOLETE_ER_SLAVE_INCIDENT : {13119,[]string{"HY000"},"The incident %s occurred on the master. Message: %s"},
	ER_NO_PARTITION_FOR_GIVEN_VALUE_SILENT : {1591,[]string{"HY000"},"Table has no partition for some existing values"},
	ER_BINLOG_UNSAFE_STATEMENT : {1592,[]string{"HY000"},"Unsafe statement written to the binary log using statement format since BINLOG_FORMAT = STATEMENT. %s"},
	ER_BINLOG_FATAL_ERROR : {1593,[]string{"HY000"},"Fatal error: %s"},
	//OBSOLETE_ER_SLAVE_RELAY_LOG_READ_FAILURE : {13121,[]string{"HY000"},"Relay log read failure: %s"},
	//OBSOLETE_ER_SLAVE_RELAY_LOG_WRITE_FAILURE : {13122,[]string{"HY000"},"Relay log write failure: %s"},
	//OBSOLETE_ER_SLAVE_CREATE_EVENT_FAILURE : {13116,[]string{"HY000"},"Failed to create %s"},
	//OBSOLETE_ER_SLAVE_MASTER_COM_FAILURE : {13120,[]string{"HY000"},"Master command %s failed: %s"},
	ER_BINLOG_LOGGING_IMPOSSIBLE : {1598,[]string{"HY000"},"Binary logging not possible. Message: %s"},
	ER_VIEW_NO_CREATION_CTX : {1599,[]string{"HY000"},"View `%-.64s`.`%-.64s` has no creation context"},
	ER_VIEW_INVALID_CREATION_CTX : {1600,[]string{"HY000"},"Creation context of view `%-.64s`.`%-.64s' is invalid"},
	//OBSOLETE_ER_SR_INVALID_CREATION_CTX : {0000,[]string{""},"Creation context of stored routine `%-.64s`.`%-.64s` is invalid"},
	ER_TRG_CORRUPTED_FILE : {1602,[]string{"HY000"},"Corrupted TRG file for table `%-.64s`.`%-.64s`"},
	ER_TRG_NO_CREATION_CTX : {1603,[]string{"HY000"},"Triggers for table `%-.64s`.`%-.64s` have no creation context"},
	ER_TRG_INVALID_CREATION_CTX : {1604,[]string{"HY000"},"Trigger creation context of table `%-.64s`.`%-.64s` is invalid"},
	ER_EVENT_INVALID_CREATION_CTX : {1605,[]string{"HY000"},"Creation context of event `%-.64s`.`%-.64s` is invalid"},
	ER_TRG_CANT_OPEN_TABLE : {1606,[]string{"HY000"},"Cannot open table for trigger `%-.64s`.`%-.64s`"},
	//OBSOLETE_ER_CANT_CREATE_SROUTINE : {0000,[]string{""},"Cannot create stored routine `%-.64s`. Check warnings"},
	//OBSOLETE_ER_NEVER_USED : {0000,[]string{""},"Ambiguous slave modes combination. %s"},
	ER_NO_FORMAT_DESCRIPTION_EVENT_BEFORE_BINLOG_STATEMENT : {1609,[]string{"HY000"},"The BINLOG statement of type `%s` was not preceded by a format description BINLOG statement."},
	ER_SLAVE_CORRUPT_EVENT : {1610,[]string{"HY000"},"Corrupted replication event was detected"},
	//OBSOLETE_ER_LOAD_DATA_INVALID_COLUMN_UNUSED : {0000,[]string{""},"Invalid column reference (%-.64s) in LOAD DATA"},
	ER_LOG_PURGE_NO_FILE : {1612,[]string{"HY000"},"Being purged log %s was not found"},
	ER_XA_RBTIMEOUT : {1613,[]string{"XA106"},"XA_RBTIMEOUT: Transaction branch was rolled back: took too long"},
	ER_XA_RBDEADLOCK : {1614,[]string{"XA102"},"XA_RBDEADLOCK: Transaction branch was rolled back: deadlock was detected"},
	ER_NEED_REPREPARE : {1615,[]string{"HY000"},"Prepared statement needs to be re-prepared"},
	//OBSOLETE_ER_DELAYED_NOT_SUPPORTED : {0000,[]string{""},"DELAYED option not supported for table '%-.192s'"},
	WARN_NO_MASTER_INFO : {1617,[]string{"HY000"},"The master info structure does not exist"},
	WARN_OPTION_IGNORED : {1618,[]string{"HY000"},"<%-.64s> option ignored"},
	ER_PLUGIN_DELETE_BUILTIN : {1619,[]string{"HY000"},"Built-in plugins cannot be deleted"},
	WARN_PLUGIN_BUSY : {1620,[]string{"HY000"},"Plugin is busy and will be uninstalled on shutdown"},
	ER_VARIABLE_IS_READONLY : {1621,[]string{"HY000"},"%s variable '%s' is read-only. Use SET %s to assign the value"},
	ER_WARN_ENGINE_TRANSACTION_ROLLBACK : {1622,[]string{"HY000"},"Storage engine %s does not support rollback for this statement. Transaction rolled back and must be restarted"},
	//OBSOLETE_ER_SLAVE_HEARTBEAT_FAILURE : {13118,[]string{"HY000"},"Unexpected master's heartbeat data: %s"},
	ER_SLAVE_HEARTBEAT_VALUE_OUT_OF_RANGE : {1624,[]string{"HY000"},"The requested value for the heartbeat period is either negative or exceeds the maximum allowed (%s seconds)."},
	ER_NDB_REPLICATION_SCHEMA_ERROR : {1625,[]string{"HY000"},"Bad schema for mysql.ndb_replication table. Message: %-.64s"},
	ER_CONFLICT_FN_PARSE_ERROR : {1626,[]string{"HY000"},"Error in parsing conflict function. Message: %-.64s"},
	ER_EXCEPTIONS_WRITE_ERROR : {1627,[]string{"HY000"},"Write to exceptions table failed. Message: %-.128s"},
	ER_TOO_LONG_TABLE_COMMENT : {1628,[]string{"HY000"},"Comment for table '%-.64s' is too long (max = %lu)"},
	ER_TOO_LONG_FIELD_COMMENT : {1629,[]string{"HY000"},"Comment for field '%-.64s' is too long (max = %lu)"},
	ER_FUNC_INEXISTENT_NAME_COLLISION : {1630,[]string{"42000"},"FUNCTION %s does not exist. Check the 'Function Name Parsing and Resolution' section in the Reference Manual"},
	ER_DATABASE_NAME : {1631,[]string{"HY000"},"Database"},
	ER_TABLE_NAME : {1632,[]string{"HY000"},"Table"},
	ER_PARTITION_NAME : {1633,[]string{"HY000"},"Partition"},
	ER_SUBPARTITION_NAME : {1634,[]string{"HY000"},"Subpartition"},
	ER_TEMPORARY_NAME : {1635,[]string{"HY000"},"Temporary"},
	ER_RENAMED_NAME : {1636,[]string{"HY000"},"Renamed"},
	ER_TOO_MANY_CONCURRENT_TRXS : {1637,[]string{"HY000"},"Too many active concurrent transactions"},
	WARN_NON_ASCII_SEPARATOR_NOT_IMPLEMENTED : {1638,[]string{"HY000"},"Non-ASCII separator arguments are not fully supported"},
	ER_DEBUG_SYNC_TIMEOUT : {1639,[]string{"HY000"},"debug sync point wait timed out"},
	ER_DEBUG_SYNC_HIT_LIMIT : {1640,[]string{"HY000"},"debug sync point hit limit reached"},
	ER_DUP_SIGNAL_SET : {1641,[]string{"42000"},"Duplicate condition information item '%s'"},
	ER_SIGNAL_WARN : {1642,[]string{"01000"},"Unhandled user-defined warning condition"},
	ER_SIGNAL_NOT_FOUND : {1643,[]string{"02000"},"Unhandled user-defined not found condition"},
	ER_SIGNAL_EXCEPTION : {1644,[]string{"HY000"},"Unhandled user-defined exception condition"},
	ER_RESIGNAL_WITHOUT_ACTIVE_HANDLER : {1645,[]string{"0K000"},"RESIGNAL when handler not active"},
	ER_SIGNAL_BAD_CONDITION_TYPE : {1646,[]string{"HY000"},"SIGNAL/RESIGNAL can only use a CONDITION defined with SQLSTATE"},
	WARN_COND_ITEM_TRUNCATED : {1647,[]string{"HY000"},"Data truncated for condition item '%s'"},
	ER_COND_ITEM_TOO_LONG : {1648,[]string{"HY000"},"Data too long for condition item '%s'"},
	ER_UNKNOWN_LOCALE : {1649,[]string{"HY000"},"Unknown locale: '%-.64s'"},
	ER_SLAVE_IGNORE_SERVER_IDS : {1650,[]string{"HY000"},"The requested server id %d clashes with the slave startup option --replicate-same-server-id"},
	//OBSOLETE_ER_QUERY_CACHE_DISABLED : {1651,[]string{"HY000"},"Query cache is disabled; restart the server with query_cache_type=1 to enable it"},
	ER_SAME_NAME_PARTITION_FIELD : {1652,[]string{"HY000"},"Duplicate partition field name '%-.192s'"},
	ER_PARTITION_COLUMN_LIST_ERROR : {1653,[]string{"HY000"},"Inconsistency in usage of column lists for partitioning"},
	ER_WRONG_TYPE_COLUMN_VALUE_ERROR : {1654,[]string{"HY000"},"Partition column values of incorrect type"},
	ER_TOO_MANY_PARTITION_FUNC_FIELDS_ERROR : {1655,[]string{"HY000"},"Too many fields in '%-.192s'"},
	ER_MAXVALUE_IN_VALUES_IN : {1656,[]string{"HY000"},"Cannot use MAXVALUE as value in VALUES IN"},
	ER_TOO_MANY_VALUES_ERROR : {1657,[]string{"HY000"},"Cannot have more than one value for this type of %-.64s partitioning"},
	ER_ROW_SINGLE_PARTITION_FIELD_ERROR : {1658,[]string{"HY000"},"Row expressions in VALUES IN only allowed for multi-field column partitioning"},
	ER_FIELD_TYPE_NOT_ALLOWED_AS_PARTITION_FIELD : {1659,[]string{"HY000"},"Field '%-.192s' is of a not allowed type for this type of partitioning"},
	ER_PARTITION_FIELDS_TOO_LONG : {1660,[]string{"HY000"},"The total length of the partitioning fields is too large"},
	ER_BINLOG_ROW_ENGINE_AND_STMT_ENGINE : {1661,[]string{"HY000"},"Cannot execute statement: impossible to write to binary log since both row-incapable engines and statement-incapable engines are involved."},
	ER_BINLOG_ROW_MODE_AND_STMT_ENGINE : {1662,[]string{"HY000"},"Cannot execute statement: impossible to write to binary log since BINLOG_FORMAT = ROW and at least one table uses a storage engine limited to statement-based logging."},
	ER_BINLOG_UNSAFE_AND_STMT_ENGINE : {1663,[]string{"HY000"},"Cannot execute statement: impossible to write to binary log since statement is unsafe, storage engine is limited to statement-based logging, and BINLOG_FORMAT = MIXED. %s"},
	ER_BINLOG_ROW_INJECTION_AND_STMT_ENGINE : {1664,[]string{"HY000"},"Cannot execute statement: impossible to write to binary log since statement is in row format and at least one table uses a storage engine limited to statement-based logging."},
	ER_BINLOG_STMT_MODE_AND_ROW_ENGINE : {1665,[]string{"HY000"},"Cannot execute statement: impossible to write to binary log since BINLOG_FORMAT = STATEMENT and at least one table uses a storage engine limited to row-based logging.%s"},
	ER_BINLOG_ROW_INJECTION_AND_STMT_MODE : {1666,[]string{"HY000"},"Cannot execute statement: impossible to write to binary log since statement is in row format and BINLOG_FORMAT = STATEMENT."},
	ER_BINLOG_MULTIPLE_ENGINES_AND_SELF_LOGGING_ENGINE : {1667,[]string{"HY000"},"Cannot execute statement: impossible to write to binary log since more than one engine is involved and at least one engine is self-logging."},
	ER_BINLOG_UNSAFE_LIMIT : {1668,[]string{"HY000"},"The statement is unsafe because it uses a LIMIT clause. This is unsafe because the set of rows included cannot be predicted."},
	//OBSOLETE_ER_UNUSED4 : {0000,[]string{""},"The statement is unsafe because it uses INSERT DELAYED. This is unsafe because the times when rows are inserted cannot be predicted."},
	ER_BINLOG_UNSAFE_SYSTEM_TABLE : {1670,[]string{"HY000"},"The statement is unsafe because it uses the general log, slow query log, or performance_schema table(s). This is unsafe because system tables may differ on slaves."},
	ER_BINLOG_UNSAFE_AUTOINC_COLUMNS : {1671,[]string{"HY000"},"Statement is unsafe because it invokes a trigger or a stored function that inserts into an AUTO_INCREMENT column. Inserted values cannot be logged correctly."},
	ER_BINLOG_UNSAFE_UDF : {1672,[]string{"HY000"},"Statement is unsafe because it uses a UDF which may not return the same value on the slave."},
	ER_BINLOG_UNSAFE_SYSTEM_VARIABLE : {1673,[]string{"HY000"},"Statement is unsafe because it uses a system variable that may have a different value on the slave."},
	ER_BINLOG_UNSAFE_SYSTEM_FUNCTION : {1674,[]string{"HY000"},"Statement is unsafe because it uses a system function that may return a different value on the slave."},
	ER_BINLOG_UNSAFE_NONTRANS_AFTER_TRANS : {1675,[]string{"HY000"},"Statement is unsafe because it accesses a non-transactional table after accessing a transactional table within the same transaction."},
	ER_MESSAGE_AND_STATEMENT : {1676,[]string{"HY000"},"%s Statement: %s"},
	//OBSOLETE_ER_SLAVE_CONVERSION_FAILED : {1677,[]string{"HY000"},"Column %d of table '%-.192s.%-.192s' cannot be converted from type '%-.32s' to type '%-.32s'"},
	ER_SLAVE_CANT_CREATE_CONVERSION : {1678,[]string{"HY000"},"Can't create conversion table for table '%-.192s.%-.192s'"},
	ER_INSIDE_TRANSACTION_PREVENTS_SWITCH_BINLOG_FORMAT : {1679,[]string{"HY000"},"Cannot modify @@session.binlog_format inside a transaction"},
	ER_PATH_LENGTH : {1680,[]string{"HY000"},"The path specified for %.64s is too long."},
	ER_WARN_DEPRECATED_SYNTAX_NO_REPLACEMENT : {1681,[]string{"HY000"},"'%s' is deprecated and will be removed in a future release."},
	ER_WRONG_NATIVE_TABLE_STRUCTURE : {1682,[]string{"HY000"},"Native table '%-.64s'.'%-.64s' has the wrong structure"},
	ER_WRONG_PERFSCHEMA_USAGE : {1683,[]string{"HY000"},"Invalid performance_schema usage."},
	ER_WARN_I_S_SKIPPED_TABLE : {1684,[]string{"HY000"},"Table '%s'.'%s' was skipped since its definition is being modified by concurrent DDL statement"},
	ER_INSIDE_TRANSACTION_PREVENTS_SWITCH_BINLOG_DIRECT : {1685,[]string{"HY000"},"Cannot modify @@session.binlog_direct_non_transactional_updates inside a transaction"},
	ER_STORED_FUNCTION_PREVENTS_SWITCH_BINLOG_DIRECT : {1686,[]string{"HY000"},"Cannot change the binlog direct flag inside a stored function or trigger"},
	ER_SPATIAL_MUST_HAVE_GEOM_COL : {1687,[]string{"42000"},"A SPATIAL index may only contain a geometrical type column"},
	ER_TOO_LONG_INDEX_COMMENT : {1688,[]string{"HY000"},"Comment for index '%-.64s' is too long (max = %lu)"},
	ER_LOCK_ABORTED : {1689,[]string{"HY000"},"Wait on a lock was aborted due to a pending exclusive lock"},
	ER_DATA_OUT_OF_RANGE : {1690,[]string{"22003"},"%s value is out of range in '%s'"},
	//OBSOLETE_ER_WRONG_SPVAR_TYPE_IN_LIMIT : {1691,[]string{"HY000"},"A variable of a non-integer based type in LIMIT clause"},
	ER_BINLOG_UNSAFE_MULTIPLE_ENGINES_AND_SELF_LOGGING_ENGINE : {1692,[]string{"HY000"},"Mixing self-logging and non-self-logging engines in a statement is unsafe."},
	ER_BINLOG_UNSAFE_MIXED_STATEMENT : {1693,[]string{"HY000"},"Statement accesses nontransactional table as well as transactional or temporary table, and writes to any of them."},
	ER_INSIDE_TRANSACTION_PREVENTS_SWITCH_SQL_LOG_BIN : {1694,[]string{"HY000"},"Cannot modify @@session.sql_log_bin inside a transaction"},
	ER_STORED_FUNCTION_PREVENTS_SWITCH_SQL_LOG_BIN : {1695,[]string{"HY000"},"Cannot change the sql_log_bin inside a stored function or trigger"},
	ER_FAILED_READ_FROM_PAR_FILE : {1696,[]string{"HY000"},"Failed to read from the .par file"},
	ER_VALUES_IS_NOT_INT_TYPE_ERROR : {1697,[]string{"HY000"},"VALUES value for partition '%-.64s' must have type INT"},
	ER_ACCESS_DENIED_NO_PASSWORD_ERROR : {1698,[]string{"28000"},"Access denied for user '%-.48s'@'%-.64s'"},
	ER_SET_PASSWORD_AUTH_PLUGIN : {1699,[]string{"HY000"},"SET PASSWORD has no significance for users authenticating via plugins"},
	//OBSOLETE_ER_GRANT_PLUGIN_USER_EXISTS : {0000,[]string{""},"GRANT with IDENTIFIED WITH is illegal because the user %-.*s already exists"},
	ER_TRUNCATE_ILLEGAL_FK : {1701,[]string{"42000"},"Cannot truncate a table referenced in a foreign key constraint (%.192s)"},
	ER_PLUGIN_IS_PERMANENT : {1702,[]string{"HY000"},"Plugin '%s' is force_plus_permanent and can not be unloaded"},
	ER_SLAVE_HEARTBEAT_VALUE_OUT_OF_RANGE_MIN : {1703,[]string{"HY000"},"The requested value for the heartbeat period is less than 1 millisecond. The value is reset to 0, meaning that heartbeating will effectively be disabled."},
	ER_SLAVE_HEARTBEAT_VALUE_OUT_OF_RANGE_MAX : {1704,[]string{"HY000"},"The requested value for the heartbeat period exceeds the value of `slave_net_timeout' seconds. A sensible value for the period should be less than the timeout."},
	ER_STMT_CACHE_FULL : {1705,[]string{"HY000"},"Multi-row statements required more than 'max_binlog_stmt_cache_size' bytes of storage; increase this mysqld variable and try again"},
	ER_MULTI_UPDATE_KEY_CONFLICT : {1706,[]string{"HY000"},"Primary key/partition key update is not allowed since the table is updated both as '%-.192s' and '%-.192s'."},
	ER_TABLE_NEEDS_REBUILD : {1707,[]string{"HY000"},"Table rebuild required. Please do \"ALTER TABLE `%-.64s` FORCE\" or dump/reload to fix it!"},
	WARN_OPTION_BELOW_LIMIT : {1708,[]string{"HY000"},"The value of '%s' should be no less than the value of '%s'"},
	ER_INDEX_COLUMN_TOO_LONG : {1709,[]string{"HY000"},"Index column size too large. The maximum column size is %lu bytes."},
	ER_ERROR_IN_TRIGGER_BODY : {1710,[]string{"HY000"},"Trigger '%-.64s' has an error in its body: '%-.256s'"},
	ER_ERROR_IN_UNKNOWN_TRIGGER_BODY : {1711,[]string{"HY000"},"Unknown trigger has an error in its body: '%-.256s'"},
	ER_INDEX_CORRUPT : {1712,[]string{"HY000"},"Index %s is corrupted"},
	ER_UNDO_RECORD_TOO_BIG : {1713,[]string{"HY000"},"Undo log record is too big."},
	ER_BINLOG_UNSAFE_INSERT_IGNORE_SELECT : {1714,[]string{"HY000"},"INSERT IGNORE... SELECT is unsafe because the order in which rows are retrieved by the SELECT determines which (if any) rows are ignored. This order cannot be predicted and may differ on master and the slave."},
	ER_BINLOG_UNSAFE_INSERT_SELECT_UPDATE : {1715,[]string{"HY000"},"INSERT... SELECT... ON DUPLICATE KEY UPDATE is unsafe because the order in which rows are retrieved by the SELECT determines which (if any) rows are updated. This order cannot be predicted and may differ on master and the slave."},
	ER_BINLOG_UNSAFE_REPLACE_SELECT : {1716,[]string{"HY000"},"REPLACE... SELECT is unsafe because the order in which rows are retrieved by the SELECT determines which (if any) rows are replaced. This order cannot be predicted and may differ on master and the slave."},
	ER_BINLOG_UNSAFE_CREATE_IGNORE_SELECT : {1717,[]string{"HY000"},"CREATE... IGNORE SELECT is unsafe because the order in which rows are retrieved by the SELECT determines which (if any) rows are ignored. This order cannot be predicted and may differ on master and the slave."},
	ER_BINLOG_UNSAFE_CREATE_REPLACE_SELECT : {1718,[]string{"HY000"},"CREATE... REPLACE SELECT is unsafe because the order in which rows are retrieved by the SELECT determines which (if any) rows are replaced. This order cannot be predicted and may differ on master and the slave."},
	ER_BINLOG_UNSAFE_UPDATE_IGNORE : {1719,[]string{"HY000"},"UPDATE IGNORE is unsafe because the order in which rows are updated determines which (if any) rows are ignored. This order cannot be predicted and may differ on master and the slave."},
	ER_PLUGIN_NO_UNINSTALL : {1720,[]string{"HY000"},"Plugin '%s' is marked as not dynamically uninstallable. You have to stop the server to uninstall it."},
	ER_PLUGIN_NO_INSTALL : {1721,[]string{"HY000"},"Plugin '%s' is marked as not dynamically installable. You have to stop the server to install it."},
	ER_BINLOG_UNSAFE_WRITE_AUTOINC_SELECT : {1722,[]string{"HY000"},"Statements writing to a table with an auto-increment column after selecting from another table are unsafe because the order in which rows are retrieved determines what (if any) rows will be written. This order cannot be predicted and may differ on master and the slave."},
	ER_BINLOG_UNSAFE_CREATE_SELECT_AUTOINC : {1723,[]string{"HY000"},"CREATE TABLE... SELECT...  on a table with an auto-increment column is unsafe because the order in which rows are retrieved by the SELECT determines which (if any) rows are inserted. This order cannot be predicted and may differ on master and the slave."},
	ER_BINLOG_UNSAFE_INSERT_TWO_KEYS : {1724,[]string{"HY000"},"INSERT... ON DUPLICATE KEY UPDATE  on a table with more than one UNIQUE KEY is unsafe"},
	ER_TABLE_IN_FK_CHECK : {1725,[]string{"HY000"},"Table is being used in foreign key check."},
	ER_UNSUPPORTED_ENGINE : {1726,[]string{"HY000"},"Storage engine '%s' does not support system tables. [%s.%s]"},
	ER_BINLOG_UNSAFE_AUTOINC_NOT_FIRST : {1727,[]string{"HY000"},"INSERT into autoincrement field which is not the first part in the composed primary key is unsafe."},
	ER_CANNOT_LOAD_FROM_TABLE_V2 : {1728,[]string{"HY000"},"Cannot load from %s.%s. The table is probably corrupted"},
	ER_MASTER_DELAY_VALUE_OUT_OF_RANGE : {1729,[]string{"HY000"},"The requested value %s for the master delay exceeds the maximum %u"},
	ER_ONLY_FD_AND_RBR_EVENTS_ALLOWED_IN_BINLOG_STATEMENT : {1730,[]string{"HY000"},"Only Format_description_log_event and row events are allowed in BINLOG statements (but %s was provided)"},
	ER_PARTITION_EXCHANGE_DIFFERENT_OPTION : {1731,[]string{"HY000"},"Non matching attribute '%-.64s' between partition and table"},
	ER_PARTITION_EXCHANGE_PART_TABLE : {1732,[]string{"HY000"},"Table to exchange with partition is partitioned: '%-.64s'"},
	ER_PARTITION_EXCHANGE_TEMP_TABLE : {1733,[]string{"HY000"},"Table to exchange with partition is temporary: '%-.64s'"},
	ER_PARTITION_INSTEAD_OF_SUBPARTITION : {1734,[]string{"HY000"},"Subpartitioned table, use subpartition instead of partition"},
	ER_UNKNOWN_PARTITION : {1735,[]string{"HY000"},"Unknown partition '%-.64s' in table '%-.64s'"},
	ER_TABLES_DIFFERENT_METADATA : {1736,[]string{"HY000"},"Tables have different definitions"},
	ER_ROW_DOES_NOT_MATCH_PARTITION : {1737,[]string{"HY000"},"Found a row that does not match the partition"},
	ER_BINLOG_CACHE_SIZE_GREATER_THAN_MAX : {1738,[]string{"HY000"},"Option binlog_cache_size (%lu) is greater than max_binlog_cache_size (%lu); setting binlog_cache_size equal to max_binlog_cache_size."},
	ER_WARN_INDEX_NOT_APPLICABLE : {1739,[]string{"HY000"},"Cannot use %-.64s access on index '%-.64s' due to type or collation conversion on field '%-.64s'"},
	ER_PARTITION_EXCHANGE_FOREIGN_KEY : {1740,[]string{"HY000"},"Table to exchange with partition has foreign key references: '%-.64s'"},
	//OBSOLETE_ER_NO_SUCH_KEY_VALUE : {0000,[]string{""},"Key value '%-.192s' was not found in table '%-.192s.%-.192s'"},
	ER_RPL_INFO_DATA_TOO_LONG : {1742,[]string{"HY000"},"Data for column '%s' too long"},
	//OBSOLETE_ER_NETWORK_READ_EVENT_CHECKSUM_FAILURE : {13115,[]string{"HY000"},"Replication event checksum verification failed while reading from network."},
	//OBSOLETE_ER_BINLOG_READ_EVENT_CHECKSUM_FAILURE : {0000,[]string{""},"Replication event checksum verification failed while reading from a log file."},
	ER_BINLOG_STMT_CACHE_SIZE_GREATER_THAN_MAX : {1745,[]string{"HY000"},"Option binlog_stmt_cache_size (%lu) is greater than max_binlog_stmt_cache_size (%lu); setting binlog_stmt_cache_size equal to max_binlog_stmt_cache_size."},
	ER_CANT_UPDATE_TABLE_IN_CREATE_TABLE_SELECT : {1746,[]string{"HY000"},"Can't update table '%-.192s' while '%-.192s' is being created."},
	ER_PARTITION_CLAUSE_ON_NONPARTITIONED : {1747,[]string{"HY000"},"PARTITION () clause on non partitioned table"},
	ER_ROW_DOES_NOT_MATCH_GIVEN_PARTITION_SET : {1748,[]string{"HY000"},"Found a row not matching the given partition set"},
	//OBSOLETE_ER_NO_SUCH_PARTITION__UNUSED : {0000,[]string{""},"partition '%-.64s' doesn't exist"},
	ER_CHANGE_RPL_INFO_REPOSITORY_FAILURE : {1750,[]string{"HY000"},"Failure while changing the type of replication repository: %s."},
	ER_WARNING_NOT_COMPLETE_ROLLBACK_WITH_CREATED_TEMP_TABLE : {1751,[]string{"HY000"},"The creation of some temporary tables could not be rolled back."},
	ER_WARNING_NOT_COMPLETE_ROLLBACK_WITH_DROPPED_TEMP_TABLE : {1752,[]string{"HY000"},"Some temporary tables were dropped, but these operations could not be rolled back."},
	ER_MTS_FEATURE_IS_NOT_SUPPORTED : {1753,[]string{"HY000"},"%s is not supported in multi-threaded slave mode. %s"},
	ER_MTS_UPDATED_DBS_GREATER_MAX : {1754,[]string{"HY000"},"The number of modified databases exceeds the maximum %d; the database names will not be included in the replication event metadata."},
	ER_MTS_CANT_PARALLEL : {1755,[]string{"HY000"},"Cannot execute the current event group in the parallel mode. Encountered event %s, relay-log name %s, position %s which prevents execution of this event group in parallel mode. Reason: %s."},
	ER_MTS_INCONSISTENT_DATA : {1756,[]string{"HY000"},"%s"},
	ER_FULLTEXT_NOT_SUPPORTED_WITH_PARTITIONING : {1757,[]string{"HY000"},"FULLTEXT index is not supported for partitioned tables."},
	ER_DA_INVALID_CONDITION_NUMBER : {1758,[]string{"35000"},"Invalid condition number"},
	ER_INSECURE_PLAIN_TEXT : {1759,[]string{"HY000"},"Sending passwords in plain text without SSL/TLS is extremely insecure."},
	ER_INSECURE_CHANGE_MASTER : {1760,[]string{"HY000"},"Storing MySQL user name or password information in the master info repository is not secure and is therefore not recommended. Please consider using the USER and PASSWORD connection options for START SLAVE; see the 'START SLAVE Syntax' in the MySQL Manual for more information."},
	ER_FOREIGN_DUPLICATE_KEY_WITH_CHILD_INFO : {1761,[]string{"23000","S1009"},"Foreign key constraint for table '%.192s', record '%-.192s' would lead to a duplicate entry in table '%.192s', key '%.192s'"},
	ER_FOREIGN_DUPLICATE_KEY_WITHOUT_CHILD_INFO : {1762,[]string{"23000","S1009"},"Foreign key constraint for table '%.192s', record '%-.192s' would lead to a duplicate entry in a child table"},
	ER_SQLTHREAD_WITH_SECURE_SLAVE : {1763,[]string{"HY000"},"Setting authentication options is not possible when only the Slave SQL Thread is being started."},
	ER_TABLE_HAS_NO_FT : {1764,[]string{"HY000"},"The table does not have FULLTEXT index to support this query"},
	ER_VARIABLE_NOT_SETTABLE_IN_SF_OR_TRIGGER : {1765,[]string{"HY000"},"The system variable %.200s cannot be set in stored functions or triggers."},
	ER_VARIABLE_NOT_SETTABLE_IN_TRANSACTION : {1766,[]string{"HY000"},"The system variable %.200s cannot be set when there is an ongoing transaction."},
	//OBSOLETE_ER_GTID_NEXT_IS_NOT_IN_GTID_NEXT_LIST : {0000,[]string{""},"The system variable @@SESSION.GTID_NEXT has the value %.200s, which is not listed in @@SESSION.GTID_NEXT_LIST."},
	//OBSOLETE_ER_CANT_CHANGE_GTID_NEXT_IN_TRANSACTION : {0000,[]string{""},"The system variable @@SESSION.GTID_NEXT cannot change inside a transaction."},
	ER_SET_STATEMENT_CANNOT_INVOKE_FUNCTION : {1769,[]string{"HY000"},"The statement 'SET %.200s' cannot invoke a stored function."},
	ER_GTID_NEXT_CANT_BE_AUTOMATIC_IF_GTID_NEXT_LIST_IS_NON_NULL : {1770,[]string{"HY000"},"The system variable @@SESSION.GTID_NEXT cannot be 'AUTOMATIC' when @@SESSION.GTID_NEXT_LIST is non-NULL."},
	//OBSOLETE_ER_SKIPPING_LOGGED_TRANSACTION : {0000,[]string{""},"Skipping transaction %.200s because it has already been executed and logged."},
	ER_MALFORMED_GTID_SET_SPECIFICATION : {1772,[]string{"HY000"},"Malformed GTID set specification '%.200s'."},
	ER_MALFORMED_GTID_SET_ENCODING : {1773,[]string{"HY000"},"Malformed GTID set encoding."},
	ER_MALFORMED_GTID_SPECIFICATION : {1774,[]string{"HY000"},"Malformed GTID specification '%.200s'."},
	ER_GNO_EXHAUSTED : {1775,[]string{"HY000"},"Impossible to generate GTID: the integer component reached the maximum value. Restart the server with a new server_uuid."},
	ER_BAD_SLAVE_AUTO_POSITION : {1776,[]string{"HY000"},"Parameters MASTER_LOG_FILE, MASTER_LOG_POS, RELAY_LOG_FILE and RELAY_LOG_POS cannot be set when MASTER_AUTO_POSITION is active."},
	ER_AUTO_POSITION_REQUIRES_GTID_MODE_NOT_OFF : {1777,[]string{"HY000"},"CHANGE MASTER TO MASTER_AUTO_POSITION = 1 cannot be executed because @@GLOBAL.GTID_MODE = OFF."},
	ER_CANT_DO_IMPLICIT_COMMIT_IN_TRX_WHEN_GTID_NEXT_IS_SET : {1778,[]string{"HY000"},"Cannot execute statements with implicit commit inside a transaction when @@SESSION.GTID_NEXT == 'UUID:NUMBER'."},
	ER_GTID_MODE_ON_REQUIRES_ENFORCE_GTID_CONSISTENCY_ON : {1779,[]string{"HY000"},"GTID_MODE = ON requires ENFORCE_GTID_CONSISTENCY = ON."},
	//OBSOLETE_ER_GTID_MODE_REQUIRES_BINLOG : {0000,[]string{""},"@@GLOBAL.GTID_MODE = ON or ON_PERMISSIVE or OFF_PERMISSIVE requires --log-bin and --log-slave-updates."},
	ER_CANT_SET_GTID_NEXT_TO_GTID_WHEN_GTID_MODE_IS_OFF : {1781,[]string{"HY000"},"@@SESSION.GTID_NEXT cannot be set to UUID:NUMBER when @@GLOBAL.GTID_MODE = OFF."},
	ER_CANT_SET_GTID_NEXT_TO_ANONYMOUS_WHEN_GTID_MODE_IS_ON : {1782,[]string{"HY000"},"@@SESSION.GTID_NEXT cannot be set to ANONYMOUS when @@GLOBAL.GTID_MODE = ON."},
	ER_CANT_SET_GTID_NEXT_LIST_TO_NON_NULL_WHEN_GTID_MODE_IS_OFF : {1783,[]string{"HY000"},"@@SESSION.GTID_NEXT_LIST cannot be set to a non-NULL value when @@GLOBAL.GTID_MODE = OFF."},
	//OBSOLETE_ER_FOUND_GTID_EVENT_WHEN_GTID_MODE_IS_OFF__UNUSED : {0000,[]string{""},"Found a Gtid_log_event when @@GLOBAL.GTID_MODE = OFF."},
	ER_GTID_UNSAFE_NON_TRANSACTIONAL_TABLE : {1785,[]string{"HY000"},"Statement violates GTID consistency: Updates to non-transactional tables can only be done in either autocommitted statements or single-statement transactions, and never in the same statement as updates to transactional tables."},
	ER_GTID_UNSAFE_CREATE_SELECT : {1786,[]string{"HY000"},"Statement violates GTID consistency: CREATE TABLE ... SELECT."},
	//OBSOLETE_ER_GTID_UNSAFE_CREATE_DROP_TEMP_TABLE_IN_TRANSACTION : {0000,[]string{""},"Statement violates GTID consistency: CREATE TEMPORARY TABLE and DROP TEMPORARY TABLE can only be executed outside transactional context.  These statements are also not allowed in a function or trigger because functions and triggers are also considered to be multi-statement transactions."},
	ER_GTID_MODE_CAN_ONLY_CHANGE_ONE_STEP_AT_A_TIME : {1788,[]string{"HY000"},"The value of @@GLOBAL.GTID_MODE can only be changed one step at a time: OFF <-> OFF_PERMISSIVE <-> ON_PERMISSIVE <-> ON. Also note that this value must be stepped up or down simultaneously on all servers. See the Manual for instructions."},
	ER_MASTER_HAS_PURGED_REQUIRED_GTIDS : {1789,[]string{"HY000"},"Cannot replicate because the master purged required binary logs. Replicate the missing transactions from elsewhere, or provision a new slave from backup. Consider increasing the master's binary log expiration period. %s"},
	ER_CANT_SET_GTID_NEXT_WHEN_OWNING_GTID : {1790,[]string{"HY000"},"@@SESSION.GTID_NEXT cannot be changed by a client that owns a GTID. The client owns %s. Ownership is released on COMMIT or ROLLBACK."},
	ER_UNKNOWN_EXPLAIN_FORMAT : {1791,[]string{"HY000"},"Unknown EXPLAIN format name: '%s'"},
	ER_CANT_EXECUTE_IN_READ_ONLY_TRANSACTION : {1792,[]string{"25006"},"Cannot execute statement in a READ ONLY transaction."},
	ER_TOO_LONG_TABLE_PARTITION_COMMENT : {1793,[]string{"HY000"},"Comment for table partition '%-.64s' is too long (max = %lu)"},
	ER_SLAVE_CONFIGURATION : {1794,[]string{"HY000"},"Slave is not configured or failed to initialize properly. You must at least set --server-id to enable either a master or a slave. Additional error messages can be found in the MySQL error log."},
	ER_INNODB_FT_LIMIT : {1795,[]string{"HY000"},"InnoDB presently supports one FULLTEXT index creation at a time"},
	ER_INNODB_NO_FT_TEMP_TABLE : {1796,[]string{"HY000"},"Cannot create FULLTEXT index on temporary InnoDB table"},
	ER_INNODB_FT_WRONG_DOCID_COLUMN : {1797,[]string{"HY000"},"Column '%-.192s' is of wrong type for an InnoDB FULLTEXT index"},
	ER_INNODB_FT_WRONG_DOCID_INDEX : {1798,[]string{"HY000"},"Index '%-.192s' is of wrong type for an InnoDB FULLTEXT index"},
	ER_INNODB_ONLINE_LOG_TOO_BIG : {1799,[]string{"HY000"},"Creating index '%-.192s' required more than 'innodb_online_alter_log_max_size' bytes of modification log. Please try again."},
	ER_UNKNOWN_ALTER_ALGORITHM : {1800,[]string{"HY000"},"Unknown ALGORITHM '%s'"},
	ER_UNKNOWN_ALTER_LOCK : {1801,[]string{"HY000"},"Unknown LOCK type '%s'"},
	ER_MTS_CHANGE_MASTER_CANT_RUN_WITH_GAPS : {1802,[]string{"HY000"},"CHANGE MASTER cannot be executed when the slave was stopped with an error or killed in MTS mode. Consider using RESET SLAVE or START SLAVE UNTIL."},
	ER_MTS_RECOVERY_FAILURE : {1803,[]string{"HY000"},"Cannot recover after SLAVE errored out in parallel execution mode. Additional error messages can be found in the MySQL error log."},
	ER_MTS_RESET_WORKERS : {1804,[]string{"HY000"},"Cannot clean up worker info tables. Additional error messages can be found in the MySQL error log."},
	ER_COL_COUNT_DOESNT_MATCH_CORRUPTED_V2 : {1805,[]string{"HY000"},"Column count of %s.%s is wrong. Expected %d, found %d. The table is probably corrupted"},
	ER_SLAVE_SILENT_RETRY_TRANSACTION : {1806,[]string{"HY000"},"Slave must silently retry current transaction"},
	ER_DISCARD_FK_CHECKS_RUNNING : {1807,[]string{"HY000"},"There is a foreign key check running on table '%-.192s'. Cannot discard the table."},
	ER_TABLE_SCHEMA_MISMATCH : {1808,[]string{"HY000"},"Schema mismatch (%s)"},
	ER_TABLE_IN_SYSTEM_TABLESPACE : {1809,[]string{"HY000"},"Table '%-.192s' in system tablespace"},
	ER_IO_READ_ERROR : {1810,[]string{"HY000"},"IO Read error: (%lu, %s) %s"},
	ER_IO_WRITE_ERROR : {1811,[]string{"HY000"},"IO Write error: (%lu, %s) %s"},
	ER_TABLESPACE_MISSING : {1812,[]string{"HY000"},"Tablespace is missing for table %s."},
	ER_TABLESPACE_EXISTS : {1813,[]string{"HY000"},"Tablespace '%-.192s' exists."},
	ER_TABLESPACE_DISCARDED : {1814,[]string{"HY000"},"Tablespace has been discarded for table '%-.192s'"},
	ER_INTERNAL_ERROR : {1815,[]string{"HY000"},"Internal error: %s"},
	ER_INNODB_IMPORT_ERROR : {1816,[]string{"HY000"},"ALTER TABLE %-.192s IMPORT TABLESPACE failed with error %lu : '%s'"},
	ER_INNODB_INDEX_CORRUPT : {1817,[]string{"HY000"},"Index corrupt: %s"},
	ER_INVALID_YEAR_COLUMN_LENGTH : {1818,[]string{"HY000"},"Invalid display width. Use YEAR instead."},
	ER_NOT_VALID_PASSWORD : {1819,[]string{"HY000"},"Your password does not satisfy the current policy requirements"},
	ER_MUST_CHANGE_PASSWORD : {1820,[]string{"HY000"},"You must reset your password using ALTER USER statement before executing this statement."},
	ER_FK_NO_INDEX_CHILD : {1821,[]string{"HY000"},"Failed to add the foreign key constraint. Missing index for constraint '%s' in the foreign table '%s'"},
	ER_FK_NO_INDEX_PARENT : {1822,[]string{"HY000"},"Failed to add the foreign key constraint. Missing index for constraint '%s' in the referenced table '%s'"},
	ER_FK_FAIL_ADD_SYSTEM : {1823,[]string{"HY000"},"Failed to add the foreign key constraint '%s' to system tables"},
	ER_FK_CANNOT_OPEN_PARENT : {1824,[]string{"HY000"},"Failed to open the referenced table '%s'"},
	ER_FK_INCORRECT_OPTION : {1825,[]string{"HY000"},"Failed to add the foreign key constraint on table '%s'. Incorrect options in FOREIGN KEY constraint '%s'"},
	ER_FK_DUP_NAME : {1826,[]string{"HY000"},"Duplicate foreign key constraint name '%s'"},
	ER_PASSWORD_FORMAT : {1827,[]string{"HY000"},"The password hash doesn't have the expected format."},
	ER_FK_COLUMN_CANNOT_DROP : {1828,[]string{"HY000"},"Cannot drop column '%-.192s': needed in a foreign key constraint '%-.192s'"},
	ER_FK_COLUMN_CANNOT_DROP_CHILD : {1829,[]string{"HY000"},"Cannot drop column '%-.192s': needed in a foreign key constraint '%-.192s' of table '%-.192s'"},
	ER_FK_COLUMN_NOT_NULL : {1830,[]string{"HY000"},"Column '%-.192s' cannot be NOT NULL: needed in a foreign key constraint '%-.192s' SET NULL"},
	ER_DUP_INDEX : {1831,[]string{"HY000"},"Duplicate index '%-.64s' defined on the table '%-.64s.%-.64s'. This is deprecated and will be disallowed in a future release."},
	ER_FK_COLUMN_CANNOT_CHANGE : {1832,[]string{"HY000"},"Cannot change column '%-.192s': used in a foreign key constraint '%-.192s'"},
	ER_FK_COLUMN_CANNOT_CHANGE_CHILD : {1833,[]string{"HY000"},"Cannot change column '%-.192s': used in a foreign key constraint '%-.192s' of table '%-.192s'"},
	//OBSOLETE_ER_UNUSED5 : {0000,[]string{""},"Cannot delete rows from table which is parent in a foreign key constraint '%-.192s' of table '%-.192s'"},
	ER_MALFORMED_PACKET : {1835,[]string{"HY000"},"Malformed communication packet."},
	ER_READ_ONLY_MODE : {1836,[]string{"HY000"},"Running in read-only mode"},
	ER_GTID_NEXT_TYPE_UNDEFINED_GTID : {1837,[]string{"HY000"},"When @@SESSION.GTID_NEXT is set to a GTID, you must explicitly set it to a different value after a COMMIT or ROLLBACK. Please check GTID_NEXT variable manual page for detailed explanation. Current @@SESSION.GTID_NEXT is '%s'."},
	ER_VARIABLE_NOT_SETTABLE_IN_SP : {1838,[]string{"HY000"},"The system variable %.200s cannot be set in stored procedures."},
	//OBSOLETE_ER_CANT_SET_GTID_PURGED_WHEN_GTID_MODE_IS_OFF : {0000,[]string{""},"@@GLOBAL.GTID_PURGED can only be set when @@GLOBAL.GTID_MODE = ON."},
	ER_CANT_SET_GTID_PURGED_WHEN_GTID_EXECUTED_IS_NOT_EMPTY : {1840,[]string{"HY000"},"@@GLOBAL.GTID_PURGED can only be set when @@GLOBAL.GTID_EXECUTED is empty."},
	ER_CANT_SET_GTID_PURGED_WHEN_OWNED_GTIDS_IS_NOT_EMPTY : {1841,[]string{"HY000"},"@@GLOBAL.GTID_PURGED can only be set when there are no ongoing transactions (not even in other clients)."},
	ER_GTID_PURGED_WAS_CHANGED : {1842,[]string{"HY000"},"@@GLOBAL.GTID_PURGED was changed from '%s' to '%s'."},
	ER_GTID_EXECUTED_WAS_CHANGED : {1843,[]string{"HY000"},"@@GLOBAL.GTID_EXECUTED was changed from '%s' to '%s'."},
	ER_BINLOG_STMT_MODE_AND_NO_REPL_TABLES : {1844,[]string{"HY000"},"Cannot execute statement: impossible to write to binary log since BINLOG_FORMAT = STATEMENT, and both replicated and non replicated tables are written to."},
	ER_ALTER_OPERATION_NOT_SUPPORTED : {1845,[]string{"0A000"},"%s is not supported for this operation. Try %s."},
	ER_ALTER_OPERATION_NOT_SUPPORTED_REASON : {1846,[]string{"0A000"},"%s is not supported. Reason: %s. Try %s."},
	ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_COPY : {1847,[]string{"HY000"},"COPY algorithm requires a lock"},
	ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_PARTITION : {1848,[]string{"HY000"},"Partition specific operations do not yet support LOCK/ALGORITHM"},
	ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FK_RENAME : {1849,[]string{"HY000"},"Columns participating in a foreign key are renamed"},
	ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_COLUMN_TYPE : {1850,[]string{"HY000"},"Cannot change column type INPLACE"},
	ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FK_CHECK : {1851,[]string{"HY000"},"Adding foreign keys needs foreign_key_checks=OFF"},
	//OBSOLETE_ER_UNUSED6 : {0000,[]string{""},"Creating unique indexes with IGNORE requires COPY algorithm to remove duplicate rows"},
	ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_NOPK : {1853,[]string{"HY000"},"Dropping a primary key is not allowed without also adding a new primary key"},
	ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_AUTOINC : {1854,[]string{"HY000"},"Adding an auto-increment column requires a lock"},
	ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_HIDDEN_FTS : {1855,[]string{"HY000"},"Cannot replace hidden FTS_DOC_ID with a user-visible one"},
	ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_CHANGE_FTS : {1856,[]string{"HY000"},"Cannot drop or rename FTS_DOC_ID"},
	ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_FTS : {1857,[]string{"HY000"},"Fulltext index creation requires a lock"},
	//OBSOLETE_ER_SQL_SLAVE_SKIP_COUNTER_NOT_SETTABLE_IN_GTID_MODE : {1858,[]string{"HY000"},"sql_slave_skip_counter can not be set when the server is running with @@GLOBAL.GTID_MODE = ON. Instead, for each transaction that you want to skip, generate an empty transaction with the same GTID as the transaction"},
	ER_DUP_UNKNOWN_IN_INDEX : {1859,[]string{"23000"},"Duplicate entry for key '%-.192s'"},
	ER_IDENT_CAUSES_TOO_LONG_PATH : {1860,[]string{"HY000"},"Long database name and identifier for object resulted in path length exceeding %d characters. Path: '%s'."},
	ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_NOT_NULL : {1861,[]string{"HY000"},"cannot silently convert NULL values, as required in this SQL_MODE"},
	ER_MUST_CHANGE_PASSWORD_LOGIN : {1862,[]string{"HY000"},"Your password has expired. To log in you must change it using a client that supports expired passwords."},
	ER_ROW_IN_WRONG_PARTITION : {1863,[]string{"HY000"},"Found a row in wrong partition %s"},
	ER_MTS_EVENT_BIGGER_PENDING_JOBS_SIZE_MAX : {1864,[]string{"HY000"},"Cannot schedule event %s, relay-log name %s, position %s to Worker thread because its size %lu exceeds %lu of slave_pending_jobs_size_max."},
	//OBSOLETE_ER_INNODB_NO_FT_USES_PARSER : {0000,[]string{""},"Cannot CREATE FULLTEXT INDEX WITH PARSER on InnoDB table"},
	ER_BINLOG_LOGICAL_CORRUPTION : {1866,[]string{"HY000"},"The binary log file '%s' is logically corrupted: %s"},
	ER_WARN_PURGE_LOG_IN_USE : {1867,[]string{"HY000"},"file %s was not purged because it was being read by %d thread(s), purged only %d out of %d files."},
	ER_WARN_PURGE_LOG_IS_ACTIVE : {1868,[]string{"HY000"},"file %s was not purged because it is the active log file."},
	ER_AUTO_INCREMENT_CONFLICT : {1869,[]string{"HY000"},"Auto-increment value in UPDATE conflicts with internally generated values"},
	WARN_ON_BLOCKHOLE_IN_RBR : {1870,[]string{"HY000"},"Row events are not logged for %s statements that modify BLACKHOLE tables in row format. Table(s): '%-.192s'"},
	ER_SLAVE_MI_INIT_REPOSITORY : {1871,[]string{"HY000"},"Slave failed to initialize master info structure from the repository"},
	ER_SLAVE_RLI_INIT_REPOSITORY : {1872,[]string{"HY000"},"Slave failed to initialize relay log info structure from the repository"},
	ER_ACCESS_DENIED_CHANGE_USER_ERROR : {1873,[]string{"28000"},"Access denied trying to change to user '%-.48s'@'%-.64s' (using password: %s). Disconnecting."},
	ER_INNODB_READ_ONLY : {1874,[]string{"HY000"},"InnoDB is in read only mode."},
	ER_STOP_SLAVE_SQL_THREAD_TIMEOUT : {1875,[]string{"HY000"},"STOP SLAVE command execution is incomplete: Slave SQL thread got the stop signal, thread is busy, SQL thread will stop once the current task is complete."},
	ER_STOP_SLAVE_IO_THREAD_TIMEOUT : {1876,[]string{"HY000"},"STOP SLAVE command execution is incomplete: Slave IO thread got the stop signal, thread is busy, IO thread will stop once the current task is complete."},
	ER_TABLE_CORRUPT : {1877,[]string{"HY000"},"Operation cannot be performed. The table '%-.64s.%-.64s' is missing, corrupt or contains bad data."},
	ER_TEMP_FILE_WRITE_FAILURE : {1878,[]string{"HY000"},"Temporary file write failure."},
	ER_INNODB_FT_AUX_NOT_HEX_ID : {1879,[]string{"HY000"},"Upgrade index name failed, please use create index(alter table) algorithm copy to rebuild index."},
	ER_OLD_TEMPORALS_UPGRADED : {1880,[]string{"HY000"},"TIME/TIMESTAMP/DATETIME columns of old format have been upgraded to the new format."},
	ER_INNODB_FORCED_RECOVERY : {1881,[]string{"HY000"},"Operation not allowed when innodb_force_recovery > 0."},
	ER_AES_INVALID_IV : {1882,[]string{"HY000"},"The initialization vector supplied to %s is too short. Must be at least %d bytes long"},
	ER_PLUGIN_CANNOT_BE_UNINSTALLED : {1883,[]string{"HY000"},"Plugin '%s' cannot be uninstalled now. %s"},
	ER_GTID_UNSAFE_BINLOG_SPLITTABLE_STATEMENT_AND_ASSIGNED_GTID : {1884,[]string{"HY000"},"Cannot execute statement because it needs to be written to the binary log as multiple statements, and this is not allowed when @@SESSION.GTID_NEXT == 'UUID:NUMBER'."},
	ER_SLAVE_HAS_MORE_GTIDS_THAN_MASTER : {1885,[]string{"HY000"},"Slave has more GTIDs than the master has, using the master's SERVER_UUID. This may indicate that the end of the binary log was truncated or that the last binary log file was lost, e.g., after a power or disk failure when sync_binlog != 1. The master may or may not have rolled back transactions that were already replicated to the slave. Suggest to replicate any transactions that master has rolled back from slave to master, and/or commit empty transactions on master to account for transactions that have been committed on master but are not included in GTID_EXECUTED."},
	ER_MISSING_KEY : {1886,[]string{"HY000"},"The table '%s.%s' does not have the necessary key(s) defined on it. Please check the table definition and create index(s) accordingly."},
	WARN_NAMED_PIPE_ACCESS_EVERYONE : {1887,[]string{"HY000"},"Setting named_pipe_full_access_group='%s' is insecure. Consider using a Windows group with fewer members."},
	ER_FILE_CORRUPT : {3000,[]string{"HY000"},"File %s is corrupted"},
	ER_ERROR_ON_MASTER : {3001,[]string{"HY000"},"Query partially completed on the master (error on master: %d) and was aborted. There is a chance that your master is inconsistent at this point. If you are sure that your master is ok, run this query manually on the slave and then restart the slave with SET GLOBAL SQL_SLAVE_SKIP_COUNTER=1; START SLAVE;. Query:'%s'"},
	//OBSOLETE_ER_INCONSISTENT_ERROR : {13113,[]string{"HY000"},"Query caused different errors on master and slave. Error on master: message (format)='%s' error code=%d; Error on slave:actual message='%s', error code=%d. Default database:'%s'. Query:'%s'"},
	ER_STORAGE_ENGINE_NOT_LOADED : {3003,[]string{"HY000"},"Storage engine for table '%s'.'%s' is not loaded."},
	ER_GET_STACKED_DA_WITHOUT_ACTIVE_HANDLER : {3004,[]string{"0Z002"},"GET STACKED DIAGNOSTICS when handler not active"},
	ER_WARN_LEGACY_SYNTAX_CONVERTED : {3005,[]string{"HY000"},"%s is no longer supported. The statement was converted to %s."},
	ER_BINLOG_UNSAFE_FULLTEXT_PLUGIN : {3006,[]string{"HY000"},"Statement is unsafe because it uses a fulltext parser plugin which may not return the same value on the slave."},
	ER_CANNOT_DISCARD_TEMPORARY_TABLE : {3007,[]string{"HY000"},"Cannot DISCARD/IMPORT tablespace associated with temporary table"},
	ER_FK_DEPTH_EXCEEDED : {3008,[]string{"HY000"},"Foreign key cascade delete/update exceeds max depth of %d."},
	ER_COL_COUNT_DOESNT_MATCH_PLEASE_UPDATE_V2 : {3009,[]string{"HY000"},"The column count of %s.%s is wrong. Expected %d, found %d. Created with MySQL %d, now running %d. Please perform the MySQL upgrade procedure."},
	ER_WARN_TRIGGER_DOESNT_HAVE_CREATED : {3010,[]string{"HY000"},"Trigger %s.%s.%s does not have CREATED attribute."},
	ER_REFERENCED_TRG_DOES_NOT_EXIST : {3011,[]string{"HY000"},"Referenced trigger '%s' for the given action time and event type does not exist."},
	ER_EXPLAIN_NOT_SUPPORTED : {3012,[]string{"HY000"},"EXPLAIN FOR CONNECTION command is supported only for SELECT/UPDATE/INSERT/DELETE/REPLACE"},
	ER_INVALID_FIELD_SIZE : {3013,[]string{"HY000"},"Invalid size for column '%-.192s'."},
	ER_MISSING_HA_CREATE_OPTION : {3014,[]string{"HY000"},"Table storage engine '%-.64s' found required create option missing"},
	ER_ENGINE_OUT_OF_MEMORY : {3015,[]string{"HY000"},"Out of memory in storage engine '%-.64s'."},
	ER_PASSWORD_EXPIRE_ANONYMOUS_USER : {3016,[]string{"HY000"},"The password for anonymous user cannot be expired."},
	ER_SLAVE_SQL_THREAD_MUST_STOP : {3017,[]string{"HY000"},"This operation cannot be performed with a running slave sql thread; run STOP SLAVE SQL_THREAD first"},
	ER_NO_FT_MATERIALIZED_SUBQUERY : {3018,[]string{"HY000"},"Cannot create FULLTEXT index on materialized subquery"},
	ER_INNODB_UNDO_LOG_FULL : {3019,[]string{"HY000"},"Undo Log error: %s"},
	ER_INVALID_ARGUMENT_FOR_LOGARITHM : {3020,[]string{"2201E"},"Invalid argument for logarithm"},
	ER_SLAVE_CHANNEL_IO_THREAD_MUST_STOP : {3021,[]string{"HY000"},"This operation cannot be performed with a running slave io thread; run STOP SLAVE IO_THREAD FOR CHANNEL '%s' first."},
	ER_WARN_OPEN_TEMP_TABLES_MUST_BE_ZERO : {3022,[]string{"HY000"},"This operation may not be safe when the slave has temporary tables. The tables will be kept open until the server restarts or until the tables are deleted by any replicated DROP statement. Suggest to wait until slave_open_temp_tables = 0."},
	ER_WARN_ONLY_MASTER_LOG_FILE_NO_POS : {3023,[]string{"HY000"},"CHANGE MASTER TO with a MASTER_LOG_FILE clause but no MASTER_LOG_POS clause may not be safe. The old position value may not be valid for the new binary log file."},
	ER_QUERY_TIMEOUT : {3024,[]string{"HY000"},"Query execution was interrupted, maximum statement execution time exceeded"},
	ER_NON_RO_SELECT_DISABLE_TIMER : {3025,[]string{"HY000"},"Select is not a read only statement, disabling timer"},
	ER_DUP_LIST_ENTRY : {3026,[]string{"HY000"},"Duplicate entry '%-.192s'."},
	//OBSOLETE_ER_SQL_MODE_NO_EFFECT : {0000,[]string{""},"'%s' mode no longer has any effect. Use STRICT_ALL_TABLES or STRICT_TRANS_TABLES instead."},
	ER_AGGREGATE_ORDER_FOR_UNION : {3028,[]string{"HY000"},"Expression #%u of ORDER BY contains aggregate function and applies to a UNION"},
	ER_AGGREGATE_ORDER_NON_AGG_QUERY : {3029,[]string{"HY000"},"Expression #%u of ORDER BY contains aggregate function and applies to the result of a non-aggregated query"},
	ER_SLAVE_WORKER_STOPPED_PREVIOUS_THD_ERROR : {3030,[]string{"HY000"},"Slave worker has stopped after at least one previous worker encountered an error when slave-preserve-commit-order was enabled. To preserve commit order, the last transaction executed by this thread has not been committed. When restarting the slave after fixing any failed threads, you should fix this worker as well."},
	ER_DONT_SUPPORT_SLAVE_PRESERVE_COMMIT_ORDER : {3031,[]string{"HY000"},"slave_preserve_commit_order is not supported %s."},
	ER_SERVER_OFFLINE_MODE : {3032,[]string{"HY000"},"The server is currently in offline mode"},
	ER_GIS_DIFFERENT_SRIDS : {3033,[]string{"HY000"},"Binary geometry function %s given two geometries of different srids: %u and %u, which should have been identical."},
	ER_GIS_UNSUPPORTED_ARGUMENT : {3034,[]string{"HY000"},"Calling geometry function %s with unsupported types of arguments."},
	ER_GIS_UNKNOWN_ERROR : {3035,[]string{"HY000"},"Unknown GIS error occurred in function %s."},
	ER_GIS_UNKNOWN_EXCEPTION : {3036,[]string{"HY000"},"Unknown exception caught in GIS function %s."},
	ER_GIS_INVALID_DATA : {3037,[]string{"22023"},"Invalid GIS data provided to function %s."},
	ER_BOOST_GEOMETRY_EMPTY_INPUT_EXCEPTION : {3038,[]string{"HY000"},"The geometry has no data in function %s."},
	ER_BOOST_GEOMETRY_CENTROID_EXCEPTION : {3039,[]string{"HY000"},"Unable to calculate centroid because geometry is empty in function %s."},
	ER_BOOST_GEOMETRY_OVERLAY_INVALID_INPUT_EXCEPTION : {3040,[]string{"HY000"},"Geometry overlay calculation error: geometry data is invalid in function %s."},
	ER_BOOST_GEOMETRY_TURN_INFO_EXCEPTION : {3041,[]string{"HY000"},"Geometry turn info calculation error: geometry data is invalid in function %s."},
	ER_BOOST_GEOMETRY_SELF_INTERSECTION_POINT_EXCEPTION : {3042,[]string{"HY000"},"Analysis procedures of intersection points interrupted unexpectedly in function %s."},
	ER_BOOST_GEOMETRY_UNKNOWN_EXCEPTION : {3043,[]string{"HY000"},"Unknown exception thrown in function %s."},
	ER_STD_BAD_ALLOC_ERROR : {3044,[]string{"HY000"},"Memory allocation error: %-.256s in function %s."},
	ER_STD_DOMAIN_ERROR : {3045,[]string{"HY000"},"Domain error: %-.256s in function %s."},
	ER_STD_LENGTH_ERROR : {3046,[]string{"HY000"},"Length error: %-.256s in function %s."},
	ER_STD_INVALID_ARGUMENT : {3047,[]string{"HY000"},"Invalid argument error: %-.256s in function %s."},
	ER_STD_OUT_OF_RANGE_ERROR : {3048,[]string{"HY000"},"Out of range error: %-.256s in function %s."},
	ER_STD_OVERFLOW_ERROR : {3049,[]string{"HY000"},"Overflow error: %-.256s in function %s."},
	ER_STD_RANGE_ERROR : {3050,[]string{"HY000"},"Range error: %-.256s in function %s."},
	ER_STD_UNDERFLOW_ERROR : {3051,[]string{"HY000"},"Underflow error: %-.256s in function %s."},
	ER_STD_LOGIC_ERROR : {3052,[]string{"HY000"},"Logic error: %-.256s in function %s."},
	ER_STD_RUNTIME_ERROR : {3053,[]string{"HY000"},"Runtime error: %-.256s in function %s."},
	ER_STD_UNKNOWN_EXCEPTION : {3054,[]string{"HY000"},"Unknown exception: %-.384s in function %s."},
	ER_GIS_DATA_WRONG_ENDIANESS : {3055,[]string{"HY000"},"Geometry byte string must be little endian."},
	ER_CHANGE_MASTER_PASSWORD_LENGTH : {3056,[]string{"HY000"},"The password provided for the replication user exceeds the maximum length of 32 characters"},
	ER_USER_LOCK_WRONG_NAME : {3057,[]string{"42000"},"Incorrect user-level lock name '%-.192s'."},
	ER_USER_LOCK_DEADLOCK : {3058,[]string{"HY000"},"Deadlock found when trying to get user-level lock; try rolling back transaction/releasing locks and restarting lock acquisition."},
	ER_REPLACE_INACCESSIBLE_ROWS : {3059,[]string{"HY000"},"REPLACE cannot be executed as it requires deleting rows that are not in the view"},
	ER_ALTER_OPERATION_NOT_SUPPORTED_REASON_GIS : {3060,[]string{"HY000"},"Do not support online operation on table with GIS index"},
	ER_ILLEGAL_USER_VAR : {3061,[]string{"42000","S1009"},"User variable name '%-.100s' is illegal"},
	ER_GTID_MODE_OFF : {3062,[]string{"HY000"},"Cannot %s when GTID_MODE = OFF."},
	//OBSOLETE_ER_UNSUPPORTED_BY_REPLICATION_THREAD : {0000,[]string{""},"Cannot %s from a replication slave thread."},
	ER_INCORRECT_TYPE : {3064,[]string{"HY000"},"Incorrect type for argument %s in function %s."},
	ER_FIELD_IN_ORDER_NOT_SELECT : {3065,[]string{"HY000"},"Expression #%u of ORDER BY clause is not in SELECT list, references column '%-.192s' which is not in SELECT list; this is incompatible with %s"},
	ER_AGGREGATE_IN_ORDER_NOT_SELECT : {3066,[]string{"HY000"},"Expression #%u of ORDER BY clause is not in SELECT list, contains aggregate function; this is incompatible with %s"},
	ER_INVALID_RPL_WILD_TABLE_FILTER_PATTERN : {3067,[]string{"HY000"},"Supplied filter list contains a value which is not in the required format 'db_pattern.table_pattern'"},
	ER_NET_OK_PACKET_TOO_LARGE : {3068,[]string{"08S01"},"OK packet too large"},
	ER_INVALID_JSON_DATA : {3069,[]string{"HY000"},"Invalid JSON data provided to function %s: %s"},
	ER_INVALID_GEOJSON_MISSING_MEMBER : {3070,[]string{"HY000"},"Invalid GeoJSON data provided to function %s: Missing required member '%s'"},
	ER_INVALID_GEOJSON_WRONG_TYPE : {3071,[]string{"HY000"},"Invalid GeoJSON data provided to function %s: Member '%s' must be of type '%s'"},
	ER_INVALID_GEOJSON_UNSPECIFIED : {3072,[]string{"HY000"},"Invalid GeoJSON data provided to function %s"},
	ER_DIMENSION_UNSUPPORTED : {3073,[]string{"HY000"},"Unsupported number of coordinate dimensions in function %s: Found %u, expected %u"},
	ER_SLAVE_CHANNEL_DOES_NOT_EXIST : {3074,[]string{"HY000"},"Slave channel '%s' does not exist."},
	//OBSOLETE_ER_SLAVE_MULTIPLE_CHANNELS_HOST_PORT : {0000,[]string{""},"A slave channel '%s' already exists for the given host and port combination."},
	ER_SLAVE_CHANNEL_NAME_INVALID_OR_TOO_LONG : {3076,[]string{"HY000"},"Couldn't create channel: Channel name is either invalid or too long."},
	ER_SLAVE_NEW_CHANNEL_WRONG_REPOSITORY : {3077,[]string{"HY000"},"To have multiple channels, repository cannot be of type FILE; Please check the repository configuration and convert them to TABLE."},
	//OBSOLETE_ER_SLAVE_CHANNEL_DELETE : {0000,[]string{""},"Cannot delete slave info objects for channel '%s'."},
	ER_SLAVE_MULTIPLE_CHANNELS_CMD : {3079,[]string{"HY000"},"Multiple channels exist on the slave. Please provide channel name as an argument."},
	ER_SLAVE_MAX_CHANNELS_EXCEEDED : {3080,[]string{"HY000"},"Maximum number of replication channels allowed exceeded."},
	ER_SLAVE_CHANNEL_MUST_STOP : {3081,[]string{"HY000"},"This operation cannot be performed with running replication threads; run STOP SLAVE FOR CHANNEL '%s' first"},
	ER_SLAVE_CHANNEL_NOT_RUNNING : {3082,[]string{"HY000"},"This operation requires running replication threads; configure slave and run START SLAVE FOR CHANNEL '%s'"},
	ER_SLAVE_CHANNEL_WAS_RUNNING : {3083,[]string{"HY000"},"Replication thread(s) for channel '%s' are already runnning."},
	ER_SLAVE_CHANNEL_WAS_NOT_RUNNING : {3084,[]string{"HY000"},"Replication thread(s) for channel '%s' are already stopped."},
	ER_SLAVE_CHANNEL_SQL_THREAD_MUST_STOP : {3085,[]string{"HY000"},"This operation cannot be performed with a running slave sql thread; run STOP SLAVE SQL_THREAD FOR CHANNEL '%s' first."},
	ER_SLAVE_CHANNEL_SQL_SKIP_COUNTER : {3086,[]string{"HY000"},"When sql_slave_skip_counter > 0, it is not allowed to start more than one SQL thread by using 'START SLAVE [SQL_THREAD]'. Value of sql_slave_skip_counter can only be used by one SQL thread at a time. Please use 'START SLAVE [SQL_THREAD] FOR CHANNEL' to start the SQL thread which will use the value of sql_slave_skip_counter."},
	ER_WRONG_FIELD_WITH_GROUP_V2 : {3087,[]string{"HY000"},"Expression #%u of %s is not in GROUP BY clause and contains nonaggregated column '%-.192s' which is not functionally dependent on columns in GROUP BY clause; this is incompatible with sql_mode=only_full_group_by"},
	ER_MIX_OF_GROUP_FUNC_AND_FIELDS_V2 : {3088,[]string{"HY000"},"In aggregated query without GROUP BY, expression #%u of %s contains nonaggregated column '%-.192s'; this is incompatible with sql_mode=only_full_group_by"},
	ER_WARN_DEPRECATED_SYSVAR_UPDATE : {3089,[]string{"HY000"},"Updating '%s' is deprecated. It will be made read-only in a future release."},
	ER_WARN_DEPRECATED_SQLMODE : {3090,[]string{"HY000"},"Changing sql mode '%s' is deprecated. It will be removed in a future release."},
	ER_CANNOT_LOG_PARTIAL_DROP_DATABASE_WITH_GTID : {3091,[]string{"HY000"},"DROP DATABASE failed; some tables may have been dropped but the database directory remains. The GTID has not been added to GTID_EXECUTED and the statement was not written to the binary log. Fix this as follows: (1) remove all files from the database directory %-.192s; (2) SET GTID_NEXT='%-.192s'; (3) DROP DATABASE `%-.192s`."},
	ER_GROUP_REPLICATION_CONFIGURATION : {3092,[]string{"HY000"},"The server is not configured properly to be an active member of the group. Please see more details on error log."},
	ER_GROUP_REPLICATION_RUNNING : {3093,[]string{"HY000"},"The START GROUP_REPLICATION command failed since the group is already running."},
	ER_GROUP_REPLICATION_APPLIER_INIT_ERROR : {3094,[]string{"HY000"},"The START GROUP_REPLICATION command failed as the applier module failed to start."},
	ER_GROUP_REPLICATION_STOP_APPLIER_THREAD_TIMEOUT : {3095,[]string{"HY000"},"The STOP GROUP_REPLICATION command execution is incomplete: The applier thread got the stop signal while it was busy. The applier thread will stop once the current task is complete."},
	ER_GROUP_REPLICATION_COMMUNICATION_LAYER_SESSION_ERROR : {3096,[]string{"HY000"},"The START GROUP_REPLICATION command failed as there was an error when initializing the group communication layer."},
	ER_GROUP_REPLICATION_COMMUNICATION_LAYER_JOIN_ERROR : {3097,[]string{"HY000"},"The START GROUP_REPLICATION command failed as there was an error when joining the communication group."},
	ER_BEFORE_DML_VALIDATION_ERROR : {3098,[]string{"HY000"},"The table does not comply with the requirements by an external plugin."},
	ER_PREVENTS_VARIABLE_WITHOUT_RBR : {3099,[]string{"HY000"},"Cannot change the value of variable %s without binary log format as ROW."},
	ER_RUN_HOOK_ERROR : {3100,[]string{"HY000"},"Error on observer while running replication hook '%s'."},
	ER_TRANSACTION_ROLLBACK_DURING_COMMIT : {3101,[]string{"40000"},"Plugin instructed the server to rollback the current transaction."},
	ER_GENERATED_COLUMN_FUNCTION_IS_NOT_ALLOWED : {3102,[]string{"HY000"},"Expression of generated column '%s' contains a disallowed function."},
	ER_UNSUPPORTED_ALTER_INPLACE_ON_VIRTUAL_COLUMN : {3103,[]string{"HY000"},"INPLACE ADD or DROP of virtual columns cannot be combined with other ALTER TABLE actions"},
	ER_WRONG_FK_OPTION_FOR_GENERATED_COLUMN : {3104,[]string{"HY000"},"Cannot define foreign key with %s clause on a generated column."},
	ER_NON_DEFAULT_VALUE_FOR_GENERATED_COLUMN : {3105,[]string{"HY000"},"The value specified for generated column '%s' in table '%s' is not allowed."},
	ER_UNSUPPORTED_ACTION_ON_GENERATED_COLUMN : {3106,[]string{"HY000"},"'%s' is not supported for generated columns."},
	ER_GENERATED_COLUMN_NON_PRIOR : {3107,[]string{"HY000"},"Generated column can refer only to generated columns defined prior to it."},
	ER_DEPENDENT_BY_GENERATED_COLUMN : {3108,[]string{"HY000"},"Column '%s' has a generated column dependency."},
	ER_GENERATED_COLUMN_REF_AUTO_INC : {3109,[]string{"HY000"},"Generated column '%s' cannot refer to auto-increment column."},
	ER_FEATURE_NOT_AVAILABLE : {3110,[]string{"HY000"},"The '%-.64s' feature is not available; you need to remove '%-.64s' or use MySQL built with '%-.64s'"},
	ER_CANT_SET_GTID_MODE : {3111,[]string{"HY000"},"SET @@GLOBAL.GTID_MODE = %-.64s is not allowed because %-.384s."},
	ER_CANT_USE_AUTO_POSITION_WITH_GTID_MODE_OFF : {3112,[]string{"HY000"},"The replication receiver thread%-.192s cannot start in AUTO_POSITION mode: this server uses @@GLOBAL.GTID_MODE = OFF."},
	//OBSOLETE_ER_CANT_REPLICATE_ANONYMOUS_WITH_AUTO_POSITION : {13148,[]string{"HY000"},"Cannot replicate anonymous transaction when AUTO_POSITION = 1, at file %.512s, position %lld."},
	//OBSOLETE_ER_CANT_REPLICATE_ANONYMOUS_WITH_GTID_MODE_ON : {13149,[]string{"HY000"},"Cannot replicate anonymous transaction when @@GLOBAL.GTID_MODE = ON, at file %.512s, position %lld."},
	//OBSOLETE_ER_CANT_REPLICATE_GTID_WITH_GTID_MODE_OFF : {13150,[]string{"HY000"},"Cannot replicate GTID-transaction when @@GLOBAL.GTID_MODE = OFF, at file %.512s, position %lld."},
	ER_CANT_ENFORCE_GTID_CONSISTENCY_WITH_ONGOING_GTID_VIOLATING_TX : {3116,[]string{"HY000"},"Cannot set ENFORCE_GTID_CONSISTENCY = ON because there are ongoing transactions that violate GTID consistency."},
	ER_ENFORCE_GTID_CONSISTENCY_WARN_WITH_ONGOING_GTID_VIOLATING_TX : {3117,[]string{"HY000"},"There are ongoing transactions that violate GTID consistency."},
	ER_ACCOUNT_HAS_BEEN_LOCKED : {3118,[]string{"HY000"},"Access denied for user '%-.48s'@'%-.64s'. Account is locked."},
	ER_WRONG_TABLESPACE_NAME : {3119,[]string{"42000"},"Incorrect tablespace name `%-.192s`"},
	ER_TABLESPACE_IS_NOT_EMPTY : {3120,[]string{"HY000"},"Tablespace `%-.192s` is not empty."},
	ER_WRONG_FILE_NAME : {3121,[]string{"HY000"},"Incorrect File Name '%s'."},
	ER_BOOST_GEOMETRY_INCONSISTENT_TURNS_EXCEPTION : {3122,[]string{"HY000"},"Inconsistent intersection points."},
	ER_WARN_OPTIMIZER_HINT_SYNTAX_ERROR : {3123,[]string{"HY000"},"Optimizer hint syntax error"},
	ER_WARN_BAD_MAX_EXECUTION_TIME : {3124,[]string{"HY000"},"Unsupported MAX_EXECUTION_TIME"},
	ER_WARN_UNSUPPORTED_MAX_EXECUTION_TIME : {3125,[]string{"HY000"},"MAX_EXECUTION_TIME hint is supported by top-level standalone SELECT statements only"},
	ER_WARN_CONFLICTING_HINT : {3126,[]string{"HY000"},"Hint %s is ignored as conflicting/duplicated"},
	ER_WARN_UNKNOWN_QB_NAME : {3127,[]string{"HY000"},"Query block name %s is not found for %s hint"},
	ER_UNRESOLVED_HINT_NAME : {3128,[]string{"HY000"},"Unresolved name %s for %s hint"},
	ER_WARN_ON_MODIFYING_GTID_EXECUTED_TABLE : {3129,[]string{"HY000"},"Please do not modify the %s table. This is a mysql internal system table to store GTIDs for committed transactions. Modifying it can lead to an inconsistent GTID state."},
	ER_PLUGGABLE_PROTOCOL_COMMAND_NOT_SUPPORTED : {3130,[]string{"HY000"},"Command not supported by pluggable protocols"},
	ER_LOCKING_SERVICE_WRONG_NAME : {3131,[]string{"42000"},"Incorrect locking service lock name '%-.192s'."},
	ER_LOCKING_SERVICE_DEADLOCK : {3132,[]string{"HY000"},"Deadlock found when trying to get locking service lock; try releasing locks and restarting lock acquisition."},
	ER_LOCKING_SERVICE_TIMEOUT : {3133,[]string{"HY000"},"Service lock wait timeout exceeded."},
	ER_GIS_MAX_POINTS_IN_GEOMETRY_OVERFLOWED : {3134,[]string{"HY000"},"Parameter %s exceeds the maximum number of points in a geometry (%lu) in function %s."},
	ER_SQL_MODE_MERGED : {3135,[]string{"HY000"},"'NO_ZERO_DATE', 'NO_ZERO_IN_DATE' and 'ERROR_FOR_DIVISION_BY_ZERO' sql modes should be used with strict mode. They will be merged with strict mode in a future release."},
	ER_VTOKEN_PLUGIN_TOKEN_MISMATCH : {3136,[]string{"HY000"},"Version token mismatch for %.*s. Correct value %.*s"},
	ER_VTOKEN_PLUGIN_TOKEN_NOT_FOUND : {3137,[]string{"HY000"},"Version token %.*s not found."},
	ER_CANT_SET_VARIABLE_WHEN_OWNING_GTID : {3138,[]string{"HY000"},"Variable %-.192s cannot be changed by a client that owns a GTID. The client owns %s. Ownership is released on COMMIT or ROLLBACK."},
	ER_SLAVE_CHANNEL_OPERATION_NOT_ALLOWED : {3139,[]string{"HY000"},"%-.192s cannot be performed on channel '%-.192s'."},
	ER_INVALID_JSON_TEXT : {3140,[]string{"22032"},"Invalid JSON text: \"%s\" at position %u in value for column '%-.200s'."},
	ER_INVALID_JSON_TEXT_IN_PARAM : {3141,[]string{"22032"},"Invalid JSON text in argument %u to function %s: \"%s\" at position %u.%-.0s"},
	ER_INVALID_JSON_BINARY_DATA : {3142,[]string{"HY000"},"The JSON binary value contains invalid data."},
	ER_INVALID_JSON_PATH : {3143,[]string{"42000"},"Invalid JSON path expression. The error is around character position %u.%-.200s"},
	ER_INVALID_JSON_CHARSET : {3144,[]string{"22032"},"Cannot create a JSON value from a string with CHARACTER SET '%s'."},
	ER_INVALID_JSON_CHARSET_IN_FUNCTION : {3145,[]string{"22032"},"Invalid JSON character data provided to function %s: '%s'; utf8 is required."},
	ER_INVALID_TYPE_FOR_JSON : {3146,[]string{"22032"},"Invalid data type for JSON data in argument %u to function %s; a JSON string or JSON type is required."},
	ER_INVALID_CAST_TO_JSON : {3147,[]string{"22032"},"Cannot CAST value to JSON."},
	ER_INVALID_JSON_PATH_CHARSET : {3148,[]string{"42000"},"A path expression must be encoded in the utf8 character set. The path expression '%-.200s' is encoded in character set '%-.200s'."},
	ER_INVALID_JSON_PATH_WILDCARD : {3149,[]string{"42000"},"In this situation, path expressions may not contain the * and ** tokens or an array range."},
	ER_JSON_VALUE_TOO_BIG : {3150,[]string{"22032"},"The JSON value is too big to be stored in a JSON column."},
	ER_JSON_KEY_TOO_BIG : {3151,[]string{"22032"},"The JSON object contains a key name that is too long."},
	ER_JSON_USED_AS_KEY : {3152,[]string{"42000"},"JSON column '%-.192s' supports indexing only via generated columns on a specified JSON path."},
	ER_JSON_VACUOUS_PATH : {3153,[]string{"42000"},"The path expression '$' is not allowed in this context."},
	ER_JSON_BAD_ONE_OR_ALL_ARG : {3154,[]string{"42000"},"The oneOrAll argument to %s may take these values: 'one' or 'all'."},
	ER_NUMERIC_JSON_VALUE_OUT_OF_RANGE : {3155,[]string{"22003"},"Out of range JSON value for CAST to %s%-.0s from column %s at row %ld"},
	ER_INVALID_JSON_VALUE_FOR_CAST : {3156,[]string{"22018"},"Invalid JSON value for CAST to %s%-.0s from column %s at row %ld"},
	ER_JSON_DOCUMENT_TOO_DEEP : {3157,[]string{"22032"},"The JSON document exceeds the maximum depth."},
	ER_JSON_DOCUMENT_NULL_KEY : {3158,[]string{"22032"},"JSON documents may not contain NULL member names."},
	ER_SECURE_TRANSPORT_REQUIRED : {3159,[]string{"HY000"},"Connections using insecure transport are prohibited while --require_secure_transport=ON."},
	ER_NO_SECURE_TRANSPORTS_CONFIGURED : {3160,[]string{"HY000"},"No secure transports (SSL or Shared Memory) are configured, unable to set --require_secure_transport=ON."},
	ER_DISABLED_STORAGE_ENGINE : {3161,[]string{"HY000"},"Storage engine %s is disabled (Table creation is disallowed)."},
	ER_USER_DOES_NOT_EXIST : {3162,[]string{"HY000"},"Authorization ID %s does not exist."},
	ER_USER_ALREADY_EXISTS : {3163,[]string{"HY000"},"Authorization ID %s already exists."},
	ER_AUDIT_API_ABORT : {3164,[]string{"HY000"},"Aborted by Audit API ('%-.48s';%d)."},
	ER_INVALID_JSON_PATH_ARRAY_CELL : {3165,[]string{"42000"},"A path expression is not a path to a cell in an array."},
	ER_BUFPOOL_RESIZE_INPROGRESS : {3166,[]string{"HY000"},"Another buffer pool resize is already in progress."},
	ER_FEATURE_DISABLED_SEE_DOC : {3167,[]string{"HY000"},"The '%s' feature is disabled; see the documentation for '%s'"},
	ER_SERVER_ISNT_AVAILABLE : {3168,[]string{"HY000"},"Server isn't available"},
	ER_SESSION_WAS_KILLED : {3169,[]string{"HY000"},"Session was killed"},
	ER_CAPACITY_EXCEEDED : {3170,[]string{"HY000"},"Memory capacity of %llu bytes for '%s' exceeded. %s"},
	ER_CAPACITY_EXCEEDED_IN_RANGE_OPTIMIZER : {3171,[]string{"HY000"},"Range optimization was not done for this query."},
	//OBSOLETE_ER_TABLE_NEEDS_UPG_PART : {0000,[]string{""},"Partitioning upgrade required. Please dump/reload to fix it or do: ALTER TABLE `%-.192s`.`%-.192s` UPGRADE PARTITIONING"},
	ER_CANT_WAIT_FOR_EXECUTED_GTID_SET_WHILE_OWNING_A_GTID : {3173,[]string{"HY000"},"The client holds ownership of the GTID %s. Therefore, WAIT_FOR_EXECUTED_GTID_SET cannot wait for this GTID."},
	ER_CANNOT_ADD_FOREIGN_BASE_COL_VIRTUAL : {3174,[]string{"HY000"},"Cannot add foreign key on the base column of indexed virtual column."},
	ER_CANNOT_CREATE_VIRTUAL_INDEX_CONSTRAINT : {3175,[]string{"HY000"},"Cannot create index on virtual column whose base column has foreign constraint."},
	ER_ERROR_ON_MODIFYING_GTID_EXECUTED_TABLE : {3176,[]string{"HY000"},"Please do not modify the %s table with an XA transaction. This is an internal system table used to store GTIDs for committed transactions. Although modifying it can lead to an inconsistent GTID state, if neccessary you can modify it with a non-XA transaction."},
	ER_LOCK_REFUSED_BY_ENGINE : {3177,[]string{"HY000"},"Lock acquisition refused by storage engine."},
	ER_UNSUPPORTED_ALTER_ONLINE_ON_VIRTUAL_COLUMN : {3178,[]string{"HY000"},"ADD COLUMN col...VIRTUAL, ADD INDEX(col)"},
	ER_MASTER_KEY_ROTATION_NOT_SUPPORTED_BY_SE : {3179,[]string{"HY000"},"Master key rotation is not supported by storage engine."},
	//OBSOLETE_ER_MASTER_KEY_ROTATION_ERROR_BY_SE : {0000,[]string{""},"Encryption key rotation error reported by SE: %s"},
	ER_MASTER_KEY_ROTATION_BINLOG_FAILED : {3181,[]string{"HY000"},"Write to binlog failed. However, master key rotation has been completed successfully."},
	ER_MASTER_KEY_ROTATION_SE_UNAVAILABLE : {3182,[]string{"HY000"},"Storage engine is not available."},
	ER_TABLESPACE_CANNOT_ENCRYPT : {3183,[]string{"HY000"},"This tablespace can't be encrypted."},
	ER_INVALID_ENCRYPTION_OPTION : {3184,[]string{"HY000"},"Invalid encryption option."},
	ER_CANNOT_FIND_KEY_IN_KEYRING : {3185,[]string{"HY000"},"Can't find master key from keyring, please check in the server log if a keyring plugin is loaded and initialized successfully."},
	ER_CAPACITY_EXCEEDED_IN_PARSER : {3186,[]string{"HY000"},"Parser bailed out for this query."},
	ER_UNSUPPORTED_ALTER_ENCRYPTION_INPLACE : {3187,[]string{"HY000"},"Cannot alter encryption attribute by inplace algorithm."},
	ER_KEYRING_UDF_KEYRING_SERVICE_ERROR : {3188,[]string{"HY000"},"Function '%s' failed because underlying keyring service returned an error. Please check if a keyring plugin is installed and that provided arguments are valid for the keyring you are using."},
	ER_USER_COLUMN_OLD_LENGTH : {3189,[]string{"HY000"},"It seems that your db schema is old. The %s column is 77 characters long and should be 93 characters long. Please perform the MySQL upgrade procedure."},
	ER_CANT_RESET_MASTER : {3190,[]string{"HY000"},"RESET MASTER is not allowed because %-.384s."},
	ER_GROUP_REPLICATION_MAX_GROUP_SIZE : {3191,[]string{"HY000"},"The START GROUP_REPLICATION command failed since the group already has 9 members."},
	ER_CANNOT_ADD_FOREIGN_BASE_COL_STORED : {3192,[]string{"HY000"},"Cannot add foreign key on the base column of stored column. "},
	ER_TABLE_REFERENCED : {3193,[]string{"HY000"},"Cannot complete the operation because table is referenced by another connection."},
	//OBSOLETE_ER_PARTITION_ENGINE_DEPRECATED_FOR_TABLE : {0000,[]string{""},"The partition engine, used by table '%-.192s.%-.192s', is deprecated and will be removed in a future release. Please use native partitioning instead."},
	//OBSOLETE_ER_WARN_USING_GEOMFROMWKB_TO_SET_SRID_ZERO : {0000,[]string{"01000"},"%.192s(geometry) is deprecated and will be replaced by st_srid(geometry, 0) in a future version. Use %.192s(st_aswkb(geometry), 0) instead."},
	//OBSOLETE_ER_WARN_USING_GEOMFROMWKB_TO_SET_SRID : {0000,[]string{"01000"},"%.192s(geometry, srid) is deprecated and will be replaced by st_srid(geometry, srid) in a future version. Use %.192s(st_aswkb(geometry), srid) instead."},
	ER_XA_RETRY : {3197,[]string{"HY000"},"The resource manager is not able to commit the transaction branch at this time. Please retry later."},
	ER_KEYRING_AWS_UDF_AWS_KMS_ERROR : {3198,[]string{"HY000"},"Function %s failed due to: %s."},
	ER_BINLOG_UNSAFE_XA : {3199,[]string{"HY000"},"Statement is unsafe because it is being used inside a XA transaction. Concurrent XA transactions may deadlock on slaves when replicated using statements."},
	ER_UDF_ERROR : {3200,[]string{"HY000"},"%s UDF failed; %s"},
	ER_KEYRING_MIGRATION_FAILURE : {3201,[]string{"HY000"},"Can not perform keyring migration : %s"},
	ER_KEYRING_ACCESS_DENIED_ERROR : {3202,[]string{"42000"},"Access denied; you need %-.128s privileges for this operation"},
	ER_KEYRING_MIGRATION_STATUS : {3203,[]string{"HY000"},"Keyring migration %s."},
	//OBSOLETE_ER_PLUGIN_FAILED_TO_OPEN_TABLES : {13154,[]string{"HY000"},"Failed to open the %s filter tables."},
	//OBSOLETE_ER_PLUGIN_FAILED_TO_OPEN_TABLE : {13155,[]string{"HY000"},"Failed to open '%s.%s' %s table."},
	//OBSOLETE_ER_AUDIT_LOG_NO_KEYRING_PLUGIN_INSTALLED : {13159,[]string{"HY000"},"No keyring plugin installed."},
	//OBSOLETE_ER_AUDIT_LOG_ENCRYPTION_PASSWORD_HAS_NOT_BEEN_SET : {13161,[]string{"HY000"},"Audit log encryption password has not been set; it will be generated automatically. Use audit_log_encryption_password_get to obtain the password or audit_log_encryption_password_set to set a new one."},
	//OBSOLETE_ER_AUDIT_LOG_COULD_NOT_CREATE_AES_KEY : {13162,[]string{"HY000"},"Could not create AES key. OpenSSL's EVP_BytesToKey function failed."},
	//OBSOLETE_ER_AUDIT_LOG_ENCRYPTION_PASSWORD_CANNOT_BE_FETCHED : {13163,[]string{"HY000"},"Audit log encryption password cannot be fetched from the keyring. Password used so far is used for encryption."},
	//OBSOLETE_ER_AUDIT_LOG_JSON_FILTERING_NOT_ENABLED : {13153,[]string{"HY000"},"Audit Log filtering has not been installed."},
	//OBSOLETE_ER_AUDIT_LOG_UDF_INSUFFICIENT_PRIVILEGE : {13158,[]string{"HY000"},"Request ignored for '%s'@'%s'. SUPER privilege or AUDIT_ADMIN role needed to perform operation"},
	//OBSOLETE_ER_AUDIT_LOG_SUPER_PRIVILEGE_REQUIRED : {13478,[]string{"HY000"},"SUPER privilege or AUDIT_ADMIN role required for '%s'@'%s' user."},
	//OBSOLETE_ER_COULD_NOT_REINITIALIZE_AUDIT_LOG_FILTERS : {13164,[]string{"HY000"},"Could not reinitialize audit log filters."},
	//OBSOLETE_ER_AUDIT_LOG_UDF_INVALID_ARGUMENT_TYPE : {13479,[]string{"HY000"},"Invalid argument type"},
	//OBSOLETE_ER_AUDIT_LOG_UDF_INVALID_ARGUMENT_COUNT : {13480,[]string{"HY000"},"Invalid argument count"},
	//OBSOLETE_ER_AUDIT_LOG_HAS_NOT_BEEN_INSTALLED : {13481,[]string{"HY000"},"audit_log plugin has not been installed using INSTALL PLUGIN syntax."},
	//OBSOLETE_ER_AUDIT_LOG_UDF_READ_INVALID_MAX_ARRAY_LENGTH_ARG_TYPE : {13482,[]string{"HY000"},"Invalid \"max_array_length\" argument type."},
	ER_AUDIT_LOG_UDF_READ_INVALID_MAX_ARRAY_LENGTH_ARG_VALUE : {3218,[]string{"HY000"},"Invalid \"max_array_length\" argument value."},
	//OBSOLETE_ER_AUDIT_LOG_JSON_FILTER_PARSING_ERROR : {13152,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_AUDIT_LOG_JSON_FILTER_NAME_CANNOT_BE_EMPTY : {13156,[]string{"HY000"},"Filter name cannot be empty."},
	//OBSOLETE_ER_AUDIT_LOG_JSON_USER_NAME_CANNOT_BE_EMPTY : {13165,[]string{"HY000"},"User cannot be empty."},
	//OBSOLETE_ER_AUDIT_LOG_JSON_FILTER_DOES_NOT_EXISTS : {0000,[]string{""},"Specified filter has not been found."},
	//OBSOLETE_ER_AUDIT_LOG_USER_FIRST_CHARACTER_MUST_BE_ALPHANUMERIC : {13166,[]string{"HY000"},"First character of the user name must be alphanumeric."},
	//OBSOLETE_ER_AUDIT_LOG_USER_NAME_INVALID_CHARACTER : {13157,[]string{"HY000"},"Invalid character in the user name."},
	//OBSOLETE_ER_AUDIT_LOG_HOST_NAME_INVALID_CHARACTER : {13160,[]string{"HY000"},"Invalid character in the host name."},
	//OBSOLETE_WARN_DEPRECATED_MAXDB_SQL_MODE_FOR_TIMESTAMP : {0000,[]string{""},"With the MAXDB SQL mode enabled, TIMESTAMP is identical with DATETIME. The MAXDB SQL mode is deprecated and will be removed in a future release. Please disable the MAXDB SQL mode and use DATETIME instead."},
	//OBSOLETE_ER_XA_REPLICATION_FILTERS : {3898,[]string{"HY000"},"The use of replication filters with XA transactions is not supported, and can lead to an undefined state in the replication slave."},
	//OBSOLETE_ER_CANT_OPEN_ERROR_LOG : {10187,[]string{"HY000"},"Could not open file '%s' for error logging%s%s"},
	//OBSOLETE_ER_GROUPING_ON_TIMESTAMP_IN_DST : {3912,[]string{"HY000"},"Grouping on temporal is non-deterministic for timezones having DST. Please consider switching to UTC for this query."},
	//OBSOLETE_ER_CANT_START_SERVER_NAMED_PIPE : {0000,[]string{""},"Can't start server : Named Pipe \"%s\" already in use."},
	ER_WRITE_SET_EXCEEDS_LIMIT : {3231,[]string{"HY000"},"The size of writeset data for the current transaction exceeds a limit imposed by an external component. If using Group Replication check 'group_replication_transaction_size_limit'."},
	ER_UNSUPPORT_COMPRESSED_TEMPORARY_TABLE : {3500,[]string{"HY000"},"CREATE TEMPORARY TABLE is not allowed with ROW_FORMAT=COMPRESSED or KEY_BLOCK_SIZE."},
	ER_ACL_OPERATION_FAILED : {3501,[]string{"HY000"},"The ACL operation failed due to the following error from SE: errcode %d - %.256s"},
	ER_UNSUPPORTED_INDEX_ALGORITHM : {3502,[]string{"HY000"},"This storage engine does not support the %s index algorithm, storage engine default was used instead."},
	ER_NO_SUCH_DB : {3503,[]string{"42Y07"},"Database '%-.192s' doesn't exist"},
	ER_TOO_BIG_ENUM : {3504,[]string{"HY000"},"Too many enumeration values for column %.192s."},
	ER_TOO_LONG_SET_ENUM_VALUE : {3505,[]string{"HY000"},"Too long enumeration/set value for column %.192s."},
	ER_INVALID_DD_OBJECT : {3506,[]string{"HY000"},"%s dictionary object is invalid. (%.192s)"},
	ER_UPDATING_DD_TABLE : {3507,[]string{"HY000"},"Failed to update %.192s dictionary object."},
	ER_INVALID_DD_OBJECT_ID : {3508,[]string{"HY000"},"Dictionary object id (%lu) does not exist."},
	ER_INVALID_DD_OBJECT_NAME : {3509,[]string{"HY000"},"Dictionary object name '%s' is invalid. (%.192s)"},
	ER_TABLESPACE_MISSING_WITH_NAME : {3510,[]string{"HY000"},"Tablespace %-.192s doesn't exist."},
	ER_TOO_LONG_ROUTINE_COMMENT : {3511,[]string{"HY000"},"Comment for routine '%-.64s' is too long (max = %lu)"},
	ER_SP_LOAD_FAILED : {3512,[]string{"HY000"},"Failed to load routine '%-.64s'."},
	ER_INVALID_BITWISE_OPERANDS_SIZE : {3513,[]string{"HY000"},"Binary operands of bitwise operators must be of equal length"},
	ER_INVALID_BITWISE_AGGREGATE_OPERANDS_SIZE : {3514,[]string{"HY000"},"Aggregate bitwise functions cannot accept arguments longer than 511 bytes; consider using the SUBSTRING() function"},
	ER_WARN_UNSUPPORTED_HINT : {3515,[]string{"HY000"},"Hints aren't supported in %s"},
	ER_UNEXPECTED_GEOMETRY_TYPE : {3516,[]string{"22S01"},"%.64s value is a geometry of unexpected type %.64s in %.64s."},
	ER_SRS_PARSE_ERROR : {3517,[]string{"SR002"},"Can't parse the spatial reference system definition of SRID %u."},
	ER_SRS_PROJ_PARAMETER_MISSING : {3518,[]string{"SR003"},"The spatial reference system definition for SRID %u does not specify the mandatory %s (EPSG %u) projection parameter."},
	ER_WARN_SRS_NOT_FOUND : {3519,[]string{"01000"},"There's no spatial reference system with SRID %u."},
	ER_SRS_NOT_CARTESIAN : {3520,[]string{"22S00"},"Function %s is only defined for Cartesian spatial reference systems, but one of its arguments is in SRID %u, which is not Cartesian."},
	ER_SRS_NOT_CARTESIAN_UNDEFINED : {3521,[]string{"SR001"},"Function %s is only defined for Cartesian spatial reference systems, but one of its arguments is in SRID %u, which has not been defined."},
	ER_PK_INDEX_CANT_BE_INVISIBLE : {3522,[]string{"HY000"},"A primary key index cannot be invisible"},
	ER_UNKNOWN_AUTHID : {3523,[]string{"HY000"},"Unknown authorization ID `%.64s`@`%.64s`"},
	ER_FAILED_ROLE_GRANT : {3524,[]string{"HY000"},"Failed to grant %.90s` to %.90s"},
	ER_OPEN_ROLE_TABLES : {3525,[]string{"HY000"},"Failed to open the security system tables"},
	ER_FAILED_DEFAULT_ROLES : {3526,[]string{"HY000"},"Failed to set default roles"},
	ER_COMPONENTS_NO_SCHEME : {3527,[]string{"HY000"},"Cannot find schema in specified URN: '%.192s'."},
	ER_COMPONENTS_NO_SCHEME_SERVICE : {3528,[]string{"HY000"},"Cannot acquire scheme load service implementation for schema '%.192s' in specified URN: '%.192s'."},
	ER_COMPONENTS_CANT_LOAD : {3529,[]string{"HY000"},"Cannot load component from specified URN: '%.192s'."},
	ER_ROLE_NOT_GRANTED : {3530,[]string{"HY000"},"`%.64s`@`%.64s` is not granted to `%.64s`@`%.64s`"},
	ER_FAILED_REVOKE_ROLE : {3531,[]string{"HY000"},"Could not revoke role from `%.64s`@`%.64s`"},
	ER_RENAME_ROLE : {3532,[]string{"HY000"},"Renaming of a role identifier is forbidden"},
	ER_COMPONENTS_CANT_ACQUIRE_SERVICE_IMPLEMENTATION : {3533,[]string{"HY000"},"Cannot acquire specified service implementation: '%.192s'."},
	ER_COMPONENTS_CANT_SATISFY_DEPENDENCY : {3534,[]string{"HY000"},"Cannot satisfy dependency for service '%.192s' required by component '%.192s'."},
	ER_COMPONENTS_LOAD_CANT_REGISTER_SERVICE_IMPLEMENTATION : {3535,[]string{"HY000"},"Cannot register service implementation '%.192s' provided by component '%.192s'."},
	ER_COMPONENTS_LOAD_CANT_INITIALIZE : {3536,[]string{"HY000"},"Initialization method provided by component '%.192s' failed."},
	ER_COMPONENTS_UNLOAD_NOT_LOADED : {3537,[]string{"HY000"},"Component specified by URN '%.192s' to unload has not been loaded before."},
	ER_COMPONENTS_UNLOAD_CANT_DEINITIALIZE : {3538,[]string{"HY000"},"De-initialization method provided by component '%.192s' failed."},
	ER_COMPONENTS_CANT_RELEASE_SERVICE : {3539,[]string{"HY000"},"Release of previously acquired service implementation failed."},
	ER_COMPONENTS_UNLOAD_CANT_UNREGISTER_SERVICE : {3540,[]string{"HY000"},"Unregistration of service implementation '%.192s' provided by component '%.192s' failed during unloading of the component."},
	ER_COMPONENTS_CANT_UNLOAD : {3541,[]string{"HY000"},"Cannot unload component from specified URN: '%.192s'."},
	ER_WARN_UNLOAD_THE_NOT_PERSISTED : {3542,[]string{"HY000"},"The Persistent Dynamic Loader was used to unload a component '%.192s', but it was not used to load that component before."},
	ER_COMPONENT_TABLE_INCORRECT : {3543,[]string{"HY000"},"The mysql.component table is missing or has an incorrect definition."},
	ER_COMPONENT_MANIPULATE_ROW_FAILED : {3544,[]string{"HY000"},"Failed to manipulate component '%.192s' persistence data. Error code %d from storage engine."},
	ER_COMPONENTS_UNLOAD_DUPLICATE_IN_GROUP : {3545,[]string{"HY000"},"The component with specified URN: '%.192s' was specified in group more than once."},
	ER_CANT_SET_GTID_PURGED_DUE_SETS_CONSTRAINTS : {3546,[]string{"HY000"},"@@GLOBAL.GTID_PURGED cannot be changed: %s"},
	ER_CANNOT_LOCK_USER_MANAGEMENT_CACHES : {3547,[]string{"HY000"},"Can not lock user management caches for processing."},
	ER_SRS_NOT_FOUND : {3548,[]string{"SR001"},"There's no spatial reference system with SRID %u."},
	ER_VARIABLE_NOT_PERSISTED : {3549,[]string{"HY000"},"Variables cannot be persisted. Please retry."},
	ER_IS_QUERY_INVALID_CLAUSE : {3550,[]string{"HY000"},"Information schema queries do not support the '%s' clause."},
	ER_UNABLE_TO_STORE_STATISTICS : {3551,[]string{"HY000"},"Unable to store dynamic %s statistics into data dictionary."},
	ER_NO_SYSTEM_SCHEMA_ACCESS : {3552,[]string{"HY000"},"Access to system schema '%.64s' is rejected."},
	ER_NO_SYSTEM_TABLESPACE_ACCESS : {3553,[]string{"HY000"},"Access to system tablespace '%.64s' is rejected."},
	ER_NO_SYSTEM_TABLE_ACCESS : {3554,[]string{"HY000"},"Access to %.64s '%.64s.%.64s' is rejected."},
	ER_NO_SYSTEM_TABLE_ACCESS_FOR_DICTIONARY_TABLE : {3555,[]string{"HY000"},"data dictionary table"},
	ER_NO_SYSTEM_TABLE_ACCESS_FOR_SYSTEM_TABLE : {3556,[]string{"HY000"},"system table"},
	ER_NO_SYSTEM_TABLE_ACCESS_FOR_TABLE : {3557,[]string{"HY000"},"table"},
	ER_INVALID_OPTION_KEY : {3558,[]string{"22023"},"Invalid option key '%.192s' in function %.192s."},
	ER_INVALID_OPTION_VALUE : {3559,[]string{"22023"},"Invalid value '%.192s' for option '%.192s' in function '%.192s'."},
	ER_INVALID_OPTION_KEY_VALUE_PAIR : {3560,[]string{"22023"},"The string '%.192s' is not a valid key %c value pair in function %.192s."},
	ER_INVALID_OPTION_START_CHARACTER : {3561,[]string{"22023"},"The options argument in function %.192s starts with the invalid character '%c'."},
	ER_INVALID_OPTION_END_CHARACTER : {3562,[]string{"22023"},"The options argument in function %.192s ends with the invalid character '%c'."},
	ER_INVALID_OPTION_CHARACTERS : {3563,[]string{"22023"},"The options argument in function %.192s contains the invalid character sequence '%.192s'."},
	ER_DUPLICATE_OPTION_KEY : {3564,[]string{"22023"},"Duplicate option key '%.192s' in funtion '%.192s'."},
	ER_WARN_SRS_NOT_FOUND_AXIS_ORDER : {3565,[]string{"01000"},"There's no spatial reference system with SRID %u. The axis order is unknown."},
	ER_NO_ACCESS_TO_NATIVE_FCT : {3566,[]string{"HY000"},"Access to native function '%.64s' is rejected."},
	ER_RESET_MASTER_TO_VALUE_OUT_OF_RANGE : {3567,[]string{"HY000"},"The requested value '%llu' for the next binary log index is out of range. Please use a value between '1' and '%lu'."},
	ER_UNRESOLVED_TABLE_LOCK : {3568,[]string{"HY000"},"Unresolved table name %s in locking clause."},
	ER_DUPLICATE_TABLE_LOCK : {3569,[]string{"HY000"},"Table %s appears in multiple locking clauses."},
	ER_BINLOG_UNSAFE_SKIP_LOCKED : {3570,[]string{"HY000"},"Statement is unsafe because it uses SKIP LOCKED. The set of inserted values is non-deterministic."},
	ER_BINLOG_UNSAFE_NOWAIT : {3571,[]string{"HY000"},"Statement is unsafe because it uses NOWAIT. Whether the command will succeed or fail is not deterministic."},
	ER_LOCK_NOWAIT : {3572,[]string{"HY000"},"Statement aborted because lock(s) could not be acquired immediately and NOWAIT is set."},
	ER_CTE_RECURSIVE_REQUIRES_UNION : {3573,[]string{"HY000"},"Recursive Common Table Expression '%s' should contain a UNION"},
	ER_CTE_RECURSIVE_REQUIRES_NONRECURSIVE_FIRST : {3574,[]string{"HY000"},"Recursive Common Table Expression '%s' should have one or more non-recursive query blocks followed by one or more recursive ones"},
	ER_CTE_RECURSIVE_FORBIDS_AGGREGATION : {3575,[]string{"HY000"},"Recursive Common Table Expression '%s' can contain neither aggregation nor window functions in recursive query block"},
	ER_CTE_RECURSIVE_FORBIDDEN_JOIN_ORDER : {3576,[]string{"HY000"},"In recursive query block of Recursive Common Table Expression '%s', the recursive table must neither be in the right argument of a LEFT JOIN, nor be forced to be non-first with join order hints"},
	ER_CTE_RECURSIVE_REQUIRES_SINGLE_REFERENCE : {3577,[]string{"HY000"},"In recursive query block of Recursive Common Table Expression '%s', the recursive table must be referenced only once, and not in any subquery"},
	ER_SWITCH_TMP_ENGINE : {3578,[]string{"HY000"},"'%s' requires @@internal_tmp_disk_storage_engine=InnoDB"},
	ER_WINDOW_NO_SUCH_WINDOW : {3579,[]string{"HY000"},"Window name '%s' is not defined."},
	ER_WINDOW_CIRCULARITY_IN_WINDOW_GRAPH : {3580,[]string{"HY000"},"There is a circularity in the window dependency graph."},
	ER_WINDOW_NO_CHILD_PARTITIONING : {3581,[]string{"HY000"},"A window which depends on another cannot define partitioning."},
	ER_WINDOW_NO_INHERIT_FRAME : {3582,[]string{"HY000"},"Window '%s' has a frame definition, so cannot be referenced by another window."},
	ER_WINDOW_NO_REDEFINE_ORDER_BY : {3583,[]string{"HY000"},"Window '%s' cannot inherit '%s' since both contain an ORDER BY clause."},
	ER_WINDOW_FRAME_START_ILLEGAL : {3584,[]string{"HY000"},"Window '%s': frame start cannot be UNBOUNDED FOLLOWING."},
	ER_WINDOW_FRAME_END_ILLEGAL : {3585,[]string{"HY000"},"Window '%s': frame end cannot be UNBOUNDED PRECEDING."},
	ER_WINDOW_FRAME_ILLEGAL : {3586,[]string{"HY000"},"Window '%s': frame start or end is negative, NULL or of non-integral type"},
	ER_WINDOW_RANGE_FRAME_ORDER_TYPE : {3587,[]string{"HY000"},"Window '%s' with RANGE N PRECEDING/FOLLOWING frame requires exactly one ORDER BY expression, of numeric or temporal type"},
	ER_WINDOW_RANGE_FRAME_TEMPORAL_TYPE : {3588,[]string{"HY000"},"Window '%s' with RANGE frame has ORDER BY expression of datetime type. Only INTERVAL bound value allowed."},
	ER_WINDOW_RANGE_FRAME_NUMERIC_TYPE : {3589,[]string{"HY000"},"Window '%s' with RANGE frame has ORDER BY expression of numeric type, INTERVAL bound value not allowed."},
	ER_WINDOW_RANGE_BOUND_NOT_CONSTANT : {3590,[]string{"HY000"},"Window '%s' has a non-constant frame bound."},
	ER_WINDOW_DUPLICATE_NAME : {3591,[]string{"HY000"},"Window '%s' is defined twice."},
	ER_WINDOW_ILLEGAL_ORDER_BY : {3592,[]string{"HY000"},"Window '%s': ORDER BY or PARTITION BY uses legacy position indication which is not supported, use expression."},
	ER_WINDOW_INVALID_WINDOW_FUNC_USE : {3593,[]string{"HY000"},"You cannot use the window function '%s' in this context.'"},
	ER_WINDOW_INVALID_WINDOW_FUNC_ALIAS_USE : {3594,[]string{"HY000"},"You cannot use the alias '%s' of an expression containing a window function in this context.'"},
	ER_WINDOW_NESTED_WINDOW_FUNC_USE_IN_WINDOW_SPEC : {3595,[]string{"HY000"},"You cannot nest a window function in the specification of window '%s'."},
	ER_WINDOW_ROWS_INTERVAL_USE : {3596,[]string{"HY000"},"Window '%s': INTERVAL can only be used with RANGE frames."},
	ER_WINDOW_NO_GROUP_ORDER_UNUSED : {3597,[]string{"HY000"},"ASC or DESC with GROUP BY isn't allowed with window functions; put ASC or DESC in ORDER BY"},
	ER_WINDOW_EXPLAIN_JSON : {3598,[]string{"HY000"},"To get information about window functions use EXPLAIN FORMAT=JSON"},
	ER_WINDOW_FUNCTION_IGNORES_FRAME : {3599,[]string{"HY000"},"Window function '%s' ignores the frame clause of window '%s' and aggregates over the whole partition"},
	ER_WL9236_NOW_UNUSED : {3600,[]string{"HY000"},"Windowing requires @@internal_tmp_mem_storage_engine=TempTable."},
	ER_INVALID_NO_OF_ARGS : {3601,[]string{"HY000"},"Too many arguments for function %s: %lu; maximum allowed is %s."},
	ER_FIELD_IN_GROUPING_NOT_GROUP_BY : {3602,[]string{"HY000"},"Argument #%u of GROUPING function is not in GROUP BY"},
	ER_TOO_LONG_TABLESPACE_COMMENT : {3603,[]string{"HY000"},"Comment for tablespace '%-.64s' is too long (max = %lu)"},
	ER_ENGINE_CANT_DROP_TABLE : {3604,[]string{"HY000"},"Storage engine can't drop table '%-.129s'"},
	ER_ENGINE_CANT_DROP_MISSING_TABLE : {3605,[]string{"HY000"},"Storage engine can't drop table '%-.129s' because it is missing. Use DROP TABLE IF EXISTS to remove it from data-dictionary."},
	ER_TABLESPACE_DUP_FILENAME : {3606,[]string{"HY000"},"Duplicate file name for tablespace '%-.64s'"},
	ER_DB_DROP_RMDIR2 : {3607,[]string{"HY000"},"Problem while dropping database. Can't remove database directory (%s). Please remove it manually."},
	ER_IMP_NO_FILES_MATCHED : {3608,[]string{"HY000"},"No SDI files matched the pattern '%s'"},
	ER_IMP_SCHEMA_DOES_NOT_EXIST : {3609,[]string{"HY000"},"Schema '%.256s', referenced in SDI, does not exist."},
	ER_IMP_TABLE_ALREADY_EXISTS : {3610,[]string{"HY000"},"Table '%.256s.%.256s', referenced in SDI, already exists."},
	ER_IMP_INCOMPATIBLE_MYSQLD_VERSION : {3611,[]string{"HY000"},"Imported mysqld_version (%llu) is not compatible with current (%llu)"},
	ER_IMP_INCOMPATIBLE_DD_VERSION : {3612,[]string{"HY000"},"Imported dd version (%u) is not compatible with current (%u)"},
	ER_IMP_INCOMPATIBLE_SDI_VERSION : {3613,[]string{"HY000"},"Imported sdi version (%llu) is not compatible with current (%llu)"},
	ER_WARN_INVALID_HINT : {3614,[]string{"HY000"},"Invalid number of arguments for hint %s"},
	ER_VAR_DOES_NOT_EXIST : {3615,[]string{"HY000"},"Variable %s does not exist in persisted config file"},
	ER_LONGITUDE_OUT_OF_RANGE : {3616,[]string{"22S02"},"Longitude %f is out of range in function %.192s. It must be within (%f, %f]."},
	ER_LATITUDE_OUT_OF_RANGE : {3617,[]string{"22S03"},"Latitude %f is out of range in function %.192s. It must be within [%f, %f]."},
	ER_NOT_IMPLEMENTED_FOR_GEOGRAPHIC_SRS : {3618,[]string{"22S00"},"%.192s(%.80s) has not been implemented for geographic spatial reference systems."},
	ER_ILLEGAL_PRIVILEGE_LEVEL : {3619,[]string{"HY000"},"Illegal privilege level specified for %s"},
	ER_NO_SYSTEM_VIEW_ACCESS : {3620,[]string{"HY000"},"Access to system view INFORMATION_SCHEMA.'%.64s' is rejected."},
	ER_COMPONENT_FILTER_FLABBERGASTED : {3621,[]string{"HY000"},"The log-filter component \"%s\" got confused at \"%s\" ..."},
	ER_PART_EXPR_TOO_LONG : {3622,[]string{"HY000"},"Partitioning expression is too long."},
	ER_UDF_DROP_DYNAMICALLY_REGISTERED : {3623,[]string{"HY000"},"DROP FUNCTION can't drop a dynamically registered user defined function"},
	ER_UNABLE_TO_STORE_COLUMN_STATISTICS : {3624,[]string{"HY000"},"Unable to store column statistics for column '%.64s' in table '%.64s'.'%.64s'"},
	ER_UNABLE_TO_UPDATE_COLUMN_STATISTICS : {3625,[]string{"HY000"},"Unable to update column statistics for column '%.64s' in table '%.64s'.'%.64s'"},
	ER_UNABLE_TO_DROP_COLUMN_STATISTICS : {3626,[]string{"HY000"},"Unable to remove column statistics for column '%.64s' in table '%.64s'.'%.64s'"},
	ER_UNABLE_TO_BUILD_HISTOGRAM : {3627,[]string{"HY000"},"Unable to build histogram statistics for column '%.64s' in table '%.64s'.'%.64s'"},
	ER_MANDATORY_ROLE : {3628,[]string{"HY000"},"The role %s is a mandatory role and can't be revoked or dropped. The restriction can be lifted by excluding the role identifier from the global variable mandatory_roles."},
	ER_MISSING_TABLESPACE_FILE : {3629,[]string{"HY000"},"Tablespace '%.192s' does not have a file named '%.192s'"},
	ER_PERSIST_ONLY_ACCESS_DENIED_ERROR : {3630,[]string{"42000"},"Access denied; you need %-.128s privileges for this operation"},
	ER_CMD_NEED_SUPER : {3631,[]string{"HY000"},"You need the SUPER privilege for command '%-.192s'"},
	ER_PATH_IN_DATADIR : {3632,[]string{"HY000"},"Path is within the current data directory '%-.192s'"},
	ER_CLONE_DDL_IN_PROGRESS : {3633,[]string{"HY000"},"Concurrent DDL is performed during clone operation. Please try again."},
	ER_CLONE_TOO_MANY_CONCURRENT_CLONES : {3634,[]string{"HY000"},"Too many concurrent clone operations. Maximum allowed - %d."},
	ER_APPLIER_LOG_EVENT_VALIDATION_ERROR : {3635,[]string{"HY000"},"The table in transaction %s does not comply with the requirements by an external plugin."},
	ER_CTE_MAX_RECURSION_DEPTH : {3636,[]string{"HY000"},"Recursive query aborted after %u iterations. Try increasing @@cte_max_recursion_depth to a larger value."},
	ER_NOT_HINT_UPDATABLE_VARIABLE : {3637,[]string{"HY000"},"Variable %s cannot be set using SET_VAR hint."},
	ER_CREDENTIALS_CONTRADICT_TO_HISTORY : {3638,[]string{"HY000"},"Cannot use these credentials for '%.*s@%.*s' because they contradict the password history policy"},
	ER_WARNING_PASSWORD_HISTORY_CLAUSES_VOID : {3639,[]string{"HY000"},"Non-zero password history clauses ignored for user '%s'@'%s' as its authentication plugin %s does not support password history"},
	ER_CLIENT_DOES_NOT_SUPPORT : {3640,[]string{"HY000"},"The client doesn't support %s"},
	ER_I_S_SKIPPED_TABLESPACE : {3641,[]string{"HY000"},"Tablespace '%s' was skipped since its definition is being modified by concurrent DDL statement"},
	ER_TABLESPACE_ENGINE_MISMATCH : {3642,[]string{"HY000"},"Engine '%.192s' does not match stored engine '%.192s' for tablespace '%.192s'"},
	ER_WRONG_SRID_FOR_COLUMN : {3643,[]string{"HY000"},"The SRID of the geometry does not match the SRID of the column '%.64s'. The SRID of the geometry is %lu, but the SRID of the column is %lu. Consider changing the SRID of the geometry or the SRID property of the column."},
	ER_CANNOT_ALTER_SRID_DUE_TO_INDEX : {3644,[]string{"HY000"},"The SRID specification on the column '%.64s' cannot be changed because there is a spatial index on the column. Please remove the spatial index before altering the SRID specification."},
	ER_WARN_BINLOG_PARTIAL_UPDATES_DISABLED : {3645,[]string{"HY000"},"When %.192s, the option binlog_row_value_options=%.192s will be ignored and updates will be written in full format to binary log."},
	ER_WARN_BINLOG_V1_ROW_EVENTS_DISABLED : {3646,[]string{"HY000"},"When %.192s, the option log_bin_use_v1_row_events=1 will be ignored and row events will be written in new format to binary log."},
	ER_WARN_BINLOG_PARTIAL_UPDATES_SUGGESTS_PARTIAL_IMAGES : {3647,[]string{"HY000"},"When %.192s, the option binlog_row_value_options=%.192s will be used only for the after-image. Full values will be written in the before-image, so the saving in disk space due to binlog_row_value_options is limited to less than 50%%."},
	ER_COULD_NOT_APPLY_JSON_DIFF : {3648,[]string{"HY000"},"Could not apply JSON diff in table %.*s, column %s."},
	ER_CORRUPTED_JSON_DIFF : {3649,[]string{"HY000"},"Corrupted JSON diff for table %.*s, column %s."},
	ER_RESOURCE_GROUP_EXISTS : {3650,[]string{"HY000"},"Resource Group '%-.192s' exists"},
	ER_RESOURCE_GROUP_NOT_EXISTS : {3651,[]string{"HY000"},"Resource Group '%-.192s' does not exist."},
	ER_INVALID_VCPU_ID : {3652,[]string{"HY000"},"Invalid cpu id %u"},
	ER_INVALID_VCPU_RANGE : {3653,[]string{"HY000"},"Invalid VCPU range %u-%u"},
	ER_INVALID_THREAD_PRIORITY : {3654,[]string{"HY000"},"Invalid thread priority value %d for %s resource group %s. Allowed range is [%d, %d]."},
	ER_DISALLOWED_OPERATION : {3655,[]string{"HY000"},"%s operation is disallowed on %s"},
	ER_RESOURCE_GROUP_BUSY : {3656,[]string{"HY000"},"Resource group %s is busy."},
	ER_RESOURCE_GROUP_DISABLED : {3657,[]string{"HY000"},"Resource group %s is disabled."},
	ER_FEATURE_UNSUPPORTED : {3658,[]string{"HY000"},"Feature %s is unsupported (%s)."},
	ER_ATTRIBUTE_IGNORED : {3659,[]string{"HY000"},"Attribute %s is ignored (%s)."},
	ER_INVALID_THREAD_ID : {3660,[]string{"HY000"},"Invalid thread id (%llu)."},
	ER_RESOURCE_GROUP_BIND_FAILED : {3661,[]string{"HY000"},"Unable to bind resource group %s with thread id (%llu).(%s)."},
	ER_INVALID_USE_OF_FORCE_OPTION : {3662,[]string{"HY000"},"Option FORCE invalid as DISABLE option is not specified."},
	ER_GROUP_REPLICATION_COMMAND_FAILURE : {3663,[]string{"HY000"},"The %s command encountered a failure. %s"},
	ER_SDI_OPERATION_FAILED : {3664,[]string{"HY000"},"Failed to %s SDI '%.192s.%.192s' in tablespace '%.192s'."},
	ER_MISSING_JSON_TABLE_VALUE : {3665,[]string{"22035"},"Missing value for JSON_TABLE column '%.192s'"},
	ER_WRONG_JSON_TABLE_VALUE : {3666,[]string{"2203F"},"Can't store an array or an object in the scalar JSON_TABLE column '%.192s'"},
	ER_TF_MUST_HAVE_ALIAS : {3667,[]string{"42000"},"Every table function must have an alias"},
	ER_TF_FORBIDDEN_JOIN_TYPE : {3668,[]string{"HY000"},"INNER or LEFT JOIN must be used for LATERAL references made by '%.192s'"},
	ER_JT_VALUE_OUT_OF_RANGE : {3669,[]string{"22003"},"Value is out of range for JSON_TABLE's column '%.192s'"},
	ER_JT_MAX_NESTED_PATH : {3670,[]string{"42000"},"More than supported %u NESTED PATHs were found in JSON_TABLE '%.192s'"},
	ER_PASSWORD_EXPIRATION_NOT_SUPPORTED_BY_AUTH_METHOD : {3671,[]string{"HY000"},"The selected authentication method %.*s does not support password expiration"},
	ER_INVALID_GEOJSON_CRS_NOT_TOP_LEVEL : {3672,[]string{"HY000"},"Invalid GeoJSON data provided to function %s: Member 'crs' must be specified in the top level object."},
	ER_BAD_NULL_ERROR_NOT_IGNORED : {3673,[]string{"23000"},"Column '%-.192s' cannot be null"},
	WARN_USELESS_SPATIAL_INDEX : {3674,[]string{"HY000"},"The spatial index on column '%.64s' will not be used by the query optimizer since the column does not have an SRID attribute. Consider adding an SRID attribute to the column."},
	ER_DISK_FULL_NOWAIT : {3675,[]string{"HY000"},"Create table/tablespace '%-.192s' failed, as disk is full"},
	ER_PARSE_ERROR_IN_DIGEST_FN : {3676,[]string{"HY000"},"Could not parse argument to digest function: \"%s\"."},
	ER_UNDISCLOSED_PARSE_ERROR_IN_DIGEST_FN : {3677,[]string{"HY000"},"Could not parse argument to digest function."},
	ER_SCHEMA_DIR_EXISTS : {3678,[]string{"HY000"},"Schema directory '%.192s' already exists. This must be resolved manually (e.g. by moving the schema directory to another location)."},
	ER_SCHEMA_DIR_MISSING : {3679,[]string{"HY000"},"Schema directory '%.192s' does not exist"},
	ER_SCHEMA_DIR_CREATE_FAILED : {3680,[]string{"HY000"},"Failed to create schema directory '%.192s' (errno: %d - %s)"},
	ER_SCHEMA_DIR_UNKNOWN : {3681,[]string{"HY000"},"Schema '%.192s' does not exist, but schema directory '%.192s' was found. This must be resolved manually (e.g. by moving the schema directory to another location)."},
	ER_ONLY_IMPLEMENTED_FOR_SRID_0_AND_4326 : {3682,[]string{"22S00"},"Function %.192s is only defined for SRID 0 and SRID 4326."},
	ER_BINLOG_EXPIRE_LOG_DAYS_AND_SECS_USED_TOGETHER : {3683,[]string{"HY000"},"The option expire_logs_days and binlog_expire_logs_seconds cannot be used together. Please use binlog_expire_logs_seconds to set the expire time (expire_logs_days is deprecated)"},
	ER_REGEXP_BUFFER_OVERFLOW : {3684,[]string{"HY000"},"The result string is larger than the result buffer."},
	ER_REGEXP_ILLEGAL_ARGUMENT : {3685,[]string{"HY000"},"Illegal argument to a regular expression."},
	ER_REGEXP_INDEX_OUTOFBOUNDS_ERROR : {3686,[]string{"HY000"},"Index out of bounds in regular expression search."},
	ER_REGEXP_INTERNAL_ERROR : {3687,[]string{"HY000"},"Internal error in the regular expression library."},
	ER_REGEXP_RULE_SYNTAX : {3688,[]string{"HY000"},"Syntax error in regular expression on line %u, character %u."},
	ER_REGEXP_BAD_ESCAPE_SEQUENCE : {3689,[]string{"HY000"},"Unrecognized escape sequence in regular expression."},
	ER_REGEXP_UNIMPLEMENTED : {3690,[]string{"HY000"},"The regular expression contains a feature that is not implemented in this library version."},
	ER_REGEXP_MISMATCHED_PAREN : {3691,[]string{"HY000"},"Mismatched parenthesis in regular expression."},
	ER_REGEXP_BAD_INTERVAL : {3692,[]string{"HY000"},"Incorrect description of a {min,max} interval."},
	ER_REGEXP_MAX_LT_MIN : {3693,[]string{"HY000"},"The maximum is less than the minumum in a {min,max} interval."},
	ER_REGEXP_INVALID_BACK_REF : {3694,[]string{"HY000"},"Invalid back-reference in regular expression."},
	ER_REGEXP_LOOK_BEHIND_LIMIT : {3695,[]string{"HY000"},"The look-behind assertion exceeds the limit in regular expression."},
	ER_REGEXP_MISSING_CLOSE_BRACKET : {3696,[]string{"HY000"},"The regular expression contains an unclosed bracket expression."},
	ER_REGEXP_INVALID_RANGE : {3697,[]string{"HY000"},"The regular expression contains an [x-y] character range where x comes after y."},
	ER_REGEXP_STACK_OVERFLOW : {3698,[]string{"HY000"},"Overflow in the regular expression backtrack stack."},
	ER_REGEXP_TIME_OUT : {3699,[]string{"HY000"},"Timeout exceeded in regular expression match."},
	ER_REGEXP_PATTERN_TOO_BIG : {3700,[]string{"HY000"},"The regular expression pattern exceeds limits on size or complexity."},
	ER_CANT_SET_ERROR_LOG_SERVICE : {3701,[]string{"HY000"},"Value for %s got confusing at or around \"%s\". Syntax may be wrong, component may not be INSTALLed, or a component that does not support instances may be listed more than once."},
	ER_EMPTY_PIPELINE_FOR_ERROR_LOG_SERVICE : {3702,[]string{"HY000"},"Setting an empty %s pipeline disables error logging!"},
	ER_COMPONENT_FILTER_DIAGNOSTICS : {3703,[]string{"HY000"},"filter %s: %s"},
	ER_NOT_IMPLEMENTED_FOR_CARTESIAN_SRS : {3704,[]string{"22S00"},"%.192s(%.80s) has not been implemented for Cartesian spatial reference systems."},
	ER_NOT_IMPLEMENTED_FOR_PROJECTED_SRS : {3705,[]string{"22S00"},"%.192s(%.80s) has not been implemented for projected spatial reference systems."},
	ER_NONPOSITIVE_RADIUS : {3706,[]string{"22003"},"Invalid radius provided to function %s: Radius must be greater than zero."},
	ER_RESTART_SERVER_FAILED : {3707,[]string{"HY000"},"Restart server failed (%s)."},
	ER_SRS_MISSING_MANDATORY_ATTRIBUTE : {3708,[]string{"SR006"},"Missing mandatory attribute %s."},
	ER_SRS_MULTIPLE_ATTRIBUTE_DEFINITIONS : {3709,[]string{"SR006"},"Multiple definitions of attribute %s."},
	ER_SRS_NAME_CANT_BE_EMPTY_OR_WHITESPACE : {3710,[]string{"SR006"},"The spatial reference system name can't be an empty string or start or end with whitespace."},
	ER_SRS_ORGANIZATION_CANT_BE_EMPTY_OR_WHITESPACE : {3711,[]string{"SR006"},"The organization name can't be an empty string or start or end with whitespace."},
	ER_SRS_ID_ALREADY_EXISTS : {3712,[]string{"SR004"},"There is already a spatial reference system with SRID %u."},
	ER_WARN_SRS_ID_ALREADY_EXISTS : {3713,[]string{"01S00"},"There is already a spatial reference system with SRID %u."},
	ER_CANT_MODIFY_SRID_0 : {3714,[]string{"SR000"},"SRID 0 is not modifiable."},
	ER_WARN_RESERVED_SRID_RANGE : {3715,[]string{"01S01"},"The SRID range [%u, %u] has been reserved for system use. SRSs in this range may be added, modified or removed without warning during upgrade."},
	ER_CANT_MODIFY_SRS_USED_BY_COLUMN : {3716,[]string{"SR005"},"Can't modify SRID %u. There is at least one column depending on it."},
	ER_SRS_INVALID_CHARACTER_IN_ATTRIBUTE : {3717,[]string{"SR006"},"Invalid character in attribute %s."},
	ER_SRS_ATTRIBUTE_STRING_TOO_LONG : {3718,[]string{"SR006"},"Attribute %s is too long. The maximum length is %u characters."},
	ER_DEPRECATED_UTF8_ALIAS : {3719,[]string{"HY000"},"'utf8' is currently an alias for the character set UTF8MB3, but will be an alias for UTF8MB4 in a future release. Please consider using UTF8MB4 in order to be unambiguous."},
	ER_DEPRECATED_NATIONAL : {3720,[]string{"HY000"},"NATIONAL/NCHAR/NVARCHAR implies the character set UTF8MB3, which will be replaced by UTF8MB4 in a future release. Please consider using CHAR(x) CHARACTER SET UTF8MB4 in order to be unambiguous."},
	ER_INVALID_DEFAULT_UTF8MB4_COLLATION : {3721,[]string{"HY000"},"Invalid default collation %s: utf8mb4_0900_ai_ci or utf8mb4_general_ci expected"},
	ER_UNABLE_TO_COLLECT_LOG_STATUS : {3722,[]string{"HY000"},"Unable to collect information for column '%-.192s': %-.192s."},
	ER_RESERVED_TABLESPACE_NAME : {3723,[]string{"HY000"},"The table '%-.192s' may not be created in the reserved tablespace '%-.192s'."},
	ER_UNABLE_TO_SET_OPTION : {3724,[]string{"HY000"},"This option cannot be set %s."},
	ER_SLAVE_POSSIBLY_DIVERGED_AFTER_DDL : {3725,[]string{"HY000"},"A commit for an atomic DDL statement was unsuccessful on the master and the slave. The slave supports atomic DDL statements but the master does not, so the action taken by the slave and master might differ. Check that their states have not diverged before proceeding."},
	ER_SRS_NOT_GEOGRAPHIC : {3726,[]string{"22S00"},"Function %s is only defined for geographic spatial reference systems, but one of its arguments is in SRID %u, which is not geographic."},
	ER_POLYGON_TOO_LARGE : {3727,[]string{"22023"},"Function %s encountered a polygon that was too large. Polygons must cover less than half the planet."},
	ER_SPATIAL_UNIQUE_INDEX : {3728,[]string{"HY000"},"Spatial indexes can't be primary or unique indexes."},
	ER_INDEX_TYPE_NOT_SUPPORTED_FOR_SPATIAL_INDEX : {3729,[]string{"HY000"},"The index type %.20s is not supported for spatial indexes."},
	ER_FK_CANNOT_DROP_PARENT : {3730,[]string{"HY000"},"Cannot drop table '%s' referenced by a foreign key constraint '%s' on table '%s'."},
	ER_GEOMETRY_PARAM_LONGITUDE_OUT_OF_RANGE : {3731,[]string{"22S02"},"A parameter of function %.192s contains a geometry with longitude %f, which is out of range. It must be within (%f, %f]."},
	ER_GEOMETRY_PARAM_LATITUDE_OUT_OF_RANGE : {3732,[]string{"22S03"},"A parameter of function %.192s contains a geometry with latitude %f, which is out of range. It must be within [%f, %f]."},
	ER_FK_CANNOT_USE_VIRTUAL_COLUMN : {3733,[]string{"HY000"},"Foreign key '%s' uses virtual column '%s' which is not supported."},
	ER_FK_NO_COLUMN_PARENT : {3734,[]string{"HY000"},"Failed to add the foreign key constraint. Missing column '%s' for constraint '%s' in the referenced table '%s'"},
	ER_CANT_SET_ERROR_SUPPRESSION_LIST : {3735,[]string{"HY000"},"%s: Could not add suppression rule for code \"%s\". Rule-set may be full, or code may not correspond to an error-log message."},
	ER_SRS_GEOGCS_INVALID_AXES : {3736,[]string{"SR002"},"The spatial reference system definition for SRID %u specifies invalid geographic axes '%.20s' and '%.20s'. One axis must be NORTH or SOUTH and the other must be EAST or WEST."},
	ER_SRS_INVALID_SEMI_MAJOR_AXIS : {3737,[]string{"SR002"},"The length of the semi-major axis must be a positive number."},
	ER_SRS_INVALID_INVERSE_FLATTENING : {3738,[]string{"SR002"},"The inverse flattening must be larger than 1.0, or 0.0 if the ellipsoid is a sphere."},
	ER_SRS_INVALID_ANGULAR_UNIT : {3739,[]string{"SR002"},"The angular unit conversion factor must be a positive number."},
	ER_SRS_INVALID_PRIME_MERIDIAN : {3740,[]string{"SR002"},"The prime meridian must be within (-180, 180] degrees, specified in the SRS angular unit."},
	ER_TRANSFORM_SOURCE_SRS_NOT_SUPPORTED : {3741,[]string{"22S00"},"Transformation from SRID %u is not supported."},
	ER_TRANSFORM_TARGET_SRS_NOT_SUPPORTED : {3742,[]string{"22S00"},"Transformation to SRID %u is not supported."},
	ER_TRANSFORM_SOURCE_SRS_MISSING_TOWGS84 : {3743,[]string{"22S00"},"Transformation from SRID %u is not supported. The spatial reference system has no TOWGS84 clause."},
	ER_TRANSFORM_TARGET_SRS_MISSING_TOWGS84 : {3744,[]string{"22S00"},"Transformation to SRID %u is not supported. The spatial reference system has no TOWGS84 clause."},
	ER_TEMP_TABLE_PREVENTS_SWITCH_SESSION_BINLOG_FORMAT : {3745,[]string{"HY000"},"Changing @@session.binlog_format is disallowed when the session has open temporary table(s). You could wait until these temporary table(s) are dropped and try again."},
	ER_TEMP_TABLE_PREVENTS_SWITCH_GLOBAL_BINLOG_FORMAT : {3746,[]string{"HY000"},"Changing @@global.binlog_format or @@persist.binlog_format is disallowed when any replication channel has open temporary table(s). You could wait until Slave_open_temp_tables = 0 and try again"},
	ER_RUNNING_APPLIER_PREVENTS_SWITCH_GLOBAL_BINLOG_FORMAT : {3747,[]string{"HY000"},"Changing @@global.binlog_format or @@persist.binlog_format is disallowed when any replication channel applier thread is running. You could execute STOP SLAVE SQL_THREAD and try again."},
	ER_CLIENT_GTID_UNSAFE_CREATE_DROP_TEMP_TABLE_IN_TRX_IN_SBR : {3748,[]string{"HY000"},"Statement violates GTID consistency: CREATE TEMPORARY TABLE and DROP TEMPORARY TABLE are not allowed inside a transaction or inside a procedure in a transactional context when @@session.binlog_format=STATEMENT."},
	//OBSOLETE_ER_XA_CANT_CREATE_MDL_BACKUP : {13477,[]string{"HY000"},"XA: Failed to take MDL Lock backup of PREPARED XA transaction during client disconnect."},
	ER_TABLE_WITHOUT_PK : {3750,[]string{"HY000"},"Unable to create or change a table without a primary key, when the system variable 'sql_require_primary_key' is set. Add a primary key to the table or unset this variable to avoid this message. Note that tables without a primary key can cause performance problems in row-based replication, so please consult your DBA before changing this setting."},
	ER_WARN_DATA_TRUNCATED_FUNCTIONAL_INDEX : {3751,[]string{"01000"},"Data truncated for functional index '%s' at row %ld"},
	ER_WARN_DATA_OUT_OF_RANGE_FUNCTIONAL_INDEX : {3752,[]string{"22003"},"Value is out of range for functional index '%s' at row %ld"},
	ER_FUNCTIONAL_INDEX_ON_JSON_OR_GEOMETRY_FUNCTION : {3753,[]string{"42000"},"Cannot create a functional index on a function that returns a JSON or GEOMETRY value."},
	ER_FUNCTIONAL_INDEX_REF_AUTO_INCREMENT : {3754,[]string{"HY000"},"Functional index '%.64s' cannot refer to an auto-increment column."},
	ER_CANNOT_DROP_COLUMN_FUNCTIONAL_INDEX : {3755,[]string{"HY000"},"Cannot drop column '%-.64s' because it is used by a functional index. In order to drop the column, you must remove the functional index."},
	ER_FUNCTIONAL_INDEX_PRIMARY_KEY : {3756,[]string{"HY000"},"The primary key cannot be a functional index"},
	ER_FUNCTIONAL_INDEX_ON_LOB : {3757,[]string{"HY000"},"Cannot create a functional index on an expression that returns a BLOB or TEXT. Please consider using CAST."},
	ER_FUNCTIONAL_INDEX_FUNCTION_IS_NOT_ALLOWED : {3758,[]string{"HY000"},"Expression of functional index '%s' contains a disallowed function."},
	ER_FULLTEXT_FUNCTIONAL_INDEX : {3759,[]string{"HY000"},"Fulltext functional index is not supported."},
	ER_SPATIAL_FUNCTIONAL_INDEX : {3760,[]string{"HY000"},"Spatial functional index is not supported."},
	ER_WRONG_KEY_COLUMN_FUNCTIONAL_INDEX : {3761,[]string{"HY000"},"The used storage engine cannot index the expression '%s'."},
	ER_FUNCTIONAL_INDEX_ON_FIELD : {3762,[]string{"HY000"},"Functional index on a column is not supported. Consider using a regular index instead."},
	ER_GENERATED_COLUMN_NAMED_FUNCTION_IS_NOT_ALLOWED : {3763,[]string{"HY000"},"Expression of generated column '%s' contains a disallowed function: %s."},
	ER_GENERATED_COLUMN_ROW_VALUE : {3764,[]string{"HY000"},"Expression of generated column '%s' cannot refer to a row value."},
	ER_GENERATED_COLUMN_VARIABLES : {3765,[]string{"HY000"},"Expression of generated column '%s' cannot refer user or system variables."},
	ER_DEPENDENT_BY_DEFAULT_GENERATED_VALUE : {3766,[]string{"HY000"},"Column '%s' of table '%s' has a default value expression dependency and cannot be dropped or renamed."},
	ER_DEFAULT_VAL_GENERATED_NON_PRIOR : {3767,[]string{"HY000"},"Default value expression of column '%s' cannot refer to a column defined after it if that column is a generated column or has an expression as default value."},
	ER_DEFAULT_VAL_GENERATED_REF_AUTO_INC : {3768,[]string{"HY000"},"Default value expression of column '%s' cannot refer to an auto-increment column."},
	ER_DEFAULT_VAL_GENERATED_FUNCTION_IS_NOT_ALLOWED : {3769,[]string{"HY000"},"Default value expression of column '%s' contains a disallowed function."},
	ER_DEFAULT_VAL_GENERATED_NAMED_FUNCTION_IS_NOT_ALLOWED : {3770,[]string{"HY000"},"Default value expression of column '%s' contains a disallowed function: %s."},
	ER_DEFAULT_VAL_GENERATED_ROW_VALUE : {3771,[]string{"HY000"},"Default value expression of column '%s' cannot refer to a row value."},
	ER_DEFAULT_VAL_GENERATED_VARIABLES : {3772,[]string{"HY000"},"Default value expression of column '%s' cannot refer user or system variables."},
	ER_DEFAULT_AS_VAL_GENERATED : {3773,[]string{"HY000"},"DEFAULT function cannot be used with default value expressions"},
	ER_UNSUPPORTED_ACTION_ON_DEFAULT_VAL_GENERATED : {3774,[]string{"HY000"},"'%s' is not supported for default value expressions."},
	ER_GTID_UNSAFE_ALTER_ADD_COL_WITH_DEFAULT_EXPRESSION : {3775,[]string{"HY000"},"Statement violates GTID consistency: ALTER TABLE ... ADD COLUMN .. with expression as DEFAULT."},
	ER_FK_CANNOT_CHANGE_ENGINE : {3776,[]string{"HY000"},"Cannot change table's storage engine because the table participates in a foreign key constraint."},
	ER_WARN_DEPRECATED_USER_SET_EXPR : {3777,[]string{"HY000"},"Setting user variables within expressions is deprecated and will be removed in a future release. Consider alternatives: 'SET variable=expression, ...', or 'SELECT expression(s) INTO variables(s)'."},
	ER_WARN_DEPRECATED_UTF8MB3_COLLATION : {3778,[]string{"HY000"},"'%-.64s' is a collation of the deprecated character set UTF8MB3. Please consider using UTF8MB4 with an appropriate collation instead."},
	ER_WARN_DEPRECATED_NESTED_COMMENT_SYNTAX : {3779,[]string{"HY000"},"Nested comment syntax is deprecated and will be removed in a future release."},
	ER_FK_INCOMPATIBLE_COLUMNS : {3780,[]string{"HY000"},"Referencing column '%s' and referenced column '%s' in foreign key constraint '%s' are incompatible."},
	ER_GR_HOLD_WAIT_TIMEOUT : {3781,[]string{"HY000"},"Timeout exceeded for held statement while new Group Replication primary member is applying backlog."},
	ER_GR_HOLD_KILLED : {3782,[]string{"HY000"},"Held statement aborted because Group Replication plugin got shut down or thread was killed while new primary member was applying backlog."},
	ER_GR_HOLD_MEMBER_STATUS_ERROR : {3783,[]string{"HY000"},"Held statement was aborted due to member being in error state, while backlog is being applied during Group Replication primary election."},
	ER_RPL_ENCRYPTION_FAILED_TO_FETCH_KEY : {3784,[]string{"HY000"},"Failed to fetch key from keyring, please check if keyring plugin is loaded."},
	ER_RPL_ENCRYPTION_KEY_NOT_FOUND : {3785,[]string{"HY000"},"Can't find key from keyring, please check in the server log if a keyring plugin is loaded and initialized successfully."},
	ER_RPL_ENCRYPTION_KEYRING_INVALID_KEY : {3786,[]string{"HY000"},"Fetched an invalid key from keyring."},
	ER_RPL_ENCRYPTION_HEADER_ERROR : {3787,[]string{"HY000"},"Error reading a replication log encryption header: %s."},
	ER_RPL_ENCRYPTION_FAILED_TO_ROTATE_LOGS : {3788,[]string{"HY000"},"Failed to rotate some logs after changing binlog encryption settings. Please fix the problem and rotate the logs manually."},
	ER_RPL_ENCRYPTION_KEY_EXISTS_UNEXPECTED : {3789,[]string{"HY000"},"Key %s exists unexpected."},
	ER_RPL_ENCRYPTION_FAILED_TO_GENERATE_KEY : {3790,[]string{"HY000"},"Failed to generate key, please check if keyring plugin is loaded."},
	ER_RPL_ENCRYPTION_FAILED_TO_STORE_KEY : {3791,[]string{"HY000"},"Failed to store key, please check if keyring plugin is loaded."},
	ER_RPL_ENCRYPTION_FAILED_TO_REMOVE_KEY : {3792,[]string{"HY000"},"Failed to remove key, please check if keyring plugin is loaded."},
	ER_RPL_ENCRYPTION_UNABLE_TO_CHANGE_OPTION : {3793,[]string{"HY000"},"Failed to change binlog_encryption value. %-.80s."},
	ER_RPL_ENCRYPTION_MASTER_KEY_RECOVERY_FAILED : {3794,[]string{"HY000"},"Unable to recover binlog encryption master key, please check if keyring plugin is loaded."},
	ER_SLOW_LOG_MODE_IGNORED_WHEN_NOT_LOGGING_TO_FILE : {3795,[]string{"HY000"},"slow query log file format changed as requested, but setting will have no effect when not actually logging to a file."},
	ER_GRP_TRX_CONSISTENCY_NOT_ALLOWED : {3796,[]string{"HY000"},"The option group_replication_consistency cannot be used on the current member state."},
	ER_GRP_TRX_CONSISTENCY_BEFORE : {3797,[]string{"HY000"},"Error while waiting for group transactions commit on group_replication_consistency= 'BEFORE'."},
	ER_GRP_TRX_CONSISTENCY_AFTER_ON_TRX_BEGIN : {3798,[]string{"HY000"},"Error while waiting for transactions with group_replication_consistency= 'AFTER' to commit."},
	ER_GRP_TRX_CONSISTENCY_BEGIN_NOT_ALLOWED : {3799,[]string{"HY000"},"The Group Replication plugin is stopping, therefore new transactions are not allowed to start."},
	ER_FUNCTIONAL_INDEX_ROW_VALUE_IS_NOT_ALLOWED : {3800,[]string{"HY000"},"Expression of functional index '%s' cannot refer to a row value."},
	ER_RPL_ENCRYPTION_FAILED_TO_ENCRYPT : {3801,[]string{"HY000"},"Failed to encrypt content to write into binlog file: %s."},
	ER_PAGE_TRACKING_NOT_STARTED : {3802,[]string{"HY000"},"Page Tracking is not started yet."},
	ER_PAGE_TRACKING_RANGE_NOT_TRACKED : {3803,[]string{"HY000"},"Tracking was not enabled for the LSN range specified"},
	ER_PAGE_TRACKING_CANNOT_PURGE : {3804,[]string{"HY000"},"Cannot purge data when concurrent clone is in progress. Try later."},
	ER_RPL_ENCRYPTION_CANNOT_ROTATE_BINLOG_MASTER_KEY : {3805,[]string{"HY000"},"Cannot rotate binary log master key when 'binlog-encryption' is off."},
	ER_BINLOG_MASTER_KEY_RECOVERY_OUT_OF_COMBINATION : {3806,[]string{"HY000"},"Unable to recover binary log master key, the combination of new_master_key_seqno=%u, master_key_seqno=%u and old_master_key_seqno=%u are wrong."},
	ER_BINLOG_MASTER_KEY_ROTATION_FAIL_TO_OPERATE_KEY : {3807,[]string{"HY000"},"Failed to operate binary log master key on keyring, please check if keyring plugin is loaded. The statement had no effect: the old binary log master key is still in use, the keyring, binary and relay log files are unchanged, and the server could not start using a new binary log master key for encrypting new binary and relay log files."},
	ER_BINLOG_MASTER_KEY_ROTATION_FAIL_TO_ROTATE_LOGS : {3808,[]string{"HY000"},"Failed to rotate one or more binary or relay log files. A new binary log master key was generated and will be used to encrypt new binary and relay log files. There may still exist binary or relay log files using the previous binary log master key."},
	ER_BINLOG_MASTER_KEY_ROTATION_FAIL_TO_REENCRYPT_LOG : {3809,[]string{"HY000"},"%s. A new binary log master key was generated and will be used to encrypt new binary and relay log files. There may still exist binary or relay log files using the previous binary log master key."},
	ER_BINLOG_MASTER_KEY_ROTATION_FAIL_TO_CLEANUP_UNUSED_KEYS : {3810,[]string{"HY000"},"Failed to remove unused binary log encryption keys from the keyring, please check if keyring plugin is loaded. The unused binary log encryption keys may still exist in the keyring, and they will be removed upon server restart or next 'ALTER INSTANCE ROTATE BINLOG MASTER KEY' execution."},
	ER_BINLOG_MASTER_KEY_ROTATION_FAIL_TO_CLEANUP_AUX_KEY : {3811,[]string{"HY000"},"Failed to remove auxiliary binary log encryption key from keyring, please check if keyring plugin is loaded. The cleanup of the binary log master key rotation process did not finish as expected and the cleanup will take place upon server restart or next 'ALTER INSTANCE ROTATE BINLOG MASTER KEY' execution."},
	ER_NON_BOOLEAN_EXPR_FOR_CHECK_CONSTRAINT : {3812,[]string{"HY000"},"An expression of non-boolean type specified to a check constraint '%-.192s'."},
	ER_COLUMN_CHECK_CONSTRAINT_REFERENCES_OTHER_COLUMN : {3813,[]string{"HY000"},"Column check constraint '%-.192s' references other column."},
	ER_CHECK_CONSTRAINT_NAMED_FUNCTION_IS_NOT_ALLOWED : {3814,[]string{"HY000"},"An expression of a check constraint '%-.192s' contains disallowed function: %s."},
	ER_CHECK_CONSTRAINT_FUNCTION_IS_NOT_ALLOWED : {3815,[]string{"HY000"},"An expression of a check constraint '%-.192s' contains disallowed function."},
	ER_CHECK_CONSTRAINT_VARIABLES : {3816,[]string{"HY000"},"An expression of a check constraint '%-.192s' cannot refer to a user or system variable."},
	ER_CHECK_CONSTRAINT_ROW_VALUE : {3817,[]string{"HY000"},"Check constraint '%-.192s' cannot refer to a row value."},
	ER_CHECK_CONSTRAINT_REFERS_AUTO_INCREMENT_COLUMN : {3818,[]string{"HY000"},"Check constraint '%-.192s' cannot refer to an auto-increment column."},
	ER_CHECK_CONSTRAINT_VIOLATED : {3819,[]string{"HY000"},"Check constraint '%-.192s' is violated."},
	ER_CHECK_CONSTRAINT_REFERS_UNKNOWN_COLUMN : {3820,[]string{"HY000"},"Check constraint '%-.192s' refers to non-existing column '%-.192s'."},
	ER_CHECK_CONSTRAINT_NOT_FOUND : {3821,[]string{"HY000"},"Check constraint '%-.192s' is not found in the table."},
	ER_CHECK_CONSTRAINT_DUP_NAME : {3822,[]string{"HY000"},"Duplicate check constraint name '%-.192s'."},
	ER_CHECK_CONSTRAINT_CLAUSE_USING_FK_REFER_ACTION_COLUMN : {3823,[]string{"HY000"},"Column '%-.192s' cannot be used in a check constraint '%-.192s': needed in a foreign key constraint '%-.192s' referential action."},
	WARN_UNENCRYPTED_TABLE_IN_ENCRYPTED_DB : {3824,[]string{"HY000"},"Creating an unencrypted table in a database with default encryption enabled."},
	ER_INVALID_ENCRYPTION_REQUEST : {3825,[]string{"HY000"},"Request to create %s table while using an %s tablespace."},
	ER_CANNOT_SET_TABLE_ENCRYPTION : {3826,[]string{"HY000"},"Table encryption differ from its database default encryption, and user doesn't have enough privilege."},
	ER_CANNOT_SET_DATABASE_ENCRYPTION : {3827,[]string{"HY000"},"Database default encryption differ from 'default_table_encryption' setting, and user doesn't have enough privilege."},
	ER_CANNOT_SET_TABLESPACE_ENCRYPTION : {3828,[]string{"HY000"},"Tablespace encryption differ from 'default_table_encryption' setting, and user doesn't have enough privilege."},
	ER_TABLESPACE_CANNOT_BE_ENCRYPTED : {3829,[]string{"HY000"},"This tablespace can't be encrypted, because one of table's schema has default encryption OFF and user doesn't have enough privilege."},
	ER_TABLESPACE_CANNOT_BE_DECRYPTED : {3830,[]string{"HY000"},"This tablespace can't be decrypted, because one of table's schema has default encryption ON and user doesn't have enough privilege."},
	ER_TABLESPACE_TYPE_UNKNOWN : {3831,[]string{"HY000"},"Cannot determine the type of the tablespace named '%s'."},
	ER_TARGET_TABLESPACE_UNENCRYPTED : {3832,[]string{"HY000"},"Source tablespace is encrypted but target tablespace is not."},
	ER_CANNOT_USE_ENCRYPTION_CLAUSE : {3833,[]string{"HY000"},"ENCRYPTION clause is not valid for %s tablespace."},
	ER_INVALID_MULTIPLE_CLAUSES : {3834,[]string{"HY000"},"Multiple %s clauses"},
	ER_UNSUPPORTED_USE_OF_GRANT_AS : {3835,[]string{"HY000"},"GRANT ... AS is currently supported only for global privileges."},
	ER_UKNOWN_AUTH_ID_OR_ACCESS_DENIED_FOR_GRANT_AS : {3836,[]string{"HY000"},"Either some of the authorization IDs in the AS clause are invalid or the current user lacks privileges to execute the statement."},
	ER_DEPENDENT_BY_FUNCTIONAL_INDEX : {3837,[]string{"HY000"},"Column '%s' has a functional index dependency and cannot be dropped or renamed."},
	ER_PLUGIN_NOT_EARLY : {3838,[]string{"HY000"},"Plugin '%s' is not to be used as an \"early\" plugin. Don't add it to --early-plugin-load, keyring migration etc."},
	ER_INNODB_REDO_LOG_ARCHIVE_START_SUBDIR_PATH : {3839,[]string{"HY000"},"Redo log archiving start prohibits path name in 'subdir' argument"},
	ER_INNODB_REDO_LOG_ARCHIVE_START_TIMEOUT : {3840,[]string{"HY000"},"Redo log archiving start timed out"},
	ER_INNODB_REDO_LOG_ARCHIVE_DIRS_INVALID : {3841,[]string{"HY000"},"Server variable 'innodb_redo_log_archive_dirs' is NULL or empty"},
	ER_INNODB_REDO_LOG_ARCHIVE_LABEL_NOT_FOUND : {3842,[]string{"HY000"},"Label '%.192s' not found in server variable 'innodb_redo_log_archive_dirs'"},
	ER_INNODB_REDO_LOG_ARCHIVE_DIR_EMPTY : {3843,[]string{"HY000"},"Directory is empty after label '%.192s' in server variable 'innodb_redo_log_archive_dirs'"},
	ER_INNODB_REDO_LOG_ARCHIVE_NO_SUCH_DIR : {3844,[]string{"HY000"},"Redo log archive directory '%.192s' does not exist or is not a directory"},
	ER_INNODB_REDO_LOG_ARCHIVE_DIR_CLASH : {3845,[]string{"HY000"},"Redo log archive directory '%.192s' is in, under, or over server directory '%.192s' - '%.192s'"},
	ER_INNODB_REDO_LOG_ARCHIVE_DIR_PERMISSIONS : {3846,[]string{"HY000"},"Redo log archive directory '%.192s' is accessible to all OS users"},
	ER_INNODB_REDO_LOG_ARCHIVE_FILE_CREATE : {3847,[]string{"HY000"},"Cannot create redo log archive file '%.512s' (OS errno: %d - %.128s)"},
	ER_INNODB_REDO_LOG_ARCHIVE_ACTIVE : {3848,[]string{"HY000"},"Redo log archiving has been started on '%.512s' - Call innodb_redo_log_archive_stop() first"},
	ER_INNODB_REDO_LOG_ARCHIVE_INACTIVE : {3849,[]string{"HY000"},"Redo log archiving is not active"},
	ER_INNODB_REDO_LOG_ARCHIVE_FAILED : {3850,[]string{"HY000"},"Redo log archiving failed: %.512s"},
	ER_INNODB_REDO_LOG_ARCHIVE_SESSION : {3851,[]string{"HY000"},"Redo log archiving has not been started by this session"},
	ER_STD_REGEX_ERROR : {3852,[]string{"HY000"},"Regex error: %-.256s in function %s."},
	ER_INVALID_JSON_TYPE : {3853,[]string{"22032"},"Invalid JSON type in argument %u to function %s; an %s is required."},
	ER_CANNOT_CONVERT_STRING : {3854,[]string{"HY000"},"Cannot convert string '%.64s' from %s to %s"},
	ER_DEPENDENT_BY_PARTITION_FUNC : {3855,[]string{"HY000"},"Column '%s' has a partitioning function dependency and cannot be dropped or renamed."},
	ER_WARN_DEPRECATED_FLOAT_AUTO_INCREMENT : {3856,[]string{"HY000"},"AUTO_INCREMENT support for FLOAT/DOUBLE columns is deprecated and will be removed in a future release. Consider removing AUTO_INCREMENT from column '%-.192s'."},
	ER_RPL_CANT_STOP_SLAVE_WHILE_LOCKED_BACKUP : {3857,[]string{"HY000"},"Cannot stop the slave SQL thread while the instance is locked for backup. Try running `UNLOCK INSTANCE` first."},
	ER_WARN_DEPRECATED_FLOAT_DIGITS : {3858,[]string{"HY000"},"Specifying number of digits for floating point data types is deprecated and will be removed in a future release."},
	ER_WARN_DEPRECATED_FLOAT_UNSIGNED : {3859,[]string{"HY000"},"UNSIGNED for decimal and floating point data types is deprecated and support for it will be removed in a future release."},
	ER_WARN_DEPRECATED_INTEGER_DISPLAY_WIDTH : {3860,[]string{"HY000"},"Integer display width is deprecated and will be removed in a future release."},
	ER_WARN_DEPRECATED_ZEROFILL : {3861,[]string{"HY000"},"The ZEROFILL attribute is deprecated and will be removed in a future release. Use the LPAD function to zero-pad numbers, or store the formatted numbers in a CHAR column."},
	ER_CLONE_DONOR : {3862,[]string{"HY000"},"Clone Donor Error: %.512s."},
	ER_CLONE_PROTOCOL : {3863,[]string{"HY000"},"Clone received unexpected response from Donor : %.512s."},
	ER_CLONE_DONOR_VERSION : {3864,[]string{"HY000"},"Clone Donor MySQL version: %.64s is different from Recipient MySQL version %.64s."},
	ER_CLONE_OS : {3865,[]string{"HY000"},"Clone Donor OS: %.64s is different from Recipient OS: %.64s."},
	ER_CLONE_PLATFORM : {3866,[]string{"HY000"},"Clone Donor platform: %.64s is different from Recipient platform: %.64s."},
	ER_CLONE_CHARSET : {3867,[]string{"HY000"},"Clone Donor collation: %.128s is unavailable in Recipient."},
	ER_CLONE_CONFIG : {3868,[]string{"HY000"},"Clone Configuration %.128s: Donor value: %.128s is different from Recipient value: %.128s."},
	ER_CLONE_SYS_CONFIG : {3869,[]string{"HY000"},"Clone system configuration: %.512s"},
	ER_CLONE_PLUGIN_MATCH : {3870,[]string{"HY000"},"Clone Donor plugin %.128s is not active in Recipient."},
	ER_CLONE_LOOPBACK : {3871,[]string{"HY000"},"Clone cannot use loop back connection while cloning into current data directory."},
	ER_CLONE_ENCRYPTION : {3872,[]string{"HY000"},"Clone needs SSL connection for encrypted table."},
	ER_CLONE_DISK_SPACE : {3873,[]string{"HY000"},"Clone estimated database size is %.64s. Available space %.64s is not enough."},
	ER_CLONE_IN_PROGRESS : {3874,[]string{"HY000"},"Concurrent clone in progress. Please try after clone is complete."},
	ER_CLONE_DISALLOWED : {3875,[]string{"HY000"},"The clone operation cannot be executed when %s."},
	ER_CANNOT_GRANT_ROLES_TO_ANONYMOUS_USER : {3876,[]string{"HY000"},"Cannot grant roles to an anonymous user."},
	ER_SECONDARY_ENGINE_PLUGIN : {3877,[]string{"HY000"},"%s"},
	ER_SECOND_PASSWORD_CANNOT_BE_EMPTY : {3878,[]string{"HY000"},"Empty password can not be retained as second password for user '%s'@'%s'."},
	ER_DB_ACCESS_DENIED : {3879,[]string{"HY000"},"Access denied for AuthId `%.64s`@`%.64s` to database '%-.192s'."},
	ER_DA_AUTH_ID_WITH_SYSTEM_USER_PRIV_IN_MANDATORY_ROLES : {3880,[]string{"HY000"},"Cannot set mandatory_roles: AuthId `%.64s`@`%.64s` has '%s' privilege."},
	ER_DA_RPL_GTID_TABLE_CANNOT_OPEN : {3881,[]string{"HY000"},"Gtid table is not ready to be used. Table '%s.%s' cannot be opened."},
	ER_GEOMETRY_IN_UNKNOWN_LENGTH_UNIT : {3882,[]string{"SU001"},"The geometry passed to function %s is in SRID 0, which doesn't specify a length unit. Can't convert to '%s'."},
	ER_DA_PLUGIN_INSTALL_ERROR : {3883,[]string{"HY000"},"Error installing plugin '%s': %s"},
	ER_NO_SESSION_TEMP : {3884,[]string{"HY000"},"Storage engine could not allocate temporary tablespace for this session."},
	ER_DA_UNKNOWN_ERROR_NUMBER : {3885,[]string{"HY000"},"Got unknown error: %d"},
	ER_COLUMN_CHANGE_SIZE : {3886,[]string{"HY000"},"Could not change column '%s' of table '%s'. The resulting size of index '%s' would exceed the max key length of %d bytes."},
	ER_REGEXP_INVALID_CAPTURE_GROUP_NAME : {3887,[]string{"HY000"},"A capture group has an invalid name."},
	ER_DA_SSL_LIBRARY_ERROR : {3888,[]string{"HY000"},"Failed to set up SSL because of the following SSL library error: %s"},
	ER_SECONDARY_ENGINE : {3889,[]string{"HY000"},"Secondary engine operation failed. %s."},
	ER_SECONDARY_ENGINE_DDL : {3890,[]string{"HY000"},"DDLs on a table with a secondary engine defined are not allowed."},
	ER_INCORRECT_CURRENT_PASSWORD : {3891,[]string{"HY000"},"Incorrect current password. Specify the correct password which has to be replaced."},
	ER_MISSING_CURRENT_PASSWORD : {3892,[]string{"HY000"},"Current password needs to be specified in the REPLACE clause in order to change it."},
	ER_CURRENT_PASSWORD_NOT_REQUIRED : {3893,[]string{"HY000"},"Do not specify the current password while changing it for other users."},
	ER_PASSWORD_CANNOT_BE_RETAINED_ON_PLUGIN_CHANGE : {3894,[]string{"HY000"},"Current password can not be retained for user '%s'@'%s' because authentication plugin is being changed."},
	ER_CURRENT_PASSWORD_CANNOT_BE_RETAINED : {3895,[]string{"HY000"},"Current password can not be retained for user '%s'@'%s' because new password is empty."},
	ER_PARTIAL_REVOKES_EXIST : {3896,[]string{"HY000"},"At least one partial revoke exists on a database. The system variable '@@partial_revokes' must be set to ON."},
	ER_CANNOT_GRANT_SYSTEM_PRIV_TO_MANDATORY_ROLE : {3897,[]string{"HY000"},"AuthId `%.64s`@`%.64s` is set as mandatory_roles. Cannot grant the '%s' privilege."},
	ER_XA_REPLICATION_FILTERS : {3898,[]string{"HY000"},"The use of replication filters with XA transactions is not supported, and can lead to an undefined state in the replication slave."},
	ER_UNSUPPORTED_SQL_MODE : {3899,[]string{"HY000"},"sql_mode=0x%08x is not supported."},
	ER_REGEXP_INVALID_FLAG : {3900,[]string{"HY000"},"Invalid match mode flag in regular expression."},
	ER_PARTIAL_REVOKE_AND_DB_GRANT_BOTH_EXISTS : {3901,[]string{"HY000"},"'%s' privilege for database '%s' exists both as partial revoke and mysql.db simultaneously. It could mean that the 'mysql' schema is corrupted."},
	ER_UNIT_NOT_FOUND : {3902,[]string{"SU001"},"There's no unit of measure named '%s'."},
	ER_INVALID_JSON_VALUE_FOR_FUNC_INDEX : {3903,[]string{"22018"},"Invalid JSON value for CAST for functional index '%-.64s'."},
	ER_JSON_VALUE_OUT_OF_RANGE_FOR_FUNC_INDEX : {3904,[]string{"22003"},"Out of range JSON value for CAST for functional index '%-.64s'."},
	ER_EXCEEDED_MV_KEYS_NUM : {3905,[]string{"HY000"},"Exceeded max number of values per record for multi-valued index '%-.64s' by %u value(s)."},
	ER_EXCEEDED_MV_KEYS_SPACE : {3906,[]string{"HY000"},"Exceeded max total length of values per record for multi-valued index '%-.64s' by %u bytes."},
	ER_FUNCTIONAL_INDEX_DATA_IS_TOO_LONG : {3907,[]string{"22001"},"Data too long for functional index '%-.64s'."},
	ER_WRONG_MVI_VALUE : {3908,[]string{"HY000"},"Cannot store an array or an object in a scalar key part of the index '%.192s'."},
	ER_WARN_FUNC_INDEX_NOT_APPLICABLE : {3909,[]string{"HY000"},"Cannot use functional index '%-.64s' due to type or collation conversion."},
	ER_GRP_RPL_UDF_ERROR : {3910,[]string{"HY000"},"The function '%s' failed. %s"},
	ER_UPDATE_GTID_PURGED_WITH_GR : {3911,[]string{"HY000"},"Cannot update GTID_PURGED with the Group Replication plugin running"},
	ER_GROUPING_ON_TIMESTAMP_IN_DST : {3912,[]string{"HY000"},"Grouping on temporal is non-deterministic for timezones having DST. Please consider switching to UTC for this query."},
	ER_TABLE_NAME_CAUSES_TOO_LONG_PATH : {3913,[]string{"HY000"},"Long database name and identifier for object resulted in a path length too long for table '%s'. Please check the path limit for your OS."},
	ER_AUDIT_LOG_INSUFFICIENT_PRIVILEGE : {3914,[]string{"HY000"},"Request ignored for '%.64s'@'%.64s'. Role needed to perform operation: '%.32s'"},
	//OBSOLETE_ER_AUDIT_LOG_PASSWORD_HAS_BEEN_COPIED : {13438,[]string{"HY000"},"'audit_log' password has been copied into '%.32s' and will be removed with first purged password."},
	ER_DA_GRP_RPL_STARTED_AUTO_REJOIN : {3916,[]string{"HY000"},"Started auto-rejoin procedure attempt %lu of %lu"},
	ER_SYSVAR_CHANGE_DURING_QUERY : {3917,[]string{"HY000"},"A plugin was loaded or unloaded during a query, a system variable table was changed."},
	ER_GLOBSTAT_CHANGE_DURING_QUERY : {3918,[]string{"HY000"},"A plugin was loaded or unloaded during a query, a global status variable was changed."},
	ER_GRP_RPL_MESSAGE_SERVICE_INIT_FAILURE : {3919,[]string{"HY000"},"The START GROUP_REPLICATION command failed to start its message service."},
	ER_CHANGE_MASTER_WRONG_COMPRESSION_ALGORITHM_CLIENT : {3920,[]string{"HY000"},"Invalid MASTER_COMPRESSION_ALGORITHMS '%.192s' for channel '%.192s'."},
	ER_CHANGE_MASTER_WRONG_COMPRESSION_LEVEL_CLIENT : {3921,[]string{"HY000"},"Invalid MASTER_ZSTD_COMPRESSION_LEVEL %u for channel '%.192s'."},
	ER_WRONG_COMPRESSION_ALGORITHM_CLIENT : {3922,[]string{"HY000"},"Invalid compression algorithm '%.192s'."},
	ER_WRONG_COMPRESSION_LEVEL_CLIENT : {3923,[]string{"HY000"},"Invalid zstd compression level for algorithm '%.192s'."},
	ER_CHANGE_MASTER_WRONG_COMPRESSION_ALGORITHM_LIST_CLIENT : {3924,[]string{"HY000"},"Specified compression algorithm list '%.192s' exceeds total count of 3 for channel '%.192s'."},
	ER_CLIENT_PRIVILEGE_CHECKS_USER_CANNOT_BE_ANONYMOUS : {3925,[]string{"HY000"},"PRIVILEGE_CHECKS_USER for replication channel '%.192s' was set to ``@`%.255s`, but anonymous users are disallowed for PRIVILEGE_CHECKS_USER."},
	ER_CLIENT_PRIVILEGE_CHECKS_USER_DOES_NOT_EXIST : {3926,[]string{"HY000"},"PRIVILEGE_CHECKS_USER for replication channel '%.192s' was set to `%.64s`@`%.255s`, but this is not an existing user."},
	ER_CLIENT_PRIVILEGE_CHECKS_USER_CORRUPT : {3927,[]string{"HY000"},"Invalid, corrupted PRIVILEGE_CHECKS_USER was found in the replication configuration repository for channel '%.192s'. Use CHANGE MASTER TO PRIVILEGE_CHECKS_USER to correct the configuration."},
	ER_CLIENT_PRIVILEGE_CHECKS_USER_NEEDS_RPL_APPLIER_PRIV : {3928,[]string{"HY000"},"PRIVILEGE_CHECKS_USER for replication channel '%.192s' was set to `%.64s`@`%.255s`, but this user does not have REPLICATION_APPLIER privilege."},
	ER_WARN_DA_PRIVILEGE_NOT_REGISTERED : {3929,[]string{"HY000"},"Dynamic privilege '%s' is not registered with the server."},
	ER_CLIENT_KEYRING_UDF_KEY_INVALID : {3930,[]string{"HY000"},"Function '%s' failed because key is invalid."},
	ER_CLIENT_KEYRING_UDF_KEY_TYPE_INVALID : {3931,[]string{"HY000"},"Function '%s' failed because key type is invalid."},
	ER_CLIENT_KEYRING_UDF_KEY_TOO_LONG : {3932,[]string{"HY000"},"Function '%s' failed because key length is too long."},
	ER_CLIENT_KEYRING_UDF_KEY_TYPE_TOO_LONG : {3933,[]string{"HY000"},"Function '%s' failed because key type is too long."},
	ER_JSON_SCHEMA_VALIDATION_ERROR_WITH_DETAILED_REPORT : {3934,[]string{"HY000"},"%s."},
	ER_DA_UDF_INVALID_CHARSET_SPECIFIED : {3935,[]string{"HY000"},"Invalid character set '%s' was specified. It must be either character set name or collation name as supported by server."},
	ER_DA_UDF_INVALID_CHARSET : {3936,[]string{"HY000"},"Invalid character set '%s' was specified. It must be a character set name as supported by server."},
	ER_DA_UDF_INVALID_COLLATION : {3937,[]string{"HY000"},"Invalid collation '%s' was specified. It must be a collation name as supported by server."},
	ER_DA_UDF_INVALID_EXTENSION_ARGUMENT_TYPE : {3938,[]string{"HY000"},"Invalid extension argument type '%s' was specified. Refer the MySQL manual for the valid UDF extension arguments type."},
	ER_MULTIPLE_CONSTRAINTS_WITH_SAME_NAME : {3939,[]string{"HY000"},"Table has multiple constraints with the name '%-.192s'. Please use constraint specific '%s' clause."},
	ER_CONSTRAINT_NOT_FOUND : {3940,[]string{"HY000"},"Constraint '%-.192s' does not exist."},
	ER_ALTER_CONSTRAINT_ENFORCEMENT_NOT_SUPPORTED : {3941,[]string{"HY000"},"Altering constraint enforcement is not supported for the constraint '%-.192s'. Enforcement state alter is not supported for the PRIMARY, UNIQUE and FOREIGN KEY type constraints."},
	ER_TABLE_VALUE_CONSTRUCTOR_MUST_HAVE_COLUMNS : {3942,[]string{"HY000"},"Each row of a VALUES clause must have at least one column, unless when used as source in an INSERT statement."},
	ER_TABLE_VALUE_CONSTRUCTOR_CANNOT_HAVE_DEFAULT : {3943,[]string{"HY000"},"A VALUES clause cannot use DEFAULT values, unless used as a source in an INSERT statement."},
	ER_CLIENT_QUERY_FAILURE_INVALID_NON_ROW_FORMAT : {3944,[]string{"HY000"},"The query does not comply with variable require_row_format restrictions."},
	ER_REQUIRE_ROW_FORMAT_INVALID_VALUE : {3945,[]string{"HY000"},"The requested value %s is invalid for REQUIRE_ROW_FORMAT, must be either 0 or 1."},
	ER_FAILED_TO_DETERMINE_IF_ROLE_IS_MANDATORY : {3946,[]string{"HY000"},"Failed to acquire lock on user management service, unable to determine if role `%s`@`%s` is mandatory"},
	ER_FAILED_TO_FETCH_MANDATORY_ROLE_LIST : {3947,[]string{"HY000"},"Failed to acquire lock on user management service, unable to fetch mandatory role list"},
	ER_CLIENT_LOCAL_FILES_DISABLED : {3948,[]string{"42000"},"Loading local data is disabled; this must be enabled on both the client and server sides"},
	ER_IMP_INCOMPATIBLE_CFG_VERSION : {3949,[]string{"HY000"},"Failed to import %s because the CFG file version (%u) is not compatible with the current version (%u)"},
	ER_DA_OOM : {3950,[]string{"HY000"},"Out of memory"},
	ER_DA_UDF_INVALID_ARGUMENT_TO_SET_CHARSET : {3951,[]string{"HY000"},"Character set can be set only for the UDF argument type STRING."},
	ER_DA_UDF_INVALID_RETURN_TYPE_TO_SET_CHARSET : {3952,[]string{"HY000"},"Character set can be set only for the UDF RETURN type STRING."},
	ER_MULTIPLE_INTO_CLAUSES : {3953,[]string{"HY000"},"Multiple INTO clauses in one query block."},
	ER_MISPLACED_INTO : {3954,[]string{"HY000"},"Misplaced INTO clause, INTO is not allowed inside subqueries, and must be placed at end of UNION clauses."},
	ER_USER_ACCESS_DENIED_FOR_USER_ACCOUNT_BLOCKED_BY_PASSWORD_LOCK : {3955,[]string{"HY000"},"Access denied for user '%-.48s'@'%-.64s'. Account is blocked for %s day(s) (%s day(s) remaining) due to %u consecutive failed logins."},
	ER_WARN_DEPRECATED_YEAR_UNSIGNED : {3956,[]string{"HY000"},"UNSIGNED for the YEAR data type is deprecated and support for it will be removed in a future release."},
	ER_CLONE_NETWORK_PACKET : {3957,[]string{"HY000"},"Clone needs max_allowed_packet value to be %u or more. Current value is %u"},
	ER_SDI_OPERATION_FAILED_MISSING_RECORD : {3958,[]string{"HY000"},"Failed to %s sdi for %s.%s in %s due to missing record."},
	ER_DEPENDENT_BY_CHECK_CONSTRAINT : {3959,[]string{"HY000"},"Check constraint '%-.192s' uses column '%-.192s', hence column cannot be dropped or renamed."},
	ER_GRP_OPERATION_NOT_ALLOWED_GR_MUST_STOP : {3960,[]string{"HY000"},"This operation cannot be performed while Group Replication is running; run STOP GROUP_REPLICATION first"},
	ER_WARN_DEPRECATED_JSON_TABLE_ON_ERROR_ON_EMPTY : {3961,[]string{"HY000"},"Specifying an ON EMPTY clause after the ON ERROR clause in a JSON_TABLE column definition is deprecated syntax and will be removed in a future release. Specify ON EMPTY before ON ERROR instead."},
	ER_WARN_DEPRECATED_INNER_INTO : {3962,[]string{"HY000"},"The INTO clause is deprecated inside query blocks of query expressions and will be removed in a future release. Please move the INTO clause to the end of statement instead."},
	ER_WARN_DEPRECATED_VALUES_FUNCTION_ALWAYS_NULL : {3963,[]string{"HY000"},"The VALUES function is deprecated and will be removed in a future release. It always returns NULL in this context. If you meant to access a value from the VALUES clause of the INSERT statement, consider using an alias (INSERT INTO ... VALUES (...) AS alias) and reference alias.col instead of VALUES(col) in the ON DUPLICATE KEY UPDATE clause."},
	ER_WARN_DEPRECATED_SQL_CALC_FOUND_ROWS : {3964,[]string{"HY000"},"SQL_CALC_FOUND_ROWS is deprecated and will be removed in a future release. Consider using two separate queries instead."},
	ER_WARN_DEPRECATED_FOUND_ROWS : {3965,[]string{"HY000"},"FOUND_ROWS() is deprecated and will be removed in a future release. Consider using COUNT(*) instead."},
	ER_MISSING_JSON_VALUE : {3966,[]string{"22035"},"No value was found by '%.192s' on the specified path."},
	ER_MULTIPLE_JSON_VALUES : {3967,[]string{"22034"},"More than one value was found by '%.192s' on the specified path."},
	ER_HOSTNAME_TOO_LONG : {3968,[]string{"HY000"},"Hostname cannot be longer than %d characters."},
	ER_WARN_CLIENT_DEPRECATED_PARTITION_PREFIX_KEY : {3969,[]string{"HY000"},"Column '%.64s.%.64s.%.64s' having prefix key part '%.64s(%u)' is ignored by the partitioning function. Use of prefixed columns in the PARTITION BY KEY() clause is deprecated and will be removed in a future release."},
	ER_GROUP_REPLICATION_USER_EMPTY_MSG : {3970,[]string{"HY000"},"The START GROUP_REPLICATION command failed since the username provided for recovery channel is empty."},
	ER_GROUP_REPLICATION_USER_MANDATORY_MSG : {3971,[]string{"HY000"},"The START GROUP_REPLICATION command failed since the USER option was not provided with PASSWORD for recovery channel."},
	ER_GROUP_REPLICATION_PASSWORD_LENGTH : {3972,[]string{"HY000"},"The START GROUP_REPLICATION command failed since the password provided for the recovery channel exceeds the maximum length of 32 characters."},
	ER_SUBQUERY_TRANSFORM_REJECTED : {3973,[]string{"HY000"},"Statement requires a transform of a subquery to a non-SET operation (like IN2EXISTS, or subquery-to-LATERAL-derived-table). This is not allowed with optimizer switch 'subquery_to_derived' on."},
	ER_DA_GRP_RPL_RECOVERY_ENDPOINT_FORMAT : {3974,[]string{"HY000"},"Invalid input value for recovery socket endpoints '%s'. Please, provide a valid, comma separated, list of endpoints (IP:port)."},
	ER_DA_GRP_RPL_RECOVERY_ENDPOINT_INVALID : {3975,[]string{"HY000"},"The server is not listening on endpoint '%s'. Only endpoints that the server is listening on are valid recovery endpoints."},
	ER_WRONG_VALUE_FOR_VAR_PLUS_ACTIONABLE_PART : {3976,[]string{"HY000"},"Variable '%-.64s' cannot be set to the value of '%-.200s'. %-.200s"},
	ER_STATEMENT_NOT_ALLOWED_AFTER_START_TRANSACTION : {3977,[]string{"HY000"},"Only BINLOG INSERT, COMMIT and ROLLBACK statements are allowed after CREATE TABLE with START TRANSACTION statement."},
	ER_FOREIGN_KEY_WITH_ATOMIC_CREATE_SELECT : {3978,[]string{"HY000"},"Foreign key creation is not allowed with CREATE TABLE as SELECT and CREATE TABLE with START TRANSACTION statement."},
	ER_NOT_ALLOWED_WITH_START_TRANSACTION : {3979,[]string{"HY000"},"START TRANSACTION clause cannot be used %.128s."},
	ER_INVALID_JSON_ATTRIBUTE : {3980,[]string{"HY000"},"Invalid json attribute, error: \"%s\" at pos %u: '%s'"},
	ER_ENGINE_ATTRIBUTE_NOT_SUPPORTED : {3981,[]string{"HY000"},"Storage engine '%s' does not support ENGINE_ATTRIBUTE."},
	ER_INVALID_USER_ATTRIBUTE_JSON : {3982,[]string{"HY000"},"The user attribute must be a valid JSON object"},
	ER_INNODB_REDO_DISABLED : {3983,[]string{"HY000"},"Cannot perform operation as InnoDB redo logging is disabled. Please retry after enabling redo log with ALTER INSTANCE"},
	ER_INNODB_REDO_ARCHIVING_ENABLED : {3984,[]string{"HY000"},"Cannot perform operation as InnoDB is archiving redo log. Please retry after stopping redo archive by invoking innodb_redo_log_archive_stop()"},
	ER_MDL_OUT_OF_RESOURCES : {3985,[]string{"HY000"},"Not enough resources to complete lock request."},
	ER_IMPLICIT_COMPARISON_FOR_JSON : {3986,[]string{"HY000"},"Evaluating a JSON value in SQL boolean context does an implicit comparison against JSON integer 0; if this is not what you want, consider converting JSON to a SQL numeric type with JSON_VALUE RETURNING"},
	ER_FUNCTION_DOES_NOT_SUPPORT_CHARACTER_SET : {3987,[]string{"HY000"},"The function %s does not support the character set '%s'."},
	ER_IMPOSSIBLE_STRING_CONVERSION : {3988,[]string{"HY000"},"Conversion from collation %s into %s impossible for %s"},
	ER_SCHEMA_READ_ONLY : {3989,[]string{"HY000"},"Schema '%s' is in read only mode."},
	ER_RPL_ASYNC_RECONNECT_GTID_MODE_OFF : {3990,[]string{"HY000"},"Failed to enable Asynchronous Replication Connection Failover feature. The CHANGE MASTER TO SOURCE_CONNECTION_AUTO_FAILOVER = 1 can only be set when @@GLOBAL.GTID_MODE = ON."},
	ER_RPL_ASYNC_RECONNECT_AUTO_POSITION_OFF : {3991,[]string{"HY000"},"Failed to enable Asynchronous Replication Connection Failover feature. The CHANGE MASTER TO SOURCE_CONNECTION_AUTO_FAILOVER = 1 can only be set when MASTER_AUTO_POSITION option of CHANGE MASTER TO is enabled."},
	ER_DISABLE_GTID_MODE_REQUIRES_ASYNC_RECONNECT_OFF : {3992,[]string{"HY000"},"The @@GLOBAL.GTID_MODE = %-.64s cannot be executed because Asynchronous Replication Connection Failover is enabled i.e. CHANGE MASTER TO SOURCE_CONNECTION_AUTO_FAILOVER = 1."},
	ER_DISABLE_AUTO_POSITION_REQUIRES_ASYNC_RECONNECT_OFF : {3993,[]string{"HY000"},"CHANGE MASTER TO MASTER_AUTO_POSITION = 0 cannot be executed because Asynchronous Replication Connection Failover is enabled i.e. CHANGE MASTER TO SOURCE_CONNECTION_AUTO_FAILOVER = 1."},
	ER_INVALID_PARAMETER_USE : {3994,[]string{"HY000"},"Invalid use of parameters in '%s'"},
	ER_CHARACTER_SET_MISMATCH : {3995,[]string{"HY000"},"Character set '%s' cannot be used in conjunction with '%s' in call to %s."},
	ER_WARN_VAR_VALUE_CHANGE_NOT_SUPPORTED : {3996,[]string{"HY000"},"Changing %s not supported on this platform. Falling back to the default."},
	ER_INVALID_TIME_ZONE_INTERVAL : {3997,[]string{"HY000"},"Invalid time zone interval: '%s'"},
	ER_INVALID_CAST : {3998,[]string{"HY000"},"Cannot cast value to %s."},
	ER_HYPERGRAPH_NOT_SUPPORTED_YET : {3999,[]string{"42000"},"The hypergraph optimizer does not yet support '%s'"},
	ER_WARN_HYPERGRAPH_EXPERIMENTAL : {4000,[]string{"HY000"},"The hypergraph optimizer is highly experimental and is meant for testing only. Do not enable it unless you are a MySQL developer."},
	ER_DA_NO_ERROR_LOG_PARSER_CONFIGURED : {4001,[]string{"HY000"},"None of the log-sinks selected with --log-error-services=... provides a log-parser. The server will not be able to make the previous runs' error-logs available in performance_schema.error_log."},
	ER_DA_ERROR_LOG_TABLE_DISABLED : {4002,[]string{"HY000"},"None of the log-sinks selected in @@global.log_error_services supports writing to the performance schema. The server will not be able to make the current runs' error events available in performance_schema.error_log. To change this, add a log-sink that supports the performance schema to @@global.log_error_services."},
	ER_DA_ERROR_LOG_MULTIPLE_FILTERS : {4003,[]string{"HY000"},"@@global.log_error_services lists more than one log-filter service. This is discouraged as it will make it hard to understand which rule in which filter affected a log-event."},
	ER_DA_CANT_OPEN_ERROR_LOG : {4004,[]string{"HY000"},"Could not open file '%s' for error logging%s%s"},
	ER_USER_REFERENCED_AS_DEFINER : {4005,[]string{"HY000"},"User %.256s is referenced as a definer account in %s."},
	ER_CANNOT_USER_REFERENCED_AS_DEFINER : {4006,[]string{"HY000"},"Operation %s failed for %.256s as it is referenced as a definer account in %s."},
	ER_REGEX_NUMBER_TOO_BIG : {4007,[]string{"HY000"},"Decimal number in regular expression is too large."},
	ER_SPVAR_NONINTEGER_TYPE : {4008,[]string{"HY000"},"The variable \"%s\" has a non-integer based type"},
	WARN_UNSUPPORTED_ACL_TABLES_READ : {4009,[]string{"HY000"},"Reads with serializable isolation/SELECT FOR SHARE are not supported for ACL tables."},
	ER_BINLOG_UNSAFE_ACL_TABLE_READ_IN_DML_DDL : {4010,[]string{"HY000"},"The statement is unsafe because it updates a table depending on ACL table read operation. As storage engine row locks are skipped for ACL table, it may not have same effect on master and slave."},
	ER_STOP_REPLICA_MONITOR_IO_THREAD_TIMEOUT : {4011,[]string{"HY000"},"STOP REPLICA command execution is incomplete: Replica Monitor thread got the stop signal, thread is busy, Monitor thread will stop once the current task is complete."},
	ER_STARTING_REPLICA_MONITOR_IO_THREAD : {4012,[]string{"HY000"},"The Replica Monitor thread failed to start."},
	ER_CANT_USE_ANONYMOUS_TO_GTID_WITH_GTID_MODE_NOT_ON : {4013,[]string{"HY000"},"Replication cannot start%.192s with ASSIGN_GTIDS_TO_ANONYMOUS_TRANSACTIONS = LOCAL|<UUID> as this server uses @@GLOBAL.GTID_MODE <> ON."},
	ER_CANT_COMBINE_ANONYMOUS_TO_GTID_AND_AUTOPOSITION : {4014,[]string{"HY000"},"The options ASSIGN_GTIDS_TO_ANONYMOUS_TRANSACTIONS = LOCAL|<UUID> and MASTER_AUTO_POSITION = 1 cannot be used together."},
	ER_ASSIGN_GTIDS_TO_ANONYMOUS_TRANSACTIONS_REQUIRES_GTID_MODE_ON : {4015,[]string{"HY000"},"CHANGE MASTER TO ASSIGN_GTIDS_TO_ANONYMOUS_TRANSACTIONS = LOCAL|<UUID> cannot be executed because @@GLOBAL.GTID_MODE <> ON."},
	ER_SQL_SLAVE_SKIP_COUNTER_USED_WITH_GTID_MODE_ON : {4016,[]string{"HY000"},"The value of sql_slave_skip_counter will only take effect for channels running with ASSIGN_GTIDS_TO_ANONYMOUS_TRANSACTIONS <> OFF."},
	ER_USING_ASSIGN_GTIDS_TO_ANONYMOUS_TRANSACTIONS_AS_LOCAL_OR_UUID : {4017,[]string{"HY000"},"Using ASSIGN_GTIDS_TO_ANONYMOUS_TRANSACTIONS creates limitations on the replication topology - you cannot fail over between downstream and upstream servers. Only use this option if it is not possible to enable GTIDs on the source, for instance, because of lack of permissions. If possible, use the procedure for enabling GTID transactions online instead, as described in the documentation."},
	ER_CANT_SET_ANONYMOUS_TO_GTID_AND_WAIT_UNTIL_SQL_THD_AFTER_GTIDS : {4018,[]string{"HY000"},"WAIT_UNTIL_SQL_THREAD_AFTER_GTIDS cannot be used on a channel configured with ASSIGN_GTIDS_TO_ANONYMOUS_TRANSACTIONS = LOCAL|<UUID>"},
	ER_CANT_SET_SQL_AFTER_OR_BEFORE_GTIDS_WITH_ANONYMOUS_TO_GTID : {4019,[]string{"HY000"},"The SQL_AFTER_GTIDS or SQL_BEFORE_GTIDS clauses for START REPLICA cannot be used when the replication channel is configured with ASSIGN_GTIDS_TO_ANONYMOUS_TRANSACTIONS = LOCAL|<UUID>."},
	ER_ANONYMOUS_TO_GTID_UUID_SAME_AS_GROUP_NAME : {4020,[]string{"HY000"},"Replication '%.192s' is configured with ASSIGN_GTIDS_TO_ANONYMOUS_TRANSACTIONS = <UUID> where the UUID value is equal to the group_replication_group_name"},
	ER_CANT_USE_SAME_UUID_AS_GROUP_NAME : {4021,[]string{"HY000"},"CHANGE MASTER TO ASSIGN_GTIDS_TO_ANONYMOUS_TRANSACTIONS = <UUID> cannot be executed because the UUID value is equal to the group_replication_group_name."},
	ER_GRP_RPL_RECOVERY_CHANNEL_STILL_RUNNING : {4022,[]string{"HY000"},"The group_replication_recovery channel is still running, most likely it is waiting for a database/table lock, which is preventing the channel from stopping. Please check database/table locks, including the ones created by backup tools."},
	ER_INNODB_INVALID_AUTOEXTEND_SIZE_VALUE : {4023,[]string{"HY000"},"AUTOEXTEND_SIZE should be a multiple of %uM"},
	ER_INNODB_INCOMPATIBLE_WITH_TABLESPACE : {4024,[]string{"HY000"},"InnoDB: \"%s\" not allowed with general tablespaces"},
	ER_INNODB_AUTOEXTEND_SIZE_OUT_OF_RANGE : {4025,[]string{"HY000"},"AUTOEXTEND_SIZE value should be between %uM and %uM"},
	ER_CANNOT_USE_AUTOEXTEND_SIZE_CLAUSE : {4026,[]string{"HY000"},"AUTOEXTEND_SIZE clause is not valid for %s tablespace."},
	ER_ROLE_GRANTED_TO_ITSELF : {4027,[]string{"HY000"},"User account %s is directly or indirectly granted to the role %s. The GRANT would create a loop"},
	ER_TABLE_MUST_HAVE_A_VISIBLE_COLUMN : {4028,[]string{"HY000"},"A table must have at least one visible column."},
	ER_INNODB_COMPRESSION_FAILURE : {4029,[]string{"HY000"},"Compression failed with the following error : %s"},
	ER_WARN_ASYNC_CONN_FAILOVER_NETWORK_NAMESPACE : {4030,[]string{"HY000"},"The parameter network_namespace is reserved for future use. Please use the CHANGE REPLICATION SOURCE command to set channel network_namespace parameter."},
	ER_PARSER_TRACE : {10000,[]string{"XX999"},"Parser saw: %s"},
	ER_BOOTSTRAP_CANT_THREAD : {10001,[]string{"HY000"},"Can't create thread to handle bootstrap (errno: %d)"},
	ER_TRIGGER_INVALID_VALUE : {10002,[]string{"HY000"},"Trigger for table '%s'.'%s': invalid %s value (%s)."},
	ER_OPT_WRONG_TREE : {10003,[]string{"HY000"},"Wrong tree: %s"},
	ER_DD_FAILSAFE : {10004,[]string{"HY000"},"Error: Invalid %s"},
	ER_DD_NO_WRITES_NO_REPOPULATION : {10005,[]string{"HY000"},"Skip re-populating collations and character sets tables in %s%sread-only mode."},
	ER_DD_VERSION_FOUND : {10006,[]string{"HY000"},"Using data dictionary with version '%d'."},
	ER_DD_VERSION_INSTALLED : {10007,[]string{"HY000"},"Installed data dictionary with version %d"},
	ER_DD_VERSION_UNSUPPORTED : {10008,[]string{"HY000"},"Data Dictionary version '%d' not supported."},
	//OBSOLETE_ER_LOG_SYSLOG_FACILITY_FAIL : {10009,[]string{"HY000"},"Failed to set syslog facility to \"%s\", setting to \"%s\" (%d) instead."},
	ER_LOG_SYSLOG_CANNOT_OPEN : {10010,[]string{"HY000"},"Cannot open %s; check privileges, or remove syseventlog from --log-error-services!"},
	ER_LOG_SLOW_CANNOT_OPEN : {10011,[]string{"HY000"}," either restart the query logging by using \"SET GLOBAL SLOW_QUERY_LOG=ON\" or"},
	ER_LOG_GENERAL_CANNOT_OPEN : {10012,[]string{"HY000"}," either restart the query logging by using \"SET GLOBAL GENERAL_LOG=ON\" or"},
	ER_LOG_CANNOT_WRITE : {10013,[]string{"HY000"},"Failed to write to %s: %s"},
	ER_RPL_ZOMBIE_ENCOUNTERED : {10014,[]string{"HY000"},"While initializing dump thread for slave with %s <%s>, found a zombie dump thread with the same %s. Master is killing the zombie dump thread(%u)."},
	ER_RPL_GTID_TABLE_CANNOT_OPEN : {10015,[]string{"HY000"},"Gtid table is not ready to be used. Table '%s.%s' cannot be opened."},
	ER_SYSTEM_SCHEMA_NOT_FOUND : {10016,[]string{"HY000"},"System schema directory does not exist."},
	ER_DD_INIT_UPGRADE_FAILED : {10017,[]string{"HY000"},"Error in initializing dictionary, upgrade will do a cleanup and exit"},
	ER_VIEW_UNKNOWN_CHARSET_OR_COLLATION : {10018,[]string{"HY000"},"View '%s'.'%s': unknown charset name and/or collation name (client: '%s'; connection: '%s')."},
	ER_DD_VIEW_CANT_ALLOC_CHARSET : {10019,[]string{"HY000"},"Error in allocating memory for character set name for view %s.%s."},
	ER_DD_INIT_FAILED : {10020,[]string{"HY000"},"Data Dictionary initialization failed."},
	ER_DD_UPDATING_PLUGIN_MD_FAILED : {10021,[]string{"HY000"},"Failed to update plugin metadata in dictionary tables."},
	ER_DD_POPULATING_TABLES_FAILED : {10022,[]string{"HY000"},"Failed to Populate DD tables."},
	ER_DD_VIEW_CANT_CREATE : {10023,[]string{"HY000"},"Error in Creating View %s.%s"},
	ER_DD_METADATA_NOT_FOUND : {10024,[]string{"HY000"},"Unable to start server. Cannot find the meta data for data dictionary table '%s'."},
	ER_DD_CACHE_NOT_EMPTY_AT_SHUTDOWN : {10025,[]string{"HY000"},"Dictionary cache not empty at shutdown."},
	ER_DD_OBJECT_REMAINS : {10026,[]string{"HY000"},"Dictionary objects used but not released."},
	ER_DD_OBJECT_REMAINS_IN_RELEASER : {10027,[]string{"HY000"},"Dictionary objects left in default releaser."},
	ER_DD_OBJECT_RELEASER_REMAINS : {10028,[]string{"HY000"},"Dictionary object auto releaser not deleted"},
	ER_DD_CANT_GET_OBJECT_KEY : {10029,[]string{"HY000"},"Error: Unable to create primary object key"},
	ER_DD_CANT_CREATE_OBJECT_KEY : {10030,[]string{"HY000"},"Error: Unable to create object key"},
	ER_CANT_CREATE_HANDLE_MGR_THREAD : {10031,[]string{"HY000"},"Can't create handle_manager thread (errno= %d)"},
	ER_RPL_REPO_HAS_GAPS : {10032,[]string{"HY000"},"It is not possible to change the type of the relay log's repository because there are workers' repositories with gaps. Please, fix the gaps first before doing such change."},
	ER_INVALID_VALUE_FOR_ENFORCE_GTID_CONSISTENCY : {10033,[]string{"HY000"},"option 'enforce-gtid-consistency': value '%s' was not recognized. Setting enforce-gtid-consistency to OFF."},
	ER_CHANGED_ENFORCE_GTID_CONSISTENCY : {10034,[]string{"HY000"},"Changed ENFORCE_GTID_CONSISTENCY from %s to %s."},
	ER_CHANGED_GTID_MODE : {10035,[]string{"HY000"},"Changed GTID_MODE from %s to %s."},
	ER_DISABLED_STORAGE_ENGINE_AS_DEFAULT : {10036,[]string{"HY000"},"%s is set to a disabled storage engine %s."},
	ER_DEBUG_SYNC_HIT : {10037,[]string{"HY000"},"Debug sync points hit:                   %22s"},
	ER_DEBUG_SYNC_EXECUTED : {10038,[]string{"HY000"},"Debug sync points executed:              %22s"},
	ER_DEBUG_SYNC_THREAD_MAX : {10039,[]string{"HY000"},"Debug sync points max active per thread: %22s"},
	ER_DEBUG_SYNC_OOM : {10040,[]string{"HY000"},"Debug Sync Facility disabled due to lack of memory."},
	ER_CANT_INIT_TC_LOG : {10041,[]string{"HY000"},"Can't init tc log"},
	ER_EVENT_CANT_INIT_QUEUE : {10042,[]string{"HY000"},"Event Scheduler: Can't initialize the execution queue"},
	ER_EVENT_PURGING_QUEUE : {10043,[]string{"HY000"},"Event Scheduler: Purging the queue. %u events"},
	ER_EVENT_LAST_EXECUTION : {10044,[]string{"HY000"},"Event Scheduler: Last execution of %s.%s. %s"},
	ER_EVENT_MESSAGE_STACK : {10045,[]string{"HY000"},"%*s"},
	ER_EVENT_EXECUTION_FAILED : {10046,[]string{"HY000"},"Event Scheduler: [%s].[%s.%s] event execution failed."},
	ER_CANT_INIT_SCHEDULER_THREAD : {10047,[]string{"HY000"},"Event Scheduler: Cannot initialize the scheduler thread"},
	ER_SCHEDULER_STOPPED : {10048,[]string{"HY000"},"Event Scheduler: Stopped"},
	ER_CANT_CREATE_SCHEDULER_THREAD : {10049,[]string{"HY000"},"Event scheduler: Failed to start scheduler, Can not create thread for event scheduler (errno=%d)"},
	ER_SCHEDULER_WAITING : {10050,[]string{"HY000"},"Event Scheduler: Waiting for the scheduler thread to reply"},
	ER_SCHEDULER_STARTED : {10051,[]string{"HY000"},"Event Scheduler: scheduler thread started with id %u"},
	ER_SCHEDULER_STOPPING_FAILED_TO_GET_EVENT : {10052,[]string{"HY000"},"Event Scheduler: Serious error during getting next event to execute. Stopping"},
	ER_SCHEDULER_STOPPING_FAILED_TO_CREATE_WORKER : {10053,[]string{"HY000"},"Event_scheduler::execute_top: Can not create event worker thread (errno=%d). Stopping event scheduler"},
	ER_SCHEDULER_KILLING : {10054,[]string{"HY000"},"Event Scheduler: Killing the scheduler thread, thread id %u"},
	ER_UNABLE_TO_RESOLVE_IP : {10055,[]string{"HY000"},"IP address '%s' could not be resolved: %s"},
	ER_UNABLE_TO_RESOLVE_HOSTNAME : {10056,[]string{"HY000"},"Host name '%s' could not be resolved: %s"},
	ER_HOSTNAME_RESEMBLES_IPV4 : {10057,[]string{"HY000"},"IP address '%s' has been resolved to the host name '%s', which resembles IPv4-address itself."},
	ER_HOSTNAME_DOESNT_RESOLVE_TO : {10058,[]string{"HY000"},"Hostname '%s' does not resolve to '%s'."},
	ER_ADDRESSES_FOR_HOSTNAME_HEADER : {10059,[]string{"HY000"},"Hostname '%s' has the following IP addresses:"},
	ER_ADDRESSES_FOR_HOSTNAME_LIST_ITEM : {10060,[]string{"HY000"}," - %s"},
	ER_TRG_WITHOUT_DEFINER : {10061,[]string{"HY000"},"Definer clause is missing in Trigger of Table %s. Rebuild Trigger to fix definer."},
	ER_TRG_NO_CLIENT_CHARSET : {10062,[]string{"HY000"},"Client character set is missing for trigger of table %s. Using default character set."},
	ER_PARSING_VIEW : {10063,[]string{"HY000"},"Error in parsing view %s.%s"},
	ER_COMPONENTS_INFRASTRUCTURE_BOOTSTRAP : {10064,[]string{"HY000"},"Failed to bootstrap components infrastructure."},
	ER_COMPONENTS_INFRASTRUCTURE_SHUTDOWN : {10065,[]string{"HY000"},"Failed to shutdown components infrastructure."},
	ER_COMPONENTS_PERSIST_LOADER_BOOTSTRAP : {10066,[]string{"HY000"},"Failed to bootstrap persistent components loader."},
	ER_DEPART_WITH_GRACE : {10067,[]string{"HY000"},"Giving %d client threads a chance to die gracefully"},
	ER_CA_SELF_SIGNED : {10068,[]string{"HY000"},"CA certificate %s is self signed."},
	ER_SSL_LIBRARY_ERROR : {10069,[]string{"HY000"},"Failed to set up SSL because of the following SSL library error: %s"},
	ER_NO_THD_NO_UUID : {10070,[]string{"HY000"},"Failed to generate a server UUID because it is failed to allocate the THD."},
	ER_UUID_SALT : {10071,[]string{"HY000"},"Salting uuid generator variables, current_pid: %lu, server_start_time: %lu, bytes_sent: %llu, "},
	ER_UUID_IS : {10072,[]string{"HY000"},"Generated uuid: '%s', server_start_time: %lu, bytes_sent: %llu"},
	ER_UUID_INVALID : {10073,[]string{"HY000"},"The server_uuid stored in auto.cnf file is not a valid UUID."},
	ER_UUID_SCRUB : {10074,[]string{"HY000"},"Garbage characters found at the end of the server_uuid value in auto.cnf file. It should be of length '%d' (UUID_LENGTH). Clear it and restart the server. "},
	ER_CREATING_NEW_UUID : {10075,[]string{"HY000"},"No existing UUID has been found, so we assume that this is the first time that this server has been started. Generating a new UUID: %s."},
	ER_CANT_CREATE_UUID : {10076,[]string{"HY000"},"Initialization of the server's UUID failed because it could not be read from the auto.cnf file. If this is a new server, the initialization failed because it was not possible to generate a new UUID."},
	ER_UNKNOWN_UNSUPPORTED_STORAGE_ENGINE : {10077,[]string{"HY000"},"Unknown/unsupported storage engine: %s"},
	ER_SECURE_AUTH_VALUE_UNSUPPORTED : {10078,[]string{"HY000"},"Unsupported value 0 for secure-auth"},
	ER_INVALID_INSTRUMENT : {10079,[]string{"HY000"},"Invalid instrument name or value for performance_schema_instrument '%s'",},
	ER_INNODB_MANDATORY : {10080,[]string{"HY000"},"The use of InnoDB is mandatory since MySQL 5.7. The former options like '--innodb=0/1/OFF/ON' or '--skip-innodb' are ignored."},
	//OBSOLETE_ER_INNODB_CANNOT_BE_IGNORED : {3704,[]string{"HY000"},"ignore-builtin-innodb is ignored and will be removed in future releases."},
	//OBSOLETE_ER_OLD_PASSWORDS_NO_MIDDLE_GROUND : {10082,[]string{"HY000"},"Invalid old_passwords mode: 1. Valid values are 2 and 0"},
	ER_VERBOSE_REQUIRES_HELP : {10083,[]string{"HY000"},"--verbose is for use with --help; did you mean --log-error-verbosity?"},
	ER_POINTLESS_WITHOUT_SLOWLOG : {10084,[]string{"HY000"},"options --log-slow-admin-statements, --log-queries-not-using-indexes and --log-slow-slave-statements have no effect if --slow-query-log is not set"},
	ER_WASTEFUL_NET_BUFFER_SIZE : {10085,[]string{"HY000"},"net_buffer_length (%lu) is set to be larger than max_allowed_packet (%lu). Please rectify."},
	ER_DEPRECATED_TIMESTAMP_IMPLICIT_DEFAULTS : {10086,[]string{"HY000"},"TIMESTAMP with implicit DEFAULT value is deprecated. Please use --explicit_defaults_for_timestamp server option (see documentation for more details)."},
	ER_FT_BOOL_SYNTAX_INVALID : {10087,[]string{"HY000"},"Invalid ft-boolean-syntax string: %s"},
	ER_CREDENTIALLESS_AUTO_USER_BAD : {10088,[]string{"HY000"},"'NO_AUTO_CREATE_USER' sql mode was not set."},
	ER_CONNECTION_HANDLING_OOM : {10089,[]string{"HY000"},"Could not allocate memory for connection handling"},
	ER_THREAD_HANDLING_OOM : {10090,[]string{"HY000"},"Could not allocate memory for thread handling"},
	ER_CANT_CREATE_TEST_FILE : {10091,[]string{"HY000"},"Can't create test file %s"},
	ER_CANT_CREATE_PID_FILE : {10092,[]string{"HY000"},"Can't start server: can't create PID file: %s"},
	ER_CANT_REMOVE_PID_FILE : {10093,[]string{"HY000"},"Unable to delete pid file: %s"},
	ER_CANT_CREATE_SHUTDOWN_THREAD : {10094,[]string{"HY000"},"Can't create thread to handle shutdown requests (errno= %d)"},
	ER_SEC_FILE_PRIV_CANT_ACCESS_DIR : {10095,[]string{"HY000"},"Failed to access directory for --secure-file-priv. Please make sure that directory exists and is accessible by MySQL Server. Supplied value : %s"},
	ER_SEC_FILE_PRIV_IGNORED : {10096,[]string{"HY000"},"Ignoring --secure-file-priv value as server is running with --initialize(-insecure)."},
	ER_SEC_FILE_PRIV_EMPTY : {10097,[]string{"HY000"},"Insecure configuration for --secure-file-priv: Current value does not restrict location of generated files. Consider setting it to a valid, non-empty path."},
	ER_SEC_FILE_PRIV_NULL : {10098,[]string{"HY000"},"--secure-file-priv is set to NULL. Operations related to importing and exporting data are disabled"},
	ER_SEC_FILE_PRIV_DIRECTORY_INSECURE : {10099,[]string{"HY000"},"Insecure configuration for --secure-file-priv: %s is accessible through --secure-file-priv. Consider choosing a different directory."},
	ER_SEC_FILE_PRIV_CANT_STAT : {10100,[]string{"HY000"},"Failed to get stat for directory pointed out by --secure-file-priv"},
	ER_SEC_FILE_PRIV_DIRECTORY_PERMISSIONS : {10101,[]string{"HY000"},"Insecure configuration for --secure-file-priv: Location is accessible to all OS users. Consider choosing a different directory."},
	ER_SEC_FILE_PRIV_ARGUMENT_TOO_LONG : {10102,[]string{"HY000"},"Value for --secure-file-priv is longer than maximum limit of %d"},
	ER_CANT_CREATE_NAMED_PIPES_THREAD : {10103,[]string{"HY000"},"Can't create thread to handle named pipes (errno= %d)"},
	ER_CANT_CREATE_TCPIP_THREAD : {10104,[]string{"HY000"},"Can't create thread to handle TCP/IP (errno= %d)"},
	ER_CANT_CREATE_SHM_THREAD : {10105,[]string{"HY000"},"Can't create thread to handle shared memory (errno= %d)"},
	ER_CANT_CREATE_INTERRUPT_THREAD : {10106,[]string{"HY000"},"Can't create interrupt-thread (error %d, errno: %d)"},
	ER_WRITABLE_CONFIG_REMOVED : {10107,[]string{"HY000"},"World-writable config file '%s' has been removed."},
	ER_CORE_VALUES : {10108,[]string{"HY000"},"setrlimit could not change the size of core files to 'infinity';  We may not be able to generate a core file on signals"},
	ER_WRONG_DATETIME_SPEC : {10109,[]string{"HY000"},"Wrong date/time format specifier: %s"},
	ER_RPL_BINLOG_FILTERS_OOM : {10110,[]string{"HY000"},"Could not allocate replication and binlog filters: %s"},
	ER_KEYCACHE_OOM : {10111,[]string{"HY000"},"Cannot allocate the keycache"},
	ER_CONFIRMING_THE_FUTURE : {10112,[]string{"HY000"},"Current time has got past year 2038. Validating current time with %d iterations before initiating the normal server shutdown process."},
	ER_BACK_IN_TIME : {10113,[]string{"HY000"},"Iteration %d: Obtained valid current time from system"},
	ER_FUTURE_DATE : {10114,[]string{"HY000"},"Iteration %d: Current time obtained from system is greater than 2038"},
	ER_UNSUPPORTED_DATE : {10115,[]string{"HY000"},"This MySQL server doesn't support dates later then 2038"},
	ER_STARTING_AS : {10116,[]string{"HY000"},"%s (mysqld %s) starting as process %lu"},
	ER_SHUTTING_DOWN_SLAVE_THREADS : {10117,[]string{"HY000"},"Shutting down slave threads"},
	ER_DISCONNECTING_REMAINING_CLIENTS : {10118,[]string{"HY000"},"Forcefully disconnecting %d remaining clients"},
	ER_ABORTING : {10119,[]string{"HY000"},"Aborting"},
	ER_BINLOG_END : {10120,[]string{"HY000"},"Binlog end"},
	ER_CALL_ME_LOCALHOST : {10121,[]string{"HY000"},"gethostname failed, using '%s' as hostname"},
	ER_USER_REQUIRES_ROOT : {10122,[]string{"HY000"},"One can only use the --user switch if running as root"},
	ER_REALLY_RUN_AS_ROOT : {10123,[]string{"HY000"},"Fatal error: Please read \"Security\" section of the manual to find out how to run mysqld as root!"},
	ER_USER_WHAT_USER : {10124,[]string{"HY000"},"Fatal error: Can't change to run as user '%s' ;  Please check that the user exists!"},
	ER_TRANSPORTS_WHAT_TRANSPORTS : {10125,[]string{"HY000"},"Server is started with --require-secure-transport=ON but no secure transports (SSL or Shared Memory) are configured."},
	ER_FAIL_SETGID : {10126,[]string{"HY000"},"setgid: %s"},
	ER_FAIL_SETUID : {10127,[]string{"HY000"},"setuid: %s"},
	ER_FAIL_SETREGID : {10128,[]string{"HY000"},"setregid: %s"},
	ER_FAIL_SETREUID : {10129,[]string{"HY000"},"setreuid: %s"},
	ER_FAIL_CHROOT : {10130,[]string{"HY000"},"chroot: %s"},
	ER_WIN_LISTEN_BUT_HOW : {10131,[]string{"HY000"},"TCP/IP, --shared-memory, or --named-pipe should be configured on NT OS"},
	ER_NOT_RIGHT_NOW : {10132,[]string{"HY000"},"CTRL-C ignored during startup"},
	ER_FIXING_CLIENT_CHARSET : {10133,[]string{"HY000"},"'%s' can not be used as client character set. '%s' will be used as default client character set."},
	ER_OOM : {10134,[]string{"HY000"},"Out of memory"},
	ER_FAILED_TO_LOCK_MEM : {10135,[]string{"HY000"},"Failed to lock memory. Errno: %d"},
	ER_MYINIT_FAILED : {10136,[]string{"HY000"},"my_init() failed."},
	ER_BEG_INITFILE : {10137,[]string{"HY000"},"Execution of init_file '%s' started."},
	ER_END_INITFILE : {10138,[]string{"HY000"},"Execution of init_file '%s' ended."},
	ER_CHANGED_MAX_OPEN_FILES : {10139,[]string{"HY000"},"Changed limits: max_open_files: %lu (requested %lu)"},
	ER_CANT_INCREASE_MAX_OPEN_FILES : {10140,[]string{"HY000"},"Could not increase number of max_open_files to more than %lu (request: %lu)"},
	ER_CHANGED_MAX_CONNECTIONS : {10141,[]string{"HY000"},"Changed limits: max_connections: %lu (requested %lu)"},
	ER_CHANGED_TABLE_OPEN_CACHE : {10142,[]string{"HY000"},"Changed limits: table_open_cache: %lu (requested %lu)"},
	ER_THE_USER_ABIDES : {10143,[]string{"HY000"},"Ignoring user change to '%s' because the user was set to '%s' earlier on the command line"},
	ER_RPL_CANT_ADD_DO_TABLE : {10144,[]string{"HY000"},"Could not add do table rule '%s'!"},
	ER_RPL_CANT_ADD_IGNORE_TABLE : {10145,[]string{"HY000"},"Could not add ignore table rule '%s'!"},
	ER_TRACK_VARIABLES_BOGUS : {10146,[]string{"HY000"},"The variable session_track_system_variables either has duplicate values or invalid values."},
	ER_EXCESS_ARGUMENTS : {10147,[]string{"HY000"},"Too many arguments (first extra is '%s')."},
	ER_VERBOSE_HINT : {10148,[]string{"HY000"},"Use --verbose --help to get a list of available options!"},
	ER_CANT_READ_ERRMSGS : {10149,[]string{"HY000"},"Unable to read errmsg.sys file"},
	ER_CANT_INIT_DBS : {10150,[]string{"HY000"},"Can't init databases"},
	ER_LOG_OUTPUT_CONTRADICTORY : {10151,[]string{"HY000"},"There were other values specified to log-output besides NONE. Disabling slow and general logs anyway."},
	ER_NO_CSV_NO_LOG_TABLES : {10152,[]string{"HY000"},"CSV engine is not present, falling back to the log files"},
	ER_RPL_REWRITEDB_MISSING_ARROW : {10153,[]string{"HY000"},"Bad syntax in replicate-rewrite-db - missing '->'!"},
	ER_RPL_REWRITEDB_EMPTY_FROM : {10154,[]string{"HY000"},"Bad syntax in replicate-rewrite-db - empty FROM db!"},
	ER_RPL_REWRITEDB_EMPTY_TO : {10155,[]string{"HY000"},"Bad syntax in replicate-rewrite-db - empty TO db!"},
	ER_LOG_FILES_GIVEN_LOG_OUTPUT_IS_TABLE : {10156,[]string{"HY000"},"Although a path was specified for the %s, log tables are used. To enable logging to files use the --log-output=file option."},
	ER_LOG_FILE_INVALID : {10157,[]string{"HY000"},"Invalid value for %s: %s"},
	ER_LOWER_CASE_TABLE_NAMES_CS_DD_ON_CI_FS_UNSUPPORTED : {10158,[]string{"HY000"},"The server option 'lower_case_table_names' is configured to use case sensitive table names but the data directory is on a case-insensitive file system which is an unsupported combination. Please consider either using a case sensitive file system for your data directory or switching to a case-insensitive table name mode."},
	ER_LOWER_CASE_TABLE_NAMES_USING_2 : {10159,[]string{"HY000"},"Setting lower_case_table_names=2 because file system for %s is case insensitive"},
	ER_LOWER_CASE_TABLE_NAMES_USING_0 : {10160,[]string{"HY000"},"lower_case_table_names was set to 2, even though your the file system '%s' is case sensitive.  Now setting lower_case_table_names to 0 to avoid future problems."},
	ER_NEED_LOG_BIN : {10161,[]string{"HY000"},"You need to use --log-bin to make %s work."},
	ER_NEED_FILE_INSTEAD_OF_DIR : {10162,[]string{"HY000"},"Path '%s' is a directory name, please specify a file name for %s option"},
	ER_LOG_BIN_BETTER_WITH_NAME : {10163,[]string{"HY000"},"No argument was provided to --log-bin, and --log-bin-index was not used; so replication may break when this MySQL server acts as a master and has his hostname changed!! Please use '--log-bin=%s' to avoid this problem."},
	ER_BINLOG_NEEDS_SERVERID : {10164,[]string{"HY000"},"You have enabled the binary log, but you haven't provided the mandatory server-id. Please refer to the proper server start-up parameters documentation"},
	ER_RPL_CANT_MAKE_PATHS : {10165,[]string{"HY000"},"Unable to create replication path names: out of memory or path names too long (path name exceeds %d or file name exceeds %d)."},
	ER_CANT_INITIALIZE_GTID : {10166,[]string{"HY000"},"Failed to initialize GTID structures."},
	ER_CANT_INITIALIZE_EARLY_PLUGINS : {10167,[]string{"HY000"},"Failed to initialize early plugins."},
	ER_CANT_INITIALIZE_BUILTIN_PLUGINS : {10168,[]string{"HY000"},"Failed to initialize builtin plugins."},
	ER_CANT_INITIALIZE_DYNAMIC_PLUGINS : {10169,[]string{"HY000"},"Failed to initialize dynamic plugins."},
	ER_PERFSCHEMA_INIT_FAILED : {10170,[]string{"HY000"},"Performance schema disabled (reason: init failed)."},
	ER_STACKSIZE_UNEXPECTED : {10171,[]string{"HY000"},"Asked for %lu thread stack, but got %ld"},
	//OBSOLETE_ER_CANT_SET_DATADIR : {10172,[]string{"HY000"},"failed to set datadir to %s"},
	ER_CANT_STAT_DATADIR : {10173,[]string{"HY000"},"Can't read data directory's stats (%d): %s. Assuming that it's not owned by the same user/group"},
	ER_CANT_CHOWN_DATADIR : {10174,[]string{"HY000"},"Can't change data directory owner to %s"},
	ER_CANT_SET_UP_PERSISTED_VALUES : {10175,[]string{"HY000"},"Setting persistent options failed."},
	ER_CANT_SAVE_GTIDS : {10176,[]string{"HY000"},"Failed to save the set of Global Transaction Identifiers of the last binary log into the mysql.gtid_executed table while the server was shutting down. The next server restart will make another attempt to save Global Transaction Identifiers into the table."},
	ER_AUTH_CANT_SET_DEFAULT_PLUGIN : {10177,[]string{"HY000"},"Can't start server: Invalid value for --default-authentication-plugin"},
	ER_CANT_JOIN_SHUTDOWN_THREAD : {10178,[]string{"HY000"},"Could not join %sthread. error:%d"},
	ER_CANT_HASH_DO_AND_IGNORE_RULES : {10179,[]string{"HY000"},"An error occurred while building do_table and ignore_table rules to hashes for global replication filter."},
	ER_CANT_OPEN_CA : {10180,[]string{"HY000"},"Error opening CA certificate file"},
	ER_CANT_ACCESS_CAPATH : {10181,[]string{"HY000"},"Error accessing directory pointed by --ssl-capath"},
	ER_SSL_TRYING_DATADIR_DEFAULTS : {10182,[]string{"HY000"},"Found %s, %s and %s in data directory. Trying to enable SSL support using them."},
	ER_AUTO_OPTIONS_FAILED : {10183,[]string{"HY000"},"Failed to create %s(file: '%s', errno %d)"},
	ER_CANT_INIT_TIMER : {10184,[]string{"HY000"},"Failed to initialize timer component (errno %d)."},
	ER_SERVERID_TOO_LARGE : {10185,[]string{"HY000"},"server-id configured is too large to represent with server-id-bits configured."},
	ER_DEFAULT_SE_UNAVAILABLE : {10186,[]string{"HY000"},"Default%s storage engine (%s) is not available"},
	ER_CANT_OPEN_ERROR_LOG : {10187,[]string{"HY000"},"Could not open file '%s' for error logging%s%s"},
	ER_INVALID_ERROR_LOG_NAME : {10188,[]string{"HY000"},"Invalid log file name after expanding symlinks: '%s'"},
	ER_RPL_INFINITY_DENIED : {10189,[]string{"HY000"},"using --replicate-same-server-id in conjunction with --log-slave-updates is impossible, it would lead to infinite loops in this server."},
	ER_RPL_INFINITY_IGNORED : {10190,[]string{"HY000"},"using --replicate-same-server-id in conjunction with --log-slave-updates would lead to infinite loops in this server. However this will be ignored as the --log-bin option is not defined or your server is running with global transaction identiers enabled."},
	//OBSOLETE_ER_NDB_TABLES_NOT_READY : {10191,[]string{"HY000"},"NDB : Tables not available after %lu seconds. Consider increasing --ndb-wait-setup value"},
	ER_TABLE_CHECK_INTACT : {10192,[]string{"HY000"},"%s"},
	ER_DD_TABLESPACE_NOT_FOUND : {10193,[]string{"HY000"},"Unable to start server. The data dictionary tablespace '%s' does not exist."},
	ER_DD_TRG_CONNECTION_COLLATION_MISSING : {10194,[]string{"HY000"},"Connection collation is missing for trigger of table %s. Using default connection collation."},
	ER_DD_TRG_DB_COLLATION_MISSING : {10195,[]string{"HY000"},"Database collation is missing for trigger of table %s. Using Default character set."},
	ER_DD_TRG_DEFINER_OOM : {10196,[]string{"HY000"},"Error in Memory allocation for Definer %s for Trigger."},
	ER_DD_TRG_FILE_UNREADABLE : {10197,[]string{"HY000"},"Error in reading %s.TRG file."},
	ER_TRG_CANT_PARSE : {10198,[]string{"HY000"},"Error in parsing Triggers from %s.TRG file."},
	ER_DD_TRG_CANT_ADD : {10199,[]string{"HY000"},"Error in creating DD entry for Trigger %s.%s"},
	ER_DD_CANT_RESOLVE_VIEW : {10200,[]string{"HY000"},"Resolving dependency for the view '%s.%s' failed. View is no more valid to use"},
	ER_DD_VIEW_WITHOUT_DEFINER : {10201,[]string{"HY000"},"%s.%s has no definer (as per an old view format). Current user is used as definer. Please recreate the view."},
	ER_PLUGIN_INIT_FAILED : {10202,[]string{"HY000"},"Plugin '%s' init function returned error."},
	ER_RPL_TRX_DELEGATES_INIT_FAILED : {10203,[]string{"HY000"},"Initialization of transaction delegates failed. Please report a bug."},
	ER_RPL_BINLOG_STORAGE_DELEGATES_INIT_FAILED : {10204,[]string{"HY000"},"Initialization binlog storage delegates failed. Please report a bug."},
	ER_RPL_BINLOG_TRANSMIT_DELEGATES_INIT_FAILED : {10205,[]string{"HY000"},"Initialization of binlog transmit delegates failed. Please report a bug."},
	ER_RPL_BINLOG_RELAY_DELEGATES_INIT_FAILED : {10206,[]string{"HY000"},"Initialization binlog relay IO delegates failed. Please report a bug."},
	ER_RPL_PLUGIN_FUNCTION_FAILED : {10207,[]string{"HY000"},"Run function '...' in plugin '%s' failed"},
	ER_SQL_HA_READ_FAILED : {10208,[]string{"HY000"},"mysql_ha_read: Got error %d when reading table '%s'"},
	ER_SR_BOGUS_VALUE : {10209,[]string{"HY000"},"Stored routine '%s'.'%s': invalid value in column %s."},
	ER_SR_INVALID_CONTEXT : {10210,[]string{"HY000"},"Invalid creation context '%s.%s'."},
	ER_READING_TABLE_FAILED : {10211,[]string{"HY000"},"Got error %d when reading table '%s'"},
	ER_DES_FILE_WRONG_KEY : {10212,[]string{"HY000"},"load_des_file:  Found wrong key_number: %c"},
	ER_CANT_SET_PERSISTED : {10213,[]string{"HY000"},"Failed to set persisted options."},
	ER_JSON_PARSE_ERROR : {10214,[]string{"HY000"},"Persisted config file is corrupt. Please ensure mysqld-auto.cnf file is valid JSON."},
	ER_CONFIG_OPTION_WITHOUT_GROUP : {10215,[]string{"HY000"},"Found option without preceding group in config file"},
	ER_VALGRIND_DO_QUICK_LEAK_CHECK : {10216,[]string{"HY000"},"VALGRIND_DO_QUICK_LEAK_CHECK"},
	ER_VALGRIND_COUNT_LEAKS : {10217,[]string{"HY000"},"VALGRIND_COUNT_LEAKS reports %lu leaked bytes for query '%.*s'"},
	ER_LOAD_DATA_INFILE_FAILED_IN_UNEXPECTED_WAY : {10218,[]string{"HY000"},"LOAD DATA INFILE in the slave SQL Thread can only read from --slave-load-tmpdir. Please, report a bug."},
	ER_UNKNOWN_ERROR_NUMBER : {10219,[]string{"HY000"},"Got unknown error: %d"},
	ER_UDF_CANT_ALLOC_FOR_STRUCTURES : {10220,[]string{"HY000"},"Can't allocate memory for udf structures"},
	ER_UDF_CANT_ALLOC_FOR_FUNCTION : {10221,[]string{"HY000"},"Can't alloc memory for udf function: '%.64s'"},
	ER_UDF_INVALID_ROW_IN_FUNCTION_TABLE : {10222,[]string{"HY000"},"Invalid row in mysql.func table for function '%.64s'"},
	ER_UDF_CANT_OPEN_FUNCTION_TABLE : {10223,[]string{"HY000"},"Could not open the mysql.func table. Please perform the MySQL upgrade procedure."},
	ER_XA_RECOVER_FOUND_TRX_IN_SE : {10224,[]string{"HY000"},"Found %d prepared transaction(s) in %s"},
	ER_XA_RECOVER_FOUND_XA_TRX : {10225,[]string{"HY000"},"Found %d prepared XA transactions"},
	ER_XA_IGNORING_XID : {10226,[]string{"HY000"},"ignore xid %s"},
	ER_XA_COMMITTING_XID : {10227,[]string{"HY000"},"commit xid %s"},
	ER_XA_ROLLING_BACK_XID : {10228,[]string{"HY000"},"rollback xid %s"},
	ER_XA_STARTING_RECOVERY : {10229,[]string{"HY000"},"Starting XA crash recovery..."},
	ER_XA_NO_MULTI_2PC_HEURISTIC_RECOVER : {10230,[]string{"HY000"},"--tc-heuristic-recover rollback strategy is not safe on systems with more than one 2-phase-commit-capable storage engine. Aborting crash recovery."},
	ER_XA_RECOVER_EXPLANATION : {10231,[]string{"HY000"},"Found %d prepared transactions! It means that mysqld was not shut down properly last time and critical recovery information (last binlog or %s file) was manually deleted after a crash. You have to start mysqld with --tc-heuristic-recover switch to commit or rollback pending transactions."},
	ER_XA_RECOVERY_DONE : {10232,[]string{"HY000"},"XA crash recovery finished."},
	ER_TRX_GTID_COLLECT_REJECT : {10233,[]string{"HY000"},"Failed to collect GTID to send in the response packet!"},
	ER_SQL_AUTHOR_DEFAULT_ROLES_FAIL : {10234,[]string{"HY000"},"MYSQL.DEFAULT_ROLES couldn't be updated for authorization identifier %s"},
	ER_SQL_USER_TABLE_CREATE_WARNING : {10235,[]string{"HY000"},"Following users were specified in CREATE USER IF NOT EXISTS but they already exist. Corresponding entry in binary log used default authentication plugin '%s' to rewrite authentication information (if any) for them: %s"},
	ER_SQL_USER_TABLE_ALTER_WARNING : {10236,[]string{"HY000"},"Following users were specified in ALTER USER IF EXISTS but they do not exist. Corresponding entry in binary log used default authentication plugin '%s' to rewrite authentication information (if any) for them: %s"},
	ER_ROW_IN_WRONG_PARTITION_PLEASE_REPAIR : {10237,[]string{"HY000"},"Table '%-192s' corrupted: row in wrong partition: %s -- Please REPAIR the table!"},
	ER_MYISAM_CRASHED_ERROR_IN_THREAD : {10238,[]string{"HY000"},"Got an error from thread_id=%u, %s:%d"},
	ER_MYISAM_CRASHED_ERROR_IN : {10239,[]string{"HY000"},"Got an error from unknown thread, %s:%d"},
	ER_TOO_MANY_STORAGE_ENGINES : {10240,[]string{"HY000"},"Too many storage engines!"},
	ER_SE_TYPECODE_CONFLICT : {10241,[]string{"HY000"},"Storage engine '%s' has conflicting typecode. Assigning value %d."},
	ER_TRX_WRITE_SET_OOM : {10242,[]string{"HY000"},"Out of memory on transaction write set extraction"},
	ER_HANDLERTON_OOM : {10243,[]string{"HY000"},"Unable to allocate memory for plugin '%s' handlerton."},
	ER_CONN_SHM_LISTENER : {10244,[]string{"HY000"},"Shared memory setting up listener"},
	ER_CONN_SHM_CANT_CREATE_SERVICE : {10245,[]string{"HY000"},"Can't create shared memory service: %s. : %s"},
	ER_CONN_SHM_CANT_CREATE_CONNECTION : {10246,[]string{"HY000"},"Can't create shared memory connection: %s. : %s"},
	ER_CONN_PIP_CANT_CREATE_EVENT : {10247,[]string{"HY000"},"Can't create event, last error=%u"},
	ER_CONN_PIP_CANT_CREATE_PIPE : {10248,[]string{"HY000"},"Can't create new named pipe!: %s"},
	ER_CONN_PER_THREAD_NO_THREAD : {10249,[]string{"HY000"},"Can't create thread to handle new connection(errno= %d)"},
	ER_CONN_TCP_NO_SOCKET : {10250,[]string{"HY000"},"Failed to create a socket for %s '%s': errno: %d."},
	ER_CONN_TCP_CREATED : {10251,[]string{"HY000"},"Server socket created on IP: '%s'."},
	ER_CONN_TCP_ADDRESS : {10252,[]string{"HY000"},"Server hostname (bind-address): '%s'; port: %d"},
	ER_CONN_TCP_IPV6_AVAILABLE : {10253,[]string{"HY000"},"IPv6 is available."},
	ER_CONN_TCP_IPV6_UNAVAILABLE : {10254,[]string{"HY000"},"IPv6 is not available."},
	ER_CONN_TCP_ERROR_WITH_STRERROR : {10255,[]string{"HY000"},"Can't create IP socket: %s"},
	ER_CONN_TCP_CANT_RESOLVE_HOSTNAME : {10256,[]string{"HY000"},"Can't start server: cannot resolve hostname!"},
	ER_CONN_TCP_IS_THERE_ANOTHER_USING_PORT : {10257,[]string{"HY000"},"Do you already have another mysqld server running on port: %d ?"},
	ER_CONN_UNIX_IS_THERE_ANOTHER_USING_SOCKET : {10258,[]string{"HY000"},"Do you already have another mysqld server running on socket: %s ?"},
	ER_CONN_UNIX_PID_CLAIMED_SOCKET_FILE : {10259,[]string{"HY000"},"Another process with pid %d is using unix socket file."},
	ER_CONN_TCP_CANT_RESET_V6ONLY : {10260,[]string{"HY000"},"Failed to reset IPV6_V6ONLY flag (error: %d). The server will listen to IPv6 addresses only."},
	ER_CONN_TCP_BIND_RETRY : {10261,[]string{"HY000"},"Retrying bind on TCP/IP port %u"},
	ER_CONN_TCP_BIND_FAIL : {10262,[]string{"HY000"},"Can't start server: Bind on TCP/IP port: %s"},
	ER_CONN_TCP_IP_NOT_LOGGED : {10263,[]string{"HY000"},"Fails to print out IP-address."},
	ER_CONN_TCP_RESOLVE_INFO : {10264,[]string{"HY000"},"  - '%s' resolves to '%s';"},
	ER_CONN_TCP_START_FAIL : {10265,[]string{"HY000"},"Can't start server: listen() on TCP/IP port: %s"},
	ER_CONN_TCP_LISTEN_FAIL : {10266,[]string{"HY000"},"listen() on TCP/IP failed with error %d"},
	ER_CONN_UNIX_PATH_TOO_LONG : {10267,[]string{"HY000"},"The socket file path is too long (> %u): %s"},
	ER_CONN_UNIX_LOCK_FILE_FAIL : {10268,[]string{"HY000"},"Unable to setup unix socket lock file."},
	ER_CONN_UNIX_NO_FD : {10269,[]string{"HY000"},"Can't start server: UNIX Socket : %s"},
	ER_CONN_UNIX_NO_BIND_NO_START : {10270,[]string{"HY000"},"Can't start server : Bind on unix socket: %s"},
	ER_CONN_UNIX_LISTEN_FAILED : {10271,[]string{"HY000"},"listen() on Unix socket failed with error %d"},
	ER_CONN_UNIX_LOCK_FILE_GIVING_UP : {10272,[]string{"HY000"},"Unable to create unix socket lock file %s after retries."},
	ER_CONN_UNIX_LOCK_FILE_CANT_CREATE : {10273,[]string{"HY000"},"Could not create unix socket lock file %s."},
	ER_CONN_UNIX_LOCK_FILE_CANT_OPEN : {10274,[]string{"HY000"},"Could not open unix socket lock file %s."},
	ER_CONN_UNIX_LOCK_FILE_CANT_READ : {10275,[]string{"HY000"},"Could not read unix socket lock file %s."},
	ER_CONN_UNIX_LOCK_FILE_EMPTY : {10276,[]string{"HY000"},"Unix socket lock file is empty %s."},
	ER_CONN_UNIX_LOCK_FILE_PIDLESS : {10277,[]string{"HY000"},"Invalid pid in unix socket lock file %s."},
	ER_CONN_UNIX_LOCK_FILE_CANT_WRITE : {10278,[]string{"HY000"},"Could not write unix socket lock file %s errno %d."},
	ER_CONN_UNIX_LOCK_FILE_CANT_DELETE : {10279,[]string{"HY000"},"Could not remove unix socket lock file %s errno %d."},
	ER_CONN_UNIX_LOCK_FILE_CANT_SYNC : {10280,[]string{"HY000"},"Could not sync unix socket lock file %s errno %d."},
	ER_CONN_UNIX_LOCK_FILE_CANT_CLOSE : {10281,[]string{"HY000"},"Could not close unix socket lock file %s errno %d."},
	ER_CONN_SOCKET_SELECT_FAILED : {10282,[]string{"HY000"},"mysqld: Got error %d from select"},
	ER_CONN_SOCKET_ACCEPT_FAILED : {10283,[]string{"HY000"},"Error in accept: %s"},
	ER_AUTH_RSA_CANT_FIND : {10284,[]string{"HY000"},"RSA %s key file not found: %s. Some authentication plugins will not work."},
	ER_AUTH_RSA_CANT_PARSE : {10285,[]string{"HY000"},"Failure to parse RSA %s key (file exists): %s: %s"},
	ER_AUTH_RSA_CANT_READ : {10286,[]string{"HY000"},"Failure to read key file: %s"},
	ER_AUTH_RSA_FILES_NOT_FOUND : {10287,[]string{"HY000"},"RSA key files not found. Some authentication plugins will not work."},
	ER_CONN_ATTR_TRUNCATED : {10288,[]string{"HY000"},"Connection attributes of length %lu were truncated (%d bytes lost) for connection %llu, user %s@%s (as %s), auth: %s"},
	ER_X509_CIPHERS_MISMATCH : {10289,[]string{"HY000"},"X.509 ciphers mismatch: should be '%s' but is '%s'"},
	ER_X509_ISSUER_MISMATCH : {10290,[]string{"HY000"},"X.509 issuer mismatch: should be '%s' but is '%s'"},
	ER_X509_SUBJECT_MISMATCH : {10291,[]string{"HY000"},"X.509 subject mismatch: should be '%s' but is '%s'"},
	ER_AUTH_CANT_ACTIVATE_ROLE : {10292,[]string{"HY000"},"Failed to activate default role %s for %s"},
	ER_X509_NEEDS_RSA_PRIVKEY : {10293,[]string{"HY000"},"Could not generate RSA private key required for X.509 certificate."},
	ER_X509_CANT_WRITE_KEY : {10294,[]string{"HY000"},"Could not write key file: %s"},
	ER_X509_CANT_CHMOD_KEY : {10295,[]string{"HY000"},"Could not set file permission for %s"},
	ER_X509_CANT_READ_CA_KEY : {10296,[]string{"HY000"},"Could not read CA key file: %s"},
	ER_X509_CANT_READ_CA_CERT : {10297,[]string{"HY000"},"Could not read CA certificate file: %s"},
	ER_X509_CANT_CREATE_CERT : {10298,[]string{"HY000"},"Could not generate X.509 certificate."},
	ER_X509_CANT_WRITE_CERT : {10299,[]string{"HY000"},"Could not write certificate file: %s"},
	ER_AUTH_CANT_CREATE_RSA_PAIR : {10300,[]string{"HY000"},"Could not generate RSA Private/Public key pair"},
	ER_AUTH_CANT_WRITE_PRIVKEY : {10301,[]string{"HY000"},"Could not write private key file: %s"},
	ER_AUTH_CANT_WRITE_PUBKEY : {10302,[]string{"HY000"},"Could not write public key file: %s"},
	ER_AUTH_SSL_CONF_PREVENTS_CERT_GENERATION : {10303,[]string{"HY000"},"Skipping generation of SSL certificates as options related to SSL are specified."},
	ER_AUTH_USING_EXISTING_CERTS : {10304,[]string{"HY000"},"Skipping generation of SSL certificates as certificate files are present in data directory."},
	ER_AUTH_CERTS_SAVED_TO_DATADIR : {10305,[]string{"HY000"},"Auto generated SSL certificates are placed in data directory."},
	ER_AUTH_CERT_GENERATION_DISABLED : {10306,[]string{"HY000"},"Skipping generation of SSL certificates as --auto_generate_certs is set to OFF."},
	ER_AUTH_RSA_CONF_PREVENTS_KEY_GENERATION : {10307,[]string{"HY000"},"Skipping generation of RSA key pair through %s as options related to RSA keys are specified."},
	ER_AUTH_KEY_GENERATION_SKIPPED_PAIR_PRESENT : {10308,[]string{"HY000"},"Skipping generation of RSA key pair through %s as key files are present in data directory."},
	ER_AUTH_KEYS_SAVED_TO_DATADIR : {10309,[]string{"HY000"},"Auto generated RSA key files through %s are placed in data directory."},
	ER_AUTH_KEY_GENERATION_DISABLED : {10310,[]string{"HY000"},"Skipping generation of RSA key pair as %s is set to OFF."},
	ER_AUTHCACHE_PROXIES_PRIV_SKIPPED_NEEDS_RESOLVE : {10311,[]string{"HY000"},"'proxies_priv' entry '%s@%s %s@%s' ignored in --skip-name-resolve mode."},
	ER_AUTHCACHE_PLUGIN_MISSING : {10312,[]string{"HY000"},"The plugin '%.*s' used to authenticate user '%s'@'%.*s' is not loaded. Nobody can currently login using this account."},
	ER_AUTHCACHE_PLUGIN_CONFIG : {10313,[]string{"HY000"},"The plugin '%s' is used to authenticate user '%s'@'%.*s', %s configured. Nobody can currently login using this account."},
	//OBSOLETE_ER_AUTHCACHE_ROLE_TABLES_DODGY : {10314,[]string{"HY000"},"Could not load mysql.role_edges and mysql.default_roles tables. ACL DDLs will not work unless the MySQL upgrade procedure is performed."},
	ER_AUTHCACHE_USER_SKIPPED_NEEDS_RESOLVE : {10315,[]string{"HY000"},"'user' entry '%s@%s' ignored in --skip-name-resolve mode."},
	ER_AUTHCACHE_USER_TABLE_DODGY : {10316,[]string{"HY000"},"Fatal error: Could not read the column 'authentication_string' from table 'mysql.user'. Please perform the MySQL upgrade procedure."},
	ER_AUTHCACHE_USER_IGNORED_DEPRECATED_PASSWORD : {10317,[]string{"HY000"},"User entry '%s'@'%s' has a deprecated pre-4.1 password. The user will be ignored and no one can login with this user anymore."},
	ER_AUTHCACHE_USER_IGNORED_NEEDS_PLUGIN : {10318,[]string{"HY000"},"User entry '%s'@'%s' has an empty plugin value. The user will be ignored and no one can login with this user anymore."},
	ER_AUTHCACHE_USER_IGNORED_INVALID_PASSWORD : {10319,[]string{"HY000"},"Found invalid password for user: '%s@%s'; Ignoring user"},
	ER_AUTHCACHE_EXPIRED_PASSWORD_UNSUPPORTED : {10320,[]string{"HY000"},"'user' entry '%s@%s' has the password ignore flag raised, but its authentication plugin doesn't support password expiration. The user id will be ignored."},
	ER_NO_SUPER_WITHOUT_USER_PLUGIN : {10321,[]string{"HY000"},"Some of the user accounts with SUPER privileges were disabled because of empty mysql.user.plugin value. If you are upgrading from MySQL 5.6 to MySQL 5.7 it means that substitution for the empty plugin column was not possible. Probably because of pre 4.1 password hash. If your account is disabled you will need to perform the MySQL upgrade procedure. For complete instructions on how to upgrade MySQL to a new version please see the 'Upgrading MySQL' section from the MySQL manual."},
	ER_AUTHCACHE_DB_IGNORED_EMPTY_NAME : {10322,[]string{"HY000"},"Found an entry in the 'db' table with empty database name; Skipped"},
	ER_AUTHCACHE_DB_SKIPPED_NEEDS_RESOLVE : {10323,[]string{"HY000"},"'db' entry '%s %s@%s' ignored in --skip-name-resolve mode."},
	ER_AUTHCACHE_DB_ENTRY_LOWERCASED_REVOKE_WILL_FAIL : {10324,[]string{"HY000"},"'db' entry '%s %s@%s' had database in mixed case that has been forced to lowercase because lower_case_table_names is set. It will not be possible to remove this privilege using REVOKE."},
	ER_AUTHCACHE_TABLE_PROXIES_PRIV_MISSING : {10325,[]string{"HY000"},"The system table mysql.proxies_priv is missing. Please perform the MySQL upgrade procedure."},
	ER_AUTHCACHE_CANT_OPEN_AND_LOCK_PRIVILEGE_TABLES : {10326,[]string{"HY000"},"Fatal error: Can't open and lock privilege tables: %s"},
	ER_AUTHCACHE_CANT_INIT_GRANT_SUBSYSTEM : {10327,[]string{"HY000"},"Fatal: can't initialize grant subsystem - '%s'"},
	ER_AUTHCACHE_PROCS_PRIV_SKIPPED_NEEDS_RESOLVE : {10328,[]string{"HY000"},"'procs_priv' entry '%s %s@%s' ignored in --skip-name-resolve mode."},
	ER_AUTHCACHE_PROCS_PRIV_ENTRY_IGNORED_BAD_ROUTINE_TYPE : {10329,[]string{"HY000"},"'procs_priv' entry '%s' ignored, bad routine type"},
	ER_AUTHCACHE_TABLES_PRIV_SKIPPED_NEEDS_RESOLVE : {10330,[]string{"HY000"},"'tables_priv' entry '%s %s@%s' ignored in --skip-name-resolve mode."},
	ER_USER_NOT_IN_EXTRA_USERS_BINLOG_POSSIBLY_INCOMPLETE : {10331,[]string{"HY000"},"Failed to add %s in extra_users. Binary log entry may miss some of the users."},
	ER_DD_SCHEMA_NOT_FOUND : {10332,[]string{"HY000"},"Unable to start server. The data dictionary schema '%s' does not exist."},
	ER_DD_TABLE_NOT_FOUND : {10333,[]string{"HY000"},"Unable to start server. The data dictionary table '%s' does not exist."},
	ER_DD_SE_INIT_FAILED : {10334,[]string{"HY000"},"Failed to initialize DD Storage Engine"},
	ER_DD_ABORTING_PARTIAL_UPGRADE : {10335,[]string{"HY000"},"Found partially upgraded DD. Aborting upgrade and deleting all DD tables. Start the upgrade process again."},
	ER_DD_FRM_EXISTS_FOR_TABLE : {10336,[]string{"HY000"},"Found .frm file with same name as one of the Dictionary Tables."},
	ER_DD_CREATED_FOR_UPGRADE : {10337,[]string{"HY000"},"Created Data Dictionary for upgrade"},
	ER_ERRMSG_CANT_FIND_FILE : {10338,[]string{"HY000"},"Can't find error-message file '%s'. Check error-message file location and 'lc-messages-dir' configuration directive."},
	ER_ERRMSG_LOADING_55_STYLE : {10339,[]string{"HY000"},"Using pre 5.5 semantics to load error messages from %s. If this is not intended, refer to the documentation for valid usage of --lc-messages-dir and --language parameters."},
	ER_ERRMSG_MISSING_IN_FILE : {10340,[]string{"HY000"},"Error message file '%s' had only %d error messages, but it should contain at least %d error messages. Check that the above file is the right version for this program!"},
	ER_ERRMSG_OOM : {10341,[]string{"HY000"},"Not enough memory for messagefile '%s'"},
	ER_ERRMSG_CANT_READ : {10342,[]string{"HY000"},"Can't read from messagefile '%s'"},
	ER_TABLE_INCOMPATIBLE_DECIMAL_FIELD : {10343,[]string{"HY000"},"Found incompatible DECIMAL field '%s' in %s; Please do \"ALTER TABLE `%s` FORCE\" to fix it!"},
	ER_TABLE_INCOMPATIBLE_YEAR_FIELD : {10344,[]string{"HY000"},"Found incompatible YEAR(x) field '%s' in %s; Please do \"ALTER TABLE `%s` FORCE\" to fix it!"},
	ER_INVALID_CHARSET_AND_DEFAULT_IS_MB : {10345,[]string{"HY000"},"'%s' had no or invalid character set, and default character set is multi-byte, so character column sizes may have changed"},
	ER_TABLE_WRONG_KEY_DEFINITION : {10346,[]string{"HY000"},"Found wrong key definition in %s; Please do \"ALTER TABLE `%s` FORCE \" to fix it!"},
	ER_CANT_OPEN_FRM_FILE : {10347,[]string{"HY000"},"Unable to open file %s"},
	ER_CANT_READ_FRM_FILE : {10348,[]string{"HY000"},"Error in reading file %s"},
	ER_TABLE_CREATED_WITH_DIFFERENT_VERSION : {10349,[]string{"HY000"},"Table '%s' was created with a different version of MySQL and cannot be read"},
	ER_VIEW_UNPARSABLE : {10350,[]string{"HY000"},"Unable to read view %s"},
	ER_FILE_TYPE_UNKNOWN : {10351,[]string{"HY000"},"File %s has unknown type in its header."},
	ER_INVALID_INFO_IN_FRM : {10352,[]string{"HY000"},"Incorrect information in file %s"},
	ER_CANT_OPEN_AND_LOCK_PRIVILEGE_TABLES : {10353,[]string{"HY000"},"Can't open and lock privilege tables: %s"},
	ER_AUDIT_PLUGIN_DOES_NOT_SUPPORT_AUDIT_AUTH_EVENTS : {10354,[]string{"HY000"},"Plugin '%s' cannot subscribe to MYSQL_AUDIT_AUTHORIZATION events. Currently not supported."},
	ER_AUDIT_PLUGIN_HAS_INVALID_DATA : {10355,[]string{"HY000"},"Plugin '%s' has invalid data."},
	ER_TZ_OOM_INITIALIZING_TIME_ZONES : {10356,[]string{"HY000"},"Fatal error: OOM while initializing time zones"},
	ER_TZ_CANT_OPEN_AND_LOCK_TIME_ZONE_TABLE : {10357,[]string{"HY000"},"Can't open and lock time zone table: %s trying to live without them"},
	ER_TZ_OOM_LOADING_LEAP_SECOND_TABLE : {10358,[]string{"HY000"},"Fatal error: Out of memory while loading mysql.time_zone_leap_second table"},
	ER_TZ_TOO_MANY_LEAPS_IN_LEAP_SECOND_TABLE : {10359,[]string{"HY000"},"Fatal error: While loading mysql.time_zone_leap_second table: too much leaps"},
	ER_TZ_ERROR_LOADING_LEAP_SECOND_TABLE : {10360,[]string{"HY000"},"Fatal error: Error while loading mysql.time_zone_leap_second table"},
	ER_TZ_UNKNOWN_OR_ILLEGAL_DEFAULT_TIME_ZONE : {10361,[]string{"HY000"},"Fatal error: Illegal or unknown default time zone '%s'"},
	ER_TZ_CANT_FIND_DESCRIPTION_FOR_TIME_ZONE : {10362,[]string{"HY000"},"Can't find description of time zone '%.*s'"},
	ER_TZ_CANT_FIND_DESCRIPTION_FOR_TIME_ZONE_ID : {10363,[]string{"HY000"},"Can't find description of time zone '%u'"},
	ER_TZ_TRANSITION_TYPE_TABLE_TYPE_TOO_LARGE : {10364,[]string{"HY000"},"Error while loading time zone description from mysql.time_zone_transition_type table: too big transition type id"},
	ER_TZ_TRANSITION_TYPE_TABLE_ABBREVIATIONS_EXCEED_SPACE : {10365,[]string{"HY000"},"Error while loading time zone description from mysql.time_zone_transition_type table: not enough room for abbreviations"},
	ER_TZ_TRANSITION_TYPE_TABLE_LOAD_ERROR : {10366,[]string{"HY000"},"Error while loading time zone description from mysql.time_zone_transition_type table"},
	ER_TZ_TRANSITION_TABLE_TOO_MANY_TRANSITIONS : {10367,[]string{"HY000"},"Error while loading time zone description from mysql.time_zone_transition table: too much transitions"},
	ER_TZ_TRANSITION_TABLE_BAD_TRANSITION_TYPE : {10368,[]string{"HY000"},"Error while loading time zone description from mysql.time_zone_transition table: bad transition type id"},
	ER_TZ_TRANSITION_TABLE_LOAD_ERROR : {10369,[]string{"HY000"},"Error while loading time zone description from mysql.time_zone_transition table"},
	ER_TZ_NO_TRANSITION_TYPES_IN_TIME_ZONE : {10370,[]string{"HY000"},"loading time zone without transition types"},
	ER_TZ_OOM_LOADING_TIME_ZONE_DESCRIPTION : {10371,[]string{"HY000"},"Out of memory while loading time zone description"},
	ER_TZ_CANT_BUILD_MKTIME_MAP : {10372,[]string{"HY000"},"Unable to build mktime map for time zone"},
	ER_TZ_OOM_WHILE_LOADING_TIME_ZONE : {10373,[]string{"HY000"},"Out of memory while loading time zone"},
	ER_TZ_OOM_WHILE_SETTING_TIME_ZONE : {10374,[]string{"HY000"},"Fatal error: Out of memory while setting new time zone"},
	ER_SLAVE_SQL_THREAD_STOPPED_UNTIL_CONDITION_BAD : {10375,[]string{"HY000"},"Slave SQL thread is stopped because UNTIL condition is bad(%s:%llu)."},
	ER_SLAVE_SQL_THREAD_STOPPED_UNTIL_POSITION_REACHED : {10376,[]string{"HY000"},"Slave SQL thread stopped because it reached its UNTIL position %llu"},
	ER_SLAVE_SQL_THREAD_STOPPED_BEFORE_GTIDS_ALREADY_APPLIED : {10377,[]string{"HY000"},"Slave SQL thread stopped because UNTIL SQL_BEFORE_GTIDS %s is already applied"},
	ER_SLAVE_SQL_THREAD_STOPPED_BEFORE_GTIDS_REACHED : {10378,[]string{"HY000"},"Slave SQL thread stopped because it reached UNTIL SQL_BEFORE_GTIDS %s"},
	ER_SLAVE_SQL_THREAD_STOPPED_AFTER_GTIDS_REACHED : {10379,[]string{"HY000"},"Slave SQL thread stopped because it reached UNTIL SQL_AFTER_GTIDS %s"},
	ER_SLAVE_SQL_THREAD_STOPPED_GAP_TRX_PROCESSED : {10380,[]string{"HY000"},"Slave SQL thread stopped according to UNTIL SQL_AFTER_MTS_GAPS as it has processed all gap transactions left from the previous slave session."},
	ER_GROUP_REPLICATION_PLUGIN_NOT_INSTALLED : {10381,[]string{"HY000"},"Group Replication plugin is not installed."},
	ER_GTID_ALREADY_ADDED_BY_USER : {10382,[]string{"HY000"},"The transaction owned GTID is already in the %s table, which is caused by an explicit modifying from user client."},
	ER_FAILED_TO_DELETE_FROM_GTID_EXECUTED_TABLE : {10383,[]string{"HY000"},"Failed to delete the row: '%s' from the gtid_executed table."},
	ER_FAILED_TO_COMPRESS_GTID_EXECUTED_TABLE : {10384,[]string{"HY000"},"Failed to compress the gtid_executed table."},
	ER_FAILED_TO_COMPRESS_GTID_EXECUTED_TABLE_OOM : {10385,[]string{"HY000"},"Failed to compress the gtid_executed table, because it is failed to allocate the THD."},
	ER_FAILED_TO_INIT_THREAD_ATTR_FOR_GTID_TABLE_COMPRESSION : {10386,[]string{"HY000"},"Failed to initialize thread attribute when creating compression thread."},
	ER_FAILED_TO_CREATE_GTID_TABLE_COMPRESSION_THREAD : {10387,[]string{"HY000"},"Can not create thread to compress gtid_executed table (errno= %d)"},
	ER_FAILED_TO_JOIN_GTID_TABLE_COMPRESSION_THREAD : {10388,[]string{"HY000"},"Could not join gtid_executed table compression thread. error:%d"},
	ER_NPIPE_FAILED_TO_INIT_SECURITY_DESCRIPTOR : {10389,[]string{"HY000"},"Can't start server : Initialize security descriptor: %s"},
	ER_NPIPE_FAILED_TO_SET_SECURITY_DESCRIPTOR : {10390,[]string{"HY000"},"Can't start server : Set security descriptor: %s"},
	ER_NPIPE_PIPE_ALREADY_IN_USE : {10391,[]string{"HY000"},"Can't start server : Named Pipe \"%s\" already in use."},
	//OBSOLETE_ER_NDB_SLAVE_SAW_EPOCH_LOWER_THAN_PREVIOUS_ON_START : {10392,[]string{"HY000"},"NDB Slave : At SQL thread start applying epoch %llu/%llu (%llu) from Master ServerId %u which is lower than previously applied epoch %llu/%llu (%llu).  Group Master Log : %s  Group Master Log Pos : %llu.  Check slave positioning."},
	//OBSOLETE_ER_NDB_SLAVE_SAW_EPOCH_LOWER_THAN_PREVIOUS : {10393,[]string{"HY000"},"NDB Slave : SQL thread stopped as applying epoch %llu/%llu (%llu) from Master ServerId %u which is lower than previously applied epoch %llu/%llu (%llu).  Group Master Log : %s  Group Master Log Pos : %llu"},
	//OBSOLETE_ER_NDB_SLAVE_SAW_ALREADY_COMMITTED_EPOCH : {10394,[]string{"HY000"},"NDB Slave : SQL thread stopped as attempted to reapply already committed epoch %llu/%llu (%llu) from server id %u.  Group Master Log : %s  Group Master Log Pos : %llu."},
	//OBSOLETE_ER_NDB_SLAVE_PREVIOUS_EPOCH_NOT_COMMITTED : {10395,[]string{"HY000"},"NDB Slave : SQL thread stopped as attempting to apply new epoch %llu/%llu (%llu) while lower received epoch %llu/%llu (%llu) has not been committed.  Master server id : %u.  Group Master Log : %s  Group Master Log Pos : %llu."},
	//OBSOLETE_ER_NDB_SLAVE_MISSING_DATA_FOR_TIMESTAMP_COLUMN : {10396,[]string{"HY000"},"NDB Slave: missing data for %s timestamp column %u."},
	//OBSOLETE_ER_NDB_SLAVE_LOGGING_EXCEPTIONS_TO : {10397,[]string{"HY000"},"NDB Slave: Table %s.%s logging exceptions to %s.%s"},
	//OBSOLETE_ER_NDB_SLAVE_LOW_EPOCH_RESOLUTION : {10398,[]string{"HY000"},"NDB Slave: Table %s.%s : %s, low epoch resolution"},
	//OBSOLETE_ER_NDB_INFO_FOUND_UNEXPECTED_FIELD_TYPE : {10399,[]string{"HY000"},"Found unexpected field type %u"},
	//OBSOLETE_ER_NDB_INFO_FAILED_TO_CREATE_NDBINFO : {10400,[]string{"HY000"},"Failed to create NdbInfo"},
	//OBSOLETE_ER_NDB_INFO_FAILED_TO_INIT_NDBINFO : {10401,[]string{"HY000"},"Failed to init NdbInfo"},
	//OBSOLETE_ER_NDB_CLUSTER_WRONG_NUMBER_OF_FUNCTION_ARGUMENTS : {10402,[]string{"HY000"},"ndb_serialize_cond: Unexpected mismatch of found and expected number of function arguments %u"},
	//OBSOLETE_ER_NDB_CLUSTER_SCHEMA_INFO : {10403,[]string{"HY000"},"%s - %s.%s"},
	//OBSOLETE_ER_NDB_CLUSTER_GENERIC_MESSAGE : {10404,[]string{"HY000"},"%s"},
	ER_RPL_CANT_OPEN_INFO_TABLE : {10405,[]string{"HY000"},"Info table is not ready to be used. Table '%s.%s' cannot be opened."},
	ER_RPL_CANT_SCAN_INFO_TABLE : {10406,[]string{"HY000"},"Info table is not ready to be used. Table '%s.%s' cannot be scanned."},
	ER_RPL_CORRUPTED_INFO_TABLE : {10407,[]string{"HY000"},"Corrupted table %s.%s. Check out table definition."},
	ER_RPL_CORRUPTED_KEYS_IN_INFO_TABLE : {10408,[]string{"HY000"},"Info table has a problem with its key field(s). Table '%s.%s' expected field #%u to be '%s' but found '%s' instead."},
	ER_RPL_WORKER_ID_IS : {10409,[]string{"HY000"},"Choosing worker id %lu, the following is going to be %lu"},
	ER_RPL_INCONSISTENT_TIMESTAMPS_IN_TRX : {10410,[]string{"HY000"},"Transaction is tagged with inconsistent logical timestamps: sequence_number (%lld) <= last_committed (%lld)"},
	ER_RPL_INCONSISTENT_SEQUENCE_NO_IN_TRX : {10411,[]string{"HY000"},"Transaction's sequence number is inconsistent with that of a preceding one: sequence_number (%lld) <= previous sequence_number (%lld)"},
	ER_RPL_CHANNELS_REQUIRE_TABLES_AS_INFO_REPOSITORIES : {10412,[]string{"HY000"},"For the creation of replication channels the master info and relay log info repositories must be set to TABLE"},
	ER_RPL_CHANNELS_REQUIRE_NON_ZERO_SERVER_ID : {10413,[]string{"HY000"},"For the creation of replication channels the server id must be different from 0"},
	ER_RPL_REPO_SHOULD_BE_TABLE : {10414,[]string{"HY000"},"Slave: Wrong repository. Repository should be TABLE"},
	ER_RPL_ERROR_CREATING_MASTER_INFO : {10415,[]string{"HY000"},"Error creating master info: %s."},
	ER_RPL_ERROR_CHANGING_MASTER_INFO_REPO_TYPE : {10416,[]string{"HY000"},"Error changing the type of master info's repository: %s."},
	ER_RPL_CHANGING_RELAY_LOG_INFO_REPO_TYPE_FAILED_DUE_TO_GAPS : {10417,[]string{"HY000"},"It is not possible to change the type of the relay log repository because there are workers repositories with possible execution gaps. The value of --relay_log_info_repository is altered to one of the found Worker repositories. The gaps have to be sorted out before resuming with the type change."},
	ER_RPL_ERROR_CREATING_RELAY_LOG_INFO : {10418,[]string{"HY000"},"Error creating relay log info: %s."},
	ER_RPL_ERROR_CHANGING_RELAY_LOG_INFO_REPO_TYPE : {10419,[]string{"HY000"},"Error changing the type of relay log info's repository: %s."},
	ER_RPL_FAILED_TO_DELETE_FROM_SLAVE_WORKERS_INFO_REPOSITORY : {10420,[]string{"HY000"},"Could not delete from Slave Workers info repository."},
	ER_RPL_FAILED_TO_RESET_STATE_IN_SLAVE_INFO_REPOSITORY : {10421,[]string{"HY000"},"Could not store the reset Slave Worker state into the slave info repository."},
	ER_RPL_ERROR_CHECKING_REPOSITORY : {10422,[]string{"HY000"},"Error in checking %s repository info type of %s."},
	ER_RPL_SLAVE_GENERIC_MESSAGE : {10423,[]string{"HY000"},"Slave: %s"},
	ER_RPL_SLAVE_COULD_NOT_CREATE_CHANNEL_LIST : {10424,[]string{"HY000"},"Slave: Could not create channel list"},
	ER_RPL_MULTISOURCE_REQUIRES_TABLE_TYPE_REPOSITORIES : {10425,[]string{"HY000"},"Slave: This slave was a multisourced slave previously which is supported only by both TABLE based master info and relay log info repositories. Found one or both of the info repos to be type FILE. Set both repos to type TABLE."},
	ER_RPL_SLAVE_FAILED_TO_INIT_A_MASTER_INFO_STRUCTURE : {10426,[]string{"HY000"},"Slave: Failed to initialize the master info structure for channel '%s'; its record may still be present in 'mysql.slave_master_info' table, consider deleting it."},
	ER_RPL_SLAVE_FAILED_TO_INIT_MASTER_INFO_STRUCTURE : {10427,[]string{"HY000"},"Failed to initialize the master info structure%s"},
	ER_RPL_SLAVE_FAILED_TO_CREATE_CHANNEL_FROM_MASTER_INFO : {10428,[]string{"HY000"},"Slave: Failed to create a channel from master info table repository."},
	ER_RPL_FAILED_TO_CREATE_NEW_INFO_FILE : {10429,[]string{"HY000"},"Failed to create a new info file (file '%s', errno %d)"},
	ER_RPL_FAILED_TO_CREATE_CACHE_FOR_INFO_FILE : {10430,[]string{"HY000"},"Failed to create a cache on info file (file '%s')"},
	ER_RPL_FAILED_TO_OPEN_INFO_FILE : {10431,[]string{"HY000"},"Failed to open the existing info file (file '%s', errno %d)"},
	ER_RPL_GTID_MEMORY_FINALLY_AVAILABLE : {10432,[]string{"HY000"},"Server overcomes the temporary 'out of memory' in '%d' tries while allocating a new chunk of intervals for storing GTIDs."},
	ER_SERVER_COST_UNKNOWN_COST_CONSTANT : {10433,[]string{"HY000"},"Unknown cost constant \"%s\" in mysql.server_cost table"},
	ER_SERVER_COST_INVALID_COST_CONSTANT : {10434,[]string{"HY000"},"Invalid value for cost constant \"%s\" in mysql.server_cost table: %.1f"},
	ER_ENGINE_COST_UNKNOWN_COST_CONSTANT : {10435,[]string{"HY000"},"Unknown cost constant \"%s\" in mysql.engine_cost table"},
	ER_ENGINE_COST_UNKNOWN_STORAGE_ENGINE : {10436,[]string{"HY000"},"Unknown storage engine \"%s\" in mysql.engine_cost table"},
	ER_ENGINE_COST_INVALID_DEVICE_TYPE_FOR_SE : {10437,[]string{"HY000"},"Invalid device type %d for \"%s\" storage engine for cost constant \"%s\" in mysql.engine_cost table"},
	ER_ENGINE_COST_INVALID_CONST_CONSTANT_FOR_SE_AND_DEVICE : {10438,[]string{"HY000"},"Invalid value for cost constant \"%s\" for \"%s\" storage engine and device type %d in mysql.engine_cost table: %.1f"},
	ER_SERVER_COST_FAILED_TO_READ : {10439,[]string{"HY000"},"Error while reading from mysql.server_cost table."},
	ER_ENGINE_COST_FAILED_TO_READ : {10440,[]string{"HY000"},"Error while reading from mysql.engine_cost table."},
	ER_FAILED_TO_OPEN_COST_CONSTANT_TABLES : {10441,[]string{"HY000"},"Failed to open optimizer cost constant tables"},
	ER_RPL_UNSUPPORTED_UNIGNORABLE_EVENT_IN_STREAM : {10442,[]string{"HY000"},"Unsupported non-ignorable event fed into the event stream."},
	ER_RPL_GTID_LOG_EVENT_IN_STREAM : {10443,[]string{"HY000"},"GTID_LOG_EVENT or ANONYMOUS_GTID_LOG_EVENT is not expected in an event stream %s."},
	ER_RPL_UNEXPECTED_BEGIN_IN_STREAM : {10444,[]string{"HY000"},"QUERY(BEGIN) is not expected in an event stream in the middle of a %s."},
	ER_RPL_UNEXPECTED_COMMIT_ROLLBACK_OR_XID_LOG_EVENT_IN_STREAM : {10445,[]string{"HY000"},"QUERY(COMMIT or ROLLBACK) or XID_LOG_EVENT is not expected in an event stream %s."},
	ER_RPL_UNEXPECTED_XA_ROLLBACK_IN_STREAM : {10446,[]string{"HY000"},"QUERY(XA ROLLBACK) is not expected in an event stream %s."},
	ER_EVENT_EXECUTION_FAILED_CANT_AUTHENTICATE_USER : {10447,[]string{"HY000"},"Event Scheduler: [%s].[%s.%s] execution failed, failed to authenticate the user."},
	ER_EVENT_EXECUTION_FAILED_USER_LOST_EVEN_PRIVILEGE : {10448,[]string{"HY000"},"Event Scheduler: [%s].[%s.%s] execution failed, user no longer has EVENT privilege."},
	ER_EVENT_ERROR_DURING_COMPILATION : {10449,[]string{"HY000"},"Event Scheduler: %serror during compilation of %s.%s"},
	ER_EVENT_DROPPING : {10450,[]string{"HY000"},"Event Scheduler: Dropping %s.%s"},
	//OBSOLETE_ER_NDB_SCHEMA_GENERIC_MESSAGE : {10451,[]string{"HY000"},"Ndb schema[%s.%s]: %s"},
	ER_RPL_INCOMPATIBLE_DECIMAL_IN_RBR : {10452,[]string{"HY000"},"In RBR mode, Slave received incompatible DECIMAL field (old-style decimal field) from Master while creating conversion table. Please consider changing datatype on Master to new style decimal by executing ALTER command for column Name: %s.%s.%s."},
	ER_INIT_ROOT_WITHOUT_PASSWORD : {10453,[]string{"HY000"},"root@localhost is created with an empty password ! Please consider switching off the --initialize-insecure option."},
	ER_INIT_GENERATING_TEMP_PASSWORD_FOR_ROOT : {10454,[]string{"HY000"},"A temporary password is generated for root@localhost: %s"},
	ER_INIT_CANT_OPEN_BOOTSTRAP_FILE : {10455,[]string{"HY000"},"Failed to open the bootstrap file %s"},
	ER_INIT_BOOTSTRAP_COMPLETE : {10456,[]string{"HY000"},"Bootstrapping complete"},
	ER_INIT_DATADIR_NOT_EMPTY_WONT_INITIALIZE : {10457,[]string{"HY000"},"--initialize specified but the data directory has files in it. Aborting."},
	ER_INIT_DATADIR_EXISTS_WONT_INITIALIZE : {10458,[]string{"HY000"},"--initialize specified on an existing data directory."},
	ER_INIT_DATADIR_EXISTS_AND_PATH_TOO_LONG_WONT_INITIALIZE : {10459,[]string{"HY000"},"--initialize specified but the data directory exists and the path is too long. Aborting."},
	ER_INIT_DATADIR_EXISTS_AND_NOT_WRITABLE_WONT_INITIALIZE : {10460,[]string{"HY000"},"--initialize specified but the data directory exists and is not writable. Aborting."},
	ER_INIT_CREATING_DD : {10461,[]string{"HY000"},"Creating the data directory %s"},
	ER_RPL_BINLOG_STARTING_DUMP : {10462,[]string{"HY000"},"Start binlog_dump to master_thread_id(%u) slave_server(%u), pos(%s, %llu)"},
	ER_RPL_BINLOG_MASTER_SENDS_HEARTBEAT : {10463,[]string{"HY000"},"master sends heartbeat message"},
	ER_RPL_BINLOG_SKIPPING_REMAINING_HEARTBEAT_INFO : {10464,[]string{"HY000"},"the rest of heartbeat info skipped ..."},
	ER_RPL_BINLOG_MASTER_USES_CHECKSUM_AND_SLAVE_CANT : {10465,[]string{"HY000"},"Master is configured to log replication events with checksum, but will not send such events to slaves that cannot process them"},
	//OBSOLETE_ER_NDB_QUERY_FAILED : {10466,[]string{"HY000"},"NDB: Query '%s' failed, error: %d: %s"},
	ER_KILLING_THREAD : {10467,[]string{"HY000"},"Killing thread %lu"},
	ER_DETACHING_SESSION_LEFT_BY_PLUGIN : {10468,[]string{"HY000"},"Plugin %s is deinitializing a thread but left a session attached. Detaching it forcefully."},
	ER_CANT_DETACH_SESSION_LEFT_BY_PLUGIN : {10469,[]string{"HY000"},"Failed to detach the session."},
	ER_DETACHED_SESSIONS_LEFT_BY_PLUGIN : {10470,[]string{"HY000"},"Closed forcefully %u session%s left opened by plugin %s"},
	ER_FAILED_TO_DECREMENT_NUMBER_OF_THREADS : {10471,[]string{"HY000"},"Failed to decrement the number of threads"},
	ER_PLUGIN_DID_NOT_DEINITIALIZE_THREADS : {10472,[]string{"HY000"},"Plugin %s did not deinitialize %u threads"},
	ER_KILLED_THREADS_OF_PLUGIN : {10473,[]string{"HY000"},"Killed %u threads of plugin %s"},
	//OBSOLETE_ER_NDB_SLAVE_MAX_REPLICATED_EPOCH_UNKNOWN : {10474,[]string{"HY000"},"NDB Slave : Could not determine maximum replicated epoch from %s.%s at Slave start, error %u %s"},
	//OBSOLETE_ER_NDB_SLAVE_MAX_REPLICATED_EPOCH_SET_TO : {10475,[]string{"HY000"},"NDB Slave : MaxReplicatedEpoch set to %llu (%u/%u) at Slave start"},
	//OBSOLETE_ER_NDB_NODE_ID_AND_MANAGEMENT_SERVER_INFO : {10476,[]string{"HY000"},"NDB: NodeID is %lu, management server '%s:%lu'"},
	//OBSOLETE_ER_NDB_DISCONNECT_INFO : {10477,[]string{"HY000"},"tid %u: node[%u] transaction_hint=%u, transaction_no_hint=%u"},
	//OBSOLETE_ER_NDB_COLUMN_DEFAULTS_DIFFER : {10478,[]string{"HY000"},"NDB Internal error: Default values differ for column %u, ndb_default: %d"},
	//OBSOLETE_ER_NDB_COLUMN_SHOULD_NOT_HAVE_NATIVE_DEFAULT : {10479,[]string{"HY000"},"NDB Internal error: Column %u has native default, but shouldn't. Flags=%u, type=%u"},
	//OBSOLETE_ER_NDB_FIELD_INFO : {10480,[]string{"HY000"},"field[ name: '%s', type: %u, real_type: %u, flags: 0x%x, is_null: %d]"},
	//OBSOLETE_ER_NDB_COLUMN_INFO : {10481,[]string{"HY000"},"ndbCol[name: '%s', type: %u, column_no: %d, nullable: %d]"},
	//OBSOLETE_ER_NDB_OOM_IN_FIX_UNIQUE_INDEX_ATTR_ORDER : {10482,[]string{"HY000"},"fix_unique_index_attr_order: my_malloc(%u) failure"},
	//OBSOLETE_ER_NDB_SLAVE_MALFORMED_EVENT_RECEIVED_ON_TABLE : {10483,[]string{"HY000"},"NDB Slave : Malformed event received on table %s cannot parse.  Stopping Slave."},
	//OBSOLETE_ER_NDB_SLAVE_CONFLICT_FUNCTION_REQUIRES_ROLE : {10484,[]string{"HY000"},"NDB Slave : Conflict function %s defined on table %s requires ndb_slave_conflict_role variable to be set.  Stopping slave."},
	//OBSOLETE_ER_NDB_SLAVE_CONFLICT_TRANSACTION_IDS : {0000,[]string{""},"NDB Slave : Transactional conflict detection defined on table %s, but events received without transaction ids.  Check --ndb-log-transaction-id setting on upstream Cluster."},
	//OBSOLETE_ER_NDB_SLAVE_BINLOG_MISSING_INFO_FOR_CONFLICT_DETECTION : {10486,[]string{"HY000"},"NDB Slave : Binlog event on table %s missing info necessary for conflict detection.  Check binlog format options on upstream cluster."},
	//OBSOLETE_ER_NDB_ERROR_IN_READAUTOINCREMENTVALUE : {10487,[]string{"HY000"},"Error %lu in readAutoIncrementValue(): %s"},
	//OBSOLETE_ER_NDB_FOUND_UNCOMMITTED_AUTOCOMMIT : {10488,[]string{"HY000"},"found uncommitted autocommit+rbwr transaction, commit status: %d"},
	//OBSOLETE_ER_NDB_SLAVE_TOO_MANY_RETRIES : {10489,[]string{"HY000"},"Ndb slave retried transaction %u time(s) in vain.  Giving up."},
	//OBSOLETE_ER_NDB_SLAVE_ERROR_IN_UPDATE_CREATE_INFO : {10490,[]string{"HY000"},"Error %lu in ::update_create_info(): %s"},
	//OBSOLETE_ER_NDB_SLAVE_CANT_ALLOCATE_TABLE_SHARE : {10491,[]string{"HY000"},"NDB: allocating table share for %s failed"},
	//OBSOLETE_ER_NDB_BINLOG_ERROR_INFO_FROM_DA : {10492,[]string{"HY000"},"NDB Binlog: (%d)%s"},
	//OBSOLETE_ER_NDB_BINLOG_CREATE_TABLE_EVENT : {10493,[]string{"HY000"},"NDB Binlog: CREATE TABLE Event: %s"},
	//OBSOLETE_ER_NDB_BINLOG_FAILED_CREATE_TABLE_EVENT_OPERATIONS : {10494,[]string{"HY000"},"NDB Binlog: FAILED CREATE TABLE event operations. Event: %s"},
	//OBSOLETE_ER_NDB_BINLOG_RENAME_EVENT : {10495,[]string{"HY000"},"NDB Binlog: RENAME Event: %s"},
	//OBSOLETE_ER_NDB_BINLOG_FAILED_CREATE_DURING_RENAME : {0000,[]string{""},"NDB Binlog: FAILED create event operations during RENAME. Event %s"},
	//OBSOLETE_ER_NDB_UNEXPECTED_RENAME_TYPE : {10497,[]string{"HY000"},"Unexpected rename case detected, sql_command: %d"},
	//OBSOLETE_ER_NDB_ERROR_IN_GET_AUTO_INCREMENT : {10498,[]string{"HY000"},"Error %lu in ::get_auto_increment(): %s"},
	//OBSOLETE_ER_NDB_CREATING_SHARE_IN_OPEN : {10499,[]string{"HY000"},"Calling ndbcluster_create_binlog_setup(%s) in ::open"},
	//OBSOLETE_ER_NDB_TABLE_OPENED_READ_ONLY : {10500,[]string{"HY000"},"table '%s' opened read only"},
	//OBSOLETE_ER_NDB_INITIALIZE_GIVEN_CLUSTER_PLUGIN_DISABLED : {10501,[]string{"HY000"},"NDB: '--initialize' -> ndbcluster plugin disabled"},
	//OBSOLETE_ER_NDB_BINLOG_FORMAT_CHANGED_FROM_STMT_TO_MIXED : {10502,[]string{"HY000"},"NDB: Changed global value of binlog_format from STATEMENT to MIXED"},
	//OBSOLETE_ER_NDB_TRAILING_SHARE_RELEASED_BY_CLOSE_CACHED_TABLES : {10503,[]string{"HY000"},"NDB_SHARE: trailing share %s, released by close_cached_tables"},
	//OBSOLETE_ER_NDB_SHARE_ALREADY_EXISTS : {10504,[]string{"HY000"},"NDB_SHARE: %s already exists use_count=%d. Moving away for safety, but possible memleak."},
	//OBSOLETE_ER_NDB_HANDLE_TRAILING_SHARE_INFO : {10505,[]string{"HY000"},"handle_trailing_share: %s use_count: %u"},
	//OBSOLETE_ER_NDB_CLUSTER_GET_SHARE_INFO : {10506,[]string{"HY000"},"ndbcluster_get_share: %s use_count: %u"},
	//OBSOLETE_ER_NDB_CLUSTER_REAL_FREE_SHARE_INFO : {10507,[]string{"HY000"},"ndbcluster_real_free_share: %s use_count: %u"},
	//OBSOLETE_ER_NDB_CLUSTER_REAL_FREE_SHARE_DROP_FAILED : {10508,[]string{"HY000"},"ndbcluster_real_free_share: %s, still open - ignored 'free' (leaked?)"},
	//OBSOLETE_ER_NDB_CLUSTER_FREE_SHARE_INFO : {10509,[]string{"HY000"},"ndbcluster_free_share: %s use_count: %u"},
	//OBSOLETE_ER_NDB_CLUSTER_MARK_SHARE_DROPPED_INFO : {10510,[]string{"HY000"},"ndbcluster_mark_share_dropped: %s use_count: %u"},
	//OBSOLETE_ER_NDB_CLUSTER_MARK_SHARE_DROPPED_DESTROYING_SHARE : {10511,[]string{"HY000"},"ndbcluster_mark_share_dropped: destroys share %s"},
	//OBSOLETE_ER_NDB_CLUSTER_OOM_THD_NDB : {10512,[]string{"HY000"},"Could not allocate Thd_ndb object"},
	//OBSOLETE_ER_NDB_BINLOG_NDB_TABLES_INITIALLY_READ_ONLY : {10513,[]string{"HY000"},"NDB Binlog: Ndb tables initially read only."},
	//OBSOLETE_ER_NDB_UTIL_THREAD_OOM : {10514,[]string{"HY000"},"ndb util thread: malloc failure, query cache not maintained properly"},
	//OBSOLETE_ER_NDB_ILLEGAL_VALUE_FOR_NDB_RECV_THREAD_CPU_MASK : {10515,[]string{"HY000"},"Trying to set ndb_recv_thread_cpu_mask to illegal value = %s, ignored"},
	//OBSOLETE_ER_NDB_TOO_MANY_CPUS_IN_NDB_RECV_THREAD_CPU_MASK : {10516,[]string{"HY000"},"Trying to set too many CPU's in ndb_recv_thread_cpu_mask, ignored this variable, erroneus value = %s"},
	ER_DBUG_CHECK_SHARES_OPEN : {10517,[]string{"HY000"},"dbug_check_shares open:"},
	ER_DBUG_CHECK_SHARES_INFO : {10518,[]string{"HY000"},"  %s.%s: state: %s(%u) use_count: %u"},
	ER_DBUG_CHECK_SHARES_DROPPED : {10519,[]string{"HY000"},"dbug_check_shares dropped:"},
	ER_INVALID_OR_OLD_TABLE_OR_DB_NAME : {10520,[]string{"HY000"},"Invalid (old?) table or database name '%s'"},
	ER_TC_RECOVERING_AFTER_CRASH_USING : {10521,[]string{"HY000"},"Recovering after a crash using %s"},
	ER_TC_CANT_AUTO_RECOVER_WITH_TC_HEURISTIC_RECOVER : {10522,[]string{"HY000"},"Cannot perform automatic crash recovery when --tc-heuristic-recover is used"},
	ER_TC_BAD_MAGIC_IN_TC_LOG : {10523,[]string{"HY000"},"Bad magic header in tc log"},
	ER_TC_NEED_N_SE_SUPPORTING_2PC_FOR_RECOVERY : {10524,[]string{"HY000"},"Recovery failed! You must enable exactly %d storage engines that support two-phase commit protocol"},
	ER_TC_RECOVERY_FAILED_THESE_ARE_YOUR_OPTIONS : {10525,[]string{"HY000"},"Crash recovery failed. Either correct the problem (if it's, for example, out of memory error) and restart, or delete tc log and start mysqld with --tc-heuristic-recover={commit|rollback}"},
	ER_TC_HEURISTIC_RECOVERY_MODE : {10526,[]string{"HY000"},"Heuristic crash recovery mode"},
	ER_TC_HEURISTIC_RECOVERY_FAILED : {10527,[]string{"HY000"},"Heuristic crash recovery failed"},
	ER_TC_RESTART_WITHOUT_TC_HEURISTIC_RECOVER : {10528,[]string{"HY000"},"Please restart mysqld without --tc-heuristic-recover"},
	ER_RPL_SLAVE_FAILED_TO_CREATE_OR_RECOVER_INFO_REPOSITORIES : {10529,[]string{"HY000"},"Failed to create or recover replication info repositories."},
	ER_RPL_SLAVE_AUTO_POSITION_IS_1_AND_GTID_MODE_IS_OFF : {10530,[]string{"HY000"},"Detected misconfiguration: replication channel '%.192s' was configured with AUTO_POSITION = 1, but the server was started with --gtid-mode=off. Either reconfigure replication using CHANGE MASTER TO MASTER_AUTO_POSITION = 0 FOR CHANNEL '%.192s', or change GTID_MODE to some value other than OFF, before starting the slave receiver thread."},
	ER_RPL_SLAVE_CANT_START_SLAVE_FOR_CHANNEL : {10531,[]string{"HY000"},"Slave: Could not start slave for channel '%s'. operation discontinued"},
	ER_RPL_SLAVE_CANT_STOP_SLAVE_FOR_CHANNEL : {10532,[]string{"HY000"},"Slave: Could not stop slave for channel '%s' operation discontinued"},
	ER_RPL_RECOVERY_NO_ROTATE_EVENT_FROM_MASTER : {10533,[]string{"HY000"},"Error during --relay-log-recovery: Could not locate rotate event from the master."},
	ER_RPL_RECOVERY_ERROR_READ_RELAY_LOG : {10534,[]string{"HY000"},"Error during --relay-log-recovery: Error reading events from relay log: %d"},
	//OBSOLETE_ER_RPL_RECOVERY_ERROR_FREEING_IO_CACHE : {10535,[]string{"HY000"},"Error during --relay-log-recovery: Error while freeing IO_CACHE object"},
	ER_RPL_RECOVERY_SKIPPED_GROUP_REPLICATION_CHANNEL : {10536,[]string{"HY000"},"Relay log recovery skipped for group replication channel."},
	ER_RPL_RECOVERY_ERROR : {10537,[]string{"HY000"},"Error during --relay-log-recovery: %s"},
	ER_RPL_RECOVERY_IO_ERROR_READING_RELAY_LOG_INDEX : {10538,[]string{"HY000"},"Error during --relay-log-recovery: Could not read relay log index file due to an IO error."},
	ER_RPL_RECOVERY_FILE_MASTER_POS_INFO : {10539,[]string{"HY000"},"Recovery from master pos %ld and file %s%s. Previous relay log pos and relay log file had been set to %lld, %s respectively."},
	ER_RPL_RECOVERY_REPLICATE_SAME_SERVER_ID_REQUIRES_POSITION : {10540,[]string{"HY000"},"Error during --relay-log-recovery: replicate_same_server_id is in use and sql thread's positions are not initialized, hence relay log recovery cannot happen."},
	ER_RPL_MTS_RECOVERY_STARTING_COORDINATOR : {10541,[]string{"HY000"},"MTS recovery: starting coordinator thread to fill MTS gaps."},
	ER_RPL_MTS_RECOVERY_FAILED_TO_START_COORDINATOR : {10542,[]string{"HY000"},"MTS recovery: failed to start the coordinator thread. Check the error log for additional details."},
	ER_RPL_MTS_AUTOMATIC_RECOVERY_FAILED : {10543,[]string{"HY000"},"MTS recovery: automatic recovery failed. Either the slave server had stopped due to an error during an earlier session or relay logs are corrupted.Fix the cause of the slave side error and restart the slave server or consider using RESET SLAVE."},
	ER_RPL_MTS_RECOVERY_CANT_OPEN_RELAY_LOG : {10544,[]string{"HY000"},"Failed to open the relay log '%s' (relay_log_pos %s)."},
	ER_RPL_MTS_RECOVERY_SUCCESSFUL : {10545,[]string{"HY000"},"MTS recovery: completed successfully."},
	ER_RPL_SERVER_ID_MISSING : {10546,[]string{"HY000"},"Server id not set, will not start slave%s"},
	ER_RPL_CANT_CREATE_SLAVE_THREAD : {10547,[]string{"HY000"},"Can't create slave thread%s."},
	ER_RPL_SLAVE_IO_THREAD_WAS_KILLED : {10548,[]string{"HY000"},"The slave IO thread%s was killed while executing initialization query '%s'"},
	ER_RPL_SLAVE_MASTER_UUID_HAS_CHANGED : {10549,[]string{"HY000"},"The master's UUID has changed, although this should not happen unless you have changed it manually. The old UUID was %s."},
	ER_RPL_SLAVE_USES_CHECKSUM_AND_MASTER_PRE_50 : {10550,[]string{"HY000"},"Found a master with MySQL server version older than 5.0. With checksums enabled on the slave, replication might not work correctly. To ensure correct replication, restart the slave server with --slave_sql_verify_checksum=0."},
	ER_RPL_SLAVE_SECONDS_BEHIND_MASTER_DUBIOUS : {10551,[]string{"HY000"},"\"SELECT UNIX_TIMESTAMP()\" failed on master, do not trust column Seconds_Behind_Master of SHOW SLAVE STATUS. Error: %s (%d)"},
	ER_RPL_SLAVE_CANT_FLUSH_MASTER_INFO_FILE : {10552,[]string{"HY000"},"Failed to flush master info file."},
	ER_RPL_SLAVE_REPORT_HOST_TOO_LONG : {10553,[]string{"HY000"},"The length of report_host is %zu. It is larger than the max length(%d), so this slave cannot be registered to the master%s."},
	ER_RPL_SLAVE_REPORT_USER_TOO_LONG : {10554,[]string{"HY000"},"The length of report_user is %zu. It is larger than the max length(%d), so this slave cannot be registered to the master%s."},
	ER_RPL_SLAVE_REPORT_PASSWORD_TOO_LONG : {10555,[]string{"HY000"},"The length of report_password is %zu. It is larger than the max length(%d), so this slave cannot be registered to the master%s."},
	ER_RPL_SLAVE_ERROR_RETRYING : {10556,[]string{"HY000"},"Error on %s: %d  %s, will retry in %d secs"},
	ER_RPL_SLAVE_ERROR_READING_FROM_SERVER : {10557,[]string{"HY000"},"Error reading packet from server%s: %s (server_errno=%d)"},
	ER_RPL_SLAVE_DUMP_THREAD_KILLED_BY_MASTER : {10558,[]string{"HY000"},"Slave%s: received end packet from server due to dump thread being killed on master. Dump threads are killed for example during master shutdown, explicitly by a user, or when the master receives a binlog send request from a duplicate server UUID <%s> : Error %s"},
	ER_RPL_MTS_STATISTICS : {10559,[]string{"HY000"},"Multi-threaded slave statistics%s: seconds elapsed = %lu; events assigned = %llu; worker queues filled over overrun level = %lu; waited due a Worker queue full = %lu; waited due the total size = %lu; waited at clock conflicts = %llu waited (count) when Workers occupied = %lu waited when Workers occupied = %llu"},
	ER_RPL_MTS_RECOVERY_COMPLETE : {10560,[]string{"HY000"},"Slave%s: MTS Recovery has completed at relay log %s, position %llu master log %s, position %llu."},
	ER_RPL_SLAVE_CANT_INIT_RELAY_LOG_POSITION : {10561,[]string{"HY000"},"Error initializing relay log position%s: %s"},
	ER_RPL_SLAVE_CONNECTED_TO_MASTER_REPLICATION_STARTED : {10562,[]string{"HY000"},"Slave I/O thread%s: connected to master '%s@%s:%d',replication started in log '%s' at position %s"},
	ER_RPL_SLAVE_IO_THREAD_KILLED : {10563,[]string{"HY000"},"Slave I/O thread%s killed while connecting to master"},
	ER_RPL_SLAVE_IO_THREAD_CANT_REGISTER_ON_MASTER : {10564,[]string{"HY000"},"Slave I/O thread couldn't register on master"},
	ER_RPL_SLAVE_FORCING_TO_RECONNECT_IO_THREAD : {10565,[]string{"HY000"},"Forcing to reconnect slave I/O thread%s"},
	ER_RPL_SLAVE_ERROR_REQUESTING_BINLOG_DUMP : {10566,[]string{"HY000"},"Failed on request_dump()%s"},
	ER_RPL_LOG_ENTRY_EXCEEDS_SLAVE_MAX_ALLOWED_PACKET : {10567,[]string{"HY000"},"Log entry on master is longer than slave_max_allowed_packet (%lu) on slave. If the entry is correct, restart the server with a higher value of slave_max_allowed_packet"},
	ER_RPL_SLAVE_STOPPING_AS_MASTER_OOM : {10568,[]string{"HY000"},"Stopping slave I/O thread due to out-of-memory error from master"},
	ER_RPL_SLAVE_IO_THREAD_ABORTED_WAITING_FOR_RELAY_LOG_SPACE : {10569,[]string{"HY000"},"Slave I/O thread aborted while waiting for relay log space"},
	ER_RPL_SLAVE_IO_THREAD_EXITING : {10570,[]string{"HY000"},"Slave I/O thread exiting%s, read up to log '%s', position %s"},
	ER_RPL_SLAVE_CANT_INITIALIZE_SLAVE_WORKER : {10571,[]string{"HY000"},"Failed during slave worker initialization%s"},
	ER_RPL_MTS_GROUP_RECOVERY_RELAY_LOG_INFO_FOR_WORKER : {10572,[]string{"HY000"},"Slave: MTS group recovery relay log info based on Worker-Id %lu, group_relay_log_name %s, group_relay_log_pos %llu group_master_log_name %s, group_master_log_pos %llu"},
	ER_RPL_ERROR_LOOKING_FOR_LOG : {10573,[]string{"HY000"},"Error looking for %s."},
	ER_RPL_MTS_GROUP_RECOVERY_RELAY_LOG_INFO : {10574,[]string{"HY000"},"Slave: MTS group recovery relay log info group_master_log_name %s, event_master_log_pos %llu."},
	ER_RPL_CANT_FIND_FOLLOWUP_FILE : {10575,[]string{"HY000"},"Error looking for file after %s."},
	ER_RPL_MTS_CHECKPOINT_PERIOD_DIFFERS_FROM_CNT : {10576,[]string{"HY000"},"This an error cnt != mts_checkpoint_period"},
	ER_RPL_SLAVE_WORKER_THREAD_CREATION_FAILED : {10577,[]string{"HY000"},"Failed during slave worker thread creation%s"},
	ER_RPL_SLAVE_WORKER_THREAD_CREATION_FAILED_WITH_ERRNO : {10578,[]string{"HY000"},"Failed during slave worker thread creation%s (errno= %d)"},
	ER_RPL_SLAVE_FAILED_TO_INIT_PARTITIONS_HASH : {10579,[]string{"HY000"},"Failed to init partitions hash"},
	//OBSOLETE_ER_RPL_SLAVE_NDB_TABLES_NOT_AVAILABLE : {10580,[]string{"HY000"},"Slave SQL thread : NDB : Tables not available after %lu seconds. Consider increasing --ndb-wait-setup value"},
	ER_RPL_SLAVE_SQL_THREAD_STARTING : {10581,[]string{"HY000"},"Slave SQL thread%s initialized, starting replication in log '%s' at position %s, relay log '%s' position: %s"},
	ER_RPL_SLAVE_SKIP_COUNTER_EXECUTED : {10582,[]string{"HY000"},"'SQL_SLAVE_SKIP_COUNTER=%ld' executed at relay_log_file='%s', relay_log_pos='%ld', master_log_name='%s', master_log_pos='%ld' and new position at relay_log_file='%s', relay_log_pos='%ld', master_log_name='%s', master_log_pos='%ld' "},
	ER_RPL_SLAVE_ADDITIONAL_ERROR_INFO_FROM_DA : {10583,[]string{"HY000"},"Slave (additional info): %s Error_code: MY-%06d"},
	ER_RPL_SLAVE_ERROR_INFO_FROM_DA : {10584,[]string{"HY000"},"Slave: %s Error_code: MY-%06d"},
	ER_RPL_SLAVE_ERROR_LOADING_USER_DEFINED_LIBRARY : {10585,[]string{"HY000"},"Error loading user-defined library, slave SQL thread aborted. Install the missing library, and restart the slave SQL thread with \"SLAVE START\". We stopped at log '%s' position %s"},
	ER_RPL_SLAVE_ERROR_RUNNING_QUERY : {10586,[]string{"HY000"},"Error running query, slave SQL thread aborted. Fix the problem, and restart the slave SQL thread with \"SLAVE START\". We stopped at log '%s' position %s"},
	ER_RPL_SLAVE_SQL_THREAD_EXITING : {10587,[]string{"HY000"},"Slave SQL thread%s exiting, replication stopped in log '%s' at position %s"},
	ER_RPL_SLAVE_READ_INVALID_EVENT_FROM_MASTER : {10588,[]string{"HY000"},"Read invalid event from master: '%s', master could be corrupt but a more likely cause of this is a bug"},
	ER_RPL_SLAVE_QUEUE_EVENT_FAILED_INVALID_CONFIGURATION : {10589,[]string{"HY000"},"The queue event failed for channel '%s' as its configuration is invalid."},
	ER_RPL_SLAVE_IO_THREAD_DETECTED_UNEXPECTED_EVENT_SEQUENCE : {10590,[]string{"HY000"},"An unexpected event sequence was detected by the IO thread while queuing the event received from master '%s' binary log file, at position %llu."},
	ER_RPL_SLAVE_CANT_USE_CHARSET : {10591,[]string{"HY000"},"'%s' can not be used as client character set. '%s' will be used as default client character set while connecting to master."},
	ER_RPL_SLAVE_CONNECTED_TO_MASTER_REPLICATION_RESUMED : {10592,[]string{"HY000"},"Slave%s: connected to master '%s@%s:%d',replication resumed in log '%s' at position %s"},
	ER_RPL_SLAVE_NEXT_LOG_IS_ACTIVE : {10593,[]string{"HY000"},"next log '%s' is active"},
	ER_RPL_SLAVE_NEXT_LOG_IS_INACTIVE : {10594,[]string{"HY000"},"next log '%s' is not active"},
	ER_RPL_SLAVE_SQL_THREAD_IO_ERROR_READING_EVENT : {10595,[]string{"HY000"},"Slave SQL thread%s: I/O error reading event (errno: %d  cur_log->error: %d)"},
	ER_RPL_SLAVE_ERROR_READING_RELAY_LOG_EVENTS : {10596,[]string{"HY000"},"Error reading relay log event%s: %s"},
	ER_SLAVE_CHANGE_MASTER_TO_EXECUTED : {10597,[]string{"HY000"},"'CHANGE MASTER TO%s executed'. Previous state master_host='%s', master_port= %u, master_log_file='%s', master_log_pos= %ld, master_bind='%s'. New state master_host='%s', master_port= %u, master_log_file='%s', master_log_pos= %ld, master_bind='%s'."},
	ER_RPL_SLAVE_NEW_MASTER_INFO_NEEDS_REPOS_TYPE_OTHER_THAN_FILE : {10598,[]string{"HY000"},"Slave: Cannot create new master info structure when repositories are of type FILE. Convert slave repositories to TABLE to replicate from multiple sources."},
	ER_RPL_FAILED_TO_STAT_LOG_IN_INDEX : {10599,[]string{"HY000"},"log %s listed in the index, but failed to stat."},
	ER_RPL_LOG_NOT_FOUND_WHILE_COUNTING_RELAY_LOG_SPACE : {10600,[]string{"HY000"},"Could not find first log while counting relay log space."},
	ER_SLAVE_CANT_USE_TEMPDIR : {10601,[]string{"HY000"},"Unable to use slave's temporary directory '%s'."},
	ER_RPL_RELAY_LOG_NEEDS_FILE_NOT_DIRECTORY : {10602,[]string{"HY000"},"Path '%s' is a directory name, please specify a file name for --relay-log option."},
	ER_RPL_RELAY_LOG_INDEX_NEEDS_FILE_NOT_DIRECTORY : {10603,[]string{"HY000"},"Path '%s' is a directory name, please specify a file name for --relay-log-index option."},
	ER_RPL_PLEASE_USE_OPTION_RELAY_LOG : {10604,[]string{"HY000"},"Neither --relay-log nor --relay-log-index were used; so replication may break when this MySQL server acts as a slave and has his hostname changed!! Please use '--relay-log=%s' to avoid this problem."},
	ER_RPL_OPEN_INDEX_FILE_FAILED : {10605,[]string{"HY000"},"Failed in open_index_file() called from Relay_log_info::rli_init_info()."},
	ER_RPL_CANT_INITIALIZE_GTID_SETS_IN_RLI_INIT_INFO : {10606,[]string{"HY000"},"Failed in init_gtid_sets() called from Relay_log_info::rli_init_info()."},
	ER_RPL_CANT_OPEN_LOG_IN_RLI_INIT_INFO : {10607,[]string{"HY000"},"Failed in open_log() called from Relay_log_info::rli_init_info()."},
	ER_RPL_ERROR_WRITING_RELAY_LOG_CONFIGURATION : {10608,[]string{"HY000"},"Error writing relay log configuration."},
	//OBSOLETE_ER_NDB_OOM_GET_NDB_BLOBS_VALUE : {10609,[]string{"HY000"},"get_ndb_blobs_value: my_malloc(%u) failed"},
	//OBSOLETE_ER_NDB_THREAD_TIMED_OUT : {10610,[]string{"HY000"},"NDB: Thread id %u timed out (30s) waiting for epoch %u/%u to be handled.  Progress : %u/%u -> %u/%u."},
	//OBSOLETE_ER_NDB_TABLE_IS_NOT_DISTRIBUTED : {10611,[]string{"HY000"},"NDB: Inconsistency detected in distributed privilege tables. Table '%s.%s' is not distributed"},
	//OBSOLETE_ER_NDB_CREATING_TABLE : {10612,[]string{"HY000"},"NDB: Creating %s.%s"},
	//OBSOLETE_ER_NDB_FLUSHING_TABLE_INFO : {10613,[]string{"HY000"},"NDB: Flushing %s.%s"},
	//OBSOLETE_ER_NDB_CLEANING_STRAY_TABLES : {10614,[]string{"HY000"},"NDB: Cleaning stray tables from database '%s'"},
	//OBSOLETE_ER_NDB_DISCOVERED_MISSING_DB : {10615,[]string{"HY000"},"NDB: Discovered missing database '%s'"},
	//OBSOLETE_ER_NDB_DISCOVERED_REMAINING_DB : {10616,[]string{"HY000"},"NDB: Discovered remaining database '%s'"},
	//OBSOLETE_ER_NDB_CLUSTER_FIND_ALL_DBS_RETRY : {10617,[]string{"HY000"},"NDB: ndbcluster_find_all_databases retry: %u - %s"},
	//OBSOLETE_ER_NDB_CLUSTER_FIND_ALL_DBS_FAIL : {10618,[]string{"HY000"},"NDB: ndbcluster_find_all_databases fail: %u - %s"},
	//OBSOLETE_ER_NDB_SKIPPING_SETUP_TABLE : {10619,[]string{"HY000"},"NDB: skipping setup table %s.%s, in state %d"},
	//OBSOLETE_ER_NDB_FAILED_TO_SET_UP_TABLE : {10620,[]string{"HY000"},"NDB: failed to setup table %s.%s, error: %d, %s"},
	//OBSOLETE_ER_NDB_MISSING_FRM_DISCOVERING : {10621,[]string{"HY000"},"NDB: missing frm for %s.%s, discovering..."},
	//OBSOLETE_ER_NDB_MISMATCH_IN_FRM_DISCOVERING : {10622,[]string{"HY000"},"NDB: mismatch in frm for %s.%s, discovering..."},
	//OBSOLETE_ER_NDB_BINLOG_CLEANING_UP_SETUP_LEFTOVERS : {10623,[]string{"HY000"},"ndb_binlog_setup: Clean up leftovers"},
	//OBSOLETE_ER_NDB_WAITING_INFO : {10624,[]string{"HY000"},"NDB %s: waiting max %u sec for %s %s.  epochs: (%u/%u,%u/%u,%u/%u)  injector proc_info: %s"},
	//OBSOLETE_ER_NDB_WAITING_INFO_WITH_MAP : {10625,[]string{"HY000"},"NDB %s: waiting max %u sec for %s %s.  epochs: (%u/%u,%u/%u,%u/%u)  injector proc_info: %s map: %x%08x"},
	//OBSOLETE_ER_NDB_TIMEOUT_WHILE_DISTRIBUTING : {10626,[]string{"HY000"},"NDB %s: distributing %s timed out. Ignoring..."},
	//OBSOLETE_ER_NDB_NOT_WAITING_FOR_DISTRIBUTING : {10627,[]string{"HY000"},"NDB %s: not waiting for distributing %s"},
	//OBSOLETE_ER_NDB_DISTRIBUTED_INFO : {10628,[]string{"HY000"},"NDB: distributed %s.%s(%u/%u) type: %s(%u) query: \'%s\' to %x%08x"},
	//OBSOLETE_ER_NDB_DISTRIBUTION_COMPLETE : {10629,[]string{"HY000"},"NDB: distribution of %s.%s(%u/%u) type: %s(%u) query: \'%s\' - complete!"},
	//OBSOLETE_ER_NDB_SCHEMA_DISTRIBUTION_FAILED : {10630,[]string{"HY000"},"NDB Schema dist: Data node: %d failed, subscriber bitmask %x%08x"},
	//OBSOLETE_ER_NDB_SCHEMA_DISTRIBUTION_REPORTS_SUBSCRIBE : {10631,[]string{"HY000"},"NDB Schema dist: Data node: %d reports subscribe from node %d, subscriber bitmask %x%08x"},
	//OBSOLETE_ER_NDB_SCHEMA_DISTRIBUTION_REPORTS_UNSUBSCRIBE : {10632,[]string{"HY000"},"NDB Schema dist: Data node: %d reports unsubscribe from node %d, subscriber bitmask %x%08x"},
	//OBSOLETE_ER_NDB_BINLOG_CANT_DISCOVER_TABLE_FROM_SCHEMA_EVENT : {10633,[]string{"HY000"},"NDB Binlog: Could not discover table '%s.%s' from binlog schema event '%s' from node %d. my_errno: %d"},
	//OBSOLETE_ER_NDB_BINLOG_SIGNALLING_UNKNOWN_VALUE : {10634,[]string{"HY000"},"NDB: unknown value for binlog signalling 0x%X, %s not logged"},
	//OBSOLETE_ER_NDB_BINLOG_REPLY_TO : {10635,[]string{"HY000"},"NDB: reply to %s.%s(%u/%u) from %s to %x%08x"},
	//OBSOLETE_ER_NDB_BINLOG_CANT_RELEASE_SLOCK : {10636,[]string{"HY000"},"NDB: Could not release slock on '%s.%s', Error code: %d Message: %s"},
	//OBSOLETE_ER_NDB_CANT_FIND_TABLE : {10637,[]string{"HY000"},"NDB schema: Could not find table '%s.%s' in NDB"},
	//OBSOLETE_ER_NDB_DISCARDING_EVENT_NO_OBJ : {10638,[]string{"HY000"},"NDB: Discarding event...no obj: %s (%u/%u)"},
	//OBSOLETE_ER_NDB_DISCARDING_EVENT_ID_VERSION_MISMATCH : {10639,[]string{"HY000"},"NDB: Discarding event...key: %s non matching id/version [%u/%u] != [%u/%u]"},
	//OBSOLETE_ER_NDB_CLEAR_SLOCK_INFO : {10640,[]string{"HY000"},"NDB: CLEAR_SLOCK key: %s(%u/%u) %x%08x, from %s to %x%08x"},
	//OBSOLETE_ER_NDB_BINLOG_SKIPPING_LOCAL_TABLE : {10641,[]string{"HY000"},"NDB Binlog: Skipping locally defined table '%s.%s' from binlog schema event '%s' from node %d."},
	//OBSOLETE_ER_NDB_BINLOG_ONLINE_ALTER_RENAME : {10642,[]string{"HY000"},"NDB Binlog: handling online alter/rename"},
	//OBSOLETE_ER_NDB_BINLOG_CANT_REOPEN_SHADOW_TABLE : {10643,[]string{"HY000"},"NDB Binlog: Failed to re-open shadow table %s.%s"},
	//OBSOLETE_ER_NDB_BINLOG_ONLINE_ALTER_RENAME_COMPLETE : {10644,[]string{"HY000"},"NDB Binlog: handling online alter/rename done"},
	//OBSOLETE_ER_NDB_BINLOG_SKIPPING_DROP_OF_LOCAL_TABLE : {10645,[]string{"HY000"},"NDB Binlog: Skipping drop of locally defined table '%s.%s' from binlog schema event '%s' from node %d. "},
	//OBSOLETE_ER_NDB_BINLOG_SKIPPING_RENAME_OF_LOCAL_TABLE : {10646,[]string{"HY000"},"NDB Binlog: Skipping renaming locally defined table '%s.%s' from binlog schema event '%s' from node %d. "},
	//OBSOLETE_ER_NDB_BINLOG_SKIPPING_DROP_OF_TABLES : {0000,[]string{""},"NDB Binlog: Skipping drop database '%s' since it contained local tables binlog schema event '%s' from node %d. "},
	//OBSOLETE_ER_NDB_BINLOG_GOT_DIST_PRIV_EVENT_FLUSHING_PRIVILEGES : {10648,[]string{"HY000"},"Got dist_priv event: %s, flushing privileges"},
	//OBSOLETE_ER_NDB_BINLOG_GOT_SCHEMA_EVENT : {10649,[]string{"HY000"},"NDB: got schema event on %s.%s(%u/%u) query: '%s' type: %s(%d) node: %u slock: %x%08x"},
	//OBSOLETE_ER_NDB_BINLOG_SKIPPING_OLD_SCHEMA_OPERATION : {10650,[]string{"HY000"},"NDB schema: Skipping old schema operation(RENAME_TABLE_NEW) on %s.%s"},
	//OBSOLETE_ER_NDB_CLUSTER_FAILURE : {10651,[]string{"HY000"},"NDB Schema dist: cluster failure at epoch %u/%u."},
	//OBSOLETE_ER_NDB_TABLES_INITIALLY_READ_ONLY_ON_RECONNECT : {10652,[]string{"HY000"},"NDB Binlog: ndb tables initially read only on reconnect."},
	//OBSOLETE_ER_NDB_IGNORING_UNKNOWN_EVENT : {10653,[]string{"HY000"},"NDB Schema dist: unknown event %u, ignoring..."},
	//OBSOLETE_ER_NDB_BINLOG_OPENING_INDEX : {10654,[]string{"HY000"},"NDB Binlog: Opening ndb_binlog_index: %d, '%s'"},
	//OBSOLETE_ER_NDB_BINLOG_CANT_LOCK_NDB_BINLOG_INDEX : {10655,[]string{"HY000"},"NDB Binlog: Unable to lock table ndb_binlog_index"},
	//OBSOLETE_ER_NDB_BINLOG_INJECTING_RANDOM_WRITE_FAILURE : {10656,[]string{"HY000"},"NDB Binlog: Injecting random write failure"},
	//OBSOLETE_ER_NDB_BINLOG_CANT_WRITE_TO_NDB_BINLOG_INDEX : {10657,[]string{"HY000"},"NDB Binlog: Failed writing to ndb_binlog_index for epoch %u/%u  orig_server_id %u orig_epoch %u/%u with error %d."},
	//OBSOLETE_ER_NDB_BINLOG_WRITING_TO_NDB_BINLOG_INDEX : {10658,[]string{"HY000"},"NDB Binlog: Writing row (%s) to ndb_binlog_index - %s"},
	//OBSOLETE_ER_NDB_BINLOG_CANT_COMMIT_TO_NDB_BINLOG_INDEX : {10659,[]string{"HY000"},"NDB Binlog: Failed committing transaction to ndb_binlog_index with error %d."},
	//OBSOLETE_ER_NDB_BINLOG_WRITE_INDEX_FAILED_AFTER_KILL : {0000,[]string{""},"NDB Binlog: Failed writing to ndb_binlog_index table while retrying after kill during shutdown"},
	//OBSOLETE_ER_NDB_BINLOG_USING_SERVER_ID_0_SLAVES_WILL_NOT : {10661,[]string{"HY000"},"NDB: server id set to zero - changes logged to bin log with server id zero will be logged with another server id by slave mysqlds"},
	//OBSOLETE_ER_NDB_SERVER_ID_RESERVED_OR_TOO_LARGE : {10662,[]string{"HY000"},"NDB: server id provided is too large to be represented in opt_server_id_bits or is reserved"},
	//OBSOLETE_ER_NDB_BINLOG_REQUIRES_V2_ROW_EVENTS : {0000,[]string{""},"NDB: --ndb-log-transaction-id requires v2 Binlog row events but server is using v1."},
	//OBSOLETE_ER_NDB_BINLOG_STATUS_FORCING_FULL_USE_WRITE : {0000,[]string{""},"NDB: ndb-log-apply-status forcing %s.%s to FULL USE_WRITE"},
	//OBSOLETE_ER_NDB_BINLOG_GENERIC_MESSAGE : {10665,[]string{"HY000"},"NDB Binlog: %s"},
	//OBSOLETE_ER_NDB_CONFLICT_GENERIC_MESSAGE : {10666,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_NDB_TRANS_DEPENDENCY_TRACKER_ERROR : {10667,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_NDB_CONFLICT_FN_PARSE_ERROR : {10668,[]string{"HY000"},"NDB Slave: Table %s.%s : Parse error on conflict fn : %s"},
	//OBSOLETE_ER_NDB_CONFLICT_FN_SETUP_ERROR : {10669,[]string{"HY000"},"NDB Slave: Table %s.%s : %s"},
	//OBSOLETE_ER_NDB_BINLOG_FAILED_TO_GET_TABLE : {10670,[]string{"HY000"},"NDB Binlog: Failed to get table %s from ndb: %s, %d"},
	//OBSOLETE_ER_NDB_BINLOG_NOT_LOGGING : {10671,[]string{"HY000"},"NDB Binlog: NOT logging %s"},
	//OBSOLETE_ER_NDB_BINLOG_CREATE_TABLE_EVENT_FAILED : {10672,[]string{"HY000"},"NDB Binlog: FAILED CREATE (DISCOVER) TABLE Event: %s"},
	//OBSOLETE_ER_NDB_BINLOG_CREATE_TABLE_EVENT_INFO : {10673,[]string{"HY000"},"NDB Binlog: CREATE (DISCOVER) TABLE Event: %s"},
	//OBSOLETE_ER_NDB_BINLOG_DISCOVER_TABLE_EVENT_INFO : {10674,[]string{"HY000"},"NDB Binlog: DISCOVER TABLE Event: %s"},
	//OBSOLETE_ER_NDB_BINLOG_BLOB_REQUIRES_PK : {10675,[]string{"HY000"},"NDB Binlog: logging of table %s with BLOB attribute and no PK is not supported"},
	//OBSOLETE_ER_NDB_BINLOG_CANT_CREATE_EVENT_IN_DB : {10676,[]string{"HY000"},"NDB Binlog: Unable to create event in database. Event: %s  Error Code: %d  Message: %s"},
	//OBSOLETE_ER_NDB_BINLOG_CANT_CREATE_EVENT_IN_DB_AND_CANT_DROP : {10677,[]string{"HY000"},"NDB Binlog: Unable to create event in database.  Attempt to correct with drop failed. Event: %s Error Code: %d Message: %s"},
	//OBSOLETE_ER_NDB_BINLOG_CANT_CREATE_EVENT_IN_DB_DROPPED : {10678,[]string{"HY000"},"NDB Binlog: Unable to create event in database.  Attempt to correct with drop ok, but create failed. Event: %s Error Code: %d Message: %s"},
	//OBSOLETE_ER_NDB_BINLOG_DISCOVER_REUSING_OLD_EVENT_OPS : {10679,[]string{"HY000"},"NDB Binlog: discover reusing old ev op"},
	//OBSOLETE_ER_NDB_BINLOG_CREATING_NDBEVENTOPERATION_FAILED : {10680,[]string{"HY000"},"NDB Binlog: Creating NdbEventOperation failed for %s"},
	//OBSOLETE_ER_NDB_BINLOG_CANT_CREATE_BLOB : {10681,[]string{"HY000"},"NDB Binlog: Creating NdbEventOperation blob field %u handles failed (code=%d) for %s"},
	//OBSOLETE_ER_NDB_BINLOG_NDBEVENT_EXECUTE_FAILED : {10682,[]string{"HY000"},"NDB Binlog: ndbevent->execute failed for %s; %d %s"},
	//OBSOLETE_ER_NDB_CREATE_EVENT_OPS_LOGGING_INFO : {10683,[]string{"HY000"},"NDB Binlog: logging %s (%s,%s)"},
	//OBSOLETE_ER_NDB_BINLOG_CANT_DROP_EVENT_FROM_DB : {10684,[]string{"HY000"},"NDB Binlog: Unable to drop event in database. Event: %s Error Code: %d Message: %s"},
	//OBSOLETE_ER_NDB_TIMED_OUT_IN_DROP_TABLE : {10685,[]string{"HY000"},"NDB %s: %s timed out. Ignoring..."},
	//OBSOLETE_ER_NDB_BINLOG_UNHANDLED_ERROR_FOR_TABLE : {10686,[]string{"HY000"},"NDB Binlog: unhandled error %d for table %s"},
	//OBSOLETE_ER_NDB_BINLOG_CLUSTER_FAILURE : {10687,[]string{"HY000"},"NDB Binlog: cluster failure for %s at epoch %u/%u."},
	//OBSOLETE_ER_NDB_BINLOG_UNKNOWN_NON_DATA_EVENT : {10688,[]string{"HY000"},"NDB Binlog: unknown non data event %d for %s. Ignoring..."},
	//OBSOLETE_ER_NDB_BINLOG_INJECTOR_DISCARDING_ROW_EVENT_METADATA : {10689,[]string{"HY000"},"NDB: Binlog Injector discarding row event meta data as server is using v1 row events. (%u %x)"},
	//OBSOLETE_ER_NDB_REMAINING_OPEN_TABLES : {10690,[]string{"HY000"},"remove_all_event_operations: Remaining open tables: "},
	//OBSOLETE_ER_NDB_REMAINING_OPEN_TABLE_INFO : {10691,[]string{"HY000"},"  %s.%s, use_count: %u"},
	//OBSOLETE_ER_NDB_COULD_NOT_GET_APPLY_STATUS_SHARE : {10692,[]string{"HY000"},"NDB: Could not get apply status share"},
	//OBSOLETE_ER_NDB_BINLOG_SERVER_SHUTDOWN_DURING_NDB_CLUSTER_START : {10693,[]string{"HY000"},"NDB Binlog: Server shutdown detected while waiting for ndbcluster to start..."},
	//OBSOLETE_ER_NDB_BINLOG_CLUSTER_RESTARTED_RESET_MASTER_SUGGESTED : {10694,[]string{"HY000"},"NDB Binlog: cluster has been restarted --initial or with older filesystem. ndb_latest_handled_binlog_epoch: %u/%u, while current epoch: %u/%u. RESET MASTER should be issued. Resetting ndb_latest_handled_binlog_epoch."},
	//OBSOLETE_ER_NDB_BINLOG_CLUSTER_HAS_RECONNECTED : {10695,[]string{"HY000"},"NDB Binlog: cluster has reconnected. Changes to the database that occurred while disconnected will not be in the binlog"},
	//OBSOLETE_ER_NDB_BINLOG_STARTING_LOG_AT_EPOCH : {10696,[]string{"HY000"},"NDB Binlog: starting log at epoch %u/%u"},
	//OBSOLETE_ER_NDB_BINLOG_NDB_TABLES_WRITABLE : {10697,[]string{"HY000"},"NDB Binlog: ndb tables writable"},
	//OBSOLETE_ER_NDB_BINLOG_SHUTDOWN_DETECTED : {10698,[]string{"HY000"},"NDB Binlog: Server shutdown detected..."},
	//OBSOLETE_ER_NDB_BINLOG_LOST_SCHEMA_CONNECTION_WAITING : {10699,[]string{"HY000"},"NDB Binlog: Just lost schema connection, hanging around"},
	//OBSOLETE_ER_NDB_BINLOG_LOST_SCHEMA_CONNECTION_CONTINUING : {10700,[]string{"HY000"},"NDB Binlog: ...and on our way"},
	//OBSOLETE_ER_NDB_BINLOG_ERROR_HANDLING_SCHEMA_EVENT : {10701,[]string{"HY000"},"NDB: error %lu (%s) on handling binlog schema event"},
	//OBSOLETE_ER_NDB_BINLOG_CANT_INJECT_APPLY_STATUS_WRITE_ROW : {10702,[]string{"HY000"},"NDB Binlog: Failed to inject apply status write row"},
	//OBSOLETE_ER_NDB_BINLOG_ERROR_DURING_GCI_ROLLBACK : {10703,[]string{"HY000"},"NDB Binlog: Error during ROLLBACK of GCI %u/%u. Error: %d"},
	//OBSOLETE_ER_NDB_BINLOG_ERROR_DURING_GCI_COMMIT : {10704,[]string{"HY000"},"NDB Binlog: Error during COMMIT of GCI. Error: %d"},
	//OBSOLETE_ER_NDB_BINLOG_LATEST_TRX_IN_EPOCH_NOT_IN_BINLOG : {10705,[]string{"HY000"},"NDB Binlog: latest transaction in epoch %u/%u not in binlog as latest handled epoch is %u/%u"},
	//OBSOLETE_ER_NDB_BINLOG_RELEASING_EXTRA_SHARE_REFERENCES : {10706,[]string{"HY000"},"NDB Binlog: Release extra share references"},
	//OBSOLETE_ER_NDB_BINLOG_REMAINING_OPEN_TABLES : {10707,[]string{"HY000"},"NDB Binlog: remaining open tables: "},
	//OBSOLETE_ER_NDB_BINLOG_REMAINING_OPEN_TABLE_INFO : {10708,[]string{"HY000"},"  %s.%s state: %u use_count: %u"},
	ER_TREE_CORRUPT_PARENT_SHOULD_POINT_AT_PARENT : {10709,[]string{"HY000"},"Wrong tree: Parent doesn't point at parent"},
	ER_TREE_CORRUPT_ROOT_SHOULD_BE_BLACK : {10710,[]string{"HY000"},"Wrong tree: Root should be black"},
	ER_TREE_CORRUPT_2_CONSECUTIVE_REDS : {10711,[]string{"HY000"},"Wrong tree: Found two red in a row"},
	ER_TREE_CORRUPT_RIGHT_IS_LEFT : {10712,[]string{"HY000"},"Wrong tree: Found right == left"},
	ER_TREE_CORRUPT_INCORRECT_BLACK_COUNT : {10713,[]string{"HY000"},"Wrong tree: Incorrect black-count: %d - %d"},
	ER_WRONG_COUNT_FOR_ORIGIN : {10714,[]string{"HY000"},"Use_count: Wrong count %lu for origin %p"},
	ER_WRONG_COUNT_FOR_KEY : {10715,[]string{"HY000"},"Use_count: Wrong count for key at %p, %lu should be %lu"},
	ER_WRONG_COUNT_OF_ELEMENTS : {10716,[]string{"HY000"},"Wrong number of elements: %u (should be %u) for tree at %p"},
	ER_RPL_ERROR_READING_SLAVE_WORKER_CONFIGURATION : {10717,[]string{"HY000"},"Error reading slave worker configuration"},
	//OBSOLETE_ER_RPL_ERROR_WRITING_SLAVE_WORKER_CONFIGURATION : {10718,[]string{"HY000"},"Error writing slave worker configuration"},
	ER_RPL_FAILED_TO_OPEN_RELAY_LOG : {10719,[]string{"HY000"},"Failed to open relay log %s, error: %s"},
	ER_RPL_WORKER_CANT_READ_RELAY_LOG : {10720,[]string{"HY000"},"Error when worker read relay log events,relay log name %s, position %llu"},
	ER_RPL_WORKER_CANT_FIND_NEXT_RELAY_LOG : {10721,[]string{"HY000"},"Failed to find next relay log when retrying the transaction, current relay log is %s"},
	ER_RPL_MTS_SLAVE_COORDINATOR_HAS_WAITED : {10722,[]string{"HY000"},"Multi-threaded slave: Coordinator has waited %lu times hitting slave_pending_jobs_size_max; current event size = %zu."},
	ER_BINLOG_FAILED_TO_WRITE_DROP_FOR_TEMP_TABLES : {10723,[]string{"HY000"},"Failed to write the DROP statement for temporary tables to binary log"},
	ER_BINLOG_OOM_WRITING_DELETE_WHILE_OPENING_HEAP_TABLE : {10724,[]string{"HY000"},"When opening HEAP table, could not allocate memory to write 'DELETE FROM `%s`.`%s`' to the binary log"},
	ER_FAILED_TO_REPAIR_TABLE : {10725,[]string{"HY000"},"Couldn't repair table: %s.%s"},
	ER_FAILED_TO_REMOVE_TEMP_TABLE : {10726,[]string{"HY000"},"Could not remove temporary table: '%s', error: %d"},
	ER_SYSTEM_TABLE_NOT_TRANSACTIONAL : {10727,[]string{"HY000"},"System table '%.*s' is expected to be transactional."},
	ER_RPL_ERROR_WRITING_MASTER_CONFIGURATION : {10728,[]string{"HY000"},"Error writing master configuration."},
	ER_RPL_ERROR_READING_MASTER_CONFIGURATION : {10729,[]string{"HY000"},"Error reading master configuration."},
	ER_RPL_SSL_INFO_IN_MASTER_INFO_IGNORED : {10730,[]string{"HY000"},"SSL information in the master info file are ignored because this MySQL slave was compiled without SSL support."},
	ER_PLUGIN_FAILED_DEINITIALIZATION : {10731,[]string{"HY000"},"Plugin '%s' of type %s failed deinitialization"},
	ER_PLUGIN_HAS_NONZERO_REFCOUNT_AFTER_DEINITIALIZATION : {10732,[]string{"HY000"},"Plugin '%s' has ref_count=%d after deinitialization."},
	ER_PLUGIN_SHUTTING_DOWN_PLUGIN : {10733,[]string{"HY000"},"Shutting down plugin '%s'"},
	ER_PLUGIN_REGISTRATION_FAILED : {10734,[]string{"HY000"},"Plugin '%s' registration as a %s failed."},
	ER_PLUGIN_CANT_OPEN_PLUGIN_TABLE : {10735,[]string{"HY000"},"Could not open the mysql.plugin table. Please perform the MySQL upgrade procedure."},
	ER_PLUGIN_CANT_LOAD : {10736,[]string{"HY000"},"Couldn't load plugin named '%s' with soname '%s'."},
	ER_PLUGIN_LOAD_PARAMETER_TOO_LONG : {10737,[]string{"HY000"},"plugin-load parameter too long"},
	ER_PLUGIN_FORCING_SHUTDOWN : {10738,[]string{"HY000"},"Plugin '%s' will be forced to shutdown"},
	ER_PLUGIN_HAS_NONZERO_REFCOUNT_AFTER_SHUTDOWN : {10739,[]string{"HY000"},"Plugin '%s' has ref_count=%d after shutdown."},
	ER_PLUGIN_UNKNOWN_VARIABLE_TYPE : {10740,[]string{"HY000"},"Unknown variable type code 0x%x in plugin '%s'."},
	ER_PLUGIN_VARIABLE_SET_READ_ONLY : {10741,[]string{"HY000"},"Server variable %s of plugin %s was forced to be read-only: string variable without update_func and PLUGIN_VAR_MEMALLOC flag"},
	ER_PLUGIN_VARIABLE_MISSING_NAME : {10742,[]string{"HY000"},"Missing variable name in plugin '%s'."},
	ER_PLUGIN_VARIABLE_NOT_ALLOCATED_THREAD_LOCAL : {10743,[]string{"HY000"},"Thread local variable '%s' not allocated in plugin '%s'."},
	ER_PLUGIN_OOM : {10744,[]string{"HY000"},"Out of memory for plugin '%s'."},
	ER_PLUGIN_BAD_OPTIONS : {10745,[]string{"HY000"},"Bad options for plugin '%s'."},
	ER_PLUGIN_PARSING_OPTIONS_FAILED : {10746,[]string{"HY000"},"Parsing options for plugin '%s' failed."},
	ER_PLUGIN_DISABLED : {10747,[]string{"HY000"},"Plugin '%s' is disabled."},
	ER_PLUGIN_HAS_CONFLICTING_SYSTEM_VARIABLES : {10748,[]string{"HY000"},"Plugin '%s' has conflicting system variables"},
	ER_PLUGIN_CANT_SET_PERSISTENT_OPTIONS : {10749,[]string{"HY000"},"Setting persistent options for plugin '%s' failed."},
	ER_MY_NET_WRITE_FAILED_FALLING_BACK_ON_STDERR : {10750,[]string{"HY000"},"Failed on my_net_write, writing to stderr instead: %s"},
	ER_RETRYING_REPAIR_WITHOUT_QUICK : {10751,[]string{"HY000"},"Retrying repair of: '%s' without quick"},
	ER_RETRYING_REPAIR_WITH_KEYCACHE : {10752,[]string{"HY000"},"Retrying repair of: '%s' with keycache"},
	ER_FOUND_ROWS_WHILE_REPAIRING : {10753,[]string{"HY000"},"Found %s of %s rows when repairing '%s'"},
	ER_ERROR_DURING_OPTIMIZE_TABLE : {10754,[]string{"HY000"},"Warning: Optimize table got errno %d on %s.%s, retrying"},
	ER_ERROR_ENABLING_KEYS : {10755,[]string{"HY000"},"Warning: Enabling keys got errno %d on %s.%s, retrying"},
	ER_CHECKING_TABLE : {10756,[]string{"HY000"},"Checking table:   '%s'"},
	ER_RECOVERING_TABLE : {10757,[]string{"HY000"},"Recovering table: '%s'"},
	ER_CANT_CREATE_TABLE_SHARE_FROM_FRM : {10758,[]string{"HY000"},"Error in creating TABLE_SHARE from %s.frm file."},
	ER_CANT_LOCK_TABLE : {10759,[]string{"HY000"},"Unable to acquire lock on %s.%s"},
	ER_CANT_ALLOC_TABLE_OBJECT : {10760,[]string{"HY000"},"Error in allocation memory for TABLE object."},
	ER_CANT_CREATE_HANDLER_OBJECT_FOR_TABLE : {10761,[]string{"HY000"},"Error in creating handler object for table %s.%s"},
	ER_CANT_SET_HANDLER_REFERENCE_FOR_TABLE : {10762,[]string{"HY000"},"Error in setting handler reference for table %s.%s"},
	ER_CANT_LOCK_TABLESPACE : {10763,[]string{"HY000"},"Unable to acquire lock on tablespace name %s"},
	ER_CANT_UPGRADE_GENERATED_COLUMNS_TO_DD : {10764,[]string{"HY000"},"Error in processing generated columns for table %s.%s"},
	ER_DD_ERROR_CREATING_ENTRY : {10765,[]string{"HY000"},"Error in Creating DD entry for %s.%s"},
	ER_DD_CANT_FETCH_TABLE_DATA : {10766,[]string{"HY000"},"Error in fetching %s.%s table data from dictionary"},
	ER_DD_CANT_FIX_SE_DATA : {10767,[]string{"HY000"},"Error in fixing SE data for %s.%s"},
	ER_DD_CANT_CREATE_SP : {10768,[]string{"HY000"},"Error in creating stored program '%s.%s'"},
	ER_CANT_OPEN_DB_OPT_USING_DEFAULT_CHARSET : {10769,[]string{"HY000"},"Unable to open db.opt file %s. Using default Character set."},
	ER_CANT_CREATE_CACHE_FOR_DB_OPT : {10770,[]string{"HY000"},"Unable to intialize IO cache to open db.opt file %s. "},
	ER_CANT_IDENTIFY_CHARSET_USING_DEFAULT : {10771,[]string{"HY000"},"Unable to identify the charset in %s. Using default character set."},
	ER_DB_OPT_NOT_FOUND_USING_DEFAULT_CHARSET : {10772,[]string{"HY000"},"db.opt file not found for %s database. Using default Character set."},
	ER_EVENT_CANT_GET_TIMEZONE_FROM_FIELD : {10773,[]string{"HY000"},"Event '%s'.'%s': invalid value in column mysql.event.time_zone."},
	ER_EVENT_CANT_FIND_TIMEZONE : {10774,[]string{"HY000"},"Event '%s'.'%s': has invalid time zone value "},
	ER_EVENT_CANT_GET_CHARSET : {10775,[]string{"HY000"},"Event '%s'.'%s': invalid value in column mysql.event.character_set_client."},
	ER_EVENT_CANT_GET_COLLATION : {10776,[]string{"HY000"},"Event '%s'.'%s': invalid value in column mysql.event.collation_connection."},
	ER_EVENT_CANT_OPEN_TABLE_MYSQL_EVENT : {10777,[]string{"HY000"},"Failed to open mysql.event Table."},
	ER_CANT_PARSE_STORED_ROUTINE_BODY : {10778,[]string{"HY000"},"Parsing '%s.%s' routine body failed.%s"},
	ER_CANT_OPEN_TABLE_MYSQL_PROC : {10779,[]string{"HY000"},"Failed to open mysql.proc Table."},
	ER_CANT_READ_TABLE_MYSQL_PROC : {10780,[]string{"HY000"},"Failed to read mysql.proc table."},
	ER_FILE_EXISTS_DURING_UPGRADE : {10781,[]string{"HY000"},"Found %s file in mysql schema. DD will create .ibd file with same name. Please rename table and start upgrade process again."},
	ER_CANT_OPEN_DATADIR_AFTER_UPGRADE_FAILURE : {10782,[]string{"HY000"},"Unable to open the data directory %s during clean up after upgrade failed"},
	ER_CANT_SET_PATH_FOR : {10783,[]string{"HY000"},"Failed to set path %s"},
	ER_CANT_OPEN_DIR : {10784,[]string{"HY000"},"Failed to open dir %s"},
	//OBSOLETE_ER_NDB_CLUSTER_CONNECTION_POOL_NODEIDS : {0000,[]string{""},"NDB: Found empty nodeid specified in --ndb-cluster-connection-pool-nodeids='%s'."},
	//OBSOLETE_ER_NDB_CANT_PARSE_NDB_CLUSTER_CONNECTION_POOL_NODEIDS : {10786,[]string{"HY000"},"NDB: Could not parse '%s' in --ndb-cluster-connection-pool-nodeids='%s'."},
	//OBSOLETE_ER_NDB_INVALID_CLUSTER_CONNECTION_POOL_NODEIDS : {0000,[]string{""},"NDB: Invalid nodeid %d in --ndb-cluster-connection-pool-nodeids='%s'."},
	//OBSOLETE_ER_NDB_DUPLICATE_CLUSTER_CONNECTION_POOL_NODEIDS : {0000,[]string{""},"NDB: Found duplicate nodeid %d in --ndb-cluster-connection-pool-nodeids='%s'."},
	//OBSOLETE_ER_NDB_POOL_SIZE_CLUSTER_CONNECTION_POOL_NODEIDS : {0000,[]string{""},"NDB: The size of the cluster connection pool must be equal to the number of nodeids in --ndb-cluster-connection-pool-nodeids='%s'."},
	//OBSOLETE_ER_NDB_NODEID_NOT_FIRST_CONNECTION_POOL_NODEIDS : {0000,[]string{""},"NDB: The nodeid specified by --ndb-nodeid must be equal to the first nodeid in --ndb-cluster-connection-pool-nodeids='%s'."},
	//OBSOLETE_ER_NDB_USING_NODEID : {10791,[]string{"HY000"},"NDB: using nodeid %u"},
	//OBSOLETE_ER_NDB_CANT_ALLOC_GLOBAL_NDB_CLUSTER_CONNECTION : {10792,[]string{"HY000"},"NDB: failed to allocate global ndb cluster connection"},
	//OBSOLETE_ER_NDB_CANT_ALLOC_GLOBAL_NDB_OBJECT : {10793,[]string{"HY000"},"NDB: failed to allocate global ndb object"},
	//OBSOLETE_ER_NDB_USING_NODEID_LIST : {10794,[]string{"HY000"},"NDB[%u]: using nodeid %u"},
	//OBSOLETE_ER_NDB_CANT_ALLOC_NDB_CLUSTER_CONNECTION : {10795,[]string{"HY000"},"NDB[%u]: failed to allocate cluster connect object"},
	//OBSOLETE_ER_NDB_STARTING_CONNECT_THREAD : {10796,[]string{"HY000"},"NDB[%u]: starting connect thread"},
	//OBSOLETE_ER_NDB_NODE_INFO : {10797,[]string{"HY000"},"NDB[%u]: NodeID: %d, %s"},
	//OBSOLETE_ER_NDB_CANT_START_CONNECT_THREAD : {10798,[]string{"HY000"},"NDB[%u]: failed to start connect thread"},
	//OBSOLETE_ER_NDB_GENERIC_ERROR : {10799,[]string{"HY000"},"NDB: error (%u) %s"},
	//OBSOLETE_ER_NDB_CPU_MASK_TOO_SHORT : {10800,[]string{"HY000"},"Ignored receive thread CPU mask, mask too short, %u CPUs needed in mask, only %u CPUs provided"},
	ER_EVENT_ERROR_CREATING_QUERY_TO_WRITE_TO_BINLOG : {10801,[]string{"HY000"},"Event Error: An error occurred while creating query string, before writing it into binary log."},
	ER_EVENT_SCHEDULER_ERROR_LOADING_FROM_DB : {10802,[]string{"HY000"},"Event Scheduler: Error while loading from disk."},
	ER_EVENT_SCHEDULER_ERROR_GETTING_EVENT_OBJECT : {10803,[]string{"HY000"},"Event Scheduler: Error getting event object."},
	ER_EVENT_SCHEDULER_GOT_BAD_DATA_FROM_TABLE : {10804,[]string{"HY000"},"Event Scheduler: Error while loading events from mysql.events.The table probably contains bad data or is corrupted"},
	ER_EVENT_CANT_GET_LOCK_FOR_DROPPING_EVENT : {10805,[]string{"HY000"},"Unable to obtain lock for dropping event %s from schema %s"},
	ER_EVENT_UNABLE_TO_DROP_EVENT : {10806,[]string{"HY000"},"Unable to drop event %s from schema %s"},
	//OBSOLETE_ER_BINLOG_ATTACHING_THREAD_MEMORY_FINALLY_AVAILABLE : {10807,[]string{"HY000"},"Server overcomes the temporary 'out of memory' in '%d' tries while attaching to session thread during the group commit phase."},
	ER_BINLOG_CANT_RESIZE_CACHE : {10808,[]string{"HY000"},"Unable to resize binlog IOCACHE auxiliary file"},
	ER_BINLOG_FILE_BEING_READ_NOT_PURGED : {10809,[]string{"HY000"},"file %s was not purged because it was being read by thread number %u"},
	ER_BINLOG_IO_ERROR_READING_HEADER : {10810,[]string{"HY000"},"I/O error reading the header from the binary log, errno=%d, io cache code=%d"},
	//OBSOLETE_ER_BINLOG_CANT_OPEN_LOG : {10811,[]string{"HY000"},"Failed to open log (file '%s', errno %d)"},
	//OBSOLETE_ER_BINLOG_CANT_CREATE_CACHE_FOR_LOG : {10812,[]string{"HY000"},"Failed to create a cache on log (file '%s')"},
	ER_BINLOG_FILE_EXTENSION_NUMBER_EXHAUSTED : {10813,[]string{"HY000"},"Log filename extension number exhausted: %06lu. Please fix this by archiving old logs and updating the index files."},
	ER_BINLOG_FILE_NAME_TOO_LONG : {10814,[]string{"HY000"},"Log filename too large: %s%s (%zu). Please fix this by archiving old logs and updating the index files."},
	ER_BINLOG_FILE_EXTENSION_NUMBER_RUNNING_LOW : {10815,[]string{"HY000"},"Next log extension: %lu. Remaining log filename extensions: %lu. Please consider archiving some logs."},
	ER_BINLOG_CANT_OPEN_FOR_LOGGING : {10816,[]string{"HY000"},"Could not open %s for logging (error %d). Turning logging off for the whole duration of the MySQL server process. To turn it on again: fix the cause, shutdown the MySQL server and restart it."},
	ER_BINLOG_FAILED_TO_SYNC_INDEX_FILE : {10817,[]string{"HY000"},"MYSQL_BIN_LOG::open_index_file failed to sync the index file."},
	ER_BINLOG_ERROR_READING_GTIDS_FROM_RELAY_LOG : {10818,[]string{"HY000"},"Error reading GTIDs from relaylog: %d"},
	ER_BINLOG_EVENTS_READ_FROM_RELAY_LOG_INFO : {10819,[]string{"HY000"},"%lu events read in relaylog file '%s' for updating Retrieved_Gtid_Set and/or IO thread transaction parser state."},
	ER_BINLOG_ERROR_READING_GTIDS_FROM_BINARY_LOG : {10820,[]string{"HY000"},"Error reading GTIDs from binary log: %d"},
	ER_BINLOG_EVENTS_READ_FROM_BINLOG_INFO : {10821,[]string{"HY000"},"Read %lu events from binary log file '%s' to determine the GTIDs purged from binary logs."},
	ER_BINLOG_CANT_GENERATE_NEW_FILE_NAME : {10822,[]string{"HY000"},"MYSQL_BIN_LOG::open failed to generate new file name."},
	ER_BINLOG_FAILED_TO_SYNC_INDEX_FILE_IN_OPEN : {10823,[]string{"HY000"},"MYSQL_BIN_LOG::open failed to sync the index file."},
	ER_BINLOG_CANT_USE_FOR_LOGGING : {10824,[]string{"HY000"},"Could not use %s for logging (error %d). Turning logging off for the whole duration of the MySQL server process. To turn it on again: fix the cause, shutdown the MySQL server and restart it."},
	ER_BINLOG_FAILED_TO_CLOSE_INDEX_FILE_WHILE_REBUILDING : {10825,[]string{"HY000"},"While rebuilding index file %s: Failed to close the index file."},
	ER_BINLOG_FAILED_TO_DELETE_INDEX_FILE_WHILE_REBUILDING : {10826,[]string{"HY000"},"While rebuilding index file %s: Failed to delete the existing index file. It could be that file is being used by some other process."},
	ER_BINLOG_FAILED_TO_RENAME_INDEX_FILE_WHILE_REBUILDING : {10827,[]string{"HY000"},"While rebuilding index file %s: Failed to rename the new index file to the existing index file."},
	ER_BINLOG_FAILED_TO_OPEN_INDEX_FILE_AFTER_REBUILDING : {10828,[]string{"HY000"},"After rebuilding the index file %s: Failed to open the index file."},
	ER_BINLOG_CANT_APPEND_LOG_TO_TMP_INDEX : {10829,[]string{"HY000"},"MYSQL_BIN_LOG::add_log_to_index failed to append log file name: %s, to crash safe index file."},
	ER_BINLOG_CANT_LOCATE_OLD_BINLOG_OR_RELAY_LOG_FILES : {10830,[]string{"HY000"},"Failed to locate old binlog or relay log files"},
	ER_BINLOG_CANT_DELETE_FILE : {10831,[]string{"HY000"},"Failed to delete file '%s'"},
	ER_BINLOG_CANT_SET_TMP_INDEX_NAME : {10832,[]string{"HY000"},"MYSQL_BIN_LOG::set_crash_safe_index_file_name failed to set file name."},
	ER_BINLOG_FAILED_TO_OPEN_TEMPORARY_INDEX_FILE : {10833,[]string{"HY000"},"MYSQL_BIN_LOG::open_crash_safe_index_file failed to open temporary index file."},
	//OBSOLETE_ER_BINLOG_ERROR_GETTING_NEXT_LOG_FROM_INDEX : {10834,[]string{"HY000"},"next log error: %d  offset: %s  log: %s included: %d"},
	ER_BINLOG_CANT_OPEN_TMP_INDEX : {10835,[]string{"HY000"},"%s failed to open the crash safe index file."},
	ER_BINLOG_CANT_COPY_INDEX_TO_TMP : {10836,[]string{"HY000"},"%s failed to copy index file to crash safe index file."},
	ER_BINLOG_CANT_CLOSE_TMP_INDEX : {10837,[]string{"HY000"},"%s failed to close the crash safe index file."},
	ER_BINLOG_CANT_MOVE_TMP_TO_INDEX : {10838,[]string{"HY000"},"%s failed to move crash safe index file to index file."},
	ER_BINLOG_PURGE_LOGS_CALLED_WITH_FILE_NOT_IN_INDEX : {10839,[]string{"HY000"},"MYSQL_BIN_LOG::purge_logs was called with file %s not listed in the index."},
	ER_BINLOG_PURGE_LOGS_CANT_SYNC_INDEX_FILE : {10840,[]string{"HY000"},"MYSQL_BIN_LOG::purge_logs failed to sync the index file."},
	ER_BINLOG_PURGE_LOGS_CANT_COPY_TO_REGISTER_FILE : {10841,[]string{"HY000"},"MYSQL_BIN_LOG::purge_logs failed to copy %s to register file."},
	ER_BINLOG_PURGE_LOGS_CANT_FLUSH_REGISTER_FILE : {10842,[]string{"HY000"},"MYSQL_BIN_LOG::purge_logs failed to flush register file."},
	ER_BINLOG_PURGE_LOGS_CANT_UPDATE_INDEX_FILE : {10843,[]string{"HY000"},"MYSQL_BIN_LOG::purge_logs failed to update the index file"},
	ER_BINLOG_PURGE_LOGS_FAILED_TO_PURGE_LOG : {10844,[]string{"HY000"},"MYSQL_BIN_LOG::purge_logs failed to process registered files that would be purged."},
	ER_BINLOG_FAILED_TO_SET_PURGE_INDEX_FILE_NAME : {10845,[]string{"HY000"},"MYSQL_BIN_LOG::set_purge_index_file_name failed to set file name."},
	ER_BINLOG_FAILED_TO_OPEN_REGISTER_FILE : {10846,[]string{"HY000"},"MYSQL_BIN_LOG::open_purge_index_file failed to open register file."},
	ER_BINLOG_FAILED_TO_REINIT_REGISTER_FILE : {10847,[]string{"HY000"},"MYSQL_BIN_LOG::purge_index_entry failed to reinit register file for read"},
	ER_BINLOG_FAILED_TO_READ_REGISTER_FILE : {10848,[]string{"HY000"},"MYSQL_BIN_LOG::purge_index_entry error %d reading from register file."},
	ER_CANT_STAT_FILE : {10849,[]string{"HY000"},"Failed to execute mysql_file_stat on file '%s'"},
	ER_BINLOG_CANT_DELETE_LOG_FILE_DOES_INDEX_MATCH_FILES : {10850,[]string{"HY000"},"Failed to delete log file '%s'; consider examining correspondence of your binlog index file to the actual binlog files"},
	ER_BINLOG_CANT_DELETE_FILE_AND_READ_BINLOG_INDEX : {10851,[]string{"HY000"},"Failed to delete file '%s' and read the binlog index file"},
	ER_BINLOG_FAILED_TO_DELETE_LOG_FILE : {10852,[]string{"HY000"},"Failed to delete log file '%s'"},
	ER_BINLOG_LOGGING_INCIDENT_TO_STOP_SLAVES : {10853,[]string{"HY000"},"%s An incident event has been written to the binary log which will stop the slaves."},
	ER_BINLOG_CANT_FIND_LOG_IN_INDEX : {10854,[]string{"HY000"},"find_log_pos() failed (error: %d)"},
	ER_BINLOG_RECOVERING_AFTER_CRASH_USING : {10855,[]string{"HY000"},"Recovering after a crash using %s"},
	ER_BINLOG_CANT_OPEN_CRASHED_BINLOG : {10856,[]string{"HY000"},"Failed to open the crashed binlog file when master server is recovering it."},
	ER_BINLOG_CANT_TRIM_CRASHED_BINLOG : {10857,[]string{"HY000"},"Failed to trim the crashed binlog file when master server is recovering it."},
	ER_BINLOG_CRASHED_BINLOG_TRIMMED : {10858,[]string{"HY000"},"Crashed binlog file %s size is %llu, but recovered up to %llu. Binlog trimmed to %llu bytes."},
	ER_BINLOG_CANT_CLEAR_IN_USE_FLAG_FOR_CRASHED_BINLOG : {10859,[]string{"HY000"},"Failed to clear LOG_EVENT_BINLOG_IN_USE_F for the crashed binlog file when master server is recovering it."},
	ER_BINLOG_FAILED_TO_RUN_AFTER_SYNC_HOOK : {10860,[]string{"HY000"},"Failed to run 'after_sync' hooks"},
	ER_TURNING_LOGGING_OFF_FOR_THE_DURATION : {10861,[]string{"HY000"},"%s Hence turning logging off for the whole duration of the MySQL server process. To turn it on again: fix the cause, shutdown the MySQL server and restart it."},
	ER_BINLOG_FAILED_TO_RUN_AFTER_FLUSH_HOOK : {10862,[]string{"HY000"},"Failed to run 'after_flush' hooks"},
	ER_BINLOG_CRASH_RECOVERY_FAILED : {10863,[]string{"HY000"},"Crash recovery failed. Either correct the problem (if it's, for example, out of memory error) and restart, or delete (or rename) binary log and start mysqld with --tc-heuristic-recover={commit|rollback}"},
	ER_BINLOG_WARNING_SUPPRESSED : {10864,[]string{"HY000"},"The following warning was suppressed %d times during the last %d seconds in the error log"},
	ER_NDB_LOG_ENTRY : {10865,[]string{"HY000"},"NDB: %s"},
	ER_NDB_LOG_ENTRY_WITH_PREFIX : {10866,[]string{"HY000"},"NDB %s: %s"},
	//OBSOLETE_ER_NDB_BINLOG_CANT_CREATE_PURGE_THD : {10867,[]string{"HY000"},"NDB: Unable to purge %s.%s File=%s (failed to setup thd)"},
	ER_INNODB_UNKNOWN_COLLATION : {10868,[]string{"HY000"},"Unknown collation #%lu."},
	ER_INNODB_INVALID_LOG_GROUP_HOME_DIR : {10869,[]string{"HY000"},"syntax error in innodb_log_group_home_dir"},
	ER_INNODB_INVALID_INNODB_UNDO_DIRECTORY : {10870,[]string{"HY000"},"syntax error in innodb_undo_directory"},
	ER_INNODB_ILLEGAL_COLON_IN_POOL : {10871,[]string{"HY000"},"InnoDB: innodb_buffer_pool_filename cannot have colon (:) in the file name."},
	ER_INNODB_INVALID_PAGE_SIZE : {10872,[]string{"HY000"},"InnoDB: Invalid page size=%lu."},
	ER_INNODB_DIRTY_WATER_MARK_NOT_LOW : {10873,[]string{"HY000"},"InnoDB: innodb_max_dirty_pages_pct_lwm cannot be set higher than innodb_max_dirty_pages_pct. Setting innodb_max_dirty_pages_pct_lwm to %lf"},
	ER_INNODB_IO_CAPACITY_EXCEEDS_MAX : {10874,[]string{"HY000"},"InnoDB: innodb_io_capacity cannot be set higher than innodb_io_capacity_max. Setting innodb_io_capacity to %lu"},
	ER_INNODB_FILES_SAME : {10875,[]string{"HY000"},"%s and %s file names seem to be the same."},
	ER_INNODB_UNREGISTERED_TRX_ACTIVE : {10876,[]string{"HY000"},"Transaction not registered for MySQL 2PC, but transaction is active"},
	ER_INNODB_CLOSING_CONNECTION_ROLLS_BACK : {10877,[]string{"HY000"},"MySQL is closing a connection that has an active InnoDB transaction. %llu row modifications will roll back."},
	ER_INNODB_TRX_XLATION_TABLE_OOM : {10878,[]string{"HY000"},"InnoDB: fail to allocate memory for index translation table. Number of Index:%lu, array size:%lu"},
	ER_INNODB_CANT_FIND_INDEX_IN_INNODB_DD : {10879,[]string{"HY000"},"Cannot find index %s in InnoDB index dictionary."},
	ER_INNODB_INDEX_COLUMN_INFO_UNLIKE_MYSQLS : {10880,[]string{"HY000"},"Found index %s whose column info does not match that of MySQL."},
	//OBSOLETE_ER_INNODB_CANT_OPEN_TABLE : {10881,[]string{"HY000"},"Failed to open table %s."},
	ER_INNODB_CANT_BUILD_INDEX_XLATION_TABLE_FOR : {10882,[]string{"HY000"},"Build InnoDB index translation table for Table %s failed"},
	ER_INNODB_PK_NOT_IN_MYSQL : {10883,[]string{"HY000"},"Table %s has a primary key in InnoDB data dictionary, but not in MySQL!"},
	ER_INNODB_PK_ONLY_IN_MYSQL : {10884,[]string{"HY000"},"Table %s has no primary key in InnoDB data dictionary, but has one in MySQL! If you created the table with a MySQL version < 3.23.54 and did not define a primary key, but defined a unique key with all non-NULL columns, then MySQL internally treats that key as the primary key. You can fix this error by dump + DROP + CREATE + reimport of the table."},
	ER_INNODB_CLUSTERED_INDEX_PRIVATE : {10885,[]string{"HY000"},"Table %s key_used_on_scan is %lu even though there is no primary key inside InnoDB."},
	//OBSOLETE_ER_INNODB_PARTITION_TABLE_LOWERCASED : {10886,[]string{"HY000"},"Partition table %s opened after converting to lower case. The table may have been moved from a case-insensitive file system. Please recreate the table in the current file system."},
	ER_ERRMSG_REPLACEMENT_DODGY : {10887,[]string{"HY000"},"Cannot replace error message (%s,%s,%s) \"%s\" with \"%s\"; wrong number or type of %% subsitutions."},
	ER_ERRMSG_REPLACEMENTS_FAILED : {10888,[]string{"HY000"},"Table for error message replacements could not be found or read, or one or more replacements could not be applied."},
	ER_NPIPE_CANT_CREATE : {10889,[]string{"HY000"},"%s: %s"},
	ER_PARTITION_MOVE_CREATED_DUPLICATE_ROW_PLEASE_FIX : {10890,[]string{"HY000"},"Table '%-192s': Delete from part %d failed with error %d. But it was already inserted into part %d, when moving the misplaced row! Please manually fix the duplicate row: %s"},
	ER_AUDIT_CANT_ABORT_COMMAND : {10891,[]string{"HY000"},"Command '%s' cannot be aborted. The trigger error was (%d) [%s]: %s"},
	ER_AUDIT_CANT_ABORT_EVENT : {10892,[]string{"HY000"},"Event '%s' cannot be aborted. The trigger error was (%d) [%s]: %s"},
	ER_AUDIT_WARNING : {10893,[]string{"HY000"},"%s. The trigger error was (%d) [%s]: %s"},
	//OBSOLETE_ER_NDB_NUMBER_OF_CHANNELS : {10894,[]string{"HY000"},"Slave SQL: Configuration with number of replication masters = %u' is not supported when applying to Ndb"},
	//OBSOLETE_ER_NDB_SLAVE_PARALLEL_WORKERS : {10895,[]string{"HY000"},"Slave SQL: Configuration 'slave_parallel_workers = %lu' is not supported when applying to Ndb"},
	//OBSOLETE_ER_NDB_DISTRIBUTING_ERR : {10896,[]string{"HY000"},"NDB %s: distributing %s err: %u"},
	ER_RPL_SLAVE_INSECURE_CHANGE_MASTER : {10897,[]string{"HY000"},"Storing MySQL user name or password information in the master info repository is not secure and is therefore not recommended. Please consider using the USER and PASSWORD connection options for START SLAVE; see the 'START SLAVE Syntax' in the MySQL Manual for more information."},
	//OBSOLETE_ER_RPL_SLAVE_FLUSH_RELAY_LOGS_NOT_ALLOWED : {10898,[]string{"HY000"},"FLUSH RELAY LOGS cannot be performed on channel '%-.192s'."},
	ER_RPL_SLAVE_INCORRECT_CHANNEL : {10899,[]string{"HY000"},"Slave channel '%s' does not exist."},
	ER_FAILED_TO_FIND_DL_ENTRY : {10900,[]string{"HY000"},"Can't find symbol '%-.128s' in library."},
	ER_FAILED_TO_OPEN_SHARED_LIBRARY : {10901,[]string{"HY000"},"Can't open shared library '%-.192s' (errno: %d %-.128s)."},
	ER_THREAD_PRIORITY_IGNORED : {10902,[]string{"HY000"},"Thread priority attribute setting in Resource Group SQL shall be ignored due to unsupported platform or insufficient privilege."},
	ER_BINLOG_CACHE_SIZE_TOO_LARGE : {10903,[]string{"HY000"},"Option binlog_cache_size (%lu) is greater than max_binlog_cache_size (%lu); setting binlog_cache_size equal to max_binlog_cache_size."},
	ER_BINLOG_STMT_CACHE_SIZE_TOO_LARGE : {10904,[]string{"HY000"},"Option binlog_stmt_cache_size (%lu) is greater than max_binlog_stmt_cache_size (%lu); setting binlog_stmt_cache_size equal to max_binlog_stmt_cache_size."},
	ER_FAILED_TO_GENERATE_UNIQUE_LOGFILE : {10905,[]string{"HY000"},"Can't generate a unique log-filename %-.200s.(1-999)."},
	ER_FAILED_TO_READ_FILE : {10906,[]string{"HY000"},"Error reading file '%-.200s' (errno: %d - %s)"},
	ER_FAILED_TO_WRITE_TO_FILE : {10907,[]string{"HY000"},"Error writing file '%-.200s' (errno: %d - %s)"},
	ER_BINLOG_UNSAFE_MESSAGE_AND_STATEMENT : {10908,[]string{"HY000"},"%s Statement: %s"},
	ER_FORCE_CLOSE_THREAD : {10909,[]string{"HY000"},"%s: Forcing close of thread %ld  user: '%-.48s'."},
	ER_SERVER_SHUTDOWN_COMPLETE : {10910,[]string{"HY000"},"%s: Shutdown complete (mysqld %s)  %s."},
	ER_RPL_CANT_HAVE_SAME_BASENAME : {10911,[]string{"HY000"},"Cannot have same base name '%s' for both binary and relay logs. Please check %s (default '%s' if --log-bin option is not used, default '%s' if --log-bin option is used without argument) and %s (default '%s') options to ensure they do not conflict."},
	ER_RPL_GTID_MODE_REQUIRES_ENFORCE_GTID_CONSISTENCY_ON : {10912,[]string{"HY000"},"GTID_MODE = ON requires ENFORCE_GTID_CONSISTENCY = ON."},
	ER_WARN_NO_SERVERID_SPECIFIED : {10913,[]string{"HY000"},"You have not provided a mandatory server-id. Servers in a replication topology must have unique server-ids. Please refer to the proper server start-up parameters documentation."},
	ER_ABORTING_USER_CONNECTION : {10914,[]string{"HY000"},"Aborted connection %u to db: '%-.192s' user: '%-.48s' host: '%-.255s' (%-.64s)."},
	ER_SQL_MODE_MERGED_WITH_STRICT_MODE : {10915,[]string{"HY000"},"'NO_ZERO_DATE', 'NO_ZERO_IN_DATE' and 'ERROR_FOR_DIVISION_BY_ZERO' sql modes should be used with strict mode. They will be merged with strict mode in a future release."},
	ER_GTID_PURGED_WAS_UPDATED : {10916,[]string{"HY000"},"@@GLOBAL.GTID_PURGED was changed from '%s' to '%s'."},
	ER_GTID_EXECUTED_WAS_UPDATED : {10917,[]string{"HY000"},"@@GLOBAL.GTID_EXECUTED was changed from '%s' to '%s'."},
	ER_DEPRECATE_MSG_WITH_REPLACEMENT : {10918,[]string{"HY000"},"'%s' is deprecated and will be removed in a future release. Please use %s instead."},
	ER_TRG_CREATION_CTX_NOT_SET : {10919,[]string{"HY000"},"Triggers for table `%-.64s`.`%-.64s` have no creation context"},
	ER_FILE_HAS_OLD_FORMAT : {10920,[]string{"HY000"},"'%-.192s' has an old format, you should re-create the '%s' object(s)"},
	ER_VIEW_CREATION_CTX_NOT_SET : {10921,[]string{"HY000"},"View `%-.64s`.`%-.64s` has no creation context"},
	//OBSOLETE_ER_TABLE_NAME_CAUSES_TOO_LONG_PATH : {3913,[]string{"HY000"},"Long database name and identifier for object resulted in a path length too long for table '%s'. Please check the path limit for your OS."},
	ER_TABLE_UPGRADE_REQUIRED : {10923,[]string{"HY000"},"Table upgrade required. Please do \"REPAIR TABLE `%-.64s`\" or dump/reload to fix it!"},
	ER_GET_ERRNO_FROM_STORAGE_ENGINE : {10924,[]string{"HY000"},"Got error %d - '%-.192s' from storage engine."},
	ER_ACCESS_DENIED_ERROR_WITHOUT_PASSWORD : {10925,[]string{"HY000"},"Access denied for user '%-.48s'@'%-.64s'"},
	ER_ACCESS_DENIED_ERROR_WITH_PASSWORD : {10926,[]string{"HY000"},"Access denied for user '%-.48s'@'%-.64s' (using password: %s)"},
	ER_ACCESS_DENIED_FOR_USER_ACCOUNT_LOCKED : {10927,[]string{"HY000"},"Access denied for user '%-.48s'@'%-.64s'. Account is locked."},
	ER_MUST_CHANGE_EXPIRED_PASSWORD : {10928,[]string{"HY000"},"Your password has expired. To log in you must change it using a client that supports expired passwords."},
	ER_SYSTEM_TABLES_NOT_SUPPORTED_BY_STORAGE_ENGINE : {10929,[]string{"HY000"},"Storage engine '%s' does not support system tables. [%s.%s]."},
	//OBSOLETE_ER_FILESORT_TERMINATED : {10930,[]string{"HY000"},"Sort aborted"},
	ER_SERVER_STARTUP_MSG : {10931,[]string{"HY000"},"%s: ready for connections. Version: '%s'  socket: '%s'  port: %d  %s."},
	ER_FAILED_TO_FIND_LOCALE_NAME : {10932,[]string{"HY000"},"Unknown locale: '%-.64s'."},
	ER_FAILED_TO_FIND_COLLATION_NAME : {10933,[]string{"HY000"},"Unknown collation: '%-.64s'."},
	ER_SERVER_OUT_OF_RESOURCES : {10934,[]string{"HY000"},"Out of memory; check if mysqld or some other process uses all available memory; if not, you may have to use 'ulimit' to allow mysqld to use more memory or you can add more swap space"},
	ER_SERVER_OUTOFMEMORY : {10935,[]string{"HY000"},"Out of memory; restart server and try again (needed %d bytes)"},
	ER_INVALID_COLLATION_FOR_CHARSET : {10936,[]string{"HY000"},"COLLATION '%s' is not valid for CHARACTER SET '%s'"},
	ER_CANT_START_ERROR_LOG_SERVICE : {10937,[]string{"HY000"},"Failed to set %s at or around \"%s\" -- service is valid, but can not be initialized; please check its configuration and make sure it can read its input(s) and write to its output(s)."},
	ER_CREATING_NEW_UUID_FIRST_START : {10938,[]string{"HY000"},"Generating a new UUID: %s."},
	ER_FAILED_TO_GET_ABSOLUTE_PATH : {10939,[]string{"HY000"},"Failed to get absolute path of program executable %s"},
	ER_PERFSCHEMA_COMPONENTS_INFRASTRUCTURE_BOOTSTRAP : {10940,[]string{"HY000"},"Failed to bootstrap performance schema components infrastructure."},
	ER_PERFSCHEMA_COMPONENTS_INFRASTRUCTURE_SHUTDOWN : {10941,[]string{"HY000"},"Failed to deinit performance schema components infrastructure."},
	ER_DUP_FD_OPEN_FAILED : {10942,[]string{"HY000"},"Could not open duplicate fd for %s: %s."},
	ER_SYSTEM_VIEW_INIT_FAILED : {10943,[]string{"HY000"},"System views initialization failed."},
	ER_RESOURCE_GROUP_POST_INIT_FAILED : {10944,[]string{"HY000"},"Resource group post initialization failed."},
	ER_RESOURCE_GROUP_SUBSYSTEM_INIT_FAILED : {10945,[]string{"HY000"},"Resource Group subsystem initialization failed."},
	ER_FAILED_START_MYSQLD_DAEMON : {10946,[]string{"HY000"},"Failed to start mysqld daemon. Check mysqld error log."},
	ER_CANNOT_CHANGE_TO_ROOT_DIR : {10947,[]string{"HY000"},"Cannot change to root directory: %s."},
	ER_PERSISTENT_PRIVILEGES_BOOTSTRAP : {10948,[]string{"HY000"},"Failed to bootstrap persistent privileges."},
	ER_BASEDIR_SET_TO : {10949,[]string{"HY000"},"Basedir set to %s."},
	ER_RPL_FILTER_ADD_WILD_DO_TABLE_FAILED : {10950,[]string{"HY000"},"Could not add wild do table rule '%s'!"},
	ER_RPL_FILTER_ADD_WILD_IGNORE_TABLE_FAILED : {10951,[]string{"HY000"},"Could not add wild ignore table rule '%s'!"},
	ER_PRIVILEGE_SYSTEM_INIT_FAILED : {10952,[]string{"HY000"},"The privilege system failed to initialize correctly. For complete instructions on how to upgrade MySQL to a new version please see the 'Upgrading MySQL' section from the MySQL manual."},
	ER_CANNOT_SET_LOG_ERROR_SERVICES : {10953,[]string{"HY000"},"Cannot set services \"%s\" requested in --log-error-services, using defaults."},
	ER_PERFSCHEMA_TABLES_INIT_FAILED : {10954,[]string{"HY000"},"Performance schema tables initialization failed."},
	ER_TX_EXTRACTION_ALGORITHM_FOR_BINLOG_TX_DEPEDENCY_TRACKING : {10955,[]string{"HY000"},"The transaction_write_set_extraction must be set to %s when binlog_transaction_dependency_tracking is %s."},
	ER_INVALID_REPLICATION_TIMESTAMPS : {10956,[]string{"HY000"},"Invalid replication timestamps: original commit timestamp is more recent than the immediate commit timestamp. This may be an issue if delayed replication is active. Make sure that servers have their clocks set to the correct time. No further message will be emitted until after timestamps become valid again."},
	ER_RPL_TIMESTAMPS_RETURNED_TO_NORMAL : {10957,[]string{"HY000"},"The replication timestamps have returned to normal values."},
	ER_BINLOG_FILE_OPEN_FAILED : {10958,[]string{"HY000"},"%s."},
	ER_BINLOG_EVENT_WRITE_TO_STMT_CACHE_FAILED : {10959,[]string{"HY000"},"Failed to write an incident event into stmt_cache."},
	ER_SLAVE_RELAY_LOG_TRUNCATE_INFO : {10960,[]string{"HY000"},"Relaylog file %s size was %llu, but was truncated at %llu."},
	ER_SLAVE_RELAY_LOG_PURGE_FAILED : {10961,[]string{"HY000"},"Unable to purge relay log files. %s:%s."},
	ER_RPL_SLAVE_FILTER_CREATE_FAILED : {10962,[]string{"HY000"},"Slave: failed in creating filter for channel '%s'."},
	ER_RPL_SLAVE_GLOBAL_FILTERS_COPY_FAILED : {10963,[]string{"HY000"},"Slave: failed in copying the global filters to its own per-channel filters on configuration for channel '%s'."},
	ER_RPL_SLAVE_RESET_FILTER_OPTIONS : {10964,[]string{"HY000"},"There are per-channel replication filter(s) configured for channel '%.192s' which does not exist. The filter(s) have been discarded."},
	ER_MISSING_GRANT_SYSTEM_TABLE : {10965,[]string{"HY000"},"The system table mysql.global_grants is missing. Please perform the MySQL upgrade procedure."},
	ER_MISSING_ACL_SYSTEM_TABLE : {10966,[]string{"HY000"},"ACL table mysql.%.*s missing. Some operations may fail."},
	ER_ANONYMOUS_AUTH_ID_NOT_ALLOWED_IN_MANDATORY_ROLES : {10967,[]string{"HY000"},"Can't set mandatory_role %s@%s: Anonymous authorization IDs are not allowed as roles."},
	ER_UNKNOWN_AUTH_ID_IN_MANDATORY_ROLE : {10968,[]string{"HY000"},"Can't set mandatory_role: There's no such authorization ID %s@%s."},
	ER_WRITE_ROW_TO_PARTITION_FAILED : {10969,[]string{"HY000"},"Table '%-192s' failed to move/insert a row from part %d into part %d: %s."},
	ER_RESOURCE_GROUP_METADATA_UPDATE_SKIPPED : {10970,[]string{"HY000"},"Skipped updating resource group metadata in InnoDB read only mode."},
	ER_FAILED_TO_PERSIST_RESOURCE_GROUP_METADATA : {10971,[]string{"HY000"},"Failed to persist resource group %s to Data Dictionary."},
	ER_FAILED_TO_DESERIALIZE_RESOURCE_GROUP : {10972,[]string{"HY000"},"Failed to deserialize resource group %s."},
	ER_FAILED_TO_UPDATE_RESOURCE_GROUP : {10973,[]string{"HY000"},"Update of resource group %s failed."},
	ER_RESOURCE_GROUP_VALIDATION_FAILED : {10974,[]string{"HY000"},"Validation of resource group %s failed. Resource group is disabled."},
	ER_FAILED_TO_ALLOCATE_MEMORY_FOR_RESOURCE_GROUP : {10975,[]string{"HY000"},"Unable to allocate memory for Resource Group %s."},
	ER_FAILED_TO_ALLOCATE_MEMORY_FOR_RESOURCE_GROUP_HASH : {10976,[]string{"HY000"},"Failed to allocate memory for resource group hash."},
	ER_FAILED_TO_ADD_RESOURCE_GROUP_TO_MAP : {10977,[]string{"HY000"},"Failed to add resource group %s to resource group map."},
	ER_RESOURCE_GROUP_IS_DISABLED : {10978,[]string{"HY000"},"Resource group feature is disabled. (Server is compiled with DISABLE_PSI_THREAD)."},
	ER_FAILED_TO_APPLY_RESOURCE_GROUP_CONTROLLER : {10979,[]string{"HY000"},"Unable to apply resource group controller %s."},
	ER_FAILED_TO_ACQUIRE_LOCK_ON_RESOURCE_GROUP : {10980,[]string{"HY000"},"Unable to acquire lock on the resource group %s. Hint to switch resource group shall be ignored."},
	ER_PFS_NOTIFICATION_FUNCTION_REGISTER_FAILED : {10981,[]string{"HY000"},"PFS %s notification function registration failed."},
	ER_RES_GRP_SET_THR_AFFINITY_FAILED : {10982,[]string{"HY000"},"Unable to bind thread id %llu to cpu id %u (error code %d - %-.192s)."},
	ER_RES_GRP_SET_THR_AFFINITY_TO_CPUS_FAILED : {10983,[]string{"HY000"},"Unable to bind thread id %llu to cpu ids (error code %d - %-.192s)."},
	ER_RES_GRP_THD_UNBIND_FROM_CPU_FAILED : {10984,[]string{"HY000"},"Unbind thread id %llu failed. (error code %d - %-.192s)."},
	ER_RES_GRP_SET_THREAD_PRIORITY_FAILED : {10985,[]string{"HY000"},"Setting thread priority %d to thread id %llu failed. (error code %d - %-.192s)."},
	ER_RES_GRP_FAILED_TO_DETERMINE_NICE_CAPABILITY : {10986,[]string{"HY000"},"Unable to determine CAP_SYS_NICE capability."},
	ER_RES_GRP_FAILED_TO_GET_THREAD_HANDLE : {10987,[]string{"HY000"},"%s failed: Failed to get handle for thread %llu."},
	ER_RES_GRP_GET_THREAD_PRIO_NOT_SUPPORTED : {10988,[]string{"HY000"},"Retrieval of thread priority unsupported on %s."},
	ER_RES_GRP_FAILED_DETERMINE_CPU_COUNT : {10989,[]string{"HY000"},"Unable to determine the number of CPUs."},
	ER_RES_GRP_FEATURE_NOT_AVAILABLE : {10990,[]string{"HY000"},"Resource group feature shall not be available. Incompatible thread handling option."},
	ER_RES_GRP_INVALID_THREAD_PRIORITY : {10991,[]string{"HY000"},"Invalid thread priority %d for a %s resource group. Allowed range is [%d, %d]."},
	ER_RES_GRP_SOLARIS_PROCESSOR_BIND_TO_CPUID_FAILED : {10992,[]string{"HY000"},"bind_to_cpu failed: processor_bind for cpuid %u failed (error code %d - %-.192s)."},
	ER_RES_GRP_SOLARIS_PROCESSOR_BIND_TO_THREAD_FAILED : {10993,[]string{"HY000"},"bind_to_cpu failed: processor_bind for thread %%llx with cpu id %u (error code %d - %-.192s)."},
	ER_RES_GRP_SOLARIS_PROCESSOR_AFFINITY_FAILED : {10994,[]string{"HY000"},"%s failed: processor_affinity failed (error code %d - %-.192s)."},
	ER_DD_UPGRADE_RENAME_IDX_STATS_FILE_FAILED : {10995,[]string{"HY000"},"Error in renaming mysql_index_stats.ibd."},
	ER_DD_UPGRADE_DD_OPEN_FAILED : {10996,[]string{"HY000"},"Error in opening data directory %s."},
	ER_DD_UPGRADE_FAILED_TO_FETCH_TABLESPACES : {10997,[]string{"HY000"},"Error in fetching list of tablespaces."},
	ER_DD_UPGRADE_FAILED_TO_ACQUIRE_TABLESPACE : {10998,[]string{"HY000"},"Error in acquiring Tablespace for SDI insertion %s."},
	ER_DD_UPGRADE_FAILED_TO_RESOLVE_TABLESPACE_ENGINE : {10999,[]string{"HY000"},"Error in resolving Engine name for tablespace %s with engine %s."},
	ER_FAILED_TO_CREATE_SDI_FOR_TABLESPACE : {11000,[]string{"HY000"},"Error in creating SDI for %s tablespace."},
	ER_FAILED_TO_STORE_SDI_FOR_TABLESPACE : {11001,[]string{"HY000"},"Error in storing SDI for %s tablespace."},
	ER_DD_UPGRADE_FAILED_TO_FETCH_TABLES : {11002,[]string{"HY000"},"Error in fetching list of tables."},
	ER_DD_UPGRADE_DD_POPULATED : {11003,[]string{"HY000"},"Finished populating Data Dictionary tables with data."},
	ER_DD_UPGRADE_INFO_FILE_OPEN_FAILED : {11004,[]string{"HY000"},"Could not open the upgrade info file '%s' in the MySQL servers datadir, errno: %d."},
	ER_DD_UPGRADE_INFO_FILE_CLOSE_FAILED : {11005,[]string{"HY000"},"Could not close the upgrade info file '%s' in the MySQL servers datadir, errno: %d."},
	ER_DD_UPGRADE_TABLESPACE_MIGRATION_FAILED : {11006,[]string{"HY000"},"Got error %d from SE while migrating tablespaces."},
	ER_DD_UPGRADE_FAILED_TO_CREATE_TABLE_STATS : {11007,[]string{"HY000"},"Error in creating TABLE statistics entry. Fix statistics data by using ANALYZE command."},
	ER_DD_UPGRADE_TABLE_STATS_MIGRATE_COMPLETED : {11008,[]string{"HY000"},"Finished migrating TABLE statistics data."},
	ER_DD_UPGRADE_FAILED_TO_CREATE_INDEX_STATS : {11009,[]string{"HY000"},"Error in creating Index statistics entry. Fix statistics data by using ANALYZE command."},
	ER_DD_UPGRADE_INDEX_STATS_MIGRATE_COMPLETED : {11010,[]string{"HY000"},"Finished migrating INDEX statistics data."},
	ER_DD_UPGRADE_FAILED_FIND_VALID_DATA_DIR : {11011,[]string{"HY000"},"Failed to find valid data directory."},
	ER_DD_UPGRADE_START : {11012,[]string{"HY000"},"Starting upgrade of data directory."},
	ER_DD_UPGRADE_FAILED_INIT_DD_SE : {11013,[]string{"HY000"},"Failed to initialize DD Storage Engine."},
	ER_DD_UPGRADE_FOUND_PARTIALLY_UPGRADED_DD_ABORT : {11014,[]string{"HY000"},"Found partially upgraded DD. Aborting upgrade and deleting all DD tables. Start the upgrade process again."},
	ER_DD_UPGRADE_FOUND_PARTIALLY_UPGRADED_DD_CONTINUE : {11015,[]string{"HY000"},"Found partially upgraded DD. Upgrade will continue and start the server."},
	ER_DD_UPGRADE_SE_LOGS_FAILED : {11016,[]string{"HY000"},"Error in upgrading engine logs."},
	ER_DD_UPGRADE_SDI_INFO_UPDATE_FAILED : {11017,[]string{"HY000"},"Error in updating SDI information."},
	ER_SKIP_UPDATING_METADATA_IN_SE_RO_MODE : {11018,[]string{"HY000"},"Skip updating %s metadata in InnoDB read-only mode."},
	ER_CREATED_SYSTEM_WITH_VERSION : {11019,[]string{"HY000"},"Created system views with I_S version %d."},
	ER_UNKNOWN_ERROR_DETECTED_IN_SE : {11020,[]string{"HY000"},"Unknown error detected %d in handler."},
	ER_READ_LOG_EVENT_FAILED : {11021,[]string{"HY000"},"Error in Log_event::read_log_event(): '%s', data_len: %lu, event_type: %d."},
	ER_ROW_DATA_TOO_BIG_TO_WRITE_IN_BINLOG : {11022,[]string{"HY000"},"The row data is greater than 4GB, which is too big to write to the binary log."},
	ER_FAILED_TO_CONSTRUCT_DROP_EVENT_QUERY : {11023,[]string{"HY000"},"Unable to construct DROP EVENT SQL query string."},
	ER_FAILED_TO_BINLOG_DROP_EVENT : {11024,[]string{"HY000"},"Unable to binlog drop event %s.%s."},
	ER_FAILED_TO_START_SLAVE_THREAD : {11025,[]string{"HY000"},"Failed to start slave threads for channel '%s'."},
	ER_RPL_IO_THREAD_KILLED : {11026,[]string{"HY000"},"%s%s."},
	ER_SLAVE_RECONNECT_FAILED : {11027,[]string{"HY000"},"Failed registering on master, reconnecting to try again, log '%s' at position %s. %s."},
	ER_SLAVE_KILLED_AFTER_RECONNECT : {11028,[]string{"HY000"},"Slave I/O thread killed during or after reconnect."},
	ER_SLAVE_NOT_STARTED_ON_SOME_CHANNELS : {11029,[]string{"HY000"},"Some of the channels are not created/initialized properly. Check for additional messages above. You will not be able to start replication on those channels until the issue is resolved and the server restarted."},
	ER_FAILED_TO_ADD_RPL_FILTER : {11030,[]string{"HY000"},"Failed to add a replication filter into filter map for channel '%.192s'."},
	ER_PER_CHANNEL_RPL_FILTER_CONF_FOR_GRP_RPL : {11031,[]string{"HY000"},"There are per-channel replication filter(s) configured for group replication channel '%.192s' which is disallowed. The filter(s) have been discarded."},
	ER_RPL_FILTERS_NOT_ATTACHED_TO_CHANNEL : {11032,[]string{"HY000"},"There are per-channel replication filter(s) configured for channel '%.192s' which does not exist. The filter(s) have been discarded."},
	ER_FAILED_TO_BUILD_DO_AND_IGNORE_TABLE_HASHES : {11033,[]string{"HY000"},"An error occurred while building do_table and ignore_table rules to hashes for per-channel filter."},
	ER_CLONE_PLUGIN_NOT_LOADED_TRACE : {11034,[]string{"HY000"},"Clone plugin cannot be loaded."},
	ER_CLONE_HANDLER_EXIST_TRACE : {11035,[]string{"HY000"},"Clone Handler exists."},
	ER_CLONE_CREATE_HANDLER_FAIL_TRACE : {11036,[]string{"HY000"},"Could not create Clone Handler."},
	ER_CYCLE_TIMER_IS_NOT_AVAILABLE : {11037,[]string{"HY000"},"The CYCLE timer is not available. WAIT events in the performance_schema will not be timed."},
	ER_NANOSECOND_TIMER_IS_NOT_AVAILABLE : {11038,[]string{"HY000"},"The NANOSECOND timer is not available. IDLE/STAGE/STATEMENT/TRANSACTION events in the performance_schema will not be timed."},
	ER_MICROSECOND_TIMER_IS_NOT_AVAILABLE : {11039,[]string{"HY000"},"The MICROSECOND timer is not available. IDLE/STAGE/STATEMENT/TRANSACTION events in the performance_schema will not be timed."},
	ER_PFS_MALLOC_ARRAY_OVERFLOW : {11040,[]string{"HY000"},"Failed to allocate memory for %zu chunks each of size %zu for buffer '%s' due to overflow."},
	ER_PFS_MALLOC_ARRAY_OOM : {11041,[]string{"HY000"},"Failed to allocate %zu bytes for buffer '%s' due to out-of-memory."},
	ER_INNODB_FAILED_TO_FIND_IDX_WITH_KEY_NO : {11042,[]string{"HY000"},"InnoDB could not find index %s key no %u for table %s through its index translation table."},
	ER_INNODB_FAILED_TO_FIND_IDX : {11043,[]string{"HY000"},"Cannot find index %s in InnoDB index translation table."},
	ER_INNODB_FAILED_TO_FIND_IDX_FROM_DICT_CACHE : {11044,[]string{"HY000"},"InnoDB could not find key no %u with name %s from dict cache for table %s."},
	ER_INNODB_ACTIVE_INDEX_CHANGE_FAILED : {11045,[]string{"HY000"},"InnoDB: change_active_index(%u) failed."},
	ER_INNODB_DIFF_IN_REF_LEN : {11046,[]string{"HY000"},"Stored ref len is %lu, but table ref len is %lu."},
	ER_WRONG_TYPE_FOR_COLUMN_PREFIX_IDX_FLD : {11047,[]string{"HY000"},"MySQL is trying to create a column prefix index field, on an inappropriate data type. Table name %s, column name %s."},
	ER_INNODB_CANNOT_CREATE_TABLE : {11048,[]string{"HY000"},"Cannot create table %s."},
	ER_INNODB_INTERNAL_INDEX : {11049,[]string{"HY000"},"Found index %s in InnoDB index list but not its MySQL index number. It could be an InnoDB internal index."},
	ER_INNODB_IDX_CNT_MORE_THAN_DEFINED_IN_MYSQL : {11050,[]string{"HY000"},"InnoDB: Table %s contains %lu indexes inside InnoDB, which is different from the number of indexes %u defined in MySQL."},
	ER_INNODB_IDX_CNT_FEWER_THAN_DEFINED_IN_MYSQL : {11051,[]string{"HY000"},"Table %s contains fewer indexes inside InnoDB than are defined in the MySQL. Have you mixed up with data dictionary from different installation?"},
	ER_INNODB_IDX_COLUMN_CNT_DIFF : {11052,[]string{"HY000"},"Index %s of %s has %lu columns unique inside InnoDB, but MySQL is asking statistics for %lu columns. Have you mixed data dictionary from different installation?"},
	ER_INNODB_USE_MONITOR_GROUP_NAME : {11053,[]string{"HY000"},"Monitor counter '%s' cannot be turned on/off individually. Please use its module name to turn on/off the counters in the module as a group."},
	ER_INNODB_MONITOR_DEFAULT_VALUE_NOT_DEFINED : {11054,[]string{"HY000"},"Default value is not defined for this set option. Please specify correct counter or module name."},
	ER_INNODB_MONITOR_IS_ENABLED : {11055,[]string{"HY000"},"InnoDB: Monitor %s is already enabled."},
	ER_INNODB_INVALID_MONITOR_COUNTER_NAME : {11056,[]string{"HY000"},"Invalid monitor counter : %s."},
	ER_WIN_LOAD_LIBRARY_FAILED : {11057,[]string{"HY000"},"LoadLibrary(\"%s\") failed: GetLastError returns %lu."},
	ER_PARTITION_HANDLER_ADMIN_MSG : {11058,[]string{"HY000"},"%s."},
	ER_RPL_RLI_INIT_INFO_MSG : {11059,[]string{"HY000"},"%s."},
	ER_DD_UPGRADE_TABLE_INTACT_ERROR : {11060,[]string{"HY000"},"%s."},
	ER_SERVER_INIT_COMPILED_IN_COMMANDS : {11061,[]string{"HY000"},"%s."},
	ER_MYISAM_CHECK_METHOD_ERROR : {11062,[]string{"HY000"},"%s."},
	ER_MYISAM_CRASHED_ERROR : {11063,[]string{"HY000"},"%s."},
	ER_WAITPID_FAILED : {11064,[]string{"HY000"},"Unable to wait for process %lld."},
	ER_FAILED_TO_FIND_MYSQLD_STATUS : {11065,[]string{"HY000"},"Unable to determine if daemon is running: %s (rc=%d)."},
	ER_INNODB_ERROR_LOGGER_MSG : {11066,[]string{"HY000"},"%s"},
	ER_INNODB_ERROR_LOGGER_FATAL_MSG : {11067,[]string{"HY000"},"[FATAL] InnoDB: %s"},
	ER_DEPRECATED_SYNTAX_WITH_REPLACEMENT : {11068,[]string{"HY000"},"The syntax '%s' is deprecated and will be removed in a future release. Please use %s instead."},
	ER_DEPRECATED_SYNTAX_NO_REPLACEMENT : {11069,[]string{"HY000"},"The syntax '%s' is deprecated and will be removed in a future release."},
	ER_DEPRECATE_MSG_NO_REPLACEMENT : {11070,[]string{"HY000"},"'%s' is deprecated and will be removed in a future release."},
	ER_LOG_PRINTF_MSG : {11071,[]string{"HY000"},"%s"},
	ER_BINLOG_LOGGING_NOT_POSSIBLE : {11072,[]string{"HY000"},"Binary logging not possible. Message: %s."},
	ER_FAILED_TO_SET_PERSISTED_OPTIONS : {11073,[]string{"HY000"},"Failed to set persisted options."},
	ER_COMPONENTS_FAILED_TO_ACQUIRE_SERVICE_IMPLEMENTATION : {11074,[]string{"HY000"},"Cannot acquire specified service implementation: '%.192s'."},
	ER_RES_GRP_INVALID_VCPU_RANGE : {11075,[]string{"HY000"},"Invalid VCPU range %u-%u."},
	ER_RES_GRP_INVALID_VCPU_ID : {11076,[]string{"HY000"},"Invalid cpu id %u."},
	ER_ERROR_DURING_FLUSH_LOG_COMMIT_PHASE : {11077,[]string{"HY000"},"Got error %d during FLUSH_LOGS."},
	ER_DROP_DATABASE_FAILED_RMDIR_MANUALLY : {11078,[]string{"HY000"},"Problem while dropping database. Can't remove database directory (%s). Please remove it manually."},
	ER_EXPIRE_LOGS_DAYS_IGNORED : {11079,[]string{"HY000"},"The option expire_logs_days cannot be used together with option binlog_expire_logs_seconds. Therefore, value of expire_logs_days is ignored."},
	ER_BINLOG_MALFORMED_OR_OLD_RELAY_LOG : {11080,[]string{"HY000"},"malformed or very old relay log which does not have FormatDescriptor."},
	ER_DD_UPGRADE_VIEW_COLUMN_NAME_TOO_LONG : {11081,[]string{"HY000"},"Upgrade of view '%s.%s' failed. Re-create the view with the explicit column name lesser than 64 characters."},
	ER_TABLE_NEEDS_DUMP_UPGRADE : {11082,[]string{"HY000"},"Table upgrade required for `%-.64s`.`%-.64s`. Please dump/reload table to fix it!"},
	ER_DD_UPGRADE_FAILED_TO_UPDATE_VER_NO_IN_TABLESPACE : {11083,[]string{"HY000"},"Error in updating version number in %s tablespace."},
	ER_KEYRING_MIGRATION_FAILED : {11084,[]string{"HY000"},"Keyring migration failed."},
	ER_KEYRING_MIGRATION_SUCCESSFUL : {11085,[]string{"HY000"},"Keyring migration successful."},
	ER_RESTART_RECEIVED_INFO : {11086,[]string{"HY000"},"Received RESTART from user %s.  Restarting mysqld (Version: %s)."},
	ER_LCTN_CHANGED : {11087,[]string{"HY000"},"Different lower_case_table_names settings for server ('%u') and data dictionary ('%u')."},
	ER_DD_INITIALIZE : {11088,[]string{"HY000"},"Data dictionary initializing version '%u'."},
	ER_DD_RESTART : {11089,[]string{"HY000"},"Data dictionary restarting version '%u'."},
	ER_DD_UPGRADE : {11090,[]string{"HY000"},"Data dictionary upgrading from version '%u' to '%u'."},
	ER_DD_UPGRADE_OFF : {11091,[]string{"HY000"},"Data dictionary upgrade prohibited by the command line option '--no_dd_upgrade'."},
	ER_DD_UPGRADE_VERSION_NOT_SUPPORTED : {11092,[]string{"HY000"},"Upgrading the data dictionary from dictionary version '%u' is not supported."},
	ER_DD_UPGRADE_SCHEMA_UNAVAILABLE : {11093,[]string{"HY000"},"Upgrading the data dictionary failed, temporary schema name '%-.192s' not available."},
	ER_DD_MINOR_DOWNGRADE : {11094,[]string{"HY000"},"Data dictionary minor downgrade from version '%u' to '%u'."},
	ER_DD_MINOR_DOWNGRADE_VERSION_NOT_SUPPORTED : {11095,[]string{"HY000"},"Minor downgrade of the Data dictionary from dictionary version '%u' is not supported."},
	ER_DD_NO_VERSION_FOUND : {11096,[]string{"HY000"},"No data dictionary version number found."},
	ER_THREAD_POOL_NOT_SUPPORTED_ON_PLATFORM : {11097,[]string{"HY000"},"Thread pool not supported, requires a minimum of %s."},
	ER_THREAD_POOL_SIZE_TOO_LOW : {11098,[]string{"HY000"},"thread_pool_size=0 means thread pool disabled, Allowed range of thread_pool_size is %d-%d."},
	ER_THREAD_POOL_SIZE_TOO_HIGH : {11099,[]string{"HY000"},"thread_pool_size=%lu is too high, %d is maximum, thread pool is disabled. Allowed range of thread_pool_size is %d-%d."},
	ER_THREAD_POOL_ALGORITHM_INVALID : {11100,[]string{"HY000"},"thread_pool_algorithm can be set to 0 and 1, 0 indicates the default low concurrency algorithm, 1 means a high concurrency algorithm."},
	ER_THREAD_POOL_INVALID_STALL_LIMIT : {11101,[]string{"HY000"},"thread_pool_stall_limit can be %d at minimum and %d at maximum, smaller values would render the thread pool fairly useless and higher values could make it possible to have undetected deadlock issues in the MySQL Server."},
	ER_THREAD_POOL_INVALID_PRIO_KICKUP_TIMER : {11102,[]string{"HY000"},"Invalid value of thread_pool_prio_kickup_timer specified. Value of thread_pool_prio_kickup_timer should be in range 0-4294967294."},
	ER_THREAD_POOL_MAX_UNUSED_THREADS_INVALID : {11103,[]string{"HY000"},"thread_pool_max_unused_threads cannot be set higher than %d."},
	ER_THREAD_POOL_CON_HANDLER_INIT_FAILED : {11104,[]string{"HY000"},"Failed to instantiate the connection handler object."},
	ER_THREAD_POOL_INIT_FAILED : {11105,[]string{"HY000"},"Failed to initialize thread pool plugin."},
	ER_THREAD_POOL_PLUGIN_STARTED : {11106,[]string{"HY000"},"Thread pool plugin started successfully with parameters: thread_pool_size = %lu, thread_pool_algorithm = %s, thread_pool_stall_limit = %u, thread_pool_prio_kickup_timer = %u, thread_pool_max_unused_threads = %u, thread_pool_high_priority_connection = %d."},
	ER_THREAD_POOL_CANNOT_SET_THREAD_SPECIFIC_DATA : {11107,[]string{"HY000"},"Can't setup connection teardown thread-specific data."},
	ER_THREAD_POOL_FAILED_TO_CREATE_CONNECT_HANDLER_THD : {11108,[]string{"HY000"},"Creation of connect handler thread failed."},
	ER_THREAD_POOL_FAILED_TO_CREATE_THD_AND_AUTH_CONN : {11109,[]string{"HY000"},"Failed to create thd and authenticate connection."},
	ER_THREAD_POOL_FAILED_PROCESS_CONNECT_EVENT : {11110,[]string{"HY000"},"Failed to process connection event."},
	ER_THREAD_POOL_FAILED_TO_CREATE_POOL : {11111,[]string{"HY000"},"Can't create pool thread (error %d, errno: %d)."},
	ER_THREAD_POOL_RATE_LIMITED_ERROR_MSGS : {11112,[]string{"HY000"},"%.*s."},
	ER_TRHEAD_POOL_LOW_LEVEL_INIT_FAILED : {11113,[]string{"HY000"},"tp_group_low_level_init() failed."},
	ER_THREAD_POOL_LOW_LEVEL_REARM_FAILED : {11114,[]string{"HY000"},"Rearm failed even after 30 seconds, can't continue without notify socket."},
	ER_THREAD_POOL_BUFFER_TOO_SMALL : {11115,[]string{"HY000"},"%s: %s buffer is too small"},
	ER_MECAB_NOT_SUPPORTED : {11116,[]string{"HY000"},"Mecab v%s is not supported, the lowest version supported is v%s."},
	ER_MECAB_NOT_VERIFIED : {11117,[]string{"HY000"},"Mecab v%s is not verified, the highest version supported is v%s."},
	ER_MECAB_CREATING_MODEL : {11118,[]string{"HY000"},"Mecab: Trying createModel(%s)."},
	ER_MECAB_FAILED_TO_CREATE_MODEL : {11119,[]string{"HY000"},"Mecab: createModel() failed: %s."},
	ER_MECAB_FAILED_TO_CREATE_TRIGGER : {11120,[]string{"HY000"},"Mecab: createTagger() failed: %s."},
	ER_MECAB_UNSUPPORTED_CHARSET : {11121,[]string{"HY000"},"Mecab: Unsupported dictionary charset %s."},
	ER_MECAB_CHARSET_LOADED : {11122,[]string{"HY000"},"Mecab: Loaded dictionary charset is %s."},
	ER_MECAB_PARSE_FAILED : {11123,[]string{"HY000"},"Mecab: parse() failed: %s."},
	ER_MECAB_OOM_WHILE_PARSING_TEXT : {11124,[]string{"HY000"},"Mecab: parse() failed: out of memory."},
	ER_MECAB_CREATE_LATTICE_FAILED : {11125,[]string{"HY000"},"Mecab: createLattice() failed: %s."},
	ER_SEMISYNC_TRACE_ENTER_FUNC : {11126,[]string{"HY000"},"---> %s enter."},
	ER_SEMISYNC_TRACE_EXIT_WITH_INT_EXIT_CODE : {11127,[]string{"HY000"},"<--- %s exit (%d)."},
	ER_SEMISYNC_TRACE_EXIT_WITH_BOOL_EXIT_CODE : {11128,[]string{"HY000"},"<--- %s exit (%s)."},
	ER_SEMISYNC_TRACE_EXIT : {11129,[]string{"HY000"},"<--- %s exit."},
	ER_SEMISYNC_RPL_INIT_FOR_TRX : {11130,[]string{"HY000"},"Semi-sync replication initialized for transactions."},
	ER_SEMISYNC_FAILED_TO_ALLOCATE_TRX_NODE : {11131,[]string{"HY000"},"%s: transaction node allocation failed for: (%s, %lu)."},
	ER_SEMISYNC_BINLOG_WRITE_OUT_OF_ORDER : {11132,[]string{"HY000"},"%s: binlog write out-of-order, tail (%s, %lu), new node (%s, %lu)."},
	ER_SEMISYNC_INSERT_LOG_INFO_IN_ENTRY : {11133,[]string{"HY000"},"%s: insert (%s, %lu) in entry(%u)."},
	ER_SEMISYNC_PROBE_LOG_INFO_IN_ENTRY : {11134,[]string{"HY000"},"%s: probe (%s, %lu) in entry(%u)."},
	ER_SEMISYNC_CLEARED_ALL_ACTIVE_TRANSACTION_NODES : {11135,[]string{"HY000"},"%s: cleared all nodes."},
	ER_SEMISYNC_CLEARED_ACTIVE_TRANSACTION_TILL_POS : {11136,[]string{"HY000"},"%s: cleared %d nodes back until pos (%s, %lu)."},
	ER_SEMISYNC_REPLY_MAGIC_NO_ERROR : {11137,[]string{"HY000"},"Read semi-sync reply magic number error."},
	ER_SEMISYNC_REPLY_PKT_LENGTH_TOO_SMALL : {11138,[]string{"HY000"},"Read semi-sync reply length error: packet is too small."},
	ER_SEMISYNC_REPLY_BINLOG_FILE_TOO_LARGE : {11139,[]string{"HY000"},"Read semi-sync reply binlog file length too large."},
	ER_SEMISYNC_SERVER_REPLY : {11140,[]string{"HY000"},"%s: Got reply(%s, %lu) from server %u."},
	ER_SEMISYNC_FUNCTION_CALLED_TWICE : {11141,[]string{"HY000"},"%s called twice."},
	ER_SEMISYNC_RPL_ENABLED_ON_MASTER : {11142,[]string{"HY000"},"Semi-sync replication enabled on the master."},
	ER_SEMISYNC_MASTER_OOM : {11143,[]string{"HY000"},"Cannot allocate memory to enable semi-sync on the master."},
	ER_SEMISYNC_DISABLED_ON_MASTER : {11144,[]string{"HY000"},"Semi-sync replication disabled on the master."},
	ER_SEMISYNC_FORCED_SHUTDOWN : {11145,[]string{"HY000"},"SEMISYNC: Forced shutdown. Some updates might not be replicated."},
	ER_SEMISYNC_MASTER_GOT_REPLY_AT_POS : {11146,[]string{"HY000"},"%s: Got reply at (%s, %lu)."},
	ER_SEMISYNC_MASTER_SIGNAL_ALL_WAITING_THREADS : {11147,[]string{"HY000"},"%s: signal all waiting threads."},
	ER_SEMISYNC_MASTER_TRX_WAIT_POS : {11148,[]string{"HY000"},"%s: wait pos (%s, %lu), repl(%d)."},
	ER_SEMISYNC_BINLOG_REPLY_IS_AHEAD : {11149,[]string{"HY000"},"%s: Binlog reply is ahead (%s, %lu)."},
	ER_SEMISYNC_MOVE_BACK_WAIT_POS : {11150,[]string{"HY000"},"%s: move back wait position (%s, %lu)."},
	ER_SEMISYNC_INIT_WAIT_POS : {11151,[]string{"HY000"},"%s: init wait position (%s, %lu)."},
	ER_SEMISYNC_WAIT_TIME_FOR_BINLOG_SENT : {11152,[]string{"HY000"},"%s: wait %lu ms for binlog sent (%s, %lu)."},
	ER_SEMISYNC_WAIT_FOR_BINLOG_TIMEDOUT : {11153,[]string{"HY000"},"Timeout waiting for reply of binlog (file: %s, pos: %lu), semi-sync up to file %s, position %lu."},
	ER_SEMISYNC_WAIT_TIME_ASSESSMENT_FOR_COMMIT_TRX_FAILED : {11154,[]string{"HY000"},"Assessment of waiting time for commitTrx failed at wait position (%s, %lu)."},
	ER_SEMISYNC_RPL_SWITCHED_OFF : {11155,[]string{"HY000"},"Semi-sync replication switched OFF."},
	ER_SEMISYNC_RPL_SWITCHED_ON : {11156,[]string{"HY000"},"Semi-sync replication switched ON at (%s, %lu)."},
	ER_SEMISYNC_NO_SPACE_IN_THE_PKT : {11157,[]string{"HY000"},"No enough space in the packet for semi-sync extra header, semi-sync replication disabled."},
	ER_SEMISYNC_SYNC_HEADER_UPDATE_INFO : {11158,[]string{"HY000"},"%s: server(%d), (%s, %lu) sync(%d), repl(%d)."},
	ER_SEMISYNC_FAILED_TO_INSERT_TRX_NODE : {11159,[]string{"HY000"},"Semi-sync failed to insert tranx_node for binlog file: %s, position: %lu."},
	ER_SEMISYNC_TRX_SKIPPED_AT_POS : {11160,[]string{"HY000"},"%s: Transaction skipped at (%s, %lu)."},
	ER_SEMISYNC_MASTER_FAILED_ON_NET_FLUSH : {11161,[]string{"HY000"},"Semi-sync master failed on net_flush() before waiting for slave reply."},
	ER_SEMISYNC_RECEIVED_ACK_IS_SMALLER : {11162,[]string{"HY000"},"The received ack is smaller than m_greatest_ack."},
	ER_SEMISYNC_ADD_ACK_TO_SLOT : {11163,[]string{"HY000"},"Add the ack into slot %u."},
	ER_SEMISYNC_UPDATE_EXISTING_SLAVE_ACK : {11164,[]string{"HY000"},"Update an exsiting ack in slot %u."},
	ER_SEMISYNC_FAILED_TO_START_ACK_RECEIVER_THD : {11165,[]string{"HY000"},"Failed to start semi-sync ACK receiver thread,  could not create thread(errno:%d)."},
	ER_SEMISYNC_STARTING_ACK_RECEIVER_THD : {11166,[]string{"HY000"},"Starting ack receiver thread."},
	ER_SEMISYNC_FAILED_TO_WAIT_ON_DUMP_SOCKET : {11167,[]string{"HY000"},"Failed to wait on semi-sync dump sockets, error: errno=%d."},
	ER_SEMISYNC_STOPPING_ACK_RECEIVER_THREAD : {11168,[]string{"HY000"},"Stopping ack receiver thread."},
	ER_SEMISYNC_FAILED_REGISTER_SLAVE_TO_RECEIVER : {11169,[]string{"HY000"},"Failed to register slave to semi-sync ACK receiver thread."},
	ER_SEMISYNC_START_BINLOG_DUMP_TO_SLAVE : {11170,[]string{"HY000"},"Start %s binlog_dump to slave (server_id: %d), pos(%s, %lu)."},
	ER_SEMISYNC_STOP_BINLOG_DUMP_TO_SLAVE : {11171,[]string{"HY000"},"Stop %s binlog_dump to slave (server_id: %d)."},
	ER_SEMISYNC_UNREGISTER_TRX_OBSERVER_FAILED : {11172,[]string{"HY000"},"unregister_trans_observer failed."},
	ER_SEMISYNC_UNREGISTER_BINLOG_STORAGE_OBSERVER_FAILED : {11173,[]string{"HY000"},"unregister_binlog_storage_observer failed."},
	ER_SEMISYNC_UNREGISTER_BINLOG_TRANSMIT_OBSERVER_FAILED : {11174,[]string{"HY000"},"unregister_binlog_transmit_observer failed."},
	ER_SEMISYNC_UNREGISTERED_REPLICATOR : {11175,[]string{"HY000"},"unregister_replicator OK."},
	ER_SEMISYNC_SOCKET_FD_TOO_LARGE : {11176,[]string{"HY000"},"Semisync slave socket fd is %u. select() cannot handle if the socket fd is bigger than %u (FD_SETSIZE)."},
	ER_SEMISYNC_SLAVE_REPLY : {11177,[]string{"HY000"},"%s: reply - %d."},
	ER_SEMISYNC_MISSING_MAGIC_NO_FOR_SEMISYNC_PKT : {11178,[]string{"HY000"},"Missing magic number for semi-sync packet, packet len: %lu."},
	ER_SEMISYNC_SLAVE_START : {11179,[]string{"HY000"},"Slave I/O thread: Start %s replication to master '%s@%s:%d' in log '%s' at position %lu."},
	ER_SEMISYNC_SLAVE_REPLY_WITH_BINLOG_INFO : {11180,[]string{"HY000"},"%s: reply (%s, %lu)."},
	ER_SEMISYNC_SLAVE_NET_FLUSH_REPLY_FAILED : {11181,[]string{"HY000"},"Semi-sync slave net_flush() reply failed."},
	ER_SEMISYNC_SLAVE_SEND_REPLY_FAILED : {11182,[]string{"HY000"},"Semi-sync slave send reply failed: %s (%d)."},
	ER_SEMISYNC_EXECUTION_FAILED_ON_MASTER : {11183,[]string{"HY000"},"Execution failed on master: %s; error %d"},
	ER_SEMISYNC_NOT_SUPPORTED_BY_MASTER : {11184,[]string{"HY000"},"Master server does not support semi-sync, fallback to asynchronous replication"},
	ER_SEMISYNC_SLAVE_SET_FAILED : {11185,[]string{"HY000"},"Set 'rpl_semi_sync_slave=1' on master failed"},
	ER_SEMISYNC_FAILED_TO_STOP_ACK_RECEIVER_THD : {11186,[]string{"HY000"},"Failed to stop ack receiver thread on my_thread_join, errno(%d)."},
	ER_FIREWALL_FAILED_TO_READ_FIREWALL_TABLES : {11187,[]string{"HY000"},"Failed to read the firewall tables"},
	ER_FIREWALL_FAILED_TO_REG_DYNAMIC_PRIVILEGES : {11188,[]string{"HY000"},"Failed to register dynamic privileges"},
	ER_FIREWALL_RECORDING_STMT_WAS_TRUNCATED : {11189,[]string{"HY000"},"Statement was truncated and not recorded: %s"},
	ER_FIREWALL_RECORDING_STMT_WITHOUT_TEXT : {11190,[]string{"HY000"},"Statement with no text was not recorded"},
	ER_FIREWALL_SUSPICIOUS_STMT : {11191,[]string{"HY000"},"SUSPICIOUS STATEMENT from '%s'. Reason: %s Statement: %s"},
	ER_FIREWALL_ACCESS_DENIED : {11192,[]string{"HY000"},"ACCESS DENIED for '%s'. Reason: %s Statement: %s"},
	ER_FIREWALL_SKIPPED_UNKNOWN_USER_MODE : {11193,[]string{"HY000"},"Skipped unknown user mode '%s'"},
	ER_FIREWALL_RELOADING_CACHE : {11194,[]string{"HY000"},"Reloading cache from disk"},
	ER_FIREWALL_RESET_FOR_USER : {11195,[]string{"HY000"},"FIREWALL RESET for '%s'"},
	ER_FIREWALL_STATUS_FLUSHED : {11196,[]string{"HY000"},"Counters are reset to zero"},
	ER_KEYRING_LOGGER_ERROR_MSG : {11197,[]string{"HY000"},"%s"},
	ER_AUDIT_LOG_FILTER_IS_NOT_INSTALLED : {11198,[]string{"HY000"},"Audit Log plugin supports a filtering, which has not been installed yet. Audit Log plugin will run in the legacy mode, which will be disabled in the next release."},
	ER_AUDIT_LOG_SWITCHING_TO_INCLUDE_LIST : {11199,[]string{"HY000"},"Previously exclude list is used, now we start using include list, exclude list is set to NULL."},
	ER_AUDIT_LOG_CANNOT_SET_LOG_POLICY_WITH_OTHER_POLICIES : {11200,[]string{"HY000"},"Cannot set audit_log_policy simultaneously with either audit_log_connection_policy or  audit_log_statement_policy, setting audit_log_connection_policy and audit_log_statement_policy based on audit_log_policy."},
	ER_AUDIT_LOG_ONLY_INCLUDE_LIST_USED : {11201,[]string{"HY000"},"Both include and exclude lists provided, include list is preferred, exclude list is set to NULL."},
	ER_AUDIT_LOG_INDEX_MAP_CANNOT_ACCESS_DIR : {11202,[]string{"HY000"},"Could not access '%s' directory."},
	ER_AUDIT_LOG_WRITER_RENAME_FILE_FAILED : {11203,[]string{"HY000"},"Could not rename file from '%s' to '%s'."},
	ER_AUDIT_LOG_WRITER_DEST_FILE_ALREADY_EXISTS : {11204,[]string{"HY000"},"File '%s' should not exist. It may be incomplete. The server crashed."},
	ER_AUDIT_LOG_WRITER_RENAME_FILE_FAILED_REMOVE_FILE_MANUALLY : {11205,[]string{"HY000"},"Could not rename file from '%s' to '%s'. Remove the file manually."},
	ER_AUDIT_LOG_WRITER_INCOMPLETE_FILE_RENAMED : {11206,[]string{"HY000"},"Incomplete file renamed from '%s' to '%s'."},
	ER_AUDIT_LOG_WRITER_FAILED_TO_WRITE_TO_FILE : {11207,[]string{"HY000"},"Error writing file '%-.200s' (errno: %d - %s)."},
	ER_AUDIT_LOG_EC_WRITER_FAILED_TO_INIT_ENCRYPTION : {11208,[]string{"HY000"},"Could not initialize audit log file encryption."},
	ER_AUDIT_LOG_EC_WRITER_FAILED_TO_INIT_COMPRESSION : {11209,[]string{"HY000"},"Could not initialize audit log file compression."},
	ER_AUDIT_LOG_EC_WRITER_FAILED_TO_CREATE_FILE : {11210,[]string{"HY000"},"Could not create '%s' file for audit logging."},
	ER_AUDIT_LOG_RENAME_LOG_FILE_BEFORE_FLUSH : {11211,[]string{"HY000"},"Audit log file (%s) must be manually renamed before audit_log_flush is set to true."},
	ER_AUDIT_LOG_FILTER_RESULT_MSG : {11212,[]string{"HY000"},"%s"},
	ER_AUDIT_LOG_JSON_READER_FAILED_TO_PARSE : {11213,[]string{"HY000"},"Error parsing JSON event. Event not accessible."},
	ER_AUDIT_LOG_JSON_READER_BUF_TOO_SMALL : {11214,[]string{"HY000"},"Buffer is too small to hold JSON event. Number of events skipped: %zu."},
	ER_AUDIT_LOG_JSON_READER_FAILED_TO_OPEN_FILE : {11215,[]string{"HY000"},"Could not open JSON file for reading. Reading next file if exists."},
	ER_AUDIT_LOG_JSON_READER_FILE_PARSING_ERROR : {11216,[]string{"HY000"},"JSON file parsing error. Reading next file if exists"},
	ER_AUDIT_LOG_FILTER_INVALID_COLUMN_COUNT : {11217,[]string{"HY000"},"Invalid column count in the '%s.%s' table."},
	ER_AUDIT_LOG_FILTER_INVALID_COLUMN_DEFINITION : {11218,[]string{"HY000"},"Invalid column definition of the '%s.%s' table."},
	ER_AUDIT_LOG_FILTER_FAILED_TO_STORE_TABLE_FLDS : {11219,[]string{"HY000"},"Could not store field of the %s table."},
	ER_AUDIT_LOG_FILTER_FAILED_TO_UPDATE_TABLE : {11220,[]string{"HY000"},"Could not update %s table."},
	ER_AUDIT_LOG_FILTER_FAILED_TO_INSERT_INTO_TABLE : {11221,[]string{"HY000"},"Could not insert into %s table."},
	ER_AUDIT_LOG_FILTER_FAILED_TO_DELETE_FROM_TABLE : {11222,[]string{"HY000"},"Could not delete from %s table."},
	ER_AUDIT_LOG_FILTER_FAILED_TO_INIT_TABLE_FOR_READ : {11223,[]string{"HY000"},"Could not initialize %s table for reading."},
	ER_AUDIT_LOG_FILTER_FAILED_TO_READ_TABLE : {11224,[]string{"HY000"},"Could not read %s table."},
	ER_AUDIT_LOG_FILTER_FAILED_TO_CLOSE_TABLE_AFTER_READING : {11225,[]string{"HY000"},"Could not close %s table reading."},
	ER_AUDIT_LOG_FILTER_USER_AND_HOST_CANNOT_BE_EMPTY : {11226,[]string{"HY000"},"Both user and host columns of %s table cannot be empty."},
	ER_AUDIT_LOG_FILTER_FLD_FILTERNAME_CANNOT_BE_EMPTY : {11227,[]string{"HY000"},"Filtername column of %s table cannot be empty."},
	ER_VALIDATE_PWD_DICT_FILE_NOT_SPECIFIED : {11228,[]string{"HY000"},"Dictionary file not specified"},
	ER_VALIDATE_PWD_DICT_FILE_NOT_LOADED : {11229,[]string{"HY000"},"Dictionary file not loaded"},
	ER_VALIDATE_PWD_DICT_FILE_TOO_BIG : {11230,[]string{"HY000"},"Dictionary file size exceeded MAX_DICTIONARY_FILE_LENGTH, not loaded"},
	ER_VALIDATE_PWD_FAILED_TO_READ_DICT_FILE : {11231,[]string{"HY000"},"Exception while reading the dictionary file"},
	ER_VALIDATE_PWD_FAILED_TO_GET_FLD_FROM_SECURITY_CTX : {11232,[]string{"HY000"},"Can't retrieve the %s from the security context"},
	ER_VALIDATE_PWD_FAILED_TO_GET_SECURITY_CTX : {11233,[]string{"HY000"},"Can't retrieve the security context"},
	ER_VALIDATE_PWD_LENGTH_CHANGED : {11234,[]string{"HY000"},"Effective value of validate_password_length is changed. New value is %d"},
	ER_REWRITER_QUERY_ERROR_MSG : {11235,[]string{"HY000"},"%s"},
	ER_REWRITER_QUERY_FAILED : {11236,[]string{"HY000"},"Rewritten query failed to parse:%s"},
	ER_XPLUGIN_STARTUP_FAILED : {11237,[]string{"HY000"},"Startup failed with error \"%s\""},
	//OBSOLETE_ER_XPLUGIN_SERVER_EXITING : {11238,[]string{"HY000"},"Exiting"},
	//OBSOLETE_ER_XPLUGIN_SERVER_EXITED : {11239,[]string{"HY000"},"Exit done"},
	ER_XPLUGIN_USING_SSL_CONF_FROM_SERVER : {11240,[]string{"HY000"},"Using SSL configuration from MySQL Server"},
	ER_XPLUGIN_USING_SSL_CONF_FROM_MYSQLX : {11241,[]string{"HY000"},"Using SSL configuration from Mysqlx Plugin"},
	ER_XPLUGIN_FAILED_TO_USE_SSL_CONF : {11242,[]string{"HY000"},"Neither MySQL Server nor Mysqlx Plugin has valid SSL configuration"},
	ER_XPLUGIN_USING_SSL_FOR_TLS_CONNECTION : {11243,[]string{"HY000"},"Using %s for TLS connections"},
	ER_XPLUGIN_REFERENCE_TO_SECURE_CONN_WITH_XPLUGIN : {11244,[]string{"HY000"},"For more information, please see the Using Secure Connections with X Plugin section in the MySQL documentation"},
	ER_XPLUGIN_ERROR_MSG : {11245,[]string{"HY000"},"%s"},
	ER_SHA_PWD_FAILED_TO_PARSE_AUTH_STRING : {11246,[]string{"HY000"},"Failed to parse stored authentication string for %s. Please check if mysql.user table not corrupted"},
	ER_SHA_PWD_FAILED_TO_GENERATE_MULTI_ROUND_HASH : {11247,[]string{"HY000"},"Error in generating multi-round hash for %s. Plugin can not perform authentication without it. This may be a transient problem"},
	ER_SHA_PWD_AUTH_REQUIRES_RSA_OR_SSL : {11248,[]string{"HY000"},"Authentication requires either RSA keys or SSL encryption"},
	ER_SHA_PWD_RSA_KEY_TOO_LONG : {11249,[]string{"HY000"},"RSA key cipher length of %u is too long. Max value is %u"},
	ER_PLUGIN_COMMON_FAILED_TO_OPEN_FILTER_TABLES : {11250,[]string{"HY000"},"Failed to open the %s filter tables"},
	ER_PLUGIN_COMMON_FAILED_TO_OPEN_TABLE : {11251,[]string{"HY000"},"Failed to open '%s.%s' %s table"},
	ER_AUTH_LDAP_ERROR_LOGGER_ERROR_MSG : {11252,[]string{"HY000"},"%s"},
	ER_CONN_CONTROL_ERROR_MSG : {11253,[]string{"HY000"},"%s"},
	ER_GRP_RPL_ERROR_MSG : {11254,[]string{"HY000"},"%s"},
	ER_SHA_PWD_SALT_FOR_USER_CORRUPT : {11255,[]string{"HY000"},"Password salt for user '%s' is corrupt"},
	ER_SYS_VAR_COMPONENT_OOM : {11256,[]string{"HY000"},"Out of memory for component system variable '%s'."},
	ER_SYS_VAR_COMPONENT_VARIABLE_SET_READ_ONLY : {11257,[]string{"HY000"},"variable %s of component %s was forced to be read-only: string variable without update_func and PLUGIN_VAR_MEMALLOC flag."},
	ER_SYS_VAR_COMPONENT_UNKNOWN_VARIABLE_TYPE : {11258,[]string{"HY000"},"Unknown variable type code 0x%x in component '%s'."},
	ER_SYS_VAR_COMPONENT_FAILED_TO_PARSE_VARIABLE_OPTIONS : {11259,[]string{"HY000"},"Parsing options for variable '%s' failed."},
	ER_SYS_VAR_COMPONENT_FAILED_TO_MAKE_VARIABLE_PERSISTENT : {11260,[]string{"HY000"},"Setting persistent options for component variable '%s' failed."},
	ER_COMPONENT_FILTER_CONFUSED : {11261,[]string{"HY000"},"The log-filter component \"%s\" got confused at \"%s\" (state: %s) ..."},
	ER_STOP_SLAVE_IO_THREAD_DISK_SPACE : {11262,[]string{"HY000"},"Waiting until I/O thread for channel '%s' finish writing to disk before stopping. Free some disk space or use 'KILL' to abort I/O thread operation. Notice that aborting the I/O thread while rotating the relay log might corrupt the relay logs, requiring a server restart to fix it."},
	ER_LOG_FILE_CANNOT_OPEN : {11263,[]string{"HY000"},"Could not use %s for logging (error %d - %s). Turning logging off for the server process. To turn it on again: fix the cause, then%s restart the MySQL server."},
	//OBSOLETE_ER_UNABLE_TO_COLLECT_LOG_STATUS : {3722,[]string{"HY000"},"Unable to collect information for column '%-.192s': %-.192s."},
	//OBSOLETE_ER_DEPRECATED_UTF8_ALIAS : {3719,[]string{"HY000"},"'utf8' is currently an alias for the character set UTF8MB3, which will be replaced by UTF8MB4 in a future release. Please consider using UTF8MB4 in order to be unambiguous."},
	//OBSOLETE_ER_DEPRECATED_NATIONAL : {3720,[]string{"HY000"},"NATIONAL/NCHAR/NVARCHAR implies the character set UTF8MB3, which will be replaced by UTF8MB4 in a future release. Please consider using CHAR(x) CHARACTER SET UTF8MB4 in order to be unambiguous."},
	//OBSOLETE_ER_SLAVE_POSSIBLY_DIVERGED_AFTER_DDL : {3725,[]string{"HY000"},"A commit for an atomic DDL statement was unsuccessful on the master and the slave. The slave supports atomic DDL statements but the master does not, so the action taken by the slave and master might differ. Check that their states have not diverged before proceeding."},
	ER_PERSIST_OPTION_STATUS : {11268,[]string{"HY000"},"Configuring persisted options failed: \"%s\"."},
	ER_NOT_IMPLEMENTED_GET_TABLESPACE_STATISTICS : {11269,[]string{"HY000"},"The storage engine '%-.192s' does not provide dynamic table statistics"},
	//OBSOLETE_ER_UNABLE_TO_SET_OPTION : {3724,[]string{"HY000"},"This option cannot be set %s."},
	//OBSOLETE_ER_RESERVED_TABLESPACE_NAME : {3723,[]string{"HY000"},"The table '%-.192s' may not be created in the reserved tablespace '%-.192s'."},
	ER_SSL_FIPS_MODE_ERROR : {11272,[]string{"HY000"},"SSL fips mode error: %s"},
	ER_CONN_INIT_CONNECT_IGNORED : {11273,[]string{"HY000"},"init_connect variable is ignored for user: %s host: %s due to expired password."},
	//OBSOLETE_ER_UNSUPPORTED_SQL_MODE : {3899,[]string{"HY000"},"sql_mode=0x%08x is not supported"},
	ER_REWRITER_OOM : {11275,[]string{"HY000"},"Out of memory."},
	ER_REWRITER_TABLE_MALFORMED_ERROR : {11276,[]string{"HY000"},"Wrong column count or names when loading rules."},
	ER_REWRITER_LOAD_FAILED : {11277,[]string{"HY000"},"Some rules failed to load."},
	ER_REWRITER_READ_FAILED : {11278,[]string{"HY000"},"Got error from storage engine while refreshing rewrite rules."},
	ER_CONN_CONTROL_EVENT_COORDINATOR_INIT_FAILED : {11279,[]string{"HY000"},"Failed to initialize Connection_event_coordinator"},
	ER_CONN_CONTROL_STAT_CONN_DELAY_TRIGGERED_UPDATE_FAILED : {11280,[]string{"HY000"},"Failed to update connection delay triggered stats"},
	ER_CONN_CONTROL_STAT_CONN_DELAY_TRIGGERED_RESET_FAILED : {11281,[]string{"HY000"},"Failed to reset connection delay triggered stats"},
	ER_CONN_CONTROL_INVALID_CONN_DELAY_TYPE : {11282,[]string{"HY000"},"Unexpected option type for connection delay."},
	ER_CONN_CONTROL_DELAY_ACTION_INIT_FAILED : {11283,[]string{"HY000"},"Failed to initialize Connection_delay_action"},
	ER_CONN_CONTROL_FAILED_TO_SET_CONN_DELAY : {11284,[]string{"HY000"},"Could not set %s delay for connection delay."},
	ER_CONN_CONTROL_FAILED_TO_UPDATE_CONN_DELAY_HASH : {11285,[]string{"HY000"},"Failed to update connection delay hash for account : %s"},
	ER_XPLUGIN_FORCE_STOP_CLIENT : {11286,[]string{"HY000"},"%s: Force stopping client because exception occurred: %s"},
	ER_XPLUGIN_MAX_AUTH_ATTEMPTS_REACHED : {11287,[]string{"HY000"},"%s.%u: Maximum number of authentication attempts reached, login failed."},
	ER_XPLUGIN_BUFFER_PAGE_ALLOC_FAILED : {11288,[]string{"HY000"},"Error allocating Buffer_page: %s"},
	ER_XPLUGIN_DETECTED_HANGING_CLIENTS : {11289,[]string{"HY000"},"Detected %u hanging client(s)"},
	ER_XPLUGIN_FAILED_TO_ACCEPT_CLIENT : {11290,[]string{"HY000"},"Error accepting client"},
	ER_XPLUGIN_FAILED_TO_SCHEDULE_CLIENT : {11291,[]string{"HY000"},"Internal error scheduling client for execution"},
	ER_XPLUGIN_FAILED_TO_PREPARE_IO_INTERFACES : {11292,[]string{"HY000"},"Preparation of I/O interfaces failed, X Protocol won't be accessible"},
	ER_XPLUGIN_SRV_SESSION_INIT_THREAD_FAILED : {11293,[]string{"HY000"},"srv_session_init_thread returned error"},
	ER_XPLUGIN_UNABLE_TO_USE_USER_SESSION_ACCOUNT : {11294,[]string{"HY000"},"Unable to use user mysql.session account when connecting the server for internal plugin requests."},
	ER_XPLUGIN_REFERENCE_TO_USER_ACCOUNT_DOC_SECTION : {11295,[]string{"HY000"},"For more information, please see the X Plugin User Account section in the MySQL documentation"},
	ER_XPLUGIN_UNEXPECTED_EXCEPTION_DISPATCHING_CMD : {11296,[]string{"HY000"},"%s: Unexpected exception dispatching command: %s"},
	ER_XPLUGIN_EXCEPTION_IN_TASK_SCHEDULER : {11297,[]string{"HY000"},"Exception in post: %s"},
	ER_XPLUGIN_TASK_SCHEDULING_FAILED : {11298,[]string{"HY000"},"Internal error scheduling task"},
	ER_XPLUGIN_EXCEPTION_IN_EVENT_LOOP : {11299,[]string{"HY000"},"Exception in event loop: \"%s\": %s"},
	ER_XPLUGIN_LISTENER_SETUP_FAILED : {11300,[]string{"HY000"},"Setup of %s failed, %s"},
	ER_XPLUING_NET_STARTUP_FAILED : {11301,[]string{"HY000"},"%s"},
	ER_XPLUGIN_FAILED_AT_SSL_CONF : {11302,[]string{"HY000"},"Failed at SSL configuration: \"%s\""},
	//OBSOLETE_ER_XPLUGIN_CLIENT_SSL_HANDSHAKE_FAILED : {11303,[]string{"HY000"},"Error during SSL handshake for client connection (%i)"},
	//OBSOLETE_ER_XPLUGIN_SSL_HANDSHAKE_WITH_SERVER_FAILED : {11304,[]string{"HY000"},"%s: Error during SSL handshake"},
	ER_XPLUGIN_FAILED_TO_CREATE_SESSION_FOR_CONN : {11305,[]string{"HY000"},"%s: Error creating session for connection from %s"},
	ER_XPLUGIN_FAILED_TO_INITIALIZE_SESSION : {11306,[]string{"HY000"},"%s: Error initializing session for connection: %s"},
	ER_XPLUGIN_MESSAGE_TOO_LONG : {11307,[]string{"HY000"},"%s: Message of size %u received, exceeding the limit of %i"},
	ER_XPLUGIN_UNINITIALIZED_MESSAGE : {11308,[]string{"HY000"},"Message is not properly initialized: %s"},
	ER_XPLUGIN_FAILED_TO_SET_MIN_NUMBER_OF_WORKERS : {11309,[]string{"HY000"},"Unable to set minimal number of workers to %u; actual value is %i"},
	ER_XPLUGIN_UNABLE_TO_ACCEPT_CONNECTION : {11310,[]string{"HY000"},"Unable to accept connection, disconnecting client"},
	ER_XPLUGIN_ALL_IO_INTERFACES_DISABLED : {11311,[]string{"HY000"},"All I/O interfaces are disabled, X Protocol won't be accessible"},
	//OBSOLETE_ER_XPLUGIN_INVALID_MSG_DURING_CLIENT_INIT : {11312,[]string{"HY000"},"%s: Invalid message %i received during client initialization"},
	//OBSOLETE_ER_XPLUGIN_CLOSING_CLIENTS_ON_SHUTDOWN : {11313,[]string{"HY000"},"%s: closing client because of shutdown (state: %i)"},
	ER_XPLUGIN_ERROR_READING_SOCKET : {11314,[]string{"HY000"},"%s: Error reading from socket %s (%i)"},
	ER_XPLUGIN_PEER_DISCONNECTED_WHILE_READING_MSG_BODY : {11315,[]string{"HY000"},"%s: peer disconnected while reading message body"},
	ER_XPLUGIN_READ_FAILED_CLOSING_CONNECTION : {11316,[]string{"HY000"},"client_id:%s - %s while reading from socket, closing connection"},
	//OBSOLETE_ER_XPLUGIN_INVALID_AUTH_METHOD : {11317,[]string{"HY000"},"%s.%u: Invalid authentication method %s"},
	//OBSOLETE_ER_XPLUGIN_UNEXPECTED_MSG_DURING_AUTHENTICATION : {11318,[]string{"HY000"},"%s: Unexpected message of type %i received during authentication"},
	//OBSOLETE_ER_XPLUGIN_ERROR_WRITING_TO_CLIENT : {11319,[]string{"HY000"},"Error writing to client: %s (%i)"},
	//OBSOLETE_ER_XPLUGIN_SCHEDULER_STARTED : {11320,[]string{"HY000"},"Scheduler \"%s\" started."},
	//OBSOLETE_ER_XPLUGIN_SCHEDULER_STOPPED : {11321,[]string{"HY000"},"Scheduler \"%s\" stopped."},
	ER_XPLUGIN_LISTENER_SYS_VARIABLE_ERROR : {11322,[]string{"HY000"},"Please see the MySQL documentation for '%s' system variables to fix the error"},
	ER_XPLUGIN_LISTENER_STATUS_MSG : {11323,[]string{"HY000"},"X Plugin ready for connections. %s"},
	ER_XPLUGIN_RETRYING_BIND_ON_PORT : {11324,[]string{"HY000"},"Retrying `bind()` on TCP/IP port %i"},
	//OBSOLETE_ER_XPLUGIN_SHUTDOWN_TRIGGERED : {11325,[]string{"HY000"},"Shutdown triggered by mysqld abort flag"},
	//OBSOLETE_ER_XPLUGIN_USER_ACCOUNT_WITH_ALL_PERMISSIONS : {11326,[]string{"HY000"},"Using %s account for authentication which has all required permissions"},
	ER_XPLUGIN_EXISTING_USER_ACCOUNT_WITH_INCOMPLETE_GRANTS : {11327,[]string{"HY000"},"Using existing %s account for authentication. Incomplete grants will be fixed"},
	//OBSOLETE_ER_XPLUGIN_SERVER_STARTS_HANDLING_CONNECTIONS : {11328,[]string{"HY000"},"Server starts handling incoming connections"},
	//OBSOLETE_ER_XPLUGIN_SERVER_STOPPED_HANDLING_CONNECTIONS : {11329,[]string{"HY000"},"Stopped handling incoming connections"},
	//OBSOLETE_ER_XPLUGIN_FAILED_TO_INTERRUPT_SESSION : {11330,[]string{"HY000"},"%s: Could not interrupt client session"},
	//OBSOLETE_ER_XPLUGIN_CLIENT_RELEASE_TRIGGERED : {11331,[]string{"HY000"},"%s: release triggered by timeout in state:%i"},
	ER_XPLUGIN_IPv6_AVAILABLE : {11332,[]string{"HY000"},"IPv6 is available"},
	//OBSOLETE_ER_XPLUGIN_UNIX_SOCKET_NOT_CONFIGURED : {11333,[]string{"HY000"},"UNIX socket not configured"},
	ER_XPLUGIN_CLIENT_KILL_MSG : {11334,[]string{"HY000"},"Kill client: %i %s"},
	ER_XPLUGIN_FAILED_TO_GET_SECURITY_CTX : {11335,[]string{"HY000"},"Could not get security context for session"},
	//OBSOLETE_ER_XPLUGIN_FAILED_TO_SWITCH_SECURITY_CTX_TO_ROOT : {11336,[]string{"HY000"},"Unable to switch security context to root"},
	ER_XPLUGIN_FAILED_TO_CLOSE_SQL_SESSION : {11337,[]string{"HY000"},"Error closing SQL session"},
	ER_XPLUGIN_FAILED_TO_EXECUTE_ADMIN_CMD : {11338,[]string{"HY000"},"Error executing admin command %s: %s"},
	ER_XPLUGIN_EMPTY_ADMIN_CMD : {11339,[]string{"HY000"},"Error executing empty admin command"},
	ER_XPLUGIN_FAILED_TO_GET_SYS_VAR : {11340,[]string{"HY000"},"Unable to retrieve system variable '%s'"},
	ER_XPLUGIN_FAILED_TO_GET_CREATION_STMT : {11341,[]string{"HY000"},"Unable to get creation stmt for collection '%s'; query result size: %lu"},
	ER_XPLUGIN_FAILED_TO_GET_ENGINE_INFO : {11342,[]string{"HY000"},"Unable to get engine info for collection '%s'; creation stmt: %s"},
	//OBSOLETE_ER_XPLUGIN_FAIL_TO_GET_RESULT_DATA : {11343,[]string{"HY000"},"Error getting result data: %s"},
	//OBSOLETE_ER_XPLUGIN_CAPABILITY_EXPIRED_PASSWORD : {11344,[]string{"HY000"},"Capability expired password failed with error: %s"},
	ER_XPLUGIN_FAILED_TO_SET_SO_REUSEADDR_FLAG : {11345,[]string{"HY000"},"Failed to set SO_REUSEADDR flag (error: %d)."},
	ER_XPLUGIN_FAILED_TO_OPEN_INTERNAL_SESSION : {11346,[]string{"HY000"},"Could not open internal MySQL session"},
	ER_XPLUGIN_FAILED_TO_SWITCH_CONTEXT : {11347,[]string{"HY000"},"Unable to switch context to user %s"},
	ER_XPLUGIN_FAILED_TO_UNREGISTER_UDF : {11348,[]string{"HY000"},"Can't unregister '%s' user defined function"},
	//OBSOLETE_ER_XPLUGIN_GET_PEER_ADDRESS_FAILED : {11349,[]string{"HY000"},"%s: get peer address failed, can't resolve IP to hostname"},
	//OBSOLETE_ER_XPLUGIN_CAPABILITY_CLIENT_INTERACTIVE_FAILED : {11350,[]string{"HY000"},"Capability client interactive failed with error: %s"},
	ER_XPLUGIN_FAILED_TO_RESET_IPV6_V6ONLY_FLAG : {11351,[]string{"HY000"},"Failed to reset IPV6_V6ONLY flag (error: %d). The server will listen to IPv6 addresses only."},
	ER_KEYRING_INVALID_KEY_TYPE : {11352,[]string{"HY000"},"Invalid key type"},
	ER_KEYRING_INVALID_KEY_LENGTH : {11353,[]string{"HY000"},"Invalid key length for given block cipher"},
	ER_KEYRING_FAILED_TO_CREATE_KEYRING_DIR : {11354,[]string{"HY000"},"Could not create keyring directory. The keyring_file will stay unusable until correct path to the keyring directory gets provided"},
	ER_KEYRING_FILE_INIT_FAILED : {11355,[]string{"HY000"},"keyring_file initialization failure. Please check if the keyring_file_data points to readable keyring file or keyring file can be created in the specified location. The keyring_file will stay unusable until correct path to the keyring file gets provided"},
	ER_KEYRING_INTERNAL_EXCEPTION_FAILED_FILE_INIT : {11356,[]string{"HY000"},"keyring_file initialization failure due to internal exception inside the plugin."},
	ER_KEYRING_FAILED_TO_GENERATE_KEY : {11357,[]string{"HY000"},"Failed to generate a key due to internal exception inside keyring_file plugin"},
	ER_KEYRING_CHECK_KEY_FAILED_DUE_TO_INVALID_KEY : {11358,[]string{"HY000"},"Error while %s key: invalid key_type"},
	ER_KEYRING_CHECK_KEY_FAILED_DUE_TO_EMPTY_KEY_ID : {11359,[]string{"HY000"},"Error while %s key: key_id cannot be empty"},
	ER_KEYRING_OPERATION_FAILED_DUE_TO_INTERNAL_ERROR : {11360,[]string{"HY000"},"Failed to %s due to internal exception inside %s plugin"},
	ER_KEYRING_INCORRECT_FILE : {11361,[]string{"HY000"},"Incorrect Keyring file"},
	ER_KEYRING_FOUND_MALFORMED_BACKUP_FILE : {11362,[]string{"HY000"},"Found malformed keyring backup file - removing it"},
	ER_KEYRING_FAILED_TO_RESTORE_FROM_BACKUP_FILE : {11363,[]string{"HY000"},"Error while restoring keyring from backup file cannot overwrite keyring with backup"},
	ER_KEYRING_FAILED_TO_FLUSH_KEYRING_TO_FILE : {11364,[]string{"HY000"},"Error while flushing in-memory keyring into keyring file"},
	ER_KEYRING_FAILED_TO_GET_FILE_STAT : {11365,[]string{"HY000"},"Error while reading stat for %s.Please check if file %s was not removed. OS returned this error: %s"},
	ER_KEYRING_FAILED_TO_REMOVE_FILE : {11366,[]string{"HY000"},"Could not remove file %s OS retuned this error: %s"},
	ER_KEYRING_FAILED_TO_TRUNCATE_FILE : {11367,[]string{"HY000"},"Could not truncate file %s. OS retuned this error: %s"},
	ER_KEYRING_UNKNOWN_ERROR : {11368,[]string{"HY000"},"Unknown error %d"},
	ER_KEYRING_FAILED_TO_SET_KEYRING_FILE_DATA : {11369,[]string{"HY000"},"keyring_file_data cannot be set to new value as the keyring file cannot be created/accessed in the provided path"},
	ER_KEYRING_FILE_IO_ERROR : {11370,[]string{"HY000"},"%s"},
	ER_KEYRING_FAILED_TO_LOAD_KEYRING_CONTENT : {11371,[]string{"HY000"},"Error while loading keyring content. The keyring might be malformed"},
	ER_KEYRING_FAILED_TO_FLUSH_KEYS_TO_KEYRING : {11372,[]string{"HY000"},"Could not flush keys to keyring"},
	ER_KEYRING_FAILED_TO_FLUSH_KEYS_TO_KEYRING_BACKUP : {11373,[]string{"HY000"},"Could not flush keys to keyring's backup"},
	ER_KEYRING_KEY_FETCH_FAILED_DUE_TO_EMPTY_KEY_ID : {11374,[]string{"HY000"},"Error while fetching key: key_id cannot be empty"},
	ER_KEYRING_FAILED_TO_REMOVE_KEY_DUE_TO_EMPTY_ID : {11375,[]string{"HY000"},"Error while removing key: key_id cannot be empty"},
	ER_KEYRING_OKV_INCORRECT_KEY_VAULT_CONFIGURED : {11376,[]string{"HY000"},"For keyring_okv to be initialized, please point keyring_okv_conf_dir variable to a directory with Oracle Key Vault configuration file and ssl materials"},
	ER_KEYRING_OKV_INIT_FAILED_DUE_TO_INCORRECT_CONF : {11377,[]string{"HY000"},"keyring_okv initialization failure. Please check that the keyring_okv_conf_dir points to a readable directory and that the directory contains Oracle Key Vault configuration file and ssl materials. Please also check that Oracle Key Vault is up and running."},
	ER_KEYRING_OKV_INIT_FAILED_DUE_TO_INTERNAL_ERROR : {11378,[]string{"HY000"},"keyring_okv initialization failure due to internal exception inside the plugin"},
	ER_KEYRING_OKV_INVALID_KEY_TYPE : {11379,[]string{"HY000"},"Invalid key type"},
	ER_KEYRING_OKV_INVALID_KEY_LENGTH_FOR_CIPHER : {11380,[]string{"HY000"},"Invalid key length for given block cipher"},
	ER_KEYRING_OKV_FAILED_TO_GENERATE_KEY_DUE_TO_INTERNAL_ERROR : {11381,[]string{"HY000"},"Failed to generate a key due to internal exception inside keyring_okv plugin"},
	ER_KEYRING_OKV_FAILED_TO_FIND_SERVER_ENTRY : {11382,[]string{"HY000"},"Could not find entry for server in configuration file %s"},
	ER_KEYRING_OKV_FAILED_TO_FIND_STANDBY_SERVER_ENTRY : {11383,[]string{"HY000"},"Could not find entry for standby server in configuration file %s"},
	ER_KEYRING_OKV_FAILED_TO_PARSE_CONF_FILE : {11384,[]string{"HY000"},"Could not parse the %s file provided"},
	ER_KEYRING_OKV_FAILED_TO_LOAD_KEY_UID : {11385,[]string{"HY000"},"Could not load keys' uids from the OKV server"},
	ER_KEYRING_OKV_FAILED_TO_INIT_SSL_LAYER : {11386,[]string{"HY000"},"Could not initialize ssl layer"},
	ER_KEYRING_OKV_FAILED_TO_INIT_CLIENT : {11387,[]string{"HY000"},"Could not initialize OKV client"},
	ER_KEYRING_OKV_CONNECTION_TO_SERVER_FAILED : {11388,[]string{"HY000"},"Could not connect to the OKV server"},
	ER_KEYRING_OKV_FAILED_TO_REMOVE_KEY : {11389,[]string{"HY000"},"Could not remove the key"},
	ER_KEYRING_OKV_FAILED_TO_ADD_ATTRIBUTE : {11390,[]string{"HY000"},"Could not add attribute, attribute_name=%s attribute value=%s"},
	ER_KEYRING_OKV_FAILED_TO_GENERATE_KEY : {11391,[]string{"HY000"},"Could not generate the key."},
	ER_KEYRING_OKV_FAILED_TO_STORE_KEY : {11392,[]string{"HY000"},"Could not store the key."},
	ER_KEYRING_OKV_FAILED_TO_ACTIVATE_KEYS : {11393,[]string{"HY000"},"Could not activate the key."},
	ER_KEYRING_OKV_FAILED_TO_FETCH_KEY : {11394,[]string{"HY000"},"Could not fetch generated key"},
	ER_KEYRING_OKV_FAILED_TO_STORE_OR_GENERATE_KEY : {11395,[]string{"HY000"},"Could not store/generate the key - failed to set key attribute x-key-ready"},
	ER_KEYRING_OKV_FAILED_TO_RETRIEVE_KEY_SIGNATURE : {11396,[]string{"HY000"},"Could not retrieve key signature from custom attributes"},
	ER_KEYRING_OKV_FAILED_TO_RETRIEVE_KEY : {11397,[]string{"HY000"},"Could not retrieve key from OKV"},
	ER_KEYRING_OKV_FAILED_TO_LOAD_SSL_TRUST_STORE : {11398,[]string{"HY000"},"Error loading trust store"},
	ER_KEYRING_OKV_FAILED_TO_SET_CERTIFICATE_FILE : {11399,[]string{"HY000"},"Error setting the certificate file."},
	ER_KEYRING_OKV_FAILED_TO_SET_KEY_FILE : {11400,[]string{"HY000"},"Error setting the key file."},
	ER_KEYRING_OKV_KEY_MISMATCH : {11401,[]string{"HY000"},"Private key does not match the certificate public key"},
	ER_KEYRING_ENCRYPTED_FILE_INCORRECT_KEYRING_FILE : {11402,[]string{"HY000"},"Incorrect Keyring file"},
	ER_KEYRING_ENCRYPTED_FILE_DECRYPTION_FAILED : {11403,[]string{"HY000"},"Keyring_encrypted_file decryption failed. Please verify --keyring-encrypted-file-password option value."},
	ER_KEYRING_ENCRYPTED_FILE_FOUND_MALFORMED_BACKUP_FILE : {11404,[]string{"HY000"},"Found malformed keyring backup file - removing it"},
	ER_KEYRING_ENCRYPTED_FILE_FAILED_TO_RESTORE_KEYRING : {11405,[]string{"HY000"},"Error while restoring keyring from backup file cannot overwrite keyring with backup"},
	ER_KEYRING_ENCRYPTED_FILE_FAILED_TO_FLUSH_KEYRING : {11406,[]string{"HY000"},"Error while flushing in-memory keyring into keyring file"},
	ER_KEYRING_ENCRYPTED_FILE_ENCRYPTION_FAILED : {11407,[]string{"HY000"},"Keyring_encrypted_file encryption failed. Please verify --keyring-encrypted-file-password option value."},
	ER_KEYRING_ENCRYPTED_FILE_INVALID_KEYRING_DIR : {11408,[]string{"HY000"},"keyring_encrypted_file_data cannot be set to new value as the keyring file cannot be created/accessed in the provided path"},
	ER_KEYRING_ENCRYPTED_FILE_FAILED_TO_CREATE_KEYRING_DIR : {11409,[]string{"HY000"},"Could not create keyring directory The keyring_encrypted_file will stay unusable until correct path to the keyring directory gets provided"},
	ER_KEYRING_ENCRYPTED_FILE_PASSWORD_IS_INVALID : {11410,[]string{"HY000"},"The keyring_encrypted_file_password must be set to a valid value."},
	ER_KEYRING_ENCRYPTED_FILE_PASSWORD_IS_TOO_LONG : {11411,[]string{"HY000"},"Too long keyring_encrypted_file_password value."},
	ER_KEYRING_ENCRYPTED_FILE_INIT_FAILURE : {11412,[]string{"HY000"},"keyring_encrypted_file initialization failure. Please check if the keyring_encrypted_file_data points to readable keyring file or keyring file can be created in the specified location or password to decrypt keyring file is correct."},
	ER_KEYRING_ENCRYPTED_FILE_INIT_FAILED_DUE_TO_INTERNAL_ERROR : {11413,[]string{"HY000"},"keyring_encrypted_file initialization failure due to internal exception inside the plugin"},
	ER_KEYRING_ENCRYPTED_FILE_GEN_KEY_FAILED_DUE_TO_INTERNAL_ERROR : {11414,[]string{"HY000"},"Failed to generate a key due to internal exception inside keyring_encrypted_file plugin"},
	ER_KEYRING_AWS_FAILED_TO_SET_CMK_ID : {11415,[]string{"HY000"},"keyring_aws_cmk_id cannot be set to the new value as AWS KMS seems to not understand the id provided. Please check that CMK id provided is correct."},
	ER_KEYRING_AWS_FAILED_TO_SET_REGION : {11416,[]string{"HY000"},"keyring_aws_region cannot be set to the new value as AWS KMS seems to not understand the region provided. Please check that region provided is correct."},
	ER_KEYRING_AWS_FAILED_TO_OPEN_CONF_FILE : {11417,[]string{"HY000"},"Could not open keyring_aws configuration file: %s. OS returned this error: %s"},
	ER_KEYRING_AWS_FAILED_TO_ACCESS_KEY_ID_FROM_CONF_FILE : {11418,[]string{"HY000"},"Could not read AWS access key id from keyring_aws configuration file: %s. OS returned this error: %s"},
	ER_KEYRING_AWS_FAILED_TO_ACCESS_KEY_FROM_CONF_FILE : {11419,[]string{"HY000"},"Could not read AWS access key from keyring_aws configuration file: %s. OS returned this error: %s"},
	ER_KEYRING_AWS_INVALID_CONF_FILE_PATH : {11420,[]string{"HY000"},"Path to keyring aws configuration file cannot be empty"},
	ER_KEYRING_AWS_INVALID_DATA_FILE_PATH : {11421,[]string{"HY000"},"Path to keyring_aws storage file cannot be empty."},
	ER_KEYRING_AWS_FAILED_TO_ACCESS_OR_CREATE_KEYRING_DIR : {11422,[]string{"HY000"},"Unable to create/access keyring directory."},
	ER_KEYRING_AWS_FAILED_TO_ACCESS_OR_CREATE_KEYRING_DATA_FILE : {11423,[]string{"HY000"},"Unable to create/access keyring_aws storage file. Please check if keyring_aws_data_file points to location where keyring file can be created/accessed. Please also make sure that MySQL server's user has high enough privileges to access this location."},
	ER_KEYRING_AWS_FAILED_TO_INIT_DUE_TO_INTERNAL_ERROR : {11424,[]string{"HY000"},"keyring_aws initialization failed due to internal error when initializing synchronization primitive - %s. OS returned this error: %s:"},
	ER_KEYRING_AWS_FAILED_TO_ACCESS_DATA_FILE : {11425,[]string{"HY000"},"Could not access keyring_aws storage file in the path provided. Please check if the keyring_aws directory can be accessed by MySQL Server"},
	ER_KEYRING_AWS_CMK_ID_NOT_SET : {11426,[]string{"HY000"},"keyring_aws_cmk_id has to be set"},
	ER_KEYRING_AWS_FAILED_TO_GET_KMS_CREDENTIAL_FROM_CONF_FILE : {11427,[]string{"HY000"},"Could not get AWS KMS credentials from the configuration file"},
	ER_KEYRING_AWS_INIT_FAILURE : {11428,[]string{"HY000"},"keyring_aws initialization failure."},
	ER_KEYRING_AWS_FAILED_TO_INIT_DUE_TO_PLUGIN_INTERNAL_ERROR : {11429,[]string{"HY000"},"keyring_aws initialization failure due to internal exception inside the plugin"},
	ER_KEYRING_AWS_INVALID_KEY_LENGTH_FOR_CIPHER : {11430,[]string{"HY000"},"Invalid key length for given block cipher"},
	ER_KEYRING_AWS_FAILED_TO_GENERATE_KEY_DUE_TO_INTERNAL_ERROR : {11431,[]string{"HY000"},"Failed to generate a key due to internal exception inside keyring_file plugin"},
	ER_KEYRING_AWS_INCORRECT_FILE : {11432,[]string{"HY000"},"Incorrect Keyring file"},
	ER_KEYRING_AWS_FOUND_MALFORMED_BACKUP_FILE : {11433,[]string{"HY000"},"Found malformed keyring backup file - removing it"},
	ER_KEYRING_AWS_FAILED_TO_RESTORE_FROM_BACKUP_FILE : {11434,[]string{"HY000"},"Error while restoring keyring from backup file cannot overwrite keyring with backup"},
	ER_KEYRING_AWS_FAILED_TO_FLUSH_KEYRING_TO_FILE : {11435,[]string{"HY000"},"Error while flushing in-memory keyring into keyring file"},
	ER_KEYRING_AWS_INCORRECT_REGION : {11436,[]string{"HY000"},"Wrong region"},
	ER_KEYRING_AWS_FAILED_TO_CONNECT_KMS : {11437,[]string{"HY000"},"Could not connect to AWS KMS with the credentials provided. Please make sure they are correct. AWS KMS returned this error: %s"},
	ER_KEYRING_AWS_FAILED_TO_GENERATE_NEW_KEY : {11438,[]string{"HY000"},"Could not generate a new key. AWS KMS returned this error: %s"},
	ER_KEYRING_AWS_FAILED_TO_ENCRYPT_KEY : {11439,[]string{"HY000"},"Could not encrypt key. AWS KMS returned this error: %s"},
	ER_KEYRING_AWS_FAILED_TO_RE_ENCRYPT_KEY : {11440,[]string{"HY000"},"Could not re-encrypt key. AWS KMS returned this error: %s"},
	ER_KEYRING_AWS_FAILED_TO_DECRYPT_KEY : {11441,[]string{"HY000"},"Could not decrypt key. AWS KMS returned this error: %s"},
	ER_KEYRING_AWS_FAILED_TO_ROTATE_CMK : {11442,[]string{"HY000"},"Could not rotate the CMK. AWS KMS returned this error: %s"},
	ER_GRP_RPL_GTID_ALREADY_USED : {11443,[]string{"HY000"},"The requested GTID '%s:%lld' was already used, the transaction will rollback."},
	ER_GRP_RPL_APPLIER_THD_KILLED : {11444,[]string{"HY000"},"The group replication applier thread was killed."},
	ER_GRP_RPL_EVENT_HANDLING_ERROR : {11445,[]string{"HY000"},"Error at event handling! Got error: %d."},
	ER_GRP_RPL_ERROR_GTID_EXECUTION_INFO : {11446,[]string{"HY000"},"Error when extracting group GTID execution information, some recovery operations may face future issues."},
	ER_GRP_RPL_CERTIFICATE_SIZE_ERROR : {11447,[]string{"HY000"},"An error occurred when trying to reduce the Certification information size for transmission."},
	ER_GRP_RPL_CREATE_APPLIER_CACHE_ERROR : {11448,[]string{"HY000"},"Failed to create group replication pipeline applier cache!"},
	ER_GRP_RPL_UNBLOCK_WAITING_THD : {11449,[]string{"HY000"},"Unblocking the group replication thread waiting for applier to start, as the start group replication was killed."},
	ER_GRP_RPL_APPLIER_PIPELINE_NOT_DISPOSED : {11450,[]string{"HY000"},"The group replication applier pipeline was not properly disposed. Check the error log for further info."},
	ER_GRP_RPL_APPLIER_THD_EXECUTION_ABORTED : {11451,[]string{"HY000"},"The applier thread execution was aborted. Unable to process more transactions, this member will now leave the group."},
	ER_GRP_RPL_APPLIER_EXECUTION_FATAL_ERROR : {11452,[]string{"HY000"},"Fatal error during execution on the Applier process of Group Replication. The server will now leave the group."},
	ER_GRP_RPL_ERROR_STOPPING_CHANNELS : {11453,[]string{"HY000"},"Error stopping all replication channels while server was leaving the group. %s"},
	ER_GRP_RPL_ERROR_SENDING_SINGLE_PRIMARY_MSSG : {11454,[]string{"HY000"},"Error sending single primary message informing that primary did apply relay logs."},
	ER_GRP_RPL_UPDATE_TRANS_SNAPSHOT_VER_ERROR : {11455,[]string{"HY000"},"Error updating transaction snapshot version after transaction being positively certified."},
	ER_GRP_RPL_SIDNO_FETCH_ERROR : {11456,[]string{"HY000"},"Error fetching transaction sidno after transaction being positively certified."},
	ER_GRP_RPL_BROADCAST_COMMIT_TRANS_MSSG_FAILED : {11457,[]string{"HY000"},"Broadcast of committed transactions message failed."},
	ER_GRP_RPL_GROUP_NAME_PARSE_ERROR : {11458,[]string{"HY000"},"Unable to parse the group_replication_group_name during the Certification module initialization."},
	ER_GRP_RPL_ADD_GRPSID_TO_GRPGTIDSID_MAP_ERROR : {11459,[]string{"HY000"},"Unable to add the group_sid in the group_gtid_sid_map during the Certification module initialization."},
	ER_GRP_RPL_UPDATE_GRPGTID_EXECUTED_ERROR : {11460,[]string{"HY000"},"Error updating group_gtid_executed GITD set during the Certification module initialization."},
	ER_GRP_RPL_DONOR_TRANS_INFO_ERROR : {11461,[]string{"HY000"},"Unable to handle the donor's transaction information when initializing the conflict detection component. Possible out of memory error."},
	ER_GRP_RPL_SERVER_CONN_ERROR : {11462,[]string{"HY000"},"Error when establishing a server connection during the Certification module initialization."},
	ER_GRP_RPL_ERROR_FETCHING_GTID_EXECUTED_SET : {11463,[]string{"HY000"},"Error when extracting this member GTID executed set. Certification module can't be properly initialized."},
	ER_GRP_RPL_ADD_GTID_TO_GRPGTID_EXECUTED_ERROR : {11464,[]string{"HY000"},"Error while adding the server GTID EXECUTED set to the group_gtid_execute during the Certification module initialization."},
	ER_GRP_RPL_ERROR_FETCHING_GTID_SET : {11465,[]string{"HY000"},"Error when extracting this member retrieved set for its applier. Certification module can't be properly initialized."},
	ER_GRP_RPL_ADD_RETRIEVED_SET_TO_GRP_GTID_EXECUTED_ERROR : {11466,[]string{"HY000"},"Error while adding the member retrieved set to the group_gtid_executed during the Certification module initialization."},
	ER_GRP_RPL_CERTIFICATION_INITIALIZATION_FAILURE : {11467,[]string{"HY000"},"Error during Certification module initialization."},
	ER_GRP_RPL_UPDATE_LAST_CONFLICT_FREE_TRANS_ERROR : {11468,[]string{"HY000"},"Unable to update last conflict free transaction, this transaction will not be tracked on performance_schema.replication_group_member_stats.last_conflict_free_transaction."},
	ER_GRP_RPL_UPDATE_TRANS_SNAPSHOT_REF_VER_ERROR : {11469,[]string{"HY000"},"Error updating transaction snapshot version reference for internal storage."},
	ER_GRP_RPL_FETCH_TRANS_SIDNO_ERROR : {11470,[]string{"HY000"},"Error fetching transaction sidno while adding to the group_gtid_executed set."},
	ER_GRP_RPL_ERROR_VERIFYING_SIDNO : {11471,[]string{"HY000"},"Error while ensuring the sidno be present in the group_gtid_executed."},
	ER_GRP_RPL_CANT_GENERATE_GTID : {11472,[]string{"HY000"},"Impossible to generate Global Transaction Identifier: the integer component reached the maximal value. Restart the group with a new group_replication_group_name."},
	ER_GRP_RPL_INVALID_GTID_SET : {11473,[]string{"HY000"},"Invalid stable transactions set."},
	ER_GRP_RPL_UPDATE_GTID_SET_ERROR : {11474,[]string{"HY000"},"Error updating stable transactions set."},
	ER_GRP_RPL_RECEIVED_SET_MISSING_GTIDS : {11475,[]string{"HY000"},"There was an error when filling the missing GTIDs on the applier channel received set. Despite not critical, on the long run this may cause performance issues."},
	ER_GRP_RPL_SKIP_COMPUTATION_TRANS_COMMITTED : {11476,[]string{"HY000"},"Skipping the computation of the Transactions_committed_all_members field as an older instance of this computation is still ongoing."},
	ER_GRP_RPL_NULL_PACKET : {11477,[]string{"HY000"},"Null packet on certifier's queue."},
	ER_GRP_RPL_CANT_READ_GTID : {11478,[]string{"HY000"},"Error reading GTIDs from the message."},
	ER_GRP_RPL_PROCESS_GTID_SET_ERROR : {11479,[]string{"HY000"},"Error processing stable transactions set."},
	ER_GRP_RPL_PROCESS_INTERSECTION_GTID_SET_ERROR : {11480,[]string{"HY000"},"Error processing intersection of stable transactions set."},
	ER_GRP_RPL_SET_STABLE_TRANS_ERROR : {11481,[]string{"HY000"},"Error setting stable transactions set."},
	ER_GRP_RPL_CANT_READ_GRP_GTID_EXTRACTED : {11482,[]string{"HY000"},"Error reading group_gtid_extracted from the View_change_log_event."},
	ER_GRP_RPL_CANT_READ_WRITE_SET_ITEM : {11483,[]string{"HY000"},"Error reading the write set item '%s' from the View_change_log_event."},
	ER_GRP_RPL_INIT_CERTIFICATION_INFO_FAILURE : {11484,[]string{"HY000"},"Error during certification_info initialization."},
	ER_GRP_RPL_CONFLICT_DETECTION_DISABLED : {11485,[]string{"HY000"},"Primary had applied all relay logs, disabled conflict detection."},
	ER_GRP_RPL_MSG_DISCARDED : {11486,[]string{"HY000"},"Message received while the plugin is not ready, message discarded."},
	ER_GRP_RPL_MISSING_GRP_RPL_APPLIER : {11487,[]string{"HY000"},"Message received without a proper group replication applier."},
	ER_GRP_RPL_CERTIFIER_MSSG_PROCESS_ERROR : {11488,[]string{"HY000"},"Error processing message in Certifier."},
	ER_GRP_RPL_SRV_NOT_ONLINE : {11489,[]string{"HY000"},"This server was not declared online since it is on status %s."},
	ER_GRP_RPL_SRV_ONLINE : {11490,[]string{"HY000"},"This server was declared online within the replication group."},
	ER_GRP_RPL_DISABLE_SRV_READ_MODE_RESTRICTED : {11491,[]string{"HY000"},"When declaring the plugin online it was not possible to disable the server read mode settings. Try to disable it manually."},
	ER_GRP_RPL_MEM_ONLINE : {11492,[]string{"HY000"},"The member with address %s:%u was declared online within the replication group."},
	ER_GRP_RPL_MEM_UNREACHABLE : {11493,[]string{"HY000"},"Member with address %s:%u has become unreachable."},
	ER_GRP_RPL_MEM_REACHABLE : {11494,[]string{"HY000"},"Member with address %s:%u is reachable again."},
	ER_GRP_RPL_SRV_BLOCKED : {11495,[]string{"HY000"},"This server is not able to reach a majority of members in the group. This server will now block all updates. The server will remain blocked until contact with the majority is restored. It is possible to use group_replication_force_members to force a new group membership."},
	ER_GRP_RPL_SRV_BLOCKED_FOR_SECS : {11496,[]string{"HY000"},"This server is not able to reach a majority of members in the group. This server will now block all updates. The server will remain blocked for the next %lu seconds. Unless contact with the majority is restored, after this time the member will error out and leave the group. It is possible to use group_replication_force_members to force a new group membership."},
	ER_GRP_RPL_CHANGE_GRP_MEM_NOT_PROCESSED : {11497,[]string{"HY000"},"A group membership change was received but the plugin is already leaving due to the configured timeout on group_replication_unreachable_majority_timeout option."},
	ER_GRP_RPL_MEMBER_CONTACT_RESTORED : {11498,[]string{"HY000"},"The member has resumed contact with a majority of the members in the group. Regular operation is restored and transactions are unblocked."},
	ER_GRP_RPL_MEMBER_REMOVED : {11499,[]string{"HY000"},"Members removed from the group: %s"},
	ER_GRP_RPL_PRIMARY_MEMBER_LEFT_GRP : {11500,[]string{"HY000"},"Primary server with address %s left the group. Electing new Primary."},
	ER_GRP_RPL_MEMBER_ADDED : {11501,[]string{"HY000"},"Members joined the group: %s"},
	ER_GRP_RPL_MEMBER_EXIT_PLUGIN_ERROR : {11502,[]string{"HY000"},"There was a previous plugin error while the member joined the group. The member will now exit the group."},
	ER_GRP_RPL_MEMBER_CHANGE : {11503,[]string{"HY000"},"Group membership changed to %s on view %s."},
	ER_GRP_RPL_MEMBER_LEFT_GRP : {11504,[]string{"HY000"},"Group membership changed: This member has left the group."},
	ER_GRP_RPL_MEMBER_EXPELLED : {11505,[]string{"HY000"},"Member was expelled from the group due to network failures, changing member status to ERROR."},
	ER_GRP_RPL_SESSION_OPEN_FAILED : {11506,[]string{"HY000"},"Unable to open session to (re)set read only mode. Skipping."},
	ER_GRP_RPL_NEW_PRIMARY_ELECTED : {11507,[]string{"HY000"},"A new primary with address %s:%u was elected. %s"},
	ER_GRP_RPL_DISABLE_READ_ONLY_FAILED : {11508,[]string{"HY000"},"Unable to disable super read only flag. Try to disable it manually"},
	ER_GRP_RPL_ENABLE_READ_ONLY_FAILED : {11509,[]string{"HY000"},"Unable to set super read only flag. Try to set it manually."},
	ER_GRP_RPL_SRV_PRIMARY_MEM : {11510,[]string{"HY000"},"This server is working as primary member."},
	ER_GRP_RPL_SRV_SECONDARY_MEM : {11511,[]string{"HY000"},"This server is working as secondary member with primary member address %s:%u."},
	ER_GRP_RPL_NO_SUITABLE_PRIMARY_MEM : {11512,[]string{"HY000"},"Unable to set any member as primary. No suitable candidate."},
	ER_GRP_RPL_SUPER_READ_ONLY_ACTIVATE_ERROR : {11513,[]string{"HY000"},"Error when activating super_read_only mode on start. The member will now exit the group."},
	ER_GRP_RPL_EXCEEDS_AUTO_INC_VALUE : {11514,[]string{"HY000"},"Group contains %lu members which is greater than group_replication_auto_increment_increment value of %lu. This can lead to a higher transactional abort rate."},
	ER_GRP_RPL_DATA_NOT_PROVIDED_BY_MEM : {11515,[]string{"HY000"},"Member with address '%s:%u' didn't provide any data during the last group change. Group information can be outdated and lead to errors on recovery."},
	ER_GRP_RPL_MEMBER_ALREADY_EXISTS : {11516,[]string{"HY000"},"There is already a member with server_uuid %s. The member will now exit the group."},
	ER_GRP_RPL_GRP_CHANGE_INFO_EXTRACT_ERROR : {11517,[]string{"HY000"},"Error when extracting information for group change. Operations and checks made to group joiners may be incomplete."},
	ER_GRP_RPL_GTID_EXECUTED_EXTRACT_ERROR : {11518,[]string{"HY000"},"Error when extracting this member GTID executed set. Operations and checks made to group joiners may be incomplete."},
	ER_GRP_RPL_GTID_SET_EXTRACT_ERROR : {11519,[]string{"HY000"},"Error when extracting this member retrieved set for its applier. Operations and checks made to group joiners may be incomplete."},
	ER_GRP_RPL_START_FAILED : {11520,[]string{"HY000"},"The START GROUP_REPLICATION command failed since the group already has 9 members."},
	ER_GRP_RPL_MEMBER_VER_INCOMPATIBLE : {11521,[]string{"HY000"},"Member version is incompatible with the group."},
	ER_GRP_RPL_TRANS_NOT_PRESENT_IN_GRP : {11522,[]string{"HY000"},"The member contains transactions not present in the group. The member will now exit the group."},
	ER_GRP_RPL_TRANS_GREATER_THAN_GRP : {11523,[]string{"HY000"},"It was not possible to assess if the member has more transactions than the group. The member will now exit the group."},
	ER_GRP_RPL_MEMBER_VERSION_LOWER_THAN_GRP : {11524,[]string{"HY000"},"Member version is lower than some group member, but since option 'group_replication_allow_local_lower_version_join is enabled, member will be allowed to join."},
	ER_GRP_RPL_LOCAL_GTID_SETS_PROCESS_ERROR : {11525,[]string{"HY000"},"Error processing local GTID sets when comparing this member transactions against the group."},
	ER_GRP_RPL_MEMBER_TRANS_GREATER_THAN_GRP : {11526,[]string{"HY000"},"This member has more executed transactions than those present in the group. Local transactions: %s > Group transactions: %s"},
	ER_GRP_RPL_BLOCK_SIZE_DIFF_FROM_GRP : {11527,[]string{"HY000"},"The member is configured with a group_replication_gtid_assignment_block_size option value '%llu' different from the group '%llu'. The member will now exit the group."},
	ER_GRP_RPL_TRANS_WRITE_SET_EXTRACT_DIFF_FROM_GRP : {11528,[]string{"HY000"},"The member is configured with a transaction-write-set-extraction option value '%s' different from the group '%s'. The member will now exit the group."},
	ER_GRP_RPL_MEMBER_CFG_INCOMPATIBLE_WITH_GRP_CFG : {11529,[]string{"HY000"},"The member configuration is not compatible with the group configuration. Variables such as group_replication_single_primary_mode or group_replication_enforce_update_everywhere_checks must have the same value on every server in the group. (member configuration option: [%s], group configuration option: [%s])."},
	ER_GRP_RPL_MEMBER_STOP_RPL_CHANNELS_ERROR : {11530,[]string{"HY000"},"Error stopping all replication channels while server was leaving the group. %s"},
	ER_GRP_RPL_PURGE_APPLIER_LOGS : {11531,[]string{"HY000"},"Detected previous RESET MASTER invocation or an issue exists in the group replication applier relay log. Purging existing applier logs."},
	ER_GRP_RPL_RESET_APPLIER_MODULE_LOGS_ERROR : {11532,[]string{"HY000"},"Unknown error occurred while resetting applier's module logs."},
	ER_GRP_RPL_APPLIER_THD_SETUP_ERROR : {11533,[]string{"HY000"},"Failed to setup the group replication applier thread."},
	ER_GRP_RPL_APPLIER_THD_START_ERROR : {11534,[]string{"HY000"},"Error while starting the group replication applier thread"},
	ER_GRP_RPL_APPLIER_THD_STOP_ERROR : {11535,[]string{"HY000"},"Failed to stop the group replication applier thread."},
	ER_GRP_RPL_FETCH_TRANS_DATA_FAILED : {11536,[]string{"HY000"},"Failed to fetch transaction data containing required transaction info for applier"},
	ER_GRP_RPL_SLAVE_IO_THD_PRIMARY_UNKNOWN : {11537,[]string{"HY000"},"Can't start slave IO THREAD of channel '%s' when group replication is running with single-primary mode and the primary member is not known."},
	ER_GRP_RPL_SALVE_IO_THD_ON_SECONDARY_MEMBER : {11538,[]string{"HY000"},"Can't start slave IO THREAD of channel '%s' when group replication is running with single-primary mode on a secondary member."},
	ER_GRP_RPL_SLAVE_SQL_THD_PRIMARY_UNKNOWN : {11539,[]string{"HY000"},"Can't start slave SQL THREAD of channel '%s' when group replication is running with single-primary mode and the primary member is not known."},
	ER_GRP_RPL_SLAVE_SQL_THD_ON_SECONDARY_MEMBER : {11540,[]string{"HY000"},"Can't start slave SQL THREAD of channel '%s' when group replication is running with single-primary mode on a secondary member."},
	ER_GRP_RPL_NEEDS_INNODB_TABLE : {11541,[]string{"HY000"},"Table %s does not use the InnoDB storage engine. This is not compatible with Group Replication."},
	ER_GRP_RPL_PRIMARY_KEY_NOT_DEFINED : {11542,[]string{"HY000"},"Table %s does not have any PRIMARY KEY. This is not compatible with Group Replication."},
	ER_GRP_RPL_FK_WITH_CASCADE_UNSUPPORTED : {11543,[]string{"HY000"},"Table %s has a foreign key with 'CASCADE', 'SET NULL' or 'SET DEFAULT' clause. This is not compatible with Group Replication."},
	ER_GRP_RPL_AUTO_INC_RESET : {11544,[]string{"HY000"},"group_replication_auto_increment_increment is reset to %lu"},
	ER_GRP_RPL_AUTO_INC_OFFSET_RESET : {11545,[]string{"HY000"},"auto_increment_offset is reset to %lu"},
	ER_GRP_RPL_AUTO_INC_SET : {11546,[]string{"HY000"},"group_replication_auto_increment_increment is set to %lu"},
	ER_GRP_RPL_AUTO_INC_OFFSET_SET : {11547,[]string{"HY000"},"auto_increment_offset is set to %lu"},
	ER_GRP_RPL_FETCH_TRANS_CONTEXT_FAILED : {11548,[]string{"HY000"},"Failed to fetch transaction context containing required transaction info for certification"},
	ER_GRP_RPL_FETCH_FORMAT_DESC_LOG_EVENT_FAILED : {11549,[]string{"HY000"},"Failed to fetch Format_description_log_event containing required server info for applier"},
	ER_GRP_RPL_FETCH_TRANS_CONTEXT_LOG_EVENT_FAILED : {11550,[]string{"HY000"},"Failed to fetch Transaction_context_log_event containing required transaction info for certification"},
	ER_GRP_RPL_FETCH_SNAPSHOT_VERSION_FAILED : {11551,[]string{"HY000"},"Failed to read snapshot version from transaction context event required for certification"},
	ER_GRP_RPL_FETCH_GTID_LOG_EVENT_FAILED : {11552,[]string{"HY000"},"Failed to fetch Gtid_log_event containing required transaction info for certification"},
	ER_GRP_RPL_UPDATE_SERV_CERTIFICATE_FAILED : {11553,[]string{"HY000"},"Unable to update certification result on server side, thread_id: %lu"},
	ER_GRP_RPL_ADD_GTID_INFO_WITH_LOCAL_GTID_FAILED : {11554,[]string{"HY000"},"Unable to add gtid information to the group_gtid_executed set when gtid was provided for local transactions"},
	ER_GRP_RPL_ADD_GTID_INFO_WITHOUT_LOCAL_GTID_FAILED : {11555,[]string{"HY000"},"Unable to add gtid information to the group_gtid_executed set when no gtid was provided for local transactions"},
	ER_GRP_RPL_NOTIFY_CERTIFICATION_OUTCOME_FAILED : {11556,[]string{"HY000"},"Failed to notify certification outcome"},
	ER_GRP_RPL_ADD_GTID_INFO_WITH_REMOTE_GTID_FAILED : {11557,[]string{"HY000"},"Unable to add gtid information to the group_gtid_executed set when gtid was provided for remote transactions"},
	ER_GRP_RPL_ADD_GTID_INFO_WITHOUT_REMOTE_GTID_FAILED : {11558,[]string{"HY000"},"Unable to add gtid information to the group_gtid_executed set when gtid was not provided for remote transactions"},
	ER_GRP_RPL_FETCH_VIEW_CHANGE_LOG_EVENT_FAILED : {11559,[]string{"HY000"},"Failed to fetch View_change_log_event containing required info for certification"},
	ER_GRP_RPL_CONTACT_WITH_SRV_FAILED : {11560,[]string{"HY000"},"Error when contacting the server to ensure the proper logging of a group change in the binlog"},
	ER_GRP_RPL_SRV_WAIT_TIME_OUT : {11561,[]string{"HY000"},"Timeout when waiting for the server to execute local transactions in order assure the group change proper logging"},
	ER_GRP_RPL_FETCH_LOG_EVENT_FAILED : {11562,[]string{"HY000"},"Failed to fetch Log_event containing required server info for applier"},
	ER_GRP_RPL_START_GRP_RPL_FAILED : {11563,[]string{"HY000"},"Unable to start Group Replication. Replication applier infrastructure is not initialized since the server was started with --initialize, --initialize-insecure or --upgrade=MINIMAL on a server upgrade."},
	ER_GRP_RPL_CONN_INTERNAL_PLUGIN_FAIL : {11564,[]string{"HY000"},"Failed to establish an internal server connection to execute plugin operations"},
	ER_GRP_RPL_SUPER_READ_ON : {11565,[]string{"HY000"},"Setting super_read_only=ON."},
	ER_GRP_RPL_SUPER_READ_OFF : {11566,[]string{"HY000"},"Setting super_read_only=OFF."},
	ER_GRP_RPL_KILLED_SESSION_ID : {11567,[]string{"HY000"},"killed session id: %d status: %d"},
	ER_GRP_RPL_KILLED_FAILED_ID : {11568,[]string{"HY000"},"killed failed id: %d failed: %d"},
	ER_GRP_RPL_INTERNAL_QUERY : {11569,[]string{"HY000"},"Internal query: %s result in error. Error number: %ld"},
	ER_GRP_RPL_COPY_FROM_EMPTY_STRING : {11570,[]string{"HY000"},"Error copying from empty string "},
	ER_GRP_RPL_QUERY_FAIL : {11571,[]string{"HY000"},"Query execution resulted in failure. errno: %d"},
	ER_GRP_RPL_CREATE_SESSION_UNABLE : {11572,[]string{"HY000"},"Unable to create a session for executing the queries on the server"},
	ER_GRP_RPL_MEMBER_NOT_FOUND : {11573,[]string{"HY000"},"The member with address %s:%u has unexpectedly disappeared, killing the current group replication recovery connection"},
	ER_GRP_RPL_MAXIMUM_CONNECTION_RETRIES_REACHED : {11574,[]string{"HY000"},"Maximum number of retries when trying to connect to a donor reached. Aborting group replication incremental recovery."},
	ER_GRP_RPL_ALL_DONORS_LEFT_ABORT_RECOVERY : {11575,[]string{"HY000"},"All donors left. Aborting group replication incremental recovery."},
	ER_GRP_RPL_ESTABLISH_RECOVERY_WITH_DONOR : {11576,[]string{"HY000"},"Establishing group recovery connection with a possible donor. Attempt %d/%d"},
	ER_GRP_RPL_ESTABLISH_RECOVERY_WITH_ANOTHER_DONOR : {11577,[]string{"HY000"},"Retrying group recovery connection with another donor. Attempt %d/%d"},
	ER_GRP_RPL_NO_VALID_DONOR : {11578,[]string{"HY000"},"No valid donors exist in the group, retrying"},
	ER_GRP_RPL_CONFIG_RECOVERY : {11579,[]string{"HY000"},"Error when configuring the asynchronous recovery channel connection to the donor."},
	ER_GRP_RPL_ESTABLISHING_CONN_GRP_REC_DONOR : {11580,[]string{"HY000"},"Establishing connection to a group replication recovery donor %s at %s port: %d."},
	ER_GRP_RPL_CREATE_GRP_RPL_REC_CHANNEL : {11581,[]string{"HY000"},"Error while creating the group replication recovery channel with donor %s at %s port: %d."},
	ER_GRP_RPL_DONOR_SERVER_CONN : {11582,[]string{"HY000"},"There was an error when connecting to the donor server. Please check that group_replication_recovery channel credentials and all MEMBER_HOST column values of performance_schema.replication_group_members table are correct and DNS resolvable."},
	ER_GRP_RPL_CHECK_STATUS_TABLE : {11583,[]string{"HY000"},"For details please check performance_schema.replication_connection_status table and error log messages of Slave I/O for channel group_replication_recovery."},
	ER_GRP_RPL_STARTING_GRP_REC : {11584,[]string{"HY000"},"Error while starting the group replication incremental recovery receiver/applier threads"},
	ER_GRP_RPL_DONOR_CONN_TERMINATION : {11585,[]string{"HY000"},"Terminating existing group replication donor connection and purging the corresponding logs."},
	ER_GRP_RPL_STOPPING_GRP_REC : {11586,[]string{"HY000"},"Error when stopping the group replication incremental recovery's donor connection"},
	ER_GRP_RPL_PURGE_REC : {11587,[]string{"HY000"},"Error when purging the group replication recovery's relay logs"},
	ER_GRP_RPL_UNABLE_TO_KILL_CONN_REC_DONOR_APPLIER : {11588,[]string{"HY000"},"Unable to kill the current group replication recovery donor connection after an applier error. Incremental recovery will shutdown."},
	ER_GRP_RPL_UNABLE_TO_KILL_CONN_REC_DONOR_FAILOVER : {11589,[]string{"HY000"},"Unable to kill the current group replication recovery donor connection during failover. Incremental recovery will shutdown."},
	ER_GRP_RPL_FAILED_TO_NOTIFY_GRP_MEMBERSHIP_EVENT : {11590,[]string{"HY000"},"Unexpected error when notifying an internal component named %s regarding a group membership event."},
	ER_GRP_RPL_FAILED_TO_BROADCAST_GRP_MEMBERSHIP_NOTIFICATION : {11591,[]string{"HY000"},"An undefined error was found while broadcasting an internal group membership notification! This is likely to happen if your components or plugins are not properly loaded or are malfunctioning!"},
	ER_GRP_RPL_FAILED_TO_BROADCAST_MEMBER_STATUS_NOTIFICATION : {11592,[]string{"HY000"},"An undefined error was found while broadcasting an internal group member status notification! This is likely to happen if your components or plugins are not properly loaded or are malfunctioning!"},
	ER_GRP_RPL_OOM_FAILED_TO_GENERATE_IDENTIFICATION_HASH : {11593,[]string{"HY000"},"No memory to generate write identification hash"},
	ER_GRP_RPL_WRITE_IDENT_HASH_BASE64_ENCODING_FAILED : {11594,[]string{"HY000"},"Base 64 encoding of the write identification hash failed"},
	ER_GRP_RPL_INVALID_BINLOG_FORMAT : {11595,[]string{"HY000"},"Binlog format should be ROW for Group Replication"},
	//OBSOLETE_ER_GRP_RPL_BINLOG_CHECKSUM_SET : {11596,[]string{"HY000"},"binlog_checksum should be NONE for Group Replication"},
	ER_GRP_RPL_TRANS_WRITE_SET_EXTRACTION_NOT_SET : {11597,[]string{"HY000"},"A transaction_write_set_extraction algorithm should be selected when running Group Replication"},
	ER_GRP_RPL_UNSUPPORTED_TRANS_ISOLATION : {11598,[]string{"HY000"},"Transaction isolation level (tx_isolation) is set to SERIALIZABLE, which is not compatible with Group Replication"},
	ER_GRP_RPL_CANNOT_EXECUTE_TRANS_WHILE_STOPPING : {11599,[]string{"HY000"},"Transaction cannot be executed while Group Replication is stopping."},
	ER_GRP_RPL_CANNOT_EXECUTE_TRANS_WHILE_RECOVERING : {11600,[]string{"HY000"},"Transaction cannot be executed while Group Replication is recovering. Try again when the server is ONLINE."},
	ER_GRP_RPL_CANNOT_EXECUTE_TRANS_IN_ERROR_STATE : {11601,[]string{"HY000"},"Transaction cannot be executed while Group Replication is on ERROR state. Check for errors and restart the plugin"},
	ER_GRP_RPL_CANNOT_EXECUTE_TRANS_IN_OFFLINE_MODE : {11602,[]string{"HY000"},"Transaction cannot be executed while Group Replication is OFFLINE. Check for errors and restart the plugin"},
	ER_GRP_RPL_MULTIPLE_CACHE_TYPE_NOT_SUPPORTED_FOR_SESSION : {11603,[]string{"HY000"},"We can only use one cache type at a time on session %u"},
	ER_GRP_RPL_FAILED_TO_REINIT_BINLOG_CACHE_FOR_READ : {11604,[]string{"HY000"},"Failed to reinit binlog cache log for read on session %u"},
	ER_GRP_RPL_FAILED_TO_CREATE_TRANS_CONTEXT : {11605,[]string{"HY000"},"Failed to create the context of the current transaction on session %u"},
	ER_GRP_RPL_FAILED_TO_EXTRACT_TRANS_WRITE_SET : {11606,[]string{"HY000"},"Failed to extract the set of items written during the execution of the current transaction on session %u"},
	ER_GRP_RPL_FAILED_TO_GATHER_TRANS_WRITE_SET : {11607,[]string{"HY000"},"Failed to gather the set of items written during the execution of the current transaction on session %u"},
	ER_GRP_RPL_TRANS_SIZE_EXCEEDS_LIMIT : {11608,[]string{"HY000"},"Error on session %u. Transaction of size %llu exceeds specified limit %lu. To increase the limit please adjust group_replication_transaction_size_limit option."},
	//OBSOLETE_ER_GRP_RPL_REINIT_OF_INTERNAL_CACHE_FOR_READ_FAILED : {11609,[]string{"HY000"},"Error while re-initializing an internal cache, for read operations, on session %u"},
	//OBSOLETE_ER_GRP_RPL_APPENDING_DATA_TO_INTERNAL_CACHE_FAILED : {11610,[]string{"HY000"},"Error while appending data to an internal cache on session %u"},
	ER_GRP_RPL_WRITE_TO_TRANSACTION_MESSAGE_FAILED : {11611,[]string{"HY000"},"Error while writing to transaction message on session %u"},
	ER_GRP_RPL_FAILED_TO_REGISTER_TRANS_OUTCOME_NOTIFICTION : {11612,[]string{"HY000"},"Unable to register for getting notifications regarding the outcome of the transaction on session %u"},
	ER_GRP_RPL_MSG_TOO_LONG_BROADCASTING_TRANS_FAILED : {11613,[]string{"HY000"},"Error broadcasting transaction to the group on session %u. Message is too big."},
	ER_GRP_RPL_BROADCASTING_TRANS_TO_GRP_FAILED : {11614,[]string{"HY000"},"Error while broadcasting the transaction to the group on session %u"},
	ER_GRP_RPL_ERROR_WHILE_WAITING_FOR_CONFLICT_DETECTION : {11615,[]string{"HY000"},"Error while waiting for conflict detection procedure to finish on session %u"},
	//OBSOLETE_ER_GRP_RPL_REINIT_OF_INTERNAL_CACHE_FOR_WRITE_FAILED : {11616,[]string{"HY000"},"Error while re-initializing an internal cache, for write operations, on session %u"},
	//OBSOLETE_ER_GRP_RPL_FAILED_TO_CREATE_COMMIT_CACHE : {11617,[]string{"HY000"},"Failed to create group replication commit cache on session %u"},
	//OBSOLETE_ER_GRP_RPL_REINIT_OF_COMMIT_CACHE_FOR_WRITE_FAILED : {11618,[]string{"HY000"},"Failed to reinit group replication commit cache for write on session %u"},
	//OBSOLETE_ER_GRP_RPL_PREV_REC_SESSION_RUNNING : {11619,[]string{"HY000"},"A previous recovery session is still running. Please stop the group replication plugin and wait for it to stop"},
	ER_GRP_RPL_FATAL_REC_PROCESS : {11620,[]string{"HY000"},"Fatal error during the incremental recovery process of Group Replication. The server will leave the group."},
	//OBSOLETE_ER_GRP_RPL_WHILE_STOPPING_REP_CHANNEL : {11621,[]string{"HY000"},"Error stopping all replication channels while server was leaving the group. %s"},
	ER_GRP_RPL_UNABLE_TO_EVALUATE_APPLIER_STATUS : {11622,[]string{"HY000"},"Unable to evaluate the group replication applier execution status. Group replication recovery will shutdown to avoid data corruption."},
	ER_GRP_RPL_ONLY_ONE_SERVER_ALIVE : {11623,[]string{"HY000"},"Only one server alive. Declaring this server as online within the replication group"},
	ER_GRP_RPL_CERTIFICATION_REC_PROCESS : {11624,[]string{"HY000"},"Error when processing certification information in the incremental recovery process"},
	ER_GRP_RPL_UNABLE_TO_ENSURE_EXECUTION_REC : {11625,[]string{"HY000"},"Unable to ensure the execution of group transactions received during recovery."},
	ER_GRP_RPL_WHILE_SENDING_MSG_REC : {11626,[]string{"HY000"},"Error while sending message in the group replication incremental recovery process."},
	ER_GRP_RPL_READ_UNABLE_FOR_SUPER_READ_ONLY : {11627,[]string{"HY000"},"Unable to read the server value for the super_read_only variable."},
	ER_GRP_RPL_READ_UNABLE_FOR_READ_ONLY_SUPER_READ_ONLY : {11628,[]string{"HY000"},"Unable to read the server values for the read_only and super_read_only variables."},
	ER_GRP_RPL_UNABLE_TO_RESET_SERVER_READ_MODE : {11629,[]string{"HY000"},"Unable to reset the server read mode settings. Try to reset them manually."},
	ER_GRP_RPL_UNABLE_TO_CERTIFY_PLUGIN_TRANS : {11630,[]string{"HY000"},"Due to a plugin error, some transactions were unable to be certified and will now rollback."},
	ER_GRP_RPL_UNBLOCK_CERTIFIED_TRANS : {11631,[]string{"HY000"},"Error when trying to unblock non certified or consistent transactions. Check for consistency errors when restarting the service"},
	//OBSOLETE_ER_GRP_RPL_SERVER_WORKING_AS_SECONDARY : {11632,[]string{"HY000"},"This server is working as secondary member with primary member address %s:%u."},
	ER_GRP_RPL_FAILED_TO_START_WITH_INVALID_SERVER_ID : {11633,[]string{"HY000"},"Unable to start Group Replication. Replication applier infrastructure is not initialized since the server was started with server_id=0. Please, restart the server with server_id larger than 0."},
	ER_GRP_RPL_FORCE_MEMBERS_MUST_BE_EMPTY : {11634,[]string{"HY000"},"group_replication_force_members must be empty on group start. Current value: '%s'"},
	ER_GRP_RPL_PLUGIN_STRUCT_INIT_NOT_POSSIBLE_ON_SERVER_START : {11635,[]string{"HY000"},"It was not possible to guarantee the initialization of plugin structures on server start"},
	ER_GRP_RPL_FAILED_TO_ENABLE_SUPER_READ_ONLY_MODE : {11636,[]string{"HY000"},"Could not enable the server read only mode and guarantee a safe recovery execution"},
	ER_GRP_RPL_FAILED_TO_INIT_COMMUNICATION_ENGINE : {11637,[]string{"HY000"},"Error on group communication engine initialization"},
	ER_GRP_RPL_FAILED_TO_START_ON_SECONDARY_WITH_ASYNC_CHANNELS : {11638,[]string{"HY000"},"Can't start group replication on secondary member with single-primary mode while asynchronous replication channels are running."},
	ER_GRP_RPL_FAILED_TO_START_COMMUNICATION_ENGINE : {11639,[]string{"HY000"},"Error on group communication engine start"},
	ER_GRP_RPL_TIMEOUT_ON_VIEW_AFTER_JOINING_GRP : {11640,[]string{"HY000"},"Timeout on wait for view after joining group"},
	ER_GRP_RPL_FAILED_TO_CALL_GRP_COMMUNICATION_INTERFACE : {11641,[]string{"HY000"},"Error calling group communication interfaces"},
	ER_GRP_RPL_MEMBER_SERVER_UUID_IS_INCOMPATIBLE_WITH_GRP : {11642,[]string{"HY000"},"Member server_uuid is incompatible with the group. Server_uuid %s matches group_replication_group_name %s."},
	ER_GRP_RPL_MEMBER_CONF_INFO : {11643,[]string{"HY000"},"Member configuration: member_id: %lu; member_uuid: \"%s\"; single-primary mode: \"%s\"; group_replication_auto_increment_increment: %lu; "},
	ER_GRP_RPL_FAILED_TO_CONFIRM_IF_SERVER_LEFT_GRP : {11644,[]string{"HY000"},"Unable to confirm whether the server has left the group or not. Check performance_schema.replication_group_members to check group membership information."},
	ER_GRP_RPL_SERVER_IS_ALREADY_LEAVING : {11645,[]string{"HY000"},"Skipping leave operation: concurrent attempt to leave the group is on-going."},
	ER_GRP_RPL_SERVER_ALREADY_LEFT : {11646,[]string{"HY000"},"Skipping leave operation: member already left the group."},
	ER_GRP_RPL_WAITING_FOR_VIEW_UPDATE : {11647,[]string{"HY000"},"Going to wait for view modification"},
	ER_GRP_RPL_TIMEOUT_RECEIVING_VIEW_CHANGE_ON_SHUTDOWN : {11648,[]string{"HY000"},"While leaving the group due to a stop, shutdown or failure there was a timeout receiving a view change. This can lead to a possible inconsistent state. Check the log for more details"},
	ER_GRP_RPL_REQUESTING_NON_MEMBER_SERVER_TO_LEAVE : {11649,[]string{"HY000"},"Requesting to leave the group despite of not being a member"},
	ER_GRP_RPL_IS_STOPPING : {11650,[]string{"HY000"},"Plugin 'group_replication' is stopping."},
	ER_GRP_RPL_IS_STOPPED : {11651,[]string{"HY000"},"Plugin 'group_replication' has been stopped."},
	ER_GRP_RPL_FAILED_TO_ENABLE_READ_ONLY_MODE_ON_SHUTDOWN : {11652,[]string{"HY000"},"On plugin shutdown it was not possible to enable the server read only mode. Local transactions will be accepted and committed."},
	ER_GRP_RPL_RECOVERY_MODULE_TERMINATION_TIMED_OUT_ON_SHUTDOWN : {11653,[]string{"HY000"},"On shutdown there was a timeout on the Group Replication recovery module termination. Check the log for more details"},
	ER_GRP_RPL_APPLIER_TERMINATION_TIMED_OUT_ON_SHUTDOWN : {11654,[]string{"HY000"},"On shutdown there was a timeout on the Group Replication applier termination."},
	ER_GRP_RPL_FAILED_TO_SHUTDOWN_REGISTRY_MODULE : {11655,[]string{"HY000"},"Unexpected failure while shutting down registry module!"},
	ER_GRP_RPL_FAILED_TO_INIT_HANDLER : {11656,[]string{"HY000"},"Failure during Group Replication handler initialization"},
	ER_GRP_RPL_FAILED_TO_REGISTER_SERVER_STATE_OBSERVER : {11657,[]string{"HY000"},"Failure when registering the server state observers"},
	ER_GRP_RPL_FAILED_TO_REGISTER_TRANS_STATE_OBSERVER : {11658,[]string{"HY000"},"Failure when registering the transactions state observers"},
	ER_GRP_RPL_FAILED_TO_REGISTER_BINLOG_STATE_OBSERVER : {11659,[]string{"HY000"},"Failure when registering the binlog state observers"},
	ER_GRP_RPL_FAILED_TO_START_ON_BOOT : {11660,[]string{"HY000"},"Unable to start Group Replication on boot"},
	ER_GRP_RPL_FAILED_TO_STOP_ON_PLUGIN_UNINSTALL : {11661,[]string{"HY000"},"Failure when stopping Group Replication on plugin uninstall"},
	ER_GRP_RPL_FAILED_TO_UNREGISTER_SERVER_STATE_OBSERVER : {11662,[]string{"HY000"},"Failure when unregistering the server state observers"},
	ER_GRP_RPL_FAILED_TO_UNREGISTER_TRANS_STATE_OBSERVER : {11663,[]string{"HY000"},"Failure when unregistering the transactions state observers"},
	ER_GRP_RPL_FAILED_TO_UNREGISTER_BINLOG_STATE_OBSERVER : {11664,[]string{"HY000"},"Failure when unregistering the binlog state observers"},
	ER_GRP_RPL_ALL_OBSERVERS_UNREGISTERED : {11665,[]string{"HY000"},"All Group Replication server observers have been successfully unregistered"},
	ER_GRP_RPL_FAILED_TO_PARSE_THE_GRP_NAME : {11666,[]string{"HY000"},"Unable to parse the group_replication_group_name."},
	ER_GRP_RPL_FAILED_TO_GENERATE_SIDNO_FOR_GRP : {11667,[]string{"HY000"},"Unable to parse the group_replication_group_name."},
	ER_GRP_RPL_APPLIER_NOT_STARTED_DUE_TO_RUNNING_PREV_SHUTDOWN : {11668,[]string{"HY000"},"Cannot start the Group Replication applier as a previous shutdown is still running: The thread will stop once its task is complete."},
	ER_GRP_RPL_FAILED_TO_INIT_APPLIER_MODULE : {11669,[]string{"HY000"},"Unable to initialize the Group Replication applier module."},
	ER_GRP_RPL_APPLIER_INITIALIZED : {11670,[]string{"HY000"},"Group Replication applier module successfully initialized!"},
	ER_GRP_RPL_COMMUNICATION_SSL_CONF_INFO : {11671,[]string{"HY000"},"Group communication SSL configuration: group_replication_ssl_mode: \"%s\"; server_key_file: \"%s\"; server_cert_file: \"%s\"; client_key_file: \"%s\"; client_cert_file: \"%s\"; ca_file: \"%s\"; ca_path: \"%s\"; cipher: \"%s\"; tls_version: \"%s\"; tls_ciphersuites: \"%s\"; crl_file: \"%s\"; crl_path: \"%s\"; ssl_fips_mode: \"%s\""},
	ER_GRP_RPL_ABORTS_AS_SSL_NOT_SUPPORTED_BY_MYSQLD : {11672,[]string{"HY000"},"MySQL server does not have SSL support and group_replication_ssl_mode is \"%s\", START GROUP_REPLICATION will abort"},
	ER_GRP_RPL_SSL_DISABLED : {11673,[]string{"HY000"},"Group communication SSL configuration: group_replication_ssl_mode: \"%s\""},
	ER_GRP_RPL_UNABLE_TO_INIT_COMMUNICATION_ENGINE : {11674,[]string{"HY000"},"Unable to initialize the group communication engine"},
	ER_GRP_RPL_BINLOG_DISABLED : {11675,[]string{"HY000"},"Binlog must be enabled for Group Replication"},
	ER_GRP_RPL_GTID_MODE_OFF : {11676,[]string{"HY000"},"Gtid mode should be ON for Group Replication"},
	ER_GRP_RPL_LOG_SLAVE_UPDATES_NOT_SET : {11677,[]string{"HY000"},"LOG_SLAVE_UPDATES should be ON for Group Replication"},
	ER_GRP_RPL_INVALID_TRANS_WRITE_SET_EXTRACTION_VALUE : {11678,[]string{"HY000"},"Extraction of transaction write sets requires an hash algorithm configuration. Please, double check that the parameter transaction-write-set-extraction is set to a valid algorithm."},
	ER_GRP_RPL_RELAY_LOG_INFO_REPO_MUST_BE_TABLE : {11679,[]string{"HY000"},"Relay log info repository must be set to TABLE"},
	ER_GRP_RPL_MASTER_INFO_REPO_MUST_BE_TABLE : {11680,[]string{"HY000"},"Master info repository must be set to TABLE."},
	ER_GRP_RPL_INCORRECT_TYPE_SET_FOR_PARALLEL_APPLIER : {11681,[]string{"HY000"},"In order to use parallel applier on Group Replication, parameter slave-parallel-type must be set to 'LOGICAL_CLOCK'."},
	ER_GRP_RPL_SLAVE_PRESERVE_COMMIT_ORDER_NOT_SET : {11682,[]string{"HY000"},"Group Replication requires slave-preserve-commit-order to be set to ON when using more than 1 applier threads."},
	ER_GRP_RPL_SINGLE_PRIM_MODE_NOT_ALLOWED_WITH_UPDATE_EVERYWHERE : {11683,[]string{"HY000"},"It is not allowed to run single primary mode with 'group_replication_enforce_update_everywhere_checks' enabled."},
	ER_GRP_RPL_MODULE_TERMINATE_ERROR : {11684,[]string{"HY000"},"error_message: %s"},
	ER_GRP_RPL_GRP_NAME_OPTION_MANDATORY : {11685,[]string{"HY000"},"The group_replication_group_name option is mandatory"},
	ER_GRP_RPL_GRP_NAME_IS_TOO_LONG : {11686,[]string{"HY000"},"The group_replication_group_name '%s' is not a valid UUID, its length is too big"},
	ER_GRP_RPL_GRP_NAME_IS_NOT_VALID_UUID : {11687,[]string{"HY000"},"The group_replication_group_name '%s' is not a valid UUID"},
	ER_GRP_RPL_FLOW_CTRL_MIN_QUOTA_GREATER_THAN_MAX_QUOTA : {11688,[]string{"HY000"},"group_replication_flow_control_min_quota cannot be larger than group_replication_flow_control_max_quota"},
	ER_GRP_RPL_FLOW_CTRL_MIN_RECOVERY_QUOTA_GREATER_THAN_MAX_QUOTA : {11689,[]string{"HY000"},"group_replication_flow_control_min_recovery_quota cannot be larger than group_replication_flow_control_max_quota"},
	ER_GRP_RPL_FLOW_CTRL_MAX_QUOTA_SMALLER_THAN_MIN_QUOTAS : {11690,[]string{"HY000"},"group_replication_flow_control_max_quota cannot be smaller than group_replication_flow_control_min_quota or group_replication_flow_control_min_recovery_quota"},
	ER_GRP_RPL_INVALID_SSL_RECOVERY_STRING : {11691,[]string{"HY000"},"The given value for recovery ssl option 'group_replication_%s' is invalid as its length is beyond the limit"},
	ER_GRP_RPL_SUPPORTS_ONLY_ONE_FORCE_MEMBERS_SET : {11692,[]string{"HY000"},"There is one group_replication_force_members operation already ongoing"},
	ER_GRP_RPL_FORCE_MEMBERS_SET_UPDATE_NOT_ALLOWED : {11693,[]string{"HY000"},"group_replication_force_members can only be updated when Group Replication is running and a majority of the members are unreachable"},
	ER_GRP_RPL_GRP_COMMUNICATION_INIT_WITH_CONF : {11694,[]string{"HY000"},"Initialized group communication with configuration: group_replication_group_name: '%s'; group_replication_local_address: '%s'; group_replication_group_seeds: '%s'; group_replication_bootstrap_group: '%s'; group_replication_poll_spin_loops: %lu; group_replication_compression_threshold: %lu; group_replication_ip_allowlist: '%s'; group_replication_communication_debug_options: '%s'; group_replication_member_expel_timeout: '%lu'; group_replication_communication_max_message_size: %lu; group_replication_message_cache_size: '%luu'"},
	ER_GRP_RPL_UNKNOWN_GRP_RPL_APPLIER_PIPELINE_REQUESTED : {11695,[]string{"HY000"},"Unknown group replication applier pipeline requested"},
	ER_GRP_RPL_FAILED_TO_BOOTSTRAP_EVENT_HANDLING_INFRASTRUCTURE : {11696,[]string{"HY000"},"Unable to bootstrap group replication event handling infrastructure. Unknown handler type: %d"},
	ER_GRP_RPL_APPLIER_HANDLER_NOT_INITIALIZED : {11697,[]string{"HY000"},"One of the group replication applier handlers is null due to an initialization error"},
	ER_GRP_RPL_APPLIER_HANDLER_IS_IN_USE : {11698,[]string{"HY000"},"A group replication applier handler, marked as unique, is already in use."},
	ER_GRP_RPL_APPLIER_HANDLER_ROLE_IS_IN_USE : {11699,[]string{"HY000"},"A group replication applier handler role, that was marked as unique, is already in use."},
	ER_GRP_RPL_FAILED_TO_INIT_APPLIER_HANDLER : {11700,[]string{"HY000"},"Error on group replication applier handler initialization"},
	ER_GRP_RPL_SQL_SERVICE_FAILED_TO_INIT_SESSION_THREAD : {11701,[]string{"HY000"},"Error when initializing a session thread for internal server connection."},
	ER_GRP_RPL_SQL_SERVICE_COMM_SESSION_NOT_INITIALIZED : {11702,[]string{"HY000"},"Error running internal SQL query: %s. The internal server communication session is not initialized"},
	ER_GRP_RPL_SQL_SERVICE_SERVER_SESSION_KILLED : {11703,[]string{"HY000"},"Error running internal SQL query: %s. The internal server session was killed or server is shutting down."},
	ER_GRP_RPL_SQL_SERVICE_FAILED_TO_RUN_SQL_QUERY : {11704,[]string{"HY000"},"Error running internal SQL query: %s. Got internal SQL error: %s(%d)"},
	ER_GRP_RPL_SQL_SERVICE_SERVER_INTERNAL_FAILURE : {11705,[]string{"HY000"},"Error running internal SQL query: %s. Internal failure."},
	ER_GRP_RPL_SQL_SERVICE_RETRIES_EXCEEDED_ON_SESSION_STATE : {11706,[]string{"HY000"},"Error, maximum number of retries exceeded when waiting for the internal server session state to be operating"},
	ER_GRP_RPL_SQL_SERVICE_FAILED_TO_FETCH_SECURITY_CTX : {11707,[]string{"HY000"},"Error when trying to fetch security context when contacting the server for internal plugin requests."},
	ER_GRP_RPL_SQL_SERVICE_SERVER_ACCESS_DENIED_FOR_USER : {11708,[]string{"HY000"},"There was an error when trying to access the server with user: %s. Make sure the user is present in the server and that the MySQL upgrade procedure was run correctly."},
	ER_GRP_RPL_SQL_SERVICE_MAX_CONN_ERROR_FROM_SERVER : {11709,[]string{"HY000"},"Failed to establish an internal server connection to execute plugin operations since the server does not have available connections, please increase @@GLOBAL.MAX_CONNECTIONS. Server error: %i."},
	ER_GRP_RPL_SQL_SERVICE_SERVER_ERROR_ON_CONN : {11710,[]string{"HY000"},"Failed to establish an internal server connection to execute plugin operations. Server error: %i. Server error message: %s"},
	ER_GRP_RPL_UNREACHABLE_MAJORITY_TIMEOUT_FOR_MEMBER : {11711,[]string{"HY000"},"This member could not reach a majority of the members for more than %ld seconds. The member will now leave the group as instructed by the group_replication_unreachable_majority_timeout option."},
	ER_GRP_RPL_SERVER_SET_TO_READ_ONLY_DUE_TO_ERRORS : {11712,[]string{"HY000"},"The server was automatically set into read only mode after an error was detected."},
	ER_GRP_RPL_GMS_LISTENER_FAILED_TO_LOG_NOTIFICATION : {11713,[]string{"HY000"},"Unable to log notification to table (errno: %lu) (res: %d)! Message: %s"},
	ER_GRP_RPL_GRP_COMMUNICATION_ENG_INIT_FAILED : {11714,[]string{"HY000"},"Failure in group communication engine '%s' initialization"},
	ER_GRP_RPL_SET_GRP_COMMUNICATION_ENG_LOGGER_FAILED : {11715,[]string{"HY000"},"Unable to set the group communication engine logger"},
	ER_GRP_RPL_DEBUG_OPTIONS : {11716,[]string{"HY000"},"Current debug options are: '%s'."},
	ER_GRP_RPL_INVALID_DEBUG_OPTIONS : {11717,[]string{"HY000"},"Some debug options in '%s' are not valid."},
	ER_GRP_RPL_EXIT_GRP_GCS_ERROR : {11718,[]string{"HY000"},"Error calling group communication interfaces while trying to leave the group"},
	ER_GRP_RPL_GRP_MEMBER_OFFLINE : {11719,[]string{"HY000"},"Member is not ONLINE, it is not possible to force a new group membership"},
	ER_GRP_RPL_GCS_INTERFACE_ERROR : {11720,[]string{"HY000"},"Error calling group communication interfaces"},
	ER_GRP_RPL_FORCE_MEMBER_VALUE_SET_ERROR : {11721,[]string{"HY000"},"Error setting group_replication_force_members value '%s' on group communication interfaces"},
	ER_GRP_RPL_FORCE_MEMBER_VALUE_SET : {11722,[]string{"HY000"},"The group_replication_force_members value '%s' was set in the group communication interfaces"},
	ER_GRP_RPL_FORCE_MEMBER_VALUE_TIME_OUT : {11723,[]string{"HY000"},"Timeout on wait for view after setting group_replication_force_members value '%s' into group communication interfaces"},
	ER_GRP_RPL_BROADCAST_COMMIT_MSSG_TOO_BIG : {11724,[]string{"HY000"},"Broadcast of committed transactions message failed. Message is too big."},
	ER_GRP_RPL_SEND_STATS_ERROR : {11725,[]string{"HY000"},"Error while sending stats message"},
	ER_GRP_RPL_MEMBER_STATS_INFO : {11726,[]string{"HY000"},"Flow control - update member stats: %s stats certifier_queue %d, applier_queue %d certified %ld (%ld), applied %ld (%ld), local %ld (%ld), quota %ld (%ld) mode=%d"},
	ER_GRP_RPL_FLOW_CONTROL_STATS : {11727,[]string{"HY000"},"Flow control: throttling to %ld commits per %ld sec, with %d writing and %d non-recovering members, min capacity %lld, lim throttle %lld"},
	ER_GRP_RPL_UNABLE_TO_CONVERT_PACKET_TO_EVENT : {11728,[]string{"HY000"},"Unable to convert a packet into an event on the applier. Error: %s"},
	ER_GRP_RPL_PIPELINE_CREATE_FAILED : {11729,[]string{"HY000"},"Failed to create group replication pipeline cache."},
	ER_GRP_RPL_PIPELINE_REINIT_FAILED_WRITE : {11730,[]string{"HY000"},"Failed to reinit group replication pipeline cache for write."},
	ER_GRP_RPL_UNABLE_TO_CONVERT_EVENT_TO_PACKET : {11731,[]string{"HY000"},"Unable to convert the event into a packet on the applier. Error: %s"},
	ER_GRP_RPL_PIPELINE_FLUSH_FAIL : {11732,[]string{"HY000"},"Failed to flush group replication pipeline cache."},
	ER_GRP_RPL_PIPELINE_REINIT_FAILED_READ : {11733,[]string{"HY000"},"Failed to reinit group replication pipeline cache for read."},
	//OBSOLETE_ER_GRP_RPL_STOP_REP_CHANNEL : {11734,[]string{"HY000"},"Error stopping all replication channels while server was leaving the group. Got error: %d. Please check the error log for more details."},
	ER_GRP_RPL_GCS_GR_ERROR_MSG : {11735,[]string{"HY000"},"%s"},
	ER_GRP_RPL_SLAVE_IO_THREAD_UNBLOCKED : {11736,[]string{"HY000"},"The slave IO thread of channel '%s' is unblocked as the member is declared ONLINE now."},
	ER_GRP_RPL_SLAVE_IO_THREAD_ERROR_OUT : {11737,[]string{"HY000"},"The slave IO thread of channel '%s' will error out as the member failed to come ONLINE."},
	ER_GRP_RPL_SLAVE_APPLIER_THREAD_UNBLOCKED : {11738,[]string{"HY000"},"The slave applier thread of channel '%s' is unblocked as the member is declared ONLINE now."},
	ER_GRP_RPL_SLAVE_APPLIER_THREAD_ERROR_OUT : {11739,[]string{"HY000"},"The slave applier thread of channel '%s' will error out as the member failed to come ONLINE."},
	ER_LDAP_AUTH_FAILED_TO_CREATE_OR_GET_CONNECTION : {11740,[]string{"HY000"},"LDAP authentication initialize: failed to create/ get connection from the pool. "},
	ER_LDAP_AUTH_DEINIT_FAILED : {11741,[]string{"HY000"},"LDAP authentication de_initialize Failed"},
	ER_LDAP_AUTH_SKIPPING_USER_GROUP_SEARCH : {11742,[]string{"HY000"},"Skipping group search, No group attribute mentioned"},
	ER_LDAP_AUTH_POOL_DISABLE_MAX_SIZE_ZERO : {11743,[]string{"HY000"},"Pool max size is 0, connection pool is disabled"},
	ER_LDAP_AUTH_FAILED_TO_CREATE_LDAP_OBJECT_CREATOR : {11744,[]string{"HY000"},"Connection pool initialization, failed to create LDAP object creator"},
	ER_LDAP_AUTH_FAILED_TO_CREATE_LDAP_OBJECT : {11745,[]string{"HY000"},"Connection pool initialization, failed to create LDAP object"},
	ER_LDAP_AUTH_TLS_CONF : {11746,[]string{"HY000"},"LDAP TLS configuration"},
	ER_LDAP_AUTH_TLS_CONNECTION : {11747,[]string{"HY000"},"LDAP TLS connection"},
	ER_LDAP_AUTH_CONN_POOL_NOT_CREATED : {11748,[]string{"HY000"},"LDAP pool is not created."},
	ER_LDAP_AUTH_CONN_POOL_INITIALIZING : {11749,[]string{"HY000"},"LDAP pool is initializing"},
	ER_LDAP_AUTH_CONN_POOL_DEINITIALIZING : {11750,[]string{"HY000"},"LDAP pool is de-initializing"},
	ER_LDAP_AUTH_ZERO_MAX_POOL_SIZE_UNCHANGED : {11751,[]string{"HY000"},"Pool max size old and new values are 0"},
	ER_LDAP_AUTH_POOL_REINITIALIZING : {11752,[]string{"HY000"},"LDAP pool is re-initializing"},
	ER_LDAP_AUTH_FAILED_TO_WRITE_PACKET : {11753,[]string{"HY000"},"Plug-in has failed to write the packet."},
	ER_LDAP_AUTH_SETTING_USERNAME : {11754,[]string{"HY000"},"Setting LDAP user name as : %s"},
	ER_LDAP_AUTH_USER_AUTH_DATA : {11755,[]string{"HY000"},"User authentication data: %s size: %lu"},
	ER_LDAP_AUTH_INFO_FOR_USER : {11756,[]string{"HY000"},"User is authenticated as: %s external user: %s"},
	ER_LDAP_AUTH_USER_GROUP_SEARCH_INFO : {11757,[]string{"HY000"},"Group search information base DN: %s scope: %d filter: %s attribute: %s"},
	ER_LDAP_AUTH_GRP_SEARCH_SPECIAL_HDL : {11758,[]string{"HY000"},"Special handling for group search, {GA} found"},
	ER_LDAP_AUTH_GRP_IS_FULL_DN : {11759,[]string{"HY000"},"Group search special handling, group full DN found. "},
	ER_LDAP_AUTH_USER_NOT_FOUND_IN_ANY_GRP : {11760,[]string{"HY000"},"User %s is not member of any group."},
	ER_LDAP_AUTH_USER_FOUND_IN_MANY_GRPS : {11761,[]string{"HY000"},"User %s is member of more than one group"},
	ER_LDAP_AUTH_USER_HAS_MULTIPLE_GRP_NAMES : {11762,[]string{"HY000"},"For user %s has multiple user group names. Please check if group attribute name is correct"},
	ER_LDAP_AUTH_SEARCHED_USER_GRP_NAME : {11763,[]string{"HY000"},"Searched group name: %s"},
	ER_LDAP_AUTH_OBJECT_CREATE_TIMESTAMP : {11764,[]string{"HY000"},"LDAP authentication object creation time_stamp: %s dn: %s"},
	ER_LDAP_AUTH_CERTIFICATE_NAME : {11765,[]string{"HY000"},"Certificate name: %s"},
	ER_LDAP_AUTH_FAILED_TO_POOL_DEINIT : {11766,[]string{"HY000"},"Failed to pool de-initialized: pool is already reconstructing"},
	ER_LDAP_AUTH_FAILED_TO_INITIALIZE_POOL_IN_RECONSTRUCTING : {11767,[]string{"HY000"},"Pool initialization failed: pool is already initialized"},
	ER_LDAP_AUTH_FAILED_TO_INITIALIZE_POOL_IN_INIT_STATE : {11768,[]string{"HY000"},"Pool initialization failed: pool is initializing"},
	ER_LDAP_AUTH_FAILED_TO_INITIALIZE_POOL_IN_DEINIT_STATE : {11769,[]string{"HY000"},"Pool initialization failed: pool is de-initializing"},
	ER_LDAP_AUTH_FAILED_TO_DEINITIALIZE_POOL_IN_RECONSTRUCT_STATE : {11770,[]string{"HY000"},"Failed to pool deinitialized: pool is already reconstructing"},
	ER_LDAP_AUTH_FAILED_TO_DEINITIALIZE_NOT_READY_POOL : {11771,[]string{"HY000"},"Failed to pool deinitialized : pool is not ready"},
	ER_LDAP_AUTH_FAILED_TO_GET_CONNECTION_AS_PLUGIN_NOT_READY : {11772,[]string{"HY000"},"Ldap_connection_pool::get: Failed to return connection as plug-in is not ready/initializing/de-initializing"},
	ER_LDAP_AUTH_CONNECTION_POOL_INIT_FAILED : {11773,[]string{"HY000"},"Connection pool has failed to initialized"},
	ER_LDAP_AUTH_MAX_ALLOWED_CONNECTION_LIMIT_HIT : {11774,[]string{"HY000"},"Ldap_connetion_pool::get LDAP maximum connection allowed size is reached. Increase the maximum limit."},
	ER_LDAP_AUTH_MAX_POOL_SIZE_SET_FAILED : {11775,[]string{"HY000"},"Set max pool size failed."},
	ER_LDAP_AUTH_PLUGIN_FAILED_TO_READ_PACKET : {11776,[]string{"HY000"},"Plug-in has failed to read the packet from client"},
	ER_LDAP_AUTH_CREATING_LDAP_CONNECTION : {11777,[]string{"HY000"},"Ldap_authentication::initialize: creating new LDAP connection. "},
	ER_LDAP_AUTH_GETTING_CONNECTION_FROM_POOL : {11778,[]string{"HY000"},"Ldap_authentication::initialize: getting connection from pool. "},
	ER_LDAP_AUTH_RETURNING_CONNECTION_TO_POOL : {11779,[]string{"HY000"},"Ldap_authentication::de_initialize putting back connection in the pool"},
	ER_LDAP_AUTH_SEARCH_USER_GROUP_ATTR_NOT_FOUND : {11780,[]string{"HY000"},"Ldap_authentication::search_user_group no group attribute found"},
	ER_LDAP_AUTH_LDAP_INFO_NULL : {11781,[]string{"HY000"},"Ldap_connetion_pool::put ldap info null"},
	ER_LDAP_AUTH_FREEING_CONNECTION : {11782,[]string{"HY000"},"Ldap_connection_pool::put connection is freeing. "},
	ER_LDAP_AUTH_CONNECTION_PUSHED_TO_POOL : {11783,[]string{"HY000"},"Ldap_connection_pool::put connection in pushed in the pool"},
	ER_LDAP_AUTH_CONNECTION_CREATOR_ENTER : {11784,[]string{"HY000"},"Ldap_connection_creator::Ldap_connection_creator"},
	ER_LDAP_AUTH_STARTING_TLS : {11785,[]string{"HY000"},"starting TLS"},
	ER_LDAP_AUTH_CONNECTION_GET_LDAP_INFO_NULL : {11786,[]string{"HY000"},"Ldap_connection_pool::get: (ldap_info == NULL)|| (*ldap_info)"},
	ER_LDAP_AUTH_DELETING_CONNECTION_KEY : {11787,[]string{"HY000"},"Ldap_connection_pool::deinit: deleting connection key %s"},
	ER_LDAP_AUTH_POOLED_CONNECTION_KEY : {11788,[]string{"HY000"}," Ldap_connection_pool::get pooled connection key: %s"},
	ER_LDAP_AUTH_CREATE_CONNECTION_KEY : {11789,[]string{"HY000"},"Ldap_connection_pool::get create connection key: %s"},
	ER_LDAP_AUTH_COMMUNICATION_HOST_INFO : {11790,[]string{"HY000"},"LDAP communication host %s port %u"},
	ER_LDAP_AUTH_METHOD_TO_CLIENT : {11791,[]string{"HY000"},"Sending authentication method to client : %s"},
	ER_LDAP_AUTH_SASL_REQUEST_FROM_CLIENT : {11792,[]string{"HY000"},"SASL request received from mysql client: %s"},
	ER_LDAP_AUTH_SASL_PROCESS_SASL : {11793,[]string{"HY000"},"Ldap_sasl_authentication::process_sasl rc: %s"},
	ER_LDAP_AUTH_SASL_BIND_SUCCESS_INFO : {11794,[]string{"HY000"},"Ldap_sasl_authentication::process_sasl sasl bind succeed. dn: %s method: %s server credential: %s"},
	ER_LDAP_AUTH_STARTED_FOR_USER : {11795,[]string{"HY000"},"LDAP authentication started for user name: %s"},
	ER_LDAP_AUTH_DISTINGUISHED_NAME : {11796,[]string{"HY000"},"%s"},
	ER_LDAP_AUTH_INIT_FAILED : {11797,[]string{"HY000"},"LDAP authentication initialize is failed with: %s"},
	ER_LDAP_AUTH_OR_GROUP_RETRIEVAL_FAILED : {11798,[]string{"HY000"},"LDAP authentication failed or group retrieval failed: %s"},
	ER_LDAP_AUTH_USER_GROUP_SEARCH_FAILED : {11799,[]string{"HY000"},"Search user group has failed: %s"},
	ER_LDAP_AUTH_USER_BIND_FAILED : {11800,[]string{"HY000"},"LDAP user bind has failed: %s"},
	ER_LDAP_AUTH_POOL_GET_FAILED_TO_CREATE_CONNECTION : {11801,[]string{"HY000"},"Connection pool get: Failed to create LDAP connection. %s"},
	ER_LDAP_AUTH_FAILED_TO_CREATE_LDAP_CONNECTION : {11802,[]string{"HY000"},"Failed to create new LDAP connection:  %s"},
	ER_LDAP_AUTH_FAILED_TO_ESTABLISH_TLS_CONNECTION : {11803,[]string{"HY000"},"Failed to establish TLS connection:  %s"},
	ER_LDAP_AUTH_FAILED_TO_SEARCH_DN : {11804,[]string{"HY000"},"Failed to search user full dn: %s"},
	ER_LDAP_AUTH_CONNECTION_POOL_REINIT_ENTER : {11805,[]string{"HY000"},"Ldap_connection_pool::reinit"},
	ER_SYSTEMD_NOTIFY_PATH_TOO_LONG : {11806,[]string{"HY000"},"The path '%s', from the NOTIFY_SOCKET environment variable, is too long. At %u bytes it exceeds the limit of %u bytes for an AF_UNIX socket."},
	ER_SYSTEMD_NOTIFY_CONNECT_FAILED : {11807,[]string{"HY000"},"Failed to connect to systemd notification socket named %s. Error: '%s'"},
	ER_SYSTEMD_NOTIFY_WRITE_FAILED : {11808,[]string{"HY000"},"Failed to write '%s' to systemd notification. Error: '%s'"},
	ER_FOUND_MISSING_GTIDS : {11809,[]string{"HY000"},"Cannot replicate to server with server_uuid='%.36s' because the present server has purged required binary logs. The connecting server needs to replicate the missing transactions from elsewhere, or be replaced by a new server created from a more recent backup. To prevent this error in the future, consider increasing the binary log expiration period on the present server. %s."},
	ER_PID_FILE_PRIV_DIRECTORY_INSECURE : {11810,[]string{"HY000"},"Insecure configuration for --pid-file: Location '%s' in the path is accessible to all OS users. Consider choosing a different directory."},
	ER_CANT_CHECK_PID_PATH : {11811,[]string{"HY000"},"Can't start server: can't check PID filepath: %s"},
	ER_VALIDATE_PWD_STATUS_VAR_REGISTRATION_FAILED : {11812,[]string{"HY000"},"validate_password status variables registration failed."},
	ER_VALIDATE_PWD_STATUS_VAR_UNREGISTRATION_FAILED : {11813,[]string{"HY000"},"validate_password status variables unregistration failed."},
	ER_VALIDATE_PWD_DICT_FILE_OPEN_FAILED : {11814,[]string{"HY000"},"Dictionary file open failed"},
	ER_VALIDATE_PWD_COULD_BE_NULL : {11815,[]string{"HY000"},"given password string could be null"},
	ER_VALIDATE_PWD_STRING_CONV_TO_LOWERCASE_FAILED : {11816,[]string{"HY000"},"failed to convert the password string to lower case"},
	ER_VALIDATE_PWD_STRING_CONV_TO_BUFFER_FAILED : {11817,[]string{"HY000"},"failed to convert the password string into a buffer"},
	ER_VALIDATE_PWD_STRING_HANDLER_MEM_ALLOCATION_FAILED : {11818,[]string{"HY000"},"memory allocation failed for string handler"},
	ER_VALIDATE_PWD_STRONG_POLICY_DICT_FILE_UNSPECIFIED : {11819,[]string{"HY000"},"Since the validate_password_policy is mentioned as Strong, dictionary file must be specified"},
	ER_VALIDATE_PWD_CONVERT_TO_BUFFER_FAILED : {11820,[]string{"HY000"},"convert_to_buffer service failed"},
	ER_VALIDATE_PWD_VARIABLE_REGISTRATION_FAILED : {11821,[]string{"HY000"},"%s variable registration failed."},
	ER_VALIDATE_PWD_VARIABLE_UNREGISTRATION_FAILED : {11822,[]string{"HY000"},"%s variable unregistration failed."},
	ER_KEYRING_MIGRATION_EXTRA_OPTIONS : {11823,[]string{"HY000"},"Please specify options specific to keyring migration. Any additional options can be ignored. NOTE: Although some options are valid, migration tool can still report error example: plugin variables for which plugin is not loaded yet."},
	//OBSOLETE_ER_INVALID_DEFAULT_UTF8MB4_COLLATION : {3721,[]string{"HY000"},"Invalid default collation %s: utf8mb4_0900_ai_ci or utf8mb4_general_ci expected"},
	ER_IB_MSG_0 : {11825,[]string{"HY000"},"%s"},
	ER_IB_MSG_1 : {11826,[]string{"HY000"},"%s"},
	ER_IB_MSG_2 : {11827,[]string{"HY000"},"%s"},
	ER_IB_MSG_3 : {11828,[]string{"HY000"},"%s"},
	ER_IB_MSG_4 : {11829,[]string{"HY000"},"%s"},
	ER_IB_MSG_5 : {11830,[]string{"HY000"},"%s"},
	ER_IB_MSG_6 : {11831,[]string{"HY000"},"%s"},
	ER_IB_MSG_7 : {11832,[]string{"HY000"},"%s"},
	ER_IB_MSG_8 : {11833,[]string{"HY000"},"%s"},
	ER_IB_MSG_9 : {11834,[]string{"HY000"},"%s"},
	ER_IB_MSG_10 : {11835,[]string{"HY000"},"%s"},
	ER_IB_MSG_11 : {11836,[]string{"HY000"},"%s"},
	ER_IB_MSG_12 : {11837,[]string{"HY000"},"%s"},
	ER_IB_MSG_13 : {11838,[]string{"HY000"},"%s"},
	ER_IB_MSG_14 : {11839,[]string{"HY000"},"%s"},
	ER_IB_MSG_15 : {11840,[]string{"HY000"},"%s"},
	ER_IB_MSG_16 : {11841,[]string{"HY000"},"%s"},
	ER_IB_MSG_17 : {11842,[]string{"HY000"},"%s"},
	ER_IB_MSG_18 : {11843,[]string{"HY000"},"%s"},
	ER_IB_MSG_19 : {11844,[]string{"HY000"},"%s"},
	ER_IB_MSG_20 : {11845,[]string{"HY000"},"%s"},
	ER_IB_MSG_21 : {11846,[]string{"HY000"},"%s"},
	ER_IB_MSG_22 : {11847,[]string{"HY000"},"%s"},
	ER_IB_MSG_23 : {11848,[]string{"HY000"},"%s"},
	ER_IB_MSG_24 : {11849,[]string{"HY000"},"%s"},
	ER_IB_MSG_25 : {11850,[]string{"HY000"},"%s"},
	ER_IB_MSG_26 : {11851,[]string{"HY000"},"%s"},
	ER_IB_MSG_27 : {11852,[]string{"HY000"},"%s"},
	ER_IB_MSG_28 : {11853,[]string{"HY000"},"%s"},
	ER_IB_MSG_29 : {11854,[]string{"HY000"},"%s"},
	ER_IB_MSG_30 : {11855,[]string{"HY000"},"%s"},
	ER_IB_MSG_31 : {11856,[]string{"HY000"},"%s"},
	ER_IB_MSG_32 : {11857,[]string{"HY000"},"%s"},
	ER_IB_MSG_33 : {11858,[]string{"HY000"},"%s"},
	ER_IB_MSG_34 : {11859,[]string{"HY000"},"%s"},
	ER_IB_MSG_35 : {11860,[]string{"HY000"},"%s"},
	ER_IB_MSG_36 : {11861,[]string{"HY000"},"%s"},
	ER_IB_MSG_37 : {11862,[]string{"HY000"},"%s"},
	ER_IB_MSG_38 : {11863,[]string{"HY000"},"%s"},
	ER_IB_MSG_39 : {11864,[]string{"HY000"},"%s"},
	ER_IB_MSG_40 : {11865,[]string{"HY000"},"%s"},
	ER_IB_MSG_41 : {11866,[]string{"HY000"},"%s"},
	ER_IB_MSG_42 : {11867,[]string{"HY000"},"%s"},
	ER_IB_MSG_43 : {11868,[]string{"HY000"},"%s"},
	ER_IB_MSG_44 : {11869,[]string{"HY000"},"%s"},
	ER_IB_MSG_45 : {11870,[]string{"HY000"},"%s"},
	ER_IB_MSG_46 : {11871,[]string{"HY000"},"%s"},
	ER_IB_MSG_47 : {11872,[]string{"HY000"},"%s"},
	ER_IB_MSG_48 : {11873,[]string{"HY000"},"%s"},
	ER_IB_MSG_49 : {11874,[]string{"HY000"},"%s"},
	ER_IB_MSG_50 : {11875,[]string{"HY000"},"%s"},
	ER_IB_MSG_51 : {11876,[]string{"HY000"},"%s"},
	ER_IB_MSG_52 : {11877,[]string{"HY000"},"%s"},
	ER_IB_MSG_53 : {11878,[]string{"HY000"},"%s"},
	ER_IB_MSG_54 : {11879,[]string{"HY000"},"%s"},
	ER_IB_MSG_55 : {11880,[]string{"HY000"},"%s"},
	ER_IB_MSG_56 : {11881,[]string{"HY000"},"%s"},
	ER_IB_MSG_57 : {11882,[]string{"HY000"},"%s"},
	ER_IB_MSG_58 : {11883,[]string{"HY000"},"%s"},
	ER_IB_MSG_59 : {11884,[]string{"HY000"},"%s"},
	ER_IB_MSG_60 : {11885,[]string{"HY000"},"%s"},
	ER_IB_MSG_61 : {11886,[]string{"HY000"},"%s"},
	ER_IB_MSG_62 : {11887,[]string{"HY000"},"%s"},
	ER_IB_MSG_63 : {11888,[]string{"HY000"},"%s"},
	ER_IB_MSG_64 : {11889,[]string{"HY000"},"%s"},
	ER_IB_MSG_65 : {11890,[]string{"HY000"},"%s"},
	ER_IB_MSG_66 : {11891,[]string{"HY000"},"%s"},
	ER_IB_MSG_67 : {11892,[]string{"HY000"},"%s"},
	ER_IB_MSG_68 : {11893,[]string{"HY000"},"%s"},
	ER_IB_MSG_69 : {11894,[]string{"HY000"},"%s"},
	ER_IB_MSG_70 : {11895,[]string{"HY000"},"%s"},
	ER_IB_MSG_71 : {11896,[]string{"HY000"},"%s"},
	ER_IB_MSG_72 : {11897,[]string{"HY000"},"%s"},
	ER_IB_MSG_73 : {11898,[]string{"HY000"},"%s"},
	ER_IB_MSG_74 : {11899,[]string{"HY000"},"%s"},
	ER_IB_MSG_75 : {11900,[]string{"HY000"},"%s"},
	ER_IB_MSG_76 : {11901,[]string{"HY000"},"%s"},
	ER_IB_MSG_77 : {11902,[]string{"HY000"},"%s"},
	ER_IB_MSG_78 : {11903,[]string{"HY000"},"%s"},
	ER_IB_MSG_79 : {11904,[]string{"HY000"},"%s"},
	ER_IB_MSG_80 : {11905,[]string{"HY000"},"%s"},
	ER_IB_MSG_81 : {11906,[]string{"HY000"},"%s"},
	ER_IB_MSG_82 : {11907,[]string{"HY000"},"%s"},
	ER_IB_MSG_83 : {11908,[]string{"HY000"},"%s"},
	ER_IB_MSG_84 : {11909,[]string{"HY000"},"%s"},
	ER_IB_MSG_85 : {11910,[]string{"HY000"},"%s"},
	ER_IB_MSG_86 : {11911,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_87 : {11912,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_88 : {11913,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_89 : {11914,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_90 : {11915,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_91 : {11916,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_92 : {11917,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_93 : {11918,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_94 : {11919,[]string{"HY000"},"%s"},
	ER_IB_MSG_95 : {11920,[]string{"HY000"},"%s"},
	ER_IB_MSG_96 : {11921,[]string{"HY000"},"%s"},
	ER_IB_MSG_97 : {11922,[]string{"HY000"},"%s"},
	ER_IB_MSG_98 : {11923,[]string{"HY000"},"%s"},
	ER_IB_MSG_99 : {11924,[]string{"HY000"},"%s"},
	ER_IB_MSG_100 : {11925,[]string{"HY000"},"%s"},
	ER_IB_MSG_101 : {11926,[]string{"HY000"},"%s"},
	ER_IB_MSG_102 : {11927,[]string{"HY000"},"%s"},
	ER_IB_MSG_103 : {11928,[]string{"HY000"},"%s"},
	ER_IB_MSG_104 : {11929,[]string{"HY000"},"%s"},
	ER_IB_MSG_105 : {11930,[]string{"HY000"},"%s"},
	ER_IB_MSG_106 : {11931,[]string{"HY000"},"%s"},
	ER_IB_MSG_107 : {11932,[]string{"HY000"},"%s"},
	ER_IB_MSG_108 : {11933,[]string{"HY000"},"%s"},
	ER_IB_MSG_109 : {11934,[]string{"HY000"},"%s"},
	ER_IB_MSG_110 : {11935,[]string{"HY000"},"%s"},
	ER_IB_MSG_111 : {11936,[]string{"HY000"},"%s"},
	ER_IB_MSG_112 : {11937,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_113 : {11938,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_114 : {11939,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_115 : {11940,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_116 : {11941,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_117 : {11942,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_118 : {11943,[]string{"HY000"},"%s"},
	ER_IB_MSG_119 : {11944,[]string{"HY000"},"%s"},
	ER_IB_MSG_120 : {11945,[]string{"HY000"},"%s"},
	ER_IB_MSG_121 : {11946,[]string{"HY000"},"%s"},
	ER_IB_MSG_122 : {11947,[]string{"HY000"},"%s"},
	ER_IB_MSG_123 : {11948,[]string{"HY000"},"%s"},
	ER_IB_MSG_124 : {11949,[]string{"HY000"},"%s"},
	ER_IB_MSG_125 : {11950,[]string{"HY000"},"%s"},
	ER_IB_MSG_126 : {11951,[]string{"HY000"},"%s"},
	ER_IB_MSG_127 : {11952,[]string{"HY000"},"%s"},
	ER_IB_MSG_128 : {11953,[]string{"HY000"},"%s"},
	ER_IB_MSG_129 : {11954,[]string{"HY000"},"%s"},
	ER_IB_MSG_130 : {11955,[]string{"HY000"},"%s"},
	ER_IB_MSG_131 : {11956,[]string{"HY000"},"%s"},
	ER_IB_MSG_132 : {11957,[]string{"HY000"},"%s"},
	ER_IB_MSG_133 : {11958,[]string{"HY000"},"%s"},
	ER_IB_MSG_134 : {11959,[]string{"HY000"},"%s"},
	ER_IB_MSG_135 : {11960,[]string{"HY000"},"%s"},
	ER_IB_MSG_136 : {11961,[]string{"HY000"},"%s"},
	ER_IB_MSG_137 : {11962,[]string{"HY000"},"%s"},
	ER_IB_MSG_138 : {11963,[]string{"HY000"},"%s"},
	ER_IB_MSG_139 : {11964,[]string{"HY000"},"%s"},
	ER_IB_MSG_140 : {11965,[]string{"HY000"},"%s"},
	ER_IB_MSG_141 : {11966,[]string{"HY000"},"%s"},
	ER_IB_MSG_142 : {11967,[]string{"HY000"},"%s"},
	ER_IB_MSG_143 : {11968,[]string{"HY000"},"%s"},
	ER_IB_MSG_144 : {11969,[]string{"HY000"},"%s"},
	ER_IB_MSG_145 : {11970,[]string{"HY000"},"%s"},
	ER_IB_MSG_146 : {11971,[]string{"HY000"},"%s"},
	ER_IB_MSG_147 : {11972,[]string{"HY000"},"%s"},
	ER_IB_MSG_148 : {11973,[]string{"HY000"},"%s"},
	ER_IB_CLONE_INTERNAL : {11974,[]string{"HY000"},"%s"},
	ER_IB_CLONE_TIMEOUT : {11975,[]string{"HY000"},"%s"},
	ER_IB_CLONE_STATUS_FILE : {11976,[]string{"HY000"},"%s"},
	ER_IB_CLONE_SQL : {11977,[]string{"HY000"},"%s"},
	ER_IB_CLONE_VALIDATE : {11978,[]string{"HY000"},"%s"},
	ER_IB_CLONE_PUNCH_HOLE : {11979,[]string{"HY000"},"%s"},
	ER_IB_CLONE_GTID_PERSIST : {11980,[]string{"HY000"},"%s"},
	ER_IB_MSG_156 : {11981,[]string{"HY000"},"%s"},
	ER_IB_MSG_157 : {11982,[]string{"HY000"},"%s"},
	ER_IB_MSG_158 : {11983,[]string{"HY000"},"%s"},
	ER_IB_MSG_159 : {11984,[]string{"HY000"},"%s"},
	ER_IB_MSG_160 : {11985,[]string{"HY000"},"%s"},
	ER_IB_MSG_161 : {11986,[]string{"HY000"},"%s"},
	ER_IB_MSG_162 : {11987,[]string{"HY000"},"%s"},
	ER_IB_MSG_163 : {11988,[]string{"HY000"},"%s"},
	ER_IB_MSG_164 : {11989,[]string{"HY000"},"%s"},
	ER_IB_MSG_165 : {11990,[]string{"HY000"},"%s"},
	ER_IB_MSG_166 : {11991,[]string{"HY000"},"%s"},
	ER_IB_MSG_167 : {11992,[]string{"HY000"},"%s"},
	ER_IB_MSG_168 : {11993,[]string{"HY000"},"%s"},
	ER_IB_MSG_169 : {11994,[]string{"HY000"},"%s"},
	ER_IB_MSG_170 : {11995,[]string{"HY000"},"%s"},
	ER_IB_MSG_171 : {11996,[]string{"HY000"},"%s"},
	ER_IB_MSG_172 : {11997,[]string{"HY000"},"%s"},
	ER_IB_MSG_173 : {11998,[]string{"HY000"},"%s"},
	ER_IB_MSG_174 : {11999,[]string{"HY000"},"%s"},
	ER_IB_MSG_175 : {12000,[]string{"HY000"},"%s"},
	ER_IB_MSG_176 : {12001,[]string{"HY000"},"%s"},
	ER_IB_MSG_177 : {12002,[]string{"HY000"},"%s"},
	ER_IB_MSG_178 : {12003,[]string{"HY000"},"%s"},
	ER_IB_MSG_179 : {12004,[]string{"HY000"},"%s"},
	ER_IB_MSG_180 : {12005,[]string{"HY000"},"%s"},
	ER_IB_MSG_181 : {12006,[]string{"HY000"},"%s"},
	ER_IB_MSG_182 : {12007,[]string{"HY000"},"%s"},
	ER_IB_MSG_183 : {12008,[]string{"HY000"},"%s"},
	ER_IB_MSG_184 : {12009,[]string{"HY000"},"%s"},
	ER_IB_MSG_185 : {12010,[]string{"HY000"},"%s"},
	ER_IB_MSG_186 : {12011,[]string{"HY000"},"%s"},
	ER_IB_MSG_187 : {12012,[]string{"HY000"},"%s"},
	ER_IB_MSG_188 : {12013,[]string{"HY000"},"%s"},
	ER_IB_MSG_189 : {12014,[]string{"HY000"},"%s"},
	ER_IB_MSG_190 : {12015,[]string{"HY000"},"%s"},
	ER_IB_MSG_191 : {12016,[]string{"HY000"},"%s"},
	ER_IB_MSG_192 : {12017,[]string{"HY000"},"%s"},
	ER_IB_MSG_193 : {12018,[]string{"HY000"},"%s"},
	ER_IB_MSG_194 : {12019,[]string{"HY000"},"%s"},
	ER_IB_MSG_195 : {12020,[]string{"HY000"},"%s"},
	ER_IB_MSG_196 : {12021,[]string{"HY000"},"%s"},
	ER_IB_MSG_197 : {12022,[]string{"HY000"},"%s"},
	ER_IB_MSG_198 : {12023,[]string{"HY000"},"%s"},
	ER_IB_MSG_199 : {12024,[]string{"HY000"},"%s"},
	ER_IB_MSG_200 : {12025,[]string{"HY000"},"%s"},
	ER_IB_MSG_201 : {12026,[]string{"HY000"},"%s"},
	ER_IB_MSG_202 : {12027,[]string{"HY000"},"%s"},
	ER_IB_MSG_203 : {12028,[]string{"HY000"},"%s"},
	ER_IB_MSG_204 : {12029,[]string{"HY000"},"%s"},
	ER_IB_MSG_205 : {12030,[]string{"HY000"},"%s"},
	ER_IB_MSG_206 : {12031,[]string{"HY000"},"%s"},
	ER_IB_MSG_207 : {12032,[]string{"HY000"},"%s"},
	ER_IB_MSG_208 : {12033,[]string{"HY000"},"%s"},
	ER_IB_MSG_209 : {12034,[]string{"HY000"},"%s"},
	ER_IB_MSG_210 : {12035,[]string{"HY000"},"%s"},
	ER_IB_MSG_211 : {12036,[]string{"HY000"},"%s"},
	ER_IB_MSG_212 : {12037,[]string{"HY000"},"%s"},
	ER_IB_MSG_213 : {12038,[]string{"HY000"},"%s"},
	ER_IB_MSG_214 : {12039,[]string{"HY000"},"%s"},
	ER_IB_MSG_215 : {12040,[]string{"HY000"},"%s"},
	ER_IB_MSG_216 : {12041,[]string{"HY000"},"%s"},
	ER_IB_MSG_217 : {12042,[]string{"HY000"},"%s"},
	ER_IB_MSG_218 : {12043,[]string{"HY000"},"%s"},
	ER_IB_MSG_219 : {12044,[]string{"HY000"},"%s"},
	ER_IB_MSG_220 : {12045,[]string{"HY000"},"%s"},
	ER_IB_MSG_221 : {12046,[]string{"HY000"},"%s"},
	ER_IB_MSG_222 : {12047,[]string{"HY000"},"%s"},
	ER_IB_MSG_223 : {12048,[]string{"HY000"},"%s"},
	ER_IB_MSG_224 : {12049,[]string{"HY000"},"%s"},
	ER_IB_MSG_225 : {12050,[]string{"HY000"},"%s"},
	ER_IB_MSG_226 : {12051,[]string{"HY000"},"%s"},
	ER_IB_MSG_227 : {12052,[]string{"HY000"},"%s"},
	ER_IB_MSG_228 : {12053,[]string{"HY000"},"%s"},
	ER_IB_MSG_229 : {12054,[]string{"HY000"},"%s"},
	ER_IB_MSG_230 : {12055,[]string{"HY000"},"%s"},
	ER_IB_MSG_231 : {12056,[]string{"HY000"},"%s"},
	ER_IB_MSG_232 : {12057,[]string{"HY000"},"%s"},
	ER_IB_MSG_233 : {12058,[]string{"HY000"},"%s"},
	ER_IB_MSG_234 : {12059,[]string{"HY000"},"%s"},
	ER_IB_MSG_235 : {12060,[]string{"HY000"},"%s"},
	ER_IB_MSG_236 : {12061,[]string{"HY000"},"%s"},
	ER_IB_MSG_237 : {12062,[]string{"HY000"},"%s"},
	ER_IB_MSG_238 : {12063,[]string{"HY000"},"%s"},
	ER_IB_MSG_239 : {12064,[]string{"HY000"},"%s"},
	ER_IB_MSG_240 : {12065,[]string{"HY000"},"%s"},
	ER_IB_MSG_241 : {12066,[]string{"HY000"},"%s"},
	ER_IB_MSG_242 : {12067,[]string{"HY000"},"%s"},
	ER_IB_MSG_243 : {12068,[]string{"HY000"},"%s"},
	ER_IB_MSG_244 : {12069,[]string{"HY000"},"%s"},
	ER_IB_MSG_245 : {12070,[]string{"HY000"},"%s"},
	ER_IB_MSG_246 : {12071,[]string{"HY000"},"%s"},
	ER_IB_MSG_247 : {12072,[]string{"HY000"},"%s"},
	ER_IB_MSG_248 : {12073,[]string{"HY000"},"%s"},
	ER_IB_MSG_249 : {12074,[]string{"HY000"},"%s"},
	ER_IB_MSG_250 : {12075,[]string{"HY000"},"%s"},
	ER_IB_MSG_251 : {12076,[]string{"HY000"},"%s"},
	ER_IB_MSG_252 : {12077,[]string{"HY000"},"%s"},
	ER_IB_MSG_253 : {12078,[]string{"HY000"},"%s"},
	ER_IB_MSG_254 : {12079,[]string{"HY000"},"%s"},
	ER_IB_MSG_255 : {12080,[]string{"HY000"},"%s"},
	ER_IB_MSG_256 : {12081,[]string{"HY000"},"%s"},
	ER_IB_MSG_257 : {12082,[]string{"HY000"},"%s"},
	ER_IB_MSG_258 : {12083,[]string{"HY000"},"%s"},
	ER_IB_MSG_259 : {12084,[]string{"HY000"},"%s"},
	ER_IB_MSG_260 : {12085,[]string{"HY000"},"%s"},
	ER_IB_MSG_261 : {12086,[]string{"HY000"},"%s"},
	ER_IB_MSG_262 : {12087,[]string{"HY000"},"%s"},
	ER_IB_MSG_263 : {12088,[]string{"HY000"},"%s"},
	ER_IB_MSG_264 : {12089,[]string{"HY000"},"%s"},
	ER_IB_MSG_265 : {12090,[]string{"HY000"},"%s"},
	ER_IB_MSG_266 : {12091,[]string{"HY000"},"%s"},
	ER_IB_MSG_267 : {12092,[]string{"HY000"},"%s"},
	ER_IB_MSG_268 : {12093,[]string{"HY000"},"%s"},
	ER_IB_MSG_269 : {12094,[]string{"HY000"},"%s"},
	ER_IB_MSG_270 : {12095,[]string{"HY000"},"%s"},
	ER_IB_MSG_271 : {12096,[]string{"HY000"},"%s"},
	ER_IB_MSG_272 : {12097,[]string{"HY000"},"Table flags are 0x%lx in the data dictionary but the flags in file %s  are 0x%llx!"},
	ER_IB_MSG_273 : {12098,[]string{"HY000"},"Can't read encryption key from file %s!"},
	ER_IB_MSG_274 : {12099,[]string{"HY000"},"Cannot close file %s, because n_pending_flushes %zu"},
	ER_IB_MSG_275 : {12100,[]string{"HY000"},"Cannot close file %s, because modification count %lld != flush count %lld"},
	ER_IB_MSG_276 : {12101,[]string{"HY000"},"Cannot close file %s, because it is in use"},
	ER_IB_MSG_277 : {12102,[]string{"HY000"},"Open file list len in shard %zu is %llu"},
	ER_IB_MSG_278 : {12103,[]string{"HY000"},"Tablespace %s, waiting for IO to stop for %lld seconds"},
	ER_IB_MSG_279 : {12104,[]string{"HY000"},"%s"},
	ER_IB_MSG_280 : {12105,[]string{"HY000"},"%s"},
	ER_IB_MSG_281 : {12106,[]string{"HY000"},"%s"},
	ER_IB_MSG_282 : {12107,[]string{"HY000"},"%s"},
	ER_IB_MSG_283 : {12108,[]string{"HY000"},"%s"},
	ER_IB_MSG_284 : {12109,[]string{"HY000"},"You must raise the value of innodb_open_files in my.cnf! Remember that InnoDB keeps all log files and all system tablespace files open for the whole time mysqld is running, and needs to open also some .ibd files if the file-per-table storage model is used. Current open files %zu, max allowed open files %zu."},
	ER_IB_MSG_285 : {12110,[]string{"HY000"},"Max tablespace id is too high, %lu"},
	ER_IB_MSG_286 : {12111,[]string{"HY000"},"Trying to access missing tablespace %lu"},
	ER_IB_MSG_287 : {12112,[]string{"HY000"},"Trying to close/delete tablespace '%s' but there are %lu pending operations on it."},
	ER_IB_MSG_288 : {12113,[]string{"HY000"},"Trying to delete/close tablespace '%s' but there are %lu flushes and %zu pending I/O's on it."},
	ER_IB_MSG_289 : {12114,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_290 : {12115,[]string{"HY000"},"Cannot delete tablespace %lu because it is not found in the tablespace memory cache."},
	ER_IB_MSG_291 : {12116,[]string{"HY000"},"While deleting tablespace %lu in DISCARD TABLESPACE. File rename/delete failed: %s"},
	ER_IB_MSG_292 : {12117,[]string{"HY000"},"Cannot delete tablespace %lu in DISCARD TABLESPACE: %s"},
	ER_IB_MSG_293 : {12118,[]string{"HY000"},"Cannot rename '%s' to '%s' for space ID %lu because the source file does not exist."},
	ER_IB_MSG_294 : {12119,[]string{"HY000"},"Cannot rename '%s' to '%s' for space ID %lu because the target file exists. Remove the target file and try again."},
	ER_IB_MSG_295 : {12120,[]string{"HY000"},"Cannot rename file '%s' (space id %lu) retried %llu times. There are either pending IOs or flushes or the file is being extended."},
	ER_IB_MSG_296 : {12121,[]string{"HY000"},"Cannot find space id %lu in the tablespace memory cache, though the file '%s' in a rename operation should have that ID."},
	ER_IB_MSG_297 : {12122,[]string{"HY000"},"Rename waiting for IO to resume"},
	ER_IB_MSG_298 : {12123,[]string{"HY000"},"Cannot find tablespace for '%s' in the tablespace memory cache"},
	ER_IB_MSG_299 : {12124,[]string{"HY000"},"Cannot find tablespace for '%s' in the tablespace memory cache"},
	ER_IB_MSG_300 : {12125,[]string{"HY000"},"Tablespace '%s' is already in the tablespace memory cache"},
	ER_IB_MSG_301 : {12126,[]string{"HY000"},"Cannot create file '%s'"},
	ER_IB_MSG_UNEXPECTED_FILE_EXISTS : {12127,[]string{"HY000"},"The file '%s' already exists though the corresponding table did not exist. Have you moved InnoDB .ibd files around without using the SQL commands DISCARD TABLESPACE and IMPORT TABLESPACE, or did mysqld crash in the middle of CREATE TABLE? You can resolve the problem by removing the file '%s' under the 'datadir' of MySQL."},
	ER_IB_MSG_303 : {12128,[]string{"HY000"},"posix_fallocate(): Failed to preallocate data for file %s, desired size %llu Operating system error number %d - %s. Check that the disk is not full or a disk quota exceeded. Make sure the file system supports this function. Refer to your operating system documentation for operating system error code information."},
	ER_IB_MSG_304 : {12129,[]string{"HY000"},"Could not write the first page to tablespace '%s'"},
	ER_IB_MSG_305 : {12130,[]string{"HY000"},"File flush of tablespace '%s' failed"},
	ER_IB_MSG_306 : {12131,[]string{"HY000"},"Could not find a valid tablespace file for `%s`. %s"},
	ER_IB_MSG_307 : {12132,[]string{"HY000"},"Ignoring data file '%s' with space ID %lu. Another data file called '%s' exists with the same space ID"},
	ER_IB_MSG_308 : {12133,[]string{"HY000"},"%s"},
	ER_IB_MSG_309 : {12134,[]string{"HY000"},"%s"},
	ER_IB_MSG_310 : {12135,[]string{"HY000"},"%s"},
	ER_IB_MSG_311 : {12136,[]string{"HY000"},"%s"},
	ER_IB_MSG_312 : {12137,[]string{"HY000"},"Can't set encryption information for tablespace %s!"},
	ER_IB_MSG_313 : {12138,[]string{"HY000"},"%s"},
	ER_IB_MSG_314 : {12139,[]string{"HY000"},"%s"},
	ER_IB_MSG_315 : {12140,[]string{"HY000"},"%s"},
	ER_IB_MSG_316 : {12141,[]string{"HY000"},"%s"},
	ER_IB_MSG_317 : {12142,[]string{"HY000"},"%s"},
	ER_IB_MSG_318 : {12143,[]string{"HY000"},"%s"},
	ER_IB_MSG_319 : {12144,[]string{"HY000"},"%s"},
	ER_IB_MSG_320 : {12145,[]string{"HY000"},"%s"},
	ER_IB_MSG_321 : {12146,[]string{"HY000"},"%s"},
	ER_IB_MSG_322 : {12147,[]string{"HY000"},"%s"},
	ER_IB_MSG_323 : {12148,[]string{"HY000"},"%s"},
	ER_IB_MSG_324 : {12149,[]string{"HY000"},"%s"},
	ER_IB_MSG_325 : {12150,[]string{"HY000"},"%s"},
	ER_IB_MSG_326 : {12151,[]string{"HY000"},"%s"},
	ER_IB_MSG_327 : {12152,[]string{"HY000"},"%s"},
	ER_IB_MSG_328 : {12153,[]string{"HY000"},"%s"},
	ER_IB_MSG_329 : {12154,[]string{"HY000"},"%s"},
	ER_IB_MSG_330 : {12155,[]string{"HY000"},"%s"},
	ER_IB_MSG_331 : {12156,[]string{"HY000"},"%s"},
	ER_IB_MSG_332 : {12157,[]string{"HY000"},"%s"},
	ER_IB_MSG_333 : {12158,[]string{"HY000"},"%s"},
	ER_IB_MSG_334 : {12159,[]string{"HY000"},"%s"},
	ER_IB_MSG_335 : {12160,[]string{"HY000"},"%s"},
	ER_IB_MSG_336 : {12161,[]string{"HY000"},"%s"},
	ER_IB_MSG_337 : {12162,[]string{"HY000"},"%s"},
	ER_IB_MSG_338 : {12163,[]string{"HY000"},"%s"},
	ER_IB_MSG_339 : {12164,[]string{"HY000"},"%s"},
	ER_IB_MSG_340 : {12165,[]string{"HY000"},"%s"},
	ER_IB_MSG_341 : {12166,[]string{"HY000"},"%s"},
	ER_IB_MSG_342 : {12167,[]string{"HY000"},"%s"},
	ER_IB_MSG_343 : {12168,[]string{"HY000"},"%s"},
	ER_IB_MSG_344 : {12169,[]string{"HY000"},"%s"},
	ER_IB_MSG_345 : {12170,[]string{"HY000"},"%s"},
	ER_IB_MSG_346 : {12171,[]string{"HY000"},"%s"},
	ER_IB_MSG_347 : {12172,[]string{"HY000"},"%s"},
	ER_IB_MSG_348 : {12173,[]string{"HY000"},"%s"},
	ER_IB_MSG_349 : {12174,[]string{"HY000"},"%s"},
	ER_IB_MSG_350 : {12175,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_351 : {12176,[]string{"HY000"},"%s"},
	ER_IB_MSG_UNPROTECTED_LOCATION_ALLOWED : {12177,[]string{"HY000"},"The datafile '%s' for tablespace %s is in an unprotected location. This file cannot be recovered after a crash until this location is added to innodb_directories."},
	//OBSOLETE_ER_IB_MSG_353 : {12178,[]string{"HY000"},"%s"},
	ER_IB_MSG_354 : {12179,[]string{"HY000"},"%s"},
	ER_IB_MSG_355 : {12180,[]string{"HY000"},"%s"},
	ER_IB_MSG_356 : {12181,[]string{"HY000"},"%s"},
	ER_IB_MSG_357 : {12182,[]string{"HY000"},"%s"},
	ER_IB_MSG_358 : {12183,[]string{"HY000"},"%s"},
	ER_IB_MSG_359 : {12184,[]string{"HY000"},"%s"},
	ER_IB_MSG_360 : {12185,[]string{"HY000"},"%s"},
	ER_IB_MSG_361 : {12186,[]string{"HY000"},"%s"},
	ER_IB_MSG_362 : {12187,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_363 : {12188,[]string{"HY000"},"%s"},
	ER_IB_MSG_364 : {12189,[]string{"HY000"},"%s"},
	ER_IB_MSG_365 : {12190,[]string{"HY000"},"%s"},
	ER_IB_MSG_IGNORE_SCAN_PATH : {12191,[]string{"HY000"},"Scan path '%s' is ignored because %s"},
	ER_IB_MSG_367 : {12192,[]string{"HY000"},"%s"},
	ER_IB_MSG_368 : {12193,[]string{"HY000"},"%s"},
	ER_IB_MSG_369 : {12194,[]string{"HY000"},"%s"},
	ER_IB_MSG_370 : {12195,[]string{"HY000"},"%s"},
	ER_IB_MSG_371 : {12196,[]string{"HY000"},"%s"},
	ER_IB_MSG_372 : {12197,[]string{"HY000"},"%s"},
	ER_IB_MSG_373 : {12198,[]string{"HY000"},"%s"},
	ER_IB_MSG_374 : {12199,[]string{"HY000"},"%s"},
	ER_IB_MSG_375 : {12200,[]string{"HY000"},"%s"},
	ER_IB_MSG_376 : {12201,[]string{"HY000"},"%s"},
	ER_IB_MSG_377 : {12202,[]string{"HY000"},"%s"},
	ER_IB_MSG_378 : {12203,[]string{"HY000"},"%s"},
	ER_IB_MSG_379 : {12204,[]string{"HY000"},"%s"},
	ER_IB_MSG_380 : {12205,[]string{"HY000"},"%s"},
	ER_IB_MSG_381 : {12206,[]string{"HY000"},"%s"},
	ER_IB_MSG_382 : {12207,[]string{"HY000"},"%s"},
	ER_IB_MSG_383 : {12208,[]string{"HY000"},"%s"},
	ER_IB_MSG_384 : {12209,[]string{"HY000"},"%s"},
	ER_IB_MSG_385 : {12210,[]string{"HY000"},"%s"},
	ER_IB_MSG_386 : {12211,[]string{"HY000"},"%s"},
	ER_IB_MSG_387 : {12212,[]string{"HY000"},"%s"},
	ER_IB_MSG_GENERAL_TABLESPACE_UNDER_DATADIR : {12213,[]string{"HY000"},"A general tablespace cannot be located under the datadir. Cannot open file '%s'."},
	ER_IB_MSG_IMPLICIT_TABLESPACE_IN_DATADIR : {12214,[]string{"HY000"},"A file-per-table tablespace cannot be located in the datadir. Cannot open file '%s'."},
	ER_IB_MSG_390 : {12215,[]string{"HY000"},"%s"},
	ER_IB_MSG_391 : {12216,[]string{"HY000"},"%s"},
	ER_IB_MSG_392 : {12217,[]string{"HY000"},"%s"},
	ER_IB_MSG_393 : {12218,[]string{"HY000"},"%s"},
	ER_IB_MSG_394 : {12219,[]string{"HY000"},"%s"},
	ER_IB_MSG_395 : {12220,[]string{"HY000"},"%s"},
	ER_IB_MSG_396 : {12221,[]string{"HY000"},"%s"},
	ER_IB_MSG_397 : {12222,[]string{"HY000"},"%s"},
	ER_IB_MSG_398 : {12223,[]string{"HY000"},"%s"},
	ER_IB_MSG_399 : {12224,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_400 : {12225,[]string{"HY000"},"%s"},
	ER_IB_MSG_401 : {12226,[]string{"HY000"},"%s"},
	ER_IB_MSG_402 : {12227,[]string{"HY000"},"%s"},
	ER_IB_MSG_403 : {12228,[]string{"HY000"},"%s"},
	ER_IB_MSG_404 : {12229,[]string{"HY000"},"%s"},
	ER_IB_MSG_405 : {12230,[]string{"HY000"},"%s"},
	ER_IB_MSG_406 : {12231,[]string{"HY000"},"%s"},
	ER_IB_MSG_407 : {12232,[]string{"HY000"},"%s"},
	ER_IB_MSG_408 : {12233,[]string{"HY000"},"%s"},
	ER_IB_MSG_409 : {12234,[]string{"HY000"},"%s"},
	ER_IB_MSG_410 : {12235,[]string{"HY000"},"%s"},
	ER_IB_MSG_411 : {12236,[]string{"HY000"},"%s"},
	ER_IB_MSG_412 : {12237,[]string{"HY000"},"%s"},
	ER_IB_MSG_413 : {12238,[]string{"HY000"},"%s"},
	ER_IB_MSG_414 : {12239,[]string{"HY000"},"%s"},
	ER_IB_MSG_415 : {12240,[]string{"HY000"},"%s"},
	ER_IB_MSG_416 : {12241,[]string{"HY000"},"%s"},
	ER_IB_MSG_417 : {12242,[]string{"HY000"},"%s"},
	ER_IB_MSG_418 : {12243,[]string{"HY000"},"%s"},
	ER_IB_MSG_419 : {12244,[]string{"HY000"},"%s"},
	ER_IB_MSG_420 : {12245,[]string{"HY000"},"%s"},
	ER_IB_MSG_421 : {12246,[]string{"HY000"},"%s"},
	ER_IB_MSG_422 : {12247,[]string{"HY000"},"%s"},
	ER_IB_MSG_423 : {12248,[]string{"HY000"},"%s"},
	ER_IB_MSG_424 : {12249,[]string{"HY000"},"%s"},
	ER_IB_MSG_425 : {12250,[]string{"HY000"},"%s"},
	ER_IB_MSG_426 : {12251,[]string{"HY000"},"%s"},
	ER_IB_MSG_427 : {12252,[]string{"HY000"},"%s"},
	ER_IB_MSG_428 : {12253,[]string{"HY000"},"%s"},
	ER_IB_MSG_429 : {12254,[]string{"HY000"},"%s"},
	ER_IB_MSG_430 : {12255,[]string{"HY000"},"%s"},
	ER_IB_MSG_431 : {12256,[]string{"HY000"},"%s"},
	ER_IB_MSG_432 : {12257,[]string{"HY000"},"%s"},
	ER_IB_MSG_433 : {12258,[]string{"HY000"},"%s"},
	ER_IB_MSG_434 : {12259,[]string{"HY000"},"%s"},
	ER_IB_MSG_435 : {12260,[]string{"HY000"},"%s"},
	ER_IB_MSG_436 : {12261,[]string{"HY000"},"%s"},
	ER_IB_MSG_437 : {12262,[]string{"HY000"},"%s"},
	ER_IB_MSG_438 : {12263,[]string{"HY000"},"%s"},
	ER_IB_MSG_439 : {12264,[]string{"HY000"},"%s"},
	ER_IB_MSG_440 : {12265,[]string{"HY000"},"%s"},
	ER_IB_MSG_441 : {12266,[]string{"HY000"},"%s"},
	ER_IB_MSG_442 : {12267,[]string{"HY000"},"%s"},
	ER_IB_MSG_443 : {12268,[]string{"HY000"},"%s"},
	ER_IB_MSG_444 : {12269,[]string{"HY000"},"%s"},
	ER_IB_MSG_445 : {12270,[]string{"HY000"},"%s"},
	ER_IB_MSG_446 : {12271,[]string{"HY000"},"%s"},
	ER_IB_MSG_447 : {12272,[]string{"HY000"},"%s"},
	ER_IB_MSG_448 : {12273,[]string{"HY000"},"%s"},
	ER_IB_MSG_449 : {12274,[]string{"HY000"},"%s"},
	ER_IB_MSG_450 : {12275,[]string{"HY000"},"%s"},
	ER_IB_MSG_451 : {12276,[]string{"HY000"},"%s"},
	ER_IB_MSG_452 : {12277,[]string{"HY000"},"%s"},
	ER_IB_MSG_453 : {12278,[]string{"HY000"},"%s"},
	ER_IB_MSG_454 : {12279,[]string{"HY000"},"%s"},
	ER_IB_MSG_455 : {12280,[]string{"HY000"},"%s"},
	ER_IB_MSG_456 : {12281,[]string{"HY000"},"%s"},
	ER_IB_MSG_457 : {12282,[]string{"HY000"},"%s"},
	ER_IB_MSG_458 : {12283,[]string{"HY000"},"%s"},
	ER_IB_MSG_459 : {12284,[]string{"HY000"},"%s"},
	ER_IB_MSG_460 : {12285,[]string{"HY000"},"%s"},
	ER_IB_MSG_461 : {12286,[]string{"HY000"},"%s"},
	ER_IB_MSG_462 : {12287,[]string{"HY000"},"%s"},
	ER_IB_MSG_463 : {12288,[]string{"HY000"},"%s"},
	ER_IB_MSG_464 : {12289,[]string{"HY000"},"%s"},
	ER_IB_MSG_465 : {12290,[]string{"HY000"},"%s"},
	ER_IB_MSG_466 : {12291,[]string{"HY000"},"%s"},
	ER_IB_MSG_467 : {12292,[]string{"HY000"},"%s"},
	ER_IB_MSG_468 : {12293,[]string{"HY000"},"%s"},
	ER_IB_MSG_469 : {12294,[]string{"HY000"},"%s"},
	ER_IB_MSG_470 : {12295,[]string{"HY000"},"%s"},
	ER_IB_MSG_471 : {12296,[]string{"HY000"},"%s"},
	ER_IB_MSG_472 : {12297,[]string{"HY000"},"%s"},
	ER_IB_MSG_473 : {12298,[]string{"HY000"},"%s"},
	ER_IB_MSG_474 : {12299,[]string{"HY000"},"%s"},
	ER_IB_MSG_475 : {12300,[]string{"HY000"},"%s"},
	ER_IB_MSG_476 : {12301,[]string{"HY000"},"%s"},
	ER_IB_MSG_477 : {12302,[]string{"HY000"},"%s"},
	ER_IB_MSG_478 : {12303,[]string{"HY000"},"%s"},
	ER_IB_MSG_479 : {12304,[]string{"HY000"},"%s"},
	ER_IB_MSG_480 : {12305,[]string{"HY000"},"%s"},
	ER_IB_MSG_481 : {12306,[]string{"HY000"},"%s"},
	ER_IB_MSG_482 : {12307,[]string{"HY000"},"%s"},
	ER_IB_MSG_483 : {12308,[]string{"HY000"},"%s"},
	ER_IB_MSG_484 : {12309,[]string{"HY000"},"%s"},
	ER_IB_MSG_485 : {12310,[]string{"HY000"},"%s"},
	ER_IB_MSG_486 : {12311,[]string{"HY000"},"%s"},
	ER_IB_MSG_487 : {12312,[]string{"HY000"},"%s"},
	ER_IB_MSG_488 : {12313,[]string{"HY000"},"%s"},
	ER_IB_MSG_489 : {12314,[]string{"HY000"},"%s"},
	ER_IB_MSG_490 : {12315,[]string{"HY000"},"%s"},
	ER_IB_MSG_491 : {12316,[]string{"HY000"},"%s"},
	ER_IB_MSG_492 : {12317,[]string{"HY000"},"%s"},
	ER_IB_MSG_493 : {12318,[]string{"HY000"},"%s"},
	ER_IB_MSG_494 : {12319,[]string{"HY000"},"%s"},
	ER_IB_MSG_495 : {12320,[]string{"HY000"},"%s"},
	ER_IB_MSG_496 : {12321,[]string{"HY000"},"%s"},
	ER_IB_MSG_497 : {12322,[]string{"HY000"},"%s"},
	ER_IB_MSG_498 : {12323,[]string{"HY000"},"%s"},
	ER_IB_MSG_499 : {12324,[]string{"HY000"},"%s"},
	ER_IB_MSG_500 : {12325,[]string{"HY000"},"%s"},
	ER_IB_MSG_501 : {12326,[]string{"HY000"},"%s"},
	ER_IB_MSG_502 : {12327,[]string{"HY000"},"%s"},
	ER_IB_MSG_503 : {12328,[]string{"HY000"},"%s"},
	ER_IB_MSG_504 : {12329,[]string{"HY000"},"%s"},
	ER_IB_MSG_505 : {12330,[]string{"HY000"},"%s"},
	ER_IB_MSG_506 : {12331,[]string{"HY000"},"%s"},
	ER_IB_MSG_507 : {12332,[]string{"HY000"},"%s"},
	ER_IB_MSG_508 : {12333,[]string{"HY000"},"%s"},
	ER_IB_MSG_509 : {12334,[]string{"HY000"},"%s"},
	ER_IB_MSG_510 : {12335,[]string{"HY000"},"%s"},
	ER_IB_MSG_511 : {12336,[]string{"HY000"},"%s"},
	ER_IB_MSG_512 : {12337,[]string{"HY000"},"%s"},
	ER_IB_MSG_513 : {12338,[]string{"HY000"},"%s"},
	ER_IB_MSG_514 : {12339,[]string{"HY000"},"%s"},
	ER_IB_MSG_515 : {12340,[]string{"HY000"},"%s"},
	ER_IB_MSG_516 : {12341,[]string{"HY000"},"%s"},
	ER_IB_MSG_517 : {12342,[]string{"HY000"},"%s"},
	ER_IB_MSG_518 : {12343,[]string{"HY000"},"%s"},
	ER_IB_MSG_519 : {12344,[]string{"HY000"},"%s"},
	ER_IB_MSG_520 : {12345,[]string{"HY000"},"%s"},
	ER_IB_MSG_521 : {12346,[]string{"HY000"},"%s"},
	ER_IB_MSG_522 : {12347,[]string{"HY000"},"%s"},
	ER_IB_MSG_523 : {12348,[]string{"HY000"},"%s"},
	ER_IB_MSG_524 : {12349,[]string{"HY000"},"%s"},
	ER_IB_MSG_525 : {12350,[]string{"HY000"},"%s"},
	ER_IB_MSG_526 : {12351,[]string{"HY000"},"%s"},
	ER_IB_MSG_527 : {12352,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_528 : {12353,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_529 : {12354,[]string{"HY000"},"%s"},
	ER_IB_MSG_530 : {12355,[]string{"HY000"},"%s"},
	ER_IB_MSG_531 : {12356,[]string{"HY000"},"%s"},
	ER_IB_MSG_532 : {12357,[]string{"HY000"},"%s"},
	ER_IB_MSG_533 : {12358,[]string{"HY000"},"%s"},
	ER_IB_MSG_534 : {12359,[]string{"HY000"},"%s"},
	ER_IB_MSG_535 : {12360,[]string{"HY000"},"%s"},
	ER_IB_MSG_536 : {12361,[]string{"HY000"},"%s"},
	ER_IB_MSG_537 : {12362,[]string{"HY000"},"%s"},
	ER_IB_MSG_538 : {12363,[]string{"HY000"},"%s"},
	ER_IB_MSG_539 : {12364,[]string{"HY000"},"%s"},
	ER_IB_MSG_540 : {12365,[]string{"HY000"},"%s"},
	ER_IB_MSG_541 : {12366,[]string{"HY000"},"%s"},
	ER_IB_MSG_542 : {12367,[]string{"HY000"},"%s"},
	ER_IB_MSG_543 : {12368,[]string{"HY000"},"%s"},
	ER_IB_MSG_544 : {12369,[]string{"HY000"},"%s"},
	ER_IB_MSG_545 : {12370,[]string{"HY000"},"%s"},
	ER_IB_MSG_546 : {12371,[]string{"HY000"},"%s"},
	ER_IB_MSG_547 : {12372,[]string{"HY000"},"%s"},
	ER_IB_MSG_548 : {12373,[]string{"HY000"},"%s"},
	ER_IB_MSG_549 : {12374,[]string{"HY000"},"%s"},
	ER_IB_MSG_550 : {12375,[]string{"HY000"},"%s"},
	ER_IB_MSG_551 : {12376,[]string{"HY000"},"%s"},
	ER_IB_MSG_552 : {12377,[]string{"HY000"},"%s"},
	ER_IB_MSG_553 : {12378,[]string{"HY000"},"%s"},
	ER_IB_MSG_554 : {12379,[]string{"HY000"},"%s"},
	ER_IB_MSG_555 : {12380,[]string{"HY000"},"%s"},
	ER_IB_MSG_556 : {12381,[]string{"HY000"},"%s"},
	ER_IB_MSG_557 : {12382,[]string{"HY000"},"%s"},
	ER_IB_MSG_558 : {12383,[]string{"HY000"},"%s"},
	ER_IB_MSG_559 : {12384,[]string{"HY000"},"%s"},
	ER_IB_MSG_560 : {12385,[]string{"HY000"},"%s"},
	ER_IB_MSG_561 : {12386,[]string{"HY000"},"%s"},
	ER_IB_MSG_562 : {12387,[]string{"HY000"},"%s"},
	ER_IB_MSG_563 : {12388,[]string{"HY000"},"%s"},
	ER_IB_MSG_564 : {12389,[]string{"HY000"},"%s"},
	ER_IB_MSG_INVALID_LOCATION_FOR_TABLE : {12390,[]string{"HY000"},"Cannot create a tablespace for table %s because the directory is not a valid location. %s"},
	ER_IB_MSG_566 : {12391,[]string{"HY000"},"%s"},
	ER_IB_MSG_567 : {12392,[]string{"HY000"},"%s"},
	ER_IB_MSG_568 : {12393,[]string{"HY000"},"%s"},
	ER_IB_MSG_569 : {12394,[]string{"HY000"},"%s"},
	ER_IB_MSG_570 : {12395,[]string{"HY000"},"%s"},
	ER_IB_MSG_571 : {12396,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_572 : {12397,[]string{"HY000"},"%s"},
	ER_IB_MSG_573 : {12398,[]string{"HY000"},"%s"},
	ER_IB_MSG_574 : {12399,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_575 : {12400,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_576 : {12401,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_577 : {12402,[]string{"HY000"},"%s"},
	ER_IB_MSG_578 : {12403,[]string{"HY000"},"%s"},
	ER_IB_MSG_579 : {12404,[]string{"HY000"},"%s"},
	ER_IB_MSG_580 : {12405,[]string{"HY000"},"%s"},
	ER_IB_MSG_581 : {12406,[]string{"HY000"},"%s"},
	ER_IB_MSG_582 : {12407,[]string{"HY000"},"%s"},
	ER_IB_MSG_583 : {12408,[]string{"HY000"},"%s"},
	ER_IB_MSG_584 : {12409,[]string{"HY000"},"%s"},
	ER_IB_MSG_585 : {12410,[]string{"HY000"},"%s"},
	ER_IB_MSG_586 : {12411,[]string{"HY000"},"%s"},
	ER_IB_MSG_587 : {12412,[]string{"HY000"},"%s"},
	ER_IB_MSG_588 : {12413,[]string{"HY000"},"%s"},
	ER_IB_MSG_589 : {12414,[]string{"HY000"},"%s"},
	ER_IB_MSG_590 : {12415,[]string{"HY000"},"%s"},
	ER_IB_MSG_591 : {12416,[]string{"HY000"},"%s"},
	ER_IB_MSG_592 : {12417,[]string{"HY000"},"%s"},
	ER_IB_MSG_593 : {12418,[]string{"HY000"},"%s"},
	ER_IB_MSG_594 : {12419,[]string{"HY000"},"%s"},
	ER_IB_MSG_595 : {12420,[]string{"HY000"},"%s"},
	ER_IB_MSG_596 : {12421,[]string{"HY000"},"%s"},
	ER_IB_MSG_597 : {12422,[]string{"HY000"},"%s"},
	ER_IB_MSG_598 : {12423,[]string{"HY000"},"%s"},
	ER_IB_MSG_599 : {12424,[]string{"HY000"},"%s"},
	ER_IB_MSG_600 : {12425,[]string{"HY000"},"%s"},
	ER_IB_MSG_601 : {12426,[]string{"HY000"},"%s"},
	ER_IB_MSG_602 : {12427,[]string{"HY000"},"%s"},
	ER_IB_MSG_603 : {12428,[]string{"HY000"},"%s"},
	ER_IB_MSG_604 : {12429,[]string{"HY000"},"%s"},
	ER_IB_MSG_605 : {12430,[]string{"HY000"},"%s"},
	ER_IB_MSG_606 : {12431,[]string{"HY000"},"%s"},
	ER_IB_MSG_607 : {12432,[]string{"HY000"},"%s"},
	ER_IB_MSG_608 : {12433,[]string{"HY000"},"%s"},
	ER_IB_MSG_609 : {12434,[]string{"HY000"},"%s"},
	ER_IB_MSG_610 : {12435,[]string{"HY000"},"%s"},
	ER_IB_MSG_611 : {12436,[]string{"HY000"},"%s"},
	ER_IB_MSG_612 : {12437,[]string{"HY000"},"%s"},
	ER_IB_MSG_613 : {12438,[]string{"HY000"},"%s"},
	ER_IB_MSG_614 : {12439,[]string{"HY000"},"%s"},
	ER_IB_MSG_615 : {12440,[]string{"HY000"},"%s"},
	ER_IB_MSG_616 : {12441,[]string{"HY000"},"%s"},
	ER_IB_MSG_617 : {12442,[]string{"HY000"},"%s"},
	ER_IB_MSG_618 : {12443,[]string{"HY000"},"%s"},
	ER_IB_MSG_619 : {12444,[]string{"HY000"},"%s"},
	ER_IB_MSG_620 : {12445,[]string{"HY000"},"%s"},
	ER_IB_MSG_621 : {12446,[]string{"HY000"},"%s"},
	ER_IB_MSG_622 : {12447,[]string{"HY000"},"%s"},
	ER_IB_MSG_623 : {12448,[]string{"HY000"},"%s"},
	ER_IB_MSG_624 : {12449,[]string{"HY000"},"%s"},
	ER_IB_MSG_625 : {12450,[]string{"HY000"},"%s"},
	ER_IB_MSG_626 : {12451,[]string{"HY000"},"%s"},
	ER_IB_MSG_627 : {12452,[]string{"HY000"},"%s"},
	ER_IB_MSG_628 : {12453,[]string{"HY000"},"%s"},
	ER_IB_MSG_629 : {12454,[]string{"HY000"},"%s"},
	ER_IB_MSG_630 : {12455,[]string{"HY000"},"%s"},
	ER_IB_MSG_631 : {12456,[]string{"HY000"},"%s"},
	ER_IB_MSG_632 : {12457,[]string{"HY000"},"%s"},
	ER_IB_MSG_633 : {12458,[]string{"HY000"},"%s"},
	ER_IB_MSG_634 : {12459,[]string{"HY000"},"%s"},
	ER_IB_MSG_635 : {12460,[]string{"HY000"},"%s"},
	ER_IB_MSG_636 : {12461,[]string{"HY000"},"%s"},
	ER_IB_MSG_637 : {12462,[]string{"HY000"},"%s"},
	ER_IB_MSG_638 : {12463,[]string{"HY000"},"%s"},
	ER_IB_MSG_639 : {12464,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_640 : {12465,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_641 : {12466,[]string{"HY000"},"%s"},
	ER_IB_MSG_642 : {12467,[]string{"HY000"},"%s"},
	ER_IB_MSG_643 : {12468,[]string{"HY000"},"%s"},
	ER_IB_MSG_644 : {12469,[]string{"HY000"},"%s"},
	ER_IB_MSG_645 : {12470,[]string{"HY000"},"%s"},
	ER_IB_MSG_646 : {12471,[]string{"HY000"},"%s"},
	ER_IB_MSG_647 : {12472,[]string{"HY000"},"%s"},
	ER_IB_MSG_648 : {12473,[]string{"HY000"},"%s"},
	ER_IB_MSG_649 : {12474,[]string{"HY000"},"%s"},
	ER_IB_MSG_650 : {12475,[]string{"HY000"},"%s"},
	ER_IB_MSG_651 : {12476,[]string{"HY000"},"%s"},
	ER_IB_MSG_652 : {12477,[]string{"HY000"},"%s"},
	ER_IB_MSG_DDL_LOG_DELETE_BY_ID_OK : {12478,[]string{"HY000"},"%s"},
	ER_IB_MSG_654 : {12479,[]string{"HY000"},"%s"},
	ER_IB_MSG_655 : {12480,[]string{"HY000"},"%s"},
	ER_IB_MSG_656 : {12481,[]string{"HY000"},"%s"},
	ER_IB_MSG_657 : {12482,[]string{"HY000"},"%s"},
	ER_IB_MSG_658 : {12483,[]string{"HY000"},"%s"},
	ER_IB_MSG_659 : {12484,[]string{"HY000"},"%s"},
	ER_IB_MSG_660 : {12485,[]string{"HY000"},"%s"},
	ER_IB_MSG_661 : {12486,[]string{"HY000"},"%s"},
	ER_IB_MSG_662 : {12487,[]string{"HY000"},"%s"},
	ER_IB_MSG_663 : {12488,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_664 : {12489,[]string{"HY000"},"The transaction log size is too large for innodb_log_buffer_size (%lu >= %lu / 2). Trying to extend it."},
	//OBSOLETE_ER_IB_MSG_665 : {12490,[]string{"HY000"},"innodb_log_buffer_size was extended to %lu bytes."},
	//OBSOLETE_ER_IB_MSG_666 : {12491,[]string{"HY000"},"The transaction log files are too small for the single transaction log (size=%lu). So, the last checkpoint age might exceed the log group capacity %llu."},
	//OBSOLETE_ER_IB_MSG_667 : {12492,[]string{"HY000"},"The age of the last checkpoint is %llu, which exceeds the log group capacity %llu."},
	//OBSOLETE_ER_IB_MSG_668 : {12493,[]string{"HY000"},"Cannot continue operation. ib_logfiles are too small for innodb_thread_concurrency %lu. The combined size of ib_logfiles should be bigger than 200 kB * innodb_thread_concurrency. To get mysqld to start up, set innodb_thread_concurrency in my.cnf to a lower value, for example, to 8. After an ERROR-FREE shutdown of mysqld you can adjust the size of ib_logfiles. %s"},
	//OBSOLETE_ER_IB_MSG_669 : {12494,[]string{"HY000"},"Redo log was encrypted, but keyring plugin is not loaded."},
	//OBSOLETE_ER_IB_MSG_670 : {12495,[]string{"HY000"},"Read redo log encryption metadata successful."},
	//OBSOLETE_ER_IB_MSG_671 : {12496,[]string{"HY000"},"Can't set redo log tablespace encryption metadata."},
	//OBSOLETE_ER_IB_MSG_672 : {12497,[]string{"HY000"},"Cannot read the encryption information in log file header, please check if keyring plugin loaded and the key file exists."},
	//OBSOLETE_ER_IB_MSG_673 : {12498,[]string{"HY000"},"Can't set redo log tablespace to be encrypted in read-only mode."},
	//OBSOLETE_ER_IB_MSG_674 : {12499,[]string{"HY000"},"Can't set redo log tablespace to be encrypted."},
	//OBSOLETE_ER_IB_MSG_675 : {12500,[]string{"HY000"},"Can't set redo log tablespace to be encrypted."},
	//OBSOLETE_ER_IB_MSG_676 : {12501,[]string{"HY000"},"Redo log encryption is enabled."},
	//OBSOLETE_ER_IB_MSG_677 : {12502,[]string{"HY000"},"Flush waiting for archiver to catch up lag LSN: %llu"},
	//OBSOLETE_ER_IB_MSG_678 : {12503,[]string{"HY000"},"Flush overwriting data to archive - wait too long (1 minute) lag LSN: %llu"},
	//OBSOLETE_ER_IB_MSG_679 : {12504,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_680 : {12505,[]string{"HY000"},"Starting shutdown..."},
	//OBSOLETE_ER_IB_MSG_681 : {12506,[]string{"HY000"},"Waiting for %s to exit"},
	//OBSOLETE_ER_IB_MSG_682 : {12507,[]string{"HY000"},"Waiting for %lu active transactions to finish"},
	//OBSOLETE_ER_IB_MSG_683 : {12508,[]string{"HY000"},"Waiting for master thread to be suspended"},
	//OBSOLETE_ER_IB_MSG_684 : {12509,[]string{"HY000"},"Waiting for page_cleaner to finish flushing of buffer pool"},
	//OBSOLETE_ER_IB_MSG_685 : {12510,[]string{"HY000"},"Pending checkpoint_writes: %lu. Pending log flush writes: %lu."},
	//OBSOLETE_ER_IB_MSG_686 : {12511,[]string{"HY000"},"Waiting for %lu buffer page I/Os to complete"},
	//OBSOLETE_ER_IB_MSG_687 : {12512,[]string{"HY000"},"MySQL has requested a very fast shutdown without flushing the InnoDB buffer pool to data files. At the next mysqld startup InnoDB will do a crash recovery!"},
	//OBSOLETE_ER_IB_MSG_688 : {12513,[]string{"HY000"},"Background thread %s woke up during shutdown"},
	//OBSOLETE_ER_IB_MSG_689 : {12514,[]string{"HY000"},"Waiting for archiver to finish archiving page and log"},
	//OBSOLETE_ER_IB_MSG_690 : {12515,[]string{"HY000"},"Background thread %s woke up during shutdown"},
	//OBSOLETE_ER_IB_MSG_691 : {12516,[]string{"HY000"},"Waiting for dirty buffer pages to be flushed"},
	//OBSOLETE_ER_IB_MSG_692 : {12517,[]string{"HY000"},"Log sequence number at shutdown %llu is lower than at startup %llu!"},
	//OBSOLETE_ER_IB_MSG_693 : {12518,[]string{"HY000"},"Waiting for archiver to finish archiving page and log"},
	ER_IB_MSG_694 : {12519,[]string{"HY000"},"############### CORRUPT LOG RECORD FOUND ###############"},
	ER_IB_MSG_695 : {12520,[]string{"HY000"},"Log record type %d, page %lu:%lu. Log parsing proceeded successfully up to %llu. Previous log record type %d, is multi %llu Recv offset %zd, prev %llu"},
	ER_IB_MSG_696 : {12521,[]string{"HY000"},"Hex dump starting %llu bytes before and ending %llu bytes after the corrupted record:"},
	ER_IB_MSG_697 : {12522,[]string{"HY000"},"Set innodb_force_recovery to ignore this error."},
	ER_IB_MSG_698 : {12523,[]string{"HY000"},"The log file may have been corrupt and it is possible that the log scan did not proceed far enough in recovery! Please run CHECK TABLE on your InnoDB tables to check that they are ok! If mysqld crashes after this recovery; %s"},
	ER_IB_MSG_699 : {12524,[]string{"HY000"},"%llu pages with log records were left unprocessed!"},
	ER_IB_MSG_700 : {12525,[]string{"HY000"},"%s"},
	ER_IB_MSG_701 : {12526,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_702 : {12527,[]string{"HY000"},"Invalid redo log header checksum."},
	//OBSOLETE_ER_IB_MSG_703 : {12528,[]string{"HY000"},"Unsupported redo log format (%lu). The redo log was created before MySQL 5.7.9"},
	ER_IB_MSG_704 : {12529,[]string{"HY000"},"Redo log format is v%lu. The redo log was created before MySQL 8.0.3."},
	ER_IB_MSG_705 : {12530,[]string{"HY000"},"Unknown redo log format (%lu). Please follow the instructions at %s upgrading-downgrading.html."},
	ER_IB_MSG_706 : {12531,[]string{"HY000"},"No valid checkpoint found (corrupted redo log). You can try --innodb-force-recovery=6 as a last resort."},
	ER_IB_MSG_707 : {12532,[]string{"HY000"},"Applying a batch of %llu redo log records ..."},
	ER_IB_MSG_708 : {12533,[]string{"HY000"},"%s"},
	ER_IB_MSG_709 : {12534,[]string{"HY000"},"%s"},
	ER_IB_MSG_710 : {12535,[]string{"HY000"},"Apply batch completed!"},
	ER_IB_MSG_711 : {12536,[]string{"HY000"},"%s"},
	ER_IB_MSG_712 : {12537,[]string{"HY000"},"%s"},
	ER_IB_MSG_713 : {12538,[]string{"HY000"},"%s"},
	ER_IB_MSG_714 : {12539,[]string{"HY000"},"%s"},
	ER_IB_MSG_715 : {12540,[]string{"HY000"},"%s"},
	ER_IB_MSG_716 : {12541,[]string{"HY000"},"%s"},
	ER_IB_MSG_717 : {12542,[]string{"HY000"},"An optimized(without redo logging) DDL operation has been performed. All modified pages may not have been flushed to the disk yet.\nThis offline backup may not be consistent"},
	ER_IB_MSG_718 : {12543,[]string{"HY000"},"Extending tablespace : %lu space name: %s to new size: %lu pages during recovery."},
	ER_IB_MSG_719 : {12544,[]string{"HY000"},"Could not extend tablespace: %lu space name: %s to new size: %lu pages during recovery."},
	ER_IB_MSG_720 : {12545,[]string{"HY000"},"Log block %llu at lsn %llu has valid header, but checksum field contains %lu, should be %lu."},
	ER_IB_MSG_721 : {12546,[]string{"HY000"},"Recovery skipped, --innodb-read-only set!"},
	ER_IB_MSG_722 : {12547,[]string{"HY000"},"Log scan progressed past the checkpoint LSN %llu."},
	ER_IB_MSG_723 : {12548,[]string{"HY000"},"Log parsing buffer overflow. Recovery may have failed! Please set log_buffer_size to a value higher than %lu."},
	ER_IB_MSG_724 : {12549,[]string{"HY000"},"Set innodb_force_recovery to ignore this error."},
	ER_IB_MSG_725 : {12550,[]string{"HY000"},"Doing recovery: scanned up to log sequence number %llu"},
	ER_IB_MSG_726 : {12551,[]string{"HY000"},"Database was not shutdown normally!"},
	ER_IB_MSG_727 : {12552,[]string{"HY000"},"Starting crash recovery."},
	ER_IB_MSG_728 : {12553,[]string{"HY000"},"The user has set SRV_FORCE_NO_LOG_REDO on, skipping log redo"},
	ER_IB_MSG_729 : {12554,[]string{"HY000"},"Cannot restore from mysqlbackup, InnoDB running in read-only mode!"},
	ER_IB_MSG_730 : {12555,[]string{"HY000"},"The log file was created by mysqlbackup --apply-log at %s. The following crash recovery is part of a normal restore."},
	ER_IB_MSG_731 : {12556,[]string{"HY000"},"Opening cloned database"},
	ER_IB_MSG_732 : {12557,[]string{"HY000"},"Redo log is from an earlier version, v%lu."},
	ER_IB_MSG_733 : {12558,[]string{"HY000"},"Redo log format v%lu not supported. Current supported format is v%lu."},
	ER_IB_MSG_734 : {12559,[]string{"HY000"},"Are you sure you are using the right ib_logfiles to start up the database? Log sequence number in the ib_logfiles is %llu, less than the log sequence number in the first system tablespace file header, %llu."},
	ER_IB_MSG_735 : {12560,[]string{"HY000"},"The log sequence number %llu in the system tablespace does not match the log sequence number %llu in the ib_logfiles!"},
	ER_IB_MSG_736 : {12561,[]string{"HY000"},"Can't initiate database recovery, running in read-only-mode."},
	ER_IB_MSG_737 : {12562,[]string{"HY000"},"We scanned the log up to %llu. A checkpoint was at %llu and the maximum LSN on a database page was %llu. It is possible that the database is now corrupt!"},
	ER_IB_MSG_738 : {12563,[]string{"HY000"},"Waiting for recv_writer to finish flushing of buffer pool"},
	ER_IB_MSG_739 : {12564,[]string{"HY000"},"Recovery parsing buffer extended to %zu."},
	ER_IB_MSG_740 : {12565,[]string{"HY000"},"Out of memory while resizing recovery parsing buffer."},
	ER_IB_MSG_741 : {12566,[]string{"HY000"},"%s"},
	ER_IB_MSG_742 : {12567,[]string{"HY000"},"%s"},
	ER_IB_MSG_743 : {12568,[]string{"HY000"},"%s"},
	ER_IB_MSG_744 : {12569,[]string{"HY000"},"%s"},
	ER_IB_MSG_745 : {12570,[]string{"HY000"},"%s"},
	ER_IB_MSG_746 : {12571,[]string{"HY000"},"%s"},
	ER_IB_MSG_747 : {12572,[]string{"HY000"},"%s"},
	ER_IB_MSG_748 : {12573,[]string{"HY000"},"%s"},
	ER_IB_MSG_749 : {12574,[]string{"HY000"},"%s"},
	ER_IB_MSG_750 : {12575,[]string{"HY000"},"%s"},
	ER_IB_MSG_751 : {12576,[]string{"HY000"},"%s"},
	ER_IB_MSG_752 : {12577,[]string{"HY000"},"%s"},
	ER_IB_MSG_753 : {12578,[]string{"HY000"},"%s"},
	ER_IB_MSG_754 : {12579,[]string{"HY000"},"%s"},
	ER_IB_MSG_755 : {12580,[]string{"HY000"},"%s"},
	ER_IB_MSG_756 : {12581,[]string{"HY000"},"%s"},
	ER_IB_MSG_757 : {12582,[]string{"HY000"},"%s"},
	ER_IB_MSG_758 : {12583,[]string{"HY000"},"%s"},
	ER_IB_MSG_759 : {12584,[]string{"HY000"},"%s"},
	ER_IB_MSG_760 : {12585,[]string{"HY000"},"%s"},
	ER_IB_MSG_761 : {12586,[]string{"HY000"},"%s"},
	ER_IB_MSG_762 : {12587,[]string{"HY000"},"%s"},
	ER_IB_MSG_763 : {12588,[]string{"HY000"},"%s"},
	ER_IB_MSG_764 : {12589,[]string{"HY000"},"%s"},
	ER_IB_MSG_765 : {12590,[]string{"HY000"},"%s"},
	ER_IB_MSG_766 : {12591,[]string{"HY000"},"%s"},
	ER_IB_MSG_767 : {12592,[]string{"HY000"},"%s"},
	ER_IB_MSG_768 : {12593,[]string{"HY000"},"%s"},
	ER_IB_MSG_769 : {12594,[]string{"HY000"},"%s"},
	ER_IB_MSG_770 : {12595,[]string{"HY000"},"%s"},
	ER_IB_MSG_771 : {12596,[]string{"HY000"},"%s"},
	ER_IB_MSG_772 : {12597,[]string{"HY000"},"%s"},
	ER_IB_MSG_773 : {12598,[]string{"HY000"},"%s"},
	ER_IB_MSG_774 : {12599,[]string{"HY000"},"%s"},
	ER_IB_MSG_775 : {12600,[]string{"HY000"},"%s"},
	ER_IB_MSG_776 : {12601,[]string{"HY000"},"%s"},
	ER_IB_MSG_777 : {12602,[]string{"HY000"},"%s"},
	ER_IB_MSG_778 : {12603,[]string{"HY000"},"%s"},
	ER_IB_MSG_779 : {12604,[]string{"HY000"},"%s"},
	ER_IB_MSG_780 : {12605,[]string{"HY000"},"%s"},
	ER_IB_MSG_781 : {12606,[]string{"HY000"},"%s"},
	ER_IB_MSG_782 : {12607,[]string{"HY000"},"%s"},
	ER_IB_MSG_783 : {12608,[]string{"HY000"},"%s"},
	ER_IB_MSG_784 : {12609,[]string{"HY000"},"%s"},
	ER_IB_MSG_785 : {12610,[]string{"HY000"},"%s"},
	ER_IB_MSG_786 : {12611,[]string{"HY000"},"%s"},
	ER_IB_MSG_787 : {12612,[]string{"HY000"},"%s"},
	ER_IB_MSG_788 : {12613,[]string{"HY000"},"%s"},
	ER_IB_MSG_789 : {12614,[]string{"HY000"},"%s"},
	ER_IB_MSG_790 : {12615,[]string{"HY000"},"%s"},
	ER_IB_MSG_791 : {12616,[]string{"HY000"},"%s"},
	ER_IB_MSG_792 : {12617,[]string{"HY000"},"%s"},
	ER_IB_MSG_793 : {12618,[]string{"HY000"},"%s"},
	ER_IB_MSG_794 : {12619,[]string{"HY000"},"%s"},
	ER_IB_MSG_795 : {12620,[]string{"HY000"},"%s"},
	ER_IB_MSG_796 : {12621,[]string{"HY000"},"%s"},
	ER_IB_MSG_797 : {12622,[]string{"HY000"},"%s"},
	ER_IB_MSG_798 : {12623,[]string{"HY000"},"%s"},
	ER_IB_MSG_799 : {12624,[]string{"HY000"},"%s"},
	ER_IB_MSG_800 : {12625,[]string{"HY000"},"%s"},
	ER_IB_MSG_801 : {12626,[]string{"HY000"},"%s"},
	ER_IB_MSG_802 : {12627,[]string{"HY000"},"%s"},
	ER_IB_MSG_803 : {12628,[]string{"HY000"},"%s"},
	ER_IB_MSG_804 : {12629,[]string{"HY000"},"%s"},
	ER_IB_MSG_805 : {12630,[]string{"HY000"},"%s"},
	ER_IB_MSG_806 : {12631,[]string{"HY000"},"%s"},
	ER_IB_MSG_807 : {12632,[]string{"HY000"},"%s"},
	ER_IB_MSG_808 : {12633,[]string{"HY000"},"%s"},
	ER_IB_MSG_809 : {12634,[]string{"HY000"},"%s"},
	ER_IB_MSG_810 : {12635,[]string{"HY000"},"%s"},
	ER_IB_MSG_811 : {12636,[]string{"HY000"},"%s"},
	ER_IB_MSG_812 : {12637,[]string{"HY000"},"%s"},
	ER_IB_MSG_813 : {12638,[]string{"HY000"},"%s"},
	ER_IB_MSG_814 : {12639,[]string{"HY000"},"%s"},
	ER_IB_MSG_815 : {12640,[]string{"HY000"},"%s"},
	ER_IB_MSG_816 : {12641,[]string{"HY000"},"%s"},
	ER_IB_MSG_817 : {12642,[]string{"HY000"},"%s"},
	ER_IB_MSG_818 : {12643,[]string{"HY000"},"%s"},
	ER_IB_MSG_819 : {12644,[]string{"HY000"},"%s"},
	ER_IB_MSG_820 : {12645,[]string{"HY000"},"%s"},
	ER_IB_MSG_821 : {12646,[]string{"HY000"},"%s"},
	ER_IB_MSG_822 : {12647,[]string{"HY000"},"%s"},
	ER_IB_MSG_823 : {12648,[]string{"HY000"},"%s"},
	ER_IB_MSG_824 : {12649,[]string{"HY000"},"%s"},
	ER_IB_MSG_825 : {12650,[]string{"HY000"},"%s"},
	ER_IB_MSG_826 : {12651,[]string{"HY000"},"%s"},
	ER_IB_MSG_827 : {12652,[]string{"HY000"},"%s"},
	ER_IB_MSG_828 : {12653,[]string{"HY000"},"%s"},
	ER_IB_MSG_829 : {12654,[]string{"HY000"},"%s"},
	ER_IB_MSG_830 : {12655,[]string{"HY000"},"%s"},
	ER_IB_MSG_831 : {12656,[]string{"HY000"},"%s"},
	ER_IB_MSG_832 : {12657,[]string{"HY000"},"%s"},
	ER_IB_MSG_833 : {12658,[]string{"HY000"},"%s"},
	ER_IB_MSG_834 : {12659,[]string{"HY000"},"%s"},
	ER_IB_MSG_835 : {12660,[]string{"HY000"},"%s"},
	ER_IB_MSG_836 : {12661,[]string{"HY000"},"%s"},
	ER_IB_MSG_837 : {12662,[]string{"HY000"},"%s"},
	ER_IB_MSG_838 : {12663,[]string{"HY000"},"%s"},
	ER_IB_MSG_839 : {12664,[]string{"HY000"},"%s"},
	ER_IB_MSG_840 : {12665,[]string{"HY000"},"%s"},
	ER_IB_MSG_841 : {12666,[]string{"HY000"},"%s"},
	ER_IB_MSG_842 : {12667,[]string{"HY000"},"%s"},
	ER_IB_MSG_843 : {12668,[]string{"HY000"},"%s"},
	ER_IB_MSG_844 : {12669,[]string{"HY000"},"%s"},
	ER_IB_MSG_845 : {12670,[]string{"HY000"},"%s"},
	ER_IB_MSG_846 : {12671,[]string{"HY000"},"%s"},
	ER_IB_MSG_847 : {12672,[]string{"HY000"},"%s"},
	ER_IB_MSG_848 : {12673,[]string{"HY000"},"%s"},
	ER_IB_MSG_849 : {12674,[]string{"HY000"},"%s"},
	ER_IB_MSG_850 : {12675,[]string{"HY000"},"%s"},
	ER_IB_MSG_851 : {12676,[]string{"HY000"},"%s"},
	ER_IB_MSG_852 : {12677,[]string{"HY000"},"%s"},
	ER_IB_MSG_853 : {12678,[]string{"HY000"},"%s"},
	ER_IB_MSG_854 : {12679,[]string{"HY000"},"%s"},
	ER_IB_MSG_855 : {12680,[]string{"HY000"},"%s"},
	ER_IB_MSG_856 : {12681,[]string{"HY000"},"%s"},
	ER_IB_MSG_857 : {12682,[]string{"HY000"},"%s"},
	ER_IB_MSG_858 : {12683,[]string{"HY000"},"%s"},
	ER_IB_MSG_859 : {12684,[]string{"HY000"},"%s"},
	ER_IB_MSG_860 : {12685,[]string{"HY000"},"%s"},
	ER_IB_MSG_861 : {12686,[]string{"HY000"},"%s"},
	ER_IB_MSG_862 : {12687,[]string{"HY000"},"%s"},
	ER_IB_MSG_863 : {12688,[]string{"HY000"},"%s"},
	ER_IB_MSG_864 : {12689,[]string{"HY000"},"%s"},
	ER_IB_MSG_865 : {12690,[]string{"HY000"},"%s"},
	ER_IB_MSG_866 : {12691,[]string{"HY000"},"%s"},
	ER_IB_MSG_867 : {12692,[]string{"HY000"},"%s"},
	ER_IB_MSG_868 : {12693,[]string{"HY000"},"%s"},
	ER_IB_MSG_869 : {12694,[]string{"HY000"},"%s"},
	ER_IB_MSG_870 : {12695,[]string{"HY000"},"%s"},
	ER_IB_MSG_871 : {12696,[]string{"HY000"},"%s"},
	ER_IB_MSG_872 : {12697,[]string{"HY000"},"%s"},
	ER_IB_MSG_873 : {12698,[]string{"HY000"},"%s"},
	ER_IB_MSG_874 : {12699,[]string{"HY000"},"%s"},
	ER_IB_MSG_875 : {12700,[]string{"HY000"},"%s"},
	ER_IB_MSG_876 : {12701,[]string{"HY000"},"%s"},
	ER_IB_MSG_877 : {12702,[]string{"HY000"},"%s"},
	ER_IB_MSG_878 : {12703,[]string{"HY000"},"%s"},
	ER_IB_MSG_879 : {12704,[]string{"HY000"},"%s"},
	ER_IB_MSG_880 : {12705,[]string{"HY000"},"%s"},
	ER_IB_MSG_881 : {12706,[]string{"HY000"},"%s"},
	ER_IB_MSG_882 : {12707,[]string{"HY000"},"%s"},
	ER_IB_MSG_883 : {12708,[]string{"HY000"},"%s"},
	ER_IB_MSG_884 : {12709,[]string{"HY000"},"%s"},
	ER_IB_MSG_885 : {12710,[]string{"HY000"},"%s"},
	ER_IB_MSG_886 : {12711,[]string{"HY000"},"%s"},
	ER_IB_MSG_887 : {12712,[]string{"HY000"},"%s"},
	ER_IB_MSG_888 : {12713,[]string{"HY000"},"%s"},
	ER_IB_MSG_889 : {12714,[]string{"HY000"},"%s"},
	ER_IB_MSG_890 : {12715,[]string{"HY000"},"%s"},
	ER_IB_MSG_891 : {12716,[]string{"HY000"},"%s"},
	ER_IB_MSG_892 : {12717,[]string{"HY000"},"%s"},
	ER_IB_MSG_893 : {12718,[]string{"HY000"},"%s"},
	ER_IB_MSG_894 : {12719,[]string{"HY000"},"%s"},
	ER_IB_MSG_895 : {12720,[]string{"HY000"},"%s"},
	ER_IB_MSG_896 : {12721,[]string{"HY000"},"%s"},
	ER_IB_MSG_897 : {12722,[]string{"HY000"},"%s"},
	ER_IB_MSG_898 : {12723,[]string{"HY000"},"%s"},
	ER_IB_MSG_899 : {12724,[]string{"HY000"},"%s"},
	ER_IB_MSG_900 : {12725,[]string{"HY000"},"%s"},
	ER_IB_MSG_901 : {12726,[]string{"HY000"},"%s"},
	ER_IB_MSG_902 : {12727,[]string{"HY000"},"%s"},
	ER_IB_MSG_903 : {12728,[]string{"HY000"},"%s"},
	ER_IB_MSG_904 : {12729,[]string{"HY000"},"%s"},
	ER_IB_MSG_905 : {12730,[]string{"HY000"},"%s"},
	ER_IB_MSG_906 : {12731,[]string{"HY000"},"%s"},
	ER_IB_MSG_907 : {12732,[]string{"HY000"},"%s"},
	ER_IB_MSG_908 : {12733,[]string{"HY000"},"%s"},
	ER_IB_MSG_909 : {12734,[]string{"HY000"},"%s"},
	ER_IB_MSG_910 : {12735,[]string{"HY000"},"%s"},
	ER_IB_MSG_911 : {12736,[]string{"HY000"},"%s"},
	ER_IB_MSG_912 : {12737,[]string{"HY000"},"%s"},
	ER_IB_MSG_913 : {12738,[]string{"HY000"},"%s"},
	ER_IB_MSG_914 : {12739,[]string{"HY000"},"%s"},
	ER_IB_MSG_915 : {12740,[]string{"HY000"},"%s"},
	ER_IB_MSG_916 : {12741,[]string{"HY000"},"%s"},
	ER_IB_MSG_917 : {12742,[]string{"HY000"},"%s"},
	ER_IB_MSG_918 : {12743,[]string{"HY000"},"%s"},
	ER_IB_MSG_919 : {12744,[]string{"HY000"},"%s"},
	ER_IB_MSG_920 : {12745,[]string{"HY000"},"%s"},
	ER_IB_MSG_921 : {12746,[]string{"HY000"},"%s"},
	ER_IB_MSG_922 : {12747,[]string{"HY000"},"%s"},
	ER_IB_MSG_923 : {12748,[]string{"HY000"},"%s"},
	ER_IB_MSG_924 : {12749,[]string{"HY000"},"%s"},
	ER_IB_MSG_925 : {12750,[]string{"HY000"},"%s"},
	ER_IB_MSG_926 : {12751,[]string{"HY000"},"%s"},
	ER_IB_MSG_927 : {12752,[]string{"HY000"},"%s"},
	ER_IB_MSG_928 : {12753,[]string{"HY000"},"%s"},
	ER_IB_MSG_929 : {12754,[]string{"HY000"},"%s"},
	ER_IB_MSG_930 : {12755,[]string{"HY000"},"%s"},
	ER_IB_MSG_931 : {12756,[]string{"HY000"},"%s"},
	ER_IB_MSG_932 : {12757,[]string{"HY000"},"%s"},
	ER_IB_MSG_933 : {12758,[]string{"HY000"},"%s"},
	ER_IB_MSG_934 : {12759,[]string{"HY000"},"%s"},
	ER_IB_MSG_935 : {12760,[]string{"HY000"},"%s"},
	ER_IB_MSG_936 : {12761,[]string{"HY000"},"%s"},
	ER_IB_MSG_937 : {12762,[]string{"HY000"},"%s"},
	ER_IB_MSG_938 : {12763,[]string{"HY000"},"%s"},
	ER_IB_MSG_939 : {12764,[]string{"HY000"},"%s"},
	ER_IB_MSG_940 : {12765,[]string{"HY000"},"%s"},
	ER_IB_MSG_941 : {12766,[]string{"HY000"},"%s"},
	ER_IB_MSG_942 : {12767,[]string{"HY000"},"%s"},
	ER_IB_MSG_943 : {12768,[]string{"HY000"},"%s"},
	ER_IB_MSG_944 : {12769,[]string{"HY000"},"%s"},
	ER_IB_MSG_945 : {12770,[]string{"HY000"},"%s"},
	ER_IB_MSG_946 : {12771,[]string{"HY000"},"%s"},
	ER_IB_MSG_947 : {12772,[]string{"HY000"},"%s"},
	ER_IB_MSG_948 : {12773,[]string{"HY000"},"%s"},
	ER_IB_MSG_949 : {12774,[]string{"HY000"},"%s"},
	ER_IB_MSG_950 : {12775,[]string{"HY000"},"%s"},
	ER_IB_MSG_951 : {12776,[]string{"HY000"},"%s"},
	ER_IB_MSG_952 : {12777,[]string{"HY000"},"%s"},
	ER_IB_MSG_953 : {12778,[]string{"HY000"},"%s"},
	ER_IB_MSG_954 : {12779,[]string{"HY000"},"%s"},
	ER_IB_MSG_955 : {12780,[]string{"HY000"},"%s"},
	ER_IB_MSG_956 : {12781,[]string{"HY000"},"%s"},
	ER_IB_MSG_957 : {12782,[]string{"HY000"},"%s"},
	ER_IB_MSG_958 : {12783,[]string{"HY000"},"%s"},
	ER_IB_MSG_959 : {12784,[]string{"HY000"},"%s"},
	ER_IB_MSG_960 : {12785,[]string{"HY000"},"%s"},
	ER_IB_MSG_961 : {12786,[]string{"HY000"},"%s"},
	ER_IB_MSG_962 : {12787,[]string{"HY000"},"%s"},
	ER_IB_MSG_963 : {12788,[]string{"HY000"},"%s"},
	ER_IB_MSG_964 : {12789,[]string{"HY000"},"%s"},
	ER_IB_MSG_965 : {12790,[]string{"HY000"},"%s"},
	ER_IB_MSG_966 : {12791,[]string{"HY000"},"%s"},
	ER_IB_MSG_967 : {12792,[]string{"HY000"},"%s"},
	ER_IB_MSG_968 : {12793,[]string{"HY000"},"%s"},
	ER_IB_MSG_969 : {12794,[]string{"HY000"},"%s"},
	ER_IB_MSG_970 : {12795,[]string{"HY000"},"%s"},
	ER_IB_MSG_971 : {12796,[]string{"HY000"},"%s"},
	ER_IB_MSG_972 : {12797,[]string{"HY000"},"%s"},
	ER_IB_MSG_973 : {12798,[]string{"HY000"},"%s"},
	ER_IB_MSG_974 : {12799,[]string{"HY000"},"%s"},
	ER_IB_MSG_975 : {12800,[]string{"HY000"},"%s"},
	ER_IB_MSG_976 : {12801,[]string{"HY000"},"%s"},
	ER_IB_MSG_977 : {12802,[]string{"HY000"},"%s"},
	ER_IB_MSG_978 : {12803,[]string{"HY000"},"%s"},
	ER_IB_MSG_979 : {12804,[]string{"HY000"},"%s"},
	ER_IB_MSG_980 : {12805,[]string{"HY000"},"%s"},
	ER_IB_MSG_981 : {12806,[]string{"HY000"},"%s"},
	ER_IB_MSG_982 : {12807,[]string{"HY000"},"%s"},
	ER_IB_MSG_983 : {12808,[]string{"HY000"},"%s"},
	ER_IB_MSG_984 : {12809,[]string{"HY000"},"%s"},
	ER_IB_MSG_985 : {12810,[]string{"HY000"},"%s"},
	ER_IB_MSG_986 : {12811,[]string{"HY000"},"%s"},
	ER_IB_MSG_987 : {12812,[]string{"HY000"},"%s"},
	ER_IB_MSG_988 : {12813,[]string{"HY000"},"%s"},
	ER_IB_MSG_989 : {12814,[]string{"HY000"},"%s"},
	ER_IB_MSG_990 : {12815,[]string{"HY000"},"%s"},
	ER_IB_MSG_991 : {12816,[]string{"HY000"},"%s"},
	ER_IB_MSG_992 : {12817,[]string{"HY000"},"%s"},
	ER_IB_MSG_993 : {12818,[]string{"HY000"},"%s"},
	ER_IB_MSG_994 : {12819,[]string{"HY000"},"%s"},
	ER_IB_MSG_995 : {12820,[]string{"HY000"},"%s"},
	ER_IB_MSG_996 : {12821,[]string{"HY000"},"%s"},
	ER_IB_MSG_997 : {12822,[]string{"HY000"},"%s"},
	ER_IB_MSG_998 : {12823,[]string{"HY000"},"%s"},
	ER_IB_MSG_999 : {12824,[]string{"HY000"},"%s"},
	ER_IB_MSG_1000 : {12825,[]string{"HY000"},"%s"},
	ER_IB_MSG_1001 : {12826,[]string{"HY000"},"%s"},
	ER_IB_MSG_1002 : {12827,[]string{"HY000"},"%s"},
	ER_IB_MSG_1003 : {12828,[]string{"HY000"},"%s"},
	ER_IB_MSG_1004 : {12829,[]string{"HY000"},"%s"},
	ER_IB_MSG_1005 : {12830,[]string{"HY000"},"%s"},
	ER_IB_MSG_1006 : {12831,[]string{"HY000"},"%s"},
	ER_IB_MSG_1007 : {12832,[]string{"HY000"},"%s"},
	ER_IB_MSG_1008 : {12833,[]string{"HY000"},"%s"},
	ER_IB_MSG_1009 : {12834,[]string{"HY000"},"%s"},
	ER_IB_MSG_1010 : {12835,[]string{"HY000"},"%s"},
	ER_IB_MSG_1011 : {12836,[]string{"HY000"},"%s"},
	ER_IB_MSG_1012 : {12837,[]string{"HY000"},"%s"},
	ER_IB_MSG_1013 : {12838,[]string{"HY000"},"%s"},
	ER_IB_MSG_1014 : {12839,[]string{"HY000"},"%s"},
	ER_IB_MSG_1015 : {12840,[]string{"HY000"},"%s"},
	ER_IB_MSG_1016 : {12841,[]string{"HY000"},"%s"},
	ER_IB_MSG_1017 : {12842,[]string{"HY000"},"%s"},
	ER_IB_MSG_1018 : {12843,[]string{"HY000"},"%s"},
	ER_IB_MSG_1019 : {12844,[]string{"HY000"},"%s"},
	ER_IB_MSG_1020 : {12845,[]string{"HY000"},"%s"},
	ER_IB_MSG_1021 : {12846,[]string{"HY000"},"%s"},
	ER_IB_MSG_1022 : {12847,[]string{"HY000"},"%s"},
	ER_IB_MSG_1023 : {12848,[]string{"HY000"},"%s"},
	ER_IB_MSG_1024 : {12849,[]string{"HY000"},"%s"},
	ER_IB_MSG_1025 : {12850,[]string{"HY000"},"%s"},
	ER_IB_MSG_1026 : {12851,[]string{"HY000"},"%s"},
	ER_IB_MSG_1027 : {12852,[]string{"HY000"},"%s"},
	ER_IB_MSG_1028 : {12853,[]string{"HY000"},"%s"},
	ER_IB_MSG_1029 : {12854,[]string{"HY000"},"%s"},
	ER_IB_MSG_1030 : {12855,[]string{"HY000"},"%s"},
	ER_IB_MSG_1031 : {12856,[]string{"HY000"},"%s"},
	ER_IB_MSG_1032 : {12857,[]string{"HY000"},"%s"},
	ER_IB_MSG_1033 : {12858,[]string{"HY000"},"%s"},
	ER_IB_MSG_1034 : {12859,[]string{"HY000"},"%s"},
	ER_IB_MSG_1035 : {12860,[]string{"HY000"},"%s"},
	ER_IB_MSG_1036 : {12861,[]string{"HY000"},"%s"},
	ER_IB_MSG_1037 : {12862,[]string{"HY000"},"%s"},
	ER_IB_MSG_1038 : {12863,[]string{"HY000"},"%s"},
	ER_IB_MSG_1039 : {12864,[]string{"HY000"},"%s"},
	ER_IB_MSG_1040 : {12865,[]string{"HY000"},"%s"},
	ER_IB_MSG_1041 : {12866,[]string{"HY000"},"%s"},
	ER_IB_MSG_1042 : {12867,[]string{"HY000"},"%s"},
	ER_IB_MSG_1043 : {12868,[]string{"HY000"},"%s"},
	ER_IB_MSG_1044 : {12869,[]string{"HY000"},"%s"},
	ER_IB_MSG_1045 : {12870,[]string{"HY000"},"%s"},
	ER_IB_MSG_1046 : {12871,[]string{"HY000"},"Old log sequence number %llu was greater than the new log sequence number %llu. Please submit a bug report to http://bugs.mysql.com"},
	ER_IB_MSG_1047 : {12872,[]string{"HY000"},"Semaphore wait has lasted > %llu seconds. We intentionally crash the server because it appears to be hung."},
	ER_IB_MSG_1048 : {12873,[]string{"HY000"},"Waiting for %llu table(s) to be dropped"},
	ER_IB_MSG_1049 : {12874,[]string{"HY000"},"Waiting for change buffer merge to complete number of bytes of change buffer just merged: %llu"},
	//OBSOLETE_ER_IB_MSG_1050 : {12875,[]string{"HY000"},"Can't set undo tablespace(s) to be encrypted since --innodb_undo_tablespaces=0."},
	ER_IB_MSG_1051 : {12876,[]string{"HY000"},"Can't set undo tablespace(s) to be encrypted in read-only-mode."},
	ER_IB_MSG_1052 : {12877,[]string{"HY000"},"Can't set undo tablespace '%s' to be encrypted."},
	ER_IB_MSG_1053 : {12878,[]string{"HY000"},"Can't set undo tablespace '%s' to be encrypted. Failed to write header page."},
	ER_IB_MSG_1054 : {12879,[]string{"HY000"},"Can't set undo tablespace '%s' to be encrypted. Error %d - %s"},
	ER_IB_MSG_1055 : {12880,[]string{"HY000"},"Encryption is enabled for undo tablespace '%s'."},
	ER_IB_MSG_1056 : {12881,[]string{"HY000"},"Can't rotate encryption on undo tablespace '%s'."},
	ER_IB_MSG_1057 : {12882,[]string{"HY000"},"Encryption is enabled for undo tablespace '%s'."},
	ER_IB_MSG_1058 : {12883,[]string{"HY000"},"os_file_get_status() failed on '%s'. Can't determine file permissions."},
	ER_IB_MSG_1059 : {12884,[]string{"HY000"},"%s can't be opened in %s mode."},
	ER_IB_MSG_1060 : {12885,[]string{"HY000"},"'%s' not a regular file."},
	ER_IB_MSG_1061 : {12886,[]string{"HY000"},"Cannot create %s"},
	ER_IB_MSG_1062 : {12887,[]string{"HY000"},"Setting log file %s size to %llu MB. Progress : %u%%"},
	ER_IB_MSG_1063 : {12888,[]string{"HY000"},"Cannot set log file %s to size %llu MB"},
	ER_IB_MSG_1064 : {12889,[]string{"HY000"},"Cannot create log files in read-only mode"},
	ER_IB_MSG_1065 : {12890,[]string{"HY000"},"Redo log encryption is enabled, but the keyring plugin is not loaded."},
	ER_IB_MSG_1066 : {12891,[]string{"HY000"},"Cannot create file for log file %s."},
	ER_IB_MSG_1067 : {12892,[]string{"HY000"},"Renaming log file %s to %s"},
	ER_IB_MSG_1068 : {12893,[]string{"HY000"},"New log files created, LSN=%llu"},
	ER_IB_MSG_1069 : {12894,[]string{"HY000"},"Unable to open '%s'."},
	ER_IB_MSG_1070 : {12895,[]string{"HY000"},"Cannot create construction log file '%s' for undo tablespace '%s'."},
	ER_IB_MSG_1071 : {12896,[]string{"HY000"},"Creating UNDO Tablespace %s"},
	ER_IB_MSG_1072 : {12897,[]string{"HY000"},"Setting file %s size to %llu MB"},
	ER_IB_MSG_1073 : {12898,[]string{"HY000"},"Physically writing the file full"},
	ER_IB_MSG_1074 : {12899,[]string{"HY000"},"Error in creating %s: probably out of disk space"},
	ER_IB_MSG_1075 : {12900,[]string{"HY000"},"Can't set encryption metadata for space %s"},
	ER_IB_MSG_1076 : {12901,[]string{"HY000"},"Cannot read first page of '%s' - %s"},
	ER_IB_MSG_1077 : {12902,[]string{"HY000"},"Undo tablespace number %lu was being truncated when mysqld quit."},
	ER_IB_MSG_1078 : {12903,[]string{"HY000"},"Cannot recover a truncated undo tablespace in read-only mode"},
	ER_IB_MSG_1079 : {12904,[]string{"HY000"},"Reconstructing undo tablespace number %lu."},
	ER_IB_MSG_1080 : {12905,[]string{"HY000"},"Cannot create %s because %s already uses Space ID=%lu! Did you change innodb_undo_directory?"},
	ER_IB_MSG_1081 : {12906,[]string{"HY000"},"UNDO tablespace %s must be %s"},
	ER_IB_MSG_1082 : {12907,[]string{"HY000"},"Error creating file for %s"},
	ER_IB_MSG_1083 : {12908,[]string{"HY000"},"Error reading encryption for %s"},
	ER_IB_MSG_CANNOT_OPEN_57_UNDO : {12909,[]string{"HY000"},"Unable to open undo tablespace number %lu"},
	ER_IB_MSG_1085 : {12910,[]string{"HY000"},"Opened %llu existing undo tablespaces."},
	ER_IB_MSG_1086 : {12911,[]string{"HY000"},"Cannot create undo tablespaces since innodb_%s has been set. Using %llu existing undo tablespaces."},
	ER_IB_MSG_1087 : {12912,[]string{"HY000"},"Cannot continue InnoDB startup in %s mode because there are no existing undo tablespaces found."},
	ER_IB_MSG_1088 : {12913,[]string{"HY000"},"Could not create undo tablespace '%s'."},
	ER_IB_MSG_1089 : {12914,[]string{"HY000"},"Error %d - %s - opening newly created undo tablespace '%s'."},
	ER_IB_MSG_1090 : {12915,[]string{"HY000"},"Created %llu undo tablespaces."},
	ER_IB_MSG_1091 : {12916,[]string{"HY000"},"Unable to create encrypted undo tablespace number %lu. please check if the keyring plugin is initialized correctly"},
	ER_IB_MSG_1092 : {12917,[]string{"HY000"},"Encryption is enabled for undo tablespace number %lu."},
	ER_IB_MSG_1093 : {12918,[]string{"HY000"},"Unable to initialize the header page in undo tablespace number %lu."},
	ER_IB_MSG_1094 : {12919,[]string{"HY000"},"Cannot delete old undo tablespaces because they contain undo logs for XA PREPARED transactions."},
	ER_IB_MSG_1095 : {12920,[]string{"HY000"},"Upgrading %zu existing undo tablespaces that were tracked in the system tablespace to %lu new independent undo tablespaces."},
	ER_IB_MSG_1096 : {12921,[]string{"HY000"},"Deleting %llu new independent undo tablespaces that we just created."},
	ER_IB_MSG_1097 : {12922,[]string{"HY000"},"Waiting for purge to start"},
	ER_IB_MSG_1098 : {12923,[]string{"HY000"},"Creating shared tablespace for temporary tables"},
	ER_IB_MSG_1099 : {12924,[]string{"HY000"},"The %s data file must be writable!"},
	ER_IB_MSG_1100 : {12925,[]string{"HY000"},"Could not create the shared %s."},
	ER_IB_MSG_1101 : {12926,[]string{"HY000"},"Unable to create the shared %s."},
	ER_IB_MSG_1102 : {12927,[]string{"HY000"},"The %s data file cannot be re-opened after check_file_spec() succeeded!"},
	ER_IB_MSG_1103 : {12928,[]string{"HY000"},"%d threads created by InnoDB had not exited at shutdown!"},
	ER_IB_MSG_1104 : {12929,[]string{"HY000"},"InnoDB Database creation was aborted %swith error %s. You may need to delete the ibdata1 file before trying to start up again."},
	ER_IB_MSG_1105 : {12930,[]string{"HY000"},"Plugin initialization aborted %swith error %s."},
	ER_IB_MSG_1106 : {12931,[]string{"HY000"},"Waiting for %llu buffer page I/Os to complete"},
	ER_IB_MSG_1107 : {12932,[]string{"HY000"},"PUNCH HOLE support available"},
	ER_IB_MSG_1108 : {12933,[]string{"HY000"},"PUNCH HOLE support not available"},
	ER_IB_MSG_1109 : {12934,[]string{"HY000"},"Size of InnoDB's ulint is %zu but size of void* is %zu. The sizes should be the same so that on a 64-bit platforms you can allocate more than 4 GB of memory."},
	ER_IB_MSG_1110 : {12935,[]string{"HY000"},"Database upgrade cannot be accomplished in read-only mode."},
	ER_IB_MSG_1111 : {12936,[]string{"HY000"},"Database upgrade cannot be accomplished with innodb_force_recovery > 0"},
	ER_IB_MSG_1112 : {12937,[]string{"HY000"},"%s"},
	ER_IB_MSG_1113 : {12938,[]string{"HY000"},"%s"},
	ER_IB_MSG_1114 : {12939,[]string{"HY000"},"%s"},
	ER_IB_MSG_1115 : {12940,[]string{"HY000"},"%s"},
	ER_IB_MSG_1116 : {12941,[]string{"HY000"},"%s"},
	ER_IB_MSG_1117 : {12942,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_1118 : {12943,[]string{"HY000"},"%s"},
	ER_IB_MSG_1119 : {12944,[]string{"HY000"},"%s"},
	ER_IB_MSG_1120 : {12945,[]string{"HY000"},"%s"},
	ER_IB_MSG_1121 : {12946,[]string{"HY000"},"%s"},
	ER_IB_MSG_1122 : {12947,[]string{"HY000"},"MySQL was built without a memory barrier capability on this architecture, which might allow a mutex/rw_lock violation under high thread concurrency. This may cause a hang."},
	ER_IB_MSG_1123 : {12948,[]string{"HY000"},"Compressed tables use zlib %s"},
	ER_IB_MSG_1124 : {12949,[]string{"HY000"},"%s"},
	ER_IB_MSG_1125 : {12950,[]string{"HY000"},"Startup called second time during the process lifetime. In the MySQL Embedded Server Library you cannot call server_init() more than once during the process lifetime."},
	ER_IB_MSG_1126 : {12951,[]string{"HY000"},"%s"},
	ER_IB_MSG_1127 : {12952,[]string{"HY000"},"Unable to create monitor file %s: %s"},
	ER_IB_MSG_1128 : {12953,[]string{"HY000"},"Disabling background log and ibuf IO write threads."},
	ER_IB_MSG_1129 : {12954,[]string{"HY000"},"Cannot initialize AIO sub-system"},
	ER_IB_MSG_1130 : {12955,[]string{"HY000"},"Initializing buffer pool, total size = %lf%c, instances = %lu, chunk size =%lf%c "},
	ER_IB_MSG_1131 : {12956,[]string{"HY000"},"Cannot allocate memory for the buffer pool"},
	ER_IB_MSG_1132 : {12957,[]string{"HY000"},"Completed initialization of buffer pool"},
	ER_IB_MSG_1133 : {12958,[]string{"HY000"},"Small buffer pool size (%lluM), the flst_validate() debug function can cause a deadlock if the buffer pool fills up."},
	ER_IB_MSG_1134 : {12959,[]string{"HY000"},"Could not open or create the system tablespace. If you tried to add new data files to the system tablespace, and it failed here, you should now edit innodb_data_file_path in my.cnf back to what it was, and remove the new ibdata files InnoDB created in this failed attempt. InnoDB only wrote those files full of zeros, but did not yet use them in any way. But be careful: do not remove old data files which contain your precious data!"},
	ER_IB_MSG_1135 : {12960,[]string{"HY000"},"Cannot create log files because data files are corrupt or the database was not shut down cleanly after creating the data files."},
	ER_IB_MSG_1136 : {12961,[]string{"HY000"},"Only one log file found"},
	ER_IB_MSG_1137 : {12962,[]string{"HY000"},"Log file %s size %llu is not a multiple of innodb_page_size"},
	ER_IB_MSG_1138 : {12963,[]string{"HY000"},"Log file %s is of different size %llu bytes than other log files %llu bytes!"},
	ER_IB_MSG_1139 : {12964,[]string{"HY000"},"Use --innodb-directories to find the tablespace files. If that fails then use --innodb-force-recovery=1 to ignore this and to permanently lose all changes to the missing tablespace(s)"},
	ER_IB_MSG_1140 : {12965,[]string{"HY000"},"The log file may have been corrupt and it is possible that the log scan or parsing did not proceed far enough in recovery. Please run CHECK TABLE on your InnoDB tables to check that they are ok! It may be safest to recover your InnoDB database from a backup!"},
	ER_IB_MSG_1141 : {12966,[]string{"HY000"},"Cannot resize log files in read-only mode."},
	ER_IB_MSG_1142 : {12967,[]string{"HY000"},"Cannot open DD tablespace."},
	ER_IB_MSG_1143 : {12968,[]string{"HY000"},"Starting to delete and rewrite log files."},
	ER_IB_MSG_1144 : {12969,[]string{"HY000"},"Undo from 5.7 found. It will be purged"},
	ER_IB_MSG_1145 : {12970,[]string{"HY000"},"%s"},
	ER_IB_MSG_1146 : {12971,[]string{"HY000"},"%s"},
	ER_IB_MSG_1147 : {12972,[]string{"HY000"},"Tablespace size stored in header is %lu pages, but the sum of data file sizes is %lu pages"},
	ER_IB_MSG_1148 : {12973,[]string{"HY000"},"Cannot start InnoDB. The tail of the system tablespace is missing. Have you edited innodb_data_file_path in my.cnf in an inappropriate way, removing ibdata files from there? You can set innodb_force_recovery=1 in my.cnf to force a startup if you are trying to recover a badly corrupt database."},
	ER_IB_MSG_1149 : {12974,[]string{"HY000"},"Tablespace size stored in header is %lu pages, but the sum of data file sizes is only %lu pages"},
	ER_IB_MSG_1150 : {12975,[]string{"HY000"},"Cannot start InnoDB. The tail of the system tablespace is missing. Have you edited innodb_data_file_path in my.cnf in an InnoDB: inappropriate way, removing ibdata files from there? You can set innodb_force_recovery=1 in my.cnf to force InnoDB: a startup if you are trying to recover a badly corrupt database."},
	ER_IB_MSG_1151 : {12976,[]string{"HY000"},"%s started; log sequence number %llu"},
	ER_IB_MSG_1152 : {12977,[]string{"HY000"},"Waiting for purge to complete"},
	//OBSOLETE_ER_IB_MSG_1153 : {12978,[]string{"HY000"},"Waiting for dict_stats_thread to exit"},
	ER_IB_MSG_1154 : {12979,[]string{"HY000"},"Query counter shows %lld queries still inside InnoDB at shutdown"},
	ER_IB_MSG_1155 : {12980,[]string{"HY000"},"Shutdown completed; log sequence number %llu"},
	ER_IB_MSG_1156 : {12981,[]string{"HY000"},"Cannot continue operation."},
	ER_IB_MSG_1157 : {12982,[]string{"HY000"},"%s"},
	ER_IB_MSG_1158 : {12983,[]string{"HY000"},"%s"},
	ER_IB_MSG_1159 : {12984,[]string{"HY000"},"%s"},
	ER_IB_MSG_1160 : {12985,[]string{"HY000"},"%s"},
	ER_IB_MSG_1161 : {12986,[]string{"HY000"},"%s"},
	ER_IB_MSG_1162 : {12987,[]string{"HY000"},"%s"},
	ER_IB_MSG_1163 : {12988,[]string{"HY000"},"%s"},
	ER_IB_MSG_1164 : {12989,[]string{"HY000"},"%s"},
	ER_IB_MSG_1165 : {12990,[]string{"HY000"},"%s"},
	ER_IB_MSG_UNDO_TRUNCATE_FAIL_TO_READ_LOG_FILE : {12991,[]string{"HY000"},"Unable to read the existing undo truncate log file '%s'. The error is %s"},
	ER_IB_MSG_UNDO_MARKED_FOR_TRUNCATE : {12992,[]string{"HY000"},"Undo tablespace %s is marked for truncate"},
	//OBSOLETE_ER_IB_MSG_UNDO_INJECT_BEFORE_MDL : {0000,[]string{""},"%s"},
	ER_IB_MSG_UNDO_TRUNCATE_START : {12994,[]string{"HY000"},"Truncating UNDO tablespace %s"},
	//OBSOLETE_ER_IB_MSG_UNDO_INJECT_BEFORE_DDL_LOG_START : {0000,[]string{""},"%s"},
	ER_IB_MSG_UNDO_TRUNCATE_DELAY_BY_LOG_CREATE : {12996,[]string{"HY000"},"Cannot create truncate log for undo tablespace '%s'."},
	//OBSOLETE_ER_IB_MSG_UNDO_INJECT_BEFORE_TRUNCATE : {0000,[]string{""},"%s"},
	ER_IB_MSG_UNDO_TRUNCATE_DELAY_BY_FAILURE : {12998,[]string{"HY000"},"Failed to truncate undo tablespace '%s'."},
	//OBSOLETE_ER_IB_MSG_UNDO_INJECT_BEFORE_STATE_UPDATE : {0000,[]string{""},"%s"},
	ER_IB_MSG_UNDO_TRUNCATE_COMPLETE : {13000,[]string{"HY000"},"Completed truncate of undo tablespace %s."},
	//OBSOLETE_ER_IB_MSG_UNDO_INJECT_TRUNCATE_DONE : {0000,[]string{""},"%s"},
	ER_IB_MSG_1177 : {13002,[]string{"HY000"},"%s"},
	ER_IB_MSG_1178 : {13003,[]string{"HY000"},"%s"},
	ER_IB_MSG_1179 : {13004,[]string{"HY000"},"%s"},
	ER_IB_MSG_1180 : {13005,[]string{"HY000"},"%s"},
	ER_IB_MSG_1181 : {13006,[]string{"HY000"},"%s"},
	ER_IB_MSG_1182 : {13007,[]string{"HY000"},"%s"},
	ER_IB_MSG_1183 : {13008,[]string{"HY000"},"%s"},
	ER_IB_MSG_1184 : {13009,[]string{"HY000"},"%s"},
	ER_IB_MSG_1185 : {13010,[]string{"HY000"},"%s"},
	ER_IB_MSG_1186 : {13011,[]string{"HY000"},"%s"},
	ER_IB_MSG_1187 : {13012,[]string{"HY000"},"%s"},
	ER_IB_MSG_1188 : {13013,[]string{"HY000"},"%s"},
	ER_IB_MSG_1189 : {13014,[]string{"HY000"},"%s"},
	ER_IB_MSG_TRX_RECOVERY_ROLLBACK_COMPLETED : {13015,[]string{"HY000"},"Rollback of non-prepared transactions completed"},
	ER_IB_MSG_1191 : {13016,[]string{"HY000"},"%s"},
	ER_IB_MSG_1192 : {13017,[]string{"HY000"},"%s"},
	ER_IB_MSG_1193 : {13018,[]string{"HY000"},"%s"},
	ER_IB_MSG_1194 : {13019,[]string{"HY000"},"%s"},
	ER_IB_MSG_1195 : {13020,[]string{"HY000"},"%s"},
	ER_IB_MSG_1196 : {13021,[]string{"HY000"},"%s"},
	ER_IB_MSG_1197 : {13022,[]string{"HY000"},"%s"},
	ER_IB_MSG_1198 : {13023,[]string{"HY000"},"%s"},
	ER_IB_MSG_1199 : {13024,[]string{"HY000"},"%s"},
	ER_IB_MSG_1200 : {13025,[]string{"HY000"},"%s"},
	ER_IB_MSG_1201 : {13026,[]string{"HY000"},"%s"},
	ER_IB_MSG_1202 : {13027,[]string{"HY000"},"%s"},
	ER_IB_MSG_1203 : {13028,[]string{"HY000"},"%s"},
	ER_IB_MSG_1204 : {13029,[]string{"HY000"},"%s"},
	ER_IB_MSG_1205 : {13030,[]string{"HY000"},"%s"},
	ER_IB_MSG_1206 : {13031,[]string{"HY000"},"%s"},
	ER_IB_MSG_1207 : {13032,[]string{"HY000"},"%s"},
	ER_IB_MSG_1208 : {13033,[]string{"HY000"},"%s"},
	ER_IB_MSG_1209 : {13034,[]string{"HY000"},"%s"},
	ER_IB_MSG_1210 : {13035,[]string{"HY000"},"%s"},
	ER_IB_MSG_1211 : {13036,[]string{"HY000"},"%s"},
	ER_IB_MSG_1212 : {13037,[]string{"HY000"},"%s"},
	ER_IB_MSG_1213 : {13038,[]string{"HY000"},"gettimeofday() failed: %s"},
	ER_IB_MSG_1214 : {13039,[]string{"HY000"},"Can't create UNDO tablespace %s %s"},
	ER_IB_MSG_1215 : {13040,[]string{"HY000"},"%s"},
	ER_IB_MSG_1216 : {13041,[]string{"HY000"},"%s"},
	ER_IB_MSG_1217 : {13042,[]string{"HY000"},"%s"},
	ER_IB_MSG_1218 : {13043,[]string{"HY000"},"%s"},
	ER_IB_MSG_1219 : {13044,[]string{"HY000"},"%s"},
	ER_IB_MSG_1220 : {13045,[]string{"HY000"},"%s"},
	ER_IB_MSG_1221 : {13046,[]string{"HY000"},"%s"},
	ER_IB_MSG_1222 : {13047,[]string{"HY000"},"%s"},
	ER_IB_MSG_1223 : {13048,[]string{"HY000"},"%s"},
	ER_IB_MSG_1224 : {13049,[]string{"HY000"},"%s"},
	ER_IB_MSG_1225 : {13050,[]string{"HY000"},"%s"},
	ER_IB_MSG_1226 : {13051,[]string{"HY000"},"%s"},
	ER_IB_MSG_1227 : {13052,[]string{"HY000"},"%s"},
	ER_IB_MSG_1228 : {13053,[]string{"HY000"},"%s"},
	ER_IB_MSG_1229 : {13054,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_1230 : {13055,[]string{"HY000"},"%s"},
	ER_IB_MSG_1231 : {13056,[]string{"HY000"},"%s"},
	ER_IB_MSG_1232 : {13057,[]string{"HY000"},"%s"},
	ER_IB_MSG_1233 : {13058,[]string{"HY000"},"%s"},
	ER_IB_MSG_1234 : {13059,[]string{"HY000"},"%s"},
	ER_IB_MSG_1235 : {13060,[]string{"HY000"},"%s"},
	ER_IB_MSG_1236 : {13061,[]string{"HY000"},"%s"},
	ER_IB_MSG_1237 : {13062,[]string{"HY000"},"%s"},
	ER_IB_MSG_1238 : {13063,[]string{"HY000"},"%s"},
	ER_IB_MSG_1239 : {13064,[]string{"HY000"},"%s"},
	ER_IB_MSG_1240 : {13065,[]string{"HY000"},"%s"},
	ER_IB_MSG_1241 : {13066,[]string{"HY000"},"%s"},
	ER_IB_MSG_1242 : {13067,[]string{"HY000"},"Can't set redo log tablespace to be encrypted in read-only mode."},
	ER_IB_MSG_1243 : {13068,[]string{"HY000"},"Can't set redo log tablespace to be encrypted."},
	ER_IB_MSG_1244 : {13069,[]string{"HY000"},"Can't set redo log tablespace to be encrypted."},
	ER_IB_MSG_1245 : {13070,[]string{"HY000"},"Redo log encryption is enabled."},
	ER_IB_MSG_1246 : {13071,[]string{"HY000"},"Waiting for archiver to finish archiving page and log"},
	ER_IB_MSG_1247 : {13072,[]string{"HY000"},"Starting shutdown..."},
	ER_IB_MSG_1248 : {13073,[]string{"HY000"},"Waiting for %s to exit."},
	ER_IB_MSG_1249 : {13074,[]string{"HY000"},"Waiting for rollback of %zu recovered transactions, before shutdown."},
	ER_IB_MSG_1250 : {13075,[]string{"HY000"},"Waiting for master thread to be suspended."},
	ER_IB_MSG_1251 : {13076,[]string{"HY000"},"Waiting for page_cleaner to finish flushing of buffer pool."},
	ER_IB_MSG_1252 : {13077,[]string{"HY000"},"Waiting for %lu buffer page I/Os to complete."},
	ER_IB_MSG_1253 : {13078,[]string{"HY000"},"MySQL has requested a very fast shutdown without flushing the InnoDB buffer pool to data files. At the next mysqld startup InnoDB will do a crash recovery!"},
	//OBSOLETE_ER_IB_MSG_1254 : {13079,[]string{"HY000"},"%s"},
	ER_IB_MSG_1255 : {13080,[]string{"HY000"},"%s"},
	ER_IB_MSG_1256 : {13081,[]string{"HY000"},"%s"},
	ER_IB_MSG_1257 : {13082,[]string{"HY000"},"%s"},
	ER_IB_MSG_1258 : {13083,[]string{"HY000"},"%s"},
	ER_IB_MSG_1259 : {13084,[]string{"HY000"},"%s"},
	ER_IB_MSG_1260 : {13085,[]string{"HY000"},"%s"},
	ER_IB_MSG_1261 : {13086,[]string{"HY000"},"%s"},
	ER_IB_MSG_1262 : {13087,[]string{"HY000"},"%s"},
	ER_IB_MSG_1263 : {13088,[]string{"HY000"},"%s"},
	ER_IB_MSG_1264 : {13089,[]string{"HY000"},"%s"},
	ER_IB_MSG_1265 : {13090,[]string{"HY000"},"%s"},
	ER_IB_MSG_1266 : {13091,[]string{"HY000"},"%s"},
	ER_IB_MSG_1267 : {13092,[]string{"HY000"},"%s"},
	ER_IB_MSG_1268 : {13093,[]string{"HY000"},"%s"},
	ER_IB_MSG_1269 : {13094,[]string{"HY000"},"%s"},
	ER_IB_MSG_1270 : {13095,[]string{"HY000"},"%s"},
	ER_RPL_SLAVE_SQL_THREAD_STOP_CMD_EXEC_TIMEOUT : {13096,[]string{"HY000"},"STOP SLAVE command execution is incomplete: Slave SQL thread got the stop signal, thread is busy, SQL thread will stop once the current task is complete."},
	ER_RPL_SLAVE_IO_THREAD_STOP_CMD_EXEC_TIMEOUT : {13097,[]string{"HY000"},"STOP SLAVE command execution is incomplete: Slave IO thread got the stop signal, thread is busy, IO thread will stop once the current task is complete."},
	ER_RPL_GTID_UNSAFE_STMT_ON_NON_TRANS_TABLE : {13098,[]string{"HY000"},"Statement violates GTID consistency: Updates to non-transactional tables can only be done in either autocommitted statements or single-statement transactions, and never in the same statement as updates to transactional tables."},
	ER_RPL_GTID_UNSAFE_STMT_CREATE_SELECT : {13099,[]string{"HY000"},"Statement violates GTID consistency: CREATE TABLE ... SELECT."},
	//OBSOLETE_ER_RPL_GTID_UNSAFE_STMT_ON_TEMPORARY_TABLE : {13100,[]string{"HY000"},"Statement violates GTID consistency: CREATE TEMPORARY TABLE and DROP TEMPORARY TABLE can only be executed outside transactional context.  These statements are also not allowed in a function or trigger because functions and triggers are also considered to be multi-statement transactions."},
	ER_BINLOG_ROW_VALUE_OPTION_IGNORED : {13101,[]string{"HY000"},"When %.192s, the option binlog_row_value_options=%.192s will be ignored and updates will be written in full format to binary log."},
	ER_BINLOG_USE_V1_ROW_EVENTS_IGNORED : {13102,[]string{"HY000"},"When %.192s, the option log_bin_use_v1_row_events=1 will be ignored and row events will be written in new format to binary log."},
	ER_BINLOG_ROW_VALUE_OPTION_USED_ONLY_FOR_AFTER_IMAGES : {13103,[]string{"HY000"},"When %.192s, the option binlog_row_value_options=%.192s will be used only for the after-image. Full values will be written in the before-image, so the saving in disk space due to binlog_row_value_options is limited to less than 50%%."},
	ER_CONNECTION_ABORTED : {13104,[]string{"HY000"},"Aborted connection %u to db: '%-.192s' user: '%-.48s' host: '%-.255s' (%-.64s)."},
	ER_NORMAL_SERVER_SHUTDOWN : {13105,[]string{"HY000"},"%s: Normal shutdown."},
	ER_KEYRING_MIGRATE_FAILED : {13106,[]string{"HY000"},"Can not perform keyring migration : %s."},
	ER_GRP_RPL_LOWER_CASE_TABLE_NAMES_DIFF_FROM_GRP : {13107,[]string{"HY000"},"The member is configured with a lower_case_table_names option value '%u' different from the group '%u'. The member will now exit the group. If there is existing data on member, it may be incompatible with group if it was created with a lower_case_table_names value different from the group."},
	ER_OOM_SAVE_GTIDS : {13108,[]string{"HY000"},"An out-of-memory error occurred while saving the set of GTIDs from the last binary log into the mysql.gtid_executed table"},
	ER_LCTN_NOT_FOUND : {13109,[]string{"HY000"},"The lower_case_table_names setting for the data dictionary was not found. Starting the server using lower_case_table_names = '%u'."},
	//OBSOLETE_ER_REGEXP_INVALID_CAPTURE_GROUP_NAME : {3887,[]string{"HY000"},"A capture group has an invalid name."},
	ER_COMPONENT_FILTER_WRONG_VALUE : {13111,[]string{"HY000"},"Variable '%-.64s' can't be set to the value of '%-.200s'"},
	ER_XPLUGIN_FAILED_TO_STOP_SERVICES : {13112,[]string{"HY000"},"Stopping services failed with error \"%s\""},
	ER_INCONSISTENT_ERROR : {13113,[]string{"HY000"},"Query caused different errors on master and slave. Error on master: message (format)='%s' error code=%d; Error on slave:actual message='%s', error code=%d. Default database:'%s'. Query:'%s'"},
	ER_SERVER_MASTER_FATAL_ERROR_READING_BINLOG : {13114,[]string{"HY000"},"Got fatal error %d from master when reading data from binary log: '%-.512s'"},
	ER_NETWORK_READ_EVENT_CHECKSUM_FAILURE : {13115,[]string{"HY000"},"Replication event checksum verification failed while reading from network."},
	ER_SLAVE_CREATE_EVENT_FAILURE : {13116,[]string{"HY000"},"Failed to create %s"},
	ER_SLAVE_FATAL_ERROR : {13117,[]string{"HY000"},"Fatal error: %s"},
	ER_SLAVE_HEARTBEAT_FAILURE : {13118,[]string{"HY000"},"Unexpected master's heartbeat data: %s"},
	ER_SLAVE_INCIDENT : {13119,[]string{"HY000"},"The incident %s occurred on the master. Message: %s"},
	ER_SLAVE_MASTER_COM_FAILURE : {13120,[]string{"HY000"},"Master command %s failed: %s"},
	ER_SLAVE_RELAY_LOG_READ_FAILURE : {13121,[]string{"HY000"},"Relay log read failure: %s"},
	ER_SLAVE_RELAY_LOG_WRITE_FAILURE : {13122,[]string{"HY000"},"Relay log write failure: %s"},
	ER_SERVER_SLAVE_MI_INIT_REPOSITORY : {13123,[]string{"HY000"},"Slave failed to initialize master info structure from the repository"},
	ER_SERVER_SLAVE_RLI_INIT_REPOSITORY : {13124,[]string{"HY000"},"Slave failed to initialize relay log info structure from the repository"},
	ER_SERVER_NET_PACKET_TOO_LARGE : {13125,[]string{"HY000"},"Got a packet bigger than 'max_allowed_packet' bytes"},
	ER_SERVER_NO_SYSTEM_TABLE_ACCESS : {13126,[]string{"HY000"},"Access to %.64s '%.64s.%.64s' is rejected."},
	ER_SERVER_UNKNOWN_ERROR : {13127,[]string{"HY000"},"Unknown error"},
	ER_SERVER_UNKNOWN_SYSTEM_VARIABLE : {13128,[]string{"HY000"},"Unknown system variable '%-.64s'"},
	ER_SERVER_NO_SESSION_TO_SEND_TO : {13129,[]string{"HY000"},"A message intended for a client cannot be sent there as no client-session is attached. Therefore, we're sending the information to the error-log instead: MY-%06d - %s"},
	ER_SERVER_NEW_ABORTING_CONNECTION : {13130,[]string{"08S01"},"Aborted connection %u to db: '%-.192s' user: '%-.48s' host: '%-.255s' (%-.64s; diagnostics area: MY-%06d - %-.64s)"},
	ER_SERVER_OUT_OF_SORTMEMORY : {13131,[]string{"HY000"},"Out of sort memory, consider increasing server sort buffer size!"},
	ER_SERVER_RECORD_FILE_FULL : {13132,[]string{"HY000"},"The table '%-.192s' is full!"},
	ER_SERVER_DISK_FULL_NOWAIT : {13133,[]string{"HY000"},"Create table/tablespace '%-.192s' failed, as disk is full."},
	ER_SERVER_HANDLER_ERROR : {13134,[]string{"HY000"},"Handler reported error %d - %s"},
	ER_SERVER_NOT_FORM_FILE : {13135,[]string{"HY000"},"Incorrect information in file: '%-.200s'"},
	ER_SERVER_CANT_OPEN_FILE : {13136,[]string{"HY000"},"Can't open file: '%-.200s' (OS errno: %d - %s)"},
	ER_SERVER_FILE_NOT_FOUND : {13137,[]string{"HY000"},"Can't find file: '%-.200s' (OS errno: %d - %s)"},
	ER_SERVER_FILE_USED : {13138,[]string{"HY000"},"'%-.192s' is locked against change (OS errno: %d - %s)"},
	ER_SERVER_CANNOT_LOAD_FROM_TABLE_V2 : {13139,[]string{"HY000"},"Cannot load from %s.%s. The table is probably corrupted!"},
	ER_ERROR_INFO_FROM_DA : {13140,[]string{"HY000"},"Error in diagnostics area: MY-%06d - %s"},
	ER_SERVER_TABLE_CHECK_FAILED : {13141,[]string{"HY000"},"Incorrect definition of table %s.%s: expected column '%s' at position %d, found '%s'."},
	ER_SERVER_COL_COUNT_DOESNT_MATCH_PLEASE_UPDATE_V2 : {13142,[]string{"HY000"},"The column count of %s.%s is wrong. Expected %d, found %d. Created with MySQL %d, now running %d. Please perform the MySQL upgrade procedure."},
	ER_SERVER_COL_COUNT_DOESNT_MATCH_CORRUPTED_V2 : {13143,[]string{"HY000"},"Column count of %s.%s is wrong. Expected %d, found %d. The table is probably corrupted"},
	ER_SERVER_ACL_TABLE_ERROR : {13144,[]string{"HY000"},""},
	ER_SERVER_SLAVE_INIT_QUERY_FAILED : {13145,[]string{"HY000"},"Slave SQL thread aborted. Can't execute init_slave query, MY-%06d - '%s'"},
	ER_SERVER_SLAVE_CONVERSION_FAILED : {13146,[]string{"HY000"},"Column %d of table '%-.192s.%-.192s' cannot be converted from type '%-.32s' to type '%-.32s'"},
	ER_SERVER_SLAVE_IGNORED_TABLE : {13147,[]string{"HY000"},"Slave SQL thread ignored the query because of replicate-*-table rules"},
	ER_CANT_REPLICATE_ANONYMOUS_WITH_AUTO_POSITION : {13148,[]string{"HY000"},"Cannot replicate anonymous transaction when AUTO_POSITION = 1, at file %.400s, position %lld."},
	ER_CANT_REPLICATE_ANONYMOUS_WITH_GTID_MODE_ON : {13149,[]string{"HY000"},"Cannot replicate anonymous transaction when @@GLOBAL.GTID_MODE = ON, at file %.400s, position %lld."},
	ER_CANT_REPLICATE_GTID_WITH_GTID_MODE_OFF : {13150,[]string{"HY000"},"Cannot replicate GTID-transaction when @@GLOBAL.GTID_MODE = OFF, at file %.400s, position %lld."},
	ER_SERVER_TEST_MESSAGE : {13151,[]string{"HY000"},"Simulated error"},
	ER_AUDIT_LOG_JSON_FILTER_PARSING_ERROR : {13152,[]string{"HY000"},"%s"},
	ER_AUDIT_LOG_JSON_FILTERING_NOT_ENABLED : {13153,[]string{"HY000"},"Audit Log filtering has not been installed."},
	ER_PLUGIN_FAILED_TO_OPEN_TABLES : {13154,[]string{"HY000"},"Failed to open the %s filter tables."},
	ER_PLUGIN_FAILED_TO_OPEN_TABLE : {13155,[]string{"HY000"},"Failed to open '%s.%s' %s table."},
	ER_AUDIT_LOG_JSON_FILTER_NAME_CANNOT_BE_EMPTY : {13156,[]string{"HY000"},"Filter name cannot be empty."},
	ER_AUDIT_LOG_USER_NAME_INVALID_CHARACTER : {13157,[]string{"HY000"},"Invalid character in the user name."},
	ER_AUDIT_LOG_UDF_INSUFFICIENT_PRIVILEGE : {13158,[]string{"HY000"},"Request ignored for '%s'@'%s'. SUPER or AUDIT_ADMIN needed to perform operation"},
	ER_AUDIT_LOG_NO_KEYRING_PLUGIN_INSTALLED : {13159,[]string{"HY000"},"No keyring plugin installed."},
	ER_AUDIT_LOG_HOST_NAME_INVALID_CHARACTER : {13160,[]string{"HY000"},"Invalid character in the host name."},
	ER_AUDIT_LOG_ENCRYPTION_PASSWORD_HAS_NOT_BEEN_SET : {13161,[]string{"HY000"},"Audit log encryption password has not been set; it will be generated automatically. Use audit_log_encryption_password_get to obtain the password or audit_log_encryption_password_set to set a new one."},
	ER_AUDIT_LOG_COULD_NOT_CREATE_AES_KEY : {13162,[]string{"HY000"},"Could not create AES key. OpenSSL's EVP_BytesToKey function failed."},
	ER_AUDIT_LOG_ENCRYPTION_PASSWORD_CANNOT_BE_FETCHED : {13163,[]string{"HY000"},"Audit log encryption password cannot be fetched from the keyring. Password used so far is used for encryption."},
	ER_COULD_NOT_REINITIALIZE_AUDIT_LOG_FILTERS : {13164,[]string{"HY000"},"Could not reinitialize audit log filters."},
	ER_AUDIT_LOG_JSON_USER_NAME_CANNOT_BE_EMPTY : {13165,[]string{"HY000"},"User cannot be empty."},
	ER_AUDIT_LOG_USER_FIRST_CHARACTER_MUST_BE_ALPHANUMERIC : {13166,[]string{"HY000"},"First character of the user name must be alphanumeric."},
	ER_AUDIT_LOG_JSON_FILTER_DOES_NOT_EXIST : {13167,[]string{"HY000"},"Specified filter has not been found."},
	ER_IB_MSG_1271 : {13168,[]string{"HY000"},"Cannot upgrade server earlier than 5.7 to 8.0"},
	ER_STARTING_INIT : {13169,[]string{"HY000"},"%s (mysqld %s) initializing of server in progress as process %lu"},
	ER_ENDING_INIT : {13170,[]string{"HY000"},"%s (mysqld %s) initializing of server has completed"},
	ER_IB_MSG_1272 : {13171,[]string{"HY000"},"Cannot boot server version %lu on data directory built by version %llu. Downgrade is not supported"},
	ER_SERVER_SHUTDOWN_INFO : {13172,[]string{"HY000"},"Received SHUTDOWN from user %s. Shutting down mysqld (Version: %s)."},
	ER_GRP_RPL_PLUGIN_ABORT : {13173,[]string{"HY000"},"The plugin encountered a critical error and will abort: %s"},
	//OBSOLETE_ER_REGEXP_INVALID_FLAG : {3900,[]string{"HY000"},"Invalid match mode flag in regular expression."},
	//OBSOLETE_ER_UPDATE_GTID_PURGED_WITH_GR : {3911,[]string{"HY000"},"Cannot update GTID_PURGED with the Group Replication plugin running"},
	ER_AUDIT_LOG_TABLE_DEFINITION_NOT_UPDATED : {13177,[]string{"HY000"},"'%s.%s' table definition has not been upgraded; Please perform the MySQL upgrade procedure."},
	ER_DD_INITIALIZE_SQL_ERROR : {13178,[]string{"HY000"},"Execution of server-side SQL statement '%s' failed with error code = %d, error message = '%s'."},
	ER_NO_PATH_FOR_SHARED_LIBRARY : {13179,[]string{"HY000"},"No paths allowed for shared library."},
	ER_UDF_ALREADY_EXISTS : {13180,[]string{"HY000"},"Function '%-.192s' already exists."},
	ER_SET_EVENT_FAILED : {13181,[]string{"HY000"},"Got Error: %ld from SetEvent."},
	ER_FAILED_TO_ALLOCATE_SSL_BIO : {13182,[]string{"HY000"},"Error allocating SSL BIO."},
	ER_IB_MSG_1273 : {13183,[]string{"HY000"},"%s"},
	ER_PID_FILEPATH_LOCATIONS_INACCESSIBLE : {13184,[]string{"HY000"},"One or several locations were inaccessible while checking PID filepath."},
	ER_UNKNOWN_VARIABLE_IN_PERSISTED_CONFIG_FILE : {13185,[]string{"HY000"},"Currently unknown variable '%s' was read from the persisted config file."},
	ER_FAILED_TO_HANDLE_DEFAULTS_FILE : {13186,[]string{"HY000"},"Fatal error in defaults handling. Program aborted!"},
	ER_DUPLICATE_SYS_VAR : {13187,[]string{"HY000"},"Duplicate variable name '%s'."},
	ER_FAILED_TO_INIT_SYS_VAR : {13188,[]string{"HY000"},"Failed to initialize system variables."},
	ER_SYS_VAR_NOT_FOUND : {13189,[]string{"HY000"},"Variable name '%s' not found."},
	ER_IB_MSG_1274 : {13190,[]string{"HY000"},"Some (%d) threads are still active"},
	ER_IB_MSG_1275 : {13191,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_TARGET_TS_UNENCRYPTED : {13192,[]string{"HY000"},"Source tablespace is encrypted but target tablespace is not."},
	ER_IB_MSG_WAIT_FOR_ENCRYPT_THREAD : {13193,[]string{"HY000"},"Waiting for tablespace_alter_encrypt_thread to exit"},
	ER_IB_MSG_1277 : {13194,[]string{"HY000"},"%s"},
	ER_IB_MSG_NO_ENCRYPT_PROGRESS_FOUND : {13195,[]string{"HY000"},"%s"},
	ER_IB_MSG_RESUME_OP_FOR_SPACE : {13196,[]string{"HY000"},"%s"},
	ER_IB_MSG_1280 : {13197,[]string{"HY000"},"%s"},
	ER_IB_MSG_1281 : {13198,[]string{"HY000"},"%s"},
	ER_IB_MSG_1282 : {13199,[]string{"HY000"},"%s"},
	ER_IB_MSG_1283 : {13200,[]string{"HY000"},"%s"},
	ER_IB_MSG_1284 : {13201,[]string{"HY000"},"%s"},
	ER_CANT_SET_ERROR_SUPPRESSION_LIST_FROM_COMMAND_LINE : {13202,[]string{"HY000"},"%s: Could not add suppression rule for code \"%s\". Rule-set may be full, or code may not correspond to an error-log message."},
	ER_INVALID_VALUE_OF_BIND_ADDRESSES : {13203,[]string{"HY000"},"Invalid value for command line option bind-addresses: '%s'"},
	ER_RELAY_LOG_SPACE_LIMIT_DISABLED : {13204,[]string{"HY000"},"Ignoring the @@global.relay_log_space_limit option because @@global.relay_log_purge is disabled."},
	ER_GRP_RPL_ERROR_GTID_SET_EXTRACTION : {13205,[]string{"HY000"},"Error when extracting GTID execution information: %s"},
	ER_GRP_RPL_MISSING_GRP_RPL_ACTION_COORDINATOR : {13206,[]string{"HY000"},"Message received without a proper group coordinator module."},
	ER_GRP_RPL_JOIN_WHEN_GROUP_ACTION_RUNNING : {13207,[]string{"HY000"},"A member cannot join the group while a group configuration operation is running."},
	ER_GRP_RPL_JOINER_EXIT_WHEN_GROUP_ACTION_RUNNING : {13208,[]string{"HY000"},"A member is joining the group while a group configuration operation is occurring. The member will now leave the group"},
	ER_GRP_RPL_CHANNEL_THREAD_WHEN_GROUP_ACTION_RUNNING : {13209,[]string{"HY000"},"Can't start slave %s when group replication is running a group configuration operation."},
	ER_GRP_RPL_APPOINTED_PRIMARY_NOT_PRESENT : {13210,[]string{"HY000"},"A primary election was invoked but the requested primary member is not in the group. Request ignored."},
	ER_GRP_RPL_ERROR_ON_MESSAGE_SENDING : {13211,[]string{"HY000"},"Error while sending message. Context: %s"},
	ER_GRP_RPL_CONFIGURATION_ACTION_ERROR : {13212,[]string{"HY000"},"Error while executing a group configuration operation: %s"},
	ER_GRP_RPL_CONFIGURATION_ACTION_LOCAL_TERMINATION : {13213,[]string{"HY000"},"Configuration operation '%s' terminated. %s"},
	ER_GRP_RPL_CONFIGURATION_ACTION_START : {13214,[]string{"HY000"},"Starting group operation local execution: %s"},
	ER_GRP_RPL_CONFIGURATION_ACTION_END : {13215,[]string{"HY000"},"Termination of group operation local execution: %s"},
	ER_GRP_RPL_CONFIGURATION_ACTION_KILLED_ERROR : {13216,[]string{"HY000"},"A configuration change was killed in this member. The member will now leave the group as its configuration may have diverged."},
	ER_GRP_RPL_PRIMARY_ELECTION_PROCESS_ERROR : {13217,[]string{"HY000"},"There was an issue on the primary election process: %s The member will now leave the group."},
	ER_GRP_RPL_PRIMARY_ELECTION_STOP_ERROR : {13218,[]string{"HY000"},"There was an issue when stopping a previous election process: %s"},
	ER_GRP_RPL_NO_STAGE_SERVICE : {13219,[]string{"HY000"},"It was not possible to initialize stage logging for this task. The operation will still run without stage tracking."},
	ER_GRP_RPL_UDF_REGISTER_ERROR : {13220,[]string{"HY000"},"Could not execute the installation of Group Replication UDF function: %s. Check if the function is already present, if so, try to remove it"},
	ER_GRP_RPL_UDF_UNREGISTER_ERROR : {13221,[]string{"HY000"},"Could not uninstall Group Replication UDF functions. Try to remove them manually if present."},
	ER_GRP_RPL_UDF_REGISTER_SERVICE_ERROR : {13222,[]string{"HY000"},"Could not execute the installation of Group Replication UDF functions. Check for other errors in the log and try to reinstall the plugin"},
	ER_GRP_RPL_SERVER_UDF_ERROR : {13223,[]string{"HY000"},"The function '%s' failed. %s"},
	//OBSOLETE_ER_CURRENT_PASSWORD_NOT_REQUIRED : {3893,[]string{"HY000"},"Do not specify the current password while changing it for other users."},
	//OBSOLETE_ER_INCORRECT_CURRENT_PASSWORD : {3891,[]string{"HY000"},"Incorrect current password. Specify the correct password which has to be replaced."},
	//OBSOLETE_ER_MISSING_CURRENT_PASSWORD : {3892,[]string{"HY000"},"Current password needs to be specified in the REPLACE clause in order to change it."},
	ER_SERVER_WRONG_VALUE_FOR_VAR : {13227,[]string{"HY000"},"Variable '%-.64s' can't be set to the value of '%-.200s'"},
	ER_COULD_NOT_CREATE_WINDOWS_REGISTRY_KEY : {13228,[]string{"HY000"},"%s was unable to create a new Windows registry key %s for %s; continuing to use the previous ident."},
	ER_SERVER_GTID_UNSAFE_CREATE_DROP_TEMP_TABLE_IN_TRX_IN_SBR : {13229,[]string{"HY000"},"Statement violates GTID consistency: CREATE TEMPORARY TABLE and DROP TEMPORARY TABLE are not allowed inside a transaction or inside a procedure in a transactional context when @@session.binlog_format=STATEMENT."},
	//OBSOLETE_ER_SECONDARY_ENGINE : {3889,[]string{"HY000"},"Secondary engine operation failed. %s."},
	//OBSOLETE_ER_SECONDARY_ENGINE_DDL : {3890,[]string{"HY000"},"DDLs on a table with a secondary engine defined are not allowed."},
	//OBSOLETE_ER_NO_SESSION_TEMP : {3884,[]string{"HY000"},"Unable to allocate temporary tablespace for this session"},
	ER_XPLUGIN_FAILED_TO_SWITCH_SECURITY_CTX : {13233,[]string{"HY000"},"Unable to switch security context to user: %s"},
	ER_RPL_GTID_UNSAFE_ALTER_ADD_COL_WITH_DEFAULT_EXPRESSION : {13234,[]string{"HY000"},"Statement violates GTID consistency: ALTER TABLE ... ADD COLUMN .. with expression as DEFAULT."},
	ER_UPGRADE_PARSE_ERROR : {13235,[]string{"HY000"},"Error in parsing %s '%s'.'%s' during upgrade. %s"},
	ER_DATA_DIRECTORY_UNUSABLE : {13236,[]string{"HY000"},"The designated data directory %s is unusable. You can remove all files that the server added to it."},
	ER_LDAP_AUTH_USER_GROUP_SEARCH_ROOT_BIND : {13237,[]string{"HY000"},"Group search rebinding via root DN: %s "},
	ER_PLUGIN_INSTALL_ERROR : {13238,[]string{"HY000"},"Error installing plugin '%s': %s"},
	ER_PLUGIN_UNINSTALL_ERROR : {13239,[]string{"HY000"},"Error uninstalling plugin '%s': %s"},
	ER_SHARED_TABLESPACE_USED_BY_PARTITIONED_TABLE : {13240,[]string{"HY000"},"Partitioned table '%s' is not allowed to use shared tablespace '%s'. Please move all partitions to file-per-table tablespaces before upgrade."},
	ER_UNKNOWN_TABLESPACE_TYPE : {13241,[]string{"HY000"},"Cannot determine the type of the tablespace named '%s'."},
	ER_WARN_DEPRECATED_UTF8_ALIAS_OPTION : {13242,[]string{"HY000"},"%s: 'utf8' is currently an alias for the character set UTF8MB3, but will be an alias for UTF8MB4 in a future release. Please consider using UTF8MB4 in order to be unambiguous."},
	ER_WARN_DEPRECATED_UTF8MB3_CHARSET_OPTION : {13243,[]string{"HY000"},"%s: The character set UTF8MB3 is deprecated and will be removed in a future release. Please consider using UTF8MB4 instead."},
	ER_WARN_DEPRECATED_UTF8MB3_COLLATION_OPTION : {13244,[]string{"HY000"},"%s: '%-.64s' is a collation of the deprecated character set UTF8MB3. Please consider using UTF8MB4 with an appropriate collation instead."},
	ER_SSL_MEMORY_INSTRUMENTATION_INIT_FAILED : {13245,[]string{"HY000"},"The SSL library function %s failed. This is typically caused by the SSL library already being used. As a result the SSL memory allocation will not be instrumented."},
	ER_IB_MSG_MADV_DONTDUMP_UNSUPPORTED : {13246,[]string{"HY000"},"Disabling @@core_file because @@innodb_buffer_pool_in_core_file is disabled, yet MADV_DONTDUMP is not supported on this platform"},
	ER_IB_MSG_MADVISE_FAILED : {13247,[]string{"HY000"},"Disabling @@core_file because @@innodb_buffer_pool_in_core_file is disabled, yet madvise(%p,%zu,%s) failed with %s"},
	//OBSOLETE_ER_COLUMN_CHANGE_SIZE : {3886,[]string{"HY000"},"Could not change column '%s' of table '%s'. The resulting size of index '%s' would exceed the max key length of %d bytes."},
	ER_WARN_REMOVED_SQL_MODE : {13249,[]string{"HY000"},"sql_mode=0x%08x has been removed and will be ignored"},
	ER_IB_MSG_FAILED_TO_ALLOCATE_WAIT : {13250,[]string{"HY000"},"Failed to allocate memory for a pool of size %zu bytes. Will wait for %zu seconds for a thread to free a resource."},
	ER_IB_MSG_NUM_POOLS : {13251,[]string{"HY000"},"Number of pools: %zu"},
	ER_IB_MSG_USING_UNDO_SPACE : {13252,[]string{"HY000"},"Using undo tablespace '%s'."},
	ER_IB_MSG_FAIL_TO_SAVE_SPACE_STATE : {13253,[]string{"HY000"},"%s Unable to save the current state of tablespace '%s' to the data dictionary"},
	ER_IB_MSG_MAX_UNDO_SPACES_REACHED : {13254,[]string{"HY000"},"Cannot create undo tablespace %s at %s because %d undo tablespaces already exist."},
	ER_IB_MSG_ERROR_OPENING_NEW_UNDO_SPACE : {13255,[]string{"HY000"},"Error %d opening newly created undo tablespace %s."},
	ER_IB_MSG_FAILED_SDI_Z_BUF_ERROR : {13256,[]string{"HY000"},"SDI Compression failed, Z_BUF_ERROR"},
	ER_IB_MSG_FAILED_SDI_Z_MEM_ERROR : {13257,[]string{"HY000"},"SDI Compression failed, Z_MEM_ERROR"},
	ER_IB_MSG_SDI_Z_STREAM_ERROR : {13258,[]string{"HY000"},"SDI Compression failed, Z_STREAM_ERROR"},
	ER_IB_MSG_SDI_Z_UNKNOWN_ERROR : {13259,[]string{"HY000"},"%s"},
	ER_IB_MSG_FOUND_WRONG_UNDO_SPACE : {13260,[]string{"HY000"},"Expected to find undo tablespace '%s' for Space ID=%lu, but found '%s' instead!  Did you change innodb_undo_directory?"},
	ER_IB_MSG_NOT_END_WITH_IBU : {13261,[]string{"HY000"},"Cannot use %s as an undo tablespace because it does not end with '.ibu'."},
	//OBSOLETE_ER_IB_MSG_UNDO_TRUNCATE_EMPTY_FILE : {0000,[]string{""},"ib_undo_trunc_empty_file"},
	//OBSOLETE_ER_IB_MSG_UNDO_INJECT_BEFORE_DD_UPDATE : {0000,[]string{""},"ib_undo_trunc_before_dd_update"},
	//OBSOLETE_ER_IB_MSG_UNDO_INJECT_BEFORE_UNDO_LOGGING : {0000,[]string{""},"ib_undo_trunc_before_done_logging"},
	//OBSOLETE_ER_IB_MSG_UNDO_INJECT_BEFORE_RSEG : {0000,[]string{""},"ib_undo_trunc_before_rsegs"},
	ER_IB_MSG_FAILED_TO_FINISH_TRUNCATE : {13266,[]string{"HY000"},"%s Failed to finish truncating Undo Tablespace '%s'"},
	ER_IB_MSG_DEPRECATED_INNODB_UNDO_TABLESPACES : {13267,[]string{"HY000"},"The setting INNODB_UNDO_TABLESPACES is deprecated and is no longer used.  InnoDB always creates 2 undo tablespaces to start with. If you need more, please use CREATE UNDO TABLESPACE."},
	ER_IB_MSG_WRONG_TABLESPACE_DIR : {13268,[]string{"HY000"},"The directory for tablespace %s does not exist or is incorrect."},
	ER_IB_MSG_LOCK_FREE_HASH_USAGE_STATS : {13269,[]string{"HY000"},"%s"},
	ER_CLONE_DONOR_TRACE : {13270,[]string{"HY000"},"Clone donor reported : %.512s."},
	ER_CLONE_PROTOCOL_TRACE : {13271,[]string{"HY000"},"Clone received unexpected response from donor : %.512s."},
	ER_CLONE_CLIENT_TRACE : {13272,[]string{"HY000"},"Client: %.512s."},
	ER_CLONE_SERVER_TRACE : {13273,[]string{"HY000"},"Server: %.512s."},
	ER_THREAD_POOL_PFS_TABLES_INIT_FAILED : {13274,[]string{"HY000"},"Failed to initialize the performance schema tables service."},
	ER_THREAD_POOL_PFS_TABLES_ADD_FAILED : {13275,[]string{"HY000"},"Failed to add thread pool performance schema tables."},
	ER_CANT_SET_DATA_DIR : {13276,[]string{"HY000"},"Failed to set datadir to '%-.200s' (OS errno: %d - %s)"},
	ER_INNODB_INVALID_INNODB_UNDO_DIRECTORY_LOCATION : {13277,[]string{"HY000"},"The innodb_undo_directory is not allowed to be an ancestor of the datadir."},
	ER_SERVER_RPL_ENCRYPTION_FAILED_TO_FETCH_KEY : {13278,[]string{"HY000"},"Failed to fetch key from keyring, please check if keyring plugin is loaded."},
	ER_SERVER_RPL_ENCRYPTION_KEY_NOT_FOUND : {13279,[]string{"HY000"},"Can't find key from keyring, please check in the server log if a keyring plugin is loaded and initialized successfully."},
	ER_SERVER_RPL_ENCRYPTION_KEYRING_INVALID_KEY : {13280,[]string{"HY000"},"Fetched an invalid key from keyring."},
	ER_SERVER_RPL_ENCRYPTION_HEADER_ERROR : {13281,[]string{"HY000"},"Error reading a replication log encryption header: %s."},
	ER_SERVER_RPL_ENCRYPTION_FAILED_TO_ROTATE_LOGS : {13282,[]string{"HY000"},"Failed to rotate some logs after changing binlog encryption settings. Please fix the problem and rotate the logs manually."},
	ER_SERVER_RPL_ENCRYPTION_KEY_EXISTS_UNEXPECTED : {13283,[]string{"HY000"},"Key %s exists unexpected."},
	ER_SERVER_RPL_ENCRYPTION_FAILED_TO_GENERATE_KEY : {13284,[]string{"HY000"},"Failed to generate key, please check if keyring plugin is loaded."},
	ER_SERVER_RPL_ENCRYPTION_FAILED_TO_STORE_KEY : {13285,[]string{"HY000"},"Failed to store key, please check if keyring plugin is loaded."},
	ER_SERVER_RPL_ENCRYPTION_FAILED_TO_REMOVE_KEY : {13286,[]string{"HY000"},"Failed to remove key, please check if keyring plugin is loaded."},
	ER_SERVER_RPL_ENCRYPTION_MASTER_KEY_RECOVERY_FAILED : {13287,[]string{"HY000"},"Unable to recover binlog encryption master key, please check if keyring plugin is loaded."},
	ER_SERVER_RPL_ENCRYPTION_UNABLE_TO_INITIALIZE : {13288,[]string{"HY000"},"Failed to initialize binlog encryption, please check if keyring plugin is loaded."},
	ER_SERVER_RPL_ENCRYPTION_UNABLE_TO_ROTATE_MASTER_KEY_AT_STARTUP : {13289,[]string{"HY000"},"Failed to rotate binlog encryption master key at startup, please check if keyring plugin is loaded."},
	ER_SERVER_RPL_ENCRYPTION_IGNORE_ROTATE_MASTER_KEY_AT_STARTUP : {13290,[]string{"HY000"},"Ignoring binlog_rotate_encryption_master_key_at_startup because binlog_encryption option is disabled."},
	ER_INVALID_ADMIN_ADDRESS : {13291,[]string{"HY000"},"Invalid value for command line option admin-address: '%s'"},
	ER_SERVER_STARTUP_ADMIN_INTERFACE : {13292,[]string{"HY000"},"Admin interface ready for connections, address: '%s'  port: %d"},
	ER_CANT_CREATE_ADMIN_THREAD : {13293,[]string{"HY000"},"Can't create thread to handle admin connections (errno= %d)"},
	ER_WARNING_RETAIN_CURRENT_PASSWORD_CLAUSE_VOID : {13294,[]string{"HY000"},"RETAIN CURRENT PASSWORD ignored for user '%s'@'%s' as its authentication plugin %s does not support multiple passwords."},
	ER_WARNING_DISCARD_OLD_PASSWORD_CLAUSE_VOID : {13295,[]string{"HY000"},"DISCARD OLD PASSWORD ignored for user '%s'@'%s' as its authentication plugin %s does not support multiple passwords."},
	//OBSOLETE_ER_SECOND_PASSWORD_CANNOT_BE_EMPTY : {3878,[]string{"HY000"},"Empty password can not be retained as second password for user '%s'@'%s'."},
	//OBSOLETE_ER_PASSWORD_CANNOT_BE_RETAINED_ON_PLUGIN_CHANGE : {3894,[]string{"HY000"},"Current password can not be retained for user '%s'@'%s' because authentication plugin is being changed."},
	//OBSOLETE_ER_CURRENT_PASSWORD_CANNOT_BE_RETAINED : {3895,[]string{"HY000"},"Current password can not be retained for user '%s'@'%s' because new password is empty."},
	ER_WARNING_AUTHCACHE_INVALID_USER_ATTRIBUTES : {13299,[]string{"HY000"},"Can not read and process value of User_attributes column from mysql.user table for user: '%s@%s'; Ignoring user."},
	ER_MYSQL_NATIVE_PASSWORD_SECOND_PASSWORD_USED_INFORMATION : {13300,[]string{"HY000"},"Second password was used for login by user: '%s'@'%s'."},
	ER_SHA256_PASSWORD_SECOND_PASSWORD_USED_INFORMATION : {13301,[]string{"HY000"},"Second password was used for login by user: '%s'@'%s'."},
	ER_CACHING_SHA2_PASSWORD_SECOND_PASSWORD_USED_INFORMATION : {13302,[]string{"HY000"},"Second password was used for login by user: '%s'@'%s'."},
	ER_GRP_RPL_SEND_TRX_PREPARED_MESSAGE_FAILED : {13303,[]string{"HY000"},"Error sending transaction '%d:%lld' prepared message from session '%u'."},
	ER_GRP_RPL_RELEASE_COMMIT_AFTER_GROUP_PREPARE_FAILED : {13304,[]string{"HY000"},"Error releasing transaction '%d:%lld' for commit on session '%u' after being prepared on all group members."},
	ER_GRP_RPL_TRX_ALREADY_EXISTS_ON_TCM_ON_AFTER_CERTIFICATION : {13305,[]string{"HY000"},"Transaction '%d:%lld' already exists on Group Replication consistency manager while being registered after conflict detection."},
	ER_GRP_RPL_FAILED_TO_INSERT_TRX_ON_TCM_ON_AFTER_CERTIFICATION : {13306,[]string{"HY000"},"Error registering transaction '%d:%lld' on Group Replication consistency manager after conflict detection."},
	ER_GRP_RPL_REGISTER_TRX_TO_WAIT_FOR_GROUP_PREPARE_FAILED : {13307,[]string{"HY000"},"Error registering transaction '%d:%lld' from session '%u' to wait for being prepared on all group members."},
	ER_GRP_RPL_TRX_WAIT_FOR_GROUP_PREPARE_FAILED : {13308,[]string{"HY000"},"Error on transaction '%d:%lld' from session '%u' while waiting for being prepared on all group members."},
	ER_GRP_RPL_TRX_DOES_NOT_EXIST_ON_TCM_ON_HANDLE_REMOTE_PREPARE : {13309,[]string{"HY000"},"Transaction '%d:%lld' does not exist on Group Replication consistency manager while receiving remote transaction prepare."},
	ER_GRP_RPL_RELEASE_BEGIN_TRX_AFTER_DEPENDENCIES_COMMIT_FAILED : {13310,[]string{"HY000"},"Error releasing transaction '%d:%lld' for execution on session '%u' after its dependencies did complete commit."},
	ER_GRP_RPL_REGISTER_TRX_TO_WAIT_FOR_DEPENDENCIES_FAILED : {13311,[]string{"HY000"},"Error registering transaction from session '%u' to wait for its dependencies to complete commit."},
	ER_GRP_RPL_WAIT_FOR_DEPENDENCIES_FAILED : {13312,[]string{"HY000"},"Error on session '%u' while waiting for its dependencies to complete commit."},
	ER_GRP_RPL_REGISTER_TRX_TO_WAIT_FOR_SYNC_BEFORE_EXECUTION_FAILED : {13313,[]string{"HY000"},"Error registering transaction from session '%u' to wait for sync before execution."},
	ER_GRP_RPL_SEND_TRX_SYNC_BEFORE_EXECUTION_FAILED : {13314,[]string{"HY000"},"Error sending sync before execution message from session '%u'."},
	ER_GRP_RPL_TRX_WAIT_FOR_SYNC_BEFORE_EXECUTION_FAILED : {13315,[]string{"HY000"},"Error on transaction from session '%u' while waiting for sync before execution."},
	ER_GRP_RPL_RELEASE_BEGIN_TRX_AFTER_WAIT_FOR_SYNC_BEFORE_EXEC : {13316,[]string{"HY000"},"Error releasing transaction for execution on session '%u' after wait for sync before execution."},
	ER_GRP_RPL_TRX_WAIT_FOR_GROUP_GTID_EXECUTED : {13317,[]string{"HY000"},"Error waiting for group executed transactions commit on session '%u'."},
	//OBSOLETE_ER_UNIT_NOT_FOUND : {3902,[]string{"SU001"},"There's no unit of measure named '%s'."},
	//OBSOLETE_ER_GEOMETRY_IN_UNKNOWN_LENGTH_UNIT : {3882,[]string{"SU001"},"The function %s uses %s as a unit, but was passed geometry without units (\"SRID 0\"). Conversion is not possible."},
	ER_WARN_PROPERTY_STRING_PARSE_FAILED : {13320,[]string{"HY000"},"Could not parse key-value pairs in property string '%s'"},
	ER_INVALID_PROPERTY_KEY : {13321,[]string{"HY000"},"Property key '%s' is invalid."},
	ER_GRP_RPL_GTID_SET_EXTRACT_ERROR_DURING_RECOVERY : {13322,[]string{"HY000"},"Error when extracting the group_replication_applier channel received transactions set. Unable to ensure the execution of group transactions received during recovery."},
	ER_SERVER_RPL_ENCRYPTION_FAILED_TO_ENCRYPT : {13323,[]string{"HY000"},"Failed to encrypt content to write into binlog file: %s."},
	ER_CANNOT_GET_SERVER_VERSION_FROM_TABLESPACE_HEADER : {13324,[]string{"HY000"},"Cannot get the server version number from the dictionary tablespace header."},
	ER_CANNOT_SET_SERVER_VERSION_IN_TABLESPACE_HEADER : {13325,[]string{"HY000"},"Cannot set the server version number in the dictionary tablespace header."},
	ER_SERVER_UPGRADE_VERSION_NOT_SUPPORTED : {13326,[]string{"HY000"},"Upgrading the server from server version '%u' is not supported."},
	ER_SERVER_UPGRADE_FROM_VERSION : {13327,[]string{"HY000"},"MySQL server upgrading from version '%u' to '%u'."},
	ER_GRP_RPL_ERROR_ON_CERT_DB_INSTALL : {13328,[]string{"HY000"},"The certification information could not be set in this server: '%s'"},
	ER_GRP_RPL_FORCE_MEMBERS_WHEN_LEAVING : {13329,[]string{"HY000"},"A request to force a new group membership was issued when the member is leaving the group."},
	ER_TRG_WRONG_ORDER : {13330,[]string{"HY000"},"Trigger %s.%s for table %s.%s is listed in wrong order. Please drop and recreate all triggers for the table."},
	//OBSOLETE_ER_SECONDARY_ENGINE_PLUGIN : {3877,[]string{"HY000"},"%s"},
	ER_LDAP_AUTH_GRP_SEARCH_NOT_SPECIAL_HDL : {13332,[]string{"HY000"},"Special handling for group search, {GA} not found"},
	ER_LDAP_AUTH_GRP_USER_OBJECT_HAS_GROUP_INFO : {13333,[]string{"HY000"},"User group retrieval: User object has group information"},
	ER_LDAP_AUTH_GRP_INFO_FOUND_IN_MANY_OBJECTS : {13334,[]string{"HY000"},"Group information found in multiple user objects. Search filter configuration is incorrect."},
	ER_LDAP_AUTH_GRP_INCORRECT_ATTRIBUTE : {13335,[]string{"HY000"},"User group retrieval: no group attribute found. Incorrect group search attribute. "},
	ER_LDAP_AUTH_GRP_NULL_ATTRIBUTE_VALUE : {13336,[]string{"HY000"},"User group retrieval: Group attribute values is NULL. "},
	ER_LDAP_AUTH_GRP_DN_PARSING_FAILED : {13337,[]string{"HY000"},"User group retrieval: parsing DN failed. "},
	ER_LDAP_AUTH_GRP_OBJECT_HAS_USER_INFO : {13338,[]string{"HY000"},"User group retrieval: Group object has user information"},
	ER_LDAP_AUTH_LDAPS : {13339,[]string{"HY000"},"Reserved port for ldaps using ldaps"},
	ER_LDAP_MAPPING_GET_USER_PROXY : {13340,[]string{"HY000"},"Get user proxy"},
	ER_LDAP_MAPPING_USER_DONT_BELONG_GROUP : {13341,[]string{"HY000"},"Get user proxy: User doesn't belongs to any group, user name will be treated as authenticated user."},
	ER_LDAP_MAPPING_INFO : {13342,[]string{"HY000"},"Get user proxy: configured mapping info: %s"},
	ER_LDAP_MAPPING_EMPTY_MAPPING : {13343,[]string{"HY000"},"Get user proxy: User doesn't have group mapping information, First LDAP group will be treated as authenticated user."},
	ER_LDAP_MAPPING_PROCESS_MAPPING : {13344,[]string{"HY000"},"Process group proxy mapping"},
	ER_LDAP_MAPPING_CHECK_DELIMI_QUOTE : {13345,[]string{"HY000"},"Check delimiter after quote"},
	ER_LDAP_MAPPING_PROCESS_DELIMITER : {13346,[]string{"HY000"},"Processing delimiter"},
	ER_LDAP_MAPPING_PROCESS_DELIMITER_EQUAL_NOT_FOUND : {13347,[]string{"HY000"},"Processing delimiter, separator = not found, resetting position"},
	ER_LDAP_MAPPING_PROCESS_DELIMITER_TRY_COMMA : {13348,[]string{"HY000"},"Processing delimiter, failed to get data for = separator try for separator ,."},
	ER_LDAP_MAPPING_PROCESS_DELIMITER_COMMA_NOT_FOUND : {13349,[]string{"HY000"},"Processing delimiter, separator , not found, resetting position"},
	ER_LDAP_MAPPING_NO_SEPEARATOR_END_OF_GROUP : {13350,[]string{"HY000"},"Processing delimiter: No mapping separator is found, end of group information"},
	ER_LDAP_MAPPING_GETTING_NEXT_MAPPING : {13351,[]string{"HY000"},"Getting next mapping information"},
	ER_LDAP_MAPPING_PARSING_CURRENT_STATE : {13352,[]string{"HY000"},"Parsing mapping, current state: %d  delimiter char: %c "},
	ER_LDAP_MAPPING_PARSING_MAPPING_INFO : {13353,[]string{"HY000"},"Parsing mapping info, LDAP group: %s MySQL proxy: %s"},
	ER_LDAP_MAPPING_PARSING_ERROR : {13354,[]string{"HY000"},"Mapping parsing error"},
	ER_LDAP_MAPPING_TRIMMING_SPACES : {13355,[]string{"HY000"},"Trimming left spaces"},
	ER_LDAP_MAPPING_IS_QUOTE : {13356,[]string{"HY000"},"Checking if current characters is quote"},
	ER_LDAP_MAPPING_NON_DESIRED_STATE : {13357,[]string{"HY000"},"Not desired state or un-defined states."},
	ER_INVALID_NAMED_PIPE_FULL_ACCESS_GROUP : {13358,[]string{"HY000"},"Invalid value for named_pipe_full_access_group."},
	ER_PREPARE_FOR_SECONDARY_ENGINE : {13359,[]string{"HY000"},"Retry the statement using a secondary storage engine."},
	ER_SERVER_WARN_DEPRECATED : {13360,[]string{"HY000"},"'%s' is deprecated and will be removed in a future release. Please use %s instead"},
	ER_AUTH_ID_WITH_SYSTEM_USER_PRIV_IN_MANDATORY_ROLES : {13361,[]string{"HY000"},"Cannot set mandatory_roles: AuthId `%.64s`@`%.64s` has '%s' privilege."},
	ER_SERVER_BINLOG_MASTER_KEY_RECOVERY_OUT_OF_COMBINATION : {13362,[]string{"HY000"},"Unable to recover binary log master key, the combination of new_master_key_seqno=%u, master_key_seqno=%u and old_master_key_seqno=%u are wrong."},
	ER_SERVER_BINLOG_MASTER_KEY_ROTATION_FAIL_TO_CLEANUP_AUX_KEY : {13363,[]string{"HY000"},"Failed to remove auxiliary binary log encryption key from keyring, please check if keyring plugin is loaded. The cleanup of the binary log master key rotation process did not finish as expected and the cleanup will take place upon server restart or next 'ALTER INSTANCE ROTATE BINLOG MASTER KEY' execution."},
	//OBSOLETE_ER_CANNOT_GRANT_SYSTEM_PRIV_TO_MANDATORY_ROLE : {3897,[]string{"HY000"},"AuthId `%.64s`@`%.64s` is set as mandatory_roles. Cannot grant the '%s' privilege."},
	//OBSOLETE_ER_PARTIAL_REVOKE_AND_DB_GRANT_BOTH_EXISTS : {3901,[]string{"HY000"},"'%s' privilege for database '%s' exists both as partial revoke and mysql.db simultaneously. It could mean 'mysql' schema is corrupted."},
	//OBSOLETE_ER_DB_ACCESS_DENIED : {3879,[]string{"HY000"},"Access denied for AuthId `%.64s`@`%.64s` to database '%-.192s'."},
	//OBSOLETE_ER_PARTIAL_REVOKES_EXIST : {3896,[]string{"HY000"},"At least one partial revoke exists on a database. The system variable '@@partial_revokes' must be set to ON."},
	ER_TURNING_ON_PARTIAL_REVOKES : {13368,[]string{"HY000"},"At least one partial revoke exists on a database. Turning ON the system variable '@@partial_revokes'."},
	ER_WARN_PARTIAL_REVOKE_AND_DB_GRANT : {13369,[]string{"HY000"},"For user '%s'@'%s', one or more privileges granted through mysql.db for database '%s', conflict with partial revoke. It could mean 'mysql' schema is corrupted."},
	ER_WARN_INCORRECT_PRIVILEGE_FOR_DB_RESTRICTIONS : {13370,[]string{"HY000"},"For user %s, ignored restrictions for privilege(s) '%s' for database '%s' as these are not valid database privileges."},
	ER_WARN_INVALID_DB_RESTRICTIONS : {13371,[]string{"HY000"},"For user %s, ignored restrictions for privilege(s) '%s' for database '%s' as corresponding global privilege(s) are not granted."},
	ER_GRP_RPL_INVALID_COMMUNICATION_PROTOCOL : {13372,[]string{"HY000"},"'%s' is an invalid value for group_replication_communication_protocol_join, please use a MySQL version between 5.7.14 and this server's version"},
	ER_GRP_RPL_STARTED_AUTO_REJOIN : {13373,[]string{"HY000"},"Started auto-rejoin procedure attempt %lu of %lu"},
	ER_GRP_RPL_TIMEOUT_RECEIVED_VC_ON_REJOIN : {13374,[]string{"HY000"},"Timeout while waiting for a view change event during the auto-rejoin procedure"},
	ER_GRP_RPL_FINISHED_AUTO_REJOIN : {13375,[]string{"HY000"},"Auto-rejoin procedure attempt %lu of %lu finished. Member was%s able to join the group."},
	ER_GRP_RPL_DEFAULT_TABLE_ENCRYPTION_DIFF_FROM_GRP : {13376,[]string{"HY000"},"The member is configured with a default_table_encryption option value '%d' different from the group '%d'. The member will now exit the group."},
	ER_SERVER_UPGRADE_OFF : {13377,[]string{"HY000"},"Server shutting down because upgrade is required, yet prohibited by the command line option '--upgrade=NONE'."},
	ER_SERVER_UPGRADE_SKIP : {13378,[]string{"HY000"},"Server upgrade is required, but skipped by command line option '--upgrade=MINIMAL'."},
	ER_SERVER_UPGRADE_PENDING : {13379,[]string{"HY000"},"Server upgrade started with version %d, but server upgrade of version %d is still pending."},
	ER_SERVER_UPGRADE_FAILED : {13380,[]string{"HY000"},"Failed to upgrade server."},
	ER_SERVER_UPGRADE_STATUS : {13381,[]string{"HY000"},"Server upgrade from '%d' to '%d' %s."},
	ER_SERVER_UPGRADE_REPAIR_REQUIRED : {13382,[]string{"HY000"},"Table '%s' requires repair."},
	ER_SERVER_UPGRADE_REPAIR_STATUS : {13383,[]string{"HY000"},"Table '%s' repair %s."},
	ER_SERVER_UPGRADE_INFO_FILE : {13384,[]string{"HY000"},"Could not open server upgrade info file '%s' for writing. Please make sure the file is writable."},
	ER_SERVER_UPGRADE_SYS_SCHEMA : {13385,[]string{"HY000"},"Upgrading the sys schema."},
	ER_SERVER_UPGRADE_MYSQL_TABLES : {13386,[]string{"HY000"},"Running queries to upgrade MySQL server."},
	ER_SERVER_UPGRADE_SYSTEM_TABLES : {13387,[]string{"HY000"},"Upgrading system table data."},
	ER_SERVER_UPGRADE_EMPTY_SYS : {13388,[]string{"HY000"},"Found empty sys database. Installing the sys schema."},
	ER_SERVER_UPGRADE_NO_SYS_VERSION : {13389,[]string{"HY000"},"A sys schema exists with no sys.version view. If you have a user created sys schema, this must be renamed for the upgrade to succeed."},
	ER_SERVER_UPGRADE_SYS_VERSION_EMPTY : {13390,[]string{"HY000"},"A sys schema exists with a sys.version view, but it returns no results."},
	ER_SERVER_UPGRADE_SYS_SCHEMA_OUTDATED : {13391,[]string{"HY000"},"Found outdated sys schema version %s."},
	ER_SERVER_UPGRADE_SYS_SCHEMA_UP_TO_DATE : {13392,[]string{"HY000"},"The sys schema is already up to date (version %s)."},
	ER_SERVER_UPGRADE_SYS_SCHEMA_OBJECT_COUNT : {13393,[]string{"HY000"},"Found %d sys %s, but expected %d. Re-installing the sys schema."},
	ER_SERVER_UPGRADE_CHECKING_DB : {13394,[]string{"HY000"},"Checking '%s' schema."},
	ER_IB_MSG_DDL_LOG_DELETE_BY_ID_TMCT : {13395,[]string{"HY000"},"Too many concurrent transactions while clearing the DDL Log. Please increase the number of Rollback Segments."},
	ER_IB_MSG_POST_RECOVER_DDL_LOG_RECOVER : {13396,[]string{"HY000"},"Error in DDL Log recovery during Post-Recovery processing."},
	ER_IB_MSG_POST_RECOVER_POST_TS_ENCRYPT : {13397,[]string{"HY000"},"Error in Post-Tablespace-Encryption during Post-Recovery processing."},
	ER_IB_MSG_DDL_LOG_FAIL_POST_DDL : {13398,[]string{"HY000"},"Error in DLL Log cleanup during Post-DDL processing."},
	ER_SERVER_BINLOG_UNSAFE_SYSTEM_FUNCTION : {13399,[]string{"HY000"},"'%s' statement is unsafe because it uses a system function that may return a different value on the slave."},
	ER_SERVER_UPGRADE_HELP_TABLE_STATUS : {13400,[]string{"HY000"},"Upgrade of help tables %s."},
	ER_GRP_RPL_SRV_GTID_WAIT_ERROR : {13401,[]string{"HY000"},"Error when waiting for the server to execute local transactions in order assure the group change proper logging"},
	ER_GRP_DELAYED_VCLE_LOGGING : {13402,[]string{"HY000"},"Unable to log the group change View log event in its exaction position in the log. This will not however affect the group replication recovery process or the overall plugin process."},
	//OBSOLETE_ER_CANNOT_GRANT_ROLES_TO_ANONYMOUS_USER : {3876,[]string{"HY000"},"Cannot grant roles to an anonymous user."},
	ER_BINLOG_UNABLE_TO_ROTATE_GTID_TABLE_READONLY : {13404,[]string{"HY000"},"Unable to create a new binlog file: Table `mysql.gtid_executed` couldn't be opened. %s"},
	ER_NETWORK_NAMESPACES_NOT_SUPPORTED : {13405,[]string{"HY000"},"Network Namespaces is not supported on this platform"},
	ER_UNKNOWN_NETWORK_NAMESPACE : {13406,[]string{"HY000"},"Unknown network namespace '%s'"},
	ER_NETWORK_NAMESPACE_NOT_ALLOWED_FOR_WILDCARD_ADDRESS : {13407,[]string{"HY000"},"Network namespace not allowed for wildcard interface address"},
	ER_SETNS_FAILED : {13408,[]string{"HY000"},"setns() failed with error '%s'"},
	ER_WILDCARD_NOT_ALLOWED_FOR_MULTIADDRESS_BIND : {13409,[]string{"HY000"},"Wildcard address value not allowed for multivalued bind address"},
	ER_NETWORK_NAMESPACE_FILE_PATH_TOO_LONG : {13410,[]string{"HY000"},"The path to a special network namespace file is too long. (got %u > max %u)"},
	ER_IB_MSG_TOO_LONG_PATH : {13411,[]string{"HY000"},"Cannot create tablespace '%s'. The filepath is too long for this OS."},
	ER_IB_RECV_FIRST_REC_GROUP_INVALID : {13412,[]string{"HY000"},"The last block of redo had corrupted first_rec_group and became fixed (%u -> %u)."},
	ER_DD_UPGRADE_COMPLETED : {13413,[]string{"HY000"},"Data dictionary upgrade from version '%u' to '%u' completed."},
	ER_SSL_SERVER_CERT_VERIFY_FAILED : {13414,[]string{"HY000"},"Server SSL certificate doesn't verify: %s"},
	ER_PERSIST_OPTION_USER_TRUNCATED : {13415,[]string{"HY000"},"Truncated a user name for %s that was too long while reading the persisted variables file"},
	ER_PERSIST_OPTION_HOST_TRUNCATED : {13416,[]string{"HY000"},"Truncated a host name for %s that was too long while reading the persisted variables file"},
	ER_NET_WAIT_ERROR : {13417,[]string{"HY000"},"The wait_timeout period was exceeded, the idle time since last command was too long."},
	ER_IB_MSG_1285 : {13418,[]string{"HY000"},"'%s' found not encrypted while '%s' is ON. Trying to encrypt it now."},
	ER_IB_MSG_CLOCK_MONOTONIC_UNSUPPORTED : {13419,[]string{"HY000"},"CLOCK_MONOTONIC is unsupported, so do not change the system time when MySQL is running !"},
	ER_IB_MSG_CLOCK_GETTIME_FAILED : {13420,[]string{"HY000"},"clock_gettime() failed: %s"},
	ER_PLUGIN_NOT_EARLY_DUP : {13421,[]string{"HY000"},"Plugin '%s' is not to be used as an \"early\" plugin. Don't add it to --early-plugin-load, keyring migration etc."},
	ER_PLUGIN_NO_INSTALL_DUP : {13422,[]string{"HY000"},"Plugin '%s' is marked as not dynamically installable. You have to stop the server to install it."},
	//OBSOLETE_ER_WARN_DEPRECATED_SQL_CALC_FOUND_ROWS : {3964,[]string{"HY000"},"SQL_CALC_FOUND_ROWS is deprecated and will be removed in a future release. Consider using two separate queries instead."},
	//OBSOLETE_ER_WARN_DEPRECATED_FOUND_ROWS : {3965,[]string{"HY000"},"FOUND_ROWS() is deprecated and will be removed in a future release. Consider using COUNT(*) instead."},
	ER_BINLOG_UNSAFE_DEFAULT_EXPRESSION_IN_SUBSTATEMENT : {13425,[]string{"HY000"},"The statement is unsafe because it invokes a trigger or a stored function that modifies a table that has a column with a DEFAULT expression that may return a different value on the slave."},
	ER_GRP_RPL_MEMBER_VER_READ_COMPATIBLE : {13426,[]string{"HY000"},"Member version is read compatible with the group."},
	ER_LOCK_ORDER_INIT_FAILED : {13427,[]string{"HY000"},"Lock order disabled (reason: init failed)."},
	ER_AUDIT_LOG_KEYRING_ID_TIMESTAMP_VALUE_IS_INVALID : {13428,[]string{"HY000"},"Keyring ID timestamp value is invalid: '%s'"},
	ER_AUDIT_LOG_FILE_NAME_TIMESTAMP_VALUE_IS_MISSING_OR_INVALID : {13429,[]string{"HY000"},"Cannot process audit log file. File name timestamp value is missing or invalid: '%s'"},
	ER_AUDIT_LOG_FILE_NAME_DOES_NOT_HAVE_REQUIRED_FORMAT : {13430,[]string{"HY000"},"Cannot process audit log file. File name does not have required format: '%s'"},
	ER_AUDIT_LOG_FILE_NAME_KEYRING_ID_VALUE_IS_MISSING : {13431,[]string{"HY000"},"Cannot process audit log file. File name keyring ID value is missing: '%s'"},
	ER_AUDIT_LOG_FILE_HAS_BEEN_SUCCESSFULLY_PROCESSED : {13432,[]string{"HY000"},"Audit log file has been successfully processed: '%s'"},
	ER_AUDIT_LOG_COULD_NOT_OPEN_FILE_FOR_READING : {13433,[]string{"HY000"},"Could not open audit log file for reading: '%s'"},
	ER_AUDIT_LOG_INVALID_FILE_CONTENT : {13434,[]string{"HY000"},"Invalid audit log file content: '%s'"},
	ER_AUDIT_LOG_CANNOT_READ_PASSWORD : {13435,[]string{"HY000"},"Cannot read password: '%.32s'"},
	ER_AUDIT_LOG_CANNOT_STORE_PASSWORD : {13436,[]string{"HY000"},"Cannot store password: '%.32s'"},
	ER_AUDIT_LOG_CANNOT_REMOVE_PASSWORD : {13437,[]string{"HY000"},"Cannot remove password: '%.32s'"},
	ER_AUDIT_LOG_PASSWORD_HAS_BEEN_COPIED : {13438,[]string{"HY000"},"'audit_log' password has been copied into '%.32s' and will be removed with first purged password."},
	//OBSOLETE_ER_AUDIT_LOG_INSUFFICIENT_PRIVILEGE : {3914,[]string{"HY000"},"Request ignored for '%.64s'@'%.64s'. Role needed to perform operation: '%.32s'"},
	//OBSOLETE_ER_WRONG_MVI_VALUE : {3908,[]string{"HY000"},"Can't store an array or an object in a scalar key part of the index '%.192s'"},
	//OBSOLETE_ER_WARN_FUNC_INDEX_NOT_APPLICABLE : {3909,[]string{"HY000"},"Cannot use functional index '%-.64s' due to type or collation conversion"},
	//OBSOLETE_ER_EXCEEDED_MV_KEYS_NUM : {3905,[]string{"HY000"},"Exceeded max number of values per record for multi-valued index '%-.64s' by %u value(s)"},
	//OBSOLETE_ER_EXCEEDED_MV_KEYS_SPACE : {3906,[]string{"HY000"},"Exceeded max total length of values per record for multi-valued index '%-.64s' by %u bytes"},
	//OBSOLETE_ER_FUNCTIONAL_INDEX_DATA_IS_TOO_LONG : {3907,[]string{"22001"},"Data too long for functional index '%-.64s'"},
	//OBSOLETE_ER_INVALID_JSON_VALUE_FOR_FUNC_INDEX : {3903,[]string{"22018"},"Invalid JSON value for CAST for functional index '%-.64s'"},
	//OBSOLETE_ER_JSON_VALUE_OUT_OF_RANGE_FOR_FUNC_INDEX : {3904,[]string{"22003"},"Out of range JSON value for CAST for functional index '%-.64s'"},
	ER_LDAP_EMPTY_USERDN_PASSWORD : {13447,[]string{"HY000"},"Empty user dn or password is not allowed, not attempting LDAP bind."},
	ER_ACL_WRONG_OR_MISSING_ACL_TABLES_LOG : {13449,[]string{"HY000"},"The current layout of the ACL tables does not conform to the server's expected layout. They're either altered, missing or not upgraded from a previous version. However a best effort attempt to read data from these tables will still be made."},
	ER_LOCK_ORDER_FAILED_WRITE_FILE : {13450,[]string{"HY000"},"LOCK_ORDER: Failed to write to file <%s>."},
	ER_LOCK_ORDER_FAILED_READ_FILE : {13451,[]string{"HY000"},"LOCK_ORDER: Failed to read from file <%s>."},
	ER_LOCK_ORDER_MESSAGE : {13452,[]string{"HY000"},"LOCK_ORDER message: %s"},
	ER_LOCK_ORDER_DEPENDENCIES_SYNTAX : {13453,[]string{"HY000"},"Lock order dependencies file <%s> (%d:%d) - (%d:%d) : %s"},
	ER_LOCK_ORDER_SCANNER_SYNTAX : {13454,[]string{"HY000"},"Lock order scanner: (%d:%d) - (%d:%d) : %s"},
	ER_DATA_DIRECTORY_UNUSABLE_DELETABLE : {13455,[]string{"HY000"},"The newly created data directory %s by --initialize is unusable. You can remove it."},
	ER_IB_MSG_BTREE_LEVEL_LIMIT_EXCEEDED : {13456,[]string{"HY000"},"No. of B-tree level created for index %s has crossed the permissible limit. If debug option innodb_limit_optimistic_insert_debug is being used try tweaking it to include more records in a page."},
	ER_IB_CLONE_START_STOP : {13457,[]string{"HY000"},"%s"},
	ER_IB_CLONE_OPERATION : {13458,[]string{"HY000"},"%s"},
	ER_IB_CLONE_RESTART : {13459,[]string{"HY000"},"%s"},
	ER_IB_CLONE_USER_DATA : {13460,[]string{"HY000"},"Clone removing all user data for provisioning: %s"},
	ER_IB_CLONE_NON_INNODB_TABLE : {13461,[]string{"HY000"},"Non innodb table: %s.%s is not cloned and is empty."},
	ER_CLONE_SHUTDOWN_TRACE : {13462,[]string{"HY000"},"Clone shutting down server as RESTART failed. Please start server to complete clone operation."},
	ER_GRP_RPL_GTID_PURGED_EXTRACT_ERROR : {13463,[]string{"HY000"},"Error when extracting this member GTID purged set. Operations and checks made to group joiners may be incomplete."},
	ER_GRP_RPL_CLONE_PROCESS_PREPARE_ERROR : {13464,[]string{"HY000"},"There was an issue when configuring the remote cloning process: %s"},
	ER_GRP_RPL_CLONE_PROCESS_EXEC_ERROR : {13465,[]string{"HY000"},"There was an issue when cloning from another server: %s"},
	ER_GRP_RPL_RECOVERY_EVAL_ERROR : {13466,[]string{"HY000"},"There was an issue when trying to evaluate the best distributed recovery strategy while joining.%s"},
	ER_GRP_RPL_NO_POSSIBLE_RECOVERY : {13467,[]string{"HY000"},"No valid or ONLINE members exist to get the missing data from the group. For cloning check if donors of the same version and with clone plugin installed exist. For incremental recovery check if you have donors where the required data was not purged from the binary logs."},
	ER_GRP_RPL_CANT_KILL_THREAD : {13468,[]string{"HY000"},"The group replication plugin could not kill the plugin routine for %s. %s"},
	ER_GRP_RPL_RECOVERY_STRAT_CLONE_THRESHOLD : {13469,[]string{"HY000"},"This member will start distributed recovery using clone. It is due to the number of missing transactions being higher than the configured threshold of %llu."},
	ER_GRP_RPL_RECOVERY_STRAT_CLONE_PURGED : {13470,[]string{"HY000"},"This member will start distributed recovery using clone. It is due to no ONLINE member has the missing data for recovering in its binary logs."},
	ER_GRP_RPL_RECOVERY_STRAT_CHOICE : {13471,[]string{"HY000"},"Distributed recovery will transfer data using: %s"},
	ER_GRP_RPL_RECOVERY_STRAT_FALLBACK : {13472,[]string{"HY000"},"Due to some issue on the previous step distributed recovery is now executing: %s"},
	ER_GRP_RPL_RECOVERY_STRAT_NO_FALLBACK : {13473,[]string{"HY000"},"Due to a critical cloning error or lack of donors, distributed recovery cannot be executed. The member will now leave the group."},
	ER_GRP_RPL_SLAVE_THREAD_ERROR_ON_CLONE : {13474,[]string{"HY000"},"The '%s' thread of channel '%s' will error out as the server will attempt to clone another server"},
	ER_UNKNOWN_TABLE_IN_UPGRADE : {13475,[]string{"HY000"},"Unknown table '%-.129s'"},
	ER_IDENT_CAUSES_TOO_LONG_PATH_IN_UPGRADE : {13476,[]string{"HY000"},"Long database name and identifier for object resulted in path length exceeding %d characters. Path: '%s'."},
	ER_XA_CANT_CREATE_MDL_BACKUP : {13477,[]string{"HY000"},"XA: Failed to take MDL Lock backup of PREPARED XA transaction during client disconnect."},
	ER_AUDIT_LOG_SUPER_PRIVILEGE_REQUIRED : {13478,[]string{"HY000"},"SUPER privilege or AUDIT_ADMIN role required for '%s'@'%s' user."},
	ER_AUDIT_LOG_UDF_INVALID_ARGUMENT_TYPE : {13479,[]string{"HY000"},"Invalid argument type"},
	ER_AUDIT_LOG_UDF_INVALID_ARGUMENT_COUNT : {13480,[]string{"HY000"},"Invalid argument count"},
	ER_AUDIT_LOG_HAS_NOT_BEEN_INSTALLED : {13481,[]string{"HY000"},"audit_log plugin has not been installed using INSTALL PLUGIN syntax."},
	ER_AUDIT_LOG_UDF_READ_INVALID_MAX_ARRAY_LENGTH_ARG_TYPE : {13482,[]string{"HY000"},"Invalid \"max_array_length\" argument type."},
	ER_LOG_CANNOT_WRITE_EXTENDED : {13483,[]string{"HY000"},"Failed to write to %s: %s (%s)"},
	//OBSOLETE_ER_UPGRADE_WITH_PARTITIONED_TABLES_REJECTED : {13484,[]string{"HY000"},"Upgrading from server version %d with partitioned tables and lower_case_table_names == 1 on a case sensitive file system may cause issues, and is therefore prohibited. To upgrade anyway, restart the new server version with the command line option 'upgrade=FORCE'. When upgrade is completed, please execute 'RENAME TABLE <part_table_name> TO <new_table_name>; RENAME TABLE <new_table_name> TO <part_table_name>;' for each of the partitioned tables. Please see the documentation for further information."},
	ER_KEYRING_AWS_INCORRECT_PROXY : {13485,[]string{"HY000"},"Incorrect environment variable %s, invalid port: %s"},
	ER_GRP_RPL_SERVER_SET_TO_OFFLINE_MODE_DUE_TO_ERRORS : {13486,[]string{"HY000"},"The server was automatically set into offline mode after an error was detected."},
	ER_GRP_RPL_MESSAGE_SERVICE_FATAL_ERROR : {13487,[]string{"HY000"},"Message delivery error on message service of Group Replication. The server will now leave the group."},
	ER_WARN_WRONG_COMPRESSION_ALGORITHM_LOG : {13488,[]string{"HY000"},"Invalid MASTER_COMPRESSION_ALGORITHMS '%.192s' found in repository for channel '%.192s'. Resetting to 'uncompressed' (no compression)."},
	ER_WARN_WRONG_COMPRESSION_LEVEL_LOG : {13489,[]string{"HY000"},"Invalid MASTER_ZSTD_COMPRESSION_LEVEL found in repository for channel '%.192s'. Resetting to %u."},
	ER_PROTOCOL_COMPRESSION_RESET_LOG : {13490,[]string{"HY000"},"Option --protocol-compression-algorithms is reset to default value."},
	ER_XPLUGIN_COMPRESSION_ERROR : {13491,[]string{"HY000"},"Fatal error while compressing outgoing data - %s"},
	ER_MYSQLBACKUP_MSG : {13492,[]string{"HY000"},"%s"},
	ER_WARN_UNKNOWN_KEYRING_AWS_REGION : {13493,[]string{"HY000"},"Unknown keyring_aws_region '%.192s'. Connection to AWS KMS may fail."},
	ER_WARN_LOG_PRIVILEGE_CHECKS_USER_DOES_NOT_EXIST : {13494,[]string{"HY000"},"PRIVILEGE_CHECKS_USER for replication channel '%.192s' was set to `%.64s`@`%.255s`, but this is not an existing user. Correct this before starting replication threads."},
	ER_WARN_LOG_PRIVILEGE_CHECKS_USER_CORRUPT : {13495,[]string{"HY000"},"Invalid, corrupted PRIVILEGE_CHECKS_USER was found in the replication configuration repository for channel '%.192s'. Use CHANGE MASTER TO PRIVILEGE_CHECKS_USER to correct the configuration."},
	ER_WARN_LOG_PRIVILEGE_CHECKS_USER_NEEDS_RPL_APPLIER_PRIV : {13496,[]string{"HY000"},"PRIVILEGE_CHECKS_USER for replication channel '%.192s' was set to `%.64s`@`%.255s`, but this user does not have REPLICATION_APPLIER privilege. Correct this before starting the replication threads."},
	ER_FILE_PRIVILEGE_FOR_REPLICATION_CHECKS : {13497,[]string{"HY000"},"The PRIVILEGE_CHECKS_USER for channel '%.192s' would need FILE privilege to execute a LOAD DATA INFILE statement replicated in statement format. Consider using binlog_format=ROW on master. If the replicated events are trusted, recover from the failure by temporarily granting FILE to the PRIVILEGE_CHECKS_USER."},
	ER_RPL_SLAVE_SQL_THREAD_STARTING_WITH_PRIVILEGE_CHECKS : {13498,[]string{"HY000"},"Slave SQL thread%s initialized, starting replication in log '%s' at position %s, relay log '%s' position: %s, user: '%.64s'@'%.255s', roles: %.512s"},
	ER_AUDIT_LOG_CANNOT_GENERATE_PASSWORD : {13499,[]string{"HY000"},"Cannot generate password: '%.32s'"},
	ER_INIT_FAILED_TO_GENERATE_ROOT_PASSWORD : {13500,[]string{"HY000"},"Failed to generate a random password for root. Probabably not enough enthropy."},
	ER_PLUGIN_LOAD_OPTIONS_IGNORED : {13501,[]string{"HY000"},"Ignoring --plugin-load[_add] list as the server is running with --initialize(-insecure)."},
	ER_WARN_AUTH_ID_WITH_SYSTEM_USER_PRIV_IN_MANDATORY_ROLES : {13502,[]string{"HY000"},"Cannot set mandatory_roles: AuthId `%.64s`@`%.64s` has '%s' privilege. AuthId(s) set in the mandatory_roles are ignored."},
	ER_IB_MSG_SKIP_HIDDEN_DIR : {13503,[]string{"HY000"},"Directory '%s' will not be scanned because it is a hidden directory."},
	ER_WARN_RPL_RECOVERY_NO_ROTATE_EVENT_FROM_MASTER_EOF : {13504,[]string{"HY000"},"Server was not able to find a rotate event from master server to initialize relay log recovery for channel '%s'. Skipping relay log recovery for the channel."},
	ER_IB_LOB_ROLLBACK_INDEX_LEN : {13505,[]string{"HY000"},"Rolling back LOB for transaction %llu undo number %llu : current index length %llu. (iteration %llu)"},
	ER_CANT_PROCESS_EXPRESSION_FOR_GENERATED_COLUMN_TO_DD : {13506,[]string{"HY000"},"Error in processing (possibly deprecated) expression or function '%.128s' for generated column %.64s.%.64s.%.64s"},
	ER_RPL_SLAVE_QUEUE_EVENT_FAILED_INVALID_NON_ROW_FORMAT : {13507,[]string{"HY000"},"The queue event failed for channel '%s' as an invalid event according to REQUIRE_ROW_FORMAT was found."},
	ER_RPL_SLAVE_APPLY_LOG_EVENT_FAILED_INVALID_NON_ROW_FORMAT : {13508,[]string{"HY000"},"The application of relay events failed for channel '%s' as an invalid event according to REQUIRE_ROW_FORMAT was found."},
	ER_LOG_PRIV_CHECKS_REQUIRE_ROW_FORMAT_NOT_SET : {13509,[]string{"HY000"},"PRIVILEGE_CHECKS_USER for replication channel '%.192s' can't be set to `%.64s`@`%.255s` unless REQUIRE_ROW_FORMAT is also set to %d."},
	ER_RPL_SLAVE_SQL_THREAD_DETECTED_UNEXPECTED_EVENT_SEQUENCE : {13510,[]string{"HY000"},"An unexpected event sequence was detected by the SQL thread while applying an event."},
	ER_IB_MSG_UPGRADE_PARTITION_FILE : {13511,[]string{"HY000"},"Updating partition file name '%s' to '%s' and all other partition files during upgrade"},
	ER_IB_MSG_DOWNGRADE_PARTITION_FILE : {13512,[]string{"HY000"},"Updating partition file name '%s' to '%s' and all other partition files during downgrade"},
	ER_IB_MSG_UPGRADE_PARTITION_FILE_IMPORT : {13513,[]string{"HY000"},"Updating partition file name '%s' to '%s' for import"},
	ER_IB_WARN_OPEN_PARTITION_FILE : {13514,[]string{"HY000"},"Unable to open partition file with new name '%s'. Please check if innodb_directories is set to include all external file paths"},
	ER_IB_MSG_FIL_STATE_MOVED_CORRECTED : {13515,[]string{"HY000"},"%s DD ID: %llu - Partition tablespace %u, name '%s' is corrected to '%s'"},
	ER_IB_MSG_FIL_STATE_MOVED_CHANGED_PATH : {13516,[]string{"HY000"},"%s DD ID: %llu - Tablespace %u, name '%s', '%s' is moved to '%s'"},
	ER_IB_MSG_FIL_STATE_MOVED_CHANGED_NAME : {13517,[]string{"HY000"},"%s DD ID: %llu - Partition tablespace %u, name '%s', '%s' is updated to '%s'"},
	ER_IB_MSG_FIL_STATE_MOVED_TOO_MANY : {13518,[]string{"HY000"},"%s Too many files have been moved, disabling logging of detailed messages"},
	ER_GR_ELECTED_PRIMARY_GTID_INFORMATION : {13519,[]string{"HY000"},"Elected primary member %s: %s"},
	ER_SCHEMA_NAME_IN_UPPER_CASE_NOT_ALLOWED : {13520,[]string{"HY000"},"Schema name '%s' containing upper case characters is not allowed with lower_case_table_names = 1."},
	ER_TABLE_NAME_IN_UPPER_CASE_NOT_ALLOWED : {13521,[]string{"HY000"},"Table name '%s.%s' containing upper case characters is not allowed with lower_case_table_names = 1."},
	ER_SCHEMA_NAME_IN_UPPER_CASE_NOT_ALLOWED_FOR_FK : {13522,[]string{"HY000"},"Schema name '%s' containing upper case characters, used by foreign key '%s' in table '%s.%s', is not allowed with lower_case_table_names = 1."},
	ER_TABLE_NAME_IN_UPPER_CASE_NOT_ALLOWED_FOR_FK : {13523,[]string{"HY000"},"Table name '%s.%s' containing upper case characters, used by foreign key '%s' in table '%s.%s', is not allowed with lower_case_table_names = 1."},
	ER_IB_MSG_DICT_PARTITION_NOT_FOUND : {13524,[]string{"HY000"},"Table Partition: %s is not found in InnoDB dictionary"},
	ER_ACCESS_DENIED_FOR_USER_ACCOUNT_BLOCKED_BY_PASSWORD_LOCK : {13525,[]string{"HY000"},"Access denied for user '%-.48s'@'%-.64s'. Account is blocked for %s day(s) (%s day(s) remaining) due to %u consecutive failed logins. Use FLUSH PRIVILEGES or ALTER USER to reset."},
	ER_INNODB_OUT_OF_RESOURCES : {13526,[]string{"HY000"},"%s"},
	ER_DD_UPGRADE_FOUND_PREPARED_XA_TRANSACTION : {13527,[]string{"HY000"},"Upgrade cannot proceed due to an existing prepared XA transaction."},
	ER_MIGRATE_TABLE_TO_DD_OOM : {13528,[]string{"HY000"},"Could not allocate memory for key_info when migrating table %s.%s"},
	ER_RPL_RELAY_LOG_RECOVERY_INFO_AFTER_CLONE : {13529,[]string{"HY000"},"Relay log information for channel '%s' was found after a clone operation. Relay log recovery will be executed to adjust positions and file information for this new server. Should that automatic procedure fail please adjust the positions through 'CHANGE MASTER TO'"},
	ER_IB_MSG_57_UNDO_SPACE_DELETE_FAIL : {13530,[]string{"HY000"},"Failed to delete 5.7 undo tablespace: %s during upgrade"},
	ER_IB_MSG_DBLWR_1285 : {13531,[]string{"HY000"},"Empty doublewrite file: %s"},
	ER_IB_MSG_DBLWR_1286 : {13532,[]string{"HY000"},"Using '%s' for doublewrite"},
	ER_IB_MSG_DBLWR_1287 : {13533,[]string{"HY000"},"Error reading doublewrite buffer from the system tablespace"},
	ER_IB_MSG_DBLWR_1288 : {13534,[]string{"HY000"},"Cannot create doublewrite buffer: you must increase your buffer pool size. Cannot continue operation."},
	ER_IB_MSG_DBLWR_1290 : {13535,[]string{"HY000"},"The page in the doublewrite file is corrupt. Cannot continue operation. You can try to recover the database with innodb_force_recovery=6"},
	ER_IB_MSG_BAD_DBLWR_FILE_NAME : {13536,[]string{"HY000"},"The doublewrite filename '%s' is incorrect."},
	//OBSOLETE_ER_IB_MSG_DBLWR_1292 : {13537,[]string{"HY000"},"%s"},
	ER_IB_MSG_DBLWR_1293 : {13538,[]string{"HY000"},"Doublewrite file create failed: %s"},
	ER_IB_MSG_DBLWR_1294 : {13539,[]string{"HY000"},"DBLWRThread: pthread_setaffinity() failed!"},
	ER_IB_MSG_DBLWR_1295 : {13540,[]string{"HY000"},"%s"},
	ER_IB_MSG_DBLWR_1296 : {13541,[]string{"HY000"},"%s"},
	ER_IB_MSG_DBLWR_1297 : {13542,[]string{"HY000"},"Doublewrite file read failed: %s"},
	ER_IB_MSG_DBLWR_1298 : {13543,[]string{"HY000"},"Dump of the data file page:"},
	ER_IB_MSG_DBLWR_1300 : {13544,[]string{"HY000"},"%s"},
	ER_IB_MSG_DBLWR_1301 : {13545,[]string{"HY000"},"%s"},
	ER_IB_MSG_DBLWR_1304 : {13546,[]string{"HY000"},"%s"},
	ER_IB_MSG_DBLWR_1305 : {13547,[]string{"HY000"},"%s"},
	ER_IB_MSG_DBLWR_1306 : {13548,[]string{"HY000"},"%s"},
	ER_IB_MSG_DBLWR_1307 : {13549,[]string{"HY000"},"%s"},
	ER_IB_MSG_DBLWR_1308 : {13550,[]string{"HY000"},"%s"},
	ER_IB_MSG_DBLWR_1309 : {13551,[]string{"HY000"},"%s"},
	ER_IB_MSG_DBLWR_1310 : {13552,[]string{"HY000"},"%s"},
	ER_IB_MSG_DBLWR_1311 : {13553,[]string{"HY000"},"%s"},
	ER_IB_MSG_DBLWR_1312 : {13554,[]string{"HY000"},"%s"},
	ER_IB_MSG_DBLWR_1313 : {13555,[]string{"HY000"},"%s"},
	ER_IB_MSG_DBLWR_1314 : {13556,[]string{"HY000"},"%s"},
	ER_IB_MSG_DBLWR_1315 : {13557,[]string{"HY000"},"%s"},
	ER_IB_MSG_DBLWR_1316 : {13558,[]string{"HY000"},"%s"},
	ER_IB_MSG_DBLWR_1317 : {13559,[]string{"HY000"},"%s"},
	ER_IB_MSG_DBLWR_1318 : {13560,[]string{"HY000"},"%s"},
	ER_IB_MSG_DBLWR_1319 : {13561,[]string{"HY000"},"Doublewrite load file %s size %lu is not a multiple of the configured page size %lu"},
	ER_IB_MSG_DBLWR_1320 : {13562,[]string{"HY000"},"Doublewrite file %s truncate failed"},
	ER_IB_MSG_DBLWR_1321 : {13563,[]string{"HY000"},"Doublewrite file %s failed to writ zeros"},
	ER_IB_MSG_DBLWR_1322 : {13564,[]string{"HY000"},"Doublewrite create file %s size %lu is not a multiple of the configured page size %lu"},
	ER_IB_MSG_DBLWR_1323 : {13565,[]string{"HY000"},"%s"},
	ER_IB_MSG_DBLWR_1324 : {13566,[]string{"HY000"},"%s"},
	ER_IB_MSG_DBLWR_1325 : {13567,[]string{"HY000"},"%s"},
	ER_IB_MSG_DBLWR_1326 : {13568,[]string{"HY000"},"%s"},
	ER_IB_MSG_DBLWR_1327 : {13569,[]string{"HY000"},"%s"},
	ER_IB_MSG_GTID_FLUSH_AT_SHUTDOWN : {13570,[]string{"HY000"},"Could not flush all GTIDs during slow shutdown. Will recover GTIDs when server restarts."},
	ER_IB_MSG_57_STAT_SPACE_DELETE_FAIL : {13571,[]string{"HY000"},"Failed to delete 5.7 stat tablespace: %s during upgrade"},
	ER_NDBINFO_UPGRADING_SCHEMA : {13572,[]string{"HY000"},"Installing ndbinfo schema version %s"},
	ER_NDBINFO_NOT_UPGRADING_SCHEMA : {13573,[]string{"HY000"},"Installed ndbinfo schema is current. Not upgrading."},
	ER_NDBINFO_UPGRADING_SCHEMA_FAIL : {13574,[]string{"HY000"},"Failed to upgrade ndbinfo schema."},
	ER_IB_MSG_CREATE_LOG_FILE : {13575,[]string{"HY000"},"Creating log file %s"},
	ER_IB_MSG_INNODB_START_INITIALIZE : {13576,[]string{"HY000"},"InnoDB initialization has started."},
	ER_IB_MSG_INNODB_END_INITIALIZE : {13577,[]string{"HY000"},"InnoDB initialization has ended."},
	ER_IB_MSG_PAGE_ARCH_NO_RESET_POINTS : {13578,[]string{"HY000"},"Could not find appropriate reset points."},
	ER_IB_WRN_PAGE_ARCH_FLUSH_DATA : {13579,[]string{"HY000"},"Unable to flush. Page archiving data may be corrupt in case of a crash."},
	ER_IB_ERR_PAGE_ARCH_INVALID_DOUBLE_WRITE_BUF : {13580,[]string{"HY000"},"Page archiver's doublewrite buffer for %ld is not valid."},
	ER_IB_ERR_PAGE_ARCH_RECOVERY_FAILED : {13581,[]string{"HY000"},"Page archiver system's recovery failed."},
	ER_IB_ERR_PAGE_ARCH_INVALID_FORMAT : {13582,[]string{"HY000"},"Invalid archived file name format. The archived file is supposed to have the format %s + [0-9]*."},
	ER_INVALID_XPLUGIN_SOCKET_SAME_AS_SERVER : {13583,[]string{"HY000"},"X Plugins UNIX socket must use different file than MySQL server. X Plugin won't be accessible through UNIX socket"},
	ER_INNODB_UNABLE_TO_ACQUIRE_DD_OBJECT : {13584,[]string{"HY000"},"%s"},
	ER_WARN_LOG_DEPRECATED_PARTITION_PREFIX_KEY : {13585,[]string{"HY000"},"Column '%.64s.%.64s.%.64s' having prefix key part '%.64s(%u)' is ignored by the partitioning function. Use of prefixed columns in the PARTITION BY KEY() clause is deprecated and will be removed in a future release."},
	ER_IB_MSG_UNDO_TRUNCATE_TOO_OFTEN : {13586,[]string{"HY000"},"Undo Truncation is occurring too often. Consider increasing --innodb-max-undo-log-size."},
	ER_GRP_RPL_IS_STARTING : {13587,[]string{"HY000"},"Plugin 'group_replication' is starting."},
	ER_IB_MSG_INVALID_LOCATION_FOR_TABLESPACE : {13588,[]string{"HY000"},"Cannot create tablespace %s because the directory is not a valid location. %s"},
	ER_IB_MSG_INVALID_LOCATION_WRONG_DB : {13589,[]string{"HY000"},"Scanned file '%s' for tablespace %s cannot be opened because it is not in a sub-directory named for the schema."},
	ER_IB_MSG_CANNOT_FIND_DD_UNDO_SPACE : {13590,[]string{"HY000"},"Cannot find undo tablespace %s with filename '%s' as indicated by the Data Dictionary. Did you move or delete this tablespace? Any undo logs in it cannot be used."},
	ER_GRP_RPL_RECOVERY_ENDPOINT_FORMAT : {13591,[]string{"HY000"},"Invalid input value for recovery socket endpoints '%s'. Please, provide a valid, comma separated, list of endpoints (IP:port)"},
	ER_GRP_RPL_RECOVERY_ENDPOINT_INVALID : {13592,[]string{"HY000"},"The server is not listening on endpoint '%s'. Only endpoints that the server is listening on are valid recovery endpoints."},
	ER_GRP_RPL_RECOVERY_ENDPOINT_INVALID_DONOR_ENDPOINT : {13593,[]string{"HY000"},"Received invalid recovery endpoints configuration from donor. This member is not a valid donor for recovery, so it will be skipped."},
	ER_GRP_RPL_RECOVERY_ENDPOINT_INTERFACES_IPS : {13594,[]string{"HY000"},"Failed to retrieve IP addresses from enabled host network interfaces."},
	ER_WARN_TLS_CHANNEL_INITIALIZATION_ERROR : {13595,[]string{"HY000"},"Failed to initialize TLS for channel: %s. See below for the description of exact issue."},
	ER_XPLUGIN_FAILED_TO_VALIDATE_ADDRESS : {13596,[]string{"HY000"},"Validation of value '%s' set to `Mysqlx_bind_address` failed: %s. Skipping this value."},
	ER_XPLUGIN_FAILED_TO_BIND_INTERFACE_ADDRESS : {13597,[]string{"HY000"},"Value '%s' set to `Mysqlx_bind_address`, X Plugin can't bind to it. Skipping this value."},
	ER_IB_ERR_RECOVERY_REDO_DISABLED : {13598,[]string{"HY000"},"Server was killed when InnoDB redo logging was disabled. Data files could be corrupt. You can try to restart the database with innodb_force_recovery=6"},
	ER_IB_WRN_FAST_SHUTDOWN_REDO_DISABLED : {13599,[]string{"HY000"},"InnoDB cannot do cold shutdown 'innodb_fast_shutdown = 2' and is forcing 'innodb_fast_shutdown = 1' as redo logging is disabled. InnoDB would flush all dirty pages to ensure physical data consistency."},
	ER_IB_WRN_REDO_DISABLED : {13600,[]string{"HY000"},"InnoDB redo logging is disabled. All data could be lost in case of a server crash."},
	ER_IB_WRN_REDO_ENABLED : {13601,[]string{"HY000"},"InnoDB redo logging is enabled. Data is now safe and can be recovered in case of a server crash."},
	ER_TLS_CONFIGURED_FOR_CHANNEL : {13602,[]string{"HY000"},"Channel %s configured to support TLS. Encrypted connections are now supported for this channel."},
	ER_TLS_CONFIGURATION_REUSED : {13603,[]string{"HY000"},"No TLS configuration was given for channel %s; re-using TLS configuration of channel %s."},
	ER_IB_TABLESPACE_PATH_VALIDATION_SKIPPED : {13604,[]string{"HY000"},"Skipping InnoDB tablespace path validation. Manually moved tablespace files will not be detected!"},
	ER_IB_CANNOT_UPGRADE_WITH_DISCARDED_TABLESPACES : {13605,[]string{"HY000"},"Upgrade failed because database contains discarded tablespaces."},
	ER_USERNAME_TRUNKATED : {13606,[]string{"HY000"},"The user name '%s' exceeds the maximum number of allowed characters %d and is trunkated."},
	ER_HOSTNAME_TRUNKATED : {13607,[]string{"HY000"},"The host name '%s' exceeds the maximum number of allowed characters %d and is trunkated."},
	ER_IB_MSG_TRX_RECOVERY_ROLLBACK_NOT_COMPLETED : {13608,[]string{"HY000"},"Rollback of non-prepared transactions not completed, due to fast shutdown"},
	ER_AUTHCACHE_ROLE_EDGES_IGNORED_EMPTY_NAME : {13609,[]string{"HY000"},"Found an entry in the 'role_edges' table with empty authorization ID; Skipped"},
	ER_AUTHCACHE_ROLE_EDGES_UNKNOWN_AUTHORIZATION_ID : {13610,[]string{"HY000"},"Found an entry in the 'role_edges' table with unknown authorization ID '%s'; Skipped"},
	ER_AUTHCACHE_DEFAULT_ROLES_IGNORED_EMPTY_NAME : {13611,[]string{"HY000"},"Found an entry in the 'default_roles' table with empty authorization ID; Skipped"},
	ER_AUTHCACHE_DEFAULT_ROLES_UNKNOWN_AUTHORIZATION_ID : {13612,[]string{"HY000"},"Found an entry in the 'default_roles' table with unknown authorization ID '%s'; Skipped"},
	ER_IB_ERR_DDL_LOG_INSERT_FAILURE : {13613,[]string{"HY000"},"Couldn't insert entry in ddl log for ddl."},
	ER_IB_LOCK_VALIDATE_LATCH_ORDER_VIOLATION : {13614,[]string{"HY000"},"%s"},
	ER_IB_RELOCK_LATCH_ORDER_VIOLATION : {13615,[]string{"HY000"},"%s"},
	//OBSOLETE_ER_IB_MSG_1352 : {0000,[]string{""},"%s"},
	//OBSOLETE_ER_IB_MSG_1353 : {0000,[]string{""},"%s"},
	//OBSOLETE_ER_IB_MSG_1354 : {0000,[]string{""},"%s"},
	//OBSOLETE_ER_IB_MSG_1355 : {0000,[]string{""},"%s"},
	//OBSOLETE_ER_IB_MSG_1356 : {0000,[]string{""},"%s"},
	ER_IB_MSG_1357 : {13621,[]string{"HY000"},"%s"},
	ER_IB_MSG_1358 : {13622,[]string{"HY000"},"%s"},
	ER_IB_MSG_1359 : {13623,[]string{"HY000"},"%s"},
	ER_IB_FAILED_TO_DELETE_TABLESPACE_FILE : {13624,[]string{"HY000"},"%s"},
	ER_IB_UNABLE_TO_EXPAND_TEMPORARY_TABLESPACE_POOL : {13625,[]string{"HY000"},"%s"},
	ER_IB_TMP_TABLESPACE_CANNOT_CREATE_DIRECTORY : {13626,[]string{"HY000"},"%s"},
	ER_IB_MSG_SCANNING_TEMP_TABLESPACE_DIR : {13627,[]string{"HY000"},"%s"},
	ER_IB_ERR_TEMP_TABLESPACE_DIR_DOESNT_EXIST : {13628,[]string{"HY000"},"%s"},
	ER_IB_ERR_TEMP_TABLESPACE_DIR_EMPTY : {13629,[]string{"HY000"},"%s"},
	ER_IB_ERR_TEMP_TABLESPACE_DIR_CONTAINS_SEMICOLON : {13630,[]string{"HY000"},"%s"},
	ER_IB_ERR_TEMP_TABLESPACE_DIR_SUBDIR_OF_DATADIR : {13631,[]string{"HY000"},"%s"},
	ER_IB_ERR_SCHED_SETAFFNINITY_FAILED : {13632,[]string{"HY000"},"%s"},
	ER_IB_ERR_UNKNOWN_PAGE_FETCH_MODE : {13633,[]string{"HY000"},"%s"},
	ER_IB_ERR_LOG_PARSING_BUFFER_OVERFLOW : {13634,[]string{"HY000"},"%s"},
	ER_IB_ERR_NOT_ENOUGH_MEMORY_FOR_PARSE_BUFFER : {13635,[]string{"HY000"},"%s"},
	ER_IB_MSG_1372 : {13636,[]string{"HY000"},"%s"},
	ER_IB_MSG_1373 : {13637,[]string{"HY000"},"%s"},
	ER_IB_MSG_1374 : {13638,[]string{"HY000"},"%s"},
	ER_IB_MSG_1375 : {13639,[]string{"HY000"},"%s"},
	ER_IB_ERR_ZLIB_UNCOMPRESS_FAILED : {13640,[]string{"HY000"},"%s"},
	ER_IB_ERR_ZLIB_BUF_ERROR : {13641,[]string{"HY000"},"%s"},
	ER_IB_ERR_ZLIB_MEM_ERROR : {13642,[]string{"HY000"},"%s"},
	ER_IB_ERR_ZLIB_DATA_ERROR : {13643,[]string{"HY000"},"%s"},
	ER_IB_ERR_ZLIB_UNKNOWN_ERROR : {13644,[]string{"HY000"},"%s"},
	ER_IB_MSG_1381 : {13645,[]string{"HY000"},"%s"},
	ER_IB_ERR_INDEX_RECORDS_WRONG_ORDER : {13646,[]string{"HY000"},"%s"},
	ER_IB_ERR_INDEX_DUPLICATE_KEY : {13647,[]string{"HY000"},"%s"},
	ER_IB_ERR_FOUND_N_DUPLICATE_KEYS : {13648,[]string{"HY000"},"%s"},
	ER_IB_ERR_FOUND_N_RECORDS_WRONG_ORDER : {13649,[]string{"HY000"},"%s"},
	ER_IB_ERR_PARALLEL_READ_OOM : {13650,[]string{"HY000"},"%s"},
	ER_IB_MSG_UNDO_MARKED_ACTIVE : {13651,[]string{"HY000"},"The state of undo tablespace %s is set to active implicitly."},
	ER_IB_MSG_UNDO_ALTERED_ACTIVE : {13652,[]string{"HY000"},"The state of undo tablespace %s is set to 'active' by ALTER TABLESPACE."},
	ER_IB_MSG_UNDO_ALTERED_INACTIVE : {13653,[]string{"HY000"},"The state of undo tablespace %s is set to 'inactive' by ALTER TABLESPACE."},
	ER_IB_MSG_UNDO_MARKED_EMPTY : {13654,[]string{"HY000"},"The state of undo tablespace %s is set to 'empty'."},
	ER_IB_MSG_UNDO_TRUNCATE_DELAY_BY_CLONE : {13655,[]string{"HY000"},"Delaying truncate of undo tablespace %s due to clone activity."},
	ER_IB_MSG_UNDO_TRUNCATE_DELAY_BY_MDL : {13656,[]string{"HY000"},"Delaying truncate of undo tablespace %s due to a metadata lock."},
	ER_IB_MSG_INJECT_CRASH : {13657,[]string{"HY000"},"Injected debug crash point: %s"},
	ER_IB_MSG_INJECT_FAILURE : {13658,[]string{"HY000"},"Injected debug failure point: %s"},
	ER_GRP_RPL_TIMEOUT_RECEIVED_VC_LEAVE_ON_REJOIN : {13659,[]string{"HY000"},"Timeout while waiting for a view change event during the leave step before a auto-rejoin attempt."},
	ER_RPL_ASYNC_RECONNECT_FAIL_NO_SOURCE : {13660,[]string{"HY000"},"Failed to automatically re-connect to a different source, for channel '%s', because %s. To fix this %s."},
	ER_UDF_REGISTER_SERVICE_ERROR : {13661,[]string{"HY000"},"Could not execute the installation of UDF functions. Check for other errors in the log"},
	ER_UDF_REGISTER_ERROR : {13662,[]string{"HY000"},"Could not execute the installation of UDF function: %s. Check if the function is already present, if so, try to remove it."},
	ER_UDF_UNREGISTER_ERROR : {13663,[]string{"HY000"},"Could not uninstall UDF functions. Try to remove them manually if present."},
	ER_EMPTY_PRIVILEGE_NAME_IGNORED : {13664,[]string{"HY000"},"An empty or illegal privilege identifier was ignored when global privileges were read from disk."},
	ER_IB_MSG_INCORRECT_SIZE : {13665,[]string{"HY000"},"%s"},
	ER_TMPDIR_PATH_TOO_LONG : {13666,[]string{"HY000"},"A tmpdir temporary path \"%s\" is too long (> %zu) for this OS. This would not leave enough space for a temporary filename of length %zu within it."},
	ER_ERROR_LOG_DESTINATION_NOT_A_FILE : {13667,[]string{"HY000"},"Error-log destination \"%s\" is not a file. Can not restore error log messages from previous run."},
	ER_NO_ERROR_LOG_PARSER_CONFIGURED : {13668,[]string{"HY000"},"None of the log-sinks selected with --log-error-services=... provides a log-parser. The server will not be able to make the previous runs' error-logs available in performance_schema.error_log."},
	ER_UPGRADE_NONEXISTENT_SCHEMA : {13669,[]string{"HY000"},"The schema \"%.64s\" referenced by %.16s \"%.128s\" does not exist. Please clean up any orphan %.16s before upgrading."},
	ER_IB_MSG_CREATED_UNDO_SPACE : {13670,[]string{"HY000"},"Created undo tablespace '%s'."},
	ER_IB_MSG_DROPPED_UNDO_SPACE : {13671,[]string{"HY000"},"Dropped undo tablespace '%s'."},
	ER_IB_MSG_MASTER_KEY_ROTATED : {13672,[]string{"HY000"},"The InnoDB Encryption Master Key has been rotated in %d tablespaces."},
	ER_IB_DBLWR_DECOMPRESS_FAILED : {13673,[]string{"HY000"},"Failed to decompress a DBLWR page (err=%d).  The original size is %d. Reporting the dblwr page as corrupted."},
	ER_IB_DBLWR_DECRYPT_FAILED : {13674,[]string{"HY000"},"Decrypting a page in doublewrite file failed: %s."},
	ER_IB_DBLWR_KEY_MISSING : {13675,[]string{"HY000"},"Encryption key missing: %s."},
	ER_INNODB_IO_WRITE_ERROR_RETRYING : {13676,[]string{"HY000"},"I/O error while writing to file: %s. Retrying ..."},
	ER_INNODB_IO_WRITE_FAILED : {13677,[]string{"HY000"},"Failed to write data to file: %s"},
	ER_LOG_COMPONENT_CANNOT_INIT : {13678,[]string{"HY000"},"Log component %s failed to initialize."},
	ER_RPL_ASYNC_CHANNEL_CANT_CONNECT : {13679,[]string{"HY000"},"The Monitor IO thread failed to connect to the source (host:%s port:%u network_namespace:%s) for channel '%s', thence it will try to connect to another source."},
	ER_RPL_ASYNC_SENDER_ADDED : {13680,[]string{"HY000"},"The source (host:%s port:%u network_namespace:%s) for channel '%s' has joined the group (group_name: %s), and so added its entry into replication_asynchronous_connection_failover table."},
	ER_RPL_ASYNC_SENDER_REMOVED : {13681,[]string{"HY000"},"The source (host:%s port:%u network_namespace:%s) for channel '%s' has left the group (group_name: %s), and so removed its entry from replication_asynchronous_connection_failover table."},
	ER_RPL_ASYNC_CHANNEL_STOPPED_QUORUM_LOST : {13682,[]string{"HY000"},"The Monitor IO thread detected that the source (host:%s port:%u network_namespace:%s) does not belong to the group majority, thence the channel '%s' will try to connect to another source."},
	ER_RPL_ASYNC_CHANNEL_CANT_CONNECT_NO_QUORUM : {13683,[]string{"HY000"},"The IO thread detected that the source (host:%s port:%u network_namespace:%s) does not belong to the group majority, thence the channel '%s' will try to connect to another source."},
	ER_RPL_ASYNC_EXECUTING_QUERY : {13684,[]string{"HY000"},"%s on the source (host:%s port:%u network_namespace:%s) for channel '%s'."},
	ER_RPL_REPLICA_MONITOR_IO_THREAD_EXITING : {13685,[]string{"HY000"},"Replica Monitor IO thread exiting."},
	ER_RPL_ASYNC_MANAGED_NAME_REMOVED : {13686,[]string{"HY000"},"The group (group_name: %s) for the channel '%s' has been removed, and so removed its entry from replication_asynchronous_connection_failover_managed and all the group members from replication_asynchronous_connection_failover table."},
	ER_RPL_ASYNC_MANAGED_NAME_ADDED : {13687,[]string{"HY000"},"The group (group_name: %s) for the channel '%s' has been added, and so added its entry in replication_asynchronous_connection_failover_managed and source to replication_asynchronous_connection_failover table."},
	ER_RPL_ASYNC_READ_FAILOVER_TABLE : {13688,[]string{"HY000"},"Error reading failover sources for channel '%s' from replication_asynchronous_connection_failover table."},
	ER_RPL_REPLICA_MONITOR_IO_THREAD_RECONNECT_CHANNEL : {13689,[]string{"HY000"},"Error %s the channel '%s', the operation will be automatically retried."},
	ER_SLAVE_ANONYMOUS_TO_GTID_IS_LOCAL_OR_UUID_AND_GTID_MODE_NOT_ON : {13690,[]string{"HY000"},"Replication channel '%.192s' is configured with ASSIGN_GTIDS_TO_ANONYMOUS_TRANSACTIONS='%s', which is invalid when GTID_MODE <> ON. If you intend to use GTID_MODE = ON everywhere, change to ASSIGN_GTIDS_TO_ANONYMOUS_TRANSACTIONS = OFF and use the procedure for enabling GTIDs online (see the documentation). If you intend to use GTIDs on this replica and cannot enable GTIDs on the source, enable GTID_MODE = ON and leave ASSIGN_GTIDS_TO_ANONYMOUS_TRANSACTIONS = LOCAL|<UUID>. If you intend to not use GTIDs at all in the replication topology, change to ASSIGN_GTIDS_TO_ANONYMOUS_TRANSACTIONS=OFF and leave GTID_MODE = '%s'."},
	ER_REPLICA_ANONYMOUS_TO_GTID_UUID_SAME_AS_GROUP_NAME : {13691,[]string{"HY000"},"Replication channel '%.192s' is configured with ASSIGN_GTIDS_TO_ANONYMOUS_TRANSACTIONS='%s' which is equal to group_replication_group_name. To fix this issue, either change the group_replication_group_name  or use a different value for ASSIGN_GTIDS_TO_ANONYMOUS_TRANSACTIONS."},
	ER_GRP_RPL_GRP_NAME_IS_SAME_AS_ANONYMOUS_TO_GTID_UUID : {13692,[]string{"HY000"},"The group_replication_group_name '%s' is the same as the UUID value for ASSIGN_GTIDS_TO_ANONYMOUS_TRANSACTIONS in a server channel"},
	ER_WARN_GTID_THRESHOLD_BREACH : {13693,[]string{"HY000"},"The integer component of the GTID number is high. Suggest restarting the server with a new server_uuid to prevent it from reaching the maximum number 2^63-1, which will make it impossible to write the binary log and invoke the behavior specified by binlog_error_action."},
	ER_HEALTH_INFO : {13694,[]string{"HY000"},"%s"},
	ER_HEALTH_WARNING : {13695,[]string{"HY000"},"%s"},
	ER_HEALTH_ERROR : {13696,[]string{"HY000"},"%s"},
	ER_HEALTH_WARNING_DISK_USAGE_LEVEL_1 : {13697,[]string{"HY000"},"Disk usage on '%s' is %.2f%%. Warning level 1 (%d%%) exceeded."},
	ER_HEALTH_WARNING_DISK_USAGE_LEVEL_2 : {13698,[]string{"HY000"},"Disk usage on '%s' is %.2f%%. Warning level 2 (%d%%) exceeded."},
	ER_HEALTH_WARNING_DISK_USAGE_LEVEL_3 : {13699,[]string{"HY000"},"Disk usage on '%s' is %.2f%%. Warning level 3 (%d%%) exceeded."},
	ER_IB_INNODB_TBSP_OUT_OF_SPACE : {13700,[]string{"HY000"},"InnoDB: Size of tablespace %s is more than the maximum size allowed."},
	ER_GRP_RPL_APPLIER_CHANNEL_STILL_RUNNING : {13701,[]string{"HY000"},"The group_replication_applier channel is still running, most likely it is waiting for a database/table lock, which is preventing the channel from stopping. Please check database/table locks, including the ones created by backup tools."},
	ER_RPL_ASYNC_RECONNECT_GTID_MODE_OFF_CHANNEL : {13702,[]string{"HY000"},"Detected misconfiguration: replication channel '%.192s' was configured with SOURCE_CONNECTION_AUTO_FAILOVER = 1, but the server was started with a value other then --gtid-mode = ON. Either reconfigure replication using CHANGE MASTER TO SOURCE_CONNECTION_AUTO_FAILOVER = 0 FOR CHANNEL '%.192s', or change GTID_MODE to value ON, before starting the replica receiver thread."},
	ER_FIREWALL_SERVICES_NOT_ACQUIRED : {13703,[]string{"HY000"},"Could not acquire required component services."},
	ER_FIREWALL_UDF_REGISTER_FAILED : {13704,[]string{"HY000"},"Automatic registration of function(s) failed."},
	ER_FIREWALL_PFS_TABLE_REGISTER_FAILED : {13705,[]string{"HY000"},"Automatic registration of Performance schema table(s) failed."},
	ER_IB_MSG_STATS_SAMPLING_TOO_LARGE : {13706,[]string{"HY000"},"%s"},

}
