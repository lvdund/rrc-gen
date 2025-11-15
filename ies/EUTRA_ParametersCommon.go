package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_ParametersCommon struct {
	mfbi_EUTRA                *EUTRA_ParametersCommon_mfbi_EUTRA         `optional`
	modifiedMPR_BehaviorEUTRA *uper.BitString                            `lb:32,ub:32,optional`
	multiNS_Pmax_EUTRA        *EUTRA_ParametersCommon_multiNS_Pmax_EUTRA `optional`
	rs_SINR_MeasEUTRA         *EUTRA_ParametersCommon_rs_SINR_MeasEUTRA  `optional`
	ne_DC                     *EUTRA_ParametersCommon_ne_DC              `optional,ext-1`
	nr_HO_ToEN_DC_r16         *EUTRA_ParametersCommon_nr_HO_ToEN_DC_r16  `optional,ext-2`
}

func (ie *EUTRA_ParametersCommon) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.ne_DC != nil || ie.nr_HO_ToEN_DC_r16 != nil
	preambleBits := []bool{hasExtensions, ie.mfbi_EUTRA != nil, ie.modifiedMPR_BehaviorEUTRA != nil, ie.multiNS_Pmax_EUTRA != nil, ie.rs_SINR_MeasEUTRA != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.mfbi_EUTRA != nil {
		if err = ie.mfbi_EUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode mfbi_EUTRA", err)
		}
	}
	if ie.modifiedMPR_BehaviorEUTRA != nil {
		if err = w.WriteBitString(ie.modifiedMPR_BehaviorEUTRA.Bytes, uint(ie.modifiedMPR_BehaviorEUTRA.NumBits), &uper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode modifiedMPR_BehaviorEUTRA", err)
		}
	}
	if ie.multiNS_Pmax_EUTRA != nil {
		if err = ie.multiNS_Pmax_EUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode multiNS_Pmax_EUTRA", err)
		}
	}
	if ie.rs_SINR_MeasEUTRA != nil {
		if err = ie.rs_SINR_MeasEUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode rs_SINR_MeasEUTRA", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.ne_DC != nil, ie.nr_HO_ToEN_DC_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap EUTRA_ParametersCommon", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.ne_DC != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ne_DC optional
			if ie.ne_DC != nil {
				if err = ie.ne_DC.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ne_DC", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.nr_HO_ToEN_DC_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode nr_HO_ToEN_DC_r16 optional
			if ie.nr_HO_ToEN_DC_r16 != nil {
				if err = ie.nr_HO_ToEN_DC_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode nr_HO_ToEN_DC_r16", err)
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

func (ie *EUTRA_ParametersCommon) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var mfbi_EUTRAPresent bool
	if mfbi_EUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var modifiedMPR_BehaviorEUTRAPresent bool
	if modifiedMPR_BehaviorEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var multiNS_Pmax_EUTRAPresent bool
	if multiNS_Pmax_EUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var rs_SINR_MeasEUTRAPresent bool
	if rs_SINR_MeasEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if mfbi_EUTRAPresent {
		ie.mfbi_EUTRA = new(EUTRA_ParametersCommon_mfbi_EUTRA)
		if err = ie.mfbi_EUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode mfbi_EUTRA", err)
		}
	}
	if modifiedMPR_BehaviorEUTRAPresent {
		var tmp_bs_modifiedMPR_BehaviorEUTRA []byte
		var n_modifiedMPR_BehaviorEUTRA uint
		if tmp_bs_modifiedMPR_BehaviorEUTRA, n_modifiedMPR_BehaviorEUTRA, err = r.ReadBitString(&uper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode modifiedMPR_BehaviorEUTRA", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_modifiedMPR_BehaviorEUTRA,
			NumBits: uint64(n_modifiedMPR_BehaviorEUTRA),
		}
		ie.modifiedMPR_BehaviorEUTRA = &tmp_bitstring
	}
	if multiNS_Pmax_EUTRAPresent {
		ie.multiNS_Pmax_EUTRA = new(EUTRA_ParametersCommon_multiNS_Pmax_EUTRA)
		if err = ie.multiNS_Pmax_EUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode multiNS_Pmax_EUTRA", err)
		}
	}
	if rs_SINR_MeasEUTRAPresent {
		ie.rs_SINR_MeasEUTRA = new(EUTRA_ParametersCommon_rs_SINR_MeasEUTRA)
		if err = ie.rs_SINR_MeasEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode rs_SINR_MeasEUTRA", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
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

			ne_DCPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ne_DC optional
			if ne_DCPresent {
				ie.ne_DC = new(EUTRA_ParametersCommon_ne_DC)
				if err = ie.ne_DC.Decode(extReader); err != nil {
					return utils.WrapError("Decode ne_DC", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			nr_HO_ToEN_DC_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode nr_HO_ToEN_DC_r16 optional
			if nr_HO_ToEN_DC_r16Present {
				ie.nr_HO_ToEN_DC_r16 = new(EUTRA_ParametersCommon_nr_HO_ToEN_DC_r16)
				if err = ie.nr_HO_ToEN_DC_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode nr_HO_ToEN_DC_r16", err)
				}
			}
		}
	}
	return nil
}
