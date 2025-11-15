package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB2_relaxedMeasurement_r16_cellEdgeEvaluation_r16 struct {
	s_SearchThresholdP_r16 ReselectionThreshold   `madatory`
	s_SearchThresholdQ_r16 *ReselectionThresholdQ `optional`
}

func (ie *SIB2_relaxedMeasurement_r16_cellEdgeEvaluation_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.s_SearchThresholdQ_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.s_SearchThresholdP_r16.Encode(w); err != nil {
		return utils.WrapError("Encode s_SearchThresholdP_r16", err)
	}
	if ie.s_SearchThresholdQ_r16 != nil {
		if err = ie.s_SearchThresholdQ_r16.Encode(w); err != nil {
			return utils.WrapError("Encode s_SearchThresholdQ_r16", err)
		}
	}
	return nil
}

func (ie *SIB2_relaxedMeasurement_r16_cellEdgeEvaluation_r16) Decode(r *uper.UperReader) error {
	var err error
	var s_SearchThresholdQ_r16Present bool
	if s_SearchThresholdQ_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.s_SearchThresholdP_r16.Decode(r); err != nil {
		return utils.WrapError("Decode s_SearchThresholdP_r16", err)
	}
	if s_SearchThresholdQ_r16Present {
		ie.s_SearchThresholdQ_r16 = new(ReselectionThresholdQ)
		if err = ie.s_SearchThresholdQ_r16.Decode(r); err != nil {
			return utils.WrapError("Decode s_SearchThresholdQ_r16", err)
		}
	}
	return nil
}
