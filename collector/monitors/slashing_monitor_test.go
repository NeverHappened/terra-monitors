package monitors

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/lidofinance/terra-monitors/collector/config"
	"github.com/stretchr/testify/suite"
)

const (
	testValAddress   = "terravalcons1ezj3lps8nqwytt42at2sgt7seq9hk708g0spyk"
	testValPublicKey = "terravalconspub1zcjduepqw2hyr7u7y70z5kdewn00xuq0wwcvnn0s7x5pjqcdpn80qsyctcpqcjhz4c"
	testConsAddress  = "terravalcons1rfaxjug6md5jrz3c0uctyt6pzd50xyxlc2tf5m"
)

type SlashingMonitorTestSuite struct {
	suite.Suite
}

func (suite *SlashingMonitorTestSuite) SetupTest() {

}

func (suite *SlashingMonitorTestSuite) TestSuccessfulRequestWithSlashing() {
	validatorInfoData, err := ioutil.ReadFile("./test_data/slashing_validator_info_jailed.json")
	suite.NoError(err)

	validatorSigningInfoData, err := ioutil.ReadFile(
		"./test_data/slashing_success_response_blocks_jailed_tombstoned.json")
	suite.NoError(err)

	whitelistedValidators, err := ioutil.ReadFile("./test_data/whitelisted_validators_response.json")
	suite.NoError(err)

	testServer := NewServerWithRoutedResponse(map[string]string{
		fmt.Sprintf("/staking/validators/%s", testValAddress):                     string(validatorInfoData),
		fmt.Sprintf("/cosmos/slashing/v1beta1/signing_infos/%s", testConsAddress): string(validatorSigningInfoData),
		fmt.Sprintf("/wasm/contracts/%s/store", HubContract):                      string(whitelistedValidators),
	})
	cfg := NewTestCollectorConfig(testServer.URL)
	cfg.BassetContractsVersion = config.V1Contracts

	logger := NewTestLogger()
	valRepository := NewValidatorsRepository(cfg, logger)
	m := NewSlashingMonitor(cfg, logger, valRepository)
	err = m.Handler(context.Background())
	suite.NoError(err)

	metrics := m.GetMetrics()
	metricVectors := m.GetMetricVectors()
	var (
		expectedNumTombstonedValidators MetricValue = &SimpleMetricValue{value: 1}
		expectedNumJailedValidators     MetricValue = &SimpleMetricValue{value: 1}
		expectedNumMissedBlocks         float64     = 5
	)
	var actualMissedBlocks float64
	for _, missedBlocks := range metricVectors[SlashingNumMissedBlocks].Labels() {
		actualMissedBlocks += metricVectors[SlashingNumMissedBlocks].Get(missedBlocks)
	}
	suite.Equal(expectedNumTombstonedValidators, metrics[SlashingNumTombstonedValidators])
	suite.Equal(expectedNumJailedValidators, metrics[SlashingNumJailedValidators])
	suite.Equal(expectedNumMissedBlocks, actualMissedBlocks)
}

func (suite *SlashingMonitorTestSuite) TestSuccessfulRequestNoSlashing() {
	validatorInfoData, err := ioutil.ReadFile("./test_data/slashing_validator_info_not_jailed.json")
	suite.NoError(err)

	validatorSigningInfoData, err := ioutil.ReadFile(
		"./test_data/slashing_success_response_no_slashing.json")
	suite.NoError(err)

	whitelistedValidators, err := ioutil.ReadFile("./test_data/whitelisted_validators_response.json")
	suite.NoError(err)

	testServer := NewServerWithRoutedResponse(map[string]string{
		fmt.Sprintf("/staking/validators/%s", testValAddress):                     string(validatorInfoData),
		fmt.Sprintf("/cosmos/slashing/v1beta1/signing_infos/%s", testConsAddress): string(validatorSigningInfoData),
		fmt.Sprintf("/wasm/contracts/%s/store", HubContract):                      string(whitelistedValidators),
	})
	cfg := NewTestCollectorConfig(testServer.URL)
	cfg.BassetContractsVersion = config.V1Contracts

	logger := NewTestLogger()
	valRepository := NewValidatorsRepository(cfg, logger)
	m := NewSlashingMonitor(cfg, logger, valRepository)
	err = m.Handler(context.Background())
	suite.NoError(err)

	metrics := m.GetMetrics()
	metricVectors := m.GetMetricVectors()

	var (
		expectedNumTombstonedValidators MetricValue = &SimpleMetricValue{value: 0}
		expectedNumJailedValidators     MetricValue = &SimpleMetricValue{value: 0}
		expectedNumMissedBlocks         float64     = 0
	)
	var actualMissedBlocks float64
	for _, missedBlocks := range metricVectors[SlashingNumMissedBlocks].Labels() {
		actualMissedBlocks += metricVectors[SlashingNumMissedBlocks].Get(missedBlocks)
	}
	suite.Equal(expectedNumTombstonedValidators, metrics[SlashingNumTombstonedValidators])
	suite.Equal(expectedNumJailedValidators, metrics[SlashingNumJailedValidators])
	suite.Equal(expectedNumMissedBlocks, actualMissedBlocks)
}

func (suite *UpdateGlobalIndexMonitorTestSuite) TestFailedSlashingRequest() {
	validatorInfoData, err := ioutil.ReadFile("./test_data/slashing_error.json")
	suite.NoError(err)

	validatorSigningInfoData, err := ioutil.ReadFile(
		"./test_data/slashing_success_response_blocks_jailed_tombstoned.json")
	suite.NoError(err)

	whitelistedValidators, err := ioutil.ReadFile("./test_data/whitelisted_validators_response.json")
	suite.NoError(err)

	testServer := NewServerWithRoutedResponse(map[string]string{
		fmt.Sprintf("/staking/validators/%s", testValAddress):                     string(validatorInfoData),
		fmt.Sprintf("/cosmos/slashing/v1beta1/signing_infos/%s", testConsAddress): string(validatorSigningInfoData),
		fmt.Sprintf("/wasm/contracts/%s/store", HubContract):                      string(whitelistedValidators),
	})
	cfg := NewTestCollectorConfig(testServer.URL)
	cfg.BassetContractsVersion = config.V1Contracts

	logger := NewTestLogger()
	valRepository := NewValidatorsRepository(cfg, logger)
	m := NewSlashingMonitor(cfg, logger, valRepository)
	err = m.Handler(context.Background())
	suite.Error(err)

	expectedErrorMessage := "failed to getValidatorsInfo"
	suite.Contains(err.Error(), expectedErrorMessage)

}
