package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCResumeComplete_v1700_IEs struct {
	needForGapNCSG_InfoNR_r17    *NeedForGapNCSG_InfoNR_r17    `optional`
	needForGapNCSG_InfoEUTRA_r17 *NeedForGapNCSG_InfoEUTRA_r17 `optional`
	nonCriticalExtension         *RRCResumeComplete_v1720_IEs  `optional`
}

func (ie *RRCResumeComplete_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.needForGapNCSG_InfoNR_r17 != nil, ie.needForGapNCSG_InfoEUTRA_r17 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.needForGapNCSG_InfoNR_r17 != nil {
		if err = ie.needForGapNCSG_InfoNR_r17.Encode(w); err != nil {
			return utils.WrapError("Encode needForGapNCSG_InfoNR_r17", err)
		}
	}
	if ie.needForGapNCSG_InfoEUTRA_r17 != nil {
		if err = ie.needForGapNCSG_InfoEUTRA_r17.Encode(w); err != nil {
			return utils.WrapError("Encode needForGapNCSG_InfoEUTRA_r17", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCResumeComplete_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var needForGapNCSG_InfoNR_r17Present bool
	if needForGapNCSG_InfoNR_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var needForGapNCSG_InfoEUTRA_r17Present bool
	if needForGapNCSG_InfoEUTRA_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if needForGapNCSG_InfoNR_r17Present {
		ie.needForGapNCSG_InfoNR_r17 = new(NeedForGapNCSG_InfoNR_r17)
		if err = ie.needForGapNCSG_InfoNR_r17.Decode(r); err != nil {
			return utils.WrapError("Decode needForGapNCSG_InfoNR_r17", err)
		}
	}
	if needForGapNCSG_InfoEUTRA_r17Present {
		ie.needForGapNCSG_InfoEUTRA_r17 = new(NeedForGapNCSG_InfoEUTRA_r17)
		if err = ie.needForGapNCSG_InfoEUTRA_r17.Decode(r); err != nil {
			return utils.WrapError("Decode needForGapNCSG_InfoEUTRA_r17", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(RRCResumeComplete_v1720_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
