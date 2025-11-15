package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpaceExt_v1700_searchSpaceType_r17_common_r17 struct {
	dci_Format4_0_r17              interface{}                                                            `optional`
	dci_Format4_1_r17              interface{}                                                            `optional`
	dci_Format4_2_r17              interface{}                                                            `optional`
	dci_Format4_1_AndFormat4_2_r17 interface{}                                                            `optional`
	dci_Format2_7_r17              *SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17 `optional`
}

func (ie *SearchSpaceExt_v1700_searchSpaceType_r17_common_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dci_Format2_7_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dci_Format2_7_r17 != nil {
		if err = ie.dci_Format2_7_r17.Encode(w); err != nil {
			return utils.WrapError("Encode dci_Format2_7_r17", err)
		}
	}
	return nil
}

func (ie *SearchSpaceExt_v1700_searchSpaceType_r17_common_r17) Decode(r *uper.UperReader) error {
	var err error
	var dci_Format2_7_r17Present bool
	if dci_Format2_7_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if dci_Format2_7_r17Present {
		ie.dci_Format2_7_r17 = new(SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17)
		if err = ie.dci_Format2_7_r17.Decode(r); err != nil {
			return utils.WrapError("Decode dci_Format2_7_r17", err)
		}
	}
	return nil
}
