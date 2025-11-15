package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC_Enum_eutra_nr      uper.Enumerated = 0
	AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC_Enum_nr            uper.Enumerated = 1
	AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC_Enum_other         uper.Enumerated = 2
	AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC_Enum_utra_nr_other uper.Enumerated = 3
	AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC_Enum_nr_other      uper.Enumerated = 4
	AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC_Enum_spare3        uper.Enumerated = 5
	AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC_Enum_spare2        uper.Enumerated = 6
	AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC_Enum_spare1        uper.Enumerated = 7
)

type AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC", err)
	}
	return nil
}

func (ie *AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
