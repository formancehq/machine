// Code generated from NumScript.g4 by ANTLR 4.9.3. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 43, 358,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 4, 6, 4, 91, 10, 4, 13, 4, 14, 4, 92, 3, 5, 6, 5, 96, 10, 5, 13,
	5, 14, 5, 97, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 107, 10,
	6, 12, 6, 14, 6, 110, 11, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 7, 3, 7, 7, 7, 121, 10, 7, 12, 7, 14, 7, 124, 11, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22,
	3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3,
	28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 7, 35, 273, 10, 35, 12, 35,
	14, 35, 276, 11, 35, 3, 35, 3, 35, 3, 36, 6, 36, 281, 10, 36, 13, 36, 14,
	36, 282, 3, 36, 5, 36, 286, 10, 36, 3, 36, 3, 36, 5, 36, 290, 10, 36, 3,
	36, 6, 36, 293, 10, 36, 13, 36, 14, 36, 294, 3, 36, 6, 36, 298, 10, 36,
	13, 36, 14, 36, 299, 3, 36, 3, 36, 6, 36, 304, 10, 36, 13, 36, 14, 36,
	305, 5, 36, 308, 10, 36, 3, 36, 5, 36, 311, 10, 36, 3, 37, 3, 37, 3, 37,
	3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 6, 38, 324, 10,
	38, 13, 38, 14, 38, 325, 3, 39, 3, 39, 3, 40, 3, 40, 6, 40, 332, 10, 40,
	13, 40, 14, 40, 333, 3, 40, 7, 40, 337, 10, 40, 12, 40, 14, 40, 340, 11,
	40, 3, 41, 3, 41, 6, 41, 344, 10, 41, 13, 41, 14, 41, 345, 3, 41, 7, 41,
	349, 10, 41, 12, 41, 14, 41, 352, 11, 41, 3, 42, 6, 42, 355, 10, 42, 13,
	42, 14, 42, 356, 4, 108, 122, 2, 43, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26,
	51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35,
	69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 3, 2, 11,
	4, 2, 12, 12, 15, 15, 4, 2, 11, 11, 34, 34, 8, 2, 34, 34, 47, 47, 50, 59,
	67, 92, 97, 97, 99, 124, 3, 2, 50, 59, 3, 2, 34, 34, 4, 2, 97, 97, 99,
	124, 5, 2, 50, 59, 97, 97, 99, 124, 5, 2, 50, 60, 97, 97, 99, 124, 4, 2,
	49, 59, 67, 92, 2, 377, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2,
	2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3,
	2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23,
	3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2,
	31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2,
	2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2,
	2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2,
	2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3,
	2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69,
	3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2,
	77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2,
	3, 85, 3, 2, 2, 2, 5, 87, 3, 2, 2, 2, 7, 90, 3, 2, 2, 2, 9, 95, 3, 2, 2,
	2, 11, 101, 3, 2, 2, 2, 13, 116, 3, 2, 2, 2, 15, 129, 3, 2, 2, 2, 17, 134,
	3, 2, 2, 2, 19, 139, 3, 2, 2, 2, 21, 151, 3, 2, 2, 2, 23, 157, 3, 2, 2,
	2, 25, 162, 3, 2, 2, 2, 27, 167, 3, 2, 2, 2, 29, 174, 3, 2, 2, 2, 31, 179,
	3, 2, 2, 2, 33, 183, 3, 2, 2, 2, 35, 195, 3, 2, 2, 2, 37, 198, 3, 2, 2,
	2, 39, 207, 3, 2, 2, 2, 41, 209, 3, 2, 2, 2, 43, 211, 3, 2, 2, 2, 45, 213,
	3, 2, 2, 2, 47, 215, 3, 2, 2, 2, 49, 217, 3, 2, 2, 2, 51, 219, 3, 2, 2,
	2, 53, 221, 3, 2, 2, 2, 55, 223, 3, 2, 2, 2, 57, 225, 3, 2, 2, 2, 59, 233,
	3, 2, 2, 2, 61, 239, 3, 2, 2, 2, 63, 246, 3, 2, 2, 2, 65, 255, 3, 2, 2,
	2, 67, 263, 3, 2, 2, 2, 69, 270, 3, 2, 2, 2, 71, 310, 3, 2, 2, 2, 73, 312,
	3, 2, 2, 2, 75, 323, 3, 2, 2, 2, 77, 327, 3, 2, 2, 2, 79, 329, 3, 2, 2,
	2, 81, 341, 3, 2, 2, 2, 83, 354, 3, 2, 2, 2, 85, 86, 7, 44, 2, 2, 86, 4,
	3, 2, 2, 2, 87, 88, 7, 46, 2, 2, 88, 6, 3, 2, 2, 2, 89, 91, 9, 2, 2, 2,
	90, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 92, 93, 3,
	2, 2, 2, 93, 8, 3, 2, 2, 2, 94, 96, 9, 3, 2, 2, 95, 94, 3, 2, 2, 2, 96,
	97, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 99, 3, 2, 2,
	2, 99, 100, 8, 5, 2, 2, 100, 10, 3, 2, 2, 2, 101, 102, 7, 49, 2, 2, 102,
	103, 7, 44, 2, 2, 103, 108, 3, 2, 2, 2, 104, 107, 5, 11, 6, 2, 105, 107,
	11, 2, 2, 2, 106, 104, 3, 2, 2, 2, 106, 105, 3, 2, 2, 2, 107, 110, 3, 2,
	2, 2, 108, 109, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 109, 111, 3, 2, 2, 2,
	110, 108, 3, 2, 2, 2, 111, 112, 7, 44, 2, 2, 112, 113, 7, 49, 2, 2, 113,
	114, 3, 2, 2, 2, 114, 115, 8, 6, 2, 2, 115, 12, 3, 2, 2, 2, 116, 117, 7,
	49, 2, 2, 117, 118, 7, 49, 2, 2, 118, 122, 3, 2, 2, 2, 119, 121, 11, 2,
	2, 2, 120, 119, 3, 2, 2, 2, 121, 124, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2,
	122, 120, 3, 2, 2, 2, 123, 125, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 125,
	126, 5, 7, 4, 2, 126, 127, 3, 2, 2, 2, 127, 128, 8, 7, 2, 2, 128, 14, 3,
	2, 2, 2, 129, 130, 7, 120, 2, 2, 130, 131, 7, 99, 2, 2, 131, 132, 7, 116,
	2, 2, 132, 133, 7, 117, 2, 2, 133, 16, 3, 2, 2, 2, 134, 135, 7, 111, 2,
	2, 135, 136, 7, 103, 2, 2, 136, 137, 7, 118, 2, 2, 137, 138, 7, 99, 2,
	2, 138, 18, 3, 2, 2, 2, 139, 140, 7, 117, 2, 2, 140, 141, 7, 103, 2, 2,
	141, 142, 7, 118, 2, 2, 142, 143, 7, 97, 2, 2, 143, 144, 7, 118, 2, 2,
	144, 145, 7, 122, 2, 2, 145, 146, 7, 97, 2, 2, 146, 147, 7, 111, 2, 2,
	147, 148, 7, 103, 2, 2, 148, 149, 7, 118, 2, 2, 149, 150, 7, 99, 2, 2,
	150, 20, 3, 2, 2, 2, 151, 152, 7, 114, 2, 2, 152, 153, 7, 116, 2, 2, 153,
	154, 7, 107, 2, 2, 154, 155, 7, 112, 2, 2, 155, 156, 7, 118, 2, 2, 156,
	22, 3, 2, 2, 2, 157, 158, 7, 104, 2, 2, 158, 159, 7, 99, 2, 2, 159, 160,
	7, 107, 2, 2, 160, 161, 7, 110, 2, 2, 161, 24, 3, 2, 2, 2, 162, 163, 7,
	117, 2, 2, 163, 164, 7, 103, 2, 2, 164, 165, 7, 112, 2, 2, 165, 166, 7,
	102, 2, 2, 166, 26, 3, 2, 2, 2, 167, 168, 7, 117, 2, 2, 168, 169, 7, 113,
	2, 2, 169, 170, 7, 119, 2, 2, 170, 171, 7, 116, 2, 2, 171, 172, 7, 101,
	2, 2, 172, 173, 7, 103, 2, 2, 173, 28, 3, 2, 2, 2, 174, 175, 7, 104, 2,
	2, 175, 176, 7, 116, 2, 2, 176, 177, 7, 113, 2, 2, 177, 178, 7, 111, 2,
	2, 178, 30, 3, 2, 2, 2, 179, 180, 7, 111, 2, 2, 180, 181, 7, 99, 2, 2,
	181, 182, 7, 122, 2, 2, 182, 32, 3, 2, 2, 2, 183, 184, 7, 102, 2, 2, 184,
	185, 7, 103, 2, 2, 185, 186, 7, 117, 2, 2, 186, 187, 7, 118, 2, 2, 187,
	188, 7, 107, 2, 2, 188, 189, 7, 112, 2, 2, 189, 190, 7, 99, 2, 2, 190,
	191, 7, 118, 2, 2, 191, 192, 7, 107, 2, 2, 192, 193, 7, 113, 2, 2, 193,
	194, 7, 112, 2, 2, 194, 34, 3, 2, 2, 2, 195, 196, 7, 118, 2, 2, 196, 197,
	7, 113, 2, 2, 197, 36, 3, 2, 2, 2, 198, 199, 7, 99, 2, 2, 199, 200, 7,
	110, 2, 2, 200, 201, 7, 110, 2, 2, 201, 202, 7, 113, 2, 2, 202, 203, 7,
	101, 2, 2, 203, 204, 7, 99, 2, 2, 204, 205, 7, 118, 2, 2, 205, 206, 7,
	103, 2, 2, 206, 38, 3, 2, 2, 2, 207, 208, 7, 45, 2, 2, 208, 40, 3, 2, 2,
	2, 209, 210, 7, 47, 2, 2, 210, 42, 3, 2, 2, 2, 211, 212, 7, 42, 2, 2, 212,
	44, 3, 2, 2, 2, 213, 214, 7, 43, 2, 2, 214, 46, 3, 2, 2, 2, 215, 216, 7,
	93, 2, 2, 216, 48, 3, 2, 2, 2, 217, 218, 7, 95, 2, 2, 218, 50, 3, 2, 2,
	2, 219, 220, 7, 125, 2, 2, 220, 52, 3, 2, 2, 2, 221, 222, 7, 127, 2, 2,
	222, 54, 3, 2, 2, 2, 223, 224, 7, 63, 2, 2, 224, 56, 3, 2, 2, 2, 225, 226,
	7, 99, 2, 2, 226, 227, 7, 101, 2, 2, 227, 228, 7, 101, 2, 2, 228, 229,
	7, 113, 2, 2, 229, 230, 7, 119, 2, 2, 230, 231, 7, 112, 2, 2, 231, 232,
	7, 118, 2, 2, 232, 58, 3, 2, 2, 2, 233, 234, 7, 99, 2, 2, 234, 235, 7,
	117, 2, 2, 235, 236, 7, 117, 2, 2, 236, 237, 7, 103, 2, 2, 237, 238, 7,
	118, 2, 2, 238, 60, 3, 2, 2, 2, 239, 240, 7, 112, 2, 2, 240, 241, 7, 119,
	2, 2, 241, 242, 7, 111, 2, 2, 242, 243, 7, 100, 2, 2, 243, 244, 7, 103,
	2, 2, 244, 245, 7, 116, 2, 2, 245, 62, 3, 2, 2, 2, 246, 247, 7, 111, 2,
	2, 247, 248, 7, 113, 2, 2, 248, 249, 7, 112, 2, 2, 249, 250, 7, 103, 2,
	2, 250, 251, 7, 118, 2, 2, 251, 252, 7, 99, 2, 2, 252, 253, 7, 116, 2,
	2, 253, 254, 7, 123, 2, 2, 254, 64, 3, 2, 2, 2, 255, 256, 7, 114, 2, 2,
	256, 257, 7, 113, 2, 2, 257, 258, 7, 116, 2, 2, 258, 259, 7, 118, 2, 2,
	259, 260, 7, 107, 2, 2, 260, 261, 7, 113, 2, 2, 261, 262, 7, 112, 2, 2,
	262, 66, 3, 2, 2, 2, 263, 264, 7, 117, 2, 2, 264, 265, 7, 118, 2, 2, 265,
	266, 7, 116, 2, 2, 266, 267, 7, 107, 2, 2, 267, 268, 7, 112, 2, 2, 268,
	269, 7, 105, 2, 2, 269, 68, 3, 2, 2, 2, 270, 274, 7, 36, 2, 2, 271, 273,
	9, 4, 2, 2, 272, 271, 3, 2, 2, 2, 273, 276, 3, 2, 2, 2, 274, 272, 3, 2,
	2, 2, 274, 275, 3, 2, 2, 2, 275, 277, 3, 2, 2, 2, 276, 274, 3, 2, 2, 2,
	277, 278, 7, 36, 2, 2, 278, 70, 3, 2, 2, 2, 279, 281, 9, 5, 2, 2, 280,
	279, 3, 2, 2, 2, 281, 282, 3, 2, 2, 2, 282, 280, 3, 2, 2, 2, 282, 283,
	3, 2, 2, 2, 283, 285, 3, 2, 2, 2, 284, 286, 9, 6, 2, 2, 285, 284, 3, 2,
	2, 2, 285, 286, 3, 2, 2, 2, 286, 287, 3, 2, 2, 2, 287, 289, 7, 49, 2, 2,
	288, 290, 9, 6, 2, 2, 289, 288, 3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290,
	292, 3, 2, 2, 2, 291, 293, 9, 5, 2, 2, 292, 291, 3, 2, 2, 2, 293, 294,
	3, 2, 2, 2, 294, 292, 3, 2, 2, 2, 294, 295, 3, 2, 2, 2, 295, 311, 3, 2,
	2, 2, 296, 298, 9, 5, 2, 2, 297, 296, 3, 2, 2, 2, 298, 299, 3, 2, 2, 2,
	299, 297, 3, 2, 2, 2, 299, 300, 3, 2, 2, 2, 300, 307, 3, 2, 2, 2, 301,
	303, 7, 48, 2, 2, 302, 304, 9, 5, 2, 2, 303, 302, 3, 2, 2, 2, 304, 305,
	3, 2, 2, 2, 305, 303, 3, 2, 2, 2, 305, 306, 3, 2, 2, 2, 306, 308, 3, 2,
	2, 2, 307, 301, 3, 2, 2, 2, 307, 308, 3, 2, 2, 2, 308, 309, 3, 2, 2, 2,
	309, 311, 7, 39, 2, 2, 310, 280, 3, 2, 2, 2, 310, 297, 3, 2, 2, 2, 311,
	72, 3, 2, 2, 2, 312, 313, 7, 116, 2, 2, 313, 314, 7, 103, 2, 2, 314, 315,
	7, 111, 2, 2, 315, 316, 7, 99, 2, 2, 316, 317, 7, 107, 2, 2, 317, 318,
	7, 112, 2, 2, 318, 319, 7, 107, 2, 2, 319, 320, 7, 112, 2, 2, 320, 321,
	7, 105, 2, 2, 321, 74, 3, 2, 2, 2, 322, 324, 9, 5, 2, 2, 323, 322, 3, 2,
	2, 2, 324, 325, 3, 2, 2, 2, 325, 323, 3, 2, 2, 2, 325, 326, 3, 2, 2, 2,
	326, 76, 3, 2, 2, 2, 327, 328, 7, 39, 2, 2, 328, 78, 3, 2, 2, 2, 329, 331,
	7, 38, 2, 2, 330, 332, 9, 7, 2, 2, 331, 330, 3, 2, 2, 2, 332, 333, 3, 2,
	2, 2, 333, 331, 3, 2, 2, 2, 333, 334, 3, 2, 2, 2, 334, 338, 3, 2, 2, 2,
	335, 337, 9, 8, 2, 2, 336, 335, 3, 2, 2, 2, 337, 340, 3, 2, 2, 2, 338,
	336, 3, 2, 2, 2, 338, 339, 3, 2, 2, 2, 339, 80, 3, 2, 2, 2, 340, 338, 3,
	2, 2, 2, 341, 343, 7, 66, 2, 2, 342, 344, 9, 7, 2, 2, 343, 342, 3, 2, 2,
	2, 344, 345, 3, 2, 2, 2, 345, 343, 3, 2, 2, 2, 345, 346, 3, 2, 2, 2, 346,
	350, 3, 2, 2, 2, 347, 349, 9, 9, 2, 2, 348, 347, 3, 2, 2, 2, 349, 352,
	3, 2, 2, 2, 350, 348, 3, 2, 2, 2, 350, 351, 3, 2, 2, 2, 351, 82, 3, 2,
	2, 2, 352, 350, 3, 2, 2, 2, 353, 355, 9, 10, 2, 2, 354, 353, 3, 2, 2, 2,
	355, 356, 3, 2, 2, 2, 356, 354, 3, 2, 2, 2, 356, 357, 3, 2, 2, 2, 357,
	84, 3, 2, 2, 2, 23, 2, 92, 97, 106, 108, 122, 274, 282, 285, 289, 294,
	299, 305, 307, 310, 325, 333, 338, 345, 350, 356, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'*'", "','", "", "", "", "", "'vars'", "'meta'", "'set_tx_meta'",
	"'print'", "'fail'", "'send'", "'source'", "'from'", "'max'", "'destination'",
	"'to'", "'allocate'", "'+'", "'-'", "'('", "')'", "'['", "']'", "'{'",
	"'}'", "'='", "'account'", "'asset'", "'number'", "'monetary'", "'portion'",
	"'string'", "", "", "'remaining'", "", "'%'",
}

