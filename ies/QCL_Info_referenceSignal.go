package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	QCL_Info_referenceSignal_Choice_nothing uint64 = iota
	QCL_Info_referenceSignal_Choice_csi_rs
	QCL_Info_referenceSignal_Choice_ssb
)

type QCL_Info_referenceSignal struct {
	Choice uint64
	csi_rs *NZP_CSI_RS_ResourceId
	ssb    *SSB_Index
}

func (ie *QCL_Info_referenceSignal) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case QCL_Info_referenceSignal_Choice_csi_rs:
		if err = ie.csi_rs.Encode(w); err != nil {
			err = utils.WrapError("Encode csi_rs", err)
		}
	case QCL_Info_referenceSignal_Choice_ssb:
		if err = ie.ssb.Encode(w); err != nil {
			err = utils.WrapError("Encode ssb", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *QCL_Info_referenceSignal) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case QCL_Info_referenceSignal_Choice_csi_rs:
		ie.csi_rs = new(NZP_CSI_RS_ResourceId)
		if err = ie.csi_rs.Decode(r); err != nil {
			return utils.WrapError("Decode csi_rs", err)
		}
	case QCL_Info_referenceSignal_Choice_ssb:
		ie.ssb = new(SSB_Index)
		if err = ie.ssb.Decode(r); err != nil {
			return utils.WrapError("Decode ssb", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
