package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PHR_Config struct {
	phr_PeriodicTimer        PHR_Config_phr_PeriodicTimer        `madatory`
	phr_ProhibitTimer        PHR_Config_phr_ProhibitTimer        `madatory`
	phr_Tx_PowerFactorChange PHR_Config_phr_Tx_PowerFactorChange `madatory`
	multiplePHR              bool                                `madatory`
	dummy                    bool                                `madatory`
	phr_Type2OtherCell       bool                                `madatory`
	phr_ModeOtherCG          PHR_Config_phr_ModeOtherCG          `madatory`
	mpe_Reporting_FR2_r16    *MPE_Config_FR2_r16                 `optional,ext-1,setuprelease`
	mpe_Reporting_FR2_r17    *MPE_Config_FR2_r17                 `optional,ext-2,setuprelease`
	twoPHRMode_r17           *PHR_Config_twoPHRMode_r17          `optional,ext-2`
}

func (ie *PHR_Config) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.mpe_Reporting_FR2_r16 != nil || ie.mpe_Reporting_FR2_r17 != nil || ie.twoPHRMode_r17 != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.phr_PeriodicTimer.Encode(w); err != nil {
		return utils.WrapError("Encode phr_PeriodicTimer", err)
	}
	if err = ie.phr_ProhibitTimer.Encode(w); err != nil {
		return utils.WrapError("Encode phr_ProhibitTimer", err)
	}
	if err = ie.phr_Tx_PowerFactorChange.Encode(w); err != nil {
		return utils.WrapError("Encode phr_Tx_PowerFactorChange", err)
	}
	if err = w.WriteBoolean(ie.multiplePHR); err != nil {
		return utils.WrapError("WriteBoolean multiplePHR", err)
	}
	if err = w.WriteBoolean(ie.dummy); err != nil {
		return utils.WrapError("WriteBoolean dummy", err)
	}
	if err = w.WriteBoolean(ie.phr_Type2OtherCell); err != nil {
		return utils.WrapError("WriteBoolean phr_Type2OtherCell", err)
	}
	if err = ie.phr_ModeOtherCG.Encode(w); err != nil {
		return utils.WrapError("Encode phr_ModeOtherCG", err)
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.mpe_Reporting_FR2_r16 != nil, ie.mpe_Reporting_FR2_r17 != nil || ie.twoPHRMode_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PHR_Config", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.mpe_Reporting_FR2_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode mpe_Reporting_FR2_r16 optional
			if ie.mpe_Reporting_FR2_r16 != nil {
				tmp_mpe_Reporting_FR2_r16 := utils.SetupRelease[*MPE_Config_FR2_r16]{
					Setup: ie.mpe_Reporting_FR2_r16,
				}
				if err = tmp_mpe_Reporting_FR2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mpe_Reporting_FR2_r16", err)
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
			optionals_ext_2 := []bool{ie.mpe_Reporting_FR2_r17 != nil, ie.twoPHRMode_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode mpe_Reporting_FR2_r17 optional
			if ie.mpe_Reporting_FR2_r17 != nil {
				tmp_mpe_Reporting_FR2_r17 := utils.SetupRelease[*MPE_Config_FR2_r17]{
					Setup: ie.mpe_Reporting_FR2_r17,
				}
				if err = tmp_mpe_Reporting_FR2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mpe_Reporting_FR2_r17", err)
				}
			}
			// encode twoPHRMode_r17 optional
			if ie.twoPHRMode_r17 != nil {
				if err = ie.twoPHRMode_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode twoPHRMode_r17", err)
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

func (ie *PHR_Config) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.phr_PeriodicTimer.Decode(r); err != nil {
		return utils.WrapError("Decode phr_PeriodicTimer", err)
	}
	if err = ie.phr_ProhibitTimer.Decode(r); err != nil {
		return utils.WrapError("Decode phr_ProhibitTimer", err)
	}
	if err = ie.phr_Tx_PowerFactorChange.Decode(r); err != nil {
		return utils.WrapError("Decode phr_Tx_PowerFactorChange", err)
	}
	var tmp_bool_multiplePHR bool
	if tmp_bool_multiplePHR, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean multiplePHR", err)
	}
	ie.multiplePHR = tmp_bool_multiplePHR
	var tmp_bool_dummy bool
	if tmp_bool_dummy, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean dummy", err)
	}
	ie.dummy = tmp_bool_dummy
	var tmp_bool_phr_Type2OtherCell bool
	if tmp_bool_phr_Type2OtherCell, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean phr_Type2OtherCell", err)
	}
	ie.phr_Type2OtherCell = tmp_bool_phr_Type2OtherCell
	if err = ie.phr_ModeOtherCG.Decode(r); err != nil {
		return utils.WrapError("Decode phr_ModeOtherCG", err)
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

			mpe_Reporting_FR2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode mpe_Reporting_FR2_r16 optional
			if mpe_Reporting_FR2_r16Present {
				tmp_mpe_Reporting_FR2_r16 := utils.SetupRelease[*MPE_Config_FR2_r16]{}
				if err = tmp_mpe_Reporting_FR2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode mpe_Reporting_FR2_r16", err)
				}
				ie.mpe_Reporting_FR2_r16 = tmp_mpe_Reporting_FR2_r16.Setup
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			mpe_Reporting_FR2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			twoPHRMode_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode mpe_Reporting_FR2_r17 optional
			if mpe_Reporting_FR2_r17Present {
				tmp_mpe_Reporting_FR2_r17 := utils.SetupRelease[*MPE_Config_FR2_r17]{}
				if err = tmp_mpe_Reporting_FR2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mpe_Reporting_FR2_r17", err)
				}
				ie.mpe_Reporting_FR2_r17 = tmp_mpe_Reporting_FR2_r17.Setup
			}
			// decode twoPHRMode_r17 optional
			if twoPHRMode_r17Present {
				ie.twoPHRMode_r17 = new(PHR_Config_twoPHRMode_r17)
				if err = ie.twoPHRMode_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode twoPHRMode_r17", err)
				}
			}
		}
	}
	return nil
}
