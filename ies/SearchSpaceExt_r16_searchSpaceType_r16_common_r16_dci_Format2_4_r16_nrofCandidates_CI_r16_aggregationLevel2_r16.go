package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_4_r16_nrofCandidates_CI_r16_aggregationLevel2_r16_Enum_n1 uper.Enumerated = 0
	SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_4_r16_nrofCandidates_CI_r16_aggregationLevel2_r16_Enum_n2 uper.Enumerated = 1
)

type SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_4_r16_nrofCandidates_CI_r16_aggregationLevel2_r16 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_4_r16_nrofCandidates_CI_r16_aggregationLevel2_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_4_r16_nrofCandidates_CI_r16_aggregationLevel2_r16", err)
	}
	return nil
}

func (ie *SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_4_r16_nrofCandidates_CI_r16_aggregationLevel2_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SearchSpaceExt_r16_searchSpaceType_r16_common_r16_dci_Format2_4_r16_nrofCandidates_CI_r16_aggregationLevel2_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
