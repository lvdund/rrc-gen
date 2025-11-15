package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpaceExt_r16_searchSpaceType_r16 struct {
	common_r16 *SearchSpaceExt_r16_searchSpaceType_r16_common_r16 `optional`
}

func (ie *SearchSpaceExt_r16_searchSpaceType_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.common_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.common_r16 != nil {
		if err = ie.common_r16.Encode(w); err != nil {
			return utils.WrapError("Encode common_r16", err)
		}
	}
	return nil
}

func (ie *SearchSpaceExt_r16_searchSpaceType_r16) Decode(r *uper.UperReader) error {
	var err error
	var common_r16Present bool
	if common_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if common_r16Present {
		ie.common_r16 = new(SearchSpaceExt_r16_searchSpaceType_r16_common_r16)
		if err = ie.common_r16.Decode(r); err != nil {
			return utils.WrapError("Decode common_r16", err)
		}
	}
	return nil
}
