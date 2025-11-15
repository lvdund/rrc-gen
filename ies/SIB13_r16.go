package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB13_r16 struct {
	sl_V2X_ConfigCommon_r16  []byte  `madatory`
	dummy                    []byte  `madatory`
	tdd_Config_r16           []byte  `madatory`
	lateNonCriticalExtension *[]byte `optional`
}

func (ie *SIB13_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteOctetString(ie.sl_V2X_ConfigCommon_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString sl_V2X_ConfigCommon_r16", err)
	}
	if err = w.WriteOctetString(ie.dummy, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString dummy", err)
	}
	if err = w.WriteOctetString(ie.tdd_Config_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteOctetString tdd_Config_r16", err)
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB13_r16) Decode(r *uper.UperReader) error {
	var err error
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_os_sl_V2X_ConfigCommon_r16 []byte
	if tmp_os_sl_V2X_ConfigCommon_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString sl_V2X_ConfigCommon_r16", err)
	}
	ie.sl_V2X_ConfigCommon_r16 = tmp_os_sl_V2X_ConfigCommon_r16
	var tmp_os_dummy []byte
	if tmp_os_dummy, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString dummy", err)
	}
	ie.dummy = tmp_os_dummy
	var tmp_os_tdd_Config_r16 []byte
	if tmp_os_tdd_Config_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadOctetString tdd_Config_r16", err)
	}
	ie.tdd_Config_r16 = tmp_os_tdd_Config_r16
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	return nil
}
