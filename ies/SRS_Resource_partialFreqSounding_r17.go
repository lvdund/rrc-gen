package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_Resource_partialFreqSounding_r17 struct {
	startRBIndexFScaling_r17 SRS_Resource_partialFreqSounding_r17_startRBIndexFScaling_r17  `lb:0,ub:1,madatory`
	enableStartRBHopping_r17 *SRS_Resource_partialFreqSounding_r17_enableStartRBHopping_r17 `optional`
}

func (ie *SRS_Resource_partialFreqSounding_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.enableStartRBHopping_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.startRBIndexFScaling_r17.Encode(w); err != nil {
		return utils.WrapError("Encode startRBIndexFScaling_r17", err)
	}
	if ie.enableStartRBHopping_r17 != nil {
		if err = ie.enableStartRBHopping_r17.Encode(w); err != nil {
			return utils.WrapError("Encode enableStartRBHopping_r17", err)
		}
	}
	return nil
}

func (ie *SRS_Resource_partialFreqSounding_r17) Decode(r *uper.UperReader) error {
	var err error
	var enableStartRBHopping_r17Present bool
	if enableStartRBHopping_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.startRBIndexFScaling_r17.Decode(r); err != nil {
		return utils.WrapError("Decode startRBIndexFScaling_r17", err)
	}
	if enableStartRBHopping_r17Present {
		ie.enableStartRBHopping_r17 = new(SRS_Resource_partialFreqSounding_r17_enableStartRBHopping_r17)
		if err = ie.enableStartRBHopping_r17.Decode(r); err != nil {
			return utils.WrapError("Decode enableStartRBHopping_r17", err)
		}
	}
	return nil
}
