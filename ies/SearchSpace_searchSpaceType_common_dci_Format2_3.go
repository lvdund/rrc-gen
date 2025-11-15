package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpace_searchSpaceType_common_dci_Format2_3 struct {
	dummy1 *SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1 `optional`
	dummy2 SearchSpace_searchSpaceType_common_dci_Format2_3_dummy2  `madatory`
}

func (ie *SearchSpace_searchSpaceType_common_dci_Format2_3) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dummy1 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dummy1 != nil {
		if err = ie.dummy1.Encode(w); err != nil {
			return utils.WrapError("Encode dummy1", err)
		}
	}
	if err = ie.dummy2.Encode(w); err != nil {
		return utils.WrapError("Encode dummy2", err)
	}
	return nil
}

func (ie *SearchSpace_searchSpaceType_common_dci_Format2_3) Decode(r *uper.UperReader) error {
	var err error
	var dummy1Present bool
	if dummy1Present, err = r.ReadBool(); err != nil {
		return err
	}
	if dummy1Present {
		ie.dummy1 = new(SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1)
		if err = ie.dummy1.Decode(r); err != nil {
			return utils.WrapError("Decode dummy1", err)
		}
	}
	if err = ie.dummy2.Decode(r); err != nil {
		return utils.WrapError("Decode dummy2", err)
	}
	return nil
}
