package helper_test

import (
	"github.com/pivotal-cf/deploy-errand/helper"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"code.cloudfoundry.org/commandrunner/fake_command_runner"

	"code.cloudfoundry.org/commandrunner/fake_command_runner/matchers"
)

var _ = Describe("Echo", func() {

	var fakeRunner *fake_command_runner.FakeCommandRunner

	BeforeEach(func(){
		fakeRunner = fake_command_runner.New()
	})

	It("invoke echo", func(){
		helper.Echo(fakeRunner)
		Expect(fakeRunner).To(fake_command_runner_matchers.HaveExecutedSerially(fake_command_runner.CommandSpec{Path: "echo", Args: []string{"shelled out"}}))
	})

})
