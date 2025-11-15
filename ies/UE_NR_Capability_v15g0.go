package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability_v15g0 struct {
	rf_Parameters_v15g0  *RF_Parameters_v15g0    `optional`
	nonCriticalExtension *UE_NR_Capability_v15j0 `optional`
}

func (ie *UE_NR_Capability_v15g0) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.rf_Parameters_v15g0 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.rf_Parameters_v15g0 != nil {
		if err = ie.rf_Parameters_v15g0.Encode(w); err != nil {
			return utils.WrapError("Encode rf_Parameters_v15g0", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability_v15g0) Decode(r *uper.UperReader) error {
	var err error
	var rf_Parameters_v15g0Present bool
	if rf_Parameters_v15g0Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if rf_Parameters_v15g0Present {
		ie.rf_Parameters_v15g0 = new(RF_Parameters_v15g0)
		if err = ie.rf_Parameters_v15g0.Decode(r); err != nil {
			return utils.WrapError("Decode rf_Parameters_v15g0", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(UE_NR_Capability_v15j0)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
