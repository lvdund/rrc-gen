package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasurementTimingConfiguration_v1550_IEs struct {
	campOnFirstSSB       bool                                      `madatory`
	psCellOnlyOnFirstSSB bool                                      `madatory`
	nonCriticalExtension *MeasurementTimingConfiguration_v1610_IEs `optional`
}

func (ie *MeasurementTimingConfiguration_v1550_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteBoolean(ie.campOnFirstSSB); err != nil {
		return utils.WrapError("WriteBoolean campOnFirstSSB", err)
	}
	if err = w.WriteBoolean(ie.psCellOnlyOnFirstSSB); err != nil {
		return utils.WrapError("WriteBoolean psCellOnlyOnFirstSSB", err)
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *MeasurementTimingConfiguration_v1550_IEs) Decode(r *uper.UperReader) error {
	var err error
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_bool_campOnFirstSSB bool
	if tmp_bool_campOnFirstSSB, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean campOnFirstSSB", err)
	}
	ie.campOnFirstSSB = tmp_bool_campOnFirstSSB
	var tmp_bool_psCellOnlyOnFirstSSB bool
	if tmp_bool_psCellOnlyOnFirstSSB, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean psCellOnlyOnFirstSSB", err)
	}
	ie.psCellOnlyOnFirstSSB = tmp_bool_psCellOnlyOnFirstSSB
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(MeasurementTimingConfiguration_v1610_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
