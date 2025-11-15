package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_CapabilityRequestFilterNR_v1710 struct {
	sidelinkRequest_r17  *UE_CapabilityRequestFilterNR_v1710_sidelinkRequest_r17 `optional`
	nonCriticalExtension interface{}                                             `optional`
}

func (ie *UE_CapabilityRequestFilterNR_v1710) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sidelinkRequest_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sidelinkRequest_r17 != nil {
		if err = ie.sidelinkRequest_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sidelinkRequest_r17", err)
		}
	}
	return nil
}

func (ie *UE_CapabilityRequestFilterNR_v1710) Decode(r *uper.UperReader) error {
	var err error
	var sidelinkRequest_r17Present bool
	if sidelinkRequest_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sidelinkRequest_r17Present {
		ie.sidelinkRequest_r17 = new(UE_CapabilityRequestFilterNR_v1710_sidelinkRequest_r17)
		if err = ie.sidelinkRequest_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sidelinkRequest_r17", err)
		}
	}
	return nil
}
