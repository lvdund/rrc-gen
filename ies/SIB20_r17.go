package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB20_r17 struct {
	mcch_Config_r17          MCCH_Config_r17          `madatory`
	cfr_ConfigMCCH_MTCH_r17  *CFR_ConfigMCCH_MTCH_r17 `optional`
	lateNonCriticalExtension *[]byte                  `optional`
}

func (ie *SIB20_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.cfr_ConfigMCCH_MTCH_r17 != nil, ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.mcch_Config_r17.Encode(w); err != nil {
		return utils.WrapError("Encode mcch_Config_r17", err)
	}
	if ie.cfr_ConfigMCCH_MTCH_r17 != nil {
		if err = ie.cfr_ConfigMCCH_MTCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode cfr_ConfigMCCH_MTCH_r17", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB20_r17) Decode(r *uper.UperReader) error {
	var err error
	var cfr_ConfigMCCH_MTCH_r17Present bool
	if cfr_ConfigMCCH_MTCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.mcch_Config_r17.Decode(r); err != nil {
		return utils.WrapError("Decode mcch_Config_r17", err)
	}
	if cfr_ConfigMCCH_MTCH_r17Present {
		ie.cfr_ConfigMCCH_MTCH_r17 = new(CFR_ConfigMCCH_MTCH_r17)
		if err = ie.cfr_ConfigMCCH_MTCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode cfr_ConfigMCCH_MTCH_r17", err)
		}
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	return nil
}
