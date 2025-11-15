package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type AffectedCarrierFreqCombInfoMRDC struct {
	victimSystemType            VictimSystemType                                             `madatory`
	interferenceDirectionMRDC   AffectedCarrierFreqCombInfoMRDC_interferenceDirectionMRDC    `madatory`
	affectedCarrierFreqCombMRDC *AffectedCarrierFreqCombInfoMRDC_affectedCarrierFreqCombMRDC `optional`
}

func (ie *AffectedCarrierFreqCombInfoMRDC) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.affectedCarrierFreqCombMRDC != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.victimSystemType.Encode(w); err != nil {
		return utils.WrapError("Encode victimSystemType", err)
	}
	if err = ie.interferenceDirectionMRDC.Encode(w); err != nil {
		return utils.WrapError("Encode interferenceDirectionMRDC", err)
	}
	if ie.affectedCarrierFreqCombMRDC != nil {
		if err = ie.affectedCarrierFreqCombMRDC.Encode(w); err != nil {
			return utils.WrapError("Encode affectedCarrierFreqCombMRDC", err)
		}
	}
	return nil
}

func (ie *AffectedCarrierFreqCombInfoMRDC) Decode(r *uper.UperReader) error {
	var err error
	var affectedCarrierFreqCombMRDCPresent bool
	if affectedCarrierFreqCombMRDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.victimSystemType.Decode(r); err != nil {
		return utils.WrapError("Decode victimSystemType", err)
	}
	if err = ie.interferenceDirectionMRDC.Decode(r); err != nil {
		return utils.WrapError("Decode interferenceDirectionMRDC", err)
	}
	if affectedCarrierFreqCombMRDCPresent {
		ie.affectedCarrierFreqCombMRDC = new(AffectedCarrierFreqCombInfoMRDC_affectedCarrierFreqCombMRDC)
		if err = ie.affectedCarrierFreqCombMRDC.Decode(r); err != nil {
			return utils.WrapError("Decode affectedCarrierFreqCombMRDC", err)
		}
	}
	return nil
}
