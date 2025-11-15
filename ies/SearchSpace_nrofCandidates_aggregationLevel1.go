package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SearchSpace_nrofCandidates_aggregationLevel1_Enum_n0 uper.Enumerated = 0
	SearchSpace_nrofCandidates_aggregationLevel1_Enum_n1 uper.Enumerated = 1
	SearchSpace_nrofCandidates_aggregationLevel1_Enum_n2 uper.Enumerated = 2
	SearchSpace_nrofCandidates_aggregationLevel1_Enum_n3 uper.Enumerated = 3
	SearchSpace_nrofCandidates_aggregationLevel1_Enum_n4 uper.Enumerated = 4
	SearchSpace_nrofCandidates_aggregationLevel1_Enum_n5 uper.Enumerated = 5
	SearchSpace_nrofCandidates_aggregationLevel1_Enum_n6 uper.Enumerated = 6
	SearchSpace_nrofCandidates_aggregationLevel1_Enum_n8 uper.Enumerated = 7
)

type SearchSpace_nrofCandidates_aggregationLevel1 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SearchSpace_nrofCandidates_aggregationLevel1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SearchSpace_nrofCandidates_aggregationLevel1", err)
	}
	return nil
}

func (ie *SearchSpace_nrofCandidates_aggregationLevel1) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SearchSpace_nrofCandidates_aggregationLevel1", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
