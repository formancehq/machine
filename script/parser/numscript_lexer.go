// Code generated from NumScript.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 39, 324,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 3,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 6, 6, 96, 10, 6, 13, 6, 14, 6, 97, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 107, 10, 7, 12, 7, 14, 7,
	110, 11, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8,
	121, 10, 8, 12, 8, 14, 8, 124, 11, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3,
	18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23,
	3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28,
	3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 31, 3, 31, 7, 31, 239, 10, 31, 12, 31, 14, 31, 242, 11, 31, 3, 31, 3,
	31, 3, 32, 6, 32, 247, 10, 32, 13, 32, 14, 32, 248, 3, 32, 5, 32, 252,
	10, 32, 3, 32, 3, 32, 5, 32, 256, 10, 32, 3, 32, 6, 32, 259, 10, 32, 13,
	32, 14, 32, 260, 3, 32, 6, 32, 264, 10, 32, 13, 32, 14, 32, 265, 3, 32,
	3, 32, 6, 32, 270, 10, 32, 13, 32, 14, 32, 271, 5, 32, 274, 10, 32, 3,
	32, 5, 32, 277, 10, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 34, 6, 34, 290, 10, 34, 13, 34, 14, 34, 291, 3,
	35, 3, 35, 3, 36, 3, 36, 6, 36, 298, 10, 36, 13, 36, 14, 36, 299, 3, 36,
	7, 36, 303, 10, 36, 12, 36, 14, 36, 306, 11, 36, 3, 37, 3, 37, 6, 37, 310,
	10, 37, 13, 37, 14, 37, 311, 3, 37, 7, 37, 315, 10, 37, 12, 37, 14, 37,
	318, 11, 37, 3, 38, 6, 38, 321, 10, 38, 13, 38, 14, 38, 322, 4, 108, 122,
	2, 39, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21,
	12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39,
	21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57,
	30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75,
	39, 3, 2, 11, 4, 2, 12, 12, 15, 15, 4, 2, 11, 11, 34, 34, 5, 2, 50, 59,
	67, 92, 99, 124, 3, 2, 50, 59, 3, 2, 34, 34, 4, 2, 97, 97, 99, 124, 5,
	2, 50, 59, 97, 97, 99, 124, 5, 2, 50, 60, 97, 97, 99, 124, 4, 2, 49, 59,
	67, 92, 2, 342, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2,
	2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2,
	2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2,
	2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3,
	2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47,
	3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2,
	55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2,
	2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2,
	2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 3, 77, 3, 2,
	2, 2, 5, 80, 3, 2, 2, 2, 7, 90, 3, 2, 2, 2, 9, 92, 3, 2, 2, 2, 11, 95,
	3, 2, 2, 2, 13, 101, 3, 2, 2, 2, 15, 116, 3, 2, 2, 2, 17, 129, 3, 2, 2,
	2, 19, 134, 3, 2, 2, 2, 21, 140, 3, 2, 2, 2, 23, 145, 3, 2, 2, 2, 25, 150,
	3, 2, 2, 2, 27, 157, 3, 2, 2, 2, 29, 169, 3, 2, 2, 2, 31, 178, 3, 2, 2,
	2, 33, 180, 3, 2, 2, 2, 35, 182, 3, 2, 2, 2, 37, 184, 3, 2, 2, 2, 39, 186,
	3, 2, 2, 2, 41, 188, 3, 2, 2, 2, 43, 190, 3, 2, 2, 2, 45, 192, 3, 2, 2,
	2, 47, 194, 3, 2, 2, 2, 49, 196, 3, 2, 2, 2, 51, 198, 3, 2, 2, 2, 53, 206,
	3, 2, 2, 2, 55, 212, 3, 2, 2, 2, 57, 219, 3, 2, 2, 2, 59, 228, 3, 2, 2,
	2, 61, 236, 3, 2, 2, 2, 63, 276, 3, 2, 2, 2, 65, 278, 3, 2, 2, 2, 67, 289,
	3, 2, 2, 2, 69, 293, 3, 2, 2, 2, 71, 295, 3, 2, 2, 2, 73, 307, 3, 2, 2,
	2, 75, 320, 3, 2, 2, 2, 77, 78, 7, 118, 2, 2, 78, 79, 7, 113, 2, 2, 79,
	4, 3, 2, 2, 2, 80, 81, 7, 111, 2, 2, 81, 82, 7, 103, 2, 2, 82, 83, 7, 118,
	2, 2, 83, 84, 7, 99, 2, 2, 84, 85, 7, 102, 2, 2, 85, 86, 7, 99, 2, 2, 86,
	87, 7, 118, 2, 2, 87, 88, 7, 99, 2, 2, 88, 89, 7, 42, 2, 2, 89, 6, 3, 2,
	2, 2, 90, 91, 7, 46, 2, 2, 91, 8, 3, 2, 2, 2, 92, 93, 9, 2, 2, 2, 93, 10,
	3, 2, 2, 2, 94, 96, 9, 3, 2, 2, 95, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2,
	97, 95, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 100, 8,
	6, 2, 2, 100, 12, 3, 2, 2, 2, 101, 102, 7, 49, 2, 2, 102, 103, 7, 44, 2,
	2, 103, 108, 3, 2, 2, 2, 104, 107, 5, 13, 7, 2, 105, 107, 11, 2, 2, 2,
	106, 104, 3, 2, 2, 2, 106, 105, 3, 2, 2, 2, 107, 110, 3, 2, 2, 2, 108,
	109, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 109, 111, 3, 2, 2, 2, 110, 108,
	3, 2, 2, 2, 111, 112, 7, 44, 2, 2, 112, 113, 7, 49, 2, 2, 113, 114, 3,
	2, 2, 2, 114, 115, 8, 7, 2, 2, 115, 14, 3, 2, 2, 2, 116, 117, 7, 49, 2,
	2, 117, 118, 7, 49, 2, 2, 118, 122, 3, 2, 2, 2, 119, 121, 11, 2, 2, 2,
	120, 119, 3, 2, 2, 2, 121, 124, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 122,
	120, 3, 2, 2, 2, 123, 125, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 125, 126,
	5, 9, 5, 2, 126, 127, 3, 2, 2, 2, 127, 128, 8, 8, 2, 2, 128, 16, 3, 2,
	2, 2, 129, 130, 7, 120, 2, 2, 130, 131, 7, 99, 2, 2, 131, 132, 7, 116,
	2, 2, 132, 133, 7, 117, 2, 2, 133, 18, 3, 2, 2, 2, 134, 135, 7, 114, 2,
	2, 135, 136, 7, 116, 2, 2, 136, 137, 7, 107, 2, 2, 137, 138, 7, 112, 2,
	2, 138, 139, 7, 118, 2, 2, 139, 20, 3, 2, 2, 2, 140, 141, 7, 104, 2, 2,
	141, 142, 7, 99, 2, 2, 142, 143, 7, 107, 2, 2, 143, 144, 7, 110, 2, 2,
	144, 22, 3, 2, 2, 2, 145, 146, 7, 117, 2, 2, 146, 147, 7, 103, 2, 2, 147,
	148, 7, 112, 2, 2, 148, 149, 7, 102, 2, 2, 149, 24, 3, 2, 2, 2, 150, 151,
	7, 117, 2, 2, 151, 152, 7, 113, 2, 2, 152, 153, 7, 119, 2, 2, 153, 154,
	7, 116, 2, 2, 154, 155, 7, 101, 2, 2, 155, 156, 7, 103, 2, 2, 156, 26,
	3, 2, 2, 2, 157, 158, 7, 102, 2, 2, 158, 159, 7, 103, 2, 2, 159, 160, 7,
	117, 2, 2, 160, 161, 7, 118, 2, 2, 161, 162, 7, 107, 2, 2, 162, 163, 7,
	112, 2, 2, 163, 164, 7, 99, 2, 2, 164, 165, 7, 118, 2, 2, 165, 166, 7,
	107, 2, 2, 166, 167, 7, 113, 2, 2, 167, 168, 7, 112, 2, 2, 168, 28, 3,
	2, 2, 2, 169, 170, 7, 99, 2, 2, 170, 171, 7, 110, 2, 2, 171, 172, 7, 110,
	2, 2, 172, 173, 7, 113, 2, 2, 173, 174, 7, 101, 2, 2, 174, 175, 7, 99,
	2, 2, 175, 176, 7, 118, 2, 2, 176, 177, 7, 103, 2, 2, 177, 30, 3, 2, 2,
	2, 178, 179, 7, 45, 2, 2, 179, 32, 3, 2, 2, 2, 180, 181, 7, 47, 2, 2, 181,
	34, 3, 2, 2, 2, 182, 183, 7, 42, 2, 2, 183, 36, 3, 2, 2, 2, 184, 185, 7,
	43, 2, 2, 185, 38, 3, 2, 2, 2, 186, 187, 7, 93, 2, 2, 187, 40, 3, 2, 2,
	2, 188, 189, 7, 95, 2, 2, 189, 42, 3, 2, 2, 2, 190, 191, 7, 125, 2, 2,
	191, 44, 3, 2, 2, 2, 192, 193, 7, 127, 2, 2, 193, 46, 3, 2, 2, 2, 194,
	195, 7, 44, 2, 2, 195, 48, 3, 2, 2, 2, 196, 197, 7, 63, 2, 2, 197, 50,
	3, 2, 2, 2, 198, 199, 7, 99, 2, 2, 199, 200, 7, 101, 2, 2, 200, 201, 7,
	101, 2, 2, 201, 202, 7, 113, 2, 2, 202, 203, 7, 119, 2, 2, 203, 204, 7,
	112, 2, 2, 204, 205, 7, 118, 2, 2, 205, 52, 3, 2, 2, 2, 206, 207, 7, 99,
	2, 2, 207, 208, 7, 117, 2, 2, 208, 209, 7, 117, 2, 2, 209, 210, 7, 103,
	2, 2, 210, 211, 7, 118, 2, 2, 211, 54, 3, 2, 2, 2, 212, 213, 7, 112, 2,
	2, 213, 214, 7, 119, 2, 2, 214, 215, 7, 111, 2, 2, 215, 216, 7, 100, 2,
	2, 216, 217, 7, 103, 2, 2, 217, 218, 7, 116, 2, 2, 218, 56, 3, 2, 2, 2,
	219, 220, 7, 111, 2, 2, 220, 221, 7, 113, 2, 2, 221, 222, 7, 112, 2, 2,
	222, 223, 7, 103, 2, 2, 223, 224, 7, 118, 2, 2, 224, 225, 7, 99, 2, 2,
	225, 226, 7, 116, 2, 2, 226, 227, 7, 123, 2, 2, 227, 58, 3, 2, 2, 2, 228,
	229, 7, 114, 2, 2, 229, 230, 7, 113, 2, 2, 230, 231, 7, 116, 2, 2, 231,
	232, 7, 118, 2, 2, 232, 233, 7, 107, 2, 2, 233, 234, 7, 113, 2, 2, 234,
	235, 7, 112, 2, 2, 235, 60, 3, 2, 2, 2, 236, 240, 7, 36, 2, 2, 237, 239,
	9, 4, 2, 2, 238, 237, 3, 2, 2, 2, 239, 242, 3, 2, 2, 2, 240, 238, 3, 2,
	2, 2, 240, 241, 3, 2, 2, 2, 241, 243, 3, 2, 2, 2, 242, 240, 3, 2, 2, 2,
	243, 244, 7, 36, 2, 2, 244, 62, 3, 2, 2, 2, 245, 247, 9, 5, 2, 2, 246,
	245, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248, 246, 3, 2, 2, 2, 248, 249,
	3, 2, 2, 2, 249, 251, 3, 2, 2, 2, 250, 252, 9, 6, 2, 2, 251, 250, 3, 2,
	2, 2, 251, 252, 3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253, 255, 7, 49, 2, 2,
	254, 256, 9, 6, 2, 2, 255, 254, 3, 2, 2, 2, 255, 256, 3, 2, 2, 2, 256,
	258, 3, 2, 2, 2, 257, 259, 9, 5, 2, 2, 258, 257, 3, 2, 2, 2, 259, 260,
	3, 2, 2, 2, 260, 258, 3, 2, 2, 2, 260, 261, 3, 2, 2, 2, 261, 277, 3, 2,
	2, 2, 262, 264, 9, 5, 2, 2, 263, 262, 3, 2, 2, 2, 264, 265, 3, 2, 2, 2,
	265, 263, 3, 2, 2, 2, 265, 266, 3, 2, 2, 2, 266, 273, 3, 2, 2, 2, 267,
	269, 7, 48, 2, 2, 268, 270, 9, 5, 2, 2, 269, 268, 3, 2, 2, 2, 270, 271,
	3, 2, 2, 2, 271, 269, 3, 2, 2, 2, 271, 272, 3, 2, 2, 2, 272, 274, 3, 2,
	2, 2, 273, 267, 3, 2, 2, 2, 273, 274, 3, 2, 2, 2, 274, 275, 3, 2, 2, 2,
	275, 277, 7, 39, 2, 2, 276, 246, 3, 2, 2, 2, 276, 263, 3, 2, 2, 2, 277,
	64, 3, 2, 2, 2, 278, 279, 7, 116, 2, 2, 279, 280, 7, 103, 2, 2, 280, 281,
	7, 111, 2, 2, 281, 282, 7, 99, 2, 2, 282, 283, 7, 107, 2, 2, 283, 284,
	7, 112, 2, 2, 284, 285, 7, 107, 2, 2, 285, 286, 7, 112, 2, 2, 286, 287,
	7, 105, 2, 2, 287, 66, 3, 2, 2, 2, 288, 290, 9, 5, 2, 2, 289, 288, 3, 2,
	2, 2, 290, 291, 3, 2, 2, 2, 291, 289, 3, 2, 2, 2, 291, 292, 3, 2, 2, 2,
	292, 68, 3, 2, 2, 2, 293, 294, 7, 39, 2, 2, 294, 70, 3, 2, 2, 2, 295, 297,
	7, 38, 2, 2, 296, 298, 9, 7, 2, 2, 297, 296, 3, 2, 2, 2, 298, 299, 3, 2,
	2, 2, 299, 297, 3, 2, 2, 2, 299, 300, 3, 2, 2, 2, 300, 304, 3, 2, 2, 2,
	301, 303, 9, 8, 2, 2, 302, 301, 3, 2, 2, 2, 303, 306, 3, 2, 2, 2, 304,
	302, 3, 2, 2, 2, 304, 305, 3, 2, 2, 2, 305, 72, 3, 2, 2, 2, 306, 304, 3,
	2, 2, 2, 307, 309, 7, 66, 2, 2, 308, 310, 9, 7, 2, 2, 309, 308, 3, 2, 2,
	2, 310, 311, 3, 2, 2, 2, 311, 309, 3, 2, 2, 2, 311, 312, 3, 2, 2, 2, 312,
	316, 3, 2, 2, 2, 313, 315, 9, 9, 2, 2, 314, 313, 3, 2, 2, 2, 315, 318,
	3, 2, 2, 2, 316, 314, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317, 74, 3, 2,
	2, 2, 318, 316, 3, 2, 2, 2, 319, 321, 9, 10, 2, 2, 320, 319, 3, 2, 2, 2,
	321, 322, 3, 2, 2, 2, 322, 320, 3, 2, 2, 2, 322, 323, 3, 2, 2, 2, 323,
	76, 3, 2, 2, 2, 22, 2, 97, 106, 108, 122, 240, 248, 251, 255, 260, 265,
	271, 273, 276, 291, 299, 304, 311, 316, 322, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'to'", "'metadata('", "','", "", "", "", "", "'vars'", "'print'",
	"'fail'", "'send'", "'source'", "'destination'", "'allocate'", "'+'", "'-'",
	"'('", "')'", "'['", "']'", "'{'", "'}'", "'*'", "'='", "'account'", "'asset'",
	"'number'", "'monetary'", "'portion'", "", "", "'remaining'", "", "'%'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "NEWLINE", "WHITESPACE", "MULTILINE_COMMENT", "LINE_COMMENT",
	"VARS", "PRINT", "FAIL", "SEND", "SOURCE", "DESTINATION", "ALLOCATE", "OP_ADD",
	"OP_SUB", "LPAREN", "RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE", "ALL",
	"EQ", "TY_ACCOUNT", "TY_ASSET", "TY_NUMBER", "TY_MONETARY", "TY_PORTION",
	"STRING", "PORTION", "PORTION_REMAINING", "NUMBER", "PERCENT", "VARIABLE_NAME",
	"ACCOUNT", "ASSET",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "NEWLINE", "WHITESPACE", "MULTILINE_COMMENT", "LINE_COMMENT",
	"VARS", "PRINT", "FAIL", "SEND", "SOURCE", "DESTINATION", "ALLOCATE", "OP_ADD",
	"OP_SUB", "LPAREN", "RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE", "ALL",
	"EQ", "TY_ACCOUNT", "TY_ASSET", "TY_NUMBER", "TY_MONETARY", "TY_PORTION",
	"STRING", "PORTION", "PORTION_REMAINING", "NUMBER", "PERCENT", "VARIABLE_NAME",
	"ACCOUNT", "ASSET",
}

type NumScriptLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewNumScriptLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *NumScriptLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewNumScriptLexer(input antlr.CharStream) *NumScriptLexer {
	l := new(NumScriptLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "NumScript.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// NumScriptLexer tokens.
const (
	NumScriptLexerT__0              = 1
	NumScriptLexerT__1              = 2
	NumScriptLexerT__2              = 3
	NumScriptLexerNEWLINE           = 4
	NumScriptLexerWHITESPACE        = 5
	NumScriptLexerMULTILINE_COMMENT = 6
	NumScriptLexerLINE_COMMENT      = 7
	NumScriptLexerVARS              = 8
	NumScriptLexerPRINT             = 9
	NumScriptLexerFAIL              = 10
	NumScriptLexerSEND              = 11
	NumScriptLexerSOURCE            = 12
	NumScriptLexerDESTINATION       = 13
	NumScriptLexerALLOCATE          = 14
	NumScriptLexerOP_ADD            = 15
	NumScriptLexerOP_SUB            = 16
	NumScriptLexerLPAREN            = 17
	NumScriptLexerRPAREN            = 18
	NumScriptLexerLBRACK            = 19
	NumScriptLexerRBRACK            = 20
	NumScriptLexerLBRACE            = 21
	NumScriptLexerRBRACE            = 22
	NumScriptLexerALL               = 23
	NumScriptLexerEQ                = 24
	NumScriptLexerTY_ACCOUNT        = 25
	NumScriptLexerTY_ASSET          = 26
	NumScriptLexerTY_NUMBER         = 27
	NumScriptLexerTY_MONETARY       = 28
	NumScriptLexerTY_PORTION        = 29
	NumScriptLexerSTRING            = 30
	NumScriptLexerPORTION           = 31
	NumScriptLexerPORTION_REMAINING = 32
	NumScriptLexerNUMBER            = 33
	NumScriptLexerPERCENT           = 34
	NumScriptLexerVARIABLE_NAME     = 35
	NumScriptLexerACCOUNT           = 36
	NumScriptLexerASSET             = 37
)
