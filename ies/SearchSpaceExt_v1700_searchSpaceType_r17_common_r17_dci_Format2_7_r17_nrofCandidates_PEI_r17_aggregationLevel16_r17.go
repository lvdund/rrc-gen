package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel16_r17_Enum_n0 uper.Enumerated = 0
	SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel16_r17_Enum_n1 uper.Enumerated = 1
)

type SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel16_r17 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel16_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel16_r17", err)
	}
	return nil
}

func (ie *SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel16_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SearchSpaceExt_v1700_searchSpaceType_r17_common_r17_dci_Format2_7_r17_nrofCandidates_PEI_r17_aggregationLevel16_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
