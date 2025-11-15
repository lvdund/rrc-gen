package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpaceExt_v1700_searchSpaceType_r17 struct {
	common_r17 *SearchSpaceExt_v1700_searchSpaceType_r17_common_r17 `optional`
}

func (ie *SearchSpaceExt_v1700_searchSpaceType_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.common_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.common_r17 != nil {
		if err = ie.common_r17.Encode(w); err != nil {
			return utils.WrapError("Encode common_r17", err)
		}
	}
	return nil
}

func (ie *SearchSpaceExt_v1700_searchSpaceType_r17) Decode(r *uper.UperReader) error {
	var err error
	var common_r17Present bool
	if common_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if common_r17Present {
		ie.common_r17 = new(SearchSpaceExt_v1700_searchSpaceType_r17_common_r17)
		if err = ie.common_r17.Decode(r); err != nil {
			return utils.WrapError("Decode common_r17", err)
		}
	}
	return nil
}
