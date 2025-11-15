package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17 struct {
	inter_SN_condPSCellChangeFDD_TDD_NRDC_r17    *MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_inter_SN_condPSCellChangeFDD_TDD_NRDC_r17    `optional`
	inter_SN_condPSCellChangeFR1_FR2_NRDC_r17    *MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_inter_SN_condPSCellChangeFR1_FR2_NRDC_r17    `optional`
	inter_SN_condPSCellChangeFDD_TDD_ENDC_r17    *MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_inter_SN_condPSCellChangeFDD_TDD_ENDC_r17    `optional`
	inter_SN_condPSCellChangeFR1_FR2_ENDC_r17    *MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_inter_SN_condPSCellChangeFR1_FR2_ENDC_r17    `optional`
	mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 *MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 `optional`
	mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 *MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 `optional`
	mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 *MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 `optional`
	sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 *MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 `optional`
	sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 *MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 `optional`
	sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 *MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 `optional`
}

func (ie *MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.inter_SN_condPSCellChangeFDD_TDD_NRDC_r17 != nil, ie.inter_SN_condPSCellChangeFR1_FR2_NRDC_r17 != nil, ie.inter_SN_condPSCellChangeFDD_TDD_ENDC_r17 != nil, ie.inter_SN_condPSCellChangeFR1_FR2_ENDC_r17 != nil, ie.mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 != nil, ie.mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 != nil, ie.mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 != nil, ie.sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 != nil, ie.sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 != nil, ie.sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.inter_SN_condPSCellChangeFDD_TDD_NRDC_r17 != nil {
		if err = ie.inter_SN_condPSCellChangeFDD_TDD_NRDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode inter_SN_condPSCellChangeFDD_TDD_NRDC_r17", err)
		}
	}
	if ie.inter_SN_condPSCellChangeFR1_FR2_NRDC_r17 != nil {
		if err = ie.inter_SN_condPSCellChangeFR1_FR2_NRDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode inter_SN_condPSCellChangeFR1_FR2_NRDC_r17", err)
		}
	}
	if ie.inter_SN_condPSCellChangeFDD_TDD_ENDC_r17 != nil {
		if err = ie.inter_SN_condPSCellChangeFDD_TDD_ENDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode inter_SN_condPSCellChangeFDD_TDD_ENDC_r17", err)
		}
	}
	if ie.inter_SN_condPSCellChangeFR1_FR2_ENDC_r17 != nil {
		if err = ie.inter_SN_condPSCellChangeFR1_FR2_ENDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode inter_SN_condPSCellChangeFR1_FR2_ENDC_r17", err)
		}
	}
	if ie.mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 != nil {
		if err = ie.mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17", err)
		}
	}
	if ie.mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 != nil {
		if err = ie.mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17", err)
		}
	}
	if ie.mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 != nil {
		if err = ie.mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17", err)
		}
	}
	if ie.sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 != nil {
		if err = ie.sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17", err)
		}
	}
	if ie.sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 != nil {
		if err = ie.sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17", err)
		}
	}
	if ie.sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 != nil {
		if err = ie.sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17) Decode(r *uper.UperReader) error {
	var err error
	var inter_SN_condPSCellChangeFDD_TDD_NRDC_r17Present bool
	if inter_SN_condPSCellChangeFDD_TDD_NRDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var inter_SN_condPSCellChangeFR1_FR2_NRDC_r17Present bool
	if inter_SN_condPSCellChangeFR1_FR2_NRDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var inter_SN_condPSCellChangeFDD_TDD_ENDC_r17Present bool
	if inter_SN_condPSCellChangeFDD_TDD_ENDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var inter_SN_condPSCellChangeFR1_FR2_ENDC_r17Present bool
	if inter_SN_condPSCellChangeFR1_FR2_ENDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17Present bool
	if mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17Present bool
	if mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17Present bool
	if mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17Present bool
	if sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17Present bool
	if sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17Present bool
	if sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if inter_SN_condPSCellChangeFDD_TDD_NRDC_r17Present {
		ie.inter_SN_condPSCellChangeFDD_TDD_NRDC_r17 = new(MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_inter_SN_condPSCellChangeFDD_TDD_NRDC_r17)
		if err = ie.inter_SN_condPSCellChangeFDD_TDD_NRDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode inter_SN_condPSCellChangeFDD_TDD_NRDC_r17", err)
		}
	}
	if inter_SN_condPSCellChangeFR1_FR2_NRDC_r17Present {
		ie.inter_SN_condPSCellChangeFR1_FR2_NRDC_r17 = new(MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_inter_SN_condPSCellChangeFR1_FR2_NRDC_r17)
		if err = ie.inter_SN_condPSCellChangeFR1_FR2_NRDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode inter_SN_condPSCellChangeFR1_FR2_NRDC_r17", err)
		}
	}
	if inter_SN_condPSCellChangeFDD_TDD_ENDC_r17Present {
		ie.inter_SN_condPSCellChangeFDD_TDD_ENDC_r17 = new(MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_inter_SN_condPSCellChangeFDD_TDD_ENDC_r17)
		if err = ie.inter_SN_condPSCellChangeFDD_TDD_ENDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode inter_SN_condPSCellChangeFDD_TDD_ENDC_r17", err)
		}
	}
	if inter_SN_condPSCellChangeFR1_FR2_ENDC_r17Present {
		ie.inter_SN_condPSCellChangeFR1_FR2_ENDC_r17 = new(MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_inter_SN_condPSCellChangeFR1_FR2_ENDC_r17)
		if err = ie.inter_SN_condPSCellChangeFR1_FR2_ENDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode inter_SN_condPSCellChangeFR1_FR2_ENDC_r17", err)
		}
	}
	if mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17Present {
		ie.mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 = new(MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17)
		if err = ie.mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17", err)
		}
	}
	if mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17Present {
		ie.mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 = new(MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17)
		if err = ie.mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17", err)
		}
	}
	if mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17Present {
		ie.mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 = new(MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17)
		if err = ie.mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17", err)
		}
	}
	if sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17Present {
		ie.sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17 = new(MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17)
		if err = ie.sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sn_InitiatedCondPSCellChange_FR1FDD_ENDC_r17", err)
		}
	}
	if sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17Present {
		ie.sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17 = new(MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17)
		if err = ie.sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sn_InitiatedCondPSCellChange_FR1TDD_ENDC_r17", err)
		}
	}
	if sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17Present {
		ie.sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17 = new(MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17_sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17)
		if err = ie.sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sn_InitiatedCondPSCellChange_FR2TDD_ENDC_r17", err)
		}
	}
	return nil
}
