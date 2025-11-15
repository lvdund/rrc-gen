package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Resourcetype_r16_SRS_PosResource_r16_periodic_r16 struct {
	periodicityAndOffset_p_r16     SRS_PeriodicityAndOffset_r16     `madatory`
	periodicityAndOffset_p_Ext_r16 *SRS_PeriodicityAndOffsetExt_r16 `optional`
}

func (ie *Resourcetype_r16_SRS_PosResource_r16_periodic_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.periodicityAndOffset_p_Ext_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.periodicityAndOffset_p_r16.Encode(w); err != nil {
		return utils.WrapError("Encode periodicityAndOffset_p_r16", err)
	}
	if ie.periodicityAndOffset_p_Ext_r16 != nil {
		if err = ie.periodicityAndOffset_p_Ext_r16.Encode(w); err != nil {
			return utils.WrapError("Encode periodicityAndOffset_p_Ext_r16", err)
		}
	}
	return nil
}

func (ie *Resourcetype_r16_SRS_PosResource_r16_periodic_r16) Decode(r *uper.UperReader) error {
	var err error
	var periodicityAndOffset_p_Ext_r16Present bool
	if periodicityAndOffset_p_Ext_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.periodicityAndOffset_p_r16.Decode(r); err != nil {
		return utils.WrapError("Decode periodicityAndOffset_p_r16", err)
	}
	if periodicityAndOffset_p_Ext_r16Present {
		ie.periodicityAndOffset_p_Ext_r16 = new(SRS_PeriodicityAndOffsetExt_r16)
		if err = ie.periodicityAndOffset_p_Ext_r16.Decode(r); err != nil {
			return utils.WrapError("Decode periodicityAndOffset_p_Ext_r16", err)
		}
	}
	return nil
}
