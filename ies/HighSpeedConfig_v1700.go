package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type HighSpeedConfig_v1700 struct {
	highSpeedMeasCA_Scell_r17  *HighSpeedConfig_v1700_highSpeedMeasCA_Scell_r17  `optional`
	highSpeedMeasInterFreq_r17 *HighSpeedConfig_v1700_highSpeedMeasInterFreq_r17 `optional`
	highSpeedDemodCA_Scell_r17 *HighSpeedConfig_v1700_highSpeedDemodCA_Scell_r17 `optional`
}

func (ie *HighSpeedConfig_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.highSpeedMeasCA_Scell_r17 != nil, ie.highSpeedMeasInterFreq_r17 != nil, ie.highSpeedDemodCA_Scell_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.highSpeedMeasCA_Scell_r17 != nil {
		if err = ie.highSpeedMeasCA_Scell_r17.Encode(w); err != nil {
			return utils.WrapError("Encode highSpeedMeasCA_Scell_r17", err)
		}
	}
	if ie.highSpeedMeasInterFreq_r17 != nil {
		if err = ie.highSpeedMeasInterFreq_r17.Encode(w); err != nil {
			return utils.WrapError("Encode highSpeedMeasInterFreq_r17", err)
		}
	}
	if ie.highSpeedDemodCA_Scell_r17 != nil {
		if err = ie.highSpeedDemodCA_Scell_r17.Encode(w); err != nil {
			return utils.WrapError("Encode highSpeedDemodCA_Scell_r17", err)
		}
	}
	return nil
}

func (ie *HighSpeedConfig_v1700) Decode(r *uper.UperReader) error {
	var err error
	var highSpeedMeasCA_Scell_r17Present bool
	if highSpeedMeasCA_Scell_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var highSpeedMeasInterFreq_r17Present bool
	if highSpeedMeasInterFreq_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var highSpeedDemodCA_Scell_r17Present bool
	if highSpeedDemodCA_Scell_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if highSpeedMeasCA_Scell_r17Present {
		ie.highSpeedMeasCA_Scell_r17 = new(HighSpeedConfig_v1700_highSpeedMeasCA_Scell_r17)
		if err = ie.highSpeedMeasCA_Scell_r17.Decode(r); err != nil {
			return utils.WrapError("Decode highSpeedMeasCA_Scell_r17", err)
		}
	}
	if highSpeedMeasInterFreq_r17Present {
		ie.highSpeedMeasInterFreq_r17 = new(HighSpeedConfig_v1700_highSpeedMeasInterFreq_r17)
		if err = ie.highSpeedMeasInterFreq_r17.Decode(r); err != nil {
			return utils.WrapError("Decode highSpeedMeasInterFreq_r17", err)
		}
	}
	if highSpeedDemodCA_Scell_r17Present {
		ie.highSpeedDemodCA_Scell_r17 = new(HighSpeedConfig_v1700_highSpeedDemodCA_Scell_r17)
		if err = ie.highSpeedDemodCA_Scell_r17.Decode(r); err != nil {
			return utils.WrapError("Decode highSpeedDemodCA_Scell_r17", err)
		}
	}
	return nil
}
