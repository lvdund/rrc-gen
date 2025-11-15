package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB1_v1610_IEs struct {
	idleModeMeasurementsEUTRA_r16 *SIB1_v1610_IEs_idleModeMeasurementsEUTRA_r16 `optional`
	idleModeMeasurementsNR_r16    *SIB1_v1610_IEs_idleModeMeasurementsNR_r16    `optional`
	posSI_SchedulingInfo_r16      *PosSI_SchedulingInfo_r16                     `optional`
	nonCriticalExtension          *SIB1_v1630_IEs                               `optional`
}

func (ie *SIB1_v1610_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.idleModeMeasurementsEUTRA_r16 != nil, ie.idleModeMeasurementsNR_r16 != nil, ie.posSI_SchedulingInfo_r16 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.idleModeMeasurementsEUTRA_r16 != nil {
		if err = ie.idleModeMeasurementsEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode idleModeMeasurementsEUTRA_r16", err)
		}
	}
	if ie.idleModeMeasurementsNR_r16 != nil {
		if err = ie.idleModeMeasurementsNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode idleModeMeasurementsNR_r16", err)
		}
	}
	if ie.posSI_SchedulingInfo_r16 != nil {
		if err = ie.posSI_SchedulingInfo_r16.Encode(w); err != nil {
			return utils.WrapError("Encode posSI_SchedulingInfo_r16", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB1_v1610_IEs) Decode(r *uper.UperReader) error {
	var err error
	var idleModeMeasurementsEUTRA_r16Present bool
	if idleModeMeasurementsEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var idleModeMeasurementsNR_r16Present bool
	if idleModeMeasurementsNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var posSI_SchedulingInfo_r16Present bool
	if posSI_SchedulingInfo_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if idleModeMeasurementsEUTRA_r16Present {
		ie.idleModeMeasurementsEUTRA_r16 = new(SIB1_v1610_IEs_idleModeMeasurementsEUTRA_r16)
		if err = ie.idleModeMeasurementsEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode idleModeMeasurementsEUTRA_r16", err)
		}
	}
	if idleModeMeasurementsNR_r16Present {
		ie.idleModeMeasurementsNR_r16 = new(SIB1_v1610_IEs_idleModeMeasurementsNR_r16)
		if err = ie.idleModeMeasurementsNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode idleModeMeasurementsNR_r16", err)
		}
	}
	if posSI_SchedulingInfo_r16Present {
		ie.posSI_SchedulingInfo_r16 = new(PosSI_SchedulingInfo_r16)
		if err = ie.posSI_SchedulingInfo_r16.Decode(r); err != nil {
			return utils.WrapError("Decode posSI_SchedulingInfo_r16", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(SIB1_v1630_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
