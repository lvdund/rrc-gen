package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_IM_Resource struct {
	csi_IM_ResourceId             CSI_IM_ResourceId                              `madatory`
	csi_IM_ResourceElementPattern *CSI_IM_Resource_csi_IM_ResourceElementPattern `lb:0,ub:12,optional`
	freqBand                      *CSI_FrequencyOccupation                       `optional`
	periodicityAndOffset          *CSI_ResourcePeriodicityAndOffset              `optional`
}

func (ie *CSI_IM_Resource) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.csi_IM_ResourceElementPattern != nil, ie.freqBand != nil, ie.periodicityAndOffset != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.csi_IM_ResourceId.Encode(w); err != nil {
		return utils.WrapError("Encode csi_IM_ResourceId", err)
	}
	if ie.csi_IM_ResourceElementPattern != nil {
		if err = ie.csi_IM_ResourceElementPattern.Encode(w); err != nil {
			return utils.WrapError("Encode csi_IM_ResourceElementPattern", err)
		}
	}
	if ie.freqBand != nil {
		if err = ie.freqBand.Encode(w); err != nil {
			return utils.WrapError("Encode freqBand", err)
		}
	}
	if ie.periodicityAndOffset != nil {
		if err = ie.periodicityAndOffset.Encode(w); err != nil {
			return utils.WrapError("Encode periodicityAndOffset", err)
		}
	}
	return nil
}

func (ie *CSI_IM_Resource) Decode(r *uper.UperReader) error {
	var err error
	var csi_IM_ResourceElementPatternPresent bool
	if csi_IM_ResourceElementPatternPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var freqBandPresent bool
	if freqBandPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var periodicityAndOffsetPresent bool
	if periodicityAndOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.csi_IM_ResourceId.Decode(r); err != nil {
		return utils.WrapError("Decode csi_IM_ResourceId", err)
	}
	if csi_IM_ResourceElementPatternPresent {
		ie.csi_IM_ResourceElementPattern = new(CSI_IM_Resource_csi_IM_ResourceElementPattern)
		if err = ie.csi_IM_ResourceElementPattern.Decode(r); err != nil {
			return utils.WrapError("Decode csi_IM_ResourceElementPattern", err)
		}
	}
	if freqBandPresent {
		ie.freqBand = new(CSI_FrequencyOccupation)
		if err = ie.freqBand.Decode(r); err != nil {
			return utils.WrapError("Decode freqBand", err)
		}
	}
	if periodicityAndOffsetPresent {
		ie.periodicityAndOffset = new(CSI_ResourcePeriodicityAndOffset)
		if err = ie.periodicityAndOffset.Decode(r); err != nil {
			return utils.WrapError("Decode periodicityAndOffset", err)
		}
	}
	return nil
}
