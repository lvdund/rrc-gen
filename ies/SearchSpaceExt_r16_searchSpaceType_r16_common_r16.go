package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpaceExt_r16_searchSpaceType_r16_common_r16 struct {
	dci_Format2_4_r16 *SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_4_r16 `optional`
	dci_Format2_5_r16 *SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16 `optional`
	dci_Format2_6_r16 interface{}                                                          `optional`
}

func (ie *SearchSpaceExt_r16_searchSpaceType_r16_common_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dci_Format2_4_r16 != nil, ie.dci_Format2_5_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dci_Format2_4_r16 != nil {
		if err = ie.dci_Format2_4_r16.Encode(w); err != nil {
			return utils.WrapError("Encode dci_Format2_4_r16", err)
		}
	}
	if ie.dci_Format2_5_r16 != nil {
		if err = ie.dci_Format2_5_r16.Encode(w); err != nil {
			return utils.WrapError("Encode dci_Format2_5_r16", err)
		}
	}
	return nil
}

func (ie *SearchSpaceExt_r16_searchSpaceType_r16_common_r16) Decode(r *uper.UperReader) error {
	var err error
	var dci_Format2_4_r16Present bool
	if dci_Format2_4_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dci_Format2_5_r16Present bool
	if dci_Format2_5_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if dci_Format2_4_r16Present {
		ie.dci_Format2_4_r16 = new(SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_4_r16)
		if err = ie.dci_Format2_4_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dci_Format2_4_r16", err)
		}
	}
	if dci_Format2_5_r16Present {
		ie.dci_Format2_5_r16 = new(SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_5_r16)
		if err = ie.dci_Format2_5_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dci_Format2_5_r16", err)
		}
	}
	return nil
}