var lexerSymbolicNames = []string{
	"", "", "", "NEWLINE", "WHITESPACE", "MULTILINE_COMMENT", "LINE_COMMENT",
	"VARS", "META", "SET_TX_META", "PRINT", "FAIL", "SEND", "SOURCE", "FROM",
	"MAX", "DESTINATION", "TO", "ALLOCATE", "OP_ADD", "OP_SUB", "LPAREN", "RPAREN",
	"LBRACK", "RBRACK", "LBRACE", "RBRACE", "EQ", "TY_ACCOUNT", "TY_ASSET",
	"TY_NUMBER", "TY_MONETARY", "TY_PORTION", "TY_STRING", "STRING", "PORTION",
	"PORTION_REMAINING", "NUMBER", "PERCENT", "VARIABLE_NAME", "ACCOUNT", "ASSET",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "NEWLINE", "WHITESPACE", "MULTILINE_COMMENT", "LINE_COMMENT",
	"VARS", "META", "SET_TX_META", "PRINT", "FAIL", "SEND", "SOURCE", "FROM",
	"MAX", "DESTINATION", "TO", "ALLOCATE", "OP_ADD", "OP_SUB", "LPAREN", "RPAREN",
	"LBRACK", "RBRACK", "LBRACE", "RBRACE", "EQ", "TY_ACCOUNT", "TY_ASSET",
	"TY_NUMBER", "TY_MONETARY", "TY_PORTION", "TY_STRING", "STRING", "PORTION",
	"PORTION_REMAINING", "NUMBER", "PERCENT", "VARIABLE_NAME", "ACCOUNT", "ASSET",
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
	NumScriptLexerNEWLINE           = 3
	NumScriptLexerWHITESPACE        = 4
	NumScriptLexerMULTILINE_COMMENT = 5
	NumScriptLexerLINE_COMMENT      = 6
	NumScriptLexerVARS              = 7
	NumScriptLexerMETA              = 8
	NumScriptLexerSET_TX_META       = 9
	NumScriptLexerPRINT             = 10
	NumScriptLexerFAIL              = 11
	NumScriptLexerSEND              = 12
	NumScriptLexerSOURCE            = 13
	NumScriptLexerFROM              = 14
	NumScriptLexerMAX               = 15
	NumScriptLexerDESTINATION       = 16
	NumScriptLexerTO                = 17
	NumScriptLexerALLOCATE          = 18
	NumScriptLexerOP_ADD            = 19
	NumScriptLexerOP_SUB            = 20
	NumScriptLexerLPAREN            = 21
	NumScriptLexerRPAREN            = 22
	NumScriptLexerLBRACK            = 23
	NumScriptLexerRBRACK            = 24
	NumScriptLexerLBRACE            = 25
	NumScriptLexerRBRACE            = 26
	NumScriptLexerEQ                = 27
	NumScriptLexerTY_ACCOUNT        = 28
	NumScriptLexerTY_ASSET          = 29
	NumScriptLexerTY_NUMBER         = 30
	NumScriptLexerTY_MONETARY       = 31
	NumScriptLexerTY_PORTION        = 32
	NumScriptLexerTY_STRING         = 33
	NumScriptLexerSTRING            = 34
	NumScriptLexerPORTION           = 35
	NumScriptLexerPORTION_REMAINING = 36
	NumScriptLexerNUMBER            = 37
	NumScriptLexerPERCENT           = 38
	NumScriptLexerVARIABLE_NAME     = 39
	NumScriptLexerACCOUNT           = 40
	NumScriptLexerASSET             = 41
)
