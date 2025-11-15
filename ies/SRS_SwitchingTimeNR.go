package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_SwitchingTimeNR struct {
	switchingTimeDL *SRS_SwitchingTimeNR_switchingTimeDL `optional`
	switchingTimeUL *SRS_SwitchingTimeNR_switchingTimeUL `optional`
}

func (ie *SRS_SwitchingTimeNR) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.switchingTimeDL != nil, ie.switchingTimeUL != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.switchingTimeDL != nil {
		if err = ie.switchingTimeDL.Encode(w); err != nil {
			return utils.WrapError("Encode switchingTimeDL", err)
		}
	}
	if ie.switchingTimeUL != nil {
		if err = ie.switchingTimeUL.Encode(w); err != nil {
			return utils.WrapError("Encode switchingTimeUL", err)
		}
	}
	return nil
}

func (ie *SRS_SwitchingTimeNR) Decode(r *uper.UperReader) error {
	var err error
	var switchingTimeDLPresent bool
	if switchingTimeDLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var switchingTimeULPresent bool
	if switchingTimeULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if switchingTimeDLPresent {
		ie.switchingTimeDL = new(SRS_SwitchingTimeNR_switchingTimeDL)
		if err = ie.switchingTimeDL.Decode(r); err != nil {
			return utils.WrapError("Decode switchingTimeDL", err)
		}
	}
	if switchingTimeULPresent {
		ie.switchingTimeUL = new(SRS_SwitchingTimeNR_switchingTimeUL)
		if err = ie.switchingTimeUL.Decode(r); err != nil {
			return utils.WrapError("Decode switchingTimeUL", err)
		}
	}
	return nil
}
