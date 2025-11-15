package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DL_PRS_QCL_Info_r17_Choice_nothing uint64 = iota
	DL_PRS_QCL_Info_r17_Choice_ssb_r17
	DL_PRS_QCL_Info_r17_Choice_dl_PRS_r17
)

type DL_PRS_QCL_Info_r17 struct {
	Choice     uint64
	ssb_r17    *DL_PRS_QCL_Info_r17_ssb_r17
	dl_PRS_r17 *DL_PRS_QCL_Info_r17_dl_PRS_r17
}

func (ie *DL_PRS_QCL_Info_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DL_PRS_QCL_Info_r17_Choice_ssb_r17:
		if err = ie.ssb_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode ssb_r17", err)
		}
	case DL_PRS_QCL_Info_r17_Choice_dl_PRS_r17:
		if err = ie.dl_PRS_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode dl_PRS_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *DL_PRS_QCL_Info_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DL_PRS_QCL_Info_r17_Choice_ssb_r17:
		ie.ssb_r17 = new(DL_PRS_QCL_Info_r17_ssb_r17)
		if err = ie.ssb_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_r17", err)
		}
	case DL_PRS_QCL_Info_r17_Choice_dl_PRS_r17:
		ie.dl_PRS_r17 = new(DL_PRS_QCL_Info_r17_dl_PRS_r17)
		if err = ie.dl_PRS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode dl_PRS_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
