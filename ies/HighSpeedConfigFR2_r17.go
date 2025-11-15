package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type HighSpeedConfigFR2_r17 struct {
	highSpeedMeasFlagFR2_r17              *HighSpeedConfigFR2_r17_highSpeedMeasFlagFR2_r17              `optional`
	highSpeedDeploymentTypeFR2_r17        *HighSpeedConfigFR2_r17_highSpeedDeploymentTypeFR2_r17        `optional`
	highSpeedLargeOneStepUL_TimingFR2_r17 *HighSpeedConfigFR2_r17_highSpeedLargeOneStepUL_TimingFR2_r17 `optional`
}

func (ie *HighSpeedConfigFR2_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.highSpeedMeasFlagFR2_r17 != nil, ie.highSpeedDeploymentTypeFR2_r17 != nil, ie.highSpeedLargeOneStepUL_TimingFR2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.highSpeedMeasFlagFR2_r17 != nil {
		if err = ie.highSpeedMeasFlagFR2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode highSpeedMeasFlagFR2_r17", err)
		}
	}
	if ie.highSpeedDeploymentTypeFR2_r17 != nil {
		if err = ie.highSpeedDeploymentTypeFR2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode highSpeedDeploymentTypeFR2_r17", err)
		}
	}
	if ie.highSpeedLargeOneStepUL_TimingFR2_r17 != nil {
		if err = ie.highSpeedLargeOneStepUL_TimingFR2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode highSpeedLargeOneStepUL_TimingFR2_r17", err)
		}
	}
	return nil
}

func (ie *HighSpeedConfigFR2_r17) Decode(r *uper.UperReader) error {
	var err error
	var highSpeedMeasFlagFR2_r17Present bool
	if highSpeedMeasFlagFR2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var highSpeedDeploymentTypeFR2_r17Present bool
	if highSpeedDeploymentTypeFR2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var highSpeedLargeOneStepUL_TimingFR2_r17Present bool
	if highSpeedLargeOneStepUL_TimingFR2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if highSpeedMeasFlagFR2_r17Present {
		ie.highSpeedMeasFlagFR2_r17 = new(HighSpeedConfigFR2_r17_highSpeedMeasFlagFR2_r17)
		if err = ie.highSpeedMeasFlagFR2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode highSpeedMeasFlagFR2_r17", err)
		}
	}
	if highSpeedDeploymentTypeFR2_r17Present {
		ie.highSpeedDeploymentTypeFR2_r17 = new(HighSpeedConfigFR2_r17_highSpeedDeploymentTypeFR2_r17)
		if err = ie.highSpeedDeploymentTypeFR2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode highSpeedDeploymentTypeFR2_r17", err)
		}
	}
	if highSpeedLargeOneStepUL_TimingFR2_r17Present {
		ie.highSpeedLargeOneStepUL_TimingFR2_r17 = new(HighSpeedConfigFR2_r17_highSpeedLargeOneStepUL_TimingFR2_r17)
		if err = ie.highSpeedLargeOneStepUL_TimingFR2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode highSpeedLargeOneStepUL_TimingFR2_r17", err)
		}
	}
	return nil
}
