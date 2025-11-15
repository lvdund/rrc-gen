package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel8_Enum_n1 uper.Enumerated = 0
	SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel8_Enum_n2 uper.Enumerated = 1
)

type SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel8 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel8) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel8", err)
	}
	return nil
}

func (ie *SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel8) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SearchSpace_searchSpaceType_common_dci_Format2_0_nrofCandidates_SFI_aggregationLevel8", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
