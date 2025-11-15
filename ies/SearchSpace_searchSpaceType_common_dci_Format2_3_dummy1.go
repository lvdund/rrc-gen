package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1_Enum_sl1  uper.Enumerated = 0
	SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1_Enum_sl2  uper.Enumerated = 1
	SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1_Enum_sl4  uper.Enumerated = 2
	SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1_Enum_sl5  uper.Enumerated = 3
	SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1_Enum_sl8  uper.Enumerated = 4
	SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1_Enum_sl10 uper.Enumerated = 5
	SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1_Enum_sl16 uper.Enumerated = 6
	SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1_Enum_sl20 uper.Enumerated = 7
)

type SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1", err)
	}
	return nil
}

func (ie *SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SearchSpace_searchSpaceType_common_dci_Format2_3_dummy1", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
