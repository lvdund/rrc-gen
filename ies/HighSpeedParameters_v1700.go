package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type HighSpeedParameters_v1700 struct {
	measurementEnhancementCA_r17        *HighSpeedParameters_v1700_measurementEnhancementCA_r17        `optional`
	measurementEnhancementInterFreq_r17 *HighSpeedParameters_v1700_measurementEnhancementInterFreq_r17 `optional`
}

func (ie *HighSpeedParameters_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measurementEnhancementCA_r17 != nil, ie.measurementEnhancementInterFreq_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measurementEnhancementCA_r17 != nil {
		if err = ie.measurementEnhancementCA_r17.Encode(w); err != nil {
			return utils.WrapError("Encode measurementEnhancementCA_r17", err)
		}
	}
	if ie.measurementEnhancementInterFreq_r17 != nil {
		if err = ie.measurementEnhancementInterFreq_r17.Encode(w); err != nil {
			return utils.WrapError("Encode measurementEnhancementInterFreq_r17", err)
		}
	}
	return nil
}

func (ie *HighSpeedParameters_v1700) Decode(r *uper.UperReader) error {
	var err error
	var measurementEnhancementCA_r17Present bool
	if measurementEnhancementCA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var measurementEnhancementInterFreq_r17Present bool
	if measurementEnhancementInterFreq_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if measurementEnhancementCA_r17Present {
		ie.measurementEnhancementCA_r17 = new(HighSpeedParameters_v1700_measurementEnhancementCA_r17)
		if err = ie.measurementEnhancementCA_r17.Decode(r); err != nil {
			return utils.WrapError("Decode measurementEnhancementCA_r17", err)
		}
	}
	if measurementEnhancementInterFreq_r17Present {
		ie.measurementEnhancementInterFreq_r17 = new(HighSpeedParameters_v1700_measurementEnhancementInterFreq_r17)
		if err = ie.measurementEnhancementInterFreq_r17.Decode(r); err != nil {
			return utils.WrapError("Decode measurementEnhancementInterFreq_r17", err)
		}
	}
	return nil
}
