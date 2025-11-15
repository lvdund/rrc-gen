package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReestablishmentComplete_v1610_IEs struct {
	ue_MeasurementsAvailable_r16 *UE_MeasurementsAvailable_r16 `optional`
	nonCriticalExtension         interface{}                   `optional`
}

func (ie *RRCReestablishmentComplete_v1610_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ue_MeasurementsAvailable_r16 != nil}
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
	return nil
}

func (ie *RRCReestablishmentComplete_v1610_IEs) Decode(r *uper.UperReader) error {
	var err error
	var ue_MeasurementsAvailable_r16Present bool
	if ue_MeasurementsAvailable_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ue_MeasurementsAvailable_r16Present {
		ie.ue_MeasurementsAvailable_r16 = new(UE_MeasurementsAvailable_r16)
		if err = ie.ue_MeasurementsAvailable_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ue_MeasurementsAvailable_r16", err)
		}
	}
	return nil
}
