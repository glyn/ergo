package cf_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"code.cloudfoundry.org/commandrunner/fake_command_runner"
	"code.cloudfoundry.org/commandrunner/fake_command_runner/matchers"
	"errors"
	"github.com/glyn/ergo/cf"
	"os/exec"
)

var _ = Describe("DisplayCfVersion", func() {

	var fakeRunner *fake_command_runner.FakeCommandRunner
	var cfWrapper cf.CF

	BeforeEach(func() {
		fakeRunner = fake_command_runner.New()
		cfWrapper = cf.New(fakeRunner)
	})

	It("invokes cf", func() {
		const testVersion = "cf version 99"
		expectedCmd := fake_command_runner.CommandSpec{Path: "cf", Args: []string{"-v"}}
		fakeRunner.WhenRunning(expectedCmd, func(cmd *exec.Cmd) error {
			cmd.Stdout.Write([]byte(testVersion))
			return nil
		})

		version, err := cfWrapper.DisplayCfVersion()

		Expect(fakeRunner).To(fake_command_runner_matchers.HaveExecutedSerially(expectedCmd))

		Expect(err).NotTo(HaveOccurred())
		Expect(version).To(Equal(testVersion))
	})

	It("propagates errors", func() {
		testError := errors.New("some error")
		expectedCmd := fake_command_runner.CommandSpec{Path: "cf", Args: []string{"-v"}}
		fakeRunner.WhenRunning(expectedCmd, func(cmd *exec.Cmd) error {
			return testError
		})

		_, err := cfWrapper.DisplayCfVersion()

		Expect(fakeRunner).To(fake_command_runner_matchers.HaveExecutedSerially(expectedCmd))

		Expect(err).To(MatchError(testError))
	})

})
