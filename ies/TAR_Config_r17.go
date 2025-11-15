package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TAR_Config_r17 struct {
	offsetThresholdTA_r17 *TAR_Config_r17_offsetThresholdTA_r17 `optional`
	timingAdvanceSR_r17   *TAR_Config_r17_timingAdvanceSR_r17   `optional`
}

func (ie *TAR_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.offsetThresholdTA_r17 != nil, ie.timingAdvanceSR_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.offsetThresholdTA_r17 != nil {
		if err = ie.offsetThresholdTA_r17.Encode(w); err != nil {
			return utils.WrapError("Encode offsetThresholdTA_r17", err)
		}
	}
	if ie.timingAdvanceSR_r17 != nil {
		if err = ie.timingAdvanceSR_r17.Encode(w); err != nil {
			return utils.WrapError("Encode timingAdvanceSR_r17", err)
		}
	}
	return nil
}

func (ie *TAR_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	var offsetThresholdTA_r17Present bool
	if offsetThresholdTA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var timingAdvanceSR_r17Present bool
	if timingAdvanceSR_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if offsetThresholdTA_r17Present {
		ie.offsetThresholdTA_r17 = new(TAR_Config_r17_offsetThresholdTA_r17)
		if err = ie.offsetThresholdTA_r17.Decode(r); err != nil {
			return utils.WrapError("Decode offsetThresholdTA_r17", err)
		}
	}
	if timingAdvanceSR_r17Present {
		ie.timingAdvanceSR_r17 = new(TAR_Config_r17_timingAdvanceSR_r17)
		if err = ie.timingAdvanceSR_r17.Decode(r); err != nil {
			return utils.WrapError("Decode timingAdvanceSR_r17", err)
		}
	}
	return nil
}
