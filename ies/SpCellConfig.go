package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SpCellConfig struct {
	servCellIndex                      *ServCellIndex                                   `optional`
	reconfigurationWithSync            *ReconfigurationWithSync                         `optional`
	rlf_TimersAndConstants             *RLF_TimersAndConstants                          `optional,setuprelease`
	rlmInSyncOutOfSyncThreshold        *SpCellConfig_rlmInSyncOutOfSyncThreshold        `optional`
	spCellConfigDedicated              *ServingCellConfig                               `optional`
	lowMobilityEvaluationConnected_r17 *SpCellConfig_lowMobilityEvaluationConnected_r17 `optional,ext-1`
	goodServingCellEvaluationRLM_r17   *GoodServingCellEvaluation_r17                   `optional,ext-1`
	goodServingCellEvaluationBFD_r17   *GoodServingCellEvaluation_r17                   `optional,ext-1`
	deactivatedSCG_Config_r17          *DeactivatedSCG_Config_r17                       `optional,ext-1,setuprelease`
}

func (ie *SpCellConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.lowMobilityEvaluationConnected_r17 != nil || ie.goodServingCellEvaluationRLM_r17 != nil || ie.goodServingCellEvaluationBFD_r17 != nil || ie.deactivatedSCG_Config_r17 != nil
	preambleBits := []bool{hasExtensions, ie.servCellIndex != nil, ie.reconfigurationWithSync != nil, ie.rlf_TimersAndConstants != nil, ie.rlmInSyncOutOfSyncThreshold != nil, ie.spCellConfigDedicated != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.servCellIndex != nil {
		if err = ie.servCellIndex.Encode(w); err != nil {
			return utils.WrapError("Encode servCellIndex", err)
		}
	}
	if ie.reconfigurationWithSync != nil {
		if err = ie.reconfigurationWithSync.Encode(w); err != nil {
			return utils.WrapError("Encode reconfigurationWithSync", err)
		}
	}
	if ie.rlf_TimersAndConstants != nil {
		tmp_rlf_TimersAndConstants := utils.SetupRelease[*RLF_TimersAndConstants]{
			Setup: ie.rlf_TimersAndConstants,
		}
		if err = tmp_rlf_TimersAndConstants.Encode(w); err != nil {
			return utils.WrapError("Encode rlf_TimersAndConstants", err)
		}
	}
	if ie.rlmInSyncOutOfSyncThreshold != nil {
		if err = ie.rlmInSyncOutOfSyncThreshold.Encode(w); err != nil {
			return utils.WrapError("Encode rlmInSyncOutOfSyncThreshold", err)
		}
	}
	if ie.spCellConfigDedicated != nil {
		if err = ie.spCellConfigDedicated.Encode(w); err != nil {
			return utils.WrapError("Encode spCellConfigDedicated", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.lowMobilityEvaluationConnected_r17 != nil || ie.goodServingCellEvaluationRLM_r17 != nil || ie.goodServingCellEvaluationBFD_r17 != nil || ie.deactivatedSCG_Config_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SpCellConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.lowMobilityEvaluationConnected_r17 != nil, ie.goodServingCellEvaluationRLM_r17 != nil, ie.goodServingCellEvaluationBFD_r17 != nil, ie.deactivatedSCG_Config_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode lowMobilityEvaluationConnected_r17 optional
			if ie.lowMobilityEvaluationConnected_r17 != nil {
				if err = ie.lowMobilityEvaluationConnected_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode lowMobilityEvaluationConnected_r17", err)
				}
			}
			// encode goodServingCellEvaluationRLM_r17 optional
			if ie.goodServingCellEvaluationRLM_r17 != nil {
				if err = ie.goodServingCellEvaluationRLM_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode goodServingCellEvaluationRLM_r17", err)
				}
			}
			// encode goodServingCellEvaluationBFD_r17 optional
			if ie.goodServingCellEvaluationBFD_r17 != nil {
				if err = ie.goodServingCellEvaluationBFD_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode goodServingCellEvaluationBFD_r17", err)
				}
			}
			// encode deactivatedSCG_Config_r17 optional
			if ie.deactivatedSCG_Config_r17 != nil {
				tmp_deactivatedSCG_Config_r17 := utils.SetupRelease[*DeactivatedSCG_Config_r17]{
					Setup: ie.deactivatedSCG_Config_r17,
				}
				if err = tmp_deactivatedSCG_Config_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode deactivatedSCG_Config_r17", err)
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

func (ie *SpCellConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var servCellIndexPresent bool
	if servCellIndexPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var reconfigurationWithSyncPresent bool
	if reconfigurationWithSyncPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var rlf_TimersAndConstantsPresent bool
	if rlf_TimersAndConstantsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var rlmInSyncOutOfSyncThresholdPresent bool
	if rlmInSyncOutOfSyncThresholdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var spCellConfigDedicatedPresent bool
	if spCellConfigDedicatedPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if servCellIndexPresent {
		ie.servCellIndex = new(ServCellIndex)
		if err = ie.servCellIndex.Decode(r); err != nil {
			return utils.WrapError("Decode servCellIndex", err)
		}
	}
	if reconfigurationWithSyncPresent {
		ie.reconfigurationWithSync = new(ReconfigurationWithSync)
		if err = ie.reconfigurationWithSync.Decode(r); err != nil {
			return utils.WrapError("Decode reconfigurationWithSync", err)
		}
	}
	if rlf_TimersAndConstantsPresent {
		tmp_rlf_TimersAndConstants := utils.SetupRelease[*RLF_TimersAndConstants]{}
		if err = tmp_rlf_TimersAndConstants.Decode(r); err != nil {
			return utils.WrapError("Decode rlf_TimersAndConstants", err)
		}
		ie.rlf_TimersAndConstants = tmp_rlf_TimersAndConstants.Setup
	}
	if rlmInSyncOutOfSyncThresholdPresent {
		ie.rlmInSyncOutOfSyncThreshold = new(SpCellConfig_rlmInSyncOutOfSyncThreshold)
		if err = ie.rlmInSyncOutOfSyncThreshold.Decode(r); err != nil {
			return utils.WrapError("Decode rlmInSyncOutOfSyncThreshold", err)
		}
	}
	if spCellConfigDedicatedPresent {
		ie.spCellConfigDedicated = new(ServingCellConfig)
		if err = ie.spCellConfigDedicated.Decode(r); err != nil {
			return utils.WrapError("Decode spCellConfigDedicated", err)
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

			lowMobilityEvaluationConnected_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			goodServingCellEvaluationRLM_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			goodServingCellEvaluationBFD_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			deactivatedSCG_Config_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode lowMobilityEvaluationConnected_r17 optional
			if lowMobilityEvaluationConnected_r17Present {
				ie.lowMobilityEvaluationConnected_r17 = new(SpCellConfig_lowMobilityEvaluationConnected_r17)
				if err = ie.lowMobilityEvaluationConnected_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode lowMobilityEvaluationConnected_r17", err)
				}
			}
			// decode goodServingCellEvaluationRLM_r17 optional
			if goodServingCellEvaluationRLM_r17Present {
				ie.goodServingCellEvaluationRLM_r17 = new(GoodServingCellEvaluation_r17)
				if err = ie.goodServingCellEvaluationRLM_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode goodServingCellEvaluationRLM_r17", err)
				}
			}
			// decode goodServingCellEvaluationBFD_r17 optional
			if goodServingCellEvaluationBFD_r17Present {
				ie.goodServingCellEvaluationBFD_r17 = new(GoodServingCellEvaluation_r17)
				if err = ie.goodServingCellEvaluationBFD_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode goodServingCellEvaluationBFD_r17", err)
				}
			}
			// decode deactivatedSCG_Config_r17 optional
			if deactivatedSCG_Config_r17Present {
				tmp_deactivatedSCG_Config_r17 := utils.SetupRelease[*DeactivatedSCG_Config_r17]{}
				if err = tmp_deactivatedSCG_Config_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode deactivatedSCG_Config_r17", err)
				}
				ie.deactivatedSCG_Config_r17 = tmp_deactivatedSCG_Config_r17.Setup
			}
		}
	}
	return nil
}
