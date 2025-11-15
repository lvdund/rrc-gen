package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SearchSpace_searchSpaceType_ue_Specific_dci_Formats_Enum_formats0_0_And_1_0 uper.Enumerated = 0
	SearchSpace_searchSpaceType_ue_Specific_dci_Formats_Enum_formats0_1_And_1_1 uper.Enumerated = 1
)

type SearchSpace_searchSpaceType_ue_Specific_dci_Formats struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SearchSpace_searchSpaceType_ue_Specific_dci_Formats) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SearchSpace_searchSpaceType_ue_Specific_dci_Formats", err)
	}
	return nil
}

func (ie *SearchSpace_searchSpaceType_ue_Specific_dci_Formats) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SearchSpace_searchSpaceType_ue_Specific_dci_Formats", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
