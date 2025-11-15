package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfiguration_v1560_IEs struct {
	mrdc_SecondaryCellGroupConfig *MRDC_SecondaryCellGroupConfig `optional,setuprelease`
	radioBearerConfig2            *[]byte                        `optional`
	sk_Counter                    *SK_Counter                    `optional`
	nonCriticalExtension          *RRCReconfiguration_v1610_IEs  `optional`
}

func (ie *RRCReconfiguration_v1560_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.mrdc_SecondaryCellGroupConfig != nil, ie.radioBearerConfig2 != nil, ie.sk_Counter != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.mrdc_SecondaryCellGroupConfig != nil {
		tmp_mrdc_SecondaryCellGroupConfig := utils.SetupRelease[*MRDC_SecondaryCellGroupConfig]{
			Setup: ie.mrdc_SecondaryCellGroupConfig,
		}
		if err = tmp_mrdc_SecondaryCellGroupConfig.Encode(w); err != nil {
			return utils.WrapError("Encode mrdc_SecondaryCellGroupConfig", err)
		}
	}
	if ie.radioBearerConfig2 != nil {
		if err = w.WriteOctetString(*ie.radioBearerConfig2, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode radioBearerConfig2", err)
		}
	}
	if ie.sk_Counter != nil {
		if err = ie.sk_Counter.Encode(w); err != nil {
			return utils.WrapError("Encode sk_Counter", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCReconfiguration_v1560_IEs) Decode(r *uper.UperReader) error {
	var err error
	var mrdc_SecondaryCellGroupConfigPresent bool
	if mrdc_SecondaryCellGroupConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var radioBearerConfig2Present bool
	if radioBearerConfig2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sk_CounterPresent bool
	if sk_CounterPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if mrdc_SecondaryCellGroupConfigPresent {
		tmp_mrdc_SecondaryCellGroupConfig := utils.SetupRelease[*MRDC_SecondaryCellGroupConfig]{}
		if err = tmp_mrdc_SecondaryCellGroupConfig.Decode(r); err != nil {
			return utils.WrapError("Decode mrdc_SecondaryCellGroupConfig", err)
		}
		ie.mrdc_SecondaryCellGroupConfig = tmp_mrdc_SecondaryCellGroupConfig.Setup
	}
	if radioBearerConfig2Present {
		var tmp_os_radioBearerConfig2 []byte
		if tmp_os_radioBearerConfig2, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode radioBearerConfig2", err)
		}
		ie.radioBearerConfig2 = &tmp_os_radioBearerConfig2
	}
	if sk_CounterPresent {
		ie.sk_Counter = new(SK_Counter)
		if err = ie.sk_Counter.Decode(r); err != nil {
			return utils.WrapError("Decode sk_Counter", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(RRCReconfiguration_v1610_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
