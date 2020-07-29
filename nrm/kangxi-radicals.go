package nrm

import (
	"bytes"
	"io"
	"unicode"
	"unicode/utf8"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gnkf/ecode"
)

var kangxiRadicals = unicode.SpecialCase{
	unicode.CaseRange{0x2f00, 0x2f00, [unicode.MaxCase]rune{0x4e00 - 0x2f00, 0, 0}}, // U+2F00 '⼀' -> U+4E00 '一'
	unicode.CaseRange{0x2f01, 0x2f01, [unicode.MaxCase]rune{0x4e28 - 0x2f01, 0, 0}}, // U+2F01 '⼁' -> U+4E28 '丨'
	unicode.CaseRange{0x2f02, 0x2f02, [unicode.MaxCase]rune{0x4e36 - 0x2f02, 0, 0}}, // U+2F02 '⼂' -> U+4E36 '丶'
	unicode.CaseRange{0x2f03, 0x2f03, [unicode.MaxCase]rune{0x4e3f - 0x2f03, 0, 0}}, // U+2F03 '⼃' -> U+4E3F '丿'
	unicode.CaseRange{0x2f04, 0x2f04, [unicode.MaxCase]rune{0x4e59 - 0x2f04, 0, 0}}, // U+2F04 '⼄' -> U+4E59 '乙'
	unicode.CaseRange{0x2f05, 0x2f05, [unicode.MaxCase]rune{0x4e85 - 0x2f05, 0, 0}}, // U+2F05 '⼅' -> U+4E85 '亅'
	unicode.CaseRange{0x2f06, 0x2f06, [unicode.MaxCase]rune{0x4e8c - 0x2f06, 0, 0}}, // U+2F06 '⼆' -> U+4E8C '二'
	unicode.CaseRange{0x2f07, 0x2f07, [unicode.MaxCase]rune{0x4ea0 - 0x2f07, 0, 0}}, // U+2F07 '⼇' -> U+4EA0 '亠'
	unicode.CaseRange{0x2f08, 0x2f08, [unicode.MaxCase]rune{0x4eba - 0x2f08, 0, 0}}, // U+2F08 '⼈' -> U+4EBA '人'
	unicode.CaseRange{0x2f09, 0x2f09, [unicode.MaxCase]rune{0x513f - 0x2f09, 0, 0}}, // U+2F09 '⼉' -> U+513F '儿'
	unicode.CaseRange{0x2f0a, 0x2f0a, [unicode.MaxCase]rune{0x5165 - 0x2f0a, 0, 0}}, // U+2F0A '⼊' -> U+5165 '入'
	unicode.CaseRange{0x2f0b, 0x2f0b, [unicode.MaxCase]rune{0x516b - 0x2f0b, 0, 0}}, // U+2F0B '⼋' -> U+516B '八'
	unicode.CaseRange{0x2f0c, 0x2f0c, [unicode.MaxCase]rune{0x5182 - 0x2f0c, 0, 0}}, // U+2F0C '⼌' -> U+5182 '冂'
	unicode.CaseRange{0x2f0d, 0x2f0d, [unicode.MaxCase]rune{0x5196 - 0x2f0d, 0, 0}}, // U+2F0D '⼍' -> U+5196 '冖'
	unicode.CaseRange{0x2f0e, 0x2f0e, [unicode.MaxCase]rune{0x51ab - 0x2f0e, 0, 0}}, // U+2F0E '⼎' -> U+51AB '冫'
	unicode.CaseRange{0x2f0f, 0x2f0f, [unicode.MaxCase]rune{0x51e0 - 0x2f0f, 0, 0}}, // U+2F0F '⼏' -> U+51E0 '几'
	unicode.CaseRange{0x2f10, 0x2f10, [unicode.MaxCase]rune{0x51f5 - 0x2f10, 0, 0}}, // U+2F10 '⼐' -> U+51F5 '凵'
	unicode.CaseRange{0x2f11, 0x2f11, [unicode.MaxCase]rune{0x5200 - 0x2f11, 0, 0}}, // U+2F11 '⼑' -> U+5200 '刀'
	unicode.CaseRange{0x2f12, 0x2f12, [unicode.MaxCase]rune{0x529b - 0x2f12, 0, 0}}, // U+2F12 '⼒' -> U+529B '力'
	unicode.CaseRange{0x2f13, 0x2f13, [unicode.MaxCase]rune{0x52f9 - 0x2f13, 0, 0}}, // U+2F13 '⼓' -> U+52F9 '勹'
	unicode.CaseRange{0x2f14, 0x2f14, [unicode.MaxCase]rune{0x5315 - 0x2f14, 0, 0}}, // U+2F14 '⼔' -> U+5315 '匕'
	unicode.CaseRange{0x2f15, 0x2f15, [unicode.MaxCase]rune{0x531a - 0x2f15, 0, 0}}, // U+2F15 '⼕' -> U+531A '匚'
	unicode.CaseRange{0x2f16, 0x2f16, [unicode.MaxCase]rune{0x5338 - 0x2f16, 0, 0}}, // U+2F16 '⼖' -> U+5338 '匸'
	unicode.CaseRange{0x2f17, 0x2f17, [unicode.MaxCase]rune{0x5341 - 0x2f17, 0, 0}}, // U+2F17 '⼗' -> U+5341 '十'
	unicode.CaseRange{0x2f18, 0x2f18, [unicode.MaxCase]rune{0x535c - 0x2f18, 0, 0}}, // U+2F18 '⼘' -> U+535C '卜'
	unicode.CaseRange{0x2f19, 0x2f19, [unicode.MaxCase]rune{0x5369 - 0x2f19, 0, 0}}, // U+2F19 '⼙' -> U+5369 '卩'
	unicode.CaseRange{0x2f1a, 0x2f1a, [unicode.MaxCase]rune{0x5382 - 0x2f1a, 0, 0}}, // U+2F1A '⼚' -> U+5382 '厂'
	unicode.CaseRange{0x2f1b, 0x2f1b, [unicode.MaxCase]rune{0x53b6 - 0x2f1b, 0, 0}}, // U+2F1B '⼛' -> U+53B6 '厶'
	unicode.CaseRange{0x2f1c, 0x2f1c, [unicode.MaxCase]rune{0x53c8 - 0x2f1c, 0, 0}}, // U+2F1C '⼜' -> U+53C8 '又'
	unicode.CaseRange{0x2f1d, 0x2f1d, [unicode.MaxCase]rune{0x53e3 - 0x2f1d, 0, 0}}, // U+2F1D '⼝' -> U+53E3 '口'
	unicode.CaseRange{0x2f1e, 0x2f1e, [unicode.MaxCase]rune{0x56d7 - 0x2f1e, 0, 0}}, // U+2F1E '⼞' -> U+56D7 '囗'
	unicode.CaseRange{0x2f1f, 0x2f1f, [unicode.MaxCase]rune{0x571f - 0x2f1f, 0, 0}}, // U+2F1F '⼟' -> U+571F '土'
	unicode.CaseRange{0x2f20, 0x2f20, [unicode.MaxCase]rune{0x58eb - 0x2f20, 0, 0}}, // U+2F20 '⼠' -> U+58EB '士'
	unicode.CaseRange{0x2f21, 0x2f21, [unicode.MaxCase]rune{0x5902 - 0x2f21, 0, 0}}, // U+2F21 '⼡' -> U+5902 '夂'
	unicode.CaseRange{0x2f22, 0x2f22, [unicode.MaxCase]rune{0x590a - 0x2f22, 0, 0}}, // U+2F22 '⼢' -> U+590A '夊'
	unicode.CaseRange{0x2f23, 0x2f23, [unicode.MaxCase]rune{0x5915 - 0x2f23, 0, 0}}, // U+2F23 '⼣' -> U+5915 '夕'
	unicode.CaseRange{0x2f24, 0x2f24, [unicode.MaxCase]rune{0x5927 - 0x2f24, 0, 0}}, // U+2F24 '⼤' -> U+5927 '大'
	unicode.CaseRange{0x2f25, 0x2f25, [unicode.MaxCase]rune{0x5973 - 0x2f25, 0, 0}}, // U+2F25 '⼥' -> U+5973 '女'
	unicode.CaseRange{0x2f26, 0x2f26, [unicode.MaxCase]rune{0x5b50 - 0x2f26, 0, 0}}, // U+2F26 '⼦' -> U+5B50 '子'
	unicode.CaseRange{0x2f27, 0x2f27, [unicode.MaxCase]rune{0x5b80 - 0x2f27, 0, 0}}, // U+2F27 '⼧' -> U+5B80 '宀'
	unicode.CaseRange{0x2f28, 0x2f28, [unicode.MaxCase]rune{0x5bf8 - 0x2f28, 0, 0}}, // U+2F28 '⼨' -> U+5BF8 '寸'
	unicode.CaseRange{0x2f29, 0x2f29, [unicode.MaxCase]rune{0x5c0f - 0x2f29, 0, 0}}, // U+2F29 '⼩' -> U+5C0F '小'
	unicode.CaseRange{0x2f2a, 0x2f2a, [unicode.MaxCase]rune{0x5c22 - 0x2f2a, 0, 0}}, // U+2F2A '⼪' -> U+5C22 '尢'
	unicode.CaseRange{0x2f2b, 0x2f2b, [unicode.MaxCase]rune{0x5c38 - 0x2f2b, 0, 0}}, // U+2F2B '⼫' -> U+5C38 '尸'
	unicode.CaseRange{0x2f2c, 0x2f2c, [unicode.MaxCase]rune{0x5c6e - 0x2f2c, 0, 0}}, // U+2F2C '⼬' -> U+5C6E '屮'
	unicode.CaseRange{0x2f2d, 0x2f2d, [unicode.MaxCase]rune{0x5c71 - 0x2f2d, 0, 0}}, // U+2F2D '⼭' -> U+5C71 '山'
	unicode.CaseRange{0x2f2e, 0x2f2e, [unicode.MaxCase]rune{0x5ddb - 0x2f2e, 0, 0}}, // U+2F2E '⼮' -> U+5DDB '巛'
	unicode.CaseRange{0x2f2f, 0x2f2f, [unicode.MaxCase]rune{0x5de5 - 0x2f2f, 0, 0}}, // U+2F2F '⼯' -> U+5DE5 '工'
	unicode.CaseRange{0x2f30, 0x2f30, [unicode.MaxCase]rune{0x5df1 - 0x2f30, 0, 0}}, // U+2F30 '⼰' -> U+5DF1 '己'
	unicode.CaseRange{0x2f31, 0x2f31, [unicode.MaxCase]rune{0x5dfe - 0x2f31, 0, 0}}, // U+2F31 '⼱' -> U+5DFE '巾'
	unicode.CaseRange{0x2f32, 0x2f32, [unicode.MaxCase]rune{0x5e72 - 0x2f32, 0, 0}}, // U+2F32 '⼲' -> U+5E72 '干'
	unicode.CaseRange{0x2f33, 0x2f33, [unicode.MaxCase]rune{0x5e7a - 0x2f33, 0, 0}}, // U+2F33 '⼳' -> U+5E7A '幺'
	unicode.CaseRange{0x2f34, 0x2f34, [unicode.MaxCase]rune{0x5e7f - 0x2f34, 0, 0}}, // U+2F34 '⼴' -> U+5E7F '广'
	unicode.CaseRange{0x2f35, 0x2f35, [unicode.MaxCase]rune{0x5ef4 - 0x2f35, 0, 0}}, // U+2F35 '⼵' -> U+5EF4 '廴'
	unicode.CaseRange{0x2f36, 0x2f36, [unicode.MaxCase]rune{0x5efe - 0x2f36, 0, 0}}, // U+2F36 '⼶' -> U+5EFE '廾'
	unicode.CaseRange{0x2f37, 0x2f37, [unicode.MaxCase]rune{0x5f0b - 0x2f37, 0, 0}}, // U+2F37 '⼷' -> U+5F0B '弋'
	unicode.CaseRange{0x2f38, 0x2f38, [unicode.MaxCase]rune{0x5f13 - 0x2f38, 0, 0}}, // U+2F38 '⼸' -> U+5F13 '弓'
	unicode.CaseRange{0x2f39, 0x2f39, [unicode.MaxCase]rune{0x5f50 - 0x2f39, 0, 0}}, // U+2F39 '⼹' -> U+5F50 '彐'
	unicode.CaseRange{0x2f3a, 0x2f3a, [unicode.MaxCase]rune{0x5f61 - 0x2f3a, 0, 0}}, // U+2F3A '⼺' -> U+5F61 '彡'
	unicode.CaseRange{0x2f3b, 0x2f3b, [unicode.MaxCase]rune{0x5f73 - 0x2f3b, 0, 0}}, // U+2F3B '⼻' -> U+5F73 '彳'
	unicode.CaseRange{0x2f3c, 0x2f3c, [unicode.MaxCase]rune{0x5fc3 - 0x2f3c, 0, 0}}, // U+2F3C '⼼' -> U+5FC3 '心'
	unicode.CaseRange{0x2f3d, 0x2f3d, [unicode.MaxCase]rune{0x6208 - 0x2f3d, 0, 0}}, // U+2F3D '⼽' -> U+6208 '戈'
	unicode.CaseRange{0x2f3e, 0x2f3e, [unicode.MaxCase]rune{0x6236 - 0x2f3e, 0, 0}}, // U+2F3E '⼾' -> U+6236 '戶'
	unicode.CaseRange{0x2f3f, 0x2f3f, [unicode.MaxCase]rune{0x624b - 0x2f3f, 0, 0}}, // U+2F3F '⼿' -> U+624B '手'
	unicode.CaseRange{0x2f40, 0x2f40, [unicode.MaxCase]rune{0x652f - 0x2f40, 0, 0}}, // U+2F40 '⽀' -> U+652F '支'
	unicode.CaseRange{0x2f41, 0x2f41, [unicode.MaxCase]rune{0x6534 - 0x2f41, 0, 0}}, // U+2F41 '⽁' -> U+6534 '攴'
	unicode.CaseRange{0x2f42, 0x2f42, [unicode.MaxCase]rune{0x6587 - 0x2f42, 0, 0}}, // U+2F42 '⽂' -> U+6587 '文'
	unicode.CaseRange{0x2f43, 0x2f43, [unicode.MaxCase]rune{0x6597 - 0x2f43, 0, 0}}, // U+2F43 '⽃' -> U+6597 '斗'
	unicode.CaseRange{0x2f44, 0x2f44, [unicode.MaxCase]rune{0x65a4 - 0x2f44, 0, 0}}, // U+2F44 '⽄' -> U+65A4 '斤'
	unicode.CaseRange{0x2f45, 0x2f45, [unicode.MaxCase]rune{0x65b9 - 0x2f45, 0, 0}}, // U+2F45 '⽅' -> U+65B9 '方'
	unicode.CaseRange{0x2f46, 0x2f46, [unicode.MaxCase]rune{0x65e0 - 0x2f46, 0, 0}}, // U+2F46 '⽆' -> U+65E0 '无'
	unicode.CaseRange{0x2f47, 0x2f47, [unicode.MaxCase]rune{0x65e5 - 0x2f47, 0, 0}}, // U+2F47 '⽇' -> U+65E5 '日'
	unicode.CaseRange{0x2f48, 0x2f48, [unicode.MaxCase]rune{0x66f0 - 0x2f48, 0, 0}}, // U+2F48 '⽈' -> U+66F0 '曰'
	unicode.CaseRange{0x2f49, 0x2f49, [unicode.MaxCase]rune{0x6708 - 0x2f49, 0, 0}}, // U+2F49 '⽉' -> U+6708 '月'
	unicode.CaseRange{0x2f4a, 0x2f4a, [unicode.MaxCase]rune{0x6728 - 0x2f4a, 0, 0}}, // U+2F4A '⽊' -> U+6728 '木'
	unicode.CaseRange{0x2f4b, 0x2f4b, [unicode.MaxCase]rune{0x6b20 - 0x2f4b, 0, 0}}, // U+2F4B '⽋' -> U+6B20 '欠'
	unicode.CaseRange{0x2f4c, 0x2f4c, [unicode.MaxCase]rune{0x6b62 - 0x2f4c, 0, 0}}, // U+2F4C '⽌' -> U+6B62 '止'
	unicode.CaseRange{0x2f4d, 0x2f4d, [unicode.MaxCase]rune{0x6b79 - 0x2f4d, 0, 0}}, // U+2F4D '⽍' -> U+6B79 '歹'
	unicode.CaseRange{0x2f4e, 0x2f4e, [unicode.MaxCase]rune{0x6bb3 - 0x2f4e, 0, 0}}, // U+2F4E '⽎' -> U+6BB3 '殳'
	unicode.CaseRange{0x2f4f, 0x2f4f, [unicode.MaxCase]rune{0x6bcb - 0x2f4f, 0, 0}}, // U+2F4F '⽏' -> U+6BCB '毋'
	unicode.CaseRange{0x2f50, 0x2f50, [unicode.MaxCase]rune{0x6bd4 - 0x2f50, 0, 0}}, // U+2F50 '⽐' -> U+6BD4 '比'
	unicode.CaseRange{0x2f51, 0x2f51, [unicode.MaxCase]rune{0x6bdb - 0x2f51, 0, 0}}, // U+2F51 '⽑' -> U+6BDB '毛'
	unicode.CaseRange{0x2f52, 0x2f52, [unicode.MaxCase]rune{0x6c0f - 0x2f52, 0, 0}}, // U+2F52 '⽒' -> U+6C0F '氏'
	unicode.CaseRange{0x2f53, 0x2f53, [unicode.MaxCase]rune{0x6c14 - 0x2f53, 0, 0}}, // U+2F53 '⽓' -> U+6C14 '气'
	unicode.CaseRange{0x2f54, 0x2f54, [unicode.MaxCase]rune{0x6c34 - 0x2f54, 0, 0}}, // U+2F54 '⽔' -> U+6C34 '水'
	unicode.CaseRange{0x2f55, 0x2f55, [unicode.MaxCase]rune{0x706b - 0x2f55, 0, 0}}, // U+2F55 '⽕' -> U+706B '火'
	unicode.CaseRange{0x2f56, 0x2f56, [unicode.MaxCase]rune{0x722a - 0x2f56, 0, 0}}, // U+2F56 '⽖' -> U+722A '爪'
	unicode.CaseRange{0x2f57, 0x2f57, [unicode.MaxCase]rune{0x7236 - 0x2f57, 0, 0}}, // U+2F57 '⽗' -> U+7236 '父'
	unicode.CaseRange{0x2f58, 0x2f58, [unicode.MaxCase]rune{0x723b - 0x2f58, 0, 0}}, // U+2F58 '⽘' -> U+723B '爻'
	unicode.CaseRange{0x2f59, 0x2f59, [unicode.MaxCase]rune{0x723f - 0x2f59, 0, 0}}, // U+2F59 '⽙' -> U+723F '爿'
	unicode.CaseRange{0x2f5a, 0x2f5a, [unicode.MaxCase]rune{0x7247 - 0x2f5a, 0, 0}}, // U+2F5A '⽚' -> U+7247 '片'
	unicode.CaseRange{0x2f5b, 0x2f5b, [unicode.MaxCase]rune{0x7259 - 0x2f5b, 0, 0}}, // U+2F5B '⽛' -> U+7259 '牙'
	unicode.CaseRange{0x2f5c, 0x2f5c, [unicode.MaxCase]rune{0x725b - 0x2f5c, 0, 0}}, // U+2F5C '⽜' -> U+725B '牛'
	unicode.CaseRange{0x2f5d, 0x2f5d, [unicode.MaxCase]rune{0x72ac - 0x2f5d, 0, 0}}, // U+2F5D '⽝' -> U+72AC '犬'
	unicode.CaseRange{0x2f5e, 0x2f5e, [unicode.MaxCase]rune{0x7384 - 0x2f5e, 0, 0}}, // U+2F5E '⽞' -> U+7384 '玄'
	unicode.CaseRange{0x2f5f, 0x2f5f, [unicode.MaxCase]rune{0x7389 - 0x2f5f, 0, 0}}, // U+2F5F '⽟' -> U+7389 '玉'
	unicode.CaseRange{0x2f60, 0x2f60, [unicode.MaxCase]rune{0x74dc - 0x2f60, 0, 0}}, // U+2F60 '⽠' -> U+74DC '瓜'
	unicode.CaseRange{0x2f61, 0x2f61, [unicode.MaxCase]rune{0x74e6 - 0x2f61, 0, 0}}, // U+2F61 '⽡' -> U+74E6 '瓦'
	unicode.CaseRange{0x2f62, 0x2f62, [unicode.MaxCase]rune{0x7518 - 0x2f62, 0, 0}}, // U+2F62 '⽢' -> U+7518 '甘'
	unicode.CaseRange{0x2f63, 0x2f63, [unicode.MaxCase]rune{0x751f - 0x2f63, 0, 0}}, // U+2F63 '⽣' -> U+751F '生'
	unicode.CaseRange{0x2f64, 0x2f64, [unicode.MaxCase]rune{0x7528 - 0x2f64, 0, 0}}, // U+2F64 '⽤' -> U+7528 '用'
	unicode.CaseRange{0x2f65, 0x2f65, [unicode.MaxCase]rune{0x7530 - 0x2f65, 0, 0}}, // U+2F65 '⽥' -> U+7530 '田'
	unicode.CaseRange{0x2f66, 0x2f66, [unicode.MaxCase]rune{0x758b - 0x2f66, 0, 0}}, // U+2F66 '⽦' -> U+758B '疋'
	unicode.CaseRange{0x2f67, 0x2f67, [unicode.MaxCase]rune{0x7592 - 0x2f67, 0, 0}}, // U+2F67 '⽧' -> U+7592 '疒'
	unicode.CaseRange{0x2f68, 0x2f68, [unicode.MaxCase]rune{0x7676 - 0x2f68, 0, 0}}, // U+2F68 '⽨' -> U+7676 '癶'
	unicode.CaseRange{0x2f69, 0x2f69, [unicode.MaxCase]rune{0x767d - 0x2f69, 0, 0}}, // U+2F69 '⽩' -> U+767D '白'
	unicode.CaseRange{0x2f6a, 0x2f6a, [unicode.MaxCase]rune{0x76ae - 0x2f6a, 0, 0}}, // U+2F6A '⽪' -> U+76AE '皮'
	unicode.CaseRange{0x2f6b, 0x2f6b, [unicode.MaxCase]rune{0x76bf - 0x2f6b, 0, 0}}, // U+2F6B '⽫' -> U+76BF '皿'
	unicode.CaseRange{0x2f6c, 0x2f6c, [unicode.MaxCase]rune{0x76ee - 0x2f6c, 0, 0}}, // U+2F6C '⽬' -> U+76EE '目'
	unicode.CaseRange{0x2f6d, 0x2f6d, [unicode.MaxCase]rune{0x77db - 0x2f6d, 0, 0}}, // U+2F6D '⽭' -> U+77DB '矛'
	unicode.CaseRange{0x2f6e, 0x2f6e, [unicode.MaxCase]rune{0x77e2 - 0x2f6e, 0, 0}}, // U+2F6E '⽮' -> U+77E2 '矢'
	unicode.CaseRange{0x2f6f, 0x2f6f, [unicode.MaxCase]rune{0x77f3 - 0x2f6f, 0, 0}}, // U+2F6F '⽯' -> U+77F3 '石'
	unicode.CaseRange{0x2f70, 0x2f70, [unicode.MaxCase]rune{0x793a - 0x2f70, 0, 0}}, // U+2F70 '⽰' -> U+793A '示'
	unicode.CaseRange{0x2f71, 0x2f71, [unicode.MaxCase]rune{0x79b8 - 0x2f71, 0, 0}}, // U+2F71 '⽱' -> U+79B8 '禸'
	unicode.CaseRange{0x2f72, 0x2f72, [unicode.MaxCase]rune{0x79be - 0x2f72, 0, 0}}, // U+2F72 '⽲' -> U+79BE '禾'
	unicode.CaseRange{0x2f73, 0x2f73, [unicode.MaxCase]rune{0x7a74 - 0x2f73, 0, 0}}, // U+2F73 '⽳' -> U+7A74 '穴'
	unicode.CaseRange{0x2f74, 0x2f74, [unicode.MaxCase]rune{0x7acb - 0x2f74, 0, 0}}, // U+2F74 '⽴' -> U+7ACB '立'
	unicode.CaseRange{0x2f75, 0x2f75, [unicode.MaxCase]rune{0x7af9 - 0x2f75, 0, 0}}, // U+2F75 '⽵' -> U+7AF9 '竹'
	unicode.CaseRange{0x2f76, 0x2f76, [unicode.MaxCase]rune{0x7c73 - 0x2f76, 0, 0}}, // U+2F76 '⽶' -> U+7C73 '米'
	unicode.CaseRange{0x2f77, 0x2f77, [unicode.MaxCase]rune{0x7cf8 - 0x2f77, 0, 0}}, // U+2F77 '⽷' -> U+7CF8 '糸'
	unicode.CaseRange{0x2f78, 0x2f78, [unicode.MaxCase]rune{0x7f36 - 0x2f78, 0, 0}}, // U+2F78 '⽸' -> U+7F36 '缶'
	unicode.CaseRange{0x2f79, 0x2f79, [unicode.MaxCase]rune{0x7f51 - 0x2f79, 0, 0}}, // U+2F79 '⽹' -> U+7F51 '网'
	unicode.CaseRange{0x2f7a, 0x2f7a, [unicode.MaxCase]rune{0x7f8a - 0x2f7a, 0, 0}}, // U+2F7A '⽺' -> U+7F8A '羊'
	unicode.CaseRange{0x2f7b, 0x2f7b, [unicode.MaxCase]rune{0x7fbd - 0x2f7b, 0, 0}}, // U+2F7B '⽻' -> U+7FBD '羽'
	unicode.CaseRange{0x2f7c, 0x2f7c, [unicode.MaxCase]rune{0x8001 - 0x2f7c, 0, 0}}, // U+2F7C '⽼' -> U+8001 '老'
	unicode.CaseRange{0x2f7d, 0x2f7d, [unicode.MaxCase]rune{0x800c - 0x2f7d, 0, 0}}, // U+2F7D '⽽' -> U+800C '而'
	unicode.CaseRange{0x2f7e, 0x2f7e, [unicode.MaxCase]rune{0x8012 - 0x2f7e, 0, 0}}, // U+2F7E '⽾' -> U+8012 '耒'
	unicode.CaseRange{0x2f7f, 0x2f7f, [unicode.MaxCase]rune{0x8033 - 0x2f7f, 0, 0}}, // U+2F7F '⽿' -> U+8033 '耳'
	unicode.CaseRange{0x2f80, 0x2f80, [unicode.MaxCase]rune{0x807f - 0x2f80, 0, 0}}, // U+2F80 '⾀' -> U+807F '聿'
	unicode.CaseRange{0x2f81, 0x2f81, [unicode.MaxCase]rune{0x8089 - 0x2f81, 0, 0}}, // U+2F81 '⾁' -> U+8089 '肉'
	unicode.CaseRange{0x2f82, 0x2f82, [unicode.MaxCase]rune{0x81e3 - 0x2f82, 0, 0}}, // U+2F82 '⾂' -> U+81E3 '臣'
	unicode.CaseRange{0x2f83, 0x2f83, [unicode.MaxCase]rune{0x81ea - 0x2f83, 0, 0}}, // U+2F83 '⾃' -> U+81EA '自'
	unicode.CaseRange{0x2f84, 0x2f84, [unicode.MaxCase]rune{0x81f3 - 0x2f84, 0, 0}}, // U+2F84 '⾄' -> U+81F3 '至'
	unicode.CaseRange{0x2f85, 0x2f85, [unicode.MaxCase]rune{0x81fc - 0x2f85, 0, 0}}, // U+2F85 '⾅' -> U+81FC '臼'
	unicode.CaseRange{0x2f86, 0x2f86, [unicode.MaxCase]rune{0x820c - 0x2f86, 0, 0}}, // U+2F86 '⾆' -> U+820C '舌'
	unicode.CaseRange{0x2f87, 0x2f87, [unicode.MaxCase]rune{0x821b - 0x2f87, 0, 0}}, // U+2F87 '⾇' -> U+821B '舛'
	unicode.CaseRange{0x2f88, 0x2f88, [unicode.MaxCase]rune{0x821f - 0x2f88, 0, 0}}, // U+2F88 '⾈' -> U+821F '舟'
	unicode.CaseRange{0x2f89, 0x2f89, [unicode.MaxCase]rune{0x826e - 0x2f89, 0, 0}}, // U+2F89 '⾉' -> U+826E '艮'
	unicode.CaseRange{0x2f8a, 0x2f8a, [unicode.MaxCase]rune{0x8272 - 0x2f8a, 0, 0}}, // U+2F8A '⾊' -> U+8272 '色'
	unicode.CaseRange{0x2f8b, 0x2f8b, [unicode.MaxCase]rune{0x8278 - 0x2f8b, 0, 0}}, // U+2F8B '⾋' -> U+8278 '艸'
	unicode.CaseRange{0x2f8c, 0x2f8c, [unicode.MaxCase]rune{0x864d - 0x2f8c, 0, 0}}, // U+2F8C '⾌' -> U+864D '虍'
	unicode.CaseRange{0x2f8d, 0x2f8d, [unicode.MaxCase]rune{0x866b - 0x2f8d, 0, 0}}, // U+2F8D '⾍' -> U+866B '虫'
	unicode.CaseRange{0x2f8e, 0x2f8e, [unicode.MaxCase]rune{0x8840 - 0x2f8e, 0, 0}}, // U+2F8E '⾎' -> U+8840 '血'
	unicode.CaseRange{0x2f8f, 0x2f8f, [unicode.MaxCase]rune{0x884c - 0x2f8f, 0, 0}}, // U+2F8F '⾏' -> U+884C '行'
	unicode.CaseRange{0x2f90, 0x2f90, [unicode.MaxCase]rune{0x8863 - 0x2f90, 0, 0}}, // U+2F90 '⾐' -> U+8863 '衣'
	unicode.CaseRange{0x2f91, 0x2f91, [unicode.MaxCase]rune{0x897e - 0x2f91, 0, 0}}, // U+2F91 '⾑' -> U+897E '襾'
	unicode.CaseRange{0x2f92, 0x2f92, [unicode.MaxCase]rune{0x898b - 0x2f92, 0, 0}}, // U+2F92 '⾒' -> U+898B '見'
	unicode.CaseRange{0x2f93, 0x2f93, [unicode.MaxCase]rune{0x89d2 - 0x2f93, 0, 0}}, // U+2F93 '⾓' -> U+89D2 '角'
	unicode.CaseRange{0x2f94, 0x2f94, [unicode.MaxCase]rune{0x8a00 - 0x2f94, 0, 0}}, // U+2F94 '⾔' -> U+8A00 '言'
	unicode.CaseRange{0x2f95, 0x2f95, [unicode.MaxCase]rune{0x8c37 - 0x2f95, 0, 0}}, // U+2F95 '⾕' -> U+8C37 '谷'
	unicode.CaseRange{0x2f96, 0x2f96, [unicode.MaxCase]rune{0x8c46 - 0x2f96, 0, 0}}, // U+2F96 '⾖' -> U+8C46 '豆'
	unicode.CaseRange{0x2f97, 0x2f97, [unicode.MaxCase]rune{0x8c55 - 0x2f97, 0, 0}}, // U+2F97 '⾗' -> U+8C55 '豕'
	unicode.CaseRange{0x2f98, 0x2f98, [unicode.MaxCase]rune{0x8c78 - 0x2f98, 0, 0}}, // U+2F98 '⾘' -> U+8C78 '豸'
	unicode.CaseRange{0x2f99, 0x2f99, [unicode.MaxCase]rune{0x8c9d - 0x2f99, 0, 0}}, // U+2F99 '⾙' -> U+8C9D '貝'
	unicode.CaseRange{0x2f9a, 0x2f9a, [unicode.MaxCase]rune{0x8d64 - 0x2f9a, 0, 0}}, // U+2F9A '⾚' -> U+8D64 '赤'
	unicode.CaseRange{0x2f9b, 0x2f9b, [unicode.MaxCase]rune{0x8d70 - 0x2f9b, 0, 0}}, // U+2F9B '⾛' -> U+8D70 '走'
	unicode.CaseRange{0x2f9c, 0x2f9c, [unicode.MaxCase]rune{0x8db3 - 0x2f9c, 0, 0}}, // U+2F9C '⾜' -> U+8DB3 '足'
	unicode.CaseRange{0x2f9d, 0x2f9d, [unicode.MaxCase]rune{0x8eab - 0x2f9d, 0, 0}}, // U+2F9D '⾝' -> U+8EAB '身'
	unicode.CaseRange{0x2f9e, 0x2f9e, [unicode.MaxCase]rune{0x8eca - 0x2f9e, 0, 0}}, // U+2F9E '⾞' -> U+8ECA '車'
	unicode.CaseRange{0x2f9f, 0x2f9f, [unicode.MaxCase]rune{0x8f9b - 0x2f9f, 0, 0}}, // U+2F9F '⾟' -> U+8F9B '辛'
	unicode.CaseRange{0x2fa0, 0x2fa0, [unicode.MaxCase]rune{0x8fb0 - 0x2fa0, 0, 0}}, // U+2FA0 '⾠' -> U+8FB0 '辰'
	unicode.CaseRange{0x2fa1, 0x2fa1, [unicode.MaxCase]rune{0x8fb5 - 0x2fa1, 0, 0}}, // U+2FA1 '⾡' -> U+8FB5 '辵'
	unicode.CaseRange{0x2fa2, 0x2fa2, [unicode.MaxCase]rune{0x9091 - 0x2fa2, 0, 0}}, // U+2FA2 '⾢' -> U+9091 '邑'
	unicode.CaseRange{0x2fa3, 0x2fa3, [unicode.MaxCase]rune{0x9149 - 0x2fa3, 0, 0}}, // U+2FA3 '⾣' -> U+9149 '酉'
	unicode.CaseRange{0x2fa4, 0x2fa4, [unicode.MaxCase]rune{0x91c6 - 0x2fa4, 0, 0}}, // U+2FA4 '⾤' -> U+91C6 '釆'
	unicode.CaseRange{0x2fa5, 0x2fa5, [unicode.MaxCase]rune{0x91cc - 0x2fa5, 0, 0}}, // U+2FA5 '⾥' -> U+91CC '里'
	unicode.CaseRange{0x2fa6, 0x2fa6, [unicode.MaxCase]rune{0x91d1 - 0x2fa6, 0, 0}}, // U+2FA6 '⾦' -> U+91D1 '金'
	unicode.CaseRange{0x2fa7, 0x2fa7, [unicode.MaxCase]rune{0x9577 - 0x2fa7, 0, 0}}, // U+2FA7 '⾧' -> U+9577 '長'
	unicode.CaseRange{0x2fa8, 0x2fa8, [unicode.MaxCase]rune{0x9580 - 0x2fa8, 0, 0}}, // U+2FA8 '⾨' -> U+9580 '門'
	unicode.CaseRange{0x2fa9, 0x2fa9, [unicode.MaxCase]rune{0x961c - 0x2fa9, 0, 0}}, // U+2FA9 '⾩' -> U+961C '阜'
	unicode.CaseRange{0x2faa, 0x2faa, [unicode.MaxCase]rune{0x96b6 - 0x2faa, 0, 0}}, // U+2FAA '⾪' -> U+96B6 '隶'
	unicode.CaseRange{0x2fab, 0x2fab, [unicode.MaxCase]rune{0x96b9 - 0x2fab, 0, 0}}, // U+2FAB '⾫' -> U+96B9 '隹'
	unicode.CaseRange{0x2fac, 0x2fac, [unicode.MaxCase]rune{0x96e8 - 0x2fac, 0, 0}}, // U+2FAC '⾬' -> U+96E8 '雨'
	unicode.CaseRange{0x2fad, 0x2fad, [unicode.MaxCase]rune{0x9751 - 0x2fad, 0, 0}}, // U+2FAD '⾭' -> U+9751 '靑'
	unicode.CaseRange{0x2fae, 0x2fae, [unicode.MaxCase]rune{0x975e - 0x2fae, 0, 0}}, // U+2FAE '⾮' -> U+975E '非'
	unicode.CaseRange{0x2faf, 0x2faf, [unicode.MaxCase]rune{0x9762 - 0x2faf, 0, 0}}, // U+2FAF '⾯' -> U+9762 '面'
	unicode.CaseRange{0x2fb0, 0x2fb0, [unicode.MaxCase]rune{0x9769 - 0x2fb0, 0, 0}}, // U+2FB0 '⾰' -> U+9769 '革'
	unicode.CaseRange{0x2fb1, 0x2fb1, [unicode.MaxCase]rune{0x97cb - 0x2fb1, 0, 0}}, // U+2FB1 '⾱' -> U+97CB '韋'
	unicode.CaseRange{0x2fb2, 0x2fb2, [unicode.MaxCase]rune{0x97ed - 0x2fb2, 0, 0}}, // U+2FB2 '⾲' -> U+97ED '韭'
	unicode.CaseRange{0x2fb3, 0x2fb3, [unicode.MaxCase]rune{0x97f3 - 0x2fb3, 0, 0}}, // U+2FB3 '⾳' -> U+97F3 '音'
	unicode.CaseRange{0x2fb4, 0x2fb4, [unicode.MaxCase]rune{0x9801 - 0x2fb4, 0, 0}}, // U+2FB4 '⾴' -> U+9801 '頁'
	unicode.CaseRange{0x2fb5, 0x2fb5, [unicode.MaxCase]rune{0x98a8 - 0x2fb5, 0, 0}}, // U+2FB5 '⾵' -> U+98A8 '風'
	unicode.CaseRange{0x2fb6, 0x2fb6, [unicode.MaxCase]rune{0x98db - 0x2fb6, 0, 0}}, // U+2FB6 '⾶' -> U+98DB '飛'
	unicode.CaseRange{0x2fb7, 0x2fb7, [unicode.MaxCase]rune{0x98df - 0x2fb7, 0, 0}}, // U+2FB7 '⾷' -> U+98DF '食'
	unicode.CaseRange{0x2fb8, 0x2fb8, [unicode.MaxCase]rune{0x9996 - 0x2fb8, 0, 0}}, // U+2FB8 '⾸' -> U+9996 '首'
	unicode.CaseRange{0x2fb9, 0x2fb9, [unicode.MaxCase]rune{0x9999 - 0x2fb9, 0, 0}}, // U+2FB9 '⾹' -> U+9999 '香'
	unicode.CaseRange{0x2fba, 0x2fba, [unicode.MaxCase]rune{0x99ac - 0x2fba, 0, 0}}, // U+2FBA '⾺' -> U+99AC '馬'
	unicode.CaseRange{0x2fbb, 0x2fbb, [unicode.MaxCase]rune{0x9aa8 - 0x2fbb, 0, 0}}, // U+2FBB '⾻' -> U+9AA8 '骨'
	unicode.CaseRange{0x2fbc, 0x2fbc, [unicode.MaxCase]rune{0x9ad8 - 0x2fbc, 0, 0}}, // U+2FBC '⾼' -> U+9AD8 '高'
	unicode.CaseRange{0x2fbd, 0x2fbd, [unicode.MaxCase]rune{0x9adf - 0x2fbd, 0, 0}}, // U+2FBD '⾽' -> U+9ADF '髟'
	unicode.CaseRange{0x2fbe, 0x2fbe, [unicode.MaxCase]rune{0x9b25 - 0x2fbe, 0, 0}}, // U+2FBE '⾾' -> U+9B25 '鬥'
	unicode.CaseRange{0x2fbf, 0x2fbf, [unicode.MaxCase]rune{0x9b2f - 0x2fbf, 0, 0}}, // U+2FBF '⾿' -> U+9B2F '鬯'
	unicode.CaseRange{0x2fc0, 0x2fc0, [unicode.MaxCase]rune{0x9b32 - 0x2fc0, 0, 0}}, // U+2FC0 '⿀' -> U+9B32 '鬲'
	unicode.CaseRange{0x2fc1, 0x2fc1, [unicode.MaxCase]rune{0x9b3c - 0x2fc1, 0, 0}}, // U+2FC1 '⿁' -> U+9B3C '鬼'
	unicode.CaseRange{0x2fc2, 0x2fc2, [unicode.MaxCase]rune{0x9b5a - 0x2fc2, 0, 0}}, // U+2FC2 '⿂' -> U+9B5A '魚'
	unicode.CaseRange{0x2fc3, 0x2fc3, [unicode.MaxCase]rune{0x9ce5 - 0x2fc3, 0, 0}}, // U+2FC3 '⿃' -> U+9CE5 '鳥'
	unicode.CaseRange{0x2fc4, 0x2fc4, [unicode.MaxCase]rune{0x9e75 - 0x2fc4, 0, 0}}, // U+2FC4 '⿄' -> U+9E75 '鹵'
	unicode.CaseRange{0x2fc5, 0x2fc5, [unicode.MaxCase]rune{0x9e7f - 0x2fc5, 0, 0}}, // U+2FC5 '⿅' -> U+9E7F '鹿'
	unicode.CaseRange{0x2fc6, 0x2fc6, [unicode.MaxCase]rune{0x9ea5 - 0x2fc6, 0, 0}}, // U+2FC6 '⿆' -> U+9EA5 '麥'
	unicode.CaseRange{0x2fc7, 0x2fc7, [unicode.MaxCase]rune{0x9ebb - 0x2fc7, 0, 0}}, // U+2FC7 '⿇' -> U+9EBB '麻'
	unicode.CaseRange{0x2fc8, 0x2fc8, [unicode.MaxCase]rune{0x9ec3 - 0x2fc8, 0, 0}}, // U+2FC8 '⿈' -> U+9EC3 '黃'
	unicode.CaseRange{0x2fc9, 0x2fc9, [unicode.MaxCase]rune{0x9ecd - 0x2fc9, 0, 0}}, // U+2FC9 '⿉' -> U+9ECD '黍'
	unicode.CaseRange{0x2fca, 0x2fca, [unicode.MaxCase]rune{0x9ed1 - 0x2fca, 0, 0}}, // U+2FCA '⿊' -> U+9ED1 '黑'
	unicode.CaseRange{0x2fcb, 0x2fcb, [unicode.MaxCase]rune{0x9ef9 - 0x2fcb, 0, 0}}, // U+2FCB '⿋' -> U+9EF9 '黹'
	unicode.CaseRange{0x2fcc, 0x2fcc, [unicode.MaxCase]rune{0x9efd - 0x2fcc, 0, 0}}, // U+2FCC '⿌' -> U+9EFD '黽'
	unicode.CaseRange{0x2fcd, 0x2fcd, [unicode.MaxCase]rune{0x9f0e - 0x2fcd, 0, 0}}, // U+2FCD '⿍' -> U+9F0E '鼎'
	unicode.CaseRange{0x2fce, 0x2fce, [unicode.MaxCase]rune{0x9f13 - 0x2fce, 0, 0}}, // U+2FCE '⿎' -> U+9F13 '鼓'
	unicode.CaseRange{0x2fcf, 0x2fcf, [unicode.MaxCase]rune{0x9f20 - 0x2fcf, 0, 0}}, // U+2FCF '⿏' -> U+9F20 '鼠'
	unicode.CaseRange{0x2fd0, 0x2fd0, [unicode.MaxCase]rune{0x9f3b - 0x2fd0, 0, 0}}, // U+2FD0 '⿐' -> U+9F3B '鼻'
	unicode.CaseRange{0x2fd1, 0x2fd1, [unicode.MaxCase]rune{0x9f4a - 0x2fd1, 0, 0}}, // U+2FD1 '⿑' -> U+9F4A '齊'
	unicode.CaseRange{0x2fd2, 0x2fd2, [unicode.MaxCase]rune{0x9f52 - 0x2fd2, 0, 0}}, // U+2FD2 '⿒' -> U+9F52 '齒'
	unicode.CaseRange{0x2fd3, 0x2fd3, [unicode.MaxCase]rune{0x9f8d - 0x2fd3, 0, 0}}, // U+2FD3 '⿓' -> U+9F8D '龍'
	unicode.CaseRange{0x2fd4, 0x2fd4, [unicode.MaxCase]rune{0x9f9c - 0x2fd4, 0, 0}}, // U+2FD4 '⿔' -> U+9F9C '龜'
	unicode.CaseRange{0x2fd5, 0x2fd5, [unicode.MaxCase]rune{0x9fa0 - 0x2fd5, 0, 0}}, // U+2FD5 '⿕' -> U+9FA0 '龠'
}

//NormKangxiRadicalsBytes dose Unicode normalization kangxi radicals only
func NormKangxiRadicalsBytes(txt []byte) ([]byte, error) {
	if !utf8.Valid(txt) {
		return nil, errs.WrapWithCause(ecode.ErrInvalidUTF8Text, nil)
	}
	return bytes.ToUpperSpecial(kangxiRadicals, txt), nil
}

//NormKangxiRadicals dose Unicode normalization kangxi radicals only
func NormKangxiRadicals(writer io.Writer, txt io.Reader) error {
	buf := &bytes.Buffer{}
	if _, err := buf.ReadFrom(txt); err != nil {
		return errs.WrapWithCause(err, nil)
	}

	dst, err := NormKangxiRadicalsBytes(buf.Bytes())
	if err != nil {
		return errs.WrapWithCause(err, nil)
	}
	if _, err := io.Copy(writer, bytes.NewReader(dst)); err != nil {
		return errs.WrapWithCause(err, nil)
	}
	return nil
}

/* Copyright 2020 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
