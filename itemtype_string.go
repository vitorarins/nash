// Code generated by "stringer -type=itemType"; DO NOT EDIT

package nash

import "fmt"

const _itemType_name = "itemErroritemEOFitemImportitemCommentitemSetEnvitemShowEnvitemVarNameitemConcatitemVariableitemListOpenitemListCloseitemListElemitemCommanditemArgitemLeftBlockitemRightBlockitemStringitemRedirRightitemRedirRBracketitemRedirLBracketitemRedirFileitemRedirNetAddritemRedirMapEqualitemRedirMapLSideitemRedirMapRSideitemIfitemElseitemComparisonitemRforkitemRforkFlagsitemCd"

var _itemType_index = [...]uint16{0, 9, 16, 26, 37, 47, 58, 69, 79, 91, 103, 116, 128, 139, 146, 159, 173, 183, 197, 214, 231, 244, 260, 277, 294, 311, 317, 325, 339, 348, 362, 368}

func (i itemType) String() string {
	i -= 2
	if i < 0 || i >= itemType(len(_itemType_index)-1) {
		return fmt.Sprintf("itemType(%d)", i+2)
	}
	return _itemType_name[_itemType_index[i]:_itemType_index[i+1]]
}
