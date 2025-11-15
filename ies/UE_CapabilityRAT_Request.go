package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_CapabilityRAT_Request struct {
	rat_Type                RAT_Type `madatory`
	capabilityRequestFilter *[]byte  `optional`
}

func (ie *UE_CapabilityRAT_Request) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.capabilityRequestFilter != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.rat_Type.Encode(w); err != nil {
		return utils.WrapError("Encode rat_Type", err)
	}
	if ie.capabilityRequestFilter != nil {
		if err = w.WriteOctetString(*ie.capabilityRequestFilter, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode capabilityRequestFilter", err)
		}
	}
	return nil
}

func (ie *UE_CapabilityRAT_Request) Decode(r *uper.UperReader) error {
	var err error
	var capabilityRequestFilterPresent bool
	if capabilityRequestFilterPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.rat_Type.Decode(r); err != nil {
		return utils.WrapError("Decode rat_Type", err)
	}
	if capabilityRequestFilterPresent {
		var tmp_os_capabilityRequestFilter []byte
		if tmp_os_capabilityRequestFilter, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode capabilityRequestFilter", err)
		}
		ie.capabilityRequestFilter = &tmp_os_capabilityRequestFilter
	}
	return nil
}
