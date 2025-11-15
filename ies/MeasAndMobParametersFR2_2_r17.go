package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersFR2_2_r17 struct {
	handoverInterF_r17            *MeasAndMobParametersFR2_2_r17_handoverInterF_r17            `optional`
	handoverLTE_EPC_r17           *MeasAndMobParametersFR2_2_r17_handoverLTE_EPC_r17           `optional`
	handoverLTE_5GC_r17           *MeasAndMobParametersFR2_2_r17_handoverLTE_5GC_r17           `optional`
	idleInactiveNR_MeasReport_r17 *MeasAndMobParametersFR2_2_r17_idleInactiveNR_MeasReport_r17 `optional`
}

func (ie *MeasAndMobParametersFR2_2_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.handoverInterF_r17 != nil, ie.handoverLTE_EPC_r17 != nil, ie.handoverLTE_5GC_r17 != nil, ie.idleInactiveNR_MeasReport_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.handoverInterF_r17 != nil {
		if err = ie.handoverInterF_r17.Encode(w); err != nil {
			return utils.WrapError("Encode handoverInterF_r17", err)
		}
	}
	if ie.handoverLTE_EPC_r17 != nil {
		if err = ie.handoverLTE_EPC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode handoverLTE_EPC_r17", err)
		}
	}
	if ie.handoverLTE_5GC_r17 != nil {
		if err = ie.handoverLTE_5GC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode handoverLTE_5GC_r17", err)
		}
	}
	if ie.idleInactiveNR_MeasReport_r17 != nil {
		if err = ie.idleInactiveNR_MeasReport_r17.Encode(w); err != nil {
			return utils.WrapError("Encode idleInactiveNR_MeasReport_r17", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersFR2_2_r17) Decode(r *uper.UperReader) error {
	var err error
	var handoverInterF_r17Present bool
	if handoverInterF_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var handoverLTE_EPC_r17Present bool
	if handoverLTE_EPC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var handoverLTE_5GC_r17Present bool
	if handoverLTE_5GC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var idleInactiveNR_MeasReport_r17Present bool
	if idleInactiveNR_MeasReport_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if handoverInterF_r17Present {
		ie.handoverInterF_r17 = new(MeasAndMobParametersFR2_2_r17_handoverInterF_r17)
		if err = ie.handoverInterF_r17.Decode(r); err != nil {
			return utils.WrapError("Decode handoverInterF_r17", err)
		}
	}
	if handoverLTE_EPC_r17Present {
		ie.handoverLTE_EPC_r17 = new(MeasAndMobParametersFR2_2_r17_handoverLTE_EPC_r17)
		if err = ie.handoverLTE_EPC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode handoverLTE_EPC_r17", err)
		}
	}
	if handoverLTE_5GC_r17Present {
		ie.handoverLTE_5GC_r17 = new(MeasAndMobParametersFR2_2_r17_handoverLTE_5GC_r17)
		if err = ie.handoverLTE_5GC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode handoverLTE_5GC_r17", err)
		}
	}
	if idleInactiveNR_MeasReport_r17Present {
		ie.idleInactiveNR_MeasReport_r17 = new(MeasAndMobParametersFR2_2_r17_idleInactiveNR_MeasReport_r17)
		if err = ie.idleInactiveNR_MeasReport_r17.Decode(r); err != nil {
			return utils.WrapError("Decode idleInactiveNR_MeasReport_r17", err)
		}
	}
	return nil
}
