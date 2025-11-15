package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ZP_CSI_RS_Resource struct {
	zp_CSI_RS_ResourceId ZP_CSI_RS_ResourceId              `madatory`
	resourceMapping      CSI_RS_ResourceMapping            `madatory`
	periodicityAndOffset *CSI_ResourcePeriodicityAndOffset `optional`
}

func (ie *ZP_CSI_RS_Resource) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.periodicityAndOffset != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.zp_CSI_RS_ResourceId.Encode(w); err != nil {
		return utils.WrapError("Encode zp_CSI_RS_ResourceId", err)
	}
	if err = ie.resourceMapping.Encode(w); err != nil {
		return utils.WrapError("Encode resourceMapping", err)
	}
	if ie.periodicityAndOffset != nil {
		if err = ie.periodicityAndOffset.Encode(w); err != nil {
			return utils.WrapError("Encode periodicityAndOffset", err)
		}
	}
	return nil
}

func (ie *ZP_CSI_RS_Resource) Decode(r *uper.UperReader) error {
	var err error
	var periodicityAndOffsetPresent bool
	if periodicityAndOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.zp_CSI_RS_ResourceId.Decode(r); err != nil {
		return utils.WrapError("Decode zp_CSI_RS_ResourceId", err)
	}
	if err = ie.resourceMapping.Decode(r); err != nil {
		return utils.WrapError("Decode resourceMapping", err)
	}
	if periodicityAndOffsetPresent {
		ie.periodicityAndOffset = new(CSI_ResourcePeriodicityAndOffset)
		if err = ie.periodicityAndOffset.Decode(r); err != nil {
			return utils.WrapError("Decode periodicityAndOffset", err)
		}
	}
	return nil
}
