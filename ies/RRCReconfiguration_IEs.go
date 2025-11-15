package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfiguration_IEs struct {
	radioBearerConfig        *RadioBearerConfig            `optional`
	secondaryCellGroup       *[]byte                       `optional`
	measConfig               *MeasConfig                   `optional`
	lateNonCriticalExtension *[]byte                       `optional`
	nonCriticalExtension     *RRCReconfiguration_v1530_IEs `optional`
}

func (ie *RRCReconfiguration_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.radioBearerConfig != nil, ie.secondaryCellGroup != nil, ie.measConfig != nil, ie.lateNonCriticalExtension != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.radioBearerConfig != nil {
		if err = ie.radioBearerConfig.Encode(w); err != nil {
			return utils.WrapError("Encode radioBearerConfig", err)
		}
	}
	if ie.secondaryCellGroup != nil {
		if err = w.WriteOctetString(*ie.secondaryCellGroup, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode secondaryCellGroup", err)
		}
	}
	if ie.measConfig != nil {
		if err = ie.measConfig.Encode(w); err != nil {
			return utils.WrapError("Encode measConfig", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCReconfiguration_IEs) Decode(r *uper.UperReader) error {
	var err error
	var radioBearerConfigPresent bool
	if radioBearerConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var secondaryCellGroupPresent bool
	if secondaryCellGroupPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measConfigPresent bool
	if measConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if radioBearerConfigPresent {
		ie.radioBearerConfig = new(RadioBearerConfig)
		if err = ie.radioBearerConfig.Decode(r); err != nil {
			return utils.WrapError("Decode radioBearerConfig", err)
		}
	}
	if secondaryCellGroupPresent {
		var tmp_os_secondaryCellGroup []byte
		if tmp_os_secondaryCellGroup, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode secondaryCellGroup", err)
		}
		ie.secondaryCellGroup = &tmp_os_secondaryCellGroup
	}
	if measConfigPresent {
		ie.measConfig = new(MeasConfig)
		if err = ie.measConfig.Decode(r); err != nil {
			return utils.WrapError("Decode measConfig", err)
		}
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(RRCReconfiguration_v1530_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
