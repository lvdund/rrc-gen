package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17 struct {
	nrofCandidates_PEI_r17 *SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17 `optional`
}

func (ie *SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.nrofCandidates_PEI_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.nrofCandidates_PEI_r17 != nil {
		if err = ie.nrofCandidates_PEI_r17.Encode(w); err != nil {
			return utils.WrapError("Encode nrofCandidates_PEI_r17", err)
		}
	}
	return nil
}

func (ie *SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17) Decode(r *uper.UperReader) error {
	var err error
	var nrofCandidates_PEI_r17Present bool
	if nrofCandidates_PEI_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if nrofCandidates_PEI_r17Present {
		ie.nrofCandidates_PEI_r17 = new(SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17)
		if err = ie.nrofCandidates_PEI_r17.Decode(r); err != nil {
			return utils.WrapError("Decode nrofCandidates_PEI_r17", err)
		}
	}
	return nil
}
