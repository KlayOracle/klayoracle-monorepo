package config

import (
	"github.com/stretchr/testify/suite"
	"os"
	"path"
	"testing"
)

type EnvironmentTestSuite struct {
	suite.Suite
	wd string
}

func (s *EnvironmentTestSuite) BeforeTest(suiteName, testName string) {
	wd, err := os.Getwd()
	if err != nil {
		s.Fail("Cannot determine working directory: ", err)
	}

	s.wd = wd
}

func (s *EnvironmentTestSuite) TestConfig() {
	testFile := path.Join(s.wd, "..", "config_test.yaml")
	configFile, err := os.Create(testFile)
	if err != nil {
		s.Fail("Cannot create file: config_test.yaml")
	}

	configFile.Write([]byte(`env: dev
service_node: 127.0.0.1
feed:
  path: feeds
  adapters:
    - KLAY_USD.json
    - RANDOM.json
organization:
  k_org_id: "PO-DMX-804-23"
  name: "KlayOracle Inc."
  website: "https://klayoracle.com"
ssl:
  key: "mono-provider.key"
  certificate: "mono-provider.crt"`))

	Load(testFile)

	s.Equal(Loaded.ServiceNode, "127.0.0.1")
	s.Equal(Loaded.Feed.Path, "feeds")
	s.EqualValues(Loaded.Feed.Adapters, []string{"KLAY_USD.json", "RANDOM.json"})
	s.Equal(Loaded.Env, "dev")
	s.EqualValues(Loaded.Organization, struct {
		ID      string
		Name    string
		Website string
	}{
		ID:      "PO-DMX-804-23",
		Name:    "KlayOracle Inc.",
		Website: "https://klayoracle.com",
	})
	s.Equal(Loaded.SSL.Key, "mono-provider.key")
	s.Equal(Loaded.SSL.Certificate, "mono-provider.crt")
}

func (s *EnvironmentTestSuite) TestLogger() {
	Loaded.Logger = NewLogger(path.Join(s.wd, "..", "test.log"))
	Loaded.Logger.Infow("Config", "running", "provider")

	content, err := os.ReadFile(path.Join(s.wd, "..", "test.log"))
	if err != nil {
		s.Fail("Cannot read file: test.log")
	}

	s.NotNil(string(content))
	s.Equal(string(content), "info	Config	{\"running\": \"provider\"}\n")
}

func TestEnvironmentSetup(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fail()
	}

	suite.Run(t, &EnvironmentTestSuite{
		wd: wd,
	})
}

func (s *EnvironmentTestSuite) AfterTest(suiteName, testName string) {
	//fmt.Println(path.Join(s.wd, "..", "config_test.yaml"))
	_ = os.Remove(path.Join(s.wd, "..", "config_test.yaml"))
	//if err != nil {
	//	fmt.Println(err)
	//	s.Fail("Could not remove config_test.yaml")
	//}
	_ = os.Remove(path.Join(s.wd, "..", "test.log"))
	//if err != nil {
	//	fmt.Println(err)
	//	s.Fail("Could not remove test.log")
	//}
}
