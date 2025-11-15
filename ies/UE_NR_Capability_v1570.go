package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability_v1570 struct {
	nrdc_Parameters_v1570 *NRDC_Parameters_v1570  `optional`
	nonCriticalExtension  *UE_NR_Capability_v1610 `optional`
}

func (ie *UE_NR_Capability_v1570) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.nrdc_Parameters_v1570 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.nrdc_Parameters_v1570 != nil {
		if err = ie.nrdc_Parameters_v1570.Encode(w); err != nil {
			return utils.WrapError("Encode nrdc_Parameters_v1570", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability_v1570) Decode(r *uper.UperReader) error {
	var err error
	var nrdc_Parameters_v1570Present bool
	if nrdc_Parameters_v1570Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if nrdc_Parameters_v1570Present {
		ie.nrdc_Parameters_v1570 = new(NRDC_Parameters_v1570)
		if err = ie.nrdc_Parameters_v1570.Decode(r); err != nil {
			return utils.WrapError("Decode nrdc_Parameters_v1570", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(UE_NR_Capability_v1610)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
