package monitors

import (
	"context"
	"fmt"
	"github.com/lidofinance/terra-monitors/collector/types"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
)

type V2ValidatorsRepositoryTestSuite struct {
	suite.Suite
}

func (suite *V2ValidatorsRepositoryTestSuite) SetupTest() {

}

func (suite *V2ValidatorsRepositoryTestSuite) TestSuccessfulRequest() {
	validatorInfoData, err := ioutil.ReadFile("./test_data/slashing_validator_info.json")
	suite.NoError(err)

	whitelistedValidators, err := ioutil.ReadFile("./test_data/validators_registry_validators_response.json")
	suite.NoError(err)

	testServer := NewServerWithRoutedResponse(map[string]string{
		fmt.Sprintf("/staking/validators/%s", testValAddress):               string(validatorInfoData),
		fmt.Sprintf("/wasm/contracts/%s/store", ValidatorsRegistryContract): string(whitelistedValidators),
	})
	cfg := NewTestCollectorConfig(testServer.URL)

	valRepository := NewV2ValidatorsRepository(cfg)

	expectedValidatorsAddresses := []string{testValAddress}
	validators, err := valRepository.GetValidatorsAddresses(context.Background())
	suite.NoError(err)
	suite.Equal(expectedValidatorsAddresses, validators)

	expectedValidatorInfo := types.ValidatorInfo{
		Address:        testValAddress,
		PubKey:         testValPublicKey,
		Moniker:        TestMoniker,
		CommissionRate: TestCommissionRate,
	}
	validatorInfo, err := valRepository.GetValidatorInfo(context.Background(), testValAddress)
	suite.NoError(err)
	suite.Equal(validatorInfo, expectedValidatorInfo)
}

func (suite *ValidatorsCommissionTestSuite) TestFailedV2ValidatorsRepository() {
	validatorInfoData, err := ioutil.ReadFile("./test_data/slashing_error.json")
	suite.NoError(err)

	testServer := NewServerWithRoutedResponse(map[string]string{
		fmt.Sprintf("/staking/validators/%s", testValAddress): string(validatorInfoData),

		// error format is the same
		fmt.Sprintf("/wasm/contracts/%s/store", ValidatorsRegistryContract): string(validatorInfoData),
	})
	cfg := NewTestCollectorConfig(testServer.URL)

	valRepository := NewV2ValidatorsRepository(cfg)

	validators, err := valRepository.GetValidatorsAddresses(context.Background())
	suite.Nil(validators)
	suite.Error(err)

	_, err = valRepository.GetValidatorInfo(context.Background(), testValAddress)
	suite.Error(err)
}