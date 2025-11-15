package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RRCReconfigurationComplete_v1560_IEs_scg_Response_Choice_nothing uint64 = iota
	RRCReconfigurationComplete_v1560_IEs_scg_Response_Choice_nr_SCG_Response
	RRCReconfigurationComplete_v1560_IEs_scg_Response_Choice_eutra_SCG_Response
)

type RRCReconfigurationComplete_v1560_IEs_scg_Response struct {
	Choice             uint64
	nr_SCG_Response    []byte `madatory`
	eutra_SCG_Response []byte `madatory`
}

func (ie *RRCReconfigurationComplete_v1560_IEs_scg_Response) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCReconfigurationComplete_v1560_IEs_scg_Response_Choice_nr_SCG_Response:
		if err = w.WriteOctetString(ie.nr_SCG_Response, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			err = utils.WrapError("Encode nr_SCG_Response", err)
		}
	case RRCReconfigurationComplete_v1560_IEs_scg_Response_Choice_eutra_SCG_Response:
		if err = w.WriteOctetString(ie.eutra_SCG_Response, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			err = utils.WrapError("Encode eutra_SCG_Response", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RRCReconfigurationComplete_v1560_IEs_scg_Response) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RRCReconfigurationComplete_v1560_IEs_scg_Response_Choice_nr_SCG_Response:
		var tmp_os_nr_SCG_Response []byte
		if tmp_os_nr_SCG_Response, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode nr_SCG_Response", err)
		}
		ie.nr_SCG_Response = tmp_os_nr_SCG_Response
	case RRCReconfigurationComplete_v1560_IEs_scg_Response_Choice_eutra_SCG_Response:
		var tmp_os_eutra_SCG_Response []byte
		if tmp_os_eutra_SCG_Response, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode eutra_SCG_Response", err)
		}
		ie.eutra_SCG_Response = tmp_os_eutra_SCG_Response
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
