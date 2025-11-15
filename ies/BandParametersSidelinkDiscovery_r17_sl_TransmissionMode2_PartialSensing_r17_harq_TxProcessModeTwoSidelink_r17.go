package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandParametersSidelinkDiscovery_r17_sl_TransmissionMode2_PartialSensing_r17_harq_TxProcessModeTwoSidelink_r17_Enum_n8  uper.Enumerated = 0
	BandParametersSidelinkDiscovery_r17_sl_TransmissionMode2_PartialSensing_r17_harq_TxProcessModeTwoSidelink_r17_Enum_n16 uper.Enumerated = 1
)

type BandParametersSidelinkDiscovery_r17_sl_TransmissionMode2_PartialSensing_r17_harq_TxProcessModeTwoSidelink_r17 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *BandParametersSidelinkDiscovery_r17_sl_TransmissionMode2_PartialSensing_r17_harq_TxProcessModeTwoSidelink_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode BandParametersSidelinkDiscovery_r17_sl_TransmissionMode2_PartialSensing_r17_harq_TxProcessModeTwoSidelink_r17", err)
	}
	return nil
}

func (ie *BandParametersSidelinkDiscovery_r17_sl_TransmissionMode2_PartialSensing_r17_harq_TxProcessModeTwoSidelink_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode BandParametersSidelinkDiscovery_r17_sl_TransmissionMode2_PartialSensing_r17_harq_TxProcessModeTwoSidelink_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
