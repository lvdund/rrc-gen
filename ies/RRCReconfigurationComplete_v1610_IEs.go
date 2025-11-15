package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfigurationComplete_v1610_IEs struct {
	ue_MeasurementsAvailable_r16 *UE_MeasurementsAvailable_r16         `optional`
	needForGapsInfoNR_r16        *NeedForGapsInfoNR_r16                `optional`
	nonCriticalExtension         *RRCReconfigurationComplete_v1640_IEs `optional`
}

func (ie *RRCReconfigurationComplete_v1610_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ue_MeasurementsAvailable_r16 != nil, ie.needForGapsInfoNR_r16 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ue_MeasurementsAvailable_r16 != nil {
		if err = ie.ue_MeasurementsAvailable_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ue_MeasurementsAvailable_r16", err)
		}
	}
	if ie.needForGapsInfoNR_r16 != nil {
		if err = ie.needForGapsInfoNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode needForGapsInfoNR_r16", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCReconfigurationComplete_v1610_IEs) Decode(r *uper.UperReader) error {
	var err error
	var ue_MeasurementsAvailable_r16Present bool
	if ue_MeasurementsAvailable_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var needForGapsInfoNR_r16Present bool
	if needForGapsInfoNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ue_MeasurementsAvailable_r16Present {
		ie.ue_MeasurementsAvailable_r16 = new(UE_MeasurementsAvailable_r16)
		if err = ie.ue_MeasurementsAvailable_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ue_MeasurementsAvailable_r16", err)
		}
	}
	if needForGapsInfoNR_r16Present {
		ie.needForGapsInfoNR_r16 = new(NeedForGapsInfoNR_r16)
		if err = ie.needForGapsInfoNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode needForGapsInfoNR_r16", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(RRCReconfigurationComplete_v1640_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
