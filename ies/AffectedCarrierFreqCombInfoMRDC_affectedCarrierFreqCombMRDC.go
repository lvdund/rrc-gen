package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type AffectedCarrierFreqCombInfoMRDC_affectedCarrierFreqCombMRDC struct {
	affectedCarrierFreqCombEUTRA *AffectedCarrierFreqCombEUTRA `optional`
	affectedCarrierFreqCombNR    AffectedCarrierFreqCombNR     `madatory`
}

func (ie *AffectedCarrierFreqCombInfoMRDC_affectedCarrierFreqCombMRDC) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.affectedCarrierFreqCombEUTRA != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.affectedCarrierFreqCombEUTRA != nil {
		if err = ie.affectedCarrierFreqCombEUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode affectedCarrierFreqCombEUTRA", err)
		}
	}
	if err = ie.affectedCarrierFreqCombNR.Encode(w); err != nil {
		return utils.WrapError("Encode affectedCarrierFreqCombNR", err)
	}
	return nil
}

func (ie *AffectedCarrierFreqCombInfoMRDC_affectedCarrierFreqCombMRDC) Decode(r *uper.UperReader) error {
	var err error
	var affectedCarrierFreqCombEUTRAPresent bool
	if affectedCarrierFreqCombEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if affectedCarrierFreqCombEUTRAPresent {
		ie.affectedCarrierFreqCombEUTRA = new(AffectedCarrierFreqCombEUTRA)
		if err = ie.affectedCarrierFreqCombEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode affectedCarrierFreqCombEUTRA", err)
		}
	}
	if err = ie.affectedCarrierFreqCombNR.Decode(r); err != nil {
		return utils.WrapError("Decode affectedCarrierFreqCombNR", err)
	}
	return nil
}
