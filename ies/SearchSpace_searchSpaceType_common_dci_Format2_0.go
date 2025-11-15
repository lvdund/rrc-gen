package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpace_searchSpaceType_common_dci_Format2_0 struct {
	nrofCandidates_SFI *SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI `optional`
}

func (ie *SearchSpace_searchSpaceType_common_dci_Format2_0) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.nrofCandidates_SFI != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.nrofCandidates_SFI != nil {
		if err = ie.nrofCandidates_SFI.Encode(w); err != nil {
			return utils.WrapError("Encode nrofCandidates_SFI", err)
		}
	}
	return nil
}

func (ie *SearchSpace_searchSpaceType_common_dci_Format2_0) Decode(r *uper.UperReader) error {
	var err error
	var nrofCandidates_SFIPresent bool
	if nrofCandidates_SFIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if nrofCandidates_SFIPresent {
		ie.nrofCandidates_SFI = new(SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI)
		if err = ie.nrofCandidates_SFI.Decode(r); err != nil {
			return utils.WrapError("Decode nrofCandidates_SFI", err)
		}
	}
	return nil
}
