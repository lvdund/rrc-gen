package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCP_ParametersSidelink_r16 struct {
	outOfOrderDeliverySidelink_r16 *PDCP_ParametersSidelink_r16_outOfOrderDeliverySidelink_r16 `optional`
}

func (ie *PDCP_ParametersSidelink_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.outOfOrderDeliverySidelink_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.outOfOrderDeliverySidelink_r16 != nil {
		if err = ie.outOfOrderDeliverySidelink_r16.Encode(w); err != nil {
			return utils.WrapError("Encode outOfOrderDeliverySidelink_r16", err)
		}
	}
	return nil
}

func (ie *PDCP_ParametersSidelink_r16) Decode(r *uper.UperReader) error {
	var err error
	var outOfOrderDeliverySidelink_r16Present bool
	if outOfOrderDeliverySidelink_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if outOfOrderDeliverySidelink_r16Present {
		ie.outOfOrderDeliverySidelink_r16 = new(PDCP_ParametersSidelink_r16_outOfOrderDeliverySidelink_r16)
		if err = ie.outOfOrderDeliverySidelink_r16.Decode(r); err != nil {
			return utils.WrapError("Decode outOfOrderDeliverySidelink_r16", err)
		}
	}
	return nil
}
