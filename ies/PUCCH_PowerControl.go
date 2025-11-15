package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_PowerControl struct {
	deltaF_PUCCH_f0              *int64                                           `lb:-16,ub:15,optional`
	deltaF_PUCCH_f1              *int64                                           `lb:-16,ub:15,optional`
	deltaF_PUCCH_f2              *int64                                           `lb:-16,ub:15,optional`
	deltaF_PUCCH_f3              *int64                                           `lb:-16,ub:15,optional`
	deltaF_PUCCH_f4              *int64                                           `lb:-16,ub:15,optional`
	p0_Set                       []P0_PUCCH                                       `lb:1,ub:maxNrofPUCCH_P0_PerSet,optional`
	pathlossReferenceRSs         []PUCCH_PathlossReferenceRS                      `lb:1,ub:maxNrofPUCCH_PathlossReferenceRSs,optional`
	twoPUCCH_PC_AdjustmentStates *PUCCH_PowerControl_twoPUCCH_PC_AdjustmentStates `optional`
	pathlossReferenceRSs_v1610   *PathlossReferenceRSs_v1610                      `optional,ext-1,setuprelease`
}

func (ie *PUCCH_PowerControl) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.pathlossReferenceRSs_v1610 != nil
	preambleBits := []bool{hasExtensions, ie.deltaF_PUCCH_f0 != nil, ie.deltaF_PUCCH_f1 != nil, ie.deltaF_PUCCH_f2 != nil, ie.deltaF_PUCCH_f3 != nil, ie.deltaF_PUCCH_f4 != nil, len(ie.p0_Set) > 0, len(ie.pathlossReferenceRSs) > 0, ie.twoPUCCH_PC_AdjustmentStates != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.deltaF_PUCCH_f0 != nil {
		if err = w.WriteInteger(*ie.deltaF_PUCCH_f0, &uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode deltaF_PUCCH_f0", err)
		}
	}
	if ie.deltaF_PUCCH_f1 != nil {
		if err = w.WriteInteger(*ie.deltaF_PUCCH_f1, &uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode deltaF_PUCCH_f1", err)
		}
	}
	if ie.deltaF_PUCCH_f2 != nil {
		if err = w.WriteInteger(*ie.deltaF_PUCCH_f2, &uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode deltaF_PUCCH_f2", err)
		}
	}
	if ie.deltaF_PUCCH_f3 != nil {
		if err = w.WriteInteger(*ie.deltaF_PUCCH_f3, &uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode deltaF_PUCCH_f3", err)
		}
	}
	if ie.deltaF_PUCCH_f4 != nil {
		if err = w.WriteInteger(*ie.deltaF_PUCCH_f4, &uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode deltaF_PUCCH_f4", err)
		}
	}
	if len(ie.p0_Set) > 0 {
		tmp_p0_Set := utils.NewSequence[*P0_PUCCH]([]*P0_PUCCH{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_P0_PerSet}, false)
		for _, i := range ie.p0_Set {
			tmp_p0_Set.Value = append(tmp_p0_Set.Value, &i)
		}
		if err = tmp_p0_Set.Encode(w); err != nil {
			return utils.WrapError("Encode p0_Set", err)
		}
	}
	if len(ie.pathlossReferenceRSs) > 0 {
		tmp_pathlossReferenceRSs := utils.NewSequence[*PUCCH_PathlossReferenceRS]([]*PUCCH_PathlossReferenceRS{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_PathlossReferenceRSs}, false)
		for _, i := range ie.pathlossReferenceRSs {
			tmp_pathlossReferenceRSs.Value = append(tmp_pathlossReferenceRSs.Value, &i)
		}
		if err = tmp_pathlossReferenceRSs.Encode(w); err != nil {
			return utils.WrapError("Encode pathlossReferenceRSs", err)
		}
	}
	if ie.twoPUCCH_PC_AdjustmentStates != nil {
		if err = ie.twoPUCCH_PC_AdjustmentStates.Encode(w); err != nil {
			return utils.WrapError("Encode twoPUCCH_PC_AdjustmentStates", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.pathlossReferenceRSs_v1610 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PUCCH_PowerControl", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.pathlossReferenceRSs_v1610 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode pathlossReferenceRSs_v1610 optional
			if ie.pathlossReferenceRSs_v1610 != nil {
				tmp_pathlossReferenceRSs_v1610 := utils.SetupRelease[*PathlossReferenceRSs_v1610]{
					Setup: ie.pathlossReferenceRSs_v1610,
				}
				if err = tmp_pathlossReferenceRSs_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pathlossReferenceRSs_v1610", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *PUCCH_PowerControl) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var deltaF_PUCCH_f0Present bool
	if deltaF_PUCCH_f0Present, err = r.ReadBool(); err != nil {
		return err
	}
	var deltaF_PUCCH_f1Present bool
	if deltaF_PUCCH_f1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var deltaF_PUCCH_f2Present bool
	if deltaF_PUCCH_f2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var deltaF_PUCCH_f3Present bool
	if deltaF_PUCCH_f3Present, err = r.ReadBool(); err != nil {
		return err
	}
	var deltaF_PUCCH_f4Present bool
	if deltaF_PUCCH_f4Present, err = r.ReadBool(); err != nil {
		return err
	}
	var p0_SetPresent bool
	if p0_SetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pathlossReferenceRSsPresent bool
	if pathlossReferenceRSsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var twoPUCCH_PC_AdjustmentStatesPresent bool
	if twoPUCCH_PC_AdjustmentStatesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if deltaF_PUCCH_f0Present {
		var tmp_int_deltaF_PUCCH_f0 int64
		if tmp_int_deltaF_PUCCH_f0, err = r.ReadInteger(&uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode deltaF_PUCCH_f0", err)
		}
		ie.deltaF_PUCCH_f0 = &tmp_int_deltaF_PUCCH_f0
	}
	if deltaF_PUCCH_f1Present {
		var tmp_int_deltaF_PUCCH_f1 int64
		if tmp_int_deltaF_PUCCH_f1, err = r.ReadInteger(&uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode deltaF_PUCCH_f1", err)
		}
		ie.deltaF_PUCCH_f1 = &tmp_int_deltaF_PUCCH_f1
	}
	if deltaF_PUCCH_f2Present {
		var tmp_int_deltaF_PUCCH_f2 int64
		if tmp_int_deltaF_PUCCH_f2, err = r.ReadInteger(&uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode deltaF_PUCCH_f2", err)
		}
		ie.deltaF_PUCCH_f2 = &tmp_int_deltaF_PUCCH_f2
	}
	if deltaF_PUCCH_f3Present {
		var tmp_int_deltaF_PUCCH_f3 int64
		if tmp_int_deltaF_PUCCH_f3, err = r.ReadInteger(&uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode deltaF_PUCCH_f3", err)
		}
		ie.deltaF_PUCCH_f3 = &tmp_int_deltaF_PUCCH_f3
	}
	if deltaF_PUCCH_f4Present {
		var tmp_int_deltaF_PUCCH_f4 int64
		if tmp_int_deltaF_PUCCH_f4, err = r.ReadInteger(&uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode deltaF_PUCCH_f4", err)
		}
		ie.deltaF_PUCCH_f4 = &tmp_int_deltaF_PUCCH_f4
	}
	if p0_SetPresent {
		tmp_p0_Set := utils.NewSequence[*P0_PUCCH]([]*P0_PUCCH{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_P0_PerSet}, false)
		fn_p0_Set := func() *P0_PUCCH {
			return new(P0_PUCCH)
		}
		if err = tmp_p0_Set.Decode(r, fn_p0_Set); err != nil {
			return utils.WrapError("Decode p0_Set", err)
		}
		ie.p0_Set = []P0_PUCCH{}
		for _, i := range tmp_p0_Set.Value {
			ie.p0_Set = append(ie.p0_Set, *i)
		}
	}
	if pathlossReferenceRSsPresent {
		tmp_pathlossReferenceRSs := utils.NewSequence[*PUCCH_PathlossReferenceRS]([]*PUCCH_PathlossReferenceRS{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_PathlossReferenceRSs}, false)
		fn_pathlossReferenceRSs := func() *PUCCH_PathlossReferenceRS {
			return new(PUCCH_PathlossReferenceRS)
		}
		if err = tmp_pathlossReferenceRSs.Decode(r, fn_pathlossReferenceRSs); err != nil {
			return utils.WrapError("Decode pathlossReferenceRSs", err)
		}
		ie.pathlossReferenceRSs = []PUCCH_PathlossReferenceRS{}
		for _, i := range tmp_pathlossReferenceRSs.Value {
			ie.pathlossReferenceRSs = append(ie.pathlossReferenceRSs, *i)
		}
	}
	if twoPUCCH_PC_AdjustmentStatesPresent {
		ie.twoPUCCH_PC_AdjustmentStates = new(PUCCH_PowerControl_twoPUCCH_PC_AdjustmentStates)
		if err = ie.twoPUCCH_PC_AdjustmentStates.Decode(r); err != nil {
			return utils.WrapError("Decode twoPUCCH_PC_AdjustmentStates", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			pathlossReferenceRSs_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode pathlossReferenceRSs_v1610 optional
			if pathlossReferenceRSs_v1610Present {
				tmp_pathlossReferenceRSs_v1610 := utils.SetupRelease[*PathlossReferenceRSs_v1610]{}
				if err = tmp_pathlossReferenceRSs_v1610.Decode(extReader); err != nil {
					return utils.WrapError("Decode pathlossReferenceRSs_v1610", err)
				}
				ie.pathlossReferenceRSs_v1610 = tmp_pathlossReferenceRSs_v1610.Setup
			}
		}
	}
	return nil
}
