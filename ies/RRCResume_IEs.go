package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCResume_IEs struct {
	radioBearerConfig        *RadioBearerConfig        `optional`
	masterCellGroup          *[]byte                   `optional`
	measConfig               *MeasConfig               `optional`
	fullConfig               *RRCResume_IEs_fullConfig `optional`
	lateNonCriticalExtension *[]byte                   `optional`
	nonCriticalExtension     *RRCResume_v1560_IEs      `optional`
}

func (ie *RRCResume_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.radioBearerConfig != nil, ie.masterCellGroup != nil, ie.measConfig != nil, ie.fullConfig != nil, ie.lateNonCriticalExtension != nil, ie.nonCriticalExtension != nil}
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
	if ie.masterCellGroup != nil {
		if err = w.WriteOctetString(*ie.masterCellGroup, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode masterCellGroup", err)
		}
	}
	if ie.measConfig != nil {
		if err = ie.measConfig.Encode(w); err != nil {
			return utils.WrapError("Encode measConfig", err)
		}
	}
	if ie.fullConfig != nil {
		if err = ie.fullConfig.Encode(w); err != nil {
			return utils.WrapError("Encode fullConfig", err)
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

func (ie *RRCResume_IEs) Decode(r *uper.UperReader) error {
	var err error
	var radioBearerConfigPresent bool
	if radioBearerConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var masterCellGroupPresent bool
	if masterCellGroupPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measConfigPresent bool
	if measConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var fullConfigPresent bool
	if fullConfigPresent, err = r.ReadBool(); err != nil {
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
	if masterCellGroupPresent {
		var tmp_os_masterCellGroup []byte
		if tmp_os_masterCellGroup, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode masterCellGroup", err)
		}
		ie.masterCellGroup = &tmp_os_masterCellGroup
	}
	if measConfigPresent {
		ie.measConfig = new(MeasConfig)
		if err = ie.measConfig.Decode(r); err != nil {
			return utils.WrapError("Decode measConfig", err)
		}
	}
	if fullConfigPresent {
		ie.fullConfig = new(RRCResume_IEs_fullConfig)
		if err = ie.fullConfig.Decode(r); err != nil {
			return utils.WrapError("Decode fullConfig", err)
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
		ie.nonCriticalExtension = new(RRCResume_v1560_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
